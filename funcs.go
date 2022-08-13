package cimgui

// #include "cimgui_wrapper.h"
import "C"

func BeginTabBar(str_id string, flags ImGuiTabBarFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTabBar(str_idArg, C.ImGuiTabBarFlags(flags)) != C.bool(true)
}

func TableSetColumnEnabled(column_n int32, v bool) {
	vArg := C.bool(v)

	C.TableSetColumnEnabled(C.int(column_n), vArg)
}

func EndListBox() {
	C.EndListBox()
}

func SetClipboardText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.SetClipboardText(textArg)
}

func DrawListSplitter_Split(self *ImDrawListSplitter, draw_list *ImDrawList, count int32) {
	C.DrawListSplitter_Split((*C.ImDrawListSplitter)(self), (*C.ImDrawList)(draw_list), C.int(count))
}

func IsAnyMouseDown() bool {
	return C.IsAnyMouseDown() != C.bool(true)
}

func IO_AddMouseWheelEvent(self *ImGuiIO, wh_x float32, wh_y float32) {
	C.IO_AddMouseWheelEvent((*C.ImGuiIO)(self), C.float(wh_x), C.float(wh_y))
}

func DrawList_GetClipRectMin(pOut *ImVec2, self *ImDrawList) {
	C.DrawList_GetClipRectMin((*C.ImVec2)(pOut), (*C.ImDrawList)(self))
}

func GetTextLineHeightWithSpacing() float32 {
	return float32(C.GetTextLineHeightWithSpacing())
}

func DrawList_ChannelsMerge(self *ImDrawList) {
	C.DrawList_ChannelsMerge((*C.ImDrawList)(self))
}

func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
	return C.BeginDragDropSource(C.ImGuiDragDropFlags(flags)) != C.bool(true)
}

func TableHeadersRow() {
	C.TableHeadersRow()
}

func TableNextColumn() bool {
	return C.TableNextColumn() != C.bool(true)
}

func EndTabItem() {
	C.EndTabItem()
}

func GetContentRegionMax(pOut *ImVec2) {
	C.GetContentRegionMax((*C.ImVec2)(pOut))
}

func FontAtlas_GetGlyphRangesCyrillic(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesCyrillic((*C.ImFontAtlas)(self)))
}

func IsWindowHovered(flags ImGuiHoveredFlags) bool {
	return C.IsWindowHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func Font_CalcWordWrapPositionA(self *ImFont, scale float32, text string, text_end string, wrap_width float32) string {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.GoString(C.Font_CalcWordWrapPositionA((*C.ImFont)(self), C.float(scale), textArg, text_endArg, C.float(wrap_width)))
}

func IsItemActivated() bool {
	return C.IsItemActivated() != C.bool(true)
}

func InputInt2(label string, v *[2]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt2(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetCursorStartPos(pOut *ImVec2) {
	C.GetCursorStartPos((*C.ImVec2)(pOut))
}

func PopItemWidth() {
	C.PopItemWidth()
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

func Spacing() {
	C.Spacing()
}

func Bullet() {
	C.Bullet()
}

func EndGroup() {
	C.EndGroup()
}

func GetScrollY() float32 {
	return float32(C.GetScrollY())
}

func IsKeyPressed(key ImGuiKey, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsKeyPressed(C.ImGuiKey(key), repeatArg) != C.bool(true)
}

func GetItemRectMax(pOut *ImVec2) {
	C.GetItemRectMax((*C.ImVec2)(pOut))
}

func GetScrollMaxY() float32 {
	return float32(C.GetScrollMaxY())
}

func InputInt4(label string, v *[4]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt4(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func FontAtlas_GetGlyphRangesChineseFull(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesChineseFull((*C.ImFontAtlas)(self)))
}

func FontAtlas_GetGlyphRangesVietnamese(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesVietnamese((*C.ImFontAtlas)(self)))
}

func BeginTooltip() {
	C.BeginTooltip()
}

func IsAnyItemActive() bool {
	return C.IsAnyItemActive() != C.bool(true)
}

func InputTextCallbackData_DeleteChars(self *ImGuiInputTextCallbackData, pos int32, bytes_count int32) {
	C.InputTextCallbackData_DeleteChars((*C.ImGuiInputTextCallbackData)(self), C.int(pos), C.int(bytes_count))
}

func DrawList_PrimReserve(self *ImDrawList, idx_count int32, vtx_count int32) {
	C.DrawList_PrimReserve((*C.ImDrawList)(self), C.int(idx_count), C.int(vtx_count))
}

func GetFrameHeightWithSpacing() float32 {
	return float32(C.GetFrameHeightWithSpacing())
}

func IsItemToggledOpen() bool {
	return C.IsItemToggledOpen() != C.bool(true)
}

func SliderInt2(label string, v *[2]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt2(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func FontGlyphRangesBuilder_SetBit(self *ImFontGlyphRangesBuilder, n uint64) {
	C.FontGlyphRangesBuilder_SetBit((*C.ImFontGlyphRangesBuilder)(self), C.ulong(n))
}

func EndChildFrame() {
	C.EndChildFrame()
}

func SetCursorPosX(local_x float32) {
	C.SetCursorPosX(C.float(local_x))
}

func GetWindowSize(pOut *ImVec2) {
	C.GetWindowSize((*C.ImVec2)(pOut))
}

func IsItemVisible() bool {
	return C.IsItemVisible() != C.bool(true)
}

func SetScrollHereY(center_y_ratio float32) {
	C.SetScrollHereY(C.float(center_y_ratio))
}

func ShowStyleEditor(ref *ImGuiStyle) {
	C.ShowStyleEditor((*C.ImGuiStyle)(ref))
}

func SetCursorPosY(local_y float32) {
	C.SetCursorPosY(C.float(local_y))
}

func IO_AddMouseButtonEvent(self *ImGuiIO, button int32, down bool) {
	downArg := C.bool(down)

	C.IO_AddMouseButtonEvent((*C.ImGuiIO)(self), C.int(button), downArg)
}

func TableNextRow(row_flags ImGuiTableRowFlags, min_row_height float32) {
	C.TableNextRow(C.ImGuiTableRowFlags(row_flags), C.float(min_row_height))
}

func DrawListSplitter_ClearFreeMemory(self *ImDrawListSplitter) {
	C.DrawListSplitter_ClearFreeMemory((*C.ImDrawListSplitter)(self))
}

func GetMousePosOnOpeningCurrentPopup(pOut *ImVec2) {
	C.GetMousePosOnOpeningCurrentPopup((*C.ImVec2)(pOut))
}

func GetMousePos(pOut *ImVec2) {
	C.GetMousePos((*C.ImVec2)(pOut))
}

func DragFloat2(label string, v *[2]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat2(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetItemRectMin(pOut *ImVec2) {
	C.GetItemRectMin((*C.ImVec2)(pOut))
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

func IsWindowCollapsed() bool {
	return C.IsWindowCollapsed() != C.bool(true)
}

func ShowDebugLogWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDebugLogWindow(p_openArg)
}

func GetStyleColorName(idx ImGuiCol) string {
	return C.GoString(C.GetStyleColorName(C.ImGuiCol(idx)))
}

func IsItemDeactivated() bool {
	return C.IsItemDeactivated() != C.bool(true)
}

func PopStyleVar(count int32) {
	C.PopStyleVar(C.int(count))
}

func Storage_SetAllInt(self *ImGuiStorage, val int32) {
	C.Storage_SetAllInt((*C.ImGuiStorage)(self), C.int(val))
}

func ArrowButton(str_id string, dir ImGuiDir) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ArrowButton(str_idArg, C.ImGuiDir(dir)) != C.bool(true)
}

func TableGetColumnCount() int {
	return int(C.TableGetColumnCount())
}

func TableGetRowIndex() int {
	return int(C.TableGetRowIndex())
}

func GetContentRegionAvail(pOut *ImVec2) {
	C.GetContentRegionAvail((*C.ImVec2)(pOut))
}

func PushItemWidth(item_width float32) {
	C.PushItemWidth(C.float(item_width))
}

func IsAnyItemFocused() bool {
	return C.IsAnyItemFocused() != C.bool(true)
}

func PushAllowKeyboardFocus(allow_keyboard_focus bool) {
	allow_keyboard_focusArg := C.bool(allow_keyboard_focus)

	C.PushAllowKeyboardFocus(allow_keyboard_focusArg)
}

func ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ShowStyleSelector(labelArg) != C.bool(true)
}

func FontAtlasCustomRect_IsPacked(self *ImFontAtlasCustomRect) bool {
	return C.FontAtlasCustomRect_IsPacked((*C.ImFontAtlasCustomRect)(self)) != C.bool(true)
}

func IsItemDeactivatedAfterEdit() bool {
	return C.IsItemDeactivatedAfterEdit() != C.bool(true)
}

func SliderFloat4(label string, v *[4]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat4(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Color_HSV(pOut *ImColor, h float32, s float32, v float32, a float32) {
	C.Color_HSV((*C.ImColor)(pOut), C.float(h), C.float(s), C.float(v), C.float(a))
}

func LogButtons() {
	C.LogButtons()
}

func EndChild() {
	C.EndChild()
}

func GetWindowWidth() float32 {
	return float32(C.GetWindowWidth())
}

func NewLine() {
	C.NewLine()
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

func DestroyPlatformWindows() {
	C.DestroyPlatformWindows()
}

func PopTextWrapPos() {
	C.PopTextWrapPos()
}

func DragInt3(label string, v *[3]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt3(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func ListClipper_Step(self *ImGuiListClipper) bool {
	return C.ListClipper_Step((*C.ImGuiListClipper)(self)) != C.bool(true)
}

func TabItemButton(label string, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TabItemButton(labelArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func IO_SetAppAcceptingEvents(self *ImGuiIO, accepting_events bool) {
	accepting_eventsArg := C.bool(accepting_events)

	C.IO_SetAppAcceptingEvents((*C.ImGuiIO)(self), accepting_eventsArg)
}

func Payload_Clear(self *ImGuiPayload) {
	C.Payload_Clear((*C.ImGuiPayload)(self))
}

func FontGlyphRangesBuilder_AddText(self *ImFontGlyphRangesBuilder, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.FontGlyphRangesBuilder_AddText((*C.ImFontGlyphRangesBuilder)(self), textArg, text_endArg)
}

func OpenPopupOnItemClick(str_id string, popup_flags ImGuiPopupFlags) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.OpenPopupOnItemClick(str_idArg, C.ImGuiPopupFlags(popup_flags))
}

func GetScrollMaxX() float32 {
	return float32(C.GetScrollMaxX())
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

func DragInt2(label string, v *[2]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt2(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DrawList_ChannelsSetCurrent(self *ImDrawList, n int32) {
	C.DrawList_ChannelsSetCurrent((*C.ImDrawList)(self), C.int(n))
}

func IO_AddInputCharactersUTF8(self *ImGuiIO, str string) {
	strArg, strFin := wrapString(str)
	defer strFin()

	C.IO_AddInputCharactersUTF8((*C.ImGuiIO)(self), strArg)
}

func BeginDragDropTarget() bool {
	return C.BeginDragDropTarget() != C.bool(true)
}

func ShowUserGuide() {
	C.ShowUserGuide()
}

func TextFilter_PassFilter(self *ImGuiTextFilter, text string, text_end string) bool {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.TextFilter_PassFilter((*C.ImGuiTextFilter)(self), textArg, text_endArg) != C.bool(true)
}

func IO_AddKeyEvent(self *ImGuiIO, key ImGuiKey, down bool) {
	downArg := C.bool(down)

	C.IO_AddKeyEvent((*C.ImGuiIO)(self), C.ImGuiKey(key), downArg)
}

func InputTextCallbackData_HasSelection(self *ImGuiInputTextCallbackData) bool {
	return C.InputTextCallbackData_HasSelection((*C.ImGuiInputTextCallbackData)(self)) != C.bool(true)
}

func IsItemActive() bool {
	return C.IsItemActive() != C.bool(true)
}

func FontAtlas_GetGlyphRangesJapanese(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesJapanese((*C.ImFontAtlas)(self)))
}

func GetColumnsCount() int {
	return int(C.GetColumnsCount())
}

func GetWindowDpiScale() float32 {
	return float32(C.GetWindowDpiScale())
}

func SetMouseCursor(cursor_type ImGuiMouseCursor) {
	C.SetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

func IO_AddFocusEvent(self *ImGuiIO, focused bool) {
	focusedArg := C.bool(focused)

	C.IO_AddFocusEvent((*C.ImGuiIO)(self), focusedArg)
}

func IsWindowDocked() bool {
	return C.IsWindowDocked() != C.bool(true)
}

func DrawList_PopTextureID(self *ImDrawList) {
	C.DrawList_PopTextureID((*C.ImDrawList)(self))
}

func SetStateStorage(storage *ImGuiStorage) {
	C.SetStateStorage((*C.ImGuiStorage)(storage))
}

func IO_AddKeyAnalogEvent(self *ImGuiIO, key ImGuiKey, down bool, v float32) {
	downArg := C.bool(down)

	C.IO_AddKeyAnalogEvent((*C.ImGuiIO)(self), C.ImGuiKey(key), downArg, C.float(v))
}

func SliderInt4(label string, v *[4]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt4(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func PopButtonRepeat() {
	C.PopButtonRepeat()
}

func SetKeyboardFocusHere(offset int32) {
	C.SetKeyboardFocusHere(C.int(offset))
}

func SliderFloat3(label string, v *[3]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat3(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsWindowFocused(flags ImGuiFocusedFlags) bool {
	return C.IsWindowFocused(C.ImGuiFocusedFlags(flags)) != C.bool(true)
}

func ResetMouseDragDelta(button ImGuiMouseButton) {
	C.ResetMouseDragDelta(C.ImGuiMouseButton(button))
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

func LogToTTY(auto_open_depth int32) {
	C.LogToTTY(C.int(auto_open_depth))
}

func BeginGroup() {
	C.BeginGroup()
}

func SliderInt3(label string, v *[3]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt3(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetMouseClickedCount(button ImGuiMouseButton) int {
	return int(C.GetMouseClickedCount(C.ImGuiMouseButton(button)))
}

func Separator() {
	C.Separator()
}

func DrawList_PopClipRect(self *ImDrawList) {
	C.DrawList_PopClipRect((*C.ImDrawList)(self))
}

func BeginPopup(str_id string, flags ImGuiWindowFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopup(str_idArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func SetColumnOffset(column_index int32, offset_x float32) {
	C.SetColumnOffset(C.int(column_index), C.float(offset_x))
}

func DestroyContext(ctx *ImGuiContext) {
	C.DestroyContext((*C.ImGuiContext)(ctx))
}

func EndDisabled() {
	C.EndDisabled()
}

func EndMainMenuBar() {
	C.EndMainMenuBar()
}

func ShowDemoWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDemoWindow(p_openArg)
}

func TextFilter_Build(self *ImGuiTextFilter) {
	C.TextFilter_Build((*C.ImGuiTextFilter)(self))
}

func EndCombo() {
	C.EndCombo()
}

func SaveIniSettingsToDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.SaveIniSettingsToDisk(ini_filenameArg)
}

func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond) {
	collapsedArg := C.bool(collapsed)

	C.SetNextWindowCollapsed(collapsedArg, C.ImGuiCond(cond))
}

func Viewport_GetWorkCenter(pOut *ImVec2, self *ImGuiViewport) {
	C.Viewport_GetWorkCenter((*C.ImVec2)(pOut), (*C.ImGuiViewport)(self))
}

func NextColumn() {
	C.NextColumn()
}

func GetFrameHeight() float32 {
	return float32(C.GetFrameHeight())
}

func EndTabBar() {
	C.EndTabBar()
}

func FontAtlas_ClearTexData(self *ImFontAtlas) {
	C.FontAtlas_ClearTexData((*C.ImFontAtlas)(self))
}

func FontAtlas_GetGlyphRangesChineseSimplifiedCommon(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesChineseSimplifiedCommon((*C.ImFontAtlas)(self)))
}

func IsItemHovered(flags ImGuiHoveredFlags) bool {
	return C.IsItemHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func ShowFontSelector(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ShowFontSelector(labelArg)
}

func LogToClipboard(auto_open_depth int32) {
	C.LogToClipboard(C.int(auto_open_depth))
}

func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	want_capture_mouseArg := C.bool(want_capture_mouse)

	C.SetNextFrameWantCaptureMouse(want_capture_mouseArg)
}

func IO_ClearInputKeys(self *ImGuiIO) {
	C.IO_ClearInputKeys((*C.ImGuiIO)(self))
}

func GetCursorPosX() float32 {
	return float32(C.GetCursorPosX())
}

func IsAnyItemHovered() bool {
	return C.IsAnyItemHovered() != C.bool(true)
}

func DrawList_ChannelsSplit(self *ImDrawList, count int32) {
	C.DrawList_ChannelsSplit((*C.ImDrawList)(self), C.int(count))
}

func TableSetupScrollFreeze(cols int32, rows int32) {
	C.TableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func DrawListSplitter_SetCurrentChannel(self *ImDrawListSplitter, draw_list *ImDrawList, channel_idx int32) {
	C.DrawListSplitter_SetCurrentChannel((*C.ImDrawListSplitter)(self), (*C.ImDrawList)(draw_list), C.int(channel_idx))
}

func Storage_Clear(self *ImGuiStorage) {
	C.Storage_Clear((*C.ImGuiStorage)(self))
}

func Font_ClearOutputData(self *ImFont) {
	C.Font_ClearOutputData((*C.ImFont)(self))
}

func Render() {
	C.Render()
}

func EndPopup() {
	C.EndPopup()
}

func TreePop() {
	C.TreePop()
}

func StyleColorsLight(dst *ImGuiStyle) {
	C.StyleColorsLight((*C.ImGuiStyle)(dst))
}

func IsMouseClicked(button ImGuiMouseButton, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsMouseClicked(C.ImGuiMouseButton(button), repeatArg) != C.bool(true)
}

func DrawList_PathClear(self *ImDrawList) {
	C.DrawList_PathClear((*C.ImDrawList)(self))
}

func InputFloat3(label string, v *[3]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat3(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
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

func BeginMenu(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	enabledArg := C.bool(enabled)

	return C.BeginMenu(labelArg, enabledArg) != C.bool(true)
}

func DrawListSplitter_Clear(self *ImDrawListSplitter) {
	C.DrawListSplitter_Clear((*C.ImDrawListSplitter)(self))
}

func PushButtonRepeat(repeat bool) {
	repeatArg := C.bool(repeat)

	C.PushButtonRepeat(repeatArg)
}

func StyleColorsClassic(dst *ImGuiStyle) {
	C.StyleColorsClassic((*C.ImGuiStyle)(dst))
}

func EndDragDropTarget() {
	C.EndDragDropTarget()
}

func TableGetColumnIndex() int {
	return int(C.TableGetColumnIndex())
}

func Viewport_GetCenter(pOut *ImVec2, self *ImGuiViewport) {
	C.Viewport_GetCenter((*C.ImVec2)(pOut), (*C.ImGuiViewport)(self))
}

func NewFrame() {
	C.NewFrame()
}

func SetScrollHereX(center_x_ratio float32) {
	C.SetScrollHereX(C.float(center_x_ratio))
}

func LoadIniSettingsFromDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.LoadIniSettingsFromDisk(ini_filenameArg)
}

func SetNextWindowBgAlpha(alpha float32) {
	C.SetNextWindowBgAlpha(C.float(alpha))
}

func TextBuffer_c_str(self *ImGuiTextBuffer) string {
	return C.GoString(C.TextBuffer_c_str((*C.ImGuiTextBuffer)(self)))
}

func BeginPopupContextItem(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextItem(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func EndTable() {
	C.EndTable()
}

func CalcTextSize(pOut *ImVec2, text string, text_end string, hide_text_after_double_hash bool, wrap_width float32) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	hide_text_after_double_hashArg := C.bool(hide_text_after_double_hash)

	C.CalcTextSize((*C.ImVec2)(pOut), textArg, text_endArg, hide_text_after_double_hashArg, C.float(wrap_width))
}

func Columns(count int32, id string, border bool) {
	idArg, idFin := wrapString(id)
	defer idFin()

	borderArg := C.bool(border)

	C.Columns(C.int(count), idArg, borderArg)
}

func InputFloat4(label string, v *[4]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat4(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func ColorEdit4(label string, col *[4]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit4(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func LogFinish() {
	C.LogFinish()
}

func GetWindowContentRegionMin(pOut *ImVec2) {
	C.GetWindowContentRegionMin((*C.ImVec2)(pOut))
}

func FontAtlas_IsBuilt(self *ImFontAtlas) bool {
	return C.FontAtlas_IsBuilt((*C.ImFontAtlas)(self)) != C.bool(true)
}

func Storage_BuildSortByKey(self *ImGuiStorage) {
	C.Storage_BuildSortByKey((*C.ImGuiStorage)(self))
}

func PopStyleColor(count int32) {
	C.PopStyleColor(C.int(count))
}

func EndDragDropSource() {
	C.EndDragDropSource()
}

func InputTextCallbackData_ClearSelection(self *ImGuiInputTextCallbackData) {
	C.InputTextCallbackData_ClearSelection((*C.ImGuiInputTextCallbackData)(self))
}

func IsItemClicked(mouse_button ImGuiMouseButton) bool {
	return C.IsItemClicked(C.ImGuiMouseButton(mouse_button)) != C.bool(true)
}

func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableHeader(labelArg)
}

func GetMouseDragDelta(pOut *ImVec2, button ImGuiMouseButton, lock_threshold float32) {
	C.GetMouseDragDelta((*C.ImVec2)(pOut), C.ImGuiMouseButton(button), C.float(lock_threshold))
}

func InputInt(label string, v *int32, step int32, step_fast int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	return C.InputInt(labelArg, vArg, C.int(step), C.int(step_fast), C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func DrawData_DeIndexAllBuffers(self *ImDrawData) {
	C.DrawData_DeIndexAllBuffers((*C.ImDrawData)(self))
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

func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return C.IsMouseDoubleClicked(C.ImGuiMouseButton(button)) != C.bool(true)
}

func PopClipRect() {
	C.PopClipRect()
}

func FontAtlas_Clear(self *ImFontAtlas) {
	C.FontAtlas_Clear((*C.ImFontAtlas)(self))
}

func FontAtlas_GetGlyphRangesDefault(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesDefault((*C.ImFontAtlas)(self)))
}

func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	want_capture_keyboardArg := C.bool(want_capture_keyboard)

	C.SetNextFrameWantCaptureKeyboard(want_capture_keyboardArg)
}

func FontGlyphRangesBuilder_GetBit(self *ImFontGlyphRangesBuilder, n uint64) bool {
	return C.FontGlyphRangesBuilder_GetBit((*C.ImFontGlyphRangesBuilder)(self), C.ulong(n)) != C.bool(true)
}

func DrawListSplitter_Merge(self *ImDrawListSplitter, draw_list *ImDrawList) {
	C.DrawListSplitter_Merge((*C.ImDrawListSplitter)(self), (*C.ImDrawList)(draw_list))
}

func ColorEdit3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func GetColumnIndex() int {
	return int(C.GetColumnIndex())
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

func GetKeyName(key ImGuiKey) string {
	return C.GoString(C.GetKeyName(C.ImGuiKey(key)))
}

func GetColumnWidth(column_index int32) float32 {
	return float32(C.GetColumnWidth(C.int(column_index)))
}

func EndTooltip() {
	C.EndTooltip()
}

func GetWindowContentRegionMax(pOut *ImVec2) {
	C.GetWindowContentRegionMax((*C.ImVec2)(pOut))
}

func FontAtlas_AddCustomRectRegular(self *ImFontAtlas, width int32, height int32) int {
	return int(C.FontAtlas_AddCustomRectRegular((*C.ImFontAtlas)(self), C.int(width), C.int(height)))
}

func DrawList_AddDrawCmd(self *ImDrawList) {
	C.DrawList_AddDrawCmd((*C.ImDrawList)(self))
}

func Style_ScaleAllSizes(self *ImGuiStyle, scale_factor float32) {
	C.Style_ScaleAllSizes((*C.ImGuiStyle)(self), C.float(scale_factor))
}

func Begin(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.Begin(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func IsItemFocused() bool {
	return C.IsItemFocused() != C.bool(true)
}

func SetNextWindowFocus() {
	C.SetNextWindowFocus()
}

func TextFilter_Draw(self *ImGuiTextFilter, label string, width float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TextFilter_Draw((*C.ImGuiTextFilter)(self), labelArg, C.float(width)) != C.bool(true)
}

func ListClipper_Begin(self *ImGuiListClipper, items_count int32, items_height float32) {
	C.ListClipper_Begin((*C.ImGuiListClipper)(self), C.int(items_count), C.float(items_height))
}

func IO_ClearInputCharacters(self *ImGuiIO) {
	C.IO_ClearInputCharacters((*C.ImGuiIO)(self))
}

func Checkbox(label string, v *bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapBool(v)
	defer vFin()

	return C.Checkbox(labelArg, vArg) != C.bool(true)
}

func GetFontTexUvWhitePixel(pOut *ImVec2) {
	C.GetFontTexUvWhitePixel((*C.ImVec2)(pOut))
}

func PopAllowKeyboardFocus() {
	C.PopAllowKeyboardFocus()
}

func DrawList_PushClipRectFullScreen(self *ImDrawList) {
	C.DrawList_PushClipRectFullScreen((*C.ImDrawList)(self))
}

func DebugTextEncoding(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.DebugTextEncoding(textArg)
}

func SetNextItemOpen(is_open bool, cond ImGuiCond) {
	is_openArg := C.bool(is_open)

	C.SetNextItemOpen(is_openArg, C.ImGuiCond(cond))
}

func BeginMainMenuBar() bool {
	return C.BeginMainMenuBar() != C.bool(true)
}

func SameLine(offset_from_start_x float32, spacing float32) {
	C.SameLine(C.float(offset_from_start_x), C.float(spacing))
}

func Font_IsLoaded(self *ImFont) bool {
	return C.Font_IsLoaded((*C.ImFont)(self)) != C.bool(true)
}

func SetCurrentContext(ctx *ImGuiContext) {
	C.SetCurrentContext((*C.ImGuiContext)(ctx))
}

func AlignTextToFramePadding() {
	C.AlignTextToFramePadding()
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

func DrawList_PrimUnreserve(self *ImDrawList, idx_count int32, vtx_count int32) {
	C.DrawList_PrimUnreserve((*C.ImDrawList)(self), C.int(idx_count), C.int(vtx_count))
}

func UpdatePlatformWindows() {
	C.UpdatePlatformWindows()
}

func SmallButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.SmallButton(labelArg) != C.bool(true)
}

func Font_GetDebugName(self *ImFont) string {
	return C.GoString(C.Font_GetDebugName((*C.ImFont)(self)))
}

func BeginCombo(label string, preview_value string, flags ImGuiComboFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	preview_valueArg, preview_valueFin := wrapString(preview_value)
	defer preview_valueFin()

	return C.BeginCombo(labelArg, preview_valueArg, C.ImGuiComboFlags(flags)) != C.bool(true)
}

func SetTabItemClosed(tab_or_docked_window_label string) {
	tab_or_docked_window_labelArg, tab_or_docked_window_labelFin := wrapString(tab_or_docked_window_label)
	defer tab_or_docked_window_labelFin()

	C.SetTabItemClosed(tab_or_docked_window_labelArg)
}

func GetKeyIndex(key ImGuiKey) int {
	return int(C.GetKeyIndex(C.ImGuiKey(key)))
}

func IO_SetKeyEventNativeData(self *ImGuiIO, key ImGuiKey, native_keycode int32, native_scancode int32, native_legacy_index int32) {
	C.IO_SetKeyEventNativeData((*C.ImGuiIO)(self), C.ImGuiKey(key), C.int(native_keycode), C.int(native_scancode), C.int(native_legacy_index))
}

func ColorPicker3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorPicker3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func GetVersion() string {
	return C.GoString(C.GetVersion())
}

func DragFloat3(label string, v *[3]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat3(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Indent(indent_w float32) {
	C.Indent(C.float(indent_w))
}

func IsKeyDown(key ImGuiKey) bool {
	return C.IsKeyDown(C.ImGuiKey(key)) != C.bool(true)
}

func GetCursorScreenPos(pOut *ImVec2) {
	C.GetCursorScreenPos((*C.ImVec2)(pOut))
}

func InputTextCallbackData_InsertChars(self *ImGuiInputTextCallbackData, pos int32, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.InputTextCallbackData_InsertChars((*C.ImGuiInputTextCallbackData)(self), C.int(pos), textArg, text_endArg)
}

func GetItemRectSize(pOut *ImVec2) {
	C.GetItemRectSize((*C.ImVec2)(pOut))
}

func Unindent(indent_w float32) {
	C.Unindent(C.float(indent_w))
}

func FontGlyphRangesBuilder_Clear(self *ImFontGlyphRangesBuilder) {
	C.FontGlyphRangesBuilder_Clear((*C.ImFontGlyphRangesBuilder)(self))
}

func End() {
	C.End()
}

func EndMenuBar() {
	C.EndMenuBar()
}

func FontAtlas_Build(self *ImFontAtlas) bool {
	return C.FontAtlas_Build((*C.ImFontAtlas)(self)) != C.bool(true)
}

func TableSetColumnIndex(column_n int32) bool {
	return C.TableSetColumnIndex(C.int(column_n)) != C.bool(true)
}

func GetWindowHeight() float32 {
	return float32(C.GetWindowHeight())
}

func PopID() {
	C.PopID()
}

func LogToFile(auto_open_depth int32, filename string) {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	C.LogToFile(C.int(auto_open_depth), filenameArg)
}

func PushFont(font *ImFont) {
	C.PushFont((*C.ImFont)(font))
}

func FontAtlas_ClearFonts(self *ImFontAtlas) {
	C.FontAtlas_ClearFonts((*C.ImFontAtlas)(self))
}

func SetColumnWidth(column_index int32, width float32) {
	C.SetColumnWidth(C.int(column_index), C.float(width))
}

func LoadIniSettingsFromMemory(ini_data string, ini_size uint64) {
	ini_dataArg, ini_dataFin := wrapString(ini_data)
	defer ini_dataFin()

	C.LoadIniSettingsFromMemory(ini_dataArg, C.ulong(ini_size))
}

func GetClipboardText() string {
	return C.GoString(C.GetClipboardText())
}

func SetNextItemWidth(item_width float32) {
	C.SetNextItemWidth(C.float(item_width))
}

func SetWindowFontScale(scale float32) {
	C.SetWindowFontScale(C.float(scale))
}

func BeginPopupModal(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginPopupModal(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func ShowAboutWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowAboutWindow(p_openArg)
}

func BeginPopupContextWindow(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextWindow(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func FontAtlas_ClearInputData(self *ImFontAtlas) {
	C.FontAtlas_ClearInputData((*C.ImFontAtlas)(self))
}

func Payload_IsDelivery(self *ImGuiPayload) bool {
	return C.Payload_IsDelivery((*C.ImGuiPayload)(self)) != C.bool(true)
}

func BeginPopupContextVoid(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextVoid(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func IO_AddMousePosEvent(self *ImGuiIO, x float32, y float32) {
	C.IO_AddMousePosEvent((*C.ImGuiIO)(self), C.float(x), C.float(y))
}

func TextUnformatted(text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.TextUnformatted(textArg, text_endArg)
}

func Payload_IsDataType(self *ImGuiPayload, typeArg string) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.Payload_IsDataType((*C.ImGuiPayload)(self), typeArgArg) != C.bool(true)
}

func DragFloat4(label string, v *[4]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat4(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Font_BuildLookupTable(self *ImFont) {
	C.Font_BuildLookupTable((*C.ImFont)(self))
}

func DrawData_Clear(self *ImDrawData) {
	C.DrawData_Clear((*C.ImDrawData)(self))
}

func BeginTabItem(label string, p_open *bool, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginTabItem(labelArg, p_openArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func IsItemEdited() bool {
	return C.IsItemEdited() != C.bool(true)
}

func EndMenu() {
	C.EndMenu()
}

func GetKeyPressedAmount(key ImGuiKey, repeat_delay float32, rate float32) int {
	return int(C.GetKeyPressedAmount(C.ImGuiKey(key), C.float(repeat_delay), C.float(rate)))
}

func FontAtlas_GetGlyphRangesKorean(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesKorean((*C.ImFontAtlas)(self)))
}

func DebugCheckVersionAndDataLayout(version_str string, sz_io uint64, sz_style uint64, sz_vec2 uint64, sz_vec4 uint64, sz_drawvert uint64, sz_drawidx uint64) bool {
	version_strArg, version_strFin := wrapString(version_str)
	defer version_strFin()

	return C.DebugCheckVersionAndDataLayout(version_strArg, C.ulong(sz_io), C.ulong(sz_style), C.ulong(sz_vec2), C.ulong(sz_vec4), C.ulong(sz_drawvert), C.ulong(sz_drawidx)) != C.bool(true)
}

func PushTextWrapPos(wrap_local_pos_x float32) {
	C.PushTextWrapPos(C.float(wrap_local_pos_x))
}

func GetColumnOffset(column_index int32) float32 {
	return float32(C.GetColumnOffset(C.int(column_index)))
}

func GetFrameCount() int {
	return int(C.GetFrameCount())
}

func ShowMetricsWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowMetricsWindow(p_openArg)
}

func Font_GrowIndex(self *ImFont, new_size int32) {
	C.Font_GrowIndex((*C.ImFont)(self), C.int(new_size))
}

func TextFilter_Clear(self *ImGuiTextFilter) {
	C.TextFilter_Clear((*C.ImGuiTextFilter)(self))
}

func BeginDisabled(disabled bool) {
	disabledArg := C.bool(disabled)

	C.BeginDisabled(disabledArg)
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

func SetColorEditOptions(flags ImGuiColorEditFlags) {
	C.SetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

func PopFont() {
	C.PopFont()
}

func StyleColorsDark(dst *ImGuiStyle) {
	C.StyleColorsDark((*C.ImGuiStyle)(dst))
}

func SetItemDefaultFocus() {
	C.SetItemDefaultFocus()
}

func TextFilter_IsActive(self *ImGuiTextFilter) bool {
	return C.TextFilter_IsActive((*C.ImGuiTextFilter)(self)) != C.bool(true)
}

func BeginMenuBar() bool {
	return C.BeginMenuBar() != C.bool(true)
}

func Color_SetHSV(self *ImColor, h float32, s float32, v float32, a float32) {
	C.Color_SetHSV((*C.ImColor)(self), C.float(h), C.float(s), C.float(v), C.float(a))
}

func ListClipper_ForceDisplayRangeByIndices(self *ImGuiListClipper, item_min int32, item_max int32) {
	C.ListClipper_ForceDisplayRangeByIndices((*C.ImGuiListClipper)(self), C.int(item_min), C.int(item_max))
}

func GetScrollX() float32 {
	return float32(C.GetScrollX())
}

func CalcItemWidth() float32 {
	return float32(C.CalcItemWidth())
}

func FontAtlas_GetGlyphRangesThai(self *ImFontAtlas) ImWchar {
	return ImWchar(C.FontAtlas_GetGlyphRangesThai((*C.ImFontAtlas)(self)))
}

func InputTextCallbackData_SelectAll(self *ImGuiInputTextCallbackData) {
	C.InputTextCallbackData_SelectAll((*C.ImGuiInputTextCallbackData)(self))
}

func Payload_IsPreview(self *ImGuiPayload) bool {
	return C.Payload_IsPreview((*C.ImGuiPayload)(self)) != C.bool(true)
}

func EndFrame() {
	C.EndFrame()
}

func GetCursorPos(pOut *ImVec2) {
	C.GetCursorPos((*C.ImVec2)(pOut))
}

func SliderFloat2(label string, v *[2]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat2(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsMouseDown(button ImGuiMouseButton) bool {
	return C.IsMouseDown(C.ImGuiMouseButton(button)) != C.bool(true)
}

func CloseCurrentPopup() {
	C.CloseCurrentPopup()
}

func GetCursorPosY() float32 {
	return float32(C.GetCursorPosY())
}

func SetItemAllowOverlap() {
	C.SetItemAllowOverlap()
}

func DrawList_GetClipRectMax(pOut *ImVec2, self *ImDrawList) {
	C.DrawList_GetClipRectMax((*C.ImVec2)(pOut), (*C.ImDrawList)(self))
}

func DragInt4(label string, v *[4]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt4(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsKeyReleased(key ImGuiKey) bool {
	return C.IsKeyReleased(C.ImGuiKey(key)) != C.bool(true)
}

func ShowStackToolWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowStackToolWindow(p_openArg)
}

func ListClipper_End(self *ImGuiListClipper) {
	C.ListClipper_End((*C.ImGuiListClipper)(self))
}

func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32) bool {
	return C.IsMouseDragging(C.ImGuiMouseButton(button), C.float(lock_threshold)) != C.bool(true)
}

func IsMouseReleased(button ImGuiMouseButton) bool {
	return C.IsMouseReleased(C.ImGuiMouseButton(button)) != C.bool(true)
}

func InputFloat2(label string, v *[2]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat2(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetWindowPos(pOut *ImVec2) {
	C.GetWindowPos((*C.ImVec2)(pOut))
}

func IsWindowAppearing() bool {
	return C.IsWindowAppearing() != C.bool(true)
}

func GetTreeNodeToLabelSpacing() float32 {
	return float32(C.GetTreeNodeToLabelSpacing())
}

func GetTextLineHeight() float32 {
	return float32(C.GetTextLineHeight())
}

func InputInt3(label string, v *[3]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt3(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetFontSize() float32 {
	return float32(C.GetFontSize())
}
