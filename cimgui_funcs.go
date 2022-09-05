package cimgui

// #include "extra_types.h"
// #include "cimgui_structs_accessor.h"
// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

func SetItemAllowOverlap() {
	C.SetItemAllowOverlap()
}

func TreeNode_Str(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TreeNode_Str(labelArg) == C.bool(true)
}

func TreeNode_StrStr(str_id string, fmt string) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	return C.TreeNode_StrStr(str_idArg, fmtArg) == C.bool(true)
}

func TreeNode_Ptr(ptr_id unsafe.Pointer, fmt string) bool {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	return C.TreeNode_Ptr(ptr_id, fmtArg) == C.bool(true)
}

func Dummy(size ImVec2) {
	C.Dummy(size.toC())
}

func EndMenuBar() {
	C.EndMenuBar()
}

func IsItemActivated() bool {
	return C.IsItemActivated() == C.bool(true)
}

// BeginPopup parameter default value hint:
// flags: 0
func BeginPopup(str_id string, flags ImGuiWindowFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopup(str_idArg, C.ImGuiWindowFlags(flags)) == C.bool(true)
}

// ProgressBar parameter default value hint:
// size_arg: ImVec2(-FLT_MIN,0)
// overlay: NULL
func ProgressBar(fraction float32, size_arg ImVec2, overlay string) {
	overlayArg, overlayFin := wrapString(overlay)
	defer overlayFin()

	C.ProgressBar(C.float(fraction), size_arg.toC(), overlayArg)
}

func ShowFontSelector(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ShowFontSelector(labelArg)
}

func (self ImFontAtlas) GetGlyphRangesChineseSimplifiedCommon() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesChineseSimplifiedCommon(self.handle()))
}

// DrawList_AddNgon parameter default value hint:
// thickness: 1.0f
func (self ImDrawList) AddNgon(center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.DrawList_AddNgon(self.handle(), center.toC(), C.float(radius), C.ImU32(col), C.int(num_segments), C.float(thickness))
}

func GetTextLineHeight() float32 {
	return float32(C.GetTextLineHeight())
}

// DestroyContext parameter default value hint:
// ctx: NULL
func DestroyContext(ctx ImGuiContext) {
	C.DestroyContext(ctx.handle())
}

func GetStyle() ImGuiStyle {
	return (ImGuiStyle)(unsafe.Pointer(C.GetStyle()))
}

func TextDisabled(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.TextDisabled(fmtArg)
}

func (self ImGuiTextBuffer) Begin() string {
	return C.GoString(C.TextBuffer_Begin(self.handle()))
}

func BulletText(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.BulletText(fmtArg)
}

func RadioButton_Bool(label string, active bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.RadioButton_Bool(labelArg, C.bool(active)) == C.bool(true)
}

func RadioButton_IntPtr(label string, v *int32, v_button int32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	return C.RadioButton_IntPtr(labelArg, vArg, C.int(v_button)) == C.bool(true)
}

// DragInt2 parameter default value hint:
// flags: 0
// format: "%d"
// v_max: 0
// v_min: 0
// v_speed: 1.0f
func DragInt2(label string, v [2]*int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt2(labelArg, (*C.int)(&vArg[0]), C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// DrawList_AddQuad parameter default value hint:
// thickness: 1.0f
func (self ImDrawList) AddQuad(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddQuad(self.handle(), p1.toC(), p2.toC(), p3.toC(), p4.toC(), C.ImU32(col), C.float(thickness))
}

// ShowStackToolWindow parameter default value hint:
// p_open: NULL
func ShowStackToolWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowStackToolWindow(p_openArg)
}

// InputInt3 parameter default value hint:
// flags: 0
func InputInt3(label string, v [3]*int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	return C.InputInt3(labelArg, (*C.int)(&vArg[0]), C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func PushAllowKeyboardFocus(allow_keyboard_focus bool) {
	C.PushAllowKeyboardFocus(C.bool(allow_keyboard_focus))
}

// ShowDebugLogWindow parameter default value hint:
// p_open: NULL
func ShowDebugLogWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDebugLogWindow(p_openArg)
}

func (self ImFontAtlas) AddCustomRectRegular(width int32, height int32) int {
	return int(C.FontAtlas_AddCustomRectRegular(self.handle(), C.int(width), C.int(height)))
}

// PopStyleColor parameter default value hint:
// count: 1
func PopStyleColor(count int32) {
	C.PopStyleColor(C.int(count))
}

func GetMousePos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMousePos(pOutArg)
}

func GetWindowContentRegionMin(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowContentRegionMin(pOutArg)
}

func (self ImGuiTextBuffer) c_str() string {
	return C.GoString(C.TextBuffer_c_str(self.handle()))
}

// DrawList_AddLine parameter default value hint:
// thickness: 1.0f
func (self ImDrawList) AddLine(p1 ImVec2, p2 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddLine(self.handle(), p1.toC(), p2.toC(), C.ImU32(col), C.float(thickness))
}

// ImageButton parameter default value hint:
// bg_col: ImVec4(0,0,0,0)
// tint_col: ImVec4(1,1,1,1)
// uv0: ImVec2(0,0)
// uv1: ImVec2(1,1)
func ImageButton(str_id string, user_texture_id ImTextureID, size ImVec2, uv0 ImVec2, uv1 ImVec2, bg_col ImVec4, tint_col ImVec4) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ImageButton(str_idArg, C.ImTextureID(user_texture_id), size.toC(), uv0.toC(), uv1.toC(), bg_col.toC(), tint_col.toC()) == C.bool(true)
}

// SliderAngle parameter default value hint:
// flags: 0
// format: "%.0f deg"
// v_degrees_max: +360.0f
// v_degrees_min: -360.0f
func SliderAngle(label string, v_rad *float32, v_degrees_min float32, v_degrees_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_radArg, v_radFin := wrapFloat(v_rad)
	defer v_radFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderAngle(labelArg, v_radArg, C.float(v_degrees_min), C.float(v_degrees_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func GetMousePosOnOpeningCurrentPopup(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMousePosOnOpeningCurrentPopup(pOutArg)
}

func Viewport_GetWorkCenter(pOut *ImVec2, self ImGuiViewport) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Viewport_GetWorkCenter(pOutArg, self.handle())
}

func EndPopup() {
	C.EndPopup()
}

func IsItemDeactivated() bool {
	return C.IsItemDeactivated() == C.bool(true)
}

func GetIO() ImGuiIO {
	return (ImGuiIO)(unsafe.Pointer(C.GetIO()))
}

// SetNextWindowCollapsed parameter default value hint:
// cond: 0
func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond) {
	C.SetNextWindowCollapsed(C.bool(collapsed), C.ImGuiCond(cond))
}

func (self ImGuiTextFilter) IsActive() bool {
	return C.TextFilter_IsActive(self.handle()) == C.bool(true)
}

func NewImFontGlyphRangesBuilder() ImFontGlyphRangesBuilder {
	return (ImFontGlyphRangesBuilder)(unsafe.Pointer(C.FontGlyphRangesBuilder_ImFontGlyphRangesBuilder()))
}

// ListClipper_Begin parameter default value hint:
// items_height: -1.0f
func (self ImGuiListClipper) Begin(items_count int32, items_height float32) {
	C.ListClipper_Begin(self.handle(), C.int(items_count), C.float(items_height))
}

func (self ImGuiIO) AddKeyEvent(key ImGuiKey, down bool) {
	C.IO_AddKeyEvent(self.handle(), C.ImGuiKey(key), C.bool(down))
}

// IsMouseClicked parameter default value hint:
// repeat: false
func IsMouseClicked(button ImGuiMouseButton, repeat bool) bool {
	return C.IsMouseClicked(C.ImGuiMouseButton(button), C.bool(repeat)) == C.bool(true)
}

func (self ImGuiViewport) Destroy() {
	C.Viewport_Destroy(self.handle())
}

func NewImFontAtlas() ImFontAtlas {
	return (ImFontAtlas)(unsafe.Pointer(C.FontAtlas_ImFontAtlas()))
}

// Font_RenderText parameter default value hint:
// cpu_fine_clip: false
// wrap_width: 0.0f
func (self ImFont) RenderText(draw_list ImDrawList, size float32, pos ImVec2, col uint32, clip_rect ImVec4, text_begin string, wrap_width float32, cpu_fine_clip bool) {
	text_beginArg, text_beginFin := wrapString(text_begin)
	defer text_beginFin()

	C.Font_RenderText(self.handle(), draw_list.handle(), C.float(size), pos.toC(), C.ImU32(col), clip_rect.toC(), text_beginArg, C.float(wrap_width), C.bool(cpu_fine_clip))
}

// Begin parameter default value hint:
// flags: 0
// p_open: NULL
func Begin(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.Begin(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) == C.bool(true)
}

func GetID_Str(str_id string) ImGuiID {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return ImGuiID(C.GetID_Str(str_idArg))
}

func GetID_StrStr(str_id_begin string, str_id_end string) ImGuiID {
	str_id_beginArg, str_id_beginFin := wrapString(str_id_begin)
	defer str_id_beginFin()

	str_id_endArg, str_id_endFin := wrapString(str_id_end)
	defer str_id_endFin()

	return ImGuiID(C.GetID_StrStr(str_id_beginArg, str_id_endArg))
}

func GetID_Ptr(ptr_id unsafe.Pointer) ImGuiID {
	return ImGuiID(C.GetID_Ptr(ptr_id))
}

func (self ImGuiTableSortSpecs) Destroy() {
	C.TableSortSpecs_Destroy(self.handle())
}

func (self ImGuiStackTool) Destroy() {
	C.StackTool_Destroy(self.handle())
}

// DrawList_PathArcTo parameter default value hint:
// num_segments: 0
func (self ImDrawList) PathArcTo(center ImVec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	C.DrawList_PathArcTo(self.handle(), center.toC(), C.float(radius), C.float(a_min), C.float(a_max), C.int(num_segments))
}

func (self ImGuiWindowClass) Destroy() {
	C.WindowClass_Destroy(self.handle())
}

func SetDrawCursorPos(local_pos ImVec2) {
	C.SetDrawCursorPos(local_pos.toC())
}

func FindViewportByID(id ImGuiID) ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.FindViewportByID(C.ImGuiID(id))))
}

func GetWindowWidth() float32 {
	return float32(C.GetWindowWidth())
}

// InputInt4 parameter default value hint:
// flags: 0
func InputInt4(label string, v [4]*int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	return C.InputInt4(labelArg, (*C.int)(&vArg[0]), C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

// Indent parameter default value hint:
// indent_w: 0.0f
func Indent(indent_w float32) {
	C.Indent(C.float(indent_w))
}

// DrawList_AddImage parameter default value hint:
// col: 4294967295
// uv_max: ImVec2(1,1)
// uv_min: ImVec2(0,0)
func (self ImDrawList) AddImage(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32) {
	C.DrawList_AddImage(self.handle(), C.ImTextureID(user_texture_id), p_min.toC(), p_max.toC(), uv_min.toC(), uv_max.toC(), C.ImU32(col))
}

func (self ImFontGlyphRangesBuilder) GetBit(n uint64) bool {
	return C.FontGlyphRangesBuilder_GetBit(self.handle(), C.xlong(n)) == C.bool(true)
}

func GetFrameHeightWithSpacing() float32 {
	return float32(C.GetFrameHeightWithSpacing())
}

// TableSetBgColor parameter default value hint:
// column_n: -1
func TableSetBgColor(target ImGuiTableBgTarget, color uint32, column_n int32) {
	C.TableSetBgColor(C.ImGuiTableBgTarget(target), C.ImU32(color), C.int(column_n))
}

func (self ImDrawList) PrimReserve(idx_count int32, vtx_count int32) {
	C.DrawList_PrimReserve(self.handle(), C.int(idx_count), C.int(vtx_count))
}

// AcceptDragDropPayload parameter default value hint:
// flags: 0
func AcceptDragDropPayload(typeArg string, flags ImGuiDragDropFlags) ImGuiPayload {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return (ImGuiPayload)(unsafe.Pointer(C.AcceptDragDropPayload(typeArgArg, C.ImGuiDragDropFlags(flags))))
}

func GetItemRectSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectSize(pOutArg)
}

func GetTime() float64 {
	return float64(C.GetTime())
}

// TabItemButton parameter default value hint:
// flags: 0
func TabItemButton(label string, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TabItemButton(labelArg, C.ImGuiTabItemFlags(flags)) == C.bool(true)
}

func (self ImDrawList) PopClipRect() {
	C.DrawList_PopClipRect(self.handle())
}

func (self ImFontAtlas) CalcCustomRectUV(rect ImFontAtlasCustomRect, out_uv_min *ImVec2, out_uv_max *ImVec2) {
	out_uv_minArg, out_uv_minFin := out_uv_min.wrap()
	defer out_uv_minFin()

	out_uv_maxArg, out_uv_maxFin := out_uv_max.wrap()
	defer out_uv_maxFin()

	C.FontAtlas_CalcCustomRectUV(self.handle(), rect.handle(), out_uv_minArg, out_uv_maxArg)
}

func GetDrawData() ImDrawData {
	return (ImDrawData)(unsafe.Pointer(C.GetDrawData()))
}

func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableHeader(labelArg)
}

func (self ImGuiNavItemData) Destroy() {
	C.NavItemData_Destroy(self.handle())
}

func GetWindowContentRegionMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowContentRegionMax(pOutArg)
}

func (self ImDrawListSplitter) Destroy() {
	C.DrawListSplitter_Destroy(self.handle())
}

func (self ImGuiOldColumns) Destroy() {
	C.OldColumns_Destroy(self.handle())
}

func Separator() {
	C.Separator()
}

func TableGetColumnCount() int {
	return int(C.TableGetColumnCount())
}

// LoadIniSettingsFromMemory parameter default value hint:
// ini_size: 0
func LoadIniSettingsFromMemory(ini_data string, ini_size uint64) {
	ini_dataArg, ini_dataFin := wrapString(ini_data)
	defer ini_dataFin()

	C.LoadIniSettingsFromMemory(ini_dataArg, C.xlong(ini_size))
}

func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	C.SetNextFrameWantCaptureKeyboard(C.bool(want_capture_keyboard))
}

// InputFloat3 parameter default value hint:
// flags: 0
// format: "%.3f"
func InputFloat3(label string, v [3]*float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat3(labelArg, (*C.float)(&vArg[0]), formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func (self ImGuiInputTextState) Destroy() {
	C.InputTextState_Destroy(self.handle())
}

func (self ImGuiTableSettings) Destroy() {
	C.TableSettings_Destroy(self.handle())
}

func GetWindowDpiScale() float32 {
	return float32(C.GetWindowDpiScale())
}

func (self ImDrawList) Destroy() {
	C.DrawList_Destroy(self.handle())
}

func IsAnyItemHovered() bool {
	return C.IsAnyItemHovered() == C.bool(true)
}

func ArrowButton(str_id string, dir ImGuiDir) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ArrowButton(str_idArg, C.ImGuiDir(dir)) == C.bool(true)
}

// Color_SetHSV parameter default value hint:
// a: 1.0f
func (self *ImColor) SetHSV(h float32, s float32, v float32, a float32) {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.Color_SetHSV(selfArg, C.float(h), C.float(s), C.float(v), C.float(a))
}

func (self ImFontAtlas) GetGlyphRangesChineseFull() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesChineseFull(self.handle()))
}

func (self ImFont) IsGlyphRangeUnused(c_begin uint32, c_last uint32) bool {
	return C.Font_IsGlyphRangeUnused(self.handle(), C.uint(c_begin), C.uint(c_last)) == C.bool(true)
}

func DebugTextEncoding(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.DebugTextEncoding(textArg)
}

func SetMouseCursor(cursor_type ImGuiMouseCursor) {
	C.SetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

func (self ImDrawList) AddRectFilledMultiColor(p_min ImVec2, p_max ImVec2, col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	C.DrawList_AddRectFilledMultiColor(self.handle(), p_min.toC(), p_max.toC(), C.ImU32(col_upr_left), C.ImU32(col_upr_right), C.ImU32(col_bot_right), C.ImU32(col_bot_left))
}

func (self ImFont) FindGlyph(c ImWchar) ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.Font_FindGlyph(self.handle(), C.ImWchar(c))))
}

// BeginPopupContextItem parameter default value hint:
// str_id: NULL
// popup_flags: 1
func BeginPopupContextItem(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextItem(str_idArg, C.ImGuiPopupFlags(popup_flags)) == C.bool(true)
}

// DockSpace parameter default value hint:
// flags: 0
// size: ImVec2(0,0)
// window_class: NULL
func DockSpace(id ImGuiID, size ImVec2, flags ImGuiDockNodeFlags, window_class ImGuiWindowClass) ImGuiID {
	return ImGuiID(C.DockSpace(C.ImGuiID(id), size.toC(), C.ImGuiDockNodeFlags(flags), window_class.handle()))
}

func SetItemDefaultFocus() {
	C.SetItemDefaultFocus()
}

func (self ImFontAtlas) SetTexID(id ImTextureID) {
	C.FontAtlas_SetTexID(self.handle(), C.ImTextureID(id))
}

func NewImGuiPlatformMonitor() ImGuiPlatformMonitor {
	return (ImGuiPlatformMonitor)(unsafe.Pointer(C.PlatformMonitor_ImGuiPlatformMonitor()))
}

func NextColumn() {
	C.NextColumn()
}

// StyleColorsDark parameter default value hint:
// dst: NULL
func StyleColorsDark(dst ImGuiStyle) {
	C.StyleColorsDark(dst.handle())
}

// Unindent parameter default value hint:
// indent_w: 0.0f
func Unindent(indent_w float32) {
	C.Unindent(C.float(indent_w))
}

func IsKeyReleased(key ImGuiKey) bool {
	return C.IsKeyReleased(C.ImGuiKey(key)) == C.bool(true)
}

func SetNextWindowBgAlpha(alpha float32) {
	C.SetNextWindowBgAlpha(C.float(alpha))
}

// SetScrollFromPosY_Float parameter default value hint:
// center_y_ratio: 0.5f
func SetScrollFromPosY_Float(local_y float32, center_y_ratio float32) {
	C.SetScrollFromPosY_Float(C.float(local_y), C.float(center_y_ratio))
}

func (self ImGuiTableColumnSettings) Destroy() {
	C.TableColumnSettings_Destroy(self.handle())
}

func GetForegroundDrawList_Nil() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetForegroundDrawList_Nil()))
}

func GetForegroundDrawList_ViewportPtr(viewport ImGuiViewport) ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetForegroundDrawList_ViewportPtr(viewport.handle())))
}

// SliderFloat3 parameter default value hint:
// flags: 0
// format: "%.3f"
func SliderFloat3(label string, v [3]*float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat3(labelArg, (*C.float)(&vArg[0]), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// TableGetColumnFlags parameter default value hint:
// column_n: -1
func TableGetColumnFlags(column_n int32) ImGuiTableColumnFlags {
	return ImGuiTableColumnFlags(C.TableGetColumnFlags(C.int(column_n)))
}

func MemFree(ptr unsafe.Pointer) {
	C.MemFree(ptr)
}

func (self ImDrawList) AddTriangleFilled(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32) {
	C.DrawList_AddTriangleFilled(self.handle(), p1.toC(), p2.toC(), p3.toC(), C.ImU32(col))
}

// ColorPicker4 parameter default value hint:
// flags: 0
// ref_col: NULL
func ColorPicker4(label string, col [4]*float32, flags ImGuiColorEditFlags, ref_col []float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := make([]C.float, len(col))
	for i, colV := range col {
		colArg[i] = C.float(*colV)
	}
	defer func() {
		for i, colV := range colArg {
			*col[i] = float32(colV)
		}
	}()

	return C.ColorPicker4(labelArg, (*C.float)(&colArg[0]), C.ImGuiColorEditFlags(flags), (*C.float)(&(ref_col[0]))) == C.bool(true)
}

func SetDrawCursorPosY(local_y float32) {
	C.SetDrawCursorPosY(C.float(local_y))
}

// SliderScalar parameter default value hint:
// flags: 0
// format: NULL
func SliderScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImDrawList) ChannelsSetCurrent(n int32) {
	C.DrawList_ChannelsSetCurrent(self.handle(), C.int(n))
}

func (self ImGuiStyleMod) Destroy() {
	C.StyleMod_Destroy(self.handle())
}

func NewImGuiTextBuffer() ImGuiTextBuffer {
	return (ImGuiTextBuffer)(unsafe.Pointer(C.TextBuffer_ImGuiTextBuffer()))
}

// Image parameter default value hint:
// border_col: ImVec4(0,0,0,0)
// tint_col: ImVec4(1,1,1,1)
// uv0: ImVec2(0,0)
// uv1: ImVec2(1,1)
func Image(user_texture_id ImTextureID, size ImVec2, uv0 ImVec2, uv1 ImVec2, tint_col ImVec4, border_col ImVec4) {
	C.Image(C.ImTextureID(user_texture_id), size.toC(), uv0.toC(), uv1.toC(), tint_col.toC(), border_col.toC())
}

// InputFloat4 parameter default value hint:
// flags: 0
// format: "%.3f"
func InputFloat4(label string, v [4]*float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat4(labelArg, (*C.float)(&vArg[0]), formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func SetNextWindowContentSize(size ImVec2) {
	C.SetNextWindowContentSize(size.toC())
}

func GetWindowHeight() float32 {
	return float32(C.GetWindowHeight())
}

func NewImGuiStyle() ImGuiStyle {
	return (ImGuiStyle)(unsafe.Pointer(C.Style_ImGuiStyle()))
}

func (self ImDrawList) PathArcToFast(center ImVec2, radius float32, a_min_of_12 int32, a_max_of_12 int32) {
	C.DrawList_PathArcToFast(self.handle(), center.toC(), C.float(radius), C.int(a_min_of_12), C.int(a_max_of_12))
}

func (self ImGuiStyle) ScaleAllSizes(scale_factor float32) {
	C.Style_ScaleAllSizes(self.handle(), C.float(scale_factor))
}

func GetDrawCursorStartPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetDrawCursorStartPos(pOutArg)
}

// InputDouble parameter default value hint:
// step_fast: 0.0
// flags: 0
// format: "%.6f"
// step: 0.0
func InputDouble(label string, v *float64, step float64, step_fast float64, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputDouble(labelArg, (*C.double)(v), C.double(step), C.double(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func (self ImGuiIO) AddInputCharacter(c uint32) {
	C.IO_AddInputCharacter(self.handle(), C.uint(c))
}

func (self *ImVec4) Destroy() {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.Vec4_Destroy(selfArg)
}

func Viewport_GetCenter(pOut *ImVec2, self ImGuiViewport) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Viewport_GetCenter(pOutArg, self.handle())
}

// DragFloat3 parameter default value hint:
// format: "%.3f"
// v_max: 0.0f
// v_min: 0.0f
// v_speed: 1.0f
// flags: 0
func DragFloat3(label string, v [3]*float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat3(labelArg, (*C.float)(&vArg[0]), C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func GetCurrentContext() ImGuiContext {
	return (ImGuiContext)(unsafe.Pointer(C.GetCurrentContext()))
}

func IsRectVisible_Nil(size ImVec2) bool {
	return C.IsRectVisible_Nil(size.toC()) == C.bool(true)
}

func IsRectVisible_Vec2(rect_min ImVec2, rect_max ImVec2) bool {
	return C.IsRectVisible_Vec2(rect_min.toC(), rect_max.toC()) == C.bool(true)
}

func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	C.SetNextFrameWantCaptureMouse(C.bool(want_capture_mouse))
}

// SetWindowCollapsed_Bool parameter default value hint:
// cond: 0
func SetWindowCollapsed_Bool(collapsed bool, cond ImGuiCond) {
	C.SetWindowCollapsed_Bool(C.bool(collapsed), C.ImGuiCond(cond))
}

// SetWindowCollapsed_Str parameter default value hint:
// cond: 0
func SetWindowCollapsed_Str(name string, collapsed bool, cond ImGuiCond) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.SetWindowCollapsed_Str(nameArg, C.bool(collapsed), C.ImGuiCond(cond))
}

// DrawList_PathBezierQuadraticCurveTo parameter default value hint:
// num_segments: 0
func (self ImDrawList) PathBezierQuadraticCurveTo(p2 ImVec2, p3 ImVec2, num_segments int32) {
	C.DrawList_PathBezierQuadraticCurveTo(self.handle(), p2.toC(), p3.toC(), C.int(num_segments))
}

func CheckboxFlags_IntPtr(label string, flags *int32, flags_value int32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	flagsArg, flagsFin := wrapInt32(flags)
	defer flagsFin()

	return C.CheckboxFlags_IntPtr(labelArg, flagsArg, C.int(flags_value)) == C.bool(true)
}

func GetTreeNodeToLabelSpacing() float32 {
	return float32(C.GetTreeNodeToLabelSpacing())
}

func (self ImGuiListClipperData) Destroy() {
	C.ListClipperData_Destroy(self.handle())
}

func GetItemRectMin(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectMin(pOutArg)
}

func (self ImDrawList) PathClear() {
	C.DrawList_PathClear(self.handle())
}

func (self ImDrawList) PushTextureID(texture_id ImTextureID) {
	C.DrawList_PushTextureID(self.handle(), C.ImTextureID(texture_id))
}

func (self ImFontAtlas) ClearInputData() {
	C.FontAtlas_ClearInputData(self.handle())
}

func (self ImDrawList) PrimRectUV(a ImVec2, b ImVec2, uv_a ImVec2, uv_b ImVec2, col uint32) {
	C.DrawList_PrimRectUV(self.handle(), a.toC(), b.toC(), uv_a.toC(), uv_b.toC(), C.ImU32(col))
}

func GetScrollMaxY() float32 {
	return float32(C.GetScrollMaxY())
}

// DragScalar parameter default value hint:
// p_min: NULL
// v_speed: 1.0f
// flags: 0
// format: NULL
// p_max: NULL
func DragScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalar(labelArg, C.ImGuiDataType(data_type), p_data, C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func EndTable() {
	C.EndTable()
}

// LogToFile parameter default value hint:
// filename: NULL
// auto_open_depth: -1
func LogToFile(auto_open_depth int32, filename string) {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	C.LogToFile(C.int(auto_open_depth), filenameArg)
}

func SetWindowFocus_Nil() {
	C.SetWindowFocus_Nil()
}

func SetWindowFocus_Str(name string) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.SetWindowFocus_Str(nameArg)
}

func NewImGuiInputTextCallbackData() ImGuiInputTextCallbackData {
	return (ImGuiInputTextCallbackData)(unsafe.Pointer(C.InputTextCallbackData_ImGuiInputTextCallbackData()))
}

func LogButtons() {
	C.LogButtons()
}

// Color_HSV parameter default value hint:
// a: 1.0f
func Color_HSV(pOut *ImColor, h float32, s float32, v float32, a float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Color_HSV(pOutArg, C.float(h), C.float(s), C.float(v), C.float(a))
}

func (self ImFont) RenderChar(draw_list ImDrawList, size float32, pos ImVec2, col uint32, c ImWchar) {
	C.Font_RenderChar(self.handle(), draw_list.handle(), C.float(size), pos.toC(), C.ImU32(col), C.ImWchar(c))
}

func PopClipRect() {
	C.PopClipRect()
}

func EndDragDropSource() {
	C.EndDragDropSource()
}

func GetMouseCursor() ImGuiMouseCursor {
	return ImGuiMouseCursor(C.GetMouseCursor())
}

// IsItemClicked parameter default value hint:
// mouse_button: 0
func IsItemClicked(mouse_button ImGuiMouseButton) bool {
	return C.IsItemClicked(C.ImGuiMouseButton(mouse_button)) == C.bool(true)
}

// DragInt4 parameter default value hint:
// flags: 0
// format: "%d"
// v_max: 0
// v_min: 0
// v_speed: 1.0f
func DragInt4(label string, v [4]*int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt4(labelArg, (*C.int)(&vArg[0]), C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImFont) FindGlyphNoFallback(c ImWchar) ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.Font_FindGlyphNoFallback(self.handle(), C.ImWchar(c))))
}

func NewImGuiViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.Viewport_ImGuiViewport()))
}

// LogToClipboard parameter default value hint:
// auto_open_depth: -1
func LogToClipboard(auto_open_depth int32) {
	C.LogToClipboard(C.int(auto_open_depth))
}

func SetScrollY_Float(scroll_y float32) {
	C.SetScrollY_Float(C.float(scroll_y))
}

func ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ShowStyleSelector(labelArg) == C.bool(true)
}

func (self ImGuiIO) AddMousePosEvent(x float32, y float32) {
	C.IO_AddMousePosEvent(self.handle(), C.float(x), C.float(y))
}

func (self ImGuiIO) AddMouseViewportEvent(id ImGuiID) {
	C.IO_AddMouseViewportEvent(self.handle(), C.ImGuiID(id))
}

// SliderInt parameter default value hint:
// flags: 0
// format: "%d"
func SliderInt(label string, v *int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImDrawList) CloneOutput() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.DrawList_CloneOutput(self.handle())))
}

func GetDrawCursorPosX() float32 {
	return float32(C.GetDrawCursorPosX())
}

// DrawList_AddText_Vec2 parameter default value hint:
// text_end: NULL
func (self ImDrawList) AddText_Vec2(pos ImVec2, col uint32, text_begin string) {
	text_beginArg, text_beginFin := wrapString(text_begin)
	defer text_beginFin()

	C.DrawList_AddText_Vec2(self.handle(), pos.toC(), C.ImU32(col), text_beginArg)
}

// DrawList_AddText_FontPtr parameter default value hint:
// cpu_fine_clip_rect: NULL
// text_end: NULL
// wrap_width: 0.0f
func (self ImDrawList) AddText_FontPtr(font ImFont, font_size float32, pos ImVec2, col uint32, text_begin string, wrap_width float32, cpu_fine_clip_rect *ImVec4) {
	text_beginArg, text_beginFin := wrapString(text_begin)
	defer text_beginFin()

	cpu_fine_clip_rectArg, cpu_fine_clip_rectFin := cpu_fine_clip_rect.wrap()
	defer cpu_fine_clip_rectFin()

	C.DrawList_AddText_FontPtr(self.handle(), font.handle(), C.float(font_size), pos.toC(), C.ImU32(col), text_beginArg, C.float(wrap_width), cpu_fine_clip_rectArg)
}

func (self ImGuiPayload) IsDelivery() bool {
	return C.Payload_IsDelivery(self.handle()) == C.bool(true)
}

// DragIntRange2 parameter default value hint:
// flags: 0
// format: "%d"
// format_max: NULL
// v_max: 0
// v_min: 0
// v_speed: 1.0f
func DragIntRange2(label string, v_current_min *int32, v_current_max *int32, v_speed float32, v_min int32, v_max int32, format string, format_max string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_current_minArg, v_current_minFin := wrapInt32(v_current_min)
	defer v_current_minFin()

	v_current_maxArg, v_current_maxFin := wrapInt32(v_current_max)
	defer v_current_maxFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	format_maxArg, format_maxFin := wrapString(format_max)
	defer format_maxFin()

	return C.DragIntRange2(labelArg, v_current_minArg, v_current_maxArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, format_maxArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImDrawList) AddPolyline(points *ImVec2, num_points int32, col uint32, flags ImDrawFlags, thickness float32) {
	pointsArg, pointsFin := points.wrap()
	defer pointsFin()

	C.DrawList_AddPolyline(self.handle(), pointsArg, C.int(num_points), C.ImU32(col), C.ImDrawFlags(flags), C.float(thickness))
}

func (self ImDrawList) PrimWriteIdx(idx ImDrawIdx) {
	C.DrawList_PrimWriteIdx(self.handle(), C.ImDrawIdx(idx))
}

func (self ImGuiIO) AddMouseWheelEvent(wh_x float32, wh_y float32) {
	C.IO_AddMouseWheelEvent(self.handle(), C.float(wh_x), C.float(wh_y))
}

func (self ImGuiTextBuffer) Empty() bool {
	return C.TextBuffer_Empty(self.handle()) == C.bool(true)
}

// IsMouseDragging parameter default value hint:
// lock_threshold: -1.0f
func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32) bool {
	return C.IsMouseDragging(C.ImGuiMouseButton(button), C.float(lock_threshold)) == C.bool(true)
}

func (self ImGuiTabItem) Destroy() {
	C.TabItem_Destroy(self.handle())
}

func SetClipboardText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.SetClipboardText(textArg)
}

func SetDrawCursorScreenPos(pos ImVec2) {
	C.SetDrawCursorScreenPos(pos.toC())
}

func (self ImGuiOldColumnData) Destroy() {
	C.OldColumnData_Destroy(self.handle())
}

func NewImGuiPlatformImeData() ImGuiPlatformImeData {
	return (ImGuiPlatformImeData)(unsafe.Pointer(C.PlatformImeData_ImGuiPlatformImeData()))
}

// Button parameter default value hint:
// size: ImVec2(0,0)
func Button(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Button(labelArg, size.toC()) == C.bool(true)
}

func SetColorEditOptions(flags ImGuiColorEditFlags) {
	C.SetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

func DrawList_GetClipRectMax(pOut *ImVec2, self ImDrawList) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.DrawList_GetClipRectMax(pOutArg, self.handle())
}

func NewImFontConfig() ImFontConfig {
	return (ImFontConfig)(unsafe.Pointer(C.FontConfig_ImFontConfig()))
}

func (self ImGuiTextBuffer) Appendf(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.TextBuffer_Appendf(self.handle(), fmtArg)
}

// SetNextItemOpen parameter default value hint:
// cond: 0
func SetNextItemOpen(is_open bool, cond ImGuiCond) {
	C.SetNextItemOpen(C.bool(is_open), C.ImGuiCond(cond))
}

func (self ImDrawListSplitter) ClearFreeMemory() {
	C.DrawListSplitter_ClearFreeMemory(self.handle())
}

func GetFontSize() float32 {
	return float32(C.GetFontSize())
}

func (self ImDrawList) AddNgonFilled(center ImVec2, radius float32, col uint32, num_segments int32) {
	C.DrawList_AddNgonFilled(self.handle(), center.toC(), C.float(radius), C.ImU32(col), C.int(num_segments))
}

func (self ImGuiMenuColumns) Destroy() {
	C.MenuColumns_Destroy(self.handle())
}

func GetFrameCount() int {
	return int(C.GetFrameCount())
}

// DrawList_PathRect parameter default value hint:
// flags: 0
// rounding: 0.0f
func (self ImDrawList) PathRect(rect_min ImVec2, rect_max ImVec2, rounding float32, flags ImDrawFlags) {
	C.DrawList_PathRect(self.handle(), rect_min.toC(), rect_max.toC(), C.float(rounding), C.ImDrawFlags(flags))
}

func (self ImGuiTextBuffer) Destroy() {
	C.TextBuffer_Destroy(self.handle())
}

// IsMouseHoveringRect parameter default value hint:
// clip: true
func IsMouseHoveringRect(r_min ImVec2, r_max ImVec2, clip bool) bool {
	return C.IsMouseHoveringRect(r_min.toC(), r_max.toC(), C.bool(clip)) == C.bool(true)
}

func IsItemDeactivatedAfterEdit() bool {
	return C.IsItemDeactivatedAfterEdit() == C.bool(true)
}

func (self ImFontAtlas) AddFont(font_cfg ImFontConfig) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFont(self.handle(), font_cfg.handle())))
}

// BeginDisabled parameter default value hint:
// disabled: true
func BeginDisabled(disabled bool) {
	C.BeginDisabled(C.bool(disabled))
}

func NewLine() {
	C.NewLine()
}

func (self ImGuiIO) SetAppAcceptingEvents(accepting_events bool) {
	C.IO_SetAppAcceptingEvents(self.handle(), C.bool(accepting_events))
}

// BeginTable parameter default value hint:
// flags: 0
// inner_width: 0.0f
// outer_size: ImVec2(0.0f,0.0f)
func BeginTable(str_id string, column int32, flags ImGuiTableFlags, outer_size ImVec2, inner_width float32) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTable(str_idArg, C.int(column), C.ImGuiTableFlags(flags), outer_size.toC(), C.float(inner_width)) == C.bool(true)
}

func (self ImGuiContext) Destroy() {
	C.Context_Destroy(self.handle())
}

func (self ImGuiWindowSettings) Destroy() {
	C.WindowSettings_Destroy(self.handle())
}

// IsItemHovered parameter default value hint:
// flags: 0
func IsItemHovered(flags ImGuiHoveredFlags) bool {
	return C.IsItemHovered(C.ImGuiHoveredFlags(flags)) == C.bool(true)
}

func (self ImGuiLastItemData) Destroy() {
	C.LastItemData_Destroy(self.handle())
}

func CloseCurrentPopup() {
	C.CloseCurrentPopup()
}

func (self ImFontAtlas) Clear() {
	C.FontAtlas_Clear(self.handle())
}

func (self ImFont) ClearOutputData() {
	C.Font_ClearOutputData(self.handle())
}

// IO_SetKeyEventNativeData parameter default value hint:
// native_legacy_index: -1
func (self ImGuiIO) SetKeyEventNativeData(key ImGuiKey, native_keycode int32, native_scancode int32, native_legacy_index int32) {
	C.IO_SetKeyEventNativeData(self.handle(), C.ImGuiKey(key), C.int(native_keycode), C.int(native_scancode), C.int(native_legacy_index))
}

func GetStyleColorVec4(idx ImGuiCol) ImVec4 {
	return newImVec4FromCPtr(C.GetStyleColorVec4(C.ImGuiCol(idx)))
}

func (self ImFontAtlas) IsBuilt() bool {
	return C.FontAtlas_IsBuilt(self.handle()) == C.bool(true)
}

func (self ImFontGlyphRangesBuilder) AddRanges(ranges *ImWchar) {
	C.FontGlyphRangesBuilder_AddRanges(self.handle(), (*C.ImWchar)(ranges))
}

func IsItemActive() bool {
	return C.IsItemActive() == C.bool(true)
}

// DragScalarN parameter default value hint:
// format: NULL
// p_max: NULL
// p_min: NULL
// v_speed: 1.0f
// flags: 0
func DragScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func GetClipboardText() string {
	return C.GoString(C.GetClipboardText())
}

// ColorButton parameter default value hint:
// flags: 0
// size: ImVec2(0,0)
func ColorButton(desc_id string, col ImVec4, flags ImGuiColorEditFlags, size ImVec2) bool {
	desc_idArg, desc_idFin := wrapString(desc_id)
	defer desc_idFin()

	return C.ColorButton(desc_idArg, col.toC(), C.ImGuiColorEditFlags(flags), size.toC()) == C.bool(true)
}

func (self ImDrawListSplitter) Merge(draw_list ImDrawList) {
	C.DrawListSplitter_Merge(self.handle(), draw_list.handle())
}

func NewImFontAtlasCustomRect() ImFontAtlasCustomRect {
	return (ImFontAtlasCustomRect)(unsafe.Pointer(C.FontAtlasCustomRect_ImFontAtlasCustomRect()))
}

func (self ImDrawListSharedData) Destroy() {
	C.DrawListSharedData_Destroy(self.handle())
}

func (self ImDrawList) ChannelsSplit(count int32) {
	C.DrawList_ChannelsSplit(self.handle(), C.int(count))
}

// InputFloat2 parameter default value hint:
// format: "%.3f"
// flags: 0
func InputFloat2(label string, v [2]*float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat2(labelArg, (*C.float)(&vArg[0]), formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

// SetNextWindowPos parameter default value hint:
// cond: 0
// pivot: ImVec2(0,0)
func SetNextWindowPos(pos ImVec2, cond ImGuiCond, pivot ImVec2) {
	C.SetNextWindowPos(pos.toC(), C.ImGuiCond(cond), pivot.toC())
}

func (self ImGuiPopupData) Destroy() {
	C.PopupData_Destroy(self.handle())
}

// TextFilter_Draw parameter default value hint:
// label: "Filter(inc,-exc)"
// width: 0.0f
func (self ImGuiTextFilter) Draw(label string, width float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TextFilter_Draw(self.handle(), labelArg, C.float(width)) == C.bool(true)
}

// SetNextWindowDockID parameter default value hint:
// cond: 0
func SetNextWindowDockID(dock_id ImGuiID, cond ImGuiCond) {
	C.SetNextWindowDockID(C.ImGuiID(dock_id), C.ImGuiCond(cond))
}

func (self ImFont) IsLoaded() bool {
	return C.Font_IsLoaded(self.handle()) == C.bool(true)
}

func NewImGuiTextFilter(default_filter string) ImGuiTextFilter {
	default_filterArg, default_filterFin := wrapString(default_filter)
	defer default_filterFin()

	return (ImGuiTextFilter)(unsafe.Pointer(C.TextFilter_ImGuiTextFilter(default_filterArg)))
}

func IsWindowCollapsed() bool {
	return C.IsWindowCollapsed() == C.bool(true)
}

// SetWindowSize_Vec2 parameter default value hint:
// cond: 0
func SetWindowSize_Vec2(size ImVec2, cond ImGuiCond) {
	C.SetWindowSize_Vec2(size.toC(), C.ImGuiCond(cond))
}

// SetWindowSize_Str parameter default value hint:
// cond: 0
func SetWindowSize_Str(name string, size ImVec2, cond ImGuiCond) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.SetWindowSize_Str(nameArg, size.toC(), C.ImGuiCond(cond))
}

func (self ImDrawList) PrimWriteVtx(pos ImVec2, uv ImVec2, col uint32) {
	C.DrawList_PrimWriteVtx(self.handle(), pos.toC(), uv.toC(), C.ImU32(col))
}

func (self ImGuiContextHook) Destroy() {
	C.ContextHook_Destroy(self.handle())
}

func LabelText(label string, fmt string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.LabelText(labelArg, fmtArg)
}

func PopItemWidth() {
	C.PopItemWidth()
}

func (self ImGuiIO) ClearInputKeys() {
	C.IO_ClearInputKeys(self.handle())
}

func (self ImGuiTextBuffer) Clear() {
	C.TextBuffer_Clear(self.handle())
}

func GetDrawCursorPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetDrawCursorPos(pOutArg)
}

// InputInt parameter default value hint:
// flags: 0
// step: 1
// step_fast: 100
func InputInt(label string, v *int32, step int32, step_fast int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	return C.InputInt(labelArg, vArg, C.int(step), C.int(step_fast), C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func TextColored(col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.TextColored(col.toC(), fmtArg)
}

func (self ImFont) GrowIndex(new_size int32) {
	C.Font_GrowIndex(self.handle(), C.int(new_size))
}

// SetScrollFromPosX_Float parameter default value hint:
// center_x_ratio: 0.5f
func SetScrollFromPosX_Float(local_x float32, center_x_ratio float32) {
	C.SetScrollFromPosX_Float(C.float(local_x), C.float(center_x_ratio))
}

func ColorConvertHSVtoRGB(h float32, s float32, v float32, out_r *float32, out_g *float32, out_b *float32) {
	out_rArg, out_rFin := wrapFloat(out_r)
	defer out_rFin()

	out_gArg, out_gFin := wrapFloat(out_g)
	defer out_gFin()

	out_bArg, out_bFin := wrapFloat(out_b)
	defer out_bFin()

	C.ColorConvertHSVtoRGB(C.float(h), C.float(s), C.float(v), out_rArg, out_gArg, out_bArg)
}

// ShowStyleEditor parameter default value hint:
// ref: NULL
func ShowStyleEditor(ref ImGuiStyle) {
	C.ShowStyleEditor(ref.handle())
}

func (self ImGuiIO) Destroy() {
	C.IO_Destroy(self.handle())
}

func NewImGuiPayload() ImGuiPayload {
	return (ImGuiPayload)(unsafe.Pointer(C.Payload_ImGuiPayload()))
}

func BeginMainMenuBar() bool {
	return C.BeginMainMenuBar() == C.bool(true)
}

// OpenPopupOnItemClick parameter default value hint:
// popup_flags: 1
// str_id: NULL
func OpenPopupOnItemClick(str_id string, popup_flags ImGuiPopupFlags) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.OpenPopupOnItemClick(str_idArg, C.ImGuiPopupFlags(popup_flags))
}

// SetKeyboardFocusHere parameter default value hint:
// offset: 0
func SetKeyboardFocusHere(offset int32) {
	C.SetKeyboardFocusHere(C.int(offset))
}

func (self ImDrawList) PrimVtx(pos ImVec2, uv ImVec2, col uint32) {
	C.DrawList_PrimVtx(self.handle(), pos.toC(), uv.toC(), C.ImU32(col))
}

func (self ImGuiDockContext) Destroy() {
	C.DockContext_Destroy(self.handle())
}

// BeginTabBar parameter default value hint:
// flags: 0
func BeginTabBar(str_id string, flags ImGuiTabBarFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTabBar(str_idArg, C.ImGuiTabBarFlags(flags)) == C.bool(true)
}

func GetFrameHeight() float32 {
	return float32(C.GetFrameHeight())
}

// DrawList_AddRectFilled parameter default value hint:
// flags: 0
// rounding: 0.0f
func (self ImDrawList) AddRectFilled(p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.DrawList_AddRectFilled(self.handle(), p_min.toC(), p_max.toC(), C.ImU32(col), C.float(rounding), C.ImDrawFlags(flags))
}

func (self ImGuiIO) ClearInputCharacters() {
	C.IO_ClearInputCharacters(self.handle())
}

func (self ImGuiTextBuffer) Reserve(capacity int32) {
	C.TextBuffer_Reserve(self.handle(), C.int(capacity))
}

func SetNextItemWidth(item_width float32) {
	C.SetNextItemWidth(C.float(item_width))
}

func TextWrapped(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.TextWrapped(fmtArg)
}

// FontAtlas_AddFontDefault parameter default value hint:
// font_cfg: NULL
func (self ImFontAtlas) AddFontDefault(font_cfg ImFontConfig) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontDefault(self.handle(), font_cfg.handle())))
}

// ColorEdit4 parameter default value hint:
// flags: 0
func ColorEdit4(label string, col [4]*float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := make([]C.float, len(col))
	for i, colV := range col {
		colArg[i] = C.float(*colV)
	}
	defer func() {
		for i, colV := range colArg {
			*col[i] = float32(colV)
		}
	}()

	return C.ColorEdit4(labelArg, (*C.float)(&colArg[0]), C.ImGuiColorEditFlags(flags)) == C.bool(true)
}

func (self ImDrawData) Clear() {
	C.DrawData_Clear(self.handle())
}

// DragFloat4 parameter default value hint:
// flags: 0
// format: "%.3f"
// v_max: 0.0f
// v_min: 0.0f
// v_speed: 1.0f
func DragFloat4(label string, v [4]*float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat4(labelArg, (*C.float)(&vArg[0]), C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self *ImVec2) Destroy() {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.Vec2_Destroy(selfArg)
}

// SameLine parameter default value hint:
// offset_from_start_x: 0.0f
// spacing: -1.0f
func SameLine(offset_from_start_x float32, spacing float32) {
	C.SameLine(C.float(offset_from_start_x), C.float(spacing))
}

// InputInt2 parameter default value hint:
// flags: 0
func InputInt2(label string, v [2]*int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	return C.InputInt2(labelArg, (*C.int)(&vArg[0]), C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func (self ImFontAtlas) GetCustomRectByIndex(index int32) ImFontAtlasCustomRect {
	return (ImFontAtlasCustomRect)(unsafe.Pointer(C.FontAtlas_GetCustomRectByIndex(self.handle(), C.int(index))))
}

func EndListBox() {
	C.EndListBox()
}

func (self ImFont) SetGlyphVisible(c ImWchar, visible bool) {
	C.Font_SetGlyphVisible(self.handle(), C.ImWchar(c), C.bool(visible))
}

func (self ImFont) GetCharAdvance(c ImWchar) float32 {
	return float32(C.Font_GetCharAdvance(self.handle(), C.ImWchar(c)))
}

func (self ImGuiTableInstanceData) Destroy() {
	C.TableInstanceData_Destroy(self.handle())
}

func (self ImDrawListSplitter) SetCurrentChannel(draw_list ImDrawList, channel_idx int32) {
	C.DrawListSplitter_SetCurrentChannel(self.handle(), draw_list.handle(), C.int(channel_idx))
}

func (self ImGuiTableColumn) Destroy() {
	C.TableColumn_Destroy(self.handle())
}

// BeginPopupModal parameter default value hint:
// flags: 0
// p_open: NULL
func BeginPopupModal(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginPopupModal(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) == C.bool(true)
}

func GetKeyPressedAmount(key ImGuiKey, repeat_delay float32, rate float32) int {
	return int(C.GetKeyPressedAmount(C.ImGuiKey(key), C.float(repeat_delay), C.float(rate)))
}

func (self ImDrawListSplitter) Clear() {
	C.DrawListSplitter_Clear(self.handle())
}

func (self ImFontGlyphRangesBuilder) Destroy() {
	C.FontGlyphRangesBuilder_Destroy(self.handle())
}

// BeginPopupContextVoid parameter default value hint:
// popup_flags: 1
// str_id: NULL
func BeginPopupContextVoid(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextVoid(str_idArg, C.ImGuiPopupFlags(popup_flags)) == C.bool(true)
}

// BeginTabItem parameter default value hint:
// p_open: NULL
// flags: 0
func BeginTabItem(label string, p_open *bool, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginTabItem(labelArg, p_openArg, C.ImGuiTabItemFlags(flags)) == C.bool(true)
}

// DragFloatRange2 parameter default value hint:
// v_speed: 1.0f
// flags: 0
// format: "%.3f"
// format_max: NULL
// v_max: 0.0f
// v_min: 0.0f
func DragFloatRange2(label string, v_current_min *float32, v_current_max *float32, v_speed float32, v_min float32, v_max float32, format string, format_max string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_current_minArg, v_current_minFin := wrapFloat(v_current_min)
	defer v_current_minFin()

	v_current_maxArg, v_current_maxFin := wrapFloat(v_current_max)
	defer v_current_maxFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	format_maxArg, format_maxFin := wrapString(format_max)
	defer format_maxFin()

	return C.DragFloatRange2(labelArg, v_current_minArg, v_current_maxArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, format_maxArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// ColorPicker3 parameter default value hint:
// flags: 0
func ColorPicker3(label string, col [3]*float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := make([]C.float, len(col))
	for i, colV := range col {
		colArg[i] = C.float(*colV)
	}
	defer func() {
		for i, colV := range colArg {
			*col[i] = float32(colV)
		}
	}()

	return C.ColorPicker3(labelArg, (*C.float)(&colArg[0]), C.ImGuiColorEditFlags(flags)) == C.bool(true)
}

func GetKeyName(key ImGuiKey) string {
	return C.GoString(C.GetKeyName(C.ImGuiKey(key)))
}

func (self *ImColor) Destroy() {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.Color_Destroy(selfArg)
}

func (self ImDrawListSplitter) Split(draw_list ImDrawList, count int32) {
	C.DrawListSplitter_Split(self.handle(), draw_list.handle(), C.int(count))
}

func (self ImFont) GetDebugName() string {
	return C.GoString(C.Font_GetDebugName(self.handle()))
}

func UpdatePlatformWindows() {
	C.UpdatePlatformWindows()
}

func SetColumnOffset(column_index int32, offset_x float32) {
	C.SetColumnOffset(C.int(column_index), C.float(offset_x))
}

func GetFontTexUvWhitePixel(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetFontTexUvWhitePixel(pOutArg)
}

// GetMouseDragDelta parameter default value hint:
// button: 0
// lock_threshold: -1.0f
func GetMouseDragDelta(pOut *ImVec2, button ImGuiMouseButton, lock_threshold float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMouseDragDelta(pOutArg, C.ImGuiMouseButton(button), C.float(lock_threshold))
}

func GetTextLineHeightWithSpacing() float32 {
	return float32(C.GetTextLineHeightWithSpacing())
}

func TreePop() {
	C.TreePop()
}

// BeginListBox parameter default value hint:
// size: ImVec2(0,0)
func BeginListBox(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.BeginListBox(labelArg, size.toC()) == C.bool(true)
}

func (self ImDrawList) PathFillConvex(col uint32) {
	C.DrawList_PathFillConvex(self.handle(), C.ImU32(col))
}

func NewImDrawList(shared_data ImDrawListSharedData) ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.DrawList_ImDrawList(shared_data.handle())))
}

// FontAtlas_AddCustomRectFontGlyph parameter default value hint:
// offset: ImVec2(0,0)
func (self ImFontAtlas) AddCustomRectFontGlyph(font ImFont, id ImWchar, width int32, height int32, advance_x float32, offset ImVec2) int {
	return int(C.FontAtlas_AddCustomRectFontGlyph(self.handle(), font.handle(), C.ImWchar(id), C.int(width), C.int(height), C.float(advance_x), offset.toC()))
}

func PushFont(font ImFont) {
	C.PushFont(font.handle())
}

func NewImDrawData() ImDrawData {
	return (ImDrawData)(unsafe.Pointer(C.DrawData_ImDrawData()))
}

func (self ImFont) CalcWordWrapPositionA(scale float32, text string, wrap_width float32) string {
	textArg, textFin := wrapString(text)
	defer textFin()

	return C.GoString(C.Font_CalcWordWrapPositionA(self.handle(), C.float(scale), textArg, C.float(wrap_width)))
}

func (self ImGuiIO) AddMouseButtonEvent(button int32, down bool) {
	C.IO_AddMouseButtonEvent(self.handle(), C.int(button), C.bool(down))
}

func (self ImGuiPtrOrIndex) Destroy() {
	C.PtrOrIndex_Destroy(self.handle())
}

func (self ImDrawData) DeIndexAllBuffers() {
	C.DrawData_DeIndexAllBuffers(self.handle())
}

// DrawList_AddBezierQuadratic parameter default value hint:
// num_segments: 0
func (self ImDrawList) AddBezierQuadratic(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.DrawList_AddBezierQuadratic(self.handle(), p1.toC(), p2.toC(), p3.toC(), C.ImU32(col), C.float(thickness), C.int(num_segments))
}

func NewImGuiIO() ImGuiIO {
	return (ImGuiIO)(unsafe.Pointer(C.IO_ImGuiIO()))
}

// DragFloat2 parameter default value hint:
// v_max: 0.0f
// v_min: 0.0f
// v_speed: 1.0f
// flags: 0
// format: "%.3f"
func DragFloat2(label string, v [2]*float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat2(labelArg, (*C.float)(&vArg[0]), C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func IsAnyItemFocused() bool {
	return C.IsAnyItemFocused() == C.bool(true)
}

// PushTextWrapPos parameter default value hint:
// wrap_local_pos_x: 0.0f
func PushTextWrapPos(wrap_local_pos_x float32) {
	C.PushTextWrapPos(C.float(wrap_local_pos_x))
}

// Combo_Str parameter default value hint:
// popup_max_height_in_items: -1
func Combo_Str(label string, current_item *int32, items_separated_by_zeros string, popup_max_height_in_items int32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	current_itemArg, current_itemFin := wrapInt32(current_item)
	defer current_itemFin()

	items_separated_by_zerosArg, items_separated_by_zerosFin := wrapString(items_separated_by_zeros)
	defer items_separated_by_zerosFin()

	return C.Combo_Str(labelArg, current_itemArg, items_separated_by_zerosArg, C.int(popup_max_height_in_items)) == C.bool(true)
}

// TextUnformatted parameter default value hint:
// text_end: NULL
func TextUnformatted(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.TextUnformatted(textArg)
}

// DrawList_PathBezierCubicCurveTo parameter default value hint:
// num_segments: 0
func (self ImDrawList) PathBezierCubicCurveTo(p2 ImVec2, p3 ImVec2, p4 ImVec2, num_segments int32) {
	C.DrawList_PathBezierCubicCurveTo(self.handle(), p2.toC(), p3.toC(), p4.toC(), C.int(num_segments))
}

func DestroyPlatformWindows() {
	C.DestroyPlatformWindows()
}

func EndMenu() {
	C.EndMenu()
}

func (self ImDrawData) ScaleClipRects(fb_scale ImVec2) {
	C.DrawData_ScaleClipRects(self.handle(), fb_scale.toC())
}

func (self ImDrawList) PushClipRectFullScreen() {
	C.DrawList_PushClipRectFullScreen(self.handle())
}

func CalcItemWidth() float32 {
	return float32(C.CalcItemWidth())
}

func GetDragDropPayload() ImGuiPayload {
	return (ImGuiPayload)(unsafe.Pointer(C.GetDragDropPayload()))
}

// TreeNodeEx_Str parameter default value hint:
// flags: 0
func TreeNodeEx_Str(label string, flags ImGuiTreeNodeFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TreeNodeEx_Str(labelArg, C.ImGuiTreeNodeFlags(flags)) == C.bool(true)
}

func TreeNodeEx_StrStr(str_id string, flags ImGuiTreeNodeFlags, fmt string) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	return C.TreeNodeEx_StrStr(str_idArg, C.ImGuiTreeNodeFlags(flags), fmtArg) == C.bool(true)
}

func TreeNodeEx_Ptr(ptr_id unsafe.Pointer, flags ImGuiTreeNodeFlags, fmt string) bool {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	return C.TreeNodeEx_Ptr(ptr_id, C.ImGuiTreeNodeFlags(flags), fmtArg) == C.bool(true)
}

func PopTextWrapPos() {
	C.PopTextWrapPos()
}

func EndChildFrame() {
	C.EndChildFrame()
}

func GetMouseClickedCount(button ImGuiMouseButton) int {
	return int(C.GetMouseClickedCount(C.ImGuiMouseButton(button)))
}

func SetColumnWidth(column_index int32, width float32) {
	C.SetColumnWidth(C.int(column_index), C.float(width))
}

// SetScrollHereX parameter default value hint:
// center_x_ratio: 0.5f
func SetScrollHereX(center_x_ratio float32) {
	C.SetScrollHereX(C.float(center_x_ratio))
}

func PopFont() {
	C.PopFont()
}

// SliderScalarN parameter default value hint:
// flags: 0
// format: NULL
func SliderScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// MenuItem_Bool parameter default value hint:
// enabled: true
// selected: false
// shortcut: NULL
func MenuItem_Bool(label string, shortcut string, selected bool, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	shortcutArg, shortcutFin := wrapString(shortcut)
	defer shortcutFin()

	return C.MenuItem_Bool(labelArg, shortcutArg, C.bool(selected), C.bool(enabled)) == C.bool(true)
}

// MenuItem_BoolPtr parameter default value hint:
// enabled: true
func MenuItem_BoolPtr(label string, shortcut string, p_selected *bool, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	shortcutArg, shortcutFin := wrapString(shortcut)
	defer shortcutFin()

	p_selectedArg, p_selectedFin := wrapBool(p_selected)
	defer p_selectedFin()

	return C.MenuItem_BoolPtr(labelArg, shortcutArg, p_selectedArg, C.bool(enabled)) == C.bool(true)
}

// SliderInt3 parameter default value hint:
// flags: 0
// format: "%d"
func SliderInt3(label string, v [3]*int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt3(labelArg, (*C.int)(&vArg[0]), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImDrawCmd) Destroy() {
	C.DrawCmd_Destroy(self.handle())
}

func GetFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.GetFont()))
}

// SaveIniSettingsToMemory parameter default value hint:
// out_ini_size: NULL
func SaveIniSettingsToMemory(out_ini_size *uint64) string {
	return C.GoString(C.SaveIniSettingsToMemory((*C.xlong)(out_ini_size)))
}

func (self ImDrawList) PopTextureID() {
	C.DrawList_PopTextureID(self.handle())
}

func GetContentRegionAvail(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetContentRegionAvail(pOutArg)
}

func (self ImDrawList) PathLineToMergeDuplicate(pos ImVec2) {
	C.DrawList_PathLineToMergeDuplicate(self.handle(), pos.toC())
}

// BeginChild_Str parameter default value hint:
// size: ImVec2(0,0)
// border: false
// flags: 0
func BeginChild_Str(str_id string, size ImVec2, border bool, flags ImGuiWindowFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginChild_Str(str_idArg, size.toC(), C.bool(border), C.ImGuiWindowFlags(flags)) == C.bool(true)
}

// BeginChild_ID parameter default value hint:
// border: false
// flags: 0
// size: ImVec2(0,0)
func BeginChild_ID(id ImGuiID, size ImVec2, border bool, flags ImGuiWindowFlags) bool {
	return C.BeginChild_ID(C.ImGuiID(id), size.toC(), C.bool(border), C.ImGuiWindowFlags(flags)) == C.bool(true)
}

func NewImGuiOnceUponAFrame() ImGuiOnceUponAFrame {
	return (ImGuiOnceUponAFrame)(unsafe.Pointer(C.OnceUponAFrame_ImGuiOnceUponAFrame()))
}

func GetColumnsCount() int {
	return int(C.GetColumnsCount())
}

func GetDrawCursorScreenPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetDrawCursorScreenPos(pOutArg)
}

func IsAnyItemActive() bool {
	return C.IsAnyItemActive() == C.bool(true)
}

func TreePush_Str(str_id string) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.TreePush_Str(str_idArg)
}

// TreePush_Ptr parameter default value hint:
// ptr_id: NULL
func TreePush_Ptr(ptr_id unsafe.Pointer) {
	C.TreePush_Ptr(ptr_id)
}

// FontAtlas_AddFontFromFileTTF parameter default value hint:
// font_cfg: NULL
// glyph_ranges: NULL
func (self ImFontAtlas) AddFontFromFileTTF(filename string, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromFileTTF(self.handle(), filenameArg, C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func IsItemVisible() bool {
	return C.IsItemVisible() == C.bool(true)
}

func (self ImGuiListClipper) End() {
	C.ListClipper_End(self.handle())
}

func (self ImDrawList) PrimRect(a ImVec2, b ImVec2, col uint32) {
	C.DrawList_PrimRect(self.handle(), a.toC(), b.toC(), C.ImU32(col))
}

func (self ImGuiListClipper) Destroy() {
	C.ListClipper_Destroy(self.handle())
}

func ColorConvertFloat4ToU32(in ImVec4) uint32 {
	return uint32(C.ColorConvertFloat4ToU32(in.toC()))
}

func End() {
	C.End()
}

func EndCombo() {
	C.EndCombo()
}

func PopID() {
	C.PopID()
}

func (self ImGuiPlatformIO) Destroy() {
	C.PlatformIO_Destroy(self.handle())
}

func GetScrollX() float32 {
	return float32(C.GetScrollX())
}

func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return C.IsMouseDoubleClicked(C.ImGuiMouseButton(button)) == C.bool(true)
}

// DrawList_PushClipRect parameter default value hint:
// intersect_with_current_clip_rect: false
func (self ImDrawList) PushClipRect(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	C.DrawList_PushClipRect(self.handle(), clip_rect_min.toC(), clip_rect_max.toC(), C.bool(intersect_with_current_clip_rect))
}

func NewImGuiTableSortSpecs() ImGuiTableSortSpecs {
	return (ImGuiTableSortSpecs)(unsafe.Pointer(C.TableSortSpecs_ImGuiTableSortSpecs()))
}

func GetContentRegionMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetContentRegionMax(pOutArg)
}

func IsItemFocused() bool {
	return C.IsItemFocused() == C.bool(true)
}

func TableSetupScrollFreeze(cols int32, rows int32) {
	C.TableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func (self ImFont) Destroy() {
	C.Font_Destroy(self.handle())
}

// RenderPlatformWindowsDefault parameter default value hint:
// platform_render_arg: NULL
// renderer_render_arg: NULL
func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer, renderer_render_arg unsafe.Pointer) {
	C.RenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg)
}

func (self ImFontAtlas) ClearFonts() {
	C.FontAtlas_ClearFonts(self.handle())
}

// ColorEdit3 parameter default value hint:
// flags: 0
func ColorEdit3(label string, col [3]*float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := make([]C.float, len(col))
	for i, colV := range col {
		colArg[i] = C.float(*colV)
	}
	defer func() {
		for i, colV := range colArg {
			*col[i] = float32(colV)
		}
	}()

	return C.ColorEdit3(labelArg, (*C.float)(&colArg[0]), C.ImGuiColorEditFlags(flags)) == C.bool(true)
}

func IsMouseReleased(button ImGuiMouseButton) bool {
	return C.IsMouseReleased(C.ImGuiMouseButton(button)) == C.bool(true)
}

func LogText(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.LogText(fmtArg)
}

func (self ImFontAtlasCustomRect) IsPacked() bool {
	return C.FontAtlasCustomRect_IsPacked(self.handle()) == C.bool(true)
}

func (self ImGuiPlatformMonitor) Destroy() {
	C.PlatformMonitor_Destroy(self.handle())
}

func TableSetColumnEnabled(column_n int32, v bool) {
	C.TableSetColumnEnabled(C.int(column_n), C.bool(v))
}

func ColorConvertRGBtoHSV(r float32, g float32, b float32, out_h *float32, out_s *float32, out_v *float32) {
	out_hArg, out_hFin := wrapFloat(out_h)
	defer out_hFin()

	out_sArg, out_sFin := wrapFloat(out_s)
	defer out_sFin()

	out_vArg, out_vFin := wrapFloat(out_v)
	defer out_vFin()

	C.ColorConvertRGBtoHSV(C.float(r), C.float(g), C.float(b), out_hArg, out_sArg, out_vArg)
}

func GetDrawListSharedData() ImDrawListSharedData {
	return (ImDrawListSharedData)(unsafe.Pointer(C.GetDrawListSharedData()))
}

// SetNextWindowSize parameter default value hint:
// cond: 0
func SetNextWindowSize(size ImVec2, cond ImGuiCond) {
	C.SetNextWindowSize(size.toC(), C.ImGuiCond(cond))
}

func (self ImGuiTextBuffer) Size() int {
	return int(C.TextBuffer_Size(self.handle()))
}

// StyleColorsClassic parameter default value hint:
// dst: NULL
func StyleColorsClassic(dst ImGuiStyle) {
	C.StyleColorsClassic(dst.handle())
}

func (self ImFontConfig) Destroy() {
	C.FontConfig_Destroy(self.handle())
}

func (self ImDrawList) ChannelsMerge() {
	C.DrawList_ChannelsMerge(self.handle())
}

func GetBackgroundDrawList_Nil() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetBackgroundDrawList_Nil()))
}

func GetBackgroundDrawList_ViewportPtr(viewport ImGuiViewport) ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetBackgroundDrawList_ViewportPtr(viewport.handle())))
}

// LogToTTY parameter default value hint:
// auto_open_depth: -1
func LogToTTY(auto_open_depth int32) {
	C.LogToTTY(C.int(auto_open_depth))
}

func (self ImGuiInputTextCallbackData) Destroy() {
	C.InputTextCallbackData_Destroy(self.handle())
}

// SetScrollHereY parameter default value hint:
// center_y_ratio: 0.5f
func SetScrollHereY(center_y_ratio float32) {
	C.SetScrollHereY(C.float(center_y_ratio))
}

// ShowDemoWindow parameter default value hint:
// p_open: NULL
func ShowDemoWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDemoWindow(p_openArg)
}

// BeginDragDropSource parameter default value hint:
// flags: 0
func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
	return C.BeginDragDropSource(C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

func PushStyleVar_Float(idx ImGuiStyleVar, val float32) {
	C.PushStyleVar_Float(C.ImGuiStyleVar(idx), C.float(val))
}

func PushStyleVar_Vec2(idx ImGuiStyleVar, val ImVec2) {
	C.PushStyleVar_Vec2(C.ImGuiStyleVar(idx), val.toC())
}

func MemAlloc(size uint64) unsafe.Pointer {
	return unsafe.Pointer(C.MemAlloc(C.xlong(size)))
}

// SliderFloat parameter default value hint:
// flags: 0
// format: "%.3f"
func SliderFloat(label string, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// DragInt3 parameter default value hint:
// format: "%d"
// v_max: 0
// v_min: 0
// v_speed: 1.0f
// flags: 0
func DragInt3(label string, v [3]*int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt3(labelArg, (*C.int)(&vArg[0]), C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func NewImGuiPlatformIO() ImGuiPlatformIO {
	return (ImGuiPlatformIO)(unsafe.Pointer(C.PlatformIO_ImGuiPlatformIO()))
}

func (self ImDrawList) AddConvexPolyFilled(points *ImVec2, num_points int32, col uint32) {
	pointsArg, pointsFin := points.wrap()
	defer pointsFin()

	C.DrawList_AddConvexPolyFilled(self.handle(), pointsArg, C.int(num_points), C.ImU32(col))
}

func (self ImGuiPayload) Clear() {
	C.Payload_Clear(self.handle())
}

func GetStyleColorName(idx ImGuiCol) string {
	return C.GoString(C.GetStyleColorName(C.ImGuiCol(idx)))
}

func GetWindowDockID() ImGuiID {
	return ImGuiID(C.GetWindowDockID())
}

func EndMainMenuBar() {
	C.EndMainMenuBar()
}

func GetWindowPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowPos(pOutArg)
}

func Bullet() {
	C.Bullet()
}

// DragInt parameter default value hint:
// v_speed: 1.0f
// flags: 0
// format: "%d"
// v_max: 0
// v_min: 0
func DragInt(label string, v *int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// IsWindowHovered parameter default value hint:
// flags: 0
func IsWindowHovered(flags ImGuiHoveredFlags) bool {
	return C.IsWindowHovered(C.ImGuiHoveredFlags(flags)) == C.bool(true)
}

func PushClipRect(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	C.PushClipRect(clip_rect_min.toC(), clip_rect_max.toC(), C.bool(intersect_with_current_clip_rect))
}

func (self ImGuiPlatformImeData) Destroy() {
	C.PlatformImeData_Destroy(self.handle())
}

func GetDrawCursorPosY() float32 {
	return float32(C.GetDrawCursorPosY())
}

// IsWindowFocused parameter default value hint:
// flags: 0
func IsWindowFocused(flags ImGuiFocusedFlags) bool {
	return C.IsWindowFocused(C.ImGuiFocusedFlags(flags)) == C.bool(true)
}

func PopButtonRepeat() {
	C.PopButtonRepeat()
}

func (self ImGuiListClipper) Step() bool {
	return C.ListClipper_Step(self.handle()) == C.bool(true)
}

func (self ImGuiPayload) IsDataType(typeArg string) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.Payload_IsDataType(self.handle(), typeArgArg) == C.bool(true)
}

func IsItemEdited() bool {
	return C.IsItemEdited() == C.bool(true)
}

func SetCurrentContext(ctx ImGuiContext) {
	C.SetCurrentContext(ctx.handle())
}

func (self ImFontAtlas) GetGlyphRangesCyrillic() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesCyrillic(self.handle()))
}

func (self ImFontAtlas) Destroy() {
	C.FontAtlas_Destroy(self.handle())
}

// VSliderInt parameter default value hint:
// flags: 0
// format: "%d"
func VSliderInt(label string, size ImVec2, v *int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderInt(labelArg, size.toC(), vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func Value_Bool(prefix string, b bool) {
	prefixArg, prefixFin := wrapString(prefix)
	defer prefixFin()

	C.Value_Bool(prefixArg, C.bool(b))
}

func Value_Int(prefix string, v int32) {
	prefixArg, prefixFin := wrapString(prefix)
	defer prefixFin()

	C.Value_Int(prefixArg, C.int(v))
}

func Value_Uint(prefix string, v uint32) {
	prefixArg, prefixFin := wrapString(prefix)
	defer prefixFin()

	C.Value_Uint(prefixArg, C.uint(v))
}

// Value_Float parameter default value hint:
// float_format: NULL
func Value_Float(prefix string, v float32, float_format string) {
	prefixArg, prefixFin := wrapString(prefix)
	defer prefixFin()

	float_formatArg, float_formatFin := wrapString(float_format)
	defer float_formatFin()

	C.Value_Float(prefixArg, C.float(v), float_formatArg)
}

func EndDragDropTarget() {
	C.EndDragDropTarget()
}

func GetPlatformIO() ImGuiPlatformIO {
	return (ImGuiPlatformIO)(unsafe.Pointer(C.GetPlatformIO()))
}

func BeginTooltip() {
	C.BeginTooltip()
}

func IsAnyMouseDown() bool {
	return C.IsAnyMouseDown() == C.bool(true)
}

// DrawList_AddRect parameter default value hint:
// flags: 0
// rounding: 0.0f
// thickness: 1.0f
func (self ImDrawList) AddRect(p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags, thickness float32) {
	C.DrawList_AddRect(self.handle(), p_min.toC(), p_max.toC(), C.ImU32(col), C.float(rounding), C.ImDrawFlags(flags), C.float(thickness))
}

func (self ImDrawList) PrimQuadUV(a ImVec2, b ImVec2, c ImVec2, d ImVec2, uv_a ImVec2, uv_b ImVec2, uv_c ImVec2, uv_d ImVec2, col uint32) {
	C.DrawList_PrimQuadUV(self.handle(), a.toC(), b.toC(), c.toC(), d.toC(), uv_a.toC(), uv_b.toC(), uv_c.toC(), uv_d.toC(), C.ImU32(col))
}

func IsKeyDown(key ImGuiKey) bool {
	return C.IsKeyDown(C.ImGuiKey(key)) == C.bool(true)
}

// SliderInt2 parameter default value hint:
// flags: 0
// format: "%d"
func SliderInt2(label string, v [2]*int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt2(labelArg, (*C.int)(&vArg[0]), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func NewFrame() {
	C.NewFrame()
}

// OpenPopup_Str parameter default value hint:
// popup_flags: 0
func OpenPopup_Str(str_id string, popup_flags ImGuiPopupFlags) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.OpenPopup_Str(str_idArg, C.ImGuiPopupFlags(popup_flags))
}

// OpenPopup_ID parameter default value hint:
// popup_flags: 0
func OpenPopup_ID(id ImGuiID, popup_flags ImGuiPopupFlags) {
	C.OpenPopup_ID(C.ImGuiID(id), C.ImGuiPopupFlags(popup_flags))
}

func (self ImGuiListClipper) ForceDisplayRangeByIndices(item_min int32, item_max int32) {
	C.ListClipper_ForceDisplayRangeByIndices(self.handle(), C.int(item_min), C.int(item_max))
}

func NewImGuiWindowClass() ImGuiWindowClass {
	return (ImGuiWindowClass)(unsafe.Pointer(C.WindowClass_ImGuiWindowClass()))
}

func SetNextWindowFocus() {
	C.SetNextWindowFocus()
}

// GetColumnWidth parameter default value hint:
// column_index: -1
func GetColumnWidth(column_index int32) float32 {
	return float32(C.GetColumnWidth(C.int(column_index)))
}

// DrawList_AddTriangle parameter default value hint:
// thickness: 1.0f
func (self ImDrawList) AddTriangle(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddTriangle(self.handle(), p1.toC(), p2.toC(), p3.toC(), C.ImU32(col), C.float(thickness))
}

// InputTextCallbackData_InsertChars parameter default value hint:
// text_end: NULL
func (self ImGuiInputTextCallbackData) InsertChars(pos int32, text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.InputTextCallbackData_InsertChars(self.handle(), C.int(pos), textArg)
}

func ColorConvertU32ToFloat4(pOut *ImVec4, in uint32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ColorConvertU32ToFloat4(pOutArg, C.ImU32(in))
}

// SliderFloat2 parameter default value hint:
// flags: 0
// format: "%.3f"
func SliderFloat2(label string, v [2]*float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat2(labelArg, (*C.float)(&vArg[0]), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func IsWindowDocked() bool {
	return C.IsWindowDocked() == C.bool(true)
}

// TableSetupColumn parameter default value hint:
// flags: 0
// init_width_or_weight: 0.0f
// user_id: 0
func TableSetupColumn(label string, flags ImGuiTableColumnFlags, init_width_or_weight float32, user_id ImGuiID) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableSetupColumn(labelArg, C.ImGuiTableColumnFlags(flags), C.float(init_width_or_weight), C.ImGuiID(user_id))
}

func NewImGuiTableColumnSortSpecs() ImGuiTableColumnSortSpecs {
	return (ImGuiTableColumnSortSpecs)(unsafe.Pointer(C.TableColumnSortSpecs_ImGuiTableColumnSortSpecs()))
}

// IsMousePosValid parameter default value hint:
// mouse_pos: NULL
func IsMousePosValid(mouse_pos *ImVec2) bool {
	mouse_posArg, mouse_posFin := mouse_pos.wrap()
	defer mouse_posFin()

	return C.IsMousePosValid(mouse_posArg) == C.bool(true)
}

// FontGlyphRangesBuilder_AddText parameter default value hint:
// text_end: NULL
func (self ImFontGlyphRangesBuilder) AddText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.FontGlyphRangesBuilder_AddText(self.handle(), textArg)
}

// Font_AddRemapChar parameter default value hint:
// overwrite_dst: true
func (self ImFont) AddRemapChar(dst ImWchar, src ImWchar, overwrite_dst bool) {
	C.Font_AddRemapChar(self.handle(), C.ImWchar(dst), C.ImWchar(src), C.bool(overwrite_dst))
}

func (self ImGuiNextItemData) Destroy() {
	C.NextItemData_Destroy(self.handle())
}

// PlotHistogram_FloatPtr parameter default value hint:
// stride: sizeof(float)
// values_offset: 0
// graph_size: ImVec2(0,0)
// overlay_text: NULL
// scale_max: FLT_MAX
// scale_min: FLT_MAX
func PlotHistogram_FloatPtr(label string, values []float32, values_count int32, values_offset int32, overlay_text string, scale_min float32, scale_max float32, graph_size ImVec2, stride int32) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	overlay_textArg, overlay_textFin := wrapString(overlay_text)
	defer overlay_textFin()

	C.PlotHistogram_FloatPtr(labelArg, (*C.float)(&(values[0])), C.int(values_count), C.int(values_offset), overlay_textArg, C.float(scale_min), C.float(scale_max), graph_size.toC(), C.int(stride))
}

func AlignTextToFramePadding() {
	C.AlignTextToFramePadding()
}

func DrawList_GetClipRectMin(pOut *ImVec2, self ImDrawList) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.DrawList_GetClipRectMin(pOutArg, self.handle())
}

func (self ImGuiPayload) IsPreview() bool {
	return C.Payload_IsPreview(self.handle()) == C.bool(true)
}

func EndChild() {
	C.EndChild()
}

// StyleColorsLight parameter default value hint:
// dst: NULL
func StyleColorsLight(dst ImGuiStyle) {
	C.StyleColorsLight(dst.handle())
}

func TableSetColumnIndex(column_n int32) bool {
	return C.TableSetColumnIndex(C.int(column_n)) == C.bool(true)
}

// ShowAboutWindow parameter default value hint:
// p_open: NULL
func ShowAboutWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowAboutWindow(p_openArg)
}

func GetWindowDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetWindowDrawList()))
}

// DragFloat parameter default value hint:
// format: "%.3f"
// v_max: 0.0f
// v_min: 0.0f
// v_speed: 1.0f
// flags: 0
func DragFloat(label string, v *float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// VSliderFloat parameter default value hint:
// flags: 0
// format: "%.3f"
func VSliderFloat(label string, size ImVec2, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderFloat(labelArg, size.toC(), vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImGuiPayload) Destroy() {
	C.Payload_Destroy(self.handle())
}

func (self ImGuiStackLevelInfo) Destroy() {
	C.StackLevelInfo_Destroy(self.handle())
}

func TableHeadersRow() {
	C.TableHeadersRow()
}

func NewImDrawListSplitter() ImDrawListSplitter {
	return (ImDrawListSplitter)(unsafe.Pointer(C.DrawListSplitter_ImDrawListSplitter()))
}

func BeginGroup() {
	C.BeginGroup()
}

// GetColorU32_Col parameter default value hint:
// alpha_mul: 1.0f
func GetColorU32_Col(idx ImGuiCol, alpha_mul float32) uint32 {
	return uint32(C.GetColorU32_Col(C.ImGuiCol(idx), C.float(alpha_mul)))
}

func GetColorU32_Vec4(col ImVec4) uint32 {
	return uint32(C.GetColorU32_Vec4(col.toC()))
}

func GetColorU32_U32(col uint32) uint32 {
	return uint32(C.GetColorU32_U32(C.ImU32(col)))
}

func GetScrollY() float32 {
	return float32(C.GetScrollY())
}

func (self ImDrawList) AddQuadFilled(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32) {
	C.DrawList_AddQuadFilled(self.handle(), p1.toC(), p2.toC(), p3.toC(), p4.toC(), C.ImU32(col))
}

// FontAtlas_AddFontFromMemoryTTF parameter default value hint:
// font_cfg: NULL
// glyph_ranges: NULL
func (self ImFontAtlas) AddFontFromMemoryTTF(font_data unsafe.Pointer, font_size int32, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromMemoryTTF(self.handle(), font_data, C.int(font_size), C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

// CalcTextSize parameter default value hint:
// hide_text_after_double_hash: false
// text_end: NULL
// wrap_width: -1.0f
func CalcTextSize(pOut *ImVec2, text string, hide_text_after_double_hash bool, wrap_width float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	textArg, textFin := wrapString(text)
	defer textFin()

	C.CalcTextSize(pOutArg, textArg, C.bool(hide_text_after_double_hash), C.float(wrap_width))
}

func (self ImGuiStackSizes) Destroy() {
	C.StackSizes_Destroy(self.handle())
}

// BeginChildFrame parameter default value hint:
// flags: 0
func BeginChildFrame(id ImGuiID, size ImVec2, flags ImGuiWindowFlags) bool {
	return C.BeginChildFrame(C.ImGuiID(id), size.toC(), C.ImGuiWindowFlags(flags)) == C.bool(true)
}

func EndTabItem() {
	C.EndTabItem()
}

func GetMainViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.GetMainViewport()))
}

func NewImDrawCmd() ImDrawCmd {
	return (ImDrawCmd)(unsafe.Pointer(C.DrawCmd_ImDrawCmd()))
}

func (self ImGuiTextBuffer) End() string {
	return C.GoString(C.TextBuffer_End(self.handle()))
}

func PushStyleColor_U32(idx ImGuiCol, col uint32) {
	C.PushStyleColor_U32(C.ImGuiCol(idx), C.ImU32(col))
}

func PushStyleColor_Vec4(idx ImGuiCol, col ImVec4) {
	C.PushStyleColor_Vec4(C.ImGuiCol(idx), col.toC())
}

func SetNextWindowViewport(viewport_id ImGuiID) {
	C.SetNextWindowViewport(C.ImGuiID(viewport_id))
}

func Spacing() {
	C.Spacing()
}

func PushItemWidth(item_width float32) {
	C.PushItemWidth(C.float(item_width))
}

func TableGetColumnIndex() int {
	return int(C.TableGetColumnIndex())
}

func EndFrame() {
	C.EndFrame()
}

// PlotLines_FloatPtr parameter default value hint:
// values_offset: 0
// graph_size: ImVec2(0,0)
// overlay_text: NULL
// scale_max: FLT_MAX
// scale_min: FLT_MAX
// stride: sizeof(float)
func PlotLines_FloatPtr(label string, values []float32, values_count int32, values_offset int32, overlay_text string, scale_min float32, scale_max float32, graph_size ImVec2, stride int32) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	overlay_textArg, overlay_textFin := wrapString(overlay_text)
	defer overlay_textFin()

	C.PlotLines_FloatPtr(labelArg, (*C.float)(&(values[0])), C.int(values_count), C.int(values_offset), overlay_textArg, C.float(scale_min), C.float(scale_max), graph_size.toC(), C.int(stride))
}

func NewImFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.Font_ImFont()))
}

func PushButtonRepeat(repeat bool) {
	C.PushButtonRepeat(C.bool(repeat))
}

func (self ImFontAtlas) GetGlyphRangesKorean() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesKorean(self.handle()))
}

// TableGetColumnName_Int parameter default value hint:
// column_n: -1
func TableGetColumnName_Int(column_n int32) string {
	return C.GoString(C.TableGetColumnName_Int(C.int(column_n)))
}

// GetColumnOffset parameter default value hint:
// column_index: -1
func GetColumnOffset(column_index int32) float32 {
	return float32(C.GetColumnOffset(C.int(column_index)))
}

// PopStyleVar parameter default value hint:
// count: 1
func PopStyleVar(count int32) {
	C.PopStyleVar(C.int(count))
}

func SetNextWindowClass(window_class ImGuiWindowClass) {
	C.SetNextWindowClass(window_class.handle())
}

func (self ImFontAtlas) Build() bool {
	return C.FontAtlas_Build(self.handle()) == C.bool(true)
}

func SetTabItemClosed(tab_or_docked_window_label string) {
	tab_or_docked_window_labelArg, tab_or_docked_window_labelFin := wrapString(tab_or_docked_window_label)
	defer tab_or_docked_window_labelFin()

	C.SetTabItemClosed(tab_or_docked_window_labelArg)
}

// SliderFloat4 parameter default value hint:
// flags: 0
// format: "%.3f"
func SliderFloat4(label string, v [4]*float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.float, len(v))
	for i, vV := range v {
		vArg[i] = C.float(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = float32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat4(labelArg, (*C.float)(&vArg[0]), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// BeginPopupContextWindow parameter default value hint:
// str_id: NULL
// popup_flags: 1
func BeginPopupContextWindow(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextWindow(str_idArg, C.ImGuiPopupFlags(popup_flags)) == C.bool(true)
}

func (self ImFontAtlas) GetGlyphRangesDefault() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesDefault(self.handle()))
}

func TableGetRowIndex() int {
	return int(C.TableGetRowIndex())
}

func (self ImGuiTableTempData) Destroy() {
	C.TableTempData_Destroy(self.handle())
}

func BeginDragDropTarget() bool {
	return C.BeginDragDropTarget() == C.bool(true)
}

// CollapsingHeader_TreeNodeFlags parameter default value hint:
// flags: 0
func CollapsingHeader_TreeNodeFlags(label string, flags ImGuiTreeNodeFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.CollapsingHeader_TreeNodeFlags(labelArg, C.ImGuiTreeNodeFlags(flags)) == C.bool(true)
}

// CollapsingHeader_BoolPtr parameter default value hint:
// flags: 0
func CollapsingHeader_BoolPtr(label string, p_visible *bool, flags ImGuiTreeNodeFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_visibleArg, p_visibleFin := wrapBool(p_visible)
	defer p_visibleFin()

	return C.CollapsingHeader_BoolPtr(labelArg, p_visibleArg, C.ImGuiTreeNodeFlags(flags)) == C.bool(true)
}

// SetWindowPos_Vec2 parameter default value hint:
// cond: 0
func SetWindowPos_Vec2(pos ImVec2, cond ImGuiCond) {
	C.SetWindowPos_Vec2(pos.toC(), C.ImGuiCond(cond))
}

// SetWindowPos_Str parameter default value hint:
// cond: 0
func SetWindowPos_Str(name string, pos ImVec2, cond ImGuiCond) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.SetWindowPos_Str(nameArg, pos.toC(), C.ImGuiCond(cond))
}

// InputScalar parameter default value hint:
// flags: 0
// format: NULL
// p_step: NULL
// p_step_fast: NULL
func InputScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func (self ImFontAtlasCustomRect) Destroy() {
	C.FontAtlasCustomRect_Destroy(self.handle())
}

func (self ImGuiNextWindowData) Destroy() {
	C.NextWindowData_Destroy(self.handle())
}

func TableNextColumn() bool {
	return C.TableNextColumn() == C.bool(true)
}

func ShowUserGuide() {
	C.ShowUserGuide()
}

func LogFinish() {
	C.LogFinish()
}

// VSliderScalar parameter default value hint:
// flags: 0
// format: NULL
func VSliderScalar(label string, size ImVec2, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderScalar(labelArg, size.toC(), C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

// FontAtlas_GetTexDataAsRGBA32 parameter default value hint:
// out_bytes_per_pixel: NULL
func (self ImFontAtlas) GetTexDataAsRGBA32(out_pixels *C.uchar, out_width *int32, out_height *int32, out_bytes_per_pixel *int32) {
	out_widthArg, out_widthFin := wrapInt32(out_width)
	defer out_widthFin()

	out_heightArg, out_heightFin := wrapInt32(out_height)
	defer out_heightFin()

	out_bytes_per_pixelArg, out_bytes_per_pixelFin := wrapInt32(out_bytes_per_pixel)
	defer out_bytes_per_pixelFin()

	C.FontAtlas_GetTexDataAsRGBA32(self.handle(), &out_pixels, out_widthArg, out_heightArg, out_bytes_per_pixelArg)
}

func (self ImFont) AddGlyph(src_cfg ImFontConfig, c ImWchar, x0 float32, y0 float32, x1 float32, y1 float32, u0 float32, v0 float32, u1 float32, v1 float32, advance_x float32) {
	C.Font_AddGlyph(self.handle(), src_cfg.handle(), C.ImWchar(c), C.float(x0), C.float(y0), C.float(x1), C.float(y1), C.float(u0), C.float(v0), C.float(u1), C.float(v1), C.float(advance_x))
}

// FontAtlas_GetTexDataAsAlpha8 parameter default value hint:
// out_bytes_per_pixel: NULL
func (self ImFontAtlas) GetTexDataAsAlpha8(out_pixels *C.uchar, out_width *int32, out_height *int32, out_bytes_per_pixel *int32) {
	out_widthArg, out_widthFin := wrapInt32(out_width)
	defer out_widthFin()

	out_heightArg, out_heightFin := wrapInt32(out_height)
	defer out_heightFin()

	out_bytes_per_pixelArg, out_bytes_per_pixelFin := wrapInt32(out_bytes_per_pixel)
	defer out_bytes_per_pixelFin()

	C.FontAtlas_GetTexDataAsAlpha8(self.handle(), &out_pixels, out_widthArg, out_heightArg, out_bytes_per_pixelArg)
}

func NewImGuiListClipper() ImGuiListClipper {
	return (ImGuiListClipper)(unsafe.Pointer(C.ListClipper_ImGuiListClipper()))
}

func IsMouseDown(button ImGuiMouseButton) bool {
	return C.IsMouseDown(C.ImGuiMouseButton(button)) == C.bool(true)
}

func SmallButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.SmallButton(labelArg) == C.bool(true)
}

func DebugCheckVersionAndDataLayout(version_str string, sz_io uint64, sz_style uint64, sz_vec2 uint64, sz_vec4 uint64, sz_drawvert uint64, sz_drawidx uint64) bool {
	version_strArg, version_strFin := wrapString(version_str)
	defer version_strFin()

	return C.DebugCheckVersionAndDataLayout(version_strArg, C.xlong(sz_io), C.xlong(sz_style), C.xlong(sz_vec2), C.xlong(sz_vec4), C.xlong(sz_drawvert), C.xlong(sz_drawidx)) == C.bool(true)
}

// ShowMetricsWindow parameter default value hint:
// p_open: NULL
func ShowMetricsWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowMetricsWindow(p_openArg)
}

func (self ImGuiComboPreviewData) Destroy() {
	C.ComboPreviewData_Destroy(self.handle())
}

func EndGroup() {
	C.EndGroup()
}

// Columns parameter default value hint:
// border: true
// count: 1
// id: NULL
func Columns(count int32, id string, border bool) {
	idArg, idFin := wrapString(id)
	defer idFin()

	C.Columns(C.int(count), idArg, C.bool(border))
}

func TableGetSortSpecs() ImGuiTableSortSpecs {
	return (ImGuiTableSortSpecs)(unsafe.Pointer(C.TableGetSortSpecs()))
}

func PopAllowKeyboardFocus() {
	C.PopAllowKeyboardFocus()
}

func (self ImGuiOnceUponAFrame) Destroy() {
	C.OnceUponAFrame_Destroy(self.handle())
}

func (self ImGuiInputTextCallbackData) DeleteChars(pos int32, bytes_count int32) {
	C.InputTextCallbackData_DeleteChars(self.handle(), C.int(pos), C.int(bytes_count))
}

func EndTooltip() {
	C.EndTooltip()
}

func SetDrawCursorPosX(local_x float32) {
	C.SetDrawCursorPosX(C.float(local_x))
}

// TableNextRow parameter default value hint:
// min_row_height: 0.0f
// row_flags: 0
func TableNextRow(row_flags ImGuiTableRowFlags, min_row_height float32) {
	C.TableNextRow(C.ImGuiTableRowFlags(row_flags), C.float(min_row_height))
}

// DrawList_AddCircle parameter default value hint:
// num_segments: 0
// thickness: 1.0f
func (self ImDrawList) AddCircle(center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.DrawList_AddCircle(self.handle(), center.toC(), C.float(radius), C.ImU32(col), C.int(num_segments), C.float(thickness))
}

func (self ImFontGlyphRangesBuilder) Clear() {
	C.FontGlyphRangesBuilder_Clear(self.handle())
}

// ResetMouseDragDelta parameter default value hint:
// button: 0
func ResetMouseDragDelta(button ImGuiMouseButton) {
	C.ResetMouseDragDelta(C.ImGuiMouseButton(button))
}

func GetItemRectMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectMax(pOutArg)
}

// DrawList_AddImageQuad parameter default value hint:
// col: 4294967295
// uv1: ImVec2(0,0)
// uv2: ImVec2(1,0)
// uv3: ImVec2(1,1)
// uv4: ImVec2(0,1)
func (self ImDrawList) AddImageQuad(user_texture_id ImTextureID, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, uv1 ImVec2, uv2 ImVec2, uv3 ImVec2, uv4 ImVec2, col uint32) {
	C.DrawList_AddImageQuad(self.handle(), C.ImTextureID(user_texture_id), p1.toC(), p2.toC(), p3.toC(), p4.toC(), uv1.toC(), uv2.toC(), uv3.toC(), uv4.toC(), C.ImU32(col))
}

func (self ImGuiTableColumnSortSpecs) Destroy() {
	C.TableColumnSortSpecs_Destroy(self.handle())
}

func Checkbox(label string, v *bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapBool(v)
	defer vFin()

	return C.Checkbox(labelArg, vArg) == C.bool(true)
}

// DockSpaceOverViewport parameter default value hint:
// flags: 0
// viewport: NULL
// window_class: NULL
func DockSpaceOverViewport(viewport ImGuiViewport, flags ImGuiDockNodeFlags, window_class ImGuiWindowClass) ImGuiID {
	return ImGuiID(C.DockSpaceOverViewport(viewport.handle(), C.ImGuiDockNodeFlags(flags), window_class.handle()))
}

func (self ImGuiIO) AddInputCharactersUTF8(str string) {
	strArg, strFin := wrapString(str)
	defer strFin()

	C.IO_AddInputCharactersUTF8(self.handle(), strArg)
}

func (self ImGuiIO) AddKeyAnalogEvent(key ImGuiKey, down bool, v float32) {
	C.IO_AddKeyAnalogEvent(self.handle(), C.ImGuiKey(key), C.bool(down), C.float(v))
}

// BeginMenu parameter default value hint:
// enabled: true
func BeginMenu(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.BeginMenu(labelArg, C.bool(enabled)) == C.bool(true)
}

func GetWindowSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowSize(pOutArg)
}

func LoadIniSettingsFromDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.LoadIniSettingsFromDisk(ini_filenameArg)
}

func (self ImFontGlyphRangesBuilder) AddChar(c ImWchar) {
	C.FontGlyphRangesBuilder_AddChar(self.handle(), C.ImWchar(c))
}

func (self ImGuiTextFilter) Build() {
	C.TextFilter_Build(self.handle())
}

func EndDisabled() {
	C.EndDisabled()
}

// InputScalarN parameter default value hint:
// flags: 0
// format: NULL
// p_step: NULL
// p_step_fast: NULL
func InputScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

// SetDragDropPayload parameter default value hint:
// cond: 0
func SetDragDropPayload(typeArg string, data unsafe.Pointer, sz uint64, cond ImGuiCond) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.SetDragDropPayload(typeArgArg, data, C.xlong(sz), C.ImGuiCond(cond)) == C.bool(true)
}

func SetWindowFontScale(scale float32) {
	C.SetWindowFontScale(C.float(scale))
}

func (self ImDrawData) Destroy() {
	C.DrawData_Destroy(self.handle())
}

func SetScrollX_Float(scroll_x float32) {
	C.SetScrollX_Float(C.float(scroll_x))
}

func (self ImDrawList) PrimUnreserve(idx_count int32, vtx_count int32) {
	C.DrawList_PrimUnreserve(self.handle(), C.int(idx_count), C.int(vtx_count))
}

func GetWindowViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.GetWindowViewport()))
}

func SetTooltip(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.SetTooltip(fmtArg)
}

func (self ImGuiInputTextCallbackData) SelectAll() {
	C.InputTextCallbackData_SelectAll(self.handle())
}

func (self ImGuiTabBar) Destroy() {
	C.TabBar_Destroy(self.handle())
}

func (self ImGuiInputEvent) Destroy() {
	C.InputEvent_Destroy(self.handle())
}

func (self ImGuiInputTextCallbackData) HasSelection() bool {
	return C.InputTextCallbackData_HasSelection(self.handle()) == C.bool(true)
}

// IsPopupOpen_Str parameter default value hint:
// flags: 0
func IsPopupOpen_Str(str_id string, flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.IsPopupOpen_Str(str_idArg, C.ImGuiPopupFlags(flags)) == C.bool(true)
}

// DrawList_AddCircleFilled parameter default value hint:
// num_segments: 0
func (self ImDrawList) AddCircleFilled(center ImVec2, radius float32, col uint32, num_segments int32) {
	C.DrawList_AddCircleFilled(self.handle(), center.toC(), C.float(radius), C.ImU32(col), C.int(num_segments))
}

// DrawList_PathStroke parameter default value hint:
// flags: 0
// thickness: 1.0f
func (self ImDrawList) PathStroke(col uint32, flags ImDrawFlags, thickness float32) {
	C.DrawList_PathStroke(self.handle(), C.ImU32(col), C.ImDrawFlags(flags), C.float(thickness))
}

func (self ImDrawCmd) GetTexID() ImTextureID {
	return ImTextureID(C.DrawCmd_GetTexID(self.handle()))
}

func BeginMenuBar() bool {
	return C.BeginMenuBar() == C.bool(true)
}

// CreateContext parameter default value hint:
// shared_font_atlas: NULL
func CreateContext(shared_font_atlas ImFontAtlas) ImGuiContext {
	return (ImGuiContext)(unsafe.Pointer(C.CreateContext(shared_font_atlas.handle())))
}

// Selectable_Bool parameter default value hint:
// selected: false
// size: ImVec2(0,0)
// flags: 0
func Selectable_Bool(label string, selected bool, flags ImGuiSelectableFlags, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Selectable_Bool(labelArg, C.bool(selected), C.ImGuiSelectableFlags(flags), size.toC()) == C.bool(true)
}

// Selectable_BoolPtr parameter default value hint:
// flags: 0
// size: ImVec2(0,0)
func Selectable_BoolPtr(label string, p_selected *bool, flags ImGuiSelectableFlags, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_selectedArg, p_selectedFin := wrapBool(p_selected)
	defer p_selectedFin()

	return C.Selectable_BoolPtr(labelArg, p_selectedArg, C.ImGuiSelectableFlags(flags), size.toC()) == C.bool(true)
}

func (self ImGuiStyle) Destroy() {
	C.Style_Destroy(self.handle())
}

func (self ImFontGlyphRangesBuilder) SetBit(n uint64) {
	C.FontGlyphRangesBuilder_SetBit(self.handle(), C.xlong(n))
}

func EndTabBar() {
	C.EndTabBar()
}

func Render() {
	C.Render()
}

// SliderInt4 parameter default value hint:
// flags: 0
// format: "%d"
func SliderInt4(label string, v [4]*int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := make([]C.int, len(v))
	for i, vV := range v {
		vArg[i] = C.int(*vV)
	}
	defer func() {
		for i, vV := range vArg {
			*v[i] = int32(vV)
		}
	}()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt4(labelArg, (*C.int)(&vArg[0]), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) == C.bool(true)
}

func (self ImFontAtlas) GetGlyphRangesJapanese() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesJapanese(self.handle()))
}

func (self ImFontAtlas) ClearTexData() {
	C.FontAtlas_ClearTexData(self.handle())
}

func IsItemToggledOpen() bool {
	return C.IsItemToggledOpen() == C.bool(true)
}

func (self ImDrawList) PathLineTo(pos ImVec2) {
	C.DrawList_PathLineTo(self.handle(), pos.toC())
}

func (self ImGuiIO) AddFocusEvent(focused bool) {
	C.IO_AddFocusEvent(self.handle(), C.bool(focused))
}

func FindViewportByPlatformHandle(platform_handle unsafe.Pointer) ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.FindViewportByPlatformHandle(platform_handle)))
}

// InvisibleButton parameter default value hint:
// flags: 0
func InvisibleButton(str_id string, size ImVec2, flags ImGuiButtonFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.InvisibleButton(str_idArg, size.toC(), C.ImGuiButtonFlags(flags)) == C.bool(true)
}

func Text(fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Text(fmtArg)
}

// FontAtlas_AddFontFromMemoryCompressedBase85TTF parameter default value hint:
// font_cfg: NULL
// glyph_ranges: NULL
func (self ImFontAtlas) AddFontFromMemoryCompressedBase85TTF(compressed_font_data_base85 string, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	compressed_font_data_base85Arg, compressed_font_data_base85Fin := wrapString(compressed_font_data_base85)
	defer compressed_font_data_base85Fin()

	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromMemoryCompressedBase85TTF(self.handle(), compressed_font_data_base85Arg, C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

// FontAtlas_AddFontFromMemoryCompressedTTF parameter default value hint:
// font_cfg: NULL
// glyph_ranges: NULL
func (self ImFontAtlas) AddFontFromMemoryCompressedTTF(compressed_font_data unsafe.Pointer, compressed_font_size int32, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromMemoryCompressedTTF(self.handle(), compressed_font_data, C.int(compressed_font_size), C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func SaveIniSettingsToDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.SaveIniSettingsToDisk(ini_filenameArg)
}

// InputFloat parameter default value hint:
// flags: 0
// format: "%.3f"
// step: 0.0f
// step_fast: 0.0f
func InputFloat(label string, v *float32, step float32, step_fast float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat(labelArg, vArg, C.float(step), C.float(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) == C.bool(true)
}

func PushID_Str(str_id string) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.PushID_Str(str_idArg)
}

func PushID_StrStr(str_id_begin string, str_id_end string) {
	str_id_beginArg, str_id_beginFin := wrapString(str_id_begin)
	defer str_id_beginFin()

	str_id_endArg, str_id_endFin := wrapString(str_id_end)
	defer str_id_endFin()

	C.PushID_StrStr(str_id_beginArg, str_id_endArg)
}

func PushID_Ptr(ptr_id unsafe.Pointer) {
	C.PushID_Ptr(ptr_id)
}

func PushID_Int(int_id int32) {
	C.PushID_Int(C.int(int_id))
}

// DrawList_AddBezierCubic parameter default value hint:
// num_segments: 0
func (self ImDrawList) AddBezierCubic(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.DrawList_AddBezierCubic(self.handle(), p1.toC(), p2.toC(), p3.toC(), p4.toC(), C.ImU32(col), C.float(thickness), C.int(num_segments))
}

func (self ImGuiTextFilter) Clear() {
	C.TextFilter_Clear(self.handle())
}

func GetScrollMaxX() float32 {
	return float32(C.GetScrollMaxX())
}

func IsWindowAppearing() bool {
	return C.IsWindowAppearing() == C.bool(true)
}

// DrawList_AddImageRounded parameter default value hint:
// flags: 0
func (self ImDrawList) AddImageRounded(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.DrawList_AddImageRounded(self.handle(), C.ImTextureID(user_texture_id), p_min.toC(), p_max.toC(), uv_min.toC(), uv_max.toC(), C.ImU32(col), C.float(rounding), C.ImDrawFlags(flags))
}

func (self ImGuiInputTextCallbackData) ClearSelection() {
	C.InputTextCallbackData_ClearSelection(self.handle())
}

func GetKeyIndex(key ImGuiKey) int {
	return int(C.GetKeyIndex(C.ImGuiKey(key)))
}

// IsKeyPressed parameter default value hint:
// repeat: true
func IsKeyPressed(key ImGuiKey, repeat bool) bool {
	return C.IsKeyPressed(C.ImGuiKey(key), C.bool(repeat)) == C.bool(true)
}

func (self ImGuiTextFilter) Destroy() {
	C.TextFilter_Destroy(self.handle())
}

func (self ImDrawList) AddDrawCmd() {
	C.DrawList_AddDrawCmd(self.handle())
}

func (self ImFontAtlas) GetGlyphRangesThai() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesThai(self.handle()))
}

// BeginCombo parameter default value hint:
// flags: 0
func BeginCombo(label string, preview_value string, flags ImGuiComboFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	preview_valueArg, preview_valueFin := wrapString(preview_value)
	defer preview_valueFin()

	return C.BeginCombo(labelArg, preview_valueArg, C.ImGuiComboFlags(flags)) == C.bool(true)
}

func GetColumnIndex() int {
	return int(C.GetColumnIndex())
}

func (self ImFontAtlas) GetGlyphRangesVietnamese() *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesVietnamese(self.handle()))
}

func (self ImFont) BuildLookupTable() {
	C.Font_BuildLookupTable(self.handle())
}

func (self ImGuiMetricsConfig) Destroy() {
	C.MetricsConfig_Destroy(self.handle())
}

// TextBuffer_Append parameter default value hint:
// str_end: NULL
func (self ImGuiTextBuffer) Append(str string, str_end string) {
	strArg, strFin := wrapString(str)
	defer strFin()

	str_endArg, str_endFin := wrapString(str_end)
	defer str_endFin()

	C.TextBuffer_Append(self.handle(), strArg, str_endArg)
}

func (self ImGuiSettingsHandler) Destroy() {
	C.SettingsHandler_Destroy(self.handle())
}

// TextFilter_PassFilter parameter default value hint:
// text_end: NULL
func (self ImGuiTextFilter) PassFilter(text string) bool {
	textArg, textFin := wrapString(text)
	defer textFin()

	return C.TextFilter_PassFilter(self.handle(), textArg) == C.bool(true)
}

func GetVersion() string {
	return C.GoString(C.GetVersion())
}

func (self ImDrawListSharedData) SetTexUvWhitePixel(v ImVec2) {
	C.ImDrawListSharedData_SetTexUvWhitePixel(self.handle(), v.toC())
}

func (self ImDrawListSharedData) GetTexUvWhitePixel() ImVec2 {
	return newImVec2FromC(C.ImDrawListSharedData_GetTexUvWhitePixel(self.handle()))
}

func (self ImDrawListSharedData) SetFont(v ImFont) {
	C.ImDrawListSharedData_SetFont(self.handle(), v.handle())
}

func (self ImDrawListSharedData) GetFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.ImDrawListSharedData_GetFont(self.handle())))
}

func (self ImDrawListSharedData) SetFontSize(v float32) {
	C.ImDrawListSharedData_SetFontSize(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) GetFontSize() float32 {
	return float32(C.ImDrawListSharedData_GetFontSize(self.handle()))
}

func (self ImDrawListSharedData) SetCurveTessellationTol(v float32) {
	C.ImDrawListSharedData_SetCurveTessellationTol(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) GetCurveTessellationTol() float32 {
	return float32(C.ImDrawListSharedData_GetCurveTessellationTol(self.handle()))
}

func (self ImDrawListSharedData) SetCircleSegmentMaxError(v float32) {
	C.ImDrawListSharedData_SetCircleSegmentMaxError(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) GetCircleSegmentMaxError() float32 {
	return float32(C.ImDrawListSharedData_GetCircleSegmentMaxError(self.handle()))
}

func (self ImDrawListSharedData) SetClipRectFullscreen(v ImVec4) {
	C.ImDrawListSharedData_SetClipRectFullscreen(self.handle(), v.toC())
}

func (self ImDrawListSharedData) GetClipRectFullscreen() ImVec4 {
	return newImVec4FromC(C.ImDrawListSharedData_GetClipRectFullscreen(self.handle()))
}

func (self ImDrawListSharedData) SetInitialFlags(v ImDrawListFlags) {
	C.ImDrawListSharedData_SetInitialFlags(self.handle(), C.ImDrawListFlags(v))
}

func (self ImDrawListSharedData) GetInitialFlags() ImDrawListFlags {
	return ImDrawListFlags(C.ImDrawListSharedData_GetInitialFlags(self.handle()))
}

func (self ImDrawListSharedData) SetArcFastRadiusCutoff(v float32) {
	C.ImDrawListSharedData_SetArcFastRadiusCutoff(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) GetArcFastRadiusCutoff() float32 {
	return float32(C.ImDrawListSharedData_GetArcFastRadiusCutoff(self.handle()))
}

func (self ImDrawListSharedData) SetTexUvLines(v *ImVec4) {
	vArg, vFin := v.wrap()
	defer vFin()

	C.ImDrawListSharedData_SetTexUvLines(self.handle(), vArg)
}

func (self ImDrawListSharedData) GetTexUvLines() ImVec4 {
	return newImVec4FromCPtr(C.ImDrawListSharedData_GetTexUvLines(self.handle()))
}

func (self ImGuiColorMod) SetCol(v ImGuiCol) {
	C.ImGuiColorMod_SetCol(self.handle(), C.ImGuiCol(v))
}

func (self ImGuiColorMod) GetCol() ImGuiCol {
	return ImGuiCol(C.ImGuiColorMod_GetCol(self.handle()))
}

func (self ImGuiColorMod) SetBackupValue(v ImVec4) {
	C.ImGuiColorMod_SetBackupValue(self.handle(), v.toC())
}

func (self ImGuiColorMod) GetBackupValue() ImVec4 {
	return newImVec4FromC(C.ImGuiColorMod_GetBackupValue(self.handle()))
}

func (self ImGuiNextItemData) SetFlags(v ImGuiNextItemDataFlags) {
	C.ImGuiNextItemData_SetFlags(self.handle(), C.ImGuiNextItemDataFlags(v))
}

func (self ImGuiNextItemData) GetFlags() ImGuiNextItemDataFlags {
	return ImGuiNextItemDataFlags(C.ImGuiNextItemData_GetFlags(self.handle()))
}

func (self ImGuiNextItemData) SetWidth(v float32) {
	C.ImGuiNextItemData_SetWidth(self.handle(), C.float(v))
}

func (self ImGuiNextItemData) GetWidth() float32 {
	return float32(C.ImGuiNextItemData_GetWidth(self.handle()))
}

func (self ImGuiNextItemData) SetFocusScopeId(v ImGuiID) {
	C.ImGuiNextItemData_SetFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNextItemData) GetFocusScopeId() ImGuiID {
	return ImGuiID(C.ImGuiNextItemData_GetFocusScopeId(self.handle()))
}

func (self ImGuiNextItemData) SetOpenCond(v ImGuiCond) {
	C.ImGuiNextItemData_SetOpenCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextItemData) GetOpenCond() ImGuiCond {
	return ImGuiCond(C.ImGuiNextItemData_GetOpenCond(self.handle()))
}

func (self ImGuiNextItemData) SetOpenVal(v bool) {
	C.ImGuiNextItemData_SetOpenVal(self.handle(), C.bool(v))
}

func (self ImGuiNextItemData) GetOpenVal() bool {
	return C.ImGuiNextItemData_GetOpenVal(self.handle()) == C.bool(true)
}

func (self ImGuiStackSizes) SetSizeOfIDStack(v int) {
	C.ImGuiStackSizes_SetSizeOfIDStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfIDStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfIDStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfColorStack(v int) {
	C.ImGuiStackSizes_SetSizeOfColorStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfColorStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfColorStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfStyleVarStack(v int) {
	C.ImGuiStackSizes_SetSizeOfStyleVarStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfStyleVarStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfStyleVarStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfFontStack(v int) {
	C.ImGuiStackSizes_SetSizeOfFontStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfFontStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfFontStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfFocusScopeStack(v int) {
	C.ImGuiStackSizes_SetSizeOfFocusScopeStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfFocusScopeStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfFocusScopeStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfGroupStack(v int) {
	C.ImGuiStackSizes_SetSizeOfGroupStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfGroupStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfGroupStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfItemFlagsStack(v int) {
	C.ImGuiStackSizes_SetSizeOfItemFlagsStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfItemFlagsStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfItemFlagsStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfBeginPopupStack(v int) {
	C.ImGuiStackSizes_SetSizeOfBeginPopupStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfBeginPopupStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfBeginPopupStack(self.handle()))
}

func (self ImGuiStackSizes) SetSizeOfDisabledStack(v int) {
	C.ImGuiStackSizes_SetSizeOfDisabledStack(self.handle(), C.short(v))
}

func (self ImGuiStackSizes) GetSizeOfDisabledStack() int {
	return int(C.ImGuiStackSizes_GetSizeOfDisabledStack(self.handle()))
}

func (self ImGuiViewport) SetID(v ImGuiID) {
	C.ImGuiViewport_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiViewport) GetID() ImGuiID {
	return ImGuiID(C.ImGuiViewport_GetID(self.handle()))
}

func (self ImGuiViewport) SetFlags(v ImGuiViewportFlags) {
	C.ImGuiViewport_SetFlags(self.handle(), C.ImGuiViewportFlags(v))
}

func (self ImGuiViewport) GetFlags() ImGuiViewportFlags {
	return ImGuiViewportFlags(C.ImGuiViewport_GetFlags(self.handle()))
}

func (self ImGuiViewport) SetPos(v ImVec2) {
	C.ImGuiViewport_SetPos(self.handle(), v.toC())
}

func (self ImGuiViewport) GetPos() ImVec2 {
	return newImVec2FromC(C.ImGuiViewport_GetPos(self.handle()))
}

func (self ImGuiViewport) SetSize(v ImVec2) {
	C.ImGuiViewport_SetSize(self.handle(), v.toC())
}

func (self ImGuiViewport) GetSize() ImVec2 {
	return newImVec2FromC(C.ImGuiViewport_GetSize(self.handle()))
}

func (self ImGuiViewport) SetWorkPos(v ImVec2) {
	C.ImGuiViewport_SetWorkPos(self.handle(), v.toC())
}

func (self ImGuiViewport) GetWorkPos() ImVec2 {
	return newImVec2FromC(C.ImGuiViewport_GetWorkPos(self.handle()))
}

func (self ImGuiViewport) SetWorkSize(v ImVec2) {
	C.ImGuiViewport_SetWorkSize(self.handle(), v.toC())
}

func (self ImGuiViewport) GetWorkSize() ImVec2 {
	return newImVec2FromC(C.ImGuiViewport_GetWorkSize(self.handle()))
}

func (self ImGuiViewport) SetDpiScale(v float32) {
	C.ImGuiViewport_SetDpiScale(self.handle(), C.float(v))
}

func (self ImGuiViewport) GetDpiScale() float32 {
	return float32(C.ImGuiViewport_GetDpiScale(self.handle()))
}

func (self ImGuiViewport) SetParentViewportId(v ImGuiID) {
	C.ImGuiViewport_SetParentViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiViewport) GetParentViewportId() ImGuiID {
	return ImGuiID(C.ImGuiViewport_GetParentViewportId(self.handle()))
}

func (self ImGuiViewport) SetDrawData(v ImDrawData) {
	C.ImGuiViewport_SetDrawData(self.handle(), v.handle())
}

func (self ImGuiViewport) GetDrawData() ImDrawData {
	return (ImDrawData)(unsafe.Pointer(C.ImGuiViewport_GetDrawData(self.handle())))
}

func (self ImGuiViewport) SetRendererUserData(v unsafe.Pointer) {
	C.ImGuiViewport_SetRendererUserData(self.handle(), v)
}

func (self ImGuiViewport) GetRendererUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiViewport_GetRendererUserData(self.handle()))
}

func (self ImGuiViewport) SetPlatformUserData(v unsafe.Pointer) {
	C.ImGuiViewport_SetPlatformUserData(self.handle(), v)
}

func (self ImGuiViewport) GetPlatformUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiViewport_GetPlatformUserData(self.handle()))
}

func (self ImGuiViewport) SetPlatformHandle(v unsafe.Pointer) {
	C.ImGuiViewport_SetPlatformHandle(self.handle(), v)
}

func (self ImGuiViewport) GetPlatformHandle() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiViewport_GetPlatformHandle(self.handle()))
}

func (self ImGuiViewport) SetPlatformHandleRaw(v unsafe.Pointer) {
	C.ImGuiViewport_SetPlatformHandleRaw(self.handle(), v)
}

func (self ImGuiViewport) GetPlatformHandleRaw() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiViewport_GetPlatformHandleRaw(self.handle()))
}

func (self ImGuiViewport) SetPlatformRequestMove(v bool) {
	C.ImGuiViewport_SetPlatformRequestMove(self.handle(), C.bool(v))
}

func (self ImGuiViewport) GetPlatformRequestMove() bool {
	return C.ImGuiViewport_GetPlatformRequestMove(self.handle()) == C.bool(true)
}

func (self ImGuiViewport) SetPlatformRequestResize(v bool) {
	C.ImGuiViewport_SetPlatformRequestResize(self.handle(), C.bool(v))
}

func (self ImGuiViewport) GetPlatformRequestResize() bool {
	return C.ImGuiViewport_GetPlatformRequestResize(self.handle()) == C.bool(true)
}

func (self ImGuiViewport) SetPlatformRequestClose(v bool) {
	C.ImGuiViewport_SetPlatformRequestClose(self.handle(), C.bool(v))
}

func (self ImGuiViewport) GetPlatformRequestClose() bool {
	return C.ImGuiViewport_GetPlatformRequestClose(self.handle()) == C.bool(true)
}

func (self ImGuiWindowTempData) SetCursorPos(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorPos(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetCursorPos() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetCursorPos(self.handle()))
}

func (self ImGuiWindowTempData) SetCursorPosPrevLine(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorPosPrevLine(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetCursorPosPrevLine() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetCursorPosPrevLine(self.handle()))
}

func (self ImGuiWindowTempData) SetCursorStartPos(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorStartPos(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetCursorStartPos() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetCursorStartPos(self.handle()))
}

func (self ImGuiWindowTempData) SetCursorMaxPos(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorMaxPos(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetCursorMaxPos() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetCursorMaxPos(self.handle()))
}

func (self ImGuiWindowTempData) SetIdealMaxPos(v ImVec2) {
	C.ImGuiWindowTempData_SetIdealMaxPos(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetIdealMaxPos() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetIdealMaxPos(self.handle()))
}

func (self ImGuiWindowTempData) SetCurrLineSize(v ImVec2) {
	C.ImGuiWindowTempData_SetCurrLineSize(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetCurrLineSize() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetCurrLineSize(self.handle()))
}

func (self ImGuiWindowTempData) SetPrevLineSize(v ImVec2) {
	C.ImGuiWindowTempData_SetPrevLineSize(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetPrevLineSize() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetPrevLineSize(self.handle()))
}

func (self ImGuiWindowTempData) SetCurrLineTextBaseOffset(v float32) {
	C.ImGuiWindowTempData_SetCurrLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) GetCurrLineTextBaseOffset() float32 {
	return float32(C.ImGuiWindowTempData_GetCurrLineTextBaseOffset(self.handle()))
}

func (self ImGuiWindowTempData) SetPrevLineTextBaseOffset(v float32) {
	C.ImGuiWindowTempData_SetPrevLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) GetPrevLineTextBaseOffset() float32 {
	return float32(C.ImGuiWindowTempData_GetPrevLineTextBaseOffset(self.handle()))
}

func (self ImGuiWindowTempData) SetIsSameLine(v bool) {
	C.ImGuiWindowTempData_SetIsSameLine(self.handle(), C.bool(v))
}

func (self ImGuiWindowTempData) GetIsSameLine() bool {
	return C.ImGuiWindowTempData_GetIsSameLine(self.handle()) == C.bool(true)
}

func (self ImGuiWindowTempData) SetCursorStartPosLossyness(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorStartPosLossyness(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetCursorStartPosLossyness() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetCursorStartPosLossyness(self.handle()))
}

func (self ImGuiWindowTempData) SetNavLayerCurrent(v ImGuiNavLayer) {
	C.ImGuiWindowTempData_SetNavLayerCurrent(self.handle(), C.ImGuiNavLayer(v))
}

func (self ImGuiWindowTempData) GetNavLayerCurrent() ImGuiNavLayer {
	return ImGuiNavLayer(C.ImGuiWindowTempData_GetNavLayerCurrent(self.handle()))
}

func (self ImGuiWindowTempData) SetNavLayersActiveMask(v int) {
	C.ImGuiWindowTempData_SetNavLayersActiveMask(self.handle(), C.short(v))
}

func (self ImGuiWindowTempData) GetNavLayersActiveMask() int {
	return int(C.ImGuiWindowTempData_GetNavLayersActiveMask(self.handle()))
}

func (self ImGuiWindowTempData) SetNavLayersActiveMaskNext(v int) {
	C.ImGuiWindowTempData_SetNavLayersActiveMaskNext(self.handle(), C.short(v))
}

func (self ImGuiWindowTempData) GetNavLayersActiveMaskNext() int {
	return int(C.ImGuiWindowTempData_GetNavLayersActiveMaskNext(self.handle()))
}

func (self ImGuiWindowTempData) SetNavFocusScopeIdCurrent(v ImGuiID) {
	C.ImGuiWindowTempData_SetNavFocusScopeIdCurrent(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowTempData) GetNavFocusScopeIdCurrent() ImGuiID {
	return ImGuiID(C.ImGuiWindowTempData_GetNavFocusScopeIdCurrent(self.handle()))
}

func (self ImGuiWindowTempData) SetNavHideHighlightOneFrame(v bool) {
	C.ImGuiWindowTempData_SetNavHideHighlightOneFrame(self.handle(), C.bool(v))
}

func (self ImGuiWindowTempData) GetNavHideHighlightOneFrame() bool {
	return C.ImGuiWindowTempData_GetNavHideHighlightOneFrame(self.handle()) == C.bool(true)
}

func (self ImGuiWindowTempData) SetNavHasScroll(v bool) {
	C.ImGuiWindowTempData_SetNavHasScroll(self.handle(), C.bool(v))
}

func (self ImGuiWindowTempData) GetNavHasScroll() bool {
	return C.ImGuiWindowTempData_GetNavHasScroll(self.handle()) == C.bool(true)
}

func (self ImGuiWindowTempData) SetMenuBarAppending(v bool) {
	C.ImGuiWindowTempData_SetMenuBarAppending(self.handle(), C.bool(v))
}

func (self ImGuiWindowTempData) GetMenuBarAppending() bool {
	return C.ImGuiWindowTempData_GetMenuBarAppending(self.handle()) == C.bool(true)
}

func (self ImGuiWindowTempData) SetMenuBarOffset(v ImVec2) {
	C.ImGuiWindowTempData_SetMenuBarOffset(self.handle(), v.toC())
}

func (self ImGuiWindowTempData) GetMenuBarOffset() ImVec2 {
	return newImVec2FromC(C.ImGuiWindowTempData_GetMenuBarOffset(self.handle()))
}

func (self ImGuiWindowTempData) GetMenuColumns() ImGuiMenuColumns {
	return newImGuiMenuColumnsFromC(C.ImGuiWindowTempData_GetMenuColumns(self.handle()))
}

func (self ImGuiWindowTempData) SetTreeDepth(v int32) {
	C.ImGuiWindowTempData_SetTreeDepth(self.handle(), C.int(v))
}

func (self ImGuiWindowTempData) GetTreeDepth() int {
	return int(C.ImGuiWindowTempData_GetTreeDepth(self.handle()))
}

func (self ImGuiWindowTempData) SetTreeJumpToParentOnPopMask(v uint32) {
	C.ImGuiWindowTempData_SetTreeJumpToParentOnPopMask(self.handle(), C.ImU32(v))
}

func (self ImGuiWindowTempData) GetTreeJumpToParentOnPopMask() uint32 {
	return uint32(C.ImGuiWindowTempData_GetTreeJumpToParentOnPopMask(self.handle()))
}

func (self ImGuiWindowTempData) SetStateStorage(v ImGuiStorage) {
	C.ImGuiWindowTempData_SetStateStorage(self.handle(), v.handle())
}

func (self ImGuiWindowTempData) GetStateStorage() ImGuiStorage {
	return (ImGuiStorage)(unsafe.Pointer(C.ImGuiWindowTempData_GetStateStorage(self.handle())))
}

func (self ImGuiWindowTempData) SetCurrentColumns(v ImGuiOldColumns) {
	C.ImGuiWindowTempData_SetCurrentColumns(self.handle(), v.handle())
}

func (self ImGuiWindowTempData) GetCurrentColumns() ImGuiOldColumns {
	return (ImGuiOldColumns)(unsafe.Pointer(C.ImGuiWindowTempData_GetCurrentColumns(self.handle())))
}

func (self ImGuiWindowTempData) SetCurrentTableIdx(v int32) {
	C.ImGuiWindowTempData_SetCurrentTableIdx(self.handle(), C.int(v))
}

func (self ImGuiWindowTempData) GetCurrentTableIdx() int {
	return int(C.ImGuiWindowTempData_GetCurrentTableIdx(self.handle()))
}

func (self ImGuiWindowTempData) SetLayoutType(v ImGuiLayoutType) {
	C.ImGuiWindowTempData_SetLayoutType(self.handle(), C.ImGuiLayoutType(v))
}

func (self ImGuiWindowTempData) GetLayoutType() ImGuiLayoutType {
	return ImGuiLayoutType(C.ImGuiWindowTempData_GetLayoutType(self.handle()))
}

func (self ImGuiWindowTempData) SetParentLayoutType(v ImGuiLayoutType) {
	C.ImGuiWindowTempData_SetParentLayoutType(self.handle(), C.ImGuiLayoutType(v))
}

func (self ImGuiWindowTempData) GetParentLayoutType() ImGuiLayoutType {
	return ImGuiLayoutType(C.ImGuiWindowTempData_GetParentLayoutType(self.handle()))
}

func (self ImGuiWindowTempData) SetItemWidth(v float32) {
	C.ImGuiWindowTempData_SetItemWidth(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) GetItemWidth() float32 {
	return float32(C.ImGuiWindowTempData_GetItemWidth(self.handle()))
}

func (self ImGuiWindowTempData) SetTextWrapPos(v float32) {
	C.ImGuiWindowTempData_SetTextWrapPos(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) GetTextWrapPos() float32 {
	return float32(C.ImGuiWindowTempData_GetTextWrapPos(self.handle()))
}

func (self ImGuiPlatformMonitor) SetMainPos(v ImVec2) {
	C.ImGuiPlatformMonitor_SetMainPos(self.handle(), v.toC())
}

func (self ImGuiPlatformMonitor) GetMainPos() ImVec2 {
	return newImVec2FromC(C.ImGuiPlatformMonitor_GetMainPos(self.handle()))
}

func (self ImGuiPlatformMonitor) SetMainSize(v ImVec2) {
	C.ImGuiPlatformMonitor_SetMainSize(self.handle(), v.toC())
}

func (self ImGuiPlatformMonitor) GetMainSize() ImVec2 {
	return newImVec2FromC(C.ImGuiPlatformMonitor_GetMainSize(self.handle()))
}

func (self ImGuiPlatformMonitor) SetWorkPos(v ImVec2) {
	C.ImGuiPlatformMonitor_SetWorkPos(self.handle(), v.toC())
}

func (self ImGuiPlatformMonitor) GetWorkPos() ImVec2 {
	return newImVec2FromC(C.ImGuiPlatformMonitor_GetWorkPos(self.handle()))
}

func (self ImGuiPlatformMonitor) SetWorkSize(v ImVec2) {
	C.ImGuiPlatformMonitor_SetWorkSize(self.handle(), v.toC())
}

func (self ImGuiPlatformMonitor) GetWorkSize() ImVec2 {
	return newImVec2FromC(C.ImGuiPlatformMonitor_GetWorkSize(self.handle()))
}

func (self ImGuiPlatformMonitor) SetDpiScale(v float32) {
	C.ImGuiPlatformMonitor_SetDpiScale(self.handle(), C.float(v))
}

func (self ImGuiPlatformMonitor) GetDpiScale() float32 {
	return float32(C.ImGuiPlatformMonitor_GetDpiScale(self.handle()))
}

func (self ImGuiPopupData) SetPopupId(v ImGuiID) {
	C.ImGuiPopupData_SetPopupId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPopupData) GetPopupId() ImGuiID {
	return ImGuiID(C.ImGuiPopupData_GetPopupId(self.handle()))
}

func (self ImGuiPopupData) SetWindow(v ImGuiWindow) {
	C.ImGuiPopupData_SetWindow(self.handle(), v.handle())
}

func (self ImGuiPopupData) GetWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiPopupData_GetWindow(self.handle())))
}

func (self ImGuiPopupData) SetSourceWindow(v ImGuiWindow) {
	C.ImGuiPopupData_SetSourceWindow(self.handle(), v.handle())
}

func (self ImGuiPopupData) GetSourceWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiPopupData_GetSourceWindow(self.handle())))
}

func (self ImGuiPopupData) SetParentNavLayer(v int32) {
	C.ImGuiPopupData_SetParentNavLayer(self.handle(), C.int(v))
}

func (self ImGuiPopupData) GetParentNavLayer() int {
	return int(C.ImGuiPopupData_GetParentNavLayer(self.handle()))
}

func (self ImGuiPopupData) SetOpenFrameCount(v int32) {
	C.ImGuiPopupData_SetOpenFrameCount(self.handle(), C.int(v))
}

func (self ImGuiPopupData) GetOpenFrameCount() int {
	return int(C.ImGuiPopupData_GetOpenFrameCount(self.handle()))
}

func (self ImGuiPopupData) SetOpenParentId(v ImGuiID) {
	C.ImGuiPopupData_SetOpenParentId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPopupData) GetOpenParentId() ImGuiID {
	return ImGuiID(C.ImGuiPopupData_GetOpenParentId(self.handle()))
}

func (self ImGuiPopupData) SetOpenPopupPos(v ImVec2) {
	C.ImGuiPopupData_SetOpenPopupPos(self.handle(), v.toC())
}

func (self ImGuiPopupData) GetOpenPopupPos() ImVec2 {
	return newImVec2FromC(C.ImGuiPopupData_GetOpenPopupPos(self.handle()))
}

func (self ImGuiPopupData) SetOpenMousePos(v ImVec2) {
	C.ImGuiPopupData_SetOpenMousePos(self.handle(), v.toC())
}

func (self ImGuiPopupData) GetOpenMousePos() ImVec2 {
	return newImVec2FromC(C.ImGuiPopupData_GetOpenMousePos(self.handle()))
}

func (self ImGuiTableColumnSortSpecs) SetColumnUserID(v ImGuiID) {
	C.ImGuiTableColumnSortSpecs_SetColumnUserID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableColumnSortSpecs) GetColumnUserID() ImGuiID {
	return ImGuiID(C.ImGuiTableColumnSortSpecs_GetColumnUserID(self.handle()))
}

func (self ImGuiTableColumnSortSpecs) SetColumnIndex(v int) {
	C.ImGuiTableColumnSortSpecs_SetColumnIndex(self.handle(), C.ImS16(v))
}

func (self ImGuiTableColumnSortSpecs) GetColumnIndex() int {
	return int(C.ImGuiTableColumnSortSpecs_GetColumnIndex(self.handle()))
}

func (self ImGuiTableColumnSortSpecs) SetSortOrder(v int) {
	C.ImGuiTableColumnSortSpecs_SetSortOrder(self.handle(), C.ImS16(v))
}

func (self ImGuiTableColumnSortSpecs) GetSortOrder() int {
	return int(C.ImGuiTableColumnSortSpecs_GetSortOrder(self.handle()))
}

func (self ImGuiTableColumnSortSpecs) SetSortDirection(v ImGuiSortDirection) {
	C.ImGuiTableColumnSortSpecs_SetSortDirection(self.handle(), C.ImGuiSortDirection(v))
}

func (self ImGuiTableColumnSortSpecs) GetSortDirection() ImGuiSortDirection {
	return ImGuiSortDirection(C.ImGuiTableColumnSortSpecs_GetSortDirection(self.handle()))
}

func (self ImGuiComboPreviewData) SetPreviewRect(v ImRect) {
	C.ImGuiComboPreviewData_SetPreviewRect(self.handle(), v.toC())
}

func (self ImGuiComboPreviewData) GetPreviewRect() ImRect {
	return newImRectFromC(C.ImGuiComboPreviewData_GetPreviewRect(self.handle()))
}

func (self ImGuiComboPreviewData) SetBackupCursorPos(v ImVec2) {
	C.ImGuiComboPreviewData_SetBackupCursorPos(self.handle(), v.toC())
}

func (self ImGuiComboPreviewData) GetBackupCursorPos() ImVec2 {
	return newImVec2FromC(C.ImGuiComboPreviewData_GetBackupCursorPos(self.handle()))
}

func (self ImGuiComboPreviewData) SetBackupCursorMaxPos(v ImVec2) {
	C.ImGuiComboPreviewData_SetBackupCursorMaxPos(self.handle(), v.toC())
}

func (self ImGuiComboPreviewData) GetBackupCursorMaxPos() ImVec2 {
	return newImVec2FromC(C.ImGuiComboPreviewData_GetBackupCursorMaxPos(self.handle()))
}

func (self ImGuiComboPreviewData) SetBackupCursorPosPrevLine(v ImVec2) {
	C.ImGuiComboPreviewData_SetBackupCursorPosPrevLine(self.handle(), v.toC())
}

func (self ImGuiComboPreviewData) GetBackupCursorPosPrevLine() ImVec2 {
	return newImVec2FromC(C.ImGuiComboPreviewData_GetBackupCursorPosPrevLine(self.handle()))
}

func (self ImGuiComboPreviewData) SetBackupPrevLineTextBaseOffset(v float32) {
	C.ImGuiComboPreviewData_SetBackupPrevLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiComboPreviewData) GetBackupPrevLineTextBaseOffset() float32 {
	return float32(C.ImGuiComboPreviewData_GetBackupPrevLineTextBaseOffset(self.handle()))
}

func (self ImGuiComboPreviewData) SetBackupLayout(v ImGuiLayoutType) {
	C.ImGuiComboPreviewData_SetBackupLayout(self.handle(), C.ImGuiLayoutType(v))
}

func (self ImGuiComboPreviewData) GetBackupLayout() ImGuiLayoutType {
	return ImGuiLayoutType(C.ImGuiComboPreviewData_GetBackupLayout(self.handle()))
}

func (self ImGuiTableColumnSettings) SetWidthOrWeight(v float32) {
	C.ImGuiTableColumnSettings_SetWidthOrWeight(self.handle(), C.float(v))
}

func (self ImGuiTableColumnSettings) GetWidthOrWeight() float32 {
	return float32(C.ImGuiTableColumnSettings_GetWidthOrWeight(self.handle()))
}

func (self ImGuiTableColumnSettings) SetUserID(v ImGuiID) {
	C.ImGuiTableColumnSettings_SetUserID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableColumnSettings) GetUserID() ImGuiID {
	return ImGuiID(C.ImGuiTableColumnSettings_GetUserID(self.handle()))
}

func (self ImGuiTableColumnSettings) SetIndex(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumnSettings_SetIndex(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumnSettings) GetIndex() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumnSettings_GetIndex(self.handle()))
}

func (self ImGuiTableColumnSettings) SetDisplayOrder(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumnSettings_SetDisplayOrder(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumnSettings) GetDisplayOrder() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumnSettings_GetDisplayOrder(self.handle()))
}

func (self ImGuiTableColumnSettings) SetSortOrder(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumnSettings_SetSortOrder(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumnSettings) GetSortOrder() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumnSettings_GetSortOrder(self.handle()))
}

func (self ImGuiTableColumnSettings) SetSortDirection(v uint) {
	C.ImGuiTableColumnSettings_SetSortDirection(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumnSettings) GetSortDirection() uint32 {
	return uint32(C.ImGuiTableColumnSettings_GetSortDirection(self.handle()))
}

func (self ImGuiTableColumnSettings) SetIsEnabled(v uint) {
	C.ImGuiTableColumnSettings_SetIsEnabled(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumnSettings) GetIsEnabled() uint32 {
	return uint32(C.ImGuiTableColumnSettings_GetIsEnabled(self.handle()))
}

func (self ImGuiTableColumnSettings) SetIsStretch(v uint) {
	C.ImGuiTableColumnSettings_SetIsStretch(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumnSettings) GetIsStretch() uint32 {
	return uint32(C.ImGuiTableColumnSettings_GetIsStretch(self.handle()))
}

func (self ImFontGlyph) SetColored(v uint32) {
	C.ImFontGlyph_SetColored(self.handle(), C.uint(v))
}

func (self ImFontGlyph) GetColored() uint32 {
	return uint32(C.ImFontGlyph_GetColored(self.handle()))
}

func (self ImFontGlyph) SetVisible(v uint32) {
	C.ImFontGlyph_SetVisible(self.handle(), C.uint(v))
}

func (self ImFontGlyph) GetVisible() uint32 {
	return uint32(C.ImFontGlyph_GetVisible(self.handle()))
}

func (self ImFontGlyph) SetCodepoint(v uint32) {
	C.ImFontGlyph_SetCodepoint(self.handle(), C.uint(v))
}

func (self ImFontGlyph) GetCodepoint() uint32 {
	return uint32(C.ImFontGlyph_GetCodepoint(self.handle()))
}

func (self ImFontGlyph) SetAdvanceX(v float32) {
	C.ImFontGlyph_SetAdvanceX(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetAdvanceX() float32 {
	return float32(C.ImFontGlyph_GetAdvanceX(self.handle()))
}

func (self ImFontGlyph) SetX0(v float32) {
	C.ImFontGlyph_SetX0(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetX0() float32 {
	return float32(C.ImFontGlyph_GetX0(self.handle()))
}

func (self ImFontGlyph) SetY0(v float32) {
	C.ImFontGlyph_SetY0(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetY0() float32 {
	return float32(C.ImFontGlyph_GetY0(self.handle()))
}

func (self ImFontGlyph) SetX1(v float32) {
	C.ImFontGlyph_SetX1(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetX1() float32 {
	return float32(C.ImFontGlyph_GetX1(self.handle()))
}

func (self ImFontGlyph) SetY1(v float32) {
	C.ImFontGlyph_SetY1(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetY1() float32 {
	return float32(C.ImFontGlyph_GetY1(self.handle()))
}

func (self ImFontGlyph) SetU0(v float32) {
	C.ImFontGlyph_SetU0(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetU0() float32 {
	return float32(C.ImFontGlyph_GetU0(self.handle()))
}

func (self ImFontGlyph) SetV0(v float32) {
	C.ImFontGlyph_SetV0(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetV0() float32 {
	return float32(C.ImFontGlyph_GetV0(self.handle()))
}

func (self ImFontGlyph) SetU1(v float32) {
	C.ImFontGlyph_SetU1(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetU1() float32 {
	return float32(C.ImFontGlyph_GetU1(self.handle()))
}

func (self ImFontGlyph) SetV1(v float32) {
	C.ImFontGlyph_SetV1(self.handle(), C.float(v))
}

func (self ImFontGlyph) GetV1() float32 {
	return float32(C.ImFontGlyph_GetV1(self.handle()))
}

func (self ImGuiInputEventMouseViewport) SetHoveredViewportID(v ImGuiID) {
	C.ImGuiInputEventMouseViewport_SetHoveredViewportID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiInputEventMouseViewport) GetHoveredViewportID() ImGuiID {
	return ImGuiID(C.ImGuiInputEventMouseViewport_GetHoveredViewportID(self.handle()))
}

func (self ImGuiListClipperRange) SetMin(v int32) {
	C.ImGuiListClipperRange_SetMin(self.handle(), C.int(v))
}

func (self ImGuiListClipperRange) GetMin() int {
	return int(C.ImGuiListClipperRange_GetMin(self.handle()))
}

func (self ImGuiListClipperRange) SetMax(v int32) {
	C.ImGuiListClipperRange_SetMax(self.handle(), C.int(v))
}

func (self ImGuiListClipperRange) GetMax() int {
	return int(C.ImGuiListClipperRange_GetMax(self.handle()))
}

func (self ImGuiListClipperRange) SetPosToIndexConvert(v bool) {
	C.ImGuiListClipperRange_SetPosToIndexConvert(self.handle(), C.bool(v))
}

func (self ImGuiListClipperRange) GetPosToIndexConvert() bool {
	return C.ImGuiListClipperRange_GetPosToIndexConvert(self.handle()) == C.bool(true)
}

func (self ImGuiListClipperRange) SetPosToIndexOffsetMin(v int) {
	C.ImGuiListClipperRange_SetPosToIndexOffsetMin(self.handle(), C.ImS8(v))
}

func (self ImGuiListClipperRange) GetPosToIndexOffsetMin() int {
	return int(C.ImGuiListClipperRange_GetPosToIndexOffsetMin(self.handle()))
}

func (self ImGuiListClipperRange) SetPosToIndexOffsetMax(v int) {
	C.ImGuiListClipperRange_SetPosToIndexOffsetMax(self.handle(), C.ImS8(v))
}

func (self ImGuiListClipperRange) GetPosToIndexOffsetMax() int {
	return int(C.ImGuiListClipperRange_GetPosToIndexOffsetMax(self.handle()))
}

func (self ImGuiOldColumnData) SetOffsetNorm(v float32) {
	C.ImGuiOldColumnData_SetOffsetNorm(self.handle(), C.float(v))
}

func (self ImGuiOldColumnData) GetOffsetNorm() float32 {
	return float32(C.ImGuiOldColumnData_GetOffsetNorm(self.handle()))
}

func (self ImGuiOldColumnData) SetOffsetNormBeforeResize(v float32) {
	C.ImGuiOldColumnData_SetOffsetNormBeforeResize(self.handle(), C.float(v))
}

func (self ImGuiOldColumnData) GetOffsetNormBeforeResize() float32 {
	return float32(C.ImGuiOldColumnData_GetOffsetNormBeforeResize(self.handle()))
}

func (self ImGuiOldColumnData) SetFlags(v ImGuiOldColumnFlags) {
	C.ImGuiOldColumnData_SetFlags(self.handle(), C.ImGuiOldColumnFlags(v))
}

func (self ImGuiOldColumnData) GetFlags() ImGuiOldColumnFlags {
	return ImGuiOldColumnFlags(C.ImGuiOldColumnData_GetFlags(self.handle()))
}

func (self ImGuiOldColumnData) SetClipRect(v ImRect) {
	C.ImGuiOldColumnData_SetClipRect(self.handle(), v.toC())
}

func (self ImGuiOldColumnData) GetClipRect() ImRect {
	return newImRectFromC(C.ImGuiOldColumnData_GetClipRect(self.handle()))
}

func (self ImGuiPlatformImeData) SetWantVisible(v bool) {
	C.ImGuiPlatformImeData_SetWantVisible(self.handle(), C.bool(v))
}

func (self ImGuiPlatformImeData) GetWantVisible() bool {
	return C.ImGuiPlatformImeData_GetWantVisible(self.handle()) == C.bool(true)
}

func (self ImGuiPlatformImeData) SetInputPos(v ImVec2) {
	C.ImGuiPlatformImeData_SetInputPos(self.handle(), v.toC())
}

func (self ImGuiPlatformImeData) GetInputPos() ImVec2 {
	return newImVec2FromC(C.ImGuiPlatformImeData_GetInputPos(self.handle()))
}

func (self ImGuiPlatformImeData) SetInputLineHeight(v float32) {
	C.ImGuiPlatformImeData_SetInputLineHeight(self.handle(), C.float(v))
}

func (self ImGuiPlatformImeData) GetInputLineHeight() float32 {
	return float32(C.ImGuiPlatformImeData_GetInputLineHeight(self.handle()))
}

func (self ImGuiStyleMod) SetVarIdx(v ImGuiStyleVar) {
	C.ImGuiStyleMod_SetVarIdx(self.handle(), C.ImGuiStyleVar(v))
}

func (self ImGuiStyleMod) GetVarIdx() ImGuiStyleVar {
	return ImGuiStyleVar(C.ImGuiStyleMod_GetVarIdx(self.handle()))
}

func (self ImGuiTableInstanceData) SetLastOuterHeight(v float32) {
	C.ImGuiTableInstanceData_SetLastOuterHeight(self.handle(), C.float(v))
}

func (self ImGuiTableInstanceData) GetLastOuterHeight() float32 {
	return float32(C.ImGuiTableInstanceData_GetLastOuterHeight(self.handle()))
}

func (self ImGuiTableInstanceData) SetLastFirstRowHeight(v float32) {
	C.ImGuiTableInstanceData_SetLastFirstRowHeight(self.handle(), C.float(v))
}

func (self ImGuiTableInstanceData) GetLastFirstRowHeight() float32 {
	return float32(C.ImGuiTableInstanceData_GetLastFirstRowHeight(self.handle()))
}

func (self ImGuiTextRange) Setb(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiTextRange_Setb(self.handle(), vArg)
}

func (self ImGuiTextRange) Getb() string {
	return C.GoString(C.ImGuiTextRange_Getb(self.handle()))
}

func (self ImGuiTextRange) Sete(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiTextRange_Sete(self.handle(), vArg)
}

func (self ImGuiTextRange) Gete() string {
	return C.GoString(C.ImGuiTextRange_Gete(self.handle()))
}

func (self ImGuiWindowClass) SetClassId(v ImGuiID) {
	C.ImGuiWindowClass_SetClassId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowClass) GetClassId() ImGuiID {
	return ImGuiID(C.ImGuiWindowClass_GetClassId(self.handle()))
}

func (self ImGuiWindowClass) SetParentViewportId(v ImGuiID) {
	C.ImGuiWindowClass_SetParentViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowClass) GetParentViewportId() ImGuiID {
	return ImGuiID(C.ImGuiWindowClass_GetParentViewportId(self.handle()))
}

func (self ImGuiWindowClass) SetViewportFlagsOverrideSet(v ImGuiViewportFlags) {
	C.ImGuiWindowClass_SetViewportFlagsOverrideSet(self.handle(), C.ImGuiViewportFlags(v))
}

func (self ImGuiWindowClass) GetViewportFlagsOverrideSet() ImGuiViewportFlags {
	return ImGuiViewportFlags(C.ImGuiWindowClass_GetViewportFlagsOverrideSet(self.handle()))
}

func (self ImGuiWindowClass) SetViewportFlagsOverrideClear(v ImGuiViewportFlags) {
	C.ImGuiWindowClass_SetViewportFlagsOverrideClear(self.handle(), C.ImGuiViewportFlags(v))
}

func (self ImGuiWindowClass) GetViewportFlagsOverrideClear() ImGuiViewportFlags {
	return ImGuiViewportFlags(C.ImGuiWindowClass_GetViewportFlagsOverrideClear(self.handle()))
}

func (self ImGuiWindowClass) SetTabItemFlagsOverrideSet(v ImGuiTabItemFlags) {
	C.ImGuiWindowClass_SetTabItemFlagsOverrideSet(self.handle(), C.ImGuiTabItemFlags(v))
}

func (self ImGuiWindowClass) GetTabItemFlagsOverrideSet() ImGuiTabItemFlags {
	return ImGuiTabItemFlags(C.ImGuiWindowClass_GetTabItemFlagsOverrideSet(self.handle()))
}

func (self ImGuiWindowClass) SetDockNodeFlagsOverrideSet(v ImGuiDockNodeFlags) {
	C.ImGuiWindowClass_SetDockNodeFlagsOverrideSet(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiWindowClass) GetDockNodeFlagsOverrideSet() ImGuiDockNodeFlags {
	return ImGuiDockNodeFlags(C.ImGuiWindowClass_GetDockNodeFlagsOverrideSet(self.handle()))
}

func (self ImGuiWindowClass) SetDockingAlwaysTabBar(v bool) {
	C.ImGuiWindowClass_SetDockingAlwaysTabBar(self.handle(), C.bool(v))
}

func (self ImGuiWindowClass) GetDockingAlwaysTabBar() bool {
	return C.ImGuiWindowClass_GetDockingAlwaysTabBar(self.handle()) == C.bool(true)
}

func (self ImGuiWindowClass) SetDockingAllowUnclassed(v bool) {
	C.ImGuiWindowClass_SetDockingAllowUnclassed(self.handle(), C.bool(v))
}

func (self ImGuiWindowClass) GetDockingAllowUnclassed() bool {
	return C.ImGuiWindowClass_GetDockingAllowUnclassed(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigFlags(v ImGuiConfigFlags) {
	C.ImGuiIO_SetConfigFlags(self.handle(), C.ImGuiConfigFlags(v))
}

func (self ImGuiIO) GetConfigFlags() ImGuiConfigFlags {
	return ImGuiConfigFlags(C.ImGuiIO_GetConfigFlags(self.handle()))
}

func (self ImGuiIO) SetBackendFlags(v ImGuiBackendFlags) {
	C.ImGuiIO_SetBackendFlags(self.handle(), C.ImGuiBackendFlags(v))
}

func (self ImGuiIO) GetBackendFlags() ImGuiBackendFlags {
	return ImGuiBackendFlags(C.ImGuiIO_GetBackendFlags(self.handle()))
}

func (self ImGuiIO) SetDisplaySize(v ImVec2) {
	C.ImGuiIO_SetDisplaySize(self.handle(), v.toC())
}

func (self ImGuiIO) GetDisplaySize() ImVec2 {
	return newImVec2FromC(C.ImGuiIO_GetDisplaySize(self.handle()))
}

func (self ImGuiIO) SetDeltaTime(v float32) {
	C.ImGuiIO_SetDeltaTime(self.handle(), C.float(v))
}

func (self ImGuiIO) GetDeltaTime() float32 {
	return float32(C.ImGuiIO_GetDeltaTime(self.handle()))
}

func (self ImGuiIO) SetIniSavingRate(v float32) {
	C.ImGuiIO_SetIniSavingRate(self.handle(), C.float(v))
}

func (self ImGuiIO) GetIniSavingRate() float32 {
	return float32(C.ImGuiIO_GetIniSavingRate(self.handle()))
}

func (self ImGuiIO) SetIniFilename(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetIniFilename(self.handle(), vArg)
}

func (self ImGuiIO) GetIniFilename() string {
	return C.GoString(C.ImGuiIO_GetIniFilename(self.handle()))
}

func (self ImGuiIO) SetLogFilename(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetLogFilename(self.handle(), vArg)
}

func (self ImGuiIO) GetLogFilename() string {
	return C.GoString(C.ImGuiIO_GetLogFilename(self.handle()))
}

func (self ImGuiIO) SetMouseDoubleClickTime(v float32) {
	C.ImGuiIO_SetMouseDoubleClickTime(self.handle(), C.float(v))
}

func (self ImGuiIO) GetMouseDoubleClickTime() float32 {
	return float32(C.ImGuiIO_GetMouseDoubleClickTime(self.handle()))
}

func (self ImGuiIO) SetMouseDoubleClickMaxDist(v float32) {
	C.ImGuiIO_SetMouseDoubleClickMaxDist(self.handle(), C.float(v))
}

func (self ImGuiIO) GetMouseDoubleClickMaxDist() float32 {
	return float32(C.ImGuiIO_GetMouseDoubleClickMaxDist(self.handle()))
}

func (self ImGuiIO) SetMouseDragThreshold(v float32) {
	C.ImGuiIO_SetMouseDragThreshold(self.handle(), C.float(v))
}

func (self ImGuiIO) GetMouseDragThreshold() float32 {
	return float32(C.ImGuiIO_GetMouseDragThreshold(self.handle()))
}

func (self ImGuiIO) SetKeyRepeatDelay(v float32) {
	C.ImGuiIO_SetKeyRepeatDelay(self.handle(), C.float(v))
}

func (self ImGuiIO) GetKeyRepeatDelay() float32 {
	return float32(C.ImGuiIO_GetKeyRepeatDelay(self.handle()))
}

func (self ImGuiIO) SetKeyRepeatRate(v float32) {
	C.ImGuiIO_SetKeyRepeatRate(self.handle(), C.float(v))
}

func (self ImGuiIO) GetKeyRepeatRate() float32 {
	return float32(C.ImGuiIO_GetKeyRepeatRate(self.handle()))
}

func (self ImGuiIO) SetUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetUserData(self.handle(), v)
}

func (self ImGuiIO) GetUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiIO_GetUserData(self.handle()))
}

func (self ImGuiIO) SetFonts(v ImFontAtlas) {
	C.ImGuiIO_SetFonts(self.handle(), v.handle())
}

func (self ImGuiIO) GetFonts() ImFontAtlas {
	return (ImFontAtlas)(unsafe.Pointer(C.ImGuiIO_GetFonts(self.handle())))
}

func (self ImGuiIO) SetFontGlobalScale(v float32) {
	C.ImGuiIO_SetFontGlobalScale(self.handle(), C.float(v))
}

func (self ImGuiIO) GetFontGlobalScale() float32 {
	return float32(C.ImGuiIO_GetFontGlobalScale(self.handle()))
}

func (self ImGuiIO) SetFontAllowUserScaling(v bool) {
	C.ImGuiIO_SetFontAllowUserScaling(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetFontAllowUserScaling() bool {
	return C.ImGuiIO_GetFontAllowUserScaling(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetFontDefault(v ImFont) {
	C.ImGuiIO_SetFontDefault(self.handle(), v.handle())
}

func (self ImGuiIO) GetFontDefault() ImFont {
	return (ImFont)(unsafe.Pointer(C.ImGuiIO_GetFontDefault(self.handle())))
}

func (self ImGuiIO) SetDisplayFramebufferScale(v ImVec2) {
	C.ImGuiIO_SetDisplayFramebufferScale(self.handle(), v.toC())
}

func (self ImGuiIO) GetDisplayFramebufferScale() ImVec2 {
	return newImVec2FromC(C.ImGuiIO_GetDisplayFramebufferScale(self.handle()))
}

func (self ImGuiIO) SetConfigDockingNoSplit(v bool) {
	C.ImGuiIO_SetConfigDockingNoSplit(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigDockingNoSplit() bool {
	return C.ImGuiIO_GetConfigDockingNoSplit(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigDockingWithShift(v bool) {
	C.ImGuiIO_SetConfigDockingWithShift(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigDockingWithShift() bool {
	return C.ImGuiIO_GetConfigDockingWithShift(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigDockingAlwaysTabBar(v bool) {
	C.ImGuiIO_SetConfigDockingAlwaysTabBar(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigDockingAlwaysTabBar() bool {
	return C.ImGuiIO_GetConfigDockingAlwaysTabBar(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigDockingTransparentPayload(v bool) {
	C.ImGuiIO_SetConfigDockingTransparentPayload(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigDockingTransparentPayload() bool {
	return C.ImGuiIO_GetConfigDockingTransparentPayload(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigViewportsNoAutoMerge(v bool) {
	C.ImGuiIO_SetConfigViewportsNoAutoMerge(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigViewportsNoAutoMerge() bool {
	return C.ImGuiIO_GetConfigViewportsNoAutoMerge(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigViewportsNoTaskBarIcon(v bool) {
	C.ImGuiIO_SetConfigViewportsNoTaskBarIcon(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigViewportsNoTaskBarIcon() bool {
	return C.ImGuiIO_GetConfigViewportsNoTaskBarIcon(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigViewportsNoDecoration(v bool) {
	C.ImGuiIO_SetConfigViewportsNoDecoration(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigViewportsNoDecoration() bool {
	return C.ImGuiIO_GetConfigViewportsNoDecoration(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigViewportsNoDefaultParent(v bool) {
	C.ImGuiIO_SetConfigViewportsNoDefaultParent(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigViewportsNoDefaultParent() bool {
	return C.ImGuiIO_GetConfigViewportsNoDefaultParent(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetMouseDrawCursor(v bool) {
	C.ImGuiIO_SetMouseDrawCursor(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetMouseDrawCursor() bool {
	return C.ImGuiIO_GetMouseDrawCursor(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigMacOSXBehaviors(v bool) {
	C.ImGuiIO_SetConfigMacOSXBehaviors(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigMacOSXBehaviors() bool {
	return C.ImGuiIO_GetConfigMacOSXBehaviors(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigInputTrickleEventQueue(v bool) {
	C.ImGuiIO_SetConfigInputTrickleEventQueue(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigInputTrickleEventQueue() bool {
	return C.ImGuiIO_GetConfigInputTrickleEventQueue(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigInputTextCursorBlink(v bool) {
	C.ImGuiIO_SetConfigInputTextCursorBlink(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigInputTextCursorBlink() bool {
	return C.ImGuiIO_GetConfigInputTextCursorBlink(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigInputTextEnterKeepActive(v bool) {
	C.ImGuiIO_SetConfigInputTextEnterKeepActive(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigInputTextEnterKeepActive() bool {
	return C.ImGuiIO_GetConfigInputTextEnterKeepActive(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigDragClickToInputText(v bool) {
	C.ImGuiIO_SetConfigDragClickToInputText(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigDragClickToInputText() bool {
	return C.ImGuiIO_GetConfigDragClickToInputText(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigWindowsResizeFromEdges(v bool) {
	C.ImGuiIO_SetConfigWindowsResizeFromEdges(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigWindowsResizeFromEdges() bool {
	return C.ImGuiIO_GetConfigWindowsResizeFromEdges(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigWindowsMoveFromTitleBarOnly(v bool) {
	C.ImGuiIO_SetConfigWindowsMoveFromTitleBarOnly(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetConfigWindowsMoveFromTitleBarOnly() bool {
	return C.ImGuiIO_GetConfigWindowsMoveFromTitleBarOnly(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetConfigMemoryCompactTimer(v float32) {
	C.ImGuiIO_SetConfigMemoryCompactTimer(self.handle(), C.float(v))
}

func (self ImGuiIO) GetConfigMemoryCompactTimer() float32 {
	return float32(C.ImGuiIO_GetConfigMemoryCompactTimer(self.handle()))
}

func (self ImGuiIO) SetBackendPlatformName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetBackendPlatformName(self.handle(), vArg)
}

func (self ImGuiIO) GetBackendPlatformName() string {
	return C.GoString(C.ImGuiIO_GetBackendPlatformName(self.handle()))
}

func (self ImGuiIO) SetBackendRendererName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetBackendRendererName(self.handle(), vArg)
}

func (self ImGuiIO) GetBackendRendererName() string {
	return C.GoString(C.ImGuiIO_GetBackendRendererName(self.handle()))
}

func (self ImGuiIO) SetBackendPlatformUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetBackendPlatformUserData(self.handle(), v)
}

func (self ImGuiIO) GetBackendPlatformUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiIO_GetBackendPlatformUserData(self.handle()))
}

func (self ImGuiIO) SetBackendRendererUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetBackendRendererUserData(self.handle(), v)
}

func (self ImGuiIO) GetBackendRendererUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiIO_GetBackendRendererUserData(self.handle()))
}

func (self ImGuiIO) SetBackendLanguageUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetBackendLanguageUserData(self.handle(), v)
}

func (self ImGuiIO) GetBackendLanguageUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiIO_GetBackendLanguageUserData(self.handle()))
}

func (self ImGuiIO) SetClipboardUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetClipboardUserData(self.handle(), v)
}

func (self ImGuiIO) GetClipboardUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiIO_GetClipboardUserData(self.handle()))
}

func (self ImGuiIO) Set_UnusedPadding(v unsafe.Pointer) {
	C.ImGuiIO_Set_UnusedPadding(self.handle(), v)
}

func (self ImGuiIO) Get_UnusedPadding() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiIO_Get_UnusedPadding(self.handle()))
}

func (self ImGuiIO) SetWantCaptureMouse(v bool) {
	C.ImGuiIO_SetWantCaptureMouse(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetWantCaptureMouse() bool {
	return C.ImGuiIO_GetWantCaptureMouse(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetWantCaptureKeyboard(v bool) {
	C.ImGuiIO_SetWantCaptureKeyboard(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetWantCaptureKeyboard() bool {
	return C.ImGuiIO_GetWantCaptureKeyboard(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetWantTextInput(v bool) {
	C.ImGuiIO_SetWantTextInput(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetWantTextInput() bool {
	return C.ImGuiIO_GetWantTextInput(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetWantSetMousePos(v bool) {
	C.ImGuiIO_SetWantSetMousePos(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetWantSetMousePos() bool {
	return C.ImGuiIO_GetWantSetMousePos(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetWantSaveIniSettings(v bool) {
	C.ImGuiIO_SetWantSaveIniSettings(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetWantSaveIniSettings() bool {
	return C.ImGuiIO_GetWantSaveIniSettings(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetNavActive(v bool) {
	C.ImGuiIO_SetNavActive(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetNavActive() bool {
	return C.ImGuiIO_GetNavActive(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetNavVisible(v bool) {
	C.ImGuiIO_SetNavVisible(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetNavVisible() bool {
	return C.ImGuiIO_GetNavVisible(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetFramerate(v float32) {
	C.ImGuiIO_SetFramerate(self.handle(), C.float(v))
}

func (self ImGuiIO) GetFramerate() float32 {
	return float32(C.ImGuiIO_GetFramerate(self.handle()))
}

func (self ImGuiIO) SetMetricsRenderVertices(v int32) {
	C.ImGuiIO_SetMetricsRenderVertices(self.handle(), C.int(v))
}

func (self ImGuiIO) GetMetricsRenderVertices() int {
	return int(C.ImGuiIO_GetMetricsRenderVertices(self.handle()))
}

func (self ImGuiIO) SetMetricsRenderIndices(v int32) {
	C.ImGuiIO_SetMetricsRenderIndices(self.handle(), C.int(v))
}

func (self ImGuiIO) GetMetricsRenderIndices() int {
	return int(C.ImGuiIO_GetMetricsRenderIndices(self.handle()))
}

func (self ImGuiIO) SetMetricsRenderWindows(v int32) {
	C.ImGuiIO_SetMetricsRenderWindows(self.handle(), C.int(v))
}

func (self ImGuiIO) GetMetricsRenderWindows() int {
	return int(C.ImGuiIO_GetMetricsRenderWindows(self.handle()))
}

func (self ImGuiIO) SetMetricsActiveWindows(v int32) {
	C.ImGuiIO_SetMetricsActiveWindows(self.handle(), C.int(v))
}

func (self ImGuiIO) GetMetricsActiveWindows() int {
	return int(C.ImGuiIO_GetMetricsActiveWindows(self.handle()))
}

func (self ImGuiIO) SetMetricsActiveAllocations(v int32) {
	C.ImGuiIO_SetMetricsActiveAllocations(self.handle(), C.int(v))
}

func (self ImGuiIO) GetMetricsActiveAllocations() int {
	return int(C.ImGuiIO_GetMetricsActiveAllocations(self.handle()))
}

func (self ImGuiIO) SetMouseDelta(v ImVec2) {
	C.ImGuiIO_SetMouseDelta(self.handle(), v.toC())
}

func (self ImGuiIO) GetMouseDelta() ImVec2 {
	return newImVec2FromC(C.ImGuiIO_GetMouseDelta(self.handle()))
}

func (self ImGuiIO) SetMousePos(v ImVec2) {
	C.ImGuiIO_SetMousePos(self.handle(), v.toC())
}

func (self ImGuiIO) GetMousePos() ImVec2 {
	return newImVec2FromC(C.ImGuiIO_GetMousePos(self.handle()))
}

func (self ImGuiIO) SetMouseWheel(v float32) {
	C.ImGuiIO_SetMouseWheel(self.handle(), C.float(v))
}

func (self ImGuiIO) GetMouseWheel() float32 {
	return float32(C.ImGuiIO_GetMouseWheel(self.handle()))
}

func (self ImGuiIO) SetMouseWheelH(v float32) {
	C.ImGuiIO_SetMouseWheelH(self.handle(), C.float(v))
}

func (self ImGuiIO) GetMouseWheelH() float32 {
	return float32(C.ImGuiIO_GetMouseWheelH(self.handle()))
}

func (self ImGuiIO) SetMouseHoveredViewport(v ImGuiID) {
	C.ImGuiIO_SetMouseHoveredViewport(self.handle(), C.ImGuiID(v))
}

func (self ImGuiIO) GetMouseHoveredViewport() ImGuiID {
	return ImGuiID(C.ImGuiIO_GetMouseHoveredViewport(self.handle()))
}

func (self ImGuiIO) SetKeyCtrl(v bool) {
	C.ImGuiIO_SetKeyCtrl(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetKeyCtrl() bool {
	return C.ImGuiIO_GetKeyCtrl(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetKeyShift(v bool) {
	C.ImGuiIO_SetKeyShift(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetKeyShift() bool {
	return C.ImGuiIO_GetKeyShift(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetKeyAlt(v bool) {
	C.ImGuiIO_SetKeyAlt(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetKeyAlt() bool {
	return C.ImGuiIO_GetKeyAlt(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetKeySuper(v bool) {
	C.ImGuiIO_SetKeySuper(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetKeySuper() bool {
	return C.ImGuiIO_GetKeySuper(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetKeyMods(v ImGuiModFlags) {
	C.ImGuiIO_SetKeyMods(self.handle(), C.ImGuiModFlags(v))
}

func (self ImGuiIO) GetKeyMods() ImGuiModFlags {
	return ImGuiModFlags(C.ImGuiIO_GetKeyMods(self.handle()))
}

func (self ImGuiIO) SetWantCaptureMouseUnlessPopupClose(v bool) {
	C.ImGuiIO_SetWantCaptureMouseUnlessPopupClose(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetWantCaptureMouseUnlessPopupClose() bool {
	return C.ImGuiIO_GetWantCaptureMouseUnlessPopupClose(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetMousePosPrev(v ImVec2) {
	C.ImGuiIO_SetMousePosPrev(self.handle(), v.toC())
}

func (self ImGuiIO) GetMousePosPrev() ImVec2 {
	return newImVec2FromC(C.ImGuiIO_GetMousePosPrev(self.handle()))
}

func (self ImGuiIO) SetPenPressure(v float32) {
	C.ImGuiIO_SetPenPressure(self.handle(), C.float(v))
}

func (self ImGuiIO) GetPenPressure() float32 {
	return float32(C.ImGuiIO_GetPenPressure(self.handle()))
}

func (self ImGuiIO) SetAppFocusLost(v bool) {
	C.ImGuiIO_SetAppFocusLost(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetAppFocusLost() bool {
	return C.ImGuiIO_GetAppFocusLost(self.handle()) == C.bool(true)
}

func (self ImGuiIO) SetBackendUsingLegacyKeyArrays(v int) {
	C.ImGuiIO_SetBackendUsingLegacyKeyArrays(self.handle(), C.ImS8(v))
}

func (self ImGuiIO) GetBackendUsingLegacyKeyArrays() int {
	return int(C.ImGuiIO_GetBackendUsingLegacyKeyArrays(self.handle()))
}

func (self ImGuiIO) SetBackendUsingLegacyNavInputArray(v bool) {
	C.ImGuiIO_SetBackendUsingLegacyNavInputArray(self.handle(), C.bool(v))
}

func (self ImGuiIO) GetBackendUsingLegacyNavInputArray() bool {
	return C.ImGuiIO_GetBackendUsingLegacyNavInputArray(self.handle()) == C.bool(true)
}

func (self ImGuiListClipperData) SetListClipper(v ImGuiListClipper) {
	C.ImGuiListClipperData_SetListClipper(self.handle(), v.handle())
}

func (self ImGuiListClipperData) GetListClipper() ImGuiListClipper {
	return (ImGuiListClipper)(unsafe.Pointer(C.ImGuiListClipperData_GetListClipper(self.handle())))
}

func (self ImGuiListClipperData) SetLossynessOffset(v float32) {
	C.ImGuiListClipperData_SetLossynessOffset(self.handle(), C.float(v))
}

func (self ImGuiListClipperData) GetLossynessOffset() float32 {
	return float32(C.ImGuiListClipperData_GetLossynessOffset(self.handle()))
}

func (self ImGuiListClipperData) SetStepNo(v int32) {
	C.ImGuiListClipperData_SetStepNo(self.handle(), C.int(v))
}

func (self ImGuiListClipperData) GetStepNo() int {
	return int(C.ImGuiListClipperData_GetStepNo(self.handle()))
}

func (self ImGuiListClipperData) SetItemsFrozen(v int32) {
	C.ImGuiListClipperData_SetItemsFrozen(self.handle(), C.int(v))
}

func (self ImGuiListClipperData) GetItemsFrozen() int {
	return int(C.ImGuiListClipperData_GetItemsFrozen(self.handle()))
}

func (self ImDrawData) SetValid(v bool) {
	C.ImDrawData_SetValid(self.handle(), C.bool(v))
}

func (self ImDrawData) GetValid() bool {
	return C.ImDrawData_GetValid(self.handle()) == C.bool(true)
}

func (self ImDrawData) SetCmdListsCount(v int32) {
	C.ImDrawData_SetCmdListsCount(self.handle(), C.int(v))
}

func (self ImDrawData) GetCmdListsCount() int {
	return int(C.ImDrawData_GetCmdListsCount(self.handle()))
}

func (self ImDrawData) SetTotalIdxCount(v int32) {
	C.ImDrawData_SetTotalIdxCount(self.handle(), C.int(v))
}

func (self ImDrawData) GetTotalIdxCount() int {
	return int(C.ImDrawData_GetTotalIdxCount(self.handle()))
}

func (self ImDrawData) SetTotalVtxCount(v int32) {
	C.ImDrawData_SetTotalVtxCount(self.handle(), C.int(v))
}

func (self ImDrawData) GetTotalVtxCount() int {
	return int(C.ImDrawData_GetTotalVtxCount(self.handle()))
}

func (self ImDrawData) SetDisplayPos(v ImVec2) {
	C.ImDrawData_SetDisplayPos(self.handle(), v.toC())
}

func (self ImDrawData) GetDisplayPos() ImVec2 {
	return newImVec2FromC(C.ImDrawData_GetDisplayPos(self.handle()))
}

func (self ImDrawData) SetDisplaySize(v ImVec2) {
	C.ImDrawData_SetDisplaySize(self.handle(), v.toC())
}

func (self ImDrawData) GetDisplaySize() ImVec2 {
	return newImVec2FromC(C.ImDrawData_GetDisplaySize(self.handle()))
}

func (self ImDrawData) SetFramebufferScale(v ImVec2) {
	C.ImDrawData_SetFramebufferScale(self.handle(), v.toC())
}

func (self ImDrawData) GetFramebufferScale() ImVec2 {
	return newImVec2FromC(C.ImDrawData_GetFramebufferScale(self.handle()))
}

func (self ImDrawData) SetOwnerViewport(v ImGuiViewport) {
	C.ImDrawData_SetOwnerViewport(self.handle(), v.handle())
}

func (self ImDrawData) GetOwnerViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.ImDrawData_GetOwnerViewport(self.handle())))
}

func (self ImGuiInputTextCallbackData) SetEventFlag(v ImGuiInputTextFlags) {
	C.ImGuiInputTextCallbackData_SetEventFlag(self.handle(), C.ImGuiInputTextFlags(v))
}

func (self ImGuiInputTextCallbackData) GetEventFlag() ImGuiInputTextFlags {
	return ImGuiInputTextFlags(C.ImGuiInputTextCallbackData_GetEventFlag(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetFlags(v ImGuiInputTextFlags) {
	C.ImGuiInputTextCallbackData_SetFlags(self.handle(), C.ImGuiInputTextFlags(v))
}

func (self ImGuiInputTextCallbackData) GetFlags() ImGuiInputTextFlags {
	return ImGuiInputTextFlags(C.ImGuiInputTextCallbackData_GetFlags(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetUserData(v unsafe.Pointer) {
	C.ImGuiInputTextCallbackData_SetUserData(self.handle(), v)
}

func (self ImGuiInputTextCallbackData) GetUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiInputTextCallbackData_GetUserData(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetEventChar(v ImWchar) {
	C.ImGuiInputTextCallbackData_SetEventChar(self.handle(), C.ImWchar(v))
}

func (self ImGuiInputTextCallbackData) SetEventKey(v ImGuiKey) {
	C.ImGuiInputTextCallbackData_SetEventKey(self.handle(), C.ImGuiKey(v))
}

func (self ImGuiInputTextCallbackData) GetEventKey() ImGuiKey {
	return ImGuiKey(C.ImGuiInputTextCallbackData_GetEventKey(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetBuf(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiInputTextCallbackData_SetBuf(self.handle(), vArg)
}

func (self ImGuiInputTextCallbackData) SetBufTextLen(v int32) {
	C.ImGuiInputTextCallbackData_SetBufTextLen(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) GetBufTextLen() int {
	return int(C.ImGuiInputTextCallbackData_GetBufTextLen(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetBufSize(v int32) {
	C.ImGuiInputTextCallbackData_SetBufSize(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) GetBufSize() int {
	return int(C.ImGuiInputTextCallbackData_GetBufSize(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetBufDirty(v bool) {
	C.ImGuiInputTextCallbackData_SetBufDirty(self.handle(), C.bool(v))
}

func (self ImGuiInputTextCallbackData) GetBufDirty() bool {
	return C.ImGuiInputTextCallbackData_GetBufDirty(self.handle()) == C.bool(true)
}

func (self ImGuiInputTextCallbackData) SetCursorPos(v int32) {
	C.ImGuiInputTextCallbackData_SetCursorPos(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) GetCursorPos() int {
	return int(C.ImGuiInputTextCallbackData_GetCursorPos(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetSelectionStart(v int32) {
	C.ImGuiInputTextCallbackData_SetSelectionStart(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) GetSelectionStart() int {
	return int(C.ImGuiInputTextCallbackData_GetSelectionStart(self.handle()))
}

func (self ImGuiInputTextCallbackData) SetSelectionEnd(v int32) {
	C.ImGuiInputTextCallbackData_SetSelectionEnd(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) GetSelectionEnd() int {
	return int(C.ImGuiInputTextCallbackData_GetSelectionEnd(self.handle()))
}

func (self ImGuiTabBar) SetFlags(v ImGuiTabBarFlags) {
	C.ImGuiTabBar_SetFlags(self.handle(), C.ImGuiTabBarFlags(v))
}

func (self ImGuiTabBar) GetFlags() ImGuiTabBarFlags {
	return ImGuiTabBarFlags(C.ImGuiTabBar_GetFlags(self.handle()))
}

func (self ImGuiTabBar) SetID(v ImGuiID) {
	C.ImGuiTabBar_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) GetID() ImGuiID {
	return ImGuiID(C.ImGuiTabBar_GetID(self.handle()))
}

func (self ImGuiTabBar) SetSelectedTabId(v ImGuiID) {
	C.ImGuiTabBar_SetSelectedTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) GetSelectedTabId() ImGuiID {
	return ImGuiID(C.ImGuiTabBar_GetSelectedTabId(self.handle()))
}

func (self ImGuiTabBar) SetNextSelectedTabId(v ImGuiID) {
	C.ImGuiTabBar_SetNextSelectedTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) GetNextSelectedTabId() ImGuiID {
	return ImGuiID(C.ImGuiTabBar_GetNextSelectedTabId(self.handle()))
}

func (self ImGuiTabBar) SetVisibleTabId(v ImGuiID) {
	C.ImGuiTabBar_SetVisibleTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) GetVisibleTabId() ImGuiID {
	return ImGuiID(C.ImGuiTabBar_GetVisibleTabId(self.handle()))
}

func (self ImGuiTabBar) SetCurrFrameVisible(v int32) {
	C.ImGuiTabBar_SetCurrFrameVisible(self.handle(), C.int(v))
}

func (self ImGuiTabBar) GetCurrFrameVisible() int {
	return int(C.ImGuiTabBar_GetCurrFrameVisible(self.handle()))
}

func (self ImGuiTabBar) SetPrevFrameVisible(v int32) {
	C.ImGuiTabBar_SetPrevFrameVisible(self.handle(), C.int(v))
}

func (self ImGuiTabBar) GetPrevFrameVisible() int {
	return int(C.ImGuiTabBar_GetPrevFrameVisible(self.handle()))
}

func (self ImGuiTabBar) SetBarRect(v ImRect) {
	C.ImGuiTabBar_SetBarRect(self.handle(), v.toC())
}

func (self ImGuiTabBar) GetBarRect() ImRect {
	return newImRectFromC(C.ImGuiTabBar_GetBarRect(self.handle()))
}

func (self ImGuiTabBar) SetCurrTabsContentsHeight(v float32) {
	C.ImGuiTabBar_SetCurrTabsContentsHeight(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetCurrTabsContentsHeight() float32 {
	return float32(C.ImGuiTabBar_GetCurrTabsContentsHeight(self.handle()))
}

func (self ImGuiTabBar) SetPrevTabsContentsHeight(v float32) {
	C.ImGuiTabBar_SetPrevTabsContentsHeight(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetPrevTabsContentsHeight() float32 {
	return float32(C.ImGuiTabBar_GetPrevTabsContentsHeight(self.handle()))
}

func (self ImGuiTabBar) SetWidthAllTabs(v float32) {
	C.ImGuiTabBar_SetWidthAllTabs(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetWidthAllTabs() float32 {
	return float32(C.ImGuiTabBar_GetWidthAllTabs(self.handle()))
}

func (self ImGuiTabBar) SetWidthAllTabsIdeal(v float32) {
	C.ImGuiTabBar_SetWidthAllTabsIdeal(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetWidthAllTabsIdeal() float32 {
	return float32(C.ImGuiTabBar_GetWidthAllTabsIdeal(self.handle()))
}

func (self ImGuiTabBar) SetScrollingAnim(v float32) {
	C.ImGuiTabBar_SetScrollingAnim(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetScrollingAnim() float32 {
	return float32(C.ImGuiTabBar_GetScrollingAnim(self.handle()))
}

func (self ImGuiTabBar) SetScrollingTarget(v float32) {
	C.ImGuiTabBar_SetScrollingTarget(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetScrollingTarget() float32 {
	return float32(C.ImGuiTabBar_GetScrollingTarget(self.handle()))
}

func (self ImGuiTabBar) SetScrollingTargetDistToVisibility(v float32) {
	C.ImGuiTabBar_SetScrollingTargetDistToVisibility(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetScrollingTargetDistToVisibility() float32 {
	return float32(C.ImGuiTabBar_GetScrollingTargetDistToVisibility(self.handle()))
}

func (self ImGuiTabBar) SetScrollingSpeed(v float32) {
	C.ImGuiTabBar_SetScrollingSpeed(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetScrollingSpeed() float32 {
	return float32(C.ImGuiTabBar_GetScrollingSpeed(self.handle()))
}

func (self ImGuiTabBar) SetScrollingRectMinX(v float32) {
	C.ImGuiTabBar_SetScrollingRectMinX(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetScrollingRectMinX() float32 {
	return float32(C.ImGuiTabBar_GetScrollingRectMinX(self.handle()))
}

func (self ImGuiTabBar) SetScrollingRectMaxX(v float32) {
	C.ImGuiTabBar_SetScrollingRectMaxX(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetScrollingRectMaxX() float32 {
	return float32(C.ImGuiTabBar_GetScrollingRectMaxX(self.handle()))
}

func (self ImGuiTabBar) SetReorderRequestTabId(v ImGuiID) {
	C.ImGuiTabBar_SetReorderRequestTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) GetReorderRequestTabId() ImGuiID {
	return ImGuiID(C.ImGuiTabBar_GetReorderRequestTabId(self.handle()))
}

func (self ImGuiTabBar) SetReorderRequestOffset(v int) {
	C.ImGuiTabBar_SetReorderRequestOffset(self.handle(), C.ImS16(v))
}

func (self ImGuiTabBar) GetReorderRequestOffset() int {
	return int(C.ImGuiTabBar_GetReorderRequestOffset(self.handle()))
}

func (self ImGuiTabBar) SetBeginCount(v int) {
	C.ImGuiTabBar_SetBeginCount(self.handle(), C.ImS8(v))
}

func (self ImGuiTabBar) GetBeginCount() int {
	return int(C.ImGuiTabBar_GetBeginCount(self.handle()))
}

func (self ImGuiTabBar) SetWantLayout(v bool) {
	C.ImGuiTabBar_SetWantLayout(self.handle(), C.bool(v))
}

func (self ImGuiTabBar) GetWantLayout() bool {
	return C.ImGuiTabBar_GetWantLayout(self.handle()) == C.bool(true)
}

func (self ImGuiTabBar) SetVisibleTabWasSubmitted(v bool) {
	C.ImGuiTabBar_SetVisibleTabWasSubmitted(self.handle(), C.bool(v))
}

func (self ImGuiTabBar) GetVisibleTabWasSubmitted() bool {
	return C.ImGuiTabBar_GetVisibleTabWasSubmitted(self.handle()) == C.bool(true)
}

func (self ImGuiTabBar) SetTabsAddedNew(v bool) {
	C.ImGuiTabBar_SetTabsAddedNew(self.handle(), C.bool(v))
}

func (self ImGuiTabBar) GetTabsAddedNew() bool {
	return C.ImGuiTabBar_GetTabsAddedNew(self.handle()) == C.bool(true)
}

func (self ImGuiTabBar) SetTabsActiveCount(v int) {
	C.ImGuiTabBar_SetTabsActiveCount(self.handle(), C.ImS16(v))
}

func (self ImGuiTabBar) GetTabsActiveCount() int {
	return int(C.ImGuiTabBar_GetTabsActiveCount(self.handle()))
}

func (self ImGuiTabBar) SetLastTabItemIdx(v int) {
	C.ImGuiTabBar_SetLastTabItemIdx(self.handle(), C.ImS16(v))
}

func (self ImGuiTabBar) GetLastTabItemIdx() int {
	return int(C.ImGuiTabBar_GetLastTabItemIdx(self.handle()))
}

func (self ImGuiTabBar) SetItemSpacingY(v float32) {
	C.ImGuiTabBar_SetItemSpacingY(self.handle(), C.float(v))
}

func (self ImGuiTabBar) GetItemSpacingY() float32 {
	return float32(C.ImGuiTabBar_GetItemSpacingY(self.handle()))
}

func (self ImGuiTabBar) SetFramePadding(v ImVec2) {
	C.ImGuiTabBar_SetFramePadding(self.handle(), v.toC())
}

func (self ImGuiTabBar) GetFramePadding() ImVec2 {
	return newImVec2FromC(C.ImGuiTabBar_GetFramePadding(self.handle()))
}

func (self ImGuiTabBar) SetBackupCursorPos(v ImVec2) {
	C.ImGuiTabBar_SetBackupCursorPos(self.handle(), v.toC())
}

func (self ImGuiTabBar) GetBackupCursorPos() ImVec2 {
	return newImVec2FromC(C.ImGuiTabBar_GetBackupCursorPos(self.handle()))
}

func (self ImGuiTabBar) GetTabsNames() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImGuiTabBar_GetTabsNames(self.handle()))
}

func (self ImGuiTable) SetID(v ImGuiID) {
	C.ImGuiTable_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTable) GetID() ImGuiID {
	return ImGuiID(C.ImGuiTable_GetID(self.handle()))
}

func (self ImGuiTable) SetFlags(v ImGuiTableFlags) {
	C.ImGuiTable_SetFlags(self.handle(), C.ImGuiTableFlags(v))
}

func (self ImGuiTable) GetFlags() ImGuiTableFlags {
	return ImGuiTableFlags(C.ImGuiTable_GetFlags(self.handle()))
}

func (self ImGuiTable) SetRawData(v unsafe.Pointer) {
	C.ImGuiTable_SetRawData(self.handle(), v)
}

func (self ImGuiTable) GetRawData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiTable_GetRawData(self.handle()))
}

func (self ImGuiTable) SetTempData(v ImGuiTableTempData) {
	C.ImGuiTable_SetTempData(self.handle(), v.handle())
}

func (self ImGuiTable) GetTempData() ImGuiTableTempData {
	return (ImGuiTableTempData)(unsafe.Pointer(C.ImGuiTable_GetTempData(self.handle())))
}

func (self ImGuiTable) SetEnabledMaskByDisplayOrder(v uint64) {
	C.ImGuiTable_SetEnabledMaskByDisplayOrder(self.handle(), C.ImU64(v))
}

func (self ImGuiTable) GetEnabledMaskByDisplayOrder() uint64 {
	return uint64(C.ImGuiTable_GetEnabledMaskByDisplayOrder(self.handle()))
}

func (self ImGuiTable) SetEnabledMaskByIndex(v uint64) {
	C.ImGuiTable_SetEnabledMaskByIndex(self.handle(), C.ImU64(v))
}

func (self ImGuiTable) GetEnabledMaskByIndex() uint64 {
	return uint64(C.ImGuiTable_GetEnabledMaskByIndex(self.handle()))
}

func (self ImGuiTable) SetVisibleMaskByIndex(v uint64) {
	C.ImGuiTable_SetVisibleMaskByIndex(self.handle(), C.ImU64(v))
}

func (self ImGuiTable) GetVisibleMaskByIndex() uint64 {
	return uint64(C.ImGuiTable_GetVisibleMaskByIndex(self.handle()))
}

func (self ImGuiTable) SetRequestOutputMaskByIndex(v uint64) {
	C.ImGuiTable_SetRequestOutputMaskByIndex(self.handle(), C.ImU64(v))
}

func (self ImGuiTable) GetRequestOutputMaskByIndex() uint64 {
	return uint64(C.ImGuiTable_GetRequestOutputMaskByIndex(self.handle()))
}

func (self ImGuiTable) SetSettingsLoadedFlags(v ImGuiTableFlags) {
	C.ImGuiTable_SetSettingsLoadedFlags(self.handle(), C.ImGuiTableFlags(v))
}

func (self ImGuiTable) GetSettingsLoadedFlags() ImGuiTableFlags {
	return ImGuiTableFlags(C.ImGuiTable_GetSettingsLoadedFlags(self.handle()))
}

func (self ImGuiTable) SetSettingsOffset(v int32) {
	C.ImGuiTable_SetSettingsOffset(self.handle(), C.int(v))
}

func (self ImGuiTable) GetSettingsOffset() int {
	return int(C.ImGuiTable_GetSettingsOffset(self.handle()))
}

func (self ImGuiTable) SetLastFrameActive(v int32) {
	C.ImGuiTable_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiTable) GetLastFrameActive() int {
	return int(C.ImGuiTable_GetLastFrameActive(self.handle()))
}

func (self ImGuiTable) SetColumnsCount(v int32) {
	C.ImGuiTable_SetColumnsCount(self.handle(), C.int(v))
}

func (self ImGuiTable) GetColumnsCount() int {
	return int(C.ImGuiTable_GetColumnsCount(self.handle()))
}

func (self ImGuiTable) SetCurrentRow(v int32) {
	C.ImGuiTable_SetCurrentRow(self.handle(), C.int(v))
}

func (self ImGuiTable) GetCurrentRow() int {
	return int(C.ImGuiTable_GetCurrentRow(self.handle()))
}

func (self ImGuiTable) SetCurrentColumn(v int32) {
	C.ImGuiTable_SetCurrentColumn(self.handle(), C.int(v))
}

func (self ImGuiTable) GetCurrentColumn() int {
	return int(C.ImGuiTable_GetCurrentColumn(self.handle()))
}

func (self ImGuiTable) SetInstanceCurrent(v int) {
	C.ImGuiTable_SetInstanceCurrent(self.handle(), C.ImS16(v))
}

func (self ImGuiTable) GetInstanceCurrent() int {
	return int(C.ImGuiTable_GetInstanceCurrent(self.handle()))
}

func (self ImGuiTable) SetInstanceInteracted(v int) {
	C.ImGuiTable_SetInstanceInteracted(self.handle(), C.ImS16(v))
}

func (self ImGuiTable) GetInstanceInteracted() int {
	return int(C.ImGuiTable_GetInstanceInteracted(self.handle()))
}

func (self ImGuiTable) SetRowPosY1(v float32) {
	C.ImGuiTable_SetRowPosY1(self.handle(), C.float(v))
}

func (self ImGuiTable) GetRowPosY1() float32 {
	return float32(C.ImGuiTable_GetRowPosY1(self.handle()))
}

func (self ImGuiTable) SetRowPosY2(v float32) {
	C.ImGuiTable_SetRowPosY2(self.handle(), C.float(v))
}

func (self ImGuiTable) GetRowPosY2() float32 {
	return float32(C.ImGuiTable_GetRowPosY2(self.handle()))
}

func (self ImGuiTable) SetRowMinHeight(v float32) {
	C.ImGuiTable_SetRowMinHeight(self.handle(), C.float(v))
}

func (self ImGuiTable) GetRowMinHeight() float32 {
	return float32(C.ImGuiTable_GetRowMinHeight(self.handle()))
}

func (self ImGuiTable) SetRowTextBaseline(v float32) {
	C.ImGuiTable_SetRowTextBaseline(self.handle(), C.float(v))
}

func (self ImGuiTable) GetRowTextBaseline() float32 {
	return float32(C.ImGuiTable_GetRowTextBaseline(self.handle()))
}

func (self ImGuiTable) SetRowIndentOffsetX(v float32) {
	C.ImGuiTable_SetRowIndentOffsetX(self.handle(), C.float(v))
}

func (self ImGuiTable) GetRowIndentOffsetX() float32 {
	return float32(C.ImGuiTable_GetRowIndentOffsetX(self.handle()))
}

func (self ImGuiTable) SetRowFlags(v ImGuiTableRowFlags) {
	C.ImGuiTable_SetRowFlags(self.handle(), C.ImGuiTableRowFlags(v))
}

func (self ImGuiTable) GetRowFlags() ImGuiTableRowFlags {
	return ImGuiTableRowFlags(C.ImGuiTable_GetRowFlags(self.handle()))
}

func (self ImGuiTable) SetLastRowFlags(v ImGuiTableRowFlags) {
	C.ImGuiTable_SetLastRowFlags(self.handle(), C.ImGuiTableRowFlags(v))
}

func (self ImGuiTable) GetLastRowFlags() ImGuiTableRowFlags {
	return ImGuiTableRowFlags(C.ImGuiTable_GetLastRowFlags(self.handle()))
}

func (self ImGuiTable) SetRowBgColorCounter(v int32) {
	C.ImGuiTable_SetRowBgColorCounter(self.handle(), C.int(v))
}

func (self ImGuiTable) GetRowBgColorCounter() int {
	return int(C.ImGuiTable_GetRowBgColorCounter(self.handle()))
}

func (self ImGuiTable) SetBorderColorStrong(v uint32) {
	C.ImGuiTable_SetBorderColorStrong(self.handle(), C.ImU32(v))
}

func (self ImGuiTable) GetBorderColorStrong() uint32 {
	return uint32(C.ImGuiTable_GetBorderColorStrong(self.handle()))
}

func (self ImGuiTable) SetBorderColorLight(v uint32) {
	C.ImGuiTable_SetBorderColorLight(self.handle(), C.ImU32(v))
}

func (self ImGuiTable) GetBorderColorLight() uint32 {
	return uint32(C.ImGuiTable_GetBorderColorLight(self.handle()))
}

func (self ImGuiTable) SetBorderX1(v float32) {
	C.ImGuiTable_SetBorderX1(self.handle(), C.float(v))
}

func (self ImGuiTable) GetBorderX1() float32 {
	return float32(C.ImGuiTable_GetBorderX1(self.handle()))
}

func (self ImGuiTable) SetBorderX2(v float32) {
	C.ImGuiTable_SetBorderX2(self.handle(), C.float(v))
}

func (self ImGuiTable) GetBorderX2() float32 {
	return float32(C.ImGuiTable_GetBorderX2(self.handle()))
}

func (self ImGuiTable) SetHostIndentX(v float32) {
	C.ImGuiTable_SetHostIndentX(self.handle(), C.float(v))
}

func (self ImGuiTable) GetHostIndentX() float32 {
	return float32(C.ImGuiTable_GetHostIndentX(self.handle()))
}

func (self ImGuiTable) SetMinColumnWidth(v float32) {
	C.ImGuiTable_SetMinColumnWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) GetMinColumnWidth() float32 {
	return float32(C.ImGuiTable_GetMinColumnWidth(self.handle()))
}

func (self ImGuiTable) SetOuterPaddingX(v float32) {
	C.ImGuiTable_SetOuterPaddingX(self.handle(), C.float(v))
}

func (self ImGuiTable) GetOuterPaddingX() float32 {
	return float32(C.ImGuiTable_GetOuterPaddingX(self.handle()))
}

func (self ImGuiTable) SetCellPaddingX(v float32) {
	C.ImGuiTable_SetCellPaddingX(self.handle(), C.float(v))
}

func (self ImGuiTable) GetCellPaddingX() float32 {
	return float32(C.ImGuiTable_GetCellPaddingX(self.handle()))
}

func (self ImGuiTable) SetCellPaddingY(v float32) {
	C.ImGuiTable_SetCellPaddingY(self.handle(), C.float(v))
}

func (self ImGuiTable) GetCellPaddingY() float32 {
	return float32(C.ImGuiTable_GetCellPaddingY(self.handle()))
}

func (self ImGuiTable) SetCellSpacingX1(v float32) {
	C.ImGuiTable_SetCellSpacingX1(self.handle(), C.float(v))
}

func (self ImGuiTable) GetCellSpacingX1() float32 {
	return float32(C.ImGuiTable_GetCellSpacingX1(self.handle()))
}

func (self ImGuiTable) SetCellSpacingX2(v float32) {
	C.ImGuiTable_SetCellSpacingX2(self.handle(), C.float(v))
}

func (self ImGuiTable) GetCellSpacingX2() float32 {
	return float32(C.ImGuiTable_GetCellSpacingX2(self.handle()))
}

func (self ImGuiTable) SetInnerWidth(v float32) {
	C.ImGuiTable_SetInnerWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) GetInnerWidth() float32 {
	return float32(C.ImGuiTable_GetInnerWidth(self.handle()))
}

func (self ImGuiTable) SetColumnsGivenWidth(v float32) {
	C.ImGuiTable_SetColumnsGivenWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) GetColumnsGivenWidth() float32 {
	return float32(C.ImGuiTable_GetColumnsGivenWidth(self.handle()))
}

func (self ImGuiTable) SetColumnsAutoFitWidth(v float32) {
	C.ImGuiTable_SetColumnsAutoFitWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) GetColumnsAutoFitWidth() float32 {
	return float32(C.ImGuiTable_GetColumnsAutoFitWidth(self.handle()))
}

func (self ImGuiTable) SetColumnsStretchSumWeights(v float32) {
	C.ImGuiTable_SetColumnsStretchSumWeights(self.handle(), C.float(v))
}

func (self ImGuiTable) GetColumnsStretchSumWeights() float32 {
	return float32(C.ImGuiTable_GetColumnsStretchSumWeights(self.handle()))
}

func (self ImGuiTable) SetResizedColumnNextWidth(v float32) {
	C.ImGuiTable_SetResizedColumnNextWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) GetResizedColumnNextWidth() float32 {
	return float32(C.ImGuiTable_GetResizedColumnNextWidth(self.handle()))
}

func (self ImGuiTable) SetResizeLockMinContentsX2(v float32) {
	C.ImGuiTable_SetResizeLockMinContentsX2(self.handle(), C.float(v))
}

func (self ImGuiTable) GetResizeLockMinContentsX2() float32 {
	return float32(C.ImGuiTable_GetResizeLockMinContentsX2(self.handle()))
}

func (self ImGuiTable) SetRefScale(v float32) {
	C.ImGuiTable_SetRefScale(self.handle(), C.float(v))
}

func (self ImGuiTable) GetRefScale() float32 {
	return float32(C.ImGuiTable_GetRefScale(self.handle()))
}

func (self ImGuiTable) SetOuterRect(v ImRect) {
	C.ImGuiTable_SetOuterRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetOuterRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetOuterRect(self.handle()))
}

func (self ImGuiTable) SetInnerRect(v ImRect) {
	C.ImGuiTable_SetInnerRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetInnerRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetInnerRect(self.handle()))
}

func (self ImGuiTable) SetWorkRect(v ImRect) {
	C.ImGuiTable_SetWorkRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetWorkRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetWorkRect(self.handle()))
}

func (self ImGuiTable) SetInnerClipRect(v ImRect) {
	C.ImGuiTable_SetInnerClipRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetInnerClipRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetInnerClipRect(self.handle()))
}

func (self ImGuiTable) SetBgClipRect(v ImRect) {
	C.ImGuiTable_SetBgClipRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetBgClipRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetBgClipRect(self.handle()))
}

func (self ImGuiTable) SetBg0ClipRectForDrawCmd(v ImRect) {
	C.ImGuiTable_SetBg0ClipRectForDrawCmd(self.handle(), v.toC())
}

func (self ImGuiTable) GetBg0ClipRectForDrawCmd() ImRect {
	return newImRectFromC(C.ImGuiTable_GetBg0ClipRectForDrawCmd(self.handle()))
}

func (self ImGuiTable) SetBg2ClipRectForDrawCmd(v ImRect) {
	C.ImGuiTable_SetBg2ClipRectForDrawCmd(self.handle(), v.toC())
}

func (self ImGuiTable) GetBg2ClipRectForDrawCmd() ImRect {
	return newImRectFromC(C.ImGuiTable_GetBg2ClipRectForDrawCmd(self.handle()))
}

func (self ImGuiTable) SetHostClipRect(v ImRect) {
	C.ImGuiTable_SetHostClipRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetHostClipRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetHostClipRect(self.handle()))
}

func (self ImGuiTable) SetHostBackupInnerClipRect(v ImRect) {
	C.ImGuiTable_SetHostBackupInnerClipRect(self.handle(), v.toC())
}

func (self ImGuiTable) GetHostBackupInnerClipRect() ImRect {
	return newImRectFromC(C.ImGuiTable_GetHostBackupInnerClipRect(self.handle()))
}

func (self ImGuiTable) SetOuterWindow(v ImGuiWindow) {
	C.ImGuiTable_SetOuterWindow(self.handle(), v.handle())
}

func (self ImGuiTable) GetOuterWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiTable_GetOuterWindow(self.handle())))
}

func (self ImGuiTable) SetInnerWindow(v ImGuiWindow) {
	C.ImGuiTable_SetInnerWindow(self.handle(), v.handle())
}

func (self ImGuiTable) GetInnerWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiTable_GetInnerWindow(self.handle())))
}

func (self ImGuiTable) GetColumnsNames() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImGuiTable_GetColumnsNames(self.handle()))
}

func (self ImGuiTable) SetDrawSplitter(v ImDrawListSplitter) {
	C.ImGuiTable_SetDrawSplitter(self.handle(), v.handle())
}

func (self ImGuiTable) GetDrawSplitter() ImDrawListSplitter {
	return (ImDrawListSplitter)(unsafe.Pointer(C.ImGuiTable_GetDrawSplitter(self.handle())))
}

func (self ImGuiTable) GetInstanceDataFirst() ImGuiTableInstanceData {
	return newImGuiTableInstanceDataFromC(C.ImGuiTable_GetInstanceDataFirst(self.handle()))
}

func (self ImGuiTable) GetSortSpecsSingle() ImGuiTableColumnSortSpecs {
	return newImGuiTableColumnSortSpecsFromC(C.ImGuiTable_GetSortSpecsSingle(self.handle()))
}

func (self ImGuiTable) GetSortSpecs() ImGuiTableSortSpecs {
	return newImGuiTableSortSpecsFromC(C.ImGuiTable_GetSortSpecs(self.handle()))
}

func (self ImGuiTable) SetSortSpecsCount(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetSortSpecsCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetSortSpecsCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetSortSpecsCount(self.handle()))
}

func (self ImGuiTable) SetColumnsEnabledCount(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetColumnsEnabledCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetColumnsEnabledCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetColumnsEnabledCount(self.handle()))
}

func (self ImGuiTable) SetColumnsEnabledFixedCount(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetColumnsEnabledFixedCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetColumnsEnabledFixedCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetColumnsEnabledFixedCount(self.handle()))
}

func (self ImGuiTable) SetDeclColumnsCount(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetDeclColumnsCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetDeclColumnsCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetDeclColumnsCount(self.handle()))
}

func (self ImGuiTable) SetHoveredColumnBody(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetHoveredColumnBody(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetHoveredColumnBody() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetHoveredColumnBody(self.handle()))
}

func (self ImGuiTable) SetHoveredColumnBorder(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetHoveredColumnBorder(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetHoveredColumnBorder() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetHoveredColumnBorder(self.handle()))
}

func (self ImGuiTable) SetAutoFitSingleColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetAutoFitSingleColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetAutoFitSingleColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetAutoFitSingleColumn(self.handle()))
}

func (self ImGuiTable) SetResizedColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetResizedColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetResizedColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetResizedColumn(self.handle()))
}

func (self ImGuiTable) SetLastResizedColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetLastResizedColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetLastResizedColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetLastResizedColumn(self.handle()))
}

func (self ImGuiTable) SetHeldHeaderColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetHeldHeaderColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetHeldHeaderColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetHeldHeaderColumn(self.handle()))
}

func (self ImGuiTable) SetReorderColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetReorderColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetReorderColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetReorderColumn(self.handle()))
}

func (self ImGuiTable) SetReorderColumnDir(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetReorderColumnDir(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetReorderColumnDir() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetReorderColumnDir(self.handle()))
}

func (self ImGuiTable) SetLeftMostEnabledColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetLeftMostEnabledColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetLeftMostEnabledColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetLeftMostEnabledColumn(self.handle()))
}

func (self ImGuiTable) SetRightMostEnabledColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetRightMostEnabledColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetRightMostEnabledColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetRightMostEnabledColumn(self.handle()))
}

func (self ImGuiTable) SetLeftMostStretchedColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetLeftMostStretchedColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetLeftMostStretchedColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetLeftMostStretchedColumn(self.handle()))
}

func (self ImGuiTable) SetRightMostStretchedColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetRightMostStretchedColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetRightMostStretchedColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetRightMostStretchedColumn(self.handle()))
}

func (self ImGuiTable) SetContextPopupColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetContextPopupColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetContextPopupColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetContextPopupColumn(self.handle()))
}

func (self ImGuiTable) SetFreezeRowsRequest(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetFreezeRowsRequest(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetFreezeRowsRequest() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetFreezeRowsRequest(self.handle()))
}

func (self ImGuiTable) SetFreezeRowsCount(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetFreezeRowsCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetFreezeRowsCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetFreezeRowsCount(self.handle()))
}

func (self ImGuiTable) SetFreezeColumnsRequest(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetFreezeColumnsRequest(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetFreezeColumnsRequest() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetFreezeColumnsRequest(self.handle()))
}

func (self ImGuiTable) SetFreezeColumnsCount(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetFreezeColumnsCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetFreezeColumnsCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetFreezeColumnsCount(self.handle()))
}

func (self ImGuiTable) SetRowCellDataCurrent(v ImGuiTableColumnIdx) {
	C.ImGuiTable_SetRowCellDataCurrent(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTable) GetRowCellDataCurrent() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTable_GetRowCellDataCurrent(self.handle()))
}

func (self ImGuiTable) SetDummyDrawChannel(v ImGuiTableDrawChannelIdx) {
	C.ImGuiTable_SetDummyDrawChannel(self.handle(), C.ImGuiTableDrawChannelIdx(v))
}

func (self ImGuiTable) GetDummyDrawChannel() ImGuiTableDrawChannelIdx {
	return ImGuiTableDrawChannelIdx(C.ImGuiTable_GetDummyDrawChannel(self.handle()))
}

func (self ImGuiTable) SetBg2DrawChannelCurrent(v ImGuiTableDrawChannelIdx) {
	C.ImGuiTable_SetBg2DrawChannelCurrent(self.handle(), C.ImGuiTableDrawChannelIdx(v))
}

func (self ImGuiTable) GetBg2DrawChannelCurrent() ImGuiTableDrawChannelIdx {
	return ImGuiTableDrawChannelIdx(C.ImGuiTable_GetBg2DrawChannelCurrent(self.handle()))
}

func (self ImGuiTable) SetBg2DrawChannelUnfrozen(v ImGuiTableDrawChannelIdx) {
	C.ImGuiTable_SetBg2DrawChannelUnfrozen(self.handle(), C.ImGuiTableDrawChannelIdx(v))
}

func (self ImGuiTable) GetBg2DrawChannelUnfrozen() ImGuiTableDrawChannelIdx {
	return ImGuiTableDrawChannelIdx(C.ImGuiTable_GetBg2DrawChannelUnfrozen(self.handle()))
}

func (self ImGuiTable) SetIsLayoutLocked(v bool) {
	C.ImGuiTable_SetIsLayoutLocked(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsLayoutLocked() bool {
	return C.ImGuiTable_GetIsLayoutLocked(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsInsideRow(v bool) {
	C.ImGuiTable_SetIsInsideRow(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsInsideRow() bool {
	return C.ImGuiTable_GetIsInsideRow(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsInitializing(v bool) {
	C.ImGuiTable_SetIsInitializing(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsInitializing() bool {
	return C.ImGuiTable_GetIsInitializing(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsSortSpecsDirty(v bool) {
	C.ImGuiTable_SetIsSortSpecsDirty(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsSortSpecsDirty() bool {
	return C.ImGuiTable_GetIsSortSpecsDirty(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsUsingHeaders(v bool) {
	C.ImGuiTable_SetIsUsingHeaders(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsUsingHeaders() bool {
	return C.ImGuiTable_GetIsUsingHeaders(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsContextPopupOpen(v bool) {
	C.ImGuiTable_SetIsContextPopupOpen(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsContextPopupOpen() bool {
	return C.ImGuiTable_GetIsContextPopupOpen(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsSettingsRequestLoad(v bool) {
	C.ImGuiTable_SetIsSettingsRequestLoad(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsSettingsRequestLoad() bool {
	return C.ImGuiTable_GetIsSettingsRequestLoad(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsSettingsDirty(v bool) {
	C.ImGuiTable_SetIsSettingsDirty(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsSettingsDirty() bool {
	return C.ImGuiTable_GetIsSettingsDirty(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsDefaultDisplayOrder(v bool) {
	C.ImGuiTable_SetIsDefaultDisplayOrder(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsDefaultDisplayOrder() bool {
	return C.ImGuiTable_GetIsDefaultDisplayOrder(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsResetAllRequest(v bool) {
	C.ImGuiTable_SetIsResetAllRequest(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsResetAllRequest() bool {
	return C.ImGuiTable_GetIsResetAllRequest(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsResetDisplayOrderRequest(v bool) {
	C.ImGuiTable_SetIsResetDisplayOrderRequest(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsResetDisplayOrderRequest() bool {
	return C.ImGuiTable_GetIsResetDisplayOrderRequest(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsUnfrozenRows(v bool) {
	C.ImGuiTable_SetIsUnfrozenRows(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsUnfrozenRows() bool {
	return C.ImGuiTable_GetIsUnfrozenRows(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetIsDefaultSizingPolicy(v bool) {
	C.ImGuiTable_SetIsDefaultSizingPolicy(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetIsDefaultSizingPolicy() bool {
	return C.ImGuiTable_GetIsDefaultSizingPolicy(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetMemoryCompacted(v bool) {
	C.ImGuiTable_SetMemoryCompacted(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetMemoryCompacted() bool {
	return C.ImGuiTable_GetMemoryCompacted(self.handle()) == C.bool(true)
}

func (self ImGuiTable) SetHostSkipItems(v bool) {
	C.ImGuiTable_SetHostSkipItems(self.handle(), C.bool(v))
}

func (self ImGuiTable) GetHostSkipItems() bool {
	return C.ImGuiTable_GetHostSkipItems(self.handle()) == C.bool(true)
}

func (self ImGuiTableTempData) SetTableIndex(v int32) {
	C.ImGuiTableTempData_SetTableIndex(self.handle(), C.int(v))
}

func (self ImGuiTableTempData) GetTableIndex() int {
	return int(C.ImGuiTableTempData_GetTableIndex(self.handle()))
}

func (self ImGuiTableTempData) SetLastTimeActive(v float32) {
	C.ImGuiTableTempData_SetLastTimeActive(self.handle(), C.float(v))
}

func (self ImGuiTableTempData) GetLastTimeActive() float32 {
	return float32(C.ImGuiTableTempData_GetLastTimeActive(self.handle()))
}

func (self ImGuiTableTempData) SetUserOuterSize(v ImVec2) {
	C.ImGuiTableTempData_SetUserOuterSize(self.handle(), v.toC())
}

func (self ImGuiTableTempData) GetUserOuterSize() ImVec2 {
	return newImVec2FromC(C.ImGuiTableTempData_GetUserOuterSize(self.handle()))
}

func (self ImGuiTableTempData) GetDrawSplitter() ImDrawListSplitter {
	return newImDrawListSplitterFromC(C.ImGuiTableTempData_GetDrawSplitter(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupWorkRect(v ImRect) {
	C.ImGuiTableTempData_SetHostBackupWorkRect(self.handle(), v.toC())
}

func (self ImGuiTableTempData) GetHostBackupWorkRect() ImRect {
	return newImRectFromC(C.ImGuiTableTempData_GetHostBackupWorkRect(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupParentWorkRect(v ImRect) {
	C.ImGuiTableTempData_SetHostBackupParentWorkRect(self.handle(), v.toC())
}

func (self ImGuiTableTempData) GetHostBackupParentWorkRect() ImRect {
	return newImRectFromC(C.ImGuiTableTempData_GetHostBackupParentWorkRect(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupPrevLineSize(v ImVec2) {
	C.ImGuiTableTempData_SetHostBackupPrevLineSize(self.handle(), v.toC())
}

func (self ImGuiTableTempData) GetHostBackupPrevLineSize() ImVec2 {
	return newImVec2FromC(C.ImGuiTableTempData_GetHostBackupPrevLineSize(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupCurrLineSize(v ImVec2) {
	C.ImGuiTableTempData_SetHostBackupCurrLineSize(self.handle(), v.toC())
}

func (self ImGuiTableTempData) GetHostBackupCurrLineSize() ImVec2 {
	return newImVec2FromC(C.ImGuiTableTempData_GetHostBackupCurrLineSize(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupCursorMaxPos(v ImVec2) {
	C.ImGuiTableTempData_SetHostBackupCursorMaxPos(self.handle(), v.toC())
}

func (self ImGuiTableTempData) GetHostBackupCursorMaxPos() ImVec2 {
	return newImVec2FromC(C.ImGuiTableTempData_GetHostBackupCursorMaxPos(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupItemWidth(v float32) {
	C.ImGuiTableTempData_SetHostBackupItemWidth(self.handle(), C.float(v))
}

func (self ImGuiTableTempData) GetHostBackupItemWidth() float32 {
	return float32(C.ImGuiTableTempData_GetHostBackupItemWidth(self.handle()))
}

func (self ImGuiTableTempData) SetHostBackupItemWidthStackSize(v int32) {
	C.ImGuiTableTempData_SetHostBackupItemWidthStackSize(self.handle(), C.int(v))
}

func (self ImGuiTableTempData) GetHostBackupItemWidthStackSize() int {
	return int(C.ImGuiTableTempData_GetHostBackupItemWidthStackSize(self.handle()))
}

func (self ImGuiInputEventKey) SetKey(v ImGuiKey) {
	C.ImGuiInputEventKey_SetKey(self.handle(), C.ImGuiKey(v))
}

func (self ImGuiInputEventKey) GetKey() ImGuiKey {
	return ImGuiKey(C.ImGuiInputEventKey_GetKey(self.handle()))
}

func (self ImGuiInputEventKey) SetDown(v bool) {
	C.ImGuiInputEventKey_SetDown(self.handle(), C.bool(v))
}

func (self ImGuiInputEventKey) GetDown() bool {
	return C.ImGuiInputEventKey_GetDown(self.handle()) == C.bool(true)
}

func (self ImGuiInputEventKey) SetAnalogValue(v float32) {
	C.ImGuiInputEventKey_SetAnalogValue(self.handle(), C.float(v))
}

func (self ImGuiInputEventKey) GetAnalogValue() float32 {
	return float32(C.ImGuiInputEventKey_GetAnalogValue(self.handle()))
}

func (self ImGuiInputEventMousePos) SetPosX(v float32) {
	C.ImGuiInputEventMousePos_SetPosX(self.handle(), C.float(v))
}

func (self ImGuiInputEventMousePos) GetPosX() float32 {
	return float32(C.ImGuiInputEventMousePos_GetPosX(self.handle()))
}

func (self ImGuiInputEventMousePos) SetPosY(v float32) {
	C.ImGuiInputEventMousePos_SetPosY(self.handle(), C.float(v))
}

func (self ImGuiInputEventMousePos) GetPosY() float32 {
	return float32(C.ImGuiInputEventMousePos_GetPosY(self.handle()))
}

func (self ImGuiInputEventMouseWheel) SetWheelX(v float32) {
	C.ImGuiInputEventMouseWheel_SetWheelX(self.handle(), C.float(v))
}

func (self ImGuiInputEventMouseWheel) GetWheelX() float32 {
	return float32(C.ImGuiInputEventMouseWheel_GetWheelX(self.handle()))
}

func (self ImGuiInputEventMouseWheel) SetWheelY(v float32) {
	C.ImGuiInputEventMouseWheel_SetWheelY(self.handle(), C.float(v))
}

func (self ImGuiInputEventMouseWheel) GetWheelY() float32 {
	return float32(C.ImGuiInputEventMouseWheel_GetWheelY(self.handle()))
}

func (self ImGuiNextWindowData) SetFlags(v ImGuiNextWindowDataFlags) {
	C.ImGuiNextWindowData_SetFlags(self.handle(), C.ImGuiNextWindowDataFlags(v))
}

func (self ImGuiNextWindowData) GetFlags() ImGuiNextWindowDataFlags {
	return ImGuiNextWindowDataFlags(C.ImGuiNextWindowData_GetFlags(self.handle()))
}

func (self ImGuiNextWindowData) SetPosCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetPosCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) GetPosCond() ImGuiCond {
	return ImGuiCond(C.ImGuiNextWindowData_GetPosCond(self.handle()))
}

func (self ImGuiNextWindowData) SetSizeCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetSizeCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) GetSizeCond() ImGuiCond {
	return ImGuiCond(C.ImGuiNextWindowData_GetSizeCond(self.handle()))
}

func (self ImGuiNextWindowData) SetCollapsedCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetCollapsedCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) GetCollapsedCond() ImGuiCond {
	return ImGuiCond(C.ImGuiNextWindowData_GetCollapsedCond(self.handle()))
}

func (self ImGuiNextWindowData) SetDockCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetDockCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) GetDockCond() ImGuiCond {
	return ImGuiCond(C.ImGuiNextWindowData_GetDockCond(self.handle()))
}

func (self ImGuiNextWindowData) SetPosVal(v ImVec2) {
	C.ImGuiNextWindowData_SetPosVal(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetPosVal() ImVec2 {
	return newImVec2FromC(C.ImGuiNextWindowData_GetPosVal(self.handle()))
}

func (self ImGuiNextWindowData) SetPosPivotVal(v ImVec2) {
	C.ImGuiNextWindowData_SetPosPivotVal(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetPosPivotVal() ImVec2 {
	return newImVec2FromC(C.ImGuiNextWindowData_GetPosPivotVal(self.handle()))
}

func (self ImGuiNextWindowData) SetSizeVal(v ImVec2) {
	C.ImGuiNextWindowData_SetSizeVal(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetSizeVal() ImVec2 {
	return newImVec2FromC(C.ImGuiNextWindowData_GetSizeVal(self.handle()))
}

func (self ImGuiNextWindowData) SetContentSizeVal(v ImVec2) {
	C.ImGuiNextWindowData_SetContentSizeVal(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetContentSizeVal() ImVec2 {
	return newImVec2FromC(C.ImGuiNextWindowData_GetContentSizeVal(self.handle()))
}

func (self ImGuiNextWindowData) SetScrollVal(v ImVec2) {
	C.ImGuiNextWindowData_SetScrollVal(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetScrollVal() ImVec2 {
	return newImVec2FromC(C.ImGuiNextWindowData_GetScrollVal(self.handle()))
}

func (self ImGuiNextWindowData) SetPosUndock(v bool) {
	C.ImGuiNextWindowData_SetPosUndock(self.handle(), C.bool(v))
}

func (self ImGuiNextWindowData) GetPosUndock() bool {
	return C.ImGuiNextWindowData_GetPosUndock(self.handle()) == C.bool(true)
}

func (self ImGuiNextWindowData) SetCollapsedVal(v bool) {
	C.ImGuiNextWindowData_SetCollapsedVal(self.handle(), C.bool(v))
}

func (self ImGuiNextWindowData) GetCollapsedVal() bool {
	return C.ImGuiNextWindowData_GetCollapsedVal(self.handle()) == C.bool(true)
}

func (self ImGuiNextWindowData) SetSizeConstraintRect(v ImRect) {
	C.ImGuiNextWindowData_SetSizeConstraintRect(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetSizeConstraintRect() ImRect {
	return newImRectFromC(C.ImGuiNextWindowData_GetSizeConstraintRect(self.handle()))
}

func (self ImGuiNextWindowData) SetSizeCallbackUserData(v unsafe.Pointer) {
	C.ImGuiNextWindowData_SetSizeCallbackUserData(self.handle(), v)
}

func (self ImGuiNextWindowData) GetSizeCallbackUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiNextWindowData_GetSizeCallbackUserData(self.handle()))
}

func (self ImGuiNextWindowData) SetBgAlphaVal(v float32) {
	C.ImGuiNextWindowData_SetBgAlphaVal(self.handle(), C.float(v))
}

func (self ImGuiNextWindowData) GetBgAlphaVal() float32 {
	return float32(C.ImGuiNextWindowData_GetBgAlphaVal(self.handle()))
}

func (self ImGuiNextWindowData) SetViewportId(v ImGuiID) {
	C.ImGuiNextWindowData_SetViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNextWindowData) GetViewportId() ImGuiID {
	return ImGuiID(C.ImGuiNextWindowData_GetViewportId(self.handle()))
}

func (self ImGuiNextWindowData) SetDockId(v ImGuiID) {
	C.ImGuiNextWindowData_SetDockId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNextWindowData) GetDockId() ImGuiID {
	return ImGuiID(C.ImGuiNextWindowData_GetDockId(self.handle()))
}

func (self ImGuiNextWindowData) GetWindowClass() ImGuiWindowClass {
	return newImGuiWindowClassFromC(C.ImGuiNextWindowData_GetWindowClass(self.handle()))
}

func (self ImGuiNextWindowData) SetMenuBarOffsetMinVal(v ImVec2) {
	C.ImGuiNextWindowData_SetMenuBarOffsetMinVal(self.handle(), v.toC())
}

func (self ImGuiNextWindowData) GetMenuBarOffsetMinVal() ImVec2 {
	return newImVec2FromC(C.ImGuiNextWindowData_GetMenuBarOffsetMinVal(self.handle()))
}

func (self ImGuiStackTool) SetLastActiveFrame(v int32) {
	C.ImGuiStackTool_SetLastActiveFrame(self.handle(), C.int(v))
}

func (self ImGuiStackTool) GetLastActiveFrame() int {
	return int(C.ImGuiStackTool_GetLastActiveFrame(self.handle()))
}

func (self ImGuiStackTool) SetStackLevel(v int32) {
	C.ImGuiStackTool_SetStackLevel(self.handle(), C.int(v))
}

func (self ImGuiStackTool) GetStackLevel() int {
	return int(C.ImGuiStackTool_GetStackLevel(self.handle()))
}

func (self ImGuiStackTool) SetQueryId(v ImGuiID) {
	C.ImGuiStackTool_SetQueryId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiStackTool) GetQueryId() ImGuiID {
	return ImGuiID(C.ImGuiStackTool_GetQueryId(self.handle()))
}

func (self ImGuiStackTool) SetCopyToClipboardOnCtrlC(v bool) {
	C.ImGuiStackTool_SetCopyToClipboardOnCtrlC(self.handle(), C.bool(v))
}

func (self ImGuiStackTool) GetCopyToClipboardOnCtrlC() bool {
	return C.ImGuiStackTool_GetCopyToClipboardOnCtrlC(self.handle()) == C.bool(true)
}

func (self ImGuiStackTool) SetCopyToClipboardLastTime(v float32) {
	C.ImGuiStackTool_SetCopyToClipboardLastTime(self.handle(), C.float(v))
}

func (self ImGuiStackTool) GetCopyToClipboardLastTime() float32 {
	return float32(C.ImGuiStackTool_GetCopyToClipboardLastTime(self.handle()))
}

func (self ImGuiStoragePair) Setkey(v ImGuiID) {
	C.ImGuiStoragePair_Setkey(self.handle(), C.ImGuiID(v))
}

func (self ImGuiStoragePair) Getkey() ImGuiID {
	return ImGuiID(C.ImGuiStoragePair_Getkey(self.handle()))
}

func (self ImGuiTextFilter) SetCountGrep(v int32) {
	C.ImGuiTextFilter_SetCountGrep(self.handle(), C.int(v))
}

func (self ImGuiTextFilter) GetCountGrep() int {
	return int(C.ImGuiTextFilter_GetCountGrep(self.handle()))
}

func (self ImGuiDockContext) GetNodes() ImGuiStorage {
	return newImGuiStorageFromC(C.ImGuiDockContext_GetNodes(self.handle()))
}

func (self ImGuiDockContext) SetWantFullRebuild(v bool) {
	C.ImGuiDockContext_SetWantFullRebuild(self.handle(), C.bool(v))
}

func (self ImGuiDockContext) GetWantFullRebuild() bool {
	return C.ImGuiDockContext_GetWantFullRebuild(self.handle()) == C.bool(true)
}

func (self ImGuiSettingsHandler) SetTypeName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiSettingsHandler_SetTypeName(self.handle(), vArg)
}

func (self ImGuiSettingsHandler) GetTypeName() string {
	return C.GoString(C.ImGuiSettingsHandler_GetTypeName(self.handle()))
}

func (self ImGuiSettingsHandler) SetTypeHash(v ImGuiID) {
	C.ImGuiSettingsHandler_SetTypeHash(self.handle(), C.ImGuiID(v))
}

func (self ImGuiSettingsHandler) GetTypeHash() ImGuiID {
	return ImGuiID(C.ImGuiSettingsHandler_GetTypeHash(self.handle()))
}

func (self ImGuiSettingsHandler) SetUserData(v unsafe.Pointer) {
	C.ImGuiSettingsHandler_SetUserData(self.handle(), v)
}

func (self ImGuiSettingsHandler) GetUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiSettingsHandler_GetUserData(self.handle()))
}

func (self ImGuiWindowSettings) SetID(v ImGuiID) {
	C.ImGuiWindowSettings_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) GetID() ImGuiID {
	return ImGuiID(C.ImGuiWindowSettings_GetID(self.handle()))
}

func (self ImGuiWindowSettings) SetViewportId(v ImGuiID) {
	C.ImGuiWindowSettings_SetViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) GetViewportId() ImGuiID {
	return ImGuiID(C.ImGuiWindowSettings_GetViewportId(self.handle()))
}

func (self ImGuiWindowSettings) SetDockId(v ImGuiID) {
	C.ImGuiWindowSettings_SetDockId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) GetDockId() ImGuiID {
	return ImGuiID(C.ImGuiWindowSettings_GetDockId(self.handle()))
}

func (self ImGuiWindowSettings) SetClassId(v ImGuiID) {
	C.ImGuiWindowSettings_SetClassId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) GetClassId() ImGuiID {
	return ImGuiID(C.ImGuiWindowSettings_GetClassId(self.handle()))
}

func (self ImGuiWindowSettings) SetDockOrder(v int) {
	C.ImGuiWindowSettings_SetDockOrder(self.handle(), C.short(v))
}

func (self ImGuiWindowSettings) GetDockOrder() int {
	return int(C.ImGuiWindowSettings_GetDockOrder(self.handle()))
}

func (self ImGuiWindowSettings) SetCollapsed(v bool) {
	C.ImGuiWindowSettings_SetCollapsed(self.handle(), C.bool(v))
}

func (self ImGuiWindowSettings) GetCollapsed() bool {
	return C.ImGuiWindowSettings_GetCollapsed(self.handle()) == C.bool(true)
}

func (self ImGuiWindowSettings) SetWantApply(v bool) {
	C.ImGuiWindowSettings_SetWantApply(self.handle(), C.bool(v))
}

func (self ImGuiWindowSettings) GetWantApply() bool {
	return C.ImGuiWindowSettings_GetWantApply(self.handle()) == C.bool(true)
}

func (self ImGuiInputEventAppFocused) SetFocused(v bool) {
	C.ImGuiInputEventAppFocused_SetFocused(self.handle(), C.bool(v))
}

func (self ImGuiInputEventAppFocused) GetFocused() bool {
	return C.ImGuiInputEventAppFocused_GetFocused(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetID(v ImGuiID) {
	C.ImGuiDockNode_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) GetID() ImGuiID {
	return ImGuiID(C.ImGuiDockNode_GetID(self.handle()))
}

func (self ImGuiDockNode) SetSharedFlags(v ImGuiDockNodeFlags) {
	C.ImGuiDockNode_SetSharedFlags(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiDockNode) GetSharedFlags() ImGuiDockNodeFlags {
	return ImGuiDockNodeFlags(C.ImGuiDockNode_GetSharedFlags(self.handle()))
}

func (self ImGuiDockNode) SetLocalFlagsInWindows(v ImGuiDockNodeFlags) {
	C.ImGuiDockNode_SetLocalFlagsInWindows(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiDockNode) GetLocalFlagsInWindows() ImGuiDockNodeFlags {
	return ImGuiDockNodeFlags(C.ImGuiDockNode_GetLocalFlagsInWindows(self.handle()))
}

func (self ImGuiDockNode) SetMergedFlags(v ImGuiDockNodeFlags) {
	C.ImGuiDockNode_SetMergedFlags(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiDockNode) GetMergedFlags() ImGuiDockNodeFlags {
	return ImGuiDockNodeFlags(C.ImGuiDockNode_GetMergedFlags(self.handle()))
}

func (self ImGuiDockNode) SetState(v ImGuiDockNodeState) {
	C.ImGuiDockNode_SetState(self.handle(), C.ImGuiDockNodeState(v))
}

func (self ImGuiDockNode) GetState() ImGuiDockNodeState {
	return ImGuiDockNodeState(C.ImGuiDockNode_GetState(self.handle()))
}

func (self ImGuiDockNode) SetParentNode(v ImGuiDockNode) {
	C.ImGuiDockNode_SetParentNode(self.handle(), v.handle())
}

func (self ImGuiDockNode) GetParentNode() ImGuiDockNode {
	return (ImGuiDockNode)(unsafe.Pointer(C.ImGuiDockNode_GetParentNode(self.handle())))
}

func (self ImGuiDockNode) SetTabBar(v ImGuiTabBar) {
	C.ImGuiDockNode_SetTabBar(self.handle(), v.handle())
}

func (self ImGuiDockNode) GetTabBar() ImGuiTabBar {
	return (ImGuiTabBar)(unsafe.Pointer(C.ImGuiDockNode_GetTabBar(self.handle())))
}

func (self ImGuiDockNode) SetPos(v ImVec2) {
	C.ImGuiDockNode_SetPos(self.handle(), v.toC())
}

func (self ImGuiDockNode) GetPos() ImVec2 {
	return newImVec2FromC(C.ImGuiDockNode_GetPos(self.handle()))
}

func (self ImGuiDockNode) SetSize(v ImVec2) {
	C.ImGuiDockNode_SetSize(self.handle(), v.toC())
}

func (self ImGuiDockNode) GetSize() ImVec2 {
	return newImVec2FromC(C.ImGuiDockNode_GetSize(self.handle()))
}

func (self ImGuiDockNode) SetSizeRef(v ImVec2) {
	C.ImGuiDockNode_SetSizeRef(self.handle(), v.toC())
}

func (self ImGuiDockNode) GetSizeRef() ImVec2 {
	return newImVec2FromC(C.ImGuiDockNode_GetSizeRef(self.handle()))
}

func (self ImGuiDockNode) SetSplitAxis(v ImGuiAxis) {
	C.ImGuiDockNode_SetSplitAxis(self.handle(), C.ImGuiAxis(v))
}

func (self ImGuiDockNode) GetSplitAxis() ImGuiAxis {
	return ImGuiAxis(C.ImGuiDockNode_GetSplitAxis(self.handle()))
}

func (self ImGuiDockNode) GetWindowClass() ImGuiWindowClass {
	return newImGuiWindowClassFromC(C.ImGuiDockNode_GetWindowClass(self.handle()))
}

func (self ImGuiDockNode) SetLastBgColor(v uint32) {
	C.ImGuiDockNode_SetLastBgColor(self.handle(), C.ImU32(v))
}

func (self ImGuiDockNode) GetLastBgColor() uint32 {
	return uint32(C.ImGuiDockNode_GetLastBgColor(self.handle()))
}

func (self ImGuiDockNode) SetHostWindow(v ImGuiWindow) {
	C.ImGuiDockNode_SetHostWindow(self.handle(), v.handle())
}

func (self ImGuiDockNode) GetHostWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiDockNode_GetHostWindow(self.handle())))
}

func (self ImGuiDockNode) SetVisibleWindow(v ImGuiWindow) {
	C.ImGuiDockNode_SetVisibleWindow(self.handle(), v.handle())
}

func (self ImGuiDockNode) GetVisibleWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiDockNode_GetVisibleWindow(self.handle())))
}

func (self ImGuiDockNode) SetCentralNode(v ImGuiDockNode) {
	C.ImGuiDockNode_SetCentralNode(self.handle(), v.handle())
}

func (self ImGuiDockNode) GetCentralNode() ImGuiDockNode {
	return (ImGuiDockNode)(unsafe.Pointer(C.ImGuiDockNode_GetCentralNode(self.handle())))
}

func (self ImGuiDockNode) SetOnlyNodeWithWindows(v ImGuiDockNode) {
	C.ImGuiDockNode_SetOnlyNodeWithWindows(self.handle(), v.handle())
}

func (self ImGuiDockNode) GetOnlyNodeWithWindows() ImGuiDockNode {
	return (ImGuiDockNode)(unsafe.Pointer(C.ImGuiDockNode_GetOnlyNodeWithWindows(self.handle())))
}

func (self ImGuiDockNode) SetCountNodeWithWindows(v int32) {
	C.ImGuiDockNode_SetCountNodeWithWindows(self.handle(), C.int(v))
}

func (self ImGuiDockNode) GetCountNodeWithWindows() int {
	return int(C.ImGuiDockNode_GetCountNodeWithWindows(self.handle()))
}

func (self ImGuiDockNode) SetLastFrameAlive(v int32) {
	C.ImGuiDockNode_SetLastFrameAlive(self.handle(), C.int(v))
}

func (self ImGuiDockNode) GetLastFrameAlive() int {
	return int(C.ImGuiDockNode_GetLastFrameAlive(self.handle()))
}

func (self ImGuiDockNode) SetLastFrameActive(v int32) {
	C.ImGuiDockNode_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiDockNode) GetLastFrameActive() int {
	return int(C.ImGuiDockNode_GetLastFrameActive(self.handle()))
}

func (self ImGuiDockNode) SetLastFrameFocused(v int32) {
	C.ImGuiDockNode_SetLastFrameFocused(self.handle(), C.int(v))
}

func (self ImGuiDockNode) GetLastFrameFocused() int {
	return int(C.ImGuiDockNode_GetLastFrameFocused(self.handle()))
}

func (self ImGuiDockNode) SetLastFocusedNodeId(v ImGuiID) {
	C.ImGuiDockNode_SetLastFocusedNodeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) GetLastFocusedNodeId() ImGuiID {
	return ImGuiID(C.ImGuiDockNode_GetLastFocusedNodeId(self.handle()))
}

func (self ImGuiDockNode) SetSelectedTabId(v ImGuiID) {
	C.ImGuiDockNode_SetSelectedTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) GetSelectedTabId() ImGuiID {
	return ImGuiID(C.ImGuiDockNode_GetSelectedTabId(self.handle()))
}

func (self ImGuiDockNode) SetWantCloseTabId(v ImGuiID) {
	C.ImGuiDockNode_SetWantCloseTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) GetWantCloseTabId() ImGuiID {
	return ImGuiID(C.ImGuiDockNode_GetWantCloseTabId(self.handle()))
}

func (self ImGuiDockNode) SetAuthorityForPos(v ImGuiDataAuthority) {
	C.ImGuiDockNode_SetAuthorityForPos(self.handle(), C.ImGuiDataAuthority(v))
}

func (self ImGuiDockNode) GetAuthorityForPos() ImGuiDataAuthority {
	return ImGuiDataAuthority(C.ImGuiDockNode_GetAuthorityForPos(self.handle()))
}

func (self ImGuiDockNode) SetAuthorityForSize(v ImGuiDataAuthority) {
	C.ImGuiDockNode_SetAuthorityForSize(self.handle(), C.ImGuiDataAuthority(v))
}

func (self ImGuiDockNode) GetAuthorityForSize() ImGuiDataAuthority {
	return ImGuiDataAuthority(C.ImGuiDockNode_GetAuthorityForSize(self.handle()))
}

func (self ImGuiDockNode) SetAuthorityForViewport(v ImGuiDataAuthority) {
	C.ImGuiDockNode_SetAuthorityForViewport(self.handle(), C.ImGuiDataAuthority(v))
}

func (self ImGuiDockNode) GetAuthorityForViewport() ImGuiDataAuthority {
	return ImGuiDataAuthority(C.ImGuiDockNode_GetAuthorityForViewport(self.handle()))
}

func (self ImGuiDockNode) SetIsVisible(v bool) {
	C.ImGuiDockNode_SetIsVisible(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetIsVisible() bool {
	return C.ImGuiDockNode_GetIsVisible(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetIsFocused(v bool) {
	C.ImGuiDockNode_SetIsFocused(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetIsFocused() bool {
	return C.ImGuiDockNode_GetIsFocused(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetIsBgDrawnThisFrame(v bool) {
	C.ImGuiDockNode_SetIsBgDrawnThisFrame(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetIsBgDrawnThisFrame() bool {
	return C.ImGuiDockNode_GetIsBgDrawnThisFrame(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetHasCloseButton(v bool) {
	C.ImGuiDockNode_SetHasCloseButton(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetHasCloseButton() bool {
	return C.ImGuiDockNode_GetHasCloseButton(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetHasWindowMenuButton(v bool) {
	C.ImGuiDockNode_SetHasWindowMenuButton(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetHasWindowMenuButton() bool {
	return C.ImGuiDockNode_GetHasWindowMenuButton(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetHasCentralNodeChild(v bool) {
	C.ImGuiDockNode_SetHasCentralNodeChild(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetHasCentralNodeChild() bool {
	return C.ImGuiDockNode_GetHasCentralNodeChild(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetWantCloseAll(v bool) {
	C.ImGuiDockNode_SetWantCloseAll(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetWantCloseAll() bool {
	return C.ImGuiDockNode_GetWantCloseAll(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetWantLockSizeOnce(v bool) {
	C.ImGuiDockNode_SetWantLockSizeOnce(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetWantLockSizeOnce() bool {
	return C.ImGuiDockNode_GetWantLockSizeOnce(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetWantMouseMove(v bool) {
	C.ImGuiDockNode_SetWantMouseMove(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetWantMouseMove() bool {
	return C.ImGuiDockNode_GetWantMouseMove(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetWantHiddenTabBarUpdate(v bool) {
	C.ImGuiDockNode_SetWantHiddenTabBarUpdate(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetWantHiddenTabBarUpdate() bool {
	return C.ImGuiDockNode_GetWantHiddenTabBarUpdate(self.handle()) == C.bool(true)
}

func (self ImGuiDockNode) SetWantHiddenTabBarToggle(v bool) {
	C.ImGuiDockNode_SetWantHiddenTabBarToggle(self.handle(), C.bool(v))
}

func (self ImGuiDockNode) GetWantHiddenTabBarToggle() bool {
	return C.ImGuiDockNode_GetWantHiddenTabBarToggle(self.handle()) == C.bool(true)
}

func (self ImGuiGroupData) SetWindowID(v ImGuiID) {
	C.ImGuiGroupData_SetWindowID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiGroupData) GetWindowID() ImGuiID {
	return ImGuiID(C.ImGuiGroupData_GetWindowID(self.handle()))
}

func (self ImGuiGroupData) SetBackupCursorPos(v ImVec2) {
	C.ImGuiGroupData_SetBackupCursorPos(self.handle(), v.toC())
}

func (self ImGuiGroupData) GetBackupCursorPos() ImVec2 {
	return newImVec2FromC(C.ImGuiGroupData_GetBackupCursorPos(self.handle()))
}

func (self ImGuiGroupData) SetBackupCursorMaxPos(v ImVec2) {
	C.ImGuiGroupData_SetBackupCursorMaxPos(self.handle(), v.toC())
}

func (self ImGuiGroupData) GetBackupCursorMaxPos() ImVec2 {
	return newImVec2FromC(C.ImGuiGroupData_GetBackupCursorMaxPos(self.handle()))
}

func (self ImGuiGroupData) SetBackupCurrLineSize(v ImVec2) {
	C.ImGuiGroupData_SetBackupCurrLineSize(self.handle(), v.toC())
}

func (self ImGuiGroupData) GetBackupCurrLineSize() ImVec2 {
	return newImVec2FromC(C.ImGuiGroupData_GetBackupCurrLineSize(self.handle()))
}

func (self ImGuiGroupData) SetBackupCurrLineTextBaseOffset(v float32) {
	C.ImGuiGroupData_SetBackupCurrLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiGroupData) GetBackupCurrLineTextBaseOffset() float32 {
	return float32(C.ImGuiGroupData_GetBackupCurrLineTextBaseOffset(self.handle()))
}

func (self ImGuiGroupData) SetBackupActiveIdIsAlive(v ImGuiID) {
	C.ImGuiGroupData_SetBackupActiveIdIsAlive(self.handle(), C.ImGuiID(v))
}

func (self ImGuiGroupData) GetBackupActiveIdIsAlive() ImGuiID {
	return ImGuiID(C.ImGuiGroupData_GetBackupActiveIdIsAlive(self.handle()))
}

func (self ImGuiGroupData) SetBackupActiveIdPreviousFrameIsAlive(v bool) {
	C.ImGuiGroupData_SetBackupActiveIdPreviousFrameIsAlive(self.handle(), C.bool(v))
}

func (self ImGuiGroupData) GetBackupActiveIdPreviousFrameIsAlive() bool {
	return C.ImGuiGroupData_GetBackupActiveIdPreviousFrameIsAlive(self.handle()) == C.bool(true)
}

func (self ImGuiGroupData) SetBackupHoveredIdIsAlive(v bool) {
	C.ImGuiGroupData_SetBackupHoveredIdIsAlive(self.handle(), C.bool(v))
}

func (self ImGuiGroupData) GetBackupHoveredIdIsAlive() bool {
	return C.ImGuiGroupData_GetBackupHoveredIdIsAlive(self.handle()) == C.bool(true)
}

func (self ImGuiGroupData) SetEmitItem(v bool) {
	C.ImGuiGroupData_SetEmitItem(self.handle(), C.bool(v))
}

func (self ImGuiGroupData) GetEmitItem() bool {
	return C.ImGuiGroupData_GetEmitItem(self.handle()) == C.bool(true)
}

func (self ImGuiInputEvent) SetType(v ImGuiInputEventType) {
	C.ImGuiInputEvent_SetType(self.handle(), C.ImGuiInputEventType(v))
}

func (self ImGuiInputEvent) GetType() ImGuiInputEventType {
	return ImGuiInputEventType(C.ImGuiInputEvent_GetType(self.handle()))
}

func (self ImGuiInputEvent) SetSource(v ImGuiInputSource) {
	C.ImGuiInputEvent_SetSource(self.handle(), C.ImGuiInputSource(v))
}

func (self ImGuiInputEvent) GetSource() ImGuiInputSource {
	return ImGuiInputSource(C.ImGuiInputEvent_GetSource(self.handle()))
}

func (self ImGuiInputEvent) SetIgnoredAsSame(v bool) {
	C.ImGuiInputEvent_SetIgnoredAsSame(self.handle(), C.bool(v))
}

func (self ImGuiInputEvent) GetIgnoredAsSame() bool {
	return C.ImGuiInputEvent_GetIgnoredAsSame(self.handle()) == C.bool(true)
}

func (self ImGuiInputEvent) SetAddedByTestEngine(v bool) {
	C.ImGuiInputEvent_SetAddedByTestEngine(self.handle(), C.bool(v))
}

func (self ImGuiInputEvent) GetAddedByTestEngine() bool {
	return C.ImGuiInputEvent_GetAddedByTestEngine(self.handle()) == C.bool(true)
}

func (self ImGuiKeyData) SetDown(v bool) {
	C.ImGuiKeyData_SetDown(self.handle(), C.bool(v))
}

func (self ImGuiKeyData) GetDown() bool {
	return C.ImGuiKeyData_GetDown(self.handle()) == C.bool(true)
}

func (self ImGuiKeyData) SetDownDuration(v float32) {
	C.ImGuiKeyData_SetDownDuration(self.handle(), C.float(v))
}

func (self ImGuiKeyData) GetDownDuration() float32 {
	return float32(C.ImGuiKeyData_GetDownDuration(self.handle()))
}

func (self ImGuiKeyData) SetDownDurationPrev(v float32) {
	C.ImGuiKeyData_SetDownDurationPrev(self.handle(), C.float(v))
}

func (self ImGuiKeyData) GetDownDurationPrev() float32 {
	return float32(C.ImGuiKeyData_GetDownDurationPrev(self.handle()))
}

func (self ImGuiKeyData) SetAnalogValue(v float32) {
	C.ImGuiKeyData_SetAnalogValue(self.handle(), C.float(v))
}

func (self ImGuiKeyData) GetAnalogValue() float32 {
	return float32(C.ImGuiKeyData_GetAnalogValue(self.handle()))
}

func (self ImGuiTableSortSpecs) SetSpecs(v ImGuiTableColumnSortSpecs) {
	C.ImGuiTableSortSpecs_SetSpecs(self.handle(), v.handle())
}

func (self ImGuiTableSortSpecs) GetSpecs() ImGuiTableColumnSortSpecs {
	return (ImGuiTableColumnSortSpecs)(unsafe.Pointer(C.ImGuiTableSortSpecs_GetSpecs(self.handle())))
}

func (self ImGuiTableSortSpecs) SetSpecsCount(v int32) {
	C.ImGuiTableSortSpecs_SetSpecsCount(self.handle(), C.int(v))
}

func (self ImGuiTableSortSpecs) GetSpecsCount() int {
	return int(C.ImGuiTableSortSpecs_GetSpecsCount(self.handle()))
}

func (self ImGuiTableSortSpecs) SetSpecsDirty(v bool) {
	C.ImGuiTableSortSpecs_SetSpecsDirty(self.handle(), C.bool(v))
}

func (self ImGuiTableSortSpecs) GetSpecsDirty() bool {
	return C.ImGuiTableSortSpecs_GetSpecsDirty(self.handle()) == C.bool(true)
}

func (self ImDrawList) SetFlags(v ImDrawListFlags) {
	C.ImDrawList_SetFlags(self.handle(), C.ImDrawListFlags(v))
}

func (self ImDrawList) GetFlags() ImDrawListFlags {
	return ImDrawListFlags(C.ImDrawList_GetFlags(self.handle()))
}

func (self ImDrawList) Set_VtxCurrentIdx(v uint32) {
	C.ImDrawList_Set_VtxCurrentIdx(self.handle(), C.uint(v))
}

func (self ImDrawList) Get_VtxCurrentIdx() uint32 {
	return uint32(C.ImDrawList_Get_VtxCurrentIdx(self.handle()))
}

func (self ImDrawList) Set_Data(v ImDrawListSharedData) {
	C.ImDrawList_Set_Data(self.handle(), v.handle())
}

func (self ImDrawList) Get_Data() ImDrawListSharedData {
	return (ImDrawListSharedData)(unsafe.Pointer(C.ImDrawList_Get_Data(self.handle())))
}

func (self ImDrawList) Set_OwnerName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImDrawList_Set_OwnerName(self.handle(), vArg)
}

func (self ImDrawList) Get_OwnerName() string {
	return C.GoString(C.ImDrawList_Get_OwnerName(self.handle()))
}

func (self ImDrawList) Set_VtxWritePtr(v ImDrawVert) {
	C.ImDrawList_Set_VtxWritePtr(self.handle(), v.handle())
}

func (self ImDrawList) Get_VtxWritePtr() ImDrawVert {
	return (ImDrawVert)(unsafe.Pointer(C.ImDrawList_Get_VtxWritePtr(self.handle())))
}

func (self ImDrawList) Get_CmdHeader() ImDrawCmdHeader {
	return newImDrawCmdHeaderFromC(C.ImDrawList_Get_CmdHeader(self.handle()))
}

func (self ImDrawList) Get_Splitter() ImDrawListSplitter {
	return newImDrawListSplitterFromC(C.ImDrawList_Get_Splitter(self.handle()))
}

func (self ImDrawList) Set_FringeScale(v float32) {
	C.ImDrawList_Set_FringeScale(self.handle(), C.float(v))
}

func (self ImDrawList) Get_FringeScale() float32 {
	return float32(C.ImDrawList_Get_FringeScale(self.handle()))
}

func (self ImFont) SetFallbackAdvanceX(v float32) {
	C.ImFont_SetFallbackAdvanceX(self.handle(), C.float(v))
}

func (self ImFont) GetFallbackAdvanceX() float32 {
	return float32(C.ImFont_GetFallbackAdvanceX(self.handle()))
}

func (self ImFont) SetFontSize(v float32) {
	C.ImFont_SetFontSize(self.handle(), C.float(v))
}

func (self ImFont) GetFontSize() float32 {
	return float32(C.ImFont_GetFontSize(self.handle()))
}

func (self ImFont) SetFallbackGlyph(v ImFontGlyph) {
	C.ImFont_SetFallbackGlyph(self.handle(), v.handle())
}

func (self ImFont) GetFallbackGlyph() ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.ImFont_GetFallbackGlyph(self.handle())))
}

func (self ImFont) SetContainerAtlas(v ImFontAtlas) {
	C.ImFont_SetContainerAtlas(self.handle(), v.handle())
}

func (self ImFont) GetContainerAtlas() ImFontAtlas {
	return (ImFontAtlas)(unsafe.Pointer(C.ImFont_GetContainerAtlas(self.handle())))
}

func (self ImFont) SetConfigData(v ImFontConfig) {
	C.ImFont_SetConfigData(self.handle(), v.handle())
}

func (self ImFont) GetConfigData() ImFontConfig {
	return (ImFontConfig)(unsafe.Pointer(C.ImFont_GetConfigData(self.handle())))
}

func (self ImFont) SetConfigDataCount(v int) {
	C.ImFont_SetConfigDataCount(self.handle(), C.short(v))
}

func (self ImFont) GetConfigDataCount() int {
	return int(C.ImFont_GetConfigDataCount(self.handle()))
}

func (self ImFont) SetFallbackChar(v ImWchar) {
	C.ImFont_SetFallbackChar(self.handle(), C.ImWchar(v))
}

func (self ImFont) SetEllipsisChar(v ImWchar) {
	C.ImFont_SetEllipsisChar(self.handle(), C.ImWchar(v))
}

func (self ImFont) SetDotChar(v ImWchar) {
	C.ImFont_SetDotChar(self.handle(), C.ImWchar(v))
}

func (self ImFont) SetDirtyLookupTables(v bool) {
	C.ImFont_SetDirtyLookupTables(self.handle(), C.bool(v))
}

func (self ImFont) GetDirtyLookupTables() bool {
	return C.ImFont_GetDirtyLookupTables(self.handle()) == C.bool(true)
}

func (self ImFont) SetScale(v float32) {
	C.ImFont_SetScale(self.handle(), C.float(v))
}

func (self ImFont) GetScale() float32 {
	return float32(C.ImFont_GetScale(self.handle()))
}

func (self ImFont) SetAscent(v float32) {
	C.ImFont_SetAscent(self.handle(), C.float(v))
}

func (self ImFont) GetAscent() float32 {
	return float32(C.ImFont_GetAscent(self.handle()))
}

func (self ImFont) SetDescent(v float32) {
	C.ImFont_SetDescent(self.handle(), C.float(v))
}

func (self ImFont) GetDescent() float32 {
	return float32(C.ImFont_GetDescent(self.handle()))
}

func (self ImFont) SetMetricsTotalSurface(v int32) {
	C.ImFont_SetMetricsTotalSurface(self.handle(), C.int(v))
}

func (self ImFont) GetMetricsTotalSurface() int {
	return int(C.ImFont_GetMetricsTotalSurface(self.handle()))
}

func (self ImGuiContextHook) SetHookId(v ImGuiID) {
	C.ImGuiContextHook_SetHookId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContextHook) GetHookId() ImGuiID {
	return ImGuiID(C.ImGuiContextHook_GetHookId(self.handle()))
}

func (self ImGuiContextHook) SetType(v ImGuiContextHookType) {
	C.ImGuiContextHook_SetType(self.handle(), C.ImGuiContextHookType(v))
}

func (self ImGuiContextHook) GetType() ImGuiContextHookType {
	return ImGuiContextHookType(C.ImGuiContextHook_GetType(self.handle()))
}

func (self ImGuiContextHook) SetOwner(v ImGuiID) {
	C.ImGuiContextHook_SetOwner(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContextHook) GetOwner() ImGuiID {
	return ImGuiID(C.ImGuiContextHook_GetOwner(self.handle()))
}

func (self ImGuiContextHook) SetUserData(v unsafe.Pointer) {
	C.ImGuiContextHook_SetUserData(self.handle(), v)
}

func (self ImGuiContextHook) GetUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiContextHook_GetUserData(self.handle()))
}

func (self ImGuiListClipper) SetDisplayStart(v int32) {
	C.ImGuiListClipper_SetDisplayStart(self.handle(), C.int(v))
}

func (self ImGuiListClipper) GetDisplayStart() int {
	return int(C.ImGuiListClipper_GetDisplayStart(self.handle()))
}

func (self ImGuiListClipper) SetDisplayEnd(v int32) {
	C.ImGuiListClipper_SetDisplayEnd(self.handle(), C.int(v))
}

func (self ImGuiListClipper) GetDisplayEnd() int {
	return int(C.ImGuiListClipper_GetDisplayEnd(self.handle()))
}

func (self ImGuiListClipper) SetItemsCount(v int32) {
	C.ImGuiListClipper_SetItemsCount(self.handle(), C.int(v))
}

func (self ImGuiListClipper) GetItemsCount() int {
	return int(C.ImGuiListClipper_GetItemsCount(self.handle()))
}

func (self ImGuiListClipper) SetItemsHeight(v float32) {
	C.ImGuiListClipper_SetItemsHeight(self.handle(), C.float(v))
}

func (self ImGuiListClipper) GetItemsHeight() float32 {
	return float32(C.ImGuiListClipper_GetItemsHeight(self.handle()))
}

func (self ImGuiListClipper) SetStartPosY(v float32) {
	C.ImGuiListClipper_SetStartPosY(self.handle(), C.float(v))
}

func (self ImGuiListClipper) GetStartPosY() float32 {
	return float32(C.ImGuiListClipper_GetStartPosY(self.handle()))
}

func (self ImGuiListClipper) SetTempData(v unsafe.Pointer) {
	C.ImGuiListClipper_SetTempData(self.handle(), v)
}

func (self ImGuiListClipper) GetTempData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiListClipper_GetTempData(self.handle()))
}

func (self ImGuiTableCellData) SetBgColor(v uint32) {
	C.ImGuiTableCellData_SetBgColor(self.handle(), C.ImU32(v))
}

func (self ImGuiTableCellData) GetBgColor() uint32 {
	return uint32(C.ImGuiTableCellData_GetBgColor(self.handle()))
}

func (self ImGuiTableCellData) SetColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTableCellData_SetColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableCellData) GetColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableCellData_GetColumn(self.handle()))
}

func (self ImDrawCmdHeader) SetClipRect(v ImVec4) {
	C.ImDrawCmdHeader_SetClipRect(self.handle(), v.toC())
}

func (self ImDrawCmdHeader) GetClipRect() ImVec4 {
	return newImVec4FromC(C.ImDrawCmdHeader_GetClipRect(self.handle()))
}

func (self ImDrawCmdHeader) SetTextureId(v ImTextureID) {
	C.ImDrawCmdHeader_SetTextureId(self.handle(), C.ImTextureID(v))
}

func (self ImDrawCmdHeader) GetTextureId() ImTextureID {
	return ImTextureID(C.ImDrawCmdHeader_GetTextureId(self.handle()))
}

func (self ImDrawCmdHeader) SetVtxOffset(v uint32) {
	C.ImDrawCmdHeader_SetVtxOffset(self.handle(), C.uint(v))
}

func (self ImDrawCmdHeader) GetVtxOffset() uint32 {
	return uint32(C.ImDrawCmdHeader_GetVtxOffset(self.handle()))
}

func (self ImFontAtlas) SetFlags(v ImFontAtlasFlags) {
	C.ImFontAtlas_SetFlags(self.handle(), C.ImFontAtlasFlags(v))
}

func (self ImFontAtlas) GetFlags() ImFontAtlasFlags {
	return ImFontAtlasFlags(C.ImFontAtlas_GetFlags(self.handle()))
}

func (self ImFontAtlas) SetTexDesiredWidth(v int32) {
	C.ImFontAtlas_SetTexDesiredWidth(self.handle(), C.int(v))
}

func (self ImFontAtlas) GetTexDesiredWidth() int {
	return int(C.ImFontAtlas_GetTexDesiredWidth(self.handle()))
}

func (self ImFontAtlas) SetTexGlyphPadding(v int32) {
	C.ImFontAtlas_SetTexGlyphPadding(self.handle(), C.int(v))
}

func (self ImFontAtlas) GetTexGlyphPadding() int {
	return int(C.ImFontAtlas_GetTexGlyphPadding(self.handle()))
}

func (self ImFontAtlas) SetLocked(v bool) {
	C.ImFontAtlas_SetLocked(self.handle(), C.bool(v))
}

func (self ImFontAtlas) GetLocked() bool {
	return C.ImFontAtlas_GetLocked(self.handle()) == C.bool(true)
}

func (self ImFontAtlas) SetTexReady(v bool) {
	C.ImFontAtlas_SetTexReady(self.handle(), C.bool(v))
}

func (self ImFontAtlas) GetTexReady() bool {
	return C.ImFontAtlas_GetTexReady(self.handle()) == C.bool(true)
}

func (self ImFontAtlas) SetTexPixelsUseColors(v bool) {
	C.ImFontAtlas_SetTexPixelsUseColors(self.handle(), C.bool(v))
}

func (self ImFontAtlas) GetTexPixelsUseColors() bool {
	return C.ImFontAtlas_GetTexPixelsUseColors(self.handle()) == C.bool(true)
}

func (self ImFontAtlas) SetTexWidth(v int32) {
	C.ImFontAtlas_SetTexWidth(self.handle(), C.int(v))
}

func (self ImFontAtlas) GetTexWidth() int {
	return int(C.ImFontAtlas_GetTexWidth(self.handle()))
}

func (self ImFontAtlas) SetTexHeight(v int32) {
	C.ImFontAtlas_SetTexHeight(self.handle(), C.int(v))
}

func (self ImFontAtlas) GetTexHeight() int {
	return int(C.ImFontAtlas_GetTexHeight(self.handle()))
}

func (self ImFontAtlas) SetTexUvScale(v ImVec2) {
	C.ImFontAtlas_SetTexUvScale(self.handle(), v.toC())
}

func (self ImFontAtlas) GetTexUvScale() ImVec2 {
	return newImVec2FromC(C.ImFontAtlas_GetTexUvScale(self.handle()))
}

func (self ImFontAtlas) SetTexUvWhitePixel(v ImVec2) {
	C.ImFontAtlas_SetTexUvWhitePixel(self.handle(), v.toC())
}

func (self ImFontAtlas) GetTexUvWhitePixel() ImVec2 {
	return newImVec2FromC(C.ImFontAtlas_GetTexUvWhitePixel(self.handle()))
}

func (self ImFontAtlas) SetFontBuilderIO(v ImFontBuilderIO) {
	C.ImFontAtlas_SetFontBuilderIO(self.handle(), v.handle())
}

func (self ImFontAtlas) GetFontBuilderIO() ImFontBuilderIO {
	return (ImFontBuilderIO)(unsafe.Pointer(C.ImFontAtlas_GetFontBuilderIO(self.handle())))
}

func (self ImFontAtlas) SetFontBuilderFlags(v uint32) {
	C.ImFontAtlas_SetFontBuilderFlags(self.handle(), C.uint(v))
}

func (self ImFontAtlas) GetFontBuilderFlags() uint32 {
	return uint32(C.ImFontAtlas_GetFontBuilderFlags(self.handle()))
}

func (self ImFontAtlas) SetPackIdMouseCursors(v int32) {
	C.ImFontAtlas_SetPackIdMouseCursors(self.handle(), C.int(v))
}

func (self ImFontAtlas) GetPackIdMouseCursors() int {
	return int(C.ImFontAtlas_GetPackIdMouseCursors(self.handle()))
}

func (self ImFontAtlas) SetPackIdLines(v int32) {
	C.ImFontAtlas_SetPackIdLines(self.handle(), C.int(v))
}

func (self ImFontAtlas) GetPackIdLines() int {
	return int(C.ImFontAtlas_GetPackIdLines(self.handle()))
}

func (self ImGuiInputEventMouseButton) SetButton(v int32) {
	C.ImGuiInputEventMouseButton_SetButton(self.handle(), C.int(v))
}

func (self ImGuiInputEventMouseButton) GetButton() int {
	return int(C.ImGuiInputEventMouseButton_GetButton(self.handle()))
}

func (self ImGuiInputEventMouseButton) SetDown(v bool) {
	C.ImGuiInputEventMouseButton_SetDown(self.handle(), C.bool(v))
}

func (self ImGuiInputEventMouseButton) GetDown() bool {
	return C.ImGuiInputEventMouseButton_GetDown(self.handle()) == C.bool(true)
}

func (self ImGuiOnceUponAFrame) SetRefFrame(v int32) {
	C.ImGuiOnceUponAFrame_SetRefFrame(self.handle(), C.int(v))
}

func (self ImGuiOnceUponAFrame) GetRefFrame() int {
	return int(C.ImGuiOnceUponAFrame_GetRefFrame(self.handle()))
}

func (self ImGuiTabItem) SetID(v ImGuiID) {
	C.ImGuiTabItem_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabItem) GetID() ImGuiID {
	return ImGuiID(C.ImGuiTabItem_GetID(self.handle()))
}

func (self ImGuiTabItem) SetFlags(v ImGuiTabItemFlags) {
	C.ImGuiTabItem_SetFlags(self.handle(), C.ImGuiTabItemFlags(v))
}

func (self ImGuiTabItem) GetFlags() ImGuiTabItemFlags {
	return ImGuiTabItemFlags(C.ImGuiTabItem_GetFlags(self.handle()))
}

func (self ImGuiTabItem) SetWindow(v ImGuiWindow) {
	C.ImGuiTabItem_SetWindow(self.handle(), v.handle())
}

func (self ImGuiTabItem) GetWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiTabItem_GetWindow(self.handle())))
}

func (self ImGuiTabItem) SetLastFrameVisible(v int32) {
	C.ImGuiTabItem_SetLastFrameVisible(self.handle(), C.int(v))
}

func (self ImGuiTabItem) GetLastFrameVisible() int {
	return int(C.ImGuiTabItem_GetLastFrameVisible(self.handle()))
}

func (self ImGuiTabItem) SetLastFrameSelected(v int32) {
	C.ImGuiTabItem_SetLastFrameSelected(self.handle(), C.int(v))
}

func (self ImGuiTabItem) GetLastFrameSelected() int {
	return int(C.ImGuiTabItem_GetLastFrameSelected(self.handle()))
}

func (self ImGuiTabItem) SetOffset(v float32) {
	C.ImGuiTabItem_SetOffset(self.handle(), C.float(v))
}

func (self ImGuiTabItem) GetOffset() float32 {
	return float32(C.ImGuiTabItem_GetOffset(self.handle()))
}

func (self ImGuiTabItem) SetWidth(v float32) {
	C.ImGuiTabItem_SetWidth(self.handle(), C.float(v))
}

func (self ImGuiTabItem) GetWidth() float32 {
	return float32(C.ImGuiTabItem_GetWidth(self.handle()))
}

func (self ImGuiTabItem) SetContentWidth(v float32) {
	C.ImGuiTabItem_SetContentWidth(self.handle(), C.float(v))
}

func (self ImGuiTabItem) GetContentWidth() float32 {
	return float32(C.ImGuiTabItem_GetContentWidth(self.handle()))
}

func (self ImGuiTabItem) SetRequestedWidth(v float32) {
	C.ImGuiTabItem_SetRequestedWidth(self.handle(), C.float(v))
}

func (self ImGuiTabItem) GetRequestedWidth() float32 {
	return float32(C.ImGuiTabItem_GetRequestedWidth(self.handle()))
}

func (self ImGuiTabItem) SetNameOffset(v int) {
	C.ImGuiTabItem_SetNameOffset(self.handle(), C.ImS32(v))
}

func (self ImGuiTabItem) GetNameOffset() int {
	return int(C.ImGuiTabItem_GetNameOffset(self.handle()))
}

func (self ImGuiTabItem) SetBeginOrder(v int) {
	C.ImGuiTabItem_SetBeginOrder(self.handle(), C.ImS16(v))
}

func (self ImGuiTabItem) GetBeginOrder() int {
	return int(C.ImGuiTabItem_GetBeginOrder(self.handle()))
}

func (self ImGuiTabItem) SetIndexDuringLayout(v int) {
	C.ImGuiTabItem_SetIndexDuringLayout(self.handle(), C.ImS16(v))
}

func (self ImGuiTabItem) GetIndexDuringLayout() int {
	return int(C.ImGuiTabItem_GetIndexDuringLayout(self.handle()))
}

func (self ImGuiTabItem) SetWantClose(v bool) {
	C.ImGuiTabItem_SetWantClose(self.handle(), C.bool(v))
}

func (self ImGuiTabItem) GetWantClose() bool {
	return C.ImGuiTabItem_GetWantClose(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetFlags(v ImGuiTableColumnFlags) {
	C.ImGuiTableColumn_SetFlags(self.handle(), C.ImGuiTableColumnFlags(v))
}

func (self ImGuiTableColumn) GetFlags() ImGuiTableColumnFlags {
	return ImGuiTableColumnFlags(C.ImGuiTableColumn_GetFlags(self.handle()))
}

func (self ImGuiTableColumn) SetWidthGiven(v float32) {
	C.ImGuiTableColumn_SetWidthGiven(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetWidthGiven() float32 {
	return float32(C.ImGuiTableColumn_GetWidthGiven(self.handle()))
}

func (self ImGuiTableColumn) SetMinX(v float32) {
	C.ImGuiTableColumn_SetMinX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetMinX() float32 {
	return float32(C.ImGuiTableColumn_GetMinX(self.handle()))
}

func (self ImGuiTableColumn) SetMaxX(v float32) {
	C.ImGuiTableColumn_SetMaxX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetMaxX() float32 {
	return float32(C.ImGuiTableColumn_GetMaxX(self.handle()))
}

func (self ImGuiTableColumn) SetWidthRequest(v float32) {
	C.ImGuiTableColumn_SetWidthRequest(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetWidthRequest() float32 {
	return float32(C.ImGuiTableColumn_GetWidthRequest(self.handle()))
}

func (self ImGuiTableColumn) SetWidthAuto(v float32) {
	C.ImGuiTableColumn_SetWidthAuto(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetWidthAuto() float32 {
	return float32(C.ImGuiTableColumn_GetWidthAuto(self.handle()))
}

func (self ImGuiTableColumn) SetStretchWeight(v float32) {
	C.ImGuiTableColumn_SetStretchWeight(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetStretchWeight() float32 {
	return float32(C.ImGuiTableColumn_GetStretchWeight(self.handle()))
}

func (self ImGuiTableColumn) SetInitStretchWeightOrWidth(v float32) {
	C.ImGuiTableColumn_SetInitStretchWeightOrWidth(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetInitStretchWeightOrWidth() float32 {
	return float32(C.ImGuiTableColumn_GetInitStretchWeightOrWidth(self.handle()))
}

func (self ImGuiTableColumn) SetClipRect(v ImRect) {
	C.ImGuiTableColumn_SetClipRect(self.handle(), v.toC())
}

func (self ImGuiTableColumn) GetClipRect() ImRect {
	return newImRectFromC(C.ImGuiTableColumn_GetClipRect(self.handle()))
}

func (self ImGuiTableColumn) SetUserID(v ImGuiID) {
	C.ImGuiTableColumn_SetUserID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableColumn) GetUserID() ImGuiID {
	return ImGuiID(C.ImGuiTableColumn_GetUserID(self.handle()))
}

func (self ImGuiTableColumn) SetWorkMinX(v float32) {
	C.ImGuiTableColumn_SetWorkMinX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetWorkMinX() float32 {
	return float32(C.ImGuiTableColumn_GetWorkMinX(self.handle()))
}

func (self ImGuiTableColumn) SetWorkMaxX(v float32) {
	C.ImGuiTableColumn_SetWorkMaxX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetWorkMaxX() float32 {
	return float32(C.ImGuiTableColumn_GetWorkMaxX(self.handle()))
}

func (self ImGuiTableColumn) SetItemWidth(v float32) {
	C.ImGuiTableColumn_SetItemWidth(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetItemWidth() float32 {
	return float32(C.ImGuiTableColumn_GetItemWidth(self.handle()))
}

func (self ImGuiTableColumn) SetContentMaxXFrozen(v float32) {
	C.ImGuiTableColumn_SetContentMaxXFrozen(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetContentMaxXFrozen() float32 {
	return float32(C.ImGuiTableColumn_GetContentMaxXFrozen(self.handle()))
}

func (self ImGuiTableColumn) SetContentMaxXUnfrozen(v float32) {
	C.ImGuiTableColumn_SetContentMaxXUnfrozen(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetContentMaxXUnfrozen() float32 {
	return float32(C.ImGuiTableColumn_GetContentMaxXUnfrozen(self.handle()))
}

func (self ImGuiTableColumn) SetContentMaxXHeadersUsed(v float32) {
	C.ImGuiTableColumn_SetContentMaxXHeadersUsed(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetContentMaxXHeadersUsed() float32 {
	return float32(C.ImGuiTableColumn_GetContentMaxXHeadersUsed(self.handle()))
}

func (self ImGuiTableColumn) SetContentMaxXHeadersIdeal(v float32) {
	C.ImGuiTableColumn_SetContentMaxXHeadersIdeal(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) GetContentMaxXHeadersIdeal() float32 {
	return float32(C.ImGuiTableColumn_GetContentMaxXHeadersIdeal(self.handle()))
}

func (self ImGuiTableColumn) SetNameOffset(v int) {
	C.ImGuiTableColumn_SetNameOffset(self.handle(), C.ImS16(v))
}

func (self ImGuiTableColumn) GetNameOffset() int {
	return int(C.ImGuiTableColumn_GetNameOffset(self.handle()))
}

func (self ImGuiTableColumn) SetDisplayOrder(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumn_SetDisplayOrder(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumn) GetDisplayOrder() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumn_GetDisplayOrder(self.handle()))
}

func (self ImGuiTableColumn) SetIndexWithinEnabledSet(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumn_SetIndexWithinEnabledSet(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumn) GetIndexWithinEnabledSet() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumn_GetIndexWithinEnabledSet(self.handle()))
}

func (self ImGuiTableColumn) SetPrevEnabledColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumn_SetPrevEnabledColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumn) GetPrevEnabledColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumn_GetPrevEnabledColumn(self.handle()))
}

func (self ImGuiTableColumn) SetNextEnabledColumn(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumn_SetNextEnabledColumn(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumn) GetNextEnabledColumn() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumn_GetNextEnabledColumn(self.handle()))
}

func (self ImGuiTableColumn) SetSortOrder(v ImGuiTableColumnIdx) {
	C.ImGuiTableColumn_SetSortOrder(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableColumn) GetSortOrder() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableColumn_GetSortOrder(self.handle()))
}

func (self ImGuiTableColumn) SetDrawChannelCurrent(v ImGuiTableDrawChannelIdx) {
	C.ImGuiTableColumn_SetDrawChannelCurrent(self.handle(), C.ImGuiTableDrawChannelIdx(v))
}

func (self ImGuiTableColumn) GetDrawChannelCurrent() ImGuiTableDrawChannelIdx {
	return ImGuiTableDrawChannelIdx(C.ImGuiTableColumn_GetDrawChannelCurrent(self.handle()))
}

func (self ImGuiTableColumn) SetDrawChannelFrozen(v ImGuiTableDrawChannelIdx) {
	C.ImGuiTableColumn_SetDrawChannelFrozen(self.handle(), C.ImGuiTableDrawChannelIdx(v))
}

func (self ImGuiTableColumn) GetDrawChannelFrozen() ImGuiTableDrawChannelIdx {
	return ImGuiTableDrawChannelIdx(C.ImGuiTableColumn_GetDrawChannelFrozen(self.handle()))
}

func (self ImGuiTableColumn) SetDrawChannelUnfrozen(v ImGuiTableDrawChannelIdx) {
	C.ImGuiTableColumn_SetDrawChannelUnfrozen(self.handle(), C.ImGuiTableDrawChannelIdx(v))
}

func (self ImGuiTableColumn) GetDrawChannelUnfrozen() ImGuiTableDrawChannelIdx {
	return ImGuiTableDrawChannelIdx(C.ImGuiTableColumn_GetDrawChannelUnfrozen(self.handle()))
}

func (self ImGuiTableColumn) SetIsEnabled(v bool) {
	C.ImGuiTableColumn_SetIsEnabled(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsEnabled() bool {
	return C.ImGuiTableColumn_GetIsEnabled(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsUserEnabled(v bool) {
	C.ImGuiTableColumn_SetIsUserEnabled(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsUserEnabled() bool {
	return C.ImGuiTableColumn_GetIsUserEnabled(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsUserEnabledNextFrame(v bool) {
	C.ImGuiTableColumn_SetIsUserEnabledNextFrame(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsUserEnabledNextFrame() bool {
	return C.ImGuiTableColumn_GetIsUserEnabledNextFrame(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsVisibleX(v bool) {
	C.ImGuiTableColumn_SetIsVisibleX(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsVisibleX() bool {
	return C.ImGuiTableColumn_GetIsVisibleX(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsVisibleY(v bool) {
	C.ImGuiTableColumn_SetIsVisibleY(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsVisibleY() bool {
	return C.ImGuiTableColumn_GetIsVisibleY(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsRequestOutput(v bool) {
	C.ImGuiTableColumn_SetIsRequestOutput(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsRequestOutput() bool {
	return C.ImGuiTableColumn_GetIsRequestOutput(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsSkipItems(v bool) {
	C.ImGuiTableColumn_SetIsSkipItems(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsSkipItems() bool {
	return C.ImGuiTableColumn_GetIsSkipItems(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetIsPreserveWidthAuto(v bool) {
	C.ImGuiTableColumn_SetIsPreserveWidthAuto(self.handle(), C.bool(v))
}

func (self ImGuiTableColumn) GetIsPreserveWidthAuto() bool {
	return C.ImGuiTableColumn_GetIsPreserveWidthAuto(self.handle()) == C.bool(true)
}

func (self ImGuiTableColumn) SetNavLayerCurrent(v int) {
	C.ImGuiTableColumn_SetNavLayerCurrent(self.handle(), C.ImS8(v))
}

func (self ImGuiTableColumn) GetNavLayerCurrent() int {
	return int(C.ImGuiTableColumn_GetNavLayerCurrent(self.handle()))
}

func (self ImGuiTableColumn) SetAutoFitQueue(v uint) {
	C.ImGuiTableColumn_SetAutoFitQueue(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumn) GetAutoFitQueue() uint32 {
	return uint32(C.ImGuiTableColumn_GetAutoFitQueue(self.handle()))
}

func (self ImGuiTableColumn) SetCannotSkipItemsQueue(v uint) {
	C.ImGuiTableColumn_SetCannotSkipItemsQueue(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumn) GetCannotSkipItemsQueue() uint32 {
	return uint32(C.ImGuiTableColumn_GetCannotSkipItemsQueue(self.handle()))
}

func (self ImGuiTableColumn) SetSortDirection(v uint) {
	C.ImGuiTableColumn_SetSortDirection(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumn) GetSortDirection() uint32 {
	return uint32(C.ImGuiTableColumn_GetSortDirection(self.handle()))
}

func (self ImGuiTableColumn) SetSortDirectionsAvailCount(v uint) {
	C.ImGuiTableColumn_SetSortDirectionsAvailCount(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumn) GetSortDirectionsAvailCount() uint32 {
	return uint32(C.ImGuiTableColumn_GetSortDirectionsAvailCount(self.handle()))
}

func (self ImGuiTableColumn) SetSortDirectionsAvailMask(v uint) {
	C.ImGuiTableColumn_SetSortDirectionsAvailMask(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumn) GetSortDirectionsAvailMask() uint32 {
	return uint32(C.ImGuiTableColumn_GetSortDirectionsAvailMask(self.handle()))
}

func (self ImGuiTableColumn) SetSortDirectionsAvailList(v uint) {
	C.ImGuiTableColumn_SetSortDirectionsAvailList(self.handle(), C.ImU8(v))
}

func (self ImGuiTableColumn) GetSortDirectionsAvailList() uint32 {
	return uint32(C.ImGuiTableColumn_GetSortDirectionsAvailList(self.handle()))
}

func (self ImGuiWindowStackData) SetWindow(v ImGuiWindow) {
	C.ImGuiWindowStackData_SetWindow(self.handle(), v.handle())
}

func (self ImGuiWindowStackData) GetWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindowStackData_GetWindow(self.handle())))
}

func (self ImGuiWindowStackData) GetParentLastItemDataBackup() ImGuiLastItemData {
	return newImGuiLastItemDataFromC(C.ImGuiWindowStackData_GetParentLastItemDataBackup(self.handle()))
}

func (self ImGuiWindowStackData) GetStackSizesOnBegin() ImGuiStackSizes {
	return newImGuiStackSizesFromC(C.ImGuiWindowStackData_GetStackSizesOnBegin(self.handle()))
}

func (self ImGuiContext) SetInitialized(v bool) {
	C.ImGuiContext_SetInitialized(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetInitialized() bool {
	return C.ImGuiContext_GetInitialized(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetFontAtlasOwnedByContext(v bool) {
	C.ImGuiContext_SetFontAtlasOwnedByContext(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetFontAtlasOwnedByContext() bool {
	return C.ImGuiContext_GetFontAtlasOwnedByContext(self.handle()) == C.bool(true)
}

func (self ImGuiContext) GetIO() ImGuiIO {
	return newImGuiIOFromC(C.ImGuiContext_GetIO(self.handle()))
}

func (self ImGuiContext) GetPlatformIO() ImGuiPlatformIO {
	return newImGuiPlatformIOFromC(C.ImGuiContext_GetPlatformIO(self.handle()))
}

func (self ImGuiContext) GetStyle() ImGuiStyle {
	return newImGuiStyleFromC(C.ImGuiContext_GetStyle(self.handle()))
}

func (self ImGuiContext) SetConfigFlagsCurrFrame(v ImGuiConfigFlags) {
	C.ImGuiContext_SetConfigFlagsCurrFrame(self.handle(), C.ImGuiConfigFlags(v))
}

func (self ImGuiContext) GetConfigFlagsCurrFrame() ImGuiConfigFlags {
	return ImGuiConfigFlags(C.ImGuiContext_GetConfigFlagsCurrFrame(self.handle()))
}

func (self ImGuiContext) SetConfigFlagsLastFrame(v ImGuiConfigFlags) {
	C.ImGuiContext_SetConfigFlagsLastFrame(self.handle(), C.ImGuiConfigFlags(v))
}

func (self ImGuiContext) GetConfigFlagsLastFrame() ImGuiConfigFlags {
	return ImGuiConfigFlags(C.ImGuiContext_GetConfigFlagsLastFrame(self.handle()))
}

func (self ImGuiContext) SetFont(v ImFont) {
	C.ImGuiContext_SetFont(self.handle(), v.handle())
}

func (self ImGuiContext) GetFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.ImGuiContext_GetFont(self.handle())))
}

func (self ImGuiContext) SetFontSize(v float32) {
	C.ImGuiContext_SetFontSize(self.handle(), C.float(v))
}

func (self ImGuiContext) GetFontSize() float32 {
	return float32(C.ImGuiContext_GetFontSize(self.handle()))
}

func (self ImGuiContext) SetFontBaseSize(v float32) {
	C.ImGuiContext_SetFontBaseSize(self.handle(), C.float(v))
}

func (self ImGuiContext) GetFontBaseSize() float32 {
	return float32(C.ImGuiContext_GetFontBaseSize(self.handle()))
}

func (self ImGuiContext) GetDrawListSharedData() ImDrawListSharedData {
	return newImDrawListSharedDataFromC(C.ImGuiContext_GetDrawListSharedData(self.handle()))
}

func (self ImGuiContext) SetTime(v float64) {
	C.ImGuiContext_SetTime(self.handle(), C.double(v))
}

func (self ImGuiContext) GetTime() float64 {
	return float64(C.ImGuiContext_GetTime(self.handle()))
}

func (self ImGuiContext) SetFrameCount(v int32) {
	C.ImGuiContext_SetFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetFrameCount() int {
	return int(C.ImGuiContext_GetFrameCount(self.handle()))
}

func (self ImGuiContext) SetFrameCountEnded(v int32) {
	C.ImGuiContext_SetFrameCountEnded(self.handle(), C.int(v))
}

func (self ImGuiContext) GetFrameCountEnded() int {
	return int(C.ImGuiContext_GetFrameCountEnded(self.handle()))
}

func (self ImGuiContext) SetFrameCountPlatformEnded(v int32) {
	C.ImGuiContext_SetFrameCountPlatformEnded(self.handle(), C.int(v))
}

func (self ImGuiContext) GetFrameCountPlatformEnded() int {
	return int(C.ImGuiContext_GetFrameCountPlatformEnded(self.handle()))
}

func (self ImGuiContext) SetFrameCountRendered(v int32) {
	C.ImGuiContext_SetFrameCountRendered(self.handle(), C.int(v))
}

func (self ImGuiContext) GetFrameCountRendered() int {
	return int(C.ImGuiContext_GetFrameCountRendered(self.handle()))
}

func (self ImGuiContext) SetWithinFrameScope(v bool) {
	C.ImGuiContext_SetWithinFrameScope(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetWithinFrameScope() bool {
	return C.ImGuiContext_GetWithinFrameScope(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetWithinFrameScopeWithImplicitWindow(v bool) {
	C.ImGuiContext_SetWithinFrameScopeWithImplicitWindow(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetWithinFrameScopeWithImplicitWindow() bool {
	return C.ImGuiContext_GetWithinFrameScopeWithImplicitWindow(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetWithinEndChild(v bool) {
	C.ImGuiContext_SetWithinEndChild(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetWithinEndChild() bool {
	return C.ImGuiContext_GetWithinEndChild(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetGcCompactAll(v bool) {
	C.ImGuiContext_SetGcCompactAll(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetGcCompactAll() bool {
	return C.ImGuiContext_GetGcCompactAll(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetTestEngineHookItems(v bool) {
	C.ImGuiContext_SetTestEngineHookItems(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetTestEngineHookItems() bool {
	return C.ImGuiContext_GetTestEngineHookItems(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetTestEngine(v unsafe.Pointer) {
	C.ImGuiContext_SetTestEngine(self.handle(), v)
}

func (self ImGuiContext) GetTestEngine() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiContext_GetTestEngine(self.handle()))
}

func (self ImGuiContext) GetWindowsById() ImGuiStorage {
	return newImGuiStorageFromC(C.ImGuiContext_GetWindowsById(self.handle()))
}

func (self ImGuiContext) SetWindowsActiveCount(v int32) {
	C.ImGuiContext_SetWindowsActiveCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetWindowsActiveCount() int {
	return int(C.ImGuiContext_GetWindowsActiveCount(self.handle()))
}

func (self ImGuiContext) SetWindowsHoverPadding(v ImVec2) {
	C.ImGuiContext_SetWindowsHoverPadding(self.handle(), v.toC())
}

func (self ImGuiContext) GetWindowsHoverPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiContext_GetWindowsHoverPadding(self.handle()))
}

func (self ImGuiContext) SetCurrentWindow(v ImGuiWindow) {
	C.ImGuiContext_SetCurrentWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetCurrentWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetCurrentWindow(self.handle())))
}

func (self ImGuiContext) SetHoveredWindow(v ImGuiWindow) {
	C.ImGuiContext_SetHoveredWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetHoveredWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetHoveredWindow(self.handle())))
}

func (self ImGuiContext) SetHoveredWindowUnderMovingWindow(v ImGuiWindow) {
	C.ImGuiContext_SetHoveredWindowUnderMovingWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetHoveredWindowUnderMovingWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetHoveredWindowUnderMovingWindow(self.handle())))
}

func (self ImGuiContext) SetMovingWindow(v ImGuiWindow) {
	C.ImGuiContext_SetMovingWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetMovingWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetMovingWindow(self.handle())))
}

func (self ImGuiContext) SetWheelingWindow(v ImGuiWindow) {
	C.ImGuiContext_SetWheelingWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetWheelingWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetWheelingWindow(self.handle())))
}

func (self ImGuiContext) SetWheelingWindowRefMousePos(v ImVec2) {
	C.ImGuiContext_SetWheelingWindowRefMousePos(self.handle(), v.toC())
}

func (self ImGuiContext) GetWheelingWindowRefMousePos() ImVec2 {
	return newImVec2FromC(C.ImGuiContext_GetWheelingWindowRefMousePos(self.handle()))
}

func (self ImGuiContext) SetWheelingWindowTimer(v float32) {
	C.ImGuiContext_SetWheelingWindowTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetWheelingWindowTimer() float32 {
	return float32(C.ImGuiContext_GetWheelingWindowTimer(self.handle()))
}

func (self ImGuiContext) SetDebugHookIdInfo(v ImGuiID) {
	C.ImGuiContext_SetDebugHookIdInfo(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetDebugHookIdInfo() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetDebugHookIdInfo(self.handle()))
}

func (self ImGuiContext) SetHoveredId(v ImGuiID) {
	C.ImGuiContext_SetHoveredId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetHoveredId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetHoveredId(self.handle()))
}

func (self ImGuiContext) SetHoveredIdPreviousFrame(v ImGuiID) {
	C.ImGuiContext_SetHoveredIdPreviousFrame(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetHoveredIdPreviousFrame() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetHoveredIdPreviousFrame(self.handle()))
}

func (self ImGuiContext) SetHoveredIdAllowOverlap(v bool) {
	C.ImGuiContext_SetHoveredIdAllowOverlap(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetHoveredIdAllowOverlap() bool {
	return C.ImGuiContext_GetHoveredIdAllowOverlap(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetHoveredIdUsingMouseWheel(v bool) {
	C.ImGuiContext_SetHoveredIdUsingMouseWheel(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetHoveredIdUsingMouseWheel() bool {
	return C.ImGuiContext_GetHoveredIdUsingMouseWheel(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetHoveredIdPreviousFrameUsingMouseWheel(v bool) {
	C.ImGuiContext_SetHoveredIdPreviousFrameUsingMouseWheel(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetHoveredIdPreviousFrameUsingMouseWheel() bool {
	return C.ImGuiContext_GetHoveredIdPreviousFrameUsingMouseWheel(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetHoveredIdDisabled(v bool) {
	C.ImGuiContext_SetHoveredIdDisabled(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetHoveredIdDisabled() bool {
	return C.ImGuiContext_GetHoveredIdDisabled(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetHoveredIdTimer(v float32) {
	C.ImGuiContext_SetHoveredIdTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetHoveredIdTimer() float32 {
	return float32(C.ImGuiContext_GetHoveredIdTimer(self.handle()))
}

func (self ImGuiContext) SetHoveredIdNotActiveTimer(v float32) {
	C.ImGuiContext_SetHoveredIdNotActiveTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetHoveredIdNotActiveTimer() float32 {
	return float32(C.ImGuiContext_GetHoveredIdNotActiveTimer(self.handle()))
}

func (self ImGuiContext) SetActiveId(v ImGuiID) {
	C.ImGuiContext_SetActiveId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetActiveId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetActiveId(self.handle()))
}

func (self ImGuiContext) SetActiveIdIsAlive(v ImGuiID) {
	C.ImGuiContext_SetActiveIdIsAlive(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetActiveIdIsAlive() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetActiveIdIsAlive(self.handle()))
}

func (self ImGuiContext) SetActiveIdTimer(v float32) {
	C.ImGuiContext_SetActiveIdTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetActiveIdTimer() float32 {
	return float32(C.ImGuiContext_GetActiveIdTimer(self.handle()))
}

func (self ImGuiContext) SetActiveIdIsJustActivated(v bool) {
	C.ImGuiContext_SetActiveIdIsJustActivated(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdIsJustActivated() bool {
	return C.ImGuiContext_GetActiveIdIsJustActivated(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdAllowOverlap(v bool) {
	C.ImGuiContext_SetActiveIdAllowOverlap(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdAllowOverlap() bool {
	return C.ImGuiContext_GetActiveIdAllowOverlap(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdNoClearOnFocusLoss(v bool) {
	C.ImGuiContext_SetActiveIdNoClearOnFocusLoss(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdNoClearOnFocusLoss() bool {
	return C.ImGuiContext_GetActiveIdNoClearOnFocusLoss(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdHasBeenPressedBefore(v bool) {
	C.ImGuiContext_SetActiveIdHasBeenPressedBefore(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdHasBeenPressedBefore() bool {
	return C.ImGuiContext_GetActiveIdHasBeenPressedBefore(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdHasBeenEditedBefore(v bool) {
	C.ImGuiContext_SetActiveIdHasBeenEditedBefore(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdHasBeenEditedBefore() bool {
	return C.ImGuiContext_GetActiveIdHasBeenEditedBefore(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdHasBeenEditedThisFrame(v bool) {
	C.ImGuiContext_SetActiveIdHasBeenEditedThisFrame(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdHasBeenEditedThisFrame() bool {
	return C.ImGuiContext_GetActiveIdHasBeenEditedThisFrame(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdClickOffset(v ImVec2) {
	C.ImGuiContext_SetActiveIdClickOffset(self.handle(), v.toC())
}

func (self ImGuiContext) GetActiveIdClickOffset() ImVec2 {
	return newImVec2FromC(C.ImGuiContext_GetActiveIdClickOffset(self.handle()))
}

func (self ImGuiContext) SetActiveIdWindow(v ImGuiWindow) {
	C.ImGuiContext_SetActiveIdWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetActiveIdWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetActiveIdWindow(self.handle())))
}

func (self ImGuiContext) SetActiveIdSource(v ImGuiInputSource) {
	C.ImGuiContext_SetActiveIdSource(self.handle(), C.ImGuiInputSource(v))
}

func (self ImGuiContext) GetActiveIdSource() ImGuiInputSource {
	return ImGuiInputSource(C.ImGuiContext_GetActiveIdSource(self.handle()))
}

func (self ImGuiContext) SetActiveIdMouseButton(v int32) {
	C.ImGuiContext_SetActiveIdMouseButton(self.handle(), C.int(v))
}

func (self ImGuiContext) GetActiveIdMouseButton() int {
	return int(C.ImGuiContext_GetActiveIdMouseButton(self.handle()))
}

func (self ImGuiContext) SetActiveIdPreviousFrame(v ImGuiID) {
	C.ImGuiContext_SetActiveIdPreviousFrame(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetActiveIdPreviousFrame() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetActiveIdPreviousFrame(self.handle()))
}

func (self ImGuiContext) SetActiveIdPreviousFrameIsAlive(v bool) {
	C.ImGuiContext_SetActiveIdPreviousFrameIsAlive(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdPreviousFrameIsAlive() bool {
	return C.ImGuiContext_GetActiveIdPreviousFrameIsAlive(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdPreviousFrameHasBeenEditedBefore(v bool) {
	C.ImGuiContext_SetActiveIdPreviousFrameHasBeenEditedBefore(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetActiveIdPreviousFrameHasBeenEditedBefore() bool {
	return C.ImGuiContext_GetActiveIdPreviousFrameHasBeenEditedBefore(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetActiveIdPreviousFrameWindow(v ImGuiWindow) {
	C.ImGuiContext_SetActiveIdPreviousFrameWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetActiveIdPreviousFrameWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetActiveIdPreviousFrameWindow(self.handle())))
}

func (self ImGuiContext) SetLastActiveId(v ImGuiID) {
	C.ImGuiContext_SetLastActiveId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetLastActiveId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetLastActiveId(self.handle()))
}

func (self ImGuiContext) SetLastActiveIdTimer(v float32) {
	C.ImGuiContext_SetLastActiveIdTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetLastActiveIdTimer() float32 {
	return float32(C.ImGuiContext_GetLastActiveIdTimer(self.handle()))
}

func (self ImGuiContext) SetActiveIdUsingNavDirMask(v uint32) {
	C.ImGuiContext_SetActiveIdUsingNavDirMask(self.handle(), C.ImU32(v))
}

func (self ImGuiContext) GetActiveIdUsingNavDirMask() uint32 {
	return uint32(C.ImGuiContext_GetActiveIdUsingNavDirMask(self.handle()))
}

func (self ImGuiContext) SetActiveIdUsingNavInputMask(v uint32) {
	C.ImGuiContext_SetActiveIdUsingNavInputMask(self.handle(), C.ImU32(v))
}

func (self ImGuiContext) GetActiveIdUsingNavInputMask() uint32 {
	return uint32(C.ImGuiContext_GetActiveIdUsingNavInputMask(self.handle()))
}

func (self ImGuiContext) SetCurrentItemFlags(v ImGuiItemFlags) {
	C.ImGuiContext_SetCurrentItemFlags(self.handle(), C.ImGuiItemFlags(v))
}

func (self ImGuiContext) GetCurrentItemFlags() ImGuiItemFlags {
	return ImGuiItemFlags(C.ImGuiContext_GetCurrentItemFlags(self.handle()))
}

func (self ImGuiContext) GetNextItemData() ImGuiNextItemData {
	return newImGuiNextItemDataFromC(C.ImGuiContext_GetNextItemData(self.handle()))
}

func (self ImGuiContext) GetLastItemData() ImGuiLastItemData {
	return newImGuiLastItemDataFromC(C.ImGuiContext_GetLastItemData(self.handle()))
}

func (self ImGuiContext) GetNextWindowData() ImGuiNextWindowData {
	return newImGuiNextWindowDataFromC(C.ImGuiContext_GetNextWindowData(self.handle()))
}

func (self ImGuiContext) SetBeginMenuCount(v int32) {
	C.ImGuiContext_SetBeginMenuCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetBeginMenuCount() int {
	return int(C.ImGuiContext_GetBeginMenuCount(self.handle()))
}

func (self ImGuiContext) SetCurrentDpiScale(v float32) {
	C.ImGuiContext_SetCurrentDpiScale(self.handle(), C.float(v))
}

func (self ImGuiContext) GetCurrentDpiScale() float32 {
	return float32(C.ImGuiContext_GetCurrentDpiScale(self.handle()))
}

func (self ImGuiContext) SetCurrentViewport(v ImGuiViewportP) {
	C.ImGuiContext_SetCurrentViewport(self.handle(), v.handle())
}

func (self ImGuiContext) GetCurrentViewport() ImGuiViewportP {
	return (ImGuiViewportP)(unsafe.Pointer(C.ImGuiContext_GetCurrentViewport(self.handle())))
}

func (self ImGuiContext) SetMouseViewport(v ImGuiViewportP) {
	C.ImGuiContext_SetMouseViewport(self.handle(), v.handle())
}

func (self ImGuiContext) GetMouseViewport() ImGuiViewportP {
	return (ImGuiViewportP)(unsafe.Pointer(C.ImGuiContext_GetMouseViewport(self.handle())))
}

func (self ImGuiContext) SetMouseLastHoveredViewport(v ImGuiViewportP) {
	C.ImGuiContext_SetMouseLastHoveredViewport(self.handle(), v.handle())
}

func (self ImGuiContext) GetMouseLastHoveredViewport() ImGuiViewportP {
	return (ImGuiViewportP)(unsafe.Pointer(C.ImGuiContext_GetMouseLastHoveredViewport(self.handle())))
}

func (self ImGuiContext) SetPlatformLastFocusedViewportId(v ImGuiID) {
	C.ImGuiContext_SetPlatformLastFocusedViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetPlatformLastFocusedViewportId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetPlatformLastFocusedViewportId(self.handle()))
}

func (self ImGuiContext) GetFallbackMonitor() ImGuiPlatformMonitor {
	return newImGuiPlatformMonitorFromC(C.ImGuiContext_GetFallbackMonitor(self.handle()))
}

func (self ImGuiContext) SetViewportFrontMostStampCount(v int32) {
	C.ImGuiContext_SetViewportFrontMostStampCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetViewportFrontMostStampCount() int {
	return int(C.ImGuiContext_GetViewportFrontMostStampCount(self.handle()))
}

func (self ImGuiContext) SetNavWindow(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetNavWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetNavWindow(self.handle())))
}

func (self ImGuiContext) SetNavId(v ImGuiID) {
	C.ImGuiContext_SetNavId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavId(self.handle()))
}

func (self ImGuiContext) SetNavFocusScopeId(v ImGuiID) {
	C.ImGuiContext_SetNavFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavFocusScopeId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavFocusScopeId(self.handle()))
}

func (self ImGuiContext) SetNavActivateId(v ImGuiID) {
	C.ImGuiContext_SetNavActivateId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavActivateId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavActivateId(self.handle()))
}

func (self ImGuiContext) SetNavActivateDownId(v ImGuiID) {
	C.ImGuiContext_SetNavActivateDownId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavActivateDownId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavActivateDownId(self.handle()))
}

func (self ImGuiContext) SetNavActivatePressedId(v ImGuiID) {
	C.ImGuiContext_SetNavActivatePressedId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavActivatePressedId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavActivatePressedId(self.handle()))
}

func (self ImGuiContext) SetNavActivateInputId(v ImGuiID) {
	C.ImGuiContext_SetNavActivateInputId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavActivateInputId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavActivateInputId(self.handle()))
}

func (self ImGuiContext) SetNavActivateFlags(v ImGuiActivateFlags) {
	C.ImGuiContext_SetNavActivateFlags(self.handle(), C.ImGuiActivateFlags(v))
}

func (self ImGuiContext) GetNavActivateFlags() ImGuiActivateFlags {
	return ImGuiActivateFlags(C.ImGuiContext_GetNavActivateFlags(self.handle()))
}

func (self ImGuiContext) SetNavJustMovedToId(v ImGuiID) {
	C.ImGuiContext_SetNavJustMovedToId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavJustMovedToId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavJustMovedToId(self.handle()))
}

func (self ImGuiContext) SetNavJustMovedToFocusScopeId(v ImGuiID) {
	C.ImGuiContext_SetNavJustMovedToFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavJustMovedToFocusScopeId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavJustMovedToFocusScopeId(self.handle()))
}

func (self ImGuiContext) SetNavJustMovedToKeyMods(v ImGuiModFlags) {
	C.ImGuiContext_SetNavJustMovedToKeyMods(self.handle(), C.ImGuiModFlags(v))
}

func (self ImGuiContext) GetNavJustMovedToKeyMods() ImGuiModFlags {
	return ImGuiModFlags(C.ImGuiContext_GetNavJustMovedToKeyMods(self.handle()))
}

func (self ImGuiContext) SetNavNextActivateId(v ImGuiID) {
	C.ImGuiContext_SetNavNextActivateId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavNextActivateId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavNextActivateId(self.handle()))
}

func (self ImGuiContext) SetNavNextActivateFlags(v ImGuiActivateFlags) {
	C.ImGuiContext_SetNavNextActivateFlags(self.handle(), C.ImGuiActivateFlags(v))
}

func (self ImGuiContext) GetNavNextActivateFlags() ImGuiActivateFlags {
	return ImGuiActivateFlags(C.ImGuiContext_GetNavNextActivateFlags(self.handle()))
}

func (self ImGuiContext) SetNavInputSource(v ImGuiInputSource) {
	C.ImGuiContext_SetNavInputSource(self.handle(), C.ImGuiInputSource(v))
}

func (self ImGuiContext) GetNavInputSource() ImGuiInputSource {
	return ImGuiInputSource(C.ImGuiContext_GetNavInputSource(self.handle()))
}

func (self ImGuiContext) SetNavLayer(v ImGuiNavLayer) {
	C.ImGuiContext_SetNavLayer(self.handle(), C.ImGuiNavLayer(v))
}

func (self ImGuiContext) GetNavLayer() ImGuiNavLayer {
	return ImGuiNavLayer(C.ImGuiContext_GetNavLayer(self.handle()))
}

func (self ImGuiContext) SetNavIdIsAlive(v bool) {
	C.ImGuiContext_SetNavIdIsAlive(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavIdIsAlive() bool {
	return C.ImGuiContext_GetNavIdIsAlive(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavMousePosDirty(v bool) {
	C.ImGuiContext_SetNavMousePosDirty(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavMousePosDirty() bool {
	return C.ImGuiContext_GetNavMousePosDirty(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavDisableHighlight(v bool) {
	C.ImGuiContext_SetNavDisableHighlight(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavDisableHighlight() bool {
	return C.ImGuiContext_GetNavDisableHighlight(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavDisableMouseHover(v bool) {
	C.ImGuiContext_SetNavDisableMouseHover(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavDisableMouseHover() bool {
	return C.ImGuiContext_GetNavDisableMouseHover(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavAnyRequest(v bool) {
	C.ImGuiContext_SetNavAnyRequest(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavAnyRequest() bool {
	return C.ImGuiContext_GetNavAnyRequest(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavInitRequest(v bool) {
	C.ImGuiContext_SetNavInitRequest(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavInitRequest() bool {
	return C.ImGuiContext_GetNavInitRequest(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavInitRequestFromMove(v bool) {
	C.ImGuiContext_SetNavInitRequestFromMove(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavInitRequestFromMove() bool {
	return C.ImGuiContext_GetNavInitRequestFromMove(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavInitResultId(v ImGuiID) {
	C.ImGuiContext_SetNavInitResultId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetNavInitResultId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetNavInitResultId(self.handle()))
}

func (self ImGuiContext) SetNavInitResultRectRel(v ImRect) {
	C.ImGuiContext_SetNavInitResultRectRel(self.handle(), v.toC())
}

func (self ImGuiContext) GetNavInitResultRectRel() ImRect {
	return newImRectFromC(C.ImGuiContext_GetNavInitResultRectRel(self.handle()))
}

func (self ImGuiContext) SetNavMoveSubmitted(v bool) {
	C.ImGuiContext_SetNavMoveSubmitted(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavMoveSubmitted() bool {
	return C.ImGuiContext_GetNavMoveSubmitted(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavMoveScoringItems(v bool) {
	C.ImGuiContext_SetNavMoveScoringItems(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavMoveScoringItems() bool {
	return C.ImGuiContext_GetNavMoveScoringItems(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavMoveForwardToNextFrame(v bool) {
	C.ImGuiContext_SetNavMoveForwardToNextFrame(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavMoveForwardToNextFrame() bool {
	return C.ImGuiContext_GetNavMoveForwardToNextFrame(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavMoveFlags(v ImGuiNavMoveFlags) {
	C.ImGuiContext_SetNavMoveFlags(self.handle(), C.ImGuiNavMoveFlags(v))
}

func (self ImGuiContext) GetNavMoveFlags() ImGuiNavMoveFlags {
	return ImGuiNavMoveFlags(C.ImGuiContext_GetNavMoveFlags(self.handle()))
}

func (self ImGuiContext) SetNavMoveScrollFlags(v ImGuiScrollFlags) {
	C.ImGuiContext_SetNavMoveScrollFlags(self.handle(), C.ImGuiScrollFlags(v))
}

func (self ImGuiContext) GetNavMoveScrollFlags() ImGuiScrollFlags {
	return ImGuiScrollFlags(C.ImGuiContext_GetNavMoveScrollFlags(self.handle()))
}

func (self ImGuiContext) SetNavMoveKeyMods(v ImGuiModFlags) {
	C.ImGuiContext_SetNavMoveKeyMods(self.handle(), C.ImGuiModFlags(v))
}

func (self ImGuiContext) GetNavMoveKeyMods() ImGuiModFlags {
	return ImGuiModFlags(C.ImGuiContext_GetNavMoveKeyMods(self.handle()))
}

func (self ImGuiContext) SetNavMoveDir(v ImGuiDir) {
	C.ImGuiContext_SetNavMoveDir(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiContext) GetNavMoveDir() ImGuiDir {
	return ImGuiDir(C.ImGuiContext_GetNavMoveDir(self.handle()))
}

func (self ImGuiContext) SetNavMoveDirForDebug(v ImGuiDir) {
	C.ImGuiContext_SetNavMoveDirForDebug(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiContext) GetNavMoveDirForDebug() ImGuiDir {
	return ImGuiDir(C.ImGuiContext_GetNavMoveDirForDebug(self.handle()))
}

func (self ImGuiContext) SetNavMoveClipDir(v ImGuiDir) {
	C.ImGuiContext_SetNavMoveClipDir(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiContext) GetNavMoveClipDir() ImGuiDir {
	return ImGuiDir(C.ImGuiContext_GetNavMoveClipDir(self.handle()))
}

func (self ImGuiContext) SetNavScoringRect(v ImRect) {
	C.ImGuiContext_SetNavScoringRect(self.handle(), v.toC())
}

func (self ImGuiContext) GetNavScoringRect() ImRect {
	return newImRectFromC(C.ImGuiContext_GetNavScoringRect(self.handle()))
}

func (self ImGuiContext) SetNavScoringNoClipRect(v ImRect) {
	C.ImGuiContext_SetNavScoringNoClipRect(self.handle(), v.toC())
}

func (self ImGuiContext) GetNavScoringNoClipRect() ImRect {
	return newImRectFromC(C.ImGuiContext_GetNavScoringNoClipRect(self.handle()))
}

func (self ImGuiContext) SetNavScoringDebugCount(v int32) {
	C.ImGuiContext_SetNavScoringDebugCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetNavScoringDebugCount() int {
	return int(C.ImGuiContext_GetNavScoringDebugCount(self.handle()))
}

func (self ImGuiContext) SetNavTabbingDir(v int32) {
	C.ImGuiContext_SetNavTabbingDir(self.handle(), C.int(v))
}

func (self ImGuiContext) GetNavTabbingDir() int {
	return int(C.ImGuiContext_GetNavTabbingDir(self.handle()))
}

func (self ImGuiContext) SetNavTabbingCounter(v int32) {
	C.ImGuiContext_SetNavTabbingCounter(self.handle(), C.int(v))
}

func (self ImGuiContext) GetNavTabbingCounter() int {
	return int(C.ImGuiContext_GetNavTabbingCounter(self.handle()))
}

func (self ImGuiContext) GetNavMoveResultLocal() ImGuiNavItemData {
	return newImGuiNavItemDataFromC(C.ImGuiContext_GetNavMoveResultLocal(self.handle()))
}

func (self ImGuiContext) GetNavMoveResultLocalVisible() ImGuiNavItemData {
	return newImGuiNavItemDataFromC(C.ImGuiContext_GetNavMoveResultLocalVisible(self.handle()))
}

func (self ImGuiContext) GetNavMoveResultOther() ImGuiNavItemData {
	return newImGuiNavItemDataFromC(C.ImGuiContext_GetNavMoveResultOther(self.handle()))
}

func (self ImGuiContext) GetNavTabbingResultFirst() ImGuiNavItemData {
	return newImGuiNavItemDataFromC(C.ImGuiContext_GetNavTabbingResultFirst(self.handle()))
}

func (self ImGuiContext) SetNavWindowingTarget(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindowingTarget(self.handle(), v.handle())
}

func (self ImGuiContext) GetNavWindowingTarget() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetNavWindowingTarget(self.handle())))
}

func (self ImGuiContext) SetNavWindowingTargetAnim(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindowingTargetAnim(self.handle(), v.handle())
}

func (self ImGuiContext) GetNavWindowingTargetAnim() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetNavWindowingTargetAnim(self.handle())))
}

func (self ImGuiContext) SetNavWindowingListWindow(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindowingListWindow(self.handle(), v.handle())
}

func (self ImGuiContext) GetNavWindowingListWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiContext_GetNavWindowingListWindow(self.handle())))
}

func (self ImGuiContext) SetNavWindowingTimer(v float32) {
	C.ImGuiContext_SetNavWindowingTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetNavWindowingTimer() float32 {
	return float32(C.ImGuiContext_GetNavWindowingTimer(self.handle()))
}

func (self ImGuiContext) SetNavWindowingHighlightAlpha(v float32) {
	C.ImGuiContext_SetNavWindowingHighlightAlpha(self.handle(), C.float(v))
}

func (self ImGuiContext) GetNavWindowingHighlightAlpha() float32 {
	return float32(C.ImGuiContext_GetNavWindowingHighlightAlpha(self.handle()))
}

func (self ImGuiContext) SetNavWindowingToggleLayer(v bool) {
	C.ImGuiContext_SetNavWindowingToggleLayer(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetNavWindowingToggleLayer() bool {
	return C.ImGuiContext_GetNavWindowingToggleLayer(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetNavWindowingAccumDeltaPos(v ImVec2) {
	C.ImGuiContext_SetNavWindowingAccumDeltaPos(self.handle(), v.toC())
}

func (self ImGuiContext) GetNavWindowingAccumDeltaPos() ImVec2 {
	return newImVec2FromC(C.ImGuiContext_GetNavWindowingAccumDeltaPos(self.handle()))
}

func (self ImGuiContext) SetNavWindowingAccumDeltaSize(v ImVec2) {
	C.ImGuiContext_SetNavWindowingAccumDeltaSize(self.handle(), v.toC())
}

func (self ImGuiContext) GetNavWindowingAccumDeltaSize() ImVec2 {
	return newImVec2FromC(C.ImGuiContext_GetNavWindowingAccumDeltaSize(self.handle()))
}

func (self ImGuiContext) SetDimBgRatio(v float32) {
	C.ImGuiContext_SetDimBgRatio(self.handle(), C.float(v))
}

func (self ImGuiContext) GetDimBgRatio() float32 {
	return float32(C.ImGuiContext_GetDimBgRatio(self.handle()))
}

func (self ImGuiContext) SetMouseCursor(v ImGuiMouseCursor) {
	C.ImGuiContext_SetMouseCursor(self.handle(), C.ImGuiMouseCursor(v))
}

func (self ImGuiContext) GetMouseCursor() ImGuiMouseCursor {
	return ImGuiMouseCursor(C.ImGuiContext_GetMouseCursor(self.handle()))
}

func (self ImGuiContext) SetDragDropActive(v bool) {
	C.ImGuiContext_SetDragDropActive(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetDragDropActive() bool {
	return C.ImGuiContext_GetDragDropActive(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetDragDropWithinSource(v bool) {
	C.ImGuiContext_SetDragDropWithinSource(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetDragDropWithinSource() bool {
	return C.ImGuiContext_GetDragDropWithinSource(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetDragDropWithinTarget(v bool) {
	C.ImGuiContext_SetDragDropWithinTarget(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetDragDropWithinTarget() bool {
	return C.ImGuiContext_GetDragDropWithinTarget(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetDragDropSourceFlags(v ImGuiDragDropFlags) {
	C.ImGuiContext_SetDragDropSourceFlags(self.handle(), C.ImGuiDragDropFlags(v))
}

func (self ImGuiContext) GetDragDropSourceFlags() ImGuiDragDropFlags {
	return ImGuiDragDropFlags(C.ImGuiContext_GetDragDropSourceFlags(self.handle()))
}

func (self ImGuiContext) SetDragDropSourceFrameCount(v int32) {
	C.ImGuiContext_SetDragDropSourceFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetDragDropSourceFrameCount() int {
	return int(C.ImGuiContext_GetDragDropSourceFrameCount(self.handle()))
}

func (self ImGuiContext) SetDragDropMouseButton(v int32) {
	C.ImGuiContext_SetDragDropMouseButton(self.handle(), C.int(v))
}

func (self ImGuiContext) GetDragDropMouseButton() int {
	return int(C.ImGuiContext_GetDragDropMouseButton(self.handle()))
}

func (self ImGuiContext) GetDragDropPayload() ImGuiPayload {
	return newImGuiPayloadFromC(C.ImGuiContext_GetDragDropPayload(self.handle()))
}

func (self ImGuiContext) SetDragDropTargetRect(v ImRect) {
	C.ImGuiContext_SetDragDropTargetRect(self.handle(), v.toC())
}

func (self ImGuiContext) GetDragDropTargetRect() ImRect {
	return newImRectFromC(C.ImGuiContext_GetDragDropTargetRect(self.handle()))
}

func (self ImGuiContext) SetDragDropTargetId(v ImGuiID) {
	C.ImGuiContext_SetDragDropTargetId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetDragDropTargetId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetDragDropTargetId(self.handle()))
}

func (self ImGuiContext) SetDragDropAcceptFlags(v ImGuiDragDropFlags) {
	C.ImGuiContext_SetDragDropAcceptFlags(self.handle(), C.ImGuiDragDropFlags(v))
}

func (self ImGuiContext) GetDragDropAcceptFlags() ImGuiDragDropFlags {
	return ImGuiDragDropFlags(C.ImGuiContext_GetDragDropAcceptFlags(self.handle()))
}

func (self ImGuiContext) SetDragDropAcceptIdCurrRectSurface(v float32) {
	C.ImGuiContext_SetDragDropAcceptIdCurrRectSurface(self.handle(), C.float(v))
}

func (self ImGuiContext) GetDragDropAcceptIdCurrRectSurface() float32 {
	return float32(C.ImGuiContext_GetDragDropAcceptIdCurrRectSurface(self.handle()))
}

func (self ImGuiContext) SetDragDropAcceptIdCurr(v ImGuiID) {
	C.ImGuiContext_SetDragDropAcceptIdCurr(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetDragDropAcceptIdCurr() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetDragDropAcceptIdCurr(self.handle()))
}

func (self ImGuiContext) SetDragDropAcceptIdPrev(v ImGuiID) {
	C.ImGuiContext_SetDragDropAcceptIdPrev(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetDragDropAcceptIdPrev() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetDragDropAcceptIdPrev(self.handle()))
}

func (self ImGuiContext) SetDragDropAcceptFrameCount(v int32) {
	C.ImGuiContext_SetDragDropAcceptFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetDragDropAcceptFrameCount() int {
	return int(C.ImGuiContext_GetDragDropAcceptFrameCount(self.handle()))
}

func (self ImGuiContext) SetDragDropHoldJustPressedId(v ImGuiID) {
	C.ImGuiContext_SetDragDropHoldJustPressedId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetDragDropHoldJustPressedId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetDragDropHoldJustPressedId(self.handle()))
}

func (self ImGuiContext) SetClipperTempDataStacked(v int32) {
	C.ImGuiContext_SetClipperTempDataStacked(self.handle(), C.int(v))
}

func (self ImGuiContext) GetClipperTempDataStacked() int {
	return int(C.ImGuiContext_GetClipperTempDataStacked(self.handle()))
}

func (self ImGuiContext) SetCurrentTable(v ImGuiTable) {
	C.ImGuiContext_SetCurrentTable(self.handle(), v.handle())
}

func (self ImGuiContext) GetCurrentTable() ImGuiTable {
	return (ImGuiTable)(unsafe.Pointer(C.ImGuiContext_GetCurrentTable(self.handle())))
}

func (self ImGuiContext) SetTablesTempDataStacked(v int32) {
	C.ImGuiContext_SetTablesTempDataStacked(self.handle(), C.int(v))
}

func (self ImGuiContext) GetTablesTempDataStacked() int {
	return int(C.ImGuiContext_GetTablesTempDataStacked(self.handle()))
}

func (self ImGuiContext) SetCurrentTabBar(v ImGuiTabBar) {
	C.ImGuiContext_SetCurrentTabBar(self.handle(), v.handle())
}

func (self ImGuiContext) GetCurrentTabBar() ImGuiTabBar {
	return (ImGuiTabBar)(unsafe.Pointer(C.ImGuiContext_GetCurrentTabBar(self.handle())))
}

func (self ImGuiContext) SetMouseLastValidPos(v ImVec2) {
	C.ImGuiContext_SetMouseLastValidPos(self.handle(), v.toC())
}

func (self ImGuiContext) GetMouseLastValidPos() ImVec2 {
	return newImVec2FromC(C.ImGuiContext_GetMouseLastValidPos(self.handle()))
}

func (self ImGuiContext) GetInputTextState() ImGuiInputTextState {
	return newImGuiInputTextStateFromC(C.ImGuiContext_GetInputTextState(self.handle()))
}

func (self ImGuiContext) GetInputTextPasswordFont() ImFont {
	return newImFontFromC(C.ImGuiContext_GetInputTextPasswordFont(self.handle()))
}

func (self ImGuiContext) SetTempInputId(v ImGuiID) {
	C.ImGuiContext_SetTempInputId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetTempInputId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetTempInputId(self.handle()))
}

func (self ImGuiContext) SetColorEditOptions(v ImGuiColorEditFlags) {
	C.ImGuiContext_SetColorEditOptions(self.handle(), C.ImGuiColorEditFlags(v))
}

func (self ImGuiContext) GetColorEditOptions() ImGuiColorEditFlags {
	return ImGuiColorEditFlags(C.ImGuiContext_GetColorEditOptions(self.handle()))
}

func (self ImGuiContext) SetColorEditLastHue(v float32) {
	C.ImGuiContext_SetColorEditLastHue(self.handle(), C.float(v))
}

func (self ImGuiContext) GetColorEditLastHue() float32 {
	return float32(C.ImGuiContext_GetColorEditLastHue(self.handle()))
}

func (self ImGuiContext) SetColorEditLastSat(v float32) {
	C.ImGuiContext_SetColorEditLastSat(self.handle(), C.float(v))
}

func (self ImGuiContext) GetColorEditLastSat() float32 {
	return float32(C.ImGuiContext_GetColorEditLastSat(self.handle()))
}

func (self ImGuiContext) SetColorEditLastColor(v uint32) {
	C.ImGuiContext_SetColorEditLastColor(self.handle(), C.ImU32(v))
}

func (self ImGuiContext) GetColorEditLastColor() uint32 {
	return uint32(C.ImGuiContext_GetColorEditLastColor(self.handle()))
}

func (self ImGuiContext) SetColorPickerRef(v ImVec4) {
	C.ImGuiContext_SetColorPickerRef(self.handle(), v.toC())
}

func (self ImGuiContext) GetColorPickerRef() ImVec4 {
	return newImVec4FromC(C.ImGuiContext_GetColorPickerRef(self.handle()))
}

func (self ImGuiContext) GetComboPreviewData() ImGuiComboPreviewData {
	return newImGuiComboPreviewDataFromC(C.ImGuiContext_GetComboPreviewData(self.handle()))
}

func (self ImGuiContext) SetSliderGrabClickOffset(v float32) {
	C.ImGuiContext_SetSliderGrabClickOffset(self.handle(), C.float(v))
}

func (self ImGuiContext) GetSliderGrabClickOffset() float32 {
	return float32(C.ImGuiContext_GetSliderGrabClickOffset(self.handle()))
}

func (self ImGuiContext) SetSliderCurrentAccum(v float32) {
	C.ImGuiContext_SetSliderCurrentAccum(self.handle(), C.float(v))
}

func (self ImGuiContext) GetSliderCurrentAccum() float32 {
	return float32(C.ImGuiContext_GetSliderCurrentAccum(self.handle()))
}

func (self ImGuiContext) SetSliderCurrentAccumDirty(v bool) {
	C.ImGuiContext_SetSliderCurrentAccumDirty(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetSliderCurrentAccumDirty() bool {
	return C.ImGuiContext_GetSliderCurrentAccumDirty(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetDragCurrentAccumDirty(v bool) {
	C.ImGuiContext_SetDragCurrentAccumDirty(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetDragCurrentAccumDirty() bool {
	return C.ImGuiContext_GetDragCurrentAccumDirty(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetDragCurrentAccum(v float32) {
	C.ImGuiContext_SetDragCurrentAccum(self.handle(), C.float(v))
}

func (self ImGuiContext) GetDragCurrentAccum() float32 {
	return float32(C.ImGuiContext_GetDragCurrentAccum(self.handle()))
}

func (self ImGuiContext) SetDragSpeedDefaultRatio(v float32) {
	C.ImGuiContext_SetDragSpeedDefaultRatio(self.handle(), C.float(v))
}

func (self ImGuiContext) GetDragSpeedDefaultRatio() float32 {
	return float32(C.ImGuiContext_GetDragSpeedDefaultRatio(self.handle()))
}

func (self ImGuiContext) SetScrollbarClickDeltaToGrabCenter(v float32) {
	C.ImGuiContext_SetScrollbarClickDeltaToGrabCenter(self.handle(), C.float(v))
}

func (self ImGuiContext) GetScrollbarClickDeltaToGrabCenter() float32 {
	return float32(C.ImGuiContext_GetScrollbarClickDeltaToGrabCenter(self.handle()))
}

func (self ImGuiContext) SetDisabledAlphaBackup(v float32) {
	C.ImGuiContext_SetDisabledAlphaBackup(self.handle(), C.float(v))
}

func (self ImGuiContext) GetDisabledAlphaBackup() float32 {
	return float32(C.ImGuiContext_GetDisabledAlphaBackup(self.handle()))
}

func (self ImGuiContext) SetDisabledStackSize(v int) {
	C.ImGuiContext_SetDisabledStackSize(self.handle(), C.short(v))
}

func (self ImGuiContext) GetDisabledStackSize() int {
	return int(C.ImGuiContext_GetDisabledStackSize(self.handle()))
}

func (self ImGuiContext) SetTooltipOverrideCount(v int) {
	C.ImGuiContext_SetTooltipOverrideCount(self.handle(), C.short(v))
}

func (self ImGuiContext) GetTooltipOverrideCount() int {
	return int(C.ImGuiContext_GetTooltipOverrideCount(self.handle()))
}

func (self ImGuiContext) SetTooltipSlowDelay(v float32) {
	C.ImGuiContext_SetTooltipSlowDelay(self.handle(), C.float(v))
}

func (self ImGuiContext) GetTooltipSlowDelay() float32 {
	return float32(C.ImGuiContext_GetTooltipSlowDelay(self.handle()))
}

func (self ImGuiContext) GetPlatformImeData() ImGuiPlatformImeData {
	return newImGuiPlatformImeDataFromC(C.ImGuiContext_GetPlatformImeData(self.handle()))
}

func (self ImGuiContext) GetPlatformImeDataPrev() ImGuiPlatformImeData {
	return newImGuiPlatformImeDataFromC(C.ImGuiContext_GetPlatformImeDataPrev(self.handle()))
}

func (self ImGuiContext) SetPlatformImeViewport(v ImGuiID) {
	C.ImGuiContext_SetPlatformImeViewport(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetPlatformImeViewport() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetPlatformImeViewport(self.handle()))
}

func (self ImGuiContext) GetDockContext() ImGuiDockContext {
	return newImGuiDockContextFromC(C.ImGuiContext_GetDockContext(self.handle()))
}

func (self ImGuiContext) SetSettingsLoaded(v bool) {
	C.ImGuiContext_SetSettingsLoaded(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetSettingsLoaded() bool {
	return C.ImGuiContext_GetSettingsLoaded(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetSettingsDirtyTimer(v float32) {
	C.ImGuiContext_SetSettingsDirtyTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) GetSettingsDirtyTimer() float32 {
	return float32(C.ImGuiContext_GetSettingsDirtyTimer(self.handle()))
}

func (self ImGuiContext) GetSettingsIniData() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImGuiContext_GetSettingsIniData(self.handle()))
}

func (self ImGuiContext) SetHookIdNext(v ImGuiID) {
	C.ImGuiContext_SetHookIdNext(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetHookIdNext() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetHookIdNext(self.handle()))
}

func (self ImGuiContext) SetLogEnabled(v bool) {
	C.ImGuiContext_SetLogEnabled(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetLogEnabled() bool {
	return C.ImGuiContext_GetLogEnabled(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetLogType(v ImGuiLogType) {
	C.ImGuiContext_SetLogType(self.handle(), C.ImGuiLogType(v))
}

func (self ImGuiContext) GetLogType() ImGuiLogType {
	return ImGuiLogType(C.ImGuiContext_GetLogType(self.handle()))
}

func (self ImGuiContext) GetLogBuffer() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImGuiContext_GetLogBuffer(self.handle()))
}

func (self ImGuiContext) SetLogNextPrefix(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiContext_SetLogNextPrefix(self.handle(), vArg)
}

func (self ImGuiContext) GetLogNextPrefix() string {
	return C.GoString(C.ImGuiContext_GetLogNextPrefix(self.handle()))
}

func (self ImGuiContext) SetLogNextSuffix(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiContext_SetLogNextSuffix(self.handle(), vArg)
}

func (self ImGuiContext) GetLogNextSuffix() string {
	return C.GoString(C.ImGuiContext_GetLogNextSuffix(self.handle()))
}

func (self ImGuiContext) SetLogLinePosY(v float32) {
	C.ImGuiContext_SetLogLinePosY(self.handle(), C.float(v))
}

func (self ImGuiContext) GetLogLinePosY() float32 {
	return float32(C.ImGuiContext_GetLogLinePosY(self.handle()))
}

func (self ImGuiContext) SetLogLineFirstItem(v bool) {
	C.ImGuiContext_SetLogLineFirstItem(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetLogLineFirstItem() bool {
	return C.ImGuiContext_GetLogLineFirstItem(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetLogDepthRef(v int32) {
	C.ImGuiContext_SetLogDepthRef(self.handle(), C.int(v))
}

func (self ImGuiContext) GetLogDepthRef() int {
	return int(C.ImGuiContext_GetLogDepthRef(self.handle()))
}

func (self ImGuiContext) SetLogDepthToExpand(v int32) {
	C.ImGuiContext_SetLogDepthToExpand(self.handle(), C.int(v))
}

func (self ImGuiContext) GetLogDepthToExpand() int {
	return int(C.ImGuiContext_GetLogDepthToExpand(self.handle()))
}

func (self ImGuiContext) SetLogDepthToExpandDefault(v int32) {
	C.ImGuiContext_SetLogDepthToExpandDefault(self.handle(), C.int(v))
}

func (self ImGuiContext) GetLogDepthToExpandDefault() int {
	return int(C.ImGuiContext_GetLogDepthToExpandDefault(self.handle()))
}

func (self ImGuiContext) SetDebugLogFlags(v ImGuiDebugLogFlags) {
	C.ImGuiContext_SetDebugLogFlags(self.handle(), C.ImGuiDebugLogFlags(v))
}

func (self ImGuiContext) GetDebugLogFlags() ImGuiDebugLogFlags {
	return ImGuiDebugLogFlags(C.ImGuiContext_GetDebugLogFlags(self.handle()))
}

func (self ImGuiContext) GetDebugLogBuf() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImGuiContext_GetDebugLogBuf(self.handle()))
}

func (self ImGuiContext) SetDebugItemPickerActive(v bool) {
	C.ImGuiContext_SetDebugItemPickerActive(self.handle(), C.bool(v))
}

func (self ImGuiContext) GetDebugItemPickerActive() bool {
	return C.ImGuiContext_GetDebugItemPickerActive(self.handle()) == C.bool(true)
}

func (self ImGuiContext) SetDebugItemPickerMouseButton(v uint) {
	C.ImGuiContext_SetDebugItemPickerMouseButton(self.handle(), C.ImU8(v))
}

func (self ImGuiContext) GetDebugItemPickerMouseButton() uint32 {
	return uint32(C.ImGuiContext_GetDebugItemPickerMouseButton(self.handle()))
}

func (self ImGuiContext) SetDebugItemPickerBreakId(v ImGuiID) {
	C.ImGuiContext_SetDebugItemPickerBreakId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) GetDebugItemPickerBreakId() ImGuiID {
	return ImGuiID(C.ImGuiContext_GetDebugItemPickerBreakId(self.handle()))
}

func (self ImGuiContext) GetDebugMetricsConfig() ImGuiMetricsConfig {
	return newImGuiMetricsConfigFromC(C.ImGuiContext_GetDebugMetricsConfig(self.handle()))
}

func (self ImGuiContext) GetDebugStackTool() ImGuiStackTool {
	return newImGuiStackToolFromC(C.ImGuiContext_GetDebugStackTool(self.handle()))
}

func (self ImGuiContext) SetDebugHoveredDockNode(v ImGuiDockNode) {
	C.ImGuiContext_SetDebugHoveredDockNode(self.handle(), v.handle())
}

func (self ImGuiContext) GetDebugHoveredDockNode() ImGuiDockNode {
	return (ImGuiDockNode)(unsafe.Pointer(C.ImGuiContext_GetDebugHoveredDockNode(self.handle())))
}

func (self ImGuiContext) SetFramerateSecPerFrameIdx(v int32) {
	C.ImGuiContext_SetFramerateSecPerFrameIdx(self.handle(), C.int(v))
}

func (self ImGuiContext) GetFramerateSecPerFrameIdx() int {
	return int(C.ImGuiContext_GetFramerateSecPerFrameIdx(self.handle()))
}

func (self ImGuiContext) SetFramerateSecPerFrameCount(v int32) {
	C.ImGuiContext_SetFramerateSecPerFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) GetFramerateSecPerFrameCount() int {
	return int(C.ImGuiContext_GetFramerateSecPerFrameCount(self.handle()))
}

func (self ImGuiContext) SetFramerateSecPerFrameAccum(v float32) {
	C.ImGuiContext_SetFramerateSecPerFrameAccum(self.handle(), C.float(v))
}

func (self ImGuiContext) GetFramerateSecPerFrameAccum() float32 {
	return float32(C.ImGuiContext_GetFramerateSecPerFrameAccum(self.handle()))
}

func (self ImGuiContext) SetWantCaptureMouseNextFrame(v int32) {
	C.ImGuiContext_SetWantCaptureMouseNextFrame(self.handle(), C.int(v))
}

func (self ImGuiContext) GetWantCaptureMouseNextFrame() int {
	return int(C.ImGuiContext_GetWantCaptureMouseNextFrame(self.handle()))
}

func (self ImGuiContext) SetWantCaptureKeyboardNextFrame(v int32) {
	C.ImGuiContext_SetWantCaptureKeyboardNextFrame(self.handle(), C.int(v))
}

func (self ImGuiContext) GetWantCaptureKeyboardNextFrame() int {
	return int(C.ImGuiContext_GetWantCaptureKeyboardNextFrame(self.handle()))
}

func (self ImGuiContext) SetWantTextInputNextFrame(v int32) {
	C.ImGuiContext_SetWantTextInputNextFrame(self.handle(), C.int(v))
}

func (self ImGuiContext) GetWantTextInputNextFrame() int {
	return int(C.ImGuiContext_GetWantTextInputNextFrame(self.handle()))
}

func (self ImGuiInputTextState) SetID(v ImGuiID) {
	C.ImGuiInputTextState_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiInputTextState) GetID() ImGuiID {
	return ImGuiID(C.ImGuiInputTextState_GetID(self.handle()))
}

func (self ImGuiInputTextState) SetCurLenW(v int32) {
	C.ImGuiInputTextState_SetCurLenW(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) GetCurLenW() int {
	return int(C.ImGuiInputTextState_GetCurLenW(self.handle()))
}

func (self ImGuiInputTextState) SetCurLenA(v int32) {
	C.ImGuiInputTextState_SetCurLenA(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) GetCurLenA() int {
	return int(C.ImGuiInputTextState_GetCurLenA(self.handle()))
}

func (self ImGuiInputTextState) SetTextAIsValid(v bool) {
	C.ImGuiInputTextState_SetTextAIsValid(self.handle(), C.bool(v))
}

func (self ImGuiInputTextState) GetTextAIsValid() bool {
	return C.ImGuiInputTextState_GetTextAIsValid(self.handle()) == C.bool(true)
}

func (self ImGuiInputTextState) SetBufCapacityA(v int32) {
	C.ImGuiInputTextState_SetBufCapacityA(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) GetBufCapacityA() int {
	return int(C.ImGuiInputTextState_GetBufCapacityA(self.handle()))
}

func (self ImGuiInputTextState) SetScrollX(v float32) {
	C.ImGuiInputTextState_SetScrollX(self.handle(), C.float(v))
}

func (self ImGuiInputTextState) GetScrollX() float32 {
	return float32(C.ImGuiInputTextState_GetScrollX(self.handle()))
}

func (self ImGuiInputTextState) SetCursorAnim(v float32) {
	C.ImGuiInputTextState_SetCursorAnim(self.handle(), C.float(v))
}

func (self ImGuiInputTextState) GetCursorAnim() float32 {
	return float32(C.ImGuiInputTextState_GetCursorAnim(self.handle()))
}

func (self ImGuiInputTextState) SetCursorFollow(v bool) {
	C.ImGuiInputTextState_SetCursorFollow(self.handle(), C.bool(v))
}

func (self ImGuiInputTextState) GetCursorFollow() bool {
	return C.ImGuiInputTextState_GetCursorFollow(self.handle()) == C.bool(true)
}

func (self ImGuiInputTextState) SetSelectedAllMouseLock(v bool) {
	C.ImGuiInputTextState_SetSelectedAllMouseLock(self.handle(), C.bool(v))
}

func (self ImGuiInputTextState) GetSelectedAllMouseLock() bool {
	return C.ImGuiInputTextState_GetSelectedAllMouseLock(self.handle()) == C.bool(true)
}

func (self ImGuiInputTextState) SetEdited(v bool) {
	C.ImGuiInputTextState_SetEdited(self.handle(), C.bool(v))
}

func (self ImGuiInputTextState) GetEdited() bool {
	return C.ImGuiInputTextState_GetEdited(self.handle()) == C.bool(true)
}

func (self ImGuiInputTextState) SetFlags(v ImGuiInputTextFlags) {
	C.ImGuiInputTextState_SetFlags(self.handle(), C.ImGuiInputTextFlags(v))
}

func (self ImGuiInputTextState) GetFlags() ImGuiInputTextFlags {
	return ImGuiInputTextFlags(C.ImGuiInputTextState_GetFlags(self.handle()))
}

func (self ImGuiLastItemData) SetID(v ImGuiID) {
	C.ImGuiLastItemData_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiLastItemData) GetID() ImGuiID {
	return ImGuiID(C.ImGuiLastItemData_GetID(self.handle()))
}

func (self ImGuiLastItemData) SetInFlags(v ImGuiItemFlags) {
	C.ImGuiLastItemData_SetInFlags(self.handle(), C.ImGuiItemFlags(v))
}

func (self ImGuiLastItemData) GetInFlags() ImGuiItemFlags {
	return ImGuiItemFlags(C.ImGuiLastItemData_GetInFlags(self.handle()))
}

func (self ImGuiLastItemData) SetStatusFlags(v ImGuiItemStatusFlags) {
	C.ImGuiLastItemData_SetStatusFlags(self.handle(), C.ImGuiItemStatusFlags(v))
}

func (self ImGuiLastItemData) GetStatusFlags() ImGuiItemStatusFlags {
	return ImGuiItemStatusFlags(C.ImGuiLastItemData_GetStatusFlags(self.handle()))
}

func (self ImGuiLastItemData) SetRect(v ImRect) {
	C.ImGuiLastItemData_SetRect(self.handle(), v.toC())
}

func (self ImGuiLastItemData) GetRect() ImRect {
	return newImRectFromC(C.ImGuiLastItemData_GetRect(self.handle()))
}

func (self ImGuiLastItemData) SetNavRect(v ImRect) {
	C.ImGuiLastItemData_SetNavRect(self.handle(), v.toC())
}

func (self ImGuiLastItemData) GetNavRect() ImRect {
	return newImRectFromC(C.ImGuiLastItemData_GetNavRect(self.handle()))
}

func (self ImGuiLastItemData) SetDisplayRect(v ImRect) {
	C.ImGuiLastItemData_SetDisplayRect(self.handle(), v.toC())
}

func (self ImGuiLastItemData) GetDisplayRect() ImRect {
	return newImRectFromC(C.ImGuiLastItemData_GetDisplayRect(self.handle()))
}

func (self ImGuiMetricsConfig) SetShowDebugLog(v bool) {
	C.ImGuiMetricsConfig_SetShowDebugLog(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowDebugLog() bool {
	return C.ImGuiMetricsConfig_GetShowDebugLog(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowStackTool(v bool) {
	C.ImGuiMetricsConfig_SetShowStackTool(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowStackTool() bool {
	return C.ImGuiMetricsConfig_GetShowStackTool(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowWindowsRects(v bool) {
	C.ImGuiMetricsConfig_SetShowWindowsRects(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowWindowsRects() bool {
	return C.ImGuiMetricsConfig_GetShowWindowsRects(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowWindowsBeginOrder(v bool) {
	C.ImGuiMetricsConfig_SetShowWindowsBeginOrder(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowWindowsBeginOrder() bool {
	return C.ImGuiMetricsConfig_GetShowWindowsBeginOrder(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowTablesRects(v bool) {
	C.ImGuiMetricsConfig_SetShowTablesRects(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowTablesRects() bool {
	return C.ImGuiMetricsConfig_GetShowTablesRects(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowDrawCmdMesh(v bool) {
	C.ImGuiMetricsConfig_SetShowDrawCmdMesh(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowDrawCmdMesh() bool {
	return C.ImGuiMetricsConfig_GetShowDrawCmdMesh(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowDrawCmdBoundingBoxes(v bool) {
	C.ImGuiMetricsConfig_SetShowDrawCmdBoundingBoxes(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowDrawCmdBoundingBoxes() bool {
	return C.ImGuiMetricsConfig_GetShowDrawCmdBoundingBoxes(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowDockingNodes(v bool) {
	C.ImGuiMetricsConfig_SetShowDockingNodes(self.handle(), C.bool(v))
}

func (self ImGuiMetricsConfig) GetShowDockingNodes() bool {
	return C.ImGuiMetricsConfig_GetShowDockingNodes(self.handle()) == C.bool(true)
}

func (self ImGuiMetricsConfig) SetShowWindowsRectsType(v int32) {
	C.ImGuiMetricsConfig_SetShowWindowsRectsType(self.handle(), C.int(v))
}

func (self ImGuiMetricsConfig) GetShowWindowsRectsType() int {
	return int(C.ImGuiMetricsConfig_GetShowWindowsRectsType(self.handle()))
}

func (self ImGuiMetricsConfig) SetShowTablesRectsType(v int32) {
	C.ImGuiMetricsConfig_SetShowTablesRectsType(self.handle(), C.int(v))
}

func (self ImGuiMetricsConfig) GetShowTablesRectsType() int {
	return int(C.ImGuiMetricsConfig_GetShowTablesRectsType(self.handle()))
}

func (self ImGuiPtrOrIndex) SetPtr(v unsafe.Pointer) {
	C.ImGuiPtrOrIndex_SetPtr(self.handle(), v)
}

func (self ImGuiPtrOrIndex) GetPtr() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiPtrOrIndex_GetPtr(self.handle()))
}

func (self ImGuiPtrOrIndex) SetIndex(v int32) {
	C.ImGuiPtrOrIndex_SetIndex(self.handle(), C.int(v))
}

func (self ImGuiPtrOrIndex) GetIndex() int {
	return int(C.ImGuiPtrOrIndex_GetIndex(self.handle()))
}

func (self ImGuiWindow) SetName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiWindow_SetName(self.handle(), vArg)
}

func (self ImGuiWindow) SetID(v ImGuiID) {
	C.ImGuiWindow_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetID() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetID(self.handle()))
}

func (self ImGuiWindow) SetFlags(v ImGuiWindowFlags) {
	C.ImGuiWindow_SetFlags(self.handle(), C.ImGuiWindowFlags(v))
}

func (self ImGuiWindow) GetFlags() ImGuiWindowFlags {
	return ImGuiWindowFlags(C.ImGuiWindow_GetFlags(self.handle()))
}

func (self ImGuiWindow) SetFlagsPreviousFrame(v ImGuiWindowFlags) {
	C.ImGuiWindow_SetFlagsPreviousFrame(self.handle(), C.ImGuiWindowFlags(v))
}

func (self ImGuiWindow) GetFlagsPreviousFrame() ImGuiWindowFlags {
	return ImGuiWindowFlags(C.ImGuiWindow_GetFlagsPreviousFrame(self.handle()))
}

func (self ImGuiWindow) GetWindowClass() ImGuiWindowClass {
	return newImGuiWindowClassFromC(C.ImGuiWindow_GetWindowClass(self.handle()))
}

func (self ImGuiWindow) SetViewport(v ImGuiViewportP) {
	C.ImGuiWindow_SetViewport(self.handle(), v.handle())
}

func (self ImGuiWindow) GetViewport() ImGuiViewportP {
	return (ImGuiViewportP)(unsafe.Pointer(C.ImGuiWindow_GetViewport(self.handle())))
}

func (self ImGuiWindow) SetViewportId(v ImGuiID) {
	C.ImGuiWindow_SetViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetViewportId() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetViewportId(self.handle()))
}

func (self ImGuiWindow) SetViewportPos(v ImVec2) {
	C.ImGuiWindow_SetViewportPos(self.handle(), v.toC())
}

func (self ImGuiWindow) GetViewportPos() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetViewportPos(self.handle()))
}

func (self ImGuiWindow) SetViewportAllowPlatformMonitorExtend(v int32) {
	C.ImGuiWindow_SetViewportAllowPlatformMonitorExtend(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetViewportAllowPlatformMonitorExtend() int {
	return int(C.ImGuiWindow_GetViewportAllowPlatformMonitorExtend(self.handle()))
}

func (self ImGuiWindow) SetPos(v ImVec2) {
	C.ImGuiWindow_SetPos(self.handle(), v.toC())
}

func (self ImGuiWindow) GetPos() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetPos(self.handle()))
}

func (self ImGuiWindow) SetSize(v ImVec2) {
	C.ImGuiWindow_SetSize(self.handle(), v.toC())
}

func (self ImGuiWindow) GetSize() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetSize(self.handle()))
}

func (self ImGuiWindow) SetSizeFull(v ImVec2) {
	C.ImGuiWindow_SetSizeFull(self.handle(), v.toC())
}

func (self ImGuiWindow) GetSizeFull() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetSizeFull(self.handle()))
}

func (self ImGuiWindow) SetContentSize(v ImVec2) {
	C.ImGuiWindow_SetContentSize(self.handle(), v.toC())
}

func (self ImGuiWindow) GetContentSize() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetContentSize(self.handle()))
}

func (self ImGuiWindow) SetContentSizeIdeal(v ImVec2) {
	C.ImGuiWindow_SetContentSizeIdeal(self.handle(), v.toC())
}

func (self ImGuiWindow) GetContentSizeIdeal() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetContentSizeIdeal(self.handle()))
}

func (self ImGuiWindow) SetContentSizeExplicit(v ImVec2) {
	C.ImGuiWindow_SetContentSizeExplicit(self.handle(), v.toC())
}

func (self ImGuiWindow) GetContentSizeExplicit() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetContentSizeExplicit(self.handle()))
}

func (self ImGuiWindow) SetWindowPadding(v ImVec2) {
	C.ImGuiWindow_SetWindowPadding(self.handle(), v.toC())
}

func (self ImGuiWindow) GetWindowPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetWindowPadding(self.handle()))
}

func (self ImGuiWindow) SetWindowRounding(v float32) {
	C.ImGuiWindow_SetWindowRounding(self.handle(), C.float(v))
}

func (self ImGuiWindow) GetWindowRounding() float32 {
	return float32(C.ImGuiWindow_GetWindowRounding(self.handle()))
}

func (self ImGuiWindow) SetWindowBorderSize(v float32) {
	C.ImGuiWindow_SetWindowBorderSize(self.handle(), C.float(v))
}

func (self ImGuiWindow) GetWindowBorderSize() float32 {
	return float32(C.ImGuiWindow_GetWindowBorderSize(self.handle()))
}

func (self ImGuiWindow) SetNameBufLen(v int32) {
	C.ImGuiWindow_SetNameBufLen(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetNameBufLen() int {
	return int(C.ImGuiWindow_GetNameBufLen(self.handle()))
}

func (self ImGuiWindow) SetMoveId(v ImGuiID) {
	C.ImGuiWindow_SetMoveId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetMoveId() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetMoveId(self.handle()))
}

func (self ImGuiWindow) SetTabId(v ImGuiID) {
	C.ImGuiWindow_SetTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetTabId() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetTabId(self.handle()))
}

func (self ImGuiWindow) SetChildId(v ImGuiID) {
	C.ImGuiWindow_SetChildId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetChildId() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetChildId(self.handle()))
}

func (self ImGuiWindow) SetScroll(v ImVec2) {
	C.ImGuiWindow_SetScroll(self.handle(), v.toC())
}

func (self ImGuiWindow) GetScroll() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetScroll(self.handle()))
}

func (self ImGuiWindow) SetScrollMax(v ImVec2) {
	C.ImGuiWindow_SetScrollMax(self.handle(), v.toC())
}

func (self ImGuiWindow) GetScrollMax() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetScrollMax(self.handle()))
}

func (self ImGuiWindow) SetScrollTarget(v ImVec2) {
	C.ImGuiWindow_SetScrollTarget(self.handle(), v.toC())
}

func (self ImGuiWindow) GetScrollTarget() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetScrollTarget(self.handle()))
}

func (self ImGuiWindow) SetScrollTargetCenterRatio(v ImVec2) {
	C.ImGuiWindow_SetScrollTargetCenterRatio(self.handle(), v.toC())
}

func (self ImGuiWindow) GetScrollTargetCenterRatio() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetScrollTargetCenterRatio(self.handle()))
}

func (self ImGuiWindow) SetScrollTargetEdgeSnapDist(v ImVec2) {
	C.ImGuiWindow_SetScrollTargetEdgeSnapDist(self.handle(), v.toC())
}

func (self ImGuiWindow) GetScrollTargetEdgeSnapDist() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetScrollTargetEdgeSnapDist(self.handle()))
}

func (self ImGuiWindow) SetScrollbarSizes(v ImVec2) {
	C.ImGuiWindow_SetScrollbarSizes(self.handle(), v.toC())
}

func (self ImGuiWindow) GetScrollbarSizes() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetScrollbarSizes(self.handle()))
}

func (self ImGuiWindow) SetScrollbarX(v bool) {
	C.ImGuiWindow_SetScrollbarX(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetScrollbarX() bool {
	return C.ImGuiWindow_GetScrollbarX(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetScrollbarY(v bool) {
	C.ImGuiWindow_SetScrollbarY(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetScrollbarY() bool {
	return C.ImGuiWindow_GetScrollbarY(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetViewportOwned(v bool) {
	C.ImGuiWindow_SetViewportOwned(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetViewportOwned() bool {
	return C.ImGuiWindow_GetViewportOwned(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetActive(v bool) {
	C.ImGuiWindow_SetActive(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetActive() bool {
	return C.ImGuiWindow_GetActive(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetWasActive(v bool) {
	C.ImGuiWindow_SetWasActive(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetWasActive() bool {
	return C.ImGuiWindow_GetWasActive(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetWriteAccessed(v bool) {
	C.ImGuiWindow_SetWriteAccessed(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetWriteAccessed() bool {
	return C.ImGuiWindow_GetWriteAccessed(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetCollapsed(v bool) {
	C.ImGuiWindow_SetCollapsed(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetCollapsed() bool {
	return C.ImGuiWindow_GetCollapsed(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetWantCollapseToggle(v bool) {
	C.ImGuiWindow_SetWantCollapseToggle(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetWantCollapseToggle() bool {
	return C.ImGuiWindow_GetWantCollapseToggle(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetSkipItems(v bool) {
	C.ImGuiWindow_SetSkipItems(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetSkipItems() bool {
	return C.ImGuiWindow_GetSkipItems(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetAppearing(v bool) {
	C.ImGuiWindow_SetAppearing(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetAppearing() bool {
	return C.ImGuiWindow_GetAppearing(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetHidden(v bool) {
	C.ImGuiWindow_SetHidden(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetHidden() bool {
	return C.ImGuiWindow_GetHidden(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetIsFallbackWindow(v bool) {
	C.ImGuiWindow_SetIsFallbackWindow(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetIsFallbackWindow() bool {
	return C.ImGuiWindow_GetIsFallbackWindow(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetIsExplicitChild(v bool) {
	C.ImGuiWindow_SetIsExplicitChild(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetIsExplicitChild() bool {
	return C.ImGuiWindow_GetIsExplicitChild(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetHasCloseButton(v bool) {
	C.ImGuiWindow_SetHasCloseButton(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetHasCloseButton() bool {
	return C.ImGuiWindow_GetHasCloseButton(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetBeginCount(v int) {
	C.ImGuiWindow_SetBeginCount(self.handle(), C.short(v))
}

func (self ImGuiWindow) GetBeginCount() int {
	return int(C.ImGuiWindow_GetBeginCount(self.handle()))
}

func (self ImGuiWindow) SetBeginOrderWithinParent(v int) {
	C.ImGuiWindow_SetBeginOrderWithinParent(self.handle(), C.short(v))
}

func (self ImGuiWindow) GetBeginOrderWithinParent() int {
	return int(C.ImGuiWindow_GetBeginOrderWithinParent(self.handle()))
}

func (self ImGuiWindow) SetBeginOrderWithinContext(v int) {
	C.ImGuiWindow_SetBeginOrderWithinContext(self.handle(), C.short(v))
}

func (self ImGuiWindow) GetBeginOrderWithinContext() int {
	return int(C.ImGuiWindow_GetBeginOrderWithinContext(self.handle()))
}

func (self ImGuiWindow) SetFocusOrder(v int) {
	C.ImGuiWindow_SetFocusOrder(self.handle(), C.short(v))
}

func (self ImGuiWindow) GetFocusOrder() int {
	return int(C.ImGuiWindow_GetFocusOrder(self.handle()))
}

func (self ImGuiWindow) SetPopupId(v ImGuiID) {
	C.ImGuiWindow_SetPopupId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetPopupId() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetPopupId(self.handle()))
}

func (self ImGuiWindow) SetAutoFitFramesX(v int) {
	C.ImGuiWindow_SetAutoFitFramesX(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetAutoFitFramesX() int {
	return int(C.ImGuiWindow_GetAutoFitFramesX(self.handle()))
}

func (self ImGuiWindow) SetAutoFitFramesY(v int) {
	C.ImGuiWindow_SetAutoFitFramesY(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetAutoFitFramesY() int {
	return int(C.ImGuiWindow_GetAutoFitFramesY(self.handle()))
}

func (self ImGuiWindow) SetAutoFitChildAxises(v int) {
	C.ImGuiWindow_SetAutoFitChildAxises(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetAutoFitChildAxises() int {
	return int(C.ImGuiWindow_GetAutoFitChildAxises(self.handle()))
}

func (self ImGuiWindow) SetAutoFitOnlyGrows(v bool) {
	C.ImGuiWindow_SetAutoFitOnlyGrows(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetAutoFitOnlyGrows() bool {
	return C.ImGuiWindow_GetAutoFitOnlyGrows(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetAutoPosLastDirection(v ImGuiDir) {
	C.ImGuiWindow_SetAutoPosLastDirection(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiWindow) GetAutoPosLastDirection() ImGuiDir {
	return ImGuiDir(C.ImGuiWindow_GetAutoPosLastDirection(self.handle()))
}

func (self ImGuiWindow) SetHiddenFramesCanSkipItems(v int) {
	C.ImGuiWindow_SetHiddenFramesCanSkipItems(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetHiddenFramesCanSkipItems() int {
	return int(C.ImGuiWindow_GetHiddenFramesCanSkipItems(self.handle()))
}

func (self ImGuiWindow) SetHiddenFramesCannotSkipItems(v int) {
	C.ImGuiWindow_SetHiddenFramesCannotSkipItems(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetHiddenFramesCannotSkipItems() int {
	return int(C.ImGuiWindow_GetHiddenFramesCannotSkipItems(self.handle()))
}

func (self ImGuiWindow) SetHiddenFramesForRenderOnly(v int) {
	C.ImGuiWindow_SetHiddenFramesForRenderOnly(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetHiddenFramesForRenderOnly() int {
	return int(C.ImGuiWindow_GetHiddenFramesForRenderOnly(self.handle()))
}

func (self ImGuiWindow) SetDisableInputsFrames(v int) {
	C.ImGuiWindow_SetDisableInputsFrames(self.handle(), C.ImS8(v))
}

func (self ImGuiWindow) GetDisableInputsFrames() int {
	return int(C.ImGuiWindow_GetDisableInputsFrames(self.handle()))
}

func (self ImGuiWindow) SetSetWindowPosAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowPosAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) GetSetWindowPosAllowFlags() ImGuiCond {
	return ImGuiCond(C.ImGuiWindow_GetSetWindowPosAllowFlags(self.handle()))
}

func (self ImGuiWindow) SetSetWindowSizeAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowSizeAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) GetSetWindowSizeAllowFlags() ImGuiCond {
	return ImGuiCond(C.ImGuiWindow_GetSetWindowSizeAllowFlags(self.handle()))
}

func (self ImGuiWindow) SetSetWindowCollapsedAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowCollapsedAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) GetSetWindowCollapsedAllowFlags() ImGuiCond {
	return ImGuiCond(C.ImGuiWindow_GetSetWindowCollapsedAllowFlags(self.handle()))
}

func (self ImGuiWindow) SetSetWindowDockAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowDockAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) GetSetWindowDockAllowFlags() ImGuiCond {
	return ImGuiCond(C.ImGuiWindow_GetSetWindowDockAllowFlags(self.handle()))
}

func (self ImGuiWindow) SetSetWindowPosVal(v ImVec2) {
	C.ImGuiWindow_SetSetWindowPosVal(self.handle(), v.toC())
}

func (self ImGuiWindow) GetSetWindowPosVal() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetSetWindowPosVal(self.handle()))
}

func (self ImGuiWindow) SetSetWindowPosPivot(v ImVec2) {
	C.ImGuiWindow_SetSetWindowPosPivot(self.handle(), v.toC())
}

func (self ImGuiWindow) GetSetWindowPosPivot() ImVec2 {
	return newImVec2FromC(C.ImGuiWindow_GetSetWindowPosPivot(self.handle()))
}

func (self ImGuiWindow) GetDC() ImGuiWindowTempData {
	return newImGuiWindowTempDataFromC(C.ImGuiWindow_GetDC(self.handle()))
}

func (self ImGuiWindow) SetOuterRectClipped(v ImRect) {
	C.ImGuiWindow_SetOuterRectClipped(self.handle(), v.toC())
}

func (self ImGuiWindow) GetOuterRectClipped() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetOuterRectClipped(self.handle()))
}

func (self ImGuiWindow) SetInnerRect(v ImRect) {
	C.ImGuiWindow_SetInnerRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetInnerRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetInnerRect(self.handle()))
}

func (self ImGuiWindow) SetInnerClipRect(v ImRect) {
	C.ImGuiWindow_SetInnerClipRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetInnerClipRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetInnerClipRect(self.handle()))
}

func (self ImGuiWindow) SetWorkRect(v ImRect) {
	C.ImGuiWindow_SetWorkRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetWorkRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetWorkRect(self.handle()))
}

func (self ImGuiWindow) SetParentWorkRect(v ImRect) {
	C.ImGuiWindow_SetParentWorkRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetParentWorkRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetParentWorkRect(self.handle()))
}

func (self ImGuiWindow) SetClipRect(v ImRect) {
	C.ImGuiWindow_SetClipRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetClipRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetClipRect(self.handle()))
}

func (self ImGuiWindow) SetContentRegionRect(v ImRect) {
	C.ImGuiWindow_SetContentRegionRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetContentRegionRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetContentRegionRect(self.handle()))
}

func (self ImGuiWindow) SetLastFrameActive(v int32) {
	C.ImGuiWindow_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetLastFrameActive() int {
	return int(C.ImGuiWindow_GetLastFrameActive(self.handle()))
}

func (self ImGuiWindow) SetLastFrameJustFocused(v int32) {
	C.ImGuiWindow_SetLastFrameJustFocused(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetLastFrameJustFocused() int {
	return int(C.ImGuiWindow_GetLastFrameJustFocused(self.handle()))
}

func (self ImGuiWindow) SetLastTimeActive(v float32) {
	C.ImGuiWindow_SetLastTimeActive(self.handle(), C.float(v))
}

func (self ImGuiWindow) GetLastTimeActive() float32 {
	return float32(C.ImGuiWindow_GetLastTimeActive(self.handle()))
}

func (self ImGuiWindow) SetItemWidthDefault(v float32) {
	C.ImGuiWindow_SetItemWidthDefault(self.handle(), C.float(v))
}

func (self ImGuiWindow) GetItemWidthDefault() float32 {
	return float32(C.ImGuiWindow_GetItemWidthDefault(self.handle()))
}

func (self ImGuiWindow) GetStateStorage() ImGuiStorage {
	return newImGuiStorageFromC(C.ImGuiWindow_GetStateStorage(self.handle()))
}

func (self ImGuiWindow) SetFontWindowScale(v float32) {
	C.ImGuiWindow_SetFontWindowScale(self.handle(), C.float(v))
}

func (self ImGuiWindow) GetFontWindowScale() float32 {
	return float32(C.ImGuiWindow_GetFontWindowScale(self.handle()))
}

func (self ImGuiWindow) SetFontDpiScale(v float32) {
	C.ImGuiWindow_SetFontDpiScale(self.handle(), C.float(v))
}

func (self ImGuiWindow) GetFontDpiScale() float32 {
	return float32(C.ImGuiWindow_GetFontDpiScale(self.handle()))
}

func (self ImGuiWindow) SetSettingsOffset(v int32) {
	C.ImGuiWindow_SetSettingsOffset(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetSettingsOffset() int {
	return int(C.ImGuiWindow_GetSettingsOffset(self.handle()))
}

func (self ImGuiWindow) SetDrawList(v ImDrawList) {
	C.ImGuiWindow_SetDrawList(self.handle(), v.handle())
}

func (self ImGuiWindow) GetDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.ImGuiWindow_GetDrawList(self.handle())))
}

func (self ImGuiWindow) GetDrawListInst() ImDrawList {
	return newImDrawListFromC(C.ImGuiWindow_GetDrawListInst(self.handle()))
}

func (self ImGuiWindow) SetParentWindow(v ImGuiWindow) {
	C.ImGuiWindow_SetParentWindow(self.handle(), v.handle())
}

func (self ImGuiWindow) GetParentWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetParentWindow(self.handle())))
}

func (self ImGuiWindow) SetParentWindowInBeginStack(v ImGuiWindow) {
	C.ImGuiWindow_SetParentWindowInBeginStack(self.handle(), v.handle())
}

func (self ImGuiWindow) GetParentWindowInBeginStack() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetParentWindowInBeginStack(self.handle())))
}

func (self ImGuiWindow) SetRootWindow(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindow(self.handle(), v.handle())
}

func (self ImGuiWindow) GetRootWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetRootWindow(self.handle())))
}

func (self ImGuiWindow) SetRootWindowPopupTree(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowPopupTree(self.handle(), v.handle())
}

func (self ImGuiWindow) GetRootWindowPopupTree() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetRootWindowPopupTree(self.handle())))
}

func (self ImGuiWindow) SetRootWindowDockTree(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowDockTree(self.handle(), v.handle())
}

func (self ImGuiWindow) GetRootWindowDockTree() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetRootWindowDockTree(self.handle())))
}

func (self ImGuiWindow) SetRootWindowForTitleBarHighlight(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowForTitleBarHighlight(self.handle(), v.handle())
}

func (self ImGuiWindow) GetRootWindowForTitleBarHighlight() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetRootWindowForTitleBarHighlight(self.handle())))
}

func (self ImGuiWindow) SetRootWindowForNav(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowForNav(self.handle(), v.handle())
}

func (self ImGuiWindow) GetRootWindowForNav() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetRootWindowForNav(self.handle())))
}

func (self ImGuiWindow) SetNavLastChildNavWindow(v ImGuiWindow) {
	C.ImGuiWindow_SetNavLastChildNavWindow(self.handle(), v.handle())
}

func (self ImGuiWindow) GetNavLastChildNavWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiWindow_GetNavLastChildNavWindow(self.handle())))
}

func (self ImGuiWindow) SetMemoryDrawListIdxCapacity(v int32) {
	C.ImGuiWindow_SetMemoryDrawListIdxCapacity(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetMemoryDrawListIdxCapacity() int {
	return int(C.ImGuiWindow_GetMemoryDrawListIdxCapacity(self.handle()))
}

func (self ImGuiWindow) SetMemoryDrawListVtxCapacity(v int32) {
	C.ImGuiWindow_SetMemoryDrawListVtxCapacity(self.handle(), C.int(v))
}

func (self ImGuiWindow) GetMemoryDrawListVtxCapacity() int {
	return int(C.ImGuiWindow_GetMemoryDrawListVtxCapacity(self.handle()))
}

func (self ImGuiWindow) SetMemoryCompacted(v bool) {
	C.ImGuiWindow_SetMemoryCompacted(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetMemoryCompacted() bool {
	return C.ImGuiWindow_GetMemoryCompacted(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetDockIsActive(v bool) {
	C.ImGuiWindow_SetDockIsActive(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetDockIsActive() bool {
	return C.ImGuiWindow_GetDockIsActive(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetDockNodeIsVisible(v bool) {
	C.ImGuiWindow_SetDockNodeIsVisible(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetDockNodeIsVisible() bool {
	return C.ImGuiWindow_GetDockNodeIsVisible(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetDockTabIsVisible(v bool) {
	C.ImGuiWindow_SetDockTabIsVisible(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetDockTabIsVisible() bool {
	return C.ImGuiWindow_GetDockTabIsVisible(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetDockTabWantClose(v bool) {
	C.ImGuiWindow_SetDockTabWantClose(self.handle(), C.bool(v))
}

func (self ImGuiWindow) GetDockTabWantClose() bool {
	return C.ImGuiWindow_GetDockTabWantClose(self.handle()) == C.bool(true)
}

func (self ImGuiWindow) SetDockOrder(v int) {
	C.ImGuiWindow_SetDockOrder(self.handle(), C.short(v))
}

func (self ImGuiWindow) GetDockOrder() int {
	return int(C.ImGuiWindow_GetDockOrder(self.handle()))
}

func (self ImGuiWindow) GetDockStyle() ImGuiWindowDockStyle {
	return newImGuiWindowDockStyleFromC(C.ImGuiWindow_GetDockStyle(self.handle()))
}

func (self ImGuiWindow) SetDockNode(v ImGuiDockNode) {
	C.ImGuiWindow_SetDockNode(self.handle(), v.handle())
}

func (self ImGuiWindow) GetDockNode() ImGuiDockNode {
	return (ImGuiDockNode)(unsafe.Pointer(C.ImGuiWindow_GetDockNode(self.handle())))
}

func (self ImGuiWindow) SetDockNodeAsHost(v ImGuiDockNode) {
	C.ImGuiWindow_SetDockNodeAsHost(self.handle(), v.handle())
}

func (self ImGuiWindow) GetDockNodeAsHost() ImGuiDockNode {
	return (ImGuiDockNode)(unsafe.Pointer(C.ImGuiWindow_GetDockNodeAsHost(self.handle())))
}

func (self ImGuiWindow) SetDockId(v ImGuiID) {
	C.ImGuiWindow_SetDockId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) GetDockId() ImGuiID {
	return ImGuiID(C.ImGuiWindow_GetDockId(self.handle()))
}

func (self ImGuiWindow) SetDockTabItemStatusFlags(v ImGuiItemStatusFlags) {
	C.ImGuiWindow_SetDockTabItemStatusFlags(self.handle(), C.ImGuiItemStatusFlags(v))
}

func (self ImGuiWindow) GetDockTabItemStatusFlags() ImGuiItemStatusFlags {
	return ImGuiItemStatusFlags(C.ImGuiWindow_GetDockTabItemStatusFlags(self.handle()))
}

func (self ImGuiWindow) SetDockTabItemRect(v ImRect) {
	C.ImGuiWindow_SetDockTabItemRect(self.handle(), v.toC())
}

func (self ImGuiWindow) GetDockTabItemRect() ImRect {
	return newImRectFromC(C.ImGuiWindow_GetDockTabItemRect(self.handle()))
}

func (self ImDrawListSplitter) Set_Current(v int32) {
	C.ImDrawListSplitter_Set_Current(self.handle(), C.int(v))
}

func (self ImDrawListSplitter) Get_Current() int {
	return int(C.ImDrawListSplitter_Get_Current(self.handle()))
}

func (self ImDrawListSplitter) Set_Count(v int32) {
	C.ImDrawListSplitter_Set_Count(self.handle(), C.int(v))
}

func (self ImDrawListSplitter) Get_Count() int {
	return int(C.ImDrawListSplitter_Get_Count(self.handle()))
}

func (self ImGuiStackLevelInfo) SetID(v ImGuiID) {
	C.ImGuiStackLevelInfo_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiStackLevelInfo) GetID() ImGuiID {
	return ImGuiID(C.ImGuiStackLevelInfo_GetID(self.handle()))
}

func (self ImGuiStackLevelInfo) SetQueryFrameCount(v int) {
	C.ImGuiStackLevelInfo_SetQueryFrameCount(self.handle(), C.ImS8(v))
}

func (self ImGuiStackLevelInfo) GetQueryFrameCount() int {
	return int(C.ImGuiStackLevelInfo_GetQueryFrameCount(self.handle()))
}

func (self ImGuiStackLevelInfo) SetQuerySuccess(v bool) {
	C.ImGuiStackLevelInfo_SetQuerySuccess(self.handle(), C.bool(v))
}

func (self ImGuiStackLevelInfo) GetQuerySuccess() bool {
	return C.ImGuiStackLevelInfo_GetQuerySuccess(self.handle()) == C.bool(true)
}

func (self ImGuiStackLevelInfo) SetDataType(v ImGuiDataType) {
	C.ImGuiStackLevelInfo_SetDataType(self.handle(), C.ImGuiDataType(v))
}

func (self ImGuiStackLevelInfo) GetDataType() ImGuiDataType {
	return ImGuiDataType(C.ImGuiStackLevelInfo_GetDataType(self.handle()))
}

func (self ImGuiTableSettings) SetID(v ImGuiID) {
	C.ImGuiTableSettings_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableSettings) GetID() ImGuiID {
	return ImGuiID(C.ImGuiTableSettings_GetID(self.handle()))
}

func (self ImGuiTableSettings) SetSaveFlags(v ImGuiTableFlags) {
	C.ImGuiTableSettings_SetSaveFlags(self.handle(), C.ImGuiTableFlags(v))
}

func (self ImGuiTableSettings) GetSaveFlags() ImGuiTableFlags {
	return ImGuiTableFlags(C.ImGuiTableSettings_GetSaveFlags(self.handle()))
}

func (self ImGuiTableSettings) SetRefScale(v float32) {
	C.ImGuiTableSettings_SetRefScale(self.handle(), C.float(v))
}

func (self ImGuiTableSettings) GetRefScale() float32 {
	return float32(C.ImGuiTableSettings_GetRefScale(self.handle()))
}

func (self ImGuiTableSettings) SetColumnsCount(v ImGuiTableColumnIdx) {
	C.ImGuiTableSettings_SetColumnsCount(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableSettings) GetColumnsCount() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableSettings_GetColumnsCount(self.handle()))
}

func (self ImGuiTableSettings) SetColumnsCountMax(v ImGuiTableColumnIdx) {
	C.ImGuiTableSettings_SetColumnsCountMax(self.handle(), C.ImGuiTableColumnIdx(v))
}

func (self ImGuiTableSettings) GetColumnsCountMax() ImGuiTableColumnIdx {
	return ImGuiTableColumnIdx(C.ImGuiTableSettings_GetColumnsCountMax(self.handle()))
}

func (self ImGuiTableSettings) SetWantApply(v bool) {
	C.ImGuiTableSettings_SetWantApply(self.handle(), C.bool(v))
}

func (self ImGuiTableSettings) GetWantApply() bool {
	return C.ImGuiTableSettings_GetWantApply(self.handle()) == C.bool(true)
}

func (self ImGuiViewportP) Get_ImGuiViewport() ImGuiViewport {
	return newImGuiViewportFromC(C.ImGuiViewportP_Get_ImGuiViewport(self.handle()))
}

func (self ImGuiViewportP) SetIdx(v int32) {
	C.ImGuiViewportP_SetIdx(self.handle(), C.int(v))
}

func (self ImGuiViewportP) GetIdx() int {
	return int(C.ImGuiViewportP_GetIdx(self.handle()))
}

func (self ImGuiViewportP) SetLastFrameActive(v int32) {
	C.ImGuiViewportP_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiViewportP) GetLastFrameActive() int {
	return int(C.ImGuiViewportP_GetLastFrameActive(self.handle()))
}

func (self ImGuiViewportP) SetLastFrontMostStampCount(v int32) {
	C.ImGuiViewportP_SetLastFrontMostStampCount(self.handle(), C.int(v))
}

func (self ImGuiViewportP) GetLastFrontMostStampCount() int {
	return int(C.ImGuiViewportP_GetLastFrontMostStampCount(self.handle()))
}

func (self ImGuiViewportP) SetLastNameHash(v ImGuiID) {
	C.ImGuiViewportP_SetLastNameHash(self.handle(), C.ImGuiID(v))
}

func (self ImGuiViewportP) GetLastNameHash() ImGuiID {
	return ImGuiID(C.ImGuiViewportP_GetLastNameHash(self.handle()))
}

func (self ImGuiViewportP) SetLastPos(v ImVec2) {
	C.ImGuiViewportP_SetLastPos(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetLastPos() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetLastPos(self.handle()))
}

func (self ImGuiViewportP) SetAlpha(v float32) {
	C.ImGuiViewportP_SetAlpha(self.handle(), C.float(v))
}

func (self ImGuiViewportP) GetAlpha() float32 {
	return float32(C.ImGuiViewportP_GetAlpha(self.handle()))
}

func (self ImGuiViewportP) SetLastAlpha(v float32) {
	C.ImGuiViewportP_SetLastAlpha(self.handle(), C.float(v))
}

func (self ImGuiViewportP) GetLastAlpha() float32 {
	return float32(C.ImGuiViewportP_GetLastAlpha(self.handle()))
}

func (self ImGuiViewportP) SetPlatformMonitor(v int) {
	C.ImGuiViewportP_SetPlatformMonitor(self.handle(), C.short(v))
}

func (self ImGuiViewportP) GetPlatformMonitor() int {
	return int(C.ImGuiViewportP_GetPlatformMonitor(self.handle()))
}

func (self ImGuiViewportP) SetPlatformWindowCreated(v bool) {
	C.ImGuiViewportP_SetPlatformWindowCreated(self.handle(), C.bool(v))
}

func (self ImGuiViewportP) GetPlatformWindowCreated() bool {
	return C.ImGuiViewportP_GetPlatformWindowCreated(self.handle()) == C.bool(true)
}

func (self ImGuiViewportP) SetWindow(v ImGuiWindow) {
	C.ImGuiViewportP_SetWindow(self.handle(), v.handle())
}

func (self ImGuiViewportP) GetWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiViewportP_GetWindow(self.handle())))
}

func (self ImGuiViewportP) GetDrawDataP() ImDrawData {
	return newImDrawDataFromC(C.ImGuiViewportP_GetDrawDataP(self.handle()))
}

func (self ImGuiViewportP) GetDrawDataBuilder() ImDrawDataBuilder {
	return newImDrawDataBuilderFromC(C.ImGuiViewportP_GetDrawDataBuilder(self.handle()))
}

func (self ImGuiViewportP) SetLastPlatformPos(v ImVec2) {
	C.ImGuiViewportP_SetLastPlatformPos(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetLastPlatformPos() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetLastPlatformPos(self.handle()))
}

func (self ImGuiViewportP) SetLastPlatformSize(v ImVec2) {
	C.ImGuiViewportP_SetLastPlatformSize(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetLastPlatformSize() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetLastPlatformSize(self.handle()))
}

func (self ImGuiViewportP) SetLastRendererSize(v ImVec2) {
	C.ImGuiViewportP_SetLastRendererSize(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetLastRendererSize() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetLastRendererSize(self.handle()))
}

func (self ImGuiViewportP) SetWorkOffsetMin(v ImVec2) {
	C.ImGuiViewportP_SetWorkOffsetMin(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetWorkOffsetMin() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetWorkOffsetMin(self.handle()))
}

func (self ImGuiViewportP) SetWorkOffsetMax(v ImVec2) {
	C.ImGuiViewportP_SetWorkOffsetMax(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetWorkOffsetMax() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetWorkOffsetMax(self.handle()))
}

func (self ImGuiViewportP) SetBuildWorkOffsetMin(v ImVec2) {
	C.ImGuiViewportP_SetBuildWorkOffsetMin(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetBuildWorkOffsetMin() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetBuildWorkOffsetMin(self.handle()))
}

func (self ImGuiViewportP) SetBuildWorkOffsetMax(v ImVec2) {
	C.ImGuiViewportP_SetBuildWorkOffsetMax(self.handle(), v.toC())
}

func (self ImGuiViewportP) GetBuildWorkOffsetMax() ImVec2 {
	return newImVec2FromC(C.ImGuiViewportP_GetBuildWorkOffsetMax(self.handle()))
}

func (self ImFontConfig) SetFontData(v unsafe.Pointer) {
	C.ImFontConfig_SetFontData(self.handle(), v)
}

func (self ImFontConfig) GetFontData() unsafe.Pointer {
	return unsafe.Pointer(C.ImFontConfig_GetFontData(self.handle()))
}

func (self ImFontConfig) SetFontDataSize(v int32) {
	C.ImFontConfig_SetFontDataSize(self.handle(), C.int(v))
}

func (self ImFontConfig) GetFontDataSize() int {
	return int(C.ImFontConfig_GetFontDataSize(self.handle()))
}

func (self ImFontConfig) SetFontDataOwnedByAtlas(v bool) {
	C.ImFontConfig_SetFontDataOwnedByAtlas(self.handle(), C.bool(v))
}

func (self ImFontConfig) GetFontDataOwnedByAtlas() bool {
	return C.ImFontConfig_GetFontDataOwnedByAtlas(self.handle()) == C.bool(true)
}

func (self ImFontConfig) SetFontNo(v int32) {
	C.ImFontConfig_SetFontNo(self.handle(), C.int(v))
}

func (self ImFontConfig) GetFontNo() int {
	return int(C.ImFontConfig_GetFontNo(self.handle()))
}

func (self ImFontConfig) SetSizePixels(v float32) {
	C.ImFontConfig_SetSizePixels(self.handle(), C.float(v))
}

func (self ImFontConfig) GetSizePixels() float32 {
	return float32(C.ImFontConfig_GetSizePixels(self.handle()))
}

func (self ImFontConfig) SetOversampleH(v int32) {
	C.ImFontConfig_SetOversampleH(self.handle(), C.int(v))
}

func (self ImFontConfig) GetOversampleH() int {
	return int(C.ImFontConfig_GetOversampleH(self.handle()))
}

func (self ImFontConfig) SetOversampleV(v int32) {
	C.ImFontConfig_SetOversampleV(self.handle(), C.int(v))
}

func (self ImFontConfig) GetOversampleV() int {
	return int(C.ImFontConfig_GetOversampleV(self.handle()))
}

func (self ImFontConfig) SetPixelSnapH(v bool) {
	C.ImFontConfig_SetPixelSnapH(self.handle(), C.bool(v))
}

func (self ImFontConfig) GetPixelSnapH() bool {
	return C.ImFontConfig_GetPixelSnapH(self.handle()) == C.bool(true)
}

func (self ImFontConfig) SetGlyphExtraSpacing(v ImVec2) {
	C.ImFontConfig_SetGlyphExtraSpacing(self.handle(), v.toC())
}

func (self ImFontConfig) GetGlyphExtraSpacing() ImVec2 {
	return newImVec2FromC(C.ImFontConfig_GetGlyphExtraSpacing(self.handle()))
}

func (self ImFontConfig) SetGlyphOffset(v ImVec2) {
	C.ImFontConfig_SetGlyphOffset(self.handle(), v.toC())
}

func (self ImFontConfig) GetGlyphOffset() ImVec2 {
	return newImVec2FromC(C.ImFontConfig_GetGlyphOffset(self.handle()))
}

func (self ImFontConfig) SetGlyphRanges(v *ImWchar) {
	C.ImFontConfig_SetGlyphRanges(self.handle(), (*C.ImWchar)(v))
}

func (self ImFontConfig) GetGlyphRanges() *ImWchar {
	return (*ImWchar)(C.ImFontConfig_GetGlyphRanges(self.handle()))
}

func (self ImFontConfig) SetGlyphMinAdvanceX(v float32) {
	C.ImFontConfig_SetGlyphMinAdvanceX(self.handle(), C.float(v))
}

func (self ImFontConfig) GetGlyphMinAdvanceX() float32 {
	return float32(C.ImFontConfig_GetGlyphMinAdvanceX(self.handle()))
}

func (self ImFontConfig) SetGlyphMaxAdvanceX(v float32) {
	C.ImFontConfig_SetGlyphMaxAdvanceX(self.handle(), C.float(v))
}

func (self ImFontConfig) GetGlyphMaxAdvanceX() float32 {
	return float32(C.ImFontConfig_GetGlyphMaxAdvanceX(self.handle()))
}

func (self ImFontConfig) SetMergeMode(v bool) {
	C.ImFontConfig_SetMergeMode(self.handle(), C.bool(v))
}

func (self ImFontConfig) GetMergeMode() bool {
	return C.ImFontConfig_GetMergeMode(self.handle()) == C.bool(true)
}

func (self ImFontConfig) SetFontBuilderFlags(v uint32) {
	C.ImFontConfig_SetFontBuilderFlags(self.handle(), C.uint(v))
}

func (self ImFontConfig) GetFontBuilderFlags() uint32 {
	return uint32(C.ImFontConfig_GetFontBuilderFlags(self.handle()))
}

func (self ImFontConfig) SetRasterizerMultiply(v float32) {
	C.ImFontConfig_SetRasterizerMultiply(self.handle(), C.float(v))
}

func (self ImFontConfig) GetRasterizerMultiply() float32 {
	return float32(C.ImFontConfig_GetRasterizerMultiply(self.handle()))
}

func (self ImFontConfig) SetEllipsisChar(v ImWchar) {
	C.ImFontConfig_SetEllipsisChar(self.handle(), C.ImWchar(v))
}

func (self ImFontConfig) SetDstFont(v ImFont) {
	C.ImFontConfig_SetDstFont(self.handle(), v.handle())
}

func (self ImFontConfig) GetDstFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.ImFontConfig_GetDstFont(self.handle())))
}

func (self ImFontAtlasCustomRect) SetWidth(v uint) {
	C.ImFontAtlasCustomRect_SetWidth(self.handle(), C.ushort(v))
}

func (self ImFontAtlasCustomRect) SetHeight(v uint) {
	C.ImFontAtlasCustomRect_SetHeight(self.handle(), C.ushort(v))
}

func (self ImFontAtlasCustomRect) SetX(v uint) {
	C.ImFontAtlasCustomRect_SetX(self.handle(), C.ushort(v))
}

func (self ImFontAtlasCustomRect) SetY(v uint) {
	C.ImFontAtlasCustomRect_SetY(self.handle(), C.ushort(v))
}

func (self ImFontAtlasCustomRect) SetGlyphID(v uint32) {
	C.ImFontAtlasCustomRect_SetGlyphID(self.handle(), C.uint(v))
}

func (self ImFontAtlasCustomRect) GetGlyphID() uint32 {
	return uint32(C.ImFontAtlasCustomRect_GetGlyphID(self.handle()))
}

func (self ImFontAtlasCustomRect) SetGlyphAdvanceX(v float32) {
	C.ImFontAtlasCustomRect_SetGlyphAdvanceX(self.handle(), C.float(v))
}

func (self ImFontAtlasCustomRect) GetGlyphAdvanceX() float32 {
	return float32(C.ImFontAtlasCustomRect_GetGlyphAdvanceX(self.handle()))
}

func (self ImFontAtlasCustomRect) SetGlyphOffset(v ImVec2) {
	C.ImFontAtlasCustomRect_SetGlyphOffset(self.handle(), v.toC())
}

func (self ImFontAtlasCustomRect) GetGlyphOffset() ImVec2 {
	return newImVec2FromC(C.ImFontAtlasCustomRect_GetGlyphOffset(self.handle()))
}

func (self ImFontAtlasCustomRect) SetFont(v ImFont) {
	C.ImFontAtlasCustomRect_SetFont(self.handle(), v.handle())
}

func (self ImFontAtlasCustomRect) GetFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.ImFontAtlasCustomRect_GetFont(self.handle())))
}

func (self ImDrawVert) Setpos(v ImVec2) {
	C.ImDrawVert_Setpos(self.handle(), v.toC())
}

func (self ImDrawVert) Getpos() ImVec2 {
	return newImVec2FromC(C.ImDrawVert_Getpos(self.handle()))
}

func (self ImDrawVert) Setuv(v ImVec2) {
	C.ImDrawVert_Setuv(self.handle(), v.toC())
}

func (self ImDrawVert) Getuv() ImVec2 {
	return newImVec2FromC(C.ImDrawVert_Getuv(self.handle()))
}

func (self ImDrawVert) Setcol(v uint32) {
	C.ImDrawVert_Setcol(self.handle(), C.ImU32(v))
}

func (self ImDrawVert) Getcol() uint32 {
	return uint32(C.ImDrawVert_Getcol(self.handle()))
}

func (self ImGuiMenuColumns) SetTotalWidth(v uint32) {
	C.ImGuiMenuColumns_SetTotalWidth(self.handle(), C.ImU32(v))
}

func (self ImGuiMenuColumns) GetTotalWidth() uint32 {
	return uint32(C.ImGuiMenuColumns_GetTotalWidth(self.handle()))
}

func (self ImGuiMenuColumns) SetNextTotalWidth(v uint32) {
	C.ImGuiMenuColumns_SetNextTotalWidth(self.handle(), C.ImU32(v))
}

func (self ImGuiMenuColumns) GetNextTotalWidth() uint32 {
	return uint32(C.ImGuiMenuColumns_GetNextTotalWidth(self.handle()))
}

func (self ImGuiMenuColumns) SetSpacing(v uint) {
	C.ImGuiMenuColumns_SetSpacing(self.handle(), C.ImU16(v))
}

func (self ImGuiMenuColumns) GetSpacing() uint32 {
	return uint32(C.ImGuiMenuColumns_GetSpacing(self.handle()))
}

func (self ImGuiMenuColumns) SetOffsetIcon(v uint) {
	C.ImGuiMenuColumns_SetOffsetIcon(self.handle(), C.ImU16(v))
}

func (self ImGuiMenuColumns) GetOffsetIcon() uint32 {
	return uint32(C.ImGuiMenuColumns_GetOffsetIcon(self.handle()))
}

func (self ImGuiMenuColumns) SetOffsetLabel(v uint) {
	C.ImGuiMenuColumns_SetOffsetLabel(self.handle(), C.ImU16(v))
}

func (self ImGuiMenuColumns) GetOffsetLabel() uint32 {
	return uint32(C.ImGuiMenuColumns_GetOffsetLabel(self.handle()))
}

func (self ImGuiMenuColumns) SetOffsetShortcut(v uint) {
	C.ImGuiMenuColumns_SetOffsetShortcut(self.handle(), C.ImU16(v))
}

func (self ImGuiMenuColumns) GetOffsetShortcut() uint32 {
	return uint32(C.ImGuiMenuColumns_GetOffsetShortcut(self.handle()))
}

func (self ImGuiMenuColumns) SetOffsetMark(v uint) {
	C.ImGuiMenuColumns_SetOffsetMark(self.handle(), C.ImU16(v))
}

func (self ImGuiMenuColumns) GetOffsetMark() uint32 {
	return uint32(C.ImGuiMenuColumns_GetOffsetMark(self.handle()))
}

func (self ImGuiNavItemData) SetWindow(v ImGuiWindow) {
	C.ImGuiNavItemData_SetWindow(self.handle(), v.handle())
}

func (self ImGuiNavItemData) GetWindow() ImGuiWindow {
	return (ImGuiWindow)(unsafe.Pointer(C.ImGuiNavItemData_GetWindow(self.handle())))
}

func (self ImGuiNavItemData) SetID(v ImGuiID) {
	C.ImGuiNavItemData_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNavItemData) GetID() ImGuiID {
	return ImGuiID(C.ImGuiNavItemData_GetID(self.handle()))
}

func (self ImGuiNavItemData) SetFocusScopeId(v ImGuiID) {
	C.ImGuiNavItemData_SetFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNavItemData) GetFocusScopeId() ImGuiID {
	return ImGuiID(C.ImGuiNavItemData_GetFocusScopeId(self.handle()))
}

func (self ImGuiNavItemData) SetRectRel(v ImRect) {
	C.ImGuiNavItemData_SetRectRel(self.handle(), v.toC())
}

func (self ImGuiNavItemData) GetRectRel() ImRect {
	return newImRectFromC(C.ImGuiNavItemData_GetRectRel(self.handle()))
}

func (self ImGuiNavItemData) SetInFlags(v ImGuiItemFlags) {
	C.ImGuiNavItemData_SetInFlags(self.handle(), C.ImGuiItemFlags(v))
}

func (self ImGuiNavItemData) GetInFlags() ImGuiItemFlags {
	return ImGuiItemFlags(C.ImGuiNavItemData_GetInFlags(self.handle()))
}

func (self ImGuiNavItemData) SetDistBox(v float32) {
	C.ImGuiNavItemData_SetDistBox(self.handle(), C.float(v))
}

func (self ImGuiNavItemData) GetDistBox() float32 {
	return float32(C.ImGuiNavItemData_GetDistBox(self.handle()))
}

func (self ImGuiNavItemData) SetDistCenter(v float32) {
	C.ImGuiNavItemData_SetDistCenter(self.handle(), C.float(v))
}

func (self ImGuiNavItemData) GetDistCenter() float32 {
	return float32(C.ImGuiNavItemData_GetDistCenter(self.handle()))
}

func (self ImGuiNavItemData) SetDistAxial(v float32) {
	C.ImGuiNavItemData_SetDistAxial(self.handle(), C.float(v))
}

func (self ImGuiNavItemData) GetDistAxial() float32 {
	return float32(C.ImGuiNavItemData_GetDistAxial(self.handle()))
}

func (self ImGuiPayload) SetData(v unsafe.Pointer) {
	C.ImGuiPayload_SetData(self.handle(), v)
}

func (self ImGuiPayload) GetData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiPayload_GetData(self.handle()))
}

func (self ImGuiPayload) SetDataSize(v int32) {
	C.ImGuiPayload_SetDataSize(self.handle(), C.int(v))
}

func (self ImGuiPayload) GetDataSize() int {
	return int(C.ImGuiPayload_GetDataSize(self.handle()))
}

func (self ImGuiPayload) SetSourceId(v ImGuiID) {
	C.ImGuiPayload_SetSourceId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPayload) GetSourceId() ImGuiID {
	return ImGuiID(C.ImGuiPayload_GetSourceId(self.handle()))
}

func (self ImGuiPayload) SetSourceParentId(v ImGuiID) {
	C.ImGuiPayload_SetSourceParentId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPayload) GetSourceParentId() ImGuiID {
	return ImGuiID(C.ImGuiPayload_GetSourceParentId(self.handle()))
}

func (self ImGuiPayload) SetDataFrameCount(v int32) {
	C.ImGuiPayload_SetDataFrameCount(self.handle(), C.int(v))
}

func (self ImGuiPayload) GetDataFrameCount() int {
	return int(C.ImGuiPayload_GetDataFrameCount(self.handle()))
}

func (self ImGuiPayload) SetPreview(v bool) {
	C.ImGuiPayload_SetPreview(self.handle(), C.bool(v))
}

func (self ImGuiPayload) GetPreview() bool {
	return C.ImGuiPayload_GetPreview(self.handle()) == C.bool(true)
}

func (self ImGuiPayload) SetDelivery(v bool) {
	C.ImGuiPayload_SetDelivery(self.handle(), C.bool(v))
}

func (self ImGuiPayload) GetDelivery() bool {
	return C.ImGuiPayload_GetDelivery(self.handle()) == C.bool(true)
}

func (self ImGuiShrinkWidthItem) SetIndex(v int32) {
	C.ImGuiShrinkWidthItem_SetIndex(self.handle(), C.int(v))
}

func (self ImGuiShrinkWidthItem) GetIndex() int {
	return int(C.ImGuiShrinkWidthItem_GetIndex(self.handle()))
}

func (self ImGuiShrinkWidthItem) SetWidth(v float32) {
	C.ImGuiShrinkWidthItem_SetWidth(self.handle(), C.float(v))
}

func (self ImGuiShrinkWidthItem) GetWidth() float32 {
	return float32(C.ImGuiShrinkWidthItem_GetWidth(self.handle()))
}

func (self ImGuiShrinkWidthItem) SetInitialWidth(v float32) {
	C.ImGuiShrinkWidthItem_SetInitialWidth(self.handle(), C.float(v))
}

func (self ImGuiShrinkWidthItem) GetInitialWidth() float32 {
	return float32(C.ImGuiShrinkWidthItem_GetInitialWidth(self.handle()))
}

func (self ImDrawCmd) SetClipRect(v ImVec4) {
	C.ImDrawCmd_SetClipRect(self.handle(), v.toC())
}

func (self ImDrawCmd) GetClipRect() ImVec4 {
	return newImVec4FromC(C.ImDrawCmd_GetClipRect(self.handle()))
}

func (self ImDrawCmd) SetTextureId(v ImTextureID) {
	C.ImDrawCmd_SetTextureId(self.handle(), C.ImTextureID(v))
}

func (self ImDrawCmd) GetTextureId() ImTextureID {
	return ImTextureID(C.ImDrawCmd_GetTextureId(self.handle()))
}

func (self ImDrawCmd) SetVtxOffset(v uint32) {
	C.ImDrawCmd_SetVtxOffset(self.handle(), C.uint(v))
}

func (self ImDrawCmd) GetVtxOffset() uint32 {
	return uint32(C.ImDrawCmd_GetVtxOffset(self.handle()))
}

func (self ImDrawCmd) SetIdxOffset(v uint32) {
	C.ImDrawCmd_SetIdxOffset(self.handle(), C.uint(v))
}

func (self ImDrawCmd) GetIdxOffset() uint32 {
	return uint32(C.ImDrawCmd_GetIdxOffset(self.handle()))
}

func (self ImDrawCmd) SetElemCount(v uint32) {
	C.ImDrawCmd_SetElemCount(self.handle(), C.uint(v))
}

func (self ImDrawCmd) GetElemCount() uint32 {
	return uint32(C.ImDrawCmd_GetElemCount(self.handle()))
}

func (self ImDrawCmd) SetUserCallbackData(v unsafe.Pointer) {
	C.ImDrawCmd_SetUserCallbackData(self.handle(), v)
}

func (self ImDrawCmd) GetUserCallbackData() unsafe.Pointer {
	return unsafe.Pointer(C.ImDrawCmd_GetUserCallbackData(self.handle()))
}

func (self ImGuiInputEventText) SetChar(v uint32) {
	C.ImGuiInputEventText_SetChar(self.handle(), C.uint(v))
}

func (self ImGuiInputEventText) GetChar() uint32 {
	return uint32(C.ImGuiInputEventText_GetChar(self.handle()))
}

func (self ImGuiOldColumns) SetID(v ImGuiID) {
	C.ImGuiOldColumns_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiOldColumns) GetID() ImGuiID {
	return ImGuiID(C.ImGuiOldColumns_GetID(self.handle()))
}

func (self ImGuiOldColumns) SetFlags(v ImGuiOldColumnFlags) {
	C.ImGuiOldColumns_SetFlags(self.handle(), C.ImGuiOldColumnFlags(v))
}

func (self ImGuiOldColumns) GetFlags() ImGuiOldColumnFlags {
	return ImGuiOldColumnFlags(C.ImGuiOldColumns_GetFlags(self.handle()))
}

func (self ImGuiOldColumns) SetIsFirstFrame(v bool) {
	C.ImGuiOldColumns_SetIsFirstFrame(self.handle(), C.bool(v))
}

func (self ImGuiOldColumns) GetIsFirstFrame() bool {
	return C.ImGuiOldColumns_GetIsFirstFrame(self.handle()) == C.bool(true)
}

func (self ImGuiOldColumns) SetIsBeingResized(v bool) {
	C.ImGuiOldColumns_SetIsBeingResized(self.handle(), C.bool(v))
}

func (self ImGuiOldColumns) GetIsBeingResized() bool {
	return C.ImGuiOldColumns_GetIsBeingResized(self.handle()) == C.bool(true)
}

func (self ImGuiOldColumns) SetCurrent(v int32) {
	C.ImGuiOldColumns_SetCurrent(self.handle(), C.int(v))
}

func (self ImGuiOldColumns) GetCurrent() int {
	return int(C.ImGuiOldColumns_GetCurrent(self.handle()))
}

func (self ImGuiOldColumns) SetCount(v int32) {
	C.ImGuiOldColumns_SetCount(self.handle(), C.int(v))
}

func (self ImGuiOldColumns) GetCount() int {
	return int(C.ImGuiOldColumns_GetCount(self.handle()))
}

func (self ImGuiOldColumns) SetOffMinX(v float32) {
	C.ImGuiOldColumns_SetOffMinX(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) GetOffMinX() float32 {
	return float32(C.ImGuiOldColumns_GetOffMinX(self.handle()))
}

func (self ImGuiOldColumns) SetOffMaxX(v float32) {
	C.ImGuiOldColumns_SetOffMaxX(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) GetOffMaxX() float32 {
	return float32(C.ImGuiOldColumns_GetOffMaxX(self.handle()))
}

func (self ImGuiOldColumns) SetLineMinY(v float32) {
	C.ImGuiOldColumns_SetLineMinY(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) GetLineMinY() float32 {
	return float32(C.ImGuiOldColumns_GetLineMinY(self.handle()))
}

func (self ImGuiOldColumns) SetLineMaxY(v float32) {
	C.ImGuiOldColumns_SetLineMaxY(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) GetLineMaxY() float32 {
	return float32(C.ImGuiOldColumns_GetLineMaxY(self.handle()))
}

func (self ImGuiOldColumns) SetHostCursorPosY(v float32) {
	C.ImGuiOldColumns_SetHostCursorPosY(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) GetHostCursorPosY() float32 {
	return float32(C.ImGuiOldColumns_GetHostCursorPosY(self.handle()))
}

func (self ImGuiOldColumns) SetHostCursorMaxPosX(v float32) {
	C.ImGuiOldColumns_SetHostCursorMaxPosX(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) GetHostCursorMaxPosX() float32 {
	return float32(C.ImGuiOldColumns_GetHostCursorMaxPosX(self.handle()))
}

func (self ImGuiOldColumns) SetHostInitialClipRect(v ImRect) {
	C.ImGuiOldColumns_SetHostInitialClipRect(self.handle(), v.toC())
}

func (self ImGuiOldColumns) GetHostInitialClipRect() ImRect {
	return newImRectFromC(C.ImGuiOldColumns_GetHostInitialClipRect(self.handle()))
}

func (self ImGuiOldColumns) SetHostBackupClipRect(v ImRect) {
	C.ImGuiOldColumns_SetHostBackupClipRect(self.handle(), v.toC())
}

func (self ImGuiOldColumns) GetHostBackupClipRect() ImRect {
	return newImRectFromC(C.ImGuiOldColumns_GetHostBackupClipRect(self.handle()))
}

func (self ImGuiOldColumns) SetHostBackupParentWorkRect(v ImRect) {
	C.ImGuiOldColumns_SetHostBackupParentWorkRect(self.handle(), v.toC())
}

func (self ImGuiOldColumns) GetHostBackupParentWorkRect() ImRect {
	return newImRectFromC(C.ImGuiOldColumns_GetHostBackupParentWorkRect(self.handle()))
}

func (self ImGuiOldColumns) GetSplitter() ImDrawListSplitter {
	return newImDrawListSplitterFromC(C.ImGuiOldColumns_GetSplitter(self.handle()))
}

func (self ImGuiSizeCallbackData) SetUserData(v unsafe.Pointer) {
	C.ImGuiSizeCallbackData_SetUserData(self.handle(), v)
}

func (self ImGuiSizeCallbackData) GetUserData() unsafe.Pointer {
	return unsafe.Pointer(C.ImGuiSizeCallbackData_GetUserData(self.handle()))
}

func (self ImGuiSizeCallbackData) SetPos(v ImVec2) {
	C.ImGuiSizeCallbackData_SetPos(self.handle(), v.toC())
}

func (self ImGuiSizeCallbackData) GetPos() ImVec2 {
	return newImVec2FromC(C.ImGuiSizeCallbackData_GetPos(self.handle()))
}

func (self ImGuiSizeCallbackData) SetCurrentSize(v ImVec2) {
	C.ImGuiSizeCallbackData_SetCurrentSize(self.handle(), v.toC())
}

func (self ImGuiSizeCallbackData) GetCurrentSize() ImVec2 {
	return newImVec2FromC(C.ImGuiSizeCallbackData_GetCurrentSize(self.handle()))
}

func (self ImGuiSizeCallbackData) SetDesiredSize(v ImVec2) {
	C.ImGuiSizeCallbackData_SetDesiredSize(self.handle(), v.toC())
}

func (self ImGuiSizeCallbackData) GetDesiredSize() ImVec2 {
	return newImVec2FromC(C.ImGuiSizeCallbackData_GetDesiredSize(self.handle()))
}

func (self ImGuiStyle) SetAlpha(v float32) {
	C.ImGuiStyle_SetAlpha(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetAlpha() float32 {
	return float32(C.ImGuiStyle_GetAlpha(self.handle()))
}

func (self ImGuiStyle) SetDisabledAlpha(v float32) {
	C.ImGuiStyle_SetDisabledAlpha(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetDisabledAlpha() float32 {
	return float32(C.ImGuiStyle_GetDisabledAlpha(self.handle()))
}

func (self ImGuiStyle) SetWindowPadding(v ImVec2) {
	C.ImGuiStyle_SetWindowPadding(self.handle(), v.toC())
}

func (self ImGuiStyle) GetWindowPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetWindowPadding(self.handle()))
}

func (self ImGuiStyle) SetWindowRounding(v float32) {
	C.ImGuiStyle_SetWindowRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetWindowRounding() float32 {
	return float32(C.ImGuiStyle_GetWindowRounding(self.handle()))
}

func (self ImGuiStyle) SetWindowBorderSize(v float32) {
	C.ImGuiStyle_SetWindowBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetWindowBorderSize() float32 {
	return float32(C.ImGuiStyle_GetWindowBorderSize(self.handle()))
}

func (self ImGuiStyle) SetWindowMinSize(v ImVec2) {
	C.ImGuiStyle_SetWindowMinSize(self.handle(), v.toC())
}

func (self ImGuiStyle) GetWindowMinSize() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetWindowMinSize(self.handle()))
}

func (self ImGuiStyle) SetWindowTitleAlign(v ImVec2) {
	C.ImGuiStyle_SetWindowTitleAlign(self.handle(), v.toC())
}

func (self ImGuiStyle) GetWindowTitleAlign() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetWindowTitleAlign(self.handle()))
}

func (self ImGuiStyle) SetWindowMenuButtonPosition(v ImGuiDir) {
	C.ImGuiStyle_SetWindowMenuButtonPosition(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiStyle) GetWindowMenuButtonPosition() ImGuiDir {
	return ImGuiDir(C.ImGuiStyle_GetWindowMenuButtonPosition(self.handle()))
}

func (self ImGuiStyle) SetChildRounding(v float32) {
	C.ImGuiStyle_SetChildRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetChildRounding() float32 {
	return float32(C.ImGuiStyle_GetChildRounding(self.handle()))
}

func (self ImGuiStyle) SetChildBorderSize(v float32) {
	C.ImGuiStyle_SetChildBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetChildBorderSize() float32 {
	return float32(C.ImGuiStyle_GetChildBorderSize(self.handle()))
}

func (self ImGuiStyle) SetPopupRounding(v float32) {
	C.ImGuiStyle_SetPopupRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetPopupRounding() float32 {
	return float32(C.ImGuiStyle_GetPopupRounding(self.handle()))
}

func (self ImGuiStyle) SetPopupBorderSize(v float32) {
	C.ImGuiStyle_SetPopupBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetPopupBorderSize() float32 {
	return float32(C.ImGuiStyle_GetPopupBorderSize(self.handle()))
}

func (self ImGuiStyle) SetFramePadding(v ImVec2) {
	C.ImGuiStyle_SetFramePadding(self.handle(), v.toC())
}

func (self ImGuiStyle) GetFramePadding() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetFramePadding(self.handle()))
}

func (self ImGuiStyle) SetFrameRounding(v float32) {
	C.ImGuiStyle_SetFrameRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetFrameRounding() float32 {
	return float32(C.ImGuiStyle_GetFrameRounding(self.handle()))
}

func (self ImGuiStyle) SetFrameBorderSize(v float32) {
	C.ImGuiStyle_SetFrameBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetFrameBorderSize() float32 {
	return float32(C.ImGuiStyle_GetFrameBorderSize(self.handle()))
}

func (self ImGuiStyle) SetItemSpacing(v ImVec2) {
	C.ImGuiStyle_SetItemSpacing(self.handle(), v.toC())
}

func (self ImGuiStyle) GetItemSpacing() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetItemSpacing(self.handle()))
}

func (self ImGuiStyle) SetItemInnerSpacing(v ImVec2) {
	C.ImGuiStyle_SetItemInnerSpacing(self.handle(), v.toC())
}

func (self ImGuiStyle) GetItemInnerSpacing() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetItemInnerSpacing(self.handle()))
}

func (self ImGuiStyle) SetCellPadding(v ImVec2) {
	C.ImGuiStyle_SetCellPadding(self.handle(), v.toC())
}

func (self ImGuiStyle) GetCellPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetCellPadding(self.handle()))
}

func (self ImGuiStyle) SetTouchExtraPadding(v ImVec2) {
	C.ImGuiStyle_SetTouchExtraPadding(self.handle(), v.toC())
}

func (self ImGuiStyle) GetTouchExtraPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetTouchExtraPadding(self.handle()))
}

func (self ImGuiStyle) SetIndentSpacing(v float32) {
	C.ImGuiStyle_SetIndentSpacing(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetIndentSpacing() float32 {
	return float32(C.ImGuiStyle_GetIndentSpacing(self.handle()))
}

func (self ImGuiStyle) SetColumnsMinSpacing(v float32) {
	C.ImGuiStyle_SetColumnsMinSpacing(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetColumnsMinSpacing() float32 {
	return float32(C.ImGuiStyle_GetColumnsMinSpacing(self.handle()))
}

func (self ImGuiStyle) SetScrollbarSize(v float32) {
	C.ImGuiStyle_SetScrollbarSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetScrollbarSize() float32 {
	return float32(C.ImGuiStyle_GetScrollbarSize(self.handle()))
}

func (self ImGuiStyle) SetScrollbarRounding(v float32) {
	C.ImGuiStyle_SetScrollbarRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetScrollbarRounding() float32 {
	return float32(C.ImGuiStyle_GetScrollbarRounding(self.handle()))
}

func (self ImGuiStyle) SetGrabMinSize(v float32) {
	C.ImGuiStyle_SetGrabMinSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetGrabMinSize() float32 {
	return float32(C.ImGuiStyle_GetGrabMinSize(self.handle()))
}

func (self ImGuiStyle) SetGrabRounding(v float32) {
	C.ImGuiStyle_SetGrabRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetGrabRounding() float32 {
	return float32(C.ImGuiStyle_GetGrabRounding(self.handle()))
}

func (self ImGuiStyle) SetLogSliderDeadzone(v float32) {
	C.ImGuiStyle_SetLogSliderDeadzone(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetLogSliderDeadzone() float32 {
	return float32(C.ImGuiStyle_GetLogSliderDeadzone(self.handle()))
}

func (self ImGuiStyle) SetTabRounding(v float32) {
	C.ImGuiStyle_SetTabRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetTabRounding() float32 {
	return float32(C.ImGuiStyle_GetTabRounding(self.handle()))
}

func (self ImGuiStyle) SetTabBorderSize(v float32) {
	C.ImGuiStyle_SetTabBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetTabBorderSize() float32 {
	return float32(C.ImGuiStyle_GetTabBorderSize(self.handle()))
}

func (self ImGuiStyle) SetTabMinWidthForCloseButton(v float32) {
	C.ImGuiStyle_SetTabMinWidthForCloseButton(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetTabMinWidthForCloseButton() float32 {
	return float32(C.ImGuiStyle_GetTabMinWidthForCloseButton(self.handle()))
}

func (self ImGuiStyle) SetColorButtonPosition(v ImGuiDir) {
	C.ImGuiStyle_SetColorButtonPosition(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiStyle) GetColorButtonPosition() ImGuiDir {
	return ImGuiDir(C.ImGuiStyle_GetColorButtonPosition(self.handle()))
}

func (self ImGuiStyle) SetButtonTextAlign(v ImVec2) {
	C.ImGuiStyle_SetButtonTextAlign(self.handle(), v.toC())
}

func (self ImGuiStyle) GetButtonTextAlign() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetButtonTextAlign(self.handle()))
}

func (self ImGuiStyle) SetSelectableTextAlign(v ImVec2) {
	C.ImGuiStyle_SetSelectableTextAlign(self.handle(), v.toC())
}

func (self ImGuiStyle) GetSelectableTextAlign() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetSelectableTextAlign(self.handle()))
}

func (self ImGuiStyle) SetDisplayWindowPadding(v ImVec2) {
	C.ImGuiStyle_SetDisplayWindowPadding(self.handle(), v.toC())
}

func (self ImGuiStyle) GetDisplayWindowPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetDisplayWindowPadding(self.handle()))
}

func (self ImGuiStyle) SetDisplaySafeAreaPadding(v ImVec2) {
	C.ImGuiStyle_SetDisplaySafeAreaPadding(self.handle(), v.toC())
}

func (self ImGuiStyle) GetDisplaySafeAreaPadding() ImVec2 {
	return newImVec2FromC(C.ImGuiStyle_GetDisplaySafeAreaPadding(self.handle()))
}

func (self ImGuiStyle) SetMouseCursorScale(v float32) {
	C.ImGuiStyle_SetMouseCursorScale(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetMouseCursorScale() float32 {
	return float32(C.ImGuiStyle_GetMouseCursorScale(self.handle()))
}

func (self ImGuiStyle) SetAntiAliasedLines(v bool) {
	C.ImGuiStyle_SetAntiAliasedLines(self.handle(), C.bool(v))
}

func (self ImGuiStyle) GetAntiAliasedLines() bool {
	return C.ImGuiStyle_GetAntiAliasedLines(self.handle()) == C.bool(true)
}

func (self ImGuiStyle) SetAntiAliasedLinesUseTex(v bool) {
	C.ImGuiStyle_SetAntiAliasedLinesUseTex(self.handle(), C.bool(v))
}

func (self ImGuiStyle) GetAntiAliasedLinesUseTex() bool {
	return C.ImGuiStyle_GetAntiAliasedLinesUseTex(self.handle()) == C.bool(true)
}

func (self ImGuiStyle) SetAntiAliasedFill(v bool) {
	C.ImGuiStyle_SetAntiAliasedFill(self.handle(), C.bool(v))
}

func (self ImGuiStyle) GetAntiAliasedFill() bool {
	return C.ImGuiStyle_GetAntiAliasedFill(self.handle()) == C.bool(true)
}

func (self ImGuiStyle) SetCurveTessellationTol(v float32) {
	C.ImGuiStyle_SetCurveTessellationTol(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetCurveTessellationTol() float32 {
	return float32(C.ImGuiStyle_GetCurveTessellationTol(self.handle()))
}

func (self ImGuiStyle) SetCircleTessellationMaxError(v float32) {
	C.ImGuiStyle_SetCircleTessellationMaxError(self.handle(), C.float(v))
}

func (self ImGuiStyle) GetCircleTessellationMaxError() float32 {
	return float32(C.ImGuiStyle_GetCircleTessellationMaxError(self.handle()))
}

func (self ImGuiDataTypeInfo) SetSize(v uint64) {
	C.ImGuiDataTypeInfo_SetSize(self.handle(), C.xlong(v))
}

func (self ImGuiDataTypeInfo) GetSize() float64 {
	return float64(C.ImGuiDataTypeInfo_GetSize(self.handle()))
}

func (self ImGuiDataTypeInfo) SetName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiDataTypeInfo_SetName(self.handle(), vArg)
}

func (self ImGuiDataTypeInfo) GetName() string {
	return C.GoString(C.ImGuiDataTypeInfo_GetName(self.handle()))
}

func (self ImGuiDataTypeInfo) SetPrintFmt(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiDataTypeInfo_SetPrintFmt(self.handle(), vArg)
}

func (self ImGuiDataTypeInfo) GetPrintFmt() string {
	return C.GoString(C.ImGuiDataTypeInfo_GetPrintFmt(self.handle()))
}

func (self ImGuiDataTypeInfo) SetScanFmt(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiDataTypeInfo_SetScanFmt(self.handle(), vArg)
}

func (self ImGuiDataTypeInfo) GetScanFmt() string {
	return C.GoString(C.ImGuiDataTypeInfo_GetScanFmt(self.handle()))
}
