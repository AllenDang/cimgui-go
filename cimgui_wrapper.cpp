#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

void DrawList_AddNgonFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddNgonFilled(self,center,radius,col,num_segments); }
bool Checkbox(const char* label,bool* v) { return igCheckbox(label,v); }
const char* GetKeyName(ImGuiKey key) { return igGetKeyName(key); }
void LogToClipboard(int auto_open_depth) { igLogToClipboard(auto_open_depth); }
int GetFrameCount() { return igGetFrameCount(); }
bool IsMouseDragging(ImGuiMouseButton button,float lock_threshold) { return igIsMouseDragging(button,lock_threshold); }
void PopClipRect() { igPopClipRect(); }
void DrawList_PathArcToFast(ImDrawList* self,const ImVec2 center,float radius,int a_min_of_12,int a_max_of_12) { ImDrawList_PathArcToFast(self,center,radius,a_min_of_12,a_max_of_12); }
const ImGuiPayload* AcceptDragDropPayload(const char* type,ImGuiDragDropFlags flags) { return igAcceptDragDropPayload(type,flags); }
float GetScrollY() { return igGetScrollY(); }
void DrawList_PopClipRect(ImDrawList* self) { ImDrawList_PopClipRect(self); }
bool BeginTabItem(const char* label,bool* p_open,ImGuiTabItemFlags flags) { return igBeginTabItem(label,p_open,flags); }
bool DragIntRange2(const char* label,int* v_current_min,int* v_current_max,float v_speed,int v_min,int v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragIntRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
bool IsItemActivated() { return igIsItemActivated(); }
bool DragScalar(const char* label,ImGuiDataType data_type,void* p_data,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalar(label,data_type,p_data,v_speed,p_min,p_max,format,flags); }
bool InputInt(const char* label,int* v,int step,int step_fast,ImGuiInputTextFlags flags) { return igInputInt(label,v,step,step_fast,flags); }
void StyleColorsDark(ImGuiStyle* dst) { igStyleColorsDark(dst); }
ImDrawData* DrawData_DrawData() { return ImDrawData_ImDrawData(); }
void AlignTextToFramePadding() { igAlignTextToFramePadding(); }
void GetMouseDragDelta(ImVec2 *pOut,ImGuiMouseButton button,float lock_threshold) { igGetMouseDragDelta(pOut,button,lock_threshold); }
float GetScrollMaxY() { return igGetScrollMaxY(); }
void LogToFile(int auto_open_depth,const char* filename) { igLogToFile(auto_open_depth,filename); }
void DrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self) { ImDrawListSplitter_ClearFreeMemory(self); }
ImGuiTableSortSpecs* TableSortSpecs_TableSortSpecs() { return ImGuiTableSortSpecs_ImGuiTableSortSpecs(); }
bool BeginDragDropSource(ImGuiDragDropFlags flags) { return igBeginDragDropSource(flags); }
void EndCombo() { igEndCombo(); }
void EndMenu() { igEndMenu(); }
int GetKeyIndex(ImGuiKey key) { return igGetKeyIndex(key); }
bool IsItemHovered(ImGuiHoveredFlags flags) { return igIsItemHovered(flags); }
void TableSetupColumn(const char* label,ImGuiTableColumnFlags flags,float init_width_or_weight,ImGuiID user_id) { igTableSetupColumn(label,flags,init_width_or_weight,user_id); }
void DrawListSplitter_Clear(ImDrawListSplitter* self) { ImDrawListSplitter_Clear(self); }
void DrawList_GetClipRectMin(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMin(pOut,self); }
void DrawList_PathArcTo(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments) { ImDrawList_PathArcTo(self,center,radius,a_min,a_max,num_segments); }
ImFont* FontAtlas_AddFont(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFont(self,font_cfg); }
void Viewport_GetCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetCenter(pOut,self); }
void EndChild() { igEndChild(); }
void ShowAboutWindow(bool* p_open) { igShowAboutWindow(p_open); }
const ImWchar* FontAtlas_GetGlyphRangesDefault(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesDefault(self); }
int* Storage_GetIntRef(ImGuiStorage* self,ImGuiID key,int default_val) { return ImGuiStorage_GetIntRef(self,key,default_val); }
const char* TextBuffer_c_str(ImGuiTextBuffer* self) { return ImGuiTextBuffer_c_str(self); }
void ListClipper_ForceDisplayRangeByIndices(ImGuiListClipper* self,int item_min,int item_max) { ImGuiListClipper_ForceDisplayRangeByIndices(self,item_min,item_max); }
void SetTabItemClosed(const char* tab_or_docked_window_label) { igSetTabItemClosed(tab_or_docked_window_label); }
ImFont* FontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self,compressed_font_data_base85,size_pixels,font_cfg,glyph_ranges); }
void EndTabBar() { igEndTabBar(); }
void GetWindowSize(ImVec2 *pOut) { igGetWindowSize(pOut); }
void SetNextWindowSize(const ImVec2 size,ImGuiCond cond) { igSetNextWindowSize(size,cond); }
bool VSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderInt(label,size,v,v_min,v_max,format,flags); }
const ImWchar* FontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesJapanese(self); }
void** Storage_GetVoidPtrRef(ImGuiStorage* self,ImGuiID key,void* default_val) { return ImGuiStorage_GetVoidPtrRef(self,key,default_val); }
ImGuiViewport* FindViewportByID(ImGuiID id) { return igFindViewportByID(id); }
void SetNextFrameWantCaptureMouse(bool want_capture_mouse) { igSetNextFrameWantCaptureMouse(want_capture_mouse); }
void NewFrame() { igNewFrame(); }
void SetMouseCursor(ImGuiMouseCursor cursor_type) { igSetMouseCursor(cursor_type); }
bool ColorEdit4(const char* label,float col[4],ImGuiColorEditFlags flags) { return igColorEdit4(label,col,flags); }
bool SmallButton(const char* label) { return igSmallButton(label); }
void StyleColorsLight(ImGuiStyle* dst) { igStyleColorsLight(dst); }
ImGuiID GetWindowDockID() { return igGetWindowDockID(); }
void EndTabItem() { igEndTabItem(); }
float GetWindowHeight() { return igGetWindowHeight(); }
bool IsItemVisible() { return igIsItemVisible(); }
void SetNextWindowViewport(ImGuiID viewport_id) { igSetNextWindowViewport(viewport_id); }
bool IsMouseClicked(ImGuiMouseButton button,bool repeat) { return igIsMouseClicked(button,repeat); }
void DrawData_ScaleClipRects(ImDrawData* self,const ImVec2 fb_scale) { ImDrawData_ScaleClipRects(self,fb_scale); }
void DrawList_AddDrawCmd(ImDrawList* self) { ImDrawList_AddDrawCmd(self); }
void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) { igSetNextFrameWantCaptureKeyboard(want_capture_keyboard); }
bool SliderInt(const char* label,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt(label,v,v_min,v_max,format,flags); }
void IO_AddKeyEvent(ImGuiIO* self,ImGuiKey key,bool down) { ImGuiIO_AddKeyEvent(self,key,down); }
bool ColorPicker3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorPicker3(label,col,flags); }
void EndDisabled() { igEndDisabled(); }
void NewLine() { igNewLine(); }
void LoadIniSettingsFromDisk(const char* ini_filename) { igLoadIniSettingsFromDisk(ini_filename); }
void GetFontTexUvWhitePixel(ImVec2 *pOut) { igGetFontTexUvWhitePixel(pOut); }
float GetCursorPosY() { return igGetCursorPosY(); }
void ShowStackToolWindow(bool* p_open) { igShowStackToolWindow(p_open); }
bool Font_IsGlyphRangeUnused(ImFont* self,unsigned int c_begin,unsigned int c_last) { return ImFont_IsGlyphRangeUnused(self,c_begin,c_last); }
void PushFont(ImFont* font) { igPushFont(font); }
bool SliderAngle(const char* label,float* v_rad,float v_degrees_min,float v_degrees_max,const char* format,ImGuiSliderFlags flags) { return igSliderAngle(label,v_rad,v_degrees_min,v_degrees_max,format,flags); }
void FontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self,ImWchar c) { ImFontGlyphRangesBuilder_AddChar(self,c); }
const char* Font_GetDebugName(ImFont* self) { return ImFont_GetDebugName(self); }
bool BeginTable(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width) { return igBeginTable(str_id,column,flags,outer_size,inner_width); }
void DrawList_AddBezierQuadratic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierQuadratic(self,p1,p2,p3,col,thickness,num_segments); }
void DrawList_AddTriangle(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness) { ImDrawList_AddTriangle(self,p1,p2,p3,col,thickness); }
void GetMousePos(ImVec2 *pOut) { igGetMousePos(pOut); }
void PopAllowKeyboardFocus() { igPopAllowKeyboardFocus(); }
bool* Storage_GetBoolRef(ImGuiStorage* self,ImGuiID key,bool default_val) { return ImGuiStorage_GetBoolRef(self,key,default_val); }
void SetNextWindowDockID(ImGuiID dock_id,ImGuiCond cond) { igSetNextWindowDockID(dock_id,cond); }
void ColorConvertRGBtoHSV(float r,float g,float b,float* out_h,float* out_s,float* out_v) { igColorConvertRGBtoHSV(r,g,b,out_h,out_s,out_v); }
void TreePop() { igTreePop(); }
ImFontAtlasCustomRect* FontAtlas_GetCustomRectByIndex(ImFontAtlas* self,int index) { return ImFontAtlas_GetCustomRectByIndex(self,index); }
void IO_AddMouseViewportEvent(ImGuiIO* self,ImGuiID id) { ImGuiIO_AddMouseViewportEvent(self,id); }
void Payload_Clear(ImGuiPayload* self) { ImGuiPayload_Clear(self); }
void FontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self,ImVector_ImWchar* out_ranges) { ImFontGlyphRangesBuilder_BuildRanges(self,out_ranges); }
void DrawList_AddageRounded(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddImageRounded(self,user_texture_id,p_min,p_max,uv_min,uv_max,col,rounding,flags); }
bool IsMousePosValid(const ImVec2* mouse_pos) { return igIsMousePosValid(mouse_pos); }
ImGuiTableSortSpecs* TableGetSortSpecs() { return igTableGetSortSpecs(); }
void TableHeader(const char* label) { igTableHeader(label); }
void TableSetBgColor(ImGuiTableBgTarget target,ImU32 color,int column_n) { igTableSetBgColor(target,color,column_n); }
void FontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self,const ImWchar* ranges) { ImFontGlyphRangesBuilder_AddRanges(self,ranges); }
void Font_RenderText(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin,const char* text_end,float wrap_width,bool cpu_fine_clip) { ImFont_RenderText(self,draw_list,size,pos,col,clip_rect,text_begin,text_end,wrap_width,cpu_fine_clip); }
void InputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self,int pos,int bytes_count) { ImGuiInputTextCallbackData_DeleteChars(self,pos,bytes_count); }
ImGuiTextFilter* TextFilter_TextFilter(const char* default_filter) { return ImGuiTextFilter_ImGuiTextFilter(default_filter); }
float GetTextLineHeight() { return igGetTextLineHeight(); }
double GetTime() { return igGetTime(); }
void SetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data) { igSetNextWindowSizeConstraints(size_min,size_max,custom_callback,custom_callback_data); }
void ShowStyleEditor(ImGuiStyle* ref) { igShowStyleEditor(ref); }
ImFontAtlasCustomRect* FontAtlasCustomRect_FontAtlasCustomRect() { return ImFontAtlasCustomRect_ImFontAtlasCustomRect(); }
bool IsAnyItemHovered() { return igIsAnyItemHovered(); }
bool ColorButton(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size) { return igColorButton(desc_id,col,flags,size); }
void TableSetupScrollFreeze(int cols,int rows) { igTableSetupScrollFreeze(cols,rows); }
int FontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x,const ImVec2 offset) { return ImFontAtlas_AddCustomRectFontGlyph(self,font,id,width,height,advance_x,offset); }
void FontAtlas_Clear(ImFontAtlas* self) { ImFontAtlas_Clear(self); }
float GetFontSize() { return igGetFontSize(); }
bool ListClipper_Step(ImGuiListClipper* self) { return ImGuiListClipper_Step(self); }
bool BeginCombo(const char* label,const char* preview_value,ImGuiComboFlags flags) { return igBeginCombo(label,preview_value,flags); }
void SetNextWindowCollapsed(bool collapsed,ImGuiCond cond) { igSetNextWindowCollapsed(collapsed,cond); }
ImGuiViewport* Viewport_Viewport() { return ImGuiViewport_ImGuiViewport(); }
void EndDragDropSource() { igEndDragDropSource(); }
ImGuiTableColumnFlags TableGetColumnFlags(int column_n) { return igTableGetColumnFlags(column_n); }
void IO_AddInputCharacterUTF16(ImGuiIO* self,ImWchar16 c) { ImGuiIO_AddInputCharacterUTF16(self,c); }
void Spacing() { igSpacing(); }
void DrawList_GetClipRectMax(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMax(pOut,self); }
void DrawList_Addage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col) { ImDrawList_AddImage(self,user_texture_id,p_min,p_max,uv_min,uv_max,col); }
void GetContentRegionAvail(ImVec2 *pOut) { igGetContentRegionAvail(pOut); }
int TableGetRowIndex() { return igTableGetRowIndex(); }
bool TableNextColumn() { return igTableNextColumn(); }
ImDrawListSplitter* DrawListSplitter_DrawListSplitter() { return ImDrawListSplitter_ImDrawListSplitter(); }
const ImFontGlyph* Font_FindGlyphNoFallback(ImFont* self,ImWchar c) { return ImFont_FindGlyphNoFallback(self,c); }
void Storage_SetBool(ImGuiStorage* self,ImGuiID key,bool val) { ImGuiStorage_SetBool(self,key,val); }
void DrawList_PushClipRectFullScreen(ImDrawList* self) { ImDrawList_PushClipRectFullScreen(self); }
void DrawList_AddCircle(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddCircle(self,center,radius,col,num_segments,thickness); }
void ListClipper_Begin(ImGuiListClipper* self,int items_count,float items_height) { ImGuiListClipper_Begin(self,items_count,items_height); }
void DrawList_AddTriangleFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col) { ImDrawList_AddTriangleFilled(self,p1,p2,p3,col); }
ImGuiID DockSpaceOverViewport(const ImGuiViewport* viewport,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpaceOverViewport(viewport,flags,window_class); }
float GetColumnWidth(int column_index) { return igGetColumnWidth(column_index); }
void Storage_BuildSortByKey(ImGuiStorage* self) { ImGuiStorage_BuildSortByKey(self); }
ImDrawListSharedData* GetDrawListSharedData() { return igGetDrawListSharedData(); }
bool InputFloat4(const char* label,float v[4],const char* format,ImGuiInputTextFlags flags) { return igInputFloat4(label,v,format,flags); }
bool InputScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalar(label,data_type,p_data,p_step,p_step_fast,format,flags); }
void SameLine(float offset_from_start_x,float spacing) { igSameLine(offset_from_start_x,spacing); }
void Font_RenderChar(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,ImWchar c) { ImFont_RenderChar(self,draw_list,size,pos,col,c); }
void DrawList_AddNgon(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddNgon(self,center,radius,col,num_segments,thickness); }
bool FontAtlas_Build(ImFontAtlas* self) { return ImFontAtlas_Build(self); }
void EndMainMenuBar() { igEndMainMenuBar(); }
bool InputScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalarN(label,data_type,p_data,components,p_step,p_step_fast,format,flags); }
bool IsWindowAppearing() { return igIsWindowAppearing(); }
const ImWchar* FontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseFull(self); }
ImGuiContext* CreateContext(ImFontAtlas* shared_font_atlas) { return igCreateContext(shared_font_atlas); }
ImGuiID DockSpace(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpace(id,size,flags,window_class); }
bool SliderFloat2(const char* label,float v[2],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat2(label,v,v_min,v_max,format,flags); }
void Storage_Clear(ImGuiStorage* self) { ImGuiStorage_Clear(self); }
ImGuiTextBuffer* TextBuffer_TextBuffer() { return ImGuiTextBuffer_ImGuiTextBuffer(); }
float GetFrameHeight() { return igGetFrameHeight(); }
bool ColorPicker4(const char* label,float col[4],ImGuiColorEditFlags flags,const float* ref_col) { return igColorPicker4(label,col,flags,ref_col); }
void DestroyPlatformWindows() { igDestroyPlatformWindows(); }
ImGuiOnceUponAFrame* OnceUponAFrame_OnceUponAFrame() { return ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(); }
float GetScrollX() { return igGetScrollX(); }
void PushTextWrapPos(float wrap_local_pos_x) { igPushTextWrapPos(wrap_local_pos_x); }
bool TabItemButton(const char* label,ImGuiTabItemFlags flags) { return igTabItemButton(label,flags); }
void DrawList_AddBezierCubic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierCubic(self,p1,p2,p3,p4,col,thickness,num_segments); }
void DebugTextEncoding(const char* text) { igDebugTextEncoding(text); }
void DrawList_PathClear(ImDrawList* self) { ImDrawList_PathClear(self); }
void InputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_SelectAll(self); }
int GetColumnIndex() { return igGetColumnIndex(); }
bool VSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igVSliderScalar(label,size,data_type,p_data,p_min,p_max,format,flags); }
bool BeginPopupModal(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBeginPopupModal(name,p_open,flags); }
void GetWindowContentRegionMax(ImVec2 *pOut) { igGetWindowContentRegionMax(pOut); }
void SetNextWindowFocus() { igSetNextWindowFocus(); }
ImFont* FontAtlas_AddFontDefault(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFontDefault(self,font_cfg); }
void ColorConvertHSVtoRGB(float h,float s,float v,float* out_r,float* out_g,float* out_b) { igColorConvertHSVtoRGB(h,s,v,out_r,out_g,out_b); }
bool InputFloat3(const char* label,float v[3],const char* format,ImGuiInputTextFlags flags) { return igInputFloat3(label,v,format,flags); }
bool IsItemToggledOpen() { return igIsItemToggledOpen(); }
void InputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self,int pos,const char* text,const char* text_end) { ImGuiInputTextCallbackData_InsertChars(self,pos,text,text_end); }
bool SliderFloat3(const char* label,float v[3],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat3(label,v,v_min,v_max,format,flags); }
bool SliderScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalar(label,data_type,p_data,p_min,p_max,format,flags); }
void SetCursorScreenPos(const ImVec2 pos) { igSetCursorScreenPos(pos); }
void TableHeadersRow() { igTableHeadersRow(); }
void DrawList_PrimUnreserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimUnreserve(self,idx_count,vtx_count); }
const ImWchar* FontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesCyrillic(self); }
bool BeginDragDropTarget() { return igBeginDragDropTarget(); }
bool IsItemEdited() { return igIsItemEdited(); }
void MemFree(void* ptr) { igMemFree(ptr); }
void StyleColorsClassic(ImGuiStyle* dst) { igStyleColorsClassic(dst); }
void SetItemDefaultFocus() { igSetItemDefaultFocus(); }
void DrawList_ChannelsSplit(ImDrawList* self,int count) { ImDrawList_ChannelsSplit(self,count); }
void GetItemRectMax(ImVec2 *pOut) { igGetItemRectMax(pOut); }
ImFont* GetFont() { return igGetFont(); }
ImGuiViewport* GetWindowViewport() { return igGetWindowViewport(); }
void SetCursorPosX(float local_x) { igSetCursorPosX(local_x); }
bool SliderFloat4(const char* label,float v[4],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat4(label,v,v_min,v_max,format,flags); }
ImDrawList* DrawList_DrawList(const ImDrawListSharedData* shared_data) { return ImDrawList_ImDrawList(shared_data); }
ImFont* FontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedTTF(self,compressed_font_data,compressed_font_size,size_pixels,font_cfg,glyph_ranges); }
ImFont* FontAtlas_AddFontFromMemoryTTF(ImFontAtlas* self,void* font_data,int font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryTTF(self,font_data,font_size,size_pixels,font_cfg,glyph_ranges); }
void IO_AddInputCharacter(ImGuiIO* self,unsigned int c) { ImGuiIO_AddInputCharacter(self,c); }
float* Storage_GetFloatRef(ImGuiStorage* self,ImGuiID key,float default_val) { return ImGuiStorage_GetFloatRef(self,key,default_val); }
void Storage_SetAllInt(ImGuiStorage* self,int val) { ImGuiStorage_SetAllInt(self,val); }
bool TextFilter_PassFilter(ImGuiTextFilter* self,const char* text,const char* text_end) { return ImGuiTextFilter_PassFilter(self,text,text_end); }
float GetScrollMaxX() { return igGetScrollMaxX(); }
void TextUnformatted(const char* text,const char* text_end) { igTextUnformatted(text,text_end); }
void DrawList_AddQuadFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col) { ImDrawList_AddQuadFilled(self,p1,p2,p3,p4,col); }
float GetWindowDpiScale() { return igGetWindowDpiScale(); }
void Storage_SetFloat(ImGuiStorage* self,ImGuiID key,float val) { ImGuiStorage_SetFloat(self,key,val); }
bool SliderInt3(const char* label,int v[3],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt3(label,v,v_min,v_max,format,flags); }
void TableSetColumnEnabled(int column_n,bool v) { igTableSetColumnEnabled(column_n,v); }
void DrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddRectFilled(self,p_min,p_max,col,rounding,flags); }
void FontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self) { ImFontGlyphRangesBuilder_Clear(self); }
void EndListBox() { igEndListBox(); }
const ImFontGlyph* Font_FindGlyph(ImFont* self,ImWchar c) { return ImFont_FindGlyph(self,c); }
float Font_GetCharAdvance(ImFont* self,ImWchar c) { return ImFont_GetCharAdvance(self,c); }
ImGuiViewport* GetMainViewport() { return igGetMainViewport(); }
void SetItemAllowOverlap() { igSetItemAllowOverlap(); }
bool BeginChildFrame(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags) { return igBeginChildFrame(id,size,flags); }
bool BeginPopupContextWindow(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextWindow(str_id,popup_flags); }
bool SliderInt2(const char* label,int v[2],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt2(label,v,v_min,v_max,format,flags); }
bool VSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderFloat(label,size,v,v_min,v_max,format,flags); }
void LogFinish() { igLogFinish(); }
void FontAtlas_CalcCustomRectUV(ImFontAtlas* self,const ImFontAtlasCustomRect* rect,ImVec2* out_uv_min,ImVec2* out_uv_max) { ImFontAtlas_CalcCustomRectUV(self,rect,out_uv_min,out_uv_max); }
ImGuiTableColumnSortSpecs* TableColumnSortSpecs_TableColumnSortSpecs() { return ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(); }
void SetNextWindowContentSize(const ImVec2 size) { igSetNextWindowContentSize(size); }
bool SetDragDropPayload(const char* type,const void* data,size_t sz,ImGuiCond cond) { return igSetDragDropPayload(type,data,sz,cond); }
void ListClipper_End(ImGuiListClipper* self) { ImGuiListClipper_End(self); }
void DrawList_PushTextureID(ImDrawList* self,ImTextureID texture_id) { ImDrawList_PushTextureID(self,texture_id); }
bool IsWindowFocused(ImGuiFocusedFlags flags) { return igIsWindowFocused(flags); }
ImFontConfig* FontConf_FontConfig() { return ImFontConfig_ImFontConfig(); }
ImGuiWindowClass* WindowClass_WindowClass() { return ImGuiWindowClass_ImGuiWindowClass(); }
void DrawList_PrimRect(ImDrawList* self,const ImVec2 a,const ImVec2 b,ImU32 col) { ImDrawList_PrimRect(self,a,b,col); }
bool DragInt4(const char* label,int v[4],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt4(label,v,v_speed,v_min,v_max,format,flags); }
void GetCursorStartPos(ImVec2 *pOut) { igGetCursorStartPos(pOut); }
int GetKeyPressedAmount(ImGuiKey key,float repeat_delay,float rate) { return igGetKeyPressedAmount(key,repeat_delay,rate); }
void Render() { igRender(); }
void SetColumnOffset(int column_index,float offset_x) { igSetColumnOffset(column_index,offset_x); }
bool DragScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalarN(label,data_type,p_data,components,v_speed,p_min,p_max,format,flags); }
void* Storage_GetVoidPtr(ImGuiStorage* self,ImGuiID key) { return ImGuiStorage_GetVoidPtr(self,key); }
void RenderPlatformWindowsDefault(void* platform_render_arg,void* renderer_render_arg) { igRenderPlatformWindowsDefault(platform_render_arg,renderer_render_arg); }
ImGuiPayload* Payload_Payload() { return ImGuiPayload_ImGuiPayload(); }
bool IsKeyReleased(ImGuiKey key) { return igIsKeyReleased(key); }
bool FontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect* self) { return ImFontAtlasCustomRect_IsPacked(self); }
ImFontGlyphRangesBuilder* FontGlyphRangesBuilder_FontGlyphRangesBuilder() { return ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(); }
void* MemAlloc(size_t size) { return igMemAlloc(size); }
bool ArrowButton(const char* str_id,ImGuiDir dir) { return igArrowButton(str_id,dir); }
bool BeginPopupContextVoid(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextVoid(str_id,popup_flags); }
void DrawList_PathStroke(ImDrawList* self,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_PathStroke(self,col,flags,thickness); }
ImU32 ColorConvertFloat4ToU32(const ImVec4 in) { return igColorConvertFloat4ToU32(in); }
bool DragFloat2(const char* label,float v[2],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat2(label,v,v_speed,v_min,v_max,format,flags); }
void DrawList_PathRect(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max,float rounding,ImDrawFlags flags) { ImDrawList_PathRect(self,rect_min,rect_max,rounding,flags); }
void SetClipboardText(const char* text) { igSetClipboardText(text); }
bool SliderInt4(const char* label,int v[4],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt4(label,v,v_min,v_max,format,flags); }
ImFontAtlas* FontAtlas_FontAtlas() { return ImFontAtlas_ImFontAtlas(); }
void EndPopup() { igEndPopup(); }
void PopStyleColor(int count) { igPopStyleColor(count); }
void DrawList_PrimQuadUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 d,const ImVec2 uv_a,const ImVec2 uv_b,const ImVec2 uv_c,const ImVec2 uv_d,ImU32 col) { ImDrawList_PrimQuadUV(self,a,b,c,d,uv_a,uv_b,uv_c,uv_d,col); }
void IO_AddKeyAnalogEvent(ImGuiIO* self,ImGuiKey key,bool down,float v) { ImGuiIO_AddKeyAnalogEvent(self,key,down,v); }
void PopID() { igPopID(); }
void DrawList_AddQuad(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness) { ImDrawList_AddQuad(self,p1,p2,p3,p4,col,thickness); }
void DrawList_PrimVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimVtx(self,pos,uv,col); }
void UpdatePlatformWindows() { igUpdatePlatformWindows(); }
ImFont* FontAtlas_AddFontFromFileTTF(ImFontAtlas* self,const char* filename,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromFileTTF(self,filename,size_pixels,font_cfg,glyph_ranges); }
void Font_GrowIndex(ImFont* self,int new_size) { ImFont_GrowIndex(self,new_size); }
void DestroyContext(ImGuiContext* ctx) { igDestroyContext(ctx); }
void DrawList_ChannelsSetCurrent(ImDrawList* self,int n) { ImDrawList_ChannelsSetCurrent(self,n); }
void DrawList_PathBezierCubicCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,int num_segments) { ImDrawList_PathBezierCubicCurveTo(self,p2,p3,p4,num_segments); }
void DrawList_PrimWriteVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimWriteVtx(self,pos,uv,col); }
void DrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness) { ImDrawList_AddLine(self,p1,p2,col,thickness); }
void InputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_ClearSelection(self); }
void SetKeyboardFocusHere(int offset) { igSetKeyboardFocusHere(offset); }
void ShowUserGuide() { igShowUserGuide(); }
void DrawList_PopTextureID(ImDrawList* self) { ImDrawList_PopTextureID(self); }
int FontAtlas_AddCustomRectRegular(ImFontAtlas* self,int width,int height) { return ImFontAtlas_AddCustomRectRegular(self,width,height); }
void FontAtlas_ClearFonts(ImFontAtlas* self) { ImFontAtlas_ClearFonts(self); }
ImGuiPlatformIO* PlatformIO_PlatformIO() { return ImGuiPlatformIO_ImGuiPlatformIO(); }
float Storage_GetFloat(ImGuiStorage* self,ImGuiID key,float default_val) { return ImGuiStorage_GetFloat(self,key,default_val); }
void BeginDisabled(bool disabled) { igBeginDisabled(disabled); }
bool InputTextMultiline(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextMultiline(label,buf,buf_size,size,flags,callback,user_data); }
void PopStyleVar(int count) { igPopStyleVar(count); }
ImGuiMouseCursor GetMouseCursor() { return igGetMouseCursor(); }
void Font_ClearOutputData(ImFont* self) { ImFont_ClearOutputData(self); }
int GetMouseClickedCount(ImGuiMouseButton button) { return igGetMouseClickedCount(button); }
void PushClipRect(const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { igPushClipRect(clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void IO_ClearInputKeys(ImGuiIO* self) { ImGuiIO_ClearInputKeys(self); }
ImGuiContext* GetCurrentContext() { return igGetCurrentContext(); }
ImGuiPlatformIO* GetPlatformIO() { return igGetPlatformIO(); }
const ImWchar* FontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesVietnamese(self); }
void SetCursorPos(const ImVec2 local_pos) { igSetCursorPos(local_pos); }
void GetWindowPos(ImVec2 *pOut) { igGetWindowPos(pOut); }
bool SliderScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalarN(label,data_type,p_data,components,p_min,p_max,format,flags); }
void TextFilter_Clear(ImGuiTextFilter* self) { ImGuiTextFilter_Clear(self); }
bool IsWindowHovered(ImGuiHoveredFlags flags) { return igIsWindowHovered(flags); }
void Separator() { igSeparator(); }
const ImWchar* FontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self); }
void GetCursorPos(ImVec2 *pOut) { igGetCursorPos(pOut); }
void LogButtons() { igLogButtons(); }
void SetColumnWidth(int column_index,float width) { igSetColumnWidth(column_index,width); }
void Color_HSV(ImColor *pOut,float h,float s,float v,float a) { ImColor_HSV(pOut,h,s,v,a); }
void OpenPopupOnItemClick(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopupOnItemClick(str_id,popup_flags); }
int TableGetColumnIndex() { return igTableGetColumnIndex(); }
void DrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness) { ImDrawList_AddRect(self,p_min,p_max,col,rounding,flags,thickness); }
bool IsItemActive() { return igIsItemActive(); }
const ImWchar* FontAtlas_GetGlyphRangesThai(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesThai(self); }
void DrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { ImDrawList_PushClipRect(self,clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
bool Payload_IsDelivery(ImGuiPayload* self) { return ImGuiPayload_IsDelivery(self); }
bool DebugCheckVersionAndDataLayout(const char* version_str,size_t sz_io,size_t sz_style,size_t sz_vec2,size_t sz_vec4,size_t sz_drawvert,size_t sz_drawidx) { return igDebugCheckVersionAndDataLayout(version_str,sz_io,sz_style,sz_vec2,sz_vec4,sz_drawvert,sz_drawidx); }
bool IsKeyDown(ImGuiKey key) { return igIsKeyDown(key); }
void SetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot) { igSetNextWindowPos(pos,cond,pivot); }
ImFont* Font_Font() { return ImFont_ImFont(); }
ImDrawData* GetDrawData() { return igGetDrawData(); }
void SaveIniSettingsToDisk(const char* ini_filename) { igSaveIniSettingsToDisk(ini_filename); }
ImDrawList* DrawList_CloneOutput(ImDrawList* self) { return ImDrawList_CloneOutput(self); }
void DrawList_PrimRectUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,ImU32 col) { ImDrawList_PrimRectUV(self,a,b,uv_a,uv_b,col); }
float GetColumnOffset(int column_index) { return igGetColumnOffset(column_index); }
float GetFrameHeightWithSpacing() { return igGetFrameHeightWithSpacing(); }
float GetTextLineHeightWithSpacing() { return igGetTextLineHeightWithSpacing(); }
bool InputFloat(const char* label,float* v,float step,float step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputFloat(label,v,step,step_fast,format,flags); }
bool ShowStyleSelector(const char* label) { return igShowStyleSelector(label); }
bool DragFloat(const char* label,float* v,float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat(label,v,v_speed,v_min,v_max,format,flags); }
void GetWindowContentRegionMin(ImVec2 *pOut) { igGetWindowContentRegionMin(pOut); }
void DrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left) { ImDrawList_AddRectFilledMultiColor(self,p_min,p_max,col_upr_left,col_upr_right,col_bot_right,col_bot_left); }
const ImWchar* FontAtlas_GetGlyphRangesKorean(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesKorean(self); }
bool DragFloatRange2(const char* label,float* v_current_min,float* v_current_max,float v_speed,float v_min,float v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragFloatRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
const char* Font_CalcWordWrapPositionA(ImFont* self,float scale,const char* text,const char* text_end,float wrap_width) { return ImFont_CalcWordWrapPositionA(self,scale,text,text_end,wrap_width); }
bool IsMouseDown(ImGuiMouseButton button) { return igIsMouseDown(button); }
ImGuiStyle* Style_Style() { return ImGuiStyle_ImGuiStyle(); }
void EndChildFrame() { igEndChildFrame(); }
void EndDragDropTarget() { igEndDragDropTarget(); }
ImGuiPlatformImeData* PlatformeData_PlatformeData() { return ImGuiPlatformImeData_ImGuiPlatformImeData(); }
void SetWindowFontScale(float scale) { igSetWindowFontScale(scale); }
void ShowMetricsWindow(bool* p_open) { igShowMetricsWindow(p_open); }
void Font_BuildLookupTable(ImFont* self) { ImFont_BuildLookupTable(self); }
void BeginTooltip() { igBeginTooltip(); }
void IO_AddMouseWheelEvent(ImGuiIO* self,float wh_x,float wh_y) { ImGuiIO_AddMouseWheelEvent(self,wh_x,wh_y); }
bool BeginPopupContextItem(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextItem(str_id,popup_flags); }
bool IsAnyMouseDown() { return igIsAnyMouseDown(); }
void PushItemWidth(float item_width) { igPushItemWidth(item_width); }
void SetNextItemOpen(bool is_open,ImGuiCond cond) { igSetNextItemOpen(is_open,cond); }
void DrawList_AddConvexPolyFilled(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col) { ImDrawList_AddConvexPolyFilled(self,points,num_points,col); }
void CloseCurrentPopup() { igCloseCurrentPopup(); }
bool IsItemDeactivated() { return igIsItemDeactivated(); }
void IO_AddMousePosEvent(ImGuiIO* self,float x,float y) { ImGuiIO_AddMousePosEvent(self,x,y); }
ImGuiStyle* GetStyle() { return igGetStyle(); }
void IO_AddMouseButtonEvent(ImGuiIO* self,int button,bool down) { ImGuiIO_AddMouseButtonEvent(self,button,down); }
ImGuiIO* IO_IO() { return ImGuiIO_ImGuiIO(); }
const char* GetClipboardText() { return igGetClipboardText(); }
void EndFrame() { igEndFrame(); }
const ImGuiPayload* GetDragDropPayload() { return igGetDragDropPayload(); }
void IO_ClearInputCharacters(ImGuiIO* self) { ImGuiIO_ClearInputCharacters(self); }
bool InputTextCallbackData_HasSelection(ImGuiInputTextCallbackData* self) { return ImGuiInputTextCallbackData_HasSelection(self); }
float CalcItemWidth() { return igCalcItemWidth(); }
bool InputFloat2(const char* label,float v[2],const char* format,ImGuiInputTextFlags flags) { return igInputFloat2(label,v,format,flags); }
void SetCursorPosY(float local_y) { igSetCursorPosY(local_y); }
bool ColorEdit3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorEdit3(label,col,flags); }
ImDrawCmd* DrawCmd_DrawCmd() { return ImDrawCmd_ImDrawCmd(); }
void IO_AddInputCharactersUTF8(ImGuiIO* self,const char* str) { ImGuiIO_AddInputCharactersUTF8(self,str); }
void EndTooltip() { igEndTooltip(); }
void ShowDebugLogWindow(bool* p_open) { igShowDebugLogWindow(p_open); }
bool Storage_GetBool(ImGuiStorage* self,ImGuiID key,bool default_val) { return ImGuiStorage_GetBool(self,key,default_val); }
ImGuiStorage* GetStateStorage() { return igGetStateStorage(); }
void SetColorEditOptions(ImGuiColorEditFlags flags) { igSetColorEditOptions(flags); }
bool FontAtlas_IsBuilt(ImFontAtlas* self) { return ImFontAtlas_IsBuilt(self); }
bool IsMouseReleased(ImGuiMouseButton button) { return igIsMouseReleased(button); }
void SetScrollHereY(float center_y_ratio) { igSetScrollHereY(center_y_ratio); }
bool InputInt3(const char* label,int v[3],ImGuiInputTextFlags flags) { return igInputInt3(label,v,flags); }
bool IsItemFocused() { return igIsItemFocused(); }
void FontAtlas_SetTexID(ImFontAtlas* self,ImTextureID id) { ImFontAtlas_SetTexID(self,id); }
bool BeginListBox(const char* label,const ImVec2 size) { return igBeginListBox(label,size); }
bool IsItemClicked(ImGuiMouseButton mouse_button) { return igIsItemClicked(mouse_button); }
void DrawList_PrimReserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimReserve(self,idx_count,vtx_count); }
void Font_SetGlyphVisible(ImFont* self,ImWchar c,bool visible) { ImFont_SetGlyphVisible(self,c,visible); }
void GetCursorScreenPos(ImVec2 *pOut) { igGetCursorScreenPos(pOut); }
float GetWindowWidth() { return igGetWindowWidth(); }
void FontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsAlpha8(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
bool BeginMenu(const char* label,bool enabled) { return igBeginMenu(label,enabled); }
void FontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self,const char* text,const char* text_end) { ImFontGlyphRangesBuilder_AddText(self,text,text_end); }
ImGuiPlatformMonitor* PlatformMonitor_PlatformMonitor() { return ImGuiPlatformMonitor_ImGuiPlatformMonitor(); }
void EndGroup() { igEndGroup(); }
void DrawList_AddageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col) { ImDrawList_AddImageQuad(self,user_texture_id,p1,p2,p3,p4,uv1,uv2,uv3,uv4,col); }
ImDrawList* GetWindowDrawList() { return igGetWindowDrawList(); }
void Dummy(const ImVec2 size) { igDummy(size); }
bool IsWindowCollapsed() { return igIsWindowCollapsed(); }
void DrawData_Clear(ImDrawData* self) { ImDrawData_Clear(self); }
void DrawList_AddCircleFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddCircleFilled(self,center,radius,col,num_segments); }
bool TextFilter_IsActive(ImGuiTextFilter* self) { return ImGuiTextFilter_IsActive(self); }
ImGuiListClipper* ListClipper_ListClipper() { return ImGuiListClipper_ImGuiListClipper(); }
void SetStateStorage(ImGuiStorage* storage) { igSetStateStorage(storage); }
void DrawList_AddCallback(ImDrawList* self,ImDrawCallback callback,void* callback_data) { ImDrawList_AddCallback(self,callback,callback_data); }
void FontAtlas_ClearTexData(ImFontAtlas* self) { ImFontAtlas_ClearTexData(self); }
void NextColumn() { igNextColumn(); }
void PopItemWidth() { igPopItemWidth(); }
bool TextFilter_Draw(ImGuiTextFilter* self,const char* label,float width) { return ImGuiTextFilter_Draw(self,label,width); }
void DrawListSplitter_Merge(ImDrawListSplitter* self,ImDrawList* draw_list) { ImDrawListSplitter_Merge(self,draw_list); }
void PopFont() { igPopFont(); }
void Storage_SetInt(ImGuiStorage* self,ImGuiID key,int val) { ImGuiStorage_SetInt(self,key,val); }
bool InputTextWithHint(const char* label,const char* hint,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextWithHint(label,hint,buf,buf_size,flags,callback,user_data); }
ImGuiViewport* FindViewportByPlatformHandle(void* platform_handle) { return igFindViewportByPlatformHandle(platform_handle); }
bool InvisibleButton(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags) { return igInvisibleButton(str_id,size,flags); }
const char* SaveIniSettingsToMemory(size_t* out_ini_size) { return igSaveIniSettingsToMemory(out_ini_size); }
void DrawList_PathFillConvex(ImDrawList* self,ImU32 col) { ImDrawList_PathFillConvex(self,col); }
void Bullet() { igBullet(); }
float GetCursorPosX() { return igGetCursorPosX(); }
const ImVec4* GetStyleColorVec4(ImGuiCol idx) { return igGetStyleColorVec4(idx); }
bool InputInt4(const char* label,int v[4],ImGuiInputTextFlags flags) { return igInputInt4(label,v,flags); }
bool IsAnyItemActive() { return igIsAnyItemActive(); }
bool InputText(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputText(label,buf,buf_size,flags,callback,user_data); }
void SetNextWindowClass(const ImGuiWindowClass* window_class) { igSetNextWindowClass(window_class); }
bool DragInt(const char* label,int* v,float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt(label,v,v_speed,v_min,v_max,format,flags); }
void ResetMouseDragDelta(ImGuiMouseButton button) { igResetMouseDragDelta(button); }
void DrawList_PathLineTo(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineTo(self,pos); }
void DrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self,ImDrawList* draw_list,int channel_idx) { ImDrawListSplitter_SetCurrentChannel(self,draw_list,channel_idx); }
void ColorConvertU32ToFloat4(ImVec4 *pOut,ImU32 in) { igColorConvertU32ToFloat4(pOut,in); }
void Font_AddGlyph(ImFont* self,const ImFontConfig* src_cfg,ImWchar c,float x0,float y0,float x1,float y1,float u0,float v0,float u1,float v1,float advance_x) { ImFont_AddGlyph(self,src_cfg,c,x0,y0,x1,y1,u0,v0,u1,v1,advance_x); }
void End() { igEnd(); }
const char* GetStyleColorName(ImGuiCol idx) { return igGetStyleColorName(idx); }
void TableNextRow(ImGuiTableRowFlags row_flags,float min_row_height) { igTableNextRow(row_flags,min_row_height); }
bool DragInt3(const char* label,int v[3],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt3(label,v,v_speed,v_min,v_max,format,flags); }
bool TableSetColumnIndex(int column_n) { return igTableSetColumnIndex(column_n); }
bool Font_IsLoaded(ImFont* self) { return ImFont_IsLoaded(self); }
bool BeginMainMenuBar() { return igBeginMainMenuBar(); }
void Font_AddRemapChar(ImFont* self,ImWchar dst,ImWchar src,bool overwrite_dst) { ImFont_AddRemapChar(self,dst,src,overwrite_dst); }
bool IsAnyItemFocused() { return igIsAnyItemFocused(); }
void LoadIniSettingsFromMemory(const char* ini_data,size_t ini_size) { igLoadIniSettingsFromMemory(ini_data,ini_size); }
void DrawList_PathLineToMergeDuplicate(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineToMergeDuplicate(self,pos); }
void LogToTTY(int auto_open_depth) { igLogToTTY(auto_open_depth); }
int TableGetColumnCount() { return igTableGetColumnCount(); }
bool BeginTabBar(const char* str_id,ImGuiTabBarFlags flags) { return igBeginTabBar(str_id,flags); }
void ShowFontSelector(const char* label) { igShowFontSelector(label); }
void TextFilter_Build(ImGuiTextFilter* self) { ImGuiTextFilter_Build(self); }
bool DragInt2(const char* label,int v[2],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt2(label,v,v_speed,v_min,v_max,format,flags); }
void GetItemRectSize(ImVec2 *pOut) { igGetItemRectSize(pOut); }
void PushAllowKeyboardFocus(bool allow_keyboard_focus) { igPushAllowKeyboardFocus(allow_keyboard_focus); }
bool Button(const char* label,const ImVec2 size) { return igButton(label,size); }
void FontAtlas_ClearInputData(ImFontAtlas* self) { ImFontAtlas_ClearInputData(self); }
void Unindent(float indent_w) { igUnindent(indent_w); }
bool IsItemDeactivatedAfterEdit() { return igIsItemDeactivatedAfterEdit(); }
void FontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsRGBA32(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
void IO_SetKeyEventNativeData(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode,int native_legacy_index) { ImGuiIO_SetKeyEventNativeData(self,key,native_keycode,native_scancode,native_legacy_index); }
bool Begin(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBegin(name,p_open,flags); }
void DrawList_PathBezierQuadraticCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,int num_segments) { ImDrawList_PathBezierQuadraticCurveTo(self,p2,p3,num_segments); }
bool BeginMenuBar() { return igBeginMenuBar(); }
int Storage_GetInt(ImGuiStorage* self,ImGuiID key,int default_val) { return ImGuiStorage_GetInt(self,key,default_val); }
void CalcTextSize(ImVec2 *pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width) { igCalcTextSize(pOut,text,text_end,hide_text_after_double_hash,wrap_width); }
void Indent(float indent_w) { igIndent(indent_w); }
void IO_SetAppAcceptingEvents(ImGuiIO* self,bool accepting_events) { ImGuiIO_SetAppAcceptingEvents(self,accepting_events); }
bool DragFloat3(const char* label,float v[3],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat3(label,v,v_speed,v_min,v_max,format,flags); }
void GetMousePosOnOpeningCurrentPopup(ImVec2 *pOut) { igGetMousePosOnOpeningCurrentPopup(pOut); }
void ShowDemoWindow(bool* p_open) { igShowDemoWindow(p_open); }
bool SliderFloat(const char* label,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat(label,v,v_min,v_max,format,flags); }
ImGuiInputTextCallbackData* InputTextCallbackData_InputTextCallbackData() { return ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(); }
bool Payload_IsDataType(ImGuiPayload* self,const char* type) { return ImGuiPayload_IsDataType(self,type); }
void BeginGroup() { igBeginGroup(); }
bool InputInt2(const char* label,int v[2],ImGuiInputTextFlags flags) { return igInputInt2(label,v,flags); }
void DrawListSplitter_Split(ImDrawListSplitter* self,ImDrawList* draw_list,int count) { ImDrawListSplitter_Split(self,draw_list,count); }
bool IsWindowDocked() { return igIsWindowDocked(); }
void SetScrollHereX(float center_x_ratio) { igSetScrollHereX(center_x_ratio); }
float GetTreeNodeToLabelSpacing() { return igGetTreeNodeToLabelSpacing(); }
int GetColumnsCount() { return igGetColumnsCount(); }
void Storage_SetVoidPtr(ImGuiStorage* self,ImGuiID key,void* val) { ImGuiStorage_SetVoidPtr(self,key,val); }
bool IsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max,bool clip) { return igIsMouseHoveringRect(r_min,r_max,clip); }
void DrawList_ChannelsMerge(ImDrawList* self) { ImDrawList_ChannelsMerge(self); }
void Viewport_GetWorkCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetWorkCenter(pOut,self); }
bool BeginPopup(const char* str_id,ImGuiWindowFlags flags) { return igBeginPopup(str_id,flags); }
void SetCurrentContext(ImGuiContext* ctx) { igSetCurrentContext(ctx); }
void DrawList_PrimWriteIdx(ImDrawList* self,ImDrawIdx idx) { ImDrawList_PrimWriteIdx(self,idx); }
void EndMenuBar() { igEndMenuBar(); }
bool InputDouble(const char* label,double* v,double step,double step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputDouble(label,v,step,step_fast,format,flags); }
void PopButtonRepeat() { igPopButtonRepeat(); }
void PopTextWrapPos() { igPopTextWrapPos(); }
void SetNextWindowBgAlpha(float alpha) { igSetNextWindowBgAlpha(alpha); }
bool FontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder* self,size_t n) { return ImFontGlyphRangesBuilder_GetBit(self,n); }
bool Payload_IsPreview(ImGuiPayload* self) { return ImGuiPayload_IsPreview(self); }
bool IsKeyPressed(ImGuiKey key,bool repeat) { return igIsKeyPressed(key,repeat); }
void DrawData_DeIndexAllBuffers(ImDrawData* self) { ImDrawData_DeIndexAllBuffers(self); }
void DrawList_AddPolyline(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_AddPolyline(self,points,num_points,col,flags,thickness); }
void Font_CalcTextSizeA(ImVec2 *pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin,const char* text_end,const char** remaining) { ImFont_CalcTextSizeA(pOut,self,size,max_width,wrap_width,text_begin,text_end,remaining); }
void Style_ScaleAllSizes(ImGuiStyle* self,float scale_factor) { ImGuiStyle_ScaleAllSizes(self,scale_factor); }
void IO_AddFocusEvent(ImGuiIO* self,bool focused) { ImGuiIO_AddFocusEvent(self,focused); }
bool DragFloat4(const char* label,float v[4],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat4(label,v,v_speed,v_min,v_max,format,flags); }
ImGuiIO* GetIO() { return igGetIO(); }
void SetNextItemWidth(float item_width) { igSetNextItemWidth(item_width); }
void Color_SetHSV(ImColor* self,float h,float s,float v,float a) { ImColor_SetHSV(self,h,s,v,a); }
void ProgressBar(float fraction,const ImVec2 size_arg,const char* overlay) { igProgressBar(fraction,size_arg,overlay); }
void GetItemRectMin(ImVec2 *pOut) { igGetItemRectMin(pOut); }
void GetContentRegionMax(ImVec2 *pOut) { igGetContentRegionMax(pOut); }
void PushButtonRepeat(bool repeat) { igPushButtonRepeat(repeat); }
bool FontAtlas_GetMouseCursorTexData(ImFontAtlas* self,ImGuiMouseCursor cursor,ImVec2* out_offset,ImVec2* out_size,ImVec2 out_uv_border[2],ImVec2 out_uv_fill[2]) { return ImFontAtlas_GetMouseCursorTexData(self,cursor,out_offset,out_size,out_uv_border,out_uv_fill); }
void FontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self,size_t n) { ImFontGlyphRangesBuilder_SetBit(self,n); }
bool IsMouseDoubleClicked(ImGuiMouseButton button) { return igIsMouseDoubleClicked(button); }
void EndTable() { igEndTable(); }
const char* GetVersion() { return igGetVersion(); }
ImTextureID DrawCmd_GetTexID(ImDrawCmd* self) { return ImDrawCmd_GetTexID(self); }
void Columns(int count,const char* id,bool border) { igColumns(count,id,border); }
