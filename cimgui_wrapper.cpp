#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

void SetWindowFontScale(float scale) { igSetWindowFontScale(scale); }
void DrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left) { ImDrawList_AddRectFilledMultiColor(self,p_min,p_max,col_upr_left,col_upr_right,col_bot_right,col_bot_left); }
void CalcTextSizeV(ImVec2 *pOut,const char* text,bool hide_text_after_double_hash,float wrap_width) { igCalcTextSize(pOut,text,0,hide_text_after_double_hash,wrap_width); }
bool SliderAngleV(const char* label,float* v_rad,float v_degrees_min,float v_degrees_max,const char* format,ImGuiSliderFlags flags) { return igSliderAngle(label,v_rad,v_degrees_min,v_degrees_max,format,flags); }
bool ColorPicker3V(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorPicker3(label,col,flags); }
const ImWchar* FontAtlas_GetGlyphRangesDefault(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesDefault(self); }
void TableColumnSortSpecs_Destroy(ImGuiTableColumnSortSpecs* self) { ImGuiTableColumnSortSpecs_destroy(self); }
void GetDrawCursorStartPos(ImVec2 *pOut) { igGetCursorStartPos(pOut); }
float GetFrameHeightWithSpacing() { return igGetFrameHeightWithSpacing(); }
void BulletText(const char* fmt) { igBulletText(fmt); }
bool InputTextCallbackData_HasSelection(ImGuiInputTextCallbackData* self) { return ImGuiInputTextCallbackData_HasSelection(self); }
bool ShowStyleSelector(const char* label) { return igShowStyleSelector(label); }
void DrawList_PathArcToFast(ImDrawList* self,const ImVec2 center,float radius,int a_min_of_12,int a_max_of_12) { ImDrawList_PathArcToFast(self,center,radius,a_min_of_12,a_max_of_12); }
void Font_BuildLookupTable(ImFont* self) { ImFont_BuildLookupTable(self); }
ImGuiID GetWindowDockID() { return igGetWindowDockID(); }
bool DragIntV(const char* label,int* v,float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt(label,v,v_speed,v_min,v_max,format,flags); }
bool SliderInt3V(const char* label,int v[3],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt3(label,v,v_min,v_max,format,flags); }
void TextBuffer_Appendf(ImGuiTextBuffer* self,const char* fmt) { ImGuiTextBuffer_appendf(self,fmt); }
void WindowClass_Destroy(ImGuiWindowClass* self) { ImGuiWindowClass_destroy(self); }
bool SliderFloat4V(const char* label,float v[4],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat4(label,v,v_min,v_max,format,flags); }
void InputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_ClearSelection(self); }
bool IsWindowCollapsed() { return igIsWindowCollapsed(); }
ImDrawData* DrawData_ImDrawData() { return ImDrawData_ImDrawData(); }
ImGuiPlatformIO* GetPlatformIO() { return igGetPlatformIO(); }
ImFont* FontAtlas_AddFontFromMemoryTTFV(ImFontAtlas* self,void* font_data,int font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryTTF(self,font_data,font_size,size_pixels,font_cfg,glyph_ranges); }
bool IsKeyPressedV(ImGuiKey key,bool repeat) { return igIsKeyPressed(key,repeat); }
void DrawList_AddImageV(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col) { ImDrawList_AddImage(self,user_texture_id,p_min,p_max,uv_min,uv_max,col); }
void ContextHook_Destroy(ImGuiContextHook* self) { ImGuiContextHook_destroy(self); }
bool Payload_IsPreview(ImGuiPayload* self) { return ImGuiPayload_IsPreview(self); }
void EndListBox() { igEndListBox(); }
void EndMainMenuBar() { igEndMainMenuBar(); }
void DrawList_PathFillConvex(ImDrawList* self,ImU32 col) { ImDrawList_PathFillConvex(self,col); }
const ImVec4* GetStyleColorVec4(ImGuiCol idx) { return igGetStyleColorVec4(idx); }
bool IsPopupOpen_StrV(const char* str_id,ImGuiPopupFlags flags) { return igIsPopupOpen_Str(str_id,flags); }
void SetMouseCursor(ImGuiMouseCursor cursor_type) { igSetMouseCursor(cursor_type); }
void SetScrollFromPosY_FloatV(float local_y,float center_y_ratio) { igSetScrollFromPosY_Float(local_y,center_y_ratio); }
void PopID() { igPopID(); }
void PushStyleColor_U32(ImGuiCol idx,ImU32 col) { igPushStyleColor_U32(idx,col); }
void PushStyleColor_Vec4(ImGuiCol idx,const ImVec4 col) { igPushStyleColor_Vec4(idx,col); }
void Vec4_Destroy(ImVec4* self) { ImVec4_destroy(self); }
void DrawList_PopClipRect(ImDrawList* self) { ImDrawList_PopClipRect(self); }
void FontAtlas_CalcCustomRectUV(ImFontAtlas* self,const ImFontAtlasCustomRect* rect,ImVec2* out_uv_min,ImVec2* out_uv_max) { ImFontAtlas_CalcCustomRectUV(self,rect,out_uv_min,out_uv_max); }
void PopClipRect() { igPopClipRect(); }
void RenderPlatformWindowsDefaultV(void* platform_render_arg,void* renderer_render_arg) { igRenderPlatformWindowsDefault(platform_render_arg,renderer_render_arg); }
bool IsMousePosValidV(const ImVec2* mouse_pos) { return igIsMousePosValid(mouse_pos); }
void DrawData_Destroy(ImDrawData* self) { ImDrawData_destroy(self); }
void TextBuffer_AppendV(ImGuiTextBuffer* self,const char* str,const char* str_end) { ImGuiTextBuffer_append(self,str,str_end); }
ImGuiContext* CreateContextV(ImFontAtlas* shared_font_atlas) { return igCreateContext(shared_font_atlas); }
void LogToClipboardV(int auto_open_depth) { igLogToClipboard(auto_open_depth); }
bool RadioButton_Bool(const char* label,bool active) { return igRadioButton_Bool(label,active); }
bool RadioButton_IntPtr(const char* label,int* v,int v_button) { return igRadioButton_IntPtr(label,v,v_button); }
void TextBuffer_Clear(ImGuiTextBuffer* self) { ImGuiTextBuffer_clear(self); }
void DebugTextEncoding(const char* text) { igDebugTextEncoding(text); }
void DrawList_PathBezierQuadraticCurveToV(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,int num_segments) { ImDrawList_PathBezierQuadraticCurveTo(self,p2,p3,num_segments); }
bool BeginTabBarV(const char* str_id,ImGuiTabBarFlags flags) { return igBeginTabBar(str_id,flags); }
bool SliderScalarNV(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalarN(label,data_type,p_data,components,p_min,p_max,format,flags); }
void TableTempData_Destroy(ImGuiTableTempData* self) { ImGuiTableTempData_destroy(self); }
bool IsItemFocused() { return igIsItemFocused(); }
int GetMouseClickedCount(ImGuiMouseButton button) { return igGetMouseClickedCount(button); }
bool IsAnyItemActive() { return igIsAnyItemActive(); }
void SetColumnOffset(int column_index,float offset_x) { igSetColumnOffset(column_index,offset_x); }
bool BeginChildFrameV(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags) { return igBeginChildFrame(id,size,flags); }
void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) { igSetNextFrameWantCaptureKeyboard(want_capture_keyboard); }
void StyleColorsDarkV(ImGuiStyle* dst) { igStyleColorsDark(dst); }
void FontAtlasCustomRect_Destroy(ImFontAtlasCustomRect* self) { ImFontAtlasCustomRect_destroy(self); }
void GetMousePos(ImVec2 *pOut) { igGetMousePos(pOut); }
ImGuiTableColumnFlags TableGetColumnFlagsV(int column_n) { return igTableGetColumnFlags(column_n); }
void EndDragDropSource() { igEndDragDropSource(); }
void ShowStyleEditorV(ImGuiStyle* ref) { igShowStyleEditor(ref); }
int FontAtlas_AddCustomRectRegular(ImFontAtlas* self,int width,int height) { return ImFontAtlas_AddCustomRectRegular(self,width,height); }
ImGuiWindowClass* WindowClass_ImGuiWindowClass() { return ImGuiWindowClass_ImGuiWindowClass(); }
bool BeginPopupV(const char* str_id,ImGuiWindowFlags flags) { return igBeginPopup(str_id,flags); }
void Color_SetHSVV(ImColor* self,float h,float s,float v,float a) { ImColor_SetHSV(self,h,s,v,a); }
void DrawList_AddTriangleV(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness) { ImDrawList_AddTriangle(self,p1,p2,p3,col,thickness); }
int TableGetRowIndex() { return igTableGetRowIndex(); }
bool IsAnyMouseDown() { return igIsAnyMouseDown(); }
void ColorConvertRGBtoHSV(float r,float g,float b,float* out_h,float* out_s,float* out_v) { igColorConvertRGBtoHSV(r,g,b,out_h,out_s,out_v); }
ImGuiID DockSpaceV(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpace(id,size,flags,window_class); }
void* MemAlloc(size_t size) { return igMemAlloc(size); }
bool TreeNode_Str(const char* label) { return igTreeNode_Str(label); }
bool TreeNode_StrStr(const char* str_id,const char* fmt) { return igTreeNode_StrStr(str_id,fmt); }
bool TreeNode_Ptr(const void* ptr_id,const char* fmt) { return igTreeNode_Ptr(ptr_id,fmt); }
void ListClipper_ForceDisplayRangeByIndices(ImGuiListClipper* self,int item_min,int item_max) { ImGuiListClipper_ForceDisplayRangeByIndices(self,item_min,item_max); }
void PushClipRect(const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { igPushClipRect(clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void SetNextWindowCollapsedV(bool collapsed,ImGuiCond cond) { igSetNextWindowCollapsed(collapsed,cond); }
bool FontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect* self) { return ImFontAtlasCustomRect_IsPacked(self); }
const ImWchar* FontAtlas_GetGlyphRangesThai(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesThai(self); }
void LogButtons() { igLogButtons(); }
void SetNextWindowSizeConstraintsV(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data) { igSetNextWindowSizeConstraints(size_min,size_max,custom_callback,custom_callback_data); }
bool IsMouseDraggingV(ImGuiMouseButton button,float lock_threshold) { return igIsMouseDragging(button,lock_threshold); }
const ImGuiPayload* AcceptDragDropPayloadV(const char* type,ImGuiDragDropFlags flags) { return igAcceptDragDropPayload(type,flags); }
void StyleColorsClassicV(ImGuiStyle* dst) { igStyleColorsClassic(dst); }
void DrawList_PathRectV(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max,float rounding,ImDrawFlags flags) { ImDrawList_PathRect(self,rect_min,rect_max,rounding,flags); }
int TextBuffer_Size(ImGuiTextBuffer* self) { return ImGuiTextBuffer_size(self); }
void PushFont(ImFont* font) { igPushFont(font); }
void DrawList_AddQuadFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col) { ImDrawList_AddQuadFilled(self,p1,p2,p3,p4,col); }
ImGuiViewport* Viewport_ImGuiViewport() { return ImGuiViewport_ImGuiViewport(); }
void SetNextWindowSizeV(const ImVec2 size,ImGuiCond cond) { igSetNextWindowSize(size,cond); }
void InputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self,int pos,int bytes_count) { ImGuiInputTextCallbackData_DeleteChars(self,pos,bytes_count); }
float GetScrollX() { return igGetScrollX(); }
bool InputFloat4V(const char* label,float v[4],const char* format,ImGuiInputTextFlags flags) { return igInputFloat4(label,v,format,flags); }
const ImFontGlyph* Font_FindGlyph(ImFont* self,ImWchar c) { return ImFont_FindGlyph(self,c); }
void GetWindowContentRegionMax(ImVec2 *pOut) { igGetWindowContentRegionMax(pOut); }
void DrawData_Clear(ImDrawData* self) { ImDrawData_Clear(self); }
void ListClipper_End(ImGuiListClipper* self) { ImGuiListClipper_End(self); }
void NextItemData_Destroy(ImGuiNextItemData* self) { ImGuiNextItemData_destroy(self); }
void SetNextWindowContentSize(const ImVec2 size) { igSetNextWindowContentSize(size); }
bool InputFloat3V(const char* label,float v[3],const char* format,ImGuiInputTextFlags flags) { return igInputFloat3(label,v,format,flags); }
void CloseCurrentPopup() { igCloseCurrentPopup(); }
bool DragScalarNV(const char* label,ImGuiDataType data_type,void* p_data,int components,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalarN(label,data_type,p_data,components,v_speed,p_min,p_max,format,flags); }
void DrawList_PathClear(ImDrawList* self) { ImDrawList_PathClear(self); }
void StackLevelInfo_Destroy(ImGuiStackLevelInfo* self) { ImGuiStackLevelInfo_destroy(self); }
void EndChildFrame() { igEndChildFrame(); }
ImGuiViewport* FindViewportByPlatformHandle(void* platform_handle) { return igFindViewportByPlatformHandle(platform_handle); }
void FontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self,ImWchar c) { ImFontGlyphRangesBuilder_AddChar(self,c); }
void TextFilter_Clear(ImGuiTextFilter* self) { ImGuiTextFilter_Clear(self); }
bool BeginPopupContextWindowV(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextWindow(str_id,popup_flags); }
void SetClipboardText(const char* text) { igSetClipboardText(text); }
void SetColorEditOptions(ImGuiColorEditFlags flags) { igSetColorEditOptions(flags); }
void TabItem_Destroy(ImGuiTabItem* self) { ImGuiTabItem_destroy(self); }
const char* TextBuffer_End(ImGuiTextBuffer* self) { return ImGuiTextBuffer_end(self); }
float GetScrollY() { return igGetScrollY(); }
bool IsWindowFocusedV(ImGuiFocusedFlags flags) { return igIsWindowFocused(flags); }
void DrawList_PathLineTo(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineTo(self,pos); }
void TableHeader(const char* label) { igTableHeader(label); }
const ImWchar* FontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseFull(self); }
ImGuiID DockSpaceOverViewportV(const ImGuiViewport* viewport,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpaceOverViewport(viewport,flags,window_class); }
bool IsWindowDocked() { return igIsWindowDocked(); }
void PlatformIO_Destroy(ImGuiPlatformIO* self) { ImGuiPlatformIO_destroy(self); }
void Vec2_Destroy(ImVec2* self) { ImVec2_destroy(self); }
void FontConfig_Destroy(ImFontConfig* self) { ImFontConfig_destroy(self); }
float GetScrollMaxY() { return igGetScrollMaxY(); }
void SetNextWindowFocus() { igSetNextWindowFocus(); }
bool CheckboxFlags_IntPtr(const char* label,int* flags,int flags_value) { return igCheckboxFlags_IntPtr(label,flags,flags_value); }
bool CheckboxFlags_UintPtr(const char* label,unsigned int* flags,unsigned int flags_value) { return igCheckboxFlags_UintPtr(label,flags,flags_value); }
const ImFontGlyph* Font_FindGlyphNoFallback(ImFont* self,ImWchar c) { return ImFont_FindGlyphNoFallback(self,c); }
void FontAtlas_Clear(ImFontAtlas* self) { ImFontAtlas_Clear(self); }
const char* GetClipboardText() { return igGetClipboardText(); }
void TextUnformattedV(const char* text) { igTextUnformatted(text,0); }
bool DebugCheckVersionAndDataLayout(const char* version_str,size_t sz_io,size_t sz_style,size_t sz_vec2,size_t sz_vec4,size_t sz_drawvert,size_t sz_drawidx) { return igDebugCheckVersionAndDataLayout(version_str,sz_io,sz_style,sz_vec2,sz_vec4,sz_drawvert,sz_drawidx); }
float GetTextLineHeightWithSpacing() { return igGetTextLineHeightWithSpacing(); }
void LoadIniSettingsFromDisk(const char* ini_filename) { igLoadIniSettingsFromDisk(ini_filename); }
bool SetDragDropPayloadV(const char* type,const void* data,size_t sz,ImGuiCond cond) { return igSetDragDropPayload(type,data,sz,cond); }
void DrawList_PrimVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimVtx(self,pos,uv,col); }
ImFont* FontAtlas_AddFontFromMemoryCompressedTTFV(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedTTF(self,compressed_font_data,compressed_font_size,size_pixels,font_cfg,glyph_ranges); }
void TreePop() { igTreePop(); }
void SetScrollX_Float(float scroll_x) { igSetScrollX_Float(scroll_x); }
void StyleColorsLightV(ImGuiStyle* dst) { igStyleColorsLight(dst); }
bool TreeNodeEx_StrV(const char* label,ImGuiTreeNodeFlags flags) { return igTreeNodeEx_Str(label,flags); }
bool TreeNodeEx_StrStr(const char* str_id,ImGuiTreeNodeFlags flags,const char* fmt) { return igTreeNodeEx_StrStr(str_id,flags,fmt); }
bool TreeNodeEx_Ptr(const void* ptr_id,ImGuiTreeNodeFlags flags,const char* fmt) { return igTreeNodeEx_Ptr(ptr_id,flags,fmt); }
void Text(const char* fmt) { igText(fmt); }
void GetMousePosOnOpeningCurrentPopup(ImVec2 *pOut) { igGetMousePosOnOpeningCurrentPopup(pOut); }
void EndPopup() { igEndPopup(); }
const ImGuiPayload* GetDragDropPayload() { return igGetDragDropPayload(); }
bool TabItemButtonV(const char* label,ImGuiTabItemFlags flags) { return igTabItemButton(label,flags); }
ImGuiTableSortSpecs* TableSortSpecs_ImGuiTableSortSpecs() { return ImGuiTableSortSpecs_ImGuiTableSortSpecs(); }
bool DragFloatRange2V(const char* label,float* v_current_min,float* v_current_max,float v_speed,float v_min,float v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragFloatRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
ImFontAtlasCustomRect* FontAtlasCustomRect_ImFontAtlasCustomRect() { return ImFontAtlasCustomRect_ImFontAtlasCustomRect(); }
ImGuiMouseCursor GetMouseCursor() { return igGetMouseCursor(); }
void NextColumn() { igNextColumn(); }
bool SmallButton(const char* label) { return igSmallButton(label); }
void DrawList_Destroy(ImDrawList* self) { ImDrawList_destroy(self); }
const ImWchar* FontAtlas_GetGlyphRangesKorean(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesKorean(self); }
void IO_AddMousePosEvent(ImGuiIO* self,float x,float y) { ImGuiIO_AddMousePosEvent(self,x,y); }
ImFontGlyphRangesBuilder* FontGlyphRangesBuilder_ImFontGlyphRangesBuilder() { return ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(); }
bool ColorEdit4V(const char* label,float col[4],ImGuiColorEditFlags flags) { return igColorEdit4(label,col,flags); }
bool IsItemClickedV(ImGuiMouseButton mouse_button) { return igIsItemClicked(mouse_button); }
void DrawListSplitter_Merge(ImDrawListSplitter* self,ImDrawList* draw_list) { ImDrawListSplitter_Merge(self,draw_list); }
void FontAtlas_ClearInputData(ImFontAtlas* self) { ImFontAtlas_ClearInputData(self); }
ImFont* GetFont() { return igGetFont(); }
float GetFrameHeight() { return igGetFrameHeight(); }
void Color_HSVV(ImColor *pOut,float h,float s,float v,float a) { ImColor_HSV(pOut,h,s,v,a); }
void PopStyleColorV(int count) { igPopStyleColor(count); }
void SameLineV(float offset_from_start_x,float spacing) { igSameLine(offset_from_start_x,spacing); }
void DrawList_PushClipRectV(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { ImDrawList_PushClipRect(self,clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void InputTextState_Destroy(ImGuiInputTextState* self) { ImGuiInputTextState_destroy(self); }
ImVec4* Vec4_ImVec4_Nil() { return ImVec4_ImVec4_Nil(); }
ImVec4* Vec4_ImVec4_Float(float _x,float _y,float _z,float _w) { return ImVec4_ImVec4_Float(_x,_y,_z,_w); }
void IO_AddFocusEvent(ImGuiIO* self,bool focused) { ImGuiIO_AddFocusEvent(self,focused); }
bool TextFilter_IsActive(ImGuiTextFilter* self) { return ImGuiTextFilter_IsActive(self); }
void EndDisabled() { igEndDisabled(); }
bool IsItemVisible() { return igIsItemVisible(); }
void TextBuffer_Destroy(ImGuiTextBuffer* self) { ImGuiTextBuffer_destroy(self); }
void MenuColumns_Destroy(ImGuiMenuColumns* self) { ImGuiMenuColumns_destroy(self); }
ImGuiViewport* GetMainViewport() { return igGetMainViewport(); }
ImGuiStyle* GetStyle() { return igGetStyle(); }
void DrawList_PushTextureID(ImDrawList* self,ImTextureID texture_id) { ImDrawList_PushTextureID(self,texture_id); }
bool IsKeyReleased(ImGuiKey key) { return igIsKeyReleased(key); }
bool DragFloat2V(const char* label,float v[2],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat2(label,v,v_speed,v_min,v_max,format,flags); }
bool DragFloatV(const char* label,float* v,float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat(label,v,v_speed,v_min,v_max,format,flags); }
void DrawList_AddDrawCmd(ImDrawList* self) { ImDrawList_AddDrawCmd(self); }
void GetWindowSize(ImVec2 *pOut) { igGetWindowSize(pOut); }
bool InvisibleButtonV(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags) { return igInvisibleButton(str_id,size,flags); }
void DrawList_ChannelsMerge(ImDrawList* self) { ImDrawList_ChannelsMerge(self); }
void Font_RenderChar(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,ImWchar c) { ImFont_RenderChar(self,draw_list,size,pos,col,c); }
void FontGlyphRangesBuilder_AddTextV(ImFontGlyphRangesBuilder* self,const char* text) { ImFontGlyphRangesBuilder_AddText(self,text,0); }
bool DragScalarV(const char* label,ImGuiDataType data_type,void* p_data,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalar(label,data_type,p_data,v_speed,p_min,p_max,format,flags); }
float GetScrollMaxX() { return igGetScrollMaxX(); }
ImGuiOnceUponAFrame* OnceUponAFrame_ImGuiOnceUponAFrame() { return ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(); }
void IO_AddInputCharacterUTF16(ImGuiIO* self,ImWchar16 c) { ImGuiIO_AddInputCharacterUTF16(self,c); }
void Font_ClearOutputData(ImFont* self) { ImFont_ClearOutputData(self); }
void IO_AddInputCharacter(ImGuiIO* self,unsigned int c) { ImGuiIO_AddInputCharacter(self,c); }
ImGuiTableColumnSortSpecs* TableColumnSortSpecs_ImGuiTableColumnSortSpecs() { return ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(); }
void IO_AddInputCharactersUTF8(ImGuiIO* self,const char* str) { ImGuiIO_AddInputCharactersUTF8(self,str); }
bool ColorPicker4V(const char* label,float col[4],ImGuiColorEditFlags flags,const float* ref_col) { return igColorPicker4(label,col,flags,ref_col); }
void DrawList_PrimUnreserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimUnreserve(self,idx_count,vtx_count); }
bool InputTextWithHintV(const char* label,const char* hint,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextWithHint(label,hint,buf,buf_size,flags,callback,user_data); }
void SetNextItemWidth(float item_width) { igSetNextItemWidth(item_width); }
bool BeginMenuBar() { return igBeginMenuBar(); }
bool InputFloat2V(const char* label,float v[2],const char* format,ImGuiInputTextFlags flags) { return igInputFloat2(label,v,format,flags); }
void EndChild() { igEndChild(); }
void IO_SetAppAcceptingEvents(ImGuiIO* self,bool accepting_events) { ImGuiIO_SetAppAcceptingEvents(self,accepting_events); }
void SetDrawCursorPosX(float local_x) { igSetCursorPosX(local_x); }
void DrawCmd_Destroy(ImDrawCmd* self) { ImDrawCmd_destroy(self); }
void DrawList_AddNgonFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddNgonFilled(self,center,radius,col,num_segments); }
bool IsMouseClickedV(ImGuiMouseButton button,bool repeat) { return igIsMouseClicked(button,repeat); }
void LogText(const char* fmt) { igLogText(fmt); }
void PlotLines_FloatPtrV(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride) { igPlotLines_FloatPtr(label,values,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size,stride); }
void PlotLines_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size) { igPlotLines_FnFloatPtr(label,values_getter,data,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size); }
void SetKeyboardFocusHereV(int offset) { igSetKeyboardFocusHere(offset); }
void Font_AddGlyph(ImFont* self,const ImFontConfig* src_cfg,ImWchar c,float x0,float y0,float x1,float y1,float u0,float v0,float u1,float v1,float advance_x) { ImFont_AddGlyph(self,src_cfg,c,x0,y0,x1,y1,u0,v0,u1,v1,advance_x); }
void SetWindowFocus_Nil() { igSetWindowFocus_Nil(); }
void SetWindowFocus_Str(const char* name) { igSetWindowFocus_Str(name); }
bool TextFilter_DrawV(ImGuiTextFilter* self,const char* label,float width) { return ImGuiTextFilter_Draw(self,label,width); }
bool BeginDragDropTarget() { return igBeginDragDropTarget(); }
void EndCombo() { igEndCombo(); }
ImTextureID DrawCmd_GetTexID(ImDrawCmd* self) { return ImDrawCmd_GetTexID(self); }
void DrawList_PushClipRectFullScreen(ImDrawList* self) { ImDrawList_PushClipRectFullScreen(self); }
ImGuiStyle* Style_ImGuiStyle() { return ImGuiStyle_ImGuiStyle(); }
void TableColumnSettings_Destroy(ImGuiTableColumnSettings* self) { ImGuiTableColumnSettings_destroy(self); }
void TableInstanceData_Destroy(ImGuiTableInstanceData* self) { ImGuiTableInstanceData_destroy(self); }
int TableGetColumnIndex() { return igTableGetColumnIndex(); }
ImFontAtlasCustomRect* FontAtlas_GetCustomRectByIndex(ImFontAtlas* self,int index) { return ImFontAtlas_GetCustomRectByIndex(self,index); }
void DestroyPlatformWindows() { igDestroyPlatformWindows(); }
bool IsItemToggledOpen() { return igIsItemToggledOpen(); }
bool CollapsingHeader_TreeNodeFlagsV(const char* label,ImGuiTreeNodeFlags flags) { return igCollapsingHeader_TreeNodeFlags(label,flags); }
bool CollapsingHeader_BoolPtrV(const char* label,bool* p_visible,ImGuiTreeNodeFlags flags) { return igCollapsingHeader_BoolPtr(label,p_visible,flags); }
bool IsItemEdited() { return igIsItemEdited(); }
void DrawListSplitter_Clear(ImDrawListSplitter* self) { ImDrawListSplitter_Clear(self); }
bool Payload_IsDelivery(ImGuiPayload* self) { return ImGuiPayload_IsDelivery(self); }
void FontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self) { ImFontGlyphRangesBuilder_Clear(self); }
void GetItemRectSize(ImVec2 *pOut) { igGetItemRectSize(pOut); }
void PushTextWrapPosV(float wrap_local_pos_x) { igPushTextWrapPos(wrap_local_pos_x); }
ImDrawList* GetBackgroundDrawList_Nil() { return igGetBackgroundDrawList_Nil(); }
ImDrawList* GetBackgroundDrawList_ViewportPtr(ImGuiViewport* viewport) { return igGetBackgroundDrawList_ViewportPtr(viewport); }
bool BeginPopupContextItemV(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextItem(str_id,popup_flags); }
float GetDrawCursorPosX() { return igGetCursorPosX(); }
bool FontAtlas_IsBuilt(ImFontAtlas* self) { return ImFontAtlas_IsBuilt(self); }
bool Font_IsGlyphRangeUnused(ImFont* self,unsigned int c_begin,unsigned int c_last) { return ImFont_IsGlyphRangeUnused(self,c_begin,c_last); }
void ShowFontSelector(const char* label) { igShowFontSelector(label); }
void EndTabBar() { igEndTabBar(); }
bool InputScalarNV(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalarN(label,data_type,p_data,components,p_step,p_step_fast,format,flags); }
void TextFilter_Destroy(ImGuiTextFilter* self) { ImGuiTextFilter_destroy(self); }
void ShowStackToolWindowV(bool* p_open) { igShowStackToolWindow(p_open); }
void ColumnsV(int count,const char* id,bool border) { igColumns(count,id,border); }
ImU32 GetColorU32_ColV(ImGuiCol idx,float alpha_mul) { return igGetColorU32_Col(idx,alpha_mul); }
ImU32 GetColorU32_Vec4(const ImVec4 col) { return igGetColorU32_Vec4(col); }
ImU32 GetColorU32_U32(ImU32 col) { return igGetColorU32_U32(col); }
ImGuiTextFilter* TextFilter_ImGuiTextFilter(const char* default_filter) { return ImGuiTextFilter_ImGuiTextFilter(default_filter); }
ImDrawData* GetDrawData() { return igGetDrawData(); }
void GetDrawCursorPos(ImVec2 *pOut) { igGetCursorPos(pOut); }
void PushItemWidth(float item_width) { igPushItemWidth(item_width); }
bool Selectable_BoolV(const char* label,bool selected,ImGuiSelectableFlags flags,const ImVec2 size) { return igSelectable_Bool(label,selected,flags,size); }
bool Selectable_BoolPtrV(const char* label,bool* p_selected,ImGuiSelectableFlags flags,const ImVec2 size) { return igSelectable_BoolPtr(label,p_selected,flags,size); }
void SetWindowCollapsed_BoolV(bool collapsed,ImGuiCond cond) { igSetWindowCollapsed_Bool(collapsed,cond); }
void SetWindowCollapsed_StrV(const char* name,bool collapsed,ImGuiCond cond) { igSetWindowCollapsed_Str(name,collapsed,cond); }
void ShowMetricsWindowV(bool* p_open) { igShowMetricsWindow(p_open); }
void PlatformMonitor_Destroy(ImGuiPlatformMonitor* self) { ImGuiPlatformMonitor_destroy(self); }
void DrawList_AddLineV(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness) { ImDrawList_AddLine(self,p1,p2,col,thickness); }
float Font_GetCharAdvance(ImFont* self,ImWchar c) { return ImFont_GetCharAdvance(self,c); }
bool SliderInt4V(const char* label,int v[4],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt4(label,v,v_min,v_max,format,flags); }
void LogFinish() { igLogFinish(); }
void SetNextWindowDockIDV(ImGuiID dock_id,ImGuiCond cond) { igSetNextWindowDockID(dock_id,cond); }
bool SliderFloatV(const char* label,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat(label,v,v_min,v_max,format,flags); }
void MemFree(void* ptr) { igMemFree(ptr); }
void ImageV(ImTextureID user_texture_id,const ImVec2 size,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 tint_col,const ImVec4 border_col) { igImage(user_texture_id,size,uv0,uv1,tint_col,border_col); }
bool BeginPopupContextVoidV(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextVoid(str_id,popup_flags); }
void Font_SetGlyphVisible(ImFont* self,ImWchar c,bool visible) { ImFont_SetGlyphVisible(self,c,visible); }
void PlatformImeData_Destroy(ImGuiPlatformImeData* self) { ImGuiPlatformImeData_destroy(self); }
void GetDrawCursorScreenPos(ImVec2 *pOut) { igGetCursorScreenPos(pOut); }
void UnindentV(float indent_w) { igUnindent(indent_w); }
void FontAtlas_GetTexDataAsAlpha8V(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsAlpha8(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
void TextFilter_Build(ImGuiTextFilter* self) { ImGuiTextFilter_Build(self); }
void StackSizes_Destroy(ImGuiStackSizes* self) { ImGuiStackSizes_destroy(self); }
void ListClipperData_Destroy(ImGuiListClipperData* self) { ImGuiListClipperData_destroy(self); }
void OnceUponAFrame_Destroy(ImGuiOnceUponAFrame* self) { ImGuiOnceUponAFrame_destroy(self); }
void EndDragDropTarget() { igEndDragDropTarget(); }
void OldColumnData_Destroy(ImGuiOldColumnData* self) { ImGuiOldColumnData_destroy(self); }
void DrawList_AddRectFilledV(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddRectFilled(self,p_min,p_max,col,rounding,flags); }
void DrawList_PathArcToV(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments) { ImDrawList_PathArcTo(self,center,radius,a_min,a_max,num_segments); }
bool DragInt4V(const char* label,int v[4],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt4(label,v,v_speed,v_min,v_max,format,flags); }
const char* GetVersion() { return igGetVersion(); }
bool IsItemActivated() { return igIsItemActivated(); }
void DrawList_AddBezierQuadraticV(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierQuadratic(self,p1,p2,p3,col,thickness,num_segments); }
void DrawList_ChannelsSplit(ImDrawList* self,int count) { ImDrawList_ChannelsSplit(self,count); }
bool InputTextV(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputText(label,buf,buf_size,flags,callback,user_data); }
int GetKeyPressedAmount(ImGuiKey key,float repeat_delay,float rate) { return igGetKeyPressedAmount(key,repeat_delay,rate); }
void DrawList_GetClipRectMin(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMin(pOut,self); }
void WindowSettings_Destroy(ImGuiWindowSettings* self) { ImGuiWindowSettings_destroy(self); }
int GetColumnsCount() { return igGetColumnsCount(); }
ImGuiIO* GetIO() { return igGetIO(); }
ImColor* Color_ImColor_Nil() { return ImColor_ImColor_Nil(); }
ImColor* Color_ImColor_Float(float r,float g,float b,float a) { return ImColor_ImColor_Float(r,g,b,a); }
ImColor* Color_ImColor_Vec4(const ImVec4 col) { return ImColor_ImColor_Vec4(col); }
ImColor* Color_ImColor_Int(int r,int g,int b,int a) { return ImColor_ImColor_Int(r,g,b,a); }
ImColor* Color_ImColor_U32(ImU32 rgba) { return ImColor_ImColor_U32(rgba); }
void IO_ClearInputCharacters(ImGuiIO* self) { ImGuiIO_ClearInputCharacters(self); }
void PopButtonRepeat() { igPopButtonRepeat(); }
void SetCurrentContext(ImGuiContext* ctx) { igSetCurrentContext(ctx); }
void InputEvent_Destroy(ImGuiInputEvent* self) { ImGuiInputEvent_destroy(self); }
ImGuiInputTextCallbackData* InputTextCallbackData_ImGuiInputTextCallbackData() { return ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(); }
bool BeginPopupModalV(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBeginPopupModal(name,p_open,flags); }
bool Font_IsLoaded(ImFont* self) { return ImFont_IsLoaded(self); }
float GetWindowDpiScale() { return igGetWindowDpiScale(); }
void SetScrollHereYV(float center_y_ratio) { igSetScrollHereY(center_y_ratio); }
ImDrawList* GetForegroundDrawList_Nil() { return igGetForegroundDrawList_Nil(); }
ImDrawList* GetForegroundDrawList_ViewportPtr(ImGuiViewport* viewport) { return igGetForegroundDrawList_ViewportPtr(viewport); }
ImDrawList* DrawList_CloneOutput(ImDrawList* self) { return ImDrawList_CloneOutput(self); }
bool IsItemDeactivated() { return igIsItemDeactivated(); }
void SetScrollY_Float(float scroll_y) { igSetScrollY_Float(scroll_y); }
void IO_Destroy(ImGuiIO* self) { ImGuiIO_destroy(self); }
ImGuiTextBuffer* TextBuffer_ImGuiTextBuffer() { return ImGuiTextBuffer_ImGuiTextBuffer(); }
bool ImageButtonV(const char* str_id,ImTextureID user_texture_id,const ImVec2 size,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 bg_col,const ImVec4 tint_col) { return igImageButton(str_id,user_texture_id,size,uv0,uv1,bg_col,tint_col); }
void FontAtlas_SetTexID(ImFontAtlas* self,ImTextureID id) { ImFontAtlas_SetTexID(self,id); }
void Viewport_GetWorkCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetWorkCenter(pOut,self); }
bool Checkbox(const char* label,bool* v) { return igCheckbox(label,v); }
ImDrawListSharedData* GetDrawListSharedData() { return igGetDrawListSharedData(); }
void FontAtlas_GetTexDataAsRGBA32V(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsRGBA32(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
void ShowAboutWindowV(bool* p_open) { igShowAboutWindow(p_open); }
ImGuiTableSortSpecs* TableGetSortSpecs() { return igTableGetSortSpecs(); }
void IO_ClearInputKeys(ImGuiIO* self) { ImGuiIO_ClearInputKeys(self); }
void TableSortSpecs_Destroy(ImGuiTableSortSpecs* self) { ImGuiTableSortSpecs_destroy(self); }
void DrawList_PrimWriteVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimWriteVtx(self,pos,uv,col); }
void Vec1_Destroy(ImVec1* self) { ImVec1_destroy(self); }
ImGuiViewport* FindViewportByID(ImGuiID id) { return igFindViewportByID(id); }
float GetWindowWidth() { return igGetWindowWidth(); }
void TableSetupScrollFreeze(int cols,int rows) { igTableSetupScrollFreeze(cols,rows); }
void OldColumns_Destroy(ImGuiOldColumns* self) { ImGuiOldColumns_destroy(self); }
bool IsAnyItemFocused() { return igIsAnyItemFocused(); }
bool DragIntRange2V(const char* label,int* v_current_min,int* v_current_max,float v_speed,int v_min,int v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragIntRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
const char* GetStyleColorName(ImGuiCol idx) { return igGetStyleColorName(idx); }
bool ListBox_Str_arrV(const char* label,int* current_item,const char* const items[],int items_count,int height_in_items) { return igListBox_Str_arr(label,current_item,items,items_count,height_in_items); }
bool ListBox_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int height_in_items) { return igListBox_FnBoolPtr(label,current_item,items_getter,data,items_count,height_in_items); }
float GetColumnWidthV(int column_index) { return igGetColumnWidth(column_index); }
void StyleMod_Destroy(ImGuiStyleMod* self) { ImGuiStyleMod_destroy(self); }
void PopStyleVarV(int count) { igPopStyleVar(count); }
void EndFrame() { igEndFrame(); }
void LogToTTYV(int auto_open_depth) { igLogToTTY(auto_open_depth); }
void Spacing() { igSpacing(); }
void GetContentRegionMax(ImVec2 *pOut) { igGetContentRegionMax(pOut); }
float GetTreeNodeToLabelSpacing() { return igGetTreeNodeToLabelSpacing(); }
bool IsItemDeactivatedAfterEdit() { return igIsItemDeactivatedAfterEdit(); }
void PopFont() { igPopFont(); }
void SetNextFrameWantCaptureMouse(bool want_capture_mouse) { igSetNextFrameWantCaptureMouse(want_capture_mouse); }
float GetDrawCursorPosY() { return igGetCursorPosY(); }
void InputTextCallbackData_InsertCharsV(ImGuiInputTextCallbackData* self,int pos,const char* text) { ImGuiInputTextCallbackData_InsertChars(self,pos,text,0); }
void Rect_Destroy(ImRect* self) { ImRect_destroy(self); }
void DestroyContextV(ImGuiContext* ctx) { igDestroyContext(ctx); }
bool TableNextColumn() { return igTableNextColumn(); }
bool BeginV(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBegin(name,p_open,flags); }
bool InputIntV(const char* label,int* v,int step,int step_fast,ImGuiInputTextFlags flags) { return igInputInt(label,v,step,step_fast,flags); }
bool TextFilter_PassFilterV(ImGuiTextFilter* self,const char* text) { return ImGuiTextFilter_PassFilter(self,text,0); }
void Viewport_GetCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetCenter(pOut,self); }
float GetWindowHeight() { return igGetWindowHeight(); }
void ShowDebugLogWindowV(bool* p_open) { igShowDebugLogWindow(p_open); }
void ListClipper_Destroy(ImGuiListClipper* self) { ImGuiListClipper_destroy(self); }
void SetItemAllowOverlap() { igSetItemAllowOverlap(); }
void DrawList_AddText_Vec2V(ImDrawList* self,const ImVec2 pos,ImU32 col,const char* text_begin) { ImDrawList_AddText_Vec2(self,pos,col,text_begin,0); }
void DrawList_AddText_FontPtrV(ImDrawList* self,const ImFont* font,float font_size,const ImVec2 pos,ImU32 col,const char* text_begin,float wrap_width,const ImVec4* cpu_fine_clip_rect) { ImDrawList_AddText_FontPtr(self,font,font_size,pos,col,text_begin,0,wrap_width,cpu_fine_clip_rect); }
bool TextBuffer_Empty(ImGuiTextBuffer* self) { return ImGuiTextBuffer_empty(self); }
void LoadIniSettingsFromMemoryV(const char* ini_data,size_t ini_size) { igLoadIniSettingsFromMemory(ini_data,ini_size); }
void EndMenu() { igEndMenu(); }
void InputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_SelectAll(self); }
bool BeginTabItemV(const char* label,bool* p_open,ImGuiTabItemFlags flags) { return igBeginTabItem(label,p_open,flags); }
void TextWrapped(const char* fmt) { igTextWrapped(fmt); }
bool FontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder* self,size_t n) { return ImFontGlyphRangesBuilder_GetBit(self,n); }
const char* TextBuffer_Begin(ImGuiTextBuffer* self) { return ImGuiTextBuffer_begin(self); }
bool IsAnyItemHovered() { return igIsAnyItemHovered(); }
bool IsWindowHoveredV(ImGuiHoveredFlags flags) { return igIsWindowHovered(flags); }
void DrawList_AddCallback(ImDrawList* self,ImDrawCallback callback,void* callback_data) { ImDrawList_AddCallback(self,callback,callback_data); }
void SetColumnWidth(int column_index,float width) { igSetColumnWidth(column_index,width); }
void GetWindowContentRegionMin(ImVec2 *pOut) { igGetWindowContentRegionMin(pOut); }
void DrawList_PrimRectUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,ImU32 col) { ImDrawList_PrimRectUV(self,a,b,uv_a,uv_b,col); }
void TextColored(const ImVec4 col,const char* fmt) { igTextColored(col,fmt); }
void Color_Destroy(ImColor* self) { ImColor_destroy(self); }
void SetDrawCursorPosY(float local_y) { igSetCursorPosY(local_y); }
void DrawList_AddNgonV(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddNgon(self,center,radius,col,num_segments,thickness); }
bool BeginChild_StrV(const char* str_id,const ImVec2 size,bool border,ImGuiWindowFlags flags) { return igBeginChild_Str(str_id,size,border,flags); }
bool BeginChild_IDV(ImGuiID id,const ImVec2 size,bool border,ImGuiWindowFlags flags) { return igBeginChild_ID(id,size,border,flags); }
float GetFontSize() { return igGetFontSize(); }
void TableSetBgColorV(ImGuiTableBgTarget target,ImU32 color,int column_n) { igTableSetBgColor(target,color,column_n); }
ImFont* Font_ImFont() { return ImFont_ImFont(); }
void PopTextWrapPos() { igPopTextWrapPos(); }
void UpdatePlatformWindows() { igUpdatePlatformWindows(); }
bool ListClipper_Step(ImGuiListClipper* self) { return ImGuiListClipper_Step(self); }
void DrawList_GetClipRectMax(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMax(pOut,self); }
const char* TextBuffer_c_str(ImGuiTextBuffer* self) { return ImGuiTextBuffer_c_str(self); }
bool Combo_Str_arrV(const char* label,int* current_item,const char* const items[],int items_count,int popup_max_height_in_items) { return igCombo_Str_arr(label,current_item,items,items_count,popup_max_height_in_items); }
bool Combo_StrV(const char* label,int* current_item,const char* items_separated_by_zeros,int popup_max_height_in_items) { return igCombo_Str(label,current_item,items_separated_by_zeros,popup_max_height_in_items); }
bool Combo_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int popup_max_height_in_items) { return igCombo_FnBoolPtr(label,current_item,items_getter,data,items_count,popup_max_height_in_items); }
void SetNextItemOpenV(bool is_open,ImGuiCond cond) { igSetNextItemOpen(is_open,cond); }
bool SliderInt2V(const char* label,int v[2],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt2(label,v,v_min,v_max,format,flags); }
void NextWindowData_Destroy(ImGuiNextWindowData* self) { ImGuiNextWindowData_destroy(self); }
bool SliderIntV(const char* label,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt(label,v,v_min,v_max,format,flags); }
void IndentV(float indent_w) { igIndent(indent_w); }
void OpenPopup_StrV(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopup_Str(str_id,popup_flags); }
void OpenPopup_IDV(ImGuiID id,ImGuiPopupFlags popup_flags) { igOpenPopup_ID(id,popup_flags); }
void ResetMouseDragDeltaV(ImGuiMouseButton button) { igResetMouseDragDelta(button); }
void IO_AddMouseButtonEvent(ImGuiIO* self,int button,bool down) { ImGuiIO_AddMouseButtonEvent(self,button,down); }
void DrawList_AddImageRoundedV(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddImageRounded(self,user_texture_id,p_min,p_max,uv_min,uv_max,col,rounding,flags); }
void FontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self,const ImWchar* ranges) { ImFontGlyphRangesBuilder_AddRanges(self,ranges); }
void FontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self,size_t n) { ImFontGlyphRangesBuilder_SetBit(self,n); }
void SaveIniSettingsToDisk(const char* ini_filename) { igSaveIniSettingsToDisk(ini_filename); }
void DrawListSplitter_Split(ImDrawListSplitter* self,ImDrawList* draw_list,int count) { ImDrawListSplitter_Split(self,draw_list,count); }
const ImWchar* FontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesVietnamese(self); }
ImFont* FontAtlas_AddFontFromMemoryCompressedBase85TTFV(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self,compressed_font_data_base85,size_pixels,font_cfg,glyph_ranges); }
const char* Font_GetDebugName(ImFont* self) { return ImFont_GetDebugName(self); }
void SetDrawCursorScreenPos(const ImVec2 pos) { igSetCursorScreenPos(pos); }
void DrawData_DeIndexAllBuffers(ImDrawData* self) { ImDrawData_DeIndexAllBuffers(self); }
void DrawList_PathLineToMergeDuplicate(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineToMergeDuplicate(self,pos); }
void LogToFileV(int auto_open_depth,const char* filename) { igLogToFile(auto_open_depth,filename); }
void SetWindowPos_Vec2V(const ImVec2 pos,ImGuiCond cond) { igSetWindowPos_Vec2(pos,cond); }
void SetWindowPos_StrV(const char* name,const ImVec2 pos,ImGuiCond cond) { igSetWindowPos_Str(name,pos,cond); }
void GetContentRegionAvail(ImVec2 *pOut) { igGetContentRegionAvail(pOut); }
void ShowUserGuide() { igShowUserGuide(); }
bool SliderFloat3V(const char* label,float v[3],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat3(label,v,v_min,v_max,format,flags); }
void DrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self,ImDrawList* draw_list,int channel_idx) { ImDrawListSplitter_SetCurrentChannel(self,draw_list,channel_idx); }
void ColorConvertHSVtoRGB(float h,float s,float v,float* out_r,float* out_g,float* out_b) { igColorConvertHSVtoRGB(h,s,v,out_r,out_g,out_b); }
void PushID_Str(const char* str_id) { igPushID_Str(str_id); }
void PushID_StrStr(const char* str_id_begin,const char* str_id_end) { igPushID_StrStr(str_id_begin,str_id_end); }
void PushID_Ptr(const void* ptr_id) { igPushID_Ptr(ptr_id); }
void PushID_Int(int int_id) { igPushID_Int(int_id); }
void Bullet() { igBullet(); }
void SetNextWindowBgAlpha(float alpha) { igSetNextWindowBgAlpha(alpha); }
void SetNextWindowPosV(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot) { igSetNextWindowPos(pos,cond,pivot); }
void DrawList_PrimRect(ImDrawList* self,const ImVec2 a,const ImVec2 b,ImU32 col) { ImDrawList_PrimRect(self,a,b,col); }
bool BeginComboV(const char* label,const char* preview_value,ImGuiComboFlags flags) { return igBeginCombo(label,preview_value,flags); }
void FontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self,ImVector_ImWchar* out_ranges) { ImFontGlyphRangesBuilder_BuildRanges(self,out_ranges); }
void ComboPreviewData_Destroy(ImGuiComboPreviewData* self) { ImGuiComboPreviewData_destroy(self); }
void IO_AddMouseWheelEvent(ImGuiIO* self,float wh_x,float wh_y) { ImGuiIO_AddMouseWheelEvent(self,wh_x,wh_y); }
void DrawList_AddRectV(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness) { ImDrawList_AddRect(self,p_min,p_max,col,rounding,flags,thickness); }
bool SliderFloat2V(const char* label,float v[2],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat2(label,v,v_min,v_max,format,flags); }
void PushAllowKeyboardFocus(bool allow_keyboard_focus) { igPushAllowKeyboardFocus(allow_keyboard_focus); }
void DrawList_AddQuadV(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness) { ImDrawList_AddQuad(self,p1,p2,p3,p4,col,thickness); }
void Payload_Clear(ImGuiPayload* self) { ImGuiPayload_Clear(self); }
void End() { igEnd(); }
const char* GetKeyName(ImGuiKey key) { return igGetKeyName(key); }
void SetTabItemClosed(const char* tab_or_docked_window_label) { igSetTabItemClosed(tab_or_docked_window_label); }
void DrawList_PrimReserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimReserve(self,idx_count,vtx_count); }
ImGuiPlatformMonitor* PlatformMonitor_ImGuiPlatformMonitor() { return ImGuiPlatformMonitor_ImGuiPlatformMonitor(); }
void AlignTextToFramePadding() { igAlignTextToFramePadding(); }
void EndMenuBar() { igEndMenuBar(); }
int FontAtlas_AddCustomRectFontGlyphV(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x,const ImVec2 offset) { return ImFontAtlas_AddCustomRectFontGlyph(self,font,id,width,height,advance_x,offset); }
void MetricsConfig_Destroy(ImGuiMetricsConfig* self) { ImGuiMetricsConfig_destroy(self); }
void PtrOrIndex_Destroy(ImGuiPtrOrIndex* self) { ImGuiPtrOrIndex_destroy(self); }
int GetKeyIndex(ImGuiKey key) { return igGetKeyIndex(key); }
ImGuiListClipper* ListClipper_ImGuiListClipper() { return ImGuiListClipper_ImGuiListClipper(); }
void Payload_Destroy(ImGuiPayload* self) { ImGuiPayload_destroy(self); }
double GetTime() { return igGetTime(); }
bool IsKeyDown(ImGuiKey key) { return igIsKeyDown(key); }
void DockContext_Destroy(ImGuiDockContext* self) { ImGuiDockContext_destroy(self); }
void NavItemData_Destroy(ImGuiNavItemData* self) { ImGuiNavItemData_destroy(self); }
void SetWindowSize_Vec2V(const ImVec2 size,ImGuiCond cond) { igSetWindowSize_Vec2(size,cond); }
void SetWindowSize_StrV(const char* name,const ImVec2 size,ImGuiCond cond) { igSetWindowSize_Str(name,size,cond); }
void Font_AddRemapCharV(ImFont* self,ImWchar dst,ImWchar src,bool overwrite_dst) { ImFont_AddRemapChar(self,dst,src,overwrite_dst); }
ImGuiContext* GetCurrentContext() { return igGetCurrentContext(); }
void GetItemRectMin(ImVec2 *pOut) { igGetItemRectMin(pOut); }
void DrawList_AddCircleV(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddCircle(self,center,radius,col,num_segments,thickness); }
void Font_Destroy(ImFont* self) { ImFont_destroy(self); }
void LastItemData_Destroy(ImGuiLastItemData* self) { ImGuiLastItemData_destroy(self); }
ImGuiPayload* Payload_ImGuiPayload() { return ImGuiPayload_ImGuiPayload(); }
ImFont* FontAtlas_AddFont(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFont(self,font_cfg); }
bool DragInt3V(const char* label,int v[3],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt3(label,v,v_speed,v_min,v_max,format,flags); }
bool IsItemActive() { return igIsItemActive(); }
bool MenuItem_BoolV(const char* label,const char* shortcut,bool selected,bool enabled) { return igMenuItem_Bool(label,shortcut,selected,enabled); }
bool MenuItem_BoolPtrV(const char* label,const char* shortcut,bool* p_selected,bool enabled) { return igMenuItem_BoolPtr(label,shortcut,p_selected,enabled); }
float GetTextLineHeight() { return igGetTextLineHeight(); }
void SetNextWindowClass(const ImGuiWindowClass* window_class) { igSetNextWindowClass(window_class); }
void SetTooltip(const char* fmt) { igSetTooltip(fmt); }
bool DragFloat4V(const char* label,float v[4],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat4(label,v,v_speed,v_min,v_max,format,flags); }
void SetScrollFromPosX_FloatV(float local_x,float center_x_ratio) { igSetScrollFromPosX_Float(local_x,center_x_ratio); }
void ColorConvertU32ToFloat4(ImVec4 *pOut,ImU32 in) { igColorConvertU32ToFloat4(pOut,in); }
void EndTabItem() { igEndTabItem(); }
bool IsWindowAppearing() { return igIsWindowAppearing(); }
const ImWchar* FontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesJapanese(self); }
bool IsMouseHoveringRectV(const ImVec2 r_min,const ImVec2 r_max,bool clip) { return igIsMouseHoveringRect(r_min,r_max,clip); }
void PushStyleVar_Float(ImGuiStyleVar idx,float val) { igPushStyleVar_Float(idx,val); }
void PushStyleVar_Vec2(ImGuiStyleVar idx,const ImVec2 val) { igPushStyleVar_Vec2(idx,val); }
bool DragInt2V(const char* label,int v[2],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt2(label,v,v_speed,v_min,v_max,format,flags); }
void NewLine() { igNewLine(); }
void PopItemWidth() { igPopItemWidth(); }
ImFontAtlas* FontAtlas_ImFontAtlas() { return ImFontAtlas_ImFontAtlas(); }
void ProgressBarV(float fraction,const ImVec2 size_arg,const char* overlay) { igProgressBar(fraction,size_arg,overlay); }
void SetDrawCursorPos(const ImVec2 local_pos) { igSetCursorPos(local_pos); }
ImDrawCmd* DrawCmd_ImDrawCmd() { return ImDrawCmd_ImDrawCmd(); }
ImGuiIO* IO_ImGuiIO() { return ImGuiIO_ImGuiIO(); }
void GetMouseDragDeltaV(ImVec2 *pOut,ImGuiMouseButton button,float lock_threshold) { igGetMouseDragDelta(pOut,button,lock_threshold); }
void IO_AddKeyEvent(ImGuiIO* self,ImGuiKey key,bool down) { ImGuiIO_AddKeyEvent(self,key,down); }
void EndTable() { igEndTable(); }
void DrawData_ScaleClipRects(ImDrawData* self,const ImVec2 fb_scale) { ImDrawData_ScaleClipRects(self,fb_scale); }
ImDrawList* GetWindowDrawList() { return igGetWindowDrawList(); }
void DrawList_AddPolyline(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_AddPolyline(self,points,num_points,col,flags,thickness); }
void DrawList_PathBezierCubicCurveToV(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,int num_segments) { ImDrawList_PathBezierCubicCurveTo(self,p2,p3,p4,num_segments); }
void TableSettings_Destroy(ImGuiTableSettings* self) { ImGuiTableSettings_destroy(self); }
bool InputTextMultilineV(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextMultiline(label,buf,buf_size,size,flags,callback,user_data); }
void NewFrame() { igNewFrame(); }
void PopupData_Destroy(ImGuiPopupData* self) { ImGuiPopupData_destroy(self); }
bool InputInt2V(const char* label,int v[2],ImGuiInputTextFlags flags) { return igInputInt2(label,v,flags); }
bool TableSetColumnIndex(int column_n) { return igTableSetColumnIndex(column_n); }
void DrawList_PrimQuadUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 d,const ImVec2 uv_a,const ImVec2 uv_b,const ImVec2 uv_c,const ImVec2 uv_d,ImU32 col) { ImDrawList_PrimQuadUV(self,a,b,c,d,uv_a,uv_b,uv_c,uv_d,col); }
ImGuiPlatformIO* PlatformIO_ImGuiPlatformIO() { return ImGuiPlatformIO_ImGuiPlatformIO(); }
void GetWindowPos(ImVec2 *pOut) { igGetWindowPos(pOut); }
bool VSliderFloatV(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderFloat(label,size,v,v_min,v_max,format,flags); }
void BeginTooltip() { igBeginTooltip(); }
const char* SaveIniSettingsToMemoryV(size_t* out_ini_size) { return igSaveIniSettingsToMemory(out_ini_size); }
void DrawList_ChannelsSetCurrent(ImDrawList* self,int n) { ImDrawList_ChannelsSetCurrent(self,n); }
void Context_Destroy(ImGuiContext* self) { ImGuiContext_destroy(self); }
ImDrawListSplitter* DrawListSplitter_ImDrawListSplitter() { return ImDrawListSplitter_ImDrawListSplitter(); }
void TextDisabled(const char* fmt) { igTextDisabled(fmt); }
ImGuiViewport* GetWindowViewport() { return igGetWindowViewport(); }
ImFontConfig* FontConfig_ImFontConfig() { return ImFontConfig_ImFontConfig(); }
bool ButtonV(const char* label,const ImVec2 size) { return igButton(label,size); }
void PushButtonRepeat(bool repeat) { igPushButtonRepeat(repeat); }
void TableColumn_Destroy(ImGuiTableColumn* self) { ImGuiTableColumn_destroy(self); }
bool InputFloatV(const char* label,float* v,float step,float step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputFloat(label,v,step,step_fast,format,flags); }
void Font_CalcTextSizeAV(ImVec2 *pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin,const char** remaining) { ImFont_CalcTextSizeA(pOut,self,size,max_width,wrap_width,text_begin,0,remaining); }
void InputTextCallbackData_Destroy(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_destroy(self); }
void BeginGroup() { igBeginGroup(); }
void EndGroup() { igEndGroup(); }
bool InputInt3V(const char* label,int v[3],ImGuiInputTextFlags flags) { return igInputInt3(label,v,flags); }
void DrawList_PopTextureID(ImDrawList* self) { ImDrawList_PopTextureID(self); }
void IO_AddMouseViewportEvent(ImGuiIO* self,ImGuiID id) { ImGuiIO_AddMouseViewportEvent(self,id); }
void SetNextWindowViewport(ImGuiID viewport_id) { igSetNextWindowViewport(viewport_id); }
void DrawList_AddCircleFilledV(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddCircleFilled(self,center,radius,col,num_segments); }
void TableNextRowV(ImGuiTableRowFlags row_flags,float min_row_height) { igTableNextRow(row_flags,min_row_height); }
void DrawList_AddTriangleFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col) { ImDrawList_AddTriangleFilled(self,p1,p2,p3,col); }
bool FontAtlas_GetMouseCursorTexData(ImFontAtlas* self,ImGuiMouseCursor cursor,ImVec2* out_offset,ImVec2* out_size,ImVec2 out_uv_border[2],ImVec2 out_uv_fill[2]) { return ImFontAtlas_GetMouseCursorTexData(self,cursor,out_offset,out_size,out_uv_border,out_uv_fill); }
bool InputInt4V(const char* label,int v[4],ImGuiInputTextFlags flags) { return igInputInt4(label,v,flags); }
bool VSliderScalarV(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igVSliderScalar(label,size,data_type,p_data,p_min,p_max,format,flags); }
void Style_Destroy(ImGuiStyle* self) { ImGuiStyle_destroy(self); }
bool ColorButtonV(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size) { return igColorButton(desc_id,col,flags,size); }
ImVec2* Vec2_ImVec2_Nil() { return ImVec2_ImVec2_Nil(); }
ImVec2* Vec2_ImVec2_Float(float _x,float _y) { return ImVec2_ImVec2_Float(_x,_y); }
const char* Font_CalcWordWrapPositionA(ImFont* self,float scale,const char* text,float wrap_width) { return ImFont_CalcWordWrapPositionA(self,scale,text,0,wrap_width); }
void TableHeadersRow() { igTableHeadersRow(); }
void Separator() { igSeparator(); }
void Font_GrowIndex(ImFont* self,int new_size) { ImFont_GrowIndex(self,new_size); }
void DrawList_PathStrokeV(ImDrawList* self,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_PathStroke(self,col,flags,thickness); }
void FontAtlas_Destroy(ImFontAtlas* self) { ImFontAtlas_destroy(self); }
void PopAllowKeyboardFocus() { igPopAllowKeyboardFocus(); }
void DrawList_AddImageQuadV(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col) { ImDrawList_AddImageQuad(self,user_texture_id,p1,p2,p3,p4,uv1,uv2,uv3,uv4,col); }
ImGuiPlatformImeData* PlatformImeData_ImGuiPlatformImeData() { return ImGuiPlatformImeData_ImGuiPlatformImeData(); }
bool IsRectVisible_Nil(const ImVec2 size) { return igIsRectVisible_Nil(size); }
bool IsRectVisible_Vec2(const ImVec2 rect_min,const ImVec2 rect_max) { return igIsRectVisible_Vec2(rect_min,rect_max); }
void TreePush_Str(const char* str_id) { igTreePush_Str(str_id); }
void TreePush_PtrV(const void* ptr_id) { igTreePush_Ptr(ptr_id); }
void SettingsHandler_Destroy(ImGuiSettingsHandler* self) { ImGuiSettingsHandler_destroy(self); }
bool IsMouseDoubleClicked(ImGuiMouseButton button) { return igIsMouseDoubleClicked(button); }
void DrawList_AddConvexPolyFilled(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col) { ImDrawList_AddConvexPolyFilled(self,points,num_points,col); }
void DrawList_PrimWriteIdx(ImDrawList* self,ImDrawIdx idx) { ImDrawList_PrimWriteIdx(self,idx); }
bool IsMouseDown(ImGuiMouseButton button) { return igIsMouseDown(button); }
bool IsItemHoveredV(ImGuiHoveredFlags flags) { return igIsItemHovered(flags); }
void IO_AddKeyAnalogEvent(ImGuiIO* self,ImGuiKey key,bool down,float v) { ImGuiIO_AddKeyAnalogEvent(self,key,down,v); }
bool BeginListBoxV(const char* label,const ImVec2 size) { return igBeginListBox(label,size); }
float CalcItemWidth() { return igCalcItemWidth(); }
bool IsMouseReleased(ImGuiMouseButton button) { return igIsMouseReleased(button); }
int TableGetColumnCount() { return igTableGetColumnCount(); }
ImFont* FontAtlas_AddFontFromFileTTFV(ImFontAtlas* self,const char* filename,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromFileTTF(self,filename,size_pixels,font_cfg,glyph_ranges); }
void IO_SetKeyEventNativeDataV(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode,int native_legacy_index) { ImGuiIO_SetKeyEventNativeData(self,key,native_keycode,native_scancode,native_legacy_index); }
void TableSetColumnEnabled(int column_n,bool v) { igTableSetColumnEnabled(column_n,v); }
const char* TableGetColumnName_IntV(int column_n) { return igTableGetColumnName_Int(column_n); }
bool BeginMainMenuBar() { return igBeginMainMenuBar(); }
ImU32 ColorConvertFloat4ToU32(const ImVec4 in) { return igColorConvertFloat4ToU32(in); }
bool InputScalarV(const char* label,ImGuiDataType data_type,void* p_data,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalar(label,data_type,p_data,p_step,p_step_fast,format,flags); }
void LabelText(const char* label,const char* fmt) { igLabelText(label,fmt); }
bool FontAtlas_Build(ImFontAtlas* self) { return ImFontAtlas_Build(self); }
void Font_RenderTextV(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin,float wrap_width,bool cpu_fine_clip) { ImFont_RenderText(self,draw_list,size,pos,col,clip_rect,text_begin,0,wrap_width,cpu_fine_clip); }
void Dummy(const ImVec2 size) { igDummy(size); }
void Viewport_Destroy(ImGuiViewport* self) { ImGuiViewport_destroy(self); }
ImGuiID GetID_Str(const char* str_id) { return igGetID_Str(str_id); }
ImGuiID GetID_StrStr(const char* str_id_begin,const char* str_id_end) { return igGetID_StrStr(str_id_begin,str_id_end); }
ImGuiID GetID_Ptr(const void* ptr_id) { return igGetID_Ptr(ptr_id); }
bool DragFloat3V(const char* label,float v[3],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat3(label,v,v_speed,v_min,v_max,format,flags); }
void ShowDemoWindowV(bool* p_open) { igShowDemoWindow(p_open); }
void StackTool_Destroy(ImGuiStackTool* self) { ImGuiStackTool_destroy(self); }
void EndTooltip() { igEndTooltip(); }
bool BeginMenuV(const char* label,bool enabled) { return igBeginMenu(label,enabled); }
void FontGlyphRangesBuilder_Destroy(ImFontGlyphRangesBuilder* self) { ImFontGlyphRangesBuilder_destroy(self); }
bool VSliderIntV(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderInt(label,size,v,v_min,v_max,format,flags); }
int GetColumnIndex() { return igGetColumnIndex(); }
void PlotHistogram_FloatPtrV(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride) { igPlotHistogram_FloatPtr(label,values,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size,stride); }
void PlotHistogram_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size) { igPlotHistogram_FnFloatPtr(label,values_getter,data,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size); }
void DrawListSharedData_Destroy(ImDrawListSharedData* self) { ImDrawListSharedData_destroy(self); }
bool ColorEdit3V(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorEdit3(label,col,flags); }
float GetColumnOffsetV(int column_index) { return igGetColumnOffset(column_index); }
void SetItemDefaultFocus() { igSetItemDefaultFocus(); }
void FontAtlas_ClearFonts(ImFontAtlas* self) { ImFontAtlas_ClearFonts(self); }
void BeginDisabledV(bool disabled) { igBeginDisabled(disabled); }
int GetFrameCount() { return igGetFrameCount(); }
void Render() { igRender(); }
void DrawListSplitter_Destroy(ImDrawListSplitter* self) { ImDrawListSplitter_destroy(self); }
void ListClipper_BeginV(ImGuiListClipper* self,int items_count,float items_height) { ImGuiListClipper_Begin(self,items_count,items_height); }
bool InputDoubleV(const char* label,double* v,double step,double step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputDouble(label,v,step,step_fast,format,flags); }
bool SliderScalarV(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalar(label,data_type,p_data,p_min,p_max,format,flags); }
ImFont* FontAtlas_AddFontDefaultV(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFontDefault(self,font_cfg); }
const ImWchar* FontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self); }
void Vec2ih_Destroy(ImVec2ih* self) { ImVec2ih_destroy(self); }
ImDrawList* DrawList_ImDrawList(const ImDrawListSharedData* shared_data) { return ImDrawList_ImDrawList(shared_data); }
bool Payload_IsDataType(ImGuiPayload* self,const char* type) { return ImGuiPayload_IsDataType(self,type); }
void TextBuffer_Reserve(ImGuiTextBuffer* self,int capacity) { ImGuiTextBuffer_reserve(self,capacity); }
bool ArrowButton(const char* str_id,ImGuiDir dir) { return igArrowButton(str_id,dir); }
void DrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self) { ImDrawListSplitter_ClearFreeMemory(self); }
void FontAtlas_ClearTexData(ImFontAtlas* self) { ImFontAtlas_ClearTexData(self); }
void GetFontTexUvWhitePixel(ImVec2 *pOut) { igGetFontTexUvWhitePixel(pOut); }
void Value_Bool(const char* prefix,bool b) { igValue_Bool(prefix,b); }
void Value_Int(const char* prefix,int v) { igValue_Int(prefix,v); }
void Value_Uint(const char* prefix,unsigned int v) { igValue_Uint(prefix,v); }
void Value_FloatV(const char* prefix,float v,const char* float_format) { igValue_Float(prefix,v,float_format); }
void SetScrollHereXV(float center_x_ratio) { igSetScrollHereX(center_x_ratio); }
void GetItemRectMax(ImVec2 *pOut) { igGetItemRectMax(pOut); }
bool BeginDragDropSourceV(ImGuiDragDropFlags flags) { return igBeginDragDropSource(flags); }
void OpenPopupOnItemClickV(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopupOnItemClick(str_id,popup_flags); }
void Style_ScaleAllSizes(ImGuiStyle* self,float scale_factor) { ImGuiStyle_ScaleAllSizes(self,scale_factor); }
bool BeginTableV(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width) { return igBeginTable(str_id,column,flags,outer_size,inner_width); }
void DrawList_AddBezierCubicV(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierCubic(self,p1,p2,p3,p4,col,thickness,num_segments); }
const ImWchar* FontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesCyrillic(self); }
void TableSetupColumnV(const char* label,ImGuiTableColumnFlags flags,float init_width_or_weight,ImGuiID user_id) { igTableSetupColumn(label,flags,init_width_or_weight,user_id); }
void TabBar_Destroy(ImGuiTabBar* self) { ImGuiTabBar_destroy(self); }
void CalcTextSize(ImVec2* pOut,const char* text) { CalcTextSizeV(pOut,text,false,-1.0f); }
bool SliderAngle(const char* label,float* v_rad) { return SliderAngleV(label,v_rad,-360.0f,+360.0f,"%.0f deg",0); }
bool ColorPicker3(const char* label,float col[3]) { return ColorPicker3V(label,col,0); }
bool DragInt(const char* label,int* v) { return DragIntV(label,v,1.0f,0,0,"%d",0); }
bool SliderInt3(const char* label,int v[3],int v_min,int v_max) { return SliderInt3V(label,v,v_min,v_max,"%d",0); }
bool SliderFloat4(const char* label,float v[4],float v_min,float v_max) { return SliderFloat4V(label,v,v_min,v_max,"%.3f",0); }
ImFont* FontAtlas_AddFontFromMemoryTTF(ImFontAtlas* self,void* font_data,int font_size,float size_pixels) { return FontAtlas_AddFontFromMemoryTTFV(self,font_data,font_size,size_pixels,NULL,NULL); }
bool IsKeyPressed(ImGuiKey key) { return IsKeyPressedV(key,true); }
void DrawList_AddImage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max) { DrawList_AddImageV(self,user_texture_id,p_min,p_max,(ImVec2){.x=0, .y=0},(ImVec2){.x=1, .y=1},4294967295); }
bool IsPopupOpen_Str(const char* str_id) { return IsPopupOpen_StrV(str_id,0); }
void SetScrollFromPosY_Float(float local_y) { SetScrollFromPosY_FloatV(local_y,0.5f); }
void RenderPlatformWindowsDefault() { RenderPlatformWindowsDefaultV(NULL,NULL); }
bool IsMousePosValid() { return IsMousePosValidV(NULL); }
void TextBuffer_Append(ImGuiTextBuffer* self,const char* str) { TextBuffer_AppendV(self,str,NULL); }
ImGuiContext* CreateContext() { return CreateContextV(NULL); }
void LogToClipboard() { LogToClipboardV(-1); }
void DrawList_PathBezierQuadraticCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3) { DrawList_PathBezierQuadraticCurveToV(self,p2,p3,0); }
bool BeginTabBar(const char* str_id) { return BeginTabBarV(str_id,0); }
bool SliderScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max) { return SliderScalarNV(label,data_type,p_data,components,p_min,p_max,NULL,0); }
bool BeginChildFrame(ImGuiID id,const ImVec2 size) { return BeginChildFrameV(id,size,0); }
void StyleColorsDark() { StyleColorsDarkV(NULL); }
ImGuiTableColumnFlags TableGetColumnFlags() { return TableGetColumnFlagsV(-1); }
void ShowStyleEditor() { ShowStyleEditorV(NULL); }
bool BeginPopup(const char* str_id) { return BeginPopupV(str_id,0); }
void Color_SetHSV(ImColor* self,float h,float s,float v) { Color_SetHSVV(self,h,s,v,1.0f); }
void DrawList_AddTriangle(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col) { DrawList_AddTriangleV(self,p1,p2,p3,col,1.0f); }
ImGuiID DockSpace(ImGuiID id) { return DockSpaceV(id,(ImVec2){.x=0, .y=0},0,NULL); }
void SetNextWindowCollapsed(bool collapsed) { SetNextWindowCollapsedV(collapsed,0); }
void SetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max) { SetNextWindowSizeConstraintsV(size_min,size_max,NULL,NULL); }
bool IsMouseDragging(ImGuiMouseButton button) { return IsMouseDraggingV(button,-1.0f); }
const ImGuiPayload* AcceptDragDropPayload(const char* type) { return AcceptDragDropPayloadV(type,0); }
void StyleColorsClassic() { StyleColorsClassicV(NULL); }
void DrawList_PathRect(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max) { DrawList_PathRectV(self,rect_min,rect_max,0.0f,0); }
void SetNextWindowSize(const ImVec2 size) { SetNextWindowSizeV(size,0); }
bool InputFloat4(const char* label,float v[4]) { return InputFloat4V(label,v,"%.3f",0); }
bool InputFloat3(const char* label,float v[3]) { return InputFloat3V(label,v,"%.3f",0); }
bool DragScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components) { return DragScalarNV(label,data_type,p_data,components,1.0f,NULL,NULL,NULL,0); }
bool BeginPopupContextWindow() { return BeginPopupContextWindowV(NULL,1); }
bool IsWindowFocused() { return IsWindowFocusedV(0); }
ImGuiID DockSpaceOverViewport() { return DockSpaceOverViewportV(NULL,0,NULL); }
void TextUnformatted(const char* text) { TextUnformattedV(text); }
bool SetDragDropPayload(const char* type,const void* data,size_t sz) { return SetDragDropPayloadV(type,data,sz,0); }
ImFont* FontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels) { return FontAtlas_AddFontFromMemoryCompressedTTFV(self,compressed_font_data,compressed_font_size,size_pixels,NULL,NULL); }
void StyleColorsLight() { StyleColorsLightV(NULL); }
bool TreeNodeEx_Str(const char* label) { return TreeNodeEx_StrV(label,0); }
bool TabItemButton(const char* label) { return TabItemButtonV(label,0); }
bool DragFloatRange2(const char* label,float* v_current_min,float* v_current_max) { return DragFloatRange2V(label,v_current_min,v_current_max,1.0f,0.0f,0.0f,"%.3f",NULL,0); }
bool ColorEdit4(const char* label,float col[4]) { return ColorEdit4V(label,col,0); }
bool IsItemClicked() { return IsItemClickedV(0); }
void Color_HSV(ImColor* pOut,float h,float s,float v) { Color_HSVV(pOut,h,s,v,1.0f); }
void PopStyleColor() { PopStyleColorV(1); }
void SameLine() { SameLineV(0.0f,-1.0f); }
void DrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max) { DrawList_PushClipRectV(self,clip_rect_min,clip_rect_max,false); }
bool DragFloat2(const char* label,float v[2]) { return DragFloat2V(label,v,1.0f,0.0f,0.0f,"%.3f",0); }
bool DragFloat(const char* label,float* v) { return DragFloatV(label,v,1.0f,0.0f,0.0f,"%.3f",0); }
bool InvisibleButton(const char* str_id,const ImVec2 size) { return InvisibleButtonV(str_id,size,0); }
void FontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self,const char* text) { FontGlyphRangesBuilder_AddTextV(self,text); }
bool DragScalar(const char* label,ImGuiDataType data_type,void* p_data) { return DragScalarV(label,data_type,p_data,1.0f,NULL,NULL,NULL,0); }
bool ColorPicker4(const char* label,float col[4]) { return ColorPicker4V(label,col,0,NULL); }
bool InputTextWithHint(const char* label,const char* hint,char* buf,size_t buf_size) { return InputTextWithHintV(label,hint,buf,buf_size,0,NULL,NULL); }
bool InputFloat2(const char* label,float v[2]) { return InputFloat2V(label,v,"%.3f",0); }
bool IsMouseClicked(ImGuiMouseButton button) { return IsMouseClickedV(button,false); }
void PlotLines_FloatPtr(const char* label,const float* values,int values_count) { PlotLines_FloatPtrV(label,values,values_count,0,NULL,igGET_FLT_MAX(),igGET_FLT_MAX(),(ImVec2){.x=0, .y=0},sizeof(float)); }
void SetKeyboardFocusHere() { SetKeyboardFocusHereV(0); }
bool TextFilter_Draw(ImGuiTextFilter* self) { return TextFilter_DrawV(self,"Filter(inc,-exc)",0.0f); }
bool CollapsingHeader_TreeNodeFlags(const char* label) { return CollapsingHeader_TreeNodeFlagsV(label,0); }
bool CollapsingHeader_BoolPtr(const char* label,bool* p_visible) { return CollapsingHeader_BoolPtrV(label,p_visible,0); }
void PushTextWrapPos() { PushTextWrapPosV(0.0f); }
bool BeginPopupContextItem() { return BeginPopupContextItemV(NULL,1); }
bool InputScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components) { return InputScalarNV(label,data_type,p_data,components,NULL,NULL,NULL,0); }
void ShowStackToolWindow() { ShowStackToolWindowV(NULL); }
void Columns() { ColumnsV(1,NULL,true); }
ImU32 GetColorU32_Col(ImGuiCol idx) { return GetColorU32_ColV(idx,1.0f); }
bool Selectable_Bool(const char* label) { return Selectable_BoolV(label,false,0,(ImVec2){.x=0, .y=0}); }
bool Selectable_BoolPtr(const char* label,bool* p_selected) { return Selectable_BoolPtrV(label,p_selected,0,(ImVec2){.x=0, .y=0}); }
void SetWindowCollapsed_Bool(bool collapsed) { SetWindowCollapsed_BoolV(collapsed,0); }
void SetWindowCollapsed_Str(const char* name,bool collapsed) { SetWindowCollapsed_StrV(name,collapsed,0); }
void ShowMetricsWindow() { ShowMetricsWindowV(NULL); }
void DrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col) { DrawList_AddLineV(self,p1,p2,col,1.0f); }
bool SliderInt4(const char* label,int v[4],int v_min,int v_max) { return SliderInt4V(label,v,v_min,v_max,"%d",0); }
void SetNextWindowDockID(ImGuiID dock_id) { SetNextWindowDockIDV(dock_id,0); }
bool SliderFloat(const char* label,float* v,float v_min,float v_max) { return SliderFloatV(label,v,v_min,v_max,"%.3f",0); }
bool BeginPopupContextVoid() { return BeginPopupContextVoidV(NULL,1); }
void Unindent() { UnindentV(0.0f); }
void FontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height) { FontAtlas_GetTexDataAsAlpha8V(self,out_pixels,out_width,out_height,NULL); }
void DrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col) { DrawList_AddRectFilledV(self,p_min,p_max,col,0.0f,0); }
void DrawList_PathArcTo(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max) { DrawList_PathArcToV(self,center,radius,a_min,a_max,0); }
bool DragInt4(const char* label,int v[4]) { return DragInt4V(label,v,1.0f,0,0,"%d",0); }
void DrawList_AddBezierQuadratic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness) { DrawList_AddBezierQuadraticV(self,p1,p2,p3,col,thickness,0); }
bool InputText(const char* label,char* buf,size_t buf_size) { return InputTextV(label,buf,buf_size,0,NULL,NULL); }
bool BeginPopupModal(const char* name) { return BeginPopupModalV(name,NULL,0); }
void SetScrollHereY() { SetScrollHereYV(0.5f); }
void FontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height) { FontAtlas_GetTexDataAsRGBA32V(self,out_pixels,out_width,out_height,NULL); }
void ShowAboutWindow() { ShowAboutWindowV(NULL); }
bool DragIntRange2(const char* label,int* v_current_min,int* v_current_max) { return DragIntRange2V(label,v_current_min,v_current_max,1.0f,0,0,"%d",NULL,0); }
bool ListBox_Str_arr(const char* label,int* current_item,const char* const items[],int items_count) { return ListBox_Str_arrV(label,current_item,items,items_count,-1); }
float GetColumnWidth() { return GetColumnWidthV(-1); }
void PopStyleVar() { PopStyleVarV(1); }
void LogToTTY() { LogToTTYV(-1); }
void InputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self,int pos,const char* text) { InputTextCallbackData_InsertCharsV(self,pos,text); }
void DestroyContext() { DestroyContextV(NULL); }
bool Begin(const char* name) { return BeginV(name,NULL,0); }
bool InputInt(const char* label,int* v) { return InputIntV(label,v,1,100,0); }
bool TextFilter_PassFilter(ImGuiTextFilter* self,const char* text) { return TextFilter_PassFilterV(self,text); }
void ShowDebugLogWindow() { ShowDebugLogWindowV(NULL); }
void DrawList_AddText_Vec2(ImDrawList* self,const ImVec2 pos,ImU32 col,const char* text_begin) { DrawList_AddText_Vec2V(self,pos,col,text_begin); }
void DrawList_AddText_FontPtr(ImDrawList* self,const ImFont* font,float font_size,const ImVec2 pos,ImU32 col,const char* text_begin) { DrawList_AddText_FontPtrV(self,font,font_size,pos,col,text_begin,0.0f,NULL); }
void LoadIniSettingsFromMemory(const char* ini_data) { LoadIniSettingsFromMemoryV(ini_data,0); }
bool BeginTabItem(const char* label) { return BeginTabItemV(label,NULL,0); }
bool IsWindowHovered() { return IsWindowHoveredV(0); }
void DrawList_AddNgon(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { DrawList_AddNgonV(self,center,radius,col,num_segments,1.0f); }
bool BeginChild_Str(const char* str_id) { return BeginChild_StrV(str_id,(ImVec2){.x=0, .y=0},false,0); }
bool BeginChild_ID(ImGuiID id) { return BeginChild_IDV(id,(ImVec2){.x=0, .y=0},false,0); }
void TableSetBgColor(ImGuiTableBgTarget target,ImU32 color) { TableSetBgColorV(target,color,-1); }
bool Combo_Str_arr(const char* label,int* current_item,const char* const items[],int items_count) { return Combo_Str_arrV(label,current_item,items,items_count,-1); }
bool Combo_Str(const char* label,int* current_item,const char* items_separated_by_zeros) { return Combo_StrV(label,current_item,items_separated_by_zeros,-1); }
void SetNextItemOpen(bool is_open) { SetNextItemOpenV(is_open,0); }
bool SliderInt2(const char* label,int v[2],int v_min,int v_max) { return SliderInt2V(label,v,v_min,v_max,"%d",0); }
bool SliderInt(const char* label,int* v,int v_min,int v_max) { return SliderIntV(label,v,v_min,v_max,"%d",0); }
void Indent() { IndentV(0.0f); }
void OpenPopup_Str(const char* str_id) { OpenPopup_StrV(str_id,0); }
void OpenPopup_ID(ImGuiID id) { OpenPopup_IDV(id,0); }
void ResetMouseDragDelta() { ResetMouseDragDeltaV(0); }
void DrawList_AddImageRounded(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding) { DrawList_AddImageRoundedV(self,user_texture_id,p_min,p_max,uv_min,uv_max,col,rounding,0); }
ImFont* FontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels) { return FontAtlas_AddFontFromMemoryCompressedBase85TTFV(self,compressed_font_data_base85,size_pixels,NULL,NULL); }
void LogToFile() { LogToFileV(-1,NULL); }
void SetWindowPos_Vec2(const ImVec2 pos) { SetWindowPos_Vec2V(pos,0); }
void SetWindowPos_Str(const char* name,const ImVec2 pos) { SetWindowPos_StrV(name,pos,0); }
bool SliderFloat3(const char* label,float v[3],float v_min,float v_max) { return SliderFloat3V(label,v,v_min,v_max,"%.3f",0); }
void SetNextWindowPos(const ImVec2 pos) { SetNextWindowPosV(pos,0,(ImVec2){.x=0, .y=0}); }
bool BeginCombo(const char* label,const char* preview_value) { return BeginComboV(label,preview_value,0); }
void DrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col) { DrawList_AddRectV(self,p_min,p_max,col,0.0f,0,1.0f); }
bool SliderFloat2(const char* label,float v[2],float v_min,float v_max) { return SliderFloat2V(label,v,v_min,v_max,"%.3f",0); }
void DrawList_AddQuad(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col) { DrawList_AddQuadV(self,p1,p2,p3,p4,col,1.0f); }
int FontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x) { return FontAtlas_AddCustomRectFontGlyphV(self,font,id,width,height,advance_x,(ImVec2){.x=0, .y=0}); }
void SetWindowSize_Vec2(const ImVec2 size) { SetWindowSize_Vec2V(size,0); }
void SetWindowSize_Str(const char* name,const ImVec2 size) { SetWindowSize_StrV(name,size,0); }
void Font_AddRemapChar(ImFont* self,ImWchar dst,ImWchar src) { Font_AddRemapCharV(self,dst,src,true); }
void DrawList_AddCircle(ImDrawList* self,const ImVec2 center,float radius,ImU32 col) { DrawList_AddCircleV(self,center,radius,col,0,1.0f); }
bool DragInt3(const char* label,int v[3]) { return DragInt3V(label,v,1.0f,0,0,"%d",0); }
bool MenuItem_Bool(const char* label) { return MenuItem_BoolV(label,NULL,false,true); }
bool MenuItem_BoolPtr(const char* label,const char* shortcut,bool* p_selected) { return MenuItem_BoolPtrV(label,shortcut,p_selected,true); }
bool DragFloat4(const char* label,float v[4]) { return DragFloat4V(label,v,1.0f,0.0f,0.0f,"%.3f",0); }
void SetScrollFromPosX_Float(float local_x) { SetScrollFromPosX_FloatV(local_x,0.5f); }
bool IsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max) { return IsMouseHoveringRectV(r_min,r_max,true); }
bool DragInt2(const char* label,int v[2]) { return DragInt2V(label,v,1.0f,0,0,"%d",0); }
void ProgressBar(float fraction) { ProgressBarV(fraction,(ImVec2){.x=-1*igGET_FLT_MIN(), .y=0},NULL); }
void GetMouseDragDelta(ImVec2* pOut) { GetMouseDragDeltaV(pOut,0,-1.0f); }
void DrawList_PathBezierCubicCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4) { DrawList_PathBezierCubicCurveToV(self,p2,p3,p4,0); }
bool InputTextMultiline(const char* label,char* buf,size_t buf_size) { return InputTextMultilineV(label,buf,buf_size,(ImVec2){.x=0, .y=0},0,NULL,NULL); }
bool InputInt2(const char* label,int v[2]) { return InputInt2V(label,v,0); }
bool VSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max) { return VSliderFloatV(label,size,v,v_min,v_max,"%.3f",0); }
const char* SaveIniSettingsToMemory() { return SaveIniSettingsToMemoryV(NULL); }
bool Button(const char* label) { return ButtonV(label,(ImVec2){.x=0, .y=0}); }
bool InputFloat(const char* label,float* v) { return InputFloatV(label,v,0.0f,0.0f,"%.3f",0); }
void Font_CalcTextSizeA(ImVec2* pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin) { Font_CalcTextSizeAV(pOut,self,size,max_width,wrap_width,text_begin,NULL); }
bool InputInt3(const char* label,int v[3]) { return InputInt3V(label,v,0); }
void DrawList_AddCircleFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col) { DrawList_AddCircleFilledV(self,center,radius,col,0); }
void TableNextRow() { TableNextRowV(0,0.0f); }
bool InputInt4(const char* label,int v[4]) { return InputInt4V(label,v,0); }
bool VSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max) { return VSliderScalarV(label,size,data_type,p_data,p_min,p_max,NULL,0); }
bool ColorButton(const char* desc_id,const ImVec4 col) { return ColorButtonV(desc_id,col,0,(ImVec2){.x=0, .y=0}); }
void DrawList_PathStroke(ImDrawList* self,ImU32 col) { DrawList_PathStrokeV(self,col,0,1.0f); }
void DrawList_AddImageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4) { DrawList_AddImageQuadV(self,user_texture_id,p1,p2,p3,p4,(ImVec2){.x=0, .y=0},(ImVec2){.x=1, .y=0},(ImVec2){.x=1, .y=1},(ImVec2){.x=0, .y=1},4294967295); }
void TreePush_Ptr() { TreePush_PtrV(NULL); }
bool IsItemHovered() { return IsItemHoveredV(0); }
bool BeginListBox(const char* label) { return BeginListBoxV(label,(ImVec2){.x=0, .y=0}); }
ImFont* FontAtlas_AddFontFromFileTTF(ImFontAtlas* self,const char* filename,float size_pixels) { return FontAtlas_AddFontFromFileTTFV(self,filename,size_pixels,NULL,NULL); }
void IO_SetKeyEventNativeData(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode) { IO_SetKeyEventNativeDataV(self,key,native_keycode,native_scancode,-1); }
const char* TableGetColumnName_Int() { return TableGetColumnName_IntV(-1); }
bool InputScalar(const char* label,ImGuiDataType data_type,void* p_data) { return InputScalarV(label,data_type,p_data,NULL,NULL,NULL,0); }
void Font_RenderText(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin) { Font_RenderTextV(self,draw_list,size,pos,col,clip_rect,text_begin,0.0f,false); }
bool DragFloat3(const char* label,float v[3]) { return DragFloat3V(label,v,1.0f,0.0f,0.0f,"%.3f",0); }
void ShowDemoWindow() { ShowDemoWindowV(NULL); }
bool BeginMenu(const char* label) { return BeginMenuV(label,true); }
bool VSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max) { return VSliderIntV(label,size,v,v_min,v_max,"%d",0); }
void PlotHistogram_FloatPtr(const char* label,const float* values,int values_count) { PlotHistogram_FloatPtrV(label,values,values_count,0,NULL,igGET_FLT_MAX(),igGET_FLT_MAX(),(ImVec2){.x=0, .y=0},sizeof(float)); }
bool ColorEdit3(const char* label,float col[3]) { return ColorEdit3V(label,col,0); }
float GetColumnOffset() { return GetColumnOffsetV(-1); }
void BeginDisabled() { BeginDisabledV(true); }
void ListClipper_Begin(ImGuiListClipper* self,int items_count) { ListClipper_BeginV(self,items_count,-1.0f); }
bool InputDouble(const char* label,double* v) { return InputDoubleV(label,v,0.0,0.0,"%.6f",0); }
bool SliderScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max) { return SliderScalarV(label,data_type,p_data,p_min,p_max,NULL,0); }
ImFont* FontAtlas_AddFontDefault(ImFontAtlas* self) { return FontAtlas_AddFontDefaultV(self,NULL); }
void Value_Float(const char* prefix,float v) { Value_FloatV(prefix,v,NULL); }
void SetScrollHereX() { SetScrollHereXV(0.5f); }
bool BeginDragDropSource() { return BeginDragDropSourceV(0); }
void OpenPopupOnItemClick() { OpenPopupOnItemClickV(NULL,1); }
bool BeginTable(const char* str_id,int column) { return BeginTableV(str_id,column,0,(ImVec2){.x=0, .y=0},0.0f); }
void DrawList_AddBezierCubic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness) { DrawList_AddBezierCubicV(self,p1,p2,p3,p4,col,thickness,0); }
void TableSetupColumn(const char* label) { TableSetupColumnV(label,0,0.0f,0); }
