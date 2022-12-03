#include "cimplot_wrapper.h"
#include "cimplot/cimplot.h"

void PlotAlignmentData_Destroy(ImPlotAlignmentData* self) { ImPlotAlignmentData_destroy(self); }
void PlotAnnotationCollection_Destroy(ImPlotAnnotationCollection* self) { ImPlotAnnotationCollection_destroy(self); }
void PlotAxis_Destroy(ImPlotAxis* self) { ImPlotAxis_destroy(self); }
void PlotColormapData_Destroy(ImPlotColormapData* self) { ImPlotColormapData_destroy(self); }
void PlotDateTimeSpec_Destroy(ImPlotDateTimeSpec* self) { ImPlotDateTimeSpec_destroy(self); }
ImPlotInputMap* PlotInputMap_ImPlotInputMap() { return ImPlotInputMap_ImPlotInputMap(); }
void PlotInputMap_Destroy(ImPlotInputMap* self) { ImPlotInputMap_destroy(self); }
void PlotItemGroup_Destroy(ImPlotItemGroup* self) { ImPlotItemGroup_destroy(self); }
void PlotLegend_Destroy(ImPlotLegend* self) { ImPlotLegend_destroy(self); }
void PlotNextItemData_Destroy(ImPlotNextItemData* self) { ImPlotNextItemData_destroy(self); }
void PlotNextPlotData_Destroy(ImPlotNextPlotData* self) { ImPlotNextPlotData_destroy(self); }
void PlotPlot_Destroy(ImPlotPlot* self) { ImPlotPlot_destroy(self); }
void PlotPointError_Destroy(ImPlotPointError* self) { ImPlotPointError_destroy(self); }
ImPlotPoint* PlotPoint_ImPlotPoint_Nil() { return ImPlotPoint_ImPlotPoint_Nil(); }
ImPlotPoint* PlotPoint_ImPlotPoint_Vec2(const ImVec2 p) { return ImPlotPoint_ImPlotPoint_Vec2(p); }
ImPlotPoint* PlotPoint_ImPlotPoint_double(double _x,double _y) { return ImPlotPoint_ImPlotPoint_double(_x,_y); }
void PlotPoint_Destroy(ImPlotPoint* self) { ImPlotPoint_destroy(self); }
double PlotRange_Clamp(ImPlotRange* self,double value) { return ImPlotRange_Clamp(self,value); }
bool PlotRange_Contains(ImPlotRange* self,double value) { return ImPlotRange_Contains(self,value); }
ImPlotRange* PlotRange_ImPlotRange_Nil() { return ImPlotRange_ImPlotRange_Nil(); }
ImPlotRange* PlotRange_ImPlotRange_double(double _min,double _max) { return ImPlotRange_ImPlotRange_double(_min,_max); }
double PlotRange_Size(ImPlotRange* self) { return ImPlotRange_Size(self); }
void PlotRange_Destroy(ImPlotRange* self) { ImPlotRange_destroy(self); }
void PlotRect_Clamp_PlotPoInt(ImPlotPoint *pOut,ImPlotRect* self,const ImPlotPoint p) { ImPlotRect_Clamp_PlotPoInt(pOut,self,p); }
void PlotRect_Clamp_double(ImPlotPoint *pOut,ImPlotRect* self,double x,double y) { ImPlotRect_Clamp_double(pOut,self,x,y); }
bool PlotRect_Contains_PlotPoInt(ImPlotRect* self,const ImPlotPoint p) { return ImPlotRect_Contains_PlotPoInt(self,p); }
bool PlotRect_Contains_double(ImPlotRect* self,double x,double y) { return ImPlotRect_Contains_double(self,x,y); }
ImPlotRect* PlotRect_ImPlotRect_Nil() { return ImPlotRect_ImPlotRect_Nil(); }
ImPlotRect* PlotRect_ImPlotRect_double(double x_min,double x_max,double y_min,double y_max) { return ImPlotRect_ImPlotRect_double(x_min,x_max,y_min,y_max); }
void PlotRect_Max(ImPlotPoint *pOut,ImPlotRect* self) { ImPlotRect_Max(pOut,self); }
void PlotRect_Min(ImPlotPoint *pOut,ImPlotRect* self) { ImPlotRect_Min(pOut,self); }
void PlotRect_Size(ImPlotPoint *pOut,ImPlotRect* self) { ImPlotRect_Size(pOut,self); }
void PlotRect_Destroy(ImPlotRect* self) { ImPlotRect_destroy(self); }
ImPlotStyle* PlotStyle_ImPlotStyle() { return ImPlotStyle_ImPlotStyle(); }
void PlotStyle_Destroy(ImPlotStyle* self) { ImPlotStyle_destroy(self); }
void PlotSubplot_Destroy(ImPlotSubplot* self) { ImPlotSubplot_destroy(self); }
void PlotTagCollection_Destroy(ImPlotTagCollection* self) { ImPlotTagCollection_destroy(self); }
void PlotTick_Destroy(ImPlotTick* self) { ImPlotTick_destroy(self); }
void PlotTicker_Destroy(ImPlotTicker* self) { ImPlotTicker_destroy(self); }
void PlotTime_Destroy(ImPlotTime* self) { ImPlotTime_destroy(self); }
ImPlotColormap Plot_AddColormap_U32PtrV(const char* name,const ImU32* cols,int size,bool qual) { return ImPlot_AddColormap_U32Ptr(name,cols,size,qual); }
ImPlotColormap Plot_AddColormap_Vec4PtrV(const char* name,const ImVec4* cols,int size,bool qual) { return ImPlot_AddColormap_Vec4Ptr(name,cols,size,qual); }
void Plot_Annotation_BoolV(double x,double y,const ImVec4 col,const ImVec2 pix_offset,bool clamp,bool round) { ImPlot_Annotation_Bool(x,y,col,pix_offset,clamp,round); }
void Plot_Annotation_Str(double x,double y,const ImVec4 col,const ImVec2 pix_offset,bool clamp,const char* fmt) { ImPlot_Annotation_Str(x,y,col,pix_offset,clamp,fmt); }
bool Plot_BeginAlignedPlotsV(const char* group_id,bool vertical) { return ImPlot_BeginAlignedPlots(group_id,vertical); }
bool Plot_BeginDragDropSourceAxisV(ImAxis axis,ImGuiDragDropFlags flags) { return ImPlot_BeginDragDropSourceAxis(axis,flags); }
bool Plot_BeginDragDropSourceItemV(const char* label_id,ImGuiDragDropFlags flags) { return ImPlot_BeginDragDropSourceItem(label_id,flags); }
bool Plot_BeginDragDropSourcePlotV(ImGuiDragDropFlags flags) { return ImPlot_BeginDragDropSourcePlot(flags); }
bool Plot_BeginDragDropTargetAxis(ImAxis axis) { return ImPlot_BeginDragDropTargetAxis(axis); }
bool Plot_BeginDragDropTargetLegend() { return ImPlot_BeginDragDropTargetLegend(); }
bool Plot_BeginDragDropTargetPlot() { return ImPlot_BeginDragDropTargetPlot(); }
bool Plot_BeginLegendPopupV(const char* label_id,ImGuiMouseButton mouse_button) { return ImPlot_BeginLegendPopup(label_id,mouse_button); }
bool Plot_BeginPlotV(const char* title_id,const ImVec2 size,ImPlotFlags flags) { return ImPlot_BeginPlot(title_id,size,flags); }
bool Plot_BeginSubplotsV(const char* title_id,int rows,int cols,const ImVec2 size,ImPlotSubplotFlags flags,float* row_ratios,float* col_ratios) { return ImPlot_BeginSubplots(title_id,rows,cols,size,flags,row_ratios,col_ratios); }
void Plot_BustColorCacheV(const char* plot_title_id) { ImPlot_BustColorCache(plot_title_id); }
void Plot_CancelPlotSelection() { ImPlot_CancelPlotSelection(); }
bool Plot_ColormapButtonV(const char* label,const ImVec2 size,ImPlotColormap cmap) { return ImPlot_ColormapButton(label,size,cmap); }
void Plot_ColormapIcon(ImPlotColormap cmap) { ImPlot_ColormapIcon(cmap); }
void Plot_ColormapScaleV(const char* label,double scale_min,double scale_max,const ImVec2 size,const char* format,ImPlotColormapScaleFlags flags,ImPlotColormap cmap) { ImPlot_ColormapScale(label,scale_min,scale_max,size,format,flags,cmap); }
bool Plot_ColormapSliderV(const char* label,float* t,ImVec4* out,const char* format,ImPlotColormap cmap) { return ImPlot_ColormapSlider(label,t,out,format,cmap); }
ImPlotContext* Plot_CreateContext() { return ImPlot_CreateContext(); }
void Plot_DestroyContextV(ImPlotContext* ctx) { ImPlot_DestroyContext(ctx); }
bool Plot_DragLineXV(int id,double* x,const ImVec4 col,float thickness,ImPlotDragToolFlags flags) { return ImPlot_DragLineX(id,x,col,thickness,flags); }
bool Plot_DragLineYV(int id,double* y,const ImVec4 col,float thickness,ImPlotDragToolFlags flags) { return ImPlot_DragLineY(id,y,col,thickness,flags); }
bool Plot_DragPointV(int id,double* x,double* y,const ImVec4 col,float size,ImPlotDragToolFlags flags) { return ImPlot_DragPoint(id,x,y,col,size,flags); }
bool Plot_DragRectV(int id,double* x_min,double* y_min,double* x_max,double* y_max,const ImVec4 col,ImPlotDragToolFlags flags) { return ImPlot_DragRect(id,x_min,y_min,x_max,y_max,col,flags); }
void Plot_EndAlignedPlots() { ImPlot_EndAlignedPlots(); }
void Plot_EndDragDropSource() { ImPlot_EndDragDropSource(); }
void Plot_EndDragDropTarget() { ImPlot_EndDragDropTarget(); }
void Plot_EndLegendPopup() { ImPlot_EndLegendPopup(); }
void Plot_EndPlot() { ImPlot_EndPlot(); }
void Plot_EndSubplots() { ImPlot_EndSubplots(); }
void Plot_GetColormapColorV(ImVec4 *pOut,int idx,ImPlotColormap cmap) { ImPlot_GetColormapColor(pOut,idx,cmap); }
int Plot_GetColormapCount() { return ImPlot_GetColormapCount(); }
ImPlotColormap Plot_GetColormapIndex(const char* name) { return ImPlot_GetColormapIndex(name); }
const char* Plot_GetColormapName(ImPlotColormap cmap) { return ImPlot_GetColormapName(cmap); }
int Plot_GetColormapSizeV(ImPlotColormap cmap) { return ImPlot_GetColormapSize(cmap); }
ImPlotContext* Plot_GetCurrentContext() { return ImPlot_GetCurrentContext(); }
ImPlotInputMap* Plot_GetInputMap() { return ImPlot_GetInputMap(); }
void Plot_GetLastItemColor(ImVec4 *pOut) { ImPlot_GetLastItemColor(pOut); }
const char* Plot_GetMarkerName(ImPlotMarker idx) { return ImPlot_GetMarkerName(idx); }
ImDrawList* Plot_GetPlotDrawList() { return ImPlot_GetPlotDrawList(); }
ImPlotRect Plot_GetPlotLimitsV(ImAxis x_axis,ImAxis y_axis) { return ImPlot_GetPlotLimits(x_axis,y_axis); }
void Plot_GetPlotMousePosV(ImPlotPoint *pOut,ImAxis x_axis,ImAxis y_axis) { ImPlot_GetPlotMousePos(pOut,x_axis,y_axis); }
void Plot_GetPlotPos(ImVec2 *pOut) { ImPlot_GetPlotPos(pOut); }
ImPlotRect Plot_GetPlotSelectionV(ImAxis x_axis,ImAxis y_axis) { return ImPlot_GetPlotSelection(x_axis,y_axis); }
void Plot_GetPlotSize(ImVec2 *pOut) { ImPlot_GetPlotSize(pOut); }
ImPlotStyle* Plot_GetStyle() { return ImPlot_GetStyle(); }
const char* Plot_GetStyleColorName(ImPlotCol idx) { return ImPlot_GetStyleColorName(idx); }
void Plot_HideNextItemV(bool hidden,ImPlotCond cond) { ImPlot_HideNextItem(hidden,cond); }
bool Plot_IsAxisHovered(ImAxis axis) { return ImPlot_IsAxisHovered(axis); }
bool Plot_IsLegendEntryHovered(const char* label_id) { return ImPlot_IsLegendEntryHovered(label_id); }
bool Plot_IsPlotHovered() { return ImPlot_IsPlotHovered(); }
bool Plot_IsPlotSelected() { return ImPlot_IsPlotSelected(); }
bool Plot_IsSubplotsHovered() { return ImPlot_IsSubplotsHovered(); }
void Plot_ItemIcon_U32(ImU32 col) { ImPlot_ItemIcon_U32(col); }
void Plot_ItemIcon_Vec4(const ImVec4 col) { ImPlot_ItemIcon_Vec4(col); }
void Plot_MapInputDefaultV(ImPlotInputMap* dst) { ImPlot_MapInputDefault(dst); }
void Plot_MapInputReverseV(ImPlotInputMap* dst) { ImPlot_MapInputReverse(dst); }
void Plot_NextColormapColor(ImVec4 *pOut) { ImPlot_NextColormapColor(pOut); }
void Plot_PixelsToPlot_FloatV(ImPlotPoint *pOut,float x,float y,ImAxis x_axis,ImAxis y_axis) { ImPlot_PixelsToPlot_Float(pOut,x,y,x_axis,y_axis); }
void Plot_PixelsToPlot_Vec2V(ImPlotPoint *pOut,const ImVec2 pix,ImAxis x_axis,ImAxis y_axis) { ImPlot_PixelsToPlot_Vec2(pOut,pix,x_axis,y_axis); }
void Plot_PlotBarGroups_FloatPtrV(const char* const label_ids[],const float* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_FloatPtr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_S16PtrV(const char* const label_ids[],const ImS16* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_S16Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_S32PtrV(const char* const label_ids[],const ImS32* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_S32Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_S64PtrV(const char* const label_ids[],const ImS64* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_S64Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_S8PtrV(const char* const label_ids[],const ImS8* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_S8Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_U16PtrV(const char* const label_ids[],const ImU16* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_U16Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_U32PtrV(const char* const label_ids[],const ImU32* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_U32Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_U64PtrV(const char* const label_ids[],const ImU64* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_U64Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_U8PtrV(const char* const label_ids[],const ImU8* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_U8Ptr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarGroups_doublePtrV(const char* const label_ids[],const double* values,int item_count,int group_count,double group_size,double shift,ImPlotBarGroupsFlags flags) { ImPlot_PlotBarGroups_doublePtr(label_ids,values,item_count,group_count,group_size,shift,flags); }
void Plot_PlotBarsGV(const char* label_id,ImPlotPoint_getter getter,void* data,int count,double bar_size,ImPlotBarsFlags flags) { ImPlot_PlotBarsG(label_id,getter,data,count,bar_size,flags); }
void Plot_PlotBars_FloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_FloatPtrFloatPtr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_FloatPtrIntV(const char* label_id,const float* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_FloatPtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_S16PtrIntV(const char* label_id,const ImS16* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S16PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_S16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S16PtrS16Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_S32PtrIntV(const char* label_id,const ImS32* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S32PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_S32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S32PtrS32Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_S64PtrIntV(const char* label_id,const ImS64* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S64PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_S64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S64PtrS64Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_S8PtrIntV(const char* label_id,const ImS8* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S8PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_S8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_S8PtrS8Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_U16PtrIntV(const char* label_id,const ImU16* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U16PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_U16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U16PtrU16Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_U32PtrIntV(const char* label_id,const ImU32* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U32PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_U32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U32PtrU32Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_U64PtrIntV(const char* label_id,const ImU64* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U64PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_U64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U64PtrU64Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_U8PtrIntV(const char* label_id,const ImU8* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U8PtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_U8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_U8PtrU8Ptr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotBars_doublePtrIntV(const char* label_id,const double* values,int count,double bar_size,double shift,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_doublePtrInt(label_id,values,count,bar_size,shift,flags,offset,stride); }
void Plot_PlotBars_doublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys,int count,double bar_size,ImPlotBarsFlags flags,int offset,int stride) { ImPlot_PlotBars_doublePtrdoublePtr(label_id,xs,ys,count,bar_size,flags,offset,stride); }
void Plot_PlotDigitalGV(const char* label_id,ImPlotPoint_getter getter,void* data,int count,ImPlotDigitalFlags flags) { ImPlot_PlotDigitalG(label_id,getter,data,count,flags); }
void Plot_PlotDigital_FloatPtrV(const char* label_id,const float* xs,const float* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_FloatPtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_S16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_S16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_S32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_S32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_S64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_S64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_S8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_S8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_U16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_U16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_U32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_U32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_U64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_U64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_U8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_U8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDigital_doublePtrV(const char* label_id,const double* xs,const double* ys,int count,ImPlotDigitalFlags flags,int offset,int stride) { ImPlot_PlotDigital_doublePtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotDummyV(const char* label_id,ImPlotDummyFlags flags) { ImPlot_PlotDummy(label_id,flags); }
void Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys,const float* neg,const float* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV(const char* label_id,const float* xs,const float* ys,const float* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_S16PtrS16PtrS16PtrIntV(const char* label_id,const ImS16* xs,const ImS16* ys,const ImS16* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S16PtrS16PtrS16PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_S16PtrS16PtrS16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,const ImS16* neg,const ImS16* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S16PtrS16PtrS16PtrS16Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_S32PtrS32PtrS32PtrIntV(const char* label_id,const ImS32* xs,const ImS32* ys,const ImS32* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S32PtrS32PtrS32PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_S32PtrS32PtrS32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,const ImS32* neg,const ImS32* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S32PtrS32PtrS32PtrS32Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_S64PtrS64PtrS64PtrIntV(const char* label_id,const ImS64* xs,const ImS64* ys,const ImS64* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,const ImS64* neg,const ImS64* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_S8PtrS8PtrS8PtrIntV(const char* label_id,const ImS8* xs,const ImS8* ys,const ImS8* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S8PtrS8PtrS8PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_S8PtrS8PtrS8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,const ImS8* neg,const ImS8* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_S8PtrS8PtrS8PtrS8Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_U16PtrU16PtrU16PtrIntV(const char* label_id,const ImU16* xs,const ImU16* ys,const ImU16* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U16PtrU16PtrU16PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_U16PtrU16PtrU16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,const ImU16* neg,const ImU16* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U16PtrU16PtrU16PtrU16Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_U32PtrU32PtrU32PtrIntV(const char* label_id,const ImU32* xs,const ImU32* ys,const ImU32* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U32PtrU32PtrU32PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_U32PtrU32PtrU32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,const ImU32* neg,const ImU32* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U32PtrU32PtrU32PtrU32Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_U64PtrU64PtrU64PtrIntV(const char* label_id,const ImU64* xs,const ImU64* ys,const ImU64* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,const ImU64* neg,const ImU64* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_U8PtrU8PtrU8PtrIntV(const char* label_id,const ImU8* xs,const ImU8* ys,const ImU8* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U8PtrU8PtrU8PtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_U8PtrU8PtrU8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,const ImU8* neg,const ImU8* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_U8PtrU8PtrU8PtrU8Ptr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotErrorBars_doublePtrdoublePtrdoublePtrIntV(const char* label_id,const double* xs,const double* ys,const double* err,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_doublePtrdoublePtrdoublePtrInt(label_id,xs,ys,err,count,flags,offset,stride); }
void Plot_PlotErrorBars_doublePtrdoublePtrdoublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys,const double* neg,const double* pos,int count,ImPlotErrorBarsFlags flags,int offset,int stride) { ImPlot_PlotErrorBars_doublePtrdoublePtrdoublePtrdoublePtr(label_id,xs,ys,neg,pos,count,flags,offset,stride); }
void Plot_PlotHeatmap_FloatPtrV(const char* label_id,const float* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_FloatPtr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_S16PtrV(const char* label_id,const ImS16* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_S16Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_S32PtrV(const char* label_id,const ImS32* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_S32Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_S64PtrV(const char* label_id,const ImS64* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_S64Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_S8PtrV(const char* label_id,const ImS8* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_S8Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_U16PtrV(const char* label_id,const ImU16* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_U16Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_U32PtrV(const char* label_id,const ImU32* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_U32Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_U64PtrV(const char* label_id,const ImU64* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_U64Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_U8PtrV(const char* label_id,const ImU8* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_U8Ptr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
void Plot_PlotHeatmap_doublePtrV(const char* label_id,const double* values,int rows,int cols,double scale_min,double scale_max,const char* label_fmt,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,ImPlotHeatmapFlags flags) { ImPlot_PlotHeatmap_doublePtr(label_id,values,rows,cols,scale_min,scale_max,label_fmt,bounds_min,bounds_max,flags); }
double Plot_PlotHistogram2D_FloatPtrV(const char* label_id,const float* xs,const float* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_FloatPtr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_S16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_S16Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_S32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_S32Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_S64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_S64Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_S8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_S8Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_U16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_U16Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_U32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_U32Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_U64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_U64Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_U8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_U8Ptr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram2D_doublePtrV(const char* label_id,const double* xs,const double* ys,int count,int x_bins,int y_bins,ImPlotRect range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram2D_doublePtr(label_id,xs,ys,count,x_bins,y_bins,range,flags); }
double Plot_PlotHistogram_FloatPtrV(const char* label_id,const float* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_FloatPtr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_S16PtrV(const char* label_id,const ImS16* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_S16Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_S32PtrV(const char* label_id,const ImS32* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_S32Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_S64PtrV(const char* label_id,const ImS64* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_S64Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_S8PtrV(const char* label_id,const ImS8* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_S8Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_U16PtrV(const char* label_id,const ImU16* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_U16Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_U32PtrV(const char* label_id,const ImU32* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_U32Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_U64PtrV(const char* label_id,const ImU64* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_U64Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_U8PtrV(const char* label_id,const ImU8* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_U8Ptr(label_id,values,count,bins,bar_scale,range,flags); }
double Plot_PlotHistogram_doublePtrV(const char* label_id,const double* values,int count,int bins,double bar_scale,ImPlotRange range,ImPlotHistogramFlags flags) { return ImPlot_PlotHistogram_doublePtr(label_id,values,count,bins,bar_scale,range,flags); }
void Plot_PlotImageV(const char* label_id,ImTextureID user_texture_id,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max,const ImVec2 uv0,const ImVec2 uv1,const ImVec4 tint_col,ImPlotImageFlags flags) { ImPlot_PlotImage(label_id,user_texture_id,bounds_min,bounds_max,uv0,uv1,tint_col,flags); }
void Plot_PlotInfLines_FloatPtrV(const char* label_id,const float* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_FloatPtr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_S16PtrV(const char* label_id,const ImS16* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_S16Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_S32PtrV(const char* label_id,const ImS32* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_S32Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_S64PtrV(const char* label_id,const ImS64* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_S64Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_S8PtrV(const char* label_id,const ImS8* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_S8Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_U16PtrV(const char* label_id,const ImU16* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_U16Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_U32PtrV(const char* label_id,const ImU32* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_U32Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_U64PtrV(const char* label_id,const ImU64* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_U64Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_U8PtrV(const char* label_id,const ImU8* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_U8Ptr(label_id,values,count,flags,offset,stride); }
void Plot_PlotInfLines_doublePtrV(const char* label_id,const double* values,int count,ImPlotInfLinesFlags flags,int offset,int stride) { ImPlot_PlotInfLines_doublePtr(label_id,values,count,flags,offset,stride); }
void Plot_PlotLineGV(const char* label_id,ImPlotPoint_getter getter,void* data,int count,ImPlotLineFlags flags) { ImPlot_PlotLineG(label_id,getter,data,count,flags); }
void Plot_PlotLine_FloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_FloatPtrFloatPtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_FloatPtrIntV(const char* label_id,const float* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_FloatPtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_S16PtrIntV(const char* label_id,const ImS16* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S16PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_S16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S16PtrS16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_S32PtrIntV(const char* label_id,const ImS32* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S32PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_S32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S32PtrS32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_S64PtrIntV(const char* label_id,const ImS64* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S64PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_S64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S64PtrS64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_S8PtrIntV(const char* label_id,const ImS8* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S8PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_S8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_S8PtrS8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_U16PtrIntV(const char* label_id,const ImU16* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U16PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_U16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U16PtrU16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_U32PtrIntV(const char* label_id,const ImU32* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U32PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_U32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U32PtrU32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_U64PtrIntV(const char* label_id,const ImU64* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U64PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_U64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U64PtrU64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_U8PtrIntV(const char* label_id,const ImU8* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U8PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_U8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_U8PtrU8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotLine_doublePtrIntV(const char* label_id,const double* values,int count,double xscale,double xstart,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_doublePtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotLine_doublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys,int count,ImPlotLineFlags flags,int offset,int stride) { ImPlot_PlotLine_doublePtrdoublePtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotPieChart_FloatPtrV(const char* const label_ids[],const float* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_FloatPtr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_S16PtrV(const char* const label_ids[],const ImS16* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_S16Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_S32PtrV(const char* const label_ids[],const ImS32* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_S32Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_S64PtrV(const char* const label_ids[],const ImS64* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_S64Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_S8PtrV(const char* const label_ids[],const ImS8* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_S8Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_U16PtrV(const char* const label_ids[],const ImU16* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_U16Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_U32PtrV(const char* const label_ids[],const ImU32* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_U32Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_U64PtrV(const char* const label_ids[],const ImU64* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_U64Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_U8PtrV(const char* const label_ids[],const ImU8* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_U8Ptr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotPieChart_doublePtrV(const char* const label_ids[],const double* values,int count,double x,double y,double radius,const char* label_fmt,double angle0,ImPlotPieChartFlags flags) { ImPlot_PlotPieChart_doublePtr(label_ids,values,count,x,y,radius,label_fmt,angle0,flags); }
void Plot_PlotScatterGV(const char* label_id,ImPlotPoint_getter getter,void* data,int count,ImPlotScatterFlags flags) { ImPlot_PlotScatterG(label_id,getter,data,count,flags); }
void Plot_PlotScatter_FloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_FloatPtrFloatPtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_FloatPtrIntV(const char* label_id,const float* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_FloatPtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_S16PtrIntV(const char* label_id,const ImS16* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S16PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_S16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S16PtrS16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_S32PtrIntV(const char* label_id,const ImS32* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S32PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_S32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S32PtrS32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_S64PtrIntV(const char* label_id,const ImS64* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S64PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_S64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S64PtrS64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_S8PtrIntV(const char* label_id,const ImS8* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S8PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_S8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_S8PtrS8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_U16PtrIntV(const char* label_id,const ImU16* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U16PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_U16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U16PtrU16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_U32PtrIntV(const char* label_id,const ImU32* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U32PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_U32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U32PtrU32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_U64PtrIntV(const char* label_id,const ImU64* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U64PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_U64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U64PtrU64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_U8PtrIntV(const char* label_id,const ImU8* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U8PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_U8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_U8PtrU8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotScatter_doublePtrIntV(const char* label_id,const double* values,int count,double xscale,double xstart,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_doublePtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotScatter_doublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys,int count,ImPlotScatterFlags flags,int offset,int stride) { ImPlot_PlotScatter_doublePtrdoublePtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotShadedGV(const char* label_id,ImPlotPoint_getter getter1,void* data1,ImPlotPoint_getter getter2,void* data2,int count,ImPlotShadedFlags flags) { ImPlot_PlotShadedG(label_id,getter1,data1,getter2,data2,count,flags); }
void Plot_PlotShaded_FloatPtrFloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys1,const float* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_FloatPtrFloatPtrIntV(const char* label_id,const float* xs,const float* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_FloatPtrFloatPtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_FloatPtrIntV(const char* label_id,const float* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_FloatPtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_S16PtrIntV(const char* label_id,const ImS16* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S16PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_S16PtrS16PtrIntV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S16PtrS16PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_S16PtrS16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys1,const ImS16* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S16PtrS16PtrS16Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_S32PtrIntV(const char* label_id,const ImS32* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S32PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_S32PtrS32PtrIntV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S32PtrS32PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_S32PtrS32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys1,const ImS32* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S32PtrS32PtrS32Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_S64PtrIntV(const char* label_id,const ImS64* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S64PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_S64PtrS64PtrIntV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S64PtrS64PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_S64PtrS64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys1,const ImS64* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S64PtrS64PtrS64Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_S8PtrIntV(const char* label_id,const ImS8* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S8PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_S8PtrS8PtrIntV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S8PtrS8PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_S8PtrS8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys1,const ImS8* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_S8PtrS8PtrS8Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_U16PtrIntV(const char* label_id,const ImU16* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U16PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_U16PtrU16PtrIntV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U16PtrU16PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_U16PtrU16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys1,const ImU16* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U16PtrU16PtrU16Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_U32PtrIntV(const char* label_id,const ImU32* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U32PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_U32PtrU32PtrIntV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U32PtrU32PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_U32PtrU32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys1,const ImU32* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U32PtrU32PtrU32Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_U64PtrIntV(const char* label_id,const ImU64* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U64PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_U64PtrU64PtrIntV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U64PtrU64PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_U64PtrU64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys1,const ImU64* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U64PtrU64PtrU64Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_U8PtrIntV(const char* label_id,const ImU8* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U8PtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_U8PtrU8PtrIntV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U8PtrU8PtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_U8PtrU8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys1,const ImU8* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_U8PtrU8PtrU8Ptr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotShaded_doublePtrIntV(const char* label_id,const double* values,int count,double yref,double xscale,double xstart,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_doublePtrInt(label_id,values,count,yref,xscale,xstart,flags,offset,stride); }
void Plot_PlotShaded_doublePtrdoublePtrIntV(const char* label_id,const double* xs,const double* ys,int count,double yref,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_doublePtrdoublePtrInt(label_id,xs,ys,count,yref,flags,offset,stride); }
void Plot_PlotShaded_doublePtrdoublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys1,const double* ys2,int count,ImPlotShadedFlags flags,int offset,int stride) { ImPlot_PlotShaded_doublePtrdoublePtrdoublePtr(label_id,xs,ys1,ys2,count,flags,offset,stride); }
void Plot_PlotStairsGV(const char* label_id,ImPlotPoint_getter getter,void* data,int count,ImPlotStairsFlags flags) { ImPlot_PlotStairsG(label_id,getter,data,count,flags); }
void Plot_PlotStairs_FloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_FloatPtrFloatPtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_FloatPtrIntV(const char* label_id,const float* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_FloatPtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_S16PtrIntV(const char* label_id,const ImS16* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S16PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_S16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S16PtrS16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_S32PtrIntV(const char* label_id,const ImS32* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S32PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_S32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S32PtrS32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_S64PtrIntV(const char* label_id,const ImS64* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S64PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_S64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S64PtrS64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_S8PtrIntV(const char* label_id,const ImS8* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S8PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_S8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_S8PtrS8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_U16PtrIntV(const char* label_id,const ImU16* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U16PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_U16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U16PtrU16Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_U32PtrIntV(const char* label_id,const ImU32* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U32PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_U32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U32PtrU32Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_U64PtrIntV(const char* label_id,const ImU64* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U64PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_U64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U64PtrU64Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_U8PtrIntV(const char* label_id,const ImU8* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U8PtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_U8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_U8PtrU8Ptr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStairs_doublePtrIntV(const char* label_id,const double* values,int count,double xscale,double xstart,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_doublePtrInt(label_id,values,count,xscale,xstart,flags,offset,stride); }
void Plot_PlotStairs_doublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys,int count,ImPlotStairsFlags flags,int offset,int stride) { ImPlot_PlotStairs_doublePtrdoublePtr(label_id,xs,ys,count,flags,offset,stride); }
void Plot_PlotStems_FloatPtrFloatPtrV(const char* label_id,const float* xs,const float* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_FloatPtrFloatPtr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_FloatPtrIntV(const char* label_id,const float* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_FloatPtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_S16PtrIntV(const char* label_id,const ImS16* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S16PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_S16PtrS16PtrV(const char* label_id,const ImS16* xs,const ImS16* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S16PtrS16Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_S32PtrIntV(const char* label_id,const ImS32* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S32PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_S32PtrS32PtrV(const char* label_id,const ImS32* xs,const ImS32* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S32PtrS32Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_S64PtrIntV(const char* label_id,const ImS64* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S64PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_S64PtrS64PtrV(const char* label_id,const ImS64* xs,const ImS64* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S64PtrS64Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_S8PtrIntV(const char* label_id,const ImS8* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S8PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_S8PtrS8PtrV(const char* label_id,const ImS8* xs,const ImS8* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_S8PtrS8Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_U16PtrIntV(const char* label_id,const ImU16* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U16PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_U16PtrU16PtrV(const char* label_id,const ImU16* xs,const ImU16* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U16PtrU16Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_U32PtrIntV(const char* label_id,const ImU32* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U32PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_U32PtrU32PtrV(const char* label_id,const ImU32* xs,const ImU32* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U32PtrU32Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_U64PtrIntV(const char* label_id,const ImU64* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U64PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_U64PtrU64PtrV(const char* label_id,const ImU64* xs,const ImU64* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U64PtrU64Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_U8PtrIntV(const char* label_id,const ImU8* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U8PtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_U8PtrU8PtrV(const char* label_id,const ImU8* xs,const ImU8* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_U8PtrU8Ptr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotStems_doublePtrIntV(const char* label_id,const double* values,int count,double ref,double scale,double start,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_doublePtrInt(label_id,values,count,ref,scale,start,flags,offset,stride); }
void Plot_PlotStems_doublePtrdoublePtrV(const char* label_id,const double* xs,const double* ys,int count,double ref,ImPlotStemsFlags flags,int offset,int stride) { ImPlot_PlotStems_doublePtrdoublePtr(label_id,xs,ys,count,ref,flags,offset,stride); }
void Plot_PlotTextV(const char* text,double x,double y,const ImVec2 pix_offset,ImPlotTextFlags flags) { ImPlot_PlotText(text,x,y,pix_offset,flags); }
void Plot_PlotToPixels_PlotPoIntV(ImVec2 *pOut,const ImPlotPoint plt,ImAxis x_axis,ImAxis y_axis) { ImPlot_PlotToPixels_PlotPoInt(pOut,plt,x_axis,y_axis); }
void Plot_PlotToPixels_doubleV(ImVec2 *pOut,double x,double y,ImAxis x_axis,ImAxis y_axis) { ImPlot_PlotToPixels_double(pOut,x,y,x_axis,y_axis); }
void Plot_PopColormapV(int count) { ImPlot_PopColormap(count); }
void Plot_PopPlotClipRect() { ImPlot_PopPlotClipRect(); }
void Plot_PopStyleColorV(int count) { ImPlot_PopStyleColor(count); }
void Plot_PopStyleVarV(int count) { ImPlot_PopStyleVar(count); }
void Plot_PushColormap_PlotColormap(ImPlotColormap cmap) { ImPlot_PushColormap_PlotColormap(cmap); }
void Plot_PushColormap_Str(const char* name) { ImPlot_PushColormap_Str(name); }
void Plot_PushPlotClipRectV(float expand) { ImPlot_PushPlotClipRect(expand); }
void Plot_PushStyleColor_U32(ImPlotCol idx,ImU32 col) { ImPlot_PushStyleColor_U32(idx,col); }
void Plot_PushStyleColor_Vec4(ImPlotCol idx,const ImVec4 col) { ImPlot_PushStyleColor_Vec4(idx,col); }
void Plot_PushStyleVar_Float(ImPlotStyleVar idx,float val) { ImPlot_PushStyleVar_Float(idx,val); }
void Plot_PushStyleVar_Int(ImPlotStyleVar idx,int val) { ImPlot_PushStyleVar_Int(idx,val); }
void Plot_PushStyleVar_Vec2(ImPlotStyleVar idx,const ImVec2 val) { ImPlot_PushStyleVar_Vec2(idx,val); }
void Plot_SampleColormapV(ImVec4 *pOut,float t,ImPlotColormap cmap) { ImPlot_SampleColormap(pOut,t,cmap); }
void Plot_SetAxes(ImAxis x_axis,ImAxis y_axis) { ImPlot_SetAxes(x_axis,y_axis); }
void Plot_SetAxis(ImAxis axis) { ImPlot_SetAxis(axis); }
void Plot_SetCurrentContext(ImPlotContext* ctx) { ImPlot_SetCurrentContext(ctx); }
void Plot_SetImGuiContext(ImGuiContext* ctx) { ImPlot_SetImGuiContext(ctx); }
void Plot_SetNextAxesLimitsV(double x_min,double x_max,double y_min,double y_max,ImPlotCond cond) { ImPlot_SetNextAxesLimits(x_min,x_max,y_min,y_max,cond); }
void Plot_SetNextAxesToFit() { ImPlot_SetNextAxesToFit(); }
void Plot_SetNextAxisLimitsV(ImAxis axis,double v_min,double v_max,ImPlotCond cond) { ImPlot_SetNextAxisLimits(axis,v_min,v_max,cond); }
void Plot_SetNextAxisLinks(ImAxis axis,double* link_min,double* link_max) { ImPlot_SetNextAxisLinks(axis,link_min,link_max); }
void Plot_SetNextAxisToFit(ImAxis axis) { ImPlot_SetNextAxisToFit(axis); }
void Plot_SetNextErrorBarStyleV(const ImVec4 col,float size,float weight) { ImPlot_SetNextErrorBarStyle(col,size,weight); }
void Plot_SetNextFillStyleV(const ImVec4 col,float alpha_mod) { ImPlot_SetNextFillStyle(col,alpha_mod); }
void Plot_SetNextLineStyleV(const ImVec4 col,float weight) { ImPlot_SetNextLineStyle(col,weight); }
void Plot_SetNextMarkerStyleV(ImPlotMarker marker,float size,const ImVec4 fill,float weight,const ImVec4 outline) { ImPlot_SetNextMarkerStyle(marker,size,fill,weight,outline); }
void Plot_SetupAxesV(const char* x_label,const char* y_label,ImPlotAxisFlags x_flags,ImPlotAxisFlags y_flags) { ImPlot_SetupAxes(x_label,y_label,x_flags,y_flags); }
void Plot_SetupAxesLimitsV(double x_min,double x_max,double y_min,double y_max,ImPlotCond cond) { ImPlot_SetupAxesLimits(x_min,x_max,y_min,y_max,cond); }
void Plot_SetupAxisV(ImAxis axis,const char* label,ImPlotAxisFlags flags) { ImPlot_SetupAxis(axis,label,flags); }
void Plot_SetupAxisFormat_PlotFormatterV(ImAxis axis,ImPlotFormatter formatter,void* data) { ImPlot_SetupAxisFormat_PlotFormatter(axis,formatter,data); }
void Plot_SetupAxisFormat_Str(ImAxis axis,const char* fmt) { ImPlot_SetupAxisFormat_Str(axis,fmt); }
void Plot_SetupAxisLimitsV(ImAxis axis,double v_min,double v_max,ImPlotCond cond) { ImPlot_SetupAxisLimits(axis,v_min,v_max,cond); }
void Plot_SetupAxisLimitsConstraints(ImAxis axis,double v_min,double v_max) { ImPlot_SetupAxisLimitsConstraints(axis,v_min,v_max); }
void Plot_SetupAxisLinks(ImAxis axis,double* link_min,double* link_max) { ImPlot_SetupAxisLinks(axis,link_min,link_max); }
void Plot_SetupAxisScale_PlotScale(ImAxis axis,ImPlotScale scale) { ImPlot_SetupAxisScale_PlotScale(axis,scale); }
void Plot_SetupAxisScale_PlotTransformV(ImAxis axis,ImPlotTransform forward,ImPlotTransform inverse,void* data) { ImPlot_SetupAxisScale_PlotTransform(axis,forward,inverse,data); }
void Plot_SetupAxisTicks_doubleV(ImAxis axis,double v_min,double v_max,int n_ticks,const char* const labels[],bool keep_default) { ImPlot_SetupAxisTicks_double(axis,v_min,v_max,n_ticks,labels,keep_default); }
void Plot_SetupAxisTicks_doublePtrV(ImAxis axis,const double* values,int n_ticks,const char* const labels[],bool keep_default) { ImPlot_SetupAxisTicks_doublePtr(axis,values,n_ticks,labels,keep_default); }
void Plot_SetupAxisZoomConstraints(ImAxis axis,double z_min,double z_max) { ImPlot_SetupAxisZoomConstraints(axis,z_min,z_max); }
void Plot_SetupFinish() { ImPlot_SetupFinish(); }
void Plot_SetupLegendV(ImPlotLocation location,ImPlotLegendFlags flags) { ImPlot_SetupLegend(location,flags); }
void Plot_SetupMouseTextV(ImPlotLocation location,ImPlotMouseTextFlags flags) { ImPlot_SetupMouseText(location,flags); }
bool Plot_ShowColormapSelector(const char* label) { return ImPlot_ShowColormapSelector(label); }
void Plot_ShowDemoWindowV(bool* p_open) { ImPlot_ShowDemoWindow(p_open); }
bool Plot_ShowInputMapSelector(const char* label) { return ImPlot_ShowInputMapSelector(label); }
void Plot_ShowMetricsWindowV(bool* p_popen) { ImPlot_ShowMetricsWindow(p_popen); }
void Plot_ShowStyleEditorV(ImPlotStyle* ref) { ImPlot_ShowStyleEditor(ref); }
bool Plot_ShowStyleSelector(const char* label) { return ImPlot_ShowStyleSelector(label); }
void Plot_ShowUserGuide() { ImPlot_ShowUserGuide(); }
void Plot_StyleColorsAutoV(ImPlotStyle* dst) { ImPlot_StyleColorsAuto(dst); }
void Plot_StyleColorsClassicV(ImPlotStyle* dst) { ImPlot_StyleColorsClassic(dst); }
void Plot_StyleColorsDarkV(ImPlotStyle* dst) { ImPlot_StyleColorsDark(dst); }
void Plot_StyleColorsLightV(ImPlotStyle* dst) { ImPlot_StyleColorsLight(dst); }
void Plot_TagX_BoolV(double x,const ImVec4 col,bool round) { ImPlot_TagX_Bool(x,col,round); }
void Plot_TagX_Str(double x,const ImVec4 col,const char* fmt) { ImPlot_TagX_Str(x,col,fmt); }
void Plot_TagY_BoolV(double y,const ImVec4 col,bool round) { ImPlot_TagY_Bool(y,col,round); }
void Plot_TagY_Str(double y,const ImVec4 col,const char* fmt) { ImPlot_TagY_Str(y,col,fmt); }
ImPlotColormap Plot_AddColormap_U32Ptr(const char* name,const ImU32* cols,int size) { return Plot_AddColormap_U32PtrV(name,cols,size,true); }
ImPlotColormap Plot_AddColormap_Vec4Ptr(const char* name,const ImVec4* cols,int size) { return Plot_AddColormap_Vec4PtrV(name,cols,size,true); }
void Plot_Annotation_Bool(double x,double y,const ImVec4 col,const ImVec2 pix_offset,bool clamp) { Plot_Annotation_BoolV(x,y,col,pix_offset,clamp,false); }
bool Plot_BeginAlignedPlots(const char* group_id) { return Plot_BeginAlignedPlotsV(group_id,true); }
bool Plot_BeginDragDropSourceAxis(ImAxis axis) { return Plot_BeginDragDropSourceAxisV(axis,0); }
bool Plot_BeginDragDropSourceItem(const char* label_id) { return Plot_BeginDragDropSourceItemV(label_id,0); }
bool Plot_BeginDragDropSourcePlot() { return Plot_BeginDragDropSourcePlotV(0); }
bool Plot_BeginLegendPopup(const char* label_id) { return Plot_BeginLegendPopupV(label_id,1); }
bool Plot_BeginPlot(const char* title_id) { return Plot_BeginPlotV(title_id,(ImVec2){.x=-1, .y=0},0); }
bool Plot_BeginSubplots(const char* title_id,int rows,int cols,const ImVec2 size) { return Plot_BeginSubplotsV(title_id,rows,cols,size,0,NULL,NULL); }
void Plot_BustColorCache() { Plot_BustColorCacheV(NULL); }
bool Plot_ColormapButton(const char* label) { return Plot_ColormapButtonV(label,(ImVec2){.x=0, .y=0},-1); }
void Plot_ColormapScale(const char* label,double scale_min,double scale_max) { Plot_ColormapScaleV(label,scale_min,scale_max,(ImVec2){.x=0, .y=0},"%g",0,-1); }
bool Plot_ColormapSlider(const char* label,float* t) { return Plot_ColormapSliderV(label,t,NULL,"",-1); }
void Plot_DestroyContext() { Plot_DestroyContextV(NULL); }
bool Plot_DragLineX(int id,double* x,const ImVec4 col) { return Plot_DragLineXV(id,x,col,1,0); }
bool Plot_DragLineY(int id,double* y,const ImVec4 col) { return Plot_DragLineYV(id,y,col,1,0); }
bool Plot_DragPoint(int id,double* x,double* y,const ImVec4 col) { return Plot_DragPointV(id,x,y,col,4,0); }
bool Plot_DragRect(int id,double* x_min,double* y_min,double* x_max,double* y_max,const ImVec4 col) { return Plot_DragRectV(id,x_min,y_min,x_max,y_max,col,0); }
void Plot_GetColormapColor(ImVec4* pOut,int idx) { Plot_GetColormapColorV(pOut,idx,-1); }
int Plot_GetColormapSize() { return Plot_GetColormapSizeV(-1); }
ImPlotRect Plot_GetPlotLimits() { return Plot_GetPlotLimitsV(-1,-1); }
void Plot_GetPlotMousePos(ImPlotPoint* pOut) { Plot_GetPlotMousePosV(pOut,-1,-1); }
ImPlotRect Plot_GetPlotSelection() { return Plot_GetPlotSelectionV(-1,-1); }
void Plot_HideNextItem() { Plot_HideNextItemV(true,ImPlotCond_Once); }
void Plot_MapInputDefault() { Plot_MapInputDefaultV(NULL); }
void Plot_MapInputReverse() { Plot_MapInputReverseV(NULL); }
void Plot_PixelsToPlot_Float(ImPlotPoint* pOut,float x,float y) { Plot_PixelsToPlot_FloatV(pOut,x,y,-1,-1); }
void Plot_PixelsToPlot_Vec2(ImPlotPoint* pOut,const ImVec2 pix) { Plot_PixelsToPlot_Vec2V(pOut,pix,-1,-1); }
void Plot_PlotBarGroups_FloatPtr(const char* const label_ids[],const float* values,int item_count,int group_count) { Plot_PlotBarGroups_FloatPtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_S16Ptr(const char* const label_ids[],const ImS16* values,int item_count,int group_count) { Plot_PlotBarGroups_S16PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_S32Ptr(const char* const label_ids[],const ImS32* values,int item_count,int group_count) { Plot_PlotBarGroups_S32PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_S64Ptr(const char* const label_ids[],const ImS64* values,int item_count,int group_count) { Plot_PlotBarGroups_S64PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_S8Ptr(const char* const label_ids[],const ImS8* values,int item_count,int group_count) { Plot_PlotBarGroups_S8PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_U16Ptr(const char* const label_ids[],const ImU16* values,int item_count,int group_count) { Plot_PlotBarGroups_U16PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_U32Ptr(const char* const label_ids[],const ImU32* values,int item_count,int group_count) { Plot_PlotBarGroups_U32PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_U64Ptr(const char* const label_ids[],const ImU64* values,int item_count,int group_count) { Plot_PlotBarGroups_U64PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_U8Ptr(const char* const label_ids[],const ImU8* values,int item_count,int group_count) { Plot_PlotBarGroups_U8PtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarGroups_doublePtr(const char* const label_ids[],const double* values,int item_count,int group_count) { Plot_PlotBarGroups_doublePtrV(label_ids,values,item_count,group_count,0.67,0,0); }
void Plot_PlotBarsG(const char* label_id,ImPlotPoint_getter getter,void* data,int count,double bar_size) { Plot_PlotBarsGV(label_id,getter,data,count,bar_size,0); }
void Plot_PlotBars_FloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys,int count,double bar_size) { Plot_PlotBars_FloatPtrFloatPtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(float)); }
void Plot_PlotBars_FloatPtrInt(const char* label_id,const float* values,int count) { Plot_PlotBars_FloatPtrIntV(label_id,values,count,0.67,0,0,0,sizeof(float)); }
void Plot_PlotBars_S16PtrInt(const char* label_id,const ImS16* values,int count) { Plot_PlotBars_S16PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImS16)); }
void Plot_PlotBars_S16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count,double bar_size) { Plot_PlotBars_S16PtrS16PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImS16)); }
void Plot_PlotBars_S32PtrInt(const char* label_id,const ImS32* values,int count) { Plot_PlotBars_S32PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImS32)); }
void Plot_PlotBars_S32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count,double bar_size) { Plot_PlotBars_S32PtrS32PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImS32)); }
void Plot_PlotBars_S64PtrInt(const char* label_id,const ImS64* values,int count) { Plot_PlotBars_S64PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImS64)); }
void Plot_PlotBars_S64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count,double bar_size) { Plot_PlotBars_S64PtrS64PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImS64)); }
void Plot_PlotBars_S8PtrInt(const char* label_id,const ImS8* values,int count) { Plot_PlotBars_S8PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImS8)); }
void Plot_PlotBars_S8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count,double bar_size) { Plot_PlotBars_S8PtrS8PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImS8)); }
void Plot_PlotBars_U16PtrInt(const char* label_id,const ImU16* values,int count) { Plot_PlotBars_U16PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImU16)); }
void Plot_PlotBars_U16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count,double bar_size) { Plot_PlotBars_U16PtrU16PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImU16)); }
void Plot_PlotBars_U32PtrInt(const char* label_id,const ImU32* values,int count) { Plot_PlotBars_U32PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImU32)); }
void Plot_PlotBars_U32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count,double bar_size) { Plot_PlotBars_U32PtrU32PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImU32)); }
void Plot_PlotBars_U64PtrInt(const char* label_id,const ImU64* values,int count) { Plot_PlotBars_U64PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImU64)); }
void Plot_PlotBars_U64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count,double bar_size) { Plot_PlotBars_U64PtrU64PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImU64)); }
void Plot_PlotBars_U8PtrInt(const char* label_id,const ImU8* values,int count) { Plot_PlotBars_U8PtrIntV(label_id,values,count,0.67,0,0,0,sizeof(ImU8)); }
void Plot_PlotBars_U8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count,double bar_size) { Plot_PlotBars_U8PtrU8PtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(ImU8)); }
void Plot_PlotBars_doublePtrInt(const char* label_id,const double* values,int count) { Plot_PlotBars_doublePtrIntV(label_id,values,count,0.67,0,0,0,sizeof(double)); }
void Plot_PlotBars_doublePtrdoublePtr(const char* label_id,const double* xs,const double* ys,int count,double bar_size) { Plot_PlotBars_doublePtrdoublePtrV(label_id,xs,ys,count,bar_size,0,0,sizeof(double)); }
void Plot_PlotDigitalG(const char* label_id,ImPlotPoint_getter getter,void* data,int count) { Plot_PlotDigitalGV(label_id,getter,data,count,0); }
void Plot_PlotDigital_FloatPtr(const char* label_id,const float* xs,const float* ys,int count) { Plot_PlotDigital_FloatPtrV(label_id,xs,ys,count,0,0,sizeof(float)); }
void Plot_PlotDigital_S16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { Plot_PlotDigital_S16PtrV(label_id,xs,ys,count,0,0,sizeof(ImS16)); }
void Plot_PlotDigital_S32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { Plot_PlotDigital_S32PtrV(label_id,xs,ys,count,0,0,sizeof(ImS32)); }
void Plot_PlotDigital_S64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { Plot_PlotDigital_S64PtrV(label_id,xs,ys,count,0,0,sizeof(ImS64)); }
void Plot_PlotDigital_S8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { Plot_PlotDigital_S8PtrV(label_id,xs,ys,count,0,0,sizeof(ImS8)); }
void Plot_PlotDigital_U16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { Plot_PlotDigital_U16PtrV(label_id,xs,ys,count,0,0,sizeof(ImU16)); }
void Plot_PlotDigital_U32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { Plot_PlotDigital_U32PtrV(label_id,xs,ys,count,0,0,sizeof(ImU32)); }
void Plot_PlotDigital_U64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { Plot_PlotDigital_U64PtrV(label_id,xs,ys,count,0,0,sizeof(ImU64)); }
void Plot_PlotDigital_U8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { Plot_PlotDigital_U8PtrV(label_id,xs,ys,count,0,0,sizeof(ImU8)); }
void Plot_PlotDigital_doublePtr(const char* label_id,const double* xs,const double* ys,int count) { Plot_PlotDigital_doublePtrV(label_id,xs,ys,count,0,0,sizeof(double)); }
void Plot_PlotDummy(const char* label_id) { Plot_PlotDummyV(label_id,0); }
void Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys,const float* neg,const float* pos,int count) { Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(float)); }
void Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(const char* label_id,const float* xs,const float* ys,const float* err,int count) { Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV(label_id,xs,ys,err,count,0,0,sizeof(float)); }
void Plot_PlotErrorBars_S16PtrS16PtrS16PtrInt(const char* label_id,const ImS16* xs,const ImS16* ys,const ImS16* err,int count) { Plot_PlotErrorBars_S16PtrS16PtrS16PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImS16)); }
void Plot_PlotErrorBars_S16PtrS16PtrS16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,const ImS16* neg,const ImS16* pos,int count) { Plot_PlotErrorBars_S16PtrS16PtrS16PtrS16PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImS16)); }
void Plot_PlotErrorBars_S32PtrS32PtrS32PtrInt(const char* label_id,const ImS32* xs,const ImS32* ys,const ImS32* err,int count) { Plot_PlotErrorBars_S32PtrS32PtrS32PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImS32)); }
void Plot_PlotErrorBars_S32PtrS32PtrS32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,const ImS32* neg,const ImS32* pos,int count) { Plot_PlotErrorBars_S32PtrS32PtrS32PtrS32PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImS32)); }
void Plot_PlotErrorBars_S64PtrS64PtrS64PtrInt(const char* label_id,const ImS64* xs,const ImS64* ys,const ImS64* err,int count) { Plot_PlotErrorBars_S64PtrS64PtrS64PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImS64)); }
void Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,const ImS64* neg,const ImS64* pos,int count) { Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImS64)); }
void Plot_PlotErrorBars_S8PtrS8PtrS8PtrInt(const char* label_id,const ImS8* xs,const ImS8* ys,const ImS8* err,int count) { Plot_PlotErrorBars_S8PtrS8PtrS8PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImS8)); }
void Plot_PlotErrorBars_S8PtrS8PtrS8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,const ImS8* neg,const ImS8* pos,int count) { Plot_PlotErrorBars_S8PtrS8PtrS8PtrS8PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImS8)); }
void Plot_PlotErrorBars_U16PtrU16PtrU16PtrInt(const char* label_id,const ImU16* xs,const ImU16* ys,const ImU16* err,int count) { Plot_PlotErrorBars_U16PtrU16PtrU16PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImU16)); }
void Plot_PlotErrorBars_U16PtrU16PtrU16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,const ImU16* neg,const ImU16* pos,int count) { Plot_PlotErrorBars_U16PtrU16PtrU16PtrU16PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImU16)); }
void Plot_PlotErrorBars_U32PtrU32PtrU32PtrInt(const char* label_id,const ImU32* xs,const ImU32* ys,const ImU32* err,int count) { Plot_PlotErrorBars_U32PtrU32PtrU32PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImU32)); }
void Plot_PlotErrorBars_U32PtrU32PtrU32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,const ImU32* neg,const ImU32* pos,int count) { Plot_PlotErrorBars_U32PtrU32PtrU32PtrU32PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImU32)); }
void Plot_PlotErrorBars_U64PtrU64PtrU64PtrInt(const char* label_id,const ImU64* xs,const ImU64* ys,const ImU64* err,int count) { Plot_PlotErrorBars_U64PtrU64PtrU64PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImU64)); }
void Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,const ImU64* neg,const ImU64* pos,int count) { Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImU64)); }
void Plot_PlotErrorBars_U8PtrU8PtrU8PtrInt(const char* label_id,const ImU8* xs,const ImU8* ys,const ImU8* err,int count) { Plot_PlotErrorBars_U8PtrU8PtrU8PtrIntV(label_id,xs,ys,err,count,0,0,sizeof(ImU8)); }
void Plot_PlotErrorBars_U8PtrU8PtrU8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,const ImU8* neg,const ImU8* pos,int count) { Plot_PlotErrorBars_U8PtrU8PtrU8PtrU8PtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(ImU8)); }
void Plot_PlotErrorBars_doublePtrdoublePtrdoublePtrInt(const char* label_id,const double* xs,const double* ys,const double* err,int count) { Plot_PlotErrorBars_doublePtrdoublePtrdoublePtrIntV(label_id,xs,ys,err,count,0,0,sizeof(double)); }
void Plot_PlotErrorBars_doublePtrdoublePtrdoublePtrdoublePtr(const char* label_id,const double* xs,const double* ys,const double* neg,const double* pos,int count) { Plot_PlotErrorBars_doublePtrdoublePtrdoublePtrdoublePtrV(label_id,xs,ys,neg,pos,count,0,0,sizeof(double)); }
void Plot_PlotHeatmap_FloatPtr(const char* label_id,const float* values,int rows,int cols) { Plot_PlotHeatmap_FloatPtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_S16Ptr(const char* label_id,const ImS16* values,int rows,int cols) { Plot_PlotHeatmap_S16PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_S32Ptr(const char* label_id,const ImS32* values,int rows,int cols) { Plot_PlotHeatmap_S32PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_S64Ptr(const char* label_id,const ImS64* values,int rows,int cols) { Plot_PlotHeatmap_S64PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_S8Ptr(const char* label_id,const ImS8* values,int rows,int cols) { Plot_PlotHeatmap_S8PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_U16Ptr(const char* label_id,const ImU16* values,int rows,int cols) { Plot_PlotHeatmap_U16PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_U32Ptr(const char* label_id,const ImU32* values,int rows,int cols) { Plot_PlotHeatmap_U32PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_U64Ptr(const char* label_id,const ImU64* values,int rows,int cols) { Plot_PlotHeatmap_U64PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_U8Ptr(const char* label_id,const ImU8* values,int rows,int cols) { Plot_PlotHeatmap_U8PtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
void Plot_PlotHeatmap_doublePtr(const char* label_id,const double* values,int rows,int cols) { Plot_PlotHeatmap_doublePtrV(label_id,values,rows,cols,0,0,"%.1f",(ImPlotPoint){.x=0, .y=0},(ImPlotPoint){.x=1, .y=1},0); }
double Plot_PlotHistogram2D_FloatPtr(const char* label_id,const float* xs,const float* ys,int count) { return Plot_PlotHistogram2D_FloatPtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_S16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { return Plot_PlotHistogram2D_S16PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_S32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { return Plot_PlotHistogram2D_S32PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_S64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { return Plot_PlotHistogram2D_S64PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_S8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { return Plot_PlotHistogram2D_S8PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_U16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { return Plot_PlotHistogram2D_U16PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_U32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { return Plot_PlotHistogram2D_U32PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_U64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { return Plot_PlotHistogram2D_U64PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_U8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { return Plot_PlotHistogram2D_U8PtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram2D_doublePtr(const char* label_id,const double* xs,const double* ys,int count) { return Plot_PlotHistogram2D_doublePtrV(label_id,xs,ys,count,ImPlotBin_Sturges,ImPlotBin_Sturges,ImPlotRect(),0); }
double Plot_PlotHistogram_FloatPtr(const char* label_id,const float* values,int count) { return Plot_PlotHistogram_FloatPtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_S16Ptr(const char* label_id,const ImS16* values,int count) { return Plot_PlotHistogram_S16PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_S32Ptr(const char* label_id,const ImS32* values,int count) { return Plot_PlotHistogram_S32PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_S64Ptr(const char* label_id,const ImS64* values,int count) { return Plot_PlotHistogram_S64PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_S8Ptr(const char* label_id,const ImS8* values,int count) { return Plot_PlotHistogram_S8PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_U16Ptr(const char* label_id,const ImU16* values,int count) { return Plot_PlotHistogram_U16PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_U32Ptr(const char* label_id,const ImU32* values,int count) { return Plot_PlotHistogram_U32PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_U64Ptr(const char* label_id,const ImU64* values,int count) { return Plot_PlotHistogram_U64PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_U8Ptr(const char* label_id,const ImU8* values,int count) { return Plot_PlotHistogram_U8PtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
double Plot_PlotHistogram_doublePtr(const char* label_id,const double* values,int count) { return Plot_PlotHistogram_doublePtrV(label_id,values,count,ImPlotBin_Sturges,1.0,ImPlotRange(),0); }
void Plot_PlotImage(const char* label_id,ImTextureID user_texture_id,const ImPlotPoint bounds_min,const ImPlotPoint bounds_max) { Plot_PlotImageV(label_id,user_texture_id,bounds_min,bounds_max,(ImVec2){.x=0, .y=0},(ImVec2){.x=1, .y=1},(ImVec4){.x=1, .y=1, .z=1, .w=1},0); }
void Plot_PlotInfLines_FloatPtr(const char* label_id,const float* values,int count) { Plot_PlotInfLines_FloatPtrV(label_id,values,count,0,0,sizeof(float)); }
void Plot_PlotInfLines_S16Ptr(const char* label_id,const ImS16* values,int count) { Plot_PlotInfLines_S16PtrV(label_id,values,count,0,0,sizeof(ImS16)); }
void Plot_PlotInfLines_S32Ptr(const char* label_id,const ImS32* values,int count) { Plot_PlotInfLines_S32PtrV(label_id,values,count,0,0,sizeof(ImS32)); }
void Plot_PlotInfLines_S64Ptr(const char* label_id,const ImS64* values,int count) { Plot_PlotInfLines_S64PtrV(label_id,values,count,0,0,sizeof(ImS64)); }
void Plot_PlotInfLines_S8Ptr(const char* label_id,const ImS8* values,int count) { Plot_PlotInfLines_S8PtrV(label_id,values,count,0,0,sizeof(ImS8)); }
void Plot_PlotInfLines_U16Ptr(const char* label_id,const ImU16* values,int count) { Plot_PlotInfLines_U16PtrV(label_id,values,count,0,0,sizeof(ImU16)); }
void Plot_PlotInfLines_U32Ptr(const char* label_id,const ImU32* values,int count) { Plot_PlotInfLines_U32PtrV(label_id,values,count,0,0,sizeof(ImU32)); }
void Plot_PlotInfLines_U64Ptr(const char* label_id,const ImU64* values,int count) { Plot_PlotInfLines_U64PtrV(label_id,values,count,0,0,sizeof(ImU64)); }
void Plot_PlotInfLines_U8Ptr(const char* label_id,const ImU8* values,int count) { Plot_PlotInfLines_U8PtrV(label_id,values,count,0,0,sizeof(ImU8)); }
void Plot_PlotInfLines_doublePtr(const char* label_id,const double* values,int count) { Plot_PlotInfLines_doublePtrV(label_id,values,count,0,0,sizeof(double)); }
void Plot_PlotLineG(const char* label_id,ImPlotPoint_getter getter,void* data,int count) { Plot_PlotLineGV(label_id,getter,data,count,0); }
void Plot_PlotLine_FloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys,int count) { Plot_PlotLine_FloatPtrFloatPtrV(label_id,xs,ys,count,0,0,sizeof(float)); }
void Plot_PlotLine_FloatPtrInt(const char* label_id,const float* values,int count) { Plot_PlotLine_FloatPtrIntV(label_id,values,count,1,0,0,0,sizeof(float)); }
void Plot_PlotLine_S16PtrInt(const char* label_id,const ImS16* values,int count) { Plot_PlotLine_S16PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS16)); }
void Plot_PlotLine_S16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { Plot_PlotLine_S16PtrS16PtrV(label_id,xs,ys,count,0,0,sizeof(ImS16)); }
void Plot_PlotLine_S32PtrInt(const char* label_id,const ImS32* values,int count) { Plot_PlotLine_S32PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS32)); }
void Plot_PlotLine_S32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { Plot_PlotLine_S32PtrS32PtrV(label_id,xs,ys,count,0,0,sizeof(ImS32)); }
void Plot_PlotLine_S64PtrInt(const char* label_id,const ImS64* values,int count) { Plot_PlotLine_S64PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS64)); }
void Plot_PlotLine_S64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { Plot_PlotLine_S64PtrS64PtrV(label_id,xs,ys,count,0,0,sizeof(ImS64)); }
void Plot_PlotLine_S8PtrInt(const char* label_id,const ImS8* values,int count) { Plot_PlotLine_S8PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS8)); }
void Plot_PlotLine_S8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { Plot_PlotLine_S8PtrS8PtrV(label_id,xs,ys,count,0,0,sizeof(ImS8)); }
void Plot_PlotLine_U16PtrInt(const char* label_id,const ImU16* values,int count) { Plot_PlotLine_U16PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU16)); }
void Plot_PlotLine_U16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { Plot_PlotLine_U16PtrU16PtrV(label_id,xs,ys,count,0,0,sizeof(ImU16)); }
void Plot_PlotLine_U32PtrInt(const char* label_id,const ImU32* values,int count) { Plot_PlotLine_U32PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU32)); }
void Plot_PlotLine_U32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { Plot_PlotLine_U32PtrU32PtrV(label_id,xs,ys,count,0,0,sizeof(ImU32)); }
void Plot_PlotLine_U64PtrInt(const char* label_id,const ImU64* values,int count) { Plot_PlotLine_U64PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU64)); }
void Plot_PlotLine_U64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { Plot_PlotLine_U64PtrU64PtrV(label_id,xs,ys,count,0,0,sizeof(ImU64)); }
void Plot_PlotLine_U8PtrInt(const char* label_id,const ImU8* values,int count) { Plot_PlotLine_U8PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU8)); }
void Plot_PlotLine_U8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { Plot_PlotLine_U8PtrU8PtrV(label_id,xs,ys,count,0,0,sizeof(ImU8)); }
void Plot_PlotLine_doublePtrInt(const char* label_id,const double* values,int count) { Plot_PlotLine_doublePtrIntV(label_id,values,count,1,0,0,0,sizeof(double)); }
void Plot_PlotLine_doublePtrdoublePtr(const char* label_id,const double* xs,const double* ys,int count) { Plot_PlotLine_doublePtrdoublePtrV(label_id,xs,ys,count,0,0,sizeof(double)); }
void Plot_PlotPieChart_FloatPtr(const char* const label_ids[],const float* values,int count,double x,double y,double radius) { Plot_PlotPieChart_FloatPtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_S16Ptr(const char* const label_ids[],const ImS16* values,int count,double x,double y,double radius) { Plot_PlotPieChart_S16PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_S32Ptr(const char* const label_ids[],const ImS32* values,int count,double x,double y,double radius) { Plot_PlotPieChart_S32PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_S64Ptr(const char* const label_ids[],const ImS64* values,int count,double x,double y,double radius) { Plot_PlotPieChart_S64PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_S8Ptr(const char* const label_ids[],const ImS8* values,int count,double x,double y,double radius) { Plot_PlotPieChart_S8PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_U16Ptr(const char* const label_ids[],const ImU16* values,int count,double x,double y,double radius) { Plot_PlotPieChart_U16PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_U32Ptr(const char* const label_ids[],const ImU32* values,int count,double x,double y,double radius) { Plot_PlotPieChart_U32PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_U64Ptr(const char* const label_ids[],const ImU64* values,int count,double x,double y,double radius) { Plot_PlotPieChart_U64PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_U8Ptr(const char* const label_ids[],const ImU8* values,int count,double x,double y,double radius) { Plot_PlotPieChart_U8PtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotPieChart_doublePtr(const char* const label_ids[],const double* values,int count,double x,double y,double radius) { Plot_PlotPieChart_doublePtrV(label_ids,values,count,x,y,radius,"%.1f",90,0); }
void Plot_PlotScatterG(const char* label_id,ImPlotPoint_getter getter,void* data,int count) { Plot_PlotScatterGV(label_id,getter,data,count,0); }
void Plot_PlotScatter_FloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys,int count) { Plot_PlotScatter_FloatPtrFloatPtrV(label_id,xs,ys,count,0,0,sizeof(float)); }
void Plot_PlotScatter_FloatPtrInt(const char* label_id,const float* values,int count) { Plot_PlotScatter_FloatPtrIntV(label_id,values,count,1,0,0,0,sizeof(float)); }
void Plot_PlotScatter_S16PtrInt(const char* label_id,const ImS16* values,int count) { Plot_PlotScatter_S16PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS16)); }
void Plot_PlotScatter_S16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { Plot_PlotScatter_S16PtrS16PtrV(label_id,xs,ys,count,0,0,sizeof(ImS16)); }
void Plot_PlotScatter_S32PtrInt(const char* label_id,const ImS32* values,int count) { Plot_PlotScatter_S32PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS32)); }
void Plot_PlotScatter_S32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { Plot_PlotScatter_S32PtrS32PtrV(label_id,xs,ys,count,0,0,sizeof(ImS32)); }
void Plot_PlotScatter_S64PtrInt(const char* label_id,const ImS64* values,int count) { Plot_PlotScatter_S64PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS64)); }
void Plot_PlotScatter_S64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { Plot_PlotScatter_S64PtrS64PtrV(label_id,xs,ys,count,0,0,sizeof(ImS64)); }
void Plot_PlotScatter_S8PtrInt(const char* label_id,const ImS8* values,int count) { Plot_PlotScatter_S8PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS8)); }
void Plot_PlotScatter_S8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { Plot_PlotScatter_S8PtrS8PtrV(label_id,xs,ys,count,0,0,sizeof(ImS8)); }
void Plot_PlotScatter_U16PtrInt(const char* label_id,const ImU16* values,int count) { Plot_PlotScatter_U16PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU16)); }
void Plot_PlotScatter_U16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { Plot_PlotScatter_U16PtrU16PtrV(label_id,xs,ys,count,0,0,sizeof(ImU16)); }
void Plot_PlotScatter_U32PtrInt(const char* label_id,const ImU32* values,int count) { Plot_PlotScatter_U32PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU32)); }
void Plot_PlotScatter_U32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { Plot_PlotScatter_U32PtrU32PtrV(label_id,xs,ys,count,0,0,sizeof(ImU32)); }
void Plot_PlotScatter_U64PtrInt(const char* label_id,const ImU64* values,int count) { Plot_PlotScatter_U64PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU64)); }
void Plot_PlotScatter_U64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { Plot_PlotScatter_U64PtrU64PtrV(label_id,xs,ys,count,0,0,sizeof(ImU64)); }
void Plot_PlotScatter_U8PtrInt(const char* label_id,const ImU8* values,int count) { Plot_PlotScatter_U8PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU8)); }
void Plot_PlotScatter_U8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { Plot_PlotScatter_U8PtrU8PtrV(label_id,xs,ys,count,0,0,sizeof(ImU8)); }
void Plot_PlotScatter_doublePtrInt(const char* label_id,const double* values,int count) { Plot_PlotScatter_doublePtrIntV(label_id,values,count,1,0,0,0,sizeof(double)); }
void Plot_PlotScatter_doublePtrdoublePtr(const char* label_id,const double* xs,const double* ys,int count) { Plot_PlotScatter_doublePtrdoublePtrV(label_id,xs,ys,count,0,0,sizeof(double)); }
void Plot_PlotShadedG(const char* label_id,ImPlotPoint_getter getter1,void* data1,ImPlotPoint_getter getter2,void* data2,int count) { Plot_PlotShadedGV(label_id,getter1,data1,getter2,data2,count,0); }
void Plot_PlotShaded_FloatPtrFloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys1,const float* ys2,int count) { Plot_PlotShaded_FloatPtrFloatPtrFloatPtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(float)); }
void Plot_PlotShaded_FloatPtrFloatPtrInt(const char* label_id,const float* xs,const float* ys,int count) { Plot_PlotShaded_FloatPtrFloatPtrIntV(label_id,xs,ys,count,0,0,0,sizeof(float)); }
void Plot_PlotShaded_FloatPtrInt(const char* label_id,const float* values,int count) { Plot_PlotShaded_FloatPtrIntV(label_id,values,count,0,1,0,0,0,sizeof(float)); }
void Plot_PlotShaded_S16PtrInt(const char* label_id,const ImS16* values,int count) { Plot_PlotShaded_S16PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS16)); }
void Plot_PlotShaded_S16PtrS16PtrInt(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { Plot_PlotShaded_S16PtrS16PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImS16)); }
void Plot_PlotShaded_S16PtrS16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys1,const ImS16* ys2,int count) { Plot_PlotShaded_S16PtrS16PtrS16PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImS16)); }
void Plot_PlotShaded_S32PtrInt(const char* label_id,const ImS32* values,int count) { Plot_PlotShaded_S32PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS32)); }
void Plot_PlotShaded_S32PtrS32PtrInt(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { Plot_PlotShaded_S32PtrS32PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImS32)); }
void Plot_PlotShaded_S32PtrS32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys1,const ImS32* ys2,int count) { Plot_PlotShaded_S32PtrS32PtrS32PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImS32)); }
void Plot_PlotShaded_S64PtrInt(const char* label_id,const ImS64* values,int count) { Plot_PlotShaded_S64PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS64)); }
void Plot_PlotShaded_S64PtrS64PtrInt(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { Plot_PlotShaded_S64PtrS64PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImS64)); }
void Plot_PlotShaded_S64PtrS64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys1,const ImS64* ys2,int count) { Plot_PlotShaded_S64PtrS64PtrS64PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImS64)); }
void Plot_PlotShaded_S8PtrInt(const char* label_id,const ImS8* values,int count) { Plot_PlotShaded_S8PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS8)); }
void Plot_PlotShaded_S8PtrS8PtrInt(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { Plot_PlotShaded_S8PtrS8PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImS8)); }
void Plot_PlotShaded_S8PtrS8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys1,const ImS8* ys2,int count) { Plot_PlotShaded_S8PtrS8PtrS8PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImS8)); }
void Plot_PlotShaded_U16PtrInt(const char* label_id,const ImU16* values,int count) { Plot_PlotShaded_U16PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU16)); }
void Plot_PlotShaded_U16PtrU16PtrInt(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { Plot_PlotShaded_U16PtrU16PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImU16)); }
void Plot_PlotShaded_U16PtrU16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys1,const ImU16* ys2,int count) { Plot_PlotShaded_U16PtrU16PtrU16PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImU16)); }
void Plot_PlotShaded_U32PtrInt(const char* label_id,const ImU32* values,int count) { Plot_PlotShaded_U32PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU32)); }
void Plot_PlotShaded_U32PtrU32PtrInt(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { Plot_PlotShaded_U32PtrU32PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImU32)); }
void Plot_PlotShaded_U32PtrU32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys1,const ImU32* ys2,int count) { Plot_PlotShaded_U32PtrU32PtrU32PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImU32)); }
void Plot_PlotShaded_U64PtrInt(const char* label_id,const ImU64* values,int count) { Plot_PlotShaded_U64PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU64)); }
void Plot_PlotShaded_U64PtrU64PtrInt(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { Plot_PlotShaded_U64PtrU64PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImU64)); }
void Plot_PlotShaded_U64PtrU64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys1,const ImU64* ys2,int count) { Plot_PlotShaded_U64PtrU64PtrU64PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImU64)); }
void Plot_PlotShaded_U8PtrInt(const char* label_id,const ImU8* values,int count) { Plot_PlotShaded_U8PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU8)); }
void Plot_PlotShaded_U8PtrU8PtrInt(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { Plot_PlotShaded_U8PtrU8PtrIntV(label_id,xs,ys,count,0,0,0,sizeof(ImU8)); }
void Plot_PlotShaded_U8PtrU8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys1,const ImU8* ys2,int count) { Plot_PlotShaded_U8PtrU8PtrU8PtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(ImU8)); }
void Plot_PlotShaded_doublePtrInt(const char* label_id,const double* values,int count) { Plot_PlotShaded_doublePtrIntV(label_id,values,count,0,1,0,0,0,sizeof(double)); }
void Plot_PlotShaded_doublePtrdoublePtrInt(const char* label_id,const double* xs,const double* ys,int count) { Plot_PlotShaded_doublePtrdoublePtrIntV(label_id,xs,ys,count,0,0,0,sizeof(double)); }
void Plot_PlotShaded_doublePtrdoublePtrdoublePtr(const char* label_id,const double* xs,const double* ys1,const double* ys2,int count) { Plot_PlotShaded_doublePtrdoublePtrdoublePtrV(label_id,xs,ys1,ys2,count,0,0,sizeof(double)); }
void Plot_PlotStairsG(const char* label_id,ImPlotPoint_getter getter,void* data,int count) { Plot_PlotStairsGV(label_id,getter,data,count,0); }
void Plot_PlotStairs_FloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys,int count) { Plot_PlotStairs_FloatPtrFloatPtrV(label_id,xs,ys,count,0,0,sizeof(float)); }
void Plot_PlotStairs_FloatPtrInt(const char* label_id,const float* values,int count) { Plot_PlotStairs_FloatPtrIntV(label_id,values,count,1,0,0,0,sizeof(float)); }
void Plot_PlotStairs_S16PtrInt(const char* label_id,const ImS16* values,int count) { Plot_PlotStairs_S16PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS16)); }
void Plot_PlotStairs_S16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { Plot_PlotStairs_S16PtrS16PtrV(label_id,xs,ys,count,0,0,sizeof(ImS16)); }
void Plot_PlotStairs_S32PtrInt(const char* label_id,const ImS32* values,int count) { Plot_PlotStairs_S32PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS32)); }
void Plot_PlotStairs_S32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { Plot_PlotStairs_S32PtrS32PtrV(label_id,xs,ys,count,0,0,sizeof(ImS32)); }
void Plot_PlotStairs_S64PtrInt(const char* label_id,const ImS64* values,int count) { Plot_PlotStairs_S64PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS64)); }
void Plot_PlotStairs_S64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { Plot_PlotStairs_S64PtrS64PtrV(label_id,xs,ys,count,0,0,sizeof(ImS64)); }
void Plot_PlotStairs_S8PtrInt(const char* label_id,const ImS8* values,int count) { Plot_PlotStairs_S8PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImS8)); }
void Plot_PlotStairs_S8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { Plot_PlotStairs_S8PtrS8PtrV(label_id,xs,ys,count,0,0,sizeof(ImS8)); }
void Plot_PlotStairs_U16PtrInt(const char* label_id,const ImU16* values,int count) { Plot_PlotStairs_U16PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU16)); }
void Plot_PlotStairs_U16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { Plot_PlotStairs_U16PtrU16PtrV(label_id,xs,ys,count,0,0,sizeof(ImU16)); }
void Plot_PlotStairs_U32PtrInt(const char* label_id,const ImU32* values,int count) { Plot_PlotStairs_U32PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU32)); }
void Plot_PlotStairs_U32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { Plot_PlotStairs_U32PtrU32PtrV(label_id,xs,ys,count,0,0,sizeof(ImU32)); }
void Plot_PlotStairs_U64PtrInt(const char* label_id,const ImU64* values,int count) { Plot_PlotStairs_U64PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU64)); }
void Plot_PlotStairs_U64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { Plot_PlotStairs_U64PtrU64PtrV(label_id,xs,ys,count,0,0,sizeof(ImU64)); }
void Plot_PlotStairs_U8PtrInt(const char* label_id,const ImU8* values,int count) { Plot_PlotStairs_U8PtrIntV(label_id,values,count,1,0,0,0,sizeof(ImU8)); }
void Plot_PlotStairs_U8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { Plot_PlotStairs_U8PtrU8PtrV(label_id,xs,ys,count,0,0,sizeof(ImU8)); }
void Plot_PlotStairs_doublePtrInt(const char* label_id,const double* values,int count) { Plot_PlotStairs_doublePtrIntV(label_id,values,count,1,0,0,0,sizeof(double)); }
void Plot_PlotStairs_doublePtrdoublePtr(const char* label_id,const double* xs,const double* ys,int count) { Plot_PlotStairs_doublePtrdoublePtrV(label_id,xs,ys,count,0,0,sizeof(double)); }
void Plot_PlotStems_FloatPtrFloatPtr(const char* label_id,const float* xs,const float* ys,int count) { Plot_PlotStems_FloatPtrFloatPtrV(label_id,xs,ys,count,0,0,0,sizeof(float)); }
void Plot_PlotStems_FloatPtrInt(const char* label_id,const float* values,int count) { Plot_PlotStems_FloatPtrIntV(label_id,values,count,0,1,0,0,0,sizeof(float)); }
void Plot_PlotStems_S16PtrInt(const char* label_id,const ImS16* values,int count) { Plot_PlotStems_S16PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS16)); }
void Plot_PlotStems_S16PtrS16Ptr(const char* label_id,const ImS16* xs,const ImS16* ys,int count) { Plot_PlotStems_S16PtrS16PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImS16)); }
void Plot_PlotStems_S32PtrInt(const char* label_id,const ImS32* values,int count) { Plot_PlotStems_S32PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS32)); }
void Plot_PlotStems_S32PtrS32Ptr(const char* label_id,const ImS32* xs,const ImS32* ys,int count) { Plot_PlotStems_S32PtrS32PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImS32)); }
void Plot_PlotStems_S64PtrInt(const char* label_id,const ImS64* values,int count) { Plot_PlotStems_S64PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS64)); }
void Plot_PlotStems_S64PtrS64Ptr(const char* label_id,const ImS64* xs,const ImS64* ys,int count) { Plot_PlotStems_S64PtrS64PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImS64)); }
void Plot_PlotStems_S8PtrInt(const char* label_id,const ImS8* values,int count) { Plot_PlotStems_S8PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImS8)); }
void Plot_PlotStems_S8PtrS8Ptr(const char* label_id,const ImS8* xs,const ImS8* ys,int count) { Plot_PlotStems_S8PtrS8PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImS8)); }
void Plot_PlotStems_U16PtrInt(const char* label_id,const ImU16* values,int count) { Plot_PlotStems_U16PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU16)); }
void Plot_PlotStems_U16PtrU16Ptr(const char* label_id,const ImU16* xs,const ImU16* ys,int count) { Plot_PlotStems_U16PtrU16PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImU16)); }
void Plot_PlotStems_U32PtrInt(const char* label_id,const ImU32* values,int count) { Plot_PlotStems_U32PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU32)); }
void Plot_PlotStems_U32PtrU32Ptr(const char* label_id,const ImU32* xs,const ImU32* ys,int count) { Plot_PlotStems_U32PtrU32PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImU32)); }
void Plot_PlotStems_U64PtrInt(const char* label_id,const ImU64* values,int count) { Plot_PlotStems_U64PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU64)); }
void Plot_PlotStems_U64PtrU64Ptr(const char* label_id,const ImU64* xs,const ImU64* ys,int count) { Plot_PlotStems_U64PtrU64PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImU64)); }
void Plot_PlotStems_U8PtrInt(const char* label_id,const ImU8* values,int count) { Plot_PlotStems_U8PtrIntV(label_id,values,count,0,1,0,0,0,sizeof(ImU8)); }
void Plot_PlotStems_U8PtrU8Ptr(const char* label_id,const ImU8* xs,const ImU8* ys,int count) { Plot_PlotStems_U8PtrU8PtrV(label_id,xs,ys,count,0,0,0,sizeof(ImU8)); }
void Plot_PlotStems_doublePtrInt(const char* label_id,const double* values,int count) { Plot_PlotStems_doublePtrIntV(label_id,values,count,0,1,0,0,0,sizeof(double)); }
void Plot_PlotStems_doublePtrdoublePtr(const char* label_id,const double* xs,const double* ys,int count) { Plot_PlotStems_doublePtrdoublePtrV(label_id,xs,ys,count,0,0,0,sizeof(double)); }
void Plot_PlotText(const char* text,double x,double y) { Plot_PlotTextV(text,x,y,(ImVec2){.x=0, .y=0},0); }
void Plot_PlotToPixels_PlotPoInt(ImVec2* pOut,const ImPlotPoint plt) { Plot_PlotToPixels_PlotPoIntV(pOut,plt,-1,-1); }
void Plot_PlotToPixels_double(ImVec2* pOut,double x,double y) { Plot_PlotToPixels_doubleV(pOut,x,y,-1,-1); }
void Plot_PopColormap() { Plot_PopColormapV(1); }
void Plot_PopStyleColor() { Plot_PopStyleColorV(1); }
void Plot_PopStyleVar() { Plot_PopStyleVarV(1); }
void Plot_PushPlotClipRect() { Plot_PushPlotClipRectV(0); }
void Plot_SampleColormap(ImVec4* pOut,float t) { Plot_SampleColormapV(pOut,t,-1); }
void Plot_SetNextAxesLimits(double x_min,double x_max,double y_min,double y_max) { Plot_SetNextAxesLimitsV(x_min,x_max,y_min,y_max,ImPlotCond_Once); }
void Plot_SetNextAxisLimits(ImAxis axis,double v_min,double v_max) { Plot_SetNextAxisLimitsV(axis,v_min,v_max,ImPlotCond_Once); }
void Plot_SetNextErrorBarStyle() { Plot_SetNextErrorBarStyleV((ImVec4){.x=0, .y=0, .z=0, .w=-1},-1,-1); }
void Plot_SetNextFillStyle() { Plot_SetNextFillStyleV((ImVec4){.x=0, .y=0, .z=0, .w=-1},-1); }
void Plot_SetNextLineStyle() { Plot_SetNextLineStyleV((ImVec4){.x=0, .y=0, .z=0, .w=-1},-1); }
void Plot_SetNextMarkerStyle() { Plot_SetNextMarkerStyleV(-1,-1,(ImVec4){.x=0, .y=0, .z=0, .w=-1},-1,(ImVec4){.x=0, .y=0, .z=0, .w=-1}); }
void Plot_SetupAxes(const char* x_label,const char* y_label) { Plot_SetupAxesV(x_label,y_label,0,0); }
void Plot_SetupAxesLimits(double x_min,double x_max,double y_min,double y_max) { Plot_SetupAxesLimitsV(x_min,x_max,y_min,y_max,ImPlotCond_Once); }
void Plot_SetupAxis(ImAxis axis) { Plot_SetupAxisV(axis,NULL,0); }
void Plot_SetupAxisFormat_PlotFormatter(ImAxis axis,ImPlotFormatter formatter) { Plot_SetupAxisFormat_PlotFormatterV(axis,formatter,NULL); }
void Plot_SetupAxisLimits(ImAxis axis,double v_min,double v_max) { Plot_SetupAxisLimitsV(axis,v_min,v_max,ImPlotCond_Once); }
void Plot_SetupAxisScale_PlotTransform(ImAxis axis,ImPlotTransform forward,ImPlotTransform inverse) { Plot_SetupAxisScale_PlotTransformV(axis,forward,inverse,NULL); }
void Plot_SetupAxisTicks_double(ImAxis axis,double v_min,double v_max,int n_ticks) { Plot_SetupAxisTicks_doubleV(axis,v_min,v_max,n_ticks,NULL,false); }
void Plot_SetupAxisTicks_doublePtr(ImAxis axis,const double* values,int n_ticks) { Plot_SetupAxisTicks_doublePtrV(axis,values,n_ticks,NULL,false); }
void Plot_SetupLegend(ImPlotLocation location) { Plot_SetupLegendV(location,0); }
void Plot_SetupMouseText(ImPlotLocation location) { Plot_SetupMouseTextV(location,0); }
void Plot_ShowDemoWindow() { Plot_ShowDemoWindowV(NULL); }
void Plot_ShowMetricsWindow() { Plot_ShowMetricsWindowV(NULL); }
void Plot_ShowStyleEditor() { Plot_ShowStyleEditorV(NULL); }
void Plot_StyleColorsAuto() { Plot_StyleColorsAutoV(NULL); }
void Plot_StyleColorsClassic() { Plot_StyleColorsClassicV(NULL); }
void Plot_StyleColorsDark() { Plot_StyleColorsDarkV(NULL); }
void Plot_StyleColorsLight() { Plot_StyleColorsLightV(NULL); }
void Plot_TagX_Bool(double x,const ImVec4 col) { Plot_TagX_BoolV(x,col,false); }
void Plot_TagY_Bool(double y,const ImVec4 col) { Plot_TagY_BoolV(y,col,false); }
