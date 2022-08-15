package cimgui

// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

func SliderFloat(label string, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func InputScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func Storage_GetInt(self ImGuiStorage, key ImGuiID, default_val int32) int {
	return int(C.Storage_GetInt(self.handle(), C.ImGuiID(key), C.int(default_val)))
}

func GetWindowDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.GetWindowDrawList()))
}

func DrawList_PushClipRect(self ImDrawList, clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

	C.DrawList_PushClipRect(self.handle(), clip_rect_min.ToC(), clip_rect_max.ToC(), intersect_with_current_clip_rectArg)
}

func EndTooltip() {
	C.EndTooltip()
}

func DrawList_AddDrawCmd(self ImDrawList) {
	C.DrawList_AddDrawCmd(self.handle())
}

func Checkbox(label string, v *bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapBool(v)
	defer vFin()

	return C.Checkbox(labelArg, vArg) != C.bool(true)
}

func Style_ScaleAllSizes(self ImGuiStyle, scale_factor float32) {
	C.Style_ScaleAllSizes(self.handle(), C.float(scale_factor))
}

func DrawList_PathLineTo(self ImDrawList, pos ImVec2) {
	C.DrawList_PathLineTo(self.handle(), pos.ToC())
}

func Button(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Button(labelArg, size.ToC()) != C.bool(true)
}

func EndTabItem() {
	C.EndTabItem()
}

func DrawList_PushClipRectFullScreen(self ImDrawList) {
	C.DrawList_PushClipRectFullScreen(self.handle())
}

func FontAtlas_CalcCustomRectUV(self ImFontAtlas, rect ImFontAtlasCustomRect, out_uv_min *ImVec2, out_uv_max *ImVec2) {
	out_uv_minArg, out_uv_minFin := out_uv_min.wrap()
	defer out_uv_minFin()

	out_uv_maxArg, out_uv_maxFin := out_uv_max.wrap()
	defer out_uv_maxFin()

	C.FontAtlas_CalcCustomRectUV(self.handle(), rect.handle(), out_uv_minArg, out_uv_maxArg)
}

func InputTextCallbackData_InputTextCallbackData() ImGuiInputTextCallbackData {
	return (ImGuiInputTextCallbackData)(unsafe.Pointer(C.InputTextCallbackData_InputTextCallbackData()))
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

func SetClipboardText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.SetClipboardText(textArg)
}

func ShowStackToolWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowStackToolWindow(p_openArg)
}

func Font_FindGlyph(self ImFont, c ImWchar) ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.Font_FindGlyph(self.handle(), C.ImWchar(c))))
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

func TableNextRow(row_flags ImGuiTableRowFlags, min_row_height float32) {
	C.TableNextRow(C.ImGuiTableRowFlags(row_flags), C.float(min_row_height))
}

func InputDouble(label string, v *float64, step float64, step_fast float64, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputDouble(labelArg, (*C.double)(v), C.double(step), C.double(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func LogToClipboard(auto_open_depth int32) {
	C.LogToClipboard(C.int(auto_open_depth))
}

func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	want_capture_keyboardArg := C.bool(want_capture_keyboard)

	C.SetNextFrameWantCaptureKeyboard(want_capture_keyboardArg)
}

func Storage_SetInt(self ImGuiStorage, key ImGuiID, val int32) {
	C.Storage_SetInt(self.handle(), C.ImGuiID(key), C.int(val))
}

func End() {
	C.End()
}

func GetItemRectMin(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectMin(pOutArg)
}

func Indent(indent_w float32) {
	C.Indent(C.float(indent_w))
}

func DrawList_PathFillConvex(self ImDrawList, col uint32) {
	C.DrawList_PathFillConvex(self.handle(), C.uint(col))
}

func SameLine(offset_from_start_x float32, spacing float32) {
	C.SameLine(C.float(offset_from_start_x), C.float(spacing))
}

func SetKeyboardFocusHere(offset int32) {
	C.SetKeyboardFocusHere(C.int(offset))
}

func ShowStyleEditor(ref ImGuiStyle) {
	C.ShowStyleEditor(ref.handle())
}

func InputTextCallbackData_DeleteChars(self ImGuiInputTextCallbackData, pos int32, bytes_count int32) {
	C.InputTextCallbackData_DeleteChars(self.handle(), C.int(pos), C.int(bytes_count))
}

func GetColumnWidth(column_index int32) float32 {
	return float32(C.GetColumnWidth(C.int(column_index)))
}

func IsKeyPressed(key ImGuiKey, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsKeyPressed(C.ImGuiKey(key), repeatArg) != C.bool(true)
}

func IO_AddKeyAnalogEvent(self ImGuiIO, key ImGuiKey, down bool, v float32) {
	downArg := C.bool(down)

	C.IO_AddKeyAnalogEvent(self.handle(), C.ImGuiKey(key), downArg, C.float(v))
}

func FontAtlas_SetTexID(self ImFontAtlas, id ImTextureID) {
	C.FontAtlas_SetTexID(self.handle(), C.ImTextureID(id))
}

func InputTextCallbackData_ClearSelection(self ImGuiInputTextCallbackData) {
	C.InputTextCallbackData_ClearSelection(self.handle())
}

func InputScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func StyleColorsDark(dst ImGuiStyle) {
	C.StyleColorsDark(dst.handle())
}

func DrawList_PathLineToMergeDuplicate(self ImDrawList, pos ImVec2) {
	C.DrawList_PathLineToMergeDuplicate(self.handle(), pos.ToC())
}

func BeginPopupModal(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginPopupModal(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func FontAtlas_AddFont(self ImFontAtlas, font_cfg ImFontConfig) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFont(self.handle(), font_cfg.handle())))
}

func Font_BuildLookupTable(self ImFont) {
	C.Font_BuildLookupTable(self.handle())
}

func TableSetupScrollFreeze(cols int32, rows int32) {
	C.TableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func DrawList_PopClipRect(self ImDrawList) {
	C.DrawList_PopClipRect(self.handle())
}

func EndMenu() {
	C.EndMenu()
}

func GetFrameCount() int {
	return int(C.GetFrameCount())
}

func TableHeadersRow() {
	C.TableHeadersRow()
}

func Viewport_GetWorkCenter(pOut *ImVec2, self ImGuiViewport) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Viewport_GetWorkCenter(pOutArg, self.handle())
}

func FontGlyphRangesBuilder_GetBit(self ImFontGlyphRangesBuilder, n uint64) bool {
	return C.FontGlyphRangesBuilder_GetBit(self.handle(), C.ulong(n)) != C.bool(true)
}

func DrawList_AddageQuad(self ImDrawList, user_texture_id ImTextureID, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, uv1 ImVec2, uv2 ImVec2, uv3 ImVec2, uv4 ImVec2, col uint32) {
	C.DrawList_AddageQuad(self.handle(), C.ImTextureID(user_texture_id), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), uv1.ToC(), uv2.ToC(), uv3.ToC(), uv4.ToC(), C.uint(col))
}

func DrawData_ScaleClipRects(self ImDrawData, fb_scale ImVec2) {
	C.DrawData_ScaleClipRects(self.handle(), fb_scale.ToC())
}

func ShowFontSelector(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ShowFontSelector(labelArg)
}

func TableSetColumnIndex(column_n int32) bool {
	return C.TableSetColumnIndex(C.int(column_n)) != C.bool(true)
}

func GetDrawListSharedData() ImDrawListSharedData {
	return (ImDrawListSharedData)(unsafe.Pointer(C.GetDrawListSharedData()))
}

func BeginDragDropTarget() bool {
	return C.BeginDragDropTarget() != C.bool(true)
}

func Payload_IsDataType(self ImGuiPayload, typeArg string) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.Payload_IsDataType(self.handle(), typeArgArg) != C.bool(true)
}

func BeginTable(str_id string, column int32, flags ImGuiTableFlags, outer_size ImVec2, inner_width float32) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTable(str_idArg, C.int(column), C.ImGuiTableFlags(flags), outer_size.ToC(), C.float(inner_width)) != C.bool(true)
}

func DrawList_PrimRectUV(self ImDrawList, a ImVec2, b ImVec2, uv_a ImVec2, uv_b ImVec2, col uint32) {
	C.DrawList_PrimRectUV(self.handle(), a.ToC(), b.ToC(), uv_a.ToC(), uv_b.ToC(), C.uint(col))
}

func BeginDisabled(disabled bool) {
	disabledArg := C.bool(disabled)

	C.BeginDisabled(disabledArg)
}

func FontAtlas_GetGlyphRangesThai(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesThai(self.handle()))
}

func InputInt4(label string, v *[4]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt4(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func SetNextWindowContentSize(size ImVec2) {
	C.SetNextWindowContentSize(size.ToC())
}

func TableSetColumnEnabled(column_n int32, v bool) {
	vArg := C.bool(v)

	C.TableSetColumnEnabled(C.int(column_n), vArg)
}

func Columns(count int32, id string, border bool) {
	idArg, idFin := wrapString(id)
	defer idFin()

	borderArg := C.bool(border)

	C.Columns(C.int(count), idArg, borderArg)
}

func TableGetRowIndex() int {
	return int(C.TableGetRowIndex())
}

func DrawList_AddPolyline(self ImDrawList, points *ImVec2, num_points int32, col uint32, flags ImDrawFlags, thickness float32) {
	pointsArg, pointsFin := points.wrap()
	defer pointsFin()

	C.DrawList_AddPolyline(self.handle(), pointsArg, C.int(num_points), C.uint(col), C.ImDrawFlags(flags), C.float(thickness))
}

func Payload_Clear(self ImGuiPayload) {
	C.Payload_Clear(self.handle())
}

func Font_FindGlyphNoFallback(self ImFont, c ImWchar) ImFontGlyph {
	return (ImFontGlyph)(unsafe.Pointer(C.Font_FindGlyphNoFallback(self.handle(), C.ImWchar(c))))
}

func FontGlyphRangesBuilder_AddRanges(self ImFontGlyphRangesBuilder, ranges *ImWchar) {
	C.FontGlyphRangesBuilder_AddRanges(self.handle(), (*C.ImWchar)(ranges))
}

func Storage_BuildSortByKey(self ImGuiStorage) {
	C.Storage_BuildSortByKey(self.handle())
}

func EndGroup() {
	C.EndGroup()
}

func InputInt2(label string, v *[2]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt2(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func DrawList_ChannelsMerge(self ImDrawList) {
	C.DrawList_ChannelsMerge(self.handle())
}

func AcceptDragDropPayload(typeArg string, flags ImGuiDragDropFlags) ImGuiPayload {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return (ImGuiPayload)(unsafe.Pointer(C.AcceptDragDropPayload(typeArgArg, C.ImGuiDragDropFlags(flags))))
}

func LogButtons() {
	C.LogButtons()
}

func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableHeader(labelArg)
}

func FontAtlas_GetTexDataAsRGBA32(self ImFontAtlas, out_pixels *C.uchar, out_width *int32, out_height *int32, out_bytes_per_pixel *int32) {
	out_widthArg, out_widthFin := wrapInt32(out_width)
	defer out_widthFin()

	out_heightArg, out_heightFin := wrapInt32(out_height)
	defer out_heightFin()

	out_bytes_per_pixelArg, out_bytes_per_pixelFin := wrapInt32(out_bytes_per_pixel)
	defer out_bytes_per_pixelFin()

	C.FontAtlas_GetTexDataAsRGBA32(self.handle(), &out_pixels, out_widthArg, out_heightArg, out_bytes_per_pixelArg)
}

func ShowUserGuide() {
	C.ShowUserGuide()
}

func TableGetColumnCount() int {
	return int(C.TableGetColumnCount())
}

func FontAtlas_ClearTexData(self ImFontAtlas) {
	C.FontAtlas_ClearTexData(self.handle())
}

func DrawList_PathArcToFast(self ImDrawList, center ImVec2, radius float32, a_min_of_12 int32, a_max_of_12 int32) {
	C.DrawList_PathArcToFast(self.handle(), center.ToC(), C.float(radius), C.int(a_min_of_12), C.int(a_max_of_12))
}

func DrawList_DrawList(shared_data ImDrawListSharedData) ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.DrawList_DrawList(shared_data.handle())))
}

func SetScrollHereY(center_y_ratio float32) {
	C.SetScrollHereY(C.float(center_y_ratio))
}

func IsMousePosValid(mouse_pos *ImVec2) bool {
	mouse_posArg, mouse_posFin := mouse_pos.wrap()
	defer mouse_posFin()

	return C.IsMousePosValid(mouse_posArg) != C.bool(true)
}

func DrawList_PushTextureID(self ImDrawList, texture_id ImTextureID) {
	C.DrawList_PushTextureID(self.handle(), C.ImTextureID(texture_id))
}

func PushItemWidth(item_width float32) {
	C.PushItemWidth(C.float(item_width))
}

func DragScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetDragDropPayload() ImGuiPayload {
	return (ImGuiPayload)(unsafe.Pointer(C.GetDragDropPayload()))
}

func FontAtlas_ClearFonts(self ImFontAtlas) {
	C.FontAtlas_ClearFonts(self.handle())
}

func SetItemAllowOverlap() {
	C.SetItemAllowOverlap()
}

func GetCursorPosY() float32 {
	return float32(C.GetCursorPosY())
}

func FontAtlas_AddFontFromFileTTF(self ImFontAtlas, filename string, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromFileTTF(self.handle(), filenameArg, C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func IO_IO() ImGuiIO {
	return (ImGuiIO)(unsafe.Pointer(C.IO_IO()))
}

func IsWindowFocused(flags ImGuiFocusedFlags) bool {
	return C.IsWindowFocused(C.ImGuiFocusedFlags(flags)) != C.bool(true)
}

func DrawList_ChannelsSetCurrent(self ImDrawList, n int32) {
	C.DrawList_ChannelsSetCurrent(self.handle(), C.int(n))
}

func IsItemDeactivatedAfterEdit() bool {
	return C.IsItemDeactivatedAfterEdit() != C.bool(true)
}

func DrawList_AddCircle(self ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.DrawList_AddCircle(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments), C.float(thickness))
}

func DrawList_AddNgonFilled(self ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32) {
	C.DrawList_AddNgonFilled(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments))
}

func DrawList_AddTriangleFilled(self ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32) {
	C.DrawList_AddTriangleFilled(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), C.uint(col))
}

func Font_SetGlyphVisible(self ImFont, c ImWchar, visible bool) {
	visibleArg := C.bool(visible)

	C.Font_SetGlyphVisible(self.handle(), C.ImWchar(c), visibleArg)
}

func Font_CalcWordWrapPositionA(self ImFont, scale float32, text string, text_end string, wrap_width float32) string {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.GoString(C.Font_CalcWordWrapPositionA(self.handle(), C.float(scale), textArg, text_endArg, C.float(wrap_width)))
}

func Viewport_Viewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.Viewport_Viewport()))
}

func FontAtlas_GetCustomRectByIndex(self ImFontAtlas, index int32) ImFontAtlasCustomRect {
	return (ImFontAtlasCustomRect)(unsafe.Pointer(C.FontAtlas_GetCustomRectByIndex(self.handle(), C.int(index))))
}

func SmallButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.SmallButton(labelArg) != C.bool(true)
}

func FontAtlas_Clear(self ImFontAtlas) {
	C.FontAtlas_Clear(self.handle())
}

func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	want_capture_mouseArg := C.bool(want_capture_mouse)

	C.SetNextFrameWantCaptureMouse(want_capture_mouseArg)
}

func ShowAboutWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowAboutWindow(p_openArg)
}

func DrawList_PathBezierCubicCurveTo(self ImDrawList, p2 ImVec2, p3 ImVec2, p4 ImVec2, num_segments int32) {
	C.DrawList_PathBezierCubicCurveTo(self.handle(), p2.ToC(), p3.ToC(), p4.ToC(), C.int(num_segments))
}

func InputFloat4(label string, v *[4]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat4(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func BeginMenu(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	enabledArg := C.bool(enabled)

	return C.BeginMenu(labelArg, enabledArg) != C.bool(true)
}

func PushTextWrapPos(wrap_local_pos_x float32) {
	C.PushTextWrapPos(C.float(wrap_local_pos_x))
}

func SliderInt2(label string, v *[2]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt2(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DragInt4(label string, v *[4]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt4(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SliderFloat3(label string, v *[3]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat3(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndTabBar() {
	C.EndTabBar()
}

func SetTabItemClosed(tab_or_docked_window_label string) {
	tab_or_docked_window_labelArg, tab_or_docked_window_labelFin := wrapString(tab_or_docked_window_label)
	defer tab_or_docked_window_labelFin()

	C.SetTabItemClosed(tab_or_docked_window_labelArg)
}

func Font_AddRemapChar(self ImFont, dst ImWchar, src ImWchar, overwrite_dst bool) {
	overwrite_dstArg := C.bool(overwrite_dst)

	C.Font_AddRemapChar(self.handle(), C.ImWchar(dst), C.ImWchar(src), overwrite_dstArg)
}

func ListClipper_Begin(self ImGuiListClipper, items_count int32, items_height float32) {
	C.ListClipper_Begin(self.handle(), C.int(items_count), C.float(items_height))
}

func IO_SetAppAcceptingEvents(self ImGuiIO, accepting_events bool) {
	accepting_eventsArg := C.bool(accepting_events)

	C.IO_SetAppAcceptingEvents(self.handle(), accepting_eventsArg)
}

func SetNextWindowBgAlpha(alpha float32) {
	C.SetNextWindowBgAlpha(C.float(alpha))
}

func StyleColorsClassic(dst ImGuiStyle) {
	C.StyleColorsClassic(dst.handle())
}

func Font_IsGlyphRangeUnused(self ImFont, c_begin uint32, c_last uint32) bool {
	return C.Font_IsGlyphRangeUnused(self.handle(), C.uint(c_begin), C.uint(c_last)) != C.bool(true)
}

func GetKeyPressedAmount(key ImGuiKey, repeat_delay float32, rate float32) int {
	return int(C.GetKeyPressedAmount(C.ImGuiKey(key), C.float(repeat_delay), C.float(rate)))
}

func IsAnyItemActive() bool {
	return C.IsAnyItemActive() != C.bool(true)
}

func GetItemRectSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectSize(pOutArg)
}

func NewFrame() {
	C.NewFrame()
}

func BeginTabBar(str_id string, flags ImGuiTabBarFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTabBar(str_idArg, C.ImGuiTabBarFlags(flags)) != C.bool(true)
}

func ColorButton(desc_id string, col ImVec4, flags ImGuiColorEditFlags, size ImVec2) bool {
	desc_idArg, desc_idFin := wrapString(desc_id)
	defer desc_idFin()

	return C.ColorButton(desc_idArg, col.ToC(), C.ImGuiColorEditFlags(flags), size.ToC()) != C.bool(true)
}

func InputTextCallbackData_HasSelection(self ImGuiInputTextCallbackData) bool {
	return C.InputTextCallbackData_HasSelection(self.handle()) != C.bool(true)
}

func CalcItemWidth() float32 {
	return float32(C.CalcItemWidth())
}

func GetColumnIndex() int {
	return int(C.GetColumnIndex())
}

func GetCursorPosX() float32 {
	return float32(C.GetCursorPosX())
}

func FontAtlas_AddFontDefault(self ImFontAtlas, font_cfg ImFontConfig) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontDefault(self.handle(), font_cfg.handle())))
}

func EndDisabled() {
	C.EndDisabled()
}

func IO_AddMouseWheelEvent(self ImGuiIO, wh_x float32, wh_y float32) {
	C.IO_AddMouseWheelEvent(self.handle(), C.float(wh_x), C.float(wh_y))
}

func TextBuffer_c_str(self ImGuiTextBuffer) string {
	return C.GoString(C.TextBuffer_c_str(self.handle()))
}

func DragScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalar(labelArg, C.ImGuiDataType(data_type), p_data, C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsItemEdited() bool {
	return C.IsItemEdited() != C.bool(true)
}

func GetKeyIndex(key ImGuiKey) int {
	return int(C.GetKeyIndex(C.ImGuiKey(key)))
}

func DrawList_AddTriangle(self ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddTriangle(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), C.uint(col), C.float(thickness))
}

func ColorEdit3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func ArrowButton(str_id string, dir ImGuiDir) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ArrowButton(str_idArg, C.ImGuiDir(dir)) != C.bool(true)
}

func PushClipRect(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

	C.PushClipRect(clip_rect_min.ToC(), clip_rect_max.ToC(), intersect_with_current_clip_rectArg)
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

func DrawList_AddBezierQuadratic(self ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.DrawList_AddBezierQuadratic(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), C.uint(col), C.float(thickness), C.int(num_segments))
}

func ListClipper_ForceDisplayRangeByIndices(self ImGuiListClipper, item_min int32, item_max int32) {
	C.ListClipper_ForceDisplayRangeByIndices(self.handle(), C.int(item_min), C.int(item_max))
}

func ColorPicker3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorPicker3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func TreePop() {
	C.TreePop()
}

func DrawCmd_DrawCmd() ImDrawCmd {
	return (ImDrawCmd)(unsafe.Pointer(C.DrawCmd_DrawCmd()))
}

func GetTextLineHeightWithSpacing() float32 {
	return float32(C.GetTextLineHeightWithSpacing())
}

func InvisibleButton(str_id string, size ImVec2, flags ImGuiButtonFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.InvisibleButton(str_idArg, size.ToC(), C.ImGuiButtonFlags(flags)) != C.bool(true)
}

func GetContentRegionAvail(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetContentRegionAvail(pOutArg)
}

func GetMousePosOnOpeningCurrentPopup(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMousePosOnOpeningCurrentPopup(pOutArg)
}

func OpenPopupOnItemClick(str_id string, popup_flags ImGuiPopupFlags) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.OpenPopupOnItemClick(str_idArg, C.ImGuiPopupFlags(popup_flags))
}

func DrawList_GetClipRectMin(pOut *ImVec2, self ImDrawList) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.DrawList_GetClipRectMin(pOutArg, self.handle())
}

func ColorEdit4(label string, col *[4]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit4(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func PopStyleVar(count int32) {
	C.PopStyleVar(C.int(count))
}

func SetColorEditOptions(flags ImGuiColorEditFlags) {
	C.SetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

func Font_RenderText(self ImFont, draw_list ImDrawList, size float32, pos ImVec2, col uint32, clip_rect ImVec4, text_begin string, text_end string, wrap_width float32, cpu_fine_clip bool) {
	text_beginArg, text_beginFin := wrapString(text_begin)
	defer text_beginFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	cpu_fine_clipArg := C.bool(cpu_fine_clip)

	C.Font_RenderText(self.handle(), draw_list.handle(), C.float(size), pos.ToC(), C.uint(col), clip_rect.ToC(), text_beginArg, text_endArg, C.float(wrap_width), cpu_fine_clipArg)
}

func EndDragDropTarget() {
	C.EndDragDropTarget()
}

func DrawListSplitter_Clear(self ImDrawListSplitter) {
	C.DrawListSplitter_Clear(self.handle())
}

func Storage_GetFloat(self ImGuiStorage, key ImGuiID, default_val float32) float32 {
	return float32(C.Storage_GetFloat(self.handle(), C.ImGuiID(key), C.float(default_val)))
}

func DrawList_PathBezierQuadraticCurveTo(self ImDrawList, p2 ImVec2, p3 ImVec2, num_segments int32) {
	C.DrawList_PathBezierQuadraticCurveTo(self.handle(), p2.ToC(), p3.ToC(), C.int(num_segments))
}

func FontAtlas_AddFontFromMemoryCompressedTTF(self ImFontAtlas, compressed_font_data unsafe.Pointer, compressed_font_size int32, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromMemoryCompressedTTF(self.handle(), compressed_font_data, C.int(compressed_font_size), C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func Font_ClearOutputData(self ImFont) {
	C.Font_ClearOutputData(self.handle())
}

func BeginListBox(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.BeginListBox(labelArg, size.ToC()) != C.bool(true)
}

func DrawList_PrimWriteIdx(self ImDrawList, idx ImDrawIdx) {
	C.DrawList_PrimWriteIdx(self.handle(), C.ImDrawIdx(idx))
}

func GetClipboardText() string {
	return C.GoString(C.GetClipboardText())
}

func FontAtlas_AddFontFromMemoryCompressedBase85TTF(self ImFontAtlas, compressed_font_data_base85 string, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	compressed_font_data_base85Arg, compressed_font_data_base85Fin := wrapString(compressed_font_data_base85)
	defer compressed_font_data_base85Fin()

	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromMemoryCompressedBase85TTF(self.handle(), compressed_font_data_base85Arg, C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func NewLine() {
	C.NewLine()
}

func DestroyContext(ctx ImGuiContext) {
	C.DestroyContext(ctx.handle())
}

func IsKeyDown(key ImGuiKey) bool {
	return C.IsKeyDown(C.ImGuiKey(key)) != C.bool(true)
}

func TextUnformatted(text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.TextUnformatted(textArg, text_endArg)
}

func PopStyleColor(count int32) {
	C.PopStyleColor(C.int(count))
}

func PushAllowKeyboardFocus(allow_keyboard_focus bool) {
	allow_keyboard_focusArg := C.bool(allow_keyboard_focus)

	C.PushAllowKeyboardFocus(allow_keyboard_focusArg)
}

func DrawList_AddCircleFilled(self ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32) {
	C.DrawList_AddCircleFilled(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments))
}

func GetScrollMaxX() float32 {
	return float32(C.GetScrollMaxX())
}

func GetMouseDragDelta(pOut *ImVec2, button ImGuiMouseButton, lock_threshold float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMouseDragDelta(pOutArg, C.ImGuiMouseButton(button), C.float(lock_threshold))
}

func DrawList_CloneOutput(self ImDrawList) ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.DrawList_CloneOutput(self.handle())))
}

func TextFilter_Draw(self ImGuiTextFilter, label string, width float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TextFilter_Draw(self.handle(), labelArg, C.float(width)) != C.bool(true)
}

func GetScrollMaxY() float32 {
	return float32(C.GetScrollMaxY())
}

func PopFont() {
	C.PopFont()
}

func GetKeyName(key ImGuiKey) string {
	return C.GoString(C.GetKeyName(C.ImGuiKey(key)))
}

func EndCombo() {
	C.EndCombo()
}

func Storage_SetBool(self ImGuiStorage, key ImGuiID, val bool) {
	valArg := C.bool(val)

	C.Storage_SetBool(self.handle(), C.ImGuiID(key), valArg)
}

func DrawList_AddRectFilledMultiColor(self ImDrawList, p_min ImVec2, p_max ImVec2, col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	C.DrawList_AddRectFilledMultiColor(self.handle(), p_min.ToC(), p_max.ToC(), C.uint(col_upr_left), C.uint(col_upr_right), C.uint(col_bot_right), C.uint(col_bot_left))
}

func ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ShowStyleSelector(labelArg) != C.bool(true)
}

func DrawList_AddBezierCubic(self ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.DrawList_AddBezierCubic(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), C.uint(col), C.float(thickness), C.int(num_segments))
}

func Spacing() {
	C.Spacing()
}

func DrawList_PrimQuadUV(self ImDrawList, a ImVec2, b ImVec2, c ImVec2, d ImVec2, uv_a ImVec2, uv_b ImVec2, uv_c ImVec2, uv_d ImVec2, col uint32) {
	C.DrawList_PrimQuadUV(self.handle(), a.ToC(), b.ToC(), c.ToC(), d.ToC(), uv_a.ToC(), uv_b.ToC(), uv_c.ToC(), uv_d.ToC(), C.uint(col))
}

func TextFilter_IsActive(self ImGuiTextFilter) bool {
	return C.TextFilter_IsActive(self.handle()) != C.bool(true)
}

func IsItemDeactivated() bool {
	return C.IsItemDeactivated() != C.bool(true)
}

func Font_GetDebugName(self ImFont) string {
	return C.GoString(C.Font_GetDebugName(self.handle()))
}

func InputFloat2(label string, v *[2]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat2(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func BeginChildFrame(id ImGuiID, size ImVec2, flags ImGuiWindowFlags) bool {
	return C.BeginChildFrame(C.ImGuiID(id), size.ToC(), C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func DrawList_AddageRounded(self ImDrawList, user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.DrawList_AddageRounded(self.handle(), C.ImTextureID(user_texture_id), p_min.ToC(), p_max.ToC(), uv_min.ToC(), uv_max.ToC(), C.uint(col), C.float(rounding), C.ImDrawFlags(flags))
}

func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32) bool {
	return C.IsMouseDragging(C.ImGuiMouseButton(button), C.float(lock_threshold)) != C.bool(true)
}

func TextFilter_Clear(self ImGuiTextFilter) {
	C.TextFilter_Clear(self.handle())
}

func Font_IsLoaded(self ImFont) bool {
	return C.Font_IsLoaded(self.handle()) != C.bool(true)
}

func Separator() {
	C.Separator()
}

func BeginTabItem(label string, p_open *bool, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginTabItem(labelArg, p_openArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func NextColumn() {
	C.NextColumn()
}

func DrawList_AddRectFilled(self ImDrawList, p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.DrawList_AddRectFilled(self.handle(), p_min.ToC(), p_max.ToC(), C.uint(col), C.float(rounding), C.ImDrawFlags(flags))
}

func EndDragDropSource() {
	C.EndDragDropSource()
}

func EndFrame() {
	C.EndFrame()
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

func FontAtlasCustomRect_IsPacked(self ImFontAtlasCustomRect) bool {
	return C.FontAtlasCustomRect_IsPacked(self.handle()) != C.bool(true)
}

func IsAnyItemFocused() bool {
	return C.IsAnyItemFocused() != C.bool(true)
}

func SliderFloat4(label string, v *[4]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat4(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DrawListSplitter_DrawListSplitter() ImDrawListSplitter {
	return (ImDrawListSplitter)(unsafe.Pointer(C.DrawListSplitter_DrawListSplitter()))
}

func FontAtlas_GetGlyphRangesDefault(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesDefault(self.handle()))
}

func SliderScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func TableColumnSortSpecs_TableColumnSortSpecs() ImGuiTableColumnSortSpecs {
	return (ImGuiTableColumnSortSpecs)(unsafe.Pointer(C.TableColumnSortSpecs_TableColumnSortSpecs()))
}

func DrawList_ChannelsSplit(self ImDrawList, count int32) {
	C.DrawList_ChannelsSplit(self.handle(), C.int(count))
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

func IO_ClearInputKeys(self ImGuiIO) {
	C.IO_ClearInputKeys(self.handle())
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

func LoadIniSettingsFromMemory(ini_data string, ini_size uint64) {
	ini_dataArg, ini_dataFin := wrapString(ini_data)
	defer ini_dataFin()

	C.LoadIniSettingsFromMemory(ini_dataArg, C.ulong(ini_size))
}

func GetStateStorage() ImGuiStorage {
	return (ImGuiStorage)(unsafe.Pointer(C.GetStateStorage()))
}

func IsWindowHovered(flags ImGuiHoveredFlags) bool {
	return C.IsWindowHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func FontGlyphRangesBuilder_AddText(self ImFontGlyphRangesBuilder, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.FontGlyphRangesBuilder_AddText(self.handle(), textArg, text_endArg)
}

func Storage_SetFloat(self ImGuiStorage, key ImGuiID, val float32) {
	C.Storage_SetFloat(self.handle(), C.ImGuiID(key), C.float(val))
}

func BeginMainMenuBar() bool {
	return C.BeginMainMenuBar() != C.bool(true)
}

func GetFont() ImFont {
	return (ImFont)(unsafe.Pointer(C.GetFont()))
}

func FontAtlas_AddCustomRectRegular(self ImFontAtlas, width int32, height int32) int {
	return int(C.FontAtlas_AddCustomRectRegular(self.handle(), C.int(width), C.int(height)))
}

func DragInt3(label string, v *[3]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt3(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func MemFree(ptr unsafe.Pointer) {
	C.MemFree(ptr)
}

func Storage_SetVoidPtr(self ImGuiStorage, key ImGuiID, val unsafe.Pointer) {
	C.Storage_SetVoidPtr(self.handle(), C.ImGuiID(key), val)
}

func SetScrollHereX(center_x_ratio float32) {
	C.SetScrollHereX(C.float(center_x_ratio))
}

func Storage_GetBool(self ImGuiStorage, key ImGuiID, default_val bool) bool {
	default_valArg := C.bool(default_val)

	return C.Storage_GetBool(self.handle(), C.ImGuiID(key), default_valArg) != C.bool(true)
}

func DrawListSplitter_ClearFreeMemory(self ImDrawListSplitter) {
	C.DrawListSplitter_ClearFreeMemory(self.handle())
}

func GetCursorScreenPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetCursorScreenPos(pOutArg)
}

func GetWindowContentRegionMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowContentRegionMax(pOutArg)
}

func TableGetColumnFlags(column_n int32) ImGuiTableColumnFlags {
	return ImGuiTableColumnFlags(C.TableGetColumnFlags(C.int(column_n)))
}

func DrawList_PrimReserve(self ImDrawList, idx_count int32, vtx_count int32) {
	C.DrawList_PrimReserve(self.handle(), C.int(idx_count), C.int(vtx_count))
}

func GetIO() ImGuiIO {
	return (ImGuiIO)(unsafe.Pointer(C.GetIO()))
}

func IsItemFocused() bool {
	return C.IsItemFocused() != C.bool(true)
}

func LogToFile(auto_open_depth int32, filename string) {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	C.LogToFile(C.int(auto_open_depth), filenameArg)
}

func IO_SetKeyEventNativeData(self ImGuiIO, key ImGuiKey, native_keycode int32, native_scancode int32, native_legacy_index int32) {
	C.IO_SetKeyEventNativeData(self.handle(), C.ImGuiKey(key), C.int(native_keycode), C.int(native_scancode), C.int(native_legacy_index))
}

func InputInt(label string, v *int32, step int32, step_fast int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	return C.InputInt(labelArg, vArg, C.int(step), C.int(step_fast), C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func FontAtlas_FontAtlas() ImFontAtlas {
	return (ImFontAtlas)(unsafe.Pointer(C.FontAtlas_FontAtlas()))
}

func SliderInt3(label string, v *[3]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt3(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DrawData_DrawData() ImDrawData {
	return (ImDrawData)(unsafe.Pointer(C.DrawData_DrawData()))
}

func GetColumnOffset(column_index int32) float32 {
	return float32(C.GetColumnOffset(C.int(column_index)))
}

func BeginPopup(str_id string, flags ImGuiWindowFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopup(str_idArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func GetContentRegionMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetContentRegionMax(pOutArg)
}

func GetScrollX() float32 {
	return float32(C.GetScrollX())
}

func GetFontSize() float32 {
	return float32(C.GetFontSize())
}

func VSliderScalar(label string, size ImVec2, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderScalar(labelArg, size.ToC(), C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Bullet() {
	C.Bullet()
}

func GetWindowPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowPos(pOutArg)
}

func GetWindowWidth() float32 {
	return float32(C.GetWindowWidth())
}

func IO_AddInputCharacter(self ImGuiIO, c uint32) {
	C.IO_AddInputCharacter(self.handle(), C.uint(c))
}

func FontAtlas_GetTexDataAsAlpha8(self ImFontAtlas, out_pixels *C.uchar, out_width *int32, out_height *int32, out_bytes_per_pixel *int32) {
	out_widthArg, out_widthFin := wrapInt32(out_width)
	defer out_widthFin()

	out_heightArg, out_heightFin := wrapInt32(out_height)
	defer out_heightFin()

	out_bytes_per_pixelArg, out_bytes_per_pixelFin := wrapInt32(out_bytes_per_pixel)
	defer out_bytes_per_pixelFin()

	C.FontAtlas_GetTexDataAsAlpha8(self.handle(), &out_pixels, out_widthArg, out_heightArg, out_bytes_per_pixelArg)
}

func IsWindowCollapsed() bool {
	return C.IsWindowCollapsed() != C.bool(true)
}

func LogFinish() {
	C.LogFinish()
}

func OnceUponAFrame_OnceUponAFrame() ImGuiOnceUponAFrame {
	return (ImGuiOnceUponAFrame)(unsafe.Pointer(C.OnceUponAFrame_OnceUponAFrame()))
}

func FontAtlas_ClearInputData(self ImFontAtlas) {
	C.FontAtlas_ClearInputData(self.handle())
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

func DrawListSplitter_Merge(self ImDrawListSplitter, draw_list ImDrawList) {
	C.DrawListSplitter_Merge(self.handle(), draw_list.handle())
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

func SetNextWindowFocus() {
	C.SetNextWindowFocus()
}

func ListClipper_Step(self ImGuiListClipper) bool {
	return C.ListClipper_Step(self.handle()) != C.bool(true)
}

func GetColumnsCount() int {
	return int(C.GetColumnsCount())
}

func GetDrawData() ImDrawData {
	return (ImDrawData)(unsafe.Pointer(C.GetDrawData()))
}

func Begin(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.Begin(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func GetMousePos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetMousePos(pOutArg)
}

func IsItemActivated() bool {
	return C.IsItemActivated() != C.bool(true)
}

func IsMouseDown(button ImGuiMouseButton) bool {
	return C.IsMouseDown(C.ImGuiMouseButton(button)) != C.bool(true)
}

func TabItemButton(label string, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TabItemButton(labelArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func DrawList_AddQuadFilled(self ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32) {
	C.DrawList_AddQuadFilled(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), C.uint(col))
}

func FontGlyphRangesBuilder_Clear(self ImFontGlyphRangesBuilder) {
	C.FontGlyphRangesBuilder_Clear(self.handle())
}

func FontGlyphRangesBuilder_FontGlyphRangesBuilder() ImFontGlyphRangesBuilder {
	return (ImFontGlyphRangesBuilder)(unsafe.Pointer(C.FontGlyphRangesBuilder_FontGlyphRangesBuilder()))
}

func GetItemRectMax(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetItemRectMax(pOutArg)
}

func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return C.IsMouseDoubleClicked(C.ImGuiMouseButton(button)) != C.bool(true)
}

func SetNextItemWidth(item_width float32) {
	C.SetNextItemWidth(C.float(item_width))
}

func BeginCombo(label string, preview_value string, flags ImGuiComboFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	preview_valueArg, preview_valueFin := wrapString(preview_value)
	defer preview_valueFin()

	return C.BeginCombo(labelArg, preview_valueArg, C.ImGuiComboFlags(flags)) != C.bool(true)
}

func DrawList_PathClear(self ImDrawList) {
	C.DrawList_PathClear(self.handle())
}

func SetCursorPosX(local_x float32) {
	C.SetCursorPosX(C.float(local_x))
}

func SetStateStorage(storage ImGuiStorage) {
	C.SetStateStorage(storage.handle())
}

func DrawList_AddRect(self ImDrawList, p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags, thickness float32) {
	C.DrawList_AddRect(self.handle(), p_min.ToC(), p_max.ToC(), C.uint(col), C.float(rounding), C.ImDrawFlags(flags), C.float(thickness))
}

func FontAtlas_GetGlyphRangesCyrillic(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesCyrillic(self.handle()))
}

func ShowMetricsWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowMetricsWindow(p_openArg)
}

func DrawData_Clear(self ImDrawData) {
	C.DrawData_Clear(self.handle())
}

func PopID() {
	C.PopID()
}

func Payload_IsPreview(self ImGuiPayload) bool {
	return C.Payload_IsPreview(self.handle()) != C.bool(true)
}

func GetScrollY() float32 {
	return float32(C.GetScrollY())
}

func GetTreeNodeToLabelSpacing() float32 {
	return float32(C.GetTreeNodeToLabelSpacing())
}

func PopTextWrapPos() {
	C.PopTextWrapPos()
}

func FontAtlas_Build(self ImFontAtlas) bool {
	return C.FontAtlas_Build(self.handle()) != C.bool(true)
}

func FontAtlas_GetGlyphRangesJapanese(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesJapanese(self.handle()))
}

func DrawList_AddQuad(self ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddQuad(self.handle(), p1.ToC(), p2.ToC(), p3.ToC(), p4.ToC(), C.uint(col), C.float(thickness))
}

func AlignTextToFramePadding() {
	C.AlignTextToFramePadding()
}

func IsMouseClicked(button ImGuiMouseButton, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsMouseClicked(C.ImGuiMouseButton(button), repeatArg) != C.bool(true)
}

func SaveIniSettingsToMemory(out_ini_size *uint64) string {
	return C.GoString(C.SaveIniSettingsToMemory((*C.ulong)(out_ini_size)))
}

func FontAtlas_AddFontFromMemoryTTF(self ImFontAtlas, font_data unsafe.Pointer, font_size int32, size_pixels float32, font_cfg ImFontConfig, glyph_ranges *ImWchar) ImFont {
	return (ImFont)(unsafe.Pointer(C.FontAtlas_AddFontFromMemoryTTF(self.handle(), font_data, C.int(font_size), C.float(size_pixels), font_cfg.handle(), (*C.ImWchar)(glyph_ranges))))
}

func SetMouseCursor(cursor_type ImGuiMouseCursor) {
	C.SetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

func FontAtlas_GetGlyphRangesVietnamese(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesVietnamese(self.handle()))
}

func DragFloat4(label string, v *[4]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat4(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DrawListSplitter_Split(self ImDrawListSplitter, draw_list ImDrawList, count int32) {
	C.DrawListSplitter_Split(self.handle(), draw_list.handle(), C.int(count))
}

func PopButtonRepeat() {
	C.PopButtonRepeat()
}

func SetDragDropPayload(typeArg string, data unsafe.Pointer, sz uint64, cond ImGuiCond) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.SetDragDropPayload(typeArgArg, data, C.ulong(sz), C.ImGuiCond(cond)) != C.bool(true)
}

func Color_SetHSV(self *ImColor, h float32, s float32, v float32, a float32) {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.Color_SetHSV(selfArg, C.float(h), C.float(s), C.float(v), C.float(a))
}

func DrawList_GetClipRectMax(pOut *ImVec2, self ImDrawList) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.DrawList_GetClipRectMax(pOutArg, self.handle())
}

func CloseCurrentPopup() {
	C.CloseCurrentPopup()
}

func SliderFloat2(label string, v *[2]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat2(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func BeginMenuBar() bool {
	return C.BeginMenuBar() != C.bool(true)
}

func Payload_Payload() ImGuiPayload {
	return (ImGuiPayload)(unsafe.Pointer(C.Payload_Payload()))
}

func ResetMouseDragDelta(button ImGuiMouseButton) {
	C.ResetMouseDragDelta(C.ImGuiMouseButton(button))
}

func SetNextWindowPos(pos ImVec2, cond ImGuiCond, pivot ImVec2) {
	C.SetNextWindowPos(pos.ToC(), C.ImGuiCond(cond), pivot.ToC())
}

func EndTable() {
	C.EndTable()
}

func FontAtlasCustomRect_FontAtlasCustomRect() ImFontAtlasCustomRect {
	return (ImFontAtlasCustomRect)(unsafe.Pointer(C.FontAtlasCustomRect_FontAtlasCustomRect()))
}

func GetVersion() string {
	return C.GoString(C.GetVersion())
}

func DrawList_Addage(self ImDrawList, user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32) {
	C.DrawList_Addage(self.handle(), C.ImTextureID(user_texture_id), p_min.ToC(), p_max.ToC(), uv_min.ToC(), uv_max.ToC(), C.uint(col))
}

func EndListBox() {
	C.EndListBox()
}

func InputFloat3(label string, v *[3]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat3(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func EndChild() {
	C.EndChild()
}

func IsWindowAppearing() bool {
	return C.IsWindowAppearing() != C.bool(true)
}

func Font_GetCharAdvance(self ImFont, c ImWchar) float32 {
	return float32(C.Font_GetCharAdvance(self.handle(), C.ImWchar(c)))
}

func GetFrameHeightWithSpacing() float32 {
	return float32(C.GetFrameHeightWithSpacing())
}

func TableSortSpecs_TableSortSpecs() ImGuiTableSortSpecs {
	return (ImGuiTableSortSpecs)(unsafe.Pointer(C.TableSortSpecs_TableSortSpecs()))
}

func GetMouseClickedCount(button ImGuiMouseButton) int {
	return int(C.GetMouseClickedCount(C.ImGuiMouseButton(button)))
}

func GetWindowHeight() float32 {
	return float32(C.GetWindowHeight())
}

func DrawList_PrimWriteVtx(self ImDrawList, pos ImVec2, uv ImVec2, col uint32) {
	C.DrawList_PrimWriteVtx(self.handle(), pos.ToC(), uv.ToC(), C.uint(col))
}

func BeginTooltip() {
	C.BeginTooltip()
}

func EndMenuBar() {
	C.EndMenuBar()
}

func EndMainMenuBar() {
	C.EndMainMenuBar()
}

func SetCursorPosY(local_y float32) {
	C.SetCursorPosY(C.float(local_y))
}

func SetNextItemOpen(is_open bool, cond ImGuiCond) {
	is_openArg := C.bool(is_open)

	C.SetNextItemOpen(is_openArg, C.ImGuiCond(cond))
}

func ColorPicker4(label string, col *[4]float32, flags ImGuiColorEditFlags, ref_col *float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	ref_colArg, ref_colFin := wrapFloat(ref_col)
	defer ref_colFin()

	return C.ColorPicker4(labelArg, colArg, C.ImGuiColorEditFlags(flags), ref_colArg) != C.bool(true)
}

func DrawList_PathArcTo(self ImDrawList, center ImVec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	C.DrawList_PathArcTo(self.handle(), center.ToC(), C.float(radius), C.float(a_min), C.float(a_max), C.int(num_segments))
}

func GetCurrentContext() ImGuiContext {
	return (ImGuiContext)(unsafe.Pointer(C.GetCurrentContext()))
}

func LoadIniSettingsFromDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.LoadIniSettingsFromDisk(ini_filenameArg)
}

func PopAllowKeyboardFocus() {
	C.PopAllowKeyboardFocus()
}

func DrawList_PrimVtx(self ImDrawList, pos ImVec2, uv ImVec2, col uint32) {
	C.DrawList_PrimVtx(self.handle(), pos.ToC(), uv.ToC(), C.uint(col))
}

func GetStyle() ImGuiStyle {
	return (ImGuiStyle)(unsafe.Pointer(C.GetStyle()))
}

func Storage_SetAllInt(self ImGuiStorage, val int32) {
	C.Storage_SetAllInt(self.handle(), C.int(val))
}

func IO_AddFocusEvent(self ImGuiIO, focused bool) {
	focusedArg := C.bool(focused)

	C.IO_AddFocusEvent(self.handle(), focusedArg)
}

func FontAtlas_IsBuilt(self ImFontAtlas) bool {
	return C.FontAtlas_IsBuilt(self.handle()) != C.bool(true)
}

func ColorConvertU32ToFloat4(pOut *ImVec4, in uint32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ColorConvertU32ToFloat4(pOutArg, C.uint(in))
}

func DrawList_AddLine(self ImDrawList, p1 ImVec2, p2 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddLine(self.handle(), p1.ToC(), p2.ToC(), C.uint(col), C.float(thickness))
}

func GetTextLineHeight() float32 {
	return float32(C.GetTextLineHeight())
}

func GetWindowSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowSize(pOutArg)
}

func IsItemClicked(mouse_button ImGuiMouseButton) bool {
	return C.IsItemClicked(C.ImGuiMouseButton(mouse_button)) != C.bool(true)
}

func SliderScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func TableGetColumnIndex() int {
	return int(C.TableGetColumnIndex())
}

func DrawList_PrimRect(self ImDrawList, a ImVec2, b ImVec2, col uint32) {
	C.DrawList_PrimRect(self.handle(), a.ToC(), b.ToC(), C.uint(col))
}

func IO_AddMouseButtonEvent(self ImGuiIO, button int32, down bool) {
	downArg := C.bool(down)

	C.IO_AddMouseButtonEvent(self.handle(), C.int(button), downArg)
}

func PushFont(font ImFont) {
	C.PushFont(font.handle())
}

func SetCursorPos(local_pos ImVec2) {
	C.SetCursorPos(local_pos.ToC())
}

func ShowDebugLogWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDebugLogWindow(p_openArg)
}

func DrawData_DeIndexAllBuffers(self ImDrawData) {
	C.DrawData_DeIndexAllBuffers(self.handle())
}

func GetFrameHeight() float32 {
	return float32(C.GetFrameHeight())
}

func DrawList_PrimUnreserve(self ImDrawList, idx_count int32, vtx_count int32) {
	C.DrawList_PrimUnreserve(self.handle(), C.int(idx_count), C.int(vtx_count))
}

func TableSetupColumn(label string, flags ImGuiTableColumnFlags, init_width_or_weight float32, user_id ImGuiID) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableSetupColumn(labelArg, C.ImGuiTableColumnFlags(flags), C.float(init_width_or_weight), C.ImGuiID(user_id))
}

func DrawList_PathRect(self ImDrawList, rect_min ImVec2, rect_max ImVec2, rounding float32, flags ImDrawFlags) {
	C.DrawList_PathRect(self.handle(), rect_min.ToC(), rect_max.ToC(), C.float(rounding), C.ImDrawFlags(flags))
}

func SetCurrentContext(ctx ImGuiContext) {
	C.SetCurrentContext(ctx.handle())
}

func ListClipper_End(self ImGuiListClipper) {
	C.ListClipper_End(self.handle())
}

func IO_AddInputCharactersUTF8(self ImGuiIO, str string) {
	strArg, strFin := wrapString(str)
	defer strFin()

	C.IO_AddInputCharactersUTF8(self.handle(), strArg)
}

func IsMouseHoveringRect(r_min ImVec2, r_max ImVec2, clip bool) bool {
	clipArg := C.bool(clip)

	return C.IsMouseHoveringRect(r_min.ToC(), r_max.ToC(), clipArg) != C.bool(true)
}

func FontGlyphRangesBuilder_AddChar(self ImFontGlyphRangesBuilder, c ImWchar) {
	C.FontGlyphRangesBuilder_AddChar(self.handle(), C.ImWchar(c))
}

func GetFontTexUvWhitePixel(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetFontTexUvWhitePixel(pOutArg)
}

func IsItemVisible() bool {
	return C.IsItemVisible() != C.bool(true)
}

func SetWindowFontScale(scale float32) {
	C.SetWindowFontScale(C.float(scale))
}

func SetColumnOffset(column_index int32, offset_x float32) {
	C.SetColumnOffset(C.int(column_index), C.float(offset_x))
}

func TextFilter_Build(self ImGuiTextFilter) {
	C.TextFilter_Build(self.handle())
}

func ShowDemoWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDemoWindow(p_openArg)
}

func IsMouseReleased(button ImGuiMouseButton) bool {
	return C.IsMouseReleased(C.ImGuiMouseButton(button)) != C.bool(true)
}

func DrawList_PathStroke(self ImDrawList, col uint32, flags ImDrawFlags, thickness float32) {
	C.DrawList_PathStroke(self.handle(), C.uint(col), C.ImDrawFlags(flags), C.float(thickness))
}

func Dummy(size ImVec2) {
	C.Dummy(size.ToC())
}

func DebugTextEncoding(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.DebugTextEncoding(textArg)
}

func SaveIniSettingsToDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.SaveIniSettingsToDisk(ini_filenameArg)
}

func Font_AddGlyph(self ImFont, src_cfg ImFontConfig, c ImWchar, x0 float32, y0 float32, x1 float32, y1 float32, u0 float32, v0 float32, u1 float32, v1 float32, advance_x float32) {
	C.Font_AddGlyph(self.handle(), src_cfg.handle(), C.ImWchar(c), C.float(x0), C.float(y0), C.float(x1), C.float(y1), C.float(u0), C.float(v0), C.float(u1), C.float(v1), C.float(advance_x))
}

func BeginGroup() {
	C.BeginGroup()
}

func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond) {
	collapsedArg := C.bool(collapsed)

	C.SetNextWindowCollapsed(collapsedArg, C.ImGuiCond(cond))
}

func LogToTTY(auto_open_depth int32) {
	C.LogToTTY(C.int(auto_open_depth))
}

func IO_ClearInputCharacters(self ImGuiIO) {
	C.IO_ClearInputCharacters(self.handle())
}

func IO_AddKeyEvent(self ImGuiIO, key ImGuiKey, down bool) {
	downArg := C.bool(down)

	C.IO_AddKeyEvent(self.handle(), C.ImGuiKey(key), downArg)
}

func TextFilter_TextFilter(default_filter string) ImGuiTextFilter {
	default_filterArg, default_filterFin := wrapString(default_filter)
	defer default_filterFin()

	return (ImGuiTextFilter)(unsafe.Pointer(C.TextFilter_TextFilter(default_filterArg)))
}

func PopClipRect() {
	C.PopClipRect()
}

func SetCursorScreenPos(pos ImVec2) {
	C.SetCursorScreenPos(pos.ToC())
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

func EndPopup() {
	C.EndPopup()
}

func EndChildFrame() {
	C.EndChildFrame()
}

func Font_Font() ImFont {
	return (ImFont)(unsafe.Pointer(C.Font_Font()))
}

func DragFloat2(label string, v *[2]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat2(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetMouseCursor() ImGuiMouseCursor {
	return ImGuiMouseCursor(C.GetMouseCursor())
}

func FontAtlas_AddCustomRectFontGlyph(self ImFontAtlas, font ImFont, id ImWchar, width int32, height int32, advance_x float32, offset ImVec2) int {
	return int(C.FontAtlas_AddCustomRectFontGlyph(self.handle(), font.handle(), C.ImWchar(id), C.int(width), C.int(height), C.float(advance_x), offset.ToC()))
}

func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
	return C.BeginDragDropSource(C.ImGuiDragDropFlags(flags)) != C.bool(true)
}

func BeginPopupContextItem(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextItem(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func GetWindowContentRegionMin(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetWindowContentRegionMin(pOutArg)
}

func InputInt3(label string, v *[3]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt3(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func DrawList_PopTextureID(self ImDrawList) {
	C.DrawList_PopTextureID(self.handle())
}

func TableGetSortSpecs() ImGuiTableSortSpecs {
	return (ImGuiTableSortSpecs)(unsafe.Pointer(C.TableGetSortSpecs()))
}

func Viewport_GetCenter(pOut *ImVec2, self ImGuiViewport) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Viewport_GetCenter(pOutArg, self.handle())
}

func DragInt2(label string, v *[2]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt2(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IO_AddMousePosEvent(self ImGuiIO, x float32, y float32) {
	C.IO_AddMousePosEvent(self.handle(), C.float(x), C.float(y))
}

func Payload_IsDelivery(self ImGuiPayload) bool {
	return C.Payload_IsDelivery(self.handle()) != C.bool(true)
}

func Font_RenderChar(self ImFont, draw_list ImDrawList, size float32, pos ImVec2, col uint32, c ImWchar) {
	C.Font_RenderChar(self.handle(), draw_list.handle(), C.float(size), pos.ToC(), C.uint(col), C.ImWchar(c))
}

func GetCursorPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetCursorPos(pOutArg)
}

func Font_GrowIndex(self ImFont, new_size int32) {
	C.Font_GrowIndex(self.handle(), C.int(new_size))
}

func BeginPopupContextWindow(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextWindow(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
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

func DragFloat3(label string, v *[3]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat3(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func FontAtlas_GetGlyphRangesKorean(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesKorean(self.handle()))
}

func TextFilter_PassFilter(self ImGuiTextFilter, text string, text_end string) bool {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.TextFilter_PassFilter(self.handle(), textArg, text_endArg) != C.bool(true)
}

func BeginPopupContextVoid(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextVoid(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func DrawList_AddConvexPolyFilled(self ImDrawList, points *ImVec2, num_points int32, col uint32) {
	pointsArg, pointsFin := points.wrap()
	defer pointsFin()

	C.DrawList_AddConvexPolyFilled(self.handle(), pointsArg, C.int(num_points), C.uint(col))
}

func InputTextCallbackData_InsertChars(self ImGuiInputTextCallbackData, pos int32, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.InputTextCallbackData_InsertChars(self.handle(), C.int(pos), textArg, text_endArg)
}

func DrawList_AddNgon(self ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.DrawList_AddNgon(self.handle(), center.ToC(), C.float(radius), C.uint(col), C.int(num_segments), C.float(thickness))
}

func DebugCheckVersionAndDataLayout(version_str string, sz_io uint64, sz_style uint64, sz_vec2 uint64, sz_vec4 uint64, sz_drawvert uint64, sz_drawidx uint64) bool {
	version_strArg, version_strFin := wrapString(version_str)
	defer version_strFin()

	return C.DebugCheckVersionAndDataLayout(version_strArg, C.ulong(sz_io), C.ulong(sz_style), C.ulong(sz_vec2), C.ulong(sz_vec4), C.ulong(sz_drawvert), C.ulong(sz_drawidx)) != C.bool(true)
}

func TextBuffer_TextBuffer() ImGuiTextBuffer {
	return (ImGuiTextBuffer)(unsafe.Pointer(C.TextBuffer_TextBuffer()))
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

func IsItemHovered(flags ImGuiHoveredFlags) bool {
	return C.IsItemHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
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

func InputTextCallbackData_SelectAll(self ImGuiInputTextCallbackData) {
	C.InputTextCallbackData_SelectAll(self.handle())
}

func StyleColorsLight(dst ImGuiStyle) {
	C.StyleColorsLight(dst.handle())
}

func SetColumnWidth(column_index int32, width float32) {
	C.SetColumnWidth(C.int(column_index), C.float(width))
}

func IsKeyReleased(key ImGuiKey) bool {
	return C.IsKeyReleased(C.ImGuiKey(key)) != C.bool(true)
}

func Render() {
	C.Render()
}

func IsItemActive() bool {
	return C.IsItemActive() != C.bool(true)
}

func SliderInt4(label string, v *[4]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt4(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func FontAtlas_GetGlyphRangesChineseFull(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesChineseFull(self.handle()))
}

func Storage_Clear(self ImGuiStorage) {
	C.Storage_Clear(self.handle())
}

func Style_Style() ImGuiStyle {
	return (ImGuiStyle)(unsafe.Pointer(C.Style_Style()))
}

func GetCursorStartPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.GetCursorStartPos(pOutArg)
}

func GetMainViewport() ImGuiViewport {
	return (ImGuiViewport)(unsafe.Pointer(C.GetMainViewport()))
}

func DrawListSplitter_SetCurrentChannel(self ImDrawListSplitter, draw_list ImDrawList, channel_idx int32) {
	C.DrawListSplitter_SetCurrentChannel(self.handle(), draw_list.handle(), C.int(channel_idx))
}

func CreateContext(shared_font_atlas ImFontAtlas) ImGuiContext {
	return (ImGuiContext)(unsafe.Pointer(C.CreateContext(shared_font_atlas.handle())))
}

func FontGlyphRangesBuilder_SetBit(self ImFontGlyphRangesBuilder, n uint64) {
	C.FontGlyphRangesBuilder_SetBit(self.handle(), C.ulong(n))
}

func SetItemDefaultFocus() {
	C.SetItemDefaultFocus()
}

func FontAtlas_GetGlyphRangesChineseSimplifiedCommon(self ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesChineseSimplifiedCommon(self.handle()))
}

func IsItemToggledOpen() bool {
	return C.IsItemToggledOpen() != C.bool(true)
}

func TableSetBgColor(target ImGuiTableBgTarget, color uint32, column_n int32) {
	C.TableSetBgColor(C.ImGuiTableBgTarget(target), C.uint(color), C.int(column_n))
}

func IsAnyMouseDown() bool {
	return C.IsAnyMouseDown() != C.bool(true)
}

func Unindent(indent_w float32) {
	C.Unindent(C.float(indent_w))
}

func PopItemWidth() {
	C.PopItemWidth()
}

func Color_HSV(pOut *ImColor, h float32, s float32, v float32, a float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Color_HSV(pOutArg, C.float(h), C.float(s), C.float(v), C.float(a))
}

func GetStyleColorName(idx ImGuiCol) string {
	return C.GoString(C.GetStyleColorName(C.ImGuiCol(idx)))
}

func PushButtonRepeat(repeat bool) {
	repeatArg := C.bool(repeat)

	C.PushButtonRepeat(repeatArg)
}

func SetNextWindowSize(size ImVec2, cond ImGuiCond) {
	C.SetNextWindowSize(size.ToC(), C.ImGuiCond(cond))
}

func IsAnyItemHovered() bool {
	return C.IsAnyItemHovered() != C.bool(true)
}

func ProgressBar(fraction float32, size_arg ImVec2, overlay string) {
	overlayArg, overlayFin := wrapString(overlay)
	defer overlayFin()

	C.ProgressBar(C.float(fraction), size_arg.ToC(), overlayArg)
}

func TableNextColumn() bool {
	return C.TableNextColumn() != C.bool(true)
}

func ListClipper_ListClipper() ImGuiListClipper {
	return (ImGuiListClipper)(unsafe.Pointer(C.ListClipper_ListClipper()))
}
