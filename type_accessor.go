package imgui

// #include "cimgui_wrapper.h"
// #include "util.h"
import "C"
import "unsafe"

func (io IO) SetMouseButtonDown(i int, down bool) {
	rawIO, rawIOFin := io.handle()
	defer rawIOFin()

	C.wrap_ImGuiIO_SetMouseButtonDown(rawIO, C.int(i), C.bool(down))
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
	for i := int32(0); i < count; i++ {
		lists[i] = d.getDrawListAt(i)
	}
	return lists
}

func (d DrawData) getDrawListAt(idx int32) DrawList {
	drawDataArg, drawDataFin := d.handle()
	defer drawDataFin()
	return *newDrawListFromC(C.wrap_DrawData_GetDrawListAt(drawDataArg, C.int(idx)))
}

func (d DrawList) GetVertexBuffer() (unsafe.Pointer, int) {
	// TODO: it is possible that this field will become available on go-side after implementing more field types
	dataArg, dataFin := d.c()
	defer dataFin()
	buffer := dataArg.VtxBuffer.Data
	bufferSize := C.sizeof_ImDrawVert * dataArg.VtxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d DrawList) GetIndexBuffer() (unsafe.Pointer, int) {
	// TODO: it is possible that this field will become available on go-side after implementing more field types
	dataArg, dataFin := d.c()
	defer dataFin()
	buffer := dataArg.IdxBuffer.Data
	bufferSize := C.sizeof_ImDrawIdx * dataArg.IdxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d DrawList) getDrawCmdAt(idx int) DrawCmd {
	dataArg, dataFin := d.handle()
	defer dataFin()

	return *newDrawCmdFromC(C.wrap_DrawList_GetDrawCmdAt(dataArg, C.int(idx)))
}

func (d DrawList) Commands() []DrawCmd {
	// TODO: it is possible that this field will become available on go-side after implementing more field types
	dataArg, dataFin := d.c()
	defer dataFin()

	count := int(dataArg.CmdBuffer.Size)
	cmds := make([]DrawCmd, count)
	for i := 0; i < count; i++ {
		cmds[i] = d.getDrawCmdAt(i)
	}
	return cmds
}

func (d DrawCmd) HasUserCallback() bool {
	dataArg, dataFin := d.handle()
	defer dataFin()

	return dataArg.UserCallback != nil
}

func (d DrawCmd) CallUserCallback(list DrawList) {
	dataArg, dataFin := d.handle()
	defer dataFin()

	listArg, listFin := list.handle()
	defer listFin()

	C.wrap_DrawCmd_CallUserCallback(listArg, dataArg)
}

func (fa FontGlyphRangesBuilder) BuildRanges(ranges GlyphRange) {
	selfArg, selfFin := fa.handle()
	defer selfFin()

	C.ImFontGlyphRangesBuilder_BuildRanges(selfArg, ranges.handle())
}
