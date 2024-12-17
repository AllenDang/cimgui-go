package backend

// extern void loopCallback();
// extern void beforeRender();
// extern void afterRender();
// extern void afterCreateContext();
// extern void beforeDestroyContext();
// extern void dropCallback(void*, int, char**);
// extern void keyCallback(void*, int, int, int, int);
// extern void sizeCallback(void*, int, int);
// #include "../imgui/extra_types.h"
// #include "../imgui/wrapper.h"
// #include "../imgui/typedefs.h"
import "C"

import (
	"errors"
	"fmt"
	"image"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
)

type voidCallbackFunc func()

var currentBackend backendCExpose

// TODO: Maybe we should get rid of it?
var textureManager TextureManager

//export loopCallback
func loopCallback() {
	if currentBackend != nil {
		if f := currentBackend.LoopFunc(); f != nil {
			f()
		}
	}
}

func LoopCallback() unsafe.Pointer {
	return C.loopCallback
}

//export beforeRender
func beforeRender() {
	if currentBackend != nil {
		if f := currentBackend.BeforeRenderHook(); f != nil {
			f()
		}
	}
}

func BeforeRender() unsafe.Pointer {
	return C.beforeRender
}

//export afterRender
func afterRender() {
	if currentBackend != nil {
		if f := currentBackend.AfterRenderHook(); f != nil {
			f()
		}
	}
}

func AfterRender() unsafe.Pointer {
	return C.afterRender
}

//export afterCreateContext
func afterCreateContext() {
	if currentBackend != nil {
		if f := currentBackend.AfterCreateHook(); f != nil {
			f()
		}
	}
}

func AfterCreateContext() unsafe.Pointer {
	return C.afterCreateContext
}

//export beforeDestroyContext
func beforeDestroyContext() {
	if currentBackend != nil {
		if f := currentBackend.BeforeDestroyHook(); f != nil {
			f()
		}
	}
}

func BeforeDestroyContext() unsafe.Pointer {
	return C.beforeDestroyContext
}

//export keyCallback
func keyCallback(_ unsafe.Pointer, key, scanCode, action, mods C.int) {
	if currentBackend != nil {
		if f := currentBackend.KeyCallback(); f != nil {
			f(int(key), int(scanCode), int(action), int(mods))
		}
	}
}

//export sizeCallback
func sizeCallback(_ unsafe.Pointer, w, h C.int) {
	if currentBackend != nil {
		if f := currentBackend.SizeCallback(); f != nil {
			f(int(w), int(h))
		}
	}
}

type (
	DropCallback       func([]string)
	KeyCallback        func(key, scanCode, action, mods int)
	SizeChangeCallback func(w, h int)
)

type WindowCloseCallback func()

//export closeCallback
func closeCallback(wnd unsafe.Pointer) {
	currentBackend.CloseCallback()(wnd)
}

//export dropCallback
func dropCallback(wnd unsafe.Pointer, count C.int, names **C.char) {
	namesSlice := make([]string, int(count))
	for i := 0; i < int(count); i++ {
		var x *C.char
		p := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(names)) + uintptr(i)*unsafe.Sizeof(x)))
		namesSlice[i] = C.GoString(*p)
	}

	currentBackend.DropCallback()(namesSlice)
}

// Backend is a special interface that implements all methods required
// to render imgui application.
type Backend[BackendFlagsT ~int] interface {
	SetAfterCreateContextHook(func())
	SetBeforeDestroyContextHook(func())
	SetBeforeRenderHook(func())
	SetAfterRenderHook(func())

	SetBgColor(color imgui.Vec4)
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
	SetCloseCallback(WindowCloseCallback)
	SetKeyCallback(KeyCallback)
	SetSizeChangeCallback(SizeChangeCallback)
	// SetWindowFlags selected hint to specified value.
	// ATTENTION: This method is able to set only one flag per call.
	SetWindowFlags(flag BackendFlagsT, value int)
	SetIcons(icons ...image.Image)
	SetSwapInterval(interval BackendFlagsT) error
	SetCursorPos(x, y float64)
	SetInputMode(mode BackendFlagsT, value BackendFlagsT)

	CreateWindow(title string, width, height int)

	TextureManager
}

// TextureManager is a part of Backend.
//
// Why I separate it? Current impl of local texture.go needs to store this somewhere, and I don't want
// to make Texture relate on BackendFlagsT.
type TextureManager interface {
	CreateTexture(pixels unsafe.Pointer, width, Height int) imgui.TextureID
	CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureID
	DeleteTexture(id imgui.TextureID)
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
	AfterCreateHook() func()
	BeforeRenderHook() func()
	LoopFunc() func()
	AfterRenderHook() func()
	BeforeDestroyHook() func()
	DropCallback() DropCallback
	CloseCallback() func(window unsafe.Pointer)
	KeyCallback() KeyCallback
	SizeCallback() SizeChangeCallback
}

var CExposerError error = errors.New("CreateBackend was unable to extract C API Exposer from your backend. This is an error if you want to use C-related backend.")

// CreateBackend should to be called before using any cimgui methods.
// This enables the following features:
//   - C-related callbacks: if your backend implements C functions, this will export some methods
//     from your backend to C.
//   - Advanced Texture Loading stuff. GL texture loading is a little complex. texture.go has some high-level logic.
//
// Error:
//   - Error could be CExposerError when your backend does not support C. THIS MAY STILL WORK CORRECTLY.
//     Specificly, this may happen if you use some backend that is not part of cimgui-go. In such a case
//     just ignore the error.
func CreateBackend[BackendFlagsT ~int](backend Backend[BackendFlagsT]) (sameBackend Backend[BackendFlagsT], err error) {
	var ok bool
	if currentBackend, ok = backend.(backendCExpose); !ok {
		err = fmt.Errorf("%T does not implement %T: %w", backend, currentBackend, CExposerError)
	}

	textureManager = backend
	return backend, err
}
