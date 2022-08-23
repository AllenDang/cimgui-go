package cimgui

// #include "util.h"
import "C"
import "unsafe"

// VertexBufferLayout returns the byte sizes necessary to select fields in a vertex buffer of a DrawList.
func VertexBufferLayout() (entrySize int, posOffset int, uvOffset int, colOffset int) {
	var entrySizeArg C.size_t
	var posOffsetArg C.size_t
	var uvOffsetArg C.size_t
	var colOffsetArg C.size_t
	C.GetVertexBufferLayout(&entrySizeArg, &posOffsetArg, &uvOffsetArg, &colOffsetArg)
	entrySize = int(entrySizeArg)
	posOffset = int(posOffsetArg)
	uvOffset = int(uvOffsetArg)
	colOffset = int(colOffsetArg)
	return
}

// IndexBufferLayout returns the byte size necessary to select fields in an index buffer of DrawList.
func IndexBufferLayout() (entrySize int) {
	var entrySizeArg C.size_t
	C.GetIndexBufferLayout(&entrySizeArg)
	entrySize = int(entrySizeArg)
	return
}

type GlyphRange uintptr

func NewGlyphRange() GlyphRange {
	return GlyphRange(unsafe.Pointer(C.NewGlyphRange()))
}

func (gr GlyphRange) handle() *C.ImVector_ImWchar {
	return (*C.ImVector_ImWchar)(unsafe.Pointer(gr))
}

func (gr GlyphRange) Destroy() {
	C.DestroyGlyphRange(gr.handle())
}

func (gr GlyphRange) Data() *ImWchar {
	return (*ImWchar)(C.GlyphRange_GetData(gr.handle()))
}

func (fa ImFontAtlas) GetFontCount() int {
	return int(C.ImFontAtlas_GetFontCount(fa.handle()))
}

func (self ImFontAtlas) GetTextureDataAsRGBA32() (pixels unsafe.Pointer, width int32, height int32, outBytesPerPixel int32) {
	var p *C.uchar
	var w C.int
	var h C.int
	var bp C.int

	C.ImFontAtlas_GetTexDataAsRGBA32(self.handle(), &p, &w, &h, &bp)

	pixels = unsafe.Pointer(p)
	width = int32(w)
	height = int32(h)
	outBytesPerPixel = int32(bp)

	return
}
