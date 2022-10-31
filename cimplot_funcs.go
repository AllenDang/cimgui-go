package cimgui

// #include "extra_types.h"
// #include "cimplot_structs_accessor.h"
// #include "cimplot_wrapper.h"
import "C"
import "unsafe"

func (self ImPlotAlignmentData) Destroy() {
	C.ImPlotAlignmentData_Destroy(self.handle())
}

func (self ImPlotAnnotationCollection) Destroy() {
	C.ImPlotAnnotationCollection_Destroy(self.handle())
}

func (self ImPlotAxis) Destroy() {
	C.ImPlotAxis_Destroy(self.handle())
}

func (self ImPlotColormapData) Destroy() {
	C.ImPlotColormapData_Destroy(self.handle())
}

func (self ImPlotDateTimeSpec) Destroy() {
	C.ImPlotDateTimeSpec_Destroy(self.handle())
}

func (self ImPlotInputMap) Destroy() {
	C.ImPlotInputMap_Destroy(self.handle())
}

func (self ImPlotItemGroup) Destroy() {
	C.ImPlotItemGroup_Destroy(self.handle())
}

func (self ImPlotLegend) Destroy() {
	C.ImPlotLegend_Destroy(self.handle())
}

func (self ImPlotNextItemData) Destroy() {
	C.ImPlotNextItemData_Destroy(self.handle())
}

func (self ImPlotNextPlotData) Destroy() {
	C.ImPlotNextPlotData_Destroy(self.handle())
}

func (self ImPlotPlot) Destroy() {
	C.ImPlotPlot_Destroy(self.handle())
}

func (self ImPlotPointError) Destroy() {
	C.ImPlotPointError_Destroy(self.handle())
}

func (self *ImPlotPoint) Destroy() {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.ImPlotPoint_Destroy(selfArg)
}

func (self ImPlotRange) Clamp(value float64) float64 {
	return float64(C.ImPlotRange_Clamp(self.handle(), C.double(value)))
}

func (self ImPlotRange) Contains(value float64) bool {
	return C.ImPlotRange_Contains(self.handle(), C.double(value)) == C.bool(true)
}

func (self ImPlotRange) Size() float64 {
	return float64(C.ImPlotRange_Size(self.handle()))
}

func (self ImPlotRange) Destroy() {
	C.ImPlotRange_Destroy(self.handle())
}

func ImPlotRect_Clamp_double(pOut *ImPlotPoint, self ImPlotRect, x float64, y float64) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlotRect_Clamp_double(pOutArg, self.handle(), C.double(x), C.double(y))
}

func (self ImPlotRect) Contains_double(x float64, y float64) bool {
	return C.ImPlotRect_Contains_double(self.handle(), C.double(x), C.double(y)) == C.bool(true)
}

func ImPlotRect_Max(pOut *ImPlotPoint, self ImPlotRect) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlotRect_Max(pOutArg, self.handle())
}

func ImPlotRect_Min(pOut *ImPlotPoint, self ImPlotRect) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlotRect_Min(pOutArg, self.handle())
}

func ImPlotRect_Size(pOut *ImPlotPoint, self ImPlotRect) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlotRect_Size(pOutArg, self.handle())
}

func (self ImPlotRect) Destroy() {
	C.ImPlotRect_Destroy(self.handle())
}

func (self ImPlotStyle) Destroy() {
	C.ImPlotStyle_Destroy(self.handle())
}

func (self ImPlotSubplot) Destroy() {
	C.ImPlotSubplot_Destroy(self.handle())
}

func (self ImPlotTagCollection) Destroy() {
	C.ImPlotTagCollection_Destroy(self.handle())
}

func (self ImPlotTick) Destroy() {
	C.ImPlotTick_Destroy(self.handle())
}

func (self ImPlotTicker) Destroy() {
	C.ImPlotTicker_Destroy(self.handle())
}

func (self ImPlotTime) Destroy() {
	C.ImPlotTime_Destroy(self.handle())
}

// ImPlot_AddColormap_Vec4PtrV parameter default value hint:
// qual: true
func ImPlot_AddColormap_Vec4PtrV(name string, cols *ImVec4, size int32, qual bool) ImPlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	colsArg, colsFin := cols.wrap()
	defer colsFin()

	return ImPlotColormap(C.ImPlot_AddColormap_Vec4PtrV(nameArg, colsArg, C.int(size), C.bool(qual)))
}

// ImPlot_Annotation_BoolV parameter default value hint:
// round: false
func ImPlot_Annotation_BoolV(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool, round bool) {
	C.ImPlot_Annotation_BoolV(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp), C.bool(round))
}

func ImPlot_Annotation_Str(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.ImPlot_Annotation_Str(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp), fmtArg)
}

// ImPlot_BeginAlignedPlotsV parameter default value hint:
// vertical: true
func ImPlot_BeginAlignedPlotsV(group_id string, vertical bool) bool {
	group_idArg, group_idFin := wrapString(group_id)
	defer group_idFin()

	return C.ImPlot_BeginAlignedPlotsV(group_idArg, C.bool(vertical)) == C.bool(true)
}

// ImPlot_BeginDragDropSourceAxisV parameter default value hint:
// flags: 0
func ImPlot_BeginDragDropSourceAxisV(axis ImAxis, flags DragDropFlags) bool {
	return C.ImPlot_BeginDragDropSourceAxisV(C.ImAxis(axis), C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

// ImPlot_BeginDragDropSourceItemV parameter default value hint:
// flags: 0
func ImPlot_BeginDragDropSourceItemV(label_id string, flags DragDropFlags) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.ImPlot_BeginDragDropSourceItemV(label_idArg, C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

// ImPlot_BeginDragDropSourcePlotV parameter default value hint:
// flags: 0
func ImPlot_BeginDragDropSourcePlotV(flags DragDropFlags) bool {
	return C.ImPlot_BeginDragDropSourcePlotV(C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

func ImPlot_BeginDragDropTargetAxis(axis ImAxis) bool {
	return C.ImPlot_BeginDragDropTargetAxis(C.ImAxis(axis)) == C.bool(true)
}

func ImPlot_BeginDragDropTargetLegend() bool {
	return C.ImPlot_BeginDragDropTargetLegend() == C.bool(true)
}

func ImPlot_BeginDragDropTargetPlot() bool {
	return C.ImPlot_BeginDragDropTargetPlot() == C.bool(true)
}

// ImPlot_BeginLegendPopupV parameter default value hint:
// mouse_button: 1
func ImPlot_BeginLegendPopupV(label_id string, mouse_button MouseButton) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.ImPlot_BeginLegendPopupV(label_idArg, C.ImGuiMouseButton(mouse_button)) == C.bool(true)
}

// ImPlot_BeginPlotV parameter default value hint:
// flags: 0
// size: ImVec2(-1,0)
func ImPlot_BeginPlotV(title_id string, size ImVec2, flags ImPlotFlags) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.ImPlot_BeginPlotV(title_idArg, size.toC(), C.ImPlotFlags(flags)) == C.bool(true)
}

// ImPlot_BeginSubplotsV parameter default value hint:
// col_ratios: ((void*)0)
// flags: 0
// row_ratios: ((void*)0)
func ImPlot_BeginSubplotsV(title_id string, rows int32, cols int32, size ImVec2, flags ImPlotSubplotFlags, row_ratios *float32, col_ratios *float32) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	row_ratiosArg, row_ratiosFin := wrapFloat(row_ratios)
	defer row_ratiosFin()

	col_ratiosArg, col_ratiosFin := wrapFloat(col_ratios)
	defer col_ratiosFin()

	return C.ImPlot_BeginSubplotsV(title_idArg, C.int(rows), C.int(cols), size.toC(), C.ImPlotSubplotFlags(flags), row_ratiosArg, col_ratiosArg) == C.bool(true)
}

// ImPlot_BustColorCacheV parameter default value hint:
// plot_title_id: ((void*)0)
func ImPlot_BustColorCacheV(plot_title_id string) {
	plot_title_idArg, plot_title_idFin := wrapString(plot_title_id)
	defer plot_title_idFin()

	C.ImPlot_BustColorCacheV(plot_title_idArg)
}

func ImPlot_CancelPlotSelection() {
	C.ImPlot_CancelPlotSelection()
}

// ImPlot_ColormapButtonV parameter default value hint:
// cmap: -1
// size: ImVec2(0,0)
func ImPlot_ColormapButtonV(label string, size ImVec2, cmap ImPlotColormap) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ImPlot_ColormapButtonV(labelArg, size.toC(), C.ImPlotColormap(cmap)) == C.bool(true)
}

func ImPlot_ColormapIcon(cmap ImPlotColormap) {
	C.ImPlot_ColormapIcon(C.ImPlotColormap(cmap))
}

// ImPlot_ColormapScaleV parameter default value hint:
// cmap: -1
// flags: 0
// format: "%g"
// size: ImVec2(0,0)
func ImPlot_ColormapScaleV(label string, scale_min float64, scale_max float64, size ImVec2, format string, flags ImPlotColormapScaleFlags, cmap ImPlotColormap) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	C.ImPlot_ColormapScaleV(labelArg, C.double(scale_min), C.double(scale_max), size.toC(), formatArg, C.ImPlotColormapScaleFlags(flags), C.ImPlotColormap(cmap))
}

// ImPlot_ColormapSliderV parameter default value hint:
// cmap: -1
// format: ""
// out: ((void*)0)
func ImPlot_ColormapSliderV(label string, t *float32, out *ImVec4, format string, cmap ImPlotColormap) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	tArg, tFin := wrapFloat(t)
	defer tFin()

	outArg, outFin := out.wrap()
	defer outFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.ImPlot_ColormapSliderV(labelArg, tArg, outArg, formatArg, C.ImPlotColormap(cmap)) == C.bool(true)
}

func ImPlot_CreateContext() ImPlotContext {
	return (ImPlotContext)(unsafe.Pointer(C.ImPlot_CreateContext()))
}

// ImPlot_DestroyContextV parameter default value hint:
// ctx: ((void*)0)
func ImPlot_DestroyContextV(ctx ImPlotContext) {
	C.ImPlot_DestroyContextV(ctx.handle())
}

// ImPlot_DragLineXV parameter default value hint:
// flags: 0
// thickness: 1
func ImPlot_DragLineXV(id int32, x *float64, col ImVec4, thickness float32, flags ImPlotDragToolFlags) bool {
	return C.ImPlot_DragLineXV(C.int(id), (*C.double)(x), col.toC(), C.float(thickness), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// ImPlot_DragLineYV parameter default value hint:
// flags: 0
// thickness: 1
func ImPlot_DragLineYV(id int32, y *float64, col ImVec4, thickness float32, flags ImPlotDragToolFlags) bool {
	return C.ImPlot_DragLineYV(C.int(id), (*C.double)(y), col.toC(), C.float(thickness), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// ImPlot_DragPointV parameter default value hint:
// flags: 0
// size: 4
func ImPlot_DragPointV(id int32, x *float64, y *float64, col ImVec4, size float32, flags ImPlotDragToolFlags) bool {
	return C.ImPlot_DragPointV(C.int(id), (*C.double)(x), (*C.double)(y), col.toC(), C.float(size), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// ImPlot_DragRectV parameter default value hint:
// flags: 0
func ImPlot_DragRectV(id int32, x_min *float64, y_min *float64, x_max *float64, y_max *float64, col ImVec4, flags ImPlotDragToolFlags) bool {
	return C.ImPlot_DragRectV(C.int(id), (*C.double)(x_min), (*C.double)(y_min), (*C.double)(x_max), (*C.double)(y_max), col.toC(), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

func ImPlot_EndAlignedPlots() {
	C.ImPlot_EndAlignedPlots()
}

func ImPlot_EndDragDropSource() {
	C.ImPlot_EndDragDropSource()
}

func ImPlot_EndDragDropTarget() {
	C.ImPlot_EndDragDropTarget()
}

func ImPlot_EndLegendPopup() {
	C.ImPlot_EndLegendPopup()
}

func ImPlot_EndPlot() {
	C.ImPlot_EndPlot()
}

func ImPlot_EndSubplots() {
	C.ImPlot_EndSubplots()
}

// ImPlot_GetColormapColorV parameter default value hint:
// cmap: -1
func ImPlot_GetColormapColorV(pOut *ImVec4, idx int32, cmap ImPlotColormap) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetColormapColorV(pOutArg, C.int(idx), C.ImPlotColormap(cmap))
}

func ImPlot_GetColormapCount() int {
	return int(C.ImPlot_GetColormapCount())
}

func ImPlot_GetColormapIndex(name string) ImPlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	return ImPlotColormap(C.ImPlot_GetColormapIndex(nameArg))
}

func ImPlot_GetColormapName(cmap ImPlotColormap) string {
	return C.GoString(C.ImPlot_GetColormapName(C.ImPlotColormap(cmap)))
}

// ImPlot_GetColormapSizeV parameter default value hint:
// cmap: -1
func ImPlot_GetColormapSizeV(cmap ImPlotColormap) int {
	return int(C.ImPlot_GetColormapSizeV(C.ImPlotColormap(cmap)))
}

func ImPlot_GetCurrentContext() ImPlotContext {
	return (ImPlotContext)(unsafe.Pointer(C.ImPlot_GetCurrentContext()))
}

func ImPlot_GetInputMap() ImPlotInputMap {
	return (ImPlotInputMap)(unsafe.Pointer(C.ImPlot_GetInputMap()))
}

func ImPlot_GetLastItemColor(pOut *ImVec4) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetLastItemColor(pOutArg)
}

func ImPlot_GetMarkerName(idx ImPlotMarker) string {
	return C.GoString(C.ImPlot_GetMarkerName(C.ImPlotMarker(idx)))
}

func ImPlot_GetPlotDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.ImPlot_GetPlotDrawList()))
}

// ImPlot_GetPlotMousePosV parameter default value hint:
// x_axis: -1
// y_axis: -1
func ImPlot_GetPlotMousePosV(pOut *ImPlotPoint, x_axis ImAxis, y_axis ImAxis) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetPlotMousePosV(pOutArg, C.ImAxis(x_axis), C.ImAxis(y_axis))
}

func ImPlot_GetPlotPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetPlotPos(pOutArg)
}

func ImPlot_GetPlotSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetPlotSize(pOutArg)
}

func ImPlot_GetStyle() ImPlotStyle {
	return (ImPlotStyle)(unsafe.Pointer(C.ImPlot_GetStyle()))
}

func ImPlot_GetStyleColorName(idx ImPlotCol) string {
	return C.GoString(C.ImPlot_GetStyleColorName(C.ImPlotCol(idx)))
}

// ImPlot_HideNextItemV parameter default value hint:
// cond: ImPlotCond_Once
// hidden: true
func ImPlot_HideNextItemV(hidden bool, cond ImPlotCond) {
	C.ImPlot_HideNextItemV(C.bool(hidden), C.ImPlotCond(cond))
}

func ImPlot_IsAxisHovered(axis ImAxis) bool {
	return C.ImPlot_IsAxisHovered(C.ImAxis(axis)) == C.bool(true)
}

func ImPlot_IsLegendEntryHovered(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.ImPlot_IsLegendEntryHovered(label_idArg) == C.bool(true)
}

func ImPlot_IsPlotHovered() bool {
	return C.ImPlot_IsPlotHovered() == C.bool(true)
}

func ImPlot_IsPlotSelected() bool {
	return C.ImPlot_IsPlotSelected() == C.bool(true)
}

func ImPlot_IsSubplotsHovered() bool {
	return C.ImPlot_IsSubplotsHovered() == C.bool(true)
}

func ImPlot_ItemIcon_U32(col uint32) {
	C.ImPlot_ItemIcon_U32(C.ImU32(col))
}

func ImPlot_ItemIcon_Vec4(col ImVec4) {
	C.ImPlot_ItemIcon_Vec4(col.toC())
}

// ImPlot_MapInputDefaultV parameter default value hint:
// dst: ((void*)0)
func ImPlot_MapInputDefaultV(dst ImPlotInputMap) {
	C.ImPlot_MapInputDefaultV(dst.handle())
}

// ImPlot_MapInputReverseV parameter default value hint:
// dst: ((void*)0)
func ImPlot_MapInputReverseV(dst ImPlotInputMap) {
	C.ImPlot_MapInputReverseV(dst.handle())
}

func ImPlot_NextColormapColor(pOut *ImVec4) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_NextColormapColor(pOutArg)
}

// ImPlot_PixelsToPlot_FloatV parameter default value hint:
// x_axis: -1
// y_axis: -1
func ImPlot_PixelsToPlot_FloatV(pOut *ImPlotPoint, x float32, y float32, x_axis ImAxis, y_axis ImAxis) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_PixelsToPlot_FloatV(pOutArg, C.float(x), C.float(y), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

// ImPlot_PixelsToPlot_Vec2V parameter default value hint:
// x_axis: -1
// y_axis: -1
func ImPlot_PixelsToPlot_Vec2V(pOut *ImPlotPoint, pix ImVec2, x_axis ImAxis, y_axis ImAxis) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_PixelsToPlot_Vec2V(pOutArg, pix.toC(), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

// ImPlot_PlotBars_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotBars_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, bar_size float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotBars_FloatPtrIntV parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(float)
func ImPlot_PlotBars_FloatPtrIntV(label_id string, values []float32, count int32, bar_size float64, shift float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotBars_S64PtrIntV parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(ImS64)
func ImPlot_PlotBars_S64PtrIntV(label_id string, values []int64, count int32, bar_size float64, shift float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotBars_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotBars_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, bar_size float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotBars_U64PtrIntV parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(ImU64)
func ImPlot_PlotBars_U64PtrIntV(label_id string, values []uint64, count int32, bar_size float64, shift float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotBars_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotBars_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, bar_size float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotDigital_FloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotDigital_FloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags ImPlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDigital_FloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotDigital_S64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotDigital_S64PtrV(label_id string, xs []int64, ys []int64, count int32, flags ImPlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDigital_S64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotDigital_U64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotDigital_U64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDigital_U64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotDummyV parameter default value hint:
// flags: 0
func ImPlot_PlotDummyV(label_id string, flags ImPlotDummyFlags) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDummyV(label_idArg, C.ImPlotDummyFlags(flags))
}

// ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, neg []float32, pos []float32, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(neg[0])), (*C.float)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV(label_id string, xs []float32, ys []float32, err []float32, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrIntV(label_id string, xs []int64, ys []int64, err []int64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrIntV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV(label_id string, xs []int64, ys []int64, neg []int64, pos []int64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(neg[0])), (*C.longlong)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrIntV(label_id string, xs []uint64, ys []uint64, err []uint64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrIntV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, neg []uint64, pos []uint64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(neg[0])), (*C.ulonglong)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotInfLines_FloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotInfLines_FloatPtrV(label_id string, values []float32, count int32, flags ImPlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotInfLines_FloatPtrV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotInfLines_S64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotInfLines_S64PtrV(label_id string, values []int64, count int32, flags ImPlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotInfLines_S64PtrV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotInfLines_U64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotInfLines_U64PtrV(label_id string, values []uint64, count int32, flags ImPlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotInfLines_U64PtrV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotLine_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotLine_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotLine_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func ImPlot_PlotLine_FloatPtrIntV(label_id string, values []float32, count int32, xscale float64, xstart float64, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotLine_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func ImPlot_PlotLine_S64PtrIntV(label_id string, values []int64, count int32, xscale float64, xstart float64, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotLine_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotLine_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotLine_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func ImPlot_PlotLine_U64PtrIntV(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotLine_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotLine_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotScatter_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotScatter_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotScatter_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func ImPlot_PlotScatter_FloatPtrIntV(label_id string, values []float32, count int32, xscale float64, xstart float64, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotScatter_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func ImPlot_PlotScatter_S64PtrIntV(label_id string, values []int64, count int32, xscale float64, xstart float64, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotScatter_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotScatter_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotScatter_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func ImPlot_PlotScatter_U64PtrIntV(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotScatter_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotScatter_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtrV(label_id string, xs []float32, ys1 []float32, ys2 []float32, count int32, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys1[0])), (*C.float)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_FloatPtrFloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// yref: 0
func ImPlot_PlotShaded_FloatPtrFloatPtrIntV(label_id string, xs []float32, ys []float32, count int32, yref float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_FloatPtrFloatPtrIntV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
// yref: 0
func ImPlot_PlotShaded_FloatPtrIntV(label_id string, values []float32, count int32, yref float64, xscale float64, xstart float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
// yref: 0
func ImPlot_PlotShaded_S64PtrIntV(label_id string, values []int64, count int32, yref float64, xscale float64, xstart float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_S64PtrS64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// yref: 0
func ImPlot_PlotShaded_S64PtrS64PtrIntV(label_id string, xs []int64, ys []int64, count int32, yref float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_S64PtrS64PtrIntV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_S64PtrS64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotShaded_S64PtrS64PtrS64PtrV(label_id string, xs []int64, ys1 []int64, ys2 []int64, count int32, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_S64PtrS64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys1[0])), (*C.longlong)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
// yref: 0
func ImPlot_PlotShaded_U64PtrIntV(label_id string, values []uint64, count int32, yref float64, xscale float64, xstart float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_U64PtrU64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// yref: 0
func ImPlot_PlotShaded_U64PtrU64PtrIntV(label_id string, xs []uint64, ys []uint64, count int32, yref float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_U64PtrU64PtrIntV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotShaded_U64PtrU64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotShaded_U64PtrU64PtrU64PtrV(label_id string, xs []uint64, ys1 []uint64, ys2 []uint64, count int32, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_U64PtrU64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys1[0])), (*C.ulonglong)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStairs_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func ImPlot_PlotStairs_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStairs_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func ImPlot_PlotStairs_FloatPtrIntV(label_id string, values []float32, count int32, xscale float64, xstart float64, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStairs_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func ImPlot_PlotStairs_S64PtrIntV(label_id string, values []int64, count int32, xscale float64, xstart float64, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStairs_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func ImPlot_PlotStairs_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStairs_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func ImPlot_PlotStairs_U64PtrIntV(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStairs_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func ImPlot_PlotStairs_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStems_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(float)
func ImPlot_PlotStems_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, ref float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStems_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(float)
func ImPlot_PlotStems_FloatPtrIntV(label_id string, values []float32, count int32, ref float64, scale float64, start float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStems_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(ImS64)
func ImPlot_PlotStems_S64PtrIntV(label_id string, values []int64, count int32, ref float64, scale float64, start float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStems_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(ImS64)
func ImPlot_PlotStems_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, ref float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStems_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(ImU64)
func ImPlot_PlotStems_U64PtrIntV(label_id string, values []uint64, count int32, ref float64, scale float64, start float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotStems_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(ImU64)
func ImPlot_PlotStems_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, ref float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// ImPlot_PlotTextV parameter default value hint:
// flags: 0
// pix_offset: ImVec2(0,0)
func ImPlot_PlotTextV(text string, x float64, y float64, pix_offset ImVec2, flags ImPlotTextFlags) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.ImPlot_PlotTextV(textArg, C.double(x), C.double(y), pix_offset.toC(), C.ImPlotTextFlags(flags))
}

// ImPlot_PlotToPixels_doubleV parameter default value hint:
// x_axis: -1
// y_axis: -1
func ImPlot_PlotToPixels_doubleV(pOut *ImVec2, x float64, y float64, x_axis ImAxis, y_axis ImAxis) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_PlotToPixels_doubleV(pOutArg, C.double(x), C.double(y), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

// ImPlot_PopColormapV parameter default value hint:
// count: 1
func ImPlot_PopColormapV(count int32) {
	C.ImPlot_PopColormapV(C.int(count))
}

func ImPlot_PopPlotClipRect() {
	C.ImPlot_PopPlotClipRect()
}

// ImPlot_PopStyleColorV parameter default value hint:
// count: 1
func ImPlot_PopStyleColorV(count int32) {
	C.ImPlot_PopStyleColorV(C.int(count))
}

// ImPlot_PopStyleVarV parameter default value hint:
// count: 1
func ImPlot_PopStyleVarV(count int32) {
	C.ImPlot_PopStyleVarV(C.int(count))
}

func ImPlot_PushColormap_PlotColormap(cmap ImPlotColormap) {
	C.ImPlot_PushColormap_PlotColormap(C.ImPlotColormap(cmap))
}

func ImPlot_PushColormap_Str(name string) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.ImPlot_PushColormap_Str(nameArg)
}

// ImPlot_PushPlotClipRectV parameter default value hint:
// expand: 0
func ImPlot_PushPlotClipRectV(expand float32) {
	C.ImPlot_PushPlotClipRectV(C.float(expand))
}

func ImPlot_PushStyleColor_U32(idx ImPlotCol, col uint32) {
	C.ImPlot_PushStyleColor_U32(C.ImPlotCol(idx), C.ImU32(col))
}

func ImPlot_PushStyleColor_Vec4(idx ImPlotCol, col ImVec4) {
	C.ImPlot_PushStyleColor_Vec4(C.ImPlotCol(idx), col.toC())
}

func ImPlot_PushStyleVar_Float(idx ImPlotStyleVar, val float32) {
	C.ImPlot_PushStyleVar_Float(C.ImPlotStyleVar(idx), C.float(val))
}

func ImPlot_PushStyleVar_Int(idx ImPlotStyleVar, val int32) {
	C.ImPlot_PushStyleVar_Int(C.ImPlotStyleVar(idx), C.int(val))
}

func ImPlot_PushStyleVar_Vec2(idx ImPlotStyleVar, val ImVec2) {
	C.ImPlot_PushStyleVar_Vec2(C.ImPlotStyleVar(idx), val.toC())
}

// ImPlot_SampleColormapV parameter default value hint:
// cmap: -1
func ImPlot_SampleColormapV(pOut *ImVec4, t float32, cmap ImPlotColormap) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_SampleColormapV(pOutArg, C.float(t), C.ImPlotColormap(cmap))
}

func ImPlot_SetAxes(x_axis ImAxis, y_axis ImAxis) {
	C.ImPlot_SetAxes(C.ImAxis(x_axis), C.ImAxis(y_axis))
}

func ImPlot_SetAxis(axis ImAxis) {
	C.ImPlot_SetAxis(C.ImAxis(axis))
}

func ImPlot_SetCurrentContext(ctx ImPlotContext) {
	C.ImPlot_SetCurrentContext(ctx.handle())
}

func ImPlot_SetImGuiContext(ctx ImGuiContext) {
	C.ImPlot_SetImGuiContext(ctx.handle())
}

// ImPlot_SetNextAxesLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func ImPlot_SetNextAxesLimitsV(x_min float64, x_max float64, y_min float64, y_max float64, cond ImPlotCond) {
	C.ImPlot_SetNextAxesLimitsV(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max), C.ImPlotCond(cond))
}

func ImPlot_SetNextAxesToFit() {
	C.ImPlot_SetNextAxesToFit()
}

// ImPlot_SetNextAxisLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func ImPlot_SetNextAxisLimitsV(axis ImAxis, v_min float64, v_max float64, cond ImPlotCond) {
	C.ImPlot_SetNextAxisLimitsV(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.ImPlotCond(cond))
}

func ImPlot_SetNextAxisLinks(axis ImAxis, link_min *float64, link_max *float64) {
	C.ImPlot_SetNextAxisLinks(C.ImAxis(axis), (*C.double)(link_min), (*C.double)(link_max))
}

func ImPlot_SetNextAxisToFit(axis ImAxis) {
	C.ImPlot_SetNextAxisToFit(C.ImAxis(axis))
}

// ImPlot_SetNextErrorBarStyleV parameter default value hint:
// col: ImVec4(0,0,0,-1)
// size: -1
// weight: -1
func ImPlot_SetNextErrorBarStyleV(col ImVec4, size float32, weight float32) {
	C.ImPlot_SetNextErrorBarStyleV(col.toC(), C.float(size), C.float(weight))
}

// ImPlot_SetNextFillStyleV parameter default value hint:
// alpha_mod: -1
// col: ImVec4(0,0,0,-1)
func ImPlot_SetNextFillStyleV(col ImVec4, alpha_mod float32) {
	C.ImPlot_SetNextFillStyleV(col.toC(), C.float(alpha_mod))
}

// ImPlot_SetNextLineStyleV parameter default value hint:
// col: ImVec4(0,0,0,-1)
// weight: -1
func ImPlot_SetNextLineStyleV(col ImVec4, weight float32) {
	C.ImPlot_SetNextLineStyleV(col.toC(), C.float(weight))
}

// ImPlot_SetNextMarkerStyleV parameter default value hint:
// fill: ImVec4(0,0,0,-1)
// marker: -1
// outline: ImVec4(0,0,0,-1)
// size: -1
// weight: -1
func ImPlot_SetNextMarkerStyleV(marker ImPlotMarker, size float32, fill ImVec4, weight float32, outline ImVec4) {
	C.ImPlot_SetNextMarkerStyleV(C.ImPlotMarker(marker), C.float(size), fill.toC(), C.float(weight), outline.toC())
}

// ImPlot_SetupAxesV parameter default value hint:
// x_flags: 0
// y_flags: 0
func ImPlot_SetupAxesV(x_label string, y_label string, x_flags ImPlotAxisFlags, y_flags ImPlotAxisFlags) {
	x_labelArg, x_labelFin := wrapString(x_label)
	defer x_labelFin()

	y_labelArg, y_labelFin := wrapString(y_label)
	defer y_labelFin()

	C.ImPlot_SetupAxesV(x_labelArg, y_labelArg, C.ImPlotAxisFlags(x_flags), C.ImPlotAxisFlags(y_flags))
}

// ImPlot_SetupAxesLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func ImPlot_SetupAxesLimitsV(x_min float64, x_max float64, y_min float64, y_max float64, cond ImPlotCond) {
	C.ImPlot_SetupAxesLimitsV(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max), C.ImPlotCond(cond))
}

// ImPlot_SetupAxisV parameter default value hint:
// flags: 0
// label: ((void*)0)
func ImPlot_SetupAxisV(axis ImAxis, label string, flags ImPlotAxisFlags) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ImPlot_SetupAxisV(C.ImAxis(axis), labelArg, C.ImPlotAxisFlags(flags))
}

func ImPlot_SetupAxisFormat_Str(axis ImAxis, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.ImPlot_SetupAxisFormat_Str(C.ImAxis(axis), fmtArg)
}

// ImPlot_SetupAxisLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func ImPlot_SetupAxisLimitsV(axis ImAxis, v_min float64, v_max float64, cond ImPlotCond) {
	C.ImPlot_SetupAxisLimitsV(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.ImPlotCond(cond))
}

func ImPlot_SetupAxisLimitsConstraints(axis ImAxis, v_min float64, v_max float64) {
	C.ImPlot_SetupAxisLimitsConstraints(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func ImPlot_SetupAxisLinks(axis ImAxis, link_min *float64, link_max *float64) {
	C.ImPlot_SetupAxisLinks(C.ImAxis(axis), (*C.double)(link_min), (*C.double)(link_max))
}

func ImPlot_SetupAxisScale_PlotScale(axis ImAxis, scale ImPlotScale) {
	C.ImPlot_SetupAxisScale_PlotScale(C.ImAxis(axis), C.ImPlotScale(scale))
}

func ImPlot_SetupAxisZoomConstraints(axis ImAxis, z_min float64, z_max float64) {
	C.ImPlot_SetupAxisZoomConstraints(C.ImAxis(axis), C.double(z_min), C.double(z_max))
}

func ImPlot_SetupFinish() {
	C.ImPlot_SetupFinish()
}

// ImPlot_SetupLegendV parameter default value hint:
// flags: 0
func ImPlot_SetupLegendV(location ImPlotLocation, flags ImPlotLegendFlags) {
	C.ImPlot_SetupLegendV(C.ImPlotLocation(location), C.ImPlotLegendFlags(flags))
}

// ImPlot_SetupMouseTextV parameter default value hint:
// flags: 0
func ImPlot_SetupMouseTextV(location ImPlotLocation, flags ImPlotMouseTextFlags) {
	C.ImPlot_SetupMouseTextV(C.ImPlotLocation(location), C.ImPlotMouseTextFlags(flags))
}

func ImPlot_ShowColormapSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ImPlot_ShowColormapSelector(labelArg) == C.bool(true)
}

// ImPlot_ShowDemoWindowV parameter default value hint:
// p_open: ((void*)0)
func ImPlot_ShowDemoWindowV(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ImPlot_ShowDemoWindowV(p_openArg)
}

func ImPlot_ShowInputMapSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ImPlot_ShowInputMapSelector(labelArg) == C.bool(true)
}

// ImPlot_ShowMetricsWindowV parameter default value hint:
// p_popen: ((void*)0)
func ImPlot_ShowMetricsWindowV(p_popen *bool) {
	p_popenArg, p_popenFin := wrapBool(p_popen)
	defer p_popenFin()

	C.ImPlot_ShowMetricsWindowV(p_popenArg)
}

// ImPlot_ShowStyleEditorV parameter default value hint:
// ref: ((void*)0)
func ImPlot_ShowStyleEditorV(ref ImPlotStyle) {
	C.ImPlot_ShowStyleEditorV(ref.handle())
}

func ImPlot_ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ImPlot_ShowStyleSelector(labelArg) == C.bool(true)
}

func ImPlot_ShowUserGuide() {
	C.ImPlot_ShowUserGuide()
}

// ImPlot_StyleColorsAutoV parameter default value hint:
// dst: ((void*)0)
func ImPlot_StyleColorsAutoV(dst ImPlotStyle) {
	C.ImPlot_StyleColorsAutoV(dst.handle())
}

// ImPlot_StyleColorsClassicV parameter default value hint:
// dst: ((void*)0)
func ImPlot_StyleColorsClassicV(dst ImPlotStyle) {
	C.ImPlot_StyleColorsClassicV(dst.handle())
}

// ImPlot_StyleColorsDarkV parameter default value hint:
// dst: ((void*)0)
func ImPlot_StyleColorsDarkV(dst ImPlotStyle) {
	C.ImPlot_StyleColorsDarkV(dst.handle())
}

// ImPlot_StyleColorsLightV parameter default value hint:
// dst: ((void*)0)
func ImPlot_StyleColorsLightV(dst ImPlotStyle) {
	C.ImPlot_StyleColorsLightV(dst.handle())
}

// ImPlot_TagX_BoolV parameter default value hint:
// round: false
func ImPlot_TagX_BoolV(x float64, col ImVec4, round bool) {
	C.ImPlot_TagX_BoolV(C.double(x), col.toC(), C.bool(round))
}

func ImPlot_TagX_Str(x float64, col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.ImPlot_TagX_Str(C.double(x), col.toC(), fmtArg)
}

// ImPlot_TagY_BoolV parameter default value hint:
// round: false
func ImPlot_TagY_BoolV(y float64, col ImVec4, round bool) {
	C.ImPlot_TagY_BoolV(C.double(y), col.toC(), C.bool(round))
}

func ImPlot_TagY_Str(y float64, col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.ImPlot_TagY_Str(C.double(y), col.toC(), fmtArg)
}

func ImPlot_AddColormap_Vec4Ptr(name string, cols *ImVec4, size int32) ImPlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	colsArg, colsFin := cols.wrap()
	defer colsFin()

	return ImPlotColormap(C.ImPlot_AddColormap_Vec4Ptr(nameArg, colsArg, C.int(size)))
}

func ImPlot_Annotation_Bool(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool) {
	C.ImPlot_Annotation_Bool(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp))
}

func ImPlot_BeginAlignedPlots(group_id string) bool {
	group_idArg, group_idFin := wrapString(group_id)
	defer group_idFin()

	return C.ImPlot_BeginAlignedPlots(group_idArg) == C.bool(true)
}

func ImPlot_BeginDragDropSourceAxis(axis ImAxis) bool {
	return C.ImPlot_BeginDragDropSourceAxis(C.ImAxis(axis)) == C.bool(true)
}

func ImPlot_BeginDragDropSourceItem(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.ImPlot_BeginDragDropSourceItem(label_idArg) == C.bool(true)
}

func ImPlot_BeginDragDropSourcePlot() bool {
	return C.ImPlot_BeginDragDropSourcePlot() == C.bool(true)
}

func ImPlot_BeginLegendPopup(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.ImPlot_BeginLegendPopup(label_idArg) == C.bool(true)
}

func ImPlot_BeginPlot(title_id string) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.ImPlot_BeginPlot(title_idArg) == C.bool(true)
}

func ImPlot_BeginSubplots(title_id string, rows int32, cols int32, size ImVec2) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.ImPlot_BeginSubplots(title_idArg, C.int(rows), C.int(cols), size.toC()) == C.bool(true)
}

func ImPlot_BustColorCache() {
	C.ImPlot_BustColorCache()
}

func ImPlot_ColormapButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ImPlot_ColormapButton(labelArg) == C.bool(true)
}

func ImPlot_ColormapScale(label string, scale_min float64, scale_max float64) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ImPlot_ColormapScale(labelArg, C.double(scale_min), C.double(scale_max))
}

func ImPlot_ColormapSlider(label string, t *float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	tArg, tFin := wrapFloat(t)
	defer tFin()

	return C.ImPlot_ColormapSlider(labelArg, tArg) == C.bool(true)
}

func ImPlot_DestroyContext() {
	C.ImPlot_DestroyContext()
}

func ImPlot_DragLineX(id int32, x *float64, col ImVec4) bool {
	return C.ImPlot_DragLineX(C.int(id), (*C.double)(x), col.toC()) == C.bool(true)
}

func ImPlot_DragLineY(id int32, y *float64, col ImVec4) bool {
	return C.ImPlot_DragLineY(C.int(id), (*C.double)(y), col.toC()) == C.bool(true)
}

func ImPlot_DragPoint(id int32, x *float64, y *float64, col ImVec4) bool {
	return C.ImPlot_DragPoint(C.int(id), (*C.double)(x), (*C.double)(y), col.toC()) == C.bool(true)
}

func ImPlot_DragRect(id int32, x_min *float64, y_min *float64, x_max *float64, y_max *float64, col ImVec4) bool {
	return C.ImPlot_DragRect(C.int(id), (*C.double)(x_min), (*C.double)(y_min), (*C.double)(x_max), (*C.double)(y_max), col.toC()) == C.bool(true)
}

func ImPlot_GetColormapColor(pOut *ImVec4, idx int32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetColormapColor(pOutArg, C.int(idx))
}

func ImPlot_GetColormapSize() int {
	return int(C.ImPlot_GetColormapSize())
}

func ImPlot_GetPlotMousePos(pOut *ImPlotPoint) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_GetPlotMousePos(pOutArg)
}

func ImPlot_HideNextItem() {
	C.ImPlot_HideNextItem()
}

func ImPlot_MapInputDefault() {
	C.ImPlot_MapInputDefault()
}

func ImPlot_MapInputReverse() {
	C.ImPlot_MapInputReverse()
}

func ImPlot_PixelsToPlot_Float(pOut *ImPlotPoint, x float32, y float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_PixelsToPlot_Float(pOutArg, C.float(x), C.float(y))
}

func ImPlot_PixelsToPlot_Vec2(pOut *ImPlotPoint, pix ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_PixelsToPlot_Vec2(pOutArg, pix.toC())
}

func ImPlot_PlotBars_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32, bar_size float64) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(bar_size))
}

func ImPlot_PlotBars_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotBars_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotBars_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, bar_size float64) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(bar_size))
}

func ImPlot_PlotBars_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotBars_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, bar_size float64) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotBars_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(bar_size))
}

func ImPlot_PlotDigital_FloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDigital_FloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func ImPlot_PlotDigital_S64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDigital_S64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotDigital_U64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDigital_U64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotDummy(label_id string) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotDummy(label_idArg)
}

func ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_id string, xs []float32, ys []float32, neg []float32, pos []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(neg[0])), (*C.float)(&(pos[0])), C.int(count))
}

func ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_id string, xs []float32, ys []float32, err []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(err[0])), C.int(count))
}

func ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_id string, xs []int64, ys []int64, err []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(err[0])), C.int(count))
}

func ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_id string, xs []int64, ys []int64, neg []int64, pos []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(neg[0])), (*C.longlong)(&(pos[0])), C.int(count))
}

func ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_id string, xs []uint64, ys []uint64, err []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(err[0])), C.int(count))
}

func ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, neg []uint64, pos []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(neg[0])), (*C.ulonglong)(&(pos[0])), C.int(count))
}

func ImPlot_PlotHeatmap_FloatPtr(label_id string, values []float32, rows int32, cols int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotHeatmap_FloatPtr(label_idArg, (*C.float)(&(values[0])), C.int(rows), C.int(cols))
}

func ImPlot_PlotHeatmap_S64Ptr(label_id string, values []int64, rows int32, cols int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotHeatmap_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(rows), C.int(cols))
}

func ImPlot_PlotHeatmap_U64Ptr(label_id string, values []uint64, rows int32, cols int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotHeatmap_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(rows), C.int(cols))
}

func ImPlot_PlotHistogram2D_FloatPtr(label_id string, xs []float32, ys []float32, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.ImPlot_PlotHistogram2D_FloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count)))
}

func ImPlot_PlotHistogram2D_S64Ptr(label_id string, xs []int64, ys []int64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.ImPlot_PlotHistogram2D_S64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count)))
}

func ImPlot_PlotHistogram2D_U64Ptr(label_id string, xs []uint64, ys []uint64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.ImPlot_PlotHistogram2D_U64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count)))
}

func ImPlot_PlotHistogram_FloatPtr(label_id string, values []float32, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.ImPlot_PlotHistogram_FloatPtr(label_idArg, (*C.float)(&(values[0])), C.int(count)))
}

func ImPlot_PlotHistogram_S64Ptr(label_id string, values []int64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.ImPlot_PlotHistogram_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(count)))
}

func ImPlot_PlotHistogram_U64Ptr(label_id string, values []uint64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.ImPlot_PlotHistogram_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count)))
}

func ImPlot_PlotInfLines_FloatPtr(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotInfLines_FloatPtr(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotInfLines_S64Ptr(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotInfLines_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotInfLines_U64Ptr(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotInfLines_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotLine_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func ImPlot_PlotLine_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotLine_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotLine_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotLine_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotLine_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotLine_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotScatter_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func ImPlot_PlotScatter_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotScatter_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotScatter_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotScatter_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotScatter_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotScatter_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_id string, xs []float32, ys1 []float32, ys2 []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys1[0])), (*C.float)(&(ys2[0])), C.int(count))
}

func ImPlot_PlotShaded_FloatPtrFloatPtrInt(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_FloatPtrFloatPtrInt(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func ImPlot_PlotShaded_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotShaded_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotShaded_S64PtrS64PtrInt(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_S64PtrS64PtrInt(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotShaded_S64PtrS64PtrS64Ptr(label_id string, xs []int64, ys1 []int64, ys2 []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_S64PtrS64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys1[0])), (*C.longlong)(&(ys2[0])), C.int(count))
}

func ImPlot_PlotShaded_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotShaded_U64PtrU64PtrInt(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_U64PtrU64PtrInt(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotShaded_U64PtrU64PtrU64Ptr(label_id string, xs []uint64, ys1 []uint64, ys2 []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotShaded_U64PtrU64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys1[0])), (*C.ulonglong)(&(ys2[0])), C.int(count))
}

func ImPlot_PlotStairs_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func ImPlot_PlotStairs_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotStairs_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotStairs_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotStairs_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotStairs_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStairs_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotStems_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func ImPlot_PlotStems_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func ImPlot_PlotStems_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func ImPlot_PlotStems_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotStems_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func ImPlot_PlotStems_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.ImPlot_PlotStems_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func ImPlot_PlotText(text string, x float64, y float64) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.ImPlot_PlotText(textArg, C.double(x), C.double(y))
}

func ImPlot_PlotToPixels_double(pOut *ImVec2, x float64, y float64) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_PlotToPixels_double(pOutArg, C.double(x), C.double(y))
}

func ImPlot_PopColormap() {
	C.ImPlot_PopColormap()
}

func ImPlot_PopStyleColor() {
	C.ImPlot_PopStyleColor()
}

func ImPlot_PopStyleVar() {
	C.ImPlot_PopStyleVar()
}

func ImPlot_PushPlotClipRect() {
	C.ImPlot_PushPlotClipRect()
}

func ImPlot_SampleColormap(pOut *ImVec4, t float32) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.ImPlot_SampleColormap(pOutArg, C.float(t))
}

func ImPlot_SetNextAxesLimits(x_min float64, x_max float64, y_min float64, y_max float64) {
	C.ImPlot_SetNextAxesLimits(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max))
}

func ImPlot_SetNextAxisLimits(axis ImAxis, v_min float64, v_max float64) {
	C.ImPlot_SetNextAxisLimits(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func ImPlot_SetNextErrorBarStyle() {
	C.ImPlot_SetNextErrorBarStyle()
}

func ImPlot_SetNextFillStyle() {
	C.ImPlot_SetNextFillStyle()
}

func ImPlot_SetNextLineStyle() {
	C.ImPlot_SetNextLineStyle()
}

func ImPlot_SetNextMarkerStyle() {
	C.ImPlot_SetNextMarkerStyle()
}

func ImPlot_SetupAxes(x_label string, y_label string) {
	x_labelArg, x_labelFin := wrapString(x_label)
	defer x_labelFin()

	y_labelArg, y_labelFin := wrapString(y_label)
	defer y_labelFin()

	C.ImPlot_SetupAxes(x_labelArg, y_labelArg)
}

func ImPlot_SetupAxesLimits(x_min float64, x_max float64, y_min float64, y_max float64) {
	C.ImPlot_SetupAxesLimits(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max))
}

func ImPlot_SetupAxis(axis ImAxis) {
	C.ImPlot_SetupAxis(C.ImAxis(axis))
}

func ImPlot_SetupAxisLimits(axis ImAxis, v_min float64, v_max float64) {
	C.ImPlot_SetupAxisLimits(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func ImPlot_SetupAxisTicks_double(axis ImAxis, v_min float64, v_max float64, n_ticks int32) {
	C.ImPlot_SetupAxisTicks_double(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.int(n_ticks))
}

func ImPlot_SetupLegend(location ImPlotLocation) {
	C.ImPlot_SetupLegend(C.ImPlotLocation(location))
}

func ImPlot_SetupMouseText(location ImPlotLocation) {
	C.ImPlot_SetupMouseText(C.ImPlotLocation(location))
}

func ImPlot_ShowDemoWindow() {
	C.ImPlot_ShowDemoWindow()
}

func ImPlot_ShowMetricsWindow() {
	C.ImPlot_ShowMetricsWindow()
}

func ImPlot_ShowStyleEditor() {
	C.ImPlot_ShowStyleEditor()
}

func ImPlot_StyleColorsAuto() {
	C.ImPlot_StyleColorsAuto()
}

func ImPlot_StyleColorsClassic() {
	C.ImPlot_StyleColorsClassic()
}

func ImPlot_StyleColorsDark() {
	C.ImPlot_StyleColorsDark()
}

func ImPlot_StyleColorsLight() {
	C.ImPlot_StyleColorsLight()
}

func ImPlot_TagX_Bool(x float64, col ImVec4) {
	C.ImPlot_TagX_Bool(C.double(x), col.toC())
}

func ImPlot_TagY_Bool(y float64, col ImVec4) {
	C.ImPlot_TagY_Bool(C.double(y), col.toC())
}

func (self ImPlotAlignmentData) SetVertical(v bool) {
	C.ImPlotAlignmentData_SetVertical(self.handle(), C.bool(v))
}

func (self ImPlotAlignmentData) GetVertical() bool {
	return C.ImPlotAlignmentData_GetVertical(self.handle()) == C.bool(true)
}

func (self ImPlotAlignmentData) SetPadA(v float32) {
	C.ImPlotAlignmentData_SetPadA(self.handle(), C.float(v))
}

func (self ImPlotAlignmentData) GetPadA() float32 {
	return float32(C.ImPlotAlignmentData_GetPadA(self.handle()))
}

func (self ImPlotAlignmentData) SetPadB(v float32) {
	C.ImPlotAlignmentData_SetPadB(self.handle(), C.float(v))
}

func (self ImPlotAlignmentData) GetPadB() float32 {
	return float32(C.ImPlotAlignmentData_GetPadB(self.handle()))
}

func (self ImPlotAlignmentData) SetPadAMax(v float32) {
	C.ImPlotAlignmentData_SetPadAMax(self.handle(), C.float(v))
}

func (self ImPlotAlignmentData) GetPadAMax() float32 {
	return float32(C.ImPlotAlignmentData_GetPadAMax(self.handle()))
}

func (self ImPlotAlignmentData) SetPadBMax(v float32) {
	C.ImPlotAlignmentData_SetPadBMax(self.handle(), C.float(v))
}

func (self ImPlotAlignmentData) GetPadBMax() float32 {
	return float32(C.ImPlotAlignmentData_GetPadBMax(self.handle()))
}

func (self ImPlotAnnotation) SetPos(v ImVec2) {
	C.ImPlotAnnotation_SetPos(self.handle(), v.toC())
}

func (self ImPlotAnnotation) GetPos() ImVec2 {
	return newImVec2FromC(C.ImPlotAnnotation_GetPos(self.handle()))
}

func (self ImPlotAnnotation) SetOffset(v ImVec2) {
	C.ImPlotAnnotation_SetOffset(self.handle(), v.toC())
}

func (self ImPlotAnnotation) GetOffset() ImVec2 {
	return newImVec2FromC(C.ImPlotAnnotation_GetOffset(self.handle()))
}

func (self ImPlotAnnotation) SetColorBg(v uint32) {
	C.ImPlotAnnotation_SetColorBg(self.handle(), C.ImU32(v))
}

func (self ImPlotAnnotation) GetColorBg() uint32 {
	return uint32(C.ImPlotAnnotation_GetColorBg(self.handle()))
}

func (self ImPlotAnnotation) SetColorFg(v uint32) {
	C.ImPlotAnnotation_SetColorFg(self.handle(), C.ImU32(v))
}

func (self ImPlotAnnotation) GetColorFg() uint32 {
	return uint32(C.ImPlotAnnotation_GetColorFg(self.handle()))
}

func (self ImPlotAnnotation) SetTextOffset(v int32) {
	C.ImPlotAnnotation_SetTextOffset(self.handle(), C.int(v))
}

func (self ImPlotAnnotation) GetTextOffset() int {
	return int(C.ImPlotAnnotation_GetTextOffset(self.handle()))
}

func (self ImPlotAnnotation) SetClamp(v bool) {
	C.ImPlotAnnotation_SetClamp(self.handle(), C.bool(v))
}

func (self ImPlotAnnotation) GetClamp() bool {
	return C.ImPlotAnnotation_GetClamp(self.handle()) == C.bool(true)
}

func (self ImPlotAnnotationCollection) GetTextBuffer() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotAnnotationCollection_GetTextBuffer(self.handle()))
}

func (self ImPlotAnnotationCollection) SetSize(v int32) {
	C.ImPlotAnnotationCollection_SetSize(self.handle(), C.int(v))
}

func (self ImPlotAnnotationCollection) GetSize() int {
	return int(C.ImPlotAnnotationCollection_GetSize(self.handle()))
}

func (self ImPlotAxis) SetID(v ImGuiID) {
	C.ImPlotAxis_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImPlotAxis) GetID() ImGuiID {
	return ImGuiID(C.ImPlotAxis_GetID(self.handle()))
}

func (self ImPlotAxis) SetFlags(v ImPlotAxisFlags) {
	C.ImPlotAxis_SetFlags(self.handle(), C.ImPlotAxisFlags(v))
}

func (self ImPlotAxis) GetFlags() ImPlotAxisFlags {
	return ImPlotAxisFlags(C.ImPlotAxis_GetFlags(self.handle()))
}

func (self ImPlotAxis) SetPreviousFlags(v ImPlotAxisFlags) {
	C.ImPlotAxis_SetPreviousFlags(self.handle(), C.ImPlotAxisFlags(v))
}

func (self ImPlotAxis) GetPreviousFlags() ImPlotAxisFlags {
	return ImPlotAxisFlags(C.ImPlotAxis_GetPreviousFlags(self.handle()))
}

func (self ImPlotAxis) GetRange() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotAxis_GetRange(self.handle()))
}

func (self ImPlotAxis) SetRangeCond(v ImPlotCond) {
	C.ImPlotAxis_SetRangeCond(self.handle(), C.ImPlotCond(v))
}

func (self ImPlotAxis) GetRangeCond() ImPlotCond {
	return ImPlotCond(C.ImPlotAxis_GetRangeCond(self.handle()))
}

func (self ImPlotAxis) SetScale(v ImPlotScale) {
	C.ImPlotAxis_SetScale(self.handle(), C.ImPlotScale(v))
}

func (self ImPlotAxis) GetScale() ImPlotScale {
	return ImPlotScale(C.ImPlotAxis_GetScale(self.handle()))
}

func (self ImPlotAxis) GetFitExtents() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotAxis_GetFitExtents(self.handle()))
}

func (self ImPlotAxis) SetOrthoAxis(v ImPlotAxis) {
	C.ImPlotAxis_SetOrthoAxis(self.handle(), v.handle())
}

func (self ImPlotAxis) GetOrthoAxis() ImPlotAxis {
	return (ImPlotAxis)(unsafe.Pointer(C.ImPlotAxis_GetOrthoAxis(self.handle())))
}

func (self ImPlotAxis) GetConstraintRange() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotAxis_GetConstraintRange(self.handle()))
}

func (self ImPlotAxis) GetConstraintZoom() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotAxis_GetConstraintZoom(self.handle()))
}

func (self ImPlotAxis) GetTicker() ImPlotTicker {
	return newImPlotTickerFromC(C.ImPlotAxis_GetTicker(self.handle()))
}

func (self ImPlotAxis) SetFormatterData(v unsafe.Pointer) {
	C.ImPlotAxis_SetFormatterData(self.handle(), v)
}

func (self ImPlotAxis) GetFormatterData() unsafe.Pointer {
	return unsafe.Pointer(C.ImPlotAxis_GetFormatterData(self.handle()))
}

func (self ImPlotAxis) SetLinkedMin(v *float64) {
	C.ImPlotAxis_SetLinkedMin(self.handle(), (*C.double)(v))
}

func (self ImPlotAxis) SetLinkedMax(v *float64) {
	C.ImPlotAxis_SetLinkedMax(self.handle(), (*C.double)(v))
}

func (self ImPlotAxis) SetPickerLevel(v int32) {
	C.ImPlotAxis_SetPickerLevel(self.handle(), C.int(v))
}

func (self ImPlotAxis) GetPickerLevel() int {
	return int(C.ImPlotAxis_GetPickerLevel(self.handle()))
}

func (self ImPlotAxis) GetPickerTimeMin() ImPlotTime {
	return newImPlotTimeFromC(C.ImPlotAxis_GetPickerTimeMin(self.handle()))
}

func (self ImPlotAxis) GetPickerTimeMax() ImPlotTime {
	return newImPlotTimeFromC(C.ImPlotAxis_GetPickerTimeMax(self.handle()))
}

func (self ImPlotAxis) SetTransformData(v unsafe.Pointer) {
	C.ImPlotAxis_SetTransformData(self.handle(), v)
}

func (self ImPlotAxis) GetTransformData() unsafe.Pointer {
	return unsafe.Pointer(C.ImPlotAxis_GetTransformData(self.handle()))
}

func (self ImPlotAxis) SetPixelMin(v float32) {
	C.ImPlotAxis_SetPixelMin(self.handle(), C.float(v))
}

func (self ImPlotAxis) GetPixelMin() float32 {
	return float32(C.ImPlotAxis_GetPixelMin(self.handle()))
}

func (self ImPlotAxis) SetPixelMax(v float32) {
	C.ImPlotAxis_SetPixelMax(self.handle(), C.float(v))
}

func (self ImPlotAxis) GetPixelMax() float32 {
	return float32(C.ImPlotAxis_GetPixelMax(self.handle()))
}

func (self ImPlotAxis) SetScaleMin(v float64) {
	C.ImPlotAxis_SetScaleMin(self.handle(), C.double(v))
}

func (self ImPlotAxis) GetScaleMin() float64 {
	return float64(C.ImPlotAxis_GetScaleMin(self.handle()))
}

func (self ImPlotAxis) SetScaleMax(v float64) {
	C.ImPlotAxis_SetScaleMax(self.handle(), C.double(v))
}

func (self ImPlotAxis) GetScaleMax() float64 {
	return float64(C.ImPlotAxis_GetScaleMax(self.handle()))
}

func (self ImPlotAxis) SetScaleToPixel(v float64) {
	C.ImPlotAxis_SetScaleToPixel(self.handle(), C.double(v))
}

func (self ImPlotAxis) GetScaleToPixel() float64 {
	return float64(C.ImPlotAxis_GetScaleToPixel(self.handle()))
}

func (self ImPlotAxis) SetDatum1(v float32) {
	C.ImPlotAxis_SetDatum1(self.handle(), C.float(v))
}

func (self ImPlotAxis) GetDatum1() float32 {
	return float32(C.ImPlotAxis_GetDatum1(self.handle()))
}

func (self ImPlotAxis) SetDatum2(v float32) {
	C.ImPlotAxis_SetDatum2(self.handle(), C.float(v))
}

func (self ImPlotAxis) GetDatum2() float32 {
	return float32(C.ImPlotAxis_GetDatum2(self.handle()))
}

func (self ImPlotAxis) SetHoverRect(v ImRect) {
	C.ImPlotAxis_SetHoverRect(self.handle(), v.toC())
}

func (self ImPlotAxis) GetHoverRect() ImRect {
	return newImRectFromC(C.ImPlotAxis_GetHoverRect(self.handle()))
}

func (self ImPlotAxis) SetLabelOffset(v int32) {
	C.ImPlotAxis_SetLabelOffset(self.handle(), C.int(v))
}

func (self ImPlotAxis) GetLabelOffset() int {
	return int(C.ImPlotAxis_GetLabelOffset(self.handle()))
}

func (self ImPlotAxis) SetColorMaj(v uint32) {
	C.ImPlotAxis_SetColorMaj(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorMaj() uint32 {
	return uint32(C.ImPlotAxis_GetColorMaj(self.handle()))
}

func (self ImPlotAxis) SetColorMin(v uint32) {
	C.ImPlotAxis_SetColorMin(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorMin() uint32 {
	return uint32(C.ImPlotAxis_GetColorMin(self.handle()))
}

func (self ImPlotAxis) SetColorTick(v uint32) {
	C.ImPlotAxis_SetColorTick(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorTick() uint32 {
	return uint32(C.ImPlotAxis_GetColorTick(self.handle()))
}

func (self ImPlotAxis) SetColorTxt(v uint32) {
	C.ImPlotAxis_SetColorTxt(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorTxt() uint32 {
	return uint32(C.ImPlotAxis_GetColorTxt(self.handle()))
}

func (self ImPlotAxis) SetColorBg(v uint32) {
	C.ImPlotAxis_SetColorBg(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorBg() uint32 {
	return uint32(C.ImPlotAxis_GetColorBg(self.handle()))
}

func (self ImPlotAxis) SetColorHov(v uint32) {
	C.ImPlotAxis_SetColorHov(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorHov() uint32 {
	return uint32(C.ImPlotAxis_GetColorHov(self.handle()))
}

func (self ImPlotAxis) SetColorAct(v uint32) {
	C.ImPlotAxis_SetColorAct(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorAct() uint32 {
	return uint32(C.ImPlotAxis_GetColorAct(self.handle()))
}

func (self ImPlotAxis) SetColorHiLi(v uint32) {
	C.ImPlotAxis_SetColorHiLi(self.handle(), C.ImU32(v))
}

func (self ImPlotAxis) GetColorHiLi() uint32 {
	return uint32(C.ImPlotAxis_GetColorHiLi(self.handle()))
}

func (self ImPlotAxis) SetEnabled(v bool) {
	C.ImPlotAxis_SetEnabled(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetEnabled() bool {
	return C.ImPlotAxis_GetEnabled(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetVertical(v bool) {
	C.ImPlotAxis_SetVertical(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetVertical() bool {
	return C.ImPlotAxis_GetVertical(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetFitThisFrame(v bool) {
	C.ImPlotAxis_SetFitThisFrame(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetFitThisFrame() bool {
	return C.ImPlotAxis_GetFitThisFrame(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetHasRange(v bool) {
	C.ImPlotAxis_SetHasRange(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetHasRange() bool {
	return C.ImPlotAxis_GetHasRange(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetHasFormatSpec(v bool) {
	C.ImPlotAxis_SetHasFormatSpec(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetHasFormatSpec() bool {
	return C.ImPlotAxis_GetHasFormatSpec(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetShowDefaultTicks(v bool) {
	C.ImPlotAxis_SetShowDefaultTicks(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetShowDefaultTicks() bool {
	return C.ImPlotAxis_GetShowDefaultTicks(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetHovered(v bool) {
	C.ImPlotAxis_SetHovered(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetHovered() bool {
	return C.ImPlotAxis_GetHovered(self.handle()) == C.bool(true)
}

func (self ImPlotAxis) SetHeld(v bool) {
	C.ImPlotAxis_SetHeld(self.handle(), C.bool(v))
}

func (self ImPlotAxis) GetHeld() bool {
	return C.ImPlotAxis_GetHeld(self.handle()) == C.bool(true)
}

func (self ImPlotColormapData) GetText() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotColormapData_GetText(self.handle()))
}

func (self ImPlotColormapData) GetMap() ImGuiStorage {
	return newImGuiStorageFromC(C.ImPlotColormapData_GetMap(self.handle()))
}

func (self ImPlotColormapData) SetCount(v int32) {
	C.ImPlotColormapData_SetCount(self.handle(), C.int(v))
}

func (self ImPlotColormapData) GetCount() int {
	return int(C.ImPlotColormapData_GetCount(self.handle()))
}

func (self ImPlotContext) SetCurrentPlot(v ImPlotPlot) {
	C.ImPlotContext_SetCurrentPlot(self.handle(), v.handle())
}

func (self ImPlotContext) GetCurrentPlot() ImPlotPlot {
	return (ImPlotPlot)(unsafe.Pointer(C.ImPlotContext_GetCurrentPlot(self.handle())))
}

func (self ImPlotContext) SetCurrentSubplot(v ImPlotSubplot) {
	C.ImPlotContext_SetCurrentSubplot(self.handle(), v.handle())
}

func (self ImPlotContext) GetCurrentSubplot() ImPlotSubplot {
	return (ImPlotSubplot)(unsafe.Pointer(C.ImPlotContext_GetCurrentSubplot(self.handle())))
}

func (self ImPlotContext) SetCurrentItems(v ImPlotItemGroup) {
	C.ImPlotContext_SetCurrentItems(self.handle(), v.handle())
}

func (self ImPlotContext) GetCurrentItems() ImPlotItemGroup {
	return (ImPlotItemGroup)(unsafe.Pointer(C.ImPlotContext_GetCurrentItems(self.handle())))
}

func (self ImPlotContext) SetCurrentItem(v ImPlotItem) {
	C.ImPlotContext_SetCurrentItem(self.handle(), v.handle())
}

func (self ImPlotContext) GetCurrentItem() ImPlotItem {
	return (ImPlotItem)(unsafe.Pointer(C.ImPlotContext_GetCurrentItem(self.handle())))
}

func (self ImPlotContext) SetPreviousItem(v ImPlotItem) {
	C.ImPlotContext_SetPreviousItem(self.handle(), v.handle())
}

func (self ImPlotContext) GetPreviousItem() ImPlotItem {
	return (ImPlotItem)(unsafe.Pointer(C.ImPlotContext_GetPreviousItem(self.handle())))
}

func (self ImPlotContext) GetCTicker() ImPlotTicker {
	return newImPlotTickerFromC(C.ImPlotContext_GetCTicker(self.handle()))
}

func (self ImPlotContext) GetAnnotations() ImPlotAnnotationCollection {
	return newImPlotAnnotationCollectionFromC(C.ImPlotContext_GetAnnotations(self.handle()))
}

func (self ImPlotContext) GetTags() ImPlotTagCollection {
	return newImPlotTagCollectionFromC(C.ImPlotContext_GetTags(self.handle()))
}

func (self ImPlotContext) SetChildWindowMade(v bool) {
	C.ImPlotContext_SetChildWindowMade(self.handle(), C.bool(v))
}

func (self ImPlotContext) GetChildWindowMade() bool {
	return C.ImPlotContext_GetChildWindowMade(self.handle()) == C.bool(true)
}

func (self ImPlotContext) GetStyle() ImPlotStyle {
	return newImPlotStyleFromC(C.ImPlotContext_GetStyle(self.handle()))
}

func (self ImPlotContext) GetColormapData() ImPlotColormapData {
	return newImPlotColormapDataFromC(C.ImPlotContext_GetColormapData(self.handle()))
}

func (self ImPlotContext) SetDigitalPlotItemCnt(v int32) {
	C.ImPlotContext_SetDigitalPlotItemCnt(self.handle(), C.int(v))
}

func (self ImPlotContext) GetDigitalPlotItemCnt() int {
	return int(C.ImPlotContext_GetDigitalPlotItemCnt(self.handle()))
}

func (self ImPlotContext) SetDigitalPlotOffset(v int32) {
	C.ImPlotContext_SetDigitalPlotOffset(self.handle(), C.int(v))
}

func (self ImPlotContext) GetDigitalPlotOffset() int {
	return int(C.ImPlotContext_GetDigitalPlotOffset(self.handle()))
}

func (self ImPlotContext) GetNextPlotData() ImPlotNextPlotData {
	return newImPlotNextPlotDataFromC(C.ImPlotContext_GetNextPlotData(self.handle()))
}

func (self ImPlotContext) GetNextItemData() ImPlotNextItemData {
	return newImPlotNextItemDataFromC(C.ImPlotContext_GetNextItemData(self.handle()))
}

func (self ImPlotContext) GetInputMap() ImPlotInputMap {
	return newImPlotInputMapFromC(C.ImPlotContext_GetInputMap(self.handle()))
}

func (self ImPlotContext) SetOpenContextThisFrame(v bool) {
	C.ImPlotContext_SetOpenContextThisFrame(self.handle(), C.bool(v))
}

func (self ImPlotContext) GetOpenContextThisFrame() bool {
	return C.ImPlotContext_GetOpenContextThisFrame(self.handle()) == C.bool(true)
}

func (self ImPlotContext) GetMousePosStringBuilder() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotContext_GetMousePosStringBuilder(self.handle()))
}

func (self ImPlotContext) SetCurrentAlignmentH(v ImPlotAlignmentData) {
	C.ImPlotContext_SetCurrentAlignmentH(self.handle(), v.handle())
}

func (self ImPlotContext) GetCurrentAlignmentH() ImPlotAlignmentData {
	return (ImPlotAlignmentData)(unsafe.Pointer(C.ImPlotContext_GetCurrentAlignmentH(self.handle())))
}

func (self ImPlotContext) SetCurrentAlignmentV(v ImPlotAlignmentData) {
	C.ImPlotContext_SetCurrentAlignmentV(self.handle(), v.handle())
}

func (self ImPlotContext) GetCurrentAlignmentV() ImPlotAlignmentData {
	return (ImPlotAlignmentData)(unsafe.Pointer(C.ImPlotContext_GetCurrentAlignmentV(self.handle())))
}

func (self ImPlotDateTimeSpec) SetDate(v ImPlotDateFmt) {
	C.ImPlotDateTimeSpec_SetDate(self.handle(), C.ImPlotDateFmt(v))
}

func (self ImPlotDateTimeSpec) GetDate() ImPlotDateFmt {
	return ImPlotDateFmt(C.ImPlotDateTimeSpec_GetDate(self.handle()))
}

func (self ImPlotDateTimeSpec) SetTime(v ImPlotTimeFmt) {
	C.ImPlotDateTimeSpec_SetTime(self.handle(), C.ImPlotTimeFmt(v))
}

func (self ImPlotDateTimeSpec) GetTime() ImPlotTimeFmt {
	return ImPlotTimeFmt(C.ImPlotDateTimeSpec_GetTime(self.handle()))
}

func (self ImPlotDateTimeSpec) SetUseISO8601(v bool) {
	C.ImPlotDateTimeSpec_SetUseISO8601(self.handle(), C.bool(v))
}

func (self ImPlotDateTimeSpec) GetUseISO8601() bool {
	return C.ImPlotDateTimeSpec_GetUseISO8601(self.handle()) == C.bool(true)
}

func (self ImPlotDateTimeSpec) SetUse24HourClock(v bool) {
	C.ImPlotDateTimeSpec_SetUse24HourClock(self.handle(), C.bool(v))
}

func (self ImPlotDateTimeSpec) GetUse24HourClock() bool {
	return C.ImPlotDateTimeSpec_GetUse24HourClock(self.handle()) == C.bool(true)
}

func (self ImPlotInputMap) SetPan(v MouseButton) {
	C.ImPlotInputMap_SetPan(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetPan() MouseButton {
	return MouseButton(C.ImPlotInputMap_GetPan(self.handle()))
}

func (self ImPlotInputMap) SetPanMod(v ModFlags) {
	C.ImPlotInputMap_SetPanMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetPanMod() ModFlags {
	return ModFlags(C.ImPlotInputMap_GetPanMod(self.handle()))
}

func (self ImPlotInputMap) SetFit(v MouseButton) {
	C.ImPlotInputMap_SetFit(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetFit() MouseButton {
	return MouseButton(C.ImPlotInputMap_GetFit(self.handle()))
}

func (self ImPlotInputMap) SetSelect(v MouseButton) {
	C.ImPlotInputMap_SetSelect(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetSelect() MouseButton {
	return MouseButton(C.ImPlotInputMap_GetSelect(self.handle()))
}

func (self ImPlotInputMap) SetSelectCancel(v MouseButton) {
	C.ImPlotInputMap_SetSelectCancel(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetSelectCancel() MouseButton {
	return MouseButton(C.ImPlotInputMap_GetSelectCancel(self.handle()))
}

func (self ImPlotInputMap) SetSelectMod(v ModFlags) {
	C.ImPlotInputMap_SetSelectMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetSelectMod() ModFlags {
	return ModFlags(C.ImPlotInputMap_GetSelectMod(self.handle()))
}

func (self ImPlotInputMap) SetSelectHorzMod(v ModFlags) {
	C.ImPlotInputMap_SetSelectHorzMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetSelectHorzMod() ModFlags {
	return ModFlags(C.ImPlotInputMap_GetSelectHorzMod(self.handle()))
}

func (self ImPlotInputMap) SetSelectVertMod(v ModFlags) {
	C.ImPlotInputMap_SetSelectVertMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetSelectVertMod() ModFlags {
	return ModFlags(C.ImPlotInputMap_GetSelectVertMod(self.handle()))
}

func (self ImPlotInputMap) SetMenu(v MouseButton) {
	C.ImPlotInputMap_SetMenu(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetMenu() MouseButton {
	return MouseButton(C.ImPlotInputMap_GetMenu(self.handle()))
}

func (self ImPlotInputMap) SetOverrideMod(v ModFlags) {
	C.ImPlotInputMap_SetOverrideMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetOverrideMod() ModFlags {
	return ModFlags(C.ImPlotInputMap_GetOverrideMod(self.handle()))
}

func (self ImPlotInputMap) SetZoomMod(v ModFlags) {
	C.ImPlotInputMap_SetZoomMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetZoomMod() ModFlags {
	return ModFlags(C.ImPlotInputMap_GetZoomMod(self.handle()))
}

func (self ImPlotInputMap) SetZoomRate(v float32) {
	C.ImPlotInputMap_SetZoomRate(self.handle(), C.float(v))
}

func (self ImPlotInputMap) GetZoomRate() float32 {
	return float32(C.ImPlotInputMap_GetZoomRate(self.handle()))
}

func (self ImPlotItem) SetID(v ImGuiID) {
	C.ImPlotItem_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImPlotItem) GetID() ImGuiID {
	return ImGuiID(C.ImPlotItem_GetID(self.handle()))
}

func (self ImPlotItem) SetColor(v uint32) {
	C.ImPlotItem_SetColor(self.handle(), C.ImU32(v))
}

func (self ImPlotItem) GetColor() uint32 {
	return uint32(C.ImPlotItem_GetColor(self.handle()))
}

func (self ImPlotItem) SetLegendHoverRect(v ImRect) {
	C.ImPlotItem_SetLegendHoverRect(self.handle(), v.toC())
}

func (self ImPlotItem) GetLegendHoverRect() ImRect {
	return newImRectFromC(C.ImPlotItem_GetLegendHoverRect(self.handle()))
}

func (self ImPlotItem) SetNameOffset(v int32) {
	C.ImPlotItem_SetNameOffset(self.handle(), C.int(v))
}

func (self ImPlotItem) GetNameOffset() int {
	return int(C.ImPlotItem_GetNameOffset(self.handle()))
}

func (self ImPlotItem) SetShow(v bool) {
	C.ImPlotItem_SetShow(self.handle(), C.bool(v))
}

func (self ImPlotItem) GetShow() bool {
	return C.ImPlotItem_GetShow(self.handle()) == C.bool(true)
}

func (self ImPlotItem) SetLegendHovered(v bool) {
	C.ImPlotItem_SetLegendHovered(self.handle(), C.bool(v))
}

func (self ImPlotItem) GetLegendHovered() bool {
	return C.ImPlotItem_GetLegendHovered(self.handle()) == C.bool(true)
}

func (self ImPlotItem) SetSeenThisFrame(v bool) {
	C.ImPlotItem_SetSeenThisFrame(self.handle(), C.bool(v))
}

func (self ImPlotItem) GetSeenThisFrame() bool {
	return C.ImPlotItem_GetSeenThisFrame(self.handle()) == C.bool(true)
}

func (self ImPlotItemGroup) SetID(v ImGuiID) {
	C.ImPlotItemGroup_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImPlotItemGroup) GetID() ImGuiID {
	return ImGuiID(C.ImPlotItemGroup_GetID(self.handle()))
}

func (self ImPlotItemGroup) GetLegend() ImPlotLegend {
	return newImPlotLegendFromC(C.ImPlotItemGroup_GetLegend(self.handle()))
}

func (self ImPlotItemGroup) SetColormapIdx(v int32) {
	C.ImPlotItemGroup_SetColormapIdx(self.handle(), C.int(v))
}

func (self ImPlotItemGroup) GetColormapIdx() int {
	return int(C.ImPlotItemGroup_GetColormapIdx(self.handle()))
}

func (self ImPlotLegend) SetFlags(v ImPlotLegendFlags) {
	C.ImPlotLegend_SetFlags(self.handle(), C.ImPlotLegendFlags(v))
}

func (self ImPlotLegend) GetFlags() ImPlotLegendFlags {
	return ImPlotLegendFlags(C.ImPlotLegend_GetFlags(self.handle()))
}

func (self ImPlotLegend) SetPreviousFlags(v ImPlotLegendFlags) {
	C.ImPlotLegend_SetPreviousFlags(self.handle(), C.ImPlotLegendFlags(v))
}

func (self ImPlotLegend) GetPreviousFlags() ImPlotLegendFlags {
	return ImPlotLegendFlags(C.ImPlotLegend_GetPreviousFlags(self.handle()))
}

func (self ImPlotLegend) SetLocation(v ImPlotLocation) {
	C.ImPlotLegend_SetLocation(self.handle(), C.ImPlotLocation(v))
}

func (self ImPlotLegend) GetLocation() ImPlotLocation {
	return ImPlotLocation(C.ImPlotLegend_GetLocation(self.handle()))
}

func (self ImPlotLegend) SetPreviousLocation(v ImPlotLocation) {
	C.ImPlotLegend_SetPreviousLocation(self.handle(), C.ImPlotLocation(v))
}

func (self ImPlotLegend) GetPreviousLocation() ImPlotLocation {
	return ImPlotLocation(C.ImPlotLegend_GetPreviousLocation(self.handle()))
}

func (self ImPlotLegend) GetLabels() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotLegend_GetLabels(self.handle()))
}

func (self ImPlotLegend) SetRect(v ImRect) {
	C.ImPlotLegend_SetRect(self.handle(), v.toC())
}

func (self ImPlotLegend) GetRect() ImRect {
	return newImRectFromC(C.ImPlotLegend_GetRect(self.handle()))
}

func (self ImPlotLegend) SetHovered(v bool) {
	C.ImPlotLegend_SetHovered(self.handle(), C.bool(v))
}

func (self ImPlotLegend) GetHovered() bool {
	return C.ImPlotLegend_GetHovered(self.handle()) == C.bool(true)
}

func (self ImPlotLegend) SetHeld(v bool) {
	C.ImPlotLegend_SetHeld(self.handle(), C.bool(v))
}

func (self ImPlotLegend) GetHeld() bool {
	return C.ImPlotLegend_GetHeld(self.handle()) == C.bool(true)
}

func (self ImPlotLegend) SetCanGoInside(v bool) {
	C.ImPlotLegend_SetCanGoInside(self.handle(), C.bool(v))
}

func (self ImPlotLegend) GetCanGoInside() bool {
	return C.ImPlotLegend_GetCanGoInside(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetLineWeight(v float32) {
	C.ImPlotNextItemData_SetLineWeight(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetLineWeight() float32 {
	return float32(C.ImPlotNextItemData_GetLineWeight(self.handle()))
}

func (self ImPlotNextItemData) SetMarker(v ImPlotMarker) {
	C.ImPlotNextItemData_SetMarker(self.handle(), C.ImPlotMarker(v))
}

func (self ImPlotNextItemData) GetMarker() ImPlotMarker {
	return ImPlotMarker(C.ImPlotNextItemData_GetMarker(self.handle()))
}

func (self ImPlotNextItemData) SetMarkerSize(v float32) {
	C.ImPlotNextItemData_SetMarkerSize(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetMarkerSize() float32 {
	return float32(C.ImPlotNextItemData_GetMarkerSize(self.handle()))
}

func (self ImPlotNextItemData) SetMarkerWeight(v float32) {
	C.ImPlotNextItemData_SetMarkerWeight(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetMarkerWeight() float32 {
	return float32(C.ImPlotNextItemData_GetMarkerWeight(self.handle()))
}

func (self ImPlotNextItemData) SetFillAlpha(v float32) {
	C.ImPlotNextItemData_SetFillAlpha(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetFillAlpha() float32 {
	return float32(C.ImPlotNextItemData_GetFillAlpha(self.handle()))
}

func (self ImPlotNextItemData) SetErrorBarSize(v float32) {
	C.ImPlotNextItemData_SetErrorBarSize(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetErrorBarSize() float32 {
	return float32(C.ImPlotNextItemData_GetErrorBarSize(self.handle()))
}

func (self ImPlotNextItemData) SetErrorBarWeight(v float32) {
	C.ImPlotNextItemData_SetErrorBarWeight(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetErrorBarWeight() float32 {
	return float32(C.ImPlotNextItemData_GetErrorBarWeight(self.handle()))
}

func (self ImPlotNextItemData) SetDigitalBitHeight(v float32) {
	C.ImPlotNextItemData_SetDigitalBitHeight(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetDigitalBitHeight() float32 {
	return float32(C.ImPlotNextItemData_GetDigitalBitHeight(self.handle()))
}

func (self ImPlotNextItemData) SetDigitalBitGap(v float32) {
	C.ImPlotNextItemData_SetDigitalBitGap(self.handle(), C.float(v))
}

func (self ImPlotNextItemData) GetDigitalBitGap() float32 {
	return float32(C.ImPlotNextItemData_GetDigitalBitGap(self.handle()))
}

func (self ImPlotNextItemData) SetRenderLine(v bool) {
	C.ImPlotNextItemData_SetRenderLine(self.handle(), C.bool(v))
}

func (self ImPlotNextItemData) GetRenderLine() bool {
	return C.ImPlotNextItemData_GetRenderLine(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetRenderFill(v bool) {
	C.ImPlotNextItemData_SetRenderFill(self.handle(), C.bool(v))
}

func (self ImPlotNextItemData) GetRenderFill() bool {
	return C.ImPlotNextItemData_GetRenderFill(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetRenderMarkerLine(v bool) {
	C.ImPlotNextItemData_SetRenderMarkerLine(self.handle(), C.bool(v))
}

func (self ImPlotNextItemData) GetRenderMarkerLine() bool {
	return C.ImPlotNextItemData_GetRenderMarkerLine(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetRenderMarkerFill(v bool) {
	C.ImPlotNextItemData_SetRenderMarkerFill(self.handle(), C.bool(v))
}

func (self ImPlotNextItemData) GetRenderMarkerFill() bool {
	return C.ImPlotNextItemData_GetRenderMarkerFill(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetHasHidden(v bool) {
	C.ImPlotNextItemData_SetHasHidden(self.handle(), C.bool(v))
}

func (self ImPlotNextItemData) GetHasHidden() bool {
	return C.ImPlotNextItemData_GetHasHidden(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetHidden(v bool) {
	C.ImPlotNextItemData_SetHidden(self.handle(), C.bool(v))
}

func (self ImPlotNextItemData) GetHidden() bool {
	return C.ImPlotNextItemData_GetHidden(self.handle()) == C.bool(true)
}

func (self ImPlotNextItemData) SetHiddenCond(v ImPlotCond) {
	C.ImPlotNextItemData_SetHiddenCond(self.handle(), C.ImPlotCond(v))
}

func (self ImPlotNextItemData) GetHiddenCond() ImPlotCond {
	return ImPlotCond(C.ImPlotNextItemData_GetHiddenCond(self.handle()))
}

func (self ImPlotPlot) SetID(v ImGuiID) {
	C.ImPlotPlot_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImPlotPlot) GetID() ImGuiID {
	return ImGuiID(C.ImPlotPlot_GetID(self.handle()))
}

func (self ImPlotPlot) SetFlags(v ImPlotFlags) {
	C.ImPlotPlot_SetFlags(self.handle(), C.ImPlotFlags(v))
}

func (self ImPlotPlot) GetFlags() ImPlotFlags {
	return ImPlotFlags(C.ImPlotPlot_GetFlags(self.handle()))
}

func (self ImPlotPlot) SetPreviousFlags(v ImPlotFlags) {
	C.ImPlotPlot_SetPreviousFlags(self.handle(), C.ImPlotFlags(v))
}

func (self ImPlotPlot) GetPreviousFlags() ImPlotFlags {
	return ImPlotFlags(C.ImPlotPlot_GetPreviousFlags(self.handle()))
}

func (self ImPlotPlot) SetMouseTextLocation(v ImPlotLocation) {
	C.ImPlotPlot_SetMouseTextLocation(self.handle(), C.ImPlotLocation(v))
}

func (self ImPlotPlot) GetMouseTextLocation() ImPlotLocation {
	return ImPlotLocation(C.ImPlotPlot_GetMouseTextLocation(self.handle()))
}

func (self ImPlotPlot) SetMouseTextFlags(v ImPlotMouseTextFlags) {
	C.ImPlotPlot_SetMouseTextFlags(self.handle(), C.ImPlotMouseTextFlags(v))
}

func (self ImPlotPlot) GetMouseTextFlags() ImPlotMouseTextFlags {
	return ImPlotMouseTextFlags(C.ImPlotPlot_GetMouseTextFlags(self.handle()))
}

func (self ImPlotPlot) GetTextBuffer() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotPlot_GetTextBuffer(self.handle()))
}

func (self ImPlotPlot) GetItems() ImPlotItemGroup {
	return newImPlotItemGroupFromC(C.ImPlotPlot_GetItems(self.handle()))
}

func (self ImPlotPlot) SetCurrentX(v ImAxis) {
	C.ImPlotPlot_SetCurrentX(self.handle(), C.ImAxis(v))
}

func (self ImPlotPlot) GetCurrentX() ImAxis {
	return ImAxis(C.ImPlotPlot_GetCurrentX(self.handle()))
}

func (self ImPlotPlot) SetCurrentY(v ImAxis) {
	C.ImPlotPlot_SetCurrentY(self.handle(), C.ImAxis(v))
}

func (self ImPlotPlot) GetCurrentY() ImAxis {
	return ImAxis(C.ImPlotPlot_GetCurrentY(self.handle()))
}

func (self ImPlotPlot) SetFrameRect(v ImRect) {
	C.ImPlotPlot_SetFrameRect(self.handle(), v.toC())
}

func (self ImPlotPlot) GetFrameRect() ImRect {
	return newImRectFromC(C.ImPlotPlot_GetFrameRect(self.handle()))
}

func (self ImPlotPlot) SetCanvasRect(v ImRect) {
	C.ImPlotPlot_SetCanvasRect(self.handle(), v.toC())
}

func (self ImPlotPlot) GetCanvasRect() ImRect {
	return newImRectFromC(C.ImPlotPlot_GetCanvasRect(self.handle()))
}

func (self ImPlotPlot) SetPlotRect(v ImRect) {
	C.ImPlotPlot_SetPlotRect(self.handle(), v.toC())
}

func (self ImPlotPlot) GetPlotRect() ImRect {
	return newImRectFromC(C.ImPlotPlot_GetPlotRect(self.handle()))
}

func (self ImPlotPlot) SetAxesRect(v ImRect) {
	C.ImPlotPlot_SetAxesRect(self.handle(), v.toC())
}

func (self ImPlotPlot) GetAxesRect() ImRect {
	return newImRectFromC(C.ImPlotPlot_GetAxesRect(self.handle()))
}

func (self ImPlotPlot) SetSelectRect(v ImRect) {
	C.ImPlotPlot_SetSelectRect(self.handle(), v.toC())
}

func (self ImPlotPlot) GetSelectRect() ImRect {
	return newImRectFromC(C.ImPlotPlot_GetSelectRect(self.handle()))
}

func (self ImPlotPlot) SetSelectStart(v ImVec2) {
	C.ImPlotPlot_SetSelectStart(self.handle(), v.toC())
}

func (self ImPlotPlot) GetSelectStart() ImVec2 {
	return newImVec2FromC(C.ImPlotPlot_GetSelectStart(self.handle()))
}

func (self ImPlotPlot) SetTitleOffset(v int32) {
	C.ImPlotPlot_SetTitleOffset(self.handle(), C.int(v))
}

func (self ImPlotPlot) GetTitleOffset() int {
	return int(C.ImPlotPlot_GetTitleOffset(self.handle()))
}

func (self ImPlotPlot) SetJustCreated(v bool) {
	C.ImPlotPlot_SetJustCreated(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetJustCreated() bool {
	return C.ImPlotPlot_GetJustCreated(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetInitialized(v bool) {
	C.ImPlotPlot_SetInitialized(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetInitialized() bool {
	return C.ImPlotPlot_GetInitialized(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetSetupLocked(v bool) {
	C.ImPlotPlot_SetSetupLocked(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetSetupLocked() bool {
	return C.ImPlotPlot_GetSetupLocked(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetFitThisFrame(v bool) {
	C.ImPlotPlot_SetFitThisFrame(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetFitThisFrame() bool {
	return C.ImPlotPlot_GetFitThisFrame(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetHovered(v bool) {
	C.ImPlotPlot_SetHovered(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetHovered() bool {
	return C.ImPlotPlot_GetHovered(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetHeld(v bool) {
	C.ImPlotPlot_SetHeld(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetHeld() bool {
	return C.ImPlotPlot_GetHeld(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetSelecting(v bool) {
	C.ImPlotPlot_SetSelecting(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetSelecting() bool {
	return C.ImPlotPlot_GetSelecting(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetSelected(v bool) {
	C.ImPlotPlot_SetSelected(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetSelected() bool {
	return C.ImPlotPlot_GetSelected(self.handle()) == C.bool(true)
}

func (self ImPlotPlot) SetContextLocked(v bool) {
	C.ImPlotPlot_SetContextLocked(self.handle(), C.bool(v))
}

func (self ImPlotPlot) GetContextLocked() bool {
	return C.ImPlotPlot_GetContextLocked(self.handle()) == C.bool(true)
}

func (self ImPlotPointError) SetX(v float64) {
	C.ImPlotPointError_SetX(self.handle(), C.double(v))
}

func (self ImPlotPointError) GetX() float64 {
	return float64(C.ImPlotPointError_GetX(self.handle()))
}

func (self ImPlotPointError) SetY(v float64) {
	C.ImPlotPointError_SetY(self.handle(), C.double(v))
}

func (self ImPlotPointError) GetY() float64 {
	return float64(C.ImPlotPointError_GetY(self.handle()))
}

func (self ImPlotPointError) SetNeg(v float64) {
	C.ImPlotPointError_SetNeg(self.handle(), C.double(v))
}

func (self ImPlotPointError) GetNeg() float64 {
	return float64(C.ImPlotPointError_GetNeg(self.handle()))
}

func (self ImPlotPointError) SetPos(v float64) {
	C.ImPlotPointError_SetPos(self.handle(), C.double(v))
}

func (self ImPlotPointError) GetPos() float64 {
	return float64(C.ImPlotPointError_GetPos(self.handle()))
}

func (self ImPlotRange) SetMin(v float64) {
	C.ImPlotRange_SetMin(self.handle(), C.double(v))
}

func (self ImPlotRange) GetMin() float64 {
	return float64(C.ImPlotRange_GetMin(self.handle()))
}

func (self ImPlotRange) SetMax(v float64) {
	C.ImPlotRange_SetMax(self.handle(), C.double(v))
}

func (self ImPlotRange) GetMax() float64 {
	return float64(C.ImPlotRange_GetMax(self.handle()))
}

func (self ImPlotRect) GetX() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotRect_GetX(self.handle()))
}

func (self ImPlotRect) GetY() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotRect_GetY(self.handle()))
}

func (self ImPlotStyle) SetLineWeight(v float32) {
	C.ImPlotStyle_SetLineWeight(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetLineWeight() float32 {
	return float32(C.ImPlotStyle_GetLineWeight(self.handle()))
}

func (self ImPlotStyle) SetMarker(v int32) {
	C.ImPlotStyle_SetMarker(self.handle(), C.int(v))
}

func (self ImPlotStyle) GetMarker() int {
	return int(C.ImPlotStyle_GetMarker(self.handle()))
}

func (self ImPlotStyle) SetMarkerSize(v float32) {
	C.ImPlotStyle_SetMarkerSize(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetMarkerSize() float32 {
	return float32(C.ImPlotStyle_GetMarkerSize(self.handle()))
}

func (self ImPlotStyle) SetMarkerWeight(v float32) {
	C.ImPlotStyle_SetMarkerWeight(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetMarkerWeight() float32 {
	return float32(C.ImPlotStyle_GetMarkerWeight(self.handle()))
}

func (self ImPlotStyle) SetFillAlpha(v float32) {
	C.ImPlotStyle_SetFillAlpha(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetFillAlpha() float32 {
	return float32(C.ImPlotStyle_GetFillAlpha(self.handle()))
}

func (self ImPlotStyle) SetErrorBarSize(v float32) {
	C.ImPlotStyle_SetErrorBarSize(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetErrorBarSize() float32 {
	return float32(C.ImPlotStyle_GetErrorBarSize(self.handle()))
}

func (self ImPlotStyle) SetErrorBarWeight(v float32) {
	C.ImPlotStyle_SetErrorBarWeight(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetErrorBarWeight() float32 {
	return float32(C.ImPlotStyle_GetErrorBarWeight(self.handle()))
}

func (self ImPlotStyle) SetDigitalBitHeight(v float32) {
	C.ImPlotStyle_SetDigitalBitHeight(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetDigitalBitHeight() float32 {
	return float32(C.ImPlotStyle_GetDigitalBitHeight(self.handle()))
}

func (self ImPlotStyle) SetDigitalBitGap(v float32) {
	C.ImPlotStyle_SetDigitalBitGap(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetDigitalBitGap() float32 {
	return float32(C.ImPlotStyle_GetDigitalBitGap(self.handle()))
}

func (self ImPlotStyle) SetPlotBorderSize(v float32) {
	C.ImPlotStyle_SetPlotBorderSize(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetPlotBorderSize() float32 {
	return float32(C.ImPlotStyle_GetPlotBorderSize(self.handle()))
}

func (self ImPlotStyle) SetMinorAlpha(v float32) {
	C.ImPlotStyle_SetMinorAlpha(self.handle(), C.float(v))
}

func (self ImPlotStyle) GetMinorAlpha() float32 {
	return float32(C.ImPlotStyle_GetMinorAlpha(self.handle()))
}

func (self ImPlotStyle) SetMajorTickLen(v ImVec2) {
	C.ImPlotStyle_SetMajorTickLen(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMajorTickLen() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMajorTickLen(self.handle()))
}

func (self ImPlotStyle) SetMinorTickLen(v ImVec2) {
	C.ImPlotStyle_SetMinorTickLen(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMinorTickLen() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMinorTickLen(self.handle()))
}

func (self ImPlotStyle) SetMajorTickSize(v ImVec2) {
	C.ImPlotStyle_SetMajorTickSize(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMajorTickSize() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMajorTickSize(self.handle()))
}

func (self ImPlotStyle) SetMinorTickSize(v ImVec2) {
	C.ImPlotStyle_SetMinorTickSize(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMinorTickSize() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMinorTickSize(self.handle()))
}

func (self ImPlotStyle) SetMajorGridSize(v ImVec2) {
	C.ImPlotStyle_SetMajorGridSize(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMajorGridSize() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMajorGridSize(self.handle()))
}

func (self ImPlotStyle) SetMinorGridSize(v ImVec2) {
	C.ImPlotStyle_SetMinorGridSize(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMinorGridSize() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMinorGridSize(self.handle()))
}

func (self ImPlotStyle) SetPlotPadding(v ImVec2) {
	C.ImPlotStyle_SetPlotPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetPlotPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetPlotPadding(self.handle()))
}

func (self ImPlotStyle) SetLabelPadding(v ImVec2) {
	C.ImPlotStyle_SetLabelPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetLabelPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetLabelPadding(self.handle()))
}

func (self ImPlotStyle) SetLegendPadding(v ImVec2) {
	C.ImPlotStyle_SetLegendPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetLegendPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetLegendPadding(self.handle()))
}

func (self ImPlotStyle) SetLegendInnerPadding(v ImVec2) {
	C.ImPlotStyle_SetLegendInnerPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetLegendInnerPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetLegendInnerPadding(self.handle()))
}

func (self ImPlotStyle) SetLegendSpacing(v ImVec2) {
	C.ImPlotStyle_SetLegendSpacing(self.handle(), v.toC())
}

func (self ImPlotStyle) GetLegendSpacing() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetLegendSpacing(self.handle()))
}

func (self ImPlotStyle) SetMousePosPadding(v ImVec2) {
	C.ImPlotStyle_SetMousePosPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetMousePosPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetMousePosPadding(self.handle()))
}

func (self ImPlotStyle) SetAnnotationPadding(v ImVec2) {
	C.ImPlotStyle_SetAnnotationPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetAnnotationPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetAnnotationPadding(self.handle()))
}

func (self ImPlotStyle) SetFitPadding(v ImVec2) {
	C.ImPlotStyle_SetFitPadding(self.handle(), v.toC())
}

func (self ImPlotStyle) GetFitPadding() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetFitPadding(self.handle()))
}

func (self ImPlotStyle) SetPlotDefaultSize(v ImVec2) {
	C.ImPlotStyle_SetPlotDefaultSize(self.handle(), v.toC())
}

func (self ImPlotStyle) GetPlotDefaultSize() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetPlotDefaultSize(self.handle()))
}

func (self ImPlotStyle) SetPlotMinSize(v ImVec2) {
	C.ImPlotStyle_SetPlotMinSize(self.handle(), v.toC())
}

func (self ImPlotStyle) GetPlotMinSize() ImVec2 {
	return newImVec2FromC(C.ImPlotStyle_GetPlotMinSize(self.handle()))
}

func (self ImPlotStyle) SetColormap(v ImPlotColormap) {
	C.ImPlotStyle_SetColormap(self.handle(), C.ImPlotColormap(v))
}

func (self ImPlotStyle) GetColormap() ImPlotColormap {
	return ImPlotColormap(C.ImPlotStyle_GetColormap(self.handle()))
}

func (self ImPlotStyle) SetUseLocalTime(v bool) {
	C.ImPlotStyle_SetUseLocalTime(self.handle(), C.bool(v))
}

func (self ImPlotStyle) GetUseLocalTime() bool {
	return C.ImPlotStyle_GetUseLocalTime(self.handle()) == C.bool(true)
}

func (self ImPlotStyle) SetUseISO8601(v bool) {
	C.ImPlotStyle_SetUseISO8601(self.handle(), C.bool(v))
}

func (self ImPlotStyle) GetUseISO8601() bool {
	return C.ImPlotStyle_GetUseISO8601(self.handle()) == C.bool(true)
}

func (self ImPlotStyle) SetUse24HourClock(v bool) {
	C.ImPlotStyle_SetUse24HourClock(self.handle(), C.bool(v))
}

func (self ImPlotStyle) GetUse24HourClock() bool {
	return C.ImPlotStyle_GetUse24HourClock(self.handle()) == C.bool(true)
}

func (self ImPlotSubplot) SetID(v ImGuiID) {
	C.ImPlotSubplot_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImPlotSubplot) GetID() ImGuiID {
	return ImGuiID(C.ImPlotSubplot_GetID(self.handle()))
}

func (self ImPlotSubplot) SetFlags(v ImPlotSubplotFlags) {
	C.ImPlotSubplot_SetFlags(self.handle(), C.ImPlotSubplotFlags(v))
}

func (self ImPlotSubplot) GetFlags() ImPlotSubplotFlags {
	return ImPlotSubplotFlags(C.ImPlotSubplot_GetFlags(self.handle()))
}

func (self ImPlotSubplot) SetPreviousFlags(v ImPlotSubplotFlags) {
	C.ImPlotSubplot_SetPreviousFlags(self.handle(), C.ImPlotSubplotFlags(v))
}

func (self ImPlotSubplot) GetPreviousFlags() ImPlotSubplotFlags {
	return ImPlotSubplotFlags(C.ImPlotSubplot_GetPreviousFlags(self.handle()))
}

func (self ImPlotSubplot) GetItems() ImPlotItemGroup {
	return newImPlotItemGroupFromC(C.ImPlotSubplot_GetItems(self.handle()))
}

func (self ImPlotSubplot) SetRows(v int32) {
	C.ImPlotSubplot_SetRows(self.handle(), C.int(v))
}

func (self ImPlotSubplot) GetRows() int {
	return int(C.ImPlotSubplot_GetRows(self.handle()))
}

func (self ImPlotSubplot) SetCols(v int32) {
	C.ImPlotSubplot_SetCols(self.handle(), C.int(v))
}

func (self ImPlotSubplot) GetCols() int {
	return int(C.ImPlotSubplot_GetCols(self.handle()))
}

func (self ImPlotSubplot) SetCurrentIdx(v int32) {
	C.ImPlotSubplot_SetCurrentIdx(self.handle(), C.int(v))
}

func (self ImPlotSubplot) GetCurrentIdx() int {
	return int(C.ImPlotSubplot_GetCurrentIdx(self.handle()))
}

func (self ImPlotSubplot) SetFrameRect(v ImRect) {
	C.ImPlotSubplot_SetFrameRect(self.handle(), v.toC())
}

func (self ImPlotSubplot) GetFrameRect() ImRect {
	return newImRectFromC(C.ImPlotSubplot_GetFrameRect(self.handle()))
}

func (self ImPlotSubplot) SetGridRect(v ImRect) {
	C.ImPlotSubplot_SetGridRect(self.handle(), v.toC())
}

func (self ImPlotSubplot) GetGridRect() ImRect {
	return newImRectFromC(C.ImPlotSubplot_GetGridRect(self.handle()))
}

func (self ImPlotSubplot) SetCellSize(v ImVec2) {
	C.ImPlotSubplot_SetCellSize(self.handle(), v.toC())
}

func (self ImPlotSubplot) GetCellSize() ImVec2 {
	return newImVec2FromC(C.ImPlotSubplot_GetCellSize(self.handle()))
}

func (self ImPlotSubplot) SetFrameHovered(v bool) {
	C.ImPlotSubplot_SetFrameHovered(self.handle(), C.bool(v))
}

func (self ImPlotSubplot) GetFrameHovered() bool {
	return C.ImPlotSubplot_GetFrameHovered(self.handle()) == C.bool(true)
}

func (self ImPlotSubplot) SetHasTitle(v bool) {
	C.ImPlotSubplot_SetHasTitle(self.handle(), C.bool(v))
}

func (self ImPlotSubplot) GetHasTitle() bool {
	return C.ImPlotSubplot_GetHasTitle(self.handle()) == C.bool(true)
}

func (self ImPlotTag) SetAxis(v ImAxis) {
	C.ImPlotTag_SetAxis(self.handle(), C.ImAxis(v))
}

func (self ImPlotTag) GetAxis() ImAxis {
	return ImAxis(C.ImPlotTag_GetAxis(self.handle()))
}

func (self ImPlotTag) SetValue(v float64) {
	C.ImPlotTag_SetValue(self.handle(), C.double(v))
}

func (self ImPlotTag) GetValue() float64 {
	return float64(C.ImPlotTag_GetValue(self.handle()))
}

func (self ImPlotTag) SetColorBg(v uint32) {
	C.ImPlotTag_SetColorBg(self.handle(), C.ImU32(v))
}

func (self ImPlotTag) GetColorBg() uint32 {
	return uint32(C.ImPlotTag_GetColorBg(self.handle()))
}

func (self ImPlotTag) SetColorFg(v uint32) {
	C.ImPlotTag_SetColorFg(self.handle(), C.ImU32(v))
}

func (self ImPlotTag) GetColorFg() uint32 {
	return uint32(C.ImPlotTag_GetColorFg(self.handle()))
}

func (self ImPlotTag) SetTextOffset(v int32) {
	C.ImPlotTag_SetTextOffset(self.handle(), C.int(v))
}

func (self ImPlotTag) GetTextOffset() int {
	return int(C.ImPlotTag_GetTextOffset(self.handle()))
}

func (self ImPlotTagCollection) GetTextBuffer() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotTagCollection_GetTextBuffer(self.handle()))
}

func (self ImPlotTagCollection) SetSize(v int32) {
	C.ImPlotTagCollection_SetSize(self.handle(), C.int(v))
}

func (self ImPlotTagCollection) GetSize() int {
	return int(C.ImPlotTagCollection_GetSize(self.handle()))
}

func (self ImPlotTick) SetPlotPos(v float64) {
	C.ImPlotTick_SetPlotPos(self.handle(), C.double(v))
}

func (self ImPlotTick) GetPlotPos() float64 {
	return float64(C.ImPlotTick_GetPlotPos(self.handle()))
}

func (self ImPlotTick) SetPixelPos(v float32) {
	C.ImPlotTick_SetPixelPos(self.handle(), C.float(v))
}

func (self ImPlotTick) GetPixelPos() float32 {
	return float32(C.ImPlotTick_GetPixelPos(self.handle()))
}

func (self ImPlotTick) SetLabelSize(v ImVec2) {
	C.ImPlotTick_SetLabelSize(self.handle(), v.toC())
}

func (self ImPlotTick) GetLabelSize() ImVec2 {
	return newImVec2FromC(C.ImPlotTick_GetLabelSize(self.handle()))
}

func (self ImPlotTick) SetTextOffset(v int32) {
	C.ImPlotTick_SetTextOffset(self.handle(), C.int(v))
}

func (self ImPlotTick) GetTextOffset() int {
	return int(C.ImPlotTick_GetTextOffset(self.handle()))
}

func (self ImPlotTick) SetMajor(v bool) {
	C.ImPlotTick_SetMajor(self.handle(), C.bool(v))
}

func (self ImPlotTick) GetMajor() bool {
	return C.ImPlotTick_GetMajor(self.handle()) == C.bool(true)
}

func (self ImPlotTick) SetShowLabel(v bool) {
	C.ImPlotTick_SetShowLabel(self.handle(), C.bool(v))
}

func (self ImPlotTick) GetShowLabel() bool {
	return C.ImPlotTick_GetShowLabel(self.handle()) == C.bool(true)
}

func (self ImPlotTick) SetLevel(v int32) {
	C.ImPlotTick_SetLevel(self.handle(), C.int(v))
}

func (self ImPlotTick) GetLevel() int {
	return int(C.ImPlotTick_GetLevel(self.handle()))
}

func (self ImPlotTick) SetIdx(v int32) {
	C.ImPlotTick_SetIdx(self.handle(), C.int(v))
}

func (self ImPlotTick) GetIdx() int {
	return int(C.ImPlotTick_GetIdx(self.handle()))
}

func (self ImPlotTicker) GetTextBuffer() ImGuiTextBuffer {
	return newImGuiTextBufferFromC(C.ImPlotTicker_GetTextBuffer(self.handle()))
}

func (self ImPlotTicker) SetMaxSize(v ImVec2) {
	C.ImPlotTicker_SetMaxSize(self.handle(), v.toC())
}

func (self ImPlotTicker) GetMaxSize() ImVec2 {
	return newImVec2FromC(C.ImPlotTicker_GetMaxSize(self.handle()))
}

func (self ImPlotTicker) SetLateSize(v ImVec2) {
	C.ImPlotTicker_SetLateSize(self.handle(), v.toC())
}

func (self ImPlotTicker) GetLateSize() ImVec2 {
	return newImVec2FromC(C.ImPlotTicker_GetLateSize(self.handle()))
}

func (self ImPlotTicker) SetLevels(v int32) {
	C.ImPlotTicker_SetLevels(self.handle(), C.int(v))
}

func (self ImPlotTicker) GetLevels() int {
	return int(C.ImPlotTicker_GetLevels(self.handle()))
}

func (self ImPlotTime) SetUs(v int32) {
	C.ImPlotTime_SetUs(self.handle(), C.int(v))
}

func (self ImPlotTime) GetUs() int {
	return int(C.ImPlotTime_GetUs(self.handle()))
}
