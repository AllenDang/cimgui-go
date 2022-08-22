package cimgui

// #cgo amd64,linux LDFLAGS: ${SRCDIR}/lib/linux/x64/libglfw3.a
// #cgo amd64,windows LDFLAGS: -L${SRCDIR}/lib/windows/x64 -l:libglfw3.a -lgdi32 -lopengl32 -limm32
// #cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo
// #cgo amd64,darwin LDFLAGS: ${SRCDIR}/lib/macos/x64/libglfw3.a
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/lib/macos/arm64/libglfw3.a
// #cgo !gles2,darwin LDFLAGS: -framework OpenGL
// #cgo gles2,darwin LDFLAGS: -lGLESv2
// extern void glfwWindowLoopCallback();
// extern void glfwBeforeRender();
// extern void glfwAfterRender();
// #include <stdint.h>
// #include "backend.h"
import "C"
import (
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
	loopFunc     voidCallbackFunc
	beforeRender voidCallbackFunc
	afterRender  voidCallbackFunc
)

type GLFWwindow uintptr

func (w GLFWwindow) handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(unsafe.Pointer(w))
}

func (w GLFWwindow) Run(loop func(), beforeRenderFunc func(), afterRenderFunc func()) {
	loopFunc = loop
	beforeRender = beforeRenderFunc
	afterRender = afterRenderFunc
	C.igRunLoop(w.handle(), C.VoidCallback(C.glfwWindowLoopCallback), C.VoidCallback(C.glfwBeforeRender), C.VoidCallback(C.glfwAfterRender))
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

func CreateGlfwWindow(title string, width, height int, flags GLFWWindowFlags) GLFWwindow {
	titleArg, titleFin := wrapString(title)
	defer titleFin()

	window := GLFWwindow(unsafe.Pointer(C.igCreateGLFWWindow(titleArg, C.int(width), C.int(height), C.GLFWWindowFlags(flags))))
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

func CreateTexture(pixels *C.uchar, width, height int) ImTextureID {
	return ImTextureID(C.igCreateTexture(pixels, C.int(width), C.int(height)))
}
