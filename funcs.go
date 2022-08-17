package cimgui

// #include "cimgui_structs_accessor.h"
// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

func IsAnyMouseDown() bool {
	return C.IsAnyMouseDown() != C.bool(true)
}

func BeginListBox(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.BeginListBox(labelArg, size.ToC()) != C.bool(true)
}

func PopStyleColor(count int32) {
	C.PopStyleColor(C.int(count))
}

func VSliderFloat(label string, size ImVec2, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderFloat(labelArg, size.ToC(), vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawList) AddImage(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32) {
	C.ImDrawList_AddImage(self.handle(), C.ImTextureID(user_texture_id), p_min.ToC(), p_max.ToC(), uv_min.ToC(), uv_max.ToC(), C.uint(col))
}

func (self ImFontAtlas) GetTexDataAsRGBA32(out_pixels *C.uchar, out_width *int32, out_height *int32, out_bytes_per_pixel *int32) {
	out_widthArg, out_widthFin := wrapInt32(out_width)
	defer out_widthFin()

	out_heightArg, out_heightFin := wrapInt32(out_height)
	defer out_heightFin()

	out_bytes_per_pixelArg, out_bytes_per_pixelFin := wrapInt32(out_bytes_per_pixel)
	defer out_bytes_per_pixelFin()

	C.ImFontAtlas_GetTexDataAsRGBA32(self.handle(), &out_pixels, out_widthArg, out_heightArg, out_bytes_per_pixelArg)
}

func GetMainViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.GetMainViewport()))
}

func GetWindowWidth() float32 {
	return float32(C.GetWindowWidth())
}

func (self ImDrawList) PrimUnreserve(idx_count int32, vtx_count int32) {
	C.ImDrawList_PrimUnreserve(self.handle(), C.int(idx_count), C.int(vtx_count))
}

func EndDragDropSource() {
	C.EndDragDropSource()
}

func (self ImFontGlyphRangesBuilder) GetBit(n uint64) bool {
	return C.ImFontGlyphRangesBuilder_GetBit(self.handle(), C.ulong(n)) != C.bool(true)
}

func DragFloat2(label string, v *[2]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat2(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndFrame() {
	C.EndFrame()
}

func NewLine() {
	C.NewLine()
}

func (self ImDrawList) AddConvexPolyFilled(points *ImVec2, num_points int32, col uint32) {
	pointsArg, pointsFin := points.wrap()
	defer pointsFin()

	C.ImDrawList_AddConvexPolyFilled(self.handle(), pointsArg, C.int(num_points), C.uint(col))
}

func (self ImDrawList) PrimVtx(pos ImVec2, uv ImVec2, col uint32) {
	C.ImDrawList_PrimVtx(self.handle(), pos.ToC(), uv.ToC(), C.uint(col))
}

func SetDragDropPayload(typeArg string, data unsafe.Pointer, sz uint64, cond ImGuiCond) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.SetDragDropPayload(typeArgArg, data, C.ulong(sz), C.ImGuiCond(cond)) != C.bool(true)
}

func ShowFontSelector(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ShowFontSelector(labelArg)
}

func (self ImDrawData) DeIndexAllBuffers() {
	C.ImDrawData_DeIndexAllBuffers(self.handle())
}

func Indent(indent_w float32) {
	C.Indent(C.float(indent_w))
}

func IsItemFocused() bool {
	return C.IsItemFocused() != C.bool(true)
}

func (self ImDrawList) PathLineTo(pos ImVec2) {
	C.ImDrawList_PathLineTo(self.handle(), pos.ToC())
}

func BeginTable(str_id string, column int32, flags ImGuiTableFlags, outer_size ImVec2, inner_width float32) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTable(str_idArg, C.int(column), C.ImGuiTableFlags(flags), outer_size.ToC(), C.float(inner_width)) != C.bool(true)
}

func (self ImDrawList) AddBezierQuadratic(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.ImDrawList_AddBezierQuadratic(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), C.uint(col), C.float(thickness), C.int(num_segments))
}

func DestroyContext(ctx ImGuiContext) {
	C.DestroyContext(ctx.handle())
}

func PushAllowKeyboardFocus(allow_keyboard_focus bool) {
	allow_keyboard_focusArg := C.bool(allow_keyboard_focus)

	C.PushAllowKeyboardFocus(allow_keyboard_focusArg)
}

func SliderInt(label string, v *int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawData) ScaleClipRects(fb_scale ImVec2) {
	C.ImDrawData_ScaleClipRects(self.handle(), fb_scale.ToC())
}

func (self ImFontAtlas) Clear() {
	C.ImFontAtlas_Clear(self.handle())
}

func (self ImFontAtlas) GetGlyphRangesDefault() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesDefault(self.handle()))
}

func SetNextItemOpen(is_open bool, cond ImGuiCond) {
	is_openArg := C.bool(is_open)

	C.SetNextItemOpen(is_openArg, C.ImGuiCond(cond))
}

func (self ImDrawList) AddTriangle(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32) {
	C.ImDrawList_AddTriangle(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), C.uint(col), C.float(thickness))
}

func ImGuiViewport_GetCenter(pOut *ImVec2, self ImGuiViewport) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImGuiViewport_GetCenter(pOutArg, self.handle())
}

func (self ImFontGlyphRangesBuilder) AddChar(c ImWchar) {
	C.ImFontGlyphRangesBuilder_AddChar(self.handle(), C.ImWchar(c))
}

func IsItemActive() bool {
	return C.IsItemActive() != C.bool(true)
}

func StyleColorsLight(dst ImGuiStyle) {
	C.StyleColorsLight(dst.handle())
}

func (self ImFontAtlas) GetGlyphRangesChineseSimplifiedCommon() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self.handle()))
}

func SetNextWindowFocus() {
	C.SetNextWindowFocus()
}

func DragInt4(label string, v *[4]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt4(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndTabBar() {
	C.EndTabBar()
}

func TableGetColumnIndex() int {
	return int(C.TableGetColumnIndex())
}

func GetVersion() string {
	return C.GoString(C.GetVersion())
}

func (self ImDrawList) AddRectFilledMultiColor(p_min ImVec2, p_max ImVec2, col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	C.ImDrawList_AddRectFilledMultiColor(self.handle(), p_min.ToC(), p_max.ToC(), C.uint(col_upr_left), C.uint(col_upr_right), C.uint(col_bot_right), C.uint(col_bot_left))
}

func SetCursorPosX(local_x float32) {
	C.SetCursorPosX(C.float(local_x))
}

func (self ImGuiInputTextCallbackData) SelectAll() {
	C.ImGuiInputTextCallbackData_SelectAll(self.handle())
}

func (self ImGuiStorage) GetInt(key ImGuiID, default_val int32) int {
	return int(C.ImGuiStorage_GetInt(self.handle(), C.ImGuiID(key), C.int(default_val)))
}

func ImDrawList_GetClipRectMin(pOut *ImVec2, self ImDrawList) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImDrawList_GetClipRectMin(pOutArg, self.handle())
}

func SameLine(offset_from_start_x float32, spacing float32) {
	C.SameLine(C.float(offset_from_start_x), C.float(spacing))
}

func SetClipboardText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.SetClipboardText(textArg)
}

func GetScrollMaxX() float32 {
	return float32(C.GetScrollMaxX())
}

func SliderFloat2(label string, v *[2]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat2(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawList) AddRectFilled(p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.ImDrawList_AddRectFilled(self.handle(), p_min.ToC(), p_max.ToC(), C.uint(col), C.float(rounding), C.ImDrawFlags(flags))
}

func InputFloat2(label string, v *[2]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat2(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func SliderFloat4(label string, v *[4]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat4(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsWindowFocused(flags ImGuiFocusedFlags) bool {
	return C.IsWindowFocused(C.ImGuiFocusedFlags(flags)) != C.bool(true)
}

func (self ImDrawList) PathBezierQuadraticCurveTo(p2 ImVec2, p3 ImVec2, num_segments int32) {
	C.ImDrawList_PathBezierQuadraticCurveTo(self.handle(), p2.ToC(), p3.ToC(), C.int(num_segments))
}

func IsWindowCollapsed() bool {
	return C.IsWindowCollapsed() != C.bool(true)
}

func TableNextColumn() bool {
	return C.TableNextColumn() != C.bool(true)
}

func (self ImFontAtlas) GetGlyphRangesJapanese() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesJapanese(self.handle()))
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

func (self ImDrawList) PushClipRectFullScreen() {
	C.ImDrawList_PushClipRectFullScreen(self.handle())
}

func (self ImFontAtlas) AddCustomRectRegular(width int32, height int32) int {
	return int(C.ImFontAtlas_AddCustomRectRegular(self.handle(), C.int(width), C.int(height)))
}

func BeginMainMenuBar() bool {
	return C.BeginMainMenuBar() != C.bool(true)
}

func GetWindowDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetWindowDrawList()))
}

func BeginMenuBar() bool {
	return C.BeginMenuBar() != C.bool(true)
}

func SmallButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.SmallButton(labelArg) != C.bool(true)
}

func (self ImGuiInputTextCallbackData) DeleteChars(pos int32, bytes_count int32) {
	C.ImGuiInputTextCallbackData_DeleteChars(self.handle(), C.int(pos), C.int(bytes_count))
}

func InvisibleButton(str_id string, size ImVec2, flags ImGuiButtonFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.InvisibleButton(str_idArg, size.ToC(), C.ImGuiButtonFlags(flags)) != C.bool(true)
}

func IsMousePosValid(mouse_pos *ImVec2) bool {
	mouse_posArg, mouse_posFin := mouse_pos.wrap()
	defer mouse_posFin()

	return C.IsMousePosValid(mouse_posArg) != C.bool(true)
}

func SliderInt4(label string, v *[4]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt4(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetNextWindowDockID(dock_id ImGuiID, cond ImGuiCond) {
	C.SetNextWindowDockID(C.ImGuiID(dock_id), C.ImGuiCond(cond))
}

func (self ImGuiStorage) SetFloat(key ImGuiID, val float32) {
	C.ImGuiStorage_SetFloat(self.handle(), C.ImGuiID(key), C.float(val))
}

func BeginPopupContextVoid(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextVoid(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func (self ImFontAtlas) Build() bool {
	return C.ImFontAtlas_Build(self.handle()) != C.bool(true)
}

func (self ImFont) IsLoaded() bool {
	return C.ImFont_IsLoaded(self.handle()) != C.bool(true)
}

func (self ImFont) GetDebugName() string {
	return C.GoString(C.ImFont_GetDebugName(self.handle()))
}

func IsKeyPressed(key ImGuiKey, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsKeyPressed(C.ImGuiKey(key), repeatArg) != C.bool(true)
}

func PopID() {
	C.PopID()
}

func (self ImFontGlyphRangesBuilder) AddRanges(ranges *ImWchar) {
	C.ImFontGlyphRangesBuilder_AddRanges(self.handle(), (*C.ImWchar)(ranges))
}

func (self ImFont) FindGlyphNoFallback(c ImWchar) ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.ImFont_FindGlyphNoFallback(self.handle(), C.ImWchar(c))))
}

func (self ImGuiIO) AddMouseViewportEvent(id ImGuiID) {
	C.ImGuiIO_AddMouseViewportEvent(self.handle(), C.ImGuiID(id))
}

func TextUnformatted(text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.TextUnformatted(textArg, text_endArg)
}

func UpdatePlatformWindows() {
	C.UpdatePlatformWindows()
}

func (self ImDrawList) PathFillConvex(col uint32) {
	C.ImDrawList_PathFillConvex(self.handle(), C.uint(col))
}

func Checkbox(label string, v *bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapBool(v)
	defer vFin()

	return C.Checkbox(labelArg, vArg) != C.bool(true)
}

func (self ImFontAtlas) GetGlyphRangesVietnamese() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesVietnamese(self.handle()))
}

func GetMousePos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMousePos(pOutArg)
}

func (self ImGuiTextFilter) Build() {
	C.ImGuiTextFilter_Build(self.handle())
}

func DebugCheckVersionAndDataLayout(version_str string, sz_io uint64, sz_style uint64, sz_vec2 uint64, sz_vec4 uint64, sz_drawvert uint64, sz_drawidx uint64) bool {
	version_strArg, version_strFin := wrapString(version_str)
	defer version_strFin()

	return C.DebugCheckVersionAndDataLayout(version_strArg, C.ulong(sz_io), C.ulong(sz_style), C.ulong(sz_vec2), C.ulong(sz_vec4), C.ulong(sz_drawvert), C.ulong(sz_drawidx)) != C.bool(true)
}

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

	return C.DragFloatRange2(labelArg, v_current_minArg, v_current_maxArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, format_maxArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawList) PathArcToFast(center ImVec2, radius float32, a_min_of_12 int32, a_max_of_12 int32) {
	C.ImDrawList_PathArcToFast(self.handle(), center.ToC(), C.float(radius), C.int(a_min_of_12), C.int(a_max_of_12))
}

func (self ImFont) FindGlyph(c ImWchar) ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.ImFont_FindGlyph(self.handle(), C.ImWchar(c))))
}

func IsAnyItemActive() bool {
	return C.IsAnyItemActive() != C.bool(true)
}

func BeginDisabled(disabled bool) {
	disabledArg := C.bool(disabled)

	C.BeginDisabled(disabledArg)
}

func SetColumnWidth(column_index int32, width float32) {
	C.SetColumnWidth(C.int(column_index), C.float(width))
}

func SetCurrentContext(ctx ImGuiContext) {
	C.SetCurrentContext(ctx.handle())
}

func (self ImFont) RenderText(draw_list ImDrawList, size float32, pos ImVec2, col uint32, clip_rect ImVec4, text_begin string, text_end string, wrap_width float32, cpu_fine_clip bool) {
	text_beginArg, text_beginFin := wrapString(text_begin)
	defer text_beginFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	cpu_fine_clipArg := C.bool(cpu_fine_clip)

	C.ImFont_RenderText(self.handle(), draw_list.handle(), C.float(size), pos.ToC(), C.uint(col), clip_rect.ToC(), text_beginArg, text_endArg, C.float(wrap_width), cpu_fine_clipArg)
}

func IsItemVisible() bool {
	return C.IsItemVisible() != C.bool(true)
}

func BeginChildFrame(id ImGuiID, size ImVec2, flags ImGuiWindowFlags) bool {
	return C.BeginChildFrame(C.ImGuiID(id), size.ToC(), C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func ColorConvertU32ToFloat4(pOut *ImVec4, in uint32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ColorConvertU32ToFloat4(pOutArg, C.uint(in))
}

func GetKeyPressedAmount(key ImGuiKey, repeat_delay float32, rate float32) int {
	return int(C.GetKeyPressedAmount(C.ImGuiKey(key), C.float(repeat_delay), C.float(rate)))
}

func GetFontSize() float32 {
	return float32(C.GetFontSize())
}

func (self ImFont) BuildLookupTable() {
	C.ImFont_BuildLookupTable(self.handle())
}

func ShowAboutWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowAboutWindow(p_openArg)
}

func (self ImDrawListSplitter) Split(draw_list ImDrawList, count int32) {
	C.ImDrawListSplitter_Split(self.handle(), draw_list.handle(), C.int(count))
}

func (self ImDrawList) AddDrawCmd() {
	C.ImDrawList_AddDrawCmd(self.handle())
}

func (self ImDrawList) AddNgon(center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.ImDrawList_AddNgon(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments), C.float(thickness))
}

func Button(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Button(labelArg, size.ToC()) != C.bool(true)
}

func (self ImGuiStyle) ScaleAllSizes(scale_factor float32) {
	C.ImGuiStyle_ScaleAllSizes(self.handle(), C.float(scale_factor))
}

func LoadIniSettingsFromMemory(ini_data string, ini_size uint64) {
	ini_dataArg, ini_dataFin := wrapString(ini_data)
	defer ini_dataFin()

	C.LoadIniSettingsFromMemory(ini_dataArg, C.ulong(ini_size))
}

func BeginPopupModal(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginPopupModal(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func GetWindowContentRegionMin(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowContentRegionMin(pOutArg)
}

func GetCursorPosX() float32 {
	return float32(C.GetCursorPosX())
}

func (self ImFontAtlas) GetGlyphRangesKorean() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesKorean(self.handle()))
}

func ColorPicker3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorPicker3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func GetColumnsCount() int {
	return int(C.GetColumnsCount())
}

func DragInt(label string, v *int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImGuiIO) AddMouseButtonEvent(button int32, down bool) {
	downArg := C.bool(down)

	C.ImGuiIO_AddMouseButtonEvent(self.handle(), C.int(button), downArg)
}

func ColorButton(desc_id string, col ImVec4, flags ImGuiColorEditFlags, size ImVec2) bool {
	desc_idArg, desc_idFin := wrapString(desc_id)
	defer desc_idFin()

	return C.ColorButton(desc_idArg, col.ToC(), C.ImGuiColorEditFlags(flags), size.ToC()) != C.bool(true)
}

func IsItemActivated() bool {
	return C.IsItemActivated() != C.bool(true)
}

func DragFloat3(label string, v *[3]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat3(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	want_capture_keyboardArg := C.bool(want_capture_keyboard)

	C.SetNextFrameWantCaptureKeyboard(want_capture_keyboardArg)
}

func Render() {
	C.Render()
}

func SetTabItemClosed(tab_or_docked_window_label string) {
	tab_or_docked_window_labelArg, tab_or_docked_window_labelFin := wrapString(tab_or_docked_window_label)
	defer tab_or_docked_window_labelFin()

	C.SetTabItemClosed(tab_or_docked_window_labelArg)
}

func (self ImDrawList) AddTriangleFilled(p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32) {
	C.ImDrawList_AddTriangleFilled(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), C.uint(col))
}

func (self ImFontAtlas) ClearFonts() {
	C.ImFontAtlas_ClearFonts(self.handle())
}

func PushTextWrapPos(wrap_local_pos_x float32) {
	C.PushTextWrapPos(C.float(wrap_local_pos_x))
}

func GetFrameHeight() float32 {
	return float32(C.GetFrameHeight())
}

func GetStateStorage() ImGuiStorage {
	return (ImGuiStorage)(unsafe.Pointer(C.GetStateStorage()))
}

func (self ImGuiStorage) Clear() {
	C.ImGuiStorage_Clear(self.handle())
}

func DragFloat(label string, v *float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DragScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalar(labelArg, C.ImGuiDataType(data_type), p_data, C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawListSplitter) Clear() {
	C.ImDrawListSplitter_Clear(self.handle())
}

func (self ImGuiIO) ClearInputKeys() {
	C.ImGuiIO_ClearInputKeys(self.handle())
}

func GetTextLineHeightWithSpacing() float32 {
	return float32(C.GetTextLineHeightWithSpacing())
}

func InputScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func SetColumnOffset(column_index int32, offset_x float32) {
	C.SetColumnOffset(C.int(column_index), C.float(offset_x))
}

func IsMouseHoveringRect(r_min ImVec2, r_max ImVec2, clip bool) bool {
	clipArg := C.bool(clip)

	return C.IsMouseHoveringRect(r_min.ToC(), r_max.ToC(), clipArg) != C.bool(true)
}

func (self ImDrawList) AddImageRounded(user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.ImDrawList_AddImageRounded(self.handle(), C.ImTextureID(user_texture_id), p_min.ToC(), p_max.ToC(), uv_min.ToC(), uv_max.ToC(), C.uint(col), C.float(rounding), C.ImDrawFlags(flags))
}

func GetColumnIndex() int {
	return int(C.GetColumnIndex())
}

func (self ImDrawList) PushTextureID(texture_id ImTextureID) {
	C.ImDrawList_PushTextureID(self.handle(), C.ImTextureID(texture_id))
}

func (self ImFont) RenderChar(draw_list ImDrawList, size float32, pos ImVec2, col uint32, c ImWchar) {
	C.ImFont_RenderChar(self.handle(), draw_list.handle(), C.float(size), pos.ToC(), C.uint(col), C.ImWchar(c))
}

func BeginCombo(label string, preview_value string, flags ImGuiComboFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	preview_valueArg, preview_valueFin := wrapString(preview_value)
	defer preview_valueFin()

	return C.BeginCombo(labelArg, preview_valueArg, C.ImGuiComboFlags(flags)) != C.bool(true)
}

func GetItemRectMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectMax(pOutArg)
}

func InputFloat(label string, v *float32, step float32, step_fast float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat(labelArg, vArg, C.float(step), C.float(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func CloseCurrentPopup() {
	C.CloseCurrentPopup()
}

func IsMouseClicked(button ImGuiMouseButton, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsMouseClicked(C.ImGuiMouseButton(button), repeatArg) != C.bool(true)
}

func ShowDemoWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDemoWindow(p_openArg)
}

func ShowMetricsWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowMetricsWindow(p_openArg)
}

func End() {
	C.End()
}

func (self ImGuiListClipper) ForceDisplayRangeByIndices(item_min int32, item_max int32) {
	C.ImGuiListClipper_ForceDisplayRangeByIndices(self.handle(), C.int(item_min), C.int(item_max))
}

func PushButtonRepeat(repeat bool) {
	repeatArg := C.bool(repeat)

	C.PushButtonRepeat(repeatArg)
}

func IsMouseReleased(button ImGuiMouseButton) bool {
	return C.IsMouseReleased(C.ImGuiMouseButton(button)) != C.bool(true)
}

func (self ImDrawList) AddBezierCubic(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.ImDrawList_AddBezierCubic(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), C.uint(col), C.float(thickness), C.int(num_segments))
}

func VSliderScalar(label string, size ImVec2, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderScalar(labelArg, size.ToC(), C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func TableSetBgColor(target ImGuiTableBgTarget, color uint32, column_n int32) {
	C.TableSetBgColor(C.ImGuiTableBgTarget(target), C.uint(color), C.int(column_n))
}

func DebugTextEncoding(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.DebugTextEncoding(textArg)
}

func SetCursorScreenPos(pos ImVec2) {
	C.SetCursorScreenPos(pos.ToC())
}

func ColorEdit3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func (self ImGuiIO) AddMousePosEvent(x float32, y float32) {
	C.ImGuiIO_AddMousePosEvent(self.handle(), C.float(x), C.float(y))
}

func GetColumnWidth(column_index int32) float32 {
	return float32(C.GetColumnWidth(C.int(column_index)))
}

func (self ImGuiStorage) SetAllInt(val int32) {
	C.ImGuiStorage_SetAllInt(self.handle(), C.int(val))
}

func PopItemWidth() {
	C.PopItemWidth()
}

func (self ImFontAtlas) ClearTexData() {
	C.ImFontAtlas_ClearTexData(self.handle())
}

func ShowStyleEditor(ref ImGuiStyle) {
	C.ShowStyleEditor(ref.handle())
}

func (self ImDrawList) PushClipRect(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

	C.ImDrawList_PushClipRect(self.handle(), clip_rect_min.ToC(), clip_rect_max.ToC(), intersect_with_current_clip_rectArg)
}

func GetWindowContentRegionMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowContentRegionMax(pOutArg)
}

func IsItemDeactivated() bool {
	return C.IsItemDeactivated() != C.bool(true)
}

func (self ImFontAtlas) GetCustomRectByIndex(index int32) ImFontAtlasCustomRect {
	return (ImFontAtlasCustomRect)(unsafe.Pointer(C.ImFontAtlas_GetCustomRectByIndex(self.handle(), C.int(index))))
}

func SetColorEditOptions(flags ImGuiColorEditFlags) {
	C.SetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

func CreateContext(shared_font_atlas ImFontAtlas) ImGuiContext {
	return (ImGuiContext)(unsafe.Pointer(C.CreateContext(shared_font_atlas.handle())))
}

func TableNextRow(row_flags ImGuiTableRowFlags, min_row_height float32) {
	C.TableNextRow(C.ImGuiTableRowFlags(row_flags), C.float(min_row_height))
}

func LogToFile(auto_open_depth int32, filename string) {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	C.LogToFile(C.int(auto_open_depth), filenameArg)
}

func PopFont() {
	C.PopFont()
}

func TableSetupScrollFreeze(cols int32, rows int32) {
	C.TableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func GetMouseCursor() ImGuiMouseCursor {
	return ImGuiMouseCursor(C.GetMouseCursor())
}

func PopStyleVar(count int32) {
	C.PopStyleVar(C.int(count))
}

func (self ImGuiIO) ClearInputCharacters() {
	C.ImGuiIO_ClearInputCharacters(self.handle())
}

func InputScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func (self ImFontAtlas) AddFont(font_cfg ImFontConfig) ImFont {
	return (ImFont)(unsafe.Pointer(C.ImFontAtlas_AddFont(self.handle(), font_cfg.handle())))
}

func IsWindowDocked() bool {
	return C.IsWindowDocked() != C.bool(true)
}

func (self ImGuiPayload) Clear() {
	C.ImGuiPayload_Clear(self.handle())
}

func CalcItemWidth() float32 {
	return float32(C.CalcItemWidth())
}

func GetItemRectMin(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectMin(pOutArg)
}

func (self ImGuiTextFilter) Clear() {
	C.ImGuiTextFilter_Clear(self.handle())
}

func IsKeyReleased(key ImGuiKey) bool {
	return C.IsKeyReleased(C.ImGuiKey(key)) != C.bool(true)
}

func StyleColorsDark(dst ImGuiStyle) {
	C.StyleColorsDark(dst.handle())
}

func (self ImDrawList) AddNgonFilled(center ImVec2, radius float32, col uint32, num_segments int32) {
	C.ImDrawList_AddNgonFilled(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments))
}

func EndTable() {
	C.EndTable()
}

func (self ImFontAtlas) AddFontFromMemoryTTF(font_data unsafe.Pointer, font_size int32, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	return (ImFont)(unsafe.Pointer(C.ImFontAtlas_AddFontFromMemoryTTF(self.handle(), font_data, C.int(font_size), C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func (self ImFont) GrowIndex(new_size int32) {
	C.ImFont_GrowIndex(self.handle(), C.int(new_size))
}

func (self ImDrawList) PathLineToMergeDuplicate(pos ImVec2) {
	C.ImDrawList_PathLineToMergeDuplicate(self.handle(), pos.ToC())
}

func BeginTabBar(str_id string, flags ImGuiTabBarFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTabBar(str_idArg, C.ImGuiTabBarFlags(flags)) != C.bool(true)
}

func (self ImDrawList) PathBezierCubicCurveTo(p2 ImVec2, p3 ImVec2, p4 ImVec2, num_segments int32) {
	C.ImDrawList_PathBezierCubicCurveTo(self.handle(), p2.ToC(), p3.ToC(), p4.ToC(), C.int(num_segments))
}

func SetMouseCursor(cursor_type ImGuiMouseCursor) {
	C.SetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

func (self ImDrawList) ChannelsSplit(count int32) {
	C.ImDrawList_ChannelsSplit(self.handle(), C.int(count))
}

func AlignTextToFramePadding() {
	C.AlignTextToFramePadding()
}

func TableGetColumnCount() int {
	return int(C.TableGetColumnCount())
}

func (self ImGuiIO) SetKeyEventNativeData(key ImGuiKey, native_keycode int32, native_scancode int32, native_legacy_index int32) {
	C.ImGuiIO_SetKeyEventNativeData(self.handle(), C.ImGuiKey(key), C.int(native_keycode), C.int(native_scancode), C.int(native_legacy_index))
}

func FindViewportByID(id ImGuiID) ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.FindViewportByID(C.ImGuiID(id))))
}

func DragScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetWindowHeight() float32 {
	return float32(C.GetWindowHeight())
}

func Unindent(indent_w float32) {
	C.Unindent(C.float(indent_w))
}

func (self ImDrawList) CloneOutput() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.ImDrawList_CloneOutput(self.handle())))
}

func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	want_capture_mouseArg := C.bool(want_capture_mouse)

	C.SetNextFrameWantCaptureMouse(want_capture_mouseArg)
}

func GetScrollMaxY() float32 {
	return float32(C.GetScrollMaxY())
}

func (self ImDrawListSplitter) SetCurrentChannel(draw_list ImDrawList, channel_idx int32) {
	C.ImDrawListSplitter_SetCurrentChannel(self.handle(), draw_list.handle(), C.int(channel_idx))
}

func DragFloat4(label string, v *[4]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat4(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetContentRegionAvail(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetContentRegionAvail(pOutArg)
}

func SaveIniSettingsToMemory(out_ini_size *uint64) string {
	return C.GoString(C.SaveIniSettingsToMemory((*C.ulong)(out_ini_size)))
}

func SetItemDefaultFocus() {
	C.SetItemDefaultFocus()
}

func SliderFloat(label string, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawList) PrimRectUV(a ImVec2, b ImVec2, uv_a ImVec2, uv_b ImVec2, col uint32) {
	C.ImDrawList_PrimRectUV(self.handle(), a.ToC(), b.ToC(), uv_a.ToC(), uv_b.ToC(), C.uint(col))
}

func Begin(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.Begin(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func InputFloat4(label string, v *[4]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat4(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetWindowSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowSize(pOutArg)
}

func (self ImGuiIO) AddKeyAnalogEvent(key ImGuiKey, down bool, v float32) {
	downArg := C.bool(down)

	C.ImGuiIO_AddKeyAnalogEvent(self.handle(), C.ImGuiKey(key), downArg, C.float(v))
}

func ImColor_SetHSV(self *ImColor, h float32, s float32, v float32, a float32) {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.ImColor_SetHSV(selfArg, C.float(h), C.float(s), C.float(v), C.float(a))
}

func SetNextWindowPos(pos ImVec2, cond ImGuiCond, pivot ImVec2) {
	C.SetNextWindowPos(pos.ToC(), C.ImGuiCond(cond), pivot.ToC())
}

func SetScrollHereX(center_x_ratio float32) {
	C.SetScrollHereX(C.float(center_x_ratio))
}

func SetNextWindowClass(window_class ImGuiWindowClass) {
	C.SetNextWindowClass(window_class.handle())
}

func GetDrawData() ImDrawData {
	return (ImDrawData)(unsafe.Pointer(C.GetDrawData()))
}

func GetScrollX() float32 {
	return float32(C.GetScrollX())
}

func (self ImDrawData) Clear() {
	C.ImDrawData_Clear(self.handle())
}

func (self ImDrawList) AddCircle(center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.ImDrawList_AddCircle(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments), C.float(thickness))
}

func (self ImGuiIO) SetAppAcceptingEvents(accepting_events bool) {
	accepting_eventsArg := C.bool(accepting_events)

	C.ImGuiIO_SetAppAcceptingEvents(self.handle(), accepting_eventsArg)
}

func GetStyle() ImGuiStyle {
	return (ImGuiStyle)(unsafe.Pointer(C.GetStyle()))
}

func PopClipRect() {
	C.PopClipRect()
}

func (self ImDrawList) PrimWriteVtx(pos ImVec2, uv ImVec2, col uint32) {
	C.ImDrawList_PrimWriteVtx(self.handle(), pos.ToC(), uv.ToC(), C.uint(col))
}

func (self ImGuiInputTextCallbackData) InsertChars(pos int32, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.ImGuiInputTextCallbackData_InsertChars(self.handle(), C.int(pos), textArg, text_endArg)
}

func GetPlatformIO() ImGuiPlatformIO {
	return (ImGuiPlatformIO)(unsafe.Pointer(C.GetPlatformIO()))
}

func LogToClipboard(auto_open_depth int32) {
	C.LogToClipboard(C.int(auto_open_depth))
}

func SetItemAllowOverlap() {
	C.SetItemAllowOverlap()
}

func TableGetRowIndex() int {
	return int(C.TableGetRowIndex())
}

func (self ImDrawList) PathRect(rect_min ImVec2, rect_max ImVec2, rounding float32, flags ImDrawFlags) {
	C.ImDrawList_PathRect(self.handle(), rect_min.ToC(), rect_max.ToC(), C.float(rounding), C.ImDrawFlags(flags))
}

func (self ImGuiIO) AddFocusEvent(focused bool) {
	focusedArg := C.bool(focused)

	C.ImGuiIO_AddFocusEvent(self.handle(), focusedArg)
}

func (self ImGuiListClipper) Step() bool {
	return C.ImGuiListClipper_Step(self.handle()) != C.bool(true)
}

func TabItemButton(label string, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TabItemButton(labelArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func (self ImDrawList) AddPolyline(points *ImVec2, num_points int32, col uint32, flags ImDrawFlags, thickness float32) {
	pointsArg, pointsFin := points.wrap()
	defer pointsFin()

	C.ImDrawList_AddPolyline(self.handle(), pointsArg, C.int(num_points), C.uint(col), C.ImDrawFlags(flags), C.float(thickness))
}

func (self ImDrawList) PathArcTo(center ImVec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	C.ImDrawList_PathArcTo(self.handle(), center.ToC(), C.float(radius), C.float(a_min), C.float(a_max), C.int(num_segments))
}

func GetDrawListSharedData() ImDrawListSharedData {
	return (ImDrawListSharedData)(unsafe.Pointer(C.GetDrawListSharedData()))
}

func EndListBox() {
	C.EndListBox()
}

func EndPopup() {
	C.EndPopup()
}

func GetFrameHeightWithSpacing() float32 {
	return float32(C.GetFrameHeightWithSpacing())
}

func SliderAngle(label string, v_rad *float32, v_degrees_min float32, v_degrees_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_radArg, v_radFin := wrapFloat(v_rad)
	defer v_radFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderAngle(labelArg, v_radArg, C.float(v_degrees_min), C.float(v_degrees_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetKeyboardFocusHere(offset int32) {
	C.SetKeyboardFocusHere(C.int(offset))
}

func (self ImGuiPayload) IsPreview() bool {
	return C.ImGuiPayload_IsPreview(self.handle()) != C.bool(true)
}

func BeginPopup(str_id string, flags ImGuiWindowFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopup(str_idArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func GetFrameCount() int {
	return int(C.GetFrameCount())
}

func GetTreeNodeToLabelSpacing() float32 {
	return float32(C.GetTreeNodeToLabelSpacing())
}

func LogFinish() {
	C.LogFinish()
}

func (self ImDrawList) PrimRect(a ImVec2, b ImVec2, col uint32) {
	C.ImDrawList_PrimRect(self.handle(), a.ToC(), b.ToC(), C.uint(col))
}

func (self ImGuiListClipper) Begin(items_count int32, items_height float32) {
	C.ImGuiListClipper_Begin(self.handle(), C.int(items_count), C.float(items_height))
}

func EndDisabled() {
	C.EndDisabled()
}

func (self ImFontAtlas) AddFontDefault(font_cfg ImFontConfig) ImFont {
	return (ImFont)(unsafe.Pointer(C.ImFontAtlas_AddFontDefault(self.handle(), font_cfg.handle())))
}

func SetNextWindowContentSize(size ImVec2) {
	C.SetNextWindowContentSize(size.ToC())
}

func (self ImDrawList) ChannelsSetCurrent(n int32) {
	C.ImDrawList_ChannelsSetCurrent(self.handle(), C.int(n))
}

func IsKeyDown(key ImGuiKey) bool {
	return C.IsKeyDown(C.ImGuiKey(key)) != C.bool(true)
}

func (self ImGuiIO) AddMouseWheelEvent(wh_x float32, wh_y float32) {
	C.ImGuiIO_AddMouseWheelEvent(self.handle(), C.float(wh_x), C.float(wh_y))
}

func BeginTabItem(label string, p_open *bool, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginTabItem(labelArg, p_openArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func (self ImGuiTextFilter) IsActive() bool {
	return C.ImGuiTextFilter_IsActive(self.handle()) != C.bool(true)
}

func (self ImFontAtlas) GetGlyphRangesCyrillic() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesCyrillic(self.handle()))
}

func EndTabItem() {
	C.EndTabItem()
}

func (self ImGuiPayload) IsDelivery() bool {
	return C.ImGuiPayload_IsDelivery(self.handle()) != C.bool(true)
}

func BeginMenu(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	enabledArg := C.bool(enabled)

	return C.BeginMenu(labelArg, enabledArg) != C.bool(true)
}

func (self ImGuiStorage) SetVoidPtr(key ImGuiID, val unsafe.Pointer) {
	C.ImGuiStorage_SetVoidPtr(self.handle(), C.ImGuiID(key), val)
}

func EndChildFrame() {
	C.EndChildFrame()
}

func ResetMouseDragDelta(button ImGuiMouseButton) {
	C.ResetMouseDragDelta(C.ImGuiMouseButton(button))
}

func (self ImFontAtlas) AddFontFromMemoryCompressedBase85TTF(compressed_font_data_base85 string, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	compressed_font_data_base85Arg, compressed_font_data_base85Fin := wrapString(compressed_font_data_base85)
	defer compressed_font_data_base85Fin()

	return (ImFont)(unsafe.Pointer(C.ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self.handle(), compressed_font_data_base85Arg, C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func (self ImFont) CalcWordWrapPositionA(scale float32, text string, text_end string, wrap_width float32) string {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.GoString(C.ImFont_CalcWordWrapPositionA(self.handle(), C.float(scale), textArg, text_endArg, C.float(wrap_width)))
}

func (self ImGuiTextBuffer) c_str() string {
	return C.GoString(C.ImGuiTextBuffer_c_str(self.handle()))
}

func ImGuiViewport_GetWorkCenter(pOut *ImVec2, self ImGuiViewport) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImGuiViewport_GetWorkCenter(pOutArg, self.handle())
}

func GetWindowDpiScale() float32 {
	return float32(C.GetWindowDpiScale())
}

func (self ImDrawList) PathStroke(col uint32, flags ImDrawFlags, thickness float32) {
	C.ImDrawList_PathStroke(self.handle(), C.uint(col), C.ImDrawFlags(flags), C.float(thickness))
}

func SetNextWindowBgAlpha(alpha float32) {
	C.SetNextWindowBgAlpha(C.float(alpha))
}

func (self ImDrawList) AddCircleFilled(center ImVec2, radius float32, col uint32, num_segments int32) {
	C.ImDrawList_AddCircleFilled(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments))
}

func (self ImGuiListClipper) End() {
	C.ImGuiListClipper_End(self.handle())
}

func GetMouseClickedCount(button ImGuiMouseButton) int {
	return int(C.GetMouseClickedCount(C.ImGuiMouseButton(button)))
}

func (self ImGuiStorage) SetBool(key ImGuiID, val bool) {
	valArg := C.bool(val)

	C.ImGuiStorage_SetBool(self.handle(), C.ImGuiID(key), valArg)
}

func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32) bool {
	return C.IsMouseDragging(C.ImGuiMouseButton(button), C.float(lock_threshold)) != C.bool(true)
}

func SliderFloat3(label string, v *[3]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat3(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImDrawList) AddQuadFilled(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32) {
	C.ImDrawList_AddQuadFilled(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), C.uint(col))
}

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

	return C.DragIntRange2(labelArg, v_current_minArg, v_current_maxArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, format_maxArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsAnyItemFocused() bool {
	return C.IsAnyItemFocused() != C.bool(true)
}

func (self ImDrawList) AddRect(p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags, thickness float32) {
	C.ImDrawList_AddRect(self.handle(), p_min.ToC(), p_max.ToC(), C.uint(col), C.float(rounding), C.ImDrawFlags(flags), C.float(thickness))
}

func (self ImDrawList) PathClear() {
	C.ImDrawList_PathClear(self.handle())
}

func SliderInt2(label string, v *[2]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt2(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func StyleColorsClassic(dst ImGuiStyle) {
	C.StyleColorsClassic(dst.handle())
}

func (self ImFontAtlas) GetGlyphRangesChineseFull() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesChineseFull(self.handle()))
}

func PopButtonRepeat() {
	C.PopButtonRepeat()
}

func EndChild() {
	C.EndChild()
}

func VSliderInt(label string, size ImVec2, v *int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderInt(labelArg, size.ToC(), vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImFontAtlas) AddFontFromFileTTF(filename string, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	return (ImFont)(unsafe.Pointer(C.ImFontAtlas_AddFontFromFileTTF(self.handle(), filenameArg, C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func BeginGroup() {
	C.BeginGroup()
}

func ImageButton(str_id string, user_texture_id ImTextureID, size ImVec2, uv0 ImVec2, uv1 ImVec2, bg_col ImVec4, tint_col ImVec4) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ImageButton(str_idArg, C.ImTextureID(user_texture_id), size.ToC(), uv0.ToC(), uv1.ToC(), bg_col.ToC(), tint_col.ToC()) != C.bool(true)
}

func TableSetupColumn(label string, flags ImGuiTableColumnFlags, init_width_or_weight float32, user_id ImGuiID) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableSetupColumn(labelArg, C.ImGuiTableColumnFlags(flags), C.float(init_width_or_weight), C.ImGuiID(user_id))
}

func PushItemWidth(item_width float32) {
	C.PushItemWidth(C.float(item_width))
}

func AcceptDragDropPayload(typeArg string, flags ImGuiDragDropFlags) ImGuiPayload {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return (ImGuiPayload)(unsafe.Pointer(C.AcceptDragDropPayload(typeArgArg, C.ImGuiDragDropFlags(flags))))
}

func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return C.IsMouseDoubleClicked(C.ImGuiMouseButton(button)) != C.bool(true)
}

func GetTextLineHeight() float32 {
	return float32(C.GetTextLineHeight())
}

func (self ImFontAtlas) GetGlyphRangesThai() *ImWchar {
	return (*ImWchar)(C.ImFontAtlas_GetGlyphRangesThai(self.handle()))
}

func TableGetSortSpecs() ImGuiTableSortSpecs {
	return (ImGuiTableSortSpecs)(unsafe.Pointer(C.TableGetSortSpecs()))
}

func TableSetColumnIndex(column_n int32) bool {
	return C.TableSetColumnIndex(C.int(column_n)) != C.bool(true)
}

func (self ImDrawListSplitter) Merge(draw_list ImDrawList) {
	C.ImDrawListSplitter_Merge(self.handle(), draw_list.handle())
}

func (self ImFontGlyphRangesBuilder) Clear() {
	C.ImFontGlyphRangesBuilder_Clear(self.handle())
}

func SliderScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImGuiTextFilter) PassFilter(text string, text_end string) bool {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.ImGuiTextFilter_PassFilter(self.handle(), textArg, text_endArg) != C.bool(true)
}

func (self ImGuiStorage) BuildSortByKey() {
	C.ImGuiStorage_BuildSortByKey(self.handle())
}

func BeginTooltip() {
	C.BeginTooltip()
}

func GetKeyIndex(key ImGuiKey) int {
	return int(C.GetKeyIndex(C.ImGuiKey(key)))
}

func (self ImFont) GetCharAdvance(c ImWchar) float32 {
	return float32(C.ImFont_GetCharAdvance(self.handle(), C.ImWchar(c)))
}

func GetIO() ImGuiIO {
	return (ImGuiIO)(unsafe.Pointer(C.GetIO()))
}

func SetNextWindowSize(size ImVec2, cond ImGuiCond) {
	C.SetNextWindowSize(size.ToC(), C.ImGuiCond(cond))
}

func SetNextWindowViewport(viewport_id ImGuiID) {
	C.SetNextWindowViewport(C.ImGuiID(viewport_id))
}

func ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ShowStyleSelector(labelArg) != C.bool(true)
}

func GetClipboardText() string {
	return C.GoString(C.GetClipboardText())
}

func CalcTextSize(pOut *ImVec2, text string, text_end string, hide_text_after_double_hash bool, wrap_width float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	hide_text_after_double_hashArg := C.bool(hide_text_after_double_hash)

	C.CalcTextSize(pOutArg, textArg, text_endArg, hide_text_after_double_hashArg, C.float(wrap_width))
}

func (self ImFontGlyphRangesBuilder) SetBit(n uint64) {
	C.ImFontGlyphRangesBuilder_SetBit(self.handle(), C.ulong(n))
}

func EndTooltip() {
	C.EndTooltip()
}

func GetCursorStartPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetCursorStartPos(pOutArg)
}

func NextColumn() {
	C.NextColumn()
}

func (self ImFontAtlas) AddFontFromMemoryCompressedTTF(compressed_font_data unsafe.Pointer, compressed_font_size int32, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	return (ImFont)(unsafe.Pointer(C.ImFontAtlas_AddFontFromMemoryCompressedTTF(self.handle(), compressed_font_data, C.int(compressed_font_size), C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func InputDouble(label string, v *float64, step float64, step_fast float64, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputDouble(labelArg, (*C.double)(v), C.double(step), C.double(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func TableGetColumnFlags(column_n int32) ImGuiTableColumnFlags {
	return ImGuiTableColumnFlags(C.TableGetColumnFlags(C.int(column_n)))
}

func (self ImGuiIO) AddInputCharactersUTF8(str string) {
	strArg, strFin := wrapString(str)
	defer strFin()

	C.ImGuiIO_AddInputCharactersUTF8(self.handle(), strArg)
}

func (self ImGuiTextFilter) Draw(label string, width float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ImGuiTextFilter_Draw(self.handle(), labelArg, C.float(width)) != C.bool(true)
}

func Image(user_texture_id ImTextureID, size ImVec2, uv0 ImVec2, uv1 ImVec2, tint_col ImVec4, border_col ImVec4) {
	C.Image(C.ImTextureID(user_texture_id), size.ToC(), uv0.ToC(), uv1.ToC(), tint_col.ToC(), border_col.ToC())
}

func Bullet() {
	C.Bullet()
}

func EndMenu() {
	C.EndMenu()
}

func (self ImFontAtlas) ClearInputData() {
	C.ImFontAtlas_ClearInputData(self.handle())
}

func (self ImGuiStorage) GetFloat(key ImGuiID, default_val float32) float32 {
	return float32(C.ImGuiStorage_GetFloat(self.handle(), C.ImGuiID(key), C.float(default_val)))
}

func GetFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.GetFont()))
}

func IsMouseDown(button ImGuiMouseButton) bool {
	return C.IsMouseDown(C.ImGuiMouseButton(button)) != C.bool(true)
}

func BeginPopupContextItem(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextItem(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func ImColor_HSV(pOut *ImColor, h float32, s float32, v float32, a float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImColor_HSV(pOutArg, C.float(h), C.float(s), C.float(v), C.float(a))
}

func EndCombo() {
	C.EndCombo()
}

func (self ImFontAtlas) CalcCustomRectUV(rect ImFontAtlasCustomRect, out_uv_min *ImVec2, out_uv_max *ImVec2) {
	out_uv_minArg, out_uv_minFin := out_uv_min.wrap()
	defer out_uv_minFin()

	out_uv_maxArg, out_uv_maxFin := out_uv_max.wrap()
	defer out_uv_maxFin()

	C.ImFontAtlas_CalcCustomRectUV(self.handle(), rect.handle(), out_uv_minArg, out_uv_maxArg)
}

func LogButtons() {
	C.LogButtons()
}

func GetMousePosOnOpeningCurrentPopup(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMousePosOnOpeningCurrentPopup(pOutArg)
}

func LoadIniSettingsFromDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.LoadIniSettingsFromDisk(ini_filenameArg)
}

func GetDragDropPayload() ImGuiPayload {
	return (ImGuiPayload)(unsafe.Pointer(C.GetDragDropPayload()))
}

func DragInt2(label string, v *[2]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt2(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndDragDropTarget() {
	C.EndDragDropTarget()
}

func EndMenuBar() {
	C.EndMenuBar()
}

func ProgressBar(fraction float32, size_arg ImVec2, overlay string) {
	overlayArg, overlayFin := wrapString(overlay)
	defer overlayFin()

	C.ProgressBar(C.float(fraction), size_arg.ToC(), overlayArg)
}

func SetScrollHereY(center_y_ratio float32) {
	C.SetScrollHereY(C.float(center_y_ratio))
}

func IsItemHovered(flags ImGuiHoveredFlags) bool {
	return C.IsItemHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func Dummy(size ImVec2) {
	C.Dummy(size.ToC())
}

func ShowUserGuide() {
	C.ShowUserGuide()
}

func InputInt(label string, v *int32, step int32, step_fast int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	return C.InputInt(labelArg, vArg, C.int(step), C.int(step_fast), C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func (self ImDrawList) PopClipRect() {
	C.ImDrawList_PopClipRect(self.handle())
}

func (self ImGuiInputTextCallbackData) HasSelection() bool {
	return C.ImGuiInputTextCallbackData_HasSelection(self.handle()) != C.bool(true)
}

func (self ImDrawList) AddLine(p1 ImVec2, p2 ImVec2, col uint32, thickness float32) {
	C.ImDrawList_AddLine(self.handle(), p1.ToC(), p2.ToC(), C.uint(col), C.float(thickness))
}

func GetScrollY() float32 {
	return float32(C.GetScrollY())
}

func IsItemDeactivatedAfterEdit() bool {
	return C.IsItemDeactivatedAfterEdit() != C.bool(true)
}

func (self ImDrawList) AddImageQuad(user_texture_id ImTextureID, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, uv1 ImVec2, uv2 ImVec2, uv3 ImVec2, uv4 ImVec2, col uint32) {
	C.ImDrawList_AddImageQuad(self.handle(), C.ImTextureID(user_texture_id), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), uv1.ToC(), uv2.ToC(), uv3.ToC(), uv4.ToC(), C.uint(col))
}

func (self ImGuiIO) AddInputCharacter(c uint32) {
	C.ImGuiIO_AddInputCharacter(self.handle(), C.uint(c))
}

func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
	return C.BeginDragDropSource(C.ImGuiDragDropFlags(flags)) != C.bool(true)
}

func ColorEdit4(label string, col *[4]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit4(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func IsItemClicked(mouse_button ImGuiMouseButton) bool {
	return C.IsItemClicked(C.ImGuiMouseButton(mouse_button)) != C.bool(true)
}

func GetColumnOffset(column_index int32) float32 {
	return float32(C.GetColumnOffset(C.int(column_index)))
}

func GetCursorPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetCursorPos(pOutArg)
}

func MemFree(ptr unsafe.Pointer) {
	C.MemFree(ptr)
}

func GetCurrentContext() ImGuiContext {
	return (ImGuiContext)(unsafe.Pointer(C.GetCurrentContext()))
}

func (self ImDrawList) AddQuad(p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32) {
	C.ImDrawList_AddQuad(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), C.uint(col), C.float(thickness))
}

func SetStateStorage(storage ImGuiStorage) {
	C.SetStateStorage(storage.handle())
}

func SetCursorPos(local_pos ImVec2) {
	C.SetCursorPos(local_pos.ToC())
}

func (self ImFontAtlas) IsBuilt() bool {
	return C.ImFontAtlas_IsBuilt(self.handle()) != C.bool(true)
}

func BeginDragDropTarget() bool {
	return C.BeginDragDropTarget() != C.bool(true)
}

func (self ImFontAtlasCustomRect) IsPacked() bool {
	return C.ImFontAtlasCustomRect_IsPacked(self.handle()) != C.bool(true)
}

func ShowStackToolWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowStackToolWindow(p_openArg)
}

func PopTextWrapPos() {
	C.PopTextWrapPos()
}

func SliderInt3(label string, v *[3]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt3(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func (self ImFontGlyphRangesBuilder) AddText(text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.ImFontGlyphRangesBuilder_AddText(self.handle(), textArg, text_endArg)
}

func InputFloat3(label string, v *[3]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat3(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func SetWindowFontScale(scale float32) {
	C.SetWindowFontScale(C.float(scale))
}

func GetCursorScreenPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetCursorScreenPos(pOutArg)
}

func (self ImDrawListSplitter) ClearFreeMemory() {
	C.ImDrawListSplitter_ClearFreeMemory(self.handle())
}

func PushFont(font ImFont) {
	C.PushFont(font.handle())
}

func IsWindowAppearing() bool {
	return C.IsWindowAppearing() != C.bool(true)
}

func (self ImGuiPayload) IsDataType(typeArg string) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.ImGuiPayload_IsDataType(self.handle(), typeArgArg) != C.bool(true)
}

func ArrowButton(str_id string, dir ImGuiDir) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ArrowButton(str_idArg, C.ImGuiDir(dir)) != C.bool(true)
}

func Spacing() {
	C.Spacing()
}

func Columns(count int32, id string, border bool) {
	idArg, idFin := wrapString(id)
	defer idFin()

	borderArg := C.bool(border)

	C.Columns(C.int(count), idArg, borderArg)
}

func (self ImFont) AddGlyph(src_cfg ImFontConfig, c ImWchar, x0 float32, y0 float32, x1 float32, y1 float32, u0 float32, v0 float32, u1 float32, v1 float32, advance_x float32) {
	C.ImFont_AddGlyph(self.handle(), src_cfg.handle(), C.ImWchar(c), C.float(x0), C.float(y0), C.float(x1), C.float(y1), C.float(u0), C.float(v0), C.float(u1), C.float(v1), C.float(advance_x))
}

func TableHeadersRow() {
	C.TableHeadersRow()
}

func LogToTTY(auto_open_depth int32) {
	C.LogToTTY(C.int(auto_open_depth))
}

func IsAnyItemHovered() bool {
	return C.IsAnyItemHovered() != C.bool(true)
}

func (self ImDrawList) PrimWriteIdx(idx ImDrawIdx) {
	C.ImDrawList_PrimWriteIdx(self.handle(), C.ImDrawIdx(idx))
}

func GetContentRegionMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetContentRegionMax(pOutArg)
}

func SetNextItemWidth(item_width float32) {
	C.SetNextItemWidth(C.float(item_width))
}

func GetWindowPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowPos(pOutArg)
}

func PopAllowKeyboardFocus() {
	C.PopAllowKeyboardFocus()
}

func GetItemRectSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectSize(pOutArg)
}

func (self ImFontAtlas) AddCustomRectFontGlyph(font ImFont, id ImWchar, width int32, height int32, advance_x float32, offset ImVec2) int {
	return int(C.ImFontAtlas_AddCustomRectFontGlyph(self.handle(), font.handle(), C.ImWchar(id), C.int(width), C.int(height), C.float(advance_x), offset.ToC()))
}

func TreePop() {
	C.TreePop()
}

func (self ImGuiIO) AddKeyEvent(key ImGuiKey, down bool) {
	downArg := C.bool(down)

	C.ImGuiIO_AddKeyEvent(self.handle(), C.ImGuiKey(key), downArg)
}

func (self ImGuiStorage) SetInt(key ImGuiID, val int32) {
	C.ImGuiStorage_SetInt(self.handle(), C.ImGuiID(key), C.int(val))
}

func ImDrawList_GetClipRectMax(pOut *ImVec2, self ImDrawList) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImDrawList_GetClipRectMax(pOutArg, self.handle())
}

func (self ImFontAtlas) GetTexDataAsAlpha8(out_pixels *C.uchar, out_width *int32, out_height *int32, out_bytes_per_pixel *int32) {
	out_widthArg, out_widthFin := wrapInt32(out_width)
	defer out_widthFin()

	out_heightArg, out_heightFin := wrapInt32(out_height)
	defer out_heightFin()

	out_bytes_per_pixelArg, out_bytes_per_pixelFin := wrapInt32(out_bytes_per_pixel)
	defer out_bytes_per_pixelFin()

	C.ImFontAtlas_GetTexDataAsAlpha8(self.handle(), &out_pixels, out_widthArg, out_heightArg, out_bytes_per_pixelArg)
}

func (self ImFont) ClearOutputData() {
	C.ImFont_ClearOutputData(self.handle())
}

func NewFrame() {
	C.NewFrame()
}

func (self ImDrawList) PopTextureID() {
	C.ImDrawList_PopTextureID(self.handle())
}

func (self ImGuiStorage) GetBool(key ImGuiID, default_val bool) bool {
	default_valArg := C.bool(default_val)

	return C.ImGuiStorage_GetBool(self.handle(), C.ImGuiID(key), default_valArg) != C.bool(true)
}

func BeginPopupContextWindow(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextWindow(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
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

func DragInt3(label string, v *[3]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt3(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func InputInt3(label string, v *[3]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt3(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func (self ImFont) AddRemapChar(dst ImWchar, src ImWchar, overwrite_dst bool) {
	overwrite_dstArg := C.bool(overwrite_dst)

	C.ImFont_AddRemapChar(self.handle(), C.ImWchar(dst), C.ImWchar(src), overwrite_dstArg)
}

func (self ImFont) SetGlyphVisible(c ImWchar, visible bool) {
	visibleArg := C.bool(visible)

	C.ImFont_SetGlyphVisible(self.handle(), C.ImWchar(c), visibleArg)
}

func GetKeyName(key ImGuiKey) string {
	return C.GoString(C.GetKeyName(C.ImGuiKey(key)))
}

func OpenPopupOnItemClick(str_id string, popup_flags ImGuiPopupFlags) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.OpenPopupOnItemClick(str_idArg, C.ImGuiPopupFlags(popup_flags))
}

func SaveIniSettingsToDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.SaveIniSettingsToDisk(ini_filenameArg)
}

func Separator() {
	C.Separator()
}

func SliderScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer, renderer_render_arg unsafe.Pointer) {
	C.RenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg)
}

func IsItemEdited() bool {
	return C.IsItemEdited() != C.bool(true)
}

func ShowDebugLogWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDebugLogWindow(p_openArg)
}

func (self ImDrawList) PrimQuadUV(a ImVec2, b ImVec2, c ImVec2, d ImVec2, uv_a ImVec2, uv_b ImVec2, uv_c ImVec2, uv_d ImVec2, col uint32) {
	C.ImDrawList_PrimQuadUV(self.handle(), a.ToC(), b.ToC(), c.ToC(), d.ToC(), uv_a.ToC(), uv_b.ToC(), uv_c.ToC(), uv_d.ToC(), C.uint(col))
}

func InputInt2(label string, v *[2]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt2(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func InputInt4(label string, v *[4]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt4(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func IsItemToggledOpen() bool {
	return C.IsItemToggledOpen() != C.bool(true)
}

func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableHeader(labelArg)
}

func TableSetColumnEnabled(column_n int32, v bool) {
	vArg := C.bool(v)

	C.TableSetColumnEnabled(C.int(column_n), vArg)
}

func (self ImDrawList) ChannelsMerge() {
	C.ImDrawList_ChannelsMerge(self.handle())
}

func FindViewportByPlatformHandle(platform_handle unsafe.Pointer) ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.FindViewportByPlatformHandle(platform_handle)))
}

func SetCursorPosY(local_y float32) {
	C.SetCursorPosY(C.float(local_y))
}

func GetWindowViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.GetWindowViewport()))
}

func GetCursorPosY() float32 {
	return float32(C.GetCursorPosY())
}

func GetStyleColorName(idx ImGuiCol) string {
	return C.GoString(C.GetStyleColorName(C.ImGuiCol(idx)))
}

func PushClipRect(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

	C.PushClipRect(clip_rect_min.ToC(), clip_rect_max.ToC(), intersect_with_current_clip_rectArg)
}

func DestroyPlatformWindows() {
	C.DestroyPlatformWindows()
}

func (self ImDrawList) PrimReserve(idx_count int32, vtx_count int32) {
	C.ImDrawList_PrimReserve(self.handle(), C.int(idx_count), C.int(vtx_count))
}

func (self ImFont) IsGlyphRangeUnused(c_begin uint32, c_last uint32) bool {
	return C.ImFont_IsGlyphRangeUnused(self.handle(), C.uint(c_begin), C.uint(c_last)) != C.bool(true)
}

func EndGroup() {
	C.EndGroup()
}

func (self ImFontAtlas) SetTexID(id ImTextureID) {
	C.ImFontAtlas_SetTexID(self.handle(), C.ImTextureID(id))
}

func EndMainMenuBar() {
	C.EndMainMenuBar()
}

func GetFontTexUvWhitePixel(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetFontTexUvWhitePixel(pOutArg)
}

func IsWindowHovered(flags ImGuiHoveredFlags) bool {
	return C.IsWindowHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func (self ImGuiInputTextCallbackData) ClearSelection() {
	C.ImGuiInputTextCallbackData_ClearSelection(self.handle())
}

func GetMouseDragDelta(pOut *ImVec2, button ImGuiMouseButton, lock_threshold float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMouseDragDelta(pOutArg, C.ImGuiMouseButton(button), C.float(lock_threshold))
}

func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond) {
	collapsedArg := C.bool(collapsed)

	C.SetNextWindowCollapsed(collapsedArg, C.ImGuiCond(cond))
}

func ColorPicker4(label string, col *[4]float32, flags ImGuiColorEditFlags, ref_col *float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	ref_colArg, ref_colFin := wrapFloat(ref_col)
	defer ref_colFin()

	return C.ColorPicker4(labelArg, colArg, C.ImGuiColorEditFlags(flags), ref_colArg) != C.bool(true)
}

func (self ImGuiInputEventMouseButton) SetButton(v int32) {
	C.ImGuiInputEventMouseButton_SetButton(self.handle(), C.int(v))
}

func (self ImGuiInputEventMouseButton) SetDown(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputEventMouseButton_SetDown(self.handle(), vArg)
}

func (self ImGuiTableSettings) SetID(v ImGuiID) {
	C.ImGuiTableSettings_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableSettings) SetSaveFlags(v ImGuiTableFlags) {
	C.ImGuiTableSettings_SetSaveFlags(self.handle(), C.ImGuiTableFlags(v))
}

func (self ImGuiTableSettings) SetRefScale(v float32) {
	C.ImGuiTableSettings_SetRefScale(self.handle(), C.float(v))
}

func (self ImGuiTableSettings) SetWantApply(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableSettings_SetWantApply(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfIDStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfIDStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfColorStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfColorStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfStyleVarStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfStyleVarStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfFontStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfFontStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfFocusScopeStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfFocusScopeStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfGroupStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfGroupStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfItemFlagsStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfItemFlagsStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfBeginPopupStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfBeginPopupStack(self.handle(), vArg)
}

func (self ImGuiStackSizes) SetSizeOfDisabledStack(v int) {
	vArg := C.short(v)

	C.ImGuiStackSizes_SetSizeOfDisabledStack(self.handle(), vArg)
}

func (self ImGuiTableColumnSortSpecs) SetColumnUserID(v ImGuiID) {
	C.ImGuiTableColumnSortSpecs_SetColumnUserID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableColumnSortSpecs) SetSortDirection(v ImGuiSortDirection) {
	C.ImGuiTableColumnSortSpecs_SetSortDirection(self.handle(), C.ImGuiSortDirection(v))
}

func (self ImGuiWindowTempData) SetCursorPos(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorPos(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetCursorPosPrevLine(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorPosPrevLine(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetCursorStartPos(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorStartPos(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetCursorMaxPos(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorMaxPos(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetIdealMaxPos(v ImVec2) {
	C.ImGuiWindowTempData_SetIdealMaxPos(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetCurrLineSize(v ImVec2) {
	C.ImGuiWindowTempData_SetCurrLineSize(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetPrevLineSize(v ImVec2) {
	C.ImGuiWindowTempData_SetPrevLineSize(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetCurrLineTextBaseOffset(v float32) {
	C.ImGuiWindowTempData_SetCurrLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) SetPrevLineTextBaseOffset(v float32) {
	C.ImGuiWindowTempData_SetPrevLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) SetIsSameLine(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowTempData_SetIsSameLine(self.handle(), vArg)
}

func (self ImGuiWindowTempData) SetCursorStartPosLossyness(v ImVec2) {
	C.ImGuiWindowTempData_SetCursorStartPosLossyness(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetNavLayerCurrent(v ImGuiNavLayer) {
	C.ImGuiWindowTempData_SetNavLayerCurrent(self.handle(), C.ImGuiNavLayer(v))
}

func (self ImGuiWindowTempData) SetNavLayersActiveMask(v int) {
	vArg := C.short(v)

	C.ImGuiWindowTempData_SetNavLayersActiveMask(self.handle(), vArg)
}

func (self ImGuiWindowTempData) SetNavLayersActiveMaskNext(v int) {
	vArg := C.short(v)

	C.ImGuiWindowTempData_SetNavLayersActiveMaskNext(self.handle(), vArg)
}

func (self ImGuiWindowTempData) SetNavFocusScopeIdCurrent(v ImGuiID) {
	C.ImGuiWindowTempData_SetNavFocusScopeIdCurrent(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowTempData) SetNavHideHighlightOneFrame(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowTempData_SetNavHideHighlightOneFrame(self.handle(), vArg)
}

func (self ImGuiWindowTempData) SetNavHasScroll(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowTempData_SetNavHasScroll(self.handle(), vArg)
}

func (self ImGuiWindowTempData) SetMenuBarAppending(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowTempData_SetMenuBarAppending(self.handle(), vArg)
}

func (self ImGuiWindowTempData) SetMenuBarOffset(v ImVec2) {
	C.ImGuiWindowTempData_SetMenuBarOffset(self.handle(), v.ToC())
}

func (self ImGuiWindowTempData) SetTreeDepth(v int32) {
	C.ImGuiWindowTempData_SetTreeDepth(self.handle(), C.int(v))
}

func (self ImGuiWindowTempData) SetTreeJumpToParentOnPopMask(v uint32) {
	C.ImGuiWindowTempData_SetTreeJumpToParentOnPopMask(self.handle(), C.uint(v))
}

func (self ImGuiWindowTempData) SetStateStorage(v ImGuiStorage) {
	C.ImGuiWindowTempData_SetStateStorage(self.handle(), v.handle())
}

func (self ImGuiWindowTempData) SetCurrentColumns(v ImGuiOldColumns) {
	C.ImGuiWindowTempData_SetCurrentColumns(self.handle(), v.handle())
}

func (self ImGuiWindowTempData) SetCurrentTableIdx(v int32) {
	C.ImGuiWindowTempData_SetCurrentTableIdx(self.handle(), C.int(v))
}

func (self ImGuiWindowTempData) SetLayoutType(v ImGuiLayoutType) {
	C.ImGuiWindowTempData_SetLayoutType(self.handle(), C.ImGuiLayoutType(v))
}

func (self ImGuiWindowTempData) SetParentLayoutType(v ImGuiLayoutType) {
	C.ImGuiWindowTempData_SetParentLayoutType(self.handle(), C.ImGuiLayoutType(v))
}

func (self ImGuiWindowTempData) SetItemWidth(v float32) {
	C.ImGuiWindowTempData_SetItemWidth(self.handle(), C.float(v))
}

func (self ImGuiWindowTempData) SetTextWrapPos(v float32) {
	C.ImGuiWindowTempData_SetTextWrapPos(self.handle(), C.float(v))
}

func (self ImFontAtlasCustomRect) SetGlyphID(v uint32) {
	C.ImFontAtlasCustomRect_SetGlyphID(self.handle(), C.uint(v))
}

func (self ImFontAtlasCustomRect) SetGlyphAdvanceX(v float32) {
	C.ImFontAtlasCustomRect_SetGlyphAdvanceX(self.handle(), C.float(v))
}

func (self ImFontAtlasCustomRect) SetGlyphOffset(v ImVec2) {
	C.ImFontAtlasCustomRect_SetGlyphOffset(self.handle(), v.ToC())
}

func (self ImFontAtlasCustomRect) SetFont(v ImFont) {
	C.ImFontAtlasCustomRect_SetFont(self.handle(), v.handle())
}

func (self ImGuiTableColumnSettings) SetWidthOrWeight(v float32) {
	C.ImGuiTableColumnSettings_SetWidthOrWeight(self.handle(), C.float(v))
}

func (self ImGuiTableColumnSettings) SetUserID(v ImGuiID) {
	C.ImGuiTableColumnSettings_SetUserID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetInitialized(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetInitialized(self.handle(), vArg)
}

func (self ImGuiContext) SetFontAtlasOwnedByContext(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetFontAtlasOwnedByContext(self.handle(), vArg)
}

func (self ImGuiContext) SetConfigFlagsCurrFrame(v ImGuiConfigFlags) {
	C.ImGuiContext_SetConfigFlagsCurrFrame(self.handle(), C.ImGuiConfigFlags(v))
}

func (self ImGuiContext) SetConfigFlagsLastFrame(v ImGuiConfigFlags) {
	C.ImGuiContext_SetConfigFlagsLastFrame(self.handle(), C.ImGuiConfigFlags(v))
}

func (self ImGuiContext) SetFont(v ImFont) {
	C.ImGuiContext_SetFont(self.handle(), v.handle())
}

func (self ImGuiContext) SetFontSize(v float32) {
	C.ImGuiContext_SetFontSize(self.handle(), C.float(v))
}

func (self ImGuiContext) SetFontBaseSize(v float32) {
	C.ImGuiContext_SetFontBaseSize(self.handle(), C.float(v))
}

func (self ImGuiContext) SetTime(v float64) {
	C.ImGuiContext_SetTime(self.handle(), C.double(v))
}

func (self ImGuiContext) SetFrameCount(v int32) {
	C.ImGuiContext_SetFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetFrameCountEnded(v int32) {
	C.ImGuiContext_SetFrameCountEnded(self.handle(), C.int(v))
}

func (self ImGuiContext) SetFrameCountPlatformEnded(v int32) {
	C.ImGuiContext_SetFrameCountPlatformEnded(self.handle(), C.int(v))
}

func (self ImGuiContext) SetFrameCountRendered(v int32) {
	C.ImGuiContext_SetFrameCountRendered(self.handle(), C.int(v))
}

func (self ImGuiContext) SetWithinFrameScope(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetWithinFrameScope(self.handle(), vArg)
}

func (self ImGuiContext) SetWithinFrameScopeWithImplicitWindow(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetWithinFrameScopeWithImplicitWindow(self.handle(), vArg)
}

func (self ImGuiContext) SetWithinEndChild(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetWithinEndChild(self.handle(), vArg)
}

func (self ImGuiContext) SetGcCompactAll(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetGcCompactAll(self.handle(), vArg)
}

func (self ImGuiContext) SetTestEngineHookItems(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetTestEngineHookItems(self.handle(), vArg)
}

func (self ImGuiContext) SetTestEngine(v unsafe.Pointer) {
	C.ImGuiContext_SetTestEngine(self.handle(), v)
}

func (self ImGuiContext) SetWindowsActiveCount(v int32) {
	C.ImGuiContext_SetWindowsActiveCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetWindowsHoverPadding(v ImVec2) {
	C.ImGuiContext_SetWindowsHoverPadding(self.handle(), v.ToC())
}

func (self ImGuiContext) SetCurrentWindow(v ImGuiWindow) {
	C.ImGuiContext_SetCurrentWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetHoveredWindow(v ImGuiWindow) {
	C.ImGuiContext_SetHoveredWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetHoveredWindowUnderMovingWindow(v ImGuiWindow) {
	C.ImGuiContext_SetHoveredWindowUnderMovingWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetHoveredDockNode(v ImGuiDockNode) {
	C.ImGuiContext_SetHoveredDockNode(self.handle(), v.handle())
}

func (self ImGuiContext) SetMovingWindow(v ImGuiWindow) {
	C.ImGuiContext_SetMovingWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetWheelingWindow(v ImGuiWindow) {
	C.ImGuiContext_SetWheelingWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetWheelingWindowRefMousePos(v ImVec2) {
	C.ImGuiContext_SetWheelingWindowRefMousePos(self.handle(), v.ToC())
}

func (self ImGuiContext) SetWheelingWindowTimer(v float32) {
	C.ImGuiContext_SetWheelingWindowTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetDebugHookIdInfo(v ImGuiID) {
	C.ImGuiContext_SetDebugHookIdInfo(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetHoveredId(v ImGuiID) {
	C.ImGuiContext_SetHoveredId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetHoveredIdPreviousFrame(v ImGuiID) {
	C.ImGuiContext_SetHoveredIdPreviousFrame(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetHoveredIdAllowOverlap(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetHoveredIdAllowOverlap(self.handle(), vArg)
}

func (self ImGuiContext) SetHoveredIdUsingMouseWheel(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetHoveredIdUsingMouseWheel(self.handle(), vArg)
}

func (self ImGuiContext) SetHoveredIdPreviousFrameUsingMouseWheel(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetHoveredIdPreviousFrameUsingMouseWheel(self.handle(), vArg)
}

func (self ImGuiContext) SetHoveredIdDisabled(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetHoveredIdDisabled(self.handle(), vArg)
}

func (self ImGuiContext) SetHoveredIdTimer(v float32) {
	C.ImGuiContext_SetHoveredIdTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetHoveredIdNotActiveTimer(v float32) {
	C.ImGuiContext_SetHoveredIdNotActiveTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetActiveId(v ImGuiID) {
	C.ImGuiContext_SetActiveId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetActiveIdIsAlive(v ImGuiID) {
	C.ImGuiContext_SetActiveIdIsAlive(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetActiveIdTimer(v float32) {
	C.ImGuiContext_SetActiveIdTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetActiveIdIsJustActivated(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdIsJustActivated(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdAllowOverlap(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdAllowOverlap(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdNoClearOnFocusLoss(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdNoClearOnFocusLoss(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdHasBeenPressedBefore(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdHasBeenPressedBefore(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdHasBeenEditedBefore(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdHasBeenEditedBefore(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdHasBeenEditedThisFrame(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdHasBeenEditedThisFrame(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdClickOffset(v ImVec2) {
	C.ImGuiContext_SetActiveIdClickOffset(self.handle(), v.ToC())
}

func (self ImGuiContext) SetActiveIdWindow(v ImGuiWindow) {
	C.ImGuiContext_SetActiveIdWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetActiveIdSource(v ImGuiInputSource) {
	C.ImGuiContext_SetActiveIdSource(self.handle(), C.ImGuiInputSource(v))
}

func (self ImGuiContext) SetActiveIdMouseButton(v int32) {
	C.ImGuiContext_SetActiveIdMouseButton(self.handle(), C.int(v))
}

func (self ImGuiContext) SetActiveIdPreviousFrame(v ImGuiID) {
	C.ImGuiContext_SetActiveIdPreviousFrame(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetActiveIdPreviousFrameIsAlive(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdPreviousFrameIsAlive(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdPreviousFrameHasBeenEditedBefore(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetActiveIdPreviousFrameHasBeenEditedBefore(self.handle(), vArg)
}

func (self ImGuiContext) SetActiveIdPreviousFrameWindow(v ImGuiWindow) {
	C.ImGuiContext_SetActiveIdPreviousFrameWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetLastActiveId(v ImGuiID) {
	C.ImGuiContext_SetLastActiveId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetLastActiveIdTimer(v float32) {
	C.ImGuiContext_SetLastActiveIdTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetActiveIdUsingNavDirMask(v uint32) {
	C.ImGuiContext_SetActiveIdUsingNavDirMask(self.handle(), C.uint(v))
}

func (self ImGuiContext) SetActiveIdUsingNavInputMask(v uint32) {
	C.ImGuiContext_SetActiveIdUsingNavInputMask(self.handle(), C.uint(v))
}

func (self ImGuiContext) SetCurrentItemFlags(v ImGuiItemFlags) {
	C.ImGuiContext_SetCurrentItemFlags(self.handle(), C.ImGuiItemFlags(v))
}

func (self ImGuiContext) SetBeginMenuCount(v int32) {
	C.ImGuiContext_SetBeginMenuCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetCurrentDpiScale(v float32) {
	C.ImGuiContext_SetCurrentDpiScale(self.handle(), C.float(v))
}

func (self ImGuiContext) SetCurrentViewport(v ImGuiViewportP) {
	C.ImGuiContext_SetCurrentViewport(self.handle(), v.handle())
}

func (self ImGuiContext) SetMouseViewport(v ImGuiViewportP) {
	C.ImGuiContext_SetMouseViewport(self.handle(), v.handle())
}

func (self ImGuiContext) SetMouseLastHoveredViewport(v ImGuiViewportP) {
	C.ImGuiContext_SetMouseLastHoveredViewport(self.handle(), v.handle())
}

func (self ImGuiContext) SetPlatformLastFocusedViewportId(v ImGuiID) {
	C.ImGuiContext_SetPlatformLastFocusedViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetViewportFrontMostStampCount(v int32) {
	C.ImGuiContext_SetViewportFrontMostStampCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetNavWindow(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetNavId(v ImGuiID) {
	C.ImGuiContext_SetNavId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavFocusScopeId(v ImGuiID) {
	C.ImGuiContext_SetNavFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavActivateId(v ImGuiID) {
	C.ImGuiContext_SetNavActivateId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavActivateDownId(v ImGuiID) {
	C.ImGuiContext_SetNavActivateDownId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavActivatePressedId(v ImGuiID) {
	C.ImGuiContext_SetNavActivatePressedId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavActivateInputId(v ImGuiID) {
	C.ImGuiContext_SetNavActivateInputId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavActivateFlags(v ImGuiActivateFlags) {
	C.ImGuiContext_SetNavActivateFlags(self.handle(), C.ImGuiActivateFlags(v))
}

func (self ImGuiContext) SetNavJustMovedToId(v ImGuiID) {
	C.ImGuiContext_SetNavJustMovedToId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavJustMovedToFocusScopeId(v ImGuiID) {
	C.ImGuiContext_SetNavJustMovedToFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavJustMovedToKeyMods(v ImGuiModFlags) {
	C.ImGuiContext_SetNavJustMovedToKeyMods(self.handle(), C.ImGuiModFlags(v))
}

func (self ImGuiContext) SetNavNextActivateId(v ImGuiID) {
	C.ImGuiContext_SetNavNextActivateId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavNextActivateFlags(v ImGuiActivateFlags) {
	C.ImGuiContext_SetNavNextActivateFlags(self.handle(), C.ImGuiActivateFlags(v))
}

func (self ImGuiContext) SetNavInputSource(v ImGuiInputSource) {
	C.ImGuiContext_SetNavInputSource(self.handle(), C.ImGuiInputSource(v))
}

func (self ImGuiContext) SetNavLayer(v ImGuiNavLayer) {
	C.ImGuiContext_SetNavLayer(self.handle(), C.ImGuiNavLayer(v))
}

func (self ImGuiContext) SetNavIdIsAlive(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavIdIsAlive(self.handle(), vArg)
}

func (self ImGuiContext) SetNavMousePosDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavMousePosDirty(self.handle(), vArg)
}

func (self ImGuiContext) SetNavDisableHighlight(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavDisableHighlight(self.handle(), vArg)
}

func (self ImGuiContext) SetNavDisableMouseHover(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavDisableMouseHover(self.handle(), vArg)
}

func (self ImGuiContext) SetNavAnyRequest(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavAnyRequest(self.handle(), vArg)
}

func (self ImGuiContext) SetNavInitRequest(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavInitRequest(self.handle(), vArg)
}

func (self ImGuiContext) SetNavInitRequestFromMove(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavInitRequestFromMove(self.handle(), vArg)
}

func (self ImGuiContext) SetNavInitResultId(v ImGuiID) {
	C.ImGuiContext_SetNavInitResultId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetNavMoveSubmitted(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavMoveSubmitted(self.handle(), vArg)
}

func (self ImGuiContext) SetNavMoveScoringItems(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavMoveScoringItems(self.handle(), vArg)
}

func (self ImGuiContext) SetNavMoveForwardToNextFrame(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavMoveForwardToNextFrame(self.handle(), vArg)
}

func (self ImGuiContext) SetNavMoveFlags(v ImGuiNavMoveFlags) {
	C.ImGuiContext_SetNavMoveFlags(self.handle(), C.ImGuiNavMoveFlags(v))
}

func (self ImGuiContext) SetNavMoveScrollFlags(v ImGuiScrollFlags) {
	C.ImGuiContext_SetNavMoveScrollFlags(self.handle(), C.ImGuiScrollFlags(v))
}

func (self ImGuiContext) SetNavMoveKeyMods(v ImGuiModFlags) {
	C.ImGuiContext_SetNavMoveKeyMods(self.handle(), C.ImGuiModFlags(v))
}

func (self ImGuiContext) SetNavMoveDir(v ImGuiDir) {
	C.ImGuiContext_SetNavMoveDir(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiContext) SetNavMoveDirForDebug(v ImGuiDir) {
	C.ImGuiContext_SetNavMoveDirForDebug(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiContext) SetNavMoveClipDir(v ImGuiDir) {
	C.ImGuiContext_SetNavMoveClipDir(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiContext) SetNavScoringDebugCount(v int32) {
	C.ImGuiContext_SetNavScoringDebugCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetNavTabbingDir(v int32) {
	C.ImGuiContext_SetNavTabbingDir(self.handle(), C.int(v))
}

func (self ImGuiContext) SetNavTabbingCounter(v int32) {
	C.ImGuiContext_SetNavTabbingCounter(self.handle(), C.int(v))
}

func (self ImGuiContext) SetNavWindowingTarget(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindowingTarget(self.handle(), v.handle())
}

func (self ImGuiContext) SetNavWindowingTargetAnim(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindowingTargetAnim(self.handle(), v.handle())
}

func (self ImGuiContext) SetNavWindowingListWindow(v ImGuiWindow) {
	C.ImGuiContext_SetNavWindowingListWindow(self.handle(), v.handle())
}

func (self ImGuiContext) SetNavWindowingTimer(v float32) {
	C.ImGuiContext_SetNavWindowingTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetNavWindowingHighlightAlpha(v float32) {
	C.ImGuiContext_SetNavWindowingHighlightAlpha(self.handle(), C.float(v))
}

func (self ImGuiContext) SetNavWindowingToggleLayer(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetNavWindowingToggleLayer(self.handle(), vArg)
}

func (self ImGuiContext) SetNavWindowingAccumDeltaPos(v ImVec2) {
	C.ImGuiContext_SetNavWindowingAccumDeltaPos(self.handle(), v.ToC())
}

func (self ImGuiContext) SetNavWindowingAccumDeltaSize(v ImVec2) {
	C.ImGuiContext_SetNavWindowingAccumDeltaSize(self.handle(), v.ToC())
}

func (self ImGuiContext) SetDimBgRatio(v float32) {
	C.ImGuiContext_SetDimBgRatio(self.handle(), C.float(v))
}

func (self ImGuiContext) SetMouseCursor(v ImGuiMouseCursor) {
	C.ImGuiContext_SetMouseCursor(self.handle(), C.ImGuiMouseCursor(v))
}

func (self ImGuiContext) SetDragDropActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetDragDropActive(self.handle(), vArg)
}

func (self ImGuiContext) SetDragDropWithinSource(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetDragDropWithinSource(self.handle(), vArg)
}

func (self ImGuiContext) SetDragDropWithinTarget(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetDragDropWithinTarget(self.handle(), vArg)
}

func (self ImGuiContext) SetDragDropSourceFlags(v ImGuiDragDropFlags) {
	C.ImGuiContext_SetDragDropSourceFlags(self.handle(), C.ImGuiDragDropFlags(v))
}

func (self ImGuiContext) SetDragDropSourceFrameCount(v int32) {
	C.ImGuiContext_SetDragDropSourceFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetDragDropMouseButton(v int32) {
	C.ImGuiContext_SetDragDropMouseButton(self.handle(), C.int(v))
}

func (self ImGuiContext) SetDragDropTargetId(v ImGuiID) {
	C.ImGuiContext_SetDragDropTargetId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetDragDropAcceptFlags(v ImGuiDragDropFlags) {
	C.ImGuiContext_SetDragDropAcceptFlags(self.handle(), C.ImGuiDragDropFlags(v))
}

func (self ImGuiContext) SetDragDropAcceptIdCurrRectSurface(v float32) {
	C.ImGuiContext_SetDragDropAcceptIdCurrRectSurface(self.handle(), C.float(v))
}

func (self ImGuiContext) SetDragDropAcceptIdCurr(v ImGuiID) {
	C.ImGuiContext_SetDragDropAcceptIdCurr(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetDragDropAcceptIdPrev(v ImGuiID) {
	C.ImGuiContext_SetDragDropAcceptIdPrev(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetDragDropAcceptFrameCount(v int32) {
	C.ImGuiContext_SetDragDropAcceptFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetDragDropHoldJustPressedId(v ImGuiID) {
	C.ImGuiContext_SetDragDropHoldJustPressedId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetClipperTempDataStacked(v int32) {
	C.ImGuiContext_SetClipperTempDataStacked(self.handle(), C.int(v))
}

func (self ImGuiContext) SetCurrentTable(v ImGuiTable) {
	C.ImGuiContext_SetCurrentTable(self.handle(), v.handle())
}

func (self ImGuiContext) SetTablesTempDataStacked(v int32) {
	C.ImGuiContext_SetTablesTempDataStacked(self.handle(), C.int(v))
}

func (self ImGuiContext) SetCurrentTabBar(v ImGuiTabBar) {
	C.ImGuiContext_SetCurrentTabBar(self.handle(), v.handle())
}

func (self ImGuiContext) SetMouseLastValidPos(v ImVec2) {
	C.ImGuiContext_SetMouseLastValidPos(self.handle(), v.ToC())
}

func (self ImGuiContext) SetTempInputId(v ImGuiID) {
	C.ImGuiContext_SetTempInputId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetColorEditOptions(v ImGuiColorEditFlags) {
	C.ImGuiContext_SetColorEditOptions(self.handle(), C.ImGuiColorEditFlags(v))
}

func (self ImGuiContext) SetColorEditLastHue(v float32) {
	C.ImGuiContext_SetColorEditLastHue(self.handle(), C.float(v))
}

func (self ImGuiContext) SetColorEditLastSat(v float32) {
	C.ImGuiContext_SetColorEditLastSat(self.handle(), C.float(v))
}

func (self ImGuiContext) SetColorEditLastColor(v uint32) {
	C.ImGuiContext_SetColorEditLastColor(self.handle(), C.uint(v))
}

func (self ImGuiContext) SetColorPickerRef(v ImVec4) {
	C.ImGuiContext_SetColorPickerRef(self.handle(), v.ToC())
}

func (self ImGuiContext) SetSliderGrabClickOffset(v float32) {
	C.ImGuiContext_SetSliderGrabClickOffset(self.handle(), C.float(v))
}

func (self ImGuiContext) SetSliderCurrentAccum(v float32) {
	C.ImGuiContext_SetSliderCurrentAccum(self.handle(), C.float(v))
}

func (self ImGuiContext) SetSliderCurrentAccumDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetSliderCurrentAccumDirty(self.handle(), vArg)
}

func (self ImGuiContext) SetDragCurrentAccumDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetDragCurrentAccumDirty(self.handle(), vArg)
}

func (self ImGuiContext) SetDragCurrentAccum(v float32) {
	C.ImGuiContext_SetDragCurrentAccum(self.handle(), C.float(v))
}

func (self ImGuiContext) SetDragSpeedDefaultRatio(v float32) {
	C.ImGuiContext_SetDragSpeedDefaultRatio(self.handle(), C.float(v))
}

func (self ImGuiContext) SetScrollbarClickDeltaToGrabCenter(v float32) {
	C.ImGuiContext_SetScrollbarClickDeltaToGrabCenter(self.handle(), C.float(v))
}

func (self ImGuiContext) SetDisabledAlphaBackup(v float32) {
	C.ImGuiContext_SetDisabledAlphaBackup(self.handle(), C.float(v))
}

func (self ImGuiContext) SetDisabledStackSize(v int) {
	vArg := C.short(v)

	C.ImGuiContext_SetDisabledStackSize(self.handle(), vArg)
}

func (self ImGuiContext) SetTooltipOverrideCount(v int) {
	vArg := C.short(v)

	C.ImGuiContext_SetTooltipOverrideCount(self.handle(), vArg)
}

func (self ImGuiContext) SetTooltipSlowDelay(v float32) {
	C.ImGuiContext_SetTooltipSlowDelay(self.handle(), C.float(v))
}

func (self ImGuiContext) SetPlatformImeViewport(v ImGuiID) {
	C.ImGuiContext_SetPlatformImeViewport(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetSettingsLoaded(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetSettingsLoaded(self.handle(), vArg)
}

func (self ImGuiContext) SetSettingsDirtyTimer(v float32) {
	C.ImGuiContext_SetSettingsDirtyTimer(self.handle(), C.float(v))
}

func (self ImGuiContext) SetHookIdNext(v ImGuiID) {
	C.ImGuiContext_SetHookIdNext(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetLogEnabled(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetLogEnabled(self.handle(), vArg)
}

func (self ImGuiContext) SetLogType(v ImGuiLogType) {
	C.ImGuiContext_SetLogType(self.handle(), C.ImGuiLogType(v))
}

func (self ImGuiContext) SetLogNextPrefix(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiContext_SetLogNextPrefix(self.handle(), vArg)
}

func (self ImGuiContext) SetLogNextSuffix(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiContext_SetLogNextSuffix(self.handle(), vArg)
}

func (self ImGuiContext) SetLogLinePosY(v float32) {
	C.ImGuiContext_SetLogLinePosY(self.handle(), C.float(v))
}

func (self ImGuiContext) SetLogLineFirstItem(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetLogLineFirstItem(self.handle(), vArg)
}

func (self ImGuiContext) SetLogDepthRef(v int32) {
	C.ImGuiContext_SetLogDepthRef(self.handle(), C.int(v))
}

func (self ImGuiContext) SetLogDepthToExpand(v int32) {
	C.ImGuiContext_SetLogDepthToExpand(self.handle(), C.int(v))
}

func (self ImGuiContext) SetLogDepthToExpandDefault(v int32) {
	C.ImGuiContext_SetLogDepthToExpandDefault(self.handle(), C.int(v))
}

func (self ImGuiContext) SetDebugLogFlags(v ImGuiDebugLogFlags) {
	C.ImGuiContext_SetDebugLogFlags(self.handle(), C.ImGuiDebugLogFlags(v))
}

func (self ImGuiContext) SetDebugItemPickerActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiContext_SetDebugItemPickerActive(self.handle(), vArg)
}

func (self ImGuiContext) SetDebugItemPickerBreakId(v ImGuiID) {
	C.ImGuiContext_SetDebugItemPickerBreakId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContext) SetFramerateSecPerFrameIdx(v int32) {
	C.ImGuiContext_SetFramerateSecPerFrameIdx(self.handle(), C.int(v))
}

func (self ImGuiContext) SetFramerateSecPerFrameCount(v int32) {
	C.ImGuiContext_SetFramerateSecPerFrameCount(self.handle(), C.int(v))
}

func (self ImGuiContext) SetFramerateSecPerFrameAccum(v float32) {
	C.ImGuiContext_SetFramerateSecPerFrameAccum(self.handle(), C.float(v))
}

func (self ImGuiContext) SetWantCaptureMouseNextFrame(v int32) {
	C.ImGuiContext_SetWantCaptureMouseNextFrame(self.handle(), C.int(v))
}

func (self ImGuiContext) SetWantCaptureKeyboardNextFrame(v int32) {
	C.ImGuiContext_SetWantCaptureKeyboardNextFrame(self.handle(), C.int(v))
}

func (self ImGuiContext) SetWantTextInputNextFrame(v int32) {
	C.ImGuiContext_SetWantTextInputNextFrame(self.handle(), C.int(v))
}

func (self ImGuiDockContext) SetWantFullRebuild(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockContext_SetWantFullRebuild(self.handle(), vArg)
}

func (self ImGuiDockNode) SetID(v ImGuiID) {
	C.ImGuiDockNode_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) SetSharedFlags(v ImGuiDockNodeFlags) {
	C.ImGuiDockNode_SetSharedFlags(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiDockNode) SetLocalFlagsInWindows(v ImGuiDockNodeFlags) {
	C.ImGuiDockNode_SetLocalFlagsInWindows(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiDockNode) SetMergedFlags(v ImGuiDockNodeFlags) {
	C.ImGuiDockNode_SetMergedFlags(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiDockNode) SetState(v ImGuiDockNodeState) {
	C.ImGuiDockNode_SetState(self.handle(), C.ImGuiDockNodeState(v))
}

func (self ImGuiDockNode) SetParentNode(v ImGuiDockNode) {
	C.ImGuiDockNode_SetParentNode(self.handle(), v.handle())
}

func (self ImGuiDockNode) SetTabBar(v ImGuiTabBar) {
	C.ImGuiDockNode_SetTabBar(self.handle(), v.handle())
}

func (self ImGuiDockNode) SetPos(v ImVec2) {
	C.ImGuiDockNode_SetPos(self.handle(), v.ToC())
}

func (self ImGuiDockNode) SetSize(v ImVec2) {
	C.ImGuiDockNode_SetSize(self.handle(), v.ToC())
}

func (self ImGuiDockNode) SetSizeRef(v ImVec2) {
	C.ImGuiDockNode_SetSizeRef(self.handle(), v.ToC())
}

func (self ImGuiDockNode) SetSplitAxis(v ImGuiAxis) {
	C.ImGuiDockNode_SetSplitAxis(self.handle(), C.ImGuiAxis(v))
}

func (self ImGuiDockNode) SetLastBgColor(v uint32) {
	C.ImGuiDockNode_SetLastBgColor(self.handle(), C.uint(v))
}

func (self ImGuiDockNode) SetHostWindow(v ImGuiWindow) {
	C.ImGuiDockNode_SetHostWindow(self.handle(), v.handle())
}

func (self ImGuiDockNode) SetVisibleWindow(v ImGuiWindow) {
	C.ImGuiDockNode_SetVisibleWindow(self.handle(), v.handle())
}

func (self ImGuiDockNode) SetCentralNode(v ImGuiDockNode) {
	C.ImGuiDockNode_SetCentralNode(self.handle(), v.handle())
}

func (self ImGuiDockNode) SetOnlyNodeWithWindows(v ImGuiDockNode) {
	C.ImGuiDockNode_SetOnlyNodeWithWindows(self.handle(), v.handle())
}

func (self ImGuiDockNode) SetCountNodeWithWindows(v int32) {
	C.ImGuiDockNode_SetCountNodeWithWindows(self.handle(), C.int(v))
}

func (self ImGuiDockNode) SetLastFrameAlive(v int32) {
	C.ImGuiDockNode_SetLastFrameAlive(self.handle(), C.int(v))
}

func (self ImGuiDockNode) SetLastFrameActive(v int32) {
	C.ImGuiDockNode_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiDockNode) SetLastFrameFocused(v int32) {
	C.ImGuiDockNode_SetLastFrameFocused(self.handle(), C.int(v))
}

func (self ImGuiDockNode) SetLastFocusedNodeId(v ImGuiID) {
	C.ImGuiDockNode_SetLastFocusedNodeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) SetSelectedTabId(v ImGuiID) {
	C.ImGuiDockNode_SetSelectedTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) SetWantCloseTabId(v ImGuiID) {
	C.ImGuiDockNode_SetWantCloseTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiDockNode) SetAuthorityForPos(v ImGuiDataAuthority) {
	C.ImGuiDockNode_SetAuthorityForPos(self.handle(), C.ImGuiDataAuthority(v))
}

func (self ImGuiDockNode) SetAuthorityForSize(v ImGuiDataAuthority) {
	C.ImGuiDockNode_SetAuthorityForSize(self.handle(), C.ImGuiDataAuthority(v))
}

func (self ImGuiDockNode) SetAuthorityForViewport(v ImGuiDataAuthority) {
	C.ImGuiDockNode_SetAuthorityForViewport(self.handle(), C.ImGuiDataAuthority(v))
}

func (self ImGuiDockNode) SetIsVisible(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetIsVisible(self.handle(), vArg)
}

func (self ImGuiDockNode) SetIsFocused(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetIsFocused(self.handle(), vArg)
}

func (self ImGuiDockNode) SetIsBgDrawnThisFrame(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetIsBgDrawnThisFrame(self.handle(), vArg)
}

func (self ImGuiDockNode) SetHasCloseButton(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetHasCloseButton(self.handle(), vArg)
}

func (self ImGuiDockNode) SetHasWindowMenuButton(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetHasWindowMenuButton(self.handle(), vArg)
}

func (self ImGuiDockNode) SetHasCentralNodeChild(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetHasCentralNodeChild(self.handle(), vArg)
}

func (self ImGuiDockNode) SetWantCloseAll(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetWantCloseAll(self.handle(), vArg)
}

func (self ImGuiDockNode) SetWantLockSizeOnce(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetWantLockSizeOnce(self.handle(), vArg)
}

func (self ImGuiDockNode) SetWantMouseMove(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetWantMouseMove(self.handle(), vArg)
}

func (self ImGuiDockNode) SetWantHiddenTabBarUpdate(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetWantHiddenTabBarUpdate(self.handle(), vArg)
}

func (self ImGuiDockNode) SetWantHiddenTabBarToggle(v bool) {
	vArg := C.bool(v)

	C.ImGuiDockNode_SetWantHiddenTabBarToggle(self.handle(), vArg)
}

func (self ImGuiGroupData) SetWindowID(v ImGuiID) {
	C.ImGuiGroupData_SetWindowID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiGroupData) SetBackupCursorPos(v ImVec2) {
	C.ImGuiGroupData_SetBackupCursorPos(self.handle(), v.ToC())
}

func (self ImGuiGroupData) SetBackupCursorMaxPos(v ImVec2) {
	C.ImGuiGroupData_SetBackupCursorMaxPos(self.handle(), v.ToC())
}

func (self ImGuiGroupData) SetBackupCurrLineSize(v ImVec2) {
	C.ImGuiGroupData_SetBackupCurrLineSize(self.handle(), v.ToC())
}

func (self ImGuiGroupData) SetBackupCurrLineTextBaseOffset(v float32) {
	C.ImGuiGroupData_SetBackupCurrLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiGroupData) SetBackupActiveIdIsAlive(v ImGuiID) {
	C.ImGuiGroupData_SetBackupActiveIdIsAlive(self.handle(), C.ImGuiID(v))
}

func (self ImGuiGroupData) SetBackupActiveIdPreviousFrameIsAlive(v bool) {
	vArg := C.bool(v)

	C.ImGuiGroupData_SetBackupActiveIdPreviousFrameIsAlive(self.handle(), vArg)
}

func (self ImGuiGroupData) SetBackupHoveredIdIsAlive(v bool) {
	vArg := C.bool(v)

	C.ImGuiGroupData_SetBackupHoveredIdIsAlive(self.handle(), vArg)
}

func (self ImGuiGroupData) SetEmitItem(v bool) {
	vArg := C.bool(v)

	C.ImGuiGroupData_SetEmitItem(self.handle(), vArg)
}

func (self ImGuiPayload) SetData(v unsafe.Pointer) {
	C.ImGuiPayload_SetData(self.handle(), v)
}

func (self ImGuiPayload) SetDataSize(v int32) {
	C.ImGuiPayload_SetDataSize(self.handle(), C.int(v))
}

func (self ImGuiPayload) SetSourceId(v ImGuiID) {
	C.ImGuiPayload_SetSourceId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPayload) SetSourceParentId(v ImGuiID) {
	C.ImGuiPayload_SetSourceParentId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPayload) SetDataFrameCount(v int32) {
	C.ImGuiPayload_SetDataFrameCount(self.handle(), C.int(v))
}

func (self ImGuiPayload) SetPreview(v bool) {
	vArg := C.bool(v)

	C.ImGuiPayload_SetPreview(self.handle(), vArg)
}

func (self ImGuiPayload) SetDelivery(v bool) {
	vArg := C.bool(v)

	C.ImGuiPayload_SetDelivery(self.handle(), vArg)
}

func (self ImDrawVert) Setpos(v ImVec2) {
	C.ImDrawVert_Setpos(self.handle(), v.ToC())
}

func (self ImDrawVert) Setuv(v ImVec2) {
	C.ImDrawVert_Setuv(self.handle(), v.ToC())
}

func (self ImDrawVert) Setcol(v uint32) {
	C.ImDrawVert_Setcol(self.handle(), C.uint(v))
}

func (self ImGuiPopupData) SetPopupId(v ImGuiID) {
	C.ImGuiPopupData_SetPopupId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPopupData) SetWindow(v ImGuiWindow) {
	C.ImGuiPopupData_SetWindow(self.handle(), v.handle())
}

func (self ImGuiPopupData) SetSourceWindow(v ImGuiWindow) {
	C.ImGuiPopupData_SetSourceWindow(self.handle(), v.handle())
}

func (self ImGuiPopupData) SetParentNavLayer(v int32) {
	C.ImGuiPopupData_SetParentNavLayer(self.handle(), C.int(v))
}

func (self ImGuiPopupData) SetOpenFrameCount(v int32) {
	C.ImGuiPopupData_SetOpenFrameCount(self.handle(), C.int(v))
}

func (self ImGuiPopupData) SetOpenParentId(v ImGuiID) {
	C.ImGuiPopupData_SetOpenParentId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiPopupData) SetOpenPopupPos(v ImVec2) {
	C.ImGuiPopupData_SetOpenPopupPos(self.handle(), v.ToC())
}

func (self ImGuiPopupData) SetOpenMousePos(v ImVec2) {
	C.ImGuiPopupData_SetOpenMousePos(self.handle(), v.ToC())
}

func (self ImGuiStyleMod) SetVarIdx(v ImGuiStyleVar) {
	C.ImGuiStyleMod_SetVarIdx(self.handle(), C.ImGuiStyleVar(v))
}

func (self ImGuiInputEventAppFocused) SetFocused(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputEventAppFocused_SetFocused(self.handle(), vArg)
}

func (self ImGuiInputEventMouseViewport) SetHoveredViewportID(v ImGuiID) {
	C.ImGuiInputEventMouseViewport_SetHoveredViewportID(self.handle(), C.ImGuiID(v))
}

func (self ImDrawListSharedData) SetTexUvWhitePixel(v ImVec2) {
	C.ImDrawListSharedData_SetTexUvWhitePixel(self.handle(), v.ToC())
}

func (self ImDrawListSharedData) SetFont(v ImFont) {
	C.ImDrawListSharedData_SetFont(self.handle(), v.handle())
}

func (self ImDrawListSharedData) SetFontSize(v float32) {
	C.ImDrawListSharedData_SetFontSize(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) SetCurveTessellationTol(v float32) {
	C.ImDrawListSharedData_SetCurveTessellationTol(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) SetCircleSegmentMaxError(v float32) {
	C.ImDrawListSharedData_SetCircleSegmentMaxError(self.handle(), C.float(v))
}

func (self ImDrawListSharedData) SetClipRectFullscreen(v ImVec4) {
	C.ImDrawListSharedData_SetClipRectFullscreen(self.handle(), v.ToC())
}

func (self ImDrawListSharedData) SetInitialFlags(v ImDrawListFlags) {
	C.ImDrawListSharedData_SetInitialFlags(self.handle(), C.ImDrawListFlags(v))
}

func (self ImDrawListSharedData) SetArcFastRadiusCutoff(v float32) {
	C.ImDrawListSharedData_SetArcFastRadiusCutoff(self.handle(), C.float(v))
}

func (self ImGuiInputTextCallbackData) SetEventFlag(v ImGuiInputTextFlags) {
	C.ImGuiInputTextCallbackData_SetEventFlag(self.handle(), C.ImGuiInputTextFlags(v))
}

func (self ImGuiInputTextCallbackData) SetFlags(v ImGuiInputTextFlags) {
	C.ImGuiInputTextCallbackData_SetFlags(self.handle(), C.ImGuiInputTextFlags(v))
}

func (self ImGuiInputTextCallbackData) SetUserData(v unsafe.Pointer) {
	C.ImGuiInputTextCallbackData_SetUserData(self.handle(), v)
}

func (self ImGuiInputTextCallbackData) SetEventChar(v ImWchar) {
	C.ImGuiInputTextCallbackData_SetEventChar(self.handle(), C.ImWchar(v))
}

func (self ImGuiInputTextCallbackData) SetEventKey(v ImGuiKey) {
	C.ImGuiInputTextCallbackData_SetEventKey(self.handle(), C.ImGuiKey(v))
}

func (self ImGuiInputTextCallbackData) SetBuf(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiInputTextCallbackData_SetBuf(self.handle(), vArg)
}

func (self ImGuiInputTextCallbackData) SetBufTextLen(v int32) {
	C.ImGuiInputTextCallbackData_SetBufTextLen(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) SetBufSize(v int32) {
	C.ImGuiInputTextCallbackData_SetBufSize(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) SetBufDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputTextCallbackData_SetBufDirty(self.handle(), vArg)
}

func (self ImGuiInputTextCallbackData) SetCursorPos(v int32) {
	C.ImGuiInputTextCallbackData_SetCursorPos(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) SetSelectionStart(v int32) {
	C.ImGuiInputTextCallbackData_SetSelectionStart(self.handle(), C.int(v))
}

func (self ImGuiInputTextCallbackData) SetSelectionEnd(v int32) {
	C.ImGuiInputTextCallbackData_SetSelectionEnd(self.handle(), C.int(v))
}

func (self ImGuiOldColumnData) SetOffsetNorm(v float32) {
	C.ImGuiOldColumnData_SetOffsetNorm(self.handle(), C.float(v))
}

func (self ImGuiOldColumnData) SetOffsetNormBeforeResize(v float32) {
	C.ImGuiOldColumnData_SetOffsetNormBeforeResize(self.handle(), C.float(v))
}

func (self ImGuiOldColumnData) SetFlags(v ImGuiOldColumnFlags) {
	C.ImGuiOldColumnData_SetFlags(self.handle(), C.ImGuiOldColumnFlags(v))
}

func (self ImGuiPlatformImeData) SetWantVisible(v bool) {
	vArg := C.bool(v)

	C.ImGuiPlatformImeData_SetWantVisible(self.handle(), vArg)
}

func (self ImGuiPlatformImeData) SetInputPos(v ImVec2) {
	C.ImGuiPlatformImeData_SetInputPos(self.handle(), v.ToC())
}

func (self ImGuiPlatformImeData) SetInputLineHeight(v float32) {
	C.ImGuiPlatformImeData_SetInputLineHeight(self.handle(), C.float(v))
}

func (self ImGuiPlatformMonitor) SetMainPos(v ImVec2) {
	C.ImGuiPlatformMonitor_SetMainPos(self.handle(), v.ToC())
}

func (self ImGuiPlatformMonitor) SetMainSize(v ImVec2) {
	C.ImGuiPlatformMonitor_SetMainSize(self.handle(), v.ToC())
}

func (self ImGuiPlatformMonitor) SetWorkPos(v ImVec2) {
	C.ImGuiPlatformMonitor_SetWorkPos(self.handle(), v.ToC())
}

func (self ImGuiPlatformMonitor) SetWorkSize(v ImVec2) {
	C.ImGuiPlatformMonitor_SetWorkSize(self.handle(), v.ToC())
}

func (self ImGuiPlatformMonitor) SetDpiScale(v float32) {
	C.ImGuiPlatformMonitor_SetDpiScale(self.handle(), C.float(v))
}

func (self ImGuiShrinkWidthItem) SetIndex(v int32) {
	C.ImGuiShrinkWidthItem_SetIndex(self.handle(), C.int(v))
}

func (self ImGuiShrinkWidthItem) SetWidth(v float32) {
	C.ImGuiShrinkWidthItem_SetWidth(self.handle(), C.float(v))
}

func (self ImGuiShrinkWidthItem) SetInitialWidth(v float32) {
	C.ImGuiShrinkWidthItem_SetInitialWidth(self.handle(), C.float(v))
}

func (self ImDrawCmdHeader) SetClipRect(v ImVec4) {
	C.ImDrawCmdHeader_SetClipRect(self.handle(), v.ToC())
}

func (self ImDrawCmdHeader) SetTextureId(v ImTextureID) {
	C.ImDrawCmdHeader_SetTextureId(self.handle(), C.ImTextureID(v))
}

func (self ImDrawCmdHeader) SetVtxOffset(v uint32) {
	C.ImDrawCmdHeader_SetVtxOffset(self.handle(), C.uint(v))
}

func (self ImDrawData) SetValid(v bool) {
	vArg := C.bool(v)

	C.ImDrawData_SetValid(self.handle(), vArg)
}

func (self ImDrawData) SetCmdListsCount(v int32) {
	C.ImDrawData_SetCmdListsCount(self.handle(), C.int(v))
}

func (self ImDrawData) SetTotalIdxCount(v int32) {
	C.ImDrawData_SetTotalIdxCount(self.handle(), C.int(v))
}

func (self ImDrawData) SetTotalVtxCount(v int32) {
	C.ImDrawData_SetTotalVtxCount(self.handle(), C.int(v))
}

func (self ImDrawData) SetDisplayPos(v ImVec2) {
	C.ImDrawData_SetDisplayPos(self.handle(), v.ToC())
}

func (self ImDrawData) SetDisplaySize(v ImVec2) {
	C.ImDrawData_SetDisplaySize(self.handle(), v.ToC())
}

func (self ImDrawData) SetFramebufferScale(v ImVec2) {
	C.ImDrawData_SetFramebufferScale(self.handle(), v.ToC())
}

func (self ImDrawData) SetOwnerViewport(v ImGuiViewport) {
	C.ImDrawData_SetOwnerViewport(self.handle(), v.handle())
}

func (self ImGuiTableTempData) SetTableIndex(v int32) {
	C.ImGuiTableTempData_SetTableIndex(self.handle(), C.int(v))
}

func (self ImGuiTableTempData) SetLastTimeActive(v float32) {
	C.ImGuiTableTempData_SetLastTimeActive(self.handle(), C.float(v))
}

func (self ImGuiTableTempData) SetUserOuterSize(v ImVec2) {
	C.ImGuiTableTempData_SetUserOuterSize(self.handle(), v.ToC())
}

func (self ImGuiTableTempData) SetHostBackupPrevLineSize(v ImVec2) {
	C.ImGuiTableTempData_SetHostBackupPrevLineSize(self.handle(), v.ToC())
}

func (self ImGuiTableTempData) SetHostBackupCurrLineSize(v ImVec2) {
	C.ImGuiTableTempData_SetHostBackupCurrLineSize(self.handle(), v.ToC())
}

func (self ImGuiTableTempData) SetHostBackupCursorMaxPos(v ImVec2) {
	C.ImGuiTableTempData_SetHostBackupCursorMaxPos(self.handle(), v.ToC())
}

func (self ImGuiTableTempData) SetHostBackupItemWidth(v float32) {
	C.ImGuiTableTempData_SetHostBackupItemWidth(self.handle(), C.float(v))
}

func (self ImGuiTableTempData) SetHostBackupItemWidthStackSize(v int32) {
	C.ImGuiTableTempData_SetHostBackupItemWidthStackSize(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiWindow_SetName(self.handle(), vArg)
}

func (self ImGuiWindow) SetID(v ImGuiID) {
	C.ImGuiWindow_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetFlags(v ImGuiWindowFlags) {
	C.ImGuiWindow_SetFlags(self.handle(), C.ImGuiWindowFlags(v))
}

func (self ImGuiWindow) SetFlagsPreviousFrame(v ImGuiWindowFlags) {
	C.ImGuiWindow_SetFlagsPreviousFrame(self.handle(), C.ImGuiWindowFlags(v))
}

func (self ImGuiWindow) SetViewport(v ImGuiViewportP) {
	C.ImGuiWindow_SetViewport(self.handle(), v.handle())
}

func (self ImGuiWindow) SetViewportId(v ImGuiID) {
	C.ImGuiWindow_SetViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetViewportPos(v ImVec2) {
	C.ImGuiWindow_SetViewportPos(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetViewportAllowPlatformMonitorExtend(v int32) {
	C.ImGuiWindow_SetViewportAllowPlatformMonitorExtend(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetPos(v ImVec2) {
	C.ImGuiWindow_SetPos(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetSize(v ImVec2) {
	C.ImGuiWindow_SetSize(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetSizeFull(v ImVec2) {
	C.ImGuiWindow_SetSizeFull(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetContentSize(v ImVec2) {
	C.ImGuiWindow_SetContentSize(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetContentSizeIdeal(v ImVec2) {
	C.ImGuiWindow_SetContentSizeIdeal(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetContentSizeExplicit(v ImVec2) {
	C.ImGuiWindow_SetContentSizeExplicit(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetWindowPadding(v ImVec2) {
	C.ImGuiWindow_SetWindowPadding(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetWindowRounding(v float32) {
	C.ImGuiWindow_SetWindowRounding(self.handle(), C.float(v))
}

func (self ImGuiWindow) SetWindowBorderSize(v float32) {
	C.ImGuiWindow_SetWindowBorderSize(self.handle(), C.float(v))
}

func (self ImGuiWindow) SetNameBufLen(v int32) {
	C.ImGuiWindow_SetNameBufLen(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetMoveId(v ImGuiID) {
	C.ImGuiWindow_SetMoveId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetTabId(v ImGuiID) {
	C.ImGuiWindow_SetTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetChildId(v ImGuiID) {
	C.ImGuiWindow_SetChildId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetScroll(v ImVec2) {
	C.ImGuiWindow_SetScroll(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetScrollMax(v ImVec2) {
	C.ImGuiWindow_SetScrollMax(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetScrollTarget(v ImVec2) {
	C.ImGuiWindow_SetScrollTarget(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetScrollTargetCenterRatio(v ImVec2) {
	C.ImGuiWindow_SetScrollTargetCenterRatio(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetScrollTargetEdgeSnapDist(v ImVec2) {
	C.ImGuiWindow_SetScrollTargetEdgeSnapDist(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetScrollbarSizes(v ImVec2) {
	C.ImGuiWindow_SetScrollbarSizes(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetScrollbarX(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetScrollbarX(self.handle(), vArg)
}

func (self ImGuiWindow) SetScrollbarY(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetScrollbarY(self.handle(), vArg)
}

func (self ImGuiWindow) SetViewportOwned(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetViewportOwned(self.handle(), vArg)
}

func (self ImGuiWindow) SetActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetActive(self.handle(), vArg)
}

func (self ImGuiWindow) SetWasActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetWasActive(self.handle(), vArg)
}

func (self ImGuiWindow) SetWriteAccessed(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetWriteAccessed(self.handle(), vArg)
}

func (self ImGuiWindow) SetCollapsed(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetCollapsed(self.handle(), vArg)
}

func (self ImGuiWindow) SetWantCollapseToggle(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetWantCollapseToggle(self.handle(), vArg)
}

func (self ImGuiWindow) SetSkipItems(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetSkipItems(self.handle(), vArg)
}

func (self ImGuiWindow) SetAppearing(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetAppearing(self.handle(), vArg)
}

func (self ImGuiWindow) SetHidden(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetHidden(self.handle(), vArg)
}

func (self ImGuiWindow) SetIsFallbackWindow(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetIsFallbackWindow(self.handle(), vArg)
}

func (self ImGuiWindow) SetIsExplicitChild(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetIsExplicitChild(self.handle(), vArg)
}

func (self ImGuiWindow) SetHasCloseButton(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetHasCloseButton(self.handle(), vArg)
}

func (self ImGuiWindow) SetBeginCount(v int) {
	vArg := C.short(v)

	C.ImGuiWindow_SetBeginCount(self.handle(), vArg)
}

func (self ImGuiWindow) SetBeginOrderWithinParent(v int) {
	vArg := C.short(v)

	C.ImGuiWindow_SetBeginOrderWithinParent(self.handle(), vArg)
}

func (self ImGuiWindow) SetBeginOrderWithinContext(v int) {
	vArg := C.short(v)

	C.ImGuiWindow_SetBeginOrderWithinContext(self.handle(), vArg)
}

func (self ImGuiWindow) SetFocusOrder(v int) {
	vArg := C.short(v)

	C.ImGuiWindow_SetFocusOrder(self.handle(), vArg)
}

func (self ImGuiWindow) SetPopupId(v ImGuiID) {
	C.ImGuiWindow_SetPopupId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetAutoFitOnlyGrows(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetAutoFitOnlyGrows(self.handle(), vArg)
}

func (self ImGuiWindow) SetAutoPosLastDirection(v ImGuiDir) {
	C.ImGuiWindow_SetAutoPosLastDirection(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiWindow) SetSetWindowPosAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowPosAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) SetSetWindowSizeAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowSizeAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) SetSetWindowCollapsedAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowCollapsedAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) SetSetWindowDockAllowFlags(v ImGuiCond) {
	C.ImGuiWindow_SetSetWindowDockAllowFlags(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiWindow) SetSetWindowPosVal(v ImVec2) {
	C.ImGuiWindow_SetSetWindowPosVal(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetSetWindowPosPivot(v ImVec2) {
	C.ImGuiWindow_SetSetWindowPosPivot(self.handle(), v.ToC())
}

func (self ImGuiWindow) SetLastFrameActive(v int32) {
	C.ImGuiWindow_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetLastFrameJustFocused(v int32) {
	C.ImGuiWindow_SetLastFrameJustFocused(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetLastTimeActive(v float32) {
	C.ImGuiWindow_SetLastTimeActive(self.handle(), C.float(v))
}

func (self ImGuiWindow) SetItemWidthDefault(v float32) {
	C.ImGuiWindow_SetItemWidthDefault(self.handle(), C.float(v))
}

func (self ImGuiWindow) SetFontWindowScale(v float32) {
	C.ImGuiWindow_SetFontWindowScale(self.handle(), C.float(v))
}

func (self ImGuiWindow) SetFontDpiScale(v float32) {
	C.ImGuiWindow_SetFontDpiScale(self.handle(), C.float(v))
}

func (self ImGuiWindow) SetSettingsOffset(v int32) {
	C.ImGuiWindow_SetSettingsOffset(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetDrawList(v ImDrawList) {
	C.ImGuiWindow_SetDrawList(self.handle(), v.handle())
}

func (self ImGuiWindow) SetParentWindow(v ImGuiWindow) {
	C.ImGuiWindow_SetParentWindow(self.handle(), v.handle())
}

func (self ImGuiWindow) SetParentWindowInBeginStack(v ImGuiWindow) {
	C.ImGuiWindow_SetParentWindowInBeginStack(self.handle(), v.handle())
}

func (self ImGuiWindow) SetRootWindow(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindow(self.handle(), v.handle())
}

func (self ImGuiWindow) SetRootWindowPopupTree(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowPopupTree(self.handle(), v.handle())
}

func (self ImGuiWindow) SetRootWindowDockTree(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowDockTree(self.handle(), v.handle())
}

func (self ImGuiWindow) SetRootWindowForTitleBarHighlight(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowForTitleBarHighlight(self.handle(), v.handle())
}

func (self ImGuiWindow) SetRootWindowForNav(v ImGuiWindow) {
	C.ImGuiWindow_SetRootWindowForNav(self.handle(), v.handle())
}

func (self ImGuiWindow) SetNavLastChildNavWindow(v ImGuiWindow) {
	C.ImGuiWindow_SetNavLastChildNavWindow(self.handle(), v.handle())
}

func (self ImGuiWindow) SetMemoryDrawListIdxCapacity(v int32) {
	C.ImGuiWindow_SetMemoryDrawListIdxCapacity(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetMemoryDrawListVtxCapacity(v int32) {
	C.ImGuiWindow_SetMemoryDrawListVtxCapacity(self.handle(), C.int(v))
}

func (self ImGuiWindow) SetMemoryCompacted(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetMemoryCompacted(self.handle(), vArg)
}

func (self ImGuiWindow) SetDockIsActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetDockIsActive(self.handle(), vArg)
}

func (self ImGuiWindow) SetDockNodeIsVisible(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetDockNodeIsVisible(self.handle(), vArg)
}

func (self ImGuiWindow) SetDockTabIsVisible(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetDockTabIsVisible(self.handle(), vArg)
}

func (self ImGuiWindow) SetDockTabWantClose(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindow_SetDockTabWantClose(self.handle(), vArg)
}

func (self ImGuiWindow) SetDockOrder(v int) {
	vArg := C.short(v)

	C.ImGuiWindow_SetDockOrder(self.handle(), vArg)
}

func (self ImGuiWindow) SetDockNode(v ImGuiDockNode) {
	C.ImGuiWindow_SetDockNode(self.handle(), v.handle())
}

func (self ImGuiWindow) SetDockNodeAsHost(v ImGuiDockNode) {
	C.ImGuiWindow_SetDockNodeAsHost(self.handle(), v.handle())
}

func (self ImGuiWindow) SetDockId(v ImGuiID) {
	C.ImGuiWindow_SetDockId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindow) SetDockTabItemStatusFlags(v ImGuiItemStatusFlags) {
	C.ImGuiWindow_SetDockTabItemStatusFlags(self.handle(), C.ImGuiItemStatusFlags(v))
}

func (self ImGuiWindowClass) SetClassId(v ImGuiID) {
	C.ImGuiWindowClass_SetClassId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowClass) SetParentViewportId(v ImGuiID) {
	C.ImGuiWindowClass_SetParentViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowClass) SetViewportFlagsOverrideSet(v ImGuiViewportFlags) {
	C.ImGuiWindowClass_SetViewportFlagsOverrideSet(self.handle(), C.ImGuiViewportFlags(v))
}

func (self ImGuiWindowClass) SetViewportFlagsOverrideClear(v ImGuiViewportFlags) {
	C.ImGuiWindowClass_SetViewportFlagsOverrideClear(self.handle(), C.ImGuiViewportFlags(v))
}

func (self ImGuiWindowClass) SetTabItemFlagsOverrideSet(v ImGuiTabItemFlags) {
	C.ImGuiWindowClass_SetTabItemFlagsOverrideSet(self.handle(), C.ImGuiTabItemFlags(v))
}

func (self ImGuiWindowClass) SetDockNodeFlagsOverrideSet(v ImGuiDockNodeFlags) {
	C.ImGuiWindowClass_SetDockNodeFlagsOverrideSet(self.handle(), C.ImGuiDockNodeFlags(v))
}

func (self ImGuiWindowClass) SetDockingAlwaysTabBar(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowClass_SetDockingAlwaysTabBar(self.handle(), vArg)
}

func (self ImGuiWindowClass) SetDockingAllowUnclassed(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowClass_SetDockingAllowUnclassed(self.handle(), vArg)
}

func (self ImGuiStackLevelInfo) SetID(v ImGuiID) {
	C.ImGuiStackLevelInfo_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiStackLevelInfo) SetQuerySuccess(v bool) {
	vArg := C.bool(v)

	C.ImGuiStackLevelInfo_SetQuerySuccess(self.handle(), vArg)
}

func (self ImGuiStackLevelInfo) SetDataType(v ImGuiDataType) {
	C.ImGuiStackLevelInfo_SetDataType(self.handle(), C.ImGuiDataType(v))
}

func (self ImGuiTableSortSpecs) SetSpecs(v ImGuiTableColumnSortSpecs) {
	C.ImGuiTableSortSpecs_SetSpecs(self.handle(), v.handle())
}

func (self ImGuiTableSortSpecs) SetSpecsCount(v int32) {
	C.ImGuiTableSortSpecs_SetSpecsCount(self.handle(), C.int(v))
}

func (self ImGuiTableSortSpecs) SetSpecsDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableSortSpecs_SetSpecsDirty(self.handle(), vArg)
}

func (self ImGuiSettingsHandler) SetTypeName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiSettingsHandler_SetTypeName(self.handle(), vArg)
}

func (self ImGuiSettingsHandler) SetTypeHash(v ImGuiID) {
	C.ImGuiSettingsHandler_SetTypeHash(self.handle(), C.ImGuiID(v))
}

func (self ImGuiSettingsHandler) SetUserData(v unsafe.Pointer) {
	C.ImGuiSettingsHandler_SetUserData(self.handle(), v)
}

func (self ImGuiTabItem) SetID(v ImGuiID) {
	C.ImGuiTabItem_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabItem) SetFlags(v ImGuiTabItemFlags) {
	C.ImGuiTabItem_SetFlags(self.handle(), C.ImGuiTabItemFlags(v))
}

func (self ImGuiTabItem) SetWindow(v ImGuiWindow) {
	C.ImGuiTabItem_SetWindow(self.handle(), v.handle())
}

func (self ImGuiTabItem) SetLastFrameVisible(v int32) {
	C.ImGuiTabItem_SetLastFrameVisible(self.handle(), C.int(v))
}

func (self ImGuiTabItem) SetLastFrameSelected(v int32) {
	C.ImGuiTabItem_SetLastFrameSelected(self.handle(), C.int(v))
}

func (self ImGuiTabItem) SetOffset(v float32) {
	C.ImGuiTabItem_SetOffset(self.handle(), C.float(v))
}

func (self ImGuiTabItem) SetWidth(v float32) {
	C.ImGuiTabItem_SetWidth(self.handle(), C.float(v))
}

func (self ImGuiTabItem) SetContentWidth(v float32) {
	C.ImGuiTabItem_SetContentWidth(self.handle(), C.float(v))
}

func (self ImGuiTabItem) SetRequestedWidth(v float32) {
	C.ImGuiTabItem_SetRequestedWidth(self.handle(), C.float(v))
}

func (self ImGuiTabItem) SetWantClose(v bool) {
	vArg := C.bool(v)

	C.ImGuiTabItem_SetWantClose(self.handle(), vArg)
}

func (self ImGuiComboPreviewData) SetBackupCursorPos(v ImVec2) {
	C.ImGuiComboPreviewData_SetBackupCursorPos(self.handle(), v.ToC())
}

func (self ImGuiComboPreviewData) SetBackupCursorMaxPos(v ImVec2) {
	C.ImGuiComboPreviewData_SetBackupCursorMaxPos(self.handle(), v.ToC())
}

func (self ImGuiComboPreviewData) SetBackupCursorPosPrevLine(v ImVec2) {
	C.ImGuiComboPreviewData_SetBackupCursorPosPrevLine(self.handle(), v.ToC())
}

func (self ImGuiComboPreviewData) SetBackupPrevLineTextBaseOffset(v float32) {
	C.ImGuiComboPreviewData_SetBackupPrevLineTextBaseOffset(self.handle(), C.float(v))
}

func (self ImGuiComboPreviewData) SetBackupLayout(v ImGuiLayoutType) {
	C.ImGuiComboPreviewData_SetBackupLayout(self.handle(), C.ImGuiLayoutType(v))
}

func (self ImGuiInputEvent) SetType(v ImGuiInputEventType) {
	C.ImGuiInputEvent_SetType(self.handle(), C.ImGuiInputEventType(v))
}

func (self ImGuiInputEvent) SetSource(v ImGuiInputSource) {
	C.ImGuiInputEvent_SetSource(self.handle(), C.ImGuiInputSource(v))
}

func (self ImGuiInputEvent) SetAddedByTestEngine(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputEvent_SetAddedByTestEngine(self.handle(), vArg)
}

func (self ImGuiListClipper) SetDisplayStart(v int32) {
	C.ImGuiListClipper_SetDisplayStart(self.handle(), C.int(v))
}

func (self ImGuiListClipper) SetDisplayEnd(v int32) {
	C.ImGuiListClipper_SetDisplayEnd(self.handle(), C.int(v))
}

func (self ImGuiListClipper) SetItemsCount(v int32) {
	C.ImGuiListClipper_SetItemsCount(self.handle(), C.int(v))
}

func (self ImGuiListClipper) SetItemsHeight(v float32) {
	C.ImGuiListClipper_SetItemsHeight(self.handle(), C.float(v))
}

func (self ImGuiListClipper) SetStartPosY(v float32) {
	C.ImGuiListClipper_SetStartPosY(self.handle(), C.float(v))
}

func (self ImGuiListClipper) SetTempData(v unsafe.Pointer) {
	C.ImGuiListClipper_SetTempData(self.handle(), v)
}

func (self ImGuiTableColumn) SetFlags(v ImGuiTableColumnFlags) {
	C.ImGuiTableColumn_SetFlags(self.handle(), C.ImGuiTableColumnFlags(v))
}

func (self ImGuiTableColumn) SetWidthGiven(v float32) {
	C.ImGuiTableColumn_SetWidthGiven(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetMinX(v float32) {
	C.ImGuiTableColumn_SetMinX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetMaxX(v float32) {
	C.ImGuiTableColumn_SetMaxX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetWidthRequest(v float32) {
	C.ImGuiTableColumn_SetWidthRequest(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetWidthAuto(v float32) {
	C.ImGuiTableColumn_SetWidthAuto(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetStretchWeight(v float32) {
	C.ImGuiTableColumn_SetStretchWeight(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetInitStretchWeightOrWidth(v float32) {
	C.ImGuiTableColumn_SetInitStretchWeightOrWidth(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetUserID(v ImGuiID) {
	C.ImGuiTableColumn_SetUserID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTableColumn) SetWorkMinX(v float32) {
	C.ImGuiTableColumn_SetWorkMinX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetWorkMaxX(v float32) {
	C.ImGuiTableColumn_SetWorkMaxX(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetItemWidth(v float32) {
	C.ImGuiTableColumn_SetItemWidth(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetContentMaxXFrozen(v float32) {
	C.ImGuiTableColumn_SetContentMaxXFrozen(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetContentMaxXUnfrozen(v float32) {
	C.ImGuiTableColumn_SetContentMaxXUnfrozen(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetContentMaxXHeadersUsed(v float32) {
	C.ImGuiTableColumn_SetContentMaxXHeadersUsed(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetContentMaxXHeadersIdeal(v float32) {
	C.ImGuiTableColumn_SetContentMaxXHeadersIdeal(self.handle(), C.float(v))
}

func (self ImGuiTableColumn) SetIsEnabled(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsEnabled(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsUserEnabled(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsUserEnabled(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsUserEnabledNextFrame(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsUserEnabledNextFrame(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsVisibleX(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsVisibleX(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsVisibleY(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsVisibleY(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsRequestOutput(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsRequestOutput(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsSkipItems(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsSkipItems(self.handle(), vArg)
}

func (self ImGuiTableColumn) SetIsPreserveWidthAuto(v bool) {
	vArg := C.bool(v)

	C.ImGuiTableColumn_SetIsPreserveWidthAuto(self.handle(), vArg)
}

func (self ImGuiContextHook) SetHookId(v ImGuiID) {
	C.ImGuiContextHook_SetHookId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContextHook) SetType(v ImGuiContextHookType) {
	C.ImGuiContextHook_SetType(self.handle(), C.ImGuiContextHookType(v))
}

func (self ImGuiContextHook) SetOwner(v ImGuiID) {
	C.ImGuiContextHook_SetOwner(self.handle(), C.ImGuiID(v))
}

func (self ImGuiContextHook) SetUserData(v unsafe.Pointer) {
	C.ImGuiContextHook_SetUserData(self.handle(), v)
}

func (self ImGuiMenuColumns) SetTotalWidth(v uint32) {
	C.ImGuiMenuColumns_SetTotalWidth(self.handle(), C.uint(v))
}

func (self ImGuiMenuColumns) SetNextTotalWidth(v uint32) {
	C.ImGuiMenuColumns_SetNextTotalWidth(self.handle(), C.uint(v))
}

func (self ImGuiTable) SetID(v ImGuiID) {
	C.ImGuiTable_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTable) SetFlags(v ImGuiTableFlags) {
	C.ImGuiTable_SetFlags(self.handle(), C.ImGuiTableFlags(v))
}

func (self ImGuiTable) SetRawData(v unsafe.Pointer) {
	C.ImGuiTable_SetRawData(self.handle(), v)
}

func (self ImGuiTable) SetTempData(v ImGuiTableTempData) {
	C.ImGuiTable_SetTempData(self.handle(), v.handle())
}

func (self ImGuiTable) SetSettingsLoadedFlags(v ImGuiTableFlags) {
	C.ImGuiTable_SetSettingsLoadedFlags(self.handle(), C.ImGuiTableFlags(v))
}

func (self ImGuiTable) SetSettingsOffset(v int32) {
	C.ImGuiTable_SetSettingsOffset(self.handle(), C.int(v))
}

func (self ImGuiTable) SetLastFrameActive(v int32) {
	C.ImGuiTable_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiTable) SetColumnsCount(v int32) {
	C.ImGuiTable_SetColumnsCount(self.handle(), C.int(v))
}

func (self ImGuiTable) SetCurrentRow(v int32) {
	C.ImGuiTable_SetCurrentRow(self.handle(), C.int(v))
}

func (self ImGuiTable) SetCurrentColumn(v int32) {
	C.ImGuiTable_SetCurrentColumn(self.handle(), C.int(v))
}

func (self ImGuiTable) SetRowPosY1(v float32) {
	C.ImGuiTable_SetRowPosY1(self.handle(), C.float(v))
}

func (self ImGuiTable) SetRowPosY2(v float32) {
	C.ImGuiTable_SetRowPosY2(self.handle(), C.float(v))
}

func (self ImGuiTable) SetRowMinHeight(v float32) {
	C.ImGuiTable_SetRowMinHeight(self.handle(), C.float(v))
}

func (self ImGuiTable) SetRowTextBaseline(v float32) {
	C.ImGuiTable_SetRowTextBaseline(self.handle(), C.float(v))
}

func (self ImGuiTable) SetRowIndentOffsetX(v float32) {
	C.ImGuiTable_SetRowIndentOffsetX(self.handle(), C.float(v))
}

func (self ImGuiTable) SetRowFlags(v ImGuiTableRowFlags) {
	C.ImGuiTable_SetRowFlags(self.handle(), C.ImGuiTableRowFlags(v))
}

func (self ImGuiTable) SetLastRowFlags(v ImGuiTableRowFlags) {
	C.ImGuiTable_SetLastRowFlags(self.handle(), C.ImGuiTableRowFlags(v))
}

func (self ImGuiTable) SetRowBgColorCounter(v int32) {
	C.ImGuiTable_SetRowBgColorCounter(self.handle(), C.int(v))
}

func (self ImGuiTable) SetBorderColorStrong(v uint32) {
	C.ImGuiTable_SetBorderColorStrong(self.handle(), C.uint(v))
}

func (self ImGuiTable) SetBorderColorLight(v uint32) {
	C.ImGuiTable_SetBorderColorLight(self.handle(), C.uint(v))
}

func (self ImGuiTable) SetBorderX1(v float32) {
	C.ImGuiTable_SetBorderX1(self.handle(), C.float(v))
}

func (self ImGuiTable) SetBorderX2(v float32) {
	C.ImGuiTable_SetBorderX2(self.handle(), C.float(v))
}

func (self ImGuiTable) SetHostIndentX(v float32) {
	C.ImGuiTable_SetHostIndentX(self.handle(), C.float(v))
}

func (self ImGuiTable) SetMinColumnWidth(v float32) {
	C.ImGuiTable_SetMinColumnWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) SetOuterPaddingX(v float32) {
	C.ImGuiTable_SetOuterPaddingX(self.handle(), C.float(v))
}

func (self ImGuiTable) SetCellPaddingX(v float32) {
	C.ImGuiTable_SetCellPaddingX(self.handle(), C.float(v))
}

func (self ImGuiTable) SetCellPaddingY(v float32) {
	C.ImGuiTable_SetCellPaddingY(self.handle(), C.float(v))
}

func (self ImGuiTable) SetCellSpacingX1(v float32) {
	C.ImGuiTable_SetCellSpacingX1(self.handle(), C.float(v))
}

func (self ImGuiTable) SetCellSpacingX2(v float32) {
	C.ImGuiTable_SetCellSpacingX2(self.handle(), C.float(v))
}

func (self ImGuiTable) SetInnerWidth(v float32) {
	C.ImGuiTable_SetInnerWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) SetColumnsGivenWidth(v float32) {
	C.ImGuiTable_SetColumnsGivenWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) SetColumnsAutoFitWidth(v float32) {
	C.ImGuiTable_SetColumnsAutoFitWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) SetColumnsStretchSumWeights(v float32) {
	C.ImGuiTable_SetColumnsStretchSumWeights(self.handle(), C.float(v))
}

func (self ImGuiTable) SetResizedColumnNextWidth(v float32) {
	C.ImGuiTable_SetResizedColumnNextWidth(self.handle(), C.float(v))
}

func (self ImGuiTable) SetResizeLockMinContentsX2(v float32) {
	C.ImGuiTable_SetResizeLockMinContentsX2(self.handle(), C.float(v))
}

func (self ImGuiTable) SetRefScale(v float32) {
	C.ImGuiTable_SetRefScale(self.handle(), C.float(v))
}

func (self ImGuiTable) SetOuterWindow(v ImGuiWindow) {
	C.ImGuiTable_SetOuterWindow(self.handle(), v.handle())
}

func (self ImGuiTable) SetInnerWindow(v ImGuiWindow) {
	C.ImGuiTable_SetInnerWindow(self.handle(), v.handle())
}

func (self ImGuiTable) SetDrawSplitter(v ImDrawListSplitter) {
	C.ImGuiTable_SetDrawSplitter(self.handle(), v.handle())
}

func (self ImGuiTable) SetIsLayoutLocked(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsLayoutLocked(self.handle(), vArg)
}

func (self ImGuiTable) SetIsInsideRow(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsInsideRow(self.handle(), vArg)
}

func (self ImGuiTable) SetIsInitializing(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsInitializing(self.handle(), vArg)
}

func (self ImGuiTable) SetIsSortSpecsDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsSortSpecsDirty(self.handle(), vArg)
}

func (self ImGuiTable) SetIsUsingHeaders(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsUsingHeaders(self.handle(), vArg)
}

func (self ImGuiTable) SetIsContextPopupOpen(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsContextPopupOpen(self.handle(), vArg)
}

func (self ImGuiTable) SetIsSettingsRequestLoad(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsSettingsRequestLoad(self.handle(), vArg)
}

func (self ImGuiTable) SetIsSettingsDirty(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsSettingsDirty(self.handle(), vArg)
}

func (self ImGuiTable) SetIsDefaultDisplayOrder(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsDefaultDisplayOrder(self.handle(), vArg)
}

func (self ImGuiTable) SetIsResetAllRequest(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsResetAllRequest(self.handle(), vArg)
}

func (self ImGuiTable) SetIsResetDisplayOrderRequest(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsResetDisplayOrderRequest(self.handle(), vArg)
}

func (self ImGuiTable) SetIsUnfrozenRows(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsUnfrozenRows(self.handle(), vArg)
}

func (self ImGuiTable) SetIsDefaultSizingPolicy(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetIsDefaultSizingPolicy(self.handle(), vArg)
}

func (self ImGuiTable) SetMemoryCompacted(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetMemoryCompacted(self.handle(), vArg)
}

func (self ImGuiTable) SetHostSkipItems(v bool) {
	vArg := C.bool(v)

	C.ImGuiTable_SetHostSkipItems(self.handle(), vArg)
}

func (self ImDrawCmd) SetClipRect(v ImVec4) {
	C.ImDrawCmd_SetClipRect(self.handle(), v.ToC())
}

func (self ImDrawCmd) SetTextureId(v ImTextureID) {
	C.ImDrawCmd_SetTextureId(self.handle(), C.ImTextureID(v))
}

func (self ImDrawCmd) SetVtxOffset(v uint32) {
	C.ImDrawCmd_SetVtxOffset(self.handle(), C.uint(v))
}

func (self ImDrawCmd) SetIdxOffset(v uint32) {
	C.ImDrawCmd_SetIdxOffset(self.handle(), C.uint(v))
}

func (self ImDrawCmd) SetElemCount(v uint32) {
	C.ImDrawCmd_SetElemCount(self.handle(), C.uint(v))
}

func (self ImDrawCmd) SetUserCallbackData(v unsafe.Pointer) {
	C.ImDrawCmd_SetUserCallbackData(self.handle(), v)
}

func (self ImGuiColorMod) SetCol(v ImGuiCol) {
	C.ImGuiColorMod_SetCol(self.handle(), C.ImGuiCol(v))
}

func (self ImGuiColorMod) SetBackupValue(v ImVec4) {
	C.ImGuiColorMod_SetBackupValue(self.handle(), v.ToC())
}

func (self ImFontConfig) SetFontData(v unsafe.Pointer) {
	C.ImFontConfig_SetFontData(self.handle(), v)
}

func (self ImFontConfig) SetFontDataSize(v int32) {
	C.ImFontConfig_SetFontDataSize(self.handle(), C.int(v))
}

func (self ImFontConfig) SetFontDataOwnedByAtlas(v bool) {
	vArg := C.bool(v)

	C.ImFontConfig_SetFontDataOwnedByAtlas(self.handle(), vArg)
}

func (self ImFontConfig) SetFontNo(v int32) {
	C.ImFontConfig_SetFontNo(self.handle(), C.int(v))
}

func (self ImFontConfig) SetSizePixels(v float32) {
	C.ImFontConfig_SetSizePixels(self.handle(), C.float(v))
}

func (self ImFontConfig) SetOversampleH(v int32) {
	C.ImFontConfig_SetOversampleH(self.handle(), C.int(v))
}

func (self ImFontConfig) SetOversampleV(v int32) {
	C.ImFontConfig_SetOversampleV(self.handle(), C.int(v))
}

func (self ImFontConfig) SetPixelSnapH(v bool) {
	vArg := C.bool(v)

	C.ImFontConfig_SetPixelSnapH(self.handle(), vArg)
}

func (self ImFontConfig) SetGlyphExtraSpacing(v ImVec2) {
	C.ImFontConfig_SetGlyphExtraSpacing(self.handle(), v.ToC())
}

func (self ImFontConfig) SetGlyphOffset(v ImVec2) {
	C.ImFontConfig_SetGlyphOffset(self.handle(), v.ToC())
}

func (self ImFontConfig) SetGlyphRanges(v *ImWchar) {
	C.ImFontConfig_SetGlyphRanges(self.handle(), (*C.ImWchar)(v))
}

func (self ImFontConfig) SetGlyphMinAdvanceX(v float32) {
	C.ImFontConfig_SetGlyphMinAdvanceX(self.handle(), C.float(v))
}

func (self ImFontConfig) SetGlyphMaxAdvanceX(v float32) {
	C.ImFontConfig_SetGlyphMaxAdvanceX(self.handle(), C.float(v))
}

func (self ImFontConfig) SetMergeMode(v bool) {
	vArg := C.bool(v)

	C.ImFontConfig_SetMergeMode(self.handle(), vArg)
}

func (self ImFontConfig) SetFontBuilderFlags(v uint32) {
	C.ImFontConfig_SetFontBuilderFlags(self.handle(), C.uint(v))
}

func (self ImFontConfig) SetRasterizerMultiply(v float32) {
	C.ImFontConfig_SetRasterizerMultiply(self.handle(), C.float(v))
}

func (self ImFontConfig) SetEllipsisChar(v ImWchar) {
	C.ImFontConfig_SetEllipsisChar(self.handle(), C.ImWchar(v))
}

func (self ImFontConfig) SetDstFont(v ImFont) {
	C.ImFontConfig_SetDstFont(self.handle(), v.handle())
}

func (self ImGuiListClipperRange) SetMin(v int32) {
	C.ImGuiListClipperRange_SetMin(self.handle(), C.int(v))
}

func (self ImGuiListClipperRange) SetMax(v int32) {
	C.ImGuiListClipperRange_SetMax(self.handle(), C.int(v))
}

func (self ImGuiListClipperRange) SetPosToIndexConvert(v bool) {
	vArg := C.bool(v)

	C.ImGuiListClipperRange_SetPosToIndexConvert(self.handle(), vArg)
}

func (self ImGuiNextWindowData) SetFlags(v ImGuiNextWindowDataFlags) {
	C.ImGuiNextWindowData_SetFlags(self.handle(), C.ImGuiNextWindowDataFlags(v))
}

func (self ImGuiNextWindowData) SetPosCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetPosCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) SetSizeCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetSizeCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) SetCollapsedCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetCollapsedCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) SetDockCond(v ImGuiCond) {
	C.ImGuiNextWindowData_SetDockCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextWindowData) SetPosVal(v ImVec2) {
	C.ImGuiNextWindowData_SetPosVal(self.handle(), v.ToC())
}

func (self ImGuiNextWindowData) SetPosPivotVal(v ImVec2) {
	C.ImGuiNextWindowData_SetPosPivotVal(self.handle(), v.ToC())
}

func (self ImGuiNextWindowData) SetSizeVal(v ImVec2) {
	C.ImGuiNextWindowData_SetSizeVal(self.handle(), v.ToC())
}

func (self ImGuiNextWindowData) SetContentSizeVal(v ImVec2) {
	C.ImGuiNextWindowData_SetContentSizeVal(self.handle(), v.ToC())
}

func (self ImGuiNextWindowData) SetScrollVal(v ImVec2) {
	C.ImGuiNextWindowData_SetScrollVal(self.handle(), v.ToC())
}

func (self ImGuiNextWindowData) SetPosUndock(v bool) {
	vArg := C.bool(v)

	C.ImGuiNextWindowData_SetPosUndock(self.handle(), vArg)
}

func (self ImGuiNextWindowData) SetCollapsedVal(v bool) {
	vArg := C.bool(v)

	C.ImGuiNextWindowData_SetCollapsedVal(self.handle(), vArg)
}

func (self ImGuiNextWindowData) SetSizeCallbackUserData(v unsafe.Pointer) {
	C.ImGuiNextWindowData_SetSizeCallbackUserData(self.handle(), v)
}

func (self ImGuiNextWindowData) SetBgAlphaVal(v float32) {
	C.ImGuiNextWindowData_SetBgAlphaVal(self.handle(), C.float(v))
}

func (self ImGuiNextWindowData) SetViewportId(v ImGuiID) {
	C.ImGuiNextWindowData_SetViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNextWindowData) SetDockId(v ImGuiID) {
	C.ImGuiNextWindowData_SetDockId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNextWindowData) SetMenuBarOffsetMinVal(v ImVec2) {
	C.ImGuiNextWindowData_SetMenuBarOffsetMinVal(self.handle(), v.ToC())
}

func (self ImDrawListSplitter) Set_Current(v int32) {
	C.ImDrawListSplitter_Set_Current(self.handle(), C.int(v))
}

func (self ImDrawListSplitter) Set_Count(v int32) {
	C.ImDrawListSplitter_Set_Count(self.handle(), C.int(v))
}

func (self ImGuiTextRange) Setb(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiTextRange_Setb(self.handle(), vArg)
}

func (self ImGuiTextRange) Sete(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiTextRange_Sete(self.handle(), vArg)
}

func (self ImGuiInputEventKey) SetKey(v ImGuiKey) {
	C.ImGuiInputEventKey_SetKey(self.handle(), C.ImGuiKey(v))
}

func (self ImGuiInputEventKey) SetDown(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputEventKey_SetDown(self.handle(), vArg)
}

func (self ImGuiInputEventKey) SetAnalogValue(v float32) {
	C.ImGuiInputEventKey_SetAnalogValue(self.handle(), C.float(v))
}

func (self ImGuiStackTool) SetLastActiveFrame(v int32) {
	C.ImGuiStackTool_SetLastActiveFrame(self.handle(), C.int(v))
}

func (self ImGuiStackTool) SetStackLevel(v int32) {
	C.ImGuiStackTool_SetStackLevel(self.handle(), C.int(v))
}

func (self ImGuiStackTool) SetQueryId(v ImGuiID) {
	C.ImGuiStackTool_SetQueryId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiStackTool) SetCopyToClipboardOnCtrlC(v bool) {
	vArg := C.bool(v)

	C.ImGuiStackTool_SetCopyToClipboardOnCtrlC(self.handle(), vArg)
}

func (self ImGuiStackTool) SetCopyToClipboardLastTime(v float32) {
	C.ImGuiStackTool_SetCopyToClipboardLastTime(self.handle(), C.float(v))
}

func (self ImGuiLastItemData) SetID(v ImGuiID) {
	C.ImGuiLastItemData_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiLastItemData) SetInFlags(v ImGuiItemFlags) {
	C.ImGuiLastItemData_SetInFlags(self.handle(), C.ImGuiItemFlags(v))
}

func (self ImGuiLastItemData) SetStatusFlags(v ImGuiItemStatusFlags) {
	C.ImGuiLastItemData_SetStatusFlags(self.handle(), C.ImGuiItemStatusFlags(v))
}

func (self ImGuiMetricsConfig) SetShowDebugLog(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowDebugLog(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowStackTool(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowStackTool(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowWindowsRects(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowWindowsRects(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowWindowsBeginOrder(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowWindowsBeginOrder(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowTablesRects(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowTablesRects(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowDrawCmdMesh(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowDrawCmdMesh(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowDrawCmdBoundingBoxes(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowDrawCmdBoundingBoxes(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowDockingNodes(v bool) {
	vArg := C.bool(v)

	C.ImGuiMetricsConfig_SetShowDockingNodes(self.handle(), vArg)
}

func (self ImGuiMetricsConfig) SetShowWindowsRectsType(v int32) {
	C.ImGuiMetricsConfig_SetShowWindowsRectsType(self.handle(), C.int(v))
}

func (self ImGuiMetricsConfig) SetShowTablesRectsType(v int32) {
	C.ImGuiMetricsConfig_SetShowTablesRectsType(self.handle(), C.int(v))
}

func (self ImGuiNextItemData) SetFlags(v ImGuiNextItemDataFlags) {
	C.ImGuiNextItemData_SetFlags(self.handle(), C.ImGuiNextItemDataFlags(v))
}

func (self ImGuiNextItemData) SetWidth(v float32) {
	C.ImGuiNextItemData_SetWidth(self.handle(), C.float(v))
}

func (self ImGuiNextItemData) SetFocusScopeId(v ImGuiID) {
	C.ImGuiNextItemData_SetFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNextItemData) SetOpenCond(v ImGuiCond) {
	C.ImGuiNextItemData_SetOpenCond(self.handle(), C.ImGuiCond(v))
}

func (self ImGuiNextItemData) SetOpenVal(v bool) {
	vArg := C.bool(v)

	C.ImGuiNextItemData_SetOpenVal(self.handle(), vArg)
}

func (self ImGuiOldColumns) SetID(v ImGuiID) {
	C.ImGuiOldColumns_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiOldColumns) SetFlags(v ImGuiOldColumnFlags) {
	C.ImGuiOldColumns_SetFlags(self.handle(), C.ImGuiOldColumnFlags(v))
}

func (self ImGuiOldColumns) SetIsFirstFrame(v bool) {
	vArg := C.bool(v)

	C.ImGuiOldColumns_SetIsFirstFrame(self.handle(), vArg)
}

func (self ImGuiOldColumns) SetIsBeingResized(v bool) {
	vArg := C.bool(v)

	C.ImGuiOldColumns_SetIsBeingResized(self.handle(), vArg)
}

func (self ImGuiOldColumns) SetCurrent(v int32) {
	C.ImGuiOldColumns_SetCurrent(self.handle(), C.int(v))
}

func (self ImGuiOldColumns) SetCount(v int32) {
	C.ImGuiOldColumns_SetCount(self.handle(), C.int(v))
}

func (self ImGuiOldColumns) SetOffMinX(v float32) {
	C.ImGuiOldColumns_SetOffMinX(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) SetOffMaxX(v float32) {
	C.ImGuiOldColumns_SetOffMaxX(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) SetLineMinY(v float32) {
	C.ImGuiOldColumns_SetLineMinY(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) SetLineMaxY(v float32) {
	C.ImGuiOldColumns_SetLineMaxY(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) SetHostCursorPosY(v float32) {
	C.ImGuiOldColumns_SetHostCursorPosY(self.handle(), C.float(v))
}

func (self ImGuiOldColumns) SetHostCursorMaxPosX(v float32) {
	C.ImGuiOldColumns_SetHostCursorMaxPosX(self.handle(), C.float(v))
}

func (self ImGuiWindowSettings) SetID(v ImGuiID) {
	C.ImGuiWindowSettings_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) SetViewportId(v ImGuiID) {
	C.ImGuiWindowSettings_SetViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) SetDockId(v ImGuiID) {
	C.ImGuiWindowSettings_SetDockId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) SetClassId(v ImGuiID) {
	C.ImGuiWindowSettings_SetClassId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiWindowSettings) SetDockOrder(v int) {
	vArg := C.short(v)

	C.ImGuiWindowSettings_SetDockOrder(self.handle(), vArg)
}

func (self ImGuiWindowSettings) SetCollapsed(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowSettings_SetCollapsed(self.handle(), vArg)
}

func (self ImGuiWindowSettings) SetWantApply(v bool) {
	vArg := C.bool(v)

	C.ImGuiWindowSettings_SetWantApply(self.handle(), vArg)
}

func (self ImGuiDataTypeInfo) SetSize(v uint64) {
	C.ImGuiDataTypeInfo_SetSize(self.handle(), C.ulong(v))
}

func (self ImGuiDataTypeInfo) SetName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiDataTypeInfo_SetName(self.handle(), vArg)
}

func (self ImGuiDataTypeInfo) SetPrintFmt(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiDataTypeInfo_SetPrintFmt(self.handle(), vArg)
}

func (self ImGuiDataTypeInfo) SetScanFmt(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiDataTypeInfo_SetScanFmt(self.handle(), vArg)
}

func (self ImGuiStyle) SetAlpha(v float32) {
	C.ImGuiStyle_SetAlpha(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetDisabledAlpha(v float32) {
	C.ImGuiStyle_SetDisabledAlpha(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetWindowPadding(v ImVec2) {
	C.ImGuiStyle_SetWindowPadding(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetWindowRounding(v float32) {
	C.ImGuiStyle_SetWindowRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetWindowBorderSize(v float32) {
	C.ImGuiStyle_SetWindowBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetWindowMinSize(v ImVec2) {
	C.ImGuiStyle_SetWindowMinSize(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetWindowTitleAlign(v ImVec2) {
	C.ImGuiStyle_SetWindowTitleAlign(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetWindowMenuButtonPosition(v ImGuiDir) {
	C.ImGuiStyle_SetWindowMenuButtonPosition(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiStyle) SetChildRounding(v float32) {
	C.ImGuiStyle_SetChildRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetChildBorderSize(v float32) {
	C.ImGuiStyle_SetChildBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetPopupRounding(v float32) {
	C.ImGuiStyle_SetPopupRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetPopupBorderSize(v float32) {
	C.ImGuiStyle_SetPopupBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetFramePadding(v ImVec2) {
	C.ImGuiStyle_SetFramePadding(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetFrameRounding(v float32) {
	C.ImGuiStyle_SetFrameRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetFrameBorderSize(v float32) {
	C.ImGuiStyle_SetFrameBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetItemSpacing(v ImVec2) {
	C.ImGuiStyle_SetItemSpacing(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetItemInnerSpacing(v ImVec2) {
	C.ImGuiStyle_SetItemInnerSpacing(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetCellPadding(v ImVec2) {
	C.ImGuiStyle_SetCellPadding(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetTouchExtraPadding(v ImVec2) {
	C.ImGuiStyle_SetTouchExtraPadding(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetIndentSpacing(v float32) {
	C.ImGuiStyle_SetIndentSpacing(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetColumnsMinSpacing(v float32) {
	C.ImGuiStyle_SetColumnsMinSpacing(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetScrollbarSize(v float32) {
	C.ImGuiStyle_SetScrollbarSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetScrollbarRounding(v float32) {
	C.ImGuiStyle_SetScrollbarRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetGrabMinSize(v float32) {
	C.ImGuiStyle_SetGrabMinSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetGrabRounding(v float32) {
	C.ImGuiStyle_SetGrabRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetLogSliderDeadzone(v float32) {
	C.ImGuiStyle_SetLogSliderDeadzone(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetTabRounding(v float32) {
	C.ImGuiStyle_SetTabRounding(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetTabBorderSize(v float32) {
	C.ImGuiStyle_SetTabBorderSize(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetTabMinWidthForCloseButton(v float32) {
	C.ImGuiStyle_SetTabMinWidthForCloseButton(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetColorButtonPosition(v ImGuiDir) {
	C.ImGuiStyle_SetColorButtonPosition(self.handle(), C.ImGuiDir(v))
}

func (self ImGuiStyle) SetButtonTextAlign(v ImVec2) {
	C.ImGuiStyle_SetButtonTextAlign(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetSelectableTextAlign(v ImVec2) {
	C.ImGuiStyle_SetSelectableTextAlign(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetDisplayWindowPadding(v ImVec2) {
	C.ImGuiStyle_SetDisplayWindowPadding(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetDisplaySafeAreaPadding(v ImVec2) {
	C.ImGuiStyle_SetDisplaySafeAreaPadding(self.handle(), v.ToC())
}

func (self ImGuiStyle) SetMouseCursorScale(v float32) {
	C.ImGuiStyle_SetMouseCursorScale(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetAntiAliasedLines(v bool) {
	vArg := C.bool(v)

	C.ImGuiStyle_SetAntiAliasedLines(self.handle(), vArg)
}

func (self ImGuiStyle) SetAntiAliasedLinesUseTex(v bool) {
	vArg := C.bool(v)

	C.ImGuiStyle_SetAntiAliasedLinesUseTex(self.handle(), vArg)
}

func (self ImGuiStyle) SetAntiAliasedFill(v bool) {
	vArg := C.bool(v)

	C.ImGuiStyle_SetAntiAliasedFill(self.handle(), vArg)
}

func (self ImGuiStyle) SetCurveTessellationTol(v float32) {
	C.ImGuiStyle_SetCurveTessellationTol(self.handle(), C.float(v))
}

func (self ImGuiStyle) SetCircleTessellationMaxError(v float32) {
	C.ImGuiStyle_SetCircleTessellationMaxError(self.handle(), C.float(v))
}

func (self ImGuiTableInstanceData) SetLastOuterHeight(v float32) {
	C.ImGuiTableInstanceData_SetLastOuterHeight(self.handle(), C.float(v))
}

func (self ImGuiTableInstanceData) SetLastFirstRowHeight(v float32) {
	C.ImGuiTableInstanceData_SetLastFirstRowHeight(self.handle(), C.float(v))
}

func (self ImGuiOnceUponAFrame) SetRefFrame(v int32) {
	C.ImGuiOnceUponAFrame_SetRefFrame(self.handle(), C.int(v))
}

func (self ImGuiStoragePair) Setkey(v ImGuiID) {
	C.ImGuiStoragePair_Setkey(self.handle(), C.ImGuiID(v))
}

func (self ImGuiInputEventText) SetChar(v uint32) {
	C.ImGuiInputEventText_SetChar(self.handle(), C.uint(v))
}

func (self ImGuiTabBar) SetFlags(v ImGuiTabBarFlags) {
	C.ImGuiTabBar_SetFlags(self.handle(), C.ImGuiTabBarFlags(v))
}

func (self ImGuiTabBar) SetID(v ImGuiID) {
	C.ImGuiTabBar_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) SetSelectedTabId(v ImGuiID) {
	C.ImGuiTabBar_SetSelectedTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) SetNextSelectedTabId(v ImGuiID) {
	C.ImGuiTabBar_SetNextSelectedTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) SetVisibleTabId(v ImGuiID) {
	C.ImGuiTabBar_SetVisibleTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) SetCurrFrameVisible(v int32) {
	C.ImGuiTabBar_SetCurrFrameVisible(self.handle(), C.int(v))
}

func (self ImGuiTabBar) SetPrevFrameVisible(v int32) {
	C.ImGuiTabBar_SetPrevFrameVisible(self.handle(), C.int(v))
}

func (self ImGuiTabBar) SetCurrTabsContentsHeight(v float32) {
	C.ImGuiTabBar_SetCurrTabsContentsHeight(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetPrevTabsContentsHeight(v float32) {
	C.ImGuiTabBar_SetPrevTabsContentsHeight(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetWidthAllTabs(v float32) {
	C.ImGuiTabBar_SetWidthAllTabs(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetWidthAllTabsIdeal(v float32) {
	C.ImGuiTabBar_SetWidthAllTabsIdeal(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetScrollingAnim(v float32) {
	C.ImGuiTabBar_SetScrollingAnim(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetScrollingTarget(v float32) {
	C.ImGuiTabBar_SetScrollingTarget(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetScrollingTargetDistToVisibility(v float32) {
	C.ImGuiTabBar_SetScrollingTargetDistToVisibility(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetScrollingSpeed(v float32) {
	C.ImGuiTabBar_SetScrollingSpeed(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetScrollingRectMinX(v float32) {
	C.ImGuiTabBar_SetScrollingRectMinX(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetScrollingRectMaxX(v float32) {
	C.ImGuiTabBar_SetScrollingRectMaxX(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetReorderRequestTabId(v ImGuiID) {
	C.ImGuiTabBar_SetReorderRequestTabId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiTabBar) SetWantLayout(v bool) {
	vArg := C.bool(v)

	C.ImGuiTabBar_SetWantLayout(self.handle(), vArg)
}

func (self ImGuiTabBar) SetVisibleTabWasSubmitted(v bool) {
	vArg := C.bool(v)

	C.ImGuiTabBar_SetVisibleTabWasSubmitted(self.handle(), vArg)
}

func (self ImGuiTabBar) SetTabsAddedNew(v bool) {
	vArg := C.bool(v)

	C.ImGuiTabBar_SetTabsAddedNew(self.handle(), vArg)
}

func (self ImGuiTabBar) SetItemSpacingY(v float32) {
	C.ImGuiTabBar_SetItemSpacingY(self.handle(), C.float(v))
}

func (self ImGuiTabBar) SetFramePadding(v ImVec2) {
	C.ImGuiTabBar_SetFramePadding(self.handle(), v.ToC())
}

func (self ImGuiTabBar) SetBackupCursorPos(v ImVec2) {
	C.ImGuiTabBar_SetBackupCursorPos(self.handle(), v.ToC())
}

func (self ImGuiWindowStackData) SetWindow(v ImGuiWindow) {
	C.ImGuiWindowStackData_SetWindow(self.handle(), v.handle())
}

func (self ImFontAtlas) SetFlags(v ImFontAtlasFlags) {
	C.ImFontAtlas_SetFlags(self.handle(), C.ImFontAtlasFlags(v))
}

func (self ImFontAtlas) SetTexDesiredWidth(v int32) {
	C.ImFontAtlas_SetTexDesiredWidth(self.handle(), C.int(v))
}

func (self ImFontAtlas) SetTexGlyphPadding(v int32) {
	C.ImFontAtlas_SetTexGlyphPadding(self.handle(), C.int(v))
}

func (self ImFontAtlas) SetLocked(v bool) {
	vArg := C.bool(v)

	C.ImFontAtlas_SetLocked(self.handle(), vArg)
}

func (self ImFontAtlas) SetTexReady(v bool) {
	vArg := C.bool(v)

	C.ImFontAtlas_SetTexReady(self.handle(), vArg)
}

func (self ImFontAtlas) SetTexPixelsUseColors(v bool) {
	vArg := C.bool(v)

	C.ImFontAtlas_SetTexPixelsUseColors(self.handle(), vArg)
}

func (self ImFontAtlas) SetTexWidth(v int32) {
	C.ImFontAtlas_SetTexWidth(self.handle(), C.int(v))
}

func (self ImFontAtlas) SetTexHeight(v int32) {
	C.ImFontAtlas_SetTexHeight(self.handle(), C.int(v))
}

func (self ImFontAtlas) SetTexUvScale(v ImVec2) {
	C.ImFontAtlas_SetTexUvScale(self.handle(), v.ToC())
}

func (self ImFontAtlas) SetTexUvWhitePixel(v ImVec2) {
	C.ImFontAtlas_SetTexUvWhitePixel(self.handle(), v.ToC())
}

func (self ImFontAtlas) SetFontBuilderIO(v ImFontBuilderIO) {
	C.ImFontAtlas_SetFontBuilderIO(self.handle(), v.handle())
}

func (self ImFontAtlas) SetFontBuilderFlags(v uint32) {
	C.ImFontAtlas_SetFontBuilderFlags(self.handle(), C.uint(v))
}

func (self ImFontAtlas) SetPackIdMouseCursors(v int32) {
	C.ImFontAtlas_SetPackIdMouseCursors(self.handle(), C.int(v))
}

func (self ImFontAtlas) SetPackIdLines(v int32) {
	C.ImFontAtlas_SetPackIdLines(self.handle(), C.int(v))
}

func (self ImFontGlyph) SetColored(v uint32) {
	C.ImFontGlyph_SetColored(self.handle(), C.uint(v))
}

func (self ImFontGlyph) SetVisible(v uint32) {
	C.ImFontGlyph_SetVisible(self.handle(), C.uint(v))
}

func (self ImFontGlyph) SetCodepoint(v uint32) {
	C.ImFontGlyph_SetCodepoint(self.handle(), C.uint(v))
}

func (self ImFontGlyph) SetAdvanceX(v float32) {
	C.ImFontGlyph_SetAdvanceX(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetX0(v float32) {
	C.ImFontGlyph_SetX0(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetY0(v float32) {
	C.ImFontGlyph_SetY0(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetX1(v float32) {
	C.ImFontGlyph_SetX1(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetY1(v float32) {
	C.ImFontGlyph_SetY1(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetU0(v float32) {
	C.ImFontGlyph_SetU0(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetV0(v float32) {
	C.ImFontGlyph_SetV0(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetU1(v float32) {
	C.ImFontGlyph_SetU1(self.handle(), C.float(v))
}

func (self ImFontGlyph) SetV1(v float32) {
	C.ImFontGlyph_SetV1(self.handle(), C.float(v))
}

func (self ImGuiIO) SetConfigFlags(v ImGuiConfigFlags) {
	C.ImGuiIO_SetConfigFlags(self.handle(), C.ImGuiConfigFlags(v))
}

func (self ImGuiIO) SetBackendFlags(v ImGuiBackendFlags) {
	C.ImGuiIO_SetBackendFlags(self.handle(), C.ImGuiBackendFlags(v))
}

func (self ImGuiIO) SetDisplaySize(v ImVec2) {
	C.ImGuiIO_SetDisplaySize(self.handle(), v.ToC())
}

func (self ImGuiIO) SetDeltaTime(v float32) {
	C.ImGuiIO_SetDeltaTime(self.handle(), C.float(v))
}

func (self ImGuiIO) SetIniSavingRate(v float32) {
	C.ImGuiIO_SetIniSavingRate(self.handle(), C.float(v))
}

func (self ImGuiIO) SetIniFilename(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetIniFilename(self.handle(), vArg)
}

func (self ImGuiIO) SetLogFilename(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetLogFilename(self.handle(), vArg)
}

func (self ImGuiIO) SetMouseDoubleClickTime(v float32) {
	C.ImGuiIO_SetMouseDoubleClickTime(self.handle(), C.float(v))
}

func (self ImGuiIO) SetMouseDoubleClickMaxDist(v float32) {
	C.ImGuiIO_SetMouseDoubleClickMaxDist(self.handle(), C.float(v))
}

func (self ImGuiIO) SetMouseDragThreshold(v float32) {
	C.ImGuiIO_SetMouseDragThreshold(self.handle(), C.float(v))
}

func (self ImGuiIO) SetKeyRepeatDelay(v float32) {
	C.ImGuiIO_SetKeyRepeatDelay(self.handle(), C.float(v))
}

func (self ImGuiIO) SetKeyRepeatRate(v float32) {
	C.ImGuiIO_SetKeyRepeatRate(self.handle(), C.float(v))
}

func (self ImGuiIO) SetUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetUserData(self.handle(), v)
}

func (self ImGuiIO) SetFonts(v ImFontAtlas) {
	C.ImGuiIO_SetFonts(self.handle(), v.handle())
}

func (self ImGuiIO) SetFontGlobalScale(v float32) {
	C.ImGuiIO_SetFontGlobalScale(self.handle(), C.float(v))
}

func (self ImGuiIO) SetFontAllowUserScaling(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetFontAllowUserScaling(self.handle(), vArg)
}

func (self ImGuiIO) SetFontDefault(v ImFont) {
	C.ImGuiIO_SetFontDefault(self.handle(), v.handle())
}

func (self ImGuiIO) SetDisplayFramebufferScale(v ImVec2) {
	C.ImGuiIO_SetDisplayFramebufferScale(self.handle(), v.ToC())
}

func (self ImGuiIO) SetConfigDockingNoSplit(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigDockingNoSplit(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigDockingWithShift(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigDockingWithShift(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigDockingAlwaysTabBar(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigDockingAlwaysTabBar(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigDockingTransparentPayload(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigDockingTransparentPayload(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigViewportsNoAutoMerge(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigViewportsNoAutoMerge(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigViewportsNoTaskBarIcon(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigViewportsNoTaskBarIcon(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigViewportsNoDecoration(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigViewportsNoDecoration(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigViewportsNoDefaultParent(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigViewportsNoDefaultParent(self.handle(), vArg)
}

func (self ImGuiIO) SetMouseDrawCursor(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetMouseDrawCursor(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigMacOSXBehaviors(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigMacOSXBehaviors(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigInputTrickleEventQueue(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigInputTrickleEventQueue(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigInputTextCursorBlink(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigInputTextCursorBlink(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigInputTextEnterKeepActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigInputTextEnterKeepActive(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigDragClickToInputText(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigDragClickToInputText(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigWindowsResizeFromEdges(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigWindowsResizeFromEdges(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigWindowsMoveFromTitleBarOnly(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetConfigWindowsMoveFromTitleBarOnly(self.handle(), vArg)
}

func (self ImGuiIO) SetConfigMemoryCompactTimer(v float32) {
	C.ImGuiIO_SetConfigMemoryCompactTimer(self.handle(), C.float(v))
}

func (self ImGuiIO) SetBackendPlatformName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetBackendPlatformName(self.handle(), vArg)
}

func (self ImGuiIO) SetBackendRendererName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImGuiIO_SetBackendRendererName(self.handle(), vArg)
}

func (self ImGuiIO) SetBackendPlatformUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetBackendPlatformUserData(self.handle(), v)
}

func (self ImGuiIO) SetBackendRendererUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetBackendRendererUserData(self.handle(), v)
}

func (self ImGuiIO) SetBackendLanguageUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetBackendLanguageUserData(self.handle(), v)
}

func (self ImGuiIO) SetClipboardUserData(v unsafe.Pointer) {
	C.ImGuiIO_SetClipboardUserData(self.handle(), v)
}

func (self ImGuiIO) Set_UnusedPadding(v unsafe.Pointer) {
	C.ImGuiIO_Set_UnusedPadding(self.handle(), v)
}

func (self ImGuiIO) SetWantCaptureMouse(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetWantCaptureMouse(self.handle(), vArg)
}

func (self ImGuiIO) SetWantCaptureKeyboard(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetWantCaptureKeyboard(self.handle(), vArg)
}

func (self ImGuiIO) SetWantTextInput(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetWantTextInput(self.handle(), vArg)
}

func (self ImGuiIO) SetWantSetMousePos(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetWantSetMousePos(self.handle(), vArg)
}

func (self ImGuiIO) SetWantSaveIniSettings(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetWantSaveIniSettings(self.handle(), vArg)
}

func (self ImGuiIO) SetNavActive(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetNavActive(self.handle(), vArg)
}

func (self ImGuiIO) SetNavVisible(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetNavVisible(self.handle(), vArg)
}

func (self ImGuiIO) SetFramerate(v float32) {
	C.ImGuiIO_SetFramerate(self.handle(), C.float(v))
}

func (self ImGuiIO) SetMetricsRenderVertices(v int32) {
	C.ImGuiIO_SetMetricsRenderVertices(self.handle(), C.int(v))
}

func (self ImGuiIO) SetMetricsRenderIndices(v int32) {
	C.ImGuiIO_SetMetricsRenderIndices(self.handle(), C.int(v))
}

func (self ImGuiIO) SetMetricsRenderWindows(v int32) {
	C.ImGuiIO_SetMetricsRenderWindows(self.handle(), C.int(v))
}

func (self ImGuiIO) SetMetricsActiveWindows(v int32) {
	C.ImGuiIO_SetMetricsActiveWindows(self.handle(), C.int(v))
}

func (self ImGuiIO) SetMetricsActiveAllocations(v int32) {
	C.ImGuiIO_SetMetricsActiveAllocations(self.handle(), C.int(v))
}

func (self ImGuiIO) SetMouseDelta(v ImVec2) {
	C.ImGuiIO_SetMouseDelta(self.handle(), v.ToC())
}

func (self ImGuiIO) SetMousePos(v ImVec2) {
	C.ImGuiIO_SetMousePos(self.handle(), v.ToC())
}

func (self ImGuiIO) SetMouseWheel(v float32) {
	C.ImGuiIO_SetMouseWheel(self.handle(), C.float(v))
}

func (self ImGuiIO) SetMouseWheelH(v float32) {
	C.ImGuiIO_SetMouseWheelH(self.handle(), C.float(v))
}

func (self ImGuiIO) SetMouseHoveredViewport(v ImGuiID) {
	C.ImGuiIO_SetMouseHoveredViewport(self.handle(), C.ImGuiID(v))
}

func (self ImGuiIO) SetKeyCtrl(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetKeyCtrl(self.handle(), vArg)
}

func (self ImGuiIO) SetKeyShift(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetKeyShift(self.handle(), vArg)
}

func (self ImGuiIO) SetKeyAlt(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetKeyAlt(self.handle(), vArg)
}

func (self ImGuiIO) SetKeySuper(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetKeySuper(self.handle(), vArg)
}

func (self ImGuiIO) SetKeyMods(v ImGuiModFlags) {
	C.ImGuiIO_SetKeyMods(self.handle(), C.ImGuiModFlags(v))
}

func (self ImGuiIO) SetWantCaptureMouseUnlessPopupClose(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetWantCaptureMouseUnlessPopupClose(self.handle(), vArg)
}

func (self ImGuiIO) SetMousePosPrev(v ImVec2) {
	C.ImGuiIO_SetMousePosPrev(self.handle(), v.ToC())
}

func (self ImGuiIO) SetPenPressure(v float32) {
	C.ImGuiIO_SetPenPressure(self.handle(), C.float(v))
}

func (self ImGuiIO) SetAppFocusLost(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetAppFocusLost(self.handle(), vArg)
}

func (self ImGuiIO) SetBackendUsingLegacyNavInputArray(v bool) {
	vArg := C.bool(v)

	C.ImGuiIO_SetBackendUsingLegacyNavInputArray(self.handle(), vArg)
}

func (self ImGuiInputEventMouseWheel) SetWheelX(v float32) {
	C.ImGuiInputEventMouseWheel_SetWheelX(self.handle(), C.float(v))
}

func (self ImGuiInputEventMouseWheel) SetWheelY(v float32) {
	C.ImGuiInputEventMouseWheel_SetWheelY(self.handle(), C.float(v))
}

func (self ImGuiListClipperData) SetListClipper(v ImGuiListClipper) {
	C.ImGuiListClipperData_SetListClipper(self.handle(), v.handle())
}

func (self ImGuiListClipperData) SetLossynessOffset(v float32) {
	C.ImGuiListClipperData_SetLossynessOffset(self.handle(), C.float(v))
}

func (self ImGuiListClipperData) SetStepNo(v int32) {
	C.ImGuiListClipperData_SetStepNo(self.handle(), C.int(v))
}

func (self ImGuiListClipperData) SetItemsFrozen(v int32) {
	C.ImGuiListClipperData_SetItemsFrozen(self.handle(), C.int(v))
}

func (self ImGuiTableCellData) SetBgColor(v uint32) {
	C.ImGuiTableCellData_SetBgColor(self.handle(), C.uint(v))
}

func (self ImFont) SetFallbackAdvanceX(v float32) {
	C.ImFont_SetFallbackAdvanceX(self.handle(), C.float(v))
}

func (self ImFont) SetFontSize(v float32) {
	C.ImFont_SetFontSize(self.handle(), C.float(v))
}

func (self ImFont) SetFallbackGlyph(v ImFontGlyph) {
	C.ImFont_SetFallbackGlyph(self.handle(), v.handle())
}

func (self ImFont) SetContainerAtlas(v ImFontAtlas) {
	C.ImFont_SetContainerAtlas(self.handle(), v.handle())
}

func (self ImFont) SetConfigData(v ImFontConfig) {
	C.ImFont_SetConfigData(self.handle(), v.handle())
}

func (self ImFont) SetConfigDataCount(v int) {
	vArg := C.short(v)

	C.ImFont_SetConfigDataCount(self.handle(), vArg)
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
	vArg := C.bool(v)

	C.ImFont_SetDirtyLookupTables(self.handle(), vArg)
}

func (self ImFont) SetScale(v float32) {
	C.ImFont_SetScale(self.handle(), C.float(v))
}

func (self ImFont) SetAscent(v float32) {
	C.ImFont_SetAscent(self.handle(), C.float(v))
}

func (self ImFont) SetDescent(v float32) {
	C.ImFont_SetDescent(self.handle(), C.float(v))
}

func (self ImFont) SetMetricsTotalSurface(v int32) {
	C.ImFont_SetMetricsTotalSurface(self.handle(), C.int(v))
}

func (self ImGuiTextFilter) SetCountGrep(v int32) {
	C.ImGuiTextFilter_SetCountGrep(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) SetID(v ImGuiID) {
	C.ImGuiInputTextState_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiInputTextState) SetCurLenW(v int32) {
	C.ImGuiInputTextState_SetCurLenW(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) SetCurLenA(v int32) {
	C.ImGuiInputTextState_SetCurLenA(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) SetTextAIsValid(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputTextState_SetTextAIsValid(self.handle(), vArg)
}

func (self ImGuiInputTextState) SetBufCapacityA(v int32) {
	C.ImGuiInputTextState_SetBufCapacityA(self.handle(), C.int(v))
}

func (self ImGuiInputTextState) SetScrollX(v float32) {
	C.ImGuiInputTextState_SetScrollX(self.handle(), C.float(v))
}

func (self ImGuiInputTextState) SetCursorAnim(v float32) {
	C.ImGuiInputTextState_SetCursorAnim(self.handle(), C.float(v))
}

func (self ImGuiInputTextState) SetCursorFollow(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputTextState_SetCursorFollow(self.handle(), vArg)
}

func (self ImGuiInputTextState) SetSelectedAllMouseLock(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputTextState_SetSelectedAllMouseLock(self.handle(), vArg)
}

func (self ImGuiInputTextState) SetEdited(v bool) {
	vArg := C.bool(v)

	C.ImGuiInputTextState_SetEdited(self.handle(), vArg)
}

func (self ImGuiInputTextState) SetFlags(v ImGuiInputTextFlags) {
	C.ImGuiInputTextState_SetFlags(self.handle(), C.ImGuiInputTextFlags(v))
}

func (self ImGuiKeyData) SetDown(v bool) {
	vArg := C.bool(v)

	C.ImGuiKeyData_SetDown(self.handle(), vArg)
}

func (self ImGuiKeyData) SetDownDuration(v float32) {
	C.ImGuiKeyData_SetDownDuration(self.handle(), C.float(v))
}

func (self ImGuiKeyData) SetDownDurationPrev(v float32) {
	C.ImGuiKeyData_SetDownDurationPrev(self.handle(), C.float(v))
}

func (self ImGuiKeyData) SetAnalogValue(v float32) {
	C.ImGuiKeyData_SetAnalogValue(self.handle(), C.float(v))
}

func (self ImGuiNavItemData) SetWindow(v ImGuiWindow) {
	C.ImGuiNavItemData_SetWindow(self.handle(), v.handle())
}

func (self ImGuiNavItemData) SetID(v ImGuiID) {
	C.ImGuiNavItemData_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNavItemData) SetFocusScopeId(v ImGuiID) {
	C.ImGuiNavItemData_SetFocusScopeId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiNavItemData) SetInFlags(v ImGuiItemFlags) {
	C.ImGuiNavItemData_SetInFlags(self.handle(), C.ImGuiItemFlags(v))
}

func (self ImGuiNavItemData) SetDistBox(v float32) {
	C.ImGuiNavItemData_SetDistBox(self.handle(), C.float(v))
}

func (self ImGuiNavItemData) SetDistCenter(v float32) {
	C.ImGuiNavItemData_SetDistCenter(self.handle(), C.float(v))
}

func (self ImGuiNavItemData) SetDistAxial(v float32) {
	C.ImGuiNavItemData_SetDistAxial(self.handle(), C.float(v))
}

func (self ImGuiPtrOrIndex) SetPtr(v unsafe.Pointer) {
	C.ImGuiPtrOrIndex_SetPtr(self.handle(), v)
}

func (self ImGuiPtrOrIndex) SetIndex(v int32) {
	C.ImGuiPtrOrIndex_SetIndex(self.handle(), C.int(v))
}

func (self ImGuiSizeCallbackData) SetUserData(v unsafe.Pointer) {
	C.ImGuiSizeCallbackData_SetUserData(self.handle(), v)
}

func (self ImGuiSizeCallbackData) SetPos(v ImVec2) {
	C.ImGuiSizeCallbackData_SetPos(self.handle(), v.ToC())
}

func (self ImGuiSizeCallbackData) SetCurrentSize(v ImVec2) {
	C.ImGuiSizeCallbackData_SetCurrentSize(self.handle(), v.ToC())
}

func (self ImGuiSizeCallbackData) SetDesiredSize(v ImVec2) {
	C.ImGuiSizeCallbackData_SetDesiredSize(self.handle(), v.ToC())
}

func (self ImDrawList) SetFlags(v ImDrawListFlags) {
	C.ImDrawList_SetFlags(self.handle(), C.ImDrawListFlags(v))
}

func (self ImDrawList) Set_VtxCurrentIdx(v uint32) {
	C.ImDrawList_Set_VtxCurrentIdx(self.handle(), C.uint(v))
}

func (self ImDrawList) Set_Data(v ImDrawListSharedData) {
	C.ImDrawList_Set_Data(self.handle(), v.handle())
}

func (self ImDrawList) Set_OwnerName(v string) {
	vArg, vFin := wrapString(v)
	defer vFin()

	C.ImDrawList_Set_OwnerName(self.handle(), vArg)
}

func (self ImDrawList) Set_VtxWritePtr(v ImDrawVert) {
	C.ImDrawList_Set_VtxWritePtr(self.handle(), v.handle())
}

func (self ImDrawList) Set_FringeScale(v float32) {
	C.ImDrawList_Set_FringeScale(self.handle(), C.float(v))
}

func (self ImGuiInputEventMousePos) SetPosX(v float32) {
	C.ImGuiInputEventMousePos_SetPosX(self.handle(), C.float(v))
}

func (self ImGuiInputEventMousePos) SetPosY(v float32) {
	C.ImGuiInputEventMousePos_SetPosY(self.handle(), C.float(v))
}

func (self ImGuiViewport) SetID(v ImGuiID) {
	C.ImGuiViewport_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImGuiViewport) SetFlags(v ImGuiViewportFlags) {
	C.ImGuiViewport_SetFlags(self.handle(), C.ImGuiViewportFlags(v))
}

func (self ImGuiViewport) SetPos(v ImVec2) {
	C.ImGuiViewport_SetPos(self.handle(), v.ToC())
}

func (self ImGuiViewport) SetSize(v ImVec2) {
	C.ImGuiViewport_SetSize(self.handle(), v.ToC())
}

func (self ImGuiViewport) SetWorkPos(v ImVec2) {
	C.ImGuiViewport_SetWorkPos(self.handle(), v.ToC())
}

func (self ImGuiViewport) SetWorkSize(v ImVec2) {
	C.ImGuiViewport_SetWorkSize(self.handle(), v.ToC())
}

func (self ImGuiViewport) SetDpiScale(v float32) {
	C.ImGuiViewport_SetDpiScale(self.handle(), C.float(v))
}

func (self ImGuiViewport) SetParentViewportId(v ImGuiID) {
	C.ImGuiViewport_SetParentViewportId(self.handle(), C.ImGuiID(v))
}

func (self ImGuiViewport) SetDrawData(v ImDrawData) {
	C.ImGuiViewport_SetDrawData(self.handle(), v.handle())
}

func (self ImGuiViewport) SetRendererUserData(v unsafe.Pointer) {
	C.ImGuiViewport_SetRendererUserData(self.handle(), v)
}

func (self ImGuiViewport) SetPlatformUserData(v unsafe.Pointer) {
	C.ImGuiViewport_SetPlatformUserData(self.handle(), v)
}

func (self ImGuiViewport) SetPlatformHandle(v unsafe.Pointer) {
	C.ImGuiViewport_SetPlatformHandle(self.handle(), v)
}

func (self ImGuiViewport) SetPlatformHandleRaw(v unsafe.Pointer) {
	C.ImGuiViewport_SetPlatformHandleRaw(self.handle(), v)
}

func (self ImGuiViewport) SetPlatformRequestMove(v bool) {
	vArg := C.bool(v)

	C.ImGuiViewport_SetPlatformRequestMove(self.handle(), vArg)
}

func (self ImGuiViewport) SetPlatformRequestResize(v bool) {
	vArg := C.bool(v)

	C.ImGuiViewport_SetPlatformRequestResize(self.handle(), vArg)
}

func (self ImGuiViewport) SetPlatformRequestClose(v bool) {
	vArg := C.bool(v)

	C.ImGuiViewport_SetPlatformRequestClose(self.handle(), vArg)
}

func (self ImGuiViewportP) SetIdx(v int32) {
	C.ImGuiViewportP_SetIdx(self.handle(), C.int(v))
}

func (self ImGuiViewportP) SetLastFrameActive(v int32) {
	C.ImGuiViewportP_SetLastFrameActive(self.handle(), C.int(v))
}

func (self ImGuiViewportP) SetLastFrontMostStampCount(v int32) {
	C.ImGuiViewportP_SetLastFrontMostStampCount(self.handle(), C.int(v))
}

func (self ImGuiViewportP) SetLastNameHash(v ImGuiID) {
	C.ImGuiViewportP_SetLastNameHash(self.handle(), C.ImGuiID(v))
}

func (self ImGuiViewportP) SetLastPos(v ImVec2) {
	C.ImGuiViewportP_SetLastPos(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetAlpha(v float32) {
	C.ImGuiViewportP_SetAlpha(self.handle(), C.float(v))
}

func (self ImGuiViewportP) SetLastAlpha(v float32) {
	C.ImGuiViewportP_SetLastAlpha(self.handle(), C.float(v))
}

func (self ImGuiViewportP) SetPlatformMonitor(v int) {
	vArg := C.short(v)

	C.ImGuiViewportP_SetPlatformMonitor(self.handle(), vArg)
}

func (self ImGuiViewportP) SetPlatformWindowCreated(v bool) {
	vArg := C.bool(v)

	C.ImGuiViewportP_SetPlatformWindowCreated(self.handle(), vArg)
}

func (self ImGuiViewportP) SetWindow(v ImGuiWindow) {
	C.ImGuiViewportP_SetWindow(self.handle(), v.handle())
}

func (self ImGuiViewportP) SetLastPlatformPos(v ImVec2) {
	C.ImGuiViewportP_SetLastPlatformPos(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetLastPlatformSize(v ImVec2) {
	C.ImGuiViewportP_SetLastPlatformSize(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetLastRendererSize(v ImVec2) {
	C.ImGuiViewportP_SetLastRendererSize(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetWorkOffsetMin(v ImVec2) {
	C.ImGuiViewportP_SetWorkOffsetMin(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetWorkOffsetMax(v ImVec2) {
	C.ImGuiViewportP_SetWorkOffsetMax(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetBuildWorkOffsetMin(v ImVec2) {
	C.ImGuiViewportP_SetBuildWorkOffsetMin(self.handle(), v.ToC())
}

func (self ImGuiViewportP) SetBuildWorkOffsetMax(v ImVec2) {
	C.ImGuiViewportP_SetBuildWorkOffsetMax(self.handle(), v.ToC())
}
