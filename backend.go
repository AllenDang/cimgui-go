package imgui

// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestoryContext();
// extern void dropCallback(void*, int, char**);
// extern void keyCallback(void*, int, int, int, int);
// extern void sizeCallback(void*, int, int);
import "C"

import (
	"image"
	"unsafe"
)

type voidCallbackFunc func()

var currentBackend backendCExpose

// TODO: Maybe we should get rid of it?
var textureManager TextureManager

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

//export keyCallback
func keyCallback(_ unsafe.Pointer, key, scanCode, action, mods C.int) {
	if currentBackend != nil {
		if f := currentBackend.keyCallback(); f != nil {
			f(int(key), int(scanCode), int(action), int(mods))
		}
	}
}

//export sizeCallback
func sizeCallback(_ unsafe.Pointer, w, h C.int) {
	if currentBackend != nil {
		if f := currentBackend.sizeCallback(); f != nil {
			f(int(w), int(h))
		}
	}
}

type DropCallback func([]string)
type KeyCallback func(key, scanCode, action, mods int)
type SizeChangeCallback func(w, h int)

type WindowCloseCallback[BackendFlagsT ~int] func(b Backend[BackendFlagsT])

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
type Backend[BackendFlagsT ~int] interface {
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
	ContentScale() (xScale, yScale float32)

	SetTargetFPS(fps uint)

	SetDropCallback(DropCallback)
	SetCloseCallback(WindowCloseCallback[BackendFlagsT])
	SetKeyCallback(KeyCallback)
	SetSizeChangeCallback(SizeChangeCallback)
	// SetWindowFlags selected hint to specified value.
	// ATTENTION: This method is able to set only one flag per call.
	SetWindowFlags(flag BackendFlagsT, value int)
	SetIcons(icons ...image.Image)

	CreateWindow(title string, width, height int)

	backendCExpose
	TextureManager
}

// TextureManager is a part of Backend.
//
// Why I separate it? Current impl of local texture.go needs to store this somewhere, and I don't want
// to make Texture relate on BackendFlagsT.
type TextureManager interface {
	CreateTexture(pixels unsafe.Pointer, width, Height int) TextureID
	CreateTextureRgba(img *image.RGBA, width, height int) TextureID
	DeleteTexture(id TextureID)
}

type backendCExpose interface {

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
	keyCallback() KeyCallback
	sizeCallback() SizeChangeCallback
}

func CreateBackend[BackendFlagsT ~int](backend Backend[BackendFlagsT]) Backend[BackendFlagsT] {
	currentBackend = backend
	textureManager = backend
	return backend
}
