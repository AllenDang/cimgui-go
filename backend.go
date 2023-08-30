package imgui

// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestoryContext();
// extern void dropCallback(void*, int, char**);
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

type DropCallback func([]string)

type WindowCloseCallback func(b Backend)

//export closeCallback
func closeCallback(wnd unsafe.Pointer) {
	currentBackend.closeCallback()(wnd)
}

//export dropCallback
func dropCallback(wnd unsafe.Pointer, count C.int, names **C.char) {
	namesSlice := make([]string, int(count))
	for i := 0; i < int(count); i++ {
		var x *C.char
		p := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(names)) + uintptr(i)*unsafe.Sizeof(x)))
		namesSlice[i] = C.GoString(*p)
	}

	currentBackend.dropCallback()(namesSlice)
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

	SetWindowPos(x, y int)
	GetWindowPos() (x, y int32)
	SetWindowSize(width, height int)
	SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int)
	SetWindowTitle(title string)
	DisplaySize() (width, height int32)
	SetShouldClose(bool)

	SetTargetFPS(fps uint)

	CreateTexture(pixels unsafe.Pointer, width, Height int) TextureID
	CreateTextureRgba(img *image.RGBA, width, height int) TextureID
	DeleteTexture(id TextureID)
	SetDropCallback(DropCallback)
	SetCloseCallback(WindowCloseCallback)
	// SetWindowHint selected hint to specified value.
	// For list of hints check GLFW source code.
	// TODO: this needs generic layer
	SetWindowHint(hint, value int)
	SetIcons(icons ...image.Image)

	// TODO: flags needs generic layer
	CreateWindow(title string, width, height int, flags GLFWWindowFlags)

	// for C callbacks
	// What happens here is a bit tricky:
	// - user sets these callbacks via Set* methods of the backend
	// - callbacks are stored in currentContext variable
	// - functions below just returns that callbacks
	// - on top of this file is a set of global function with names similar to these below
	// - these functions are exported to C
	// - backend implementation uses C references to Go callbacks to share them (again ;-) )
	//   into backend code.
	// As you can see this is all to convert Go callback int C callback
	afterCreateHook() func()
	beforeRenderHook() func()
	loopFunc() func()
	afterRenderHook() func()
	beforeDestroyHook() func()
	dropCallback() DropCallback
	closeCallback() func(window unsafe.Pointer)
}

func CreateBackend(backend Backend) Backend {
	currentBackend = backend
	return currentBackend
}

func GetBackend() Backend {
	return currentBackend
}
