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
// #include <stdlib.h>
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
	"image/draw"
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
	closeCB              func(pointer unsafe.Pointer)
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

func (b *GLFWBackend) closeCallback() func(wnd unsafe.Pointer) {
	return b.closeCB
}

func (b *GLFWBackend) SetWindowPos(x, y int) {
	C.igGLFWWindow_SetWindowPos(b.handle(), C.int(x), C.int(y))
}

func (b *GLFWBackend) GetWindowPos() (x, y int32) {
	xArg, xFin := WrapNumberPtr[C.int, int32](&x)
	defer xFin()

	yArg, yFin := WrapNumberPtr[C.int, int32](&y)
	defer yFin()

	C.igGLFWWindow_GetWindowPos(b.handle(), xArg, yArg)

	return
}

func (b *GLFWBackend) SetWindowSize(width, height int) {
	C.igGLFWWindow_SetSize(b.handle(), C.int(width), C.int(height))
}

func (b GLFWBackend) DisplaySize() (width int32, height int32) {
	widthArg, widthFin := WrapNumberPtr[C.int, int32](&width)
	defer widthFin()

	heightArg, heightFin := WrapNumberPtr[C.int, int32](&height)
	defer heightFin()

	C.igGLFWWindow_GetDisplaySize(b.handle(), widthArg, heightArg)

	return
}

func (b *GLFWBackend) SetWindowTitle(title string) {
	titleArg, titleFin := WrapString(title)
	defer titleFin()

	C.igGLFWWindow_SetTitle(b.handle(), titleArg)
}

// The minimum and maximum size of the content area of a windowed mode window.
// To specify only a minimum size or only a maximum one, set the other pair to -1
// e.g. SetWindowSizeLimits(640, 480, -1, -1)
func (b *GLFWBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	C.igGLFWWindow_SetSizeLimits(b.handle(), C.int(minWidth), C.int(minHeight), C.int(maxWidth), C.int(maxHeight))
}

func (b GLFWBackend) SetShouldClose(value bool) {
	C.igGLFWWindow_SetShouldClose(b.handle(), C.int(CastBool(value)))
}

func (b *GLFWBackend) CreateWindow(title string, width, height int, flags GLFWWindowFlags) {
	titleArg, titleFin := WrapString(title)
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

func (b *GLFWBackend) SetCloseCallback(cbfun WindowCloseCallback) {
	b.closeCB = func(_ unsafe.Pointer) {
		cbfun(b)
	}

	C.igGLFWWindow_SetCloseCallback(b.handle())
}

// SetIcons sets icons for the window.
// THIS CODE COMES FROM https://github.com/go-gl/glfw (BSD-3 clause) - Copyright (c) 2012 The glfw3-go Authors. All rights reserved.
func (b *GLFWBackend) SetIcons(images []image.Image) {
	/*
		cIcons := make([]C.GLFWimage, len(icons))
		for _, i := range icons {
			x, y := i.Bounds().Dx(), i.Bounds().Dy()
			cIcon := C.GLFWimage{
				width:  C.int(x),
				height: C.int(y),
			}

			pixels := make([]C.uchar, x*y)
			for imgY := 0; imgY < y; imgY++ {
				for imgX := 0; imgX < x; imgX++ {
					r, g, b, _ := i.At(imgX, imgY).RGBA()
					pixels[imgY*x+imgX] = (r*6/256)*36 + (g*6/256)*6 + (b * 6 / 256)
				}
			}

			cIcon.pixels = (*C.uchar)(unsafe.Pointer(&pixels[0]))
			cIcons = append(cIcons, cIcon)
		}

		C.igGLFWWindow_SetIcon(b.handle(), C.int(len(cIcons)), &cIcons[0])
	*/

	count := len(images)
	cimages := make([]C.CImage, count)
	freePixels := make([]func(), count)

	for i, img := range images {
		var pixels []uint8
		b := img.Bounds()

		switch img := img.(type) {
		case *image.NRGBA:
			pixels = img.Pix
		default:
			m := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
			draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
			pixels = m.Pix
		}

		pix, free := func(origin []byte) (pointer *uint8, free func()) {
			n := len(origin)
			if n == 0 {
				return nil, func() {}
			}

			ptr := C.CBytes(origin)
			return (*uint8)(ptr), func() { C.free(ptr) }
		}(pixels)

		freePixels[i] = free

		cimages[i].width = C.int(b.Dx())
		cimages[i].height = C.int(b.Dy())
		cimages[i].pixels = (*C.uchar)(pix)
	}

	var p *C.CImage
	if count > 0 {
		p = &cimages[0]
	}
	C.igGLFWWindow_SetIcon(b.handle(), C.int(count), p)

	for _, v := range freePixels {
		v()
	}
}
