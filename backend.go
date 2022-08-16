package cimgui

// #cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/lib/macos/arm64/libglfw3.a
// #cgo !gles2,darwin LDFLAGS: -framework OpenGL
// #cgo gles2,darwin LDFLAGS: -lGLESv2
// extern void glfwWindowLoopCallback();
// #include <stdint.h>
// #include "backend.h"
import "C"
import (
	"unsafe"
)

type GLFWwindow uintptr
type voidCallbackFunc func()

var (
	loopFunc voidCallbackFunc
)

func (w GLFWwindow) handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(unsafe.Pointer(w))
}

func (w GLFWwindow) Run(loop func()) {
	loopFunc = loop
	C.igRunLoop(w.handle(), C.VoidCallback(C.glfwWindowLoopCallback))
}

//export glfwWindowLoopCallback
func glfwWindowLoopCallback() {
	if loopFunc != nil {
		loopFunc()
	}
}

func CreateGlfwWindow(title string, width, height int) GLFWwindow {
	titleArg, titleFin := wrapString(title)
	defer titleFin()

	window := GLFWwindow(unsafe.Pointer(C.igCreateGlfwWindow(titleArg, C.int(width), C.int(height))))
	if window == 0 {
		panic("Failed to create GLFW window")
	}

	return window
}
