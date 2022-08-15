package cimgui

// #include "cimgui_wrapper.h"
// #include "cimgui_structs_accessor.h"
// #include "util.h"
import "C"
import "unsafe"

func (io ImGuiIO) SetBackendFlags(flags ImGuiBackendFlags) {
	C.ImGuiIO_SetBackendFlags(io.handle(), C.ImGuiBackendFlags(flags))
}

func (io ImGuiIO) GetBackendFlags() ImGuiBackendFlags {
	return ImGuiBackendFlags(io.C().BackendFlags)
}

func (io ImGuiIO) SetBackendPlatformName(name string) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.ImGuiIO_SetBackendPlatformName(io.handle(), nameArg)
}

func (io ImGuiIO) SetDisplaySize(width, height float32) {
	C.ImGuiIO_SetDisplaySize(io.handle(), ImVec2{X: width, Y: height}.ToC())
}

func (io ImGuiIO) SetDeltaTime(delta float32) {
	C.ImGuiIO_SetDeltaTime(io.handle(), C.float(delta))
}

func (io ImGuiIO) SetMousePos(x, y float32) {
	C.ImGuiIO_SetMousePos(io.handle(), ImVec2{X: x, Y: y}.ToC())
}

func (io ImGuiIO) SetMouseButtonDown(i int, down bool) {
	//TODO: implement this
}

func (io ImGuiIO) GetConfigFlags() ImGuiConfigFlags {
	return ImGuiConfigFlags(io.C().ConfigFlags)
}

func (io ImGuiIO) GetMouseDrawCursor() bool {
	return bool(io.C().MouseDrawCursor)
}

func (io ImGuiIO) AddMouseWheelDelta(horizontal, vertical float32) {
	ioC := io.C()
	ioC.MouseWheel += C.float(vertical)
	ioC.MouseWheelH += C.float(horizontal)
}

func (io ImGuiIO) AddFocusEvent(focused bool) {
	//TODO: implement this
}

func (io ImGuiIO) AddKeyEvent(key ImGuiKey, down bool) {
	IO_AddKeyEvent(io, key, down)
}

func (io ImGuiIO) AddInputCharacters(cs string) {
	IO_AddInputCharactersUTF8(io, cs)
}

func (io ImGuiIO) GetFonts() ImFontAtlas {
	return (ImFontAtlas)(unsafe.Pointer(io.C().Fonts))
}

func (d ImDrawData) ScaleClipRects(width, height float32) {
	// TODO: implement this
	// DrawData_ScaleClipRects(d, ImVec2{X: width, Y: height}.ToC())
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
	bufferSize := C.sizeof_ImDrawIdx * d.C().VtxBuffer.Size
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

func (d ImDrawCmd) TextureID() uintptr {
	return uintptr(d.C().TextureId)
}

func (f ImFontAtlas) SetTextureID(id ImTextureID) {
	FontAtlas_SetTexID(f, id)
}
