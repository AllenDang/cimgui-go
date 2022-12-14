package cimgui

// #include "extra_types.h"
// #include "cimplot_structs_accessor.h"
// #include "cimplot_wrapper.h"
import "C"
import "unsafe"

func (self ImPlotAlignmentData) Destroy() {
	C.PlotAlignmentData_Destroy(self.handle())
}

func (self ImPlotAnnotationCollection) Destroy() {
	C.PlotAnnotationCollection_Destroy(self.handle())
}

func (self ImPlotAxis) Destroy() {
	C.PlotAxis_Destroy(self.handle())
}

func (self ImPlotColormapData) Destroy() {
	C.PlotColormapData_Destroy(self.handle())
}

func (self ImPlotDateTimeSpec) Destroy() {
	C.PlotDateTimeSpec_Destroy(self.handle())
}

func NewImPlotInputMap() ImPlotInputMap {
	return (ImPlotInputMap)(unsafe.Pointer(C.PlotInputMap_ImPlotInputMap()))
}

func (self ImPlotInputMap) Destroy() {
	C.PlotInputMap_Destroy(self.handle())
}

func (self ImPlotItemGroup) Destroy() {
	C.PlotItemGroup_Destroy(self.handle())
}

func (self ImPlotLegend) Destroy() {
	C.PlotLegend_Destroy(self.handle())
}

func (self ImPlotNextItemData) Destroy() {
	C.PlotNextItemData_Destroy(self.handle())
}

func (self ImPlotNextPlotData) Destroy() {
	C.PlotNextPlotData_Destroy(self.handle())
}

func (self ImPlotPlot) Destroy() {
	C.PlotPlot_Destroy(self.handle())
}

func (self ImPlotPointError) Destroy() {
	C.PlotPointError_Destroy(self.handle())
}

func (self *ImPlotPoint) Destroy() {
	selfArg, selfFin := self.wrap()
	defer selfFin()

	C.PlotPoint_Destroy(selfArg)
}

func (self ImPlotRange) Clamp(value float64) float64 {
	return float64(C.PlotRange_Clamp(self.handle(), C.double(value)))
}

func (self ImPlotRange) Contains(value float64) bool {
	return C.PlotRange_Contains(self.handle(), C.double(value)) == C.bool(true)
}

func NewImPlotRangeNil() ImPlotRange {
	return (ImPlotRange)(unsafe.Pointer(C.PlotRange_ImPlotRange_Nil()))
}

func NewImPlotRangedouble(_min float64, _max float64) ImPlotRange {
	return (ImPlotRange)(unsafe.Pointer(C.PlotRange_ImPlotRange_double(C.double(_min), C.double(_max))))
}

func (self ImPlotRange) Size() float64 {
	return float64(C.PlotRange_Size(self.handle()))
}

func (self ImPlotRange) Destroy() {
	C.PlotRange_Destroy(self.handle())
}

func (self ImPlotRect) Contains_double(x float64, y float64) bool {
	return C.PlotRect_Contains_double(self.handle(), C.double(x), C.double(y)) == C.bool(true)
}

func NewImPlotRectNil() ImPlotRect {
	return (ImPlotRect)(unsafe.Pointer(C.PlotRect_ImPlotRect_Nil()))
}

func NewImPlotRectdouble(x_min float64, x_max float64, y_min float64, y_max float64) ImPlotRect {
	return (ImPlotRect)(unsafe.Pointer(C.PlotRect_ImPlotRect_double(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max))))
}

func (self ImPlotRect) Destroy() {
	C.PlotRect_Destroy(self.handle())
}

func NewImPlotStyle() ImPlotStyle {
	return (ImPlotStyle)(unsafe.Pointer(C.PlotStyle_ImPlotStyle()))
}

func (self ImPlotStyle) Destroy() {
	C.PlotStyle_Destroy(self.handle())
}

func (self ImPlotSubplot) Destroy() {
	C.PlotSubplot_Destroy(self.handle())
}

func (self ImPlotTagCollection) Destroy() {
	C.PlotTagCollection_Destroy(self.handle())
}

func (self ImPlotTick) Destroy() {
	C.PlotTick_Destroy(self.handle())
}

func (self ImPlotTicker) Destroy() {
	C.PlotTicker_Destroy(self.handle())
}

func (self ImPlotTime) Destroy() {
	C.PlotTime_Destroy(self.handle())
}

// Plot_AddColormap_Vec4PtrV parameter default value hint:
// qual: true
func Plot_AddColormap_Vec4PtrV(name string, cols *ImVec4, size int32, qual bool) PlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	colsArg, colsFin := cols.wrap()
	defer colsFin()

	return PlotColormap(C.Plot_AddColormap_Vec4PtrV(nameArg, colsArg, C.int(size), C.bool(qual)))
}

// Plot_Annotation_BoolV parameter default value hint:
// round: false
func Plot_Annotation_BoolV(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool, round bool) {
	C.Plot_Annotation_BoolV(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp), C.bool(round))
}

func Plot_Annotation_Str(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_Annotation_Str(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp), fmtArg)
}

// Plot_BeginAlignedPlotsV parameter default value hint:
// vertical: true
func Plot_BeginAlignedPlotsV(group_id string, vertical bool) bool {
	group_idArg, group_idFin := wrapString(group_id)
	defer group_idFin()

	return C.Plot_BeginAlignedPlotsV(group_idArg, C.bool(vertical)) == C.bool(true)
}

// Plot_BeginDragDropSourceAxisV parameter default value hint:
// flags: 0
func Plot_BeginDragDropSourceAxisV(axis ImAxis, flags DragDropFlags) bool {
	return C.Plot_BeginDragDropSourceAxisV(C.ImAxis(axis), C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

// Plot_BeginDragDropSourceItemV parameter default value hint:
// flags: 0
func Plot_BeginDragDropSourceItemV(label_id string, flags DragDropFlags) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_BeginDragDropSourceItemV(label_idArg, C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

// Plot_BeginDragDropSourcePlotV parameter default value hint:
// flags: 0
func Plot_BeginDragDropSourcePlotV(flags DragDropFlags) bool {
	return C.Plot_BeginDragDropSourcePlotV(C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

func Plot_BeginDragDropTargetAxis(axis ImAxis) bool {
	return C.Plot_BeginDragDropTargetAxis(C.ImAxis(axis)) == C.bool(true)
}

func Plot_BeginDragDropTargetLegend() bool {
	return C.Plot_BeginDragDropTargetLegend() == C.bool(true)
}

func Plot_BeginDragDropTargetPlot() bool {
	return C.Plot_BeginDragDropTargetPlot() == C.bool(true)
}

// Plot_BeginLegendPopupV parameter default value hint:
// mouse_button: 1
func Plot_BeginLegendPopupV(label_id string, mouse_button MouseButton) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_BeginLegendPopupV(label_idArg, C.ImGuiMouseButton(mouse_button)) == C.bool(true)
}

// Plot_BeginPlotV parameter default value hint:
// flags: 0
// size: ImVec2(-1,0)
func Plot_BeginPlotV(title_id string, size ImVec2, flags PlotFlags) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.Plot_BeginPlotV(title_idArg, size.toC(), C.ImPlotFlags(flags)) == C.bool(true)
}

// Plot_BeginSubplotsV parameter default value hint:
// col_ratios: ((void*)0)
// flags: 0
// row_ratios: ((void*)0)
func Plot_BeginSubplotsV(title_id string, rows int32, cols int32, size ImVec2, flags PlotSubplotFlags, row_ratios *float32, col_ratios *float32) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	row_ratiosArg, row_ratiosFin := wrapNumberPtr[C.float, float32](row_ratios)
	defer row_ratiosFin()

	col_ratiosArg, col_ratiosFin := wrapNumberPtr[C.float, float32](col_ratios)
	defer col_ratiosFin()

	return C.Plot_BeginSubplotsV(title_idArg, C.int(rows), C.int(cols), size.toC(), C.ImPlotSubplotFlags(flags), row_ratiosArg, col_ratiosArg) == C.bool(true)
}

// Plot_BustColorCacheV parameter default value hint:
// plot_title_id: ((void*)0)
func Plot_BustColorCacheV(plot_title_id string) {
	plot_title_idArg, plot_title_idFin := wrapString(plot_title_id)
	defer plot_title_idFin()

	C.Plot_BustColorCacheV(plot_title_idArg)
}

func Plot_CancelPlotSelection() {
	C.Plot_CancelPlotSelection()
}

// Plot_ColormapButtonV parameter default value hint:
// cmap: -1
// size: ImVec2(0,0)
func Plot_ColormapButtonV(label string, size ImVec2, cmap PlotColormap) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ColormapButtonV(labelArg, size.toC(), C.ImPlotColormap(cmap)) == C.bool(true)
}

func Plot_ColormapIcon(cmap PlotColormap) {
	C.Plot_ColormapIcon(C.ImPlotColormap(cmap))
}

// Plot_ColormapScaleV parameter default value hint:
// cmap: -1
// flags: 0
// format: "%g"
// size: ImVec2(0,0)
func Plot_ColormapScaleV(label string, scale_min float64, scale_max float64, size ImVec2, format string, flags PlotColormapScaleFlags, cmap PlotColormap) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	C.Plot_ColormapScaleV(labelArg, C.double(scale_min), C.double(scale_max), size.toC(), formatArg, C.ImPlotColormapScaleFlags(flags), C.ImPlotColormap(cmap))
}

// Plot_ColormapSliderV parameter default value hint:
// cmap: -1
// format: ""
// out: ((void*)0)
func Plot_ColormapSliderV(label string, t *float32, out *ImVec4, format string, cmap PlotColormap) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	tArg, tFin := wrapNumberPtr[C.float, float32](t)
	defer tFin()

	outArg, outFin := out.wrap()
	defer outFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.Plot_ColormapSliderV(labelArg, tArg, outArg, formatArg, C.ImPlotColormap(cmap)) == C.bool(true)
}

func Plot_CreateContext() ImPlotContext {
	return (ImPlotContext)(unsafe.Pointer(C.Plot_CreateContext()))
}

// Plot_DestroyContextV parameter default value hint:
// ctx: ((void*)0)
func Plot_DestroyContextV(ctx ImPlotContext) {
	C.Plot_DestroyContextV(ctx.handle())
}

// Plot_DragLineXV parameter default value hint:
// flags: 0
// thickness: 1
func Plot_DragLineXV(id int32, x *float64, col ImVec4, thickness float32, flags PlotDragToolFlags) bool {
	return C.Plot_DragLineXV(C.int(id), (*C.double)(x), col.toC(), C.float(thickness), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// Plot_DragLineYV parameter default value hint:
// flags: 0
// thickness: 1
func Plot_DragLineYV(id int32, y *float64, col ImVec4, thickness float32, flags PlotDragToolFlags) bool {
	return C.Plot_DragLineYV(C.int(id), (*C.double)(y), col.toC(), C.float(thickness), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// Plot_DragPointV parameter default value hint:
// flags: 0
// size: 4
func Plot_DragPointV(id int32, x *float64, y *float64, col ImVec4, size float32, flags PlotDragToolFlags) bool {
	return C.Plot_DragPointV(C.int(id), (*C.double)(x), (*C.double)(y), col.toC(), C.float(size), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// Plot_DragRectV parameter default value hint:
// flags: 0
func Plot_DragRectV(id int32, x_min *float64, y_min *float64, x_max *float64, y_max *float64, col ImVec4, flags PlotDragToolFlags) bool {
	return C.Plot_DragRectV(C.int(id), (*C.double)(x_min), (*C.double)(y_min), (*C.double)(x_max), (*C.double)(y_max), col.toC(), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

func Plot_EndAlignedPlots() {
	C.Plot_EndAlignedPlots()
}

func Plot_EndDragDropSource() {
	C.Plot_EndDragDropSource()
}

func Plot_EndDragDropTarget() {
	C.Plot_EndDragDropTarget()
}

func Plot_EndLegendPopup() {
	C.Plot_EndLegendPopup()
}

func Plot_EndPlot() {
	C.Plot_EndPlot()
}

func Plot_EndSubplots() {
	C.Plot_EndSubplots()
}

// Plot_GetColormapColorV parameter default value hint:
// cmap: -1
func Plot_GetColormapColorV(idx int32, cmap PlotColormap) ImVec4 {
	pOut := ImVec4{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetColormapColorV(pOutArg, C.int(idx), C.ImPlotColormap(cmap))
	return pOut
}

func Plot_GetColormapCount() int {
	return int(C.Plot_GetColormapCount())
}

func Plot_GetColormapIndex(name string) PlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	return PlotColormap(C.Plot_GetColormapIndex(nameArg))
}

func Plot_GetColormapName(cmap PlotColormap) string {
	return C.GoString(C.Plot_GetColormapName(C.ImPlotColormap(cmap)))
}

// Plot_GetColormapSizeV parameter default value hint:
// cmap: -1
func Plot_GetColormapSizeV(cmap PlotColormap) int {
	return int(C.Plot_GetColormapSizeV(C.ImPlotColormap(cmap)))
}

func Plot_GetCurrentContext() ImPlotContext {
	return (ImPlotContext)(unsafe.Pointer(C.Plot_GetCurrentContext()))
}

func Plot_GetInputMap() ImPlotInputMap {
	return (ImPlotInputMap)(unsafe.Pointer(C.Plot_GetInputMap()))
}

func Plot_GetLastItemColor() ImVec4 {
	pOut := ImVec4{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetLastItemColor(pOutArg)
	return pOut
}

func Plot_GetMarkerName(idx PlotMarker) string {
	return C.GoString(C.Plot_GetMarkerName(C.ImPlotMarker(idx)))
}

func Plot_GetPlotDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.Plot_GetPlotDrawList()))
}

func Plot_GetPlotPos() ImVec2 {
	pOut := ImVec2{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetPlotPos(pOutArg)
	return pOut
}

func Plot_GetPlotSize() ImVec2 {
	pOut := ImVec2{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetPlotSize(pOutArg)
	return pOut
}

func Plot_GetStyle() ImPlotStyle {
	return (ImPlotStyle)(unsafe.Pointer(C.Plot_GetStyle()))
}

func Plot_GetStyleColorName(idx PlotCol) string {
	return C.GoString(C.Plot_GetStyleColorName(C.ImPlotCol(idx)))
}

// Plot_HideNextItemV parameter default value hint:
// cond: ImPlotCond_Once
// hidden: true
func Plot_HideNextItemV(hidden bool, cond PlotCond) {
	C.Plot_HideNextItemV(C.bool(hidden), C.ImPlotCond(cond))
}

func Plot_IsAxisHovered(axis ImAxis) bool {
	return C.Plot_IsAxisHovered(C.ImAxis(axis)) == C.bool(true)
}

func Plot_IsLegendEntryHovered(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_IsLegendEntryHovered(label_idArg) == C.bool(true)
}

func Plot_IsPlotHovered() bool {
	return C.Plot_IsPlotHovered() == C.bool(true)
}

func Plot_IsPlotSelected() bool {
	return C.Plot_IsPlotSelected() == C.bool(true)
}

func Plot_IsSubplotsHovered() bool {
	return C.Plot_IsSubplotsHovered() == C.bool(true)
}

func Plot_ItemIcon_U32(col uint32) {
	C.Plot_ItemIcon_U32(C.ImU32(col))
}

func Plot_ItemIcon_Vec4(col ImVec4) {
	C.Plot_ItemIcon_Vec4(col.toC())
}

// Plot_MapInputDefaultV parameter default value hint:
// dst: ((void*)0)
func Plot_MapInputDefaultV(dst ImPlotInputMap) {
	C.Plot_MapInputDefaultV(dst.handle())
}

// Plot_MapInputReverseV parameter default value hint:
// dst: ((void*)0)
func Plot_MapInputReverseV(dst ImPlotInputMap) {
	C.Plot_MapInputReverseV(dst.handle())
}

func Plot_NextColormapColor() ImVec4 {
	pOut := ImVec4{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_NextColormapColor(pOutArg)
	return pOut
}

// Plot_PlotBars_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotBars_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, bar_size float64, flags PlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_FloatPtrIntV parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(float)
func Plot_PlotBars_FloatPtrIntV(label_id string, values []float32, count int32, bar_size float64, shift float64, flags PlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_S64PtrIntV parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(ImS64)
func Plot_PlotBars_S64PtrIntV(label_id string, values []int64, count int32, bar_size float64, shift float64, flags PlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotBars_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, bar_size float64, flags PlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_U64PtrIntV parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(ImU64)
func Plot_PlotBars_U64PtrIntV(label_id string, values []uint64, count int32, bar_size float64, shift float64, flags PlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotBars_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, bar_size float64, flags PlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotDigital_FloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotDigital_FloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags PlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_FloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotDigital_S64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotDigital_S64PtrV(label_id string, xs []int64, ys []int64, count int32, flags PlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_S64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotDigital_U64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotDigital_U64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags PlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_U64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotDummyV parameter default value hint:
// flags: 0
func Plot_PlotDummyV(label_id string, flags PlotDummyFlags) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDummyV(label_idArg, C.ImPlotDummyFlags(flags))
}

// Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, neg []float32, pos []float32, count int32, flags PlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(neg[0])), (*C.float)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV(label_id string, xs []float32, ys []float32, err []float32, count int32, flags PlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrIntV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_S64PtrS64PtrS64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotErrorBars_S64PtrS64PtrS64PtrIntV(label_id string, xs []int64, ys []int64, err []int64, count int32, flags PlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_S64PtrS64PtrS64PtrIntV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV(label_id string, xs []int64, ys []int64, neg []int64, pos []int64, count int32, flags PlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(neg[0])), (*C.longlong)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_U64PtrU64PtrU64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotErrorBars_U64PtrU64PtrU64PtrIntV(label_id string, xs []uint64, ys []uint64, err []uint64, count int32, flags PlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_U64PtrU64PtrU64PtrIntV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, neg []uint64, pos []uint64, count int32, flags PlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(neg[0])), (*C.ulonglong)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotInfLines_FloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotInfLines_FloatPtrV(label_id string, values []float32, count int32, flags PlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_FloatPtrV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotInfLines_S64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotInfLines_S64PtrV(label_id string, values []int64, count int32, flags PlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_S64PtrV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotInfLines_U64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotInfLines_U64PtrV(label_id string, values []uint64, count int32, flags PlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_U64PtrV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotLine_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags PlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func Plot_PlotLine_FloatPtrIntV(label_id string, values []float32, count int32, xscale float64, xstart float64, flags PlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func Plot_PlotLine_S64PtrIntV(label_id string, values []int64, count int32, xscale float64, xstart float64, flags PlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotLine_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, flags PlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func Plot_PlotLine_U64PtrIntV(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags PlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotLine_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags PlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotScatter_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags PlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func Plot_PlotScatter_FloatPtrIntV(label_id string, values []float32, count int32, xscale float64, xstart float64, flags PlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func Plot_PlotScatter_S64PtrIntV(label_id string, values []int64, count int32, xscale float64, xstart float64, flags PlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotScatter_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, flags PlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func Plot_PlotScatter_U64PtrIntV(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags PlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotScatter_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags PlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_FloatPtrFloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotShaded_FloatPtrFloatPtrFloatPtrV(label_id string, xs []float32, ys1 []float32, ys2 []float32, count int32, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_FloatPtrFloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys1[0])), (*C.float)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_FloatPtrFloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// yref: 0
func Plot_PlotShaded_FloatPtrFloatPtrIntV(label_id string, xs []float32, ys []float32, count int32, yref float64, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_FloatPtrFloatPtrIntV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
// yref: 0
func Plot_PlotShaded_FloatPtrIntV(label_id string, values []float32, count int32, yref float64, xscale float64, xstart float64, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
// yref: 0
func Plot_PlotShaded_S64PtrIntV(label_id string, values []int64, count int32, yref float64, xscale float64, xstart float64, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_S64PtrS64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// yref: 0
func Plot_PlotShaded_S64PtrS64PtrIntV(label_id string, xs []int64, ys []int64, count int32, yref float64, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrS64PtrIntV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_S64PtrS64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotShaded_S64PtrS64PtrS64PtrV(label_id string, xs []int64, ys1 []int64, ys2 []int64, count int32, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrS64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys1[0])), (*C.longlong)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
// yref: 0
func Plot_PlotShaded_U64PtrIntV(label_id string, values []uint64, count int32, yref float64, xscale float64, xstart float64, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_U64PtrU64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// yref: 0
func Plot_PlotShaded_U64PtrU64PtrIntV(label_id string, xs []uint64, ys []uint64, count int32, yref float64, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrU64PtrIntV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_U64PtrU64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotShaded_U64PtrU64PtrU64PtrV(label_id string, xs []uint64, ys1 []uint64, ys2 []uint64, count int32, flags PlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrU64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys1[0])), (*C.ulonglong)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotStairs_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, flags PlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func Plot_PlotStairs_FloatPtrIntV(label_id string, values []float32, count int32, xscale float64, xstart float64, flags PlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func Plot_PlotStairs_S64PtrIntV(label_id string, values []int64, count int32, xscale float64, xstart float64, flags PlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotStairs_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, flags PlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func Plot_PlotStairs_U64PtrIntV(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags PlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotStairs_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, flags PlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_FloatPtrFloatPtrV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(float)
func Plot_PlotStems_FloatPtrFloatPtrV(label_id string, xs []float32, ys []float32, count int32, ref float64, flags PlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_FloatPtrFloatPtrV(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_FloatPtrIntV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(float)
func Plot_PlotStems_FloatPtrIntV(label_id string, values []float32, count int32, ref float64, scale float64, start float64, flags PlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_FloatPtrIntV(label_idArg, (*C.float)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_S64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(ImS64)
func Plot_PlotStems_S64PtrIntV(label_id string, values []int64, count int32, ref float64, scale float64, start float64, flags PlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_S64PtrIntV(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_S64PtrS64PtrV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(ImS64)
func Plot_PlotStems_S64PtrS64PtrV(label_id string, xs []int64, ys []int64, count int32, ref float64, flags PlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_S64PtrS64PtrV(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_U64PtrIntV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(ImU64)
func Plot_PlotStems_U64PtrIntV(label_id string, values []uint64, count int32, ref float64, scale float64, start float64, flags PlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_U64PtrIntV(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_U64PtrU64PtrV parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(ImU64)
func Plot_PlotStems_U64PtrU64PtrV(label_id string, xs []uint64, ys []uint64, count int32, ref float64, flags PlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_U64PtrU64PtrV(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotTextV parameter default value hint:
// flags: 0
// pix_offset: ImVec2(0,0)
func Plot_PlotTextV(text string, x float64, y float64, pix_offset ImVec2, flags PlotTextFlags) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.Plot_PlotTextV(textArg, C.double(x), C.double(y), pix_offset.toC(), C.ImPlotTextFlags(flags))
}

// Plot_PlotToPixels_doubleV parameter default value hint:
// x_axis: -1
// y_axis: -1
func Plot_PlotToPixels_doubleV(x float64, y float64, x_axis ImAxis, y_axis ImAxis) ImVec2 {
	pOut := ImVec2{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_PlotToPixels_doubleV(pOutArg, C.double(x), C.double(y), C.ImAxis(x_axis), C.ImAxis(y_axis))
	return pOut
}

// Plot_PopColormapV parameter default value hint:
// count: 1
func Plot_PopColormapV(count int32) {
	C.Plot_PopColormapV(C.int(count))
}

func Plot_PopPlotClipRect() {
	C.Plot_PopPlotClipRect()
}

// Plot_PopStyleColorV parameter default value hint:
// count: 1
func Plot_PopStyleColorV(count int32) {
	C.Plot_PopStyleColorV(C.int(count))
}

// Plot_PopStyleVarV parameter default value hint:
// count: 1
func Plot_PopStyleVarV(count int32) {
	C.Plot_PopStyleVarV(C.int(count))
}

func Plot_PushColormap_PlotColormap(cmap PlotColormap) {
	C.Plot_PushColormap_PlotColormap(C.ImPlotColormap(cmap))
}

func Plot_PushColormap_Str(name string) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.Plot_PushColormap_Str(nameArg)
}

// Plot_PushPlotClipRectV parameter default value hint:
// expand: 0
func Plot_PushPlotClipRectV(expand float32) {
	C.Plot_PushPlotClipRectV(C.float(expand))
}

func Plot_PushStyleColor_U32(idx PlotCol, col uint32) {
	C.Plot_PushStyleColor_U32(C.ImPlotCol(idx), C.ImU32(col))
}

func Plot_PushStyleColor_Vec4(idx PlotCol, col ImVec4) {
	C.Plot_PushStyleColor_Vec4(C.ImPlotCol(idx), col.toC())
}

func Plot_PushStyleVar_Float(idx PlotStyleVar, val float32) {
	C.Plot_PushStyleVar_Float(C.ImPlotStyleVar(idx), C.float(val))
}

func Plot_PushStyleVar_Int(idx PlotStyleVar, val int32) {
	C.Plot_PushStyleVar_Int(C.ImPlotStyleVar(idx), C.int(val))
}

func Plot_PushStyleVar_Vec2(idx PlotStyleVar, val ImVec2) {
	C.Plot_PushStyleVar_Vec2(C.ImPlotStyleVar(idx), val.toC())
}

// Plot_SampleColormapV parameter default value hint:
// cmap: -1
func Plot_SampleColormapV(t float32, cmap PlotColormap) ImVec4 {
	pOut := ImVec4{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_SampleColormapV(pOutArg, C.float(t), C.ImPlotColormap(cmap))
	return pOut
}

func Plot_SetAxes(x_axis ImAxis, y_axis ImAxis) {
	C.Plot_SetAxes(C.ImAxis(x_axis), C.ImAxis(y_axis))
}

func Plot_SetAxis(axis ImAxis) {
	C.Plot_SetAxis(C.ImAxis(axis))
}

func Plot_SetCurrentContext(ctx ImPlotContext) {
	C.Plot_SetCurrentContext(ctx.handle())
}

func Plot_SetImGuiContext(ctx ImGuiContext) {
	C.Plot_SetImGuiContext(ctx.handle())
}

// Plot_SetNextAxesLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetNextAxesLimitsV(x_min float64, x_max float64, y_min float64, y_max float64, cond PlotCond) {
	C.Plot_SetNextAxesLimitsV(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max), C.ImPlotCond(cond))
}

func Plot_SetNextAxesToFit() {
	C.Plot_SetNextAxesToFit()
}

// Plot_SetNextAxisLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetNextAxisLimitsV(axis ImAxis, v_min float64, v_max float64, cond PlotCond) {
	C.Plot_SetNextAxisLimitsV(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.ImPlotCond(cond))
}

func Plot_SetNextAxisLinks(axis ImAxis, link_min *float64, link_max *float64) {
	C.Plot_SetNextAxisLinks(C.ImAxis(axis), (*C.double)(link_min), (*C.double)(link_max))
}

func Plot_SetNextAxisToFit(axis ImAxis) {
	C.Plot_SetNextAxisToFit(C.ImAxis(axis))
}

// Plot_SetNextErrorBarStyleV parameter default value hint:
// col: ImVec4(0,0,0,-1)
// size: -1
// weight: -1
func Plot_SetNextErrorBarStyleV(col ImVec4, size float32, weight float32) {
	C.Plot_SetNextErrorBarStyleV(col.toC(), C.float(size), C.float(weight))
}

// Plot_SetNextFillStyleV parameter default value hint:
// alpha_mod: -1
// col: ImVec4(0,0,0,-1)
func Plot_SetNextFillStyleV(col ImVec4, alpha_mod float32) {
	C.Plot_SetNextFillStyleV(col.toC(), C.float(alpha_mod))
}

// Plot_SetNextLineStyleV parameter default value hint:
// col: ImVec4(0,0,0,-1)
// weight: -1
func Plot_SetNextLineStyleV(col ImVec4, weight float32) {
	C.Plot_SetNextLineStyleV(col.toC(), C.float(weight))
}

// Plot_SetNextMarkerStyleV parameter default value hint:
// fill: ImVec4(0,0,0,-1)
// marker: -1
// outline: ImVec4(0,0,0,-1)
// size: -1
// weight: -1
func Plot_SetNextMarkerStyleV(marker PlotMarker, size float32, fill ImVec4, weight float32, outline ImVec4) {
	C.Plot_SetNextMarkerStyleV(C.ImPlotMarker(marker), C.float(size), fill.toC(), C.float(weight), outline.toC())
}

// Plot_SetupAxesV parameter default value hint:
// x_flags: 0
// y_flags: 0
func Plot_SetupAxesV(x_label string, y_label string, x_flags PlotAxisFlags, y_flags PlotAxisFlags) {
	x_labelArg, x_labelFin := wrapString(x_label)
	defer x_labelFin()

	y_labelArg, y_labelFin := wrapString(y_label)
	defer y_labelFin()

	C.Plot_SetupAxesV(x_labelArg, y_labelArg, C.ImPlotAxisFlags(x_flags), C.ImPlotAxisFlags(y_flags))
}

// Plot_SetupAxesLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetupAxesLimitsV(x_min float64, x_max float64, y_min float64, y_max float64, cond PlotCond) {
	C.Plot_SetupAxesLimitsV(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max), C.ImPlotCond(cond))
}

// Plot_SetupAxisV parameter default value hint:
// flags: 0
// label: ((void*)0)
func Plot_SetupAxisV(axis ImAxis, label string, flags PlotAxisFlags) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.Plot_SetupAxisV(C.ImAxis(axis), labelArg, C.ImPlotAxisFlags(flags))
}

func Plot_SetupAxisFormat_Str(axis ImAxis, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_SetupAxisFormat_Str(C.ImAxis(axis), fmtArg)
}

// Plot_SetupAxisLimitsV parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetupAxisLimitsV(axis ImAxis, v_min float64, v_max float64, cond PlotCond) {
	C.Plot_SetupAxisLimitsV(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.ImPlotCond(cond))
}

func Plot_SetupAxisLimitsConstraints(axis ImAxis, v_min float64, v_max float64) {
	C.Plot_SetupAxisLimitsConstraints(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func Plot_SetupAxisLinks(axis ImAxis, link_min *float64, link_max *float64) {
	C.Plot_SetupAxisLinks(C.ImAxis(axis), (*C.double)(link_min), (*C.double)(link_max))
}

func Plot_SetupAxisScale_PlotScale(axis ImAxis, scale PlotScale) {
	C.Plot_SetupAxisScale_PlotScale(C.ImAxis(axis), C.ImPlotScale(scale))
}

func Plot_SetupAxisZoomConstraints(axis ImAxis, z_min float64, z_max float64) {
	C.Plot_SetupAxisZoomConstraints(C.ImAxis(axis), C.double(z_min), C.double(z_max))
}

func Plot_SetupFinish() {
	C.Plot_SetupFinish()
}

// Plot_SetupLegendV parameter default value hint:
// flags: 0
func Plot_SetupLegendV(location PlotLocation, flags PlotLegendFlags) {
	C.Plot_SetupLegendV(C.ImPlotLocation(location), C.ImPlotLegendFlags(flags))
}

// Plot_SetupMouseTextV parameter default value hint:
// flags: 0
func Plot_SetupMouseTextV(location PlotLocation, flags PlotMouseTextFlags) {
	C.Plot_SetupMouseTextV(C.ImPlotLocation(location), C.ImPlotMouseTextFlags(flags))
}

func Plot_ShowColormapSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ShowColormapSelector(labelArg) == C.bool(true)
}

// Plot_ShowDemoWindowV parameter default value hint:
// p_open: ((void*)0)
func Plot_ShowDemoWindowV(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.Plot_ShowDemoWindowV(p_openArg)
}

func Plot_ShowInputMapSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ShowInputMapSelector(labelArg) == C.bool(true)
}

// Plot_ShowMetricsWindowV parameter default value hint:
// p_popen: ((void*)0)
func Plot_ShowMetricsWindowV(p_popen *bool) {
	p_popenArg, p_popenFin := wrapBool(p_popen)
	defer p_popenFin()

	C.Plot_ShowMetricsWindowV(p_popenArg)
}

// Plot_ShowStyleEditorV parameter default value hint:
// ref: ((void*)0)
func Plot_ShowStyleEditorV(ref ImPlotStyle) {
	C.Plot_ShowStyleEditorV(ref.handle())
}

func Plot_ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ShowStyleSelector(labelArg) == C.bool(true)
}

func Plot_ShowUserGuide() {
	C.Plot_ShowUserGuide()
}

// Plot_StyleColorsAutoV parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsAutoV(dst ImPlotStyle) {
	C.Plot_StyleColorsAutoV(dst.handle())
}

// Plot_StyleColorsClassicV parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsClassicV(dst ImPlotStyle) {
	C.Plot_StyleColorsClassicV(dst.handle())
}

// Plot_StyleColorsDarkV parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsDarkV(dst ImPlotStyle) {
	C.Plot_StyleColorsDarkV(dst.handle())
}

// Plot_StyleColorsLightV parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsLightV(dst ImPlotStyle) {
	C.Plot_StyleColorsLightV(dst.handle())
}

// Plot_TagX_BoolV parameter default value hint:
// round: false
func Plot_TagX_BoolV(x float64, col ImVec4, round bool) {
	C.Plot_TagX_BoolV(C.double(x), col.toC(), C.bool(round))
}

func Plot_TagX_Str(x float64, col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_TagX_Str(C.double(x), col.toC(), fmtArg)
}

// Plot_TagY_BoolV parameter default value hint:
// round: false
func Plot_TagY_BoolV(y float64, col ImVec4, round bool) {
	C.Plot_TagY_BoolV(C.double(y), col.toC(), C.bool(round))
}

func Plot_TagY_Str(y float64, col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_TagY_Str(C.double(y), col.toC(), fmtArg)
}

func Plot_AddColormap_Vec4Ptr(name string, cols *ImVec4, size int32) PlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	colsArg, colsFin := cols.wrap()
	defer colsFin()

	return PlotColormap(C.Plot_AddColormap_Vec4Ptr(nameArg, colsArg, C.int(size)))
}

func Plot_Annotation_Bool(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool) {
	C.Plot_Annotation_Bool(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp))
}

func Plot_BeginAlignedPlots(group_id string) bool {
	group_idArg, group_idFin := wrapString(group_id)
	defer group_idFin()

	return C.Plot_BeginAlignedPlots(group_idArg) == C.bool(true)
}

func Plot_BeginDragDropSourceAxis(axis ImAxis) bool {
	return C.Plot_BeginDragDropSourceAxis(C.ImAxis(axis)) == C.bool(true)
}

func Plot_BeginDragDropSourceItem(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_BeginDragDropSourceItem(label_idArg) == C.bool(true)
}

func Plot_BeginDragDropSourcePlot() bool {
	return C.Plot_BeginDragDropSourcePlot() == C.bool(true)
}

func Plot_BeginLegendPopup(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_BeginLegendPopup(label_idArg) == C.bool(true)
}

func Plot_BeginPlot(title_id string) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.Plot_BeginPlot(title_idArg) == C.bool(true)
}

func Plot_BeginSubplots(title_id string, rows int32, cols int32, size ImVec2) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.Plot_BeginSubplots(title_idArg, C.int(rows), C.int(cols), size.toC()) == C.bool(true)
}

func Plot_BustColorCache() {
	C.Plot_BustColorCache()
}

func Plot_ColormapButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ColormapButton(labelArg) == C.bool(true)
}

func Plot_ColormapScale(label string, scale_min float64, scale_max float64) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.Plot_ColormapScale(labelArg, C.double(scale_min), C.double(scale_max))
}

func Plot_ColormapSlider(label string, t *float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	tArg, tFin := wrapNumberPtr[C.float, float32](t)
	defer tFin()

	return C.Plot_ColormapSlider(labelArg, tArg) == C.bool(true)
}

func Plot_DestroyContext() {
	C.Plot_DestroyContext()
}

func Plot_DragLineX(id int32, x *float64, col ImVec4) bool {
	return C.Plot_DragLineX(C.int(id), (*C.double)(x), col.toC()) == C.bool(true)
}

func Plot_DragLineY(id int32, y *float64, col ImVec4) bool {
	return C.Plot_DragLineY(C.int(id), (*C.double)(y), col.toC()) == C.bool(true)
}

func Plot_DragPoint(id int32, x *float64, y *float64, col ImVec4) bool {
	return C.Plot_DragPoint(C.int(id), (*C.double)(x), (*C.double)(y), col.toC()) == C.bool(true)
}

func Plot_DragRect(id int32, x_min *float64, y_min *float64, x_max *float64, y_max *float64, col ImVec4) bool {
	return C.Plot_DragRect(C.int(id), (*C.double)(x_min), (*C.double)(y_min), (*C.double)(x_max), (*C.double)(y_max), col.toC()) == C.bool(true)
}

func Plot_GetColormapColor(idx int32) ImVec4 {
	pOut := ImVec4{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetColormapColor(pOutArg, C.int(idx))
	return pOut
}

func Plot_GetColormapSize() int {
	return int(C.Plot_GetColormapSize())
}

func Plot_HideNextItem() {
	C.Plot_HideNextItem()
}

func Plot_MapInputDefault() {
	C.Plot_MapInputDefault()
}

func Plot_MapInputReverse() {
	C.Plot_MapInputReverse()
}

func Plot_PlotBars_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32, bar_size float64) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count), C.double(bar_size))
}

func Plot_PlotBars_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotBars_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotBars_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, bar_size float64) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(bar_size))
}

func Plot_PlotBars_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotBars_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, bar_size float64) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(bar_size))
}

func Plot_PlotDigital_FloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_FloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func Plot_PlotDigital_S64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_S64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func Plot_PlotDigital_U64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_U64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func Plot_PlotDummy(label_id string) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDummy(label_idArg)
}

func Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_id string, xs []float32, ys []float32, neg []float32, pos []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(neg[0])), (*C.float)(&(pos[0])), C.int(count))
}

func Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_id string, xs []float32, ys []float32, err []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), (*C.float)(&(err[0])), C.int(count))
}

func Plot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_id string, xs []int64, ys []int64, err []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(err[0])), C.int(count))
}

func Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_id string, xs []int64, ys []int64, neg []int64, pos []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(neg[0])), (*C.longlong)(&(pos[0])), C.int(count))
}

func Plot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_id string, xs []uint64, ys []uint64, err []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(err[0])), C.int(count))
}

func Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, neg []uint64, pos []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(neg[0])), (*C.ulonglong)(&(pos[0])), C.int(count))
}

func Plot_PlotHeatmap_FloatPtr(label_id string, values []float32, rows int32, cols int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotHeatmap_FloatPtr(label_idArg, (*C.float)(&(values[0])), C.int(rows), C.int(cols))
}

func Plot_PlotHeatmap_S64Ptr(label_id string, values []int64, rows int32, cols int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotHeatmap_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(rows), C.int(cols))
}

func Plot_PlotHeatmap_U64Ptr(label_id string, values []uint64, rows int32, cols int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotHeatmap_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(rows), C.int(cols))
}

func Plot_PlotHistogram2D_FloatPtr(label_id string, xs []float32, ys []float32, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.Plot_PlotHistogram2D_FloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count)))
}

func Plot_PlotHistogram2D_S64Ptr(label_id string, xs []int64, ys []int64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.Plot_PlotHistogram2D_S64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count)))
}

func Plot_PlotHistogram2D_U64Ptr(label_id string, xs []uint64, ys []uint64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.Plot_PlotHistogram2D_U64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count)))
}

func Plot_PlotHistogram_FloatPtr(label_id string, values []float32, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.Plot_PlotHistogram_FloatPtr(label_idArg, (*C.float)(&(values[0])), C.int(count)))
}

func Plot_PlotHistogram_S64Ptr(label_id string, values []int64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.Plot_PlotHistogram_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(count)))
}

func Plot_PlotHistogram_U64Ptr(label_id string, values []uint64, count int32) float64 {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return float64(C.Plot_PlotHistogram_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count)))
}

func Plot_PlotInfLines_FloatPtr(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_FloatPtr(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotInfLines_S64Ptr(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotInfLines_U64Ptr(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotLine_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func Plot_PlotLine_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotLine_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotLine_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func Plot_PlotLine_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotLine_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func Plot_PlotScatter_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func Plot_PlotScatter_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotScatter_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotScatter_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func Plot_PlotScatter_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotScatter_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func Plot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_id string, xs []float32, ys1 []float32, ys2 []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys1[0])), (*C.float)(&(ys2[0])), C.int(count))
}

func Plot_PlotShaded_FloatPtrFloatPtrInt(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_FloatPtrFloatPtrInt(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func Plot_PlotShaded_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotShaded_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotShaded_S64PtrS64PtrInt(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrS64PtrInt(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func Plot_PlotShaded_S64PtrS64PtrS64Ptr(label_id string, xs []int64, ys1 []int64, ys2 []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrS64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys1[0])), (*C.longlong)(&(ys2[0])), C.int(count))
}

func Plot_PlotShaded_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotShaded_U64PtrU64PtrInt(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrU64PtrInt(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func Plot_PlotShaded_U64PtrU64PtrU64Ptr(label_id string, xs []uint64, ys1 []uint64, ys2 []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrU64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys1[0])), (*C.ulonglong)(&(ys2[0])), C.int(count))
}

func Plot_PlotStairs_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func Plot_PlotStairs_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotStairs_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotStairs_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func Plot_PlotStairs_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotStairs_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func Plot_PlotStems_FloatPtrFloatPtr(label_id string, xs []float32, ys []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_FloatPtrFloatPtr(label_idArg, (*C.float)(&(xs[0])), (*C.float)(&(ys[0])), C.int(count))
}

func Plot_PlotStems_FloatPtrInt(label_id string, values []float32, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_FloatPtrInt(label_idArg, (*C.float)(&(values[0])), C.int(count))
}

func Plot_PlotStems_S64PtrInt(label_id string, values []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count))
}

func Plot_PlotStems_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count))
}

func Plot_PlotStems_U64PtrInt(label_id string, values []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count))
}

func Plot_PlotStems_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count))
}

func Plot_PlotText(text string, x float64, y float64) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.Plot_PlotText(textArg, C.double(x), C.double(y))
}

func Plot_PlotToPixels_double(x float64, y float64) ImVec2 {
	pOut := ImVec2{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_PlotToPixels_double(pOutArg, C.double(x), C.double(y))
	return pOut
}

func Plot_PopColormap() {
	C.Plot_PopColormap()
}

func Plot_PopStyleColor() {
	C.Plot_PopStyleColor()
}

func Plot_PopStyleVar() {
	C.Plot_PopStyleVar()
}

func Plot_PushPlotClipRect() {
	C.Plot_PushPlotClipRect()
}

func Plot_SampleColormap(t float32) ImVec4 {
	pOut := ImVec4{}
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_SampleColormap(pOutArg, C.float(t))
	return pOut
}

func Plot_SetNextAxesLimits(x_min float64, x_max float64, y_min float64, y_max float64) {
	C.Plot_SetNextAxesLimits(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max))
}

func Plot_SetNextAxisLimits(axis ImAxis, v_min float64, v_max float64) {
	C.Plot_SetNextAxisLimits(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func Plot_SetNextErrorBarStyle() {
	C.Plot_SetNextErrorBarStyle()
}

func Plot_SetNextFillStyle() {
	C.Plot_SetNextFillStyle()
}

func Plot_SetNextLineStyle() {
	C.Plot_SetNextLineStyle()
}

func Plot_SetNextMarkerStyle() {
	C.Plot_SetNextMarkerStyle()
}

func Plot_SetupAxes(x_label string, y_label string) {
	x_labelArg, x_labelFin := wrapString(x_label)
	defer x_labelFin()

	y_labelArg, y_labelFin := wrapString(y_label)
	defer y_labelFin()

	C.Plot_SetupAxes(x_labelArg, y_labelArg)
}

func Plot_SetupAxesLimits(x_min float64, x_max float64, y_min float64, y_max float64) {
	C.Plot_SetupAxesLimits(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max))
}

func Plot_SetupAxis(axis ImAxis) {
	C.Plot_SetupAxis(C.ImAxis(axis))
}

func Plot_SetupAxisLimits(axis ImAxis, v_min float64, v_max float64) {
	C.Plot_SetupAxisLimits(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func Plot_SetupAxisTicks_double(axis ImAxis, v_min float64, v_max float64, n_ticks int32) {
	C.Plot_SetupAxisTicks_double(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.int(n_ticks))
}

func Plot_SetupLegend(location PlotLocation) {
	C.Plot_SetupLegend(C.ImPlotLocation(location))
}

func Plot_SetupMouseText(location PlotLocation) {
	C.Plot_SetupMouseText(C.ImPlotLocation(location))
}

func Plot_ShowDemoWindow() {
	C.Plot_ShowDemoWindow()
}

func Plot_ShowMetricsWindow() {
	C.Plot_ShowMetricsWindow()
}

func Plot_ShowStyleEditor() {
	C.Plot_ShowStyleEditor()
}

func Plot_StyleColorsAuto() {
	C.Plot_StyleColorsAuto()
}

func Plot_StyleColorsClassic() {
	C.Plot_StyleColorsClassic()
}

func Plot_StyleColorsDark() {
	C.Plot_StyleColorsDark()
}

func Plot_StyleColorsLight() {
	C.Plot_StyleColorsLight()
}

func Plot_TagX_Bool(x float64, col ImVec4) {
	C.Plot_TagX_Bool(C.double(x), col.toC())
}

func Plot_TagY_Bool(y float64, col ImVec4) {
	C.Plot_TagY_Bool(C.double(y), col.toC())
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

func (self ImPlotAxis) SetFlags(v PlotAxisFlags) {
	C.ImPlotAxis_SetFlags(self.handle(), C.ImPlotAxisFlags(v))
}

func (self ImPlotAxis) GetFlags() PlotAxisFlags {
	return PlotAxisFlags(C.ImPlotAxis_GetFlags(self.handle()))
}

func (self ImPlotAxis) SetPreviousFlags(v PlotAxisFlags) {
	C.ImPlotAxis_SetPreviousFlags(self.handle(), C.ImPlotAxisFlags(v))
}

func (self ImPlotAxis) GetPreviousFlags() PlotAxisFlags {
	return PlotAxisFlags(C.ImPlotAxis_GetPreviousFlags(self.handle()))
}

func (self ImPlotAxis) GetRange() ImPlotRange {
	return newImPlotRangeFromC(C.ImPlotAxis_GetRange(self.handle()))
}

func (self ImPlotAxis) SetRangeCond(v PlotCond) {
	C.ImPlotAxis_SetRangeCond(self.handle(), C.ImPlotCond(v))
}

func (self ImPlotAxis) GetRangeCond() PlotCond {
	return PlotCond(C.ImPlotAxis_GetRangeCond(self.handle()))
}

func (self ImPlotAxis) SetScale(v PlotScale) {
	C.ImPlotAxis_SetScale(self.handle(), C.ImPlotScale(v))
}

func (self ImPlotAxis) GetScale() PlotScale {
	return PlotScale(C.ImPlotAxis_GetScale(self.handle()))
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

func (self ImPlotDateTimeSpec) SetDate(v PlotDateFmt) {
	C.ImPlotDateTimeSpec_SetDate(self.handle(), C.ImPlotDateFmt(v))
}

func (self ImPlotDateTimeSpec) GetDate() PlotDateFmt {
	return PlotDateFmt(C.ImPlotDateTimeSpec_GetDate(self.handle()))
}

func (self ImPlotDateTimeSpec) SetTime(v PlotTimeFmt) {
	C.ImPlotDateTimeSpec_SetTime(self.handle(), C.ImPlotTimeFmt(v))
}

func (self ImPlotDateTimeSpec) GetTime() PlotTimeFmt {
	return PlotTimeFmt(C.ImPlotDateTimeSpec_GetTime(self.handle()))
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

func (self ImPlotLegend) SetFlags(v PlotLegendFlags) {
	C.ImPlotLegend_SetFlags(self.handle(), C.ImPlotLegendFlags(v))
}

func (self ImPlotLegend) GetFlags() PlotLegendFlags {
	return PlotLegendFlags(C.ImPlotLegend_GetFlags(self.handle()))
}

func (self ImPlotLegend) SetPreviousFlags(v PlotLegendFlags) {
	C.ImPlotLegend_SetPreviousFlags(self.handle(), C.ImPlotLegendFlags(v))
}

func (self ImPlotLegend) GetPreviousFlags() PlotLegendFlags {
	return PlotLegendFlags(C.ImPlotLegend_GetPreviousFlags(self.handle()))
}

func (self ImPlotLegend) SetLocation(v PlotLocation) {
	C.ImPlotLegend_SetLocation(self.handle(), C.ImPlotLocation(v))
}

func (self ImPlotLegend) GetLocation() PlotLocation {
	return PlotLocation(C.ImPlotLegend_GetLocation(self.handle()))
}

func (self ImPlotLegend) SetPreviousLocation(v PlotLocation) {
	C.ImPlotLegend_SetPreviousLocation(self.handle(), C.ImPlotLocation(v))
}

func (self ImPlotLegend) GetPreviousLocation() PlotLocation {
	return PlotLocation(C.ImPlotLegend_GetPreviousLocation(self.handle()))
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

func (self ImPlotNextItemData) SetMarker(v PlotMarker) {
	C.ImPlotNextItemData_SetMarker(self.handle(), C.ImPlotMarker(v))
}

func (self ImPlotNextItemData) GetMarker() PlotMarker {
	return PlotMarker(C.ImPlotNextItemData_GetMarker(self.handle()))
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

func (self ImPlotNextItemData) SetHiddenCond(v PlotCond) {
	C.ImPlotNextItemData_SetHiddenCond(self.handle(), C.ImPlotCond(v))
}

func (self ImPlotNextItemData) GetHiddenCond() PlotCond {
	return PlotCond(C.ImPlotNextItemData_GetHiddenCond(self.handle()))
}

func (self ImPlotPlot) SetID(v ImGuiID) {
	C.ImPlotPlot_SetID(self.handle(), C.ImGuiID(v))
}

func (self ImPlotPlot) GetID() ImGuiID {
	return ImGuiID(C.ImPlotPlot_GetID(self.handle()))
}

func (self ImPlotPlot) SetFlags(v PlotFlags) {
	C.ImPlotPlot_SetFlags(self.handle(), C.ImPlotFlags(v))
}

func (self ImPlotPlot) GetFlags() PlotFlags {
	return PlotFlags(C.ImPlotPlot_GetFlags(self.handle()))
}

func (self ImPlotPlot) SetPreviousFlags(v PlotFlags) {
	C.ImPlotPlot_SetPreviousFlags(self.handle(), C.ImPlotFlags(v))
}

func (self ImPlotPlot) GetPreviousFlags() PlotFlags {
	return PlotFlags(C.ImPlotPlot_GetPreviousFlags(self.handle()))
}

func (self ImPlotPlot) SetMouseTextLocation(v PlotLocation) {
	C.ImPlotPlot_SetMouseTextLocation(self.handle(), C.ImPlotLocation(v))
}

func (self ImPlotPlot) GetMouseTextLocation() PlotLocation {
	return PlotLocation(C.ImPlotPlot_GetMouseTextLocation(self.handle()))
}

func (self ImPlotPlot) SetMouseTextFlags(v PlotMouseTextFlags) {
	C.ImPlotPlot_SetMouseTextFlags(self.handle(), C.ImPlotMouseTextFlags(v))
}

func (self ImPlotPlot) GetMouseTextFlags() PlotMouseTextFlags {
	return PlotMouseTextFlags(C.ImPlotPlot_GetMouseTextFlags(self.handle()))
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

func (self ImPlotStyle) SetColormap(v PlotColormap) {
	C.ImPlotStyle_SetColormap(self.handle(), C.ImPlotColormap(v))
}

func (self ImPlotStyle) GetColormap() PlotColormap {
	return PlotColormap(C.ImPlotStyle_GetColormap(self.handle()))
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

func (self ImPlotSubplot) SetFlags(v PlotSubplotFlags) {
	C.ImPlotSubplot_SetFlags(self.handle(), C.ImPlotSubplotFlags(v))
}

func (self ImPlotSubplot) GetFlags() PlotSubplotFlags {
	return PlotSubplotFlags(C.ImPlotSubplot_GetFlags(self.handle()))
}

func (self ImPlotSubplot) SetPreviousFlags(v PlotSubplotFlags) {
	C.ImPlotSubplot_SetPreviousFlags(self.handle(), C.ImPlotSubplotFlags(v))
}

func (self ImPlotSubplot) GetPreviousFlags() PlotSubplotFlags {
	return PlotSubplotFlags(C.ImPlotSubplot_GetPreviousFlags(self.handle()))
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
