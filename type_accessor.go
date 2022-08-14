package cimgui

// #include "cimgui_wrapper.h"
// #include "util.h"
import "C"
import "unsafe"

func (v ImVec2) V() (x float32, y float32) {
	return float32(v.x), float32(v.y)
}

func (v ImVec4) V() (x float32, y float32, z float32, w float32) {
	return float32(v.x), float32(v.y), float32(v.z), float32(v.w)
}

func (io *ImGuiIO) SetBackendFlags(flags ImGuiBackendFlags) {
	io.BackendFlags = C.ImGuiBackendFlags(flags)
}

func (io *ImGuiIO) GetBackendFlags() ImGuiBackendFlags {
	return ImGuiBackendFlags(io.BackendFlags)
}

func (io *ImGuiIO) SetBackendPlatformName(name string) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	io.BackendPlatformName = nameArg
}

func (io *ImGuiIO) SetDisplaySize(width, height float32) {
	io.DisplaySize.x = C.float(width)
	io.DisplaySize.y = C.float(height)
}

func (io ImGuiIO) SetDeltaTime(delta float32) {
	io.DeltaTime = C.float(delta)
}

func (io ImGuiIO) SetMousePos(x, y float32) {
	io.MousePos.x = C.float(x)
	io.MousePos.y = C.float(y)
}

func (io ImGuiIO) SetMouseButtonDown(i int, down bool) {
	io.MouseDown[i] = C.bool(down)
}

func (io ImGuiIO) GetConfigFlags() ImGuiConfigFlags {
	return ImGuiConfigFlags(io.ConfigFlags)
}

func (io ImGuiIO) GetMouseDrawCursor() bool {
	return bool(io.MouseDrawCursor)
}

func (io ImGuiIO) AddMouseWheelDelta(horizontal, vertical float32) {
	io.MouseWheel += C.float(vertical)
	io.MouseWheelH += C.float(horizontal)
}

func (io ImGuiIO) AddFocusEvent(focused bool) {
	//TODO: implement this
}

func (io *ImGuiIO) AddKeyEvent(key ImGuiKey, down bool) {
	IO_AddKeyEvent(io, key, down)
}

func (io *ImGuiIO) AddInputCharacters(cs string) {
	IO_AddInputCharactersUTF8(io, cs)
}

func (io *ImGuiIO) GetFonts() *ImFontAtlas {
	return (*ImFontAtlas)(io.Fonts)
}

func (d *ImDrawData) ScaleClipRects(width, height float32) {
	DrawData_ScaleClipRects(d, ImVec2{x: C.float(width), y: C.float(height)})
}

// Commands returns the list of draw commands.
// Typically 1 command = 1 GPU draw call, unless the command is a callback.
func (d *ImDrawData) CommandLists() []*ImDrawList {
	count := int(d.CmdListsCount)
	lists := make([]*ImDrawList, count)
	for i := 0; i < count; i++ {
		lists[i] = d.getDrawListAt(i)
	}
	return lists
}

func (d *ImDrawData) getDrawListAt(idx int) *ImDrawList {
	return (*ImDrawList)(C.DrawData_GetDrawListAt((*C.ImDrawData)(d), C.int(idx)))
}

func (d *ImDrawList) GetVertexBuffer() (unsafe.Pointer, int) {
	buffer := d.VtxBuffer.Data
	bufferSize := C.sizeof_ImDrawIdx * d.VtxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d *ImDrawList) GetIndexBuffer() (unsafe.Pointer, int) {
	buffer := d.IdxBuffer.Data
	bufferSize := C.sizeof_ImDrawIdx * d.IdxBuffer.Size
	return unsafe.Pointer(buffer), int(bufferSize)
}

func (d *ImDrawList) getDrawCmdAt(idx int) *ImDrawCmd {
	return (*ImDrawCmd)(C.DrawList_GetDrawCmdAt((*C.ImDrawList)(d), C.int(idx)))
}

func (d *ImDrawList) Commands() []ImDrawCmd {
	count := int(d.CmdBuffer.Size)
	cmds := make([]ImDrawCmd, count)
	for i := 0; i < count; i++ {
		cmds[i] = *(d.getDrawCmdAt(i))
	}
	return cmds
}

func (d *ImDrawCmd) HasUserCallback() bool {
	return d.UserCallback != nil
}

func (d *ImDrawCmd) CallUserCallback(list *ImDrawList) {
	C.DrawCmd_CallUserCallback((*C.ImDrawList)(list), (*C.ImDrawCmd)(d))
}

func (d *ImDrawCmd) TextureID() uintptr {
	return uintptr(d.TextureId)
}

func (f *ImFontAtlas) SetTextureID(id ImTextureID) {
	FontAtlas_SetTexID(f, id)
}
