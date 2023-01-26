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
// extern void glfwWindowLoopCallback();
// extern void glfwBeforeRender();
// extern void glfwAfterRender();
// extern void glfwAfterCreateContext();
// extern void glfwBeforeDestoryContext();
// #include <stdint.h>
// #include "backend.h"
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

var (
	afterCreateContext   voidCallbackFunc
	loopFunc             voidCallbackFunc
	beforeRender         voidCallbackFunc
	afterRender          voidCallbackFunc
	beforeDestoryContext voidCallbackFunc
)

type GLFWwindow uintptr

func (w GLFWwindow) handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(unsafe.Pointer(w))
}

func SetAfterCreateContextHook(hook func()) {
	afterCreateContext = hook
}

func SetBeforeDestroyContextHook(hook func()) {
	beforeDestoryContext = hook
}

func SetBeforeRenderHook(hook func()) {
	beforeRender = hook
}

func SetAfterRenderHook(hook func()) {
	afterRender = hook
}

func SetBgColor(color Vec4) {
	C.igSetBgColor(color.toC())
}

func (w GLFWwindow) Run(loop func()) {
	loopFunc = loop
	C.igRunLoop(w.handle(), C.VoidCallback(C.glfwWindowLoopCallback), C.VoidCallback(C.glfwBeforeRender), C.VoidCallback(C.glfwAfterRender), C.VoidCallback(C.glfwBeforeDestoryContext))
}

func (w GLFWwindow) DisplaySize() (width int32, height int32) {
	widthArg, widthFin := wrapNumberPtr[C.int, int32](&width)
	defer widthFin()

	heightArg, heightFin := wrapNumberPtr[C.int, int32](&height)
	defer heightFin()

	C.igGLFWWindow_GetDisplaySize(w.handle(), widthArg, heightArg)

	return
}

func (w GLFWwindow) SetWindowShouldClose(value bool) {
	var valueInt int

	if value {
		valueInt = 1
	}

	C.igGLFWWindow_SetShouldClose(w.handle(), C.int(valueInt))
}

//export glfwWindowLoopCallback
func glfwWindowLoopCallback() {
	if loopFunc != nil {
		loopFunc()
	}
}

//export glfwBeforeRender
func glfwBeforeRender() {
	if beforeRender != nil {
		beforeRender()
	}
}

//export glfwAfterRender
func glfwAfterRender() {
	if afterRender != nil {
		afterRender()
	}
}

//export glfwAfterCreateContext
func glfwAfterCreateContext() {
	if afterCreateContext != nil {
		afterCreateContext()
	}
}

//export glfwBeforeDestoryContext
func glfwBeforeDestoryContext() {
	if beforeDestoryContext != nil {
		beforeDestoryContext()
	}
}

func CreateGlfwWindow(title string, width, height int, flags GLFWWindowFlags) GLFWwindow {
	titleArg, titleFin := wrapString(title)
	defer titleFin()

	window := GLFWwindow(unsafe.Pointer(C.igCreateGLFWWindow(titleArg, C.int(width), C.int(height), C.GLFWWindowFlags(flags), C.VoidCallback(C.glfwAfterCreateContext))))
	if window == 0 {
		panic("Failed to create GLFW window")
	}

	return window
}

func SetTargetFPS(fps uint) {
	C.igSetTargetFPS(C.uint(fps))
}

func Refresh() {
	C.igRefresh()
}

func CreateTexture(pixels unsafe.Pointer, width, height int) TextureID {
	return TextureID(C.igCreateTexture((*C.uchar)(pixels), C.int(width), C.int(height)))
}

func CreateTextureRgba(img *image.RGBA, width, height int) TextureID {
	return TextureID(C.igCreateTexture((*C.uchar)(&(img.Pix[0])), C.int(width), C.int(height)))
}

func DeleteTexture(id TextureID) {
	C.igDeleteTexture(C.ImTextureID(id))
}
