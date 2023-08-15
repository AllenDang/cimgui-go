package imgui

// #include "util.h"
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// VertexBufferLayout returns the byte sizes necessary to select fields in a vertex buffer of a DrawList.
func VertexBufferLayout() (entrySize int, posOffset int, uvOffset int, colOffset int) {
	var entrySizeArg C.size_t
	var posOffsetArg C.size_t
	var uvOffsetArg C.size_t
	var colOffsetArg C.size_t
	C.wrap_GetVertexBufferLayout(&entrySizeArg, &posOffsetArg, &uvOffsetArg, &colOffsetArg)
	entrySize = int(entrySizeArg)
	posOffset = int(posOffsetArg)
	uvOffset = int(uvOffsetArg)
	colOffset = int(colOffsetArg)
	return
}

// IndexBufferLayout returns the byte size necessary to select fields in an index buffer of DrawList.
func IndexBufferLayout() (entrySize int) {
	var entrySizeArg C.size_t
	C.wrap_GetIndexBufferLayout(&entrySizeArg)
	entrySize = int(entrySizeArg)
	return
}

type GlyphRange uintptr

func NewGlyphRange() GlyphRange {
	return GlyphRange(unsafe.Pointer(C.wrap_NewGlyphRange()))
}

func (gr GlyphRange) handle() *C.ImVector_ImWchar {
	return (*C.ImVector_ImWchar)(unsafe.Pointer(gr))
}

func (gr GlyphRange) Destroy() {
	C.wrap_DestroyGlyphRange(gr.handle())
}

func (gr GlyphRange) Data() *Wchar {
	return (*Wchar)(C.wrap_GlyphRange_GetData(gr.handle()))
}

func (fa FontAtlas) FontCount() int {
	selfArg, selfFin := fa.handle()
	defer selfFin()

	return int(C.wrap_ImFontAtlas_GetFontCount(selfArg))
}

func (self FontAtlas) TextureDataAsAlpha8() (pixels unsafe.Pointer, width int32, height int32, outBytesPerPixel int32) {
	var p *C.uchar
	var w C.int
	var h C.int
	var bp C.int

	selfArg, selfFin := self.handle()
	defer selfFin()

	C.ImFontAtlas_GetTexDataAsAlpha8(selfArg, &p, &w, &h, &bp)

	pixels = unsafe.Pointer(p)
	width = int32(w)
	height = int32(h)
	outBytesPerPixel = int32(bp)

	return
}

func (self FontAtlas) GetTextureDataAsRGBA32() (pixels unsafe.Pointer, width int32, height int32, outBytesPerPixel int32) {
	var p *C.uchar
	var w C.int
	var h C.int
	var bp C.int

	selfArg, selfFin := self.handle()
	defer selfFin()

	C.ImFontAtlas_GetTexDataAsRGBA32(selfArg, &p, &w, &h, &bp)

	pixels = unsafe.Pointer(p)
	width = int32(w)
	height = int32(h)
	outBytesPerPixel = int32(bp)

	return
}

// Ptr takes a slice or pointer (to a singular scalar value or the first
// element of an array or slice) and returns its GL-compatible address.
//
// For example:
//
//	var data []uint8
//	...
//	gl.TexImage2D(gl.TEXTURE_2D, ..., gl.UNSIGNED_BYTE, gl.Ptr(&data[0]))
func Ptr(data interface{}) unsafe.Pointer {
	if data == nil {
		return unsafe.Pointer(nil)
	}

	var addr unsafe.Pointer
	v := reflect.ValueOf(data)
	switch v.Type().Kind() {
	case reflect.Ptr:
		e := v.Elem()
		switch e.Kind() {
		case
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			addr = unsafe.Pointer(e.UnsafeAddr())
		default:
			panic(fmt.Errorf("unsupported pointer to type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", e.Kind()))
		}
	case reflect.Uintptr:
		addr = unsafe.Pointer(data.(uintptr))
	case reflect.Slice:
		addr = unsafe.Pointer(v.Index(0).UnsafeAddr())
	default:
		panic(fmt.Errorf("unsupported type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", v.Type()))
	}

	return addr
}
