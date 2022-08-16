package cimgui

// #cgo pkg-config: glfw3
// #cgo !gles2,darwin LDFLAGS: -framework OpenGL
// #cgo gles2,darwin LDFLAGS: -lGLESv2
// #include "backend.h"
import "C"

func Run() {
	C.Launch_GLFW_OpenGL3()
}
