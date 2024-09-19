package imgui

// #include "util.h"
import "C"

import (
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

func (gr GlyphRange) Handle() *C.ImVector_ImWchar {
	return (*C.ImVector_ImWchar)(unsafe.Pointer(gr))
}

func (gr GlyphRange) Destroy() {
	C.wrap_DestroyGlyphRange(gr.Handle())
}

func (gr GlyphRange) Data() *Wchar {
	return (*Wchar)(C.wrap_GlyphRange_GetData(gr.Handle()))
}

func (fa FontAtlas) FontCount() int {
	selfArg, selfFin := fa.Handle()
	defer selfFin()

	return int(C.wrap_ImFontAtlas_GetFontCount(selfArg))
}

func (self FontAtlas) TextureDataAsAlpha8() (pixels unsafe.Pointer, width int32, height int32, outBytesPerPixel int32) {
	var p *C.uchar
	var w C.int
	var h C.int
	var bp C.int

	selfArg, selfFin := self.Handle()
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

	selfArg, selfFin := self.Handle()
	defer selfFin()

	C.ImFontAtlas_GetTexDataAsRGBA32(selfArg, &p, &w, &h, &bp)

	pixels = unsafe.Pointer(p)
	width = int32(w)
	height = int32(h)
	outBytesPerPixel = int32(bp)

	return
}
