package cimgui

// #include "extra_types.h"
// #include "cimplot_structs_accessor.h"
// #include "cimplot_wrapper.h"
import "C"
import "unsafe"

func Plot_SetCurrentContext(ctx ImPlotContext) {
	C.Plot_SetCurrentContext(ctx.handle())
}

func Plot_SetNextAxesToFit() {
	C.Plot_SetNextAxesToFit()
}

func (self ImPlotRect) Contains_double(x float64, y float64) bool {
	return C.PlotRect_Contains_double(self.handle(), C.double(x), C.double(y)) == C.bool(true)
}

func (self ImPlotStyle) Destroy() {
	C.PlotStyle_Destroy(self.handle())
}

func Plot_PopPlotClipRect() {
	C.Plot_PopPlotClipRect()
}

func Plot_ShowInputMapSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ShowInputMapSelector(labelArg) == C.bool(true)
}

func (self ImPlotRange) Size() float64 {
	return float64(C.PlotRange_Size(self.handle()))
}

// Plot_BeginDragDropSourceItem parameter default value hint:
// flags: 0
func Plot_BeginDragDropSourceItem(label_id string, flags ImGuiDragDropFlags) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_BeginDragDropSourceItem(label_idArg, C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

func Plot_BeginDragDropTargetPlot() bool {
	return C.Plot_BeginDragDropTargetPlot() == C.bool(true)
}

// Plot_ColormapButton parameter default value hint:
// cmap: -1
// size: ImVec2(0,0)
func Plot_ColormapButton(label string, size ImVec2, cmap ImPlotColormap) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ColormapButton(labelArg, size.toC(), C.ImPlotColormap(cmap)) == C.bool(true)
}

// Plot_ShowMetricsWindow parameter default value hint:
// p_popen: ((void*)0)
func Plot_ShowMetricsWindow(p_popen *bool) {
	p_popenArg, p_popenFin := wrapBool(p_popen)
	defer p_popenFin()

	C.Plot_ShowMetricsWindow(p_popenArg)
}

func Plot_EndLegendPopup() {
	C.Plot_EndLegendPopup()
}

// Plot_TagX_Bool parameter default value hint:
// round: false
func Plot_TagX_Bool(x float64, col ImVec4, round bool) {
	C.Plot_TagX_Bool(C.double(x), col.toC(), C.bool(round))
}

func Plot_TagX_Str(x float64, col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_TagX_Str(C.double(x), col.toC(), fmtArg)
}

// Plot_GetColormapSize parameter default value hint:
// cmap: -1
func Plot_GetColormapSize(cmap ImPlotColormap) int {
	return int(C.Plot_GetColormapSize(C.ImPlotColormap(cmap)))
}

// Plot_PlotDigital_FloatPtr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotDigital_FloatPtr(label_id string, xs *float32, ys *float32, count int32, flags ImPlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotDigital_FloatPtr(label_idArg, xsArg, ysArg, C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotDigital_S64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotDigital_S64Ptr(label_id string, xs []int64, ys []int64, count int32, flags ImPlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_S64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotDigital_U64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotDigital_U64Ptr(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotDigitalFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDigital_U64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotDigitalFlags(flags), C.int(offset), C.int(stride))
}

func Plot_ShowUserGuide() {
	C.Plot_ShowUserGuide()
}

func Plot_SetupAxisScale_PlotScale(axis ImAxis, scale ImPlotScale) {
	C.Plot_SetupAxisScale_PlotScale(C.ImAxis(axis), C.ImPlotScale(scale))
}

// Plot_StyleColorsClassic parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsClassic(dst ImPlotStyle) {
	C.Plot_StyleColorsClassic(dst.handle())
}

func (self ImPlotDateTimeSpec) Destroy() {
	C.PlotDateTimeSpec_Destroy(self.handle())
}

func (self ImPlotRange) Destroy() {
	C.PlotRange_Destroy(self.handle())
}

func Plot_GetColormapCount() int {
	return int(C.Plot_GetColormapCount())
}

// Plot_PlotText parameter default value hint:
// flags: 0
// pix_offset: ImVec2(0,0)
func Plot_PlotText(text string, x float64, y float64, pix_offset ImVec2, flags ImPlotTextFlags) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.Plot_PlotText(textArg, C.double(x), C.double(y), pix_offset.toC(), C.ImPlotTextFlags(flags))
}

// Plot_PlotToPixels_double parameter default value hint:
// x_axis: -1
// y_axis: -1
func Plot_PlotToPixels_double(pOut *ImVec2, x float64, y float64, x_axis ImAxis, y_axis ImAxis) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_PlotToPixels_double(pOutArg, C.double(x), C.double(y), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

// Plot_PushPlotClipRect parameter default value hint:
// expand: 0
func Plot_PushPlotClipRect(expand float32) {
	C.Plot_PushPlotClipRect(C.float(expand))
}

func (self ImPlotTime) Destroy() {
	C.PlotTime_Destroy(self.handle())
}

func Plot_ColormapIcon(cmap ImPlotColormap) {
	C.Plot_ColormapIcon(C.ImPlotColormap(cmap))
}

func Plot_SetupFinish() {
	C.Plot_SetupFinish()
}

// Plot_StyleColorsLight parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsLight(dst ImPlotStyle) {
	C.Plot_StyleColorsLight(dst.handle())
}

// Plot_PlotDummy parameter default value hint:
// flags: 0
func Plot_PlotDummy(label_id string, flags ImPlotDummyFlags) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotDummy(label_idArg, C.ImPlotDummyFlags(flags))
}

// Plot_BeginSubplots parameter default value hint:
// col_ratios: ((void*)0)
// flags: 0
// row_ratios: ((void*)0)
func Plot_BeginSubplots(title_id string, rows int32, cols int32, size ImVec2, flags ImPlotSubplotFlags, row_ratios *float32, col_ratios *float32) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	row_ratiosArg, row_ratiosFin := wrapFloat(row_ratios)
	defer row_ratiosFin()

	col_ratiosArg, col_ratiosFin := wrapFloat(col_ratios)
	defer col_ratiosFin()

	return C.Plot_BeginSubplots(title_idArg, C.int(rows), C.int(cols), size.toC(), C.ImPlotSubplotFlags(flags), row_ratiosArg, col_ratiosArg) == C.bool(true)
}

func PlotRect_Min(pOut ImPlotPoint, self ImPlotRect) {
	C.PlotRect_Min(pOut.handle(), self.handle())
}

func Plot_BeginDragDropTargetLegend() bool {
	return C.Plot_BeginDragDropTargetLegend() == C.bool(true)
}

// Plot_BeginPlot parameter default value hint:
// flags: 0
// size: ImVec2(-1,0)
func Plot_BeginPlot(title_id string, size ImVec2, flags ImPlotFlags) bool {
	title_idArg, title_idFin := wrapString(title_id)
	defer title_idFin()

	return C.Plot_BeginPlot(title_idArg, size.toC(), C.ImPlotFlags(flags)) == C.bool(true)
}

// Plot_SetNextAxesLimits parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetNextAxesLimits(x_min float64, x_max float64, y_min float64, y_max float64, cond ImPlotCond) {
	C.Plot_SetNextAxesLimits(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max), C.ImPlotCond(cond))
}

func Plot_GetStyleColorName(idx ImPlotCol) string {
	return C.GoString(C.Plot_GetStyleColorName(C.ImPlotCol(idx)))
}

func Plot_GetLastItemColor(pOut *ImVec4) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetLastItemColor(pOutArg)
}

func NewImPlotInputMap() ImPlotInputMap {
	return (ImPlotInputMap)(unsafe.Pointer(C.PlotInputMap_ImPlotInputMap()))
}

func (self ImPlotPlot) Destroy() {
	C.PlotPlot_Destroy(self.handle())
}

func Plot_SetAxes(x_axis ImAxis, y_axis ImAxis) {
	C.Plot_SetAxes(C.ImAxis(x_axis), C.ImAxis(y_axis))
}

// Plot_SetupLegend parameter default value hint:
// flags: 0
func Plot_SetupLegend(location ImPlotLocation, flags ImPlotLegendFlags) {
	C.Plot_SetupLegend(C.ImPlotLocation(location), C.ImPlotLegendFlags(flags))
}

// Plot_ShowStyleEditor parameter default value hint:
// ref: ((void*)0)
func Plot_ShowStyleEditor(ref ImPlotStyle) {
	C.Plot_ShowStyleEditor(ref.handle())
}

// Plot_SampleColormap parameter default value hint:
// cmap: -1
func Plot_SampleColormap(pOut *ImVec4, t float32, cmap ImPlotColormap) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_SampleColormap(pOutArg, C.float(t), C.ImPlotColormap(cmap))
}

func Plot_SetImGuiContext(ctx ImGuiContext) {
	C.Plot_SetImGuiContext(ctx.handle())
}

// Plot_SetupAxesLimits parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetupAxesLimits(x_min float64, x_max float64, y_min float64, y_max float64, cond ImPlotCond) {
	C.Plot_SetupAxesLimits(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max), C.ImPlotCond(cond))
}

// Plot_TagY_Bool parameter default value hint:
// round: false
func Plot_TagY_Bool(y float64, col ImVec4, round bool) {
	C.Plot_TagY_Bool(C.double(y), col.toC(), C.bool(round))
}

func Plot_TagY_Str(y float64, col ImVec4, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_TagY_Str(C.double(y), col.toC(), fmtArg)
}

func (self ImPlotAlignmentData) Destroy() {
	C.PlotAlignmentData_Destroy(self.handle())
}

func Plot_EndSubplots() {
	C.Plot_EndSubplots()
}

func Plot_SetupAxisLinks(axis ImAxis, link_min *float64, link_max *float64) {
	C.Plot_SetupAxisLinks(C.ImAxis(axis), (*C.double)(link_min), (*C.double)(link_max))
}

// Plot_PixelsToPlot_Vec2 parameter default value hint:
// x_axis: -1
// y_axis: -1
func Plot_PixelsToPlot_Vec2(pOut ImPlotPoint, pix ImVec2, x_axis ImAxis, y_axis ImAxis) {
	C.Plot_PixelsToPlot_Vec2(pOut.handle(), pix.toC(), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

// Plot_PixelsToPlot_Float parameter default value hint:
// x_axis: -1
// y_axis: -1
func Plot_PixelsToPlot_Float(pOut ImPlotPoint, x float32, y float32, x_axis ImAxis, y_axis ImAxis) {
	C.Plot_PixelsToPlot_Float(pOut.handle(), C.float(x), C.float(y), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

func (self ImPlotTicker) Destroy() {
	C.PlotTicker_Destroy(self.handle())
}

func Plot_EndPlot() {
	C.Plot_EndPlot()
}

func Plot_GetPlotSize(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetPlotSize(pOutArg)
}

func (self ImPlotRange) Contains(value float64) bool {
	return C.PlotRange_Contains(self.handle(), C.double(value)) == C.bool(true)
}

func (self ImPlotTick) Destroy() {
	C.PlotTick_Destroy(self.handle())
}

// Plot_Annotation_Bool parameter default value hint:
// round: false
func Plot_Annotation_Bool(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool, round bool) {
	C.Plot_Annotation_Bool(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp), C.bool(round))
}

func Plot_Annotation_Str(x float64, y float64, col ImVec4, pix_offset ImVec2, clamp bool, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_Annotation_Str(C.double(x), C.double(y), col.toC(), pix_offset.toC(), C.bool(clamp), fmtArg)
}

func (self ImPlotTagCollection) Destroy() {
	C.PlotTagCollection_Destroy(self.handle())
}

// Plot_BeginAlignedPlots parameter default value hint:
// vertical: true
func Plot_BeginAlignedPlots(group_id string, vertical bool) bool {
	group_idArg, group_idFin := wrapString(group_id)
	defer group_idFin()

	return C.Plot_BeginAlignedPlots(group_idArg, C.bool(vertical)) == C.bool(true)
}

// Plot_BeginDragDropSourcePlot parameter default value hint:
// flags: 0
func Plot_BeginDragDropSourcePlot(flags ImGuiDragDropFlags) bool {
	return C.Plot_BeginDragDropSourcePlot(C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

// Plot_StyleColorsAuto parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsAuto(dst ImPlotStyle) {
	C.Plot_StyleColorsAuto(dst.handle())
}

func (self ImPlotLegend) Destroy() {
	C.PlotLegend_Destroy(self.handle())
}

// Plot_SetNextFillStyle parameter default value hint:
// alpha_mod: -1
// col: ImVec4(0,0,0,-1)
func Plot_SetNextFillStyle(col ImVec4, alpha_mod float32) {
	C.Plot_SetNextFillStyle(col.toC(), C.float(alpha_mod))
}

func Plot_CancelPlotSelection() {
	C.Plot_CancelPlotSelection()
}

func Plot_EndDragDropSource() {
	C.Plot_EndDragDropSource()
}

// Plot_GetPlotMousePos parameter default value hint:
// x_axis: -1
// y_axis: -1
func Plot_GetPlotMousePos(pOut ImPlotPoint, x_axis ImAxis, y_axis ImAxis) {
	C.Plot_GetPlotMousePos(pOut.handle(), C.ImAxis(x_axis), C.ImAxis(y_axis))
}

func (self ImPlotNextPlotData) Destroy() {
	C.PlotNextPlotData_Destroy(self.handle())
}

func (self ImPlotSubplot) Destroy() {
	C.PlotSubplot_Destroy(self.handle())
}

func Plot_GetColormapIndex(name string) ImPlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	return ImPlotColormap(C.Plot_GetColormapIndex(nameArg))
}

// Plot_PlotStairs_FloatPtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func Plot_PlotStairs_FloatPtrInt(label_id string, values *float32, count int32, xscale float64, xstart float64, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotStairs_FloatPtrInt(label_idArg, valuesArg, C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_S64PtrInt parameter default value hint:
// xstart: 0
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
func Plot_PlotStairs_S64PtrInt(label_id string, values []int64, count int32, xscale float64, xstart float64, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_U64PtrInt parameter default value hint:
// xstart: 0
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
func Plot_PlotStairs_U64PtrInt(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_FloatPtrFloatPtr parameter default value hint:
// stride: sizeof(float)
// flags: 0
// offset: 0
func Plot_PlotStairs_FloatPtrFloatPtr(label_id string, xs *float32, ys *float32, count int32, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotStairs_FloatPtrFloatPtr(label_idArg, xsArg, ysArg, C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_S64PtrS64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotStairs_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStairs_U64PtrU64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotStairs_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotStairsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStairs_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotStairsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_SetupAxisLimits parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetupAxisLimits(axis ImAxis, v_min float64, v_max float64, cond ImPlotCond) {
	C.Plot_SetupAxisLimits(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.ImPlotCond(cond))
}

func (self ImPlotNextItemData) Destroy() {
	C.PlotNextItemData_Destroy(self.handle())
}

// Plot_GetColormapColor parameter default value hint:
// cmap: -1
func Plot_GetColormapColor(pOut *ImVec4, idx int32, cmap ImPlotColormap) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetColormapColor(pOutArg, C.int(idx), C.ImPlotColormap(cmap))
}

// Plot_PopColormap parameter default value hint:
// count: 1
func Plot_PopColormap(count int32) {
	C.Plot_PopColormap(C.int(count))
}

// Plot_SetNextErrorBarStyle parameter default value hint:
// col: ImVec4(0,0,0,-1)
// size: -1
// weight: -1
func Plot_SetNextErrorBarStyle(col ImVec4, size float32, weight float32) {
	C.Plot_SetNextErrorBarStyle(col.toC(), C.float(size), C.float(weight))
}

func Plot_GetCurrentContext() ImPlotContext {
	return (ImPlotContext)(unsafe.Pointer(C.Plot_GetCurrentContext()))
}

func Plot_IsSubplotsHovered() bool {
	return C.Plot_IsSubplotsHovered() == C.bool(true)
}

func (self ImPlotInputMap) Destroy() {
	C.PlotInputMap_Destroy(self.handle())
}

func Plot_CreateContext() ImPlotContext {
	return (ImPlotContext)(unsafe.Pointer(C.Plot_CreateContext()))
}

func (self ImPlotRange) Clamp(value float64) float64 {
	return float64(C.PlotRange_Clamp(self.handle(), C.double(value)))
}

func Plot_SetAxis(axis ImAxis) {
	C.Plot_SetAxis(C.ImAxis(axis))
}

func NewImPlotPointNil() ImPlotPoint {
	return (ImPlotPoint)(unsafe.Pointer(C.PlotPoint_ImPlotPoint_Nil()))
}

func NewImPlotPointdouble(_x float64, _y float64) ImPlotPoint {
	return (ImPlotPoint)(unsafe.Pointer(C.PlotPoint_ImPlotPoint_double(C.double(_x), C.double(_y))))
}

func NewImPlotPointVec2(p ImVec2) ImPlotPoint {
	return (ImPlotPoint)(unsafe.Pointer(C.PlotPoint_ImPlotPoint_Vec2(p.toC())))
}

func Plot_GetStyle() ImPlotStyle {
	return (ImPlotStyle)(unsafe.Pointer(C.Plot_GetStyle()))
}

// Plot_HideNextItem parameter default value hint:
// cond: ImPlotCond_Once
// hidden: true
func Plot_HideNextItem(hidden bool, cond ImPlotCond) {
	C.Plot_HideNextItem(C.bool(hidden), C.ImPlotCond(cond))
}

func Plot_IsAxisHovered(axis ImAxis) bool {
	return C.Plot_IsAxisHovered(C.ImAxis(axis)) == C.bool(true)
}

// Plot_SetupAxis parameter default value hint:
// flags: 0
// label: ((void*)0)
func Plot_SetupAxis(axis ImAxis, label string, flags ImPlotAxisFlags) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.Plot_SetupAxis(C.ImAxis(axis), labelArg, C.ImPlotAxisFlags(flags))
}

func Plot_SetNextAxisToFit(axis ImAxis) {
	C.Plot_SetNextAxisToFit(C.ImAxis(axis))
}

// Plot_PopStyleVar parameter default value hint:
// count: 1
func Plot_PopStyleVar(count int32) {
	C.Plot_PopStyleVar(C.int(count))
}

func Plot_PushColormap_PlotColormap(cmap ImPlotColormap) {
	C.Plot_PushColormap_PlotColormap(C.ImPlotColormap(cmap))
}

func Plot_PushColormap_Str(name string) {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	C.Plot_PushColormap_Str(nameArg)
}

// Plot_SetNextAxisLimits parameter default value hint:
// cond: ImPlotCond_Once
func Plot_SetNextAxisLimits(axis ImAxis, v_min float64, v_max float64, cond ImPlotCond) {
	C.Plot_SetNextAxisLimits(C.ImAxis(axis), C.double(v_min), C.double(v_max), C.ImPlotCond(cond))
}

// Plot_MapInputReverse parameter default value hint:
// dst: ((void*)0)
func Plot_MapInputReverse(dst ImPlotInputMap) {
	C.Plot_MapInputReverse(dst.handle())
}

func Plot_PushStyleColor_U32(idx ImPlotCol, col uint32) {
	C.Plot_PushStyleColor_U32(C.ImPlotCol(idx), C.ImU32(col))
}

func Plot_PushStyleColor_Vec4(idx ImPlotCol, col ImVec4) {
	C.Plot_PushStyleColor_Vec4(C.ImPlotCol(idx), col.toC())
}

func Plot_GetPlotPos(pOut *ImVec2) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_GetPlotPos(pOutArg)
}

// Plot_MapInputDefault parameter default value hint:
// dst: ((void*)0)
func Plot_MapInputDefault(dst ImPlotInputMap) {
	C.Plot_MapInputDefault(dst.handle())
}

func (self ImPlotRect) Destroy() {
	C.PlotRect_Destroy(self.handle())
}

// Plot_PlotScatter_FloatPtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
func Plot_PlotScatter_FloatPtrInt(label_id string, values *float32, count int32, xscale float64, xstart float64, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotScatter_FloatPtrInt(label_idArg, valuesArg, C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_S64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func Plot_PlotScatter_S64PtrInt(label_id string, values []int64, count int32, xscale float64, xstart float64, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_U64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
func Plot_PlotScatter_U64PtrInt(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_FloatPtrFloatPtr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotScatter_FloatPtrFloatPtr(label_id string, xs *float32, ys *float32, count int32, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotScatter_FloatPtrFloatPtr(label_idArg, xsArg, ysArg, C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_S64PtrS64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotScatter_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotScatter_U64PtrU64Ptr parameter default value hint:
// offset: 0
// stride: sizeof(ImU64)
// flags: 0
func Plot_PlotScatter_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotScatterFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotScatter_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotScatterFlags(flags), C.int(offset), C.int(stride))
}

func Plot_SetNextAxisLinks(axis ImAxis, link_min *float64, link_max *float64) {
	C.Plot_SetNextAxisLinks(C.ImAxis(axis), (*C.double)(link_min), (*C.double)(link_max))
}

func PlotRect_Size(pOut ImPlotPoint, self ImPlotRect) {
	C.PlotRect_Size(pOut.handle(), self.handle())
}

// Plot_BustColorCache parameter default value hint:
// plot_title_id: ((void*)0)
func Plot_BustColorCache(plot_title_id string) {
	plot_title_idArg, plot_title_idFin := wrapString(plot_title_id)
	defer plot_title_idFin()

	C.Plot_BustColorCache(plot_title_idArg)
}

func Plot_EndDragDropTarget() {
	C.Plot_EndDragDropTarget()
}

// Plot_PlotShaded_FloatPtrInt parameter default value hint:
// offset: 0
// stride: sizeof(float)
// xscale: 1
// xstart: 0
// yref: 0
// flags: 0
func Plot_PlotShaded_FloatPtrInt(label_id string, values *float32, count int32, yref float64, xscale float64, xstart float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotShaded_FloatPtrInt(label_idArg, valuesArg, C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_S64PtrInt parameter default value hint:
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
// yref: 0
// flags: 0
func Plot_PlotShaded_S64PtrInt(label_id string, values []int64, count int32, yref float64, xscale float64, xstart float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_U64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
// yref: 0
func Plot_PlotShaded_U64PtrInt(label_id string, values []uint64, count int32, yref float64, xscale float64, xstart float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(yref), C.double(xscale), C.double(xstart), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_FloatPtrFloatPtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
// yref: 0
func Plot_PlotShaded_FloatPtrFloatPtrInt(label_id string, xs *float32, ys *float32, count int32, yref float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotShaded_FloatPtrFloatPtrInt(label_idArg, xsArg, ysArg, C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_S64PtrS64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// yref: 0
func Plot_PlotShaded_S64PtrS64PtrInt(label_id string, xs []int64, ys []int64, count int32, yref float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrS64PtrInt(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_U64PtrU64PtrInt parameter default value hint:
// stride: sizeof(ImU64)
// yref: 0
// flags: 0
// offset: 0
func Plot_PlotShaded_U64PtrU64PtrInt(label_id string, xs []uint64, ys []uint64, count int32, yref float64, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrU64PtrInt(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(yref), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_FloatPtrFloatPtrFloatPtr parameter default value hint:
// stride: sizeof(float)
// flags: 0
// offset: 0
func Plot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_id string, xs *float32, ys1 *float32, ys2 *float32, count int32, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ys1Arg, ys1Fin := wrapFloat(ys1)
	defer ys1Fin()

	ys2Arg, ys2Fin := wrapFloat(ys2)
	defer ys2Fin()

	C.Plot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_idArg, xsArg, ys1Arg, ys2Arg, C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_S64PtrS64PtrS64Ptr parameter default value hint:
// stride: sizeof(ImS64)
// flags: 0
// offset: 0
func Plot_PlotShaded_S64PtrS64PtrS64Ptr(label_id string, xs []int64, ys1 []int64, ys2 []int64, count int32, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_S64PtrS64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys1[0])), (*C.longlong)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotShaded_U64PtrU64PtrU64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotShaded_U64PtrU64PtrU64Ptr(label_id string, xs []uint64, ys1 []uint64, ys2 []uint64, count int32, flags ImPlotShadedFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotShaded_U64PtrU64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys1[0])), (*C.ulonglong)(&(ys2[0])), C.int(count), C.ImPlotShadedFlags(flags), C.int(offset), C.int(stride))
}

func Plot_GetColormapName(cmap ImPlotColormap) string {
	return C.GoString(C.Plot_GetColormapName(C.ImPlotColormap(cmap)))
}

// Plot_PlotInfLines_FloatPtr parameter default value hint:
// stride: sizeof(float)
// flags: 0
// offset: 0
func Plot_PlotInfLines_FloatPtr(label_id string, values *float32, count int32, flags ImPlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotInfLines_FloatPtr(label_idArg, valuesArg, C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotInfLines_S64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotInfLines_S64Ptr(label_id string, values []int64, count int32, flags ImPlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_S64Ptr(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotInfLines_U64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotInfLines_U64Ptr(label_id string, values []uint64, count int32, flags ImPlotInfLinesFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotInfLines_U64Ptr(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.ImPlotInfLinesFlags(flags), C.int(offset), C.int(stride))
}

// Plot_SetNextMarkerStyle parameter default value hint:
// fill: ImVec4(0,0,0,-1)
// marker: -1
// outline: ImVec4(0,0,0,-1)
// size: -1
// weight: -1
func Plot_SetNextMarkerStyle(marker ImPlotMarker, size float32, fill ImVec4, weight float32, outline ImVec4) {
	C.Plot_SetNextMarkerStyle(C.ImPlotMarker(marker), C.float(size), fill.toC(), C.float(weight), outline.toC())
}

// Plot_SetNextLineStyle parameter default value hint:
// weight: -1
// col: ImVec4(0,0,0,-1)
func Plot_SetNextLineStyle(col ImVec4, weight float32) {
	C.Plot_SetNextLineStyle(col.toC(), C.float(weight))
}

func (self ImPlotAxis) Destroy() {
	C.PlotAxis_Destroy(self.handle())
}

// Plot_PlotLine_FloatPtrInt parameter default value hint:
// xscale: 1
// xstart: 0
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotLine_FloatPtrInt(label_id string, values *float32, count int32, xscale float64, xstart float64, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotLine_FloatPtrInt(label_idArg, valuesArg, C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_S64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
// xscale: 1
// xstart: 0
func Plot_PlotLine_S64PtrInt(label_id string, values []int64, count int32, xscale float64, xstart float64, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_U64PtrInt parameter default value hint:
// stride: sizeof(ImU64)
// xscale: 1
// xstart: 0
// flags: 0
// offset: 0
func Plot_PlotLine_U64PtrInt(label_id string, values []uint64, count int32, xscale float64, xstart float64, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(xscale), C.double(xstart), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_FloatPtrFloatPtr parameter default value hint:
// stride: sizeof(float)
// flags: 0
// offset: 0
func Plot_PlotLine_FloatPtrFloatPtr(label_id string, xs *float32, ys *float32, count int32, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotLine_FloatPtrFloatPtr(label_idArg, xsArg, ysArg, C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_S64PtrS64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotLine_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotLine_U64PtrU64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotLine_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, flags ImPlotLineFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotLine_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.ImPlotLineFlags(flags), C.int(offset), C.int(stride))
}

func Plot_IsPlotHovered() bool {
	return C.Plot_IsPlotHovered() == C.bool(true)
}

func Plot_SetupAxisLimitsConstraints(axis ImAxis, v_min float64, v_max float64) {
	C.Plot_SetupAxisLimitsConstraints(C.ImAxis(axis), C.double(v_min), C.double(v_max))
}

func (self ImPlotAnnotationCollection) Destroy() {
	C.PlotAnnotationCollection_Destroy(self.handle())
}

func (self ImPlotColormapData) Destroy() {
	C.PlotColormapData_Destroy(self.handle())
}

func NewImPlotRangeNil() ImPlotRange {
	return (ImPlotRange)(unsafe.Pointer(C.PlotRange_ImPlotRange_Nil()))
}

func NewImPlotRangedouble(_min float64, _max float64) ImPlotRange {
	return (ImPlotRange)(unsafe.Pointer(C.PlotRange_ImPlotRange_double(C.double(_min), C.double(_max))))
}

// Plot_AddColormap_Vec4Ptr parameter default value hint:
// qual: true
func Plot_AddColormap_Vec4Ptr(name string, cols *ImVec4, size int32, qual bool) ImPlotColormap {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	colsArg, colsFin := cols.wrap()
	defer colsFin()

	return ImPlotColormap(C.Plot_AddColormap_Vec4Ptr(nameArg, colsArg, C.int(size), C.bool(qual)))
}

func Plot_IsLegendEntryHovered(label_id string) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_IsLegendEntryHovered(label_idArg) == C.bool(true)
}

func NewImPlotRectNil() ImPlotRect {
	return (ImPlotRect)(unsafe.Pointer(C.PlotRect_ImPlotRect_Nil()))
}

func NewImPlotRectdouble(x_min float64, x_max float64, y_min float64, y_max float64) ImPlotRect {
	return (ImPlotRect)(unsafe.Pointer(C.PlotRect_ImPlotRect_double(C.double(x_min), C.double(x_max), C.double(y_min), C.double(y_max))))
}

func Plot_BeginDragDropTargetAxis(axis ImAxis) bool {
	return C.Plot_BeginDragDropTargetAxis(C.ImAxis(axis)) == C.bool(true)
}

func Plot_EndAlignedPlots() {
	C.Plot_EndAlignedPlots()
}

func Plot_IsPlotSelected() bool {
	return C.Plot_IsPlotSelected() == C.bool(true)
}

func (self ImPlotPointError) Destroy() {
	C.PlotPointError_Destroy(self.handle())
}

func Plot_GetPlotDrawList() ImDrawList {
	return (ImDrawList)(unsafe.Pointer(C.Plot_GetPlotDrawList()))
}

// Plot_SetupAxes parameter default value hint:
// x_flags: 0
// y_flags: 0
func Plot_SetupAxes(x_label string, y_label string, x_flags ImPlotAxisFlags, y_flags ImPlotAxisFlags) {
	x_labelArg, x_labelFin := wrapString(x_label)
	defer x_labelFin()

	y_labelArg, y_labelFin := wrapString(y_label)
	defer y_labelFin()

	C.Plot_SetupAxes(x_labelArg, y_labelArg, C.ImPlotAxisFlags(x_flags), C.ImPlotAxisFlags(y_flags))
}

func Plot_PushStyleVar_Float(idx ImPlotStyleVar, val float32) {
	C.Plot_PushStyleVar_Float(C.ImPlotStyleVar(idx), C.float(val))
}

func Plot_PushStyleVar_Int(idx ImPlotStyleVar, val int32) {
	C.Plot_PushStyleVar_Int(C.ImPlotStyleVar(idx), C.int(val))
}

func Plot_PushStyleVar_Vec2(idx ImPlotStyleVar, val ImVec2) {
	C.Plot_PushStyleVar_Vec2(C.ImPlotStyleVar(idx), val.toC())
}

func PlotRect_Clamp_double(pOut ImPlotPoint, self ImPlotRect, x float64, y float64) {
	C.PlotRect_Clamp_double(pOut.handle(), self.handle(), C.double(x), C.double(y))
}

// Plot_BeginLegendPopup parameter default value hint:
// mouse_button: 1
func Plot_BeginLegendPopup(label_id string, mouse_button ImGuiMouseButton) bool {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	return C.Plot_BeginLegendPopup(label_idArg, C.ImGuiMouseButton(mouse_button)) == C.bool(true)
}

// Plot_DragLineY parameter default value hint:
// flags: 0
// thickness: 1
func Plot_DragLineY(id int32, y *float64, col ImVec4, thickness float32, flags ImPlotDragToolFlags) bool {
	return C.Plot_DragLineY(C.int(id), (*C.double)(y), col.toC(), C.float(thickness), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

func Plot_ShowColormapSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ShowColormapSelector(labelArg) == C.bool(true)
}

// Plot_ColormapScale parameter default value hint:
// cmap: -1
// flags: 0
// format: "%g"
// size: ImVec2(0,0)
func Plot_ColormapScale(label string, scale_min float64, scale_max float64, size ImVec2, format string, flags ImPlotColormapScaleFlags, cmap ImPlotColormap) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	C.Plot_ColormapScale(labelArg, C.double(scale_min), C.double(scale_max), size.toC(), formatArg, C.ImPlotColormapScaleFlags(flags), C.ImPlotColormap(cmap))
}

// Plot_ColormapSlider parameter default value hint:
// cmap: -1
// format: ""
// out: ((void*)0)
func Plot_ColormapSlider(label string, t *float32, out *ImVec4, format string, cmap ImPlotColormap) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	tArg, tFin := wrapFloat(t)
	defer tFin()

	outArg, outFin := out.wrap()
	defer outFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.Plot_ColormapSlider(labelArg, tArg, outArg, formatArg, C.ImPlotColormap(cmap)) == C.bool(true)
}

func PlotRect_Max(pOut ImPlotPoint, self ImPlotRect) {
	C.PlotRect_Max(pOut.handle(), self.handle())
}

func NewImPlotStyle() ImPlotStyle {
	return (ImPlotStyle)(unsafe.Pointer(C.PlotStyle_ImPlotStyle()))
}

// Plot_BeginDragDropSourceAxis parameter default value hint:
// flags: 0
func Plot_BeginDragDropSourceAxis(axis ImAxis, flags ImGuiDragDropFlags) bool {
	return C.Plot_BeginDragDropSourceAxis(C.ImAxis(axis), C.ImGuiDragDropFlags(flags)) == C.bool(true)
}

func Plot_GetMarkerName(idx ImPlotMarker) string {
	return C.GoString(C.Plot_GetMarkerName(C.ImPlotMarker(idx)))
}

// Plot_PlotStems_FloatPtrInt parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(float)
func Plot_PlotStems_FloatPtrInt(label_id string, values *float32, count int32, ref float64, scale float64, start float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotStems_FloatPtrInt(label_idArg, valuesArg, C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_S64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(ImS64)
func Plot_PlotStems_S64PtrInt(label_id string, values []int64, count int32, ref float64, scale float64, start float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_U64PtrInt parameter default value hint:
// offset: 0
// ref: 0
// scale: 1
// start: 0
// stride: sizeof(ImU64)
// flags: 0
func Plot_PlotStems_U64PtrInt(label_id string, values []uint64, count int32, ref float64, scale float64, start float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(ref), C.double(scale), C.double(start), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_FloatPtrFloatPtr parameter default value hint:
// stride: sizeof(float)
// flags: 0
// offset: 0
// ref: 0
func Plot_PlotStems_FloatPtrFloatPtr(label_id string, xs *float32, ys *float32, count int32, ref float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotStems_FloatPtrFloatPtr(label_idArg, xsArg, ysArg, C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_S64PtrS64Ptr parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(ImS64)
func Plot_PlotStems_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, ref float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotStems_U64PtrU64Ptr parameter default value hint:
// flags: 0
// offset: 0
// ref: 0
// stride: sizeof(ImU64)
func Plot_PlotStems_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, ref float64, flags ImPlotStemsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotStems_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(ref), C.ImPlotStemsFlags(flags), C.int(offset), C.int(stride))
}

func (self ImPlotItemGroup) Destroy() {
	C.PlotItemGroup_Destroy(self.handle())
}

func Plot_GetInputMap() ImPlotInputMap {
	return (ImPlotInputMap)(unsafe.Pointer(C.Plot_GetInputMap()))
}

func Plot_ItemIcon_Vec4(col ImVec4) {
	C.Plot_ItemIcon_Vec4(col.toC())
}

func Plot_ItemIcon_U32(col uint32) {
	C.Plot_ItemIcon_U32(C.ImU32(col))
}

// Plot_PlotBars_FloatPtrInt parameter default value hint:
// offset: 0
// shift: 0
// stride: sizeof(float)
// bar_size: 0.67
// flags: 0
func Plot_PlotBars_FloatPtrInt(label_id string, values *float32, count int32, bar_size float64, shift float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	valuesArg, valuesFin := wrapFloat(values)
	defer valuesFin()

	C.Plot_PlotBars_FloatPtrInt(label_idArg, valuesArg, C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_S64PtrInt parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(ImS64)
func Plot_PlotBars_S64PtrInt(label_id string, values []int64, count int32, bar_size float64, shift float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_S64PtrInt(label_idArg, (*C.longlong)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_U64PtrInt parameter default value hint:
// bar_size: 0.67
// flags: 0
// offset: 0
// shift: 0
// stride: sizeof(ImU64)
func Plot_PlotBars_U64PtrInt(label_id string, values []uint64, count int32, bar_size float64, shift float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_U64PtrInt(label_idArg, (*C.ulonglong)(&(values[0])), C.int(count), C.double(bar_size), C.double(shift), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_FloatPtrFloatPtr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotBars_FloatPtrFloatPtr(label_id string, xs *float32, ys *float32, count int32, bar_size float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	C.Plot_PlotBars_FloatPtrFloatPtr(label_idArg, xsArg, ysArg, C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_S64PtrS64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotBars_S64PtrS64Ptr(label_id string, xs []int64, ys []int64, count int32, bar_size float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_S64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotBars_U64PtrU64Ptr parameter default value hint:
// stride: sizeof(ImU64)
// flags: 0
// offset: 0
func Plot_PlotBars_U64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, count int32, bar_size float64, flags ImPlotBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotBars_U64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), C.int(count), C.double(bar_size), C.ImPlotBarsFlags(flags), C.int(offset), C.int(stride))
}

func Plot_SetupAxisFormat_Str(axis ImAxis, fmt string) {
	fmtArg, fmtFin := wrapString(fmt)
	defer fmtFin()

	C.Plot_SetupAxisFormat_Str(C.ImAxis(axis), fmtArg)
}

// Plot_DestroyContext parameter default value hint:
// ctx: ((void*)0)
func Plot_DestroyContext(ctx ImPlotContext) {
	C.Plot_DestroyContext(ctx.handle())
}

// Plot_DragLineX parameter default value hint:
// flags: 0
// thickness: 1
func Plot_DragLineX(id int32, x *float64, col ImVec4, thickness float32, flags ImPlotDragToolFlags) bool {
	return C.Plot_DragLineX(C.int(id), (*C.double)(x), col.toC(), C.float(thickness), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// Plot_SetupMouseText parameter default value hint:
// flags: 0
func Plot_SetupMouseText(location ImPlotLocation, flags ImPlotMouseTextFlags) {
	C.Plot_SetupMouseText(C.ImPlotLocation(location), C.ImPlotMouseTextFlags(flags))
}

// Plot_ShowDemoWindow parameter default value hint:
// p_open: ((void*)0)
func Plot_ShowDemoWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.Plot_ShowDemoWindow(p_openArg)
}

func Plot_ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Plot_ShowStyleSelector(labelArg) == C.bool(true)
}

// Plot_DragRect parameter default value hint:
// flags: 0
func Plot_DragRect(id int32, x_min *float64, y_min *float64, x_max *float64, y_max *float64, col ImVec4, flags ImPlotDragToolFlags) bool {
	return C.Plot_DragRect(C.int(id), (*C.double)(x_min), (*C.double)(y_min), (*C.double)(x_max), (*C.double)(y_max), col.toC(), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

// Plot_PopStyleColor parameter default value hint:
// count: 1
func Plot_PopStyleColor(count int32) {
	C.Plot_PopStyleColor(C.int(count))
}

func Plot_SetupAxisZoomConstraints(axis ImAxis, z_min float64, z_max float64) {
	C.Plot_SetupAxisZoomConstraints(C.ImAxis(axis), C.double(z_min), C.double(z_max))
}

// Plot_StyleColorsDark parameter default value hint:
// dst: ((void*)0)
func Plot_StyleColorsDark(dst ImPlotStyle) {
	C.Plot_StyleColorsDark(dst.handle())
}

func Plot_NextColormapColor(pOut *ImVec4) {
	pOutArg, pOutFin := pOut.wrap()
	defer pOutFin()

	C.Plot_NextColormapColor(pOutArg)
}

// Plot_DragPoint parameter default value hint:
// flags: 0
// size: 4
func Plot_DragPoint(id int32, x *float64, y *float64, col ImVec4, size float32, flags ImPlotDragToolFlags) bool {
	return C.Plot_DragPoint(C.int(id), (*C.double)(x), (*C.double)(y), col.toC(), C.float(size), C.ImPlotDragToolFlags(flags)) == C.bool(true)
}

func (self ImPlotPoint) Destroy() {
	C.PlotPoint_Destroy(self.handle())
}

// Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(float)
func Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_id string, xs *float32, ys *float32, err *float32, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	errArg, errFin := wrapFloat(err)
	defer errFin()

	C.Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_idArg, xsArg, ysArg, errArg, C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_S64PtrS64PtrS64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_id string, xs []int64, ys []int64, err []int64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_S64PtrS64PtrS64PtrInt(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_U64PtrU64PtrU64PtrInt parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImU64)
func Plot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_id string, xs []uint64, ys []uint64, err []uint64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_U64PtrU64PtrU64PtrInt(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(err[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr parameter default value hint:
// offset: 0
// stride: sizeof(float)
// flags: 0
func Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_id string, xs *float32, ys *float32, neg *float32, pos *float32, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	xsArg, xsFin := wrapFloat(xs)
	defer xsFin()

	ysArg, ysFin := wrapFloat(ys)
	defer ysFin()

	negArg, negFin := wrapFloat(neg)
	defer negFin()

	posArg, posFin := wrapFloat(pos)
	defer posFin()

	C.Plot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_idArg, xsArg, ysArg, negArg, posArg, C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr parameter default value hint:
// flags: 0
// offset: 0
// stride: sizeof(ImS64)
func Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_id string, xs []int64, ys []int64, neg []int64, pos []int64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_S64PtrS64PtrS64PtrS64Ptr(label_idArg, (*C.longlong)(&(xs[0])), (*C.longlong)(&(ys[0])), (*C.longlong)(&(neg[0])), (*C.longlong)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
}

// Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr parameter default value hint:
// stride: sizeof(ImU64)
// flags: 0
// offset: 0
func Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_id string, xs []uint64, ys []uint64, neg []uint64, pos []uint64, count int32, flags ImPlotErrorBarsFlags, offset int32, stride int32) {
	label_idArg, label_idFin := wrapString(label_id)
	defer label_idFin()

	C.Plot_PlotErrorBars_U64PtrU64PtrU64PtrU64Ptr(label_idArg, (*C.ulonglong)(&(xs[0])), (*C.ulonglong)(&(ys[0])), (*C.ulonglong)(&(neg[0])), (*C.ulonglong)(&(pos[0])), C.int(count), C.ImPlotErrorBarsFlags(flags), C.int(offset), C.int(stride))
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

func (self ImPlotTime) SetUs(v int32) {
	C.ImPlotTime_SetUs(self.handle(), C.int(v))
}

func (self ImPlotTime) GetUs() int {
	return int(C.ImPlotTime_GetUs(self.handle()))
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

func (self ImPlotInputMap) SetPan(v ImGuiMouseButton) {
	C.ImPlotInputMap_SetPan(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetPan() ImGuiMouseButton {
	return ImGuiMouseButton(C.ImPlotInputMap_GetPan(self.handle()))
}

func (self ImPlotInputMap) SetPanMod(v ImGuiModFlags) {
	C.ImPlotInputMap_SetPanMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetPanMod() ImGuiModFlags {
	return ImGuiModFlags(C.ImPlotInputMap_GetPanMod(self.handle()))
}

func (self ImPlotInputMap) SetFit(v ImGuiMouseButton) {
	C.ImPlotInputMap_SetFit(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetFit() ImGuiMouseButton {
	return ImGuiMouseButton(C.ImPlotInputMap_GetFit(self.handle()))
}

func (self ImPlotInputMap) SetSelect(v ImGuiMouseButton) {
	C.ImPlotInputMap_SetSelect(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetSelect() ImGuiMouseButton {
	return ImGuiMouseButton(C.ImPlotInputMap_GetSelect(self.handle()))
}

func (self ImPlotInputMap) SetSelectCancel(v ImGuiMouseButton) {
	C.ImPlotInputMap_SetSelectCancel(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetSelectCancel() ImGuiMouseButton {
	return ImGuiMouseButton(C.ImPlotInputMap_GetSelectCancel(self.handle()))
}

func (self ImPlotInputMap) SetSelectMod(v ImGuiModFlags) {
	C.ImPlotInputMap_SetSelectMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetSelectMod() ImGuiModFlags {
	return ImGuiModFlags(C.ImPlotInputMap_GetSelectMod(self.handle()))
}

func (self ImPlotInputMap) SetSelectHorzMod(v ImGuiModFlags) {
	C.ImPlotInputMap_SetSelectHorzMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetSelectHorzMod() ImGuiModFlags {
	return ImGuiModFlags(C.ImPlotInputMap_GetSelectHorzMod(self.handle()))
}

func (self ImPlotInputMap) SetSelectVertMod(v ImGuiModFlags) {
	C.ImPlotInputMap_SetSelectVertMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetSelectVertMod() ImGuiModFlags {
	return ImGuiModFlags(C.ImPlotInputMap_GetSelectVertMod(self.handle()))
}

func (self ImPlotInputMap) SetMenu(v ImGuiMouseButton) {
	C.ImPlotInputMap_SetMenu(self.handle(), C.ImGuiMouseButton(v))
}

func (self ImPlotInputMap) GetMenu() ImGuiMouseButton {
	return ImGuiMouseButton(C.ImPlotInputMap_GetMenu(self.handle()))
}

func (self ImPlotInputMap) SetOverrideMod(v ImGuiModFlags) {
	C.ImPlotInputMap_SetOverrideMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetOverrideMod() ImGuiModFlags {
	return ImGuiModFlags(C.ImPlotInputMap_GetOverrideMod(self.handle()))
}

func (self ImPlotInputMap) SetZoomMod(v ImGuiModFlags) {
	C.ImPlotInputMap_SetZoomMod(self.handle(), C.ImGuiModFlags(v))
}

func (self ImPlotInputMap) GetZoomMod() ImGuiModFlags {
	return ImGuiModFlags(C.ImPlotInputMap_GetZoomMod(self.handle()))
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

func (self ImPlotPoint) Setx(v float64) {
	C.ImPlotPoint_Setx(self.handle(), C.double(v))
}

func (self ImPlotPoint) Getx() float64 {
	return float64(C.ImPlotPoint_Getx(self.handle()))
}

func (self ImPlotPoint) Sety(v float64) {
	C.ImPlotPoint_Sety(self.handle(), C.double(v))
}

func (self ImPlotPoint) Gety() float64 {
	return float64(C.ImPlotPoint_Gety(self.handle()))
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
