#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

bool DragScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalarN(label,data_type,p_data,components,v_speed,p_min,p_max,format,flags); }
float GetWindowDpiScale() { return igGetWindowDpiScale(); }
void DrawList_AddageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col) { ImDrawList_AddImageQuad(self,user_texture_id,p1,p2,p3,p4,uv1,uv2,uv3,uv4,col); }
void DrawList_PushTextureID(ImDrawList* self,ImTextureID texture_id) { ImDrawList_PushTextureID(self,texture_id); }
float GetCursorPosX() { return igGetCursorPosX(); }
bool Payload_IsDataType(ImGuiPayload* self,const char* type) { return ImGuiPayload_IsDataType(self,type); }
void DrawList_PopClipRect(ImDrawList* self) { ImDrawList_PopClipRect(self); }
bool BeginChildFrame(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags) { return igBeginChildFrame(id,size,flags); }
int GetColumnIndex() { return igGetColumnIndex(); }
void DrawData_ScaleClipRects(ImDrawData* self,const ImVec2 fb_scale) { ImDrawData_ScaleClipRects(self,fb_scale); }
void DrawList_AddQuad(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness) { ImDrawList_AddQuad(self,p1,p2,p3,p4,col,thickness); }
void Font_AddGlyph(ImFont* self,const ImFontConfig* src_cfg,ImWchar c,float x0,float y0,float x1,float y1,float u0,float v0,float u1,float v1,float advance_x) { ImFont_AddGlyph(self,src_cfg,c,x0,y0,x1,y1,u0,v0,u1,v1,advance_x); }
bool BeginTable(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width) { return igBeginTable(str_id,column,flags,outer_size,inner_width); }
void ColorConvertHSVtoRGB(float h,float s,float v,float* out_r,float* out_g,float* out_b) { igColorConvertHSVtoRGB(h,s,v,out_r,out_g,out_b); }
float GetFrameHeightWithSpacing() { return igGetFrameHeightWithSpacing(); }
bool IsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max,bool clip) { return igIsMouseHoveringRect(r_min,r_max,clip); }
bool IsMouseDragging(ImGuiMouseButton button,float lock_threshold) { return igIsMouseDragging(button,lock_threshold); }
ImGuiListClipper* ListClipper_ListClipper() { return ImGuiListClipper_ImGuiListClipper(); }
bool VSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderInt(label,size,v,v_min,v_max,format,flags); }
ImFont* FontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedTTF(self,compressed_font_data,compressed_font_size,size_pixels,font_cfg,glyph_ranges); }
const ImWchar* FontAtlas_GetGlyphRangesDefault(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesDefault(self); }
bool ColorButton(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size) { return igColorButton(desc_id,col,flags,size); }
void PopButtonRepeat() { igPopButtonRepeat(); }
ImFontAtlasCustomRect* FontAtlas_GetCustomRectByIndex(ImFontAtlas* self,int index) { return ImFontAtlas_GetCustomRectByIndex(self,index); }
ImGuiPlatformIO* PlatformIO_PlatformIO() { return ImGuiPlatformIO_ImGuiPlatformIO(); }
float GetScrollMaxY() { return igGetScrollMaxY(); }
void DrawList_AddTriangle(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness) { ImDrawList_AddTriangle(self,p1,p2,p3,col,thickness); }
void PushClipRect(const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { igPushClipRect(clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
bool TextFilter_IsActive(ImGuiTextFilter* self) { return ImGuiTextFilter_IsActive(self); }
void EndDisabled() { igEndDisabled(); }
bool IsWindowCollapsed() { return igIsWindowCollapsed(); }
ImFont* FontAtlas_AddFontFromFileTTF(ImFontAtlas* self,const char* filename,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromFileTTF(self,filename,size_pixels,font_cfg,glyph_ranges); }
void DebugTextEncoding(const char* text) { igDebugTextEncoding(text); }
void LogToClipboard(int auto_open_depth) { igLogToClipboard(auto_open_depth); }
void DrawList_GetClipRectMax(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMax(pOut,self); }
bool InputTextCallbackData_HasSelection(ImGuiInputTextCallbackData* self) { return ImGuiInputTextCallbackData_HasSelection(self); }
void Storage_Clear(ImGuiStorage* self) { ImGuiStorage_Clear(self); }
bool IsKeyDown(ImGuiKey key) { return igIsKeyDown(key); }
void DrawList_AddTriangleFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col) { ImDrawList_AddTriangleFilled(self,p1,p2,p3,col); }
void Style_ScaleAllSizes(ImGuiStyle* self,float scale_factor) { ImGuiStyle_ScaleAllSizes(self,scale_factor); }
bool IsItemFocused() { return igIsItemFocused(); }
ImGuiTableSortSpecs* TableGetSortSpecs() { return igTableGetSortSpecs(); }
void DrawList_PathArcTo(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments) { ImDrawList_PathArcTo(self,center,radius,a_min,a_max,num_segments); }
void Storage_SetBool(ImGuiStorage* self,ImGuiID key,bool val) { ImGuiStorage_SetBool(self,key,val); }
bool ColorPicker3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorPicker3(label,col,flags); }
bool DragFloat4(const char* label,float v[4],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat4(label,v,v_speed,v_min,v_max,format,flags); }
void BeginTooltip() { igBeginTooltip(); }
bool IsAnyItemFocused() { return igIsAnyItemFocused(); }
ImGuiInputTextCallbackData* InputTextCallbackData_InputTextCallbackData() { return ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(); }
const char* GetKeyName(ImGuiKey key) { return igGetKeyName(key); }
void PopTextWrapPos() { igPopTextWrapPos(); }
void TableSetColumnEnabled(int column_n,bool v) { igTableSetColumnEnabled(column_n,v); }
const ImFontGlyph* Font_FindGlyph(ImFont* self,ImWchar c) { return ImFont_FindGlyph(self,c); }
void TextFilter_Clear(ImGuiTextFilter* self) { ImGuiTextFilter_Clear(self); }
float GetColumnWidth(int column_index) { return igGetColumnWidth(column_index); }
ImDrawListSharedData* GetDrawListSharedData() { return igGetDrawListSharedData(); }
void IO_AddMousePosEvent(ImGuiIO* self,float x,float y) { ImGuiIO_AddMousePosEvent(self,x,y); }
void EndDragDropSource() { igEndDragDropSource(); }
void ShowUserGuide() { igShowUserGuide(); }
void IO_AddInputCharacterUTF16(ImGuiIO* self,ImWchar16 c) { ImGuiIO_AddInputCharacterUTF16(self,c); }
void GetWindowPos(ImVec2 *pOut) { igGetWindowPos(pOut); }
void DrawList_PathClear(ImDrawList* self) { ImDrawList_PathClear(self); }
void ShowStackToolWindow(bool* p_open) { igShowStackToolWindow(p_open); }
bool SliderInt(const char* label,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt(label,v,v_min,v_max,format,flags); }
void Render() { igRender(); }
void DrawList_PathBezierCubicCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,int num_segments) { ImDrawList_PathBezierCubicCurveTo(self,p2,p3,p4,num_segments); }
void TreePop() { igTreePop(); }
void ColorConvertU32ToFloat4(ImVec4 *pOut,ImU32 in) { igColorConvertU32ToFloat4(pOut,in); }
void SetCurrentContext(ImGuiContext* ctx) { igSetCurrentContext(ctx); }
void DrawListSplitter_Merge(ImDrawListSplitter* self,ImDrawList* draw_list) { ImDrawListSplitter_Merge(self,draw_list); }
bool BeginPopup(const char* str_id,ImGuiWindowFlags flags) { return igBeginPopup(str_id,flags); }
bool SliderInt4(const char* label,int v[4],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt4(label,v,v_min,v_max,format,flags); }
void DrawList_PrimReserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimReserve(self,idx_count,vtx_count); }
void SetCursorScreenPos(const ImVec2 pos) { igSetCursorScreenPos(pos); }
bool TextFilter_Draw(ImGuiTextFilter* self,const char* label,float width) { return ImGuiTextFilter_Draw(self,label,width); }
void Font_RenderText(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin,const char* text_end,float wrap_width,bool cpu_fine_clip) { ImFont_RenderText(self,draw_list,size,pos,col,clip_rect,text_begin,text_end,wrap_width,cpu_fine_clip); }
bool IsWindowAppearing() { return igIsWindowAppearing(); }
void SetNextWindowViewport(ImGuiID viewport_id) { igSetNextWindowViewport(viewport_id); }
void DrawList_AddConvexPolyFilled(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col) { ImDrawList_AddConvexPolyFilled(self,points,num_points,col); }
void Bullet() { igBullet(); }
ImFont* FontAtlas_AddFontFromMemoryTTF(ImFontAtlas* self,void* font_data,int font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryTTF(self,font_data,font_size,size_pixels,font_cfg,glyph_ranges); }
float Storage_GetFloat(ImGuiStorage* self,ImGuiID key,float default_val) { return ImGuiStorage_GetFloat(self,key,default_val); }
void FontAtlas_ClearInputData(ImFontAtlas* self) { ImFontAtlas_ClearInputData(self); }
void TableSetupScrollFreeze(int cols,int rows) { igTableSetupScrollFreeze(cols,rows); }
void FontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsAlpha8(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
bool Storage_GetBool(ImGuiStorage* self,ImGuiID key,bool default_val) { return ImGuiStorage_GetBool(self,key,default_val); }
void* Storage_GetVoidPtr(ImGuiStorage* self,ImGuiID key) { return ImGuiStorage_GetVoidPtr(self,key); }
bool InputScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalarN(label,data_type,p_data,components,p_step,p_step_fast,format,flags); }
void TableSetBgColor(ImGuiTableBgTarget target,ImU32 color,int column_n) { igTableSetBgColor(target,color,column_n); }
ImGuiStorage* GetStateStorage() { return igGetStateStorage(); }
const char* SaveIniSettingsToMemory(size_t* out_ini_size) { return igSaveIniSettingsToMemory(out_ini_size); }
void DrawListSplitter_Split(ImDrawListSplitter* self,ImDrawList* draw_list,int count) { ImDrawListSplitter_Split(self,draw_list,count); }
void DrawList_PrimWriteIdx(ImDrawList* self,ImDrawIdx idx) { ImDrawList_PrimWriteIdx(self,idx); }
void Indent(float indent_w) { igIndent(indent_w); }
void EndListBox() { igEndListBox(); }
void GetCursorScreenPos(ImVec2 *pOut) { igGetCursorScreenPos(pOut); }
bool InputInt2(const char* label,int v[2],ImGuiInputTextFlags flags) { return igInputInt2(label,v,flags); }
void SetItemAllowOverlap() { igSetItemAllowOverlap(); }
void IO_AddFocusEvent(ImGuiIO* self,bool focused) { ImGuiIO_AddFocusEvent(self,focused); }
bool InputInt(const char* label,int* v,int step,int step_fast,ImGuiInputTextFlags flags) { return igInputInt(label,v,step,step_fast,flags); }
void SetStateStorage(ImGuiStorage* storage) { igSetStateStorage(storage); }
void DrawList_AddageRounded(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddImageRounded(self,user_texture_id,p_min,p_max,uv_min,uv_max,col,rounding,flags); }
void Storage_SetAllInt(ImGuiStorage* self,int val) { ImGuiStorage_SetAllInt(self,val); }
bool InputInt3(const char* label,int v[3],ImGuiInputTextFlags flags) { return igInputInt3(label,v,flags); }
void SetTabItemClosed(const char* tab_or_docked_window_label) { igSetTabItemClosed(tab_or_docked_window_label); }
bool FontAtlas_Build(ImFontAtlas* self) { return ImFontAtlas_Build(self); }
int FontAtlas_AddCustomRectRegular(ImFontAtlas* self,int width,int height) { return ImFontAtlas_AddCustomRectRegular(self,width,height); }
bool Font_IsGlyphRangeUnused(ImFont* self,unsigned int c_begin,unsigned int c_last) { return ImFont_IsGlyphRangeUnused(self,c_begin,c_last); }
bool SliderFloat(const char* label,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat(label,v,v_min,v_max,format,flags); }
void SetNextWindowClass(const ImGuiWindowClass* window_class) { igSetNextWindowClass(window_class); }
ImGuiViewport* GetMainViewport() { return igGetMainViewport(); }
bool SetDragDropPayload(const char* type,const void* data,size_t sz,ImGuiCond cond) { return igSetDragDropPayload(type,data,sz,cond); }
bool BeginMenuBar() { return igBeginMenuBar(); }
bool BeginTabItem(const char* label,bool* p_open,ImGuiTabItemFlags flags) { return igBeginTabItem(label,p_open,flags); }
bool SliderInt3(const char* label,int v[3],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt3(label,v,v_min,v_max,format,flags); }
void DrawList_PathFillConvex(ImDrawList* self,ImU32 col) { ImDrawList_PathFillConvex(self,col); }
bool* Storage_GetBoolRef(ImGuiStorage* self,ImGuiID key,bool default_val) { return ImGuiStorage_GetBoolRef(self,key,default_val); }
void DrawData_Clear(ImDrawData* self) { ImDrawData_Clear(self); }
bool IsItemEdited() { return igIsItemEdited(); }
bool FontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect* self) { return ImFontAtlasCustomRect_IsPacked(self); }
void IO_SetAppAcceptingEvents(ImGuiIO* self,bool accepting_events) { ImGuiIO_SetAppAcceptingEvents(self,accepting_events); }
bool BeginMainMenuBar() { return igBeginMainMenuBar(); }
void SetCursorPosY(float local_y) { igSetCursorPosY(local_y); }
const ImGuiPayload* GetDragDropPayload() { return igGetDragDropPayload(); }
void PopStyleColor(int count) { igPopStyleColor(count); }
void DrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left) { ImDrawList_AddRectFilledMultiColor(self,p_min,p_max,col_upr_left,col_upr_right,col_bot_right,col_bot_left); }
void TableHeadersRow() { igTableHeadersRow(); }
void IO_ClearInputCharacters(ImGuiIO* self) { ImGuiIO_ClearInputCharacters(self); }
int GetColumnsCount() { return igGetColumnsCount(); }
void ListClipper_ForceDisplayRangeByIndices(ImGuiListClipper* self,int item_min,int item_max) { ImGuiListClipper_ForceDisplayRangeByIndices(self,item_min,item_max); }
void GetMousePos(ImVec2 *pOut) { igGetMousePos(pOut); }
bool SliderAngle(const char* label,float* v_rad,float v_degrees_min,float v_degrees_max,const char* format,ImGuiSliderFlags flags) { return igSliderAngle(label,v_rad,v_degrees_min,v_degrees_max,format,flags); }
void DrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddRectFilled(self,p_min,p_max,col,rounding,flags); }
void DrawList_PrimVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimVtx(self,pos,uv,col); }
bool BeginPopupModal(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBeginPopupModal(name,p_open,flags); }
bool InputTextMultiline(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextMultiline(label,buf,buf_size,size,flags,callback,user_data); }
bool VSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igVSliderScalar(label,size,data_type,p_data,p_min,p_max,format,flags); }
ImGuiPlatformImeData* PlatformeData_PlatformeData() { return ImGuiPlatformImeData_ImGuiPlatformImeData(); }
float GetWindowWidth() { return igGetWindowWidth(); }
void PushButtonRepeat(bool repeat) { igPushButtonRepeat(repeat); }
ImGuiViewport* Viewport_Viewport() { return ImGuiViewport_ImGuiViewport(); }
void TextFilter_Build(ImGuiTextFilter* self) { ImGuiTextFilter_Build(self); }
bool IsItemClicked(ImGuiMouseButton mouse_button) { return igIsItemClicked(mouse_button); }
ImDrawCmd* DrawCmd_DrawCmd() { return ImDrawCmd_ImDrawCmd(); }
void GetWindowContentRegionMin(ImVec2 *pOut) { igGetWindowContentRegionMin(pOut); }
void DrawList_AddCallback(ImDrawList* self,ImDrawCallback callback,void* callback_data) { ImDrawList_AddCallback(self,callback,callback_data); }
const ImWchar* FontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self); }
float CalcItemWidth() { return igCalcItemWidth(); }
ImU32 ColorConvertFloat4ToU32(const ImVec4 in) { return igColorConvertFloat4ToU32(in); }
bool Begin(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBegin(name,p_open,flags); }
ImGuiID DockSpaceOverViewport(const ImGuiViewport* viewport,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpaceOverViewport(viewport,flags,window_class); }
void LogToTTY(int auto_open_depth) { igLogToTTY(auto_open_depth); }
bool SliderInt2(const char* label,int v[2],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt2(label,v,v_min,v_max,format,flags); }
ImGuiTableColumnSortSpecs* TableColumnSortSpecs_TableColumnSortSpecs() { return ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(); }
void* MemAlloc(size_t size) { return igMemAlloc(size); }
bool IsItemHovered(ImGuiHoveredFlags flags) { return igIsItemHovered(flags); }
void DrawList_AddDrawCmd(ImDrawList* self) { ImDrawList_AddDrawCmd(self); }
void DrawList_PathBezierQuadraticCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,int num_segments) { ImDrawList_PathBezierQuadraticCurveTo(self,p2,p3,num_segments); }
void FontAtlas_ClearTexData(ImFontAtlas* self) { ImFontAtlas_ClearTexData(self); }
bool InputScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalar(label,data_type,p_data,p_step,p_step_fast,format,flags); }
bool InvisibleButton(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags) { return igInvisibleButton(str_id,size,flags); }
void SetColorEditOptions(ImGuiColorEditFlags flags) { igSetColorEditOptions(flags); }
void Storage_SetFloat(ImGuiStorage* self,ImGuiID key,float val) { ImGuiStorage_SetFloat(self,key,val); }
bool DragFloatRange2(const char* label,float* v_current_min,float* v_current_max,float v_speed,float v_min,float v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragFloatRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
void EndTabBar() { igEndTabBar(); }
ImGuiTextFilter* TextFilter_TextFilter(const char* default_filter) { return ImGuiTextFilter_ImGuiTextFilter(default_filter); }
const ImWchar* FontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesJapanese(self); }
bool IsWindowDocked() { return igIsWindowDocked(); }
void FontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsRGBA32(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
int GetKeyPressedAmount(ImGuiKey key,float repeat_delay,float rate) { return igGetKeyPressedAmount(key,repeat_delay,rate); }
void DrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self,ImDrawList* draw_list,int channel_idx) { ImDrawListSplitter_SetCurrentChannel(self,draw_list,channel_idx); }
ImGuiMouseCursor GetMouseCursor() { return igGetMouseCursor(); }
float GetScrollMaxX() { return igGetScrollMaxX(); }
ImGuiViewport* GetWindowViewport() { return igGetWindowViewport(); }
void SetNextItemWidth(float item_width) { igSetNextItemWidth(item_width); }
void DrawList_PushClipRectFullScreen(ImDrawList* self) { ImDrawList_PushClipRectFullScreen(self); }
void TextUnformatted(const char* text,const char* text_end) { igTextUnformatted(text,text_end); }
void FontAtlas_Clear(ImFontAtlas* self) { ImFontAtlas_Clear(self); }
bool FontAtlas_IsBuilt(ImFontAtlas* self) { return ImFontAtlas_IsBuilt(self); }
bool IsAnyItemActive() { return igIsAnyItemActive(); }
void NewLine() { igNewLine(); }
bool ShowStyleSelector(const char* label) { return igShowStyleSelector(label); }
bool IsItemToggledOpen() { return igIsItemToggledOpen(); }
void PushItemWidth(float item_width) { igPushItemWidth(item_width); }
void EndMenuBar() { igEndMenuBar(); }
float GetTextLineHeightWithSpacing() { return igGetTextLineHeightWithSpacing(); }
bool IsWindowHovered(ImGuiHoveredFlags flags) { return igIsWindowHovered(flags); }
void IO_SetKeyEventNativeData(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode,int native_legacy_index) { ImGuiIO_SetKeyEventNativeData(self,key,native_keycode,native_scancode,native_legacy_index); }
void EndTable() { igEndTable(); }
void PopID() { igPopID(); }
void Color_HSV(ImColor *pOut,float h,float s,float v,float a) { ImColor_HSV(pOut,h,s,v,a); }
const ImFontGlyph* Font_FindGlyphNoFallback(ImFont* self,ImWchar c) { return ImFont_FindGlyphNoFallback(self,c); }
float GetFontSize() { return igGetFontSize(); }
void SameLine(float offset_from_start_x,float spacing) { igSameLine(offset_from_start_x,spacing); }
void SetCursorPosX(float local_x) { igSetCursorPosX(local_x); }
void DrawList_PrimRect(ImDrawList* self,const ImVec2 a,const ImVec2 b,ImU32 col) { ImDrawList_PrimRect(self,a,b,col); }
bool BeginPopupContextWindow(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextWindow(str_id,popup_flags); }
void DestroyContext(ImGuiContext* ctx) { igDestroyContext(ctx); }
ImGuiStyle* GetStyle() { return igGetStyle(); }
bool SliderFloat2(const char* label,float v[2],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat2(label,v,v_min,v_max,format,flags); }
void SetScrollHereY(float center_y_ratio) { igSetScrollHereY(center_y_ratio); }
void FontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self,const ImWchar* ranges) { ImFontGlyphRangesBuilder_AddRanges(self,ranges); }
ImGuiIO* GetIO() { return igGetIO(); }
ImFontGlyphRangesBuilder* FontGlyphRangesBuilder_FontGlyphRangesBuilder() { return ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(); }
bool ListClipper_Step(ImGuiListClipper* self) { return ImGuiListClipper_Step(self); }
float* Storage_GetFloatRef(ImGuiStorage* self,ImGuiID key,float default_val) { return ImGuiStorage_GetFloatRef(self,key,default_val); }
void PushTextWrapPos(float wrap_local_pos_x) { igPushTextWrapPos(wrap_local_pos_x); }
void SetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot) { igSetNextWindowPos(pos,cond,pivot); }
void DrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness) { ImDrawList_AddLine(self,p1,p2,col,thickness); }
bool Payload_IsPreview(ImGuiPayload* self) { return ImGuiPayload_IsPreview(self); }
bool SliderFloat3(const char* label,float v[3],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat3(label,v,v_min,v_max,format,flags); }
void CalcTextSize(ImVec2 *pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width) { igCalcTextSize(pOut,text,text_end,hide_text_after_double_hash,wrap_width); }
bool IsKeyReleased(ImGuiKey key) { return igIsKeyReleased(key); }
bool SliderScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalar(label,data_type,p_data,p_min,p_max,format,flags); }
ImGuiTextBuffer* TextBuffer_TextBuffer() { return ImGuiTextBuffer_ImGuiTextBuffer(); }
ImGuiWindowClass* WindowClass_WindowClass() { return ImGuiWindowClass_ImGuiWindowClass(); }
ImGuiPlatformMonitor* PlatformMonitor_PlatformMonitor() { return ImGuiPlatformMonitor_ImGuiPlatformMonitor(); }
bool DragInt(const char* label,int* v,float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt(label,v,v_speed,v_min,v_max,format,flags); }
void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) { igSetNextFrameWantCaptureKeyboard(want_capture_keyboard); }
void Font_ClearOutputData(ImFont* self) { ImFont_ClearOutputData(self); }
const char* GetVersion() { return igGetVersion(); }
void DrawList_PathRect(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max,float rounding,ImDrawFlags flags) { ImDrawList_PathRect(self,rect_min,rect_max,rounding,flags); }
ImFont* Font_Font() { return ImFont_ImFont(); }
void IO_AddMouseButtonEvent(ImGuiIO* self,int button,bool down) { ImGuiIO_AddMouseButtonEvent(self,button,down); }
const ImWchar* FontAtlas_GetGlyphRangesKorean(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesKorean(self); }
void EndChild() { igEndChild(); }
const ImVec4* GetStyleColorVec4(ImGuiCol idx) { return igGetStyleColorVec4(idx); }
const char* Font_CalcWordWrapPositionA(ImFont* self,float scale,const char* text,const char* text_end,float wrap_width) { return ImFont_CalcWordWrapPositionA(self,scale,text,text_end,wrap_width); }
void EndTooltip() { igEndTooltip(); }
bool DebugCheckVersionAndDataLayout(const char* version_str,size_t sz_io,size_t sz_style,size_t sz_vec2,size_t sz_vec4,size_t sz_drawvert,size_t sz_drawidx) { return igDebugCheckVersionAndDataLayout(version_str,sz_io,sz_style,sz_vec2,sz_vec4,sz_drawvert,sz_drawidx); }
bool InputFloat4(const char* label,float v[4],const char* format,ImGuiInputTextFlags flags) { return igInputFloat4(label,v,format,flags); }
void SetNextWindowContentSize(const ImVec2 size) { igSetNextWindowContentSize(size); }
void StyleColorsClassic(ImGuiStyle* dst) { igStyleColorsClassic(dst); }
ImGuiOnceUponAFrame* OnceUponAFrame_OnceUponAFrame() { return ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(); }
bool DragFloat(const char* label,float* v,float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat(label,v,v_speed,v_min,v_max,format,flags); }
int GetMouseClickedCount(ImGuiMouseButton button) { return igGetMouseClickedCount(button); }
void TableHeader(const char* label) { igTableHeader(label); }
void ShowStyleEditor(ImGuiStyle* ref) { igShowStyleEditor(ref); }
void FontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self,size_t n) { ImFontGlyphRangesBuilder_SetBit(self,n); }
void EndMainMenuBar() { igEndMainMenuBar(); }
void DrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness) { ImDrawList_AddRect(self,p_min,p_max,col,rounding,flags,thickness); }
void EndDragDropTarget() { igEndDragDropTarget(); }
void DrawList_PathStroke(ImDrawList* self,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_PathStroke(self,col,flags,thickness); }
bool DragFloat3(const char* label,float v[3],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat3(label,v,v_speed,v_min,v_max,format,flags); }
void GetItemRectMax(ImVec2 *pOut) { igGetItemRectMax(pOut); }
void SetNextFrameWantCaptureMouse(bool want_capture_mouse) { igSetNextFrameWantCaptureMouse(want_capture_mouse); }
void ShowFontSelector(const char* label) { igShowFontSelector(label); }
void Font_SetGlyphVisible(ImFont* self,ImWchar c,bool visible) { ImFont_SetGlyphVisible(self,c,visible); }
bool IsItemVisible() { return igIsItemVisible(); }
bool InputTextWithHint(const char* label,const char* hint,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextWithHint(label,hint,buf,buf_size,flags,callback,user_data); }
void DrawList_PrimWriteVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimWriteVtx(self,pos,uv,col); }
void Payload_Clear(ImGuiPayload* self) { ImGuiPayload_Clear(self); }
void GetCursorPos(ImVec2 *pOut) { igGetCursorPos(pOut); }
int TableGetColumnCount() { return igTableGetColumnCount(); }
void InputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self,int pos,const char* text,const char* text_end) { ImGuiInputTextCallbackData_InsertChars(self,pos,text,text_end); }
void ColorConvertRGBtoHSV(float r,float g,float b,float* out_h,float* out_s,float* out_v) { igColorConvertRGBtoHSV(r,g,b,out_h,out_s,out_v); }
void Storage_SetInt(ImGuiStorage* self,ImGuiID key,int val) { ImGuiStorage_SetInt(self,key,val); }
ImGuiViewport* FindViewportByID(ImGuiID id) { return igFindViewportByID(id); }
bool IsAnyItemHovered() { return igIsAnyItemHovered(); }
void DrawList_ChannelsMerge(ImDrawList* self) { ImDrawList_ChannelsMerge(self); }
void** Storage_GetVoidPtrRef(ImGuiStorage* self,ImGuiID key,void* default_val) { return ImGuiStorage_GetVoidPtrRef(self,key,default_val); }
bool IsMousePosValid(const ImVec2* mouse_pos) { return igIsMousePosValid(mouse_pos); }
void IO_AddKeyEvent(ImGuiIO* self,ImGuiKey key,bool down) { ImGuiIO_AddKeyEvent(self,key,down); }
void DrawList_ChannelsSplit(ImDrawList* self,int count) { ImDrawList_ChannelsSplit(self,count); }
bool SliderScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalarN(label,data_type,p_data,components,p_min,p_max,format,flags); }
bool TextFilter_PassFilter(ImGuiTextFilter* self,const char* text,const char* text_end) { return ImGuiTextFilter_PassFilter(self,text,text_end); }
void DrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self) { ImDrawListSplitter_ClearFreeMemory(self); }
void FontAtlas_ClearFonts(ImFontAtlas* self) { ImFontAtlas_ClearFonts(self); }
void InputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_SelectAll(self); }
bool InputInt4(const char* label,int v[4],ImGuiInputTextFlags flags) { return igInputInt4(label,v,flags); }
void IO_AddInputCharacter(ImGuiIO* self,unsigned int c) { ImGuiIO_AddInputCharacter(self,c); }
bool BeginDragDropSource(ImGuiDragDropFlags flags) { return igBeginDragDropSource(flags); }
void EndMenu() { igEndMenu(); }
const ImGuiPayload* AcceptDragDropPayload(const char* type,ImGuiDragDropFlags flags) { return igAcceptDragDropPayload(type,flags); }
void DestroyPlatformWindows() { igDestroyPlatformWindows(); }
void AlignTextToFramePadding() { igAlignTextToFramePadding(); }
ImFontAtlas* FontAtlas_FontAtlas() { return ImFontAtlas_ImFontAtlas(); }
void FontAtlas_SetTexID(ImFontAtlas* self,ImTextureID id) { ImFontAtlas_SetTexID(self,id); }
void Dummy(const ImVec2 size) { igDummy(size); }
bool SliderFloat4(const char* label,float v[4],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat4(label,v,v_min,v_max,format,flags); }
void SetScrollHereX(float center_x_ratio) { igSetScrollHereX(center_x_ratio); }
bool Font_IsLoaded(ImFont* self) { return ImFont_IsLoaded(self); }
void GetWindowContentRegionMax(ImVec2 *pOut) { igGetWindowContentRegionMax(pOut); }
bool BeginDragDropTarget() { return igBeginDragDropTarget(); }
void Viewport_GetCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetCenter(pOut,self); }
void EndCombo() { igEndCombo(); }
void SetNextWindowSize(const ImVec2 size,ImGuiCond cond) { igSetNextWindowSize(size,cond); }
void UpdatePlatformWindows() { igUpdatePlatformWindows(); }
void SetItemDefaultFocus() { igSetItemDefaultFocus(); }
void ResetMouseDragDelta(ImGuiMouseButton button) { igResetMouseDragDelta(button); }
int TableGetRowIndex() { return igTableGetRowIndex(); }
void TableSetupColumn(const char* label,ImGuiTableColumnFlags flags,float init_width_or_weight,ImGuiID user_id) { igTableSetupColumn(label,flags,init_width_or_weight,user_id); }
void Color_SetHSV(ImColor* self,float h,float s,float v,float a) { ImColor_SetHSV(self,h,s,v,a); }
void DrawList_PathLineToMergeDuplicate(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineToMergeDuplicate(self,pos); }
void DrawList_AddBezierCubic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierCubic(self,p1,p2,p3,p4,col,thickness,num_segments); }
void End() { igEnd(); }
void ProgressBar(float fraction,const ImVec2 size_arg,const char* overlay) { igProgressBar(fraction,size_arg,overlay); }
void FontAtlas_CalcCustomRectUV(ImFontAtlas* self,const ImFontAtlasCustomRect* rect,ImVec2* out_uv_min,ImVec2* out_uv_max) { ImFontAtlas_CalcCustomRectUV(self,rect,out_uv_min,out_uv_max); }
void Font_GrowIndex(ImFont* self,int new_size) { ImFont_GrowIndex(self,new_size); }
void ListClipper_End(ImGuiListClipper* self) { ImGuiListClipper_End(self); }
ImFont* FontAtlas_AddFontDefault(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFontDefault(self,font_cfg); }
void GetMousePosOnOpeningCurrentPopup(ImVec2 *pOut) { igGetMousePosOnOpeningCurrentPopup(pOut); }
bool TableNextColumn() { return igTableNextColumn(); }
void PopFont() { igPopFont(); }
void Font_CalcTextSizeA(ImVec2 *pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin,const char* text_end,const char** remaining) { ImFont_CalcTextSizeA(pOut,self,size,max_width,wrap_width,text_begin,text_end,remaining); }
void SetNextWindowFocus() { igSetNextWindowFocus(); }
void IO_AddMouseViewportEvent(ImGuiIO* self,ImGuiID id) { ImGuiIO_AddMouseViewportEvent(self,id); }
ImGuiContext* CreateContext(ImFontAtlas* shared_font_atlas) { return igCreateContext(shared_font_atlas); }
float GetFrameHeight() { return igGetFrameHeight(); }
void InputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_ClearSelection(self); }
bool BeginTabBar(const char* str_id,ImGuiTabBarFlags flags) { return igBeginTabBar(str_id,flags); }
void GetItemRectSize(ImVec2 *pOut) { igGetItemRectSize(pOut); }
void GetWindowSize(ImVec2 *pOut) { igGetWindowSize(pOut); }
bool IsKeyPressed(ImGuiKey key,bool repeat) { return igIsKeyPressed(key,repeat); }
ImFont* FontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self,compressed_font_data_base85,size_pixels,font_cfg,glyph_ranges); }
void FontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self,const char* text,const char* text_end) { ImFontGlyphRangesBuilder_AddText(self,text,text_end); }
void SetMouseCursor(ImGuiMouseCursor cursor_type) { igSetMouseCursor(cursor_type); }
bool InputFloat(const char* label,float* v,float step,float step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputFloat(label,v,step,step_fast,format,flags); }
bool FontAtlas_GetMouseCursorTexData(ImFontAtlas* self,ImGuiMouseCursor cursor,ImVec2* out_offset,ImVec2* out_size,ImVec2 out_uv_border[2],ImVec2 out_uv_fill[2]) { return ImFontAtlas_GetMouseCursorTexData(self,cursor,out_offset,out_size,out_uv_border,out_uv_fill); }
bool DragIntRange2(const char* label,int* v_current_min,int* v_current_max,float v_speed,int v_min,int v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragIntRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
void Unindent(float indent_w) { igUnindent(indent_w); }
void GetContentRegionAvail(ImVec2 *pOut) { igGetContentRegionAvail(pOut); }
void DrawList_AddCircle(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddCircle(self,center,radius,col,num_segments,thickness); }
void Font_BuildLookupTable(ImFont* self) { ImFont_BuildLookupTable(self); }
bool IsAnyMouseDown() { return igIsAnyMouseDown(); }
void IO_ClearInputKeys(ImGuiIO* self) { ImGuiIO_ClearInputKeys(self); }
bool IsItemActive() { return igIsItemActive(); }
void LogButtons() { igLogButtons(); }
void SetKeyboardFocusHere(int offset) { igSetKeyboardFocusHere(offset); }
void SetNextItemOpen(bool is_open,ImGuiCond cond) { igSetNextItemOpen(is_open,cond); }
void DrawList_AddNgonFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddNgonFilled(self,center,radius,col,num_segments); }
void Font_AddRemapChar(ImFont* self,ImWchar dst,ImWchar src,bool overwrite_dst) { ImFont_AddRemapChar(self,dst,src,overwrite_dst); }
void LoadIniSettingsFromMemory(const char* ini_data,size_t ini_size) { igLoadIniSettingsFromMemory(ini_data,ini_size); }
void DrawList_PrimQuadUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 d,const ImVec2 uv_a,const ImVec2 uv_b,const ImVec2 uv_c,const ImVec2 uv_d,ImU32 col) { ImDrawList_PrimQuadUV(self,a,b,c,d,uv_a,uv_b,uv_c,uv_d,col); }
void PopClipRect() { igPopClipRect(); }
void FontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self,ImWchar c) { ImFontGlyphRangesBuilder_AddChar(self,c); }
void EndPopup() { igEndPopup(); }
bool IsMouseDown(ImGuiMouseButton button) { return igIsMouseDown(button); }
void SetNextWindowCollapsed(bool collapsed,ImGuiCond cond) { igSetNextWindowCollapsed(collapsed,cond); }
bool BeginListBox(const char* label,const ImVec2 size) { return igBeginListBox(label,size); }
void RenderPlatformWindowsDefault(void* platform_render_arg,void* renderer_render_arg) { igRenderPlatformWindowsDefault(platform_render_arg,renderer_render_arg); }
bool ColorPicker4(const char* label,float col[4],ImGuiColorEditFlags flags,const float* ref_col) { return igColorPicker4(label,col,flags,ref_col); }
ImTextureID DrawCmd_GetTexID(ImDrawCmd* self) { return ImDrawCmd_GetTexID(self); }
ImFont* FontAtlas_AddFont(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFont(self,font_cfg); }
int* Storage_GetIntRef(ImGuiStorage* self,ImGuiID key,int default_val) { return ImGuiStorage_GetIntRef(self,key,default_val); }
bool BeginMenu(const char* label,bool enabled) { return igBeginMenu(label,enabled); }
void GetContentRegionMax(ImVec2 *pOut) { igGetContentRegionMax(pOut); }
void SetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data) { igSetNextWindowSizeConstraints(size_min,size_max,custom_callback,custom_callback_data); }
float Font_GetCharAdvance(ImFont* self,ImWchar c) { return ImFont_GetCharAdvance(self,c); }
float GetCursorPosY() { return igGetCursorPosY(); }
void MemFree(void* ptr) { igMemFree(ptr); }
ImGuiStyle* Style_Style() { return ImGuiStyle_ImGuiStyle(); }
ImGuiID DockSpace(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpace(id,size,flags,window_class); }
bool IsMouseReleased(ImGuiMouseButton button) { return igIsMouseReleased(button); }
void SetNextWindowBgAlpha(float alpha) { igSetNextWindowBgAlpha(alpha); }
void BeginDisabled(bool disabled) { igBeginDisabled(disabled); }
void PushFont(ImFont* font) { igPushFont(font); }
void SetColumnWidth(int column_index,float width) { igSetColumnWidth(column_index,width); }
void DrawList_PopTextureID(ImDrawList* self) { ImDrawList_PopTextureID(self); }
const char* Font_GetDebugName(ImFont* self) { return ImFont_GetDebugName(self); }
int Storage_GetInt(ImGuiStorage* self,ImGuiID key,int default_val) { return ImGuiStorage_GetInt(self,key,default_val); }
bool ColorEdit3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorEdit3(label,col,flags); }
int GetFrameCount() { return igGetFrameCount(); }
ImGuiPayload* Payload_Payload() { return ImGuiPayload_ImGuiPayload(); }
void InputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self,int pos,int bytes_count) { ImGuiInputTextCallbackData_DeleteChars(self,pos,bytes_count); }
bool Checkbox(const char* label,bool* v) { return igCheckbox(label,v); }
void FontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self) { ImFontGlyphRangesBuilder_Clear(self); }
void DrawList_AddPolyline(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_AddPolyline(self,points,num_points,col,flags,thickness); }
void Storage_SetVoidPtr(ImGuiStorage* self,ImGuiID key,void* val) { ImGuiStorage_SetVoidPtr(self,key,val); }
void DrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { ImDrawList_PushClipRect(self,clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void DrawData_DeIndexAllBuffers(ImDrawData* self) { ImDrawData_DeIndexAllBuffers(self); }
bool Payload_IsDelivery(ImGuiPayload* self) { return ImGuiPayload_IsDelivery(self); }
float GetTextLineHeight() { return igGetTextLineHeight(); }
void EndChildFrame() { igEndChildFrame(); }
void DrawList_AddQuadFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col) { ImDrawList_AddQuadFilled(self,p1,p2,p3,p4,col); }
void Font_RenderChar(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,ImWchar c) { ImFont_RenderChar(self,draw_list,size,pos,col,c); }
bool BeginPopupContextItem(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextItem(str_id,popup_flags); }
ImFont* GetFont() { return igGetFont(); }
void GetCursorStartPos(ImVec2 *pOut) { igGetCursorStartPos(pOut); }
void DrawListSplitter_Clear(ImDrawListSplitter* self) { ImDrawListSplitter_Clear(self); }
void DrawList_PrimUnreserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimUnreserve(self,idx_count,vtx_count); }
bool InputFloat2(const char* label,float v[2],const char* format,ImGuiInputTextFlags flags) { return igInputFloat2(label,v,format,flags); }
void DrawList_ChannelsSetCurrent(ImDrawList* self,int n) { ImDrawList_ChannelsSetCurrent(self,n); }
void PopStyleVar(int count) { igPopStyleVar(count); }
void TableNextRow(ImGuiTableRowFlags row_flags,float min_row_height) { igTableNextRow(row_flags,min_row_height); }
bool ArrowButton(const char* str_id,ImGuiDir dir) { return igArrowButton(str_id,dir); }
void DrawList_AddNgon(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddNgon(self,center,radius,col,num_segments,thickness); }
ImFontAtlasCustomRect* FontAtlasCustomRect_FontAtlasCustomRect() { return ImFontAtlasCustomRect_ImFontAtlasCustomRect(); }
const ImWchar* FontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesCyrillic(self); }
ImFontConfig* FontConf_FontConfig() { return ImFontConfig_ImFontConfig(); }
void SetNextWindowDockID(ImGuiID dock_id,ImGuiCond cond) { igSetNextWindowDockID(dock_id,cond); }
void ShowDemoWindow(bool* p_open) { igShowDemoWindow(p_open); }
void NewFrame() { igNewFrame(); }
void SetColumnOffset(int column_index,float offset_x) { igSetColumnOffset(column_index,offset_x); }
ImDrawData* DrawData_DrawData() { return ImDrawData_ImDrawData(); }
float GetScrollY() { return igGetScrollY(); }
void DrawList_AddCircleFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddCircleFilled(self,center,radius,col,num_segments); }
void IO_AddMouseWheelEvent(ImGuiIO* self,float wh_x,float wh_y) { ImGuiIO_AddMouseWheelEvent(self,wh_x,wh_y); }
bool IsMouseClicked(ImGuiMouseButton button,bool repeat) { return igIsMouseClicked(button,repeat); }
bool DragScalar(const char* label,ImGuiDataType data_type,void* p_data,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalar(label,data_type,p_data,v_speed,p_min,p_max,format,flags); }
ImGuiContext* GetCurrentContext() { return igGetCurrentContext(); }
void PushAllowKeyboardFocus(bool allow_keyboard_focus) { igPushAllowKeyboardFocus(allow_keyboard_focus); }
bool TableSetColumnIndex(int column_n) { return igTableSetColumnIndex(column_n); }
void Spacing() { igSpacing(); }
void SetCursorPos(const ImVec2 local_pos) { igSetCursorPos(local_pos); }
int FontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x,const ImVec2 offset) { return ImFontAtlas_AddCustomRectFontGlyph(self,font,id,width,height,advance_x,offset); }
bool Button(const char* label,const ImVec2 size) { return igButton(label,size); }
ImGuiIO* IO_IO() { return ImGuiIO_ImGuiIO(); }
bool ColorEdit4(const char* label,float col[4],ImGuiColorEditFlags flags) { return igColorEdit4(label,col,flags); }
const char* GetClipboardText() { return igGetClipboardText(); }
float GetWindowHeight() { return igGetWindowHeight(); }
bool IsItemDeactivated() { return igIsItemDeactivated(); }
ImGuiTableColumnFlags TableGetColumnFlags(int column_n) { return igTableGetColumnFlags(column_n); }
void EndTabItem() { igEndTabItem(); }
int GetKeyIndex(ImGuiKey key) { return igGetKeyIndex(key); }
void StyleColorsLight(ImGuiStyle* dst) { igStyleColorsLight(dst); }
void DrawList_PrimRectUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,ImU32 col) { ImDrawList_PrimRectUV(self,a,b,uv_a,uv_b,col); }
const ImWchar* FontAtlas_GetGlyphRangesThai(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesThai(self); }
bool InputText(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputText(label,buf,buf_size,flags,callback,user_data); }
void ShowAboutWindow(bool* p_open) { igShowAboutWindow(p_open); }
bool FontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder* self,size_t n) { return ImFontGlyphRangesBuilder_GetBit(self,n); }
bool IsWindowFocused(ImGuiFocusedFlags flags) { return igIsWindowFocused(flags); }
void SetClipboardText(const char* text) { igSetClipboardText(text); }
ImGuiTableSortSpecs* TableSortSpecs_TableSortSpecs() { return ImGuiTableSortSpecs_ImGuiTableSortSpecs(); }
bool InputDouble(const char* label,double* v,double step,double step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputDouble(label,v,step,step_fast,format,flags); }
void SaveIniSettingsToDisk(const char* ini_filename) { igSaveIniSettingsToDisk(ini_filename); }
void BeginGroup() { igBeginGroup(); }
bool DragInt4(const char* label,int v[4],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt4(label,v,v_speed,v_min,v_max,format,flags); }
void DrawList_Addage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col) { ImDrawList_AddImage(self,user_texture_id,p_min,p_max,uv_min,uv_max,col); }
ImGuiPlatformIO* GetPlatformIO() { return igGetPlatformIO(); }
void ShowMetricsWindow(bool* p_open) { igShowMetricsWindow(p_open); }
bool BeginCombo(const char* label,const char* preview_value,ImGuiComboFlags flags) { return igBeginCombo(label,preview_value,flags); }
bool TabItemButton(const char* label,ImGuiTabItemFlags flags) { return igTabItemButton(label,flags); }
ImDrawList* GetWindowDrawList() { return igGetWindowDrawList(); }
void DrawList_AddBezierQuadratic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierQuadratic(self,p1,p2,p3,col,thickness,num_segments); }
void LogFinish() { igLogFinish(); }
bool VSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderFloat(label,size,v,v_min,v_max,format,flags); }
bool BeginPopupContextVoid(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextVoid(str_id,popup_flags); }
void StyleColorsDark(ImGuiStyle* dst) { igStyleColorsDark(dst); }
ImDrawList* DrawList_CloneOutput(ImDrawList* self) { return ImDrawList_CloneOutput(self); }
bool InputFloat3(const char* label,float v[3],const char* format,ImGuiInputTextFlags flags) { return igInputFloat3(label,v,format,flags); }
void DrawList_PathArcToFast(ImDrawList* self,const ImVec2 center,float radius,int a_min_of_12,int a_max_of_12) { ImDrawList_PathArcToFast(self,center,radius,a_min_of_12,a_max_of_12); }
void PopAllowKeyboardFocus() { igPopAllowKeyboardFocus(); }
bool DragFloat2(const char* label,float v[2],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat2(label,v,v_speed,v_min,v_max,format,flags); }
float GetScrollX() { return igGetScrollX(); }
void GetFontTexUvWhitePixel(ImVec2 *pOut) { igGetFontTexUvWhitePixel(pOut); }
int TableGetColumnIndex() { return igTableGetColumnIndex(); }
void DrawList_PathLineTo(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineTo(self,pos); }
ImGuiViewport* FindViewportByPlatformHandle(void* platform_handle) { return igFindViewportByPlatformHandle(platform_handle); }
bool DragInt2(const char* label,int v[2],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt2(label,v,v_speed,v_min,v_max,format,flags); }
void PopItemWidth() { igPopItemWidth(); }
const char* TextBuffer_c_str(ImGuiTextBuffer* self) { return ImGuiTextBuffer_c_str(self); }
void GetMouseDragDelta(ImVec2 *pOut,ImGuiMouseButton button,float lock_threshold) { igGetMouseDragDelta(pOut,button,lock_threshold); }
bool IsMouseDoubleClicked(ImGuiMouseButton button) { return igIsMouseDoubleClicked(button); }
ImDrawList* DrawList_DrawList(const ImDrawListSharedData* shared_data) { return ImDrawList_ImDrawList(shared_data); }
double GetTime() { return igGetTime(); }
bool IsItemDeactivatedAfterEdit() { return igIsItemDeactivatedAfterEdit(); }
void EndFrame() { igEndFrame(); }
bool IsItemActivated() { return igIsItemActivated(); }
void DrawList_GetClipRectMin(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMin(pOut,self); }
float GetTreeNodeToLabelSpacing() { return igGetTreeNodeToLabelSpacing(); }
void OpenPopupOnItemClick(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopupOnItemClick(str_id,popup_flags); }
void FontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self,ImVector_ImWchar* out_ranges) { ImFontGlyphRangesBuilder_BuildRanges(self,out_ranges); }
void LoadIniSettingsFromDisk(const char* ini_filename) { igLoadIniSettingsFromDisk(ini_filename); }
const ImWchar* FontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesVietnamese(self); }
void Storage_BuildSortByKey(ImGuiStorage* self) { ImGuiStorage_BuildSortByKey(self); }
void EndGroup() { igEndGroup(); }
ImDrawData* GetDrawData() { return igGetDrawData(); }
void CloseCurrentPopup() { igCloseCurrentPopup(); }
void Columns(int count,const char* id,bool border) { igColumns(count,id,border); }
void ShowDebugLogWindow(bool* p_open) { igShowDebugLogWindow(p_open); }
void NextColumn() { igNextColumn(); }
void Viewport_GetWorkCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetWorkCenter(pOut,self); }
ImDrawListSplitter* DrawListSplitter_DrawListSplitter() { return ImDrawListSplitter_ImDrawListSplitter(); }
void ListClipper_Begin(ImGuiListClipper* self,int items_count,float items_height) { ImGuiListClipper_Begin(self,items_count,items_height); }
bool DragInt3(const char* label,int v[3],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt3(label,v,v_speed,v_min,v_max,format,flags); }
bool SmallButton(const char* label) { return igSmallButton(label); }
const ImWchar* FontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseFull(self); }
const char* GetStyleColorName(ImGuiCol idx) { return igGetStyleColorName(idx); }
void IO_AddKeyAnalogEvent(ImGuiIO* self,ImGuiKey key,bool down,float v) { ImGuiIO_AddKeyAnalogEvent(self,key,down,v); }
void GetItemRectMin(ImVec2 *pOut) { igGetItemRectMin(pOut); }
void LogToFile(int auto_open_depth,const char* filename) { igLogToFile(auto_open_depth,filename); }
void Separator() { igSeparator(); }
float GetColumnOffset(int column_index) { return igGetColumnOffset(column_index); }
void SetWindowFontScale(float scale) { igSetWindowFontScale(scale); }
void IO_AddInputCharactersUTF8(ImGuiIO* self,const char* str) { ImGuiIO_AddInputCharactersUTF8(self,str); }
ImGuiID GetWindowDockID() { return igGetWindowDockID(); }
