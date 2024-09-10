package ebitenbackend

import (
	"fmt"
	"image"
	"runtime"
	"slices"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EbitenBackendFlags int

const (
	// EbitenBackendFlagsCursorMode sets the cursor mode.
	// refer ebiten.CursorModeType
	// CursorModeVisible, CursorModeHidden, CursorModeCaptured
	EbitenBackendFlagsCursorMode EbitenBackendFlags = iota
	// EbitenBackendFlagsCursorShape sets the cursor shape.
	// refer ebiten.CursorShapeType
	// CursorShapeDefault, CursorShapeText, CursorShapeCrosshair, CursorShapePointer,
	// CursorShapeEWResize, CursorShapeNSResize, CursorShapeNESWResize, CursorShapeNWSEResize,
	// CursorShapeMove, CursorShapeNotAllowed CursorShapeType, EbitenBackendFlagsCursorShape
	EbitenBackendFlagsCursorShape
	EbitenBackendFlagsResizingMode
	EbitenBackendFlagsFPSMode
	EbitenBackendFlagsDecorated
	EbitenBackendFlagsFloating
	EbitenBackendFlagsMaximized
	EbitenBackendFlagsMinimized
	EbitenBackendFlagsClosingHandled
	EbitenBackendFlagsMousePassthrough
)

var (
	_ imgui.Backend[EbitenBackendFlags] = &EbitenBackend{}
	_ ebiten.Game                       = &EbitenBackend{}
)

type EbitenBackend struct {
	afterCreateContext,
	beforeRender,
	afterRender,
	beforeDestroy,
	loop func()

	// callbacks
	closeCb imgui.WindowCloseCallback[EbitenBackendFlags]

	// ebiten stuff
	dscale  float64
	retina  bool
	w, h    int
	manager *Manager
	fps     uint
	bgColor imgui.Vec4

	cache TextureCache
	ctx   *imgui.Context
}

// this is "pointer" to the first texture used for font atlas texture.
// it needs to be a var and cannot be more private because of CGO stuff.
var id1 = 1

// NewEbitenBackend creates a new Ebiten backend.
// it takes font atlas which could be nil
// - TODO: make font atlas a factory method and move some stuff to CreateWindow
func NewEbitenBackend(fontAtlas *imgui.FontAtlas) *EbitenBackend {
	var imctx *imgui.Context

	if fontAtlas != nil {
		imctx = imgui.CreateContextV(fontAtlas)
	} else {
		imctx = imgui.CreateContext()
	}

	result := &EbitenBackend{
		ctx:     imctx,
		cache:   NewCache(),
		manager: NewManager(nil),
		fps:     60,
	}

	// Build texture atlas
	fonts := imgui.CurrentIO().Fonts()
	_, _, _, _ = fonts.GetTextureDataAsRGBA32() // call this to force imgui to build the font atlas cache

	texID := imgui.TextureID{}
	texID.Data = uintptr(unsafe.Pointer(&id1)) // TODO: this shit will cause -race issues
	fonts.SetTexID(texID)

	result.cache.SetFontAtlasTextureID(texID)

	runtime.SetFinalizer(result, (*EbitenBackend).onfinalize)

	return result
}

func (b *EbitenBackend) SetAfterCreateContextHook(fn func()) {
	b.afterCreateContext = fn
}

func (b *EbitenBackend) SetBeforeDestroyContextHook(fn func()) {
	b.beforeDestroy = fn
}

func (b *EbitenBackend) SetBeforeRenderHook(fn func()) {
	b.beforeRender = fn
}

func (b *EbitenBackend) SetAfterRenderHook(fn func()) {
	b.afterRender = fn
}

func (b *EbitenBackend) SetBgColor(color imgui.Vec4) {
	b.bgColor = color
}

func (b *EbitenBackend) Run(loop func()) {
	b.loop = loop
	b.dscale = ebiten.DeviceScaleFactor()

	if b.afterCreateContext != nil {
		b.afterCreateContext()
	}

	ebiten.RunGame(b)
}

func (b *EbitenBackend) Refresh() {}

func (b *EbitenBackend) SetWindowPos(x, y int) {
	ebiten.SetWindowPosition(x, y)
}

func (b *EbitenBackend) GetWindowPos() (x, y int32) {
	xInt, yInt := ebiten.WindowPosition()
	return int32(xInt), int32(yInt)
}

func (b *EbitenBackend) SetWindowSize(width, height int) {
	ebiten.SetWindowSize(width, height)
	b.w = width
	b.h = height
}

func (b *EbitenBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {}

func (b *EbitenBackend) SetWindowTitle(title string) {
	ebiten.SetWindowTitle(title)
}

func (b *EbitenBackend) DisplaySize() (width, height int32)     { return 0, 0 }
func (b *EbitenBackend) SetShouldClose(bool)                    {}
func (b *EbitenBackend) ContentScale() (xScale, yScale float32) { return 1, 1 }

func (e *EbitenBackend) SetTargetFPS(fps uint) {
	e.fps = fps
}

func (b *EbitenBackend) SetDropCallback(imgui.DropCallback) {}

func (b *EbitenBackend) SetCloseCallback(cb imgui.WindowCloseCallback[EbitenBackendFlags]) {
	b.closeCb = cb
}

func (b *EbitenBackend) SetKeyCallback(imgui.KeyCallback)               {}
func (b *EbitenBackend) SetSizeChangeCallback(imgui.SizeChangeCallback) {}

func (b *EbitenBackend) SetWindowFlags(flag EbitenBackendFlags, value int) {
	switch flag {
	case EbitenBackendFlagsCursorMode:
		ebiten.SetCursorMode(ebiten.CursorModeType(value))
	case EbitenBackendFlagsCursorShape:
		ebiten.SetCursorShape(ebiten.CursorShapeType(value))
	case EbitenBackendFlagsFPSMode:
		ebiten.SetVsyncEnabled(value == 0)
	case EbitenBackendFlagsResizingMode:
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeType(value))
	case EbitenBackendFlagsDecorated:
		ebiten.SetWindowDecorated(value != 0)
	case EbitenBackendFlagsFloating:
		ebiten.SetWindowFloating(value != 0)
	case EbitenBackendFlagsMaximized:
		if value != 0 {
			ebiten.MaximizeWindow()
		} else {
			ebiten.RestoreWindow()
		}
	case EbitenBackendFlagsMinimized:
		if value != 0 {
			ebiten.MinimizeWindow()
		} else {
			ebiten.RestoreWindow()
		}
	case EbitenBackendFlagsClosingHandled:
		ebiten.SetWindowClosingHandled(value != 0)
	case EbitenBackendFlagsMousePassthrough:
		ebiten.SetWindowMousePassthrough(value != 0)
	default:
		panic("Invalid flag for SetWindowFlags.")
	}
}

func (b *EbitenBackend) SetIcons(icons ...image.Image) {
	ebiten.SetWindowIcon(icons)
}

func (b *EbitenBackend) SetSwapInterval(interval EbitenBackendFlags) error { return nil }
func (b *EbitenBackend) SetCursorPos(x, y float64)                         {}
func (b *EbitenBackend) SetInputMode(mode, value EbitenBackendFlags)       {}

func (e *EbitenBackend) CreateWindow(title string, width, height int) {
	e.SetWindowTitle(title)
	e.SetWindowSize(width, height)
}

func (e *EbitenBackend) CreateTexture(pixels unsafe.Pointer, width, height int) imgui.TextureID {
	eimg := ebiten.NewImage(width, height)
	eimg.WritePixels(premultiplyPixels(pixels, width, height))

	tid := imgui.TextureID{Data: uintptr(e.cache.NextId())}
	e.cache.SetTexture(tid, eimg)
	return tid
}

func (b *EbitenBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureID {
	pix := img.Pix
	return b.CreateTexture(unsafe.Pointer(&pix[0]), width, height)
}

func (e *EbitenBackend) DeleteTexture(id imgui.TextureID) {
	e.cache.RemoveTexture(id)
}

// ebiten

// Draw draws the generated imgui frame to the screen.
// This is usually called inside the game's Draw() function.
func (e *EbitenBackend) Draw(screen *ebiten.Image) {
	// add background color
	bgRect := slices.Repeat([]byte{
		byte(e.bgColor.X * 255),
		byte(e.bgColor.Y * 255),
		byte(e.bgColor.Z * 255),
		byte(e.bgColor.W * 255),
	}, e.w*e.h)
	screen.WritePixels(bgRect)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f\nFPS: %.2f\n[C]lipMask: %t", ebiten.ActualTPS(), ebiten.ActualFPS(), e.ClipMask()), 10, 2)

	e.manager.screenWidth = screen.Bounds().Dx()
	e.manager.screenHeight = screen.Bounds().Dy()
	imgui.Render()
	if e.manager.ClipMask {
		if e.manager.lmask == nil {
			w, h := screen.Size()
			e.manager.lmask = ebiten.NewImage(w, h)
		} else {
			w1, h1 := screen.Size()
			w2, h2 := e.manager.lmask.Size()
			if w1 != w2 || h1 != h2 {
				e.manager.lmask.Dispose()
				e.manager.lmask = ebiten.NewImage(w1, h1)
			}
		}
		RenderMasked(screen, e.manager.lmask, imgui.CurrentDrawData(), e.cache, e.manager.Filter)
	} else {
		Render(screen, imgui.CurrentDrawData(), e.cache, e.manager.Filter)
	}
}

// Update needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
// delta is the time in seconds since the last frame.
func (e *EbitenBackend) Update() error {
	if ebiten.IsWindowBeingClosed() {
		e.closeCb(e)
	}

	if e.beforeRender != nil {
		e.beforeRender()
	}

	io := imgui.CurrentIO()
	if e.manager.width > 0 || e.manager.height > 0 {
		io.SetDisplaySize(imgui.Vec2{X: e.manager.width, Y: e.manager.height})
	} else if e.manager.screenWidth > 0 || e.manager.screenHeight > 0 {
		io.SetDisplaySize(imgui.Vec2{X: float32(e.manager.screenWidth), Y: float32(e.manager.screenHeight)})
	}

	io.SetDeltaTime(1.0 / float32(e.fps))
	if e.manager.SyncCursor {
		if e.manager.GetCursor != nil {
			x, y := e.manager.GetCursor()
			io.SetMousePos(imgui.Vec2{X: x, Y: y})
		} else {
			mx, my := ebiten.CursorPosition()
			io.SetMousePos(imgui.Vec2{X: float32(mx), Y: float32(my)})
		}
		io.SetMouseButtonDown(0, ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))
		io.SetMouseButtonDown(1, ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight))
		io.SetMouseButtonDown(2, ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle))
		xoff, yoff := ebiten.Wheel()
		io.AddMouseWheelDelta(float32(xoff), float32(yoff))
		e.manager.controlCursorShape()
	}

	if e.manager.SyncInputs {
		if e.manager.SyncInputsFn != nil {
			e.manager.SyncInputsFn()
		} else {
			e.manager.inputChars = sendInput(imgui.CurrentIO(), e.manager.inputChars)
		}
	}

	// TODO: what is that?
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		e.SetClipMask(!e.ClipMask())
	}

	e.BeginFrame()
	defer e.EndFrame()

	e.loop()

	defer func() {
		if e.afterRender != nil {
			e.afterRender()
		}
	}()

	return nil
}

func (e *EbitenBackend) Layout(outsideWidth, outsideHeight int) (int, int) {
	if e.retina {
		m := ebiten.DeviceScaleFactor()
		e.w = int(float64(outsideWidth) * m)
		e.h = int(float64(outsideHeight) * m)
	} else {
		e.w = outsideWidth
		e.h = outsideHeight
	}

	e.SetDisplaySize(float32(e.w), float32(e.h))

	return e.w, e.h
}

// Text implements imgui clipboard
func (e *EbitenBackend) Text() (string, error) {
	return e.manager.cliptxt, nil
}

// SetText implements imgui clipboard
func (e *EbitenBackend) SetText(text string) {
	e.manager.cliptxt = text
}

func (e *EbitenBackend) onfinalize() {
	if e.beforeDestroy != nil {
		e.beforeDestroy()
	}

	runtime.SetFinalizer(e, nil)
	e.ctx.Destroy()
}
