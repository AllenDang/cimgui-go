package cimgui

// #include "cimgui_wrapper.h"
// #include "util.h"
import "C"
import "unsafe"

func (io ImGuiIO) GetBackendFlags() ImGuiBackendFlags {
	return ImGuiBackendFlags(io.C().BackendFlags)
}

func (io ImGuiIO) GetConfigFlags() ImGuiConfigFlags {
	return ImGuiConfigFlags(io.C().ConfigFlags)
}

func (io ImGuiIO) SetMouseButtonDown(i int, down bool) {
	C.ImGuiIO_SetMouseButtonDown(io.handle(), C.int(i), C.bool(down))
}

func (io ImGuiIO) GetMouseDrawCursor() bool {
	return bool(io.C().MouseDrawCursor)
}

func (io ImGuiIO) AddMouseWheelDelta(horizontal, vertical float32) {
	ioC := io.C()
	v := ioC.MouseWheel + C.float(vertical)
	h := ioC.MouseWheelH + C.float(horizontal)
	io.SetMouseWheel(float32(v))
	io.SetMouseWheelH(float32(h))
}

func (io ImGuiIO) GetFonts() ImFontAtlas {
	return (ImFontAtlas)(unsafe.Pointer(io.C().Fonts))
}

// Commands returns the list of draw commands.
// Typically 1 command = 1 GPU draw call, unless the command is a callback.
func (d ImDrawData) CommandLists() []ImDrawList {
	count := int(d.C().CmdListsCount)
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
	buffer := d.C().VtxBuffer.Data
	bufferSize := C.sizeof_ImDrawVert * d.C().VtxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d ImDrawList) GetIndexBuffer() (unsafe.Pointer, int) {
	buffer := d.C().IdxBuffer.Data
	bufferSize := C.sizeof_ImDrawIdx * d.C().IdxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d ImDrawList) getDrawCmdAt(idx int) ImDrawCmd {
	return (ImDrawCmd)(unsafe.Pointer(C.DrawList_GetDrawCmdAt(d.handle(), C.int(idx))))
}

func (d ImDrawList) Commands() []ImDrawCmd {
	count := int(d.C().CmdBuffer.Size)
	cmds := make([]ImDrawCmd, count)
	for i := 0; i < count; i++ {
		cmds[i] = d.getDrawCmdAt(i)
	}
	return cmds
}

func (d ImDrawCmd) HasUserCallback() bool {
	return d.C().UserCallback != nil
}

func (d ImDrawCmd) CallUserCallback(list ImDrawList) {
	C.DrawCmd_CallUserCallback(list.handle(), d.handle())
}

func (d ImDrawCmd) TextureID() ImTextureID {
	return ImTextureID(d.C().TextureId)
}
