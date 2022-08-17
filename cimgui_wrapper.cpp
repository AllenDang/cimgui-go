#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

bool IsAnyMouseDown() { return igIsAnyMouseDown(); }
ImGuiPlatformIO* ImGuiPlatformIO_ImGuiPlatformIO() { return ImGuiPlatformIO_ImGuiPlatformIO(); }
bool BeginListBox(const char* label,const ImVec2 size) { return igBeginListBox(label,size); }
void PopStyleColor(int count) { igPopStyleColor(count); }
bool VSliderFloat(const char* label,const ImVec2 size,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderFloat(label,size,v,v_min,v_max,format,flags); }
void ImDrawList_AddImage(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col) { ImDrawList_AddImage(self,user_texture_id,p_min,p_max,uv_min,uv_max,col); }
void ImFontAtlas_GetTexDataAsRGBA32(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsRGBA32(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
ImGuiViewport* GetMainViewport() { return igGetMainViewport(); }
float GetWindowWidth() { return igGetWindowWidth(); }
void ImDrawList_PrimUnreserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimUnreserve(self,idx_count,vtx_count); }
void EndDragDropSource() { igEndDragDropSource(); }
bool ImFontGlyphRangesBuilder_GetBit(ImFontGlyphRangesBuilder* self,size_t n) { return ImFontGlyphRangesBuilder_GetBit(self,n); }
bool DragFloat2(const char* label,float v[2],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat2(label,v,v_speed,v_min,v_max,format,flags); }
void EndFrame() { igEndFrame(); }
void NewLine() { igNewLine(); }
ImDrawData* ImDrawData_ImDrawData() { return ImDrawData_ImDrawData(); }
void ImDrawList_AddConvexPolyFilled(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col) { ImDrawList_AddConvexPolyFilled(self,points,num_points,col); }
void ImDrawList_PrimVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimVtx(self,pos,uv,col); }
bool SetDragDropPayload(const char* type,const void* data,size_t sz,ImGuiCond cond) { return igSetDragDropPayload(type,data,sz,cond); }
void ShowFontSelector(const char* label) { igShowFontSelector(label); }
void ImDrawData_DeIndexAllBuffers(ImDrawData* self) { ImDrawData_DeIndexAllBuffers(self); }
void Indent(float indent_w) { igIndent(indent_w); }
bool IsItemFocused() { return igIsItemFocused(); }
void ImDrawList_PathLineTo(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineTo(self,pos); }
ImGuiStyle* ImGuiStyle_ImGuiStyle() { return ImGuiStyle_ImGuiStyle(); }
bool BeginTable(const char* str_id,int column,ImGuiTableFlags flags,const ImVec2 outer_size,float inner_width) { return igBeginTable(str_id,column,flags,outer_size,inner_width); }
void ImDrawList_AddBezierQuadratic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierQuadratic(self,p1,p2,p3,col,thickness,num_segments); }
ImGuiViewport* ImGuiViewport_ImGuiViewport() { return ImGuiViewport_ImGuiViewport(); }
void DestroyContext(ImGuiContext* ctx) { igDestroyContext(ctx); }
void PushAllowKeyboardFocus(bool allow_keyboard_focus) { igPushAllowKeyboardFocus(allow_keyboard_focus); }
bool SliderInt(const char* label,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt(label,v,v_min,v_max,format,flags); }
void ImDrawData_ScaleClipRects(ImDrawData* self,const ImVec2 fb_scale) { ImDrawData_ScaleClipRects(self,fb_scale); }
void ImFontAtlas_Clear(ImFontAtlas* self) { ImFontAtlas_Clear(self); }
const ImWchar* ImFontAtlas_GetGlyphRangesDefault(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesDefault(self); }
void SetNextItemOpen(bool is_open,ImGuiCond cond) { igSetNextItemOpen(is_open,cond); }
void ImDrawList_AddTriangle(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col,float thickness) { ImDrawList_AddTriangle(self,p1,p2,p3,col,thickness); }
void ImGuiViewport_GetCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetCenter(pOut,self); }
void ImFontGlyphRangesBuilder_AddChar(ImFontGlyphRangesBuilder* self,ImWchar c) { ImFontGlyphRangesBuilder_AddChar(self,c); }
bool IsItemActive() { return igIsItemActive(); }
void StyleColorsLight(ImGuiStyle* dst) { igStyleColorsLight(dst); }
const ImWchar* ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseSimplifiedCommon(self); }
void SetNextWindowFocus() { igSetNextWindowFocus(); }
bool DragInt4(const char* label,int v[4],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt4(label,v,v_speed,v_min,v_max,format,flags); }
void EndTabBar() { igEndTabBar(); }
ImGuiID GetWindowDockID() { return igGetWindowDockID(); }
int TableGetColumnIndex() { return igTableGetColumnIndex(); }
const char* GetVersion() { return igGetVersion(); }
void ImDrawList_AddRectFilledMultiColor(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col_upr_left,ImU32 col_upr_right,ImU32 col_bot_right,ImU32 col_bot_left) { ImDrawList_AddRectFilledMultiColor(self,p_min,p_max,col_upr_left,col_upr_right,col_bot_right,col_bot_left); }
ImGuiTableColumnSortSpecs* ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs() { return ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs(); }
void SetCursorPosX(float local_x) { igSetCursorPosX(local_x); }
void ImGuiInputTextCallbackData_SelectAll(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_SelectAll(self); }
int ImGuiStorage_GetInt(ImGuiStorage* self,ImGuiID key,int default_val) { return ImGuiStorage_GetInt(self,key,default_val); }
void ImDrawList_GetClipRectMin(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMin(pOut,self); }
void SameLine(float offset_from_start_x,float spacing) { igSameLine(offset_from_start_x,spacing); }
void SetClipboardText(const char* text) { igSetClipboardText(text); }
float GetScrollMaxX() { return igGetScrollMaxX(); }
bool SliderFloat2(const char* label,float v[2],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat2(label,v,v_min,v_max,format,flags); }
void ImDrawList_AddRectFilled(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddRectFilled(self,p_min,p_max,col,rounding,flags); }
const ImVec4* GetStyleColorVec4(ImGuiCol idx) { return igGetStyleColorVec4(idx); }
bool InputFloat2(const char* label,float v[2],const char* format,ImGuiInputTextFlags flags) { return igInputFloat2(label,v,format,flags); }
bool SliderFloat4(const char* label,float v[4],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat4(label,v,v_min,v_max,format,flags); }
bool IsWindowFocused(ImGuiFocusedFlags flags) { return igIsWindowFocused(flags); }
void ImDrawList_PathBezierQuadraticCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,int num_segments) { ImDrawList_PathBezierQuadraticCurveTo(self,p2,p3,num_segments); }
bool IsWindowCollapsed() { return igIsWindowCollapsed(); }
bool TableNextColumn() { return igTableNextColumn(); }
const ImWchar* ImFontAtlas_GetGlyphRangesJapanese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesJapanese(self); }
void ColorConvertRGBtoHSV(float r,float g,float b,float* out_h,float* out_s,float* out_v) { igColorConvertRGBtoHSV(r,g,b,out_h,out_s,out_v); }
void ImDrawList_PushClipRectFullScreen(ImDrawList* self) { ImDrawList_PushClipRectFullScreen(self); }
int ImFontAtlas_AddCustomRectRegular(ImFontAtlas* self,int width,int height) { return ImFontAtlas_AddCustomRectRegular(self,width,height); }
void* MemAlloc(size_t size) { return igMemAlloc(size); }
bool BeginMainMenuBar() { return igBeginMainMenuBar(); }
ImDrawList* GetWindowDrawList() { return igGetWindowDrawList(); }
bool BeginMenuBar() { return igBeginMenuBar(); }
bool SmallButton(const char* label) { return igSmallButton(label); }
void ImGuiInputTextCallbackData_DeleteChars(ImGuiInputTextCallbackData* self,int pos,int bytes_count) { ImGuiInputTextCallbackData_DeleteChars(self,pos,bytes_count); }
bool InvisibleButton(const char* str_id,const ImVec2 size,ImGuiButtonFlags flags) { return igInvisibleButton(str_id,size,flags); }
bool IsMousePosValid(const ImVec2* mouse_pos) { return igIsMousePosValid(mouse_pos); }
ImFont* ImFont_ImFont() { return ImFont_ImFont(); }
bool SliderInt4(const char* label,int v[4],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt4(label,v,v_min,v_max,format,flags); }
ImDrawCmd* ImDrawCmd_ImDrawCmd() { return ImDrawCmd_ImDrawCmd(); }
void SetNextWindowDockID(ImGuiID dock_id,ImGuiCond cond) { igSetNextWindowDockID(dock_id,cond); }
void ImGuiStorage_SetFloat(ImGuiStorage* self,ImGuiID key,float val) { ImGuiStorage_SetFloat(self,key,val); }
bool BeginPopupContextVoid(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextVoid(str_id,popup_flags); }
ImGuiID DockSpace(ImGuiID id,const ImVec2 size,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpace(id,size,flags,window_class); }
bool ImFontAtlas_Build(ImFontAtlas* self) { return ImFontAtlas_Build(self); }
bool ImFont_IsLoaded(ImFont* self) { return ImFont_IsLoaded(self); }
const char* ImFont_GetDebugName(ImFont* self) { return ImFont_GetDebugName(self); }
bool IsKeyPressed(ImGuiKey key,bool repeat) { return igIsKeyPressed(key,repeat); }
void** ImGuiStorage_GetVoidPtrRef(ImGuiStorage* self,ImGuiID key,void* default_val) { return ImGuiStorage_GetVoidPtrRef(self,key,default_val); }
void PopID() { igPopID(); }
void ImFontGlyphRangesBuilder_AddRanges(ImFontGlyphRangesBuilder* self,const ImWchar* ranges) { ImFontGlyphRangesBuilder_AddRanges(self,ranges); }
const ImFontGlyph* ImFont_FindGlyphNoFallback(ImFont* self,ImWchar c) { return ImFont_FindGlyphNoFallback(self,c); }
void ImGuiIO_AddMouseViewportEvent(ImGuiIO* self,ImGuiID id) { ImGuiIO_AddMouseViewportEvent(self,id); }
void TextUnformatted(const char* text,const char* text_end) { igTextUnformatted(text,text_end); }
void UpdatePlatformWindows() { igUpdatePlatformWindows(); }
void ImDrawList_PathFillConvex(ImDrawList* self,ImU32 col) { ImDrawList_PathFillConvex(self,col); }
bool Checkbox(const char* label,bool* v) { return igCheckbox(label,v); }
const ImWchar* ImFontAtlas_GetGlyphRangesVietnamese(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesVietnamese(self); }
void GetMousePos(ImVec2 *pOut) { igGetMousePos(pOut); }
void ImGuiTextFilter_Build(ImGuiTextFilter* self) { ImGuiTextFilter_Build(self); }
bool DebugCheckVersionAndDataLayout(const char* version_str,size_t sz_io,size_t sz_style,size_t sz_vec2,size_t sz_vec4,size_t sz_drawvert,size_t sz_drawidx) { return igDebugCheckVersionAndDataLayout(version_str,sz_io,sz_style,sz_vec2,sz_vec4,sz_drawvert,sz_drawidx); }
bool DragFloatRange2(const char* label,float* v_current_min,float* v_current_max,float v_speed,float v_min,float v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragFloatRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
void ImDrawList_PathArcToFast(ImDrawList* self,const ImVec2 center,float radius,int a_min_of_12,int a_max_of_12) { ImDrawList_PathArcToFast(self,center,radius,a_min_of_12,a_max_of_12); }
const ImFontGlyph* ImFont_FindGlyph(ImFont* self,ImWchar c) { return ImFont_FindGlyph(self,c); }
bool IsAnyItemActive() { return igIsAnyItemActive(); }
void BeginDisabled(bool disabled) { igBeginDisabled(disabled); }
void SetColumnWidth(int column_index,float width) { igSetColumnWidth(column_index,width); }
void SetCurrentContext(ImGuiContext* ctx) { igSetCurrentContext(ctx); }
void ImFont_RenderText(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,const ImVec4 clip_rect,const char* text_begin,const char* text_end,float wrap_width,bool cpu_fine_clip) { ImFont_RenderText(self,draw_list,size,pos,col,clip_rect,text_begin,text_end,wrap_width,cpu_fine_clip); }
ImGuiIO* ImGuiIO_ImGuiIO() { return ImGuiIO_ImGuiIO(); }
bool IsItemVisible() { return igIsItemVisible(); }
bool BeginChildFrame(ImGuiID id,const ImVec2 size,ImGuiWindowFlags flags) { return igBeginChildFrame(id,size,flags); }
void ColorConvertU32ToFloat4(ImVec4 *pOut,ImU32 in) { igColorConvertU32ToFloat4(pOut,in); }
int GetKeyPressedAmount(ImGuiKey key,float repeat_delay,float rate) { return igGetKeyPressedAmount(key,repeat_delay,rate); }
ImDrawListSplitter* ImDrawListSplitter_ImDrawListSplitter() { return ImDrawListSplitter_ImDrawListSplitter(); }
float GetFontSize() { return igGetFontSize(); }
void ImFont_BuildLookupTable(ImFont* self) { ImFont_BuildLookupTable(self); }
void ShowAboutWindow(bool* p_open) { igShowAboutWindow(p_open); }
void ImDrawListSplitter_Split(ImDrawListSplitter* self,ImDrawList* draw_list,int count) { ImDrawListSplitter_Split(self,draw_list,count); }
void ImDrawList_AddDrawCmd(ImDrawList* self) { ImDrawList_AddDrawCmd(self); }
void ImDrawList_AddNgon(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddNgon(self,center,radius,col,num_segments,thickness); }
bool Button(const char* label,const ImVec2 size) { return igButton(label,size); }
void ImGuiStyle_ScaleAllSizes(ImGuiStyle* self,float scale_factor) { ImGuiStyle_ScaleAllSizes(self,scale_factor); }
void LoadIniSettingsFromMemory(const char* ini_data,size_t ini_size) { igLoadIniSettingsFromMemory(ini_data,ini_size); }
ImGuiListClipper* ImGuiListClipper_ImGuiListClipper() { return ImGuiListClipper_ImGuiListClipper(); }
bool BeginPopupModal(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBeginPopupModal(name,p_open,flags); }
void GetWindowContentRegionMin(ImVec2 *pOut) { igGetWindowContentRegionMin(pOut); }
float GetCursorPosX() { return igGetCursorPosX(); }
const ImWchar* ImFontAtlas_GetGlyphRangesKorean(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesKorean(self); }
bool ColorPicker3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorPicker3(label,col,flags); }
int GetColumnsCount() { return igGetColumnsCount(); }
bool DragInt(const char* label,int* v,float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt(label,v,v_speed,v_min,v_max,format,flags); }
void ImGuiIO_AddMouseButtonEvent(ImGuiIO* self,int button,bool down) { ImGuiIO_AddMouseButtonEvent(self,button,down); }
bool ColorButton(const char* desc_id,const ImVec4 col,ImGuiColorEditFlags flags,const ImVec2 size) { return igColorButton(desc_id,col,flags,size); }
bool IsItemActivated() { return igIsItemActivated(); }
bool DragFloat3(const char* label,float v[3],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat3(label,v,v_speed,v_min,v_max,format,flags); }
void SetNextFrameWantCaptureKeyboard(bool want_capture_keyboard) { igSetNextFrameWantCaptureKeyboard(want_capture_keyboard); }
void Render() { igRender(); }
void SetTabItemClosed(const char* tab_or_docked_window_label) { igSetTabItemClosed(tab_or_docked_window_label); }
void ImDrawList_AddTriangleFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,ImU32 col) { ImDrawList_AddTriangleFilled(self,p1,p2,p3,col); }
void ImFontAtlas_ClearFonts(ImFontAtlas* self) { ImFontAtlas_ClearFonts(self); }
void PushTextWrapPos(float wrap_local_pos_x) { igPushTextWrapPos(wrap_local_pos_x); }
float GetFrameHeight() { return igGetFrameHeight(); }
ImGuiStorage* GetStateStorage() { return igGetStateStorage(); }
void ImGuiStorage_Clear(ImGuiStorage* self) { ImGuiStorage_Clear(self); }
bool DragFloat(const char* label,float* v,float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat(label,v,v_speed,v_min,v_max,format,flags); }
bool DragScalar(const char* label,ImGuiDataType data_type,void* p_data,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalar(label,data_type,p_data,v_speed,p_min,p_max,format,flags); }
void ImDrawListSplitter_Clear(ImDrawListSplitter* self) { ImDrawListSplitter_Clear(self); }
void ImGuiIO_ClearInputKeys(ImGuiIO* self) { ImGuiIO_ClearInputKeys(self); }
bool* ImGuiStorage_GetBoolRef(ImGuiStorage* self,ImGuiID key,bool default_val) { return ImGuiStorage_GetBoolRef(self,key,default_val); }
void SetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data) { igSetNextWindowSizeConstraints(size_min,size_max,custom_callback,custom_callback_data); }
float GetTextLineHeightWithSpacing() { return igGetTextLineHeightWithSpacing(); }
bool InputScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalarN(label,data_type,p_data,components,p_step,p_step_fast,format,flags); }
void SetColumnOffset(int column_index,float offset_x) { igSetColumnOffset(column_index,offset_x); }
bool IsMouseHoveringRect(const ImVec2 r_min,const ImVec2 r_max,bool clip) { return igIsMouseHoveringRect(r_min,r_max,clip); }
void ImDrawList_AddImageRounded(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p_min,const ImVec2 p_max,const ImVec2 uv_min,const ImVec2 uv_max,ImU32 col,float rounding,ImDrawFlags flags) { ImDrawList_AddImageRounded(self,user_texture_id,p_min,p_max,uv_min,uv_max,col,rounding,flags); }
int GetColumnIndex() { return igGetColumnIndex(); }
void ImDrawList_PushTextureID(ImDrawList* self,ImTextureID texture_id) { ImDrawList_PushTextureID(self,texture_id); }
void ImFont_RenderChar(ImFont* self,ImDrawList* draw_list,float size,const ImVec2 pos,ImU32 col,ImWchar c) { ImFont_RenderChar(self,draw_list,size,pos,col,c); }
bool BeginCombo(const char* label,const char* preview_value,ImGuiComboFlags flags) { return igBeginCombo(label,preview_value,flags); }
void GetItemRectMax(ImVec2 *pOut) { igGetItemRectMax(pOut); }
bool InputFloat(const char* label,float* v,float step,float step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputFloat(label,v,step,step_fast,format,flags); }
void CloseCurrentPopup() { igCloseCurrentPopup(); }
bool IsMouseClicked(ImGuiMouseButton button,bool repeat) { return igIsMouseClicked(button,repeat); }
void ShowDemoWindow(bool* p_open) { igShowDemoWindow(p_open); }
void ShowMetricsWindow(bool* p_open) { igShowMetricsWindow(p_open); }
void End() { igEnd(); }
void ImGuiListClipper_ForceDisplayRangeByIndices(ImGuiListClipper* self,int item_min,int item_max) { ImGuiListClipper_ForceDisplayRangeByIndices(self,item_min,item_max); }
void PushButtonRepeat(bool repeat) { igPushButtonRepeat(repeat); }
bool IsMouseReleased(ImGuiMouseButton button) { return igIsMouseReleased(button); }
ImFontGlyphRangesBuilder* ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder() { return ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder(); }
ImGuiOnceUponAFrame* ImGuiOnceUponAFrame_ImGuiOnceUponAFrame() { return ImGuiOnceUponAFrame_ImGuiOnceUponAFrame(); }
void ImDrawList_AddBezierCubic(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness,int num_segments) { ImDrawList_AddBezierCubic(self,p1,p2,p3,p4,col,thickness,num_segments); }
bool VSliderScalar(const char* label,const ImVec2 size,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igVSliderScalar(label,size,data_type,p_data,p_min,p_max,format,flags); }
void TableSetBgColor(ImGuiTableBgTarget target,ImU32 color,int column_n) { igTableSetBgColor(target,color,column_n); }
void DebugTextEncoding(const char* text) { igDebugTextEncoding(text); }
void SetCursorScreenPos(const ImVec2 pos) { igSetCursorScreenPos(pos); }
bool ColorEdit3(const char* label,float col[3],ImGuiColorEditFlags flags) { return igColorEdit3(label,col,flags); }
void ImGuiIO_AddMousePosEvent(ImGuiIO* self,float x,float y) { ImGuiIO_AddMousePosEvent(self,x,y); }
float GetColumnWidth(int column_index) { return igGetColumnWidth(column_index); }
void ImGuiStorage_SetAllInt(ImGuiStorage* self,int val) { ImGuiStorage_SetAllInt(self,val); }
void PopItemWidth() { igPopItemWidth(); }
void ImFontAtlas_ClearTexData(ImFontAtlas* self) { ImFontAtlas_ClearTexData(self); }
void ShowStyleEditor(ImGuiStyle* ref) { igShowStyleEditor(ref); }
void ImDrawList_PushClipRect(ImDrawList* self,const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { ImDrawList_PushClipRect(self,clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void GetWindowContentRegionMax(ImVec2 *pOut) { igGetWindowContentRegionMax(pOut); }
bool IsItemDeactivated() { return igIsItemDeactivated(); }
bool ImFontAtlas_GetMouseCursorTexData(ImFontAtlas* self,ImGuiMouseCursor cursor,ImVec2* out_offset,ImVec2* out_size,ImVec2 out_uv_border[2],ImVec2 out_uv_fill[2]) { return ImFontAtlas_GetMouseCursorTexData(self,cursor,out_offset,out_size,out_uv_border,out_uv_fill); }
ImFontAtlasCustomRect* ImFontAtlas_GetCustomRectByIndex(ImFontAtlas* self,int index) { return ImFontAtlas_GetCustomRectByIndex(self,index); }
void ImGuiIO_AddInputCharacterUTF16(ImGuiIO* self,ImWchar16 c) { ImGuiIO_AddInputCharacterUTF16(self,c); }
void SetColorEditOptions(ImGuiColorEditFlags flags) { igSetColorEditOptions(flags); }
ImGuiContext* CreateContext(ImFontAtlas* shared_font_atlas) { return igCreateContext(shared_font_atlas); }
void TableNextRow(ImGuiTableRowFlags row_flags,float min_row_height) { igTableNextRow(row_flags,min_row_height); }
void LogToFile(int auto_open_depth,const char* filename) { igLogToFile(auto_open_depth,filename); }
void PopFont() { igPopFont(); }
void TableSetupScrollFreeze(int cols,int rows) { igTableSetupScrollFreeze(cols,rows); }
ImGuiMouseCursor GetMouseCursor() { return igGetMouseCursor(); }
void PopStyleVar(int count) { igPopStyleVar(count); }
void ImGuiIO_ClearInputCharacters(ImGuiIO* self) { ImGuiIO_ClearInputCharacters(self); }
bool InputScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_step,const void* p_step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputScalar(label,data_type,p_data,p_step,p_step_fast,format,flags); }
ImFont* ImFontAtlas_AddFont(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFont(self,font_cfg); }
bool IsWindowDocked() { return igIsWindowDocked(); }
void ImGuiPayload_Clear(ImGuiPayload* self) { ImGuiPayload_Clear(self); }
float CalcItemWidth() { return igCalcItemWidth(); }
void GetItemRectMin(ImVec2 *pOut) { igGetItemRectMin(pOut); }
void ImGuiTextFilter_Clear(ImGuiTextFilter* self) { ImGuiTextFilter_Clear(self); }
bool IsKeyReleased(ImGuiKey key) { return igIsKeyReleased(key); }
void StyleColorsDark(ImGuiStyle* dst) { igStyleColorsDark(dst); }
void ImDrawList_AddNgonFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddNgonFilled(self,center,radius,col,num_segments); }
void EndTable() { igEndTable(); }
ImFont* ImFontAtlas_AddFontFromMemoryTTF(ImFontAtlas* self,void* font_data,int font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryTTF(self,font_data,font_size,size_pixels,font_cfg,glyph_ranges); }
void ImFont_GrowIndex(ImFont* self,int new_size) { ImFont_GrowIndex(self,new_size); }
void ImDrawList_PathLineToMergeDuplicate(ImDrawList* self,const ImVec2 pos) { ImDrawList_PathLineToMergeDuplicate(self,pos); }
bool BeginTabBar(const char* str_id,ImGuiTabBarFlags flags) { return igBeginTabBar(str_id,flags); }
void ImDrawList_AddCallback(ImDrawList* self,ImDrawCallback callback,void* callback_data) { ImDrawList_AddCallback(self,callback,callback_data); }
void ImDrawList_PathBezierCubicCurveTo(ImDrawList* self,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,int num_segments) { ImDrawList_PathBezierCubicCurveTo(self,p2,p3,p4,num_segments); }
void SetMouseCursor(ImGuiMouseCursor cursor_type) { igSetMouseCursor(cursor_type); }
void ImDrawList_ChannelsSplit(ImDrawList* self,int count) { ImDrawList_ChannelsSplit(self,count); }
void AlignTextToFramePadding() { igAlignTextToFramePadding(); }
int TableGetColumnCount() { return igTableGetColumnCount(); }
void* ImGuiStorage_GetVoidPtr(ImGuiStorage* self,ImGuiID key) { return ImGuiStorage_GetVoidPtr(self,key); }
void ImGuiIO_SetKeyEventNativeData(ImGuiIO* self,ImGuiKey key,int native_keycode,int native_scancode,int native_legacy_index) { ImGuiIO_SetKeyEventNativeData(self,key,native_keycode,native_scancode,native_legacy_index); }
ImGuiViewport* FindViewportByID(ImGuiID id) { return igFindViewportByID(id); }
bool DragScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,float v_speed,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igDragScalarN(label,data_type,p_data,components,v_speed,p_min,p_max,format,flags); }
float GetWindowHeight() { return igGetWindowHeight(); }
void Unindent(float indent_w) { igUnindent(indent_w); }
ImDrawList* ImDrawList_CloneOutput(ImDrawList* self) { return ImDrawList_CloneOutput(self); }
void SetNextFrameWantCaptureMouse(bool want_capture_mouse) { igSetNextFrameWantCaptureMouse(want_capture_mouse); }
float GetScrollMaxY() { return igGetScrollMaxY(); }
void ImDrawListSplitter_SetCurrentChannel(ImDrawListSplitter* self,ImDrawList* draw_list,int channel_idx) { ImDrawListSplitter_SetCurrentChannel(self,draw_list,channel_idx); }
bool DragFloat4(const char* label,float v[4],float v_speed,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igDragFloat4(label,v,v_speed,v_min,v_max,format,flags); }
void GetContentRegionAvail(ImVec2 *pOut) { igGetContentRegionAvail(pOut); }
const char* SaveIniSettingsToMemory(size_t* out_ini_size) { return igSaveIniSettingsToMemory(out_ini_size); }
void SetItemDefaultFocus() { igSetItemDefaultFocus(); }
bool SliderFloat(const char* label,float* v,float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat(label,v,v_min,v_max,format,flags); }
void ImDrawList_PrimRectUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 uv_a,const ImVec2 uv_b,ImU32 col) { ImDrawList_PrimRectUV(self,a,b,uv_a,uv_b,col); }
bool Begin(const char* name,bool* p_open,ImGuiWindowFlags flags) { return igBegin(name,p_open,flags); }
bool InputFloat4(const char* label,float v[4],const char* format,ImGuiInputTextFlags flags) { return igInputFloat4(label,v,format,flags); }
void GetWindowSize(ImVec2 *pOut) { igGetWindowSize(pOut); }
void ImGuiIO_AddKeyAnalogEvent(ImGuiIO* self,ImGuiKey key,bool down,float v) { ImGuiIO_AddKeyAnalogEvent(self,key,down,v); }
void ImColor_SetHSV(ImColor* self,float h,float s,float v,float a) { ImColor_SetHSV(self,h,s,v,a); }
void SetNextWindowPos(const ImVec2 pos,ImGuiCond cond,const ImVec2 pivot) { igSetNextWindowPos(pos,cond,pivot); }
void SetScrollHereX(float center_x_ratio) { igSetScrollHereX(center_x_ratio); }
void SetNextWindowClass(const ImGuiWindowClass* window_class) { igSetNextWindowClass(window_class); }
ImDrawData* GetDrawData() { return igGetDrawData(); }
float GetScrollX() { return igGetScrollX(); }
void ImDrawData_Clear(ImDrawData* self) { ImDrawData_Clear(self); }
void ImDrawList_AddCircle(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments,float thickness) { ImDrawList_AddCircle(self,center,radius,col,num_segments,thickness); }
void ImGuiIO_SetAppAcceptingEvents(ImGuiIO* self,bool accepting_events) { ImGuiIO_SetAppAcceptingEvents(self,accepting_events); }
ImGuiStyle* GetStyle() { return igGetStyle(); }
void PopClipRect() { igPopClipRect(); }
void ImDrawList_PrimWriteVtx(ImDrawList* self,const ImVec2 pos,const ImVec2 uv,ImU32 col) { ImDrawList_PrimWriteVtx(self,pos,uv,col); }
void ImGuiInputTextCallbackData_InsertChars(ImGuiInputTextCallbackData* self,int pos,const char* text,const char* text_end) { ImGuiInputTextCallbackData_InsertChars(self,pos,text,text_end); }
ImGuiPlatformIO* GetPlatformIO() { return igGetPlatformIO(); }
double GetTime() { return igGetTime(); }
void LogToClipboard(int auto_open_depth) { igLogToClipboard(auto_open_depth); }
void SetItemAllowOverlap() { igSetItemAllowOverlap(); }
int TableGetRowIndex() { return igTableGetRowIndex(); }
void ImDrawList_PathRect(ImDrawList* self,const ImVec2 rect_min,const ImVec2 rect_max,float rounding,ImDrawFlags flags) { ImDrawList_PathRect(self,rect_min,rect_max,rounding,flags); }
void ImGuiIO_AddFocusEvent(ImGuiIO* self,bool focused) { ImGuiIO_AddFocusEvent(self,focused); }
bool ImGuiListClipper_Step(ImGuiListClipper* self) { return ImGuiListClipper_Step(self); }
bool TabItemButton(const char* label,ImGuiTabItemFlags flags) { return igTabItemButton(label,flags); }
void ImDrawList_AddPolyline(ImDrawList* self,const ImVec2* points,int num_points,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_AddPolyline(self,points,num_points,col,flags,thickness); }
void ImDrawList_PathArcTo(ImDrawList* self,const ImVec2 center,float radius,float a_min,float a_max,int num_segments) { ImDrawList_PathArcTo(self,center,radius,a_min,a_max,num_segments); }
ImDrawListSharedData* GetDrawListSharedData() { return igGetDrawListSharedData(); }
void EndListBox() { igEndListBox(); }
void EndPopup() { igEndPopup(); }
float GetFrameHeightWithSpacing() { return igGetFrameHeightWithSpacing(); }
bool SliderAngle(const char* label,float* v_rad,float v_degrees_min,float v_degrees_max,const char* format,ImGuiSliderFlags flags) { return igSliderAngle(label,v_rad,v_degrees_min,v_degrees_max,format,flags); }
void SetKeyboardFocusHere(int offset) { igSetKeyboardFocusHere(offset); }
bool ImGuiPayload_IsPreview(ImGuiPayload* self) { return ImGuiPayload_IsPreview(self); }
bool BeginPopup(const char* str_id,ImGuiWindowFlags flags) { return igBeginPopup(str_id,flags); }
int GetFrameCount() { return igGetFrameCount(); }
ImFontConfig* ImFontConf_ImFontConfig() { return ImFontConfig_ImFontConfig(); }
float GetTreeNodeToLabelSpacing() { return igGetTreeNodeToLabelSpacing(); }
void LogFinish() { igLogFinish(); }
bool InputTextWithHint(const char* label,const char* hint,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextWithHint(label,hint,buf,buf_size,flags,callback,user_data); }
void ImDrawList_PrimRect(ImDrawList* self,const ImVec2 a,const ImVec2 b,ImU32 col) { ImDrawList_PrimRect(self,a,b,col); }
void ImGuiListClipper_Begin(ImGuiListClipper* self,int items_count,float items_height) { ImGuiListClipper_Begin(self,items_count,items_height); }
void EndDisabled() { igEndDisabled(); }
ImFont* ImFontAtlas_AddFontDefault(ImFontAtlas* self,const ImFontConfig* font_cfg) { return ImFontAtlas_AddFontDefault(self,font_cfg); }
void SetNextWindowContentSize(const ImVec2 size) { igSetNextWindowContentSize(size); }
void ImDrawList_ChannelsSetCurrent(ImDrawList* self,int n) { ImDrawList_ChannelsSetCurrent(self,n); }
bool IsKeyDown(ImGuiKey key) { return igIsKeyDown(key); }
void ImGuiIO_AddMouseWheelEvent(ImGuiIO* self,float wh_x,float wh_y) { ImGuiIO_AddMouseWheelEvent(self,wh_x,wh_y); }
bool BeginTabItem(const char* label,bool* p_open,ImGuiTabItemFlags flags) { return igBeginTabItem(label,p_open,flags); }
bool ImGuiTextFilter_IsActive(ImGuiTextFilter* self) { return ImGuiTextFilter_IsActive(self); }
const ImWchar* ImFontAtlas_GetGlyphRangesCyrillic(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesCyrillic(self); }
void EndTabItem() { igEndTabItem(); }
bool ImGuiPayload_IsDelivery(ImGuiPayload* self) { return ImGuiPayload_IsDelivery(self); }
bool BeginMenu(const char* label,bool enabled) { return igBeginMenu(label,enabled); }
ImGuiPayload* ImGuiPayload_ImGuiPayload() { return ImGuiPayload_ImGuiPayload(); }
void ImGuiStorage_SetVoidPtr(ImGuiStorage* self,ImGuiID key,void* val) { ImGuiStorage_SetVoidPtr(self,key,val); }
void EndChildFrame() { igEndChildFrame(); }
void ResetMouseDragDelta(ImGuiMouseButton button) { igResetMouseDragDelta(button); }
ImFont* ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(ImFontAtlas* self,const char* compressed_font_data_base85,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(self,compressed_font_data_base85,size_pixels,font_cfg,glyph_ranges); }
const char* ImFont_CalcWordWrapPositionA(ImFont* self,float scale,const char* text,const char* text_end,float wrap_width) { return ImFont_CalcWordWrapPositionA(self,scale,text,text_end,wrap_width); }
const char* ImGuiTextBuffer_c_str(ImGuiTextBuffer* self) { return ImGuiTextBuffer_c_str(self); }
void ImGuiViewport_GetWorkCenter(ImVec2 *pOut,ImGuiViewport* self) { ImGuiViewport_GetWorkCenter(pOut,self); }
float GetWindowDpiScale() { return igGetWindowDpiScale(); }
void ImDrawList_PathStroke(ImDrawList* self,ImU32 col,ImDrawFlags flags,float thickness) { ImDrawList_PathStroke(self,col,flags,thickness); }
void SetNextWindowBgAlpha(float alpha) { igSetNextWindowBgAlpha(alpha); }
void ImDrawList_AddCircleFilled(ImDrawList* self,const ImVec2 center,float radius,ImU32 col,int num_segments) { ImDrawList_AddCircleFilled(self,center,radius,col,num_segments); }
ImFontAtlasCustomRect* ImFontAtlasCustomRect_ImFontAtlasCustomRect() { return ImFontAtlasCustomRect_ImFontAtlasCustomRect(); }
void ImGuiListClipper_End(ImGuiListClipper* self) { ImGuiListClipper_End(self); }
int GetMouseClickedCount(ImGuiMouseButton button) { return igGetMouseClickedCount(button); }
void ImGuiStorage_SetBool(ImGuiStorage* self,ImGuiID key,bool val) { ImGuiStorage_SetBool(self,key,val); }
bool IsMouseDragging(ImGuiMouseButton button,float lock_threshold) { return igIsMouseDragging(button,lock_threshold); }
bool SliderFloat3(const char* label,float v[3],float v_min,float v_max,const char* format,ImGuiSliderFlags flags) { return igSliderFloat3(label,v,v_min,v_max,format,flags); }
void ImDrawList_AddQuadFilled(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col) { ImDrawList_AddQuadFilled(self,p1,p2,p3,p4,col); }
bool DragIntRange2(const char* label,int* v_current_min,int* v_current_max,float v_speed,int v_min,int v_max,const char* format,const char* format_max,ImGuiSliderFlags flags) { return igDragIntRange2(label,v_current_min,v_current_max,v_speed,v_min,v_max,format,format_max,flags); }
ImGuiWindowClass* ImGuiWindowClass_ImGuiWindowClass() { return ImGuiWindowClass_ImGuiWindowClass(); }
bool IsAnyItemFocused() { return igIsAnyItemFocused(); }
void ImDrawList_AddRect(ImDrawList* self,const ImVec2 p_min,const ImVec2 p_max,ImU32 col,float rounding,ImDrawFlags flags,float thickness) { ImDrawList_AddRect(self,p_min,p_max,col,rounding,flags,thickness); }
void ImDrawList_PathClear(ImDrawList* self) { ImDrawList_PathClear(self); }
bool SliderInt2(const char* label,int v[2],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt2(label,v,v_min,v_max,format,flags); }
void StyleColorsClassic(ImGuiStyle* dst) { igStyleColorsClassic(dst); }
const ImWchar* ImFontAtlas_GetGlyphRangesChineseFull(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesChineseFull(self); }
ImGuiTableSortSpecs* ImGuiTableSortSpecs_ImGuiTableSortSpecs() { return ImGuiTableSortSpecs_ImGuiTableSortSpecs(); }
void PopButtonRepeat() { igPopButtonRepeat(); }
void EndChild() { igEndChild(); }
bool VSliderInt(const char* label,const ImVec2 size,int* v,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igVSliderInt(label,size,v,v_min,v_max,format,flags); }
ImFont* ImFontAtlas_AddFontFromFileTTF(ImFontAtlas* self,const char* filename,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromFileTTF(self,filename,size_pixels,font_cfg,glyph_ranges); }
void BeginGroup() { igBeginGroup(); }
bool ImageButton(const char* str_id,ImTextureID user_texture_id,const ImVec2 size,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 bg_col,const ImVec4 tint_col) { return igImageButton(str_id,user_texture_id,size,uv0,uv1,bg_col,tint_col); }
void TableSetupColumn(const char* label,ImGuiTableColumnFlags flags,float init_width_or_weight,ImGuiID user_id) { igTableSetupColumn(label,flags,init_width_or_weight,user_id); }
void PushItemWidth(float item_width) { igPushItemWidth(item_width); }
const ImGuiPayload* AcceptDragDropPayload(const char* type,ImGuiDragDropFlags flags) { return igAcceptDragDropPayload(type,flags); }
bool IsMouseDoubleClicked(ImGuiMouseButton button) { return igIsMouseDoubleClicked(button); }
ImFontAtlas* ImFontAtlas_ImFontAtlas() { return ImFontAtlas_ImFontAtlas(); }
float GetTextLineHeight() { return igGetTextLineHeight(); }
const ImWchar* ImFontAtlas_GetGlyphRangesThai(ImFontAtlas* self) { return ImFontAtlas_GetGlyphRangesThai(self); }
ImGuiTableSortSpecs* TableGetSortSpecs() { return igTableGetSortSpecs(); }
bool TableSetColumnIndex(int column_n) { return igTableSetColumnIndex(column_n); }
void ImDrawListSplitter_Merge(ImDrawListSplitter* self,ImDrawList* draw_list) { ImDrawListSplitter_Merge(self,draw_list); }
void ImFontGlyphRangesBuilder_Clear(ImFontGlyphRangesBuilder* self) { ImFontGlyphRangesBuilder_Clear(self); }
bool SliderScalarN(const char* label,ImGuiDataType data_type,void* p_data,int components,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalarN(label,data_type,p_data,components,p_min,p_max,format,flags); }
ImGuiTextFilter* ImGuiTextFilter_ImGuiTextFilter(const char* default_filter) { return ImGuiTextFilter_ImGuiTextFilter(default_filter); }
bool ImGuiTextFilter_PassFilter(ImGuiTextFilter* self,const char* text,const char* text_end) { return ImGuiTextFilter_PassFilter(self,text,text_end); }
void ImGuiStorage_BuildSortByKey(ImGuiStorage* self) { ImGuiStorage_BuildSortByKey(self); }
void BeginTooltip() { igBeginTooltip(); }
int GetKeyIndex(ImGuiKey key) { return igGetKeyIndex(key); }
float ImFont_GetCharAdvance(ImFont* self,ImWchar c) { return ImFont_GetCharAdvance(self,c); }
ImGuiIO* GetIO() { return igGetIO(); }
void SetNextWindowSize(const ImVec2 size,ImGuiCond cond) { igSetNextWindowSize(size,cond); }
void SetNextWindowViewport(ImGuiID viewport_id) { igSetNextWindowViewport(viewport_id); }
bool ShowStyleSelector(const char* label) { return igShowStyleSelector(label); }
const char* GetClipboardText() { return igGetClipboardText(); }
void CalcTextSize(ImVec2 *pOut,const char* text,const char* text_end,bool hide_text_after_double_hash,float wrap_width) { igCalcTextSize(pOut,text,text_end,hide_text_after_double_hash,wrap_width); }
void ImFontGlyphRangesBuilder_SetBit(ImFontGlyphRangesBuilder* self,size_t n) { ImFontGlyphRangesBuilder_SetBit(self,n); }
void EndTooltip() { igEndTooltip(); }
void GetCursorStartPos(ImVec2 *pOut) { igGetCursorStartPos(pOut); }
void NextColumn() { igNextColumn(); }
ImFont* ImFontAtlas_AddFontFromMemoryCompressedTTF(ImFontAtlas* self,const void* compressed_font_data,int compressed_font_size,float size_pixels,const ImFontConfig* font_cfg,const ImWchar* glyph_ranges) { return ImFontAtlas_AddFontFromMemoryCompressedTTF(self,compressed_font_data,compressed_font_size,size_pixels,font_cfg,glyph_ranges); }
bool InputDouble(const char* label,double* v,double step,double step_fast,const char* format,ImGuiInputTextFlags flags) { return igInputDouble(label,v,step,step_fast,format,flags); }
ImGuiTableColumnFlags TableGetColumnFlags(int column_n) { return igTableGetColumnFlags(column_n); }
void ImFont_CalcTextSizeA(ImVec2 *pOut,ImFont* self,float size,float max_width,float wrap_width,const char* text_begin,const char* text_end,const char** remaining) { ImFont_CalcTextSizeA(pOut,self,size,max_width,wrap_width,text_begin,text_end,remaining); }
void ImGuiIO_AddInputCharactersUTF8(ImGuiIO* self,const char* str) { ImGuiIO_AddInputCharactersUTF8(self,str); }
bool ImGuiTextFilter_Draw(ImGuiTextFilter* self,const char* label,float width) { return ImGuiTextFilter_Draw(self,label,width); }
ImDrawList* ImDrawList_ImDrawList(const ImDrawListSharedData* shared_data) { return ImDrawList_ImDrawList(shared_data); }
void Image(ImTextureID user_texture_id,const ImVec2 size,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 tint_col,const ImVec4 border_col) { igImage(user_texture_id,size,uv0,uv1,tint_col,border_col); }
void Bullet() { igBullet(); }
void EndMenu() { igEndMenu(); }
void ImFontAtlas_ClearInputData(ImFontAtlas* self) { ImFontAtlas_ClearInputData(self); }
ImGuiInputTextCallbackData* ImGuiInputTextCallbackData_ImGuiInputTextCallbackData() { return ImGuiInputTextCallbackData_ImGuiInputTextCallbackData(); }
float ImGuiStorage_GetFloat(ImGuiStorage* self,ImGuiID key,float default_val) { return ImGuiStorage_GetFloat(self,key,default_val); }
ImFont* GetFont() { return igGetFont(); }
bool IsMouseDown(ImGuiMouseButton button) { return igIsMouseDown(button); }
bool BeginPopupContextItem(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextItem(str_id,popup_flags); }
void ImColor_HSV(ImColor *pOut,float h,float s,float v,float a) { ImColor_HSV(pOut,h,s,v,a); }
void EndCombo() { igEndCombo(); }
ImGuiTextBuffer* ImGuiTextBuffer_ImGuiTextBuffer() { return ImGuiTextBuffer_ImGuiTextBuffer(); }
void ImFontAtlas_CalcCustomRectUV(ImFontAtlas* self,const ImFontAtlasCustomRect* rect,ImVec2* out_uv_min,ImVec2* out_uv_max) { ImFontAtlas_CalcCustomRectUV(self,rect,out_uv_min,out_uv_max); }
void LogButtons() { igLogButtons(); }
void GetMousePosOnOpeningCurrentPopup(ImVec2 *pOut) { igGetMousePosOnOpeningCurrentPopup(pOut); }
void LoadIniSettingsFromDisk(const char* ini_filename) { igLoadIniSettingsFromDisk(ini_filename); }
const ImGuiPayload* GetDragDropPayload() { return igGetDragDropPayload(); }
bool DragInt2(const char* label,int v[2],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt2(label,v,v_speed,v_min,v_max,format,flags); }
void EndDragDropTarget() { igEndDragDropTarget(); }
void EndMenuBar() { igEndMenuBar(); }
void ProgressBar(float fraction,const ImVec2 size_arg,const char* overlay) { igProgressBar(fraction,size_arg,overlay); }
bool InputText(const char* label,char* buf,size_t buf_size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputText(label,buf,buf_size,flags,callback,user_data); }
void SetScrollHereY(float center_y_ratio) { igSetScrollHereY(center_y_ratio); }
bool IsItemHovered(ImGuiHoveredFlags flags) { return igIsItemHovered(flags); }
void Dummy(const ImVec2 size) { igDummy(size); }
int* ImGuiStorage_GetIntRef(ImGuiStorage* self,ImGuiID key,int default_val) { return ImGuiStorage_GetIntRef(self,key,default_val); }
void ShowUserGuide() { igShowUserGuide(); }
bool InputInt(const char* label,int* v,int step,int step_fast,ImGuiInputTextFlags flags) { return igInputInt(label,v,step,step_fast,flags); }
void ImDrawList_PopClipRect(ImDrawList* self) { ImDrawList_PopClipRect(self); }
ImGuiPlatformMonitor* ImGuiPlatformMonitor_ImGuiPlatformMonitor() { return ImGuiPlatformMonitor_ImGuiPlatformMonitor(); }
bool ImGuiInputTextCallbackData_HasSelection(ImGuiInputTextCallbackData* self) { return ImGuiInputTextCallbackData_HasSelection(self); }
void ImDrawList_AddLine(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,ImU32 col,float thickness) { ImDrawList_AddLine(self,p1,p2,col,thickness); }
bool InputTextMultiline(const char* label,char* buf,size_t buf_size,const ImVec2 size,ImGuiInputTextFlags flags,ImGuiInputTextCallback callback,void* user_data) { return igInputTextMultiline(label,buf,buf_size,size,flags,callback,user_data); }
float GetScrollY() { return igGetScrollY(); }
bool IsItemDeactivatedAfterEdit() { return igIsItemDeactivatedAfterEdit(); }
void ImDrawList_AddImageQuad(ImDrawList* self,ImTextureID user_texture_id,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,const ImVec2 uv1,const ImVec2 uv2,const ImVec2 uv3,const ImVec2 uv4,ImU32 col) { ImDrawList_AddImageQuad(self,user_texture_id,p1,p2,p3,p4,uv1,uv2,uv3,uv4,col); }
void ImGuiIO_AddInputCharacter(ImGuiIO* self,unsigned int c) { ImGuiIO_AddInputCharacter(self,c); }
bool BeginDragDropSource(ImGuiDragDropFlags flags) { return igBeginDragDropSource(flags); }
bool ColorEdit4(const char* label,float col[4],ImGuiColorEditFlags flags) { return igColorEdit4(label,col,flags); }
bool IsItemClicked(ImGuiMouseButton mouse_button) { return igIsItemClicked(mouse_button); }
float GetColumnOffset(int column_index) { return igGetColumnOffset(column_index); }
void GetCursorPos(ImVec2 *pOut) { igGetCursorPos(pOut); }
void MemFree(void* ptr) { igMemFree(ptr); }
ImGuiContext* GetCurrentContext() { return igGetCurrentContext(); }
void ImDrawList_AddQuad(ImDrawList* self,const ImVec2 p1,const ImVec2 p2,const ImVec2 p3,const ImVec2 p4,ImU32 col,float thickness) { ImDrawList_AddQuad(self,p1,p2,p3,p4,col,thickness); }
void SetStateStorage(ImGuiStorage* storage) { igSetStateStorage(storage); }
void SetCursorPos(const ImVec2 local_pos) { igSetCursorPos(local_pos); }
bool ImFontAtlas_IsBuilt(ImFontAtlas* self) { return ImFontAtlas_IsBuilt(self); }
bool BeginDragDropTarget() { return igBeginDragDropTarget(); }
bool ImFontAtlasCustomRect_IsPacked(ImFontAtlasCustomRect* self) { return ImFontAtlasCustomRect_IsPacked(self); }
void ShowStackToolWindow(bool* p_open) { igShowStackToolWindow(p_open); }
float* ImGuiStorage_GetFloatRef(ImGuiStorage* self,ImGuiID key,float default_val) { return ImGuiStorage_GetFloatRef(self,key,default_val); }
void PopTextWrapPos() { igPopTextWrapPos(); }
bool SliderInt3(const char* label,int v[3],int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igSliderInt3(label,v,v_min,v_max,format,flags); }
void ImFontGlyphRangesBuilder_AddText(ImFontGlyphRangesBuilder* self,const char* text,const char* text_end) { ImFontGlyphRangesBuilder_AddText(self,text,text_end); }
bool InputFloat3(const char* label,float v[3],const char* format,ImGuiInputTextFlags flags) { return igInputFloat3(label,v,format,flags); }
void SetWindowFontScale(float scale) { igSetWindowFontScale(scale); }
void GetCursorScreenPos(ImVec2 *pOut) { igGetCursorScreenPos(pOut); }
void ImDrawListSplitter_ClearFreeMemory(ImDrawListSplitter* self) { ImDrawListSplitter_ClearFreeMemory(self); }
void PushFont(ImFont* font) { igPushFont(font); }
bool IsWindowAppearing() { return igIsWindowAppearing(); }
bool ImGuiPayload_IsDataType(ImGuiPayload* self,const char* type) { return ImGuiPayload_IsDataType(self,type); }
bool ArrowButton(const char* str_id,ImGuiDir dir) { return igArrowButton(str_id,dir); }
void Spacing() { igSpacing(); }
void Columns(int count,const char* id,bool border) { igColumns(count,id,border); }
void ImFont_AddGlyph(ImFont* self,const ImFontConfig* src_cfg,ImWchar c,float x0,float y0,float x1,float y1,float u0,float v0,float u1,float v1,float advance_x) { ImFont_AddGlyph(self,src_cfg,c,x0,y0,x1,y1,u0,v0,u1,v1,advance_x); }
void TableHeadersRow() { igTableHeadersRow(); }
void LogToTTY(int auto_open_depth) { igLogToTTY(auto_open_depth); }
bool IsAnyItemHovered() { return igIsAnyItemHovered(); }
void ImDrawList_PrimWriteIdx(ImDrawList* self,ImDrawIdx idx) { ImDrawList_PrimWriteIdx(self,idx); }
void GetContentRegionMax(ImVec2 *pOut) { igGetContentRegionMax(pOut); }
ImGuiID DockSpaceOverViewport(const ImGuiViewport* viewport,ImGuiDockNodeFlags flags,const ImGuiWindowClass* window_class) { return igDockSpaceOverViewport(viewport,flags,window_class); }
void SetNextItemWidth(float item_width) { igSetNextItemWidth(item_width); }
void GetWindowPos(ImVec2 *pOut) { igGetWindowPos(pOut); }
void PopAllowKeyboardFocus() { igPopAllowKeyboardFocus(); }
void GetItemRectSize(ImVec2 *pOut) { igGetItemRectSize(pOut); }
int ImFontAtlas_AddCustomRectFontGlyph(ImFontAtlas* self,ImFont* font,ImWchar id,int width,int height,float advance_x,const ImVec2 offset) { return ImFontAtlas_AddCustomRectFontGlyph(self,font,id,width,height,advance_x,offset); }
void TreePop() { igTreePop(); }
void ImGuiIO_AddKeyEvent(ImGuiIO* self,ImGuiKey key,bool down) { ImGuiIO_AddKeyEvent(self,key,down); }
void ImGuiStorage_SetInt(ImGuiStorage* self,ImGuiID key,int val) { ImGuiStorage_SetInt(self,key,val); }
void ImDrawList_GetClipRectMax(ImVec2 *pOut,ImDrawList* self) { ImDrawList_GetClipRectMax(pOut,self); }
void ImFontAtlas_GetTexDataAsAlpha8(ImFontAtlas* self,unsigned char** out_pixels,int* out_width,int* out_height,int* out_bytes_per_pixel) { ImFontAtlas_GetTexDataAsAlpha8(self,out_pixels,out_width,out_height,out_bytes_per_pixel); }
void ImFont_ClearOutputData(ImFont* self) { ImFont_ClearOutputData(self); }
void NewFrame() { igNewFrame(); }
void ImDrawList_PopTextureID(ImDrawList* self) { ImDrawList_PopTextureID(self); }
bool ImGuiStorage_GetBool(ImGuiStorage* self,ImGuiID key,bool default_val) { return ImGuiStorage_GetBool(self,key,default_val); }
bool BeginPopupContextWindow(const char* str_id,ImGuiPopupFlags popup_flags) { return igBeginPopupContextWindow(str_id,popup_flags); }
void ColorConvertHSVtoRGB(float h,float s,float v,float* out_r,float* out_g,float* out_b) { igColorConvertHSVtoRGB(h,s,v,out_r,out_g,out_b); }
bool DragInt3(const char* label,int v[3],float v_speed,int v_min,int v_max,const char* format,ImGuiSliderFlags flags) { return igDragInt3(label,v,v_speed,v_min,v_max,format,flags); }
bool InputInt3(const char* label,int v[3],ImGuiInputTextFlags flags) { return igInputInt3(label,v,flags); }
void ImFont_AddRemapChar(ImFont* self,ImWchar dst,ImWchar src,bool overwrite_dst) { ImFont_AddRemapChar(self,dst,src,overwrite_dst); }
void ImFont_SetGlyphVisible(ImFont* self,ImWchar c,bool visible) { ImFont_SetGlyphVisible(self,c,visible); }
const char* GetKeyName(ImGuiKey key) { return igGetKeyName(key); }
void ImFontGlyphRangesBuilder_BuildRanges(ImFontGlyphRangesBuilder* self,ImVector_ImWchar* out_ranges) { ImFontGlyphRangesBuilder_BuildRanges(self,out_ranges); }
void OpenPopupOnItemClick(const char* str_id,ImGuiPopupFlags popup_flags) { igOpenPopupOnItemClick(str_id,popup_flags); }
void SaveIniSettingsToDisk(const char* ini_filename) { igSaveIniSettingsToDisk(ini_filename); }
void Separator() { igSeparator(); }
bool SliderScalar(const char* label,ImGuiDataType data_type,void* p_data,const void* p_min,const void* p_max,const char* format,ImGuiSliderFlags flags) { return igSliderScalar(label,data_type,p_data,p_min,p_max,format,flags); }
void RenderPlatformWindowsDefault(void* platform_render_arg,void* renderer_render_arg) { igRenderPlatformWindowsDefault(platform_render_arg,renderer_render_arg); }
bool IsItemEdited() { return igIsItemEdited(); }
void ShowDebugLogWindow(bool* p_open) { igShowDebugLogWindow(p_open); }
void ImDrawList_PrimQuadUV(ImDrawList* self,const ImVec2 a,const ImVec2 b,const ImVec2 c,const ImVec2 d,const ImVec2 uv_a,const ImVec2 uv_b,const ImVec2 uv_c,const ImVec2 uv_d,ImU32 col) { ImDrawList_PrimQuadUV(self,a,b,c,d,uv_a,uv_b,uv_c,uv_d,col); }
bool InputInt2(const char* label,int v[2],ImGuiInputTextFlags flags) { return igInputInt2(label,v,flags); }
bool InputInt4(const char* label,int v[4],ImGuiInputTextFlags flags) { return igInputInt4(label,v,flags); }
bool IsItemToggledOpen() { return igIsItemToggledOpen(); }
void TableHeader(const char* label) { igTableHeader(label); }
void TableSetColumnEnabled(int column_n,bool v) { igTableSetColumnEnabled(column_n,v); }
void ImDrawList_ChannelsMerge(ImDrawList* self) { ImDrawList_ChannelsMerge(self); }
ImGuiViewport* FindViewportByPlatformHandle(void* platform_handle) { return igFindViewportByPlatformHandle(platform_handle); }
void SetCursorPosY(float local_y) { igSetCursorPosY(local_y); }
ImGuiViewport* GetWindowViewport() { return igGetWindowViewport(); }
float GetCursorPosY() { return igGetCursorPosY(); }
const char* GetStyleColorName(ImGuiCol idx) { return igGetStyleColorName(idx); }
void PushClipRect(const ImVec2 clip_rect_min,const ImVec2 clip_rect_max,bool intersect_with_current_clip_rect) { igPushClipRect(clip_rect_min,clip_rect_max,intersect_with_current_clip_rect); }
void DestroyPlatformWindows() { igDestroyPlatformWindows(); }
ImTextureID ImDrawCmd_GetTexID(ImDrawCmd* self) { return ImDrawCmd_GetTexID(self); }
void ImDrawList_PrimReserve(ImDrawList* self,int idx_count,int vtx_count) { ImDrawList_PrimReserve(self,idx_count,vtx_count); }
bool ImFont_IsGlyphRangeUnused(ImFont* self,unsigned int c_begin,unsigned int c_last) { return ImFont_IsGlyphRangeUnused(self,c_begin,c_last); }
void EndGroup() { igEndGroup(); }
void ImFontAtlas_SetTexID(ImFontAtlas* self,ImTextureID id) { ImFontAtlas_SetTexID(self,id); }
void EndMainMenuBar() { igEndMainMenuBar(); }
void GetFontTexUvWhitePixel(ImVec2 *pOut) { igGetFontTexUvWhitePixel(pOut); }
ImGuiPlatformImeData* ImGuiPlatformImeData_ImGuiPlatformImeData() { return ImGuiPlatformImeData_ImGuiPlatformImeData(); }
bool IsWindowHovered(ImGuiHoveredFlags flags) { return igIsWindowHovered(flags); }
void ImGuiInputTextCallbackData_ClearSelection(ImGuiInputTextCallbackData* self) { ImGuiInputTextCallbackData_ClearSelection(self); }
void GetMouseDragDelta(ImVec2 *pOut,ImGuiMouseButton button,float lock_threshold) { igGetMouseDragDelta(pOut,button,lock_threshold); }
void SetNextWindowCollapsed(bool collapsed,ImGuiCond cond) { igSetNextWindowCollapsed(collapsed,cond); }
ImU32 ColorConvertFloat4ToU32(const ImVec4 in) { return igColorConvertFloat4ToU32(in); }
bool ColorPicker4(const char* label,float col[4],ImGuiColorEditFlags flags,const float* ref_col) { return igColorPicker4(label,col,flags,ref_col); }
