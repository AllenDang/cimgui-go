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
	"unsafe"
)

type GLFWWindowFlags int

const (
	GLFWWindowFlagsNone         GLFWWindowFlags = GLFWWindowFlags(C.GLFWWindowFlagsNone)
	GLFWWindowFlagsNotResizable GLFWWindowFlags = GLFWWindowFlags(C.GLFWWindowFlagsNotResizable)
	GLFWWindowFlagsMaximized    GLFWWindowFlags = GLFWWindowFlags(C.GLFWWindowFlagsMaximized)
	GLFWWindowFlagsFloating     GLFWWindowFlags = GLFWWindowFlags(C.GLFWWindowFlagsFloating)
	GLFWWindowFlagsFrameless    GLFWWindowFlags = GLFWWindowFlags(C.GLFWWindowFlagsFrameless)
	GLFWWindowFlagsTransparent  GLFWWindowFlags = GLFWWindowFlags(C.GLFWWindowFlagsTransparent)
)

type voidCallbackFunc func()

type GLFWBackend struct {
	afterCreateContext   voidCallbackFunc
	loop                 voidCallbackFunc
	beforeRender         voidCallbackFunc
	afterRender          voidCallbackFunc
	beforeDestoryContext voidCallbackFunc
	dropCB               DropCallback
	window               uintptr
}

func NewGLFWBackend() *GLFWBackend {
	return &GLFWBackend{}
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
	C.igRunLoop(b.handle(), C.VoidCallback(C.loopCallback), C.VoidCallback(C.beforeRender), C.VoidCallback(C.afterRender), C.VoidCallback(C.beforeDestoryContext))
}

func (b *GLFWBackend) loopFunc() func() {
	return b.loop
}

func (b *GLFWBackend) dropCallback() DropCallback {
	return b.dropCB
}

func (b *GLFWBackend) SetWindowPos(x, y int) {
	C.igGLFWWindow_SetWindowPos(b.handle(), C.int(x), C.int(y))
}

func (b *GLFWBackend) SetWindowSize(width, height int) {
	C.igGLFWWindow_SetSize(b.handle(), C.int(width), C.int(height))
}

func (b GLFWBackend) DisplaySize() (width int32, height int32) {
	widthArg, widthFin := wrapNumberPtr[C.int, int32](&width)
	defer widthFin()

	heightArg, heightFin := wrapNumberPtr[C.int, int32](&height)
	defer heightFin()

	C.igGLFWWindow_GetDisplaySize(b.handle(), widthArg, heightArg)

	return
}

func (b *GLFWBackend) SetWindowTitle(title string) {
	titleArg, titleFin := wrapString(title)
	defer titleFin()

	C.igGLFWWindow_SetTitle(b.handle(), titleArg)
}

func (b GLFWBackend) SetShouldClose(value bool) {
	C.igGLFWWindow_SetShouldClose(b.handle(), C.int(castBool(value)))
}

func (b *GLFWBackend) CreateWindow(title string, width, height int, flags GLFWWindowFlags) {
	titleArg, titleFin := wrapString(title)
	defer titleFin()

	// b.window = uintptr(unsafe.Pointer(C.igCreateGLFWWindow(titleArg, C.int(width), C.int(height), C.GLFWWindowFlags(flags), C.VoidCallback(C.afterCreateContext))))
	b.window = uintptr(unsafe.Pointer(C.igCreateGLFWWindow(
		titleArg,
		C.int(width),
		C.int(height),
		C.GLFWWindowFlags(flags),
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
	return TextureID(C.igCreateTexture((*C.uchar)(pixels), C.int(width), C.int(height)))
}

func (b *GLFWBackend) CreateTextureRgba(img *image.RGBA, width, height int) TextureID {
	return TextureID(C.igCreateTexture((*C.uchar)(&(img.Pix[0])), C.int(width), C.int(height)))
}

func (b *GLFWBackend) DeleteTexture(id TextureID) {
	C.igDeleteTexture(C.ImTextureID(id))
}

// SetDropCallback sets the drop callback which is called when an object
// is dropped over the window.
func (b *GLFWBackend) SetDropCallback(cbfun DropCallback) {
	b.dropCB = cbfun
	C.igGLFWWindow_SetDropCallbackCB(b.handle())
}
