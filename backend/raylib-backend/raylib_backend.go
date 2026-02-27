package raylibbackend

import (
	"image"
	"unsafe"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	_ backend.Backend[RaylibBackendFlags] = &RaylibBackend{}
)

// raylibClipboardHandler implements imgui.ClipboardHandler for raylib
type raylibClipboardHandler struct{}

// GetClipboard returns the current clipboard text from raylib
func (h *raylibClipboardHandler) GetClipboard() string {
	return rl.GetClipboardText()
}

// SetClipboard sets the clipboard text using raylib
func (h *raylibClipboardHandler) SetClipboard(text string) { rl.SetClipboardText(text) }

// RaylibBackend implements the imgui.Backend interface for raylib
type RaylibBackend struct {
	customFontAtlas *imgui.FontAtlas
	customCtx       *imgui.Context

	afterCreateContext func()
	beforeImGuiRender  func()
	beforeRender       func()
	afterRender        func()
	beforeDestroy      func()
	loop               func()

	closeCb  backend.WindowCloseCallback
	resizeCb backend.SizeChangeCallback
	dropCb   backend.DropCallback
	keyCb    backend.KeyCallback

	shouldClose bool

	cache TextureCache
	ctx   *imgui.Context

	bgColor imgui.Vec4

	lastTime    float64
	lastW       int
	lastH       int
	lastFocused bool

	texRefCache map[uint32]imgui.TextureRef

	disabledScissorFrames int

	matCache map[uint32]rl.Material

	// cache vertex layout
	cachedLayout      *VertexLayout
	cachedLayoutValid bool

	// reusable mesh buffers
	meshVerts []float32
	meshUVs   []float32
	meshCols  []uint8
	meshIdxs  []uint16
}

// NewRaylibBackend creates a new raylib backend implementation
func NewRaylibBackend() *RaylibBackend {
	return &RaylibBackend{
		cache:       NewTextureCache(),
		lastTime:    0.0,
		bgColor:     imgui.NewVec4(0.1, 0.1, 0.1, 1.0),
		texRefCache: make(map[uint32]imgui.TextureRef),
		matCache:    make(map[uint32]rl.Material),
	}
}

// SetCustomFontAtlas allows setting a custom font atlas before CreateWindow
func (r *RaylibBackend) SetCustomFontAtlas(atlas *imgui.FontAtlas) {
	r.customFontAtlas = atlas
}

// SetCustomContext allows setting a custom ImGui context before CreateWindow
func (r *RaylibBackend) SetCustomContext(ctx *imgui.Context) {
	r.customCtx = ctx
}

// initImGuiContext initializes ImGui context and flags
func (r *RaylibBackend) initImGuiContext() {
	var imctx *imgui.Context
	switch {
	case r.customCtx != nil:
		imctx = r.customCtx
	case r.customFontAtlas != nil:
		imctx = imgui.CreateContextV(r.customFontAtlas)
	default:
		imctx = imgui.CreateContext()
	}

	r.ctx = imctx

	io := imgui.CurrentIO()
	io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsRendererHasTextures)
	io.SetBackendRendererName("cimgui-go-raylib")

	// Set up clipboard
	platformIO := imgui.CurrentPlatformIO()
	platformIO.SetClipboardHandler(&raylibClipboardHandler{})

	if fa := io.Fonts(); fa != nil {
		fa.SetTexDesiredFormat(imgui.TextureFormatRGBA32)
	}
}

// Dispose cleans up resources
func (r *RaylibBackend) Dispose() {
	// Nothing yet
}

// NewFrame is called by at the start of each frame
func (r *RaylibBackend) NewFrame() {
	io := imgui.CurrentIO()

	// Ensure font atlas is updated and has a valid TexID before NewFrame
	if fa := io.Fonts(); fa != nil && fa.TexIsBuilt() {
		if texData := fa.TexData(); texData != nil {
			r.cache.UpdateTexture(*texData)
		}
	}

	// Set display size
	w, h := int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight())
	io.SetDisplaySize(imgui.Vec2{X: float32(w), Y: float32(h)})

	// Sync imgui and raylib dpi. Default to 1 if unavailable.
	dpi := rl.GetWindowScaleDPI()
	if dpi.X <= 0 {
		dpi.X = 1
	}

	if dpi.Y <= 0 {
		dpi.Y = 1
	}

	io.SetDisplayFramebufferScale(imgui.Vec2{X: dpi.X, Y: dpi.Y})

	// Set delta time
	currentTime := rl.GetTime()
	if r.lastTime > 0 {
		io.SetDeltaTime(float32(currentTime - r.lastTime))
	}

	r.lastTime = currentTime

	// Update mouse position
	io.SetMousePos(imgui.Vec2{X: rl.GetMousePosition().X, Y: rl.GetMousePosition().Y})

	// Update mouse buttons
	io.SetMouseButtonDown(0, rl.IsMouseButtonDown(rl.MouseButtonLeft))
	io.SetMouseButtonDown(1, rl.IsMouseButtonDown(rl.MouseButtonRight))
	io.SetMouseButtonDown(2, rl.IsMouseButtonDown(rl.MouseButtonMiddle))
	io.SetMouseButtonDown(3, rl.IsMouseButtonDown(rl.MouseButtonSide))
	io.SetMouseButtonDown(4, rl.IsMouseButtonDown(rl.MouseButtonExtra))

	// Update mouse wheel
	io.AddMouseWheelDelta(0, rl.GetMouseWheelMove())

	// Feed keyboard and text input
	sendRaylibInput(io)

	// Only poll touch/gamepad if they are active
	if rl.GetTouchPointCount() > 0 || GetPrevTouchDown() {
		sendRaylibTouch(io)
	}

	// Only poll gamepad if one is connected
	if rl.IsGamepadAvailable(0) {
		sendRaylibGamepad(io)
	} else {
		// Clear gamepad flag when disconnected
		io.SetBackendFlags(io.BackendFlags() &^ imgui.BackendFlagsHasGamepad)
	}

	// Update focus state
	focused := rl.IsWindowFocused()
	if focused != r.lastFocused {
		io.AddFocusEvent(focused)
		r.lastFocused = focused
	}

	// Update cursor
	r.updateMouseCursor(io)
}

// updateMouseCursor changes the OS cursor
func (r *RaylibBackend) updateMouseCursor(io *imgui.IO) {
	// If ImGui wants to draw the cursor itself, hide the OS cursor
	if io.MouseDrawCursor() {
		rl.HideCursor()
		return
	}

	// Otherwise show OS cursor and set shape
	rl.ShowCursor()

	cursor := imgui.CurrentMouseCursor()
	switch cursor {
	case imgui.MouseCursorNone:
		rl.HideCursor()
	case imgui.MouseCursorArrow:
		rl.SetMouseCursor(rl.MouseCursorArrow)
	case imgui.MouseCursorTextInput:
		rl.SetMouseCursor(rl.MouseCursorIBeam)
	case imgui.MouseCursorResizeAll:
		rl.SetMouseCursor(rl.MouseCursorResizeAll)
	case imgui.MouseCursorResizeNS:
		rl.SetMouseCursor(rl.MouseCursorResizeNWSE)
	case imgui.MouseCursorResizeEW:
		rl.SetMouseCursor(rl.MouseCursorResizeNESW)
	case imgui.MouseCursorResizeNESW:
		rl.SetMouseCursor(rl.MouseCursorResizeNESW)
	case imgui.MouseCursorResizeNWSE:
		rl.SetMouseCursor(rl.MouseCursorResizeNWSE)
	case imgui.MouseCursorHand:
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	case imgui.MouseCursorNotAllowed:
		rl.SetMouseCursor(rl.MouseCursorNotAllowed)
	default:
		rl.SetMouseCursor(rl.MouseCursorArrow)
	}
}

// Render is the core drawing logic
func (r *RaylibBackend) Render(drawData imgui.DrawData) {
	// Flush any pending raylib draw calls
	rl.DrawRenderBatchActive()

	// Set up orthographic projection for screen space rendering
	displaySize := drawData.DisplaySize()
	fbScale := drawData.FramebufferScale()

	if fbScale.X == 0 {
		fbScale.X = 1
	}

	if fbScale.Y == 0 {
		fbScale.Y = 1
	}

	rl.MatrixMode(rl.Projection)
	rl.PushMatrix()
	rl.LoadIdentity()
	rl.Ortho(0, float64(displaySize.X*fbScale.X), float64(displaySize.Y*fbScale.Y), 0, -1, 1)

	rl.MatrixMode(rl.Modelview)
	rl.PushMatrix()
	rl.LoadIdentity()

	r.renderRL(&drawData)

	// Flush all ImGui draw calls before restoring matrices
	rl.DrawRenderBatchActive()

	rl.MatrixMode(rl.Projection)
	rl.PopMatrix()
	rl.MatrixMode(rl.Modelview)
	rl.PopMatrix()
}

// SetAfterCreateContextHook sets the callback to run after ImGui context is created
func (r *RaylibBackend) SetAfterCreateContextHook(f func()) { r.afterCreateContext = f }

// SetBeforeDestroyContextHook sets the callback to run before ImGui context is destroyed
func (r *RaylibBackend) SetBeforeDestroyContextHook(f func()) { r.beforeDestroy = f }

// SetBeforeImGuiRenderHook sets the callback to run before ImGui rendering begins
func (r *RaylibBackend) SetBeforeImGuiRenderHook(f func()) { r.beforeImGuiRender = f }

// SetBeforeRenderHook sets the callback to run before rendering
func (r *RaylibBackend) SetBeforeRenderHook(f func()) { r.beforeRender = f }

// SetAfterRenderHook sets the callback to run after rendering
func (r *RaylibBackend) SetAfterRenderHook(f func()) { r.afterRender = f }

// SetBgColor sets the background color for the window
func (r *RaylibBackend) SetBgColor(color imgui.Vec4) { r.bgColor = color }

// Run starts the main loop of the backend
func (r *RaylibBackend) Run(f func()) {
	r.loop = f

	if r.afterCreateContext != nil {
		r.afterCreateContext()
	}

	for !rl.WindowShouldClose() && !r.shouldClose {
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(
			uint8(r.bgColor.X*255),
			uint8(r.bgColor.Y*255),
			uint8(r.bgColor.Z*255),
			uint8(r.bgColor.W*255),
		))

		if r.beforeImGuiRender != nil {
			r.beforeImGuiRender()
			rl.DrawRenderBatchActive()
		}

		r.NewFrame()
		imgui.NewFrame()

		if r.beforeRender != nil {
			r.beforeRender()
		}

		if r.loop != nil {
			r.loop()
		}

		imgui.Render()

		dd := imgui.CurrentDrawData()
		if dd != nil {
			r.Render(*dd)
		}

		if r.afterRender != nil {
			r.afterRender()
		}

		rl.EndDrawing()

		// Check for file drops
		if r.dropCb != nil && rl.IsFileDropped() {
			files := rl.LoadDroppedFiles()
			if len(files) > 0 {
				r.dropCb(files)
			}

			rl.UnloadDroppedFiles()
		}

		if r.resizeCb != nil {
			w, h := rl.GetScreenWidth(), rl.GetScreenHeight()
			if r.lastW != w || r.lastH != h {
				r.lastW, r.lastH = w, h
				r.resizeCb(r.lastW, r.lastH)
			}
		}
	}

	if r.closeCb != nil {
		r.closeCb()
	}

	if r.beforeDestroy != nil {
		r.beforeDestroy()
	}

	rl.CloseWindow()
}

// Refresh does nothing as raylib draws each frame in Run loop
func (r *RaylibBackend) Refresh() { /* raylib draws each frame in Run loop */ }

// SetWindowPos sets the window position
func (r *RaylibBackend) SetWindowPos(x, y int) { rl.SetWindowPosition(x, y) }

// GetWindowPos returns the current window position
func (r *RaylibBackend) GetWindowPos() (x, y int32) {
	pos := rl.GetWindowPosition()

	return int32(pos.X), int32(pos.Y)
}

// SetWindowSize sets the window size
func (r *RaylibBackend) SetWindowSize(width, height int) { rl.SetWindowSize(width, height) }

// SetWindowSizeLimits sets the minimum and maximum window size limits
func (r *RaylibBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	if minWidth > 0 && minHeight > 0 {
		rl.SetWindowMinSize(minWidth, minHeight)
	}

	if maxWidth > 0 && maxHeight > 0 {
		rl.SetWindowMaxSize(maxWidth, maxHeight)
	}
}

// SetWindowTitle sets the window title
func (r *RaylibBackend) SetWindowTitle(title string) { rl.SetWindowTitle(title) }

// DisplaySize returns the current display/screen size
func (r *RaylibBackend) DisplaySize() (width, height int32) {
	return int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight())
}

// SetShouldClose sets whether the window should close
func (r *RaylibBackend) SetShouldClose(b bool) { r.shouldClose = b }

// ContentScale returns the DPI scale factors for the window
func (r *RaylibBackend) ContentScale() (xScale, yScale float32) {
	s := rl.GetWindowScaleDPI()

	return s.X, s.Y
}

// SetTargetFPS sets the target frames per second
func (r *RaylibBackend) SetTargetFPS(fps uint) { rl.SetTargetFPS(int32(fps)) }

// SetDropCallback sets the callback for file drop events
func (r *RaylibBackend) SetDropCallback(callback backend.DropCallback) { r.dropCb = callback }

// SetCloseCallback sets the callback for window close events
func (r *RaylibBackend) SetCloseCallback(callback backend.WindowCloseCallback) { r.closeCb = callback }

// SetKeyCallback sets the callback for keyboard events
func (r *RaylibBackend) SetKeyCallback(callback backend.KeyCallback) { r.keyCb = callback }

// SetSizeChangeCallback sets the callback for window size change events
func (r *RaylibBackend) SetSizeChangeCallback(callback backend.SizeChangeCallback) {
	r.resizeCb = callback
}

// SetWindowFlags sets or clears a window flag
func (r *RaylibBackend) SetWindowFlags(flag RaylibBackendFlags, value int) {
	// Map backend flags to raylib flags
	var rlFlag uint32

	switch flag {
	case RaylibBackendFlagsVsyncHint:
		rlFlag = rl.FlagVsyncHint
	case RaylibBackendFlagsFullscreen:
		rlFlag = rl.FlagFullscreenMode
	case RaylibBackendFlagsResizable:
		rlFlag = rl.FlagWindowResizable
	case RaylibBackendFlagsUndecorated:
		rlFlag = rl.FlagWindowUndecorated
	case RaylibBackendFlagsHidden:
		rlFlag = rl.FlagWindowHidden
	case RaylibBackendFlagsMinimized:
		rlFlag = rl.FlagWindowMinimized
	case RaylibBackendFlagsMaximized:
		rlFlag = rl.FlagWindowMaximized
	case RaylibBackendFlagsUnfocused:
		rlFlag = rl.FlagWindowUnfocused
	case RaylibBackendFlagsTopmost:
		rlFlag = rl.FlagWindowTopmost
	case RaylibBackendFlagsAlwaysRun:
		rlFlag = rl.FlagWindowAlwaysRun
	case RaylibBackendFlagsTransparent:
		rlFlag = rl.FlagWindowTransparent
	case RaylibBackendFlagsHighDPI:
		rlFlag = rl.FlagWindowHighdpi
	case RaylibBackendFlagsMousePassthrough:
		rlFlag = rl.FlagWindowMousePassthrough
	case RaylibBackendFlagsBorderlessWindowed:
		rlFlag = rl.FlagBorderlessWindowedMode
	case RaylibBackendFlagsMSAA4X:
		rlFlag = rl.FlagMsaa4xHint
	case RaylibBackendFlagsInterlaced:
		rlFlag = rl.FlagInterlacedHint
	default:
		return
	}

	// Set or clear the window state
	if value == 0 {
		rl.ClearWindowState(rlFlag)
	} else {
		rl.SetWindowState(rlFlag)
	}
}

// SetIcons sets the window icon (uses the first icon if multiple is provided)
func (r *RaylibBackend) SetIcons(icons ...image.Image) {
	if len(icons) == 0 {
		return
	}

	// Use the first icon (raylib SetWindowIcon takes single image)
	rlImg := rl.NewImageFromImage(icons[0])
	rl.SetWindowIcon(*rlImg)
	rl.UnloadImage(rlImg)
}

// SetSwapInterval sets the swap interval for VSync (RaylibSwapIntervalDisabled = off, RaylibSwapIntervalEnabled = on)
func (r *RaylibBackend) SetSwapInterval(interval RaylibBackendFlags) error {
	switch interval {
	case RaylibSwapIntervalDisabled:
		rl.ClearWindowState(rl.FlagVsyncHint)
	case RaylibSwapIntervalEnabled:
		rl.SetWindowState(rl.FlagVsyncHint)
	default:
		panic("unhandled default case")
	}
	return nil
}

// SetCursorPos sets the mouse cursor position
func (r *RaylibBackend) SetCursorPos(x, y float64) { rl.SetMousePosition(int(x), int(y)) }

// SetInputMode sets the input mode for cursor visibility and enabling
func (r *RaylibBackend) SetInputMode(mode, value RaylibBackendFlags) {
	switch mode {
	case RaylibInputModeCursorVisible:
		switch value {
		case RaylibInputModeDisabled:
			rl.HideCursor()
		case RaylibInputModeEnabled:
			rl.ShowCursor()
		default:
			panic("unhandled default case")
		}
	case RaylibInputModeCursorEnabled:
		switch value {
		case RaylibInputModeDisabled:
			rl.DisableCursor()
		case RaylibInputModeEnabled:
			rl.EnableCursor()
		default:
			panic("unhandled default case")
		}
	default:
		panic("unhandled default case")
	}
}

// SetConfigFlags sets raylib config flags
func (r *RaylibBackend) SetConfigFlags(flags ...RaylibBackendFlags) {
	var rlFlags uint32

	for _, flag := range flags {
		switch flag {
		case RaylibBackendFlagsVsyncHint:
			rlFlags |= rl.FlagVsyncHint
		case RaylibBackendFlagsFullscreen:
			rlFlags |= rl.FlagFullscreenMode
		case RaylibBackendFlagsResizable:
			rlFlags |= rl.FlagWindowResizable
		case RaylibBackendFlagsUndecorated:
			rlFlags |= rl.FlagWindowUndecorated
		case RaylibBackendFlagsHidden:
			rlFlags |= rl.FlagWindowHidden
		case RaylibBackendFlagsMinimized:
			rlFlags |= rl.FlagWindowMinimized
		case RaylibBackendFlagsMaximized:
			rlFlags |= rl.FlagWindowMaximized
		case RaylibBackendFlagsUnfocused:
			rlFlags |= rl.FlagWindowUnfocused
		case RaylibBackendFlagsTopmost:
			rlFlags |= rl.FlagWindowTopmost
		case RaylibBackendFlagsAlwaysRun:
			rlFlags |= rl.FlagWindowAlwaysRun
		case RaylibBackendFlagsTransparent:
			rlFlags |= rl.FlagWindowTransparent
		case RaylibBackendFlagsHighDPI:
			rlFlags |= rl.FlagWindowHighdpi
		case RaylibBackendFlagsMousePassthrough:
			rlFlags |= rl.FlagWindowMousePassthrough
		case RaylibBackendFlagsBorderlessWindowed:
			rlFlags |= rl.FlagBorderlessWindowedMode
		case RaylibBackendFlagsMSAA4X:
			rlFlags |= rl.FlagMsaa4xHint
		case RaylibBackendFlagsInterlaced:
			rlFlags |= rl.FlagInterlacedHint
		default:
			panic("Unhandled RaylibBackendFlag")
		}
	}

	rl.SetConfigFlags(rlFlags)
}

// CreateWindow creates and initializes the raylib window and ImGui context
func (r *RaylibBackend) CreateWindow(title string, width, height int) {
	// initialize ImGui
	r.initImGuiContext()

	// reduce log noise from raylib (suppress GLFW Wayland warnings)
	rl.SetTraceLogLevel(rl.LogError)

	// init raylib window
	rl.InitWindow(int32(width), int32(height), title)

	rl.SetWindowState(rl.FlagVsyncHint)
	rl.SetTargetFPS(60)

	r.lastW, r.lastH = width, height
}

// CreateTexture creates a new ImGui texture reference.
func (r *RaylibBackend) CreateTexture(pixels unsafe.Pointer, width, height int) imgui.TextureRef {
	// Upload raw RGBA32 pixels to a raylib texture
	n := width * height * 4
	src := (*[1 << 28]uint8)(pixels)[:n:n]

	rlImg := rl.NewImageFromImage(imageRGBA(width, height, src))
	tx := rl.LoadTextureFromImage(rlImg)
	rl.UnloadImage(rlImg)

	rl.SetTextureFilter(tx, rl.FilterBilinear)
	rl.SetTextureWrap(tx, rl.WrapClamp)

	tid := imgui.TextureID(r.cache.NextId())
	r.cache.SetTexture(tid, tx)

	return *imgui.NewTextureRefTextureID(tid)
}

// CreateTextureRgba creates an ImGui texture reference from an RGBA image
func (r *RaylibBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureRef {
	rlImg := rl.NewImageFromImage(imageRGBA(width, height, img.Pix))
	tx := rl.LoadTextureFromImage(rlImg)

	rl.SetTextureFilter(tx, rl.FilterBilinear)
	rl.SetTextureWrap(tx, rl.WrapClamp)
	rl.UnloadImage(rlImg)

	tid := imgui.TextureID(r.cache.NextId())
	r.cache.SetTexture(tid, tx)

	return *imgui.NewTextureRefTextureID(tid)
}

// DeleteTexture removes a texture from the cache
func (r *RaylibBackend) DeleteTexture(id imgui.TextureRef) { r.cache.RemoveTexture(id.TexID()) }

// SetTextureFilterForRef sets the texture filter mode for a texture reference
func (r *RaylibBackend) SetTextureFilterForRef(id imgui.TextureRef, filter rl.TextureFilterMode) {
	if r.cache.HasTexture(id.TexID()) {
		tex := r.cache.GetTexture(id.TexID())
		rl.SetTextureFilter(tex, filter)
	}
}

// SetTextureWrapForRef sets the texture wrap mode for a texture reference
func (r *RaylibBackend) SetTextureWrapForRef(id imgui.TextureRef, wrap rl.TextureWrapMode) {
	if r.cache.HasTexture(id.TexID()) {
		tex := r.cache.GetTexture(id.TexID())
		rl.SetTextureWrap(tex, wrap)
	}
}

// CreateTextureFromTexture2D registers an existing raylib Texture2D in the ImGui texture cache
func (r *RaylibBackend) CreateTextureFromTexture2D(tex rl.Texture2D) imgui.TextureRef {
	if ref, ok := r.texRefCache[tex.ID]; ok {
		return ref
	}

	rl.SetTextureFilter(tex, rl.FilterBilinear)
	rl.SetTextureWrap(tex, rl.WrapClamp)

	tid := imgui.TextureID(r.cache.NextId())
	r.cache.SetTexture(tid, tex)

	ref := *imgui.NewTextureRefTextureID(tid)
	r.texRefCache[tex.ID] = ref
	return ref
}

// AfterCreateHook returns the after-create-context hook function
func (r *RaylibBackend) AfterCreateHook() func() { return r.afterCreateContext }

// BeforeRenderHook returns the before-render hook function
func (r *RaylibBackend) BeforeRenderHook() func() { return r.beforeRender }

// LoopFunc returns the main loop function
func (r *RaylibBackend) LoopFunc() func() { return r.loop }

// AfterRenderHook returns the after-render hook function
func (r *RaylibBackend) AfterRenderHook() func() { return r.afterRender }

// BeforeDestroyHook returns the before-destroy hook function
func (r *RaylibBackend) BeforeDestroyHook() func() { return r.beforeDestroy }

// DropCallback returns the file drop callback
func (r *RaylibBackend) DropCallback() backend.DropCallback { return r.dropCb }

// CloseCallback returns the window close callback
func (r *RaylibBackend) CloseCallback() func(window unsafe.Pointer) {
	return func(_ unsafe.Pointer) {
		if r.closeCb != nil {
			r.closeCb()
		}
	}
}

// KeyCallback returns the keyboard callback
func (r *RaylibBackend) KeyCallback() backend.KeyCallback { return r.keyCb }

// SizeCallback returns the size change callback
func (r *RaylibBackend) SizeCallback() backend.SizeChangeCallback { return r.resizeCb }

// MaximizeWindow maximizes the window
func (r *RaylibBackend) MaximizeWindow() { rl.MaximizeWindow() }

// MinimizeWindow minimizes the window
func (r *RaylibBackend) MinimizeWindow() { rl.MinimizeWindow() }

// RestoreWindow restores the window from a minimized or maximized state
func (r *RaylibBackend) RestoreWindow() { rl.RestoreWindow() }

// SetWindowOpacity sets the window opacity
func (r *RaylibBackend) SetWindowOpacity(opacity float32) { rl.SetWindowOpacity(opacity) }

// ToggleFullscreen toggles fullscreen mode
func (r *RaylibBackend) ToggleFullscreen() { rl.ToggleFullscreen() }

// ToggleBorderlessWindowed toggles borderless windowed mode
func (r *RaylibBackend) ToggleBorderlessWindowed() { rl.ToggleBorderlessWindowed() }

// IsWindowState checks if a window state flag is set
func (r *RaylibBackend) IsWindowState(flag uint32) bool { return rl.IsWindowState(flag) }

// IsWindowFullscreen checks if the window is in fullscreen mode
func (r *RaylibBackend) IsWindowFullscreen() bool { return rl.IsWindowFullscreen() }

// IsWindowMaximized checks if the window is maximized
func (r *RaylibBackend) IsWindowMaximized() bool { return rl.IsWindowMaximized() }

// IsWindowMinimized checks if the window is minimized
func (r *RaylibBackend) IsWindowMinimized() bool { return rl.IsWindowMinimized() }

// IsWindowFocused checks if the window has focus
func (r *RaylibBackend) IsWindowFocused() bool { return rl.IsWindowFocused() }

// IsWindowHidden checks if the window is hidden
func (r *RaylibBackend) IsWindowHidden() bool { return rl.IsWindowHidden() }

// IsWindowResized checks if the window was resized
func (r *RaylibBackend) IsWindowResized() bool { return rl.IsWindowResized() }

// ShowCursor shows the mouse cursor
func (r *RaylibBackend) ShowCursor() { rl.ShowCursor() }

// HideCursor hides the mouse cursor
func (r *RaylibBackend) HideCursor() { rl.HideCursor() }

// EnableCursor enables the mouse cursor
func (r *RaylibBackend) EnableCursor() { rl.EnableCursor() }

// DisableCursor disables the mouse cursor
func (r *RaylibBackend) DisableCursor() { rl.DisableCursor() }

// IsCursorHidden checks if the cursor is hidden
func (r *RaylibBackend) IsCursorHidden() bool { return rl.IsCursorHidden() }

// IsCursorOnScreen checks if the cursor is on screen
func (r *RaylibBackend) IsCursorOnScreen() bool { return rl.IsCursorOnScreen() }

// SetMouseCursor sets the mouse cursor style
func (r *RaylibBackend) SetMouseCursor(cursor int32) { rl.SetMouseCursor(cursor) }

// GetClipboardText returns the clipboard text
func (r *RaylibBackend) GetClipboardText() string { return rl.GetClipboardText() }

// SetClipboardText sets the clipboard text
func (r *RaylibBackend) SetClipboardText(text string) { rl.SetClipboardText(text) }

// GetCurrentMonitor returns the current monitor index
func (r *RaylibBackend) GetCurrentMonitor() int { return rl.GetCurrentMonitor() }

// GetMonitorCount returns the number of monitors
func (r *RaylibBackend) GetMonitorCount() int { return rl.GetMonitorCount() }

// GetMonitorWidth returns the width of the specified monitor
func (r *RaylibBackend) GetMonitorWidth(monitor int) int { return rl.GetMonitorWidth(monitor) }

// GetMonitorHeight returns the height of the specified monitor
func (r *RaylibBackend) GetMonitorHeight(monitor int) int { return rl.GetMonitorHeight(monitor) }

// GetMonitorPhysicalWidth returns the physical width of the specified monitor in millimeters
func (r *RaylibBackend) GetMonitorPhysicalWidth(monitor int) int {
	return rl.GetMonitorPhysicalWidth(monitor)
}

// GetMonitorPhysicalHeight returns the physical height of the specified monitor in millimeters
func (r *RaylibBackend) GetMonitorPhysicalHeight(monitor int) int {
	return rl.GetMonitorPhysicalHeight(monitor)
}

// GetMonitorRefreshRate returns the refresh rate of the specified monitor
func (r *RaylibBackend) GetMonitorRefreshRate(monitor int) int {
	return rl.GetMonitorRefreshRate(monitor)
}

// GetMonitorName returns the name of the specified monitor
func (r *RaylibBackend) GetMonitorName(monitor int) string { return rl.GetMonitorName(monitor) }

// GetMonitorPosition returns the position of the specified monitor
func (r *RaylibBackend) GetMonitorPosition(monitor int) (x, y int32) {
	pos := rl.GetMonitorPosition(monitor)

	return int32(pos.X), int32(pos.Y)
}

// SetWindowMonitor sets the monitor that the window should use
func (r *RaylibBackend) SetWindowMonitor(monitor int) { rl.SetWindowMonitor(monitor) }

// GetFPS returns the current frames per second
func (r *RaylibBackend) GetFPS() int32 { return rl.GetFPS() }

// GetFrameTime returns the time in seconds for the last frame drawn
func (r *RaylibBackend) GetFrameTime() float32 { return rl.GetFrameTime() }

// GetTime returns the elapsed time in seconds since initialization
func (r *RaylibBackend) GetTime() float64 { return rl.GetTime() }

// TakeScreenshot takes a screenshot and saves it to the specified file
func (r *RaylibBackend) TakeScreenshot(filename string) { rl.TakeScreenshot(filename) }

// OpenURL opens a URL in the default browser
func (r *RaylibBackend) OpenURL(url string) { rl.OpenURL(url) }

// GetRenderWidth returns the render width (framebuffer width)
func (r *RaylibBackend) GetRenderWidth() int { return rl.GetRenderWidth() }

// GetRenderHeight returns the render height (framebuffer height)
func (r *RaylibBackend) GetRenderHeight() int { return rl.GetRenderHeight() }

// SetExitKey sets the key to exit the application
func (r *RaylibBackend) SetExitKey(key int32) { rl.SetExitKey(key) }

// GetMouseX returns the mouse X position
func (r *RaylibBackend) GetMouseX() int32 { return rl.GetMouseX() }

// GetMouseY returns the mouse Y position
func (r *RaylibBackend) GetMouseY() int32 { return rl.GetMouseY() }

// GetMouseDelta returns the mouse position delta between frames
func (r *RaylibBackend) GetMouseDelta() (x, y float32) {
	delta := rl.GetMouseDelta()

	return delta.X, delta.Y
}

// GetMouseWheelMove returns the mouse wheel movement for Y axis
func (r *RaylibBackend) GetMouseWheelMove() float32 { return rl.GetMouseWheelMove() }

// GetMouseWheelMoveV returns the mouse wheel movement for both X and Y axes
func (r *RaylibBackend) GetMouseWheelMoveV() (x, y float32) {
	wheel := rl.GetMouseWheelMoveV()

	return wheel.X, wheel.Y
}

// GetKeyPressed returns the last key pressed
func (r *RaylibBackend) GetKeyPressed() int32 { return rl.GetKeyPressed() }

// GetCharPressed returns the last character pressed
func (r *RaylibBackend) GetCharPressed() int32 { return rl.GetCharPressed() }

// IsKeyDown checks if a key is currently down
func (r *RaylibBackend) IsKeyDown(key int32) bool { return rl.IsKeyDown(key) }

// IsKeyUp checks if a key is currently up
func (r *RaylibBackend) IsKeyUp(key int32) bool { return rl.IsKeyUp(key) }

// IsKeyPressed checks if a key was pressed once
func (r *RaylibBackend) IsKeyPressed(key int32) bool { return rl.IsKeyPressed(key) }

// IsKeyReleased checks if a key was released once
func (r *RaylibBackend) IsKeyReleased(key int32) bool { return rl.IsKeyReleased(key) }

// SetGesturesEnabled enables gesture detection for specific gesture types
func (r *RaylibBackend) SetGesturesEnabled(flags uint32) { rl.SetGesturesEnabled(flags) }

// IsGestureDetected checks if a specific gesture was detected
func (r *RaylibBackend) IsGestureDetected(gesture rl.Gestures) bool {
	return rl.IsGestureDetected(gesture)
}

// GetGestureDetected returns the latest detected gesture
func (r *RaylibBackend) GetGestureDetected() rl.Gestures { return rl.GetGestureDetected() }

// GetRandomValue generates a random integer between min and max
func (r *RaylibBackend) GetRandomValue(min, max int32) int32 { return rl.GetRandomValue(min, max) }
