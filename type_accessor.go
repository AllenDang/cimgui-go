package imgui

// #include "cimgui_wrapper.h"
// #include "util.h"
import "C"
import "unsafe"

func (io ImGuiIO) SetMouseButtonDown(i int, down bool) {
	C.ImGuiIO_SetMouseButtonDown(io.handle(), C.int(i), C.bool(down))
}

func (io ImGuiIO) AddMouseWheelDelta(horizontal, vertical float32) {
	v := io.GetMouseWheel() + vertical
	h := io.GetMouseWheelH() + horizontal
	io.SetMouseWheel(v)
	io.SetMouseWheelH(h)
}

// Commands returns the list of draw commands.
// Typically 1 command = 1 GPU draw call, unless the command is a callback.
func (d ImDrawData) CommandLists() []ImDrawList {
	count := d.GetCmdListsCount()
	lists := make([]ImDrawList, count)
	for i := 0; i < count; i++ {
		lists[i] = d.getDrawListAt(i)
	}
	return lists
}

func (d ImDrawData) getDrawListAt(idx int) ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.DrawData_GetDrawListAt(d.handle(), C.int(idx))))
}

func (d ImDrawList) GetVertexBuffer() (unsafe.Pointer, int) {
	buffer := d.c().VtxBuffer.Data
	bufferSize := C.sizeof_ImDrawVert * d.c().VtxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d ImDrawList) GetIndexBuffer() (unsafe.Pointer, int) {
	buffer := d.c().IdxBuffer.Data
	bufferSize := C.sizeof_ImDrawIdx * d.c().IdxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d ImDrawList) getDrawCmdAt(idx int) ImDrawCmd {
	return (ImDrawCmd)(unsafe.Pointer(C.DrawList_GetDrawCmdAt(d.handle(), C.int(idx))))
}

func (d ImDrawList) Commands() []ImDrawCmd {
	count := int(d.c().CmdBuffer.Size)
	cmds := make([]ImDrawCmd, count)
	for i := 0; i < count; i++ {
		cmds[i] = d.getDrawCmdAt(i)
	}
	return cmds
}

func (d ImDrawCmd) HasUserCallback() bool {
	return d.c().UserCallback != nil
}

func (d ImDrawCmd) CallUserCallback(list ImDrawList) {
	C.DrawCmd_CallUserCallback(list.handle(), d.handle())
}

func (fa ImFontGlyphRangesBuilder) BuildRanges(ranges GlyphRange) {
	C.ImFontGlyphRangesBuilder_BuildRanges(fa.handle(), ranges.handle())
}
