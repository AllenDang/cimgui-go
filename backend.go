package imgui

// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestoryContext();
import "C"

import (
	"image"
	"unsafe"
)

var currentBackend Backend

//export loopCallback
func loopCallback() {
	if currentBackend != nil {
		if f := currentBackend.loopFunc(); f != nil {
			f()
		}
	}
}

//export beforeRender
func beforeRender() {
	if currentBackend != nil {
		if f := currentBackend.beforeRenderHook(); f != nil {
			f()
		}
	}
}

//export afterRender
func afterRender() {
	if currentBackend != nil {
		if f := currentBackend.afterRenderHook(); f != nil {
			f()
		}
	}
}

//export afterCreateContext
func afterCreateContext() {
	if currentBackend != nil {
		if f := currentBackend.afterCreateHook(); f != nil {
			f()
		}
	}
}

//export beforeDestoryContext
func beforeDestoryContext() {
	if currentBackend != nil {
		if f := currentBackend.beforeDestroyHook(); f != nil {
			f()
		}
	}
}

// Backend is a special interface that implements all methods required
// to render imgui application.
type Backend interface {
	SetAfterCreateContextHook(func())
	SetBeforeDestroyContextHook(func())
	SetBeforeRenderHook(func())
	SetAfterRenderHook(func())

	SetBgColor(color Vec4)
	Run(func())
	Refresh()

	DisplaySize() (width, height int32)
	SetShouldClose(bool)

	SetTargetFPS(fps uint)

	CreateTexture(pixels unsafe.Pointer, width, Height int) TextureID
	CreateTextureRgba(img *image.RGBA, width, height int) TextureID
	DeleteTexture(id TextureID)

	// TODO: flags needs generic layer
	CreateWindow(title string, width, height int, flags GLFWWindowFlags)

	// for C callbacks
	afterCreateHook() func()
	beforeRenderHook() func()
	loopFunc() func()
	afterRenderHook() func()
	beforeDestroyHook() func()
}

func CreateBackend( /*TODO: backend type*/ ) Backend {
	currentBackend = &GLFWBackend{}
	return currentBackend
}

func GetBackend() Backend {
	return currentBackend
}
