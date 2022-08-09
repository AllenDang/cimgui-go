package cimgui

// #include "cimgui_wrapper.h"
import "C"

func NewLine() {
C.NewLine()
}

func IsMouseHoveringRect(r_min ImVec2,r_max ImVec2,clip bool) bool {
clipArg := C.bool(clip)

return C.IsMouseHoveringRect(r_min.ToC(),r_max.ToC(),clipArg) != C.bool(true)
}

func TableSetupScrollFreeze(cols int32,rows int32) {
C.TableSetupScrollFreeze(C.int(cols),C.int(rows))
}

func Begin(name string,p_open *bool,flags ImGuiWindowFlags) bool {
nameArg, nameFin := wrapString(name)
defer nameFin()

p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

return C.Begin(nameArg,p_openArg,C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func BeginTable(str_id string,column int32,flags ImGuiTableFlags,outer_size ImVec2,inner_width float32) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.BeginTable(str_idArg,C.int(column),C.ImGuiTableFlags(flags),outer_size.ToC(),C.float(inner_width)) != C.bool(true)
}

func GetVersion() string {
return C.GoString(C.GetVersion())}

func SetCursorPosY(local_y float32) {
C.SetCursorPosY(C.float(local_y))
}

func GetItemRectSize(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetItemRectSize(pOutArg)
}

func LogToFile(auto_open_depth int32,filename string) {
filenameArg, filenameFin := wrapString(filename)
defer filenameFin()

C.LogToFile(C.int(auto_open_depth),filenameArg)
}

func PopFont() {
C.PopFont()
}

func AlignTextToFramePadding() {
C.AlignTextToFramePadding()
}

func GetCursorPosX() float32 {
return float32(C.GetCursorPosX())}

func GetKeyPressedAmount(key ImGuiKey,repeat_delay float32,rate float32) int {
return int(C.GetKeyPressedAmount(C.ImGuiKey(key),C.float(repeat_delay),C.float(rate)))}

func TableNextColumn() bool {
return C.TableNextColumn() != C.bool(true)
}

func DragIntRange2(label string,v_current_min *int32,v_current_max *int32,v_speed float32,v_min int32,v_max int32,format string,format_max string,flags ImGuiSliderFlags) bool {
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

return C.DragIntRange2(labelArg,v_current_minArg,v_current_maxArg,C.float(v_speed),C.int(v_min),C.int(v_max),formatArg,format_maxArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetColumnOffset(column_index int32) float32 {
return float32(C.GetColumnOffset(C.int(column_index)))}

func BeginMenuBar() bool {
return C.BeginMenuBar() != C.bool(true)
}

func SetScrollHereY(center_y_ratio float32) {
C.SetScrollHereY(C.float(center_y_ratio))
}

func Dummy(size ImVec2) {
C.Dummy(size.ToC())
}

func EndMenuBar() {
C.EndMenuBar()
}

func EndPopup() {
C.EndPopup()
}

func PopStyleColor(count int32) {
C.PopStyleColor(C.int(count))
}

func SliderFloat4(label string,v *[4]float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderFloat4(labelArg,vArg,C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetFrameCount() int {
return int(C.GetFrameCount())}

func BeginMainMenuBar() bool {
return C.BeginMainMenuBar() != C.bool(true)
}

func ShowUserGuide() {
C.ShowUserGuide()
}

func ColorButton(desc_id string,col ImVec4,flags ImGuiColorEditFlags,size ImVec2) bool {
desc_idArg, desc_idFin := wrapString(desc_id)
defer desc_idFin()

return C.ColorButton(desc_idArg,col.ToC(),C.ImGuiColorEditFlags(flags),size.ToC()) != C.bool(true)
}

func Separator() {
C.Separator()
}

func GetWindowHeight() float32 {
return float32(C.GetWindowHeight())}

func SetColumnOffset(column_index int32,offset_x float32) {
C.SetColumnOffset(C.int(column_index),C.float(offset_x))
}

func LoadIniSettingsFromDisk(ini_filename string) {
ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
defer ini_filenameFin()

C.LoadIniSettingsFromDisk(ini_filenameArg)
}

func ShowFontSelector(label string) {
labelArg, labelFin := wrapString(label)
defer labelFin()

C.ShowFontSelector(labelArg)
}

func NextColumn() {
C.NextColumn()
}

func GetMouseDragDelta(pOut *ImVec2,button ImGuiMouseButton,lock_threshold float32) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetMouseDragDelta(pOutArg,C.ImGuiMouseButton(button),C.float(lock_threshold))
}

func InputFloat4(label string,v *[4]float32,format string,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.InputFloat4(labelArg,vArg,formatArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func IsItemDeactivatedAfterEdit() bool {
return C.IsItemDeactivatedAfterEdit() != C.bool(true)
}

func IsKeyReleased(key ImGuiKey) bool {
return C.IsKeyReleased(C.ImGuiKey(key)) != C.bool(true)
}

func GetItemRectMax(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetItemRectMax(pOutArg)
}

func SetColumnWidth(column_index int32,width float32) {
C.SetColumnWidth(C.int(column_index),C.float(width))
}

func EndDragDropSource() {
C.EndDragDropSource()
}

func PushTextWrapPos(wrap_local_pos_x float32) {
C.PushTextWrapPos(C.float(wrap_local_pos_x))
}

func Spacing() {
C.Spacing()
}

func Columns(count int32,id string,border bool) {
idArg, idFin := wrapString(id)
defer idFin()

borderArg := C.bool(border)

C.Columns(C.int(count),idArg,borderArg)
}

func PopID() {
C.PopID()
}

func ResetMouseDragDelta(button ImGuiMouseButton) {
C.ResetMouseDragDelta(C.ImGuiMouseButton(button))
}

func InputInt2(label string,v *[2]int32,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

return C.InputInt2(labelArg,vArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func BeginGroup() {
C.BeginGroup()
}

func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
return C.IsMouseDoubleClicked(C.ImGuiMouseButton(button)) != C.bool(true)
}

func EndMainMenuBar() {
C.EndMainMenuBar()
}

func GetContentRegionMax(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetContentRegionMax(pOutArg)
}

func TreePop() {
C.TreePop()
}

func SetNextItemWidth(item_width float32) {
C.SetNextItemWidth(C.float(item_width))
}

func GetColumnsCount() int {
return int(C.GetColumnsCount())}

func EndDisabled() {
C.EndDisabled()
}

func BeginPopupContextWindow(str_id string,popup_flags ImGuiPopupFlags) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.BeginPopupContextWindow(str_idArg,C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func Checkbox(label string,v *bool) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapBool(v)
defer vFin()

return C.Checkbox(labelArg,vArg) != C.bool(true)
}

func GetColumnIndex() int {
return int(C.GetColumnIndex())}

func LogButtons() {
C.LogButtons()
}

func IsItemDeactivated() bool {
return C.IsItemDeactivated() != C.bool(true)
}

func SetNextItemOpen(is_open bool,cond ImGuiCond) {
is_openArg := C.bool(is_open)

C.SetNextItemOpen(is_openArg,C.ImGuiCond(cond))
}

func BeginPopupModal(name string,p_open *bool,flags ImGuiWindowFlags) bool {
nameArg, nameFin := wrapString(name)
defer nameFin()

p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

return C.BeginPopupModal(nameArg,p_openArg,C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func IsItemToggledOpen() bool {
return C.IsItemToggledOpen() != C.bool(true)
}

func IsMouseReleased(button ImGuiMouseButton) bool {
return C.IsMouseReleased(C.ImGuiMouseButton(button)) != C.bool(true)
}

func PushAllowKeyboardFocus(allow_keyboard_focus bool) {
allow_keyboard_focusArg := C.bool(allow_keyboard_focus)

C.PushAllowKeyboardFocus(allow_keyboard_focusArg)
}

func IsMouseDown(button ImGuiMouseButton) bool {
return C.IsMouseDown(C.ImGuiMouseButton(button)) != C.bool(true)
}

func BeginCombo(label string,preview_value string,flags ImGuiComboFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

preview_valueArg, preview_valueFin := wrapString(preview_value)
defer preview_valueFin()

return C.BeginCombo(labelArg,preview_valueArg,C.ImGuiComboFlags(flags)) != C.bool(true)
}

func BeginTabBar(str_id string,flags ImGuiTabBarFlags) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.BeginTabBar(str_idArg,C.ImGuiTabBarFlags(flags)) != C.bool(true)
}

func Bullet() {
C.Bullet()
}

func PopClipRect() {
C.PopClipRect()
}

func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
want_capture_keyboardArg := C.bool(want_capture_keyboard)

C.SetNextFrameWantCaptureKeyboard(want_capture_keyboardArg)
}

func DebugTextEncoding(text string) {
textArg, textFin := wrapString(text)
defer textFin()

C.DebugTextEncoding(textArg)
}

func EndTable() {
C.EndTable()
}

func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
want_capture_mouseArg := C.bool(want_capture_mouse)

C.SetNextFrameWantCaptureMouse(want_capture_mouseArg)
}

func SliderFloat2(label string,v *[2]float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderFloat2(labelArg,vArg,C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Button(label string,size ImVec2) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

return C.Button(labelArg,size.ToC()) != C.bool(true)
}

func IsItemHovered(flags ImGuiHoveredFlags) bool {
return C.IsItemHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func ColorPicker3(label string,col *[3]float32,flags ImGuiColorEditFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

colArg := (*C.float)(&col[0])

return C.ColorPicker3(labelArg,colArg,C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func SliderFloat3(label string,v *[3]float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderFloat3(labelArg,vArg,C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SliderInt3(label string,v *[3]int32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderInt3(labelArg,vArg,C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func VSliderInt(label string,size ImVec2,v *int32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapInt32(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.VSliderInt(labelArg,size.ToC(),vArg,C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsItemEdited() bool {
return C.IsItemEdited() != C.bool(true)
}

func DragInt3(label string,v *[3]int32,v_speed float32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragInt3(labelArg,vArg,C.float(v_speed),C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetCursorPos(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetCursorPos(pOutArg)
}

func TableSetColumnIndex(column_n int32) bool {
return C.TableSetColumnIndex(C.int(column_n)) != C.bool(true)
}

func GetTextLineHeightWithSpacing() float32 {
return float32(C.GetTextLineHeightWithSpacing())}

func BeginPopup(str_id string,flags ImGuiWindowFlags) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.BeginPopup(str_idArg,C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func ColorEdit3(label string,col *[3]float32,flags ImGuiColorEditFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

colArg := (*C.float)(&col[0])

return C.ColorEdit3(labelArg,colArg,C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func IsItemVisible() bool {
return C.IsItemVisible() != C.bool(true)
}

func SetKeyboardFocusHere(offset int32) {
C.SetKeyboardFocusHere(C.int(offset))
}

func TableNextRow(row_flags ImGuiTableRowFlags,min_row_height float32) {
C.TableNextRow(C.ImGuiTableRowFlags(row_flags),C.float(min_row_height))
}

func SliderFloat(label string,v *float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapFloat(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderFloat(labelArg,vArg,C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetMousePos(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetMousePos(pOutArg)
}

func DragInt4(label string,v *[4]int32,v_speed float32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragInt4(labelArg,vArg,C.float(v_speed),C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetWindowWidth() float32 {
return float32(C.GetWindowWidth())}

func IsMouseClicked(button ImGuiMouseButton,repeat bool) bool {
repeatArg := C.bool(repeat)

return C.IsMouseClicked(C.ImGuiMouseButton(button),repeatArg) != C.bool(true)
}

func GetCursorStartPos(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetCursorStartPos(pOutArg)
}

func GetMouseClickedCount(button ImGuiMouseButton) int {
return int(C.GetMouseClickedCount(C.ImGuiMouseButton(button)))}

func ProgressBar(fraction float32,size_arg ImVec2,overlay string) {
overlayArg, overlayFin := wrapString(overlay)
defer overlayFin()

C.ProgressBar(C.float(fraction),size_arg.ToC(),overlayArg)
}

func ArrowButton(str_id string,dir ImGuiDir) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.ArrowButton(str_idArg,C.ImGuiDir(dir)) != C.bool(true)
}

func IsWindowDocked() bool {
return C.IsWindowDocked() != C.bool(true)
}

func IsMouseDragging(button ImGuiMouseButton,lock_threshold float32) bool {
return C.IsMouseDragging(C.ImGuiMouseButton(button),C.float(lock_threshold)) != C.bool(true)
}

func SetTabItemClosed(tab_or_docked_window_label string) {
tab_or_docked_window_labelArg, tab_or_docked_window_labelFin := wrapString(tab_or_docked_window_label)
defer tab_or_docked_window_labelFin()

C.SetTabItemClosed(tab_or_docked_window_labelArg)
}

func IsItemActivated() bool {
return C.IsItemActivated() != C.bool(true)
}

func SetItemDefaultFocus() {
C.SetItemDefaultFocus()
}

func ColorConvertRGBtoHSV(r float32,g float32,b float32,out_h *float32,out_s *float32,out_v *float32) {
out_hArg, out_hFin := wrapFloat(out_h)
defer out_hFin()

out_sArg, out_sFin := wrapFloat(out_s)
defer out_sFin()

out_vArg, out_vFin := wrapFloat(out_v)
defer out_vFin()

C.ColorConvertRGBtoHSV(C.float(r),C.float(g),C.float(b),out_hArg,out_sArg,out_vArg)
}

func LogToTTY(auto_open_depth int32) {
C.LogToTTY(C.int(auto_open_depth))
}

func Render() {
C.Render()
}

func BeginTabItem(label string,p_open *bool,flags ImGuiTabItemFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

return C.BeginTabItem(labelArg,p_openArg,C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func IsAnyMouseDown() bool {
return C.IsAnyMouseDown() != C.bool(true)
}

func GetKeyName(key ImGuiKey) string {
return C.GoString(C.GetKeyName(C.ImGuiKey(key)))}

func ShowMetricsWindow(p_open *bool) {
p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

C.ShowMetricsWindow(p_openArg)
}

func DragFloat2(label string,v *[2]float32,v_speed float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragFloat2(labelArg,vArg,C.float(v_speed),C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func ShowStackToolWindow(p_open *bool) {
p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

C.ShowStackToolWindow(p_openArg)
}

func GetContentRegionAvail(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetContentRegionAvail(pOutArg)
}

func EndFrame() {
C.EndFrame()
}

func IsWindowFocused(flags ImGuiFocusedFlags) bool {
return C.IsWindowFocused(C.ImGuiFocusedFlags(flags)) != C.bool(true)
}

func SmallButton(label string) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

return C.SmallButton(labelArg) != C.bool(true)
}

func InputInt4(label string,v *[4]int32,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

return C.InputInt4(labelArg,vArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func VSliderFloat(label string,size ImVec2,v *float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapFloat(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.VSliderFloat(labelArg,size.ToC(),vArg,C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsAnyItemActive() bool {
return C.IsAnyItemActive() != C.bool(true)
}

func PopTextWrapPos() {
C.PopTextWrapPos()
}

func SetNextWindowPos(pos ImVec2,cond ImGuiCond,pivot ImVec2) {
C.SetNextWindowPos(pos.ToC(),C.ImGuiCond(cond),pivot.ToC())
}

func EndGroup() {
C.EndGroup()
}

func Unindent(indent_w float32) {
C.Unindent(C.float(indent_w))
}

func SetScrollHereX(center_x_ratio float32) {
C.SetScrollHereX(C.float(center_x_ratio))
}

func PopItemWidth() {
C.PopItemWidth()
}

func SliderInt2(label string,v *[2]int32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderInt2(labelArg,vArg,C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetClipboardText() string {
return C.GoString(C.GetClipboardText())}

func InputFloat2(label string,v *[2]float32,format string,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.InputFloat2(labelArg,vArg,formatArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetWindowPos(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetWindowPos(pOutArg)
}

func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
return C.BeginDragDropSource(C.ImGuiDragDropFlags(flags)) != C.bool(true)
}

func DragInt(label string,v *int32,v_speed float32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapInt32(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragInt(labelArg,vArg,C.float(v_speed),C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func OpenPopupOnItemClick(str_id string,popup_flags ImGuiPopupFlags) {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

C.OpenPopupOnItemClick(str_idArg,C.ImGuiPopupFlags(popup_flags))
}

func IsWindowAppearing() bool {
return C.IsWindowAppearing() != C.bool(true)
}

func SetItemAllowOverlap() {
C.SetItemAllowOverlap()
}

func TableSetColumnEnabled(column_n int32,v bool) {
vArg := C.bool(v)

C.TableSetColumnEnabled(C.int(column_n),vArg)
}

func IsWindowCollapsed() bool {
return C.IsWindowCollapsed() != C.bool(true)
}

func DragFloat4(label string,v *[4]float32,v_speed float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragFloat4(labelArg,vArg,C.float(v_speed),C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetMousePosOnOpeningCurrentPopup(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetMousePosOnOpeningCurrentPopup(pOutArg)
}

func GetFontTexUvWhitePixel(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetFontTexUvWhitePixel(pOutArg)
}

func GetScrollMaxY() float32 {
return float32(C.GetScrollMaxY())}

func Indent(indent_w float32) {
C.Indent(C.float(indent_w))
}

func IsKeyDown(key ImGuiKey) bool {
return C.IsKeyDown(C.ImGuiKey(key)) != C.bool(true)
}

func NewFrame() {
C.NewFrame()
}

func DragFloat(label string,v *float32,v_speed float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapFloat(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragFloat(labelArg,vArg,C.float(v_speed),C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndListBox() {
C.EndListBox()
}

func GetFrameHeightWithSpacing() float32 {
return float32(C.GetFrameHeightWithSpacing())}

func BeginPopupContextItem(str_id string,popup_flags ImGuiPopupFlags) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.BeginPopupContextItem(str_idArg,C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func TableGetColumnCount() int {
return int(C.TableGetColumnCount())}

func InvisibleButton(str_id string,size ImVec2,flags ImGuiButtonFlags) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.InvisibleButton(str_idArg,size.ToC(),C.ImGuiButtonFlags(flags)) != C.bool(true)
}

func BeginMenu(label string,enabled bool) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

enabledArg := C.bool(enabled)

return C.BeginMenu(labelArg,enabledArg) != C.bool(true)
}

func IsAnyItemHovered() bool {
return C.IsAnyItemHovered() != C.bool(true)
}

func PushClipRect(clip_rect_min ImVec2,clip_rect_max ImVec2,intersect_with_current_clip_rect bool) {
intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

C.PushClipRect(clip_rect_min.ToC(),clip_rect_max.ToC(),intersect_with_current_clip_rectArg)
}

func CloseCurrentPopup() {
C.CloseCurrentPopup()
}

func EndCombo() {
C.EndCombo()
}

func GetFontSize() float32 {
return float32(C.GetFontSize())}

func PopStyleVar(count int32) {
C.PopStyleVar(C.int(count))
}

func SliderInt(label string,v *int32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapInt32(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderInt(labelArg,vArg,C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DragFloat3(label string,v *[3]float32,v_speed float32,v_min float32,v_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragFloat3(labelArg,vArg,C.float(v_speed),C.float(v_min),C.float(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndTabBar() {
C.EndTabBar()
}

func LogToClipboard(auto_open_depth int32) {
C.LogToClipboard(C.int(auto_open_depth))
}

func ColorConvertHSVtoRGB(h float32,s float32,v float32,out_r *float32,out_g *float32,out_b *float32) {
out_rArg, out_rFin := wrapFloat(out_r)
defer out_rFin()

out_gArg, out_gFin := wrapFloat(out_g)
defer out_gFin()

out_bArg, out_bFin := wrapFloat(out_b)
defer out_bFin()

C.ColorConvertHSVtoRGB(C.float(h),C.float(s),C.float(v),out_rArg,out_gArg,out_bArg)
}

func ShowDebugLogWindow(p_open *bool) {
p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

C.ShowDebugLogWindow(p_openArg)
}

func GetWindowContentRegionMin(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetWindowContentRegionMin(pOutArg)
}

func GetScrollY() float32 {
return float32(C.GetScrollY())}

func InputFloat(label string,v *float32,step float32,step_fast float32,format string,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapFloat(v)
defer vFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.InputFloat(labelArg,vArg,C.float(step),C.float(step_fast),formatArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetFrameHeight() float32 {
return float32(C.GetFrameHeight())}

func IsAnyItemFocused() bool {
return C.IsAnyItemFocused() != C.bool(true)
}

func SetNextWindowSize(size ImVec2,cond ImGuiCond) {
C.SetNextWindowSize(size.ToC(),C.ImGuiCond(cond))
}

func TableGetRowIndex() int {
return int(C.TableGetRowIndex())}

func DragFloatRange2(label string,v_current_min *float32,v_current_max *float32,v_speed float32,v_min float32,v_max float32,format string,format_max string,flags ImGuiSliderFlags) bool {
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

return C.DragFloatRange2(labelArg,v_current_minArg,v_current_maxArg,C.float(v_speed),C.float(v_min),C.float(v_max),formatArg,format_maxArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetCursorPosY() float32 {
return float32(C.GetCursorPosY())}

func GetWindowSize(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetWindowSize(pOutArg)
}

func EndDragDropTarget() {
C.EndDragDropTarget()
}

func PopAllowKeyboardFocus() {
C.PopAllowKeyboardFocus()
}

func GetWindowContentRegionMax(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetWindowContentRegionMax(pOutArg)
}

func InputFloat3(label string,v *[3]float32,format string,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.float)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.InputFloat3(labelArg,vArg,formatArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func IsWindowHovered(flags ImGuiHoveredFlags) bool {
return C.IsWindowHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func SetClipboardText(text string) {
textArg, textFin := wrapString(text)
defer textFin()

C.SetClipboardText(textArg)
}

func SetMouseCursor(cursor_type ImGuiMouseCursor) {
C.SetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

func ColorEdit4(label string,col *[4]float32,flags ImGuiColorEditFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

colArg := (*C.float)(&col[0])

return C.ColorEdit4(labelArg,colArg,C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func GetTextLineHeight() float32 {
return float32(C.GetTextLineHeight())}

func SaveIniSettingsToDisk(ini_filename string) {
ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
defer ini_filenameFin()

C.SaveIniSettingsToDisk(ini_filenameArg)
}

func GetColumnWidth(column_index int32) float32 {
return float32(C.GetColumnWidth(C.int(column_index)))}

func GetScrollX() float32 {
return float32(C.GetScrollX())}

func ShowDemoWindow(p_open *bool) {
p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

C.ShowDemoWindow(p_openArg)
}

func EndTabItem() {
C.EndTabItem()
}

func SetNextWindowCollapsed(collapsed bool,cond ImGuiCond) {
collapsedArg := C.bool(collapsed)

C.SetNextWindowCollapsed(collapsedArg,C.ImGuiCond(cond))
}

func SetWindowFontScale(scale float32) {
C.SetWindowFontScale(C.float(scale))
}

func ShowStyleSelector(label string) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

return C.ShowStyleSelector(labelArg) != C.bool(true)
}

func ShowAboutWindow(p_open *bool) {
p_openArg, p_openFin := wrapBool(p_open)
defer p_openFin()

C.ShowAboutWindow(p_openArg)
}

func BeginTooltip() {
C.BeginTooltip()
}

func End() {
C.End()
}

func InputInt3(label string,v *[3]int32,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

return C.InputInt3(labelArg,vArg,C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func BeginDragDropTarget() bool {
return C.BeginDragDropTarget() != C.bool(true)
}

func DestroyPlatformWindows() {
C.DestroyPlatformWindows()
}

func TabItemButton(label string,flags ImGuiTabItemFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

return C.TabItemButton(labelArg,C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func GetItemRectMin(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetItemRectMin(pOutArg)
}

func SetCursorScreenPos(pos ImVec2) {
C.SetCursorScreenPos(pos.ToC())
}

func TableHeadersRow() {
C.TableHeadersRow()
}

func TableGetColumnIndex() int {
return int(C.TableGetColumnIndex())}

func DragInt2(label string,v *[2]int32,v_speed float32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.DragInt2(labelArg,vArg,C.float(v_speed),C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func EndTooltip() {
C.EndTooltip()
}

func PushItemWidth(item_width float32) {
C.PushItemWidth(C.float(item_width))
}

func CalcTextSize(pOut *ImVec2,text string,text_end string,hide_text_after_double_hash bool,wrap_width float32) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

textArg, textFin := wrapString(text)
defer textFin()

text_endArg, text_endFin := wrapString(text_end)
defer text_endFin()

hide_text_after_double_hashArg := C.bool(hide_text_after_double_hash)

C.CalcTextSize(pOutArg,textArg,text_endArg,hide_text_after_double_hashArg,C.float(wrap_width))
}

func SetNextWindowBgAlpha(alpha float32) {
C.SetNextWindowBgAlpha(C.float(alpha))
}

func TextUnformatted(text string,text_end string) {
textArg, textFin := wrapString(text)
defer textFin()

text_endArg, text_endFin := wrapString(text_end)
defer text_endFin()

C.TextUnformatted(textArg,text_endArg)
}

func EndMenu() {
C.EndMenu()
}

func GetWindowDpiScale() float32 {
return float32(C.GetWindowDpiScale())}

func InputInt(label string,v *int32,step int32,step_fast int32,flags ImGuiInputTextFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg, vFin := wrapInt32(v)
defer vFin()

return C.InputInt(labelArg,vArg,C.int(step),C.int(step_fast),C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetTreeNodeToLabelSpacing() float32 {
return float32(C.GetTreeNodeToLabelSpacing())}

func LoadIniSettingsFromMemory(ini_data string,ini_size uint64) {
ini_dataArg, ini_dataFin := wrapString(ini_data)
defer ini_dataFin()

C.LoadIniSettingsFromMemory(ini_dataArg,C.ulong(ini_size))
}

func SliderAngle(label string,v_rad *float32,v_degrees_min float32,v_degrees_max float32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

v_radArg, v_radFin := wrapFloat(v_rad)
defer v_radFin()

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderAngle(labelArg,v_radArg,C.float(v_degrees_min),C.float(v_degrees_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetNextWindowContentSize(size ImVec2) {
C.SetNextWindowContentSize(size.ToC())
}

func BeginDisabled(disabled bool) {
disabledArg := C.bool(disabled)

C.BeginDisabled(disabledArg)
}

func IsItemClicked(mouse_button ImGuiMouseButton) bool {
return C.IsItemClicked(C.ImGuiMouseButton(mouse_button)) != C.bool(true)
}

func SetNextWindowFocus() {
C.SetNextWindowFocus()
}

func BeginListBox(label string,size ImVec2) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

return C.BeginListBox(labelArg,size.ToC()) != C.bool(true)
}

func DebugCheckVersionAndDataLayout(version_str string,sz_io uint64,sz_style uint64,sz_vec2 uint64,sz_vec4 uint64,sz_drawvert uint64,sz_drawidx uint64) bool {
version_strArg, version_strFin := wrapString(version_str)
defer version_strFin()

return C.DebugCheckVersionAndDataLayout(version_strArg,C.ulong(sz_io),C.ulong(sz_style),C.ulong(sz_vec2),C.ulong(sz_vec4),C.ulong(sz_drawvert),C.ulong(sz_drawidx)) != C.bool(true)
}

func PushButtonRepeat(repeat bool) {
repeatArg := C.bool(repeat)

C.PushButtonRepeat(repeatArg)
}

func EndChild() {
C.EndChild()
}

func GetCursorScreenPos(pOut *ImVec2) {
pOutArg, pOutFin := wrapImVec2(pOut)
defer pOutFin()

C.GetCursorScreenPos(pOutArg)
}

func IsItemActive() bool {
return C.IsItemActive() != C.bool(true)
}

func SetColorEditOptions(flags ImGuiColorEditFlags) {
C.SetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

func IsItemFocused() bool {
return C.IsItemFocused() != C.bool(true)
}

func CalcItemWidth() float32 {
return float32(C.CalcItemWidth())}

func GetStyleColorName(idx ImGuiCol) string {
return C.GoString(C.GetStyleColorName(C.ImGuiCol(idx)))}

func EndChildFrame() {
C.EndChildFrame()
}

func GetScrollMaxX() float32 {
return float32(C.GetScrollMaxX())}

func PopButtonRepeat() {
C.PopButtonRepeat()
}

func GetKeyIndex(key ImGuiKey) int {
return int(C.GetKeyIndex(C.ImGuiKey(key)))}

func TableHeader(label string) {
labelArg, labelFin := wrapString(label)
defer labelFin()

C.TableHeader(labelArg)
}

func UpdatePlatformWindows() {
C.UpdatePlatformWindows()
}

func BeginPopupContextVoid(str_id string,popup_flags ImGuiPopupFlags) bool {
str_idArg, str_idFin := wrapString(str_id)
defer str_idFin()

return C.BeginPopupContextVoid(str_idArg,C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func LogFinish() {
C.LogFinish()
}

func SameLine(offset_from_start_x float32,spacing float32) {
C.SameLine(C.float(offset_from_start_x),C.float(spacing))
}

func SliderInt4(label string,v *[4]int32,v_min int32,v_max int32,format string,flags ImGuiSliderFlags) bool {
labelArg, labelFin := wrapString(label)
defer labelFin()

vArg := (*C.int)(&v[0])

formatArg, formatFin := wrapString(format)
defer formatFin()

return C.SliderInt4(labelArg,vArg,C.int(v_min),C.int(v_max),formatArg,C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsKeyPressed(key ImGuiKey,repeat bool) bool {
repeatArg := C.bool(repeat)

return C.IsKeyPressed(C.ImGuiKey(key),repeatArg) != C.bool(true)
}

func SetCursorPosX(local_x float32) {
C.SetCursorPosX(C.float(local_x))
}

func SetCursorPos(local_pos ImVec2) {
C.SetCursorPos(local_pos.ToC())
}

