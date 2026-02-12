//go:build darwin

package metalbackend

// #cgo CPPFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #cgo CPPFLAGS: -DIMGUI_USE_WCHAR32 -DIMGUI_DISABLE_OBSOLETE_FUNCTIONS=1
// #cgo CPPFLAGS: -I${SRCDIR}/../../cwrappers/imgui -I${SRCDIR}/../../cwrappers/imgui/backends
// #cgo CXXFLAGS: -std=c++11 -fobjc-arc -x objective-c++ -Wno-nullability-completeness -Wno-deprecated-declarations
// #cgo amd64,darwin LDFLAGS: ${SRCDIR}/../../lib/macos/x64/cimgui.a
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/../../lib/macos/arm64/cimgui.a
// #cgo CPPFLAGS: -DCIMGUI_GO_USE_METAL
// #cgo darwin LDFLAGS: -framework AppKit -framework Metal -framework MetalKit -framework QuartzCore -framework GameController
// #include <stdlib.h>
// #include "metal_backend.h"
import "C"

import (
	"image"
	"unsafe"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/internal"
)

type voidCallbackFunc func()

type MetalWindowFlags int

const (
	MetalWindowFlagResizable           = MetalWindowFlags(C.MetalWindowFlagResizable)
	MetalWindowFlagTitled              = MetalWindowFlags(C.MetalWindowFlagTitled)
	MetalWindowFlagClosable            = MetalWindowFlags(C.MetalWindowFlagClosable)
	MetalWindowFlagMiniaturizable      = MetalWindowFlags(C.MetalWindowFlagMiniaturizable)
	MetalWindowFlagFullSizeContentView = MetalWindowFlags(C.MetalWindowFlagFullSizeContentView)
	MetalWindowFlagTransparent         = MetalWindowFlags(C.MetalWindowFlagTransparent)
)

var _ backend.Backend[MetalWindowFlags] = &MetalBackend{}

type MetalBackend struct {
	afterCreateContext   voidCallbackFunc
	loop                 voidCallbackFunc
	beforeRender         voidCallbackFunc
	afterRender          voidCallbackFunc
	beforeDestroyContext voidCallbackFunc
	dropCB               backend.DropCallback
	closeCB              func(pointer unsafe.Pointer)
	keyCb                backend.KeyCallback
	sizeCb               backend.SizeChangeCallback
	window               uintptr
	bgColor              imgui.Vec4
	hasBgColor           bool
	targetFPS            uint
}

func NewMetalBackend() *MetalBackend {
	if C.igInitMetal() == 0 {
		panic("Failed to initialize Metal backend")
	}

	return &MetalBackend{targetFPS: 60}
}

func (b *MetalBackend) handle() *C.MetalBackendContext {
	return (*C.MetalBackendContext)(unsafe.Pointer(b.window))
}

func (b *MetalBackend) SetAfterCreateContextHook(hook func()) {
	b.afterCreateContext = hook
}

func (b *MetalBackend) AfterCreateHook() func() {
	return b.afterCreateContext
}

func (b *MetalBackend) SetBeforeDestroyContextHook(hook func()) {
	b.beforeDestroyContext = hook
}

func (b *MetalBackend) BeforeDestroyHook() func() {
	return b.beforeDestroyContext
}

func (b *MetalBackend) SetBeforeRenderHook(hook func()) {
	b.beforeRender = hook
}

func (b *MetalBackend) BeforeRenderHook() func() {
	return b.beforeRender
}

func (b *MetalBackend) SetAfterRenderHook(hook func()) {
	b.afterRender = hook
}

func (b *MetalBackend) AfterRenderHook() func() {
	return b.afterRender
}

func (b *MetalBackend) SetBgColor(color imgui.Vec4) {
	b.bgColor = color
	b.hasBgColor = true
	if b.window == 0 {
		return
	}

	c := color.ToC()
	C.igMetalSetBgColor(b.handle(), *((*C.ImVec4)(unsafe.Pointer(&c))))
}

func (b *MetalBackend) Run(loop func()) {
	b.loop = loop
	C.igMetalRunLoop(
		b.handle(),
		C.VoidCallback(backend.LoopCallback()),
		C.VoidCallback(backend.BeforeRender()),
		C.VoidCallback(backend.AfterRender()),
		C.VoidCallback(backend.BeforeDestroyContext()),
	)
}

func (b *MetalBackend) LoopFunc() func() {
	return b.loop
}

func (b *MetalBackend) DropCallback() backend.DropCallback {
	return b.dropCB
}

func (b *MetalBackend) CloseCallback() func(wnd unsafe.Pointer) {
	return b.closeCB
}

func (b *MetalBackend) SetWindowPos(x, y int) {
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetWindowPos(b.handle(), C.int(x), C.int(y))
}

func (b *MetalBackend) GetWindowPos() (x, y int32) {
	if b.window == 0 {
		return
	}

	xArg, xFin := internal.WrapNumberPtr[C.int, int32](&x)
	defer xFin()

	yArg, yFin := internal.WrapNumberPtr[C.int, int32](&y)
	defer yFin()

	C.igMetalWindow_GetWindowPos(b.handle(), xArg, yArg)
	return
}

func (b *MetalBackend) SetWindowSize(width, height int) {
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetSize(b.handle(), C.int(width), C.int(height))
}

func (b *MetalBackend) DisplaySize() (width int32, height int32) {
	if b.window == 0 {
		return
	}

	widthArg, widthFin := internal.WrapNumberPtr[C.int, int32](&width)
	defer widthFin()

	heightArg, heightFin := internal.WrapNumberPtr[C.int, int32](&height)
	defer heightFin()

	C.igMetalWindow_GetDisplaySize(b.handle(), widthArg, heightArg)
	return
}

func (b *MetalBackend) ContentScale() (width, height float32) {
	if b.window == 0 {
		return
	}

	widthArg, widthFin := internal.WrapNumberPtr[C.float, float32](&width)
	defer widthFin()

	heightArg, heightFin := internal.WrapNumberPtr[C.float, float32](&height)
	defer heightFin()

	C.igMetalWindow_GetContentScale(b.handle(), widthArg, heightArg)
	return
}

func (b *MetalBackend) SetWindowTitle(title string) {
	if b.window == 0 {
		return
	}

	titleArg, titleFin := internal.WrapString[C.char](title)
	defer titleFin()
	C.igMetalWindow_SetTitle(b.handle(), (*C.char)(titleArg))
}

func (b *MetalBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetSizeLimits(b.handle(), C.int(minWidth), C.int(minHeight), C.int(maxWidth), C.int(maxHeight))
}

func (b *MetalBackend) SetShouldClose(value bool) {
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetShouldClose(b.handle(), C.bool(value))
}

func (b *MetalBackend) CreateWindow(title string, width, height int) {
	titleArg, titleFin := internal.WrapString[C.char](title)
	defer titleFin()

	b.window = uintptr(unsafe.Pointer(C.igCreateMetalWindow(
		(*C.char)(titleArg),
		C.int(width),
		C.int(height),
	)))

	if b.window == 0 {
		panic("Failed to create Metal window")
	}

	if b.hasBgColor {
		c := b.bgColor.ToC()
		C.igMetalSetBgColor(b.handle(), *((*C.ImVec4)(unsafe.Pointer(&c))))
	}

	C.igMetalSetTargetFPS(b.handle(), C.uint(b.targetFPS))

	if b.dropCB != nil {
		C.igMetalWindow_SetDropCallback(b.handle())
	}
	if b.closeCB != nil {
		C.igMetalWindow_SetCloseCallback(b.handle())
	}
	if b.sizeCb != nil {
		C.igMetalWindow_SetSizeCallback(b.handle())
	}

	if b.afterCreateContext != nil {
		b.afterCreateContext()
	}
}

func (b *MetalBackend) SetTargetFPS(fps uint) {
	b.targetFPS = fps
	if b.window == 0 {
		return
	}
	C.igMetalSetTargetFPS(b.handle(), C.uint(fps))
}

func (b *MetalBackend) Refresh() {
	if b.window == 0 {
		return
	}
	C.igMetalRefresh(b.handle())
}

func (b *MetalBackend) CreateTexture(pixels unsafe.Pointer, width, height int) imgui.TextureRef {
	if b.window == 0 {
		panic("Metal backend not initialized")
	}
	tex := C.igMetalCreateTexture(b.handle(), (*C.uchar)(pixels), C.int(width), C.int(height))
	return *imgui.NewTextureRefTextureID(*imgui.NewTextureIDFromC(&tex))
}

func (b *MetalBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureRef {
	if len(img.Pix) == 0 {
		panic("CreateTextureRgba called with empty pixel buffer")
	}
	if b.window == 0 {
		panic("Metal backend not initialized")
	}
	return b.CreateTexture(unsafe.Pointer(&img.Pix[0]), width, height)
}

func (b *MetalBackend) DeleteTexture(id imgui.TextureRef) {
	C.igMetalDeleteTexture(C.ImTextureID(id.TexID()))
}

func (b *MetalBackend) SetDropCallback(cbfun backend.DropCallback) {
	b.dropCB = cbfun
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetDropCallback(b.handle())
}

func (b *MetalBackend) SetCloseCallback(cbfun backend.WindowCloseCallback) {
	b.closeCB = func(_ unsafe.Pointer) {
		cbfun()
	}
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetCloseCallback(b.handle())
}

func (b *MetalBackend) SetKeyCallback(cbfun backend.KeyCallback) {
	// TODO: Key event forwarding is not wired yet on the Cocoa side.
	b.keyCb = cbfun
}

func (b *MetalBackend) KeyCallback() backend.KeyCallback {
	return b.keyCb
}

func (b *MetalBackend) SetSizeChangeCallback(cbfun backend.SizeChangeCallback) {
	b.sizeCb = cbfun
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetSizeCallback(b.handle())
}

func (b *MetalBackend) SizeCallback() backend.SizeChangeCallback {
	return b.sizeCb
}

func (b *MetalBackend) SetWindowFlags(flag MetalWindowFlags, value int) {
	if b.window == 0 {
		return
	}
	C.igMetalWindow_SetFlag(b.handle(), C.MetalWindowFlag(flag), C.int(value))
}

func (b *MetalBackend) SetIcons(_ ...image.Image) {
	// TODO: no-op for now; macOS app icons are usually set via Info.plist.
}

func (b *MetalBackend) SetSwapInterval(_ MetalWindowFlags) error {
	// Metal presents are synced by the system.
	return nil
}

func (b *MetalBackend) SetCursorPos(_x, _y float64) {
	// TODO: Not implemented for Metal backend.
}

func (b *MetalBackend) SetInputMode(_mode MetalWindowFlags, _value MetalWindowFlags) {
	// TODO: Not implemented for Metal backend.
}
