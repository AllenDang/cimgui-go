package cimgui

// #include "cimgui_wrapper.h"
// #include "util.h"
import "C"
import "unsafe"

func (io IO) SetMouseButtonDown(i int, down bool) {
	C.ImGuiIO_SetMouseButtonDown(io.handle(), C.int(i), C.bool(down))
}

func (io IO) AddMouseWheelDelta(horizontal, vertical float32) {
	v := io.MouseWheel() + vertical
	h := io.MouseWheelH() + horizontal
	io.SetMouseWheel(v)
	io.SetMouseWheelH(h)
}

// Commands returns the list of draw commands.
// Typically 1 command = 1 GPU draw call, unless the command is a callback.
func (d DrawData) CommandLists() []DrawList {
	count := d.CmdListsCount()
	lists := make([]DrawList, count)
	for i := 0; i < count; i++ {
		lists[i] = d.getDrawListAt(i)
	}
	return lists
}

func (d DrawData) getDrawListAt(idx int) DrawList {
	return (DrawList)(unsafe.Pointer(C.DrawData_GetDrawListAt(d.handle(), C.int(idx))))
}

func (d DrawList) VertexBuffer() (unsafe.Pointer, int) {
	buffer := d.c().VtxBuffer.Data
	bufferSize := C.sizeof_ImDrawVert * d.c().VtxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d DrawList) IndexBuffer() (unsafe.Pointer, int) {
	buffer := d.c().IdxBuffer.Data
	bufferSize := C.sizeof_ImDrawIdx * d.c().IdxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d DrawList) getDrawCmdAt(idx int) DrawCmd {
	return (DrawCmd)(unsafe.Pointer(C.DrawList_GetDrawCmdAt(d.handle(), C.int(idx))))
}

func (d DrawList) Commands() []DrawCmd {
	count := int(d.c().CmdBuffer.Size)
	cmds := make([]DrawCmd, count)
	for i := 0; i < count; i++ {
		cmds[i] = d.getDrawCmdAt(i)
	}
	return cmds
}

func (d DrawCmd) HasUserCallback() bool {
	return d.c().UserCallback != nil
}

func (d DrawCmd) CallUserCallback(list DrawList) {
	C.DrawCmd_CallUserCallback(list.handle(), d.handle())
}

func (fa FontGlyphRangesBuilder) BuildRanges(ranges GlyphRange) {
	C.ImFontGlyphRangesBuilder_BuildRanges(fa.handle(), ranges.handle())
}
