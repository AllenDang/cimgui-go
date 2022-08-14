package cimgui

import (
	"fmt"
	"image"
	"math"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var GlfwDontCare int = glfw.DontCare

type GLFWClipboard struct {
	window *glfw.Window
}

func NewGLFWClipboard(w *glfw.Window) *GLFWClipboard {
	return &GLFWClipboard{window: w}
}

func (c *GLFWClipboard) Text() (string, error) {
	return c.window.GetClipboardString(), nil
}

func (c *GLFWClipboard) SetText(text string) {
	c.window.SetClipboardString(text)
}

type GLFWWindowFlags uint8

const (
	GLFWWindowFlagsNotResizable GLFWWindowFlags = 1 << iota
	GLFWWindowFlagsMaximized
	GLFWWindowFlagsFloating
	GLFWWindowFlagsFrameless
	GLFWWindowFlagsTransparent
)

// GLFW implements a platform based on github.com/go-gl/glfw (v3.3).
type GLFW struct {
	imguiIO ImGuiIO

	window *glfw.Window

	tps              int
	time             float64
	mouseJustPressed [3]bool

	mouseCursors map[ImGuiMouseCursor]*glfw.Cursor

	posChangeCallback  func(int, int)
	sizeChangeCallback func(int, int)
	dropCallback       func([]string)
	inputCallback      func(key glfw.Key, mods glfw.ModifierKey, action glfw.Action)
	closeCallback      func() bool
}

// NewGLFW attempts to initialize a GLFW context.
func NewGLFW(io ImGuiIO, title string, width, height int, flags GLFWWindowFlags) (*GLFW, error) {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize glfw: %v", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)

	if flags&GLFWWindowFlagsNotResizable != 0 {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	if flags&GLFWWindowFlagsMaximized != 0 {
		glfw.WindowHint(glfw.Maximized, glfw.True)
	}

	if flags&GLFWWindowFlagsFloating != 0 {
		glfw.WindowHint(glfw.Floating, glfw.True)
	}

	if flags&GLFWWindowFlagsFrameless != 0 {
		glfw.WindowHint(glfw.Decorated, glfw.False)
	}

	if flags&GLFWWindowFlagsTransparent != 0 {
		glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("failed to create window: %v", err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	platform := &GLFW{
		imguiIO: io,
		window:  window,

		tps: 60,
	}
	platform.installCallbacks()

	// Create mosue cursors
	platform.mouseCursors = make(map[ImGuiMouseCursor]*glfw.Cursor)
	platform.mouseCursors[ImGuiMouseCursor_Arrow] = glfw.CreateStandardCursor(glfw.ArrowCursor)
	platform.mouseCursors[ImGuiMouseCursor_TextInput] = glfw.CreateStandardCursor(glfw.IBeamCursor)
	platform.mouseCursors[ImGuiMouseCursor_ResizeAll] = glfw.CreateStandardCursor(glfw.CrosshairCursor)
	platform.mouseCursors[ImGuiMouseCursor_Hand] = glfw.CreateStandardCursor(glfw.HandCursor)
	platform.mouseCursors[ImGuiMouseCursor_ResizeEW] = glfw.CreateStandardCursor(glfw.HResizeCursor)
	platform.mouseCursors[ImGuiMouseCursor_ResizeNS] = glfw.CreateStandardCursor(glfw.VResizeCursor)

	// io.SetClipboard(NewGLFWClipboard(window))

	if flags&GLFWWindowFlagsMaximized == 0 {
		// Center window to monitor
		platform.centerWindow()
	}

	platform.window.Show()

	return platform, nil
}

// Dispose cleans up the resources.
func (platform *GLFW) Dispose() {
	platform.window.Destroy()
	glfw.Terminate()
}

func (platform *GLFW) GetContentScale() float32 {
	x, _ := platform.window.GetContentScale()

	// Do not scale on MacOS
	if runtime.GOOS == "darwin" {
		x = 1
	}

	return x
}

func (platform *GLFW) GetWindow() *glfw.Window {
	return platform.window
}

func (platform *GLFW) GetPos() (x, y int) {
	return platform.window.GetPos()
}

func (platform *GLFW) centerWindow() {
	monitor := platform.getBestMonitor()
	if monitor == nil {
		return
	}

	mode := monitor.GetVideoMode()
	if mode == nil {
		return
	}

	monitorX, monitorY := monitor.GetPos()
	windowWidth, windowHeight := platform.window.GetSize()

	platform.window.SetPos(monitorX+(mode.Width-windowWidth)/2, monitorY+(mode.Height-windowHeight)/2)
}

func (platform *GLFW) getBestMonitor() *glfw.Monitor {
	monitors := glfw.GetMonitors()

	if len(monitors) == 0 {
		return nil
	}

	width, height := platform.window.GetSize()
	x, y := platform.window.GetPos()

	var bestMonitor *glfw.Monitor
	var bestArea int

	for _, m := range monitors {
		monitorX, monitorY := m.GetPos()
		mode := m.GetVideoMode()
		if mode == nil {
			continue
		}

		areaMinX := int(math.Max(float64(x), float64(monitorX)))
		areaMinY := int(math.Max(float64(y), float64(monitorY)))

		areaMaxX := int(math.Min(float64(x+width), float64(monitorX+mode.Width)))
		areaMaxY := int(math.Min(float64(y+height), float64(monitorY+mode.Height)))

		area := (areaMaxX - areaMinX) * (areaMaxY - areaMinY)

		if area > bestArea {
			bestArea = area
			bestMonitor = m
		}
	}

	return bestMonitor
}

// ShouldStop returns true if the window is to be closed.
func (platform *GLFW) ShouldStop() bool {
	return platform.window.ShouldClose()
}

// SetShouldStop sets whether window should be closed
func (platform *GLFW) SetShouldStop(v bool) {
	platform.window.SetShouldClose(v)
}

// ProcessEvents handles all pending window events.
func (platform *GLFW) ProcessEvents() {
	glfw.WaitEventsTimeout(1)
	glfw.PollEvents()
}

// DisplaySize returns the dimension of the display.
func (platform *GLFW) DisplaySize() [2]float32 {
	w, h := platform.window.GetSize()
	return [2]float32{float32(w), float32(h)}
}

// FramebufferSize returns the dimension of the framebuffer.
func (platform *GLFW) FramebufferSize() [2]float32 {
	w, h := platform.window.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}

// NewFrame marks the begin of a render pass. It forwards all current state to imgui IO.
func (platform *GLFW) NewFrame() {
	// Setup display size (every frame to accommodate for window resizing)
	displaySize := platform.DisplaySize()
	platform.imguiIO.SetDisplaySize(displaySize[0], displaySize[1])

	// Setup time step
	currentTime := glfw.GetTime()
	if platform.time > 0 {
		platform.imguiIO.SetDeltaTime(float32(currentTime - platform.time))
	}
	platform.time = currentTime

	// Setup inputs
	if platform.window.GetAttrib(glfw.Focused) != 0 {
		x, y := platform.window.GetCursorPos()
		platform.imguiIO.SetMousePos(float32(x), float32(y))
	} else {
		platform.imguiIO.SetMousePos(-math.MaxFloat32, -math.MaxFloat32)
	}

	for i := 0; i < len(platform.mouseJustPressed); i++ {
		down := platform.mouseJustPressed[i] || (platform.window.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		platform.imguiIO.SetMouseButtonDown(i, down)
		platform.mouseJustPressed[i] = false
	}

	platform.updateMouseCursor()
}

// PostRender performs a buffer swap.
func (platform *GLFW) PostRender() {
	platform.window.SwapBuffers()
}

func (platform *GLFW) SetPosChangeCallback(cb func(int, int)) {
	platform.posChangeCallback = cb
}

func (platform *GLFW) SetSizeChangeCallback(cb func(int, int)) {
	platform.sizeChangeCallback = cb
}

func (platform *GLFW) Update() {
	glfw.PostEmptyEvent()
}

func (platform *GLFW) SetDropCallback(cb func(names []string)) {
	platform.dropCallback = cb
}

func (platform *GLFW) SetInputCallback(cb func(key glfw.Key, mods glfw.ModifierKey, action glfw.Action)) {
	platform.inputCallback = cb
}

func (platform *GLFW) SetCloseCallback(cb func() bool) {
	platform.closeCallback = cb
}

func (platform *GLFW) updateMouseCursor() {
	io := platform.imguiIO
	if (io.GetConfigFlags()&ImGuiConfigFlags_NoMouseCursorChange) == 1 || platform.window.GetInputMode(glfw.CursorMode) == glfw.CursorDisabled {
		return
	}

	cursor := GetMouseCursor()
	if cursor == ImGuiMouseCursor_None || io.GetMouseDrawCursor() {
		platform.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	} else {
		gCursor := platform.mouseCursors[ImGuiMouseCursor_Arrow]
		if c, ok := platform.mouseCursors[cursor]; ok {
			gCursor = c
		}
		platform.window.SetCursor(gCursor)
		platform.window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func (platform *GLFW) keyToModifier(key glfw.Key) glfw.ModifierKey {
	if key == glfw.KeyLeftControl || key == glfw.KeyRightControl {
		return glfw.ModControl
	}

	if key == glfw.KeyLeftShift || key == glfw.KeyRightShift {
		return glfw.ModShift
	}

	if key == glfw.KeyLeftAlt || key == glfw.KeyRightAlt {
		return glfw.ModAlt
	}

	if key == glfw.KeyLeftSuper || key == glfw.KeyRightSuper {
		return glfw.ModSuper
	}

	return 0
}

func (platform *GLFW) keyToImGuiKey(key glfw.Key) ImGuiKey {
	switch key {
	case glfw.KeyTab:
		return ImGuiKey_Tab
	case glfw.KeyLeft:
		return ImGuiKey_LeftArrow
	case glfw.KeyRight:
		return ImGuiKey_RightArrow
	case glfw.KeyUp:
		return ImGuiKey_UpArrow
	case glfw.KeyDown:
		return ImGuiKey_DownArrow
	case glfw.KeyPageUp:
		return ImGuiKey_PageUp
	case glfw.KeyPageDown:
		return ImGuiKey_PageDown
	case glfw.KeyHome:
		return ImGuiKey_Home
	case glfw.KeyEnd:
		return ImGuiKey_End
	case glfw.KeyInsert:
		return ImGuiKey_Insert
	case glfw.KeyDelete:
		return ImGuiKey_Delete
	case glfw.KeyBackspace:
		return ImGuiKey_Backspace
	case glfw.KeySpace:
		return ImGuiKey_Space
	case glfw.KeyEnter:
		return ImGuiKey_Enter
	case glfw.KeyEscape:
		return ImGuiKey_Escape
	case glfw.KeyApostrophe:
		return ImGuiKey_Apostrophe
	case glfw.KeyComma:
		return ImGuiKey_Comma
	case glfw.KeyMinus:
		return ImGuiKey_Minus
	case glfw.KeyPeriod:
		return ImGuiKey_Period
	case glfw.KeySlash:
		return ImGuiKey_Slash
	case glfw.KeySemicolon:
		return ImGuiKey_Semicolon
	case glfw.KeyEqual:
		return ImGuiKey_Equal
	case glfw.KeyLeftBracket:
		return ImGuiKey_LeftBracket
	case glfw.KeyBackslash:
		return ImGuiKey_Backslash
	case glfw.KeyRightBracket:
		return ImGuiKey_RightBracket
	case glfw.KeyGraveAccent:
		return ImGuiKey_GraveAccent
	case glfw.KeyCapsLock:
		return ImGuiKey_CapsLock
	case glfw.KeyScrollLock:
		return ImGuiKey_ScrollLock
	case glfw.KeyNumLock:
		return ImGuiKey_NumLock
	case glfw.KeyPrintScreen:
		return ImGuiKey_PrintScreen
	case glfw.KeyPause:
		return ImGuiKey_Pause
	case glfw.KeyKP0:
		return ImGuiKey_Keypad0
	case glfw.KeyKP1:
		return ImGuiKey_Keypad1
	case glfw.KeyKP2:
		return ImGuiKey_Keypad2
	case glfw.KeyKP3:
		return ImGuiKey_Keypad3
	case glfw.KeyKP4:
		return ImGuiKey_Keypad4
	case glfw.KeyKP5:
		return ImGuiKey_Keypad5
	case glfw.KeyKP6:
		return ImGuiKey_Keypad6
	case glfw.KeyKP7:
		return ImGuiKey_Keypad7
	case glfw.KeyKP8:
		return ImGuiKey_Keypad8
	case glfw.KeyKP9:
		return ImGuiKey_Keypad9
	case glfw.KeyKPDecimal:
		return ImGuiKey_KeypadDecimal
	case glfw.KeyKPDivide:
		return ImGuiKey_KeypadDivide
	case glfw.KeyKPMultiply:
		return ImGuiKey_KeypadMultiply
	case glfw.KeyKPSubtract:
		return ImGuiKey_KeypadSubtract
	case glfw.KeyKPAdd:
		return ImGuiKey_KeypadAdd
	case glfw.KeyKPEnter:
		return ImGuiKey_KeypadEnter
	case glfw.KeyKPEqual:
		return ImGuiKey_KeypadEqual
	case glfw.KeyLeftShift:
		return ImGuiKey_LeftShift
	case glfw.KeyLeftControl:
		return ImGuiKey_LeftCtrl
	case glfw.KeyLeftAlt:
		return ImGuiKey_LeftAlt
	case glfw.KeyLeftSuper:
		return ImGuiKey_LeftSuper
	case glfw.KeyRightShift:
		return ImGuiKey_RightShift
	case glfw.KeyRightControl:
		return ImGuiKey_RightCtrl
	case glfw.KeyRightAlt:
		return ImGuiKey_RightAlt
	case glfw.KeyRightSuper:
		return ImGuiKey_RightSuper
	case glfw.KeyMenu:
		return ImGuiKey_Menu
	case glfw.Key0:
		return ImGuiKey_0
	case glfw.Key1:
		return ImGuiKey_1
	case glfw.Key2:
		return ImGuiKey_2
	case glfw.Key3:
		return ImGuiKey_3
	case glfw.Key4:
		return ImGuiKey_4
	case glfw.Key5:
		return ImGuiKey_5
	case glfw.Key6:
		return ImGuiKey_6
	case glfw.Key7:
		return ImGuiKey_7
	case glfw.Key8:
		return ImGuiKey_8
	case glfw.Key9:
		return ImGuiKey_9
	case glfw.KeyA:
		return ImGuiKey_A
	case glfw.KeyB:
		return ImGuiKey_B
	case glfw.KeyC:
		return ImGuiKey_C
	case glfw.KeyD:
		return ImGuiKey_D
	case glfw.KeyE:
		return ImGuiKey_E
	case glfw.KeyF:
		return ImGuiKey_F
	case glfw.KeyG:
		return ImGuiKey_G
	case glfw.KeyH:
		return ImGuiKey_H
	case glfw.KeyI:
		return ImGuiKey_I
	case glfw.KeyJ:
		return ImGuiKey_J
	case glfw.KeyK:
		return ImGuiKey_K
	case glfw.KeyL:
		return ImGuiKey_L
	case glfw.KeyM:
		return ImGuiKey_M
	case glfw.KeyN:
		return ImGuiKey_N
	case glfw.KeyO:
		return ImGuiKey_O
	case glfw.KeyP:
		return ImGuiKey_P
	case glfw.KeyQ:
		return ImGuiKey_Q
	case glfw.KeyR:
		return ImGuiKey_R
	case glfw.KeyS:
		return ImGuiKey_S
	case glfw.KeyT:
		return ImGuiKey_T
	case glfw.KeyU:
		return ImGuiKey_U
	case glfw.KeyV:
		return ImGuiKey_V
	case glfw.KeyW:
		return ImGuiKey_W
	case glfw.KeyX:
		return ImGuiKey_X
	case glfw.KeyY:
		return ImGuiKey_Y
	case glfw.KeyZ:
		return ImGuiKey_Z
	case glfw.KeyF1:
		return ImGuiKey_F1
	case glfw.KeyF2:
		return ImGuiKey_F2
	case glfw.KeyF3:
		return ImGuiKey_F3
	case glfw.KeyF4:
		return ImGuiKey_F4
	case glfw.KeyF5:
		return ImGuiKey_F5
	case glfw.KeyF6:
		return ImGuiKey_F6
	case glfw.KeyF7:
		return ImGuiKey_F7
	case glfw.KeyF8:
		return ImGuiKey_F8
	case glfw.KeyF9:
		return ImGuiKey_F9
	case glfw.KeyF10:
		return ImGuiKey_F10
	case glfw.KeyF11:
		return ImGuiKey_F11
	case glfw.KeyF12:
		return ImGuiKey_F12
	default:
		return ImGuiKey_None
	}
}

func (platform *GLFW) installCallbacks() {
	platform.window.SetMouseButtonCallback(platform.mouseButtonChange)
	platform.window.SetScrollCallback(platform.mouseScrollChange)
	platform.window.SetKeyCallback(platform.keyChange)
	platform.window.SetCharCallback(platform.charChange)
	platform.window.SetSizeCallback(platform.sizeChange)
	platform.window.SetDropCallback(platform.onDrop)
	platform.window.SetPosCallback(platform.posChange)
	platform.window.SetCloseCallback(platform.onClose)
	platform.window.SetFocusCallback(platform.onFocus)
}

var glfwButtonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: 0,
	glfw.MouseButton2: 1,
	glfw.MouseButton3: 2,
}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	0: glfw.MouseButton1,
	1: glfw.MouseButton2,
	2: glfw.MouseButton3,
}

func (platform *GLFW) onFocus(window *glfw.Window, focused bool) {
	//TODO: implement there
	// platform.imguiIO.AddFocusEvent(focused)
}

func (platform *GLFW) onClose(window *glfw.Window) {
	if platform.closeCallback != nil {
		shouldClose := platform.closeCallback()
		window.SetShouldClose(shouldClose)
	}
}

func (platform *GLFW) onDrop(window *glfw.Window, names []string) {
	window.Focus()

	if platform.dropCallback != nil {
		platform.dropCallback(names)
	}
}

func (platform *GLFW) posChange(window *glfw.Window, x, y int) {
	// Notfy pos changed and redraw.
	if platform.posChangeCallback != nil {
		platform.posChangeCallback(x, y)
	}

	platform.Update()
}

func (platform *GLFW) sizeChange(window *glfw.Window, width, height int) {
	// Notify size changed and redraw.
	if platform.sizeChangeCallback != nil {
		platform.sizeChangeCallback(width, height)
	}

	platform.Update()
}

func (platform *GLFW) mouseButtonChange(window *glfw.Window, rawButton glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	buttonIndex, known := glfwButtonIndexByID[rawButton]

	if known && (action == glfw.Press) {
		platform.mouseJustPressed[buttonIndex] = true
	}

	platform.Update()
}

func (platform *GLFW) mouseScrollChange(window *glfw.Window, x, y float64) {
	platform.imguiIO.AddMouseWheelDelta(float32(x), float32(y))
	platform.Update()
}

func (platform *GLFW) updateKeyModifiers(mods glfw.ModifierKey) {
	platform.imguiIO.AddKeyEvent(ImGuiKey_ModCtrl, (mods&glfw.ModControl) != 0)
	platform.imguiIO.AddKeyEvent(ImGuiKey_ModShift, (mods&glfw.ModShift) != 0)
	platform.imguiIO.AddKeyEvent(ImGuiKey_ModAlt, (mods&glfw.ModAlt) != 0)
	platform.imguiIO.AddKeyEvent(ImGuiKey_ModSuper, (mods&glfw.ModSuper) != 0)
}

func (platform *GLFW) keyChange(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	platform.Update()

	if action != glfw.Press && action != glfw.Release {
		return
	}

	if mkey := platform.keyToModifier(key); mkey != 0 {
		if action == glfw.Press {
			mods |= mkey
		} else {
			mods &= mkey
		}
	}

	platform.updateKeyModifiers(mods)

	imguiKey := platform.keyToImGuiKey(key)
	platform.imguiIO.AddKeyEvent(imguiKey, action == glfw.Press)

	if platform.inputCallback != nil {
		platform.inputCallback(key, mods, action)
	}

	platform.Update()
}

func (platform *GLFW) charChange(window *glfw.Window, char rune) {
	platform.Update()
	platform.imguiIO.AddInputCharacters(string(char))
	platform.Update()
}

func (platform *GLFW) GetClipboard() string {
	return platform.window.GetClipboardString()
}

func (platform *GLFW) SetClipboard(content string) {
	platform.window.SetClipboardString(content)
}

func (platform *GLFW) GetTPS() int {
	return platform.tps
}

func (platform *GLFW) SetTPS(tps int) {
	platform.tps = tps
}

// SetIcon sets the icon of the specified window. If passed an array of candidate images,
// those of or closest to the sizes desired by the system are selected. If no images are
// specified, the window reverts to its default icon.
//
// The image is ideally provided in the form of *image.NRGBA.
// The pixels are 32-bit, little-endian, non-premultiplied RGBA, i.e. eight
// bits per channel with the red channel first. They are arranged canonically
// as packed sequential rows, starting from the top-left corner. If the image
// type is not *image.NRGBA, it will be converted to it.
//
// The desired image sizes varies depending on platform and system settings. The selected
// images will be rescaled as needed. Good sizes include 16x16, 32x32 and 48x48.
func (platform *GLFW) SetIcon(icons []image.Image) {
	platform.window.SetIcon(icons)
}

// SetSizeLimits sets the size limits of the client area of the specified window.
// If the window is full screen or not resizable, this function does nothing.
//
// The size limits are applied immediately and may cause the window to be resized.
// To specify only a minimum size or only a maximum one, set the other pair to GLFW_DONT_CARE.
// To disable size limits for a window, set them all to GLFW_DONT_CARE.
func (platform *GLFW) SetSizeLimits(minw, minh, maxw, maxh int) {
	platform.window.SetSizeLimits(minw, minh, maxw, maxh)
}

// SetTitle sets the title of window.
func (platform *GLFW) SetTitle(title string) {
	platform.window.SetTitle(title)
}

// IsMinimized checks whether window is minimized.
func (platform *GLFW) IsMinimized() bool {
	return glfw.True == platform.window.GetAttrib(glfw.Iconified)
}

// IsVisible checks whether window is visible.
func (platform *GLFW) IsVisible() bool {
	return glfw.True == platform.window.GetAttrib(glfw.Visible)
}
