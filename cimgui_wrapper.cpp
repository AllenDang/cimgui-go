#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

float GetFrameHeight() { return igGetFrameHeight(); }
void DrawList_AddTriangle(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness) { ImDrawList_AddTriangle(self,p1,p2,p3,col,thickness); }
const ImFontGlyph* Font_FindGlyphNoFallback(ImFont* self,ImWchar c) { return ImFont_FindGlyphNoFallback(self,c); }
bool ShowStyleSelector(const char* label) { return igShowStyleSelector(label); }
void DrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness) { ImDrawList_AddRect(self,p_min,p_max,col,rounding,flags,thickness); }
ImFont* FontAtlas_AddFont(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFont(self,font_cfg); }
ImFontAtlas* FontAtlas_ImFontAtlas() { return ImFontAtlas_ImFontAtlas(); }
ImDrawListSharedData* GetDrawListSharedData() { return igGetDrawListSharedData(); }
bool IsMouseDown(ImGuiMouseButton button) { return igIsMouseDown(button); }
ImDrawCmd* DrawCmd_ImDrawCmd() { return ImDrawCmd_ImDrawCmd(); }
void RenderPlatformWindowsDefault(void* platform_render_arg,void* renderer_render_arg) { igRenderPlatformWindowsDefault(platform_render_arg,renderer_render_arg); }
ImDrawList* GetBackgroundDrawList_Nil() { return igGetBackgroundDrawList_Nil(); }
ImDrawList* GetBackgroundDrawList_ViewportPtr(ImGuiViewport* viewport) { return igGetBackgroundDrawList_ViewportPtr(viewport); }
void SetNextItemWidth(float item_width) { igSetNextItemWidth(item_width); }
ImGuiOnceUponAFrame* OnceUponAFrame_ImGuiOnceUponAFrame() { return ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(); }
void SetNextWindowCollapsed(bool collapsed,ImGuiCond cond) { igSetNextWindowCollapsed(collapsed,cond); }
bool SliderFloat2(const char* label,float v[2],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat2(label,v,v_min,v_max,format,flags); }
void FontAtlas_ClearFonts(ImFontAtlas* self) { ImFontAtlas_ClearFonts(self); }
bool ListBox_Str_arr(const char* label,int* current_item,const char* const items[],int items_count,int height_in_items) { return igListBox_Str_arr(label,current_item,items,items_count,height_in_items); }
bool ListBox_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int height_in_items) { return igListBox_FnBoolPtr(label,current_item,items_getter,data,items_count,height_in_items); }
void PopFont() { igPopFont(); }
void SetNextWindowClass(const ImGuiWindowClass* window_class) { igSetNextWindowClass(window_class); }
ImGuiViewport* GetWindowViewport() { return igGetWindowViewport(); }
void FontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self) { ImFontGlyphRangesBuilder_Clear(self); }
void IO_AddMouseWheelEvent(ImGuiIO* self,float wh_x,float wh_y) { ImGuiIO_AddMouseWheelEvent(self,wh_x,wh_y); }
float GetFontSize() { return igGetFontSize(); }
void SetDrawCursorPos(const ImVec2 local_pos) { igSetCursorPos(local_pos); }
void PopClipRect() { igPopClipRect(); }
void PopItemWidth() { igPopItemWidth(); }
bool FontAtlas_IsBuilt(ImFontAtlas* self) { return ImFontAtlas_IsBuilt(self); }
void Color_SetHSV(ImColor* self,float h,float s,float v,float a) { ImColor_SetHSV(self,h,s,v,a); }
void DrawList_AddImageRounded(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddImageRounded(self,user_texture_id,p_min,p_max,uv_min,uv_max,col,rounding,flags); }
void DrawList_PathBezierQuadraticCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,int num_segments) { ImDrawList_PathBezierQuadraticCurveTo(self,p2,p3,num_segments); }
bool IsMouseReleased(ImGuiMouseButton button) { return igIsMouseReleased(button); }
bool SliderScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalarN(label,data_type,p_data,components,p_min,p_max,format,flags); }
void StyleColorsDark(ImGuiStyle* dst) { igStyleColorsDark(dst); }
ImDrawList* DrawList_ImDrawList(const ImDrawListSharedData* shared_data) { return ImDrawList_ImDrawList(shared_data); }
const ImWchar* FontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseFull(self); }
const char* Font_CalcWordWrapPositionA(ImFont* self,float scale,const char* text,const char* text_end,float wrap_width) { return ImFont_CalcWordWrapPositionA(self,scale,text,text_end,wrap_width); }
bool Font_IsGlyphRangeUnused(ImFont* self,unsigned int c_begin,unsigned int c_last) { return ImFont_IsGlyphRangeUnused(self,c_begin,c_last); }
void IO_AddKeyEvent(ImGuiIO* self,ImGuiKey key,bool down) { ImGuiIO_AddKeyEvent(self,key,down); }
void LoadIniSettingsFromDisk(const char* ini_filename) { igLoadIniSettingsFromDisk(ini_filename); }
void DrawListSplitter_Clear(ImDrawListSplitter* self) { ImDrawListSplitter_Clear(self); }
ImGuiPlatformIO* GetPlatformIO() { return igGetPlatformIO(); }
void DrawList_GetClipRectMin(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMin(pOut,self); }
void GetWindowContentRegionMin(ImVec2 *pOut) { igGetWindowContentRegionMin(pOut); }
bool TextFilter_IsActive(ImGuiTextFilter* self) { return ImGuiTextFilter_IsActive(self); }
bool BeginMenu(const char* label,bool enabled) { return igBeginMenu(label,enabled); }
ImGuiContext* CreateContext(ImFontAtlas* shared_font_atlas) { return igCreateContext(shared_font_atlas); }
int GetKeyPressedAmount(ImGuiKey key,float repeat_delay,float rate) { return igGetKeyPressedAmount(key,repeat_delay,rate); }
ImFont* Font_ImFont() { return ImFont_ImFont(); }
void SetTabItemClosed(const char* tab_or_docked_window_label) { igSetTabItemClosed(tab_or_docked_window_label); }
void Font_BuildLookupTable(ImFont* self) { ImFont_BuildLookupTable(self); }
void BeginDisabled(bool disabled) { igBeginDisabled(disabled); }
void TableHeader(const char* label) { igTableHeader(label); }
void TreePop() { igTreePop(); }
int TableGetRowIndex() { return igTableGetRowIndex(); }
void DrawList_PushTextureID(ImDrawList* self,ImTextureID texture_id) { ImDrawList_PushTextureID(self,texture_id); }
ImFont* FontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self,compressed_font_data_base85,size_pixels,font_cfg,glyph_ranges); }
ImGuiContext* GetCurrentContext() { return igGetCurrentContext(); }
bool TabItemButton(const char* label,ImGuiTabItemFlags flags) { return igTabItemButton(label,flags); }
void DrawList_PrimWriteIdx(ImDrawList* self,ImDrawIdx idx) { ImDrawList_PrimWriteIdx(self,idx); }
ImGuiPlatformImeData* PlatformImeData_ImGuiPlatformImeData() { return ImGuiPlatformImeData_ImGuiPlatformImeData(); }
void PushItemWidth(float item_width) { igPushItemWidth(item_width); }
void InputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self,int pos,const char* text,const char* text_end) { ImGuiInputTextCallbackData_InsertChars(self,pos,text,text_end); }
bool ListClipper_Step(ImGuiListClipper* self) { return ImGuiListClipper_Step(self); }
const char* GetKeyName(ImGuiKey key) { return igGetKeyName(key); }
bool FontAtlas_Build(ImFontAtlas* self) { return ImFontAtlas_Build(self); }
ImGuiID GetID_Str(const char* str_id) { return igGetID_Str(str_id); }
ImGuiID GetID_StrStr(const char* str_id_begin,const char* str_id_end) { return igGetID_StrStr(str_id_begin,str_id_end); }
ImGuiID GetID_Ptr(const void* ptr_id) { return igGetID_Ptr(ptr_id); }
float GetTextLineHeightWithSpacing() { return igGetTextLineHeightWithSpacing(); }
float GetWindowWidth() { return igGetWindowWidth(); }
void PopStyleVar(int count) { igPopStyleVar(count); }
void DrawList_PrimWriteVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimWriteVtx(self,pos,uv,col); }
void SetNextWindowContentSize(const ImVec2 size) { igSetNextWindowContentSize(size); }
void DrawList_PopTextureID(ImDrawList* self) { ImDrawList_PopTextureID(self); }
void DrawList_PrimUnreserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimUnreserve(self,idx_count,vtx_count); }
bool SliderAngle(const char* label,float* v_rad,float v_degrees_min,float v_degrees_max,const char* format,ImGuiSliderFlags flags) { return igSliderAngle(label,v_rad,v_degrees_min,v_degrees_max,format,flags); }
void GetWindowSize(ImVec2 *pOut) { igGetWindowSize(pOut); }
bool IsRectVisible_Nil(const ImVec2 size) { return igIsRectVisible_Nil(size); }
bool IsRectVisible_Vec2(const ImVec2 rect_min,const ImVec2 rect_max) { return igIsRectVisible_Vec2(rect_min,rect_max); }
void FontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self,const ImWchar* ranges) { ImFontGlyphRangesBuilder_AddRanges(self,ranges); }
void IO_ClearInputCharacters(ImGuiIO* self) { ImGuiIO_ClearInputCharacters(self); }
bool IsAnyItemHovered() { return igIsAnyItemHovered(); }
void DebugTextEncoding(const char* text) { igDebugTextEncoding(text); }
void PushTextWrapPos(float wrap_local_pos_x) { igPushTextWrapPos(wrap_local_pos_x); }
void DrawList_AddImageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col) { ImDrawList_AddImageQuad(self,user_texture_id,p1,p2,p3,p4,uv1,uv2,uv3,uv4,col); }
float GetDrawCursorPosY() { return igGetCursorPosY(); }
bool SetDragDropPayload(const char* type,const void* data,size_t sz,ImGuiCond cond) { return igSetDragDropPayload(type,data,sz,cond); }
bool IsItemActivated() { return igIsItemActivated(); }
bool IsPopupOpen_Str(const char* str_id,ImGuiPopupFlags flags) { return igIsPopupOpen_Str(str_id,flags); }
int GetColumnsCount() { return igGetColumnsCount(); }
void DrawList_PrimVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimVtx(self,pos,uv,col); }
bool MenuItem_Bool(const char* label,const char* shortcut,bool selected,bool enabled) { return igMenuItem_Bool(label,shortcut,selected,enabled); }
bool MenuItem_BoolPtr(const char* label,const char* shortcut,bool* p_selected,bool enabled) { return igMenuItem_BoolPtr(label,shortcut,p_selected,enabled); }
ImFontAtlasCustomRect* FontAtlasCustomRect_ImFontAtlasCustomRect() { return ImFontAtlasCustomRect_ImFontAtlasCustomRect(); }
void GetMousePosOnOpeningCurrentPopup(ImVec2 *pOut) { igGetMousePosOnOpeningCurrentPopup(pOut); }
void TreePush_Str(const char* str_id) { igTreePush_Str(str_id); }
void TreePush_Ptr(const void* ptr_id) { igTreePush_Ptr(ptr_id); }
int FontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x,const ImVec2 offset) { return ImFontAtlas_AddCustomRectFontGlyph(self,font,id,width,height,advance_x,offset); }
bool FontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect* self) { return ImFontAtlasCustomRect_IsPacked(self); }
void IO_AddInputCharacter(ImGuiIO* self,unsigned int c) { ImGuiIO_AddInputCharacter(self,c); }
void EndListBox() { igEndListBox(); }
void FontAtlas_Clear(ImFontAtlas* self) { ImFontAtlas_Clear(self); }
bool DragInt4(const char* label,int v[4],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt4(label,v,v_speed,v_min,v_max,format,flags); }
void LogToFile(int auto_open_depth,const char* filename) { igLogToFile(auto_open_depth,filename); }
bool InputScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalarN(label,data_type,p_data,components,p_step,p_step_fast,format,flags); }
void Font_GrowIndex(ImFont* self,int new_size) { ImFont_GrowIndex(self,new_size); }
void DrawList_ChannelsSetCurrent(ImDrawList* self,int n) { ImDrawList_ChannelsSetCurrent(self,n); }
void SetDrawCursorPosX(float local_x) { igSetCursorPosX(local_x); }
void InputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_SelectAll(self); }
float GetFrameHeightWithSpacing() { return igGetFrameHeightWithSpacing(); }
bool InputInt4(const char* label,int v[4],ImGuiInputTextFlags flags) { return igInputInt4(label,v,flags); }
void LogButtons() { igLogButtons(); }
bool Payload_IsDataType(ImGuiPayload* self,const char* type) { return ImGuiPayload_IsDataType(self,type); }
void EndPopup() { igEndPopup(); }
bool InputTextMultiline(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextMultiline(label,buf,buf_size,size,flags,callback,user_data); }
ImDrawList* DrawList_CloneOutput(ImDrawList* self) { return ImDrawList_CloneOutput(self); }
float GetColumnWidth(int column_index) { return igGetColumnWidth(column_index); }
ImGuiID DockSpaceOverViewport(const ImGuiViewport* viewport,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpaceOverViewport(viewport,flags,window_class); }
ImDrawList* GetForegroundDrawList_Nil() { return igGetForegroundDrawList_Nil(); }
ImDrawList* GetForegroundDrawList_ViewportPtr(ImGuiViewport* viewport) { return igGetForegroundDrawList_ViewportPtr(viewport); }
void SetColorEditOptions(ImGuiColorEditFlags flags) { igSetColorEditOptions(flags); }
void SetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot) { igSetNextWindowPos(pos,cond,pivot); }
void SetScrollHereX(float center_x_ratio) { igSetScrollHereX(center_x_ratio); }
ImGuiTableColumnSortSpecs* TableColumnSortSpecs_ImGuiTableColumnSortSpecs() { return ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(); }
float GetScrollY() { return igGetScrollY(); }
bool BeginPopupContextItem(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextItem(str_id,popup_flags); }
ImGuiTextBuffer* TextBuffer_ImGuiTextBuffer() { return ImGuiTextBuffer_ImGuiTextBuffer(); }
void SetWindowCollapsed_Bool(bool collapsed,ImGuiCond cond) { igSetWindowCollapsed_Bool(collapsed,cond); }
void SetWindowCollapsed_Str(const char* name,bool collapsed,ImGuiCond cond) { igSetWindowCollapsed_Str(name,collapsed,cond); }
bool SliderFloat3(const char* label,float v[3],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat3(label,v,v_min,v_max,format,flags); }
bool ArrowButton(const char* str_id,ImGuiDir dir) { return igArrowButton(str_id,dir); }
bool ColorEdit3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorEdit3(label,col,flags); }
void SetDrawCursorPosY(float local_y) { igSetCursorPosY(local_y); }
void SetNextItemOpen(bool is_open,ImGuiCond cond) { igSetNextItemOpen(is_open,cond); }
void PopID() { igPopID(); }
void ShowUserGuide() { igShowUserGuide(); }
void FontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsRGBA32(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
const ImGuiPayload* GetDragDropPayload() { return igGetDragDropPayload(); }
int GetKeyIndex(ImGuiKey key) { return igGetKeyIndex(key); }
void ShowDebugLogWindow(bool* p_open) { igShowDebugLogWindow(p_open); }
void BeginTooltip() { igBeginTooltip(); }
void EndMenu() { igEndMenu(); }
void IO_AddMousePosEvent(ImGuiIO* self,float x,float y) { ImGuiIO_AddMousePosEvent(self,x,y); }
void ColorConvertU32ToFloat4(ImVec4 *pOut,ImU32 in) { igColorConvertU32ToFloat4(pOut,in); }
float Font_GetCharAdvance(ImFont* self,ImWchar c) { return ImFont_GetCharAdvance(self,c); }
bool InputDouble(const char* label,double* v,double step,double step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputDouble(label,v,step,step_fast,format,flags); }
bool SliderInt4(const char* label,int v[4],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt4(label,v,v_min,v_max,format,flags); }
const ImWchar* FontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesJapanese(self); }
bool BeginMenuBar() { return igBeginMenuBar(); }
void SetItemDefaultFocus() { igSetItemDefaultFocus(); }
void LogFinish() { igLogFinish(); }
void EndMenuBar() { igEndMenuBar(); }
const char* Font_GetDebugName(ImFont* self) { return ImFont_GetDebugName(self); }
void InputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_ClearSelection(self); }
bool ColorPicker3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorPicker3(label,col,flags); }
ImGuiStyle* GetStyle() { return igGetStyle(); }
void IO_AddMouseViewportEvent(ImGuiIO* self,ImGuiID id) { ImGuiIO_AddMouseViewportEvent(self,id); }
void ColorConvertRGBtoHSV(float r,float g,float b,float* out_h,float* out_s,float* out_v) { igColorConvertRGBtoHSV(r,g,b,out_h,out_s,out_v); }
void GetContentRegionMax(ImVec2 *pOut) { igGetContentRegionMax(pOut); }
void IO_ClearInputKeys(ImGuiIO* self) { ImGuiIO_ClearInputKeys(self); }
void Separator() { igSeparator(); }
void TextDisabled(const char* fmt) { igTextDisabled(fmt); }
void GetItemRectMax(ImVec2 *pOut) { igGetItemRectMax(pOut); }
void LabelText(const char* label,const char* fmt) { igLabelText(label,fmt); }
bool SliderInt3(const char* label,int v[3],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt3(label,v,v_min,v_max,format,flags); }
bool IsMousePosValid(const ImVec2* mouse_pos) { return igIsMousePosValid(mouse_pos); }
void FontAtlas_ClearTexData(ImFontAtlas* self) { ImFontAtlas_ClearTexData(self); }
void ShowFontSelector(const char* label) { igShowFontSelector(label); }
void TableSetBgColor(ImGuiTableBgTarget target,ImU32 color,int column_n) { igTableSetBgColor(target,color,column_n); }
void Dummy(const ImVec2 size) { igDummy(size); }
void SaveIniSettingsToDisk(const char* ini_filename) { igSaveIniSettingsToDisk(ini_filename); }
void DrawList_PopClipRect(ImDrawList* self) { ImDrawList_PopClipRect(self); }
bool IsItemVisible() { return igIsItemVisible(); }
void DrawList_AddTriangleFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col) { ImDrawList_AddTriangleFilled(self,p1,p2,p3,col); }
void LogToTTY(int auto_open_depth) { igLogToTTY(auto_open_depth); }
void Font_RenderChar(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,ImWchar c) { ImFont_RenderChar(self,draw_list,size,pos,col,c); }
void DrawList_PrimRectUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,ImU32 col) { ImDrawList_PrimRectUV(self,a,b,uv_a,uv_b,col); }
ImFont* FontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedTTF(self,compressed_font_data,compressed_font_size,size_pixels,font_cfg,glyph_ranges); }
bool InputText(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputText(label,buf,buf_size,flags,callback,user_data); }
void Render() { igRender(); }
void SetItemAllowOverlap() { igSetItemAllowOverlap(); }
int FontAtlas_AddCustomRectRegular(ImFontAtlas* self,int width,int height) { return ImFontAtlas_AddCustomRectRegular(self,width,height); }
float GetScrollMaxY() { return igGetScrollMaxY(); }
bool InputInt3(const char* label,int v[3],ImGuiInputTextFlags flags) { return igInputInt3(label,v,flags); }
bool SmallButton(const char* label) { return igSmallButton(label); }
void DrawData_ScaleClipRects(ImDrawData* self,const ImVec2 fb_scale) { ImDrawData_ScaleClipRects(self,fb_scale); }
void FontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self,ImVector_ImWchar* out_ranges) { ImFontGlyphRangesBuilder_BuildRanges(self,out_ranges); }
void EndDragDropTarget() { igEndDragDropTarget(); }
float GetScrollMaxX() { return igGetScrollMaxX(); }
void End() { igEnd(); }
void SetScrollY_Float(float scroll_y) { igSetScrollY_Float(scroll_y); }
void InputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self,int pos,int bytes_count) { ImGuiInputTextCallbackData_DeleteChars(self,pos,bytes_count); }
bool DragInt(const char* label,int* v,float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt(label,v,v_speed,v_min,v_max,format,flags); }
float GetWindowHeight() { return igGetWindowHeight(); }
bool IsWindowFocused(ImGuiFocusedFlags flags) { return igIsWindowFocused(flags); }
const ImWchar* FontAtlas_GetGlyphRangesThai(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesThai(self); }
void SetNextWindowSize(const ImVec2 size,ImGuiCond cond) { igSetNextWindowSize(size,cond); }
void DrawList_AddConvexPolyFilled(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col) { ImDrawList_AddConvexPolyFilled(self,points,num_points,col); }
const ImVec4* GetStyleColorVec4(ImGuiCol idx) { return igGetStyleColorVec4(idx); }
ImDrawListSplitter* DrawListSplitter_ImDrawListSplitter() { return ImDrawListSplitter_ImDrawListSplitter(); }
void FontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self,const char* text,const char* text_end) { ImFontGlyphRangesBuilder_AddText(self,text,text_end); }
float GetTextLineHeight() { return igGetTextLineHeight(); }
bool SliderFloat(const char* label,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat(label,v,v_min,v_max,format,flags); }
const ImWchar* FontAtlas_GetGlyphRangesKorean(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesKorean(self); }
void ResetMouseDragDelta(ImGuiMouseButton button) { igResetMouseDragDelta(button); }
void ListClipper_ForceDisplayRangeByIndices(ImGuiListClipper* self,int item_min,int item_max) { ImGuiListClipper_ForceDisplayRangeByIndices(self,item_min,item_max); }
void ColorConvertHSVtoRGB(float h,float s,float v,float* out_r,float* out_g,float* out_b) { igColorConvertHSVtoRGB(h,s,v,out_r,out_g,out_b); }
ImColor* Color_ImColor_Nil() { return ImColor_ImColor_Nil(); }
ImColor* Color_ImColor_Float(float r,float g,float b,float a) { return ImColor_ImColor_Float(r,g,b,a); }
ImColor* Color_ImColor_Vec4(const ImVec4 col) { return ImColor_ImColor_Vec4(col); }
ImColor* Color_ImColor_Int(int r,int g,int b,int a) { return ImColor_ImColor_Int(r,g,b,a); }
ImColor* Color_ImColor_U32(ImU32 rgba) { return ImColor_ImColor_U32(rgba); }
void DrawList_AddCallback(ImDrawList* self,ImDrawCallback callback,void* callback_data) { ImDrawList_AddCallback(self,callback,callback_data); }
float GetColumnOffset(int column_index) { return igGetColumnOffset(column_index); }
void Font_AddGlyph(ImFont* self,const ImFontConfig* src_cfg,ImWchar c,float x0,float y0,float x1,float y1,float u0,float v0,float u1,float v1,float advance_x) { ImFont_AddGlyph(self,src_cfg,c,x0,y0,x1,y1,u0,v0,u1,v1,advance_x); }
bool BeginMainMenuBar() { return igBeginMainMenuBar(); }
bool BeginTabItem(const char* label,bool* p_open,ImGuiTabItemFlags flags) { return igBeginTabItem(label,p_open,flags); }
bool InputInt(const char* label,int* v,int step,int step_fast,ImGuiInputTextFlags flags) { return igInputInt(label,v,step,step_fast,flags); }
bool IsKeyDown(ImGuiKey key) { return igIsKeyDown(key); }
bool IsMouseDoubleClicked(ImGuiMouseButton button) { return igIsMouseDoubleClicked(button); }
void DrawList_AddPolyline(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_AddPolyline(self,points,num_points,col,flags,thickness); }
ImGuiListClipper* ListClipper_ImGuiListClipper() { return ImGuiListClipper_ImGuiListClipper(); }
void DestroyPlatformWindows() { igDestroyPlatformWindows(); }
bool IsAnyItemFocused() { return igIsAnyItemFocused(); }
bool BeginCombo(const char* label,const char* preview_value,ImGuiComboFlags flags) { return igBeginCombo(label,preview_value,flags); }
bool DragFloat4(const char* label,float v[4],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat4(label,v,v_speed,v_min,v_max,format,flags); }
void EndChildFrame() { igEndChildFrame(); }
void MemFree(void* ptr) { igMemFree(ptr); }
void ShowMetricsWindow(bool* p_open) { igShowMetricsWindow(p_open); }
void DrawList_AddQuadFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col) { ImDrawList_AddQuadFilled(self,p1,p2,p3,p4,col); }
float GetScrollX() { return igGetScrollX(); }
void ProgressBar(float fraction,const ImVec2 size_arg,const char* overlay) { igProgressBar(fraction,size_arg,overlay); }
ImGuiTableSortSpecs* TableSortSpecs_ImGuiTableSortSpecs() { return ImGuiTableSortSpecs_ImGuiTableSortSpecs(); }
void Viewport_GetWorkCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetWorkCenter(pOut,self); }
bool ColorButton(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size) { return igColorButton(desc_id,col,flags,size); }
bool IsItemFocused() { return igIsItemFocused(); }
ImFontAtlasCustomRect* FontAtlas_GetCustomRectByIndex(ImFontAtlas* self,int index) { return ImFontAtlas_GetCustomRectByIndex(self,index); }
const char* SaveIniSettingsToMemory(size_t* out_ini_size) { return igSaveIniSettingsToMemory(out_ini_size); }
void SetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data) { igSetNextWindowSizeConstraints(size_min,size_max,custom_callback,custom_callback_data); }
bool BeginDragDropSource(ImGuiDragDropFlags flags) { return igBeginDragDropSource(flags); }
void SetColumnOffset(int column_index,float offset_x) { igSetColumnOffset(column_index,offset_x); }
void Value_Bool(const char* prefix,bool b) { igValue_Bool(prefix,b); }
void Value_Int(const char* prefix,int v) { igValue_Int(prefix,v); }
void Value_Uint(const char* prefix,unsigned int v) { igValue_Uint(prefix,v); }
void Value_Float(const char* prefix,float v,const char* float_format) { igValue_Float(prefix,v,float_format); }
void Font_SetGlyphVisible(ImFont* self,ImWchar c,bool visible) { ImFont_SetGlyphVisible(self,c,visible); }
int GetFrameCount() { return igGetFrameCount(); }
const char* GetVersion() { return igGetVersion(); }
bool RadioButton_Bool(const char* label,bool active) { return igRadioButton_Bool(label,active); }
bool RadioButton_IntPtr(const char* label,int* v,int v_button) { return igRadioButton_IntPtr(label,v,v_button); }
void TableSetColumnEnabled(int column_n,bool v) { igTableSetColumnEnabled(column_n,v); }
void Indent(float indent_w) { igIndent(indent_w); }
bool IsAnyItemActive() { return igIsAnyItemActive(); }
void Unindent(float indent_w) { igUnindent(indent_w); }
void ShowAboutWindow(bool* p_open) { igShowAboutWindow(p_open); }
bool TextFilter_Draw(ImGuiTextFilter* self,const char* label,float width) { return ImGuiTextFilter_Draw(self,label,width); }
bool InputFloat(const char* label,float* v,float step,float step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputFloat(label,v,step,step_fast,format,flags); }
void FontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsAlpha8(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
bool InputFloat2(const char* label,float v[2],const char* format,ImGuiInputTextFlags flags) { return igInputFloat2(label,v,format,flags); }
void ShowStackToolWindow(bool* p_open) { igShowStackToolWindow(p_open); }
bool VSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igVSliderScalar(label,size,data_type,p_data,p_min,p_max,format,flags); }
void DrawList_PathArcToFast(ImDrawList* self,const ImVec2 center,float radius,int a_min_of_12,int a_max_of_12) { ImDrawList_PathArcToFast(self,center,radius,a_min_of_12,a_max_of_12); }
void BulletText(const char* fmt) { igBulletText(fmt); }
const ImWchar* FontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesCyrillic(self); }
void StyleColorsLight(ImGuiStyle* dst) { igStyleColorsLight(dst); }
void SetWindowFocus_Nil() { igSetWindowFocus_Nil(); }
void SetWindowFocus_Str(const char* name) { igSetWindowFocus_Str(name); }
ImGuiPlatformMonitor* PlatformMonitor_ImGuiPlatformMonitor() { return ImGuiPlatformMonitor_ImGuiPlatformMonitor(); }
bool DragScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalarN(label,data_type,p_data,components,v_speed,p_min,p_max,format,flags); }
float GetDrawCursorPosX() { return igGetCursorPosX(); }
bool IsItemDeactivatedAfterEdit() { return igIsItemDeactivatedAfterEdit(); }
bool IsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max,bool clip) { return igIsMouseHoveringRect(r_min,r_max,clip); }
void DrawData_Clear(ImDrawData* self) { ImDrawData_Clear(self); }
void DrawList_AddImage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col) { ImDrawList_AddImage(self,user_texture_id,p_min,p_max,uv_min,uv_max,col); }
bool BeginChildFrame(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags) { return igBeginChildFrame(id,size,flags); }
int TableGetColumnIndex() { return igTableGetColumnIndex(); }
ImDrawData* DrawData_ImDrawData() { return ImDrawData_ImDrawData(); }
void GetDrawCursorPos(ImVec2 *pOut) { igGetCursorPos(pOut); }
void GetItemRectSize(ImVec2 *pOut) { igGetItemRectSize(pOut); }
ImFont* GetFont() { return igGetFont(); }
void PopButtonRepeat() { igPopButtonRepeat(); }
void DrawList_AddDrawCmd(ImDrawList* self) { ImDrawList_AddDrawCmd(self); }
bool VSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderInt(label,size,v,v_min,v_max,format,flags); }
ImFontConfig* FontConfig_ImFontConfig() { return ImFontConfig_ImFontConfig(); }
void SetTooltip(const char* fmt) { igSetTooltip(fmt); }
bool TreeNodeEx_Str(const char* label,ImGuiTreeNodeFlags flags) { return igTreeNodeEx_Str(label,flags); }
bool TreeNodeEx_StrStr(const char* str_id,ImGuiTreeNodeFlags flags,const char* fmt) { return igTreeNodeEx_StrStr(str_id,flags,fmt); }
bool TreeNodeEx_Ptr(const void* ptr_id,ImGuiTreeNodeFlags flags,const char* fmt) { return igTreeNodeEx_Ptr(ptr_id,flags,fmt); }
bool IsItemClicked(ImGuiMouseButton mouse_button) { return igIsItemClicked(mouse_button); }
bool IsWindowDocked() { return igIsWindowDocked(); }
void TableSetupScrollFreeze(int cols,int rows) { igTableSetupScrollFreeze(cols,rows); }
const ImWchar* FontAtlas_GetGlyphRangesDefault(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesDefault(self); }
void EndGroup() { igEndGroup(); }
void GetContentRegionAvail(ImVec2 *pOut) { igGetContentRegionAvail(pOut); }
void Columns(int count,const char* id,bool border) { igColumns(count,id,border); }
const char* GetStyleColorName(ImGuiCol idx) { return igGetStyleColorName(idx); }
bool IsItemDeactivated() { return igIsItemDeactivated(); }
bool BeginPopupModal(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBeginPopupModal(name,p_open,flags); }
float CalcItemWidth() { return igCalcItemWidth(); }
void NextColumn() { igNextColumn(); }
void OpenPopupOnItemClick(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopupOnItemClick(str_id,popup_flags); }
void EndTooltip() { igEndTooltip(); }
ImU32 ColorConvertFloat4ToU32(const ImVec4 in) { return igColorConvertFloat4ToU32(in); }
ImGuiViewport* FindViewportByPlatformHandle(void* platform_handle) { return igFindViewportByPlatformHandle(platform_handle); }
void DrawList_PathStroke(ImDrawList* self,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_PathStroke(self,col,flags,thickness); }
bool IsWindowHovered(ImGuiHoveredFlags flags) { return igIsWindowHovered(flags); }
void DrawList_AddBezierQuadratic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierQuadratic(self,p1,p2,p3,col,thickness,num_segments); }
bool DebugCheckVersionAndDataLayout(const char* version_str,size_t sz_io,size_t sz_style,size_t sz_vec2,size_t sz_vec4,size_t sz_drawvert,size_t sz_drawidx) { return igDebugCheckVersionAndDataLayout(version_str,sz_io,sz_style,sz_vec2,sz_vec4,sz_drawvert,sz_drawidx); }
int GetColumnIndex() { return igGetColumnIndex(); }
ImGuiPlatformIO* PlatformIO_ImGuiPlatformIO() { return ImGuiPlatformIO_ImGuiPlatformIO(); }
void TextFilter_Clear(ImGuiTextFilter* self) { ImGuiTextFilter_Clear(self); }
void SetDrawCursorScreenPos(const ImVec2 pos) { igSetCursorScreenPos(pos); }
bool DragFloat(const char* label,float* v,float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat(label,v,v_speed,v_min,v_max,format,flags); }
ImDrawData* GetDrawData() { return igGetDrawData(); }
void GetMousePos(ImVec2 *pOut) { igGetMousePos(pOut); }
bool IsKeyReleased(ImGuiKey key) { return igIsKeyReleased(key); }
void DrawList_PathFillConvex(ImDrawList* self,ImU32 col) { ImDrawList_PathFillConvex(self,col); }
void Font_AddRemapChar(ImFont* self,ImWchar dst,ImWchar src,bool overwrite_dst) { ImFont_AddRemapChar(self,dst,src,overwrite_dst); }
void EndChild() { igEndChild(); }
bool IsItemHovered(ImGuiHoveredFlags flags) { return igIsItemHovered(flags); }
bool IsItemToggledOpen() { return igIsItemToggledOpen(); }
const ImWchar* FontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesVietnamese(self); }
void DestroyContext(ImGuiContext* ctx) { igDestroyContext(ctx); }
void DrawList_AddCircle(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddCircle(self,center,radius,col,num_segments,thickness); }
ImFont* FontAtlas_AddFontFromFileTTF(ImFontAtlas* self,const char* filename,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromFileTTF(self,filename,size_pixels,font_cfg,glyph_ranges); }
void FontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self,ImWchar c) { ImFontGlyphRangesBuilder_AddChar(self,c); }
void Payload_Clear(ImGuiPayload* self) { ImGuiPayload_Clear(self); }
void DrawList_ChannelsMerge(ImDrawList* self) { ImDrawList_ChannelsMerge(self); }
ImGuiStyle* Style_ImGuiStyle() { return ImGuiStyle_ImGuiStyle(); }
void TableSetupColumn(const char* label,ImGuiTableColumnFlags flags,float init_width_or_weight,ImGuiID user_id) { igTableSetupColumn(label,flags,init_width_or_weight,user_id); }
void CalcTextSize(ImVec2 *pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width) { igCalcTextSize(pOut,text,text_end,hide_text_after_double_hash,wrap_width); }
bool InputScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalar(label,data_type,p_data,p_step,p_step_fast,format,flags); }
void PushAllowKeyboardFocus(bool allow_keyboard_focus) { igPushAllowKeyboardFocus(allow_keyboard_focus); }
void PushStyleColor_U32(ImGuiCol idx,ImU32 col) { igPushStyleColor_U32(idx,col); }
void PushStyleColor_Vec4(ImGuiCol idx,const ImVec4 col) { igPushStyleColor_Vec4(idx,col); }
void DrawData_DeIndexAllBuffers(ImDrawData* self) { ImDrawData_DeIndexAllBuffers(self); }
void IO_AddFocusEvent(ImGuiIO* self,bool focused) { ImGuiIO_AddFocusEvent(self,focused); }
void IO_SetAppAcceptingEvents(ImGuiIO* self,bool accepting_events) { ImGuiIO_SetAppAcceptingEvents(self,accepting_events); }
bool IsMouseDragging(ImGuiMouseButton button,float lock_threshold) { return igIsMouseDragging(button,lock_threshold); }
void EndFrame() { igEndFrame(); }
void Spacing() { igSpacing(); }
ImGuiInputTextCallbackData* InputTextCallbackData_ImGuiInputTextCallbackData() { return ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(); }
void PushFont(ImFont* font) { igPushFont(font); }
void SetNextWindowBgAlpha(float alpha) { igSetNextWindowBgAlpha(alpha); }
void DrawList_GetClipRectMax(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMax(pOut,self); }
bool CollapsingHeader_TreeNodeFlags(const char* label,ImGuiTreeNodeFlags flags) { return igCollapsingHeader_TreeNodeFlags(label,flags); }
bool CollapsingHeader_BoolPtr(const char* label,bool* p_visible,ImGuiTreeNodeFlags flags) { return igCollapsingHeader_BoolPtr(label,p_visible,flags); }
void PushStyleVar_Float(ImGuiStyleVar idx,float val) { igPushStyleVar_Float(idx,val); }
void PushStyleVar_Vec2(ImGuiStyleVar idx,const ImVec2 val) { igPushStyleVar_Vec2(idx,val); }
ImGuiTableSortSpecs* TableGetSortSpecs() { return igTableGetSortSpecs(); }
void EndTable() { igEndTable(); }
void GetMouseDragDelta(ImVec2 *pOut,ImGuiMouseButton button,float lock_threshold) { igGetMouseDragDelta(pOut,button,lock_threshold); }
void DrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self,ImDrawList* draw_list,int channel_idx) { ImDrawListSplitter_SetCurrentChannel(self,draw_list,channel_idx); }
void DrawList_PathClear(ImDrawList* self) { ImDrawList_PathClear(self); }
bool DragFloatRange2(const char* label,float* v_current_min,float* v_current_max,float v_speed,float v_min,float v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragFloatRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
bool FontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder* self,size_t n) { return ImFontGlyphRangesBuilder_GetBit(self,n); }
void EndDisabled() { igEndDisabled(); }
void GetFontTexUvWhitePixel(ImVec2 *pOut) { igGetFontTexUvWhitePixel(pOut); }
float GetTreeNodeToLabelSpacing() { return igGetTreeNodeToLabelSpacing(); }
void DrawList_PathLineTo(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineTo(self,pos); }
ImFont* FontAtlas_AddFontDefault(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFontDefault(self,font_cfg); }
bool DragFloat3(const char* label,float v[3],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat3(label,v,v_speed,v_min,v_max,format,flags); }
void EndTabBar() { igEndTabBar(); }
bool InputTextWithHint(const char* label,const char* hint,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextWithHint(label,hint,buf,buf_size,flags,callback,user_data); }
void DrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self) { ImDrawListSplitter_ClearFreeMemory(self); }
void DrawListSplitter_Merge(ImDrawListSplitter* self,ImDrawList* draw_list) { ImDrawListSplitter_Merge(self,draw_list); }
bool BeginPopupContextWindow(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextWindow(str_id,popup_flags); }
void LoadIniSettingsFromMemory(const char* ini_data,size_t ini_size) { igLoadIniSettingsFromMemory(ini_data,ini_size); }
void ListClipper_End(ImGuiListClipper* self) { ImGuiListClipper_End(self); }
bool BeginPopup(const char* str_id,ImGuiWindowFlags flags) { return igBeginPopup(str_id,flags); }
void TextUnformatted(const char* text,const char* text_end) { igTextUnformatted(text,text_end); }
void PushButtonRepeat(bool repeat) { igPushButtonRepeat(repeat); }
void SetClipboardText(const char* text) { igSetClipboardText(text); }
bool TableNextColumn() { return igTableNextColumn(); }
void TextFilter_Build(ImGuiTextFilter* self) { ImGuiTextFilter_Build(self); }
void DrawList_AddQuad(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness) { ImDrawList_AddQuad(self,p1,p2,p3,p4,col,thickness); }
int TableGetColumnCount() { return igTableGetColumnCount(); }
void FontAtlas_CalcCustomRectUV(ImFontAtlas* self,const ImFontAtlasCustomRect* rect,ImVec2* out_uv_min,ImVec2* out_uv_max) { ImFontAtlas_CalcCustomRectUV(self,rect,out_uv_min,out_uv_max); }
ImGuiViewport* Viewport_ImGuiViewport() { return ImGuiViewport_ImGuiViewport(); }
ImTextureID DrawCmd_GetTexID(ImDrawCmd* self) { return ImDrawCmd_GetTexID(self); }
void SetNextWindowViewport(ImGuiID viewport_id) { igSetNextWindowViewport(viewport_id); }
bool DragInt3(const char* label,int v[3],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt3(label,v,v_speed,v_min,v_max,format,flags); }
const ImFontGlyph* Font_FindGlyph(ImFont* self,ImWchar c) { return ImFont_FindGlyph(self,c); }
bool BeginChild_Str(const char* str_id,const ImVec2 size,bool border,ImGuiWindowFlags flags) { return igBeginChild_Str(str_id,size,border,flags); }
bool BeginChild_ID(ImGuiID id,const ImVec2 size,bool border,ImGuiWindowFlags flags) { return igBeginChild_ID(id,size,border,flags); }
bool BeginTable(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width) { return igBeginTable(str_id,column,flags,outer_size,inner_width); }
void PopTextWrapPos() { igPopTextWrapPos(); }
bool SliderInt2(const char* label,int v[2],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt2(label,v,v_min,v_max,format,flags); }
bool Font_IsLoaded(ImFont* self) { return ImFont_IsLoaded(self); }
ImFont* FontAtlas_AddFontFromMemoryTTF(ImFontAtlas* self,void* font_data,int font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryTTF(self,font_data,font_size,size_pixels,font_cfg,glyph_ranges); }
bool BeginDragDropTarget() { return igBeginDragDropTarget(); }
void DrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left) { ImDrawList_AddRectFilledMultiColor(self,p_min,p_max,col_upr_left,col_upr_right,col_bot_right,col_bot_left); }
void DrawList_PrimReserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimReserve(self,idx_count,vtx_count); }
int GetMouseClickedCount(ImGuiMouseButton button) { return igGetMouseClickedCount(button); }
bool InputFloat3(const char* label,float v[3],const char* format,ImGuiInputTextFlags flags) { return igInputFloat3(label,v,format,flags); }
void DrawList_PrimQuadUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 d,const ImVec2 uv_a,const ImVec2 uv_b,const ImVec2 uv_c,const ImVec2 uv_d,ImU32 col) { ImDrawList_PrimQuadUV(self,a,b,c,d,uv_a,uv_b,uv_c,uv_d,col); }
bool FontAtlas_GetMouseCursorTexData(ImFontAtlas* self,ImGuiMouseCursor cursor,ImVec2* out_offset,ImVec2* out_size,ImVec2 out_uv_border[2],ImVec2 out_uv_fill[2]) { return ImFontAtlas_GetMouseCursorTexData(self,cursor,out_offset,out_size,out_uv_border,out_uv_fill); }
void BeginGroup() { igBeginGroup(); }
bool DragIntRange2(const char* label,int* v_current_min,int* v_current_max,float v_speed,int v_min,int v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragIntRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
const char* TextBuffer_c_str(ImGuiTextBuffer* self) { return ImGuiTextBuffer_c_str(self); }
void SetWindowPos_Vec2(const ImVec2 pos,ImGuiCond cond) { igSetWindowPos_Vec2(pos,cond); }
void SetWindowPos_Str(const char* name,const ImVec2 pos,ImGuiCond cond) { igSetWindowPos_Str(name,pos,cond); }
void TextWrapped(const char* fmt) { igTextWrapped(fmt); }
void AlignTextToFramePadding() { igAlignTextToFramePadding(); }
bool SliderScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalar(label,data_type,p_data,p_min,p_max,format,flags); }
ImGuiID DockSpace(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpace(id,size,flags,window_class); }
void FontAtlas_SetTexID(ImFontAtlas* self,ImTextureID id) { ImFontAtlas_SetTexID(self,id); }
void SetScrollX_Float(float scroll_x) { igSetScrollX_Float(scroll_x); }
bool SliderInt(const char* label,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt(label,v,v_min,v_max,format,flags); }
ImVec2* Vec2_ImVec2_Nil() { return ImVec2_ImVec2_Nil(); }
ImVec2* Vec2_ImVec2_Float(float _x,float _y) { return ImVec2_ImVec2_Float(_x,_y); }
bool Checkbox(const char* label,bool* v) { return igCheckbox(label,v); }
bool BeginTabBar(const char* str_id,ImGuiTabBarFlags flags) { return igBeginTabBar(str_id,flags); }
bool InputFloat4(const char* label,float v[4],const char* format,ImGuiInputTextFlags flags) { return igInputFloat4(label,v,format,flags); }
bool IsMouseClicked(ImGuiMouseButton button,bool repeat) { return igIsMouseClicked(button,repeat); }
void DrawList_AddNgon(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddNgon(self,center,radius,col,num_segments,thickness); }
void CloseCurrentPopup() { igCloseCurrentPopup(); }
void SetColumnWidth(int column_index,float width) { igSetColumnWidth(column_index,width); }
void TableNextRow(ImGuiTableRowFlags row_flags,float min_row_height) { igTableNextRow(row_flags,min_row_height); }
bool TreeNode_Str(const char* label) { return igTreeNode_Str(label); }
bool TreeNode_StrStr(const char* str_id,const char* fmt) { return igTreeNode_StrStr(str_id,fmt); }
bool TreeNode_Ptr(const void* ptr_id,const char* fmt) { return igTreeNode_Ptr(ptr_id,fmt); }
void IO_AddInputCharactersUTF8(ImGuiIO* self,const char* str) { ImGuiIO_AddInputCharactersUTF8(self,str); }
bool InputTextCallbackData_HasSelection(ImGuiInputTextCallbackData* self) { return ImGuiInputTextCallbackData_HasSelection(self); }
void StyleColorsClassic(ImGuiStyle* dst) { igStyleColorsClassic(dst); }
void TextColored(const ImVec4 col,const char* fmt) { igTextColored(col,fmt); }
void UpdatePlatformWindows() { igUpdatePlatformWindows(); }
bool IsKeyPressed(ImGuiKey key,bool repeat) { return igIsKeyPressed(key,repeat); }
void DrawList_AddBezierCubic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierCubic(self,p1,p2,p3,p4,col,thickness,num_segments); }
bool BeginPopupContextVoid(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextVoid(str_id,popup_flags); }
void PlotHistogram_FloatPtr(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride) { igPlotHistogram_FloatPtr(label,values,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size,stride); }
void PlotHistogram_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size) { igPlotHistogram_FnFloatPtr(label,values_getter,data,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size); }
bool Combo_Str_arr(const char* label,int* current_item,const char* const items[],int items_count,int popup_max_height_in_items) { return igCombo_Str_arr(label,current_item,items,items_count,popup_max_height_in_items); }
bool Combo_Str(const char* label,int* current_item,const char* items_separated_by_zeros,int popup_max_height_in_items) { return igCombo_Str(label,current_item,items_separated_by_zeros,popup_max_height_in_items); }
bool Combo_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int popup_max_height_in_items) { return igCombo_FnBoolPtr(label,current_item,items_getter,data,items_count,popup_max_height_in_items); }
bool DragInt2(const char* label,int v[2],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt2(label,v,v_speed,v_min,v_max,format,flags); }
bool Selectable_Bool(const char* label,bool selected,ImGuiSelectableFlags flags,const ImVec2 size) { return igSelectable_Bool(label,selected,flags,size); }
bool Selectable_BoolPtr(const char* label,bool* p_selected,ImGuiSelectableFlags flags,const ImVec2 size) { return igSelectable_BoolPtr(label,p_selected,flags,size); }
void DrawList_PathArcTo(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments) { ImDrawList_PathArcTo(self,center,radius,a_min,a_max,num_segments); }
bool Button(const char* label,const ImVec2 size) { return igButton(label,size); }
void EndTabItem() { igEndTabItem(); }
void SetScrollFromPosX_Float(float local_x,float center_x_ratio) { igSetScrollFromPosX_Float(local_x,center_x_ratio); }
void IO_AddKeyAnalogEvent(ImGuiIO* self,ImGuiKey key,bool down,float v) { ImGuiIO_AddKeyAnalogEvent(self,key,down,v); }
void ShowStyleEditor(ImGuiStyle* ref) { igShowStyleEditor(ref); }
bool Payload_IsPreview(ImGuiPayload* self) { return ImGuiPayload_IsPreview(self); }
bool TextFilter_PassFilter(ImGuiTextFilter* self,const char* text,const char* text_end) { return ImGuiTextFilter_PassFilter(self,text,text_end); }
ImGuiID GetWindowDockID() { return igGetWindowDockID(); }
void DrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { ImDrawList_PushClipRect(self,clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void FontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self,size_t n) { ImFontGlyphRangesBuilder_SetBit(self,n); }
bool IsWindowCollapsed() { return igIsWindowCollapsed(); }
void FontAtlas_ClearInputData(ImFontAtlas* self) { ImFontAtlas_ClearInputData(self); }
void NewFrame() { igNewFrame(); }
void DrawList_PrimRect(ImDrawList* self,const ImVec2 a,const ImVec2 b,ImU32 col) { ImDrawList_PrimRect(self,a,b,col); }
ImGuiViewport* FindViewportByID(ImGuiID id) { return igFindViewportByID(id); }
void DrawList_PathRect(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max,float rounding,ImDrawFlags flags) { ImDrawList_PathRect(self,rect_min,rect_max,rounding,flags); }
void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) { igSetNextFrameWantCaptureKeyboard(want_capture_keyboard); }
double GetTime() { return igGetTime(); }
ImGuiIO* GetIO() { return igGetIO(); }
void GetItemRectMin(ImVec2 *pOut) { igGetItemRectMin(pOut); }
void PushID_Str(const char* str_id) { igPushID_Str(str_id); }
void PushID_StrStr(const char* str_id_begin,const char* str_id_end) { igPushID_StrStr(str_id_begin,str_id_end); }
void PushID_Ptr(const void* ptr_id) { igPushID_Ptr(ptr_id); }
void PushID_Int(int int_id) { igPushID_Int(int_id); }
ImGuiTextFilter* TextFilter_ImGuiTextFilter(const char* default_filter) { return ImGuiTextFilter_ImGuiTextFilter(default_filter); }
const char* GetClipboardText() { return igGetClipboardText(); }
void IO_SetKeyEventNativeData(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode,int native_legacy_index) { ImGuiIO_SetKeyEventNativeData(self,key,native_keycode,native_scancode,native_legacy_index); }
ImFontGlyphRangesBuilder* FontGlyphRangesBuilder_ImFontGlyphRangesBuilder() { return ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(); }
bool SliderFloat4(const char* label,float v[4],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat4(label,v,v_min,v_max,format,flags); }
const char* TableGetColumnName_Int(int column_n) { return igTableGetColumnName_Int(column_n); }
void PlotLines_FloatPtr(const char* label,const float* values,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size,int stride) { igPlotLines_FloatPtr(label,values,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size,stride); }
void PlotLines_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size) { igPlotLines_FnFloatPtr(label,values_getter,data,values_count,values_offset,overlay_text,scale_min,scale_max,graph_size); }
bool InvisibleButton(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags) { return igInvisibleButton(str_id,size,flags); }
void SameLine(float offset_from_start_x,float spacing) { igSameLine(offset_from_start_x,spacing); }
void DrawList_PushClipRectFullScreen(ImDrawList* self) { ImDrawList_PushClipRectFullScreen(self); }
bool Begin(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBegin(name,p_open,flags); }
void SetWindowSize_Vec2(const ImVec2 size,ImGuiCond cond) { igSetWindowSize_Vec2(size,cond); }
void SetWindowSize_Str(const char* name,const ImVec2 size,ImGuiCond cond) { igSetWindowSize_Str(name,size,cond); }
ImVec4* Vec4_ImVec4_Nil() { return ImVec4_ImVec4_Nil(); }
ImVec4* Vec4_ImVec4_Float(float _x,float _y,float _z,float _w) { return ImVec4_ImVec4_Float(_x,_y,_z,_w); }
void DrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddRectFilled(self,p_min,p_max,col,rounding,flags); }
void Viewport_GetCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetCenter(pOut,self); }
void EndDragDropSource() { igEndDragDropSource(); }
void NewLine() { igNewLine(); }
void TableHeadersRow() { igTableHeadersRow(); }
void Color_HSV(ImColor *pOut,float h,float s,float v,float a) { ImColor_HSV(pOut,h,s,v,a); }
void Font_CalcTextSizeA(ImVec2 *pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin,const char* text_end,const char** remaining) { ImFont_CalcTextSizeA(pOut,self,size,max_width,wrap_width,text_begin,text_end,remaining); }
void SetNextWindowDockID(ImGuiID dock_id,ImGuiCond cond) { igSetNextWindowDockID(dock_id,cond); }
ImGuiMouseCursor GetMouseCursor() { return igGetMouseCursor(); }
void IO_AddInputCharacterUTF16(ImGuiIO* self,ImWchar16 c) { ImGuiIO_AddInputCharacterUTF16(self,c); }
void IO_AddMouseButtonEvent(ImGuiIO* self,int button,bool down) { ImGuiIO_AddMouseButtonEvent(self,button,down); }
const ImGuiPayload* AcceptDragDropPayload(const char* type,ImGuiDragDropFlags flags) { return igAcceptDragDropPayload(type,flags); }
void Bullet() { igBullet(); }
bool CheckboxFlags_IntPtr(const char* label,int* flags,int flags_value) { return igCheckboxFlags_IntPtr(label,flags,flags_value); }
bool CheckboxFlags_UintPtr(const char* label,unsigned int* flags,unsigned int flags_value) { return igCheckboxFlags_UintPtr(label,flags,flags_value); }
bool InputInt2(const char* label,int v[2],ImGuiInputTextFlags flags) { return igInputInt2(label,v,flags); }
bool IsItemActive() { return igIsItemActive(); }
const ImWchar* FontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self); }
void EndCombo() { igEndCombo(); }
void SetWindowFontScale(float scale) { igSetWindowFontScale(scale); }
void GetWindowPos(ImVec2 *pOut) { igGetWindowPos(pOut); }
void SetCurrentContext(ImGuiContext* ctx) { igSetCurrentContext(ctx); }
bool TableSetColumnIndex(int column_n) { return igTableSetColumnIndex(column_n); }
ImDrawList* GetWindowDrawList() { return igGetWindowDrawList(); }
bool IsItemEdited() { return igIsItemEdited(); }
void DrawList_AddText_Vec2(ImDrawList* self,const ImVec2 pos,ImU32 col,const char* text_begin,const char* text_end) { ImDrawList_AddText_Vec2(self,pos,col,text_begin,text_end); }
void DrawList_AddText_FontPtr(ImDrawList* self,const ImFont* font,float font_size,const ImVec2 pos,ImU32 col,const char* text_begin,const char* text_end,float wrap_width,const ImVec4* cpu_fine_clip_rect) { ImDrawList_AddText_FontPtr(self,font,font_size,pos,col,text_begin,text_end,wrap_width,cpu_fine_clip_rect); }
bool ColorPicker4(const char* label,float col[4],ImGuiColorEditFlags flags,const float* ref_col) { return igColorPicker4(label,col,flags,ref_col); }
ImGuiTableColumnFlags TableGetColumnFlags(int column_n) { return igTableGetColumnFlags(column_n); }
void Style_ScaleAllSizes(ImGuiStyle* self,float scale_factor) { ImGuiStyle_ScaleAllSizes(self,scale_factor); }
void GetDrawCursorStartPos(ImVec2 *pOut) { igGetCursorStartPos(pOut); }
void DrawListSplitter_Split(ImDrawListSplitter* self,ImDrawList* draw_list,int count) { ImDrawListSplitter_Split(self,draw_list,count); }
void DrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness) { ImDrawList_AddLine(self,p1,p2,col,thickness); }
void EndMainMenuBar() { igEndMainMenuBar(); }
bool ColorEdit4(const char* label,float col[4],ImGuiColorEditFlags flags) { return igColorEdit4(label,col,flags); }
ImGuiViewport* GetMainViewport() { return igGetMainViewport(); }
void SetNextWindowFocus() { igSetNextWindowFocus(); }
void LogToClipboard(int auto_open_depth) { igLogToClipboard(auto_open_depth); }
void GetWindowContentRegionMax(ImVec2 *pOut) { igGetWindowContentRegionMax(pOut); }
void DrawList_AddNgonFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddNgonFilled(self,center,radius,col,num_segments); }
void SetScrollHereY(float center_y_ratio) { igSetScrollHereY(center_y_ratio); }
ImGuiPayload* Payload_ImGuiPayload() { return ImGuiPayload_ImGuiPayload(); }
bool BeginListBox(const char* label,const ImVec2 size) { return igBeginListBox(label,size); }
bool DragFloat2(const char* label,float v[2],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat2(label,v,v_speed,v_min,v_max,format,flags); }
ImU32 GetColorU32_Col(ImGuiCol idx,float alpha_mul) { return igGetColorU32_Col(idx,alpha_mul); }
ImU32 GetColorU32_Vec4(const ImVec4 col) { return igGetColorU32_Vec4(col); }
ImU32 GetColorU32_U32(ImU32 col) { return igGetColorU32_U32(col); }
void SetNextFrameWantCaptureMouse(bool want_capture_mouse) { igSetNextFrameWantCaptureMouse(want_capture_mouse); }
bool Payload_IsDelivery(ImGuiPayload* self) { return ImGuiPayload_IsDelivery(self); }
void DrawList_PathBezierCubicCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,int num_segments) { ImDrawList_PathBezierCubicCurveTo(self,p2,p3,p4,num_segments); }
void ShowDemoWindow(bool* p_open) { igShowDemoWindow(p_open); }
bool VSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderFloat(label,size,v,v_min,v_max,format,flags); }
void DrawList_AddCircleFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddCircleFilled(self,center,radius,col,num_segments); }
void DrawList_PathLineToMergeDuplicate(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineToMergeDuplicate(self,pos); }
void ListClipper_Begin(ImGuiListClipper* self,int items_count,float items_height) { ImGuiListClipper_Begin(self,items_count,items_height); }
ImGuiWindowClass* WindowClass_ImGuiWindowClass() { return ImGuiWindowClass_ImGuiWindowClass(); }
void OpenPopup_Str(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopup_Str(str_id,popup_flags); }
void OpenPopup_ID(ImGuiID id,ImGuiPopupFlags popup_flags) { igOpenPopup_ID(id,popup_flags); }
void PopAllowKeyboardFocus() { igPopAllowKeyboardFocus(); }
ImGuiIO* IO_ImGuiIO() { return ImGuiIO_ImGuiIO(); }
bool IsWindowAppearing() { return igIsWindowAppearing(); }
void PopStyleColor(int count) { igPopStyleColor(count); }
void Font_RenderText(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin,const char* text_end,float wrap_width,bool cpu_fine_clip) { ImFont_RenderText(self,draw_list,size,pos,col,clip_rect,text_begin,text_end,wrap_width,cpu_fine_clip); }
bool DragScalar(const char* label,ImGuiDataType data_type,void* p_data,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalar(label,data_type,p_data,v_speed,p_min,p_max,format,flags); }
void* MemAlloc(size_t size) { return igMemAlloc(size); }
void SetScrollFromPosY_Float(float local_y,float center_y_ratio) { igSetScrollFromPosY_Float(local_y,center_y_ratio); }
void PushClipRect(const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { igPushClipRect(clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void Text(const char* fmt) { igText(fmt); }
void SetMouseCursor(ImGuiMouseCursor cursor_type) { igSetMouseCursor(cursor_type); }
void GetDrawCursorScreenPos(ImVec2 *pOut) { igGetCursorScreenPos(pOut); }
void LogText(const char* fmt) { igLogText(fmt); }
void SetKeyboardFocusHere(int offset) { igSetKeyboardFocusHere(offset); }
void DrawList_ChannelsSplit(ImDrawList* self,int count) { ImDrawList_ChannelsSplit(self,count); }
void Font_ClearOutputData(ImFont* self) { ImFont_ClearOutputData(self); }
float GetWindowDpiScale() { return igGetWindowDpiScale(); }
bool IsAnyMouseDown() { return igIsAnyMouseDown(); }
