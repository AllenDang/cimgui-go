//go:build !exclude_cimgui_glfw
// +build !exclude_cimgui_glfw

package imgui

// #cgo amd64,linux LDFLAGS: ${SRCDIR}/lib/linux/x64/libglfw3.a -ldl -lGL -lX11
// #cgo amd64,windows LDFLAGS: -L${SRCDIR}/lib/windows/x64 -l:libglfw3.a -lgdi32 -lopengl32 -limm32
// #cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo
// #cgo amd64,darwin LDFLAGS: ${SRCDIR}/lib/macos/x64/libglfw3.a
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/lib/macos/arm64/libglfw3.a
// #cgo !gles2,darwin LDFLAGS: -framework OpenGL
// #cgo gles2,darwin LDFLAGS: -lGLESv2
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_GLFW
// #include <stdlib.h>
// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestoryContext();
// #include <stdint.h>
// #include "glfw_backend.h"
import "C"

import (
	"image"
	"image/draw"
	"unsafe"
)

type GLFWWindowFlags int

// Window flags
const (
	GLFWWindowFlagsNone        = GLFWWindowFlags(C.GLFWWindowNone)
	GLFWWindowFlagsResizable   = GLFWWindowFlags(C.GLFWWindowResizable)
	GLFWWindowFlagsMaximized   = GLFWWindowFlags(C.GLFWWindowMaximized)
	GLFWWindowFlagsDecorated   = GLFWWindowFlags(C.GLFWWindowDecorated)
	GLFWWindowFlagsTransparent = GLFWWindowFlags(C.GLFWWindowTransparentFramebuffer)
	GLFWWindowFlagsVisible     = GLFWWindowFlags(C.GLFWWindowVisible)
	GLFWWindowFlagsFloating    = GLFWWindowFlags(C.GLFWWindowFloating)
	GLFWWindowFlagsFocused     = GLFWWindowFlags(C.GLFWWindowFocused)
	GLFWWindowFlagsIconified   = GLFWWindowFlags(C.GLFWWindowIconified)
	GLFWWindowFlagsAutoIconify = GLFWWindowFlags(C.GLFWWindowAutoIconify)
)

// SwapInterval values
const (
	GLFWSwapIntervalImmediate = GLFWWindowFlags(0)
	GLFWSwapIntervalVsync     = GLFWWindowFlags(1)
)

// Input modes
const (
	GLFWInputModeCursor         = GLFWWindowFlags(C.GLFW_CURSOR)
	GLFWInputModeCursorNormal   = GLFWWindowFlags(C.GLFW_CURSOR_NORMAL)
	GLFWInputModeCursorHidden   = GLFWWindowFlags(C.GLFW_CURSOR_HIDDEN)
	GLFWInputModeCursorDisabled = GLFWWindowFlags(C.GLFW_CURSOR_DISABLED)
	GLFWInputModeRawMouseMotion = GLFWWindowFlags(C.GLFW_RAW_MOUSE_MOTION)
)

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

	/* Function keys */
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

var _ Backend[GLFWWindowFlags] = &GLFWBackend{}

type GLFWBackend struct {
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

func NewGLFWBackend() *GLFWBackend {
	b := &GLFWBackend{}
	if C.igInitGLFW() == 0 {
		panic("Failed to initialize GLFW")
	}

	return b
}

func (b *GLFWBackend) handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(unsafe.Pointer(b.window))
}

func (b *GLFWBackend) SetAfterCreateContextHook(hook func()) {
	b.afterCreateContext = hook
}

func (b *GLFWBackend) afterCreateHook() func() {
	return b.afterCreateContext
}

func (b *GLFWBackend) SetBeforeDestroyContextHook(hook func()) {
	b.beforeDestoryContext = hook
}

func (b *GLFWBackend) beforeDestroyHook() func() {
	return b.beforeDestoryContext
}

func (b *GLFWBackend) SetBeforeRenderHook(hook func()) {
	b.beforeRender = hook
}

func (b *GLFWBackend) beforeRenderHook() func() {
	return b.beforeRender
}

func (b *GLFWBackend) SetAfterRenderHook(hook func()) {
	b.afterRender = hook
}

func (b *GLFWBackend) afterRenderHook() func() {
	return b.afterRender
}

func (b *GLFWBackend) SetBgColor(color Vec4) {
	C.igSetBgColor(color.toC())
}

func (b *GLFWBackend) Run(loop func()) {
	b.loop = loop
	C.igGLFWRunLoop(b.handle(), C.VoidCallback(C.loopCallback), C.VoidCallback(C.beforeRender), C.VoidCallback(C.afterRender), C.VoidCallback(C.beforeDestoryContext))
}

func (b *GLFWBackend) loopFunc() func() {
	return b.loop
}

func (b *GLFWBackend) dropCallback() DropCallback {
	return b.dropCB
}

func (b *GLFWBackend) closeCallback() func(wnd unsafe.Pointer) {
	return b.closeCB
}

func (b *GLFWBackend) SetWindowPos(x, y int) {
	C.igGLFWWindow_SetWindowPos(b.handle(), C.int(x), C.int(y))
}

func (b *GLFWBackend) GetWindowPos() (x, y int32) {
	xArg, xFin := WrapNumberPtr[C.int, int32](&x)
	defer xFin()

	yArg, yFin := WrapNumberPtr[C.int, int32](&y)
	defer yFin()

	C.igGLFWWindow_GetWindowPos(b.handle(), xArg, yArg)

	return
}

func (b *GLFWBackend) SetWindowSize(width, height int) {
	C.igGLFWWindow_SetSize(b.handle(), C.int(width), C.int(height))
}

func (b *GLFWBackend) DisplaySize() (width int32, height int32) {
	widthArg, widthFin := WrapNumberPtr[C.int, int32](&width)
	defer widthFin()

	heightArg, heightFin := WrapNumberPtr[C.int, int32](&height)
	defer heightFin()

	C.igGLFWWindow_GetDisplaySize(b.handle(), widthArg, heightArg)

	return
}

func (b *GLFWBackend) ContentScale() (width, height float32) {
	widthArg, widthFin := WrapNumberPtr[C.float, float32](&width)
	defer widthFin()

	heightArg, heightFin := WrapNumberPtr[C.float, float32](&height)
	defer heightFin()

	C.igGLFWWindow_GetContentScale(b.handle(), widthArg, heightArg)

	return
}

func (b *GLFWBackend) SetWindowTitle(title string) {
	titleArg, titleFin := WrapString(title)
	defer titleFin()

	C.igGLFWWindow_SetTitle(b.handle(), titleArg)
}

// The minimum and maximum size of the content area of a windowed mode window.
// To specify only a minimum size or only a maximum one, set the other pair to -1
// e.g. SetWindowSizeLimits(640, 480, -1, -1)
func (b *GLFWBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	C.igGLFWWindow_SetSizeLimits(b.handle(), C.int(minWidth), C.int(minHeight), C.int(maxWidth), C.int(maxHeight))
}

func (b *GLFWBackend) SetShouldClose(value bool) {
	C.igGLFWWindow_SetShouldClose(b.handle(), C.int(CastBool(value)))
}

func (b *GLFWBackend) CreateWindow(title string, width, height int) {
	titleArg, titleFin := WrapString(title)
	defer titleFin()

	b.window = uintptr(unsafe.Pointer(C.igCreateGLFWWindow(
		titleArg,
		C.int(width),
		C.int(height),
		C.VoidCallback(C.afterCreateContext),
	)))
	if b.window == 0 {
		panic("Failed to create GLFW window")
	}
}

func (b *GLFWBackend) SetTargetFPS(fps uint) {
	C.igSetTargetFPS(C.uint(fps))
}

func (b *GLFWBackend) Refresh() {
	C.igRefresh()
}

func (b *GLFWBackend) CreateTexture(pixels unsafe.Pointer, width, height int) TextureID {
	tex := C.igCreateTexture((*C.uchar)(pixels), C.int(width), C.int(height))
	return *newTextureIDFromC(&tex)
}

func (b *GLFWBackend) CreateTextureRgba(img *image.RGBA, width, height int) TextureID {
	tex := C.igCreateTexture((*C.uchar)(&(img.Pix[0])), C.int(width), C.int(height))
	return *newTextureIDFromC(&tex)
}

func (b *GLFWBackend) DeleteTexture(id TextureID) {
	C.igDeleteTexture(C.ImTextureID(id.Data))
}

// SetDropCallback sets the drop callback which is called when an object
// is dropped over the window.
func (b *GLFWBackend) SetDropCallback(cbfun DropCallback) {
	b.dropCB = cbfun
	C.igGLFWWindow_SetDropCallbackCB(b.handle())
}

func (b *GLFWBackend) SetCloseCallback(cbfun WindowCloseCallback[GLFWWindowFlags]) {
	b.closeCB = func(_ unsafe.Pointer) {
		cbfun(b)
	}

	C.igGLFWWindow_SetCloseCallback(b.handle())
}

// SetWindowHint applies to next CreateWindow call
// so use it before CreateWindow call ;-)
func (b *GLFWBackend) SetWindowFlags(flag GLFWWindowFlags, value int) {
	C.igWindowHint(C.GLFWWindowFlags(flag), C.int(value))
}

// SetIcons sets icons for the window.
// THIS CODE COMES FROM https://github.com/go-gl/glfw (BSD-3 clause) - Copyright (c) 2012 The glfw3-go Authors. All rights reserved.
func (b *GLFWBackend) SetIcons(images ...image.Image) {
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
	C.igGLFWWindow_SetIcon(b.handle(), C.int(count), p)

	for _, v := range freePixels {
		v()
	}
}

func (b *GLFWBackend) SetKeyCallback(cbfun KeyCallback) {
	b.keyCb = func(k, s, a, m int) {
		C.iggImplGlfw_KeyCallback(b.handle(), C.int(k), C.int(s), C.int(a), C.int(m))
		cbfun(k, s, a, m)
	}
	C.igGLFWWindow_SetKeyCallback(b.handle())
}

func (b *GLFWBackend) keyCallback() KeyCallback {
	return b.keyCb
}

func (b *GLFWBackend) SetSizeChangeCallback(cbfun SizeChangeCallback) {
	b.sizeCb = cbfun
	C.igGLFWWindow_SetSizeCallback(b.handle())
}

func (b *GLFWBackend) sizeCallback() SizeChangeCallback {
	return b.sizeCb
}

func (b *GLFWBackend) SetSwapInterval(interval GLFWWindowFlags) error {
	C.glfwSwapInterval(C.int(interval))
	return nil
}

func (b *GLFWBackend) SetCursorPos(x, y float64) {
	C.glfwSetCursorPos(b.handle(), C.double(x), C.double(y))
}

func (b *GLFWBackend) SetInputMode(mode GLFWWindowFlags, value GLFWWindowFlags) {
	C.glfwSetInputMode(b.handle(), C.int(mode), C.int(value))
}
