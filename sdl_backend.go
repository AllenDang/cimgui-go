//go:build !exclude_cimgui_sdl && !darwin

package imgui

// #cgo amd64,linux LDFLAGS: ${SRCDIR}/lib/linux/x64/libSDL2.a -ldl -lGL -lX11
// #cgo amd64,windows LDFLAGS: -L${SRCDIR}/lib/windows/x64 -l:libSDL2.a -lgdi32 -lopengl32 -limm32         -mwindows  -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lsetupapi -lversion -luuid
// #cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo
// #cgo amd64,darwin LDFLAGS: ${SRCDIR}/lib/macos/x64/libSDL2.a
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/lib/macos/arm64/libSDL2.a
// #cgo !gles2,darwin LDFLAGS: -framework OpenGL
// #cgo gles2,darwin LDFLAGS: -lGLESv2
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_SDL2
// #include <stdlib.h>
// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestoryContext();
// #include <stdint.h>
// #include "sdl_backend.h"
import "C"

import (
	"errors"
	"image"
	"image/draw"
	"unsafe"
)

type SDLWindowFlags int

// SDL bool values
const (
	SDLTrue  = SDLWindowFlags(1)
	SDLFalse = SDLWindowFlags(0)
)

// Window flags
const (
	SDLWindowFlagsNone              = SDLWindowFlags(0) // Clear all flags
	SDLWindowFlagsFullScreen        = SDLWindowFlags(C.SDL_WINDOW_FULLSCREEN)
	SDLWindowFlagsOpengl            = SDLWindowFlags(C.SDL_WINDOW_OPENGL)
	SDLWindowFlagsDecorated         = SDLWindowFlags(C.SDL_WINDOW_SHOWN)
	SDLWindowFlagsTransparent       = SDLWindowFlags(C.SDL_WINDOW_HIDDEN)
	SDLWindowFlagsVisible           = SDLWindowFlags(C.SDL_WINDOW_BORDERLESS)
	SDLWindowFlagsResizable         = SDLWindowFlags(C.SDL_WINDOW_RESIZABLE)
	SDLWindowFlagsMinimized         = SDLWindowFlags(C.SDL_WINDOW_MINIMIZED)
	SDLWindowFlagsMaximized         = SDLWindowFlags(C.SDL_WINDOW_MAXIMIZED)
	SDLWindowFlagsMouseGrabbed      = SDLWindowFlags(C.SDL_WINDOW_MOUSE_GRABBED)
	SDLWindowFlagsInputFocus        = SDLWindowFlags(C.SDL_WINDOW_INPUT_FOCUS)
	SDLWindowFlagsMouseFocus        = SDLWindowFlags(C.SDL_WINDOW_MOUSE_FOCUS)
	SDLWindowFlagsFullscreenDesktop = SDLWindowFlags(C.SDL_WINDOW_FULLSCREEN_DESKTOP)
	SDLWindowFlagsWindowForeign     = SDLWindowFlags(C.SDL_WINDOW_FOREIGN)
	SDLWindowFlagsAllowHighDPI      = SDLWindowFlags(C.SDL_WINDOW_ALLOW_HIGHDPI)
	SDLWindowFlagsMouseCapture      = SDLWindowFlags(C.SDL_WINDOW_MOUSE_CAPTURE)
	SDLWindowFlagsAlwaysOnTop       = SDLWindowFlags(C.SDL_WINDOW_ALWAYS_ON_TOP)
	SDLWindowFlagsSkipTaskbar       = SDLWindowFlags(C.SDL_WINDOW_SKIP_TASKBAR)
	SDLWindowFlagsUtility           = SDLWindowFlags(C.SDL_WINDOW_UTILITY)
	SDLWindowFlagsTooltip           = SDLWindowFlags(C.SDL_WINDOW_TOOLTIP)
	SDLWindowFlagsPopupMenu         = SDLWindowFlags(C.SDL_WINDOW_POPUP_MENU)
	SDLWindowFlagsKeyboardGrabbed   = SDLWindowFlags(C.SDL_WINDOW_KEYBOARD_GRABBED)
	SDLWindowFlagsWindowVulkan      = SDLWindowFlags(C.SDL_WINDOW_VULKAN)
	SDLWindowFlagsWindowMetal       = SDLWindowFlags(C.SDL_WINDOW_METAL)
)

// Swap interval flags
const (
	SDLSwapIntervalImmediate    = SDLWindowFlags(0)
	SDLSwapIntervalVsync        = SDLWindowFlags(1)
	SDLSwapIntervalAdaptiveSync = SDLWindowFlags(-1)
)

// Input mode flags
const (
	SDLInputModeRelativeMouse = SDLWindowFlags(iota)
	SDLInputModeGrab
)

/*
type GLFWKey int

const (
	GLFWKeySpace        = GLFWKey(C.GLFWKeySpace)
	GLFWKeyApostrophe   = GLFWKey(C.GLFWKeyApostrophe)
	GLFWKeyComma        = GLFWKey(C.GLFWKeyComma)
	GLFWKeyMinus        = GLFWKey(C.GLFWKeyMinus)
	GLFWKeyPeriod       = GLFWKey(C.GLFWKeyPeriod)
	GLFWKeySlash        = GLFWKey(C.GLFWKeySlash)
	GLFWKey0            = GLFWKey(C.GLFWKey0)
	GLFWKey1            = GLFWKey(C.GLFWKey1)
	GLFWKey2            = GLFWKey(C.GLFWKey2)
	GLFWKey3            = GLFWKey(C.GLFWKey3)
	GLFWKey4            = GLFWKey(C.GLFWKey4)
	GLFWKey5            = GLFWKey(C.GLFWKey5)
	GLFWKey6            = GLFWKey(C.GLFWKey6)
	GLFWKey7            = GLFWKey(C.GLFWKey7)
	GLFWKey8            = GLFWKey(C.GLFWKey8)
	GLFWKey9            = GLFWKey(C.GLFWKey9)
	GLFWKeySemicolon    = GLFWKey(C.GLFWKeySemicolon)
	GLFWKeyEqual        = GLFWKey(C.GLFWKeyEqual)
	GLFWKeyA            = GLFWKey(C.GLFWKeyA)
	GLFWKeyB            = GLFWKey(C.GLFWKeyB)
	GLFWKeyC            = GLFWKey(C.GLFWKeyC)
	GLFWKeyD            = GLFWKey(C.GLFWKeyD)
	GLFWKeyE            = GLFWKey(C.GLFWKeyE)
	GLFWKeyF            = GLFWKey(C.GLFWKeyF)
	GLFWKeyG            = GLFWKey(C.GLFWKeyG)
	GLFWKeyH            = GLFWKey(C.GLFWKeyH)
	GLFWKeyI            = GLFWKey(C.GLFWKeyI)
	GLFWKeyJ            = GLFWKey(C.GLFWKeyJ)
	GLFWKeyK            = GLFWKey(C.GLFWKeyK)
	GLFWKeyL            = GLFWKey(C.GLFWKeyL)
	GLFWKeyM            = GLFWKey(C.GLFWKeyM)
	GLFWKeyN            = GLFWKey(C.GLFWKeyN)
	GLFWKeyO            = GLFWKey(C.GLFWKeyO)
	GLFWKeyP            = GLFWKey(C.GLFWKeyP)
	GLFWKeyQ            = GLFWKey(C.GLFWKeyQ)
	GLFWKeyR            = GLFWKey(C.GLFWKeyR)
	GLFWKeyS            = GLFWKey(C.GLFWKeyS)
	GLFWKeyT            = GLFWKey(C.GLFWKeyT)
	GLFWKeyU            = GLFWKey(C.GLFWKeyU)
	GLFWKeyV            = GLFWKey(C.GLFWKeyV)
	GLFWKeyW            = GLFWKey(C.GLFWKeyW)
	GLFWKeyX            = GLFWKey(C.GLFWKeyX)
	GLFWKeyY            = GLFWKey(C.GLFWKeyY)
	GLFWKeyZ            = GLFWKey(C.GLFWKeyZ)
	GLFWKeyLeftBracket  = GLFWKey(C.GLFWKeyLeftBracket)
	GLFWKeyBackslash    = GLFWKey(C.GLFWKeyBackslash)
	GLFWKeyRightBracket = GLFWKey(C.GLFWKeyRightBracket)
	GLFWKeyGraveAccent  = GLFWKey(C.GLFWKeyGraveAccent)
	GLFWKeyWorld1       = GLFWKey(C.GLFWKeyWorld1)
	GLFWKeyWorld2       = GLFWKey(C.GLFWKeyWorld2)

	// Function keys
	GLFWKeyEscape       = GLFWKey(C.GLFWKeyEscape)
	GLFWKeyEnter        = GLFWKey(C.GLFWKeyEnter)
	GLFWKeyTab          = GLFWKey(C.GLFWKeyTab)
	GLFWKeyBackspace    = GLFWKey(C.GLFWKeyBackspace)
	GLFWKeyInsert       = GLFWKey(C.GLFWKeyInsert)
	GLFWKeyDelete       = GLFWKey(C.GLFWKeyDelete)
	GLFWKeyRight        = GLFWKey(C.GLFWKeyRight)
	GLFWKeyLeft         = GLFWKey(C.GLFWKeyLeft)
	GLFWKeyDown         = GLFWKey(C.GLFWKeyDown)
	GLFWKeyUp           = GLFWKey(C.GLFWKeyUp)
	GLFWKeyPageUp       = GLFWKey(C.GLFWKeyPageUp)
	GLFWKeyPageDown     = GLFWKey(C.GLFWKeyPageDown)
	GLFWKeyHome         = GLFWKey(C.GLFWKeyHome)
	GLFWKeyEnd          = GLFWKey(C.GLFWKeyEnd)
	GLFWKeyCapsLock     = GLFWKey(C.GLFWKeyCapsLock)
	GLFWKeyScrollLock   = GLFWKey(C.GLFWKeyScrollLock)
	GLFWKeyNumLock      = GLFWKey(C.GLFWKeyNumLock)
	GLFWKeyPrintScreen  = GLFWKey(C.GLFWKeyPrintScreen)
	GLFWKeyPause        = GLFWKey(C.GLFWKeyPause)
	GLFWKeyF1           = GLFWKey(C.GLFWKeyF1)
	GLFWKeyF2           = GLFWKey(C.GLFWKeyF2)
	GLFWKeyF3           = GLFWKey(C.GLFWKeyF3)
	GLFWKeyF4           = GLFWKey(C.GLFWKeyF4)
	GLFWKeyF5           = GLFWKey(C.GLFWKeyF5)
	GLFWKeyF6           = GLFWKey(C.GLFWKeyF6)
	GLFWKeyF7           = GLFWKey(C.GLFWKeyF7)
	GLFWKeyF8           = GLFWKey(C.GLFWKeyF8)
	GLFWKeyF9           = GLFWKey(C.GLFWKeyF9)
	GLFWKeyF10          = GLFWKey(C.GLFWKeyF10)
	GLFWKeyF11          = GLFWKey(C.GLFWKeyF11)
	GLFWKeyF12          = GLFWKey(C.GLFWKeyF12)
	GLFWKeyF13          = GLFWKey(C.GLFWKeyF13)
	GLFWKeyF14          = GLFWKey(C.GLFWKeyF14)
	GLFWKeyF15          = GLFWKey(C.GLFWKeyF15)
	GLFWKeyF16          = GLFWKey(C.GLFWKeyF16)
	GLFWKeyF17          = GLFWKey(C.GLFWKeyF17)
	GLFWKeyF18          = GLFWKey(C.GLFWKeyF18)
	GLFWKeyF19          = GLFWKey(C.GLFWKeyF19)
	GLFWKeyF20          = GLFWKey(C.GLFWKeyF20)
	GLFWKeyF21          = GLFWKey(C.GLFWKeyF21)
	GLFWKeyF22          = GLFWKey(C.GLFWKeyF22)
	GLFWKeyF23          = GLFWKey(C.GLFWKeyF23)
	GLFWKeyF24          = GLFWKey(C.GLFWKeyF24)
	GLFWKeyF25          = GLFWKey(C.GLFWKeyF25)
	GLFWKeyKp0          = GLFWKey(C.GLFWKeyKp0)
	GLFWKeyKp1          = GLFWKey(C.GLFWKeyKp1)
	GLFWKeyKp2          = GLFWKey(C.GLFWKeyKp2)
	GLFWKeyKp3          = GLFWKey(C.GLFWKeyKp3)
	GLFWKeyKp4          = GLFWKey(C.GLFWKeyKp4)
	GLFWKeyKp5          = GLFWKey(C.GLFWKeyKp5)
	GLFWKeyKp6          = GLFWKey(C.GLFWKeyKp6)
	GLFWKeyKp7          = GLFWKey(C.GLFWKeyKp7)
	GLFWKeyKp8          = GLFWKey(C.GLFWKeyKp8)
	GLFWKeyKp9          = GLFWKey(C.GLFWKeyKp9)
	GLFWKeyKpDecimal    = GLFWKey(C.GLFWKeyKpDecimal)
	GLFWKeyKpDivide     = GLFWKey(C.GLFWKeyKpDivide)
	GLFWKeyKpMultiply   = GLFWKey(C.GLFWKeyKpMultiply)
	GLFWKeyKpSubtract   = GLFWKey(C.GLFWKeyKpSubtract)
	GLFWKeyKpAdd        = GLFWKey(C.GLFWKeyKpAdd)
	GLFWKeyKpEnter      = GLFWKey(C.GLFWKeyKpEnter)
	GLFWKeyKpEqual      = GLFWKey(C.GLFWKeyKpEqual)
	GLFWKeyLeftShift    = GLFWKey(C.GLFWKeyLeftShift)
	GLFWKeyLeftControl  = GLFWKey(C.GLFWKeyLeftControl)
	GLFWKeyLeftAlt      = GLFWKey(C.GLFWKeyLeftAlt)
	GLFWKeyLeftSuper    = GLFWKey(C.GLFWKeyLeftSuper)
	GLFWKeyRightShift   = GLFWKey(C.GLFWKeyRightShift)
	GLFWKeyRightControl = GLFWKey(C.GLFWKeyRightControl)
	GLFWKeyRightAlt     = GLFWKey(C.GLFWKeyRightAlt)
	GLFWKeyRightSuper   = GLFWKey(C.GLFWKeyRightSuper)
	GLFWKeyMenu         = GLFWKey(C.GLFWKeyMenu)
)

type GLFWModifierKey int

const (
	GLFWModShift    = GLFWModifierKey(C.GLFWModShift)
	GLFWModControl  = GLFWModifierKey(C.GLFWModControl)
	GLFWModAlt      = GLFWModifierKey(C.GLFWModAlt)
	GLFWModSuper    = GLFWModifierKey(C.GLFWModSuper)
	GLFWModCapsLock = GLFWModifierKey(C.GLFWModCapsLock)
	GLFWModNumLock  = GLFWModifierKey(C.GLFWModNumLock)
)
*/

var _ Backend[SDLWindowFlags] = &SDLBackend{}

type SDLBackend struct {
	afterCreateContext   voidCallbackFunc
	loop                 voidCallbackFunc
	beforeRender         voidCallbackFunc
	afterRender          voidCallbackFunc
	beforeDestoryContext voidCallbackFunc
	dropCB               DropCallback
	closeCB              func(pointer unsafe.Pointer)
	keyCb                KeyCallback
	sizeCb               SizeChangeCallback
	window               uintptr
}

func SDLGetError() string {
	return C.GoString(C.SDL_GetError())
}

func NewSDLBackend() *SDLBackend {
	b := &SDLBackend{}
	// \returns 0 on success or a negative error code on failure; call
	// SDL_GetError() for more information.
	if C.igInitSDL() != 0 {
		panic(SDLGetError())
	}

	return b
}

func (b *SDLBackend) handle() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(b.window))
}

func (b *SDLBackend) SetAfterCreateContextHook(hook func()) {
	b.afterCreateContext = hook
}

func (b *SDLBackend) afterCreateHook() func() {
	return b.afterCreateContext
}

func (b *SDLBackend) SetBeforeDestroyContextHook(hook func()) {
	b.beforeDestoryContext = hook
}

func (b *SDLBackend) beforeDestroyHook() func() {
	return b.beforeDestoryContext
}

func (b *SDLBackend) SetBeforeRenderHook(hook func()) {
	b.beforeRender = hook
}

func (b *SDLBackend) beforeRenderHook() func() {
	return b.beforeRender
}

func (b *SDLBackend) SetAfterRenderHook(hook func()) {
	b.afterRender = hook
}

func (b *SDLBackend) afterRenderHook() func() {
	return b.afterRender
}

func (b *SDLBackend) SetBgColor(color Vec4) {
	C.igSetBgColor(color.toC())
}

func (b *SDLBackend) Run(loop func()) {
	b.loop = loop
	C.igSDLRunLoop(b.handle(), C.VoidCallback(C.loopCallback), C.VoidCallback(C.beforeRender), C.VoidCallback(C.afterRender), C.VoidCallback(C.beforeDestoryContext))
}

func (b *SDLBackend) loopFunc() func() {
	return b.loop
}

func (b *SDLBackend) dropCallback() DropCallback {
	return b.dropCB
}

func (b *SDLBackend) closeCallback() func(wnd unsafe.Pointer) {
	return b.closeCB
}

func (b *SDLBackend) SetWindowPos(x, y int) {
	C.igSDLWindow_SetWindowPos(b.handle(), C.int(x), C.int(y))
}

func (b *SDLBackend) GetWindowPos() (x, y int32) {
	xArg, xFin := WrapNumberPtr[C.int, int32](&x)
	defer xFin()

	yArg, yFin := WrapNumberPtr[C.int, int32](&y)
	defer yFin()

	C.igSDLWindow_GetWindowPos(b.handle(), xArg, yArg)

	return
}

func (b *SDLBackend) SetWindowSize(width, height int) {
	C.igSDLWindow_SetSize(b.handle(), C.int(width), C.int(height))
}

func (b *SDLBackend) DisplaySize() (width int32, height int32) {
	widthArg, widthFin := WrapNumberPtr[C.int, int32](&width)
	defer widthFin()

	heightArg, heightFin := WrapNumberPtr[C.int, int32](&height)
	defer heightFin()

	C.igSDLWindow_GetDisplaySize(b.handle(), widthArg, heightArg)

	return
}

func (b *SDLBackend) ContentScale() (width, height float32) {
	widthArg, widthFin := WrapNumberPtr[C.float, float32](&width)
	defer widthFin()

	heightArg, heightFin := WrapNumberPtr[C.float, float32](&height)
	defer heightFin()

	C.igSDLWindow_GetContentScale(b.handle(), widthArg, heightArg)

	return
}

func (b *SDLBackend) SetWindowTitle(title string) {
	titleArg, titleFin := WrapString(title)
	defer titleFin()

	C.igSDLWindow_SetTitle(b.handle(), titleArg)
}

// The minimum and maximum size of the content area of a windowed mode window.
// To specify only a minimum size or only a maximum one, set the other pair to -1
// e.g. SetWindowSizeLimits(640, 480, -1, -1)
func (b *SDLBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	C.igSDLWindow_SetSizeLimits(b.handle(), C.int(minWidth), C.int(minHeight), C.int(maxWidth), C.int(maxHeight))
}

func (b *SDLBackend) SetShouldClose(value bool) {
	// TODO: not implemented
	// C.igSDLWindow_SetShouldClose(b.handle(), C.int(CastBool(value)))
}

func (b *SDLBackend) CreateWindow(title string, width, height int) {
	titleArg, titleFin := WrapString(title)
	defer titleFin()

	b.window = uintptr(unsafe.Pointer(C.igCreateSDLWindow(
		titleArg,
		C.int(width),
		C.int(height),
		C.VoidCallback(C.afterCreateContext),
	)))
	if b.window == 0 {
		panic("Failed to create SDL window")
	}
}

func (b *SDLBackend) SetTargetFPS(fps uint) {
	C.igSetTargetFPS(C.uint(fps))
}

func (b *SDLBackend) Refresh() {
	C.igRefresh()
}

func (b *SDLBackend) CreateTexture(pixels unsafe.Pointer, width, height int) TextureID {
	tex := C.igCreateTexture((*C.uchar)(pixels), C.int(width), C.int(height))
	return *newTextureIDFromC(&tex)
}

func (b *SDLBackend) CreateTextureRgba(img *image.RGBA, width, height int) TextureID {
	tex := C.igCreateTexture((*C.uchar)(&(img.Pix[0])), C.int(width), C.int(height))
	return *newTextureIDFromC(&tex)
}

func (b *SDLBackend) DeleteTexture(id TextureID) {
	C.igDeleteTexture(C.ImTextureID(id.Data))
}

// SetDropCallback sets the drop callback which is called when an object
// is dropped over the window.
func (b *SDLBackend) SetDropCallback(cbfun DropCallback) {
	b.dropCB = cbfun
	// TODO: not implemented
	// C.igSDLWindow_SetDropCallbackCB(b.handle())
}

func (b *SDLBackend) SetCloseCallback(cbfun WindowCloseCallback[SDLWindowFlags]) {
	b.closeCB = func(_ unsafe.Pointer) {
		cbfun(b)
	}

	// TODO: not implemented
	// C.igSDLWindow_SetCloseCallback(b.handle())
}

// SetWindowHint applies to next CreateWindow call
// so use it before CreateWindow call ;-)
// default applied flags: SDLWindowFlagsOpengl | SDLWindowFlagsResizable | SDLWindowFlagsAllowHighDPI
// set flag if value is 1, clear flag if value is 0
func (b *SDLBackend) SetWindowFlags(flag SDLWindowFlags, value int) {
	C.igSDLWindowHint(C.SDL_WindowFlags(flag), C.int(value))
}

// SetIcons sets icons for the window.
// THIS CODE COMES FROM https://github.com/go-gl/glfw (BSD-3 clause) - Copyright (c) 2012 The glfw3-go Authors. All rights reserved.
func (b *SDLBackend) SetIcons(images ...image.Image) {
	count := len(images)
	cimages := make([]C.CImage, count)
	freePixels := make([]func(), count)

	for i, img := range images {
		var pixels []uint8
		b := img.Bounds()

		switch img := img.(type) {
		case *image.NRGBA:
			pixels = img.Pix
		default:
			m := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
			draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
			pixels = m.Pix
		}

		pix, free := func(origin []byte) (pointer *uint8, free func()) {
			n := len(origin)
			if n == 0 {
				return nil, func() {}
			}

			ptr := C.CBytes(origin)
			return (*uint8)(ptr), func() { C.free(ptr) }
		}(pixels)

		freePixels[i] = free

		cimages[i].width = C.int(b.Dx())
		cimages[i].height = C.int(b.Dy())
		cimages[i].pixels = (*C.uchar)(pix)
	}

	var p *C.CImage
	if count > 0 {
		p = &cimages[0]
	}

	_ = p
	// TODO: not implemented
	// C.igSDLWindow_SetIcon(b.handle(), C.int(count), p)

	for _, v := range freePixels {
		v()
	}
}

func (b *SDLBackend) SetKeyCallback(cbfun KeyCallback) {
	b.keyCb = cbfun
	// TODO: not implemented
	// C.igSDLWindow_SetKeyCallback(b.handle())
}

func (b *SDLBackend) keyCallback() KeyCallback {
	return b.keyCb
}

func (b *SDLBackend) SetSizeChangeCallback(cbfun SizeChangeCallback) {
	b.sizeCb = cbfun
	// TODO: notttt pimlemented
	// C.igSDLWindow_SetSizeCallback(b.handle())
}

func (b *SDLBackend) sizeCallback() SizeChangeCallback {
	return b.sizeCb
}

func (b *SDLBackend) SetSwapInterval(interval SDLWindowFlags) error {
	if C.SDL_GL_SetSwapInterval(C.int(interval)) != 0 {
		return errors.New(SDLGetError())
	}
	return nil
}

func (b *SDLBackend) SetCursorPos(x, y float64) {
	C.SDL_WarpMouseInWindow(b.handle(), C.int(x), C.int(y))
}

func (b *SDLBackend) SetInputMode(mode SDLWindowFlags, value SDLWindowFlags) {
	switch mode {
	case SDLInputModeRelativeMouse:
		C.SDL_SetRelativeMouseMode(C.SDL_bool(value))
	case SDLInputModeGrab:
		C.SDL_SetWindowGrab(b.handle(), C.SDL_bool(value))
	}
}
