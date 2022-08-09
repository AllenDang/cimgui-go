#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

void GetFontTexUvWhitePixel(ImVec2 *pOut) { igGetFontTexUvWhitePixel(pOut); }
bool IsAnyItemHovered() { return igIsAnyItemHovered(); }
bool BeginMenuBar() { return igBeginMenuBar(); }
ImGuiStorage *GetStateStorage() { return igGetStateStorage(); }
void GuiListClipper_Begin(ImGuiListClipper *self, int items_count, float items_height) { ImGuiListClipper_Begin(self, items_count, items_height); }
bool ArrowButton(const char *str_id, ImGuiDir dir) { return igArrowButton(str_id, dir); }
bool IsItemEdited() { return igIsItemEdited(); }
void GuiStyle_destroy(ImGuiStyle *self) { ImGuiStyle_destroy(self); }
bool InputDouble(const char *label, double *v, double step, double step_fast, const char *format, ImGuiInputTextFlags flags) { return igInputDouble(label, v, step, step_fast, format, flags); }
void GuiIO_AddInputCharacter(ImGuiIO *self, unsigned int c) { ImGuiIO_AddInputCharacter(self, c); }
void DrawList_AddImageRounded(ImDrawList *self, ImTextureID user_texture_id, const ImVec2 p_min, const ImVec2 p_max, const ImVec2 uv_min, const ImVec2 uv_max, ImU32 col, float rounding, ImDrawFlags flags) { ImDrawList_AddImageRounded(self, user_texture_id, p_min, p_max, uv_min, uv_max, col, rounding, flags); }
void GuiIO_AddInputCharactersUTF8(ImGuiIO *self, const char *str) { ImGuiIO_AddInputCharactersUTF8(self, str); }
void DrawListSharedData_destroy(ImDrawListSharedData *self) { ImDrawListSharedData_destroy(self); }
bool BeginPopupContextWindow(const char *str_id, ImGuiPopupFlags popup_flags) { return igBeginPopupContextWindow(str_id, popup_flags); }
void SetWindowFontScale(float scale) { igSetWindowFontScale(scale); }
bool GuiTextFilter_PassFilter(ImGuiTextFilter *self, const char *text, const char *text_end) { return ImGuiTextFilter_PassFilter(self, text, text_end); }
void DebugTextEncoding(const char *text) { igDebugTextEncoding(text); }
const ImGuiPayload *AcceptDragDropPayload(const char *type, ImGuiDragDropFlags flags) { return igAcceptDragDropPayload(type, flags); }
void DrawList_PathBezierCubicCurveTo(ImDrawList *self, const ImVec2 p2, const ImVec2 p3, const ImVec2 p4, int num_segments) { ImDrawList_PathBezierCubicCurveTo(self, p2, p3, p4, num_segments); }
void Font_GrowIndex(ImFont *self, int new_size) { ImFont_GrowIndex(self, new_size); }
void EndChildFrame() { igEndChildFrame(); }
void EndTooltip() { igEndTooltip(); }
void DrawList_PathLineToMergeDuplicate(ImDrawList *self, const ImVec2 pos) { ImDrawList_PathLineToMergeDuplicate(self, pos); }
bool FontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder *self, size_t n) { return ImFontGlyphRangesBuilder_GetBit(self, n); }
void ShowFontSelector(const char *label) { igShowFontSelector(label); }
void TableSetupScrollFreeze(int cols, int rows) { igTableSetupScrollFreeze(cols, rows); }
void LogFinish() { igLogFinish(); }
void FontAtlas_GetTexDataAsRGBA32(ImFontAtlas *self, unsigned char **out_pixels, int *out_width, int *out_height, int *out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsRGBA32(self, out_pixels, out_width, out_height, out_bytes_per_pixel); }
void GuiTextFilter_Build(ImGuiTextFilter *self) { ImGuiTextFilter_Build(self); }
ImGuiViewport *GetWindowViewport() { return igGetWindowViewport(); }
void GetCursorScreenPos(ImVec2 *pOut) { igGetCursorScreenPos(pOut); }
float GetScrollY() { return igGetScrollY(); }
bool InputText(const char *label, char *buf, size_t buf_size, ImGuiInputTextFlags flags, ImGuiInputTextCallback callback, void *user_data) { return igInputText(label, buf, buf_size, flags, callback, user_data); }
void SetColumnWidth(int column_index, float width) { igSetColumnWidth(column_index, width); }
void GetMousePosOnOpeningCurrentPopup(ImVec2 *pOut) { igGetMousePosOnOpeningCurrentPopup(pOut); }
bool GuiListClipper_Step(ImGuiListClipper *self) { return ImGuiListClipper_Step(self); }
void DestroyPlatformWindows() { igDestroyPlatformWindows(); }
void GuiWindowClass_destroy(ImGuiWindowClass *self) { ImGuiWindowClass_destroy(self); }
bool InputFloat3(const char *label, float v[3], const char *format, ImGuiInputTextFlags flags) { return igInputFloat3(label, v, format, flags); }
void EndMainMenuBar() { igEndMainMenuBar(); }
int FontAtlas_AddCustomRectRegular(ImFontAtlas *self, int width, int height) { return ImFontAtlas_AddCustomRectRegular(self, width, height); }
void GuiOnceUponAFrame_destroy(ImGuiOnceUponAFrame *self) { ImGuiOnceUponAFrame_destroy(self); }
void EndGroup() { igEndGroup(); }
ImFontGlyphRangesBuilder *FontGlyphRangesBuilder_ImFontGlyphRangesBuilder() { return ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(); }
bool IsMouseHoveringRect(const ImVec2 r_min, const ImVec2 r_max, bool clip) { return igIsMouseHoveringRect(r_min, r_max, clip); }
ImFontAtlasCustomRect *FontAtlas_GetCustomRectByIndex(ImFontAtlas *self, int index) { return ImFontAtlas_GetCustomRectByIndex(self, index); }
bool BeginTabItem(const char *label, bool *p_open, ImGuiTabItemFlags flags) { return igBeginTabItem(label, p_open, flags); }
void EndChild() { igEndChild(); }
int TableGetColumnIndex() { return igTableGetColumnIndex(); }
void TreePop() { igTreePop(); }
void DrawList_AddDrawCmd(ImDrawList *self) { ImDrawList_AddDrawCmd(self); }
bool IsWindowFocused(ImGuiFocusedFlags flags) { return igIsWindowFocused(flags); }
void NewFrame() { igNewFrame(); }
void PopTextWrapPos() { igPopTextWrapPos(); }
void GuiIO_destroy(ImGuiIO *self) { ImGuiIO_destroy(self); }
void ColorConvertU32ToFloat4(ImVec4 *pOut, ImU32 in) { igColorConvertU32ToFloat4(pOut, in); }
void LogToTTY(int auto_open_depth) { igLogToTTY(auto_open_depth); }
void EndDisabled() { igEndDisabled(); }
void DrawList_ChannelsSetCurrent(ImDrawList *self, int n) { ImDrawList_ChannelsSetCurrent(self, n); }
void DrawList_PathFillConvex(ImDrawList *self, ImU32 col) { ImDrawList_PathFillConvex(self, col); }
bool GuiTextBuffer_empty(ImGuiTextBuffer *self) { return ImGuiTextBuffer_empty(self); }
ImFont *FontAtlas_AddFontFromFileTTF(ImFontAtlas *self, const char *filename, float size_pixels, const ImFontConfig *font_cfg, const ImWchar *glyph_ranges) { return ImFontAtlas_AddFontFromFileTTF(self, filename, size_pixels, font_cfg, glyph_ranges); }
void DrawList_ChannelsMerge(ImDrawList *self) { ImDrawList_ChannelsMerge(self); }
void SetScrollHereY(float center_y_ratio) { igSetScrollHereY(center_y_ratio); }
void GuiStyleMod_destroy(ImGuiStyleMod *self) { ImGuiStyleMod_destroy(self); }
bool IsItemVisible() { return igIsItemVisible(); }
bool IsMouseDoubleClicked(ImGuiMouseButton button) { return igIsMouseDoubleClicked(button); }
void DrawList_PushClipRectFullScreen(ImDrawList *self) { ImDrawList_PushClipRectFullScreen(self); }
void ColorConvertRGBtoHSV(float r, float g, float b, float *out_h, float *out_s, float *out_v) { igColorConvertRGBtoHSV(r, g, b, out_h, out_s, out_v); }
void GuiContextHook_destroy(ImGuiContextHook *self) { ImGuiContextHook_destroy(self); }
void GuiTableSettings_destroy(ImGuiTableSettings *self) { ImGuiTableSettings_destroy(self); }
bool InputInt2(const char *label, int v[2], ImGuiInputTextFlags flags) { return igInputInt2(label, v, flags); }
void PopButtonRepeat() { igPopButtonRepeat(); }
void SetCursorPosY(float local_y) { igSetCursorPosY(local_y); }
int GuiTextBuffer_size(ImGuiTextBuffer *self) { return ImGuiTextBuffer_size(self); }
float GetTextLineHeight() { return igGetTextLineHeight(); }
ImGuiViewport *FindViewportByPlatformHandle(void *platform_handle) { return igFindViewportByPlatformHandle(platform_handle); }
void GuiTableSortSpecs_destroy(ImGuiTableSortSpecs *self) { ImGuiTableSortSpecs_destroy(self); }
void EndTabBar() { igEndTabBar(); }
bool DragFloat2(const char *label, float v[2], float v_speed, float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igDragFloat2(label, v, v_speed, v_min, v_max, format, flags); }
bool DragInt4(const char *label, int v[4], float v_speed, int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igDragInt4(label, v, v_speed, v_min, v_max, format, flags); }
void ShowStyleEditor(ImGuiStyle *ref) { igShowStyleEditor(ref); }
void FontAtlas_ClearInputData(ImFontAtlas *self) { ImFontAtlas_ClearInputData(self); }
void Indent(float indent_w) { igIndent(indent_w); }
bool SetDragDropPayload(const char *type, const void *data, size_t sz, ImGuiCond cond) { return igSetDragDropPayload(type, data, sz, cond); }
bool DragFloat3(const char *label, float v[3], float v_speed, float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igDragFloat3(label, v, v_speed, v_min, v_max, format, flags); }
void GetItemRectSize(ImVec2 *pOut) { igGetItemRectSize(pOut); }
bool IsMouseDragging(ImGuiMouseButton button, float lock_threshold) { return igIsMouseDragging(button, lock_threshold); }
bool SliderScalar(const char *label, ImGuiDataType data_type, void *p_data, const void *p_min, const void *p_max, const char *format, ImGuiSliderFlags flags) { return igSliderScalar(label, data_type, p_data, p_min, p_max, format, flags); }
void DrawList_PathStroke(ImDrawList *self, ImU32 col, ImDrawFlags flags, float thickness) { ImDrawList_PathStroke(self, col, flags, thickness); }
void BeginTooltip() { igBeginTooltip(); }
void Font_CalcTextSizeA(ImVec2 *pOut, ImFont *self, float size, float max_width, float wrap_width, const char *text_begin, const char *text_end, const char **remaining) { ImFont_CalcTextSizeA(pOut, self, size, max_width, wrap_width, text_begin, text_end, remaining); }
bool GuiInputTextCallbackData_HasSelection(ImGuiInputTextCallbackData *self) { return ImGuiInputTextCallbackData_HasSelection(self); }
bool InputFloat2(const char *label, float v[2], const char *format, ImGuiInputTextFlags flags) { return igInputFloat2(label, v, format, flags); }
void GuiStorage_SetInt(ImGuiStorage *self, ImGuiID key, int val) { ImGuiStorage_SetInt(self, key, val); }
void DrawList_PopClipRect(ImDrawList *self) { ImDrawList_PopClipRect(self); }
bool BeginPopup(const char *str_id, ImGuiWindowFlags flags) { return igBeginPopup(str_id, flags); }
void PushFont(ImFont *font) { igPushFont(font); }
void TableSetupColumn(const char *label, ImGuiTableColumnFlags flags, float init_width_or_weight, ImGuiID user_id) { igTableSetupColumn(label, flags, init_width_or_weight, user_id); }
bool DragScalar(const char *label, ImGuiDataType data_type, void *p_data, float v_speed, const void *p_min, const void *p_max, const char *format, ImGuiSliderFlags flags) { return igDragScalar(label, data_type, p_data, v_speed, p_min, p_max, format, flags); }
ImFont *FontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas *self, const void *compressed_font_data, int compressed_font_size, float size_pixels, const ImFontConfig *font_cfg, const ImWchar *glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedTTF(self, compressed_font_data, compressed_font_size, size_pixels, font_cfg, glyph_ranges); }
void Font_SetGlyphVisible(ImFont *self, ImWchar c, bool visible) { ImFont_SetGlyphVisible(self, c, visible); }
void FontAtlas_CalcCustomRectUV(ImFontAtlas *self, const ImFontAtlasCustomRect *rect, ImVec2 *out_uv_min, ImVec2 *out_uv_max) { ImFontAtlas_CalcCustomRectUV(self, rect, out_uv_min, out_uv_max); }
void DrawList_AddBezierCubic(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, const ImVec2 p4, ImU32 col, float thickness, int num_segments) { ImDrawList_AddBezierCubic(self, p1, p2, p3, p4, col, thickness, num_segments); }
void DrawList_AddLine(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, ImU32 col, float thickness) { ImDrawList_AddLine(self, p1, p2, col, thickness); }
void GuiIO_AddMouseViewportEvent(ImGuiIO *self, ImGuiID id) { ImGuiIO_AddMouseViewportEvent(self, id); }
bool IsItemClicked(ImGuiMouseButton mouse_button) { return igIsItemClicked(mouse_button); }
void DrawCmd_destroy(ImDrawCmd *self) { ImDrawCmd_destroy(self); }
const char *GetKeyName(ImGuiKey key) { return igGetKeyName(key); }
void DrawList_ChannelsSplit(ImDrawList *self, int count) { ImDrawList_ChannelsSplit(self, count); }
bool GuiPayload_IsDelivery(ImGuiPayload *self) { return ImGuiPayload_IsDelivery(self); }
void EndDragDropSource() { igEndDragDropSource(); }
void SetNextWindowContentSize(const ImVec2 size) { igSetNextWindowContentSize(size); }
ImGuiTableColumnFlags TableGetColumnFlags(int column_n) { return igTableGetColumnFlags(column_n); }
ImGuiViewport *GetMainViewport() { return igGetMainViewport(); }
void PushAllowKeyboardFocus(bool allow_keyboard_focus) { igPushAllowKeyboardFocus(allow_keyboard_focus); }
void Vec1_destroy(ImVec1 *self) { ImVec1_destroy(self); }
void GuiContext_destroy(ImGuiContext *self) { ImGuiContext_destroy(self); }
void GuiInputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData *self, int pos, int bytes_count) { ImGuiInputTextCallbackData_DeleteChars(self, pos, bytes_count); }
ImGuiTableColumnSortSpecs *GuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs() { return ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(); }
void PopAllowKeyboardFocus() { igPopAllowKeyboardFocus(); }
void SetColorEditOptions(ImGuiColorEditFlags flags) { igSetColorEditOptions(flags); }
bool Button(const char *label, const ImVec2 size) { return igButton(label, size); }
void DrawList_AddCircle(ImDrawList *self, const ImVec2 center, float radius, ImU32 col, int num_segments, float thickness) { ImDrawList_AddCircle(self, center, radius, col, num_segments, thickness); }
ImFontAtlas *FontAtlas_ImFontAtlas() { return ImFontAtlas_ImFontAtlas(); }
float GetCursorPosX() { return igGetCursorPosX(); }
void GetWindowContentRegionMin(ImVec2 *pOut) { igGetWindowContentRegionMin(pOut); }
bool ColorPicker3(const char *label, float col[3], ImGuiColorEditFlags flags) { return igColorPicker3(label, col, flags); }
void PushItemWidth(float item_width) { igPushItemWidth(item_width); }
void GetMouseDragDelta(ImVec2 *pOut, ImGuiMouseButton button, float lock_threshold) { igGetMouseDragDelta(pOut, button, lock_threshold); }
void FontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder *self, const ImWchar *ranges) { ImFontGlyphRangesBuilder_AddRanges(self, ranges); }
void AlignTextToFramePadding() { igAlignTextToFramePadding(); }
ImDrawData *DrawData_ImDrawData() { return ImDrawData_ImDrawData(); }
const ImWchar *FontAtlas_GetGlyphRangesVietnamese(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesVietnamese(self); }
void GuiStorage_SetVoidPtr(ImGuiStorage *self, ImGuiID key, void *val) { ImGuiStorage_SetVoidPtr(self, key, val); }
void TableSetColumnEnabled(int column_n, bool v) { igTableSetColumnEnabled(column_n, v); }
void SetNextItemOpen(bool is_open, ImGuiCond cond) { igSetNextItemOpen(is_open, cond); }
bool SliderFloat2(const char *label, float v[2], float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igSliderFloat2(label, v, v_min, v_max, format, flags); }
float GuiStorage_GetFloat(ImGuiStorage *self, ImGuiID key, float default_val) { return ImGuiStorage_GetFloat(self, key, default_val); }
bool DragIntRange2(const char *label, int *v_current_min, int *v_current_max, float v_speed, int v_min, int v_max, const char *format, const char *format_max, ImGuiSliderFlags flags) { return igDragIntRange2(label, v_current_min, v_current_max, v_speed, v_min, v_max, format, format_max, flags); }
void GuiComboPreviewData_destroy(ImGuiComboPreviewData *self) { ImGuiComboPreviewData_destroy(self); }
void PopFont() { igPopFont(); }
void DrawList_AddTriangleFilled(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, ImU32 col) { ImDrawList_AddTriangleFilled(self, p1, p2, p3, col); }
void GuiStyle_ScaleAllSizes(ImGuiStyle *self, float scale_factor) { ImGuiStyle_ScaleAllSizes(self, scale_factor); }
bool ColorEdit3(const char *label, float col[3], ImGuiColorEditFlags flags) { return igColorEdit3(label, col, flags); }
void PushClipRect(const ImVec2 clip_rect_min, const ImVec2 clip_rect_max, bool intersect_with_current_clip_rect) { igPushClipRect(clip_rect_min, clip_rect_max, intersect_with_current_clip_rect); }
const char *Font_CalcWordWrapPositionA(ImFont *self, float scale, const char *text, const char *text_end, float wrap_width) { return ImFont_CalcWordWrapPositionA(self, scale, text, text_end, wrap_width); }
bool BeginPopupContextItem(const char *str_id, ImGuiPopupFlags popup_flags) { return igBeginPopupContextItem(str_id, popup_flags); }
void LoadIniSettingsFromDisk(const char *ini_filename) { igLoadIniSettingsFromDisk(ini_filename); }
void FontAtlas_Clear(ImFontAtlas *self) { ImFontAtlas_Clear(self); }
void GuiViewport_GetCenter(ImVec2 *pOut, ImGuiViewport *self) { ImGuiViewport_GetCenter(pOut, self); }
void GetContentRegionAvail(ImVec2 *pOut) { igGetContentRegionAvail(pOut); }
void FontAtlas_destroy(ImFontAtlas *self) { ImFontAtlas_destroy(self); }
void GuiStorage_BuildSortByKey(ImGuiStorage *self) { ImGuiStorage_BuildSortByKey(self); }
ImGuiWindowClass *GuiWindowClass_ImGuiWindowClass() { return ImGuiWindowClass_ImGuiWindowClass(); }
void Color_HSV(ImColor *pOut, float h, float s, float v, float a) { ImColor_HSV(pOut, h, s, v, a); }
int GetMouseClickedCount(ImGuiMouseButton button) { return igGetMouseClickedCount(button); }
ImFont *FontAtlas_AddFontDefault(ImFontAtlas *self, const ImFontConfig *font_cfg) { return ImFontAtlas_AddFontDefault(self, font_cfg); }
bool IsKeyDown(ImGuiKey key) { return igIsKeyDown(key); }
bool GuiTextRange_empty(ImGuiTextRange *self) { return ImGuiTextRange_empty(self); }
void EndTabItem() { igEndTabItem(); }
void GuiListClipper_ForceDisplayRangeByIndices(ImGuiListClipper *self, int item_min, int item_max) { ImGuiListClipper_ForceDisplayRangeByIndices(self, item_min, item_max); }
void GuiTabBar_destroy(ImGuiTabBar *self) { ImGuiTabBar_destroy(self); }
void GuiTextBuffer_clear(ImGuiTextBuffer *self) { ImGuiTextBuffer_clear(self); }
void DrawList_AddTriangle(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, ImU32 col, float thickness) { ImDrawList_AddTriangle(self, p1, p2, p3, col, thickness); }
void DrawList_destroy(ImDrawList *self) { ImDrawList_destroy(self); }
void CloseCurrentPopup() { igCloseCurrentPopup(); }
bool DragFloatRange2(const char *label, float *v_current_min, float *v_current_max, float v_speed, float v_min, float v_max, const char *format, const char *format_max, ImGuiSliderFlags flags) { return igDragFloatRange2(label, v_current_min, v_current_max, v_speed, v_min, v_max, format, format_max, flags); }
void DrawList_AddPolyline(ImDrawList *self, const ImVec2 *points, int num_points, ImU32 col, ImDrawFlags flags, float thickness) { ImDrawList_AddPolyline(self, points, num_points, col, flags, thickness); }
int FontAtlas_AddCustomRectFontGlyph(ImFontAtlas *self, ImFont *font, ImWchar id, int width, int height, float advance_x, const ImVec2 offset) { return ImFontAtlas_AddCustomRectFontGlyph(self, font, id, width, height, advance_x, offset); }
bool BeginListBox(const char *label, const ImVec2 size) { return igBeginListBox(label, size); }
void StyleColorsLight(ImGuiStyle *dst) { igStyleColorsLight(dst); }
void EndDragDropTarget() { igEndDragDropTarget(); }
void EndTable() { igEndTable(); }
bool IsAnyItemActive() { return igIsAnyItemActive(); }
void Color_destroy(ImColor *self) { ImColor_destroy(self); }
ImFont *FontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas *self, const char *compressed_font_data_base85, float size_pixels, const ImFontConfig *font_cfg, const ImWchar *glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self, compressed_font_data_base85, size_pixels, font_cfg, glyph_ranges); }
ImFontAtlasCustomRect *FontAtlasCustomRect_ImFontAtlasCustomRect() { return ImFontAtlasCustomRect_ImFontAtlasCustomRect(); }
bool IsMouseDown(ImGuiMouseButton button) { return igIsMouseDown(button); }
ImFont *FontAtlas_AddFontFromMemoryTTF(ImFontAtlas *self, void *font_data, int font_size, float size_pixels, const ImFontConfig *font_cfg, const ImWchar *glyph_ranges) { return ImFontAtlas_AddFontFromMemoryTTF(self, font_data, font_size, size_pixels, font_cfg, glyph_ranges); }
bool FontAtlas_Build(ImFontAtlas *self) { return ImFontAtlas_Build(self); }
bool SliderFloat3(const char *label, float v[3], float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igSliderFloat3(label, v, v_min, v_max, format, flags); }
void GuiIO_AddInputCharacterUTF16(ImGuiIO *self, ImWchar16 c) { ImGuiIO_AddInputCharacterUTF16(self, c); }
void DrawList_AddRectFilled(ImDrawList *self, const ImVec2 p_min, const ImVec2 p_max, ImU32 col, float rounding, ImDrawFlags flags) { ImDrawList_AddRectFilled(self, p_min, p_max, col, rounding, flags); }
void GuiListClipperData_destroy(ImGuiListClipperData *self) { ImGuiListClipperData_destroy(self); }
ImGuiContext *CreateContext(ImFontAtlas *shared_font_atlas) { return igCreateContext(shared_font_atlas); }
void PopID() { igPopID(); }
ImGuiTableSortSpecs *TableGetSortSpecs() { return igTableGetSortSpecs(); }
bool BeginMainMenuBar() { return igBeginMainMenuBar(); }
void DrawListSplitter_ClearFreeMemory(ImDrawListSplitter *self) { ImDrawListSplitter_ClearFreeMemory(self); }
void DrawList_AddNgon(ImDrawList *self, const ImVec2 center, float radius, ImU32 col, int num_segments, float thickness) { ImDrawList_AddNgon(self, center, radius, col, num_segments, thickness); }
void DrawList_PrimUnreserve(ImDrawList *self, int idx_count, int vtx_count) { ImDrawList_PrimUnreserve(self, idx_count, vtx_count); }
ImGuiIO *GuiIO_ImGuiIO() { return ImGuiIO_ImGuiIO(); }
void GuiInputTextCallbackData_SelectAll(ImGuiInputTextCallbackData *self) { ImGuiInputTextCallbackData_SelectAll(self); }
void PopStyleVar(int count) { igPopStyleVar(count); }
void PushTextWrapPos(float wrap_local_pos_x) { igPushTextWrapPos(wrap_local_pos_x); }
void SetItemDefaultFocus() { igSetItemDefaultFocus(); }
bool SliderFloat4(const char *label, float v[4], float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igSliderFloat4(label, v, v_min, v_max, format, flags); }
void DrawList_PathClear(ImDrawList *self) { ImDrawList_PathClear(self); }
void GetWindowSize(ImVec2 *pOut) { igGetWindowSize(pOut); }
void DrawList_PrimReserve(ImDrawList *self, int idx_count, int vtx_count) { ImDrawList_PrimReserve(self, idx_count, vtx_count); }
void GuiIO_ClearInputCharacters(ImGuiIO *self) { ImGuiIO_ClearInputCharacters(self); }
float GetColumnOffset(int column_index) { return igGetColumnOffset(column_index); }
bool TabItemButton(const char *label, ImGuiTabItemFlags flags) { return igTabItemButton(label, flags); }
void TableSetBgColor(ImGuiTableBgTarget target, ImU32 color, int column_n) { igTableSetBgColor(target, color, column_n); }
void GuiIO_AddMousePosEvent(ImGuiIO *self, float x, float y) { ImGuiIO_AddMousePosEvent(self, x, y); }
bool BeginTabBar(const char *str_id, ImGuiTabBarFlags flags) { return igBeginTabBar(str_id, flags); }
float CalcItemWidth() { return igCalcItemWidth(); }
bool IsKeyReleased(ImGuiKey key) { return igIsKeyReleased(key); }
void SetNextWindowDockID(ImGuiID dock_id, ImGuiCond cond) { igSetNextWindowDockID(dock_id, cond); }
void DrawList_AddNgonFilled(ImDrawList *self, const ImVec2 center, float radius, ImU32 col, int num_segments) { ImDrawList_AddNgonFilled(self, center, radius, col, num_segments); }
void FontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder *self, const char *text, const char *text_end) { ImFontGlyphRangesBuilder_AddText(self, text, text_end); }
void LogButtons() { igLogButtons(); }
void DrawList_PrimRect(ImDrawList *self, const ImVec2 a, const ImVec2 b, ImU32 col) { ImDrawList_PrimRect(self, a, b, col); }
void DrawList_PrimVtx(ImDrawList *self, const ImVec2 pos, const ImVec2 uv, ImU32 col) { ImDrawList_PrimVtx(self, pos, uv, col); }
void UpdatePlatformWindows() { igUpdatePlatformWindows(); }
void GuiTableColumnSettings_destroy(ImGuiTableColumnSettings *self) { ImGuiTableColumnSettings_destroy(self); }
bool DragFloat4(const char *label, float v[4], float v_speed, float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igDragFloat4(label, v, v_speed, v_min, v_max, format, flags); }
void Render() { igRender(); }
void SetNextFrameWantCaptureMouse(bool want_capture_mouse) { igSetNextFrameWantCaptureMouse(want_capture_mouse); }
bool FontAtlas_IsBuilt(ImFontAtlas *self) { return ImFontAtlas_IsBuilt(self); }
bool BeginPopupContextVoid(const char *str_id, ImGuiPopupFlags popup_flags) { return igBeginPopupContextVoid(str_id, popup_flags); }
void LogToFile(int auto_open_depth, const char *filename) { igLogToFile(auto_open_depth, filename); }
void GuiStorage_Clear(ImGuiStorage *self) { ImGuiStorage_Clear(self); }
const char *GuiTextBuffer_begin(ImGuiTextBuffer *self) { return ImGuiTextBuffer_begin(self); }
void GuiViewport_GetWorkCenter(ImVec2 *pOut, ImGuiViewport *self) { ImGuiViewport_GetWorkCenter(pOut, self); }
bool InputScalarN(const char *label, ImGuiDataType data_type, void *p_data, int components, const void *p_step, const void *p_step_fast, const char *format, ImGuiInputTextFlags flags) { return igInputScalarN(label, data_type, p_data, components, p_step, p_step_fast, format, flags); }
void PopStyleColor(int count) { igPopStyleColor(count); }
void GuiStoragePair_destroy(ImGuiStoragePair *self) { ImGuiStoragePair_destroy(self); }
ImGuiViewport *GuiViewport_ImGuiViewport() { return ImGuiViewport_ImGuiViewport(); }
ImGuiMouseCursor GetMouseCursor() { return igGetMouseCursor(); }
bool InputInt(const char *label, int *v, int step, int step_fast, ImGuiInputTextFlags flags) { return igInputInt(label, v, step, step_fast, flags); }
void DrawListSplitter_destroy(ImDrawListSplitter *self) { ImDrawListSplitter_destroy(self); }
ImGuiPlatformIO *GuiPlatformIO_ImGuiPlatformIO() { return ImGuiPlatformIO_ImGuiPlatformIO(); }
bool IsWindowDocked() { return igIsWindowDocked(); }
bool IsAnyMouseDown() { return igIsAnyMouseDown(); }
void FontAtlasCustomRect_destroy(ImFontAtlasCustomRect *self) { ImFontAtlasCustomRect_destroy(self); }
bool IsWindowAppearing() { return igIsWindowAppearing(); }
const ImWchar *FontAtlas_GetGlyphRangesChineseFull(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesChineseFull(self); }
void Vec2_destroy(ImVec2 *self) { ImVec2_destroy(self); }
bool Font_IsGlyphRangeUnused(ImFont *self, unsigned int c_begin, unsigned int c_last) { return ImFont_IsGlyphRangeUnused(self, c_begin, c_last); }
const char *GuiTextBuffer_end(ImGuiTextBuffer *self) { return ImGuiTextBuffer_end(self); }
bool DebugCheckVersionAndDataLayout(const char *version_str, size_t sz_io, size_t sz_style, size_t sz_vec2, size_t sz_vec4, size_t sz_drawvert, size_t sz_drawidx) { return igDebugCheckVersionAndDataLayout(version_str, sz_io, sz_style, sz_vec2, sz_vec4, sz_drawvert, sz_drawidx); }
bool SliderScalarN(const char *label, ImGuiDataType data_type, void *p_data, int components, const void *p_min, const void *p_max, const char *format, ImGuiSliderFlags flags) { return igSliderScalarN(label, data_type, p_data, components, p_min, p_max, format, flags); }
void ShowDemoWindow(bool *p_open) { igShowDemoWindow(p_open); }
void GuiPtrOrIndex_destroy(ImGuiPtrOrIndex *self) { ImGuiPtrOrIndex_destroy(self); }
void GetCursorStartPos(ImVec2 *pOut) { igGetCursorStartPos(pOut); }
bool IsItemHovered(ImGuiHoveredFlags flags) { return igIsItemHovered(flags); }
bool IsItemToggledOpen() { return igIsItemToggledOpen(); }
void ResetMouseDragDelta(ImGuiMouseButton button) { igResetMouseDragDelta(button); }
double GetTime() { return igGetTime(); }
void DrawListSplitter_Merge(ImDrawListSplitter *self, ImDrawList *draw_list) { ImDrawListSplitter_Merge(self, draw_list); }
void GuiTextBuffer_append(ImGuiTextBuffer *self, const char *str, const char *str_end) { ImGuiTextBuffer_append(self, str, str_end); }
float GetCursorPosY() { return igGetCursorPosY(); }
bool TableNextColumn() { return igTableNextColumn(); }
void Font_destroy(ImFont *self) { ImFont_destroy(self); }
void GuiIO_SetKeyEventNativeData(ImGuiIO *self, ImGuiKey key, int native_keycode, int native_scancode, int native_legacy_index) { ImGuiIO_SetKeyEventNativeData(self, key, native_keycode, native_scancode, native_legacy_index); }
void GuiNextItemData_destroy(ImGuiNextItemData *self) { ImGuiNextItemData_destroy(self); }
void FontConf_destroy(ImFontConfig *self) { ImFontConfig_destroy(self); }
void DestroyContext(ImGuiContext *ctx) { igDestroyContext(ctx); }
void DrawList_AddCircleFilled(ImDrawList *self, const ImVec2 center, float radius, ImU32 col, int num_segments) { ImDrawList_AddCircleFilled(self, center, radius, col, num_segments); }
void Bullet() { igBullet(); }
ImDrawList *DrawList_ImDrawList(const ImDrawListSharedData *shared_data) { return ImDrawList_ImDrawList(shared_data); }
void DrawList_PathArcToFast(ImDrawList *self, const ImVec2 center, float radius, int a_min_of_12, int a_max_of_12) { ImDrawList_PathArcToFast(self, center, radius, a_min_of_12, a_max_of_12); }
void GuiPlatformMonitor_destroy(ImGuiPlatformMonitor *self) { ImGuiPlatformMonitor_destroy(self); }
void GuiTextRange_split(ImGuiTextRange *self, char separator, ImVector_ImGuiTextRange *out) { ImGuiTextRange_split(self, separator, out); }
void DrawData_Clear(ImDrawData *self) { ImDrawData_Clear(self); }
void GuiInputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData *self) { ImGuiInputTextCallbackData_ClearSelection(self); }
void Font_AddGlyph(ImFont *self, const ImFontConfig *src_cfg, ImWchar c, float x0, float y0, float x1, float y1, float u0, float v0, float u1, float v1, float advance_x) { ImFont_AddGlyph(self, src_cfg, c, x0, y0, x1, y1, u0, v0, u1, v1, advance_x); }
void DrawList_AddImageQuad(ImDrawList *self, ImTextureID user_texture_id, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, const ImVec2 p4, const ImVec2 uv1, const ImVec2 uv2, const ImVec2 uv3, const ImVec2 uv4, ImU32 col) { ImDrawList_AddImageQuad(self, user_texture_id, p1, p2, p3, p4, uv1, uv2, uv3, uv4, col); }
void SetItemAllowOverlap() { igSetItemAllowOverlap(); }
bool GuiTextFilter_IsActive(ImGuiTextFilter *self) { return ImGuiTextFilter_IsActive(self); }
bool BeginTable(const char *str_id, int column, ImGuiTableFlags flags, const ImVec2 outer_size, float inner_width) { return igBeginTable(str_id, column, flags, outer_size, inner_width); }
void FontAtlas_SetTexID(ImFontAtlas *self, ImTextureID id) { ImFontAtlas_SetTexID(self, id); }
void GetCursorPos(ImVec2 *pOut) { igGetCursorPos(pOut); }
float GetWindowDpiScale() { return igGetWindowDpiScale(); }
ImGuiIO *GetIO() { return igGetIO(); }
ImDrawList *GetWindowDrawList() { return igGetWindowDrawList(); }
void GuiIO_ClearInputKeys(ImGuiIO *self) { ImGuiIO_ClearInputKeys(self); }
void CalcTextSize(ImVec2 *pOut, const char *text, const char *text_end, bool hide_text_after_double_hash, float wrap_width) { igCalcTextSize(pOut, text, text_end, hide_text_after_double_hash, wrap_width); }
float GetScrollMaxX() { return igGetScrollMaxX(); }
bool SmallButton(const char *label) { return igSmallButton(label); }
bool ColorEdit4(const char *label, float col[4], ImGuiColorEditFlags flags) { return igColorEdit4(label, col, flags); }
void DrawListSplitter_SetCurrentChannel(ImDrawListSplitter *self, ImDrawList *draw_list, int channel_idx) { ImDrawListSplitter_SetCurrentChannel(self, draw_list, channel_idx); }
void BeginDisabled(bool disabled) { igBeginDisabled(disabled); }
void EndMenuBar() { igEndMenuBar(); }
float GetTextLineHeightWithSpacing() { return igGetTextLineHeightWithSpacing(); }
void DrawData_DeIndexAllBuffers(ImDrawData *self) { ImDrawData_DeIndexAllBuffers(self); }
void GuiViewport_destroy(ImGuiViewport *self) { ImGuiViewport_destroy(self); }
void SetNextWindowSizeConstraints(const ImVec2 size_min, const ImVec2 size_max, ImGuiSizeCallback custom_callback, void *custom_callback_data) { igSetNextWindowSizeConstraints(size_min, size_max, custom_callback, custom_callback_data); }
void FontAtlas_ClearFonts(ImFontAtlas *self) { ImFontAtlas_ClearFonts(self); }
ImGuiStyle *GuiStyle_ImGuiStyle() { return ImGuiStyle_ImGuiStyle(); }
float GetScrollMaxY() { return igGetScrollMaxY(); }
void GuiListClipper_End(ImGuiListClipper *self) { ImGuiListClipper_End(self); }
void GuiTableInstanceData_destroy(ImGuiTableInstanceData *self) { ImGuiTableInstanceData_destroy(self); }
void GuiTextFilter_destroy(ImGuiTextFilter *self) { ImGuiTextFilter_destroy(self); }
void FontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder *self, size_t n) { ImFontGlyphRangesBuilder_SetBit(self, n); }
int GetColumnIndex() { return igGetColumnIndex(); }
ImFont *GetFont() { return igGetFont(); }
void OpenPopupOnItemClick(const char *str_id, ImGuiPopupFlags popup_flags) { igOpenPopupOnItemClick(str_id, popup_flags); }
ImGuiListClipper *GuiListClipper_ImGuiListClipper() { return ImGuiListClipper_ImGuiListClipper(); }
void Font_RenderText(ImFont *self, ImDrawList *draw_list, float size, const ImVec2 pos, ImU32 col, const ImVec4 clip_rect, const char *text_begin, const char *text_end, float wrap_width, bool cpu_fine_clip) { ImFont_RenderText(self, draw_list, size, pos, col, clip_rect, text_begin, text_end, wrap_width, cpu_fine_clip); }
void GuiNavItemData_destroy(ImGuiNavItemData *self) { ImGuiNavItemData_destroy(self); }
bool BeginMenu(const char *label, bool enabled) { return igBeginMenu(label, enabled); }
bool Checkbox(const char *label, bool *v) { return igCheckbox(label, v); }
const char *GetClipboardText() { return igGetClipboardText(); }
void ShowStackToolWindow(bool *p_open) { igShowStackToolWindow(p_open); }
void GuiPopupData_destroy(ImGuiPopupData *self) { ImGuiPopupData_destroy(self); }
void GuiStorage_SetFloat(ImGuiStorage *self, ImGuiID key, float val) { ImGuiStorage_SetFloat(self, key, val); }
const ImFontGlyph *Font_FindGlyph(ImFont *self, ImWchar c) { return ImFont_FindGlyph(self, c); }
bool IsWindowHovered(ImGuiHoveredFlags flags) { return igIsWindowHovered(flags); }
void LogToClipboard(int auto_open_depth) { igLogToClipboard(auto_open_depth); }
void SetTabItemClosed(const char *tab_or_docked_window_label) { igSetTabItemClosed(tab_or_docked_window_label); }
bool BeginPopupModal(const char *name, bool *p_open, ImGuiWindowFlags flags) { return igBeginPopupModal(name, p_open, flags); }
void GetItemRectMax(ImVec2 *pOut) { igGetItemRectMax(pOut); }
void SetAllocatorFunctions(ImGuiMemAllocFunc alloc_func, ImGuiMemFreeFunc free_func, void *user_data) { igSetAllocatorFunctions(alloc_func, free_func, user_data); }
ImGuiID DockSpace(ImGuiID id, const ImVec2 size, ImGuiDockNodeFlags flags, const ImGuiWindowClass *window_class) { return igDockSpace(id, size, flags, window_class); }
const ImWchar *FontAtlas_GetGlyphRangesThai(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesThai(self); }
bool IsItemDeactivatedAfterEdit() { return igIsItemDeactivatedAfterEdit(); }
void DrawList_AddCallback(ImDrawList *self, ImDrawCallback callback, void *callback_data) { ImDrawList_AddCallback(self, callback, callback_data); }
void FontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder *self, ImVector_ImWchar *out_ranges) { ImFontGlyphRangesBuilder_BuildRanges(self, out_ranges); }
void GuiOldColumnData_destroy(ImGuiOldColumnData *self) { ImGuiOldColumnData_destroy(self); }
float GetWindowWidth() { return igGetWindowWidth(); }
const char *SaveIniSettingsToMemory(size_t *out_ini_size) { return igSaveIniSettingsToMemory(out_ini_size); }
ImDrawList *DrawList_CloneOutput(ImDrawList *self) { return ImDrawList_CloneOutput(self); }
void FontAtlas_ClearTexData(ImFontAtlas *self) { ImFontAtlas_ClearTexData(self); }
void GuiStackLevelInfo_destroy(ImGuiStackLevelInfo *self) { ImGuiStackLevelInfo_destroy(self); }
bool *GuiStorage_GetBoolRef(ImGuiStorage *self, ImGuiID key, bool default_val) { return ImGuiStorage_GetBoolRef(self, key, default_val); }
ImFont *FontAtlas_AddFont(ImFontAtlas *self, const ImFontConfig *font_cfg) { return ImFontAtlas_AddFont(self, font_cfg); }
void End() { igEnd(); }
float GetWindowHeight() { return igGetWindowHeight(); }
void GuiStackSizes_destroy(ImGuiStackSizes *self) { ImGuiStackSizes_destroy(self); }
float GetFrameHeight() { return igGetFrameHeight(); }
void SetStateStorage(ImGuiStorage *storage) { igSetStateStorage(storage); }
bool TableSetColumnIndex(int column_n) { return igTableSetColumnIndex(column_n); }
bool GuiPayload_IsPreview(ImGuiPayload *self) { return ImGuiPayload_IsPreview(self); }
void GuiPayload_destroy(ImGuiPayload *self) { ImGuiPayload_destroy(self); }
int *GuiStorage_GetIntRef(ImGuiStorage *self, ImGuiID key, int default_val) { return ImGuiStorage_GetIntRef(self, key, default_val); }
ImU32 ColorConvertFloat4ToU32(const ImVec4 in) { return igColorConvertFloat4ToU32(in); }
void Columns(int count, const char *id, bool border) { igColumns(count, id, border); }
float GetTreeNodeToLabelSpacing() { return igGetTreeNodeToLabelSpacing(); }
ImGuiID DockSpaceOverViewport(const ImGuiViewport *viewport, ImGuiDockNodeFlags flags, const ImGuiWindowClass *window_class) { return igDockSpaceOverViewport(viewport, flags, window_class); }
void SetCurrentContext(ImGuiContext *ctx) { igSetCurrentContext(ctx); }
void DrawList_PathArcTo(ImDrawList *self, const ImVec2 center, float radius, float a_min, float a_max, int num_segments) { ImDrawList_PathArcTo(self, center, radius, a_min, a_max, num_segments); }
void GuiMenuColumns_destroy(ImGuiMenuColumns *self) { ImGuiMenuColumns_destroy(self); }
void BeginGroup() { igBeginGroup(); }
bool IsItemActivated() { return igIsItemActivated(); }
void DrawList_AddImage(ImDrawList *self, ImTextureID user_texture_id, const ImVec2 p_min, const ImVec2 p_max, const ImVec2 uv_min, const ImVec2 uv_max, ImU32 col) { ImDrawList_AddImage(self, user_texture_id, p_min, p_max, uv_min, uv_max, col); }
void DrawList_AddRect(ImDrawList *self, const ImVec2 p_min, const ImVec2 p_max, ImU32 col, float rounding, ImDrawFlags flags, float thickness) { ImDrawList_AddRect(self, p_min, p_max, col, rounding, flags, thickness); }
void GuiListClipper_destroy(ImGuiListClipper *self) { ImGuiListClipper_destroy(self); }
bool VSliderInt(const char *label, const ImVec2 size, int *v, int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igVSliderInt(label, size, v, v_min, v_max, format, flags); }
void FontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder *self, ImWchar c) { ImFontGlyphRangesBuilder_AddChar(self, c); }
const char *GuiTextBuffer_c_str(ImGuiTextBuffer *self) { return ImGuiTextBuffer_c_str(self); }
bool GuiPayload_IsDataType(ImGuiPayload *self, const char *type) { return ImGuiPayload_IsDataType(self, type); }
bool DragInt2(const char *label, int v[2], float v_speed, int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igDragInt2(label, v, v_speed, v_min, v_max, format, flags); }
void DrawData_ScaleClipRects(ImDrawData *self, const ImVec2 fb_scale) { ImDrawData_ScaleClipRects(self, fb_scale); }
void SetCursorScreenPos(const ImVec2 pos) { igSetCursorScreenPos(pos); }
void DrawList_AddQuadFilled(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, const ImVec2 p4, ImU32 col) { ImDrawList_AddQuadFilled(self, p1, p2, p3, p4, col); }
void SetClipboardText(const char *text) { igSetClipboardText(text); }
void StyleColorsClassic(ImGuiStyle *dst) { igStyleColorsClassic(dst); }
void GuiLastItemData_destroy(ImGuiLastItemData *self) { ImGuiLastItemData_destroy(self); }
void GuiTextFilter_Clear(ImGuiTextFilter *self) { ImGuiTextFilter_Clear(self); }
bool BeginDragDropSource(ImGuiDragDropFlags flags) { return igBeginDragDropSource(flags); }
bool DragFloat(const char *label, float *v, float v_speed, float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igDragFloat(label, v, v_speed, v_min, v_max, format, flags); }
void GuiIO_AddFocusEvent(ImGuiIO *self, bool focused) { ImGuiIO_AddFocusEvent(self, focused); }
void GuiPlatformIO_destroy(ImGuiPlatformIO *self) { ImGuiPlatformIO_destroy(self); }
void ShowAboutWindow(bool *p_open) { igShowAboutWindow(p_open); }
bool IsItemDeactivated() { return igIsItemDeactivated(); }
bool SliderInt3(const char *label, int v[3], int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igSliderInt3(label, v, v_min, v_max, format, flags); }
void GuiStorage_SetAllInt(ImGuiStorage *self, int val) { ImGuiStorage_SetAllInt(self, val); }
ImDrawData *GetDrawData() { return igGetDrawData(); }
bool IsAnyItemFocused() { return igIsAnyItemFocused(); }
ImGuiID GetWindowDockID() { return igGetWindowDockID(); }
void MemFree(void *ptr) { igMemFree(ptr); }
void DrawList_AddConvexPolyFilled(ImDrawList *self, const ImVec2 *points, int num_points, ImU32 col) { ImDrawList_AddConvexPolyFilled(self, points, num_points, col); }
bool BeginCombo(const char *label, const char *preview_value, ImGuiComboFlags flags) { return igBeginCombo(label, preview_value, flags); }
int GetColumnsCount() { return igGetColumnsCount(); }
float *GuiStorage_GetFloatRef(ImGuiStorage *self, ImGuiID key, float default_val) { return ImGuiStorage_GetFloatRef(self, key, default_val); }
ImGuiTextBuffer *GuiTextBuffer_ImGuiTextBuffer() { return ImGuiTextBuffer_ImGuiTextBuffer(); }
void DrawList_AddBezierQuadratic(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, ImU32 col, float thickness, int num_segments) { ImDrawList_AddBezierQuadratic(self, p1, p2, p3, col, thickness, num_segments); }
bool FontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect *self) { return ImFontAtlasCustomRect_IsPacked(self); }
bool VSliderFloat(const char *label, const ImVec2 size, float *v, float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igVSliderFloat(label, size, v, v_min, v_max, format, flags); }
bool SliderInt(const char *label, int *v, int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igSliderInt(label, v, v_min, v_max, format, flags); }
void GuiInputEvent_destroy(ImGuiInputEvent *self) { ImGuiInputEvent_destroy(self); }
void GuiStackTool_destroy(ImGuiStackTool *self) { ImGuiStackTool_destroy(self); }
bool IsKeyPressed(ImGuiKey key, bool repeat) { return igIsKeyPressed(key, repeat); }
void GuiMetricsConf_destroy(ImGuiMetricsConfig *self) { ImGuiMetricsConfig_destroy(self); }
void *GuiStorage_GetVoidPtr(ImGuiStorage *self, ImGuiID key) { return ImGuiStorage_GetVoidPtr(self, key); }
ImGuiContext *GetCurrentContext() { return igGetCurrentContext(); }
bool DragInt(const char *label, int *v, float v_speed, int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igDragInt(label, v, v_speed, v_min, v_max, format, flags); }
void StyleColorsDark(ImGuiStyle *dst) { igStyleColorsDark(dst); }
bool InputTextMultiline(const char *label, char *buf, size_t buf_size, const ImVec2 size, ImGuiInputTextFlags flags, ImGuiInputTextCallback callback, void *user_data) { return igInputTextMultiline(label, buf, buf_size, size, flags, callback, user_data); }
void SetNextWindowBgAlpha(float alpha) { igSetNextWindowBgAlpha(alpha); }
const ImWchar *FontAtlas_GetGlyphRangesJapanese(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesJapanese(self); }
void GuiTextRange_destroy(ImGuiTextRange *self) { ImGuiTextRange_destroy(self); }
bool ColorPicker4(const char *label, float col[4], ImGuiColorEditFlags flags, const float *ref_col) { return igColorPicker4(label, col, flags, ref_col); }
void TableHeader(const char *label) { igTableHeader(label); }
const ImVec4 *GetStyleColorVec4(ImGuiCol idx) { return igGetStyleColorVec4(idx); }
void GuiNextWindowData_destroy(ImGuiNextWindowData *self) { ImGuiNextWindowData_destroy(self); }
ImGuiTableSortSpecs *GuiTableSortSpecs_ImGuiTableSortSpecs() { return ImGuiTableSortSpecs_ImGuiTableSortSpecs(); }
bool Begin(const char *name, bool *p_open, ImGuiWindowFlags flags) { return igBegin(name, p_open, flags); }
void EndMenu() { igEndMenu(); }
void GuiTextBuffer_reserve(ImGuiTextBuffer *self, int capacity) { ImGuiTextBuffer_reserve(self, capacity); }
ImGuiPlatformIO *GetPlatformIO() { return igGetPlatformIO(); }
void DrawList_PushTextureID(ImDrawList *self, ImTextureID texture_id) { ImDrawList_PushTextureID(self, texture_id); }
const char *Font_GetDebugName(ImFont *self) { return ImFont_GetDebugName(self); }
bool ColorButton(const char *desc_id, const ImVec4 col, ImGuiColorEditFlags flags, const ImVec2 size) { return igColorButton(desc_id, col, flags, size); }
void ShowMetricsWindow(bool *p_open) { igShowMetricsWindow(p_open); }
void Color_SetHSV(ImColor *self, float h, float s, float v, float a) { ImColor_SetHSV(self, h, s, v, a); }
const ImWchar *FontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self); }
void EndListBox() { igEndListBox(); }
void GetWindowPos(ImVec2 *pOut) { igGetWindowPos(pOut); }
void GuiStorage_SetBool(ImGuiStorage *self, ImGuiID key, bool val) { ImGuiStorage_SetBool(self, key, val); }
bool IsMouseClicked(ImGuiMouseButton button, bool repeat) { return igIsMouseClicked(button, repeat); }
bool InputInt3(const char *label, int v[3], ImGuiInputTextFlags flags) { return igInputInt3(label, v, flags); }
void GuiIO_AddMouseWheelEvent(ImGuiIO *self, float wh_x, float wh_y) { ImGuiIO_AddMouseWheelEvent(self, wh_x, wh_y); }
void NextColumn() { igNextColumn(); }
bool DragScalarN(const char *label, ImGuiDataType data_type, void *p_data, int components, float v_speed, const void *p_min, const void *p_max, const char *format, ImGuiSliderFlags flags) { return igDragScalarN(label, data_type, p_data, components, v_speed, p_min, p_max, format, flags); }
const char *GetVersion() { return igGetVersion(); }
bool IsWindowCollapsed() { return igIsWindowCollapsed(); }
void *MemAlloc(size_t size) { return igMemAlloc(size); }
bool SliderFloat(const char *label, float *v, float v_min, float v_max, const char *format, ImGuiSliderFlags flags) { return igSliderFloat(label, v, v_min, v_max, format, flags); }
void GetAllocatorFunctions(ImGuiMemAllocFunc *p_alloc_func, ImGuiMemFreeFunc *p_free_func, void **p_user_data) { igGetAllocatorFunctions(p_alloc_func, p_free_func, p_user_data); }
ImGuiStyle *GetStyle() { return igGetStyle(); }
bool InputFloat(const char *label, float *v, float step, float step_fast, const char *format, ImGuiInputTextFlags flags) { return igInputFloat(label, v, step, step_fast, format, flags); }
bool FontAtlas_GetMouseCursorTexData(ImFontAtlas *self, ImGuiMouseCursor cursor, ImVec2 *out_offset, ImVec2 *out_size, ImVec2 out_uv_border[2], ImVec2 out_uv_fill[2]) { return ImFontAtlas_GetMouseCursorTexData(self, cursor, out_offset, out_size, out_uv_border, out_uv_fill); }
ImFont *Font_ImFont() { return ImFont_ImFont(); }
void GuiDockContext_destroy(ImGuiDockContext *self) { ImGuiDockContext_destroy(self); }
void EndFrame() { igEndFrame(); }
void ColorConvertHSVtoRGB(float h, float s, float v, float *out_r, float *out_g, float *out_b) { igColorConvertHSVtoRGB(h, s, v, out_r, out_g, out_b); }
void SetNextWindowSize(const ImVec2 size, ImGuiCond cond) { igSetNextWindowSize(size, cond); }
void DrawList_PopTextureID(ImDrawList *self) { ImDrawList_PopTextureID(self); }
void Vec4_destroy(ImVec4 *self) { ImVec4_destroy(self); }
bool BeginChildFrame(ImGuiID id, const ImVec2 size, ImGuiWindowFlags flags) { return igBeginChildFrame(id, size, flags); }
void **GuiStorage_GetVoidPtrRef(ImGuiStorage *self, ImGuiID key, void *default_val) { return ImGuiStorage_GetVoidPtrRef(self, key, default_val); }
void Vec2ih_destroy(ImVec2ih *self) { ImVec2ih_destroy(self); }
void SaveIniSettingsToDisk(const char *ini_filename) { igSaveIniSettingsToDisk(ini_filename); }
void DrawList_GetClipRectMax(ImVec2 *pOut, ImDrawList *self) { ImDrawList_GetClipRectMax(pOut, self); }
void FontGlyphRangesBuilder_destroy(ImFontGlyphRangesBuilder *self) { ImFontGlyphRangesBuilder_destroy(self); }
bool InputTextWithHint(const char *label, const char *hint, char *buf, size_t buf_size, ImGuiInputTextFlags flags, ImGuiInputTextCallback callback, void *user_data) { return igInputTextWithHint(label, hint, buf, buf_size, flags, callback, user_data); }
bool SliderInt4(const char *label, int v[4], int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igSliderInt4(label, v, v_min, v_max, format, flags); }
void DrawListSplitter_Clear(ImDrawListSplitter *self) { ImDrawListSplitter_Clear(self); }
const ImWchar *FontAtlas_GetGlyphRangesDefault(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesDefault(self); }
void GuiSettingsHandler_destroy(ImGuiSettingsHandler *self) { ImGuiSettingsHandler_destroy(self); }
void DrawList_AddQuad(ImDrawList *self, const ImVec2 p1, const ImVec2 p2, const ImVec2 p3, const ImVec2 p4, ImU32 col, float thickness) { ImDrawList_AddQuad(self, p1, p2, p3, p4, col, thickness); }
void LoadIniSettingsFromMemory(const char *ini_data, size_t ini_size) { igLoadIniSettingsFromMemory(ini_data, ini_size); }
bool Font_IsLoaded(ImFont *self) { return ImFont_IsLoaded(self); }
void Font_RenderChar(ImFont *self, ImDrawList *draw_list, float size, const ImVec2 pos, ImU32 col, ImWchar c) { ImFont_RenderChar(self, draw_list, size, pos, col, c); }
void Dummy(const ImVec2 size) { igDummy(size); }
const char *GetStyleColorName(ImGuiCol idx) { return igGetStyleColorName(idx); }
bool InputScalar(const char *label, ImGuiDataType data_type, void *p_data, const void *p_step, const void *p_step_fast, const char *format, ImGuiInputTextFlags flags) { return igInputScalar(label, data_type, p_data, p_step, p_step_fast, format, flags); }
void SetCursorPos(const ImVec2 local_pos) { igSetCursorPos(local_pos); }
void DrawList_PrimWriteVtx(ImDrawList *self, const ImVec2 pos, const ImVec2 uv, ImU32 col) { ImDrawList_PrimWriteVtx(self, pos, uv, col); }
bool IsMousePosValid(const ImVec2 *mouse_pos) { return igIsMousePosValid(mouse_pos); }
void SetKeyboardFocusHere(int offset) { igSetKeyboardFocusHere(offset); }
void TableHeadersRow() { igTableHeadersRow(); }
bool InputFloat4(const char *label, float v[4], const char *format, ImGuiInputTextFlags flags) { return igInputFloat4(label, v, format, flags); }
void ShowDebugLogWindow(bool *p_open) { igShowDebugLogWindow(p_open); }
bool GuiStorage_GetBool(ImGuiStorage *self, ImGuiID key, bool default_val) { return ImGuiStorage_GetBool(self, key, default_val); }
void GuiPlatformImeData_destroy(ImGuiPlatformImeData *self) { ImGuiPlatformImeData_destroy(self); }
ImGuiPlatformImeData *GuiPlatformImeData_ImGuiPlatformImeData() { return ImGuiPlatformImeData_ImGuiPlatformImeData(); }
void GuiTableColumn_destroy(ImGuiTableColumn *self) { ImGuiTableColumn_destroy(self); }
void Rect_destroy(ImRect *self) { ImRect_destroy(self); }
void EndPopup() { igEndPopup(); }
bool IsMouseReleased(ImGuiMouseButton button) { return igIsMouseReleased(button); }
ImGuiPlatformMonitor *GuiPlatformMonitor_ImGuiPlatformMonitor() { return ImGuiPlatformMonitor_ImGuiPlatformMonitor(); }
bool InvisibleButton(const char *str_id, const ImVec2 size, ImGuiButtonFlags flags) { return igInvisibleButton(str_id, size, flags); }
void DrawList_PrimWriteIdx(ImDrawList *self, ImDrawIdx idx) { ImDrawList_PrimWriteIdx(self, idx); }
ImGuiPayload *GuiPayload_ImGuiPayload() { return ImGuiPayload_ImGuiPayload(); }
float GetScrollX() { return igGetScrollX(); }
void DrawList_PathBezierQuadraticCurveTo(ImDrawList *self, const ImVec2 p2, const ImVec2 p3, int num_segments) { ImDrawList_PathBezierQuadraticCurveTo(self, p2, p3, num_segments); }
void Font_AddRemapChar(ImFont *self, ImWchar dst, ImWchar src, bool overwrite_dst) { ImFont_AddRemapChar(self, dst, src, overwrite_dst); }
float GetFontSize() { return igGetFontSize(); }
void EndCombo() { igEndCombo(); }
const ImWchar *FontAtlas_GetGlyphRangesKorean(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesKorean(self); }
void GuiTableTempData_destroy(ImGuiTableTempData *self) { ImGuiTableTempData_destroy(self); }
bool SliderInt2(const char *label, int v[2], int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igSliderInt2(label, v, v_min, v_max, format, flags); }
int TableGetRowIndex() { return igTableGetRowIndex(); }
void SetCursorPosX(float local_x) { igSetCursorPosX(local_x); }
void GuiIO_SetAppAcceptingEvents(ImGuiIO *self, bool accepting_events) { ImGuiIO_SetAppAcceptingEvents(self, accepting_events); }
void RenderPlatformWindowsDefault(void *platform_render_arg, void *renderer_render_arg) { igRenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg); }
void Font_ClearOutputData(ImFont *self) { ImFont_ClearOutputData(self); }
bool GuiTextFilter_Draw(ImGuiTextFilter *self, const char *label, float width) { return ImGuiTextFilter_Draw(self, label, width); }
bool BeginDragDropTarget() { return igBeginDragDropTarget(); }
void SetNextItemWidth(float item_width) { igSetNextItemWidth(item_width); }
float Font_GetCharAdvance(ImFont *self, ImWchar c) { return ImFont_GetCharAdvance(self, c); }
int TableGetColumnCount() { return igTableGetColumnCount(); }
void TextUnformatted(const char *text, const char *text_end) { igTextUnformatted(text, text_end); }
void GuiWindowSettings_destroy(ImGuiWindowSettings *self) { ImGuiWindowSettings_destroy(self); }
void PopItemWidth() { igPopItemWidth(); }
void Spacing() { igSpacing(); }
void GuiIO_AddKeyEvent(ImGuiIO *self, ImGuiKey key, bool down) { ImGuiIO_AddKeyEvent(self, key, down); }
void ProgressBar(float fraction, const ImVec2 size_arg, const char *overlay) { igProgressBar(fraction, size_arg, overlay); }
void PushButtonRepeat(bool repeat) { igPushButtonRepeat(repeat); }
void SetNextWindowViewport(ImGuiID viewport_id) { igSetNextWindowViewport(viewport_id); }
void SetNextWindowCollapsed(bool collapsed, ImGuiCond cond) { igSetNextWindowCollapsed(collapsed, cond); }
void TableNextRow(ImGuiTableRowFlags row_flags, float min_row_height) { igTableNextRow(row_flags, min_row_height); }
void DrawList_PathRect(ImDrawList *self, const ImVec2 rect_min, const ImVec2 rect_max, float rounding, ImDrawFlags flags) { ImDrawList_PathRect(self, rect_min, rect_max, rounding, flags); }
void DrawList_AddRectFilledMultiColor(ImDrawList *self, const ImVec2 p_min, const ImVec2 p_max, ImU32 col_upr_left, ImU32 col_upr_right, ImU32 col_bot_right, ImU32 col_bot_left) { ImDrawList_AddRectFilledMultiColor(self, p_min, p_max, col_upr_left, col_upr_right, col_bot_right, col_bot_left); }
void GetMousePos(ImVec2 *pOut) { igGetMousePos(pOut); }
float GetColumnWidth(int column_index) { return igGetColumnWidth(column_index); }
int GetKeyPressedAmount(ImGuiKey key, float repeat_delay, float rate) { return igGetKeyPressedAmount(key, repeat_delay, rate); }
void DrawList_PathLineTo(ImDrawList *self, const ImVec2 pos) { ImDrawList_PathLineTo(self, pos); }
void FontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder *self) { ImFontGlyphRangesBuilder_Clear(self); }
void SetNextWindowPos(const ImVec2 pos, ImGuiCond cond, const ImVec2 pivot) { igSetNextWindowPos(pos, cond, pivot); }
void DrawData_destroy(ImDrawData *self) { ImDrawData_destroy(self); }
void GuiIO_AddKeyAnalogEvent(ImGuiIO *self, ImGuiKey key, bool down, float v) { ImGuiIO_AddKeyAnalogEvent(self, key, down, v); }
bool DragInt3(const char *label, int v[3], float v_speed, int v_min, int v_max, const char *format, ImGuiSliderFlags flags) { return igDragInt3(label, v, v_speed, v_min, v_max, format, flags); }
void DrawListSplitter_Split(ImDrawListSplitter *self, ImDrawList *draw_list, int count) { ImDrawListSplitter_Split(self, draw_list, count); }
void GetItemRectMin(ImVec2 *pOut) { igGetItemRectMin(pOut); }
void GuiTextBuffer_destroy(ImGuiTextBuffer *self) { ImGuiTextBuffer_destroy(self); }
const ImGuiPayload *GetDragDropPayload() { return igGetDragDropPayload(); }
void GetWindowContentRegionMax(ImVec2 *pOut) { igGetWindowContentRegionMax(pOut); }
void PopClipRect() { igPopClipRect(); }
bool ShowStyleSelector(const char *label) { return igShowStyleSelector(label); }
void SetNextWindowClass(const ImGuiWindowClass *window_class) { igSetNextWindowClass(window_class); }
float GetFrameHeightWithSpacing() { return igGetFrameHeightWithSpacing(); }
bool IsItemActive() { return igIsItemActive(); }
void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) { igSetNextFrameWantCaptureKeyboard(want_capture_keyboard); }
ImDrawListSplitter *DrawListSplitter_ImDrawListSplitter() { return ImDrawListSplitter_ImDrawListSplitter(); }
void FontAtlas_GetTexDataAsAlpha8(ImFontAtlas *self, unsigned char **out_pixels, int *out_width, int *out_height, int *out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsAlpha8(self, out_pixels, out_width, out_height, out_bytes_per_pixel); }
void GuiInputTextCallbackData_InsertChars(ImGuiInputTextCallbackData *self, int pos, const char *text, const char *text_end) { ImGuiInputTextCallbackData_InsertChars(self, pos, text, text_end); }
int GetFrameCount() { return igGetFrameCount(); }
int GetKeyIndex(ImGuiKey key) { return igGetKeyIndex(key); }
const ImFontGlyph *Font_FindGlyphNoFallback(ImFont *self, ImWchar c) { return ImFont_FindGlyphNoFallback(self, c); }
void Font_BuildLookupTable(ImFont *self) { ImFont_BuildLookupTable(self); }
ImGuiTextFilter *GuiTextFilter_ImGuiTextFilter(const char *default_filter) { return ImGuiTextFilter_ImGuiTextFilter(default_filter); }
void ShowUserGuide() { igShowUserGuide(); }
ImGuiInputTextCallbackData *GuiInputTextCallbackData_ImGuiInputTextCallbackData() { return ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(); }
void GuiOldColumns_destroy(ImGuiOldColumns *self) { ImGuiOldColumns_destroy(self); }
bool InputInt4(const char *label, int v[4], ImGuiInputTextFlags flags) { return igInputInt4(label, v, flags); }
void Separator() { igSeparator(); }
void NewLine() { igNewLine(); }
void SameLine(float offset_from_start_x, float spacing) { igSameLine(offset_from_start_x, spacing); }
void DrawList_GetClipRectMin(ImVec2 *pOut, ImDrawList *self) { ImDrawList_GetClipRectMin(pOut, self); }
void DrawList_PushClipRect(ImDrawList *self, const ImVec2 clip_rect_min, const ImVec2 clip_rect_max, bool intersect_with_current_clip_rect) { ImDrawList_PushClipRect(self, clip_rect_min, clip_rect_max, intersect_with_current_clip_rect); }
void SetScrollHereX(float center_x_ratio) { igSetScrollHereX(center_x_ratio); }
ImTextureID DrawCmd_GetTexID(ImDrawCmd *self) { return ImDrawCmd_GetTexID(self); }
void DrawList_PrimRectUV(ImDrawList *self, const ImVec2 a, const ImVec2 b, const ImVec2 uv_a, const ImVec2 uv_b, ImU32 col) { ImDrawList_PrimRectUV(self, a, b, uv_a, uv_b, col); }
void GuiInputTextState_destroy(ImGuiInputTextState *self) { ImGuiInputTextState_destroy(self); }
bool SliderAngle(const char *label, float *v_rad, float v_degrees_min, float v_degrees_max, const char *format, ImGuiSliderFlags flags) { return igSliderAngle(label, v_rad, v_degrees_min, v_degrees_max, format, flags); }
ImGuiViewport *FindViewportByID(ImGuiID id) { return igFindViewportByID(id); }
ImDrawCmd *DrawCmd_ImDrawCmd() { return ImDrawCmd_ImDrawCmd(); }
void GuiPayload_Clear(ImGuiPayload *self) { ImGuiPayload_Clear(self); }
void Unindent(float indent_w) { igUnindent(indent_w); }
const ImWchar *FontAtlas_GetGlyphRangesCyrillic(ImFontAtlas *self) { return ImFontAtlas_GetGlyphRangesCyrillic(self); }
void SetColumnOffset(int column_index, float offset_x) { igSetColumnOffset(column_index, offset_x); }
bool IsItemFocused() { return igIsItemFocused(); }
bool VSliderScalar(const char *label, const ImVec2 size, ImGuiDataType data_type, void *p_data, const void *p_min, const void *p_max, const char *format, ImGuiSliderFlags flags) { return igVSliderScalar(label, size, data_type, p_data, p_min, p_max, format, flags); }
ImFontConfig *FontConf_ImFontConfig() { return ImFontConfig_ImFontConfig(); }
void GuiIO_AddMouseButtonEvent(ImGuiIO *self, int button, bool down) { ImGuiIO_AddMouseButtonEvent(self, button, down); }
void GuiTabItem_destroy(ImGuiTabItem *self) { ImGuiTabItem_destroy(self); }
ImDrawListSharedData *GetDrawListSharedData() { return igGetDrawListSharedData(); }
void GuiInputTextCallbackData_destroy(ImGuiInputTextCallbackData *self) { ImGuiInputTextCallbackData_destroy(self); }
int GuiStorage_GetInt(ImGuiStorage *self, ImGuiID key, int default_val) { return ImGuiStorage_GetInt(self, key, default_val); }
void GuiTableColumnSortSpecs_destroy(ImGuiTableColumnSortSpecs *self) { ImGuiTableColumnSortSpecs_destroy(self); }
void SetMouseCursor(ImGuiMouseCursor cursor_type) { igSetMouseCursor(cursor_type); }
void DrawList_PrimQuadUV(ImDrawList *self, const ImVec2 a, const ImVec2 b, const ImVec2 c, const ImVec2 d, const ImVec2 uv_a, const ImVec2 uv_b, const ImVec2 uv_c, const ImVec2 uv_d, ImU32 col) { ImDrawList_PrimQuadUV(self, a, b, c, d, uv_a, uv_b, uv_c, uv_d, col); }
ImGuiOnceUponAFrame *GuiOnceUponAFrame_ImGuiOnceUponAFrame() { return ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(); }
void GetContentRegionMax(ImVec2 *pOut) { igGetContentRegionMax(pOut); }
void SetNextWindowFocus() { igSetNextWindowFocus(); }
