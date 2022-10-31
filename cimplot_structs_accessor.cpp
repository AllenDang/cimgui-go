
#include "cimplot_wrapper.h"
#include "cimplot_structs_accessor.h"

void Formatter_Time_Data_SetTime(Formatter_Time_Data *Formatter_Time_DataPtr, ImPlotTime v) { Formatter_Time_DataPtr->Time = v; }
ImPlotTime Formatter_Time_Data_GetTime(Formatter_Time_Data *self) { return self->Time; }
void Formatter_Time_Data_SetSpec(Formatter_Time_Data *Formatter_Time_DataPtr, ImPlotDateTimeSpec v) { Formatter_Time_DataPtr->Spec = v; }
ImPlotDateTimeSpec Formatter_Time_Data_GetSpec(Formatter_Time_Data *self) { return self->Spec; }
void Formatter_Time_Data_SetUserFormatter(Formatter_Time_Data *Formatter_Time_DataPtr, ImPlotFormatter v) { Formatter_Time_DataPtr->UserFormatter = v; }
ImPlotFormatter Formatter_Time_Data_GetUserFormatter(Formatter_Time_Data *self) { return self->UserFormatter; }
void Formatter_Time_Data_SetUserFormatterData(Formatter_Time_Data *Formatter_Time_DataPtr, void* v) { Formatter_Time_DataPtr->UserFormatterData = v; }
void* Formatter_Time_Data_GetUserFormatterData(Formatter_Time_Data *self) { return self->UserFormatterData; }
void ImPlotAlignmentData_SetVertical(ImPlotAlignmentData *ImPlotAlignmentDataPtr, bool v) { ImPlotAlignmentDataPtr->Vertical = v; }
bool ImPlotAlignmentData_GetVertical(ImPlotAlignmentData *self) { return self->Vertical; }
void ImPlotAlignmentData_SetPadA(ImPlotAlignmentData *ImPlotAlignmentDataPtr, float v) { ImPlotAlignmentDataPtr->PadA = v; }
float ImPlotAlignmentData_GetPadA(ImPlotAlignmentData *self) { return self->PadA; }
void ImPlotAlignmentData_SetPadB(ImPlotAlignmentData *ImPlotAlignmentDataPtr, float v) { ImPlotAlignmentDataPtr->PadB = v; }
float ImPlotAlignmentData_GetPadB(ImPlotAlignmentData *self) { return self->PadB; }
void ImPlotAlignmentData_SetPadAMax(ImPlotAlignmentData *ImPlotAlignmentDataPtr, float v) { ImPlotAlignmentDataPtr->PadAMax = v; }
float ImPlotAlignmentData_GetPadAMax(ImPlotAlignmentData *self) { return self->PadAMax; }
void ImPlotAlignmentData_SetPadBMax(ImPlotAlignmentData *ImPlotAlignmentDataPtr, float v) { ImPlotAlignmentDataPtr->PadBMax = v; }
float ImPlotAlignmentData_GetPadBMax(ImPlotAlignmentData *self) { return self->PadBMax; }
void ImPlotAnnotation_SetPos(ImPlotAnnotation *ImPlotAnnotationPtr, ImVec2 v) { ImPlotAnnotationPtr->Pos = v; }
ImVec2 ImPlotAnnotation_GetPos(ImPlotAnnotation *self) { return self->Pos; }
void ImPlotAnnotation_SetOffset(ImPlotAnnotation *ImPlotAnnotationPtr, ImVec2 v) { ImPlotAnnotationPtr->Offset = v; }
ImVec2 ImPlotAnnotation_GetOffset(ImPlotAnnotation *self) { return self->Offset; }
void ImPlotAnnotation_SetColorBg(ImPlotAnnotation *ImPlotAnnotationPtr, ImU32 v) { ImPlotAnnotationPtr->ColorBg = v; }
ImU32 ImPlotAnnotation_GetColorBg(ImPlotAnnotation *self) { return self->ColorBg; }
void ImPlotAnnotation_SetColorFg(ImPlotAnnotation *ImPlotAnnotationPtr, ImU32 v) { ImPlotAnnotationPtr->ColorFg = v; }
ImU32 ImPlotAnnotation_GetColorFg(ImPlotAnnotation *self) { return self->ColorFg; }
void ImPlotAnnotation_SetTextOffset(ImPlotAnnotation *ImPlotAnnotationPtr, int v) { ImPlotAnnotationPtr->TextOffset = v; }
int ImPlotAnnotation_GetTextOffset(ImPlotAnnotation *self) { return self->TextOffset; }
void ImPlotAnnotation_SetClamp(ImPlotAnnotation *ImPlotAnnotationPtr, bool v) { ImPlotAnnotationPtr->Clamp = v; }
bool ImPlotAnnotation_GetClamp(ImPlotAnnotation *self) { return self->Clamp; }
void ImPlotAnnotationCollection_SetAnnotations(ImPlotAnnotationCollection *ImPlotAnnotationCollectionPtr, ImVector_ImPlotAnnotation v) { ImPlotAnnotationCollectionPtr->Annotations = v; }
ImVector_ImPlotAnnotation ImPlotAnnotationCollection_GetAnnotations(ImPlotAnnotationCollection *self) { return self->Annotations; }
void ImPlotAnnotationCollection_SetTextBuffer(ImPlotAnnotationCollection *ImPlotAnnotationCollectionPtr, ImGuiTextBuffer v) { ImPlotAnnotationCollectionPtr->TextBuffer = v; }
ImGuiTextBuffer ImPlotAnnotationCollection_GetTextBuffer(ImPlotAnnotationCollection *self) { return self->TextBuffer; }
void ImPlotAnnotationCollection_SetSize(ImPlotAnnotationCollection *ImPlotAnnotationCollectionPtr, int v) { ImPlotAnnotationCollectionPtr->Size = v; }
int ImPlotAnnotationCollection_GetSize(ImPlotAnnotationCollection *self) { return self->Size; }
void ImPlotAxis_SetID(ImPlotAxis *ImPlotAxisPtr, ImGuiID v) { ImPlotAxisPtr->ID = v; }
ImGuiID ImPlotAxis_GetID(ImPlotAxis *self) { return self->ID; }
void ImPlotAxis_SetFlags(ImPlotAxis *ImPlotAxisPtr, ImPlotAxisFlags v) { ImPlotAxisPtr->Flags = v; }
ImPlotAxisFlags ImPlotAxis_GetFlags(ImPlotAxis *self) { return self->Flags; }
void ImPlotAxis_SetPreviousFlags(ImPlotAxis *ImPlotAxisPtr, ImPlotAxisFlags v) { ImPlotAxisPtr->PreviousFlags = v; }
ImPlotAxisFlags ImPlotAxis_GetPreviousFlags(ImPlotAxis *self) { return self->PreviousFlags; }
void ImPlotAxis_SetRange(ImPlotAxis *ImPlotAxisPtr, ImPlotRange v) { ImPlotAxisPtr->Range = v; }
ImPlotRange ImPlotAxis_GetRange(ImPlotAxis *self) { return self->Range; }
void ImPlotAxis_SetRangeCond(ImPlotAxis *ImPlotAxisPtr, ImPlotCond v) { ImPlotAxisPtr->RangeCond = v; }
ImPlotCond ImPlotAxis_GetRangeCond(ImPlotAxis *self) { return self->RangeCond; }
void ImPlotAxis_SetScale(ImPlotAxis *ImPlotAxisPtr, ImPlotScale v) { ImPlotAxisPtr->Scale = v; }
ImPlotScale ImPlotAxis_GetScale(ImPlotAxis *self) { return self->Scale; }
void ImPlotAxis_SetFitExtents(ImPlotAxis *ImPlotAxisPtr, ImPlotRange v) { ImPlotAxisPtr->FitExtents = v; }
ImPlotRange ImPlotAxis_GetFitExtents(ImPlotAxis *self) { return self->FitExtents; }
void ImPlotAxis_SetOrthoAxis(ImPlotAxis *ImPlotAxisPtr, ImPlotAxis* v) { ImPlotAxisPtr->OrthoAxis = v; }
ImPlotAxis* ImPlotAxis_GetOrthoAxis(ImPlotAxis *self) { return self->OrthoAxis; }
void ImPlotAxis_SetConstraintRange(ImPlotAxis *ImPlotAxisPtr, ImPlotRange v) { ImPlotAxisPtr->ConstraintRange = v; }
ImPlotRange ImPlotAxis_GetConstraintRange(ImPlotAxis *self) { return self->ConstraintRange; }
void ImPlotAxis_SetConstraintZoom(ImPlotAxis *ImPlotAxisPtr, ImPlotRange v) { ImPlotAxisPtr->ConstraintZoom = v; }
ImPlotRange ImPlotAxis_GetConstraintZoom(ImPlotAxis *self) { return self->ConstraintZoom; }
void ImPlotAxis_SetTicker(ImPlotAxis *ImPlotAxisPtr, ImPlotTicker v) { ImPlotAxisPtr->Ticker = v; }
ImPlotTicker ImPlotAxis_GetTicker(ImPlotAxis *self) { return self->Ticker; }
void ImPlotAxis_SetFormatter(ImPlotAxis *ImPlotAxisPtr, ImPlotFormatter v) { ImPlotAxisPtr->Formatter = v; }
ImPlotFormatter ImPlotAxis_GetFormatter(ImPlotAxis *self) { return self->Formatter; }
void ImPlotAxis_SetFormatterData(ImPlotAxis *ImPlotAxisPtr, void* v) { ImPlotAxisPtr->FormatterData = v; }
void* ImPlotAxis_GetFormatterData(ImPlotAxis *self) { return self->FormatterData; }
void ImPlotAxis_SetLocator(ImPlotAxis *ImPlotAxisPtr, ImPlotLocator v) { ImPlotAxisPtr->Locator = v; }
ImPlotLocator ImPlotAxis_GetLocator(ImPlotAxis *self) { return self->Locator; }
void ImPlotAxis_SetLinkedMin(ImPlotAxis *ImPlotAxisPtr, double* v) { ImPlotAxisPtr->LinkedMin = v; }
double* ImPlotAxis_GetLinkedMin(ImPlotAxis *self) { return self->LinkedMin; }
void ImPlotAxis_SetLinkedMax(ImPlotAxis *ImPlotAxisPtr, double* v) { ImPlotAxisPtr->LinkedMax = v; }
double* ImPlotAxis_GetLinkedMax(ImPlotAxis *self) { return self->LinkedMax; }
void ImPlotAxis_SetPickerLevel(ImPlotAxis *ImPlotAxisPtr, int v) { ImPlotAxisPtr->PickerLevel = v; }
int ImPlotAxis_GetPickerLevel(ImPlotAxis *self) { return self->PickerLevel; }
void ImPlotAxis_SetPickerTimeMin(ImPlotAxis *ImPlotAxisPtr, ImPlotTime v) { ImPlotAxisPtr->PickerTimeMin = v; }
ImPlotTime ImPlotAxis_GetPickerTimeMin(ImPlotAxis *self) { return self->PickerTimeMin; }
void ImPlotAxis_SetPickerTimeMax(ImPlotAxis *ImPlotAxisPtr, ImPlotTime v) { ImPlotAxisPtr->PickerTimeMax = v; }
ImPlotTime ImPlotAxis_GetPickerTimeMax(ImPlotAxis *self) { return self->PickerTimeMax; }
void ImPlotAxis_SetTransformForward(ImPlotAxis *ImPlotAxisPtr, ImPlotTransform v) { ImPlotAxisPtr->TransformForward = v; }
ImPlotTransform ImPlotAxis_GetTransformForward(ImPlotAxis *self) { return self->TransformForward; }
void ImPlotAxis_SetTransformInverse(ImPlotAxis *ImPlotAxisPtr, ImPlotTransform v) { ImPlotAxisPtr->TransformInverse = v; }
ImPlotTransform ImPlotAxis_GetTransformInverse(ImPlotAxis *self) { return self->TransformInverse; }
void ImPlotAxis_SetTransformData(ImPlotAxis *ImPlotAxisPtr, void* v) { ImPlotAxisPtr->TransformData = v; }
void* ImPlotAxis_GetTransformData(ImPlotAxis *self) { return self->TransformData; }
void ImPlotAxis_SetPixelMin(ImPlotAxis *ImPlotAxisPtr, float v) { ImPlotAxisPtr->PixelMin = v; }
float ImPlotAxis_GetPixelMin(ImPlotAxis *self) { return self->PixelMin; }
void ImPlotAxis_SetPixelMax(ImPlotAxis *ImPlotAxisPtr, float v) { ImPlotAxisPtr->PixelMax = v; }
float ImPlotAxis_GetPixelMax(ImPlotAxis *self) { return self->PixelMax; }
void ImPlotAxis_SetScaleMin(ImPlotAxis *ImPlotAxisPtr, double v) { ImPlotAxisPtr->ScaleMin = v; }
double ImPlotAxis_GetScaleMin(ImPlotAxis *self) { return self->ScaleMin; }
void ImPlotAxis_SetScaleMax(ImPlotAxis *ImPlotAxisPtr, double v) { ImPlotAxisPtr->ScaleMax = v; }
double ImPlotAxis_GetScaleMax(ImPlotAxis *self) { return self->ScaleMax; }
void ImPlotAxis_SetScaleToPixel(ImPlotAxis *ImPlotAxisPtr, double v) { ImPlotAxisPtr->ScaleToPixel = v; }
double ImPlotAxis_GetScaleToPixel(ImPlotAxis *self) { return self->ScaleToPixel; }
void ImPlotAxis_SetDatum1(ImPlotAxis *ImPlotAxisPtr, float v) { ImPlotAxisPtr->Datum1 = v; }
float ImPlotAxis_GetDatum1(ImPlotAxis *self) { return self->Datum1; }
void ImPlotAxis_SetDatum2(ImPlotAxis *ImPlotAxisPtr, float v) { ImPlotAxisPtr->Datum2 = v; }
float ImPlotAxis_GetDatum2(ImPlotAxis *self) { return self->Datum2; }
void ImPlotAxis_SetHoverRect(ImPlotAxis *ImPlotAxisPtr, ImRect v) { ImPlotAxisPtr->HoverRect = v; }
ImRect ImPlotAxis_GetHoverRect(ImPlotAxis *self) { return self->HoverRect; }
void ImPlotAxis_SetLabelOffset(ImPlotAxis *ImPlotAxisPtr, int v) { ImPlotAxisPtr->LabelOffset = v; }
int ImPlotAxis_GetLabelOffset(ImPlotAxis *self) { return self->LabelOffset; }
void ImPlotAxis_SetColorMaj(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorMaj = v; }
ImU32 ImPlotAxis_GetColorMaj(ImPlotAxis *self) { return self->ColorMaj; }
void ImPlotAxis_SetColorMin(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorMin = v; }
ImU32 ImPlotAxis_GetColorMin(ImPlotAxis *self) { return self->ColorMin; }
void ImPlotAxis_SetColorTick(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorTick = v; }
ImU32 ImPlotAxis_GetColorTick(ImPlotAxis *self) { return self->ColorTick; }
void ImPlotAxis_SetColorTxt(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorTxt = v; }
ImU32 ImPlotAxis_GetColorTxt(ImPlotAxis *self) { return self->ColorTxt; }
void ImPlotAxis_SetColorBg(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorBg = v; }
ImU32 ImPlotAxis_GetColorBg(ImPlotAxis *self) { return self->ColorBg; }
void ImPlotAxis_SetColorHov(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorHov = v; }
ImU32 ImPlotAxis_GetColorHov(ImPlotAxis *self) { return self->ColorHov; }
void ImPlotAxis_SetColorAct(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorAct = v; }
ImU32 ImPlotAxis_GetColorAct(ImPlotAxis *self) { return self->ColorAct; }
void ImPlotAxis_SetColorHiLi(ImPlotAxis *ImPlotAxisPtr, ImU32 v) { ImPlotAxisPtr->ColorHiLi = v; }
ImU32 ImPlotAxis_GetColorHiLi(ImPlotAxis *self) { return self->ColorHiLi; }
void ImPlotAxis_SetEnabled(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->Enabled = v; }
bool ImPlotAxis_GetEnabled(ImPlotAxis *self) { return self->Enabled; }
void ImPlotAxis_SetVertical(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->Vertical = v; }
bool ImPlotAxis_GetVertical(ImPlotAxis *self) { return self->Vertical; }
void ImPlotAxis_SetFitThisFrame(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->FitThisFrame = v; }
bool ImPlotAxis_GetFitThisFrame(ImPlotAxis *self) { return self->FitThisFrame; }
void ImPlotAxis_SetHasRange(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->HasRange = v; }
bool ImPlotAxis_GetHasRange(ImPlotAxis *self) { return self->HasRange; }
void ImPlotAxis_SetHasFormatSpec(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->HasFormatSpec = v; }
bool ImPlotAxis_GetHasFormatSpec(ImPlotAxis *self) { return self->HasFormatSpec; }
void ImPlotAxis_SetShowDefaultTicks(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->ShowDefaultTicks = v; }
bool ImPlotAxis_GetShowDefaultTicks(ImPlotAxis *self) { return self->ShowDefaultTicks; }
void ImPlotAxis_SetHovered(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->Hovered = v; }
bool ImPlotAxis_GetHovered(ImPlotAxis *self) { return self->Hovered; }
void ImPlotAxis_SetHeld(ImPlotAxis *ImPlotAxisPtr, bool v) { ImPlotAxisPtr->Held = v; }
bool ImPlotAxis_GetHeld(ImPlotAxis *self) { return self->Held; }
void ImPlotColormapData_SetKeyCounts(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_int v) { ImPlotColormapDataPtr->KeyCounts = v; }
ImVector_int ImPlotColormapData_GetKeyCounts(ImPlotColormapData *self) { return self->KeyCounts; }
void ImPlotColormapData_SetKeyOffsets(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_int v) { ImPlotColormapDataPtr->KeyOffsets = v; }
ImVector_int ImPlotColormapData_GetKeyOffsets(ImPlotColormapData *self) { return self->KeyOffsets; }
void ImPlotColormapData_SetTables(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_ImU32 v) { ImPlotColormapDataPtr->Tables = v; }
ImVector_ImU32 ImPlotColormapData_GetTables(ImPlotColormapData *self) { return self->Tables; }
void ImPlotColormapData_SetTableSizes(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_int v) { ImPlotColormapDataPtr->TableSizes = v; }
ImVector_int ImPlotColormapData_GetTableSizes(ImPlotColormapData *self) { return self->TableSizes; }
void ImPlotColormapData_SetTableOffsets(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_int v) { ImPlotColormapDataPtr->TableOffsets = v; }
ImVector_int ImPlotColormapData_GetTableOffsets(ImPlotColormapData *self) { return self->TableOffsets; }
void ImPlotColormapData_SetText(ImPlotColormapData *ImPlotColormapDataPtr, ImGuiTextBuffer v) { ImPlotColormapDataPtr->Text = v; }
ImGuiTextBuffer ImPlotColormapData_GetText(ImPlotColormapData *self) { return self->Text; }
void ImPlotColormapData_SetTextOffsets(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_int v) { ImPlotColormapDataPtr->TextOffsets = v; }
ImVector_int ImPlotColormapData_GetTextOffsets(ImPlotColormapData *self) { return self->TextOffsets; }
void ImPlotColormapData_SetQuals(ImPlotColormapData *ImPlotColormapDataPtr, ImVector_bool v) { ImPlotColormapDataPtr->Quals = v; }
ImVector_bool ImPlotColormapData_GetQuals(ImPlotColormapData *self) { return self->Quals; }
void ImPlotColormapData_SetMap(ImPlotColormapData *ImPlotColormapDataPtr, ImGuiStorage v) { ImPlotColormapDataPtr->Map = v; }
ImGuiStorage ImPlotColormapData_GetMap(ImPlotColormapData *self) { return self->Map; }
void ImPlotColormapData_SetCount(ImPlotColormapData *ImPlotColormapDataPtr, int v) { ImPlotColormapDataPtr->Count = v; }
int ImPlotColormapData_GetCount(ImPlotColormapData *self) { return self->Count; }
void ImPlotContext_SetPlots(ImPlotContext *ImPlotContextPtr, ImPool_ImPlotPlot v) { ImPlotContextPtr->Plots = v; }
ImPool_ImPlotPlot ImPlotContext_GetPlots(ImPlotContext *self) { return self->Plots; }
void ImPlotContext_SetSubplots(ImPlotContext *ImPlotContextPtr, ImPool_ImPlotSubplot v) { ImPlotContextPtr->Subplots = v; }
ImPool_ImPlotSubplot ImPlotContext_GetSubplots(ImPlotContext *self) { return self->Subplots; }
void ImPlotContext_SetCurrentPlot(ImPlotContext *ImPlotContextPtr, ImPlotPlot* v) { ImPlotContextPtr->CurrentPlot = v; }
ImPlotPlot* ImPlotContext_GetCurrentPlot(ImPlotContext *self) { return self->CurrentPlot; }
void ImPlotContext_SetCurrentSubplot(ImPlotContext *ImPlotContextPtr, ImPlotSubplot* v) { ImPlotContextPtr->CurrentSubplot = v; }
ImPlotSubplot* ImPlotContext_GetCurrentSubplot(ImPlotContext *self) { return self->CurrentSubplot; }
void ImPlotContext_SetCurrentItems(ImPlotContext *ImPlotContextPtr, ImPlotItemGroup* v) { ImPlotContextPtr->CurrentItems = v; }
ImPlotItemGroup* ImPlotContext_GetCurrentItems(ImPlotContext *self) { return self->CurrentItems; }
void ImPlotContext_SetCurrentItem(ImPlotContext *ImPlotContextPtr, ImPlotItem* v) { ImPlotContextPtr->CurrentItem = v; }
ImPlotItem* ImPlotContext_GetCurrentItem(ImPlotContext *self) { return self->CurrentItem; }
void ImPlotContext_SetPreviousItem(ImPlotContext *ImPlotContextPtr, ImPlotItem* v) { ImPlotContextPtr->PreviousItem = v; }
ImPlotItem* ImPlotContext_GetPreviousItem(ImPlotContext *self) { return self->PreviousItem; }
void ImPlotContext_SetCTicker(ImPlotContext *ImPlotContextPtr, ImPlotTicker v) { ImPlotContextPtr->CTicker = v; }
ImPlotTicker ImPlotContext_GetCTicker(ImPlotContext *self) { return self->CTicker; }
void ImPlotContext_SetAnnotations(ImPlotContext *ImPlotContextPtr, ImPlotAnnotationCollection v) { ImPlotContextPtr->Annotations = v; }
ImPlotAnnotationCollection ImPlotContext_GetAnnotations(ImPlotContext *self) { return self->Annotations; }
void ImPlotContext_SetTags(ImPlotContext *ImPlotContextPtr, ImPlotTagCollection v) { ImPlotContextPtr->Tags = v; }
ImPlotTagCollection ImPlotContext_GetTags(ImPlotContext *self) { return self->Tags; }
void ImPlotContext_SetChildWindowMade(ImPlotContext *ImPlotContextPtr, bool v) { ImPlotContextPtr->ChildWindowMade = v; }
bool ImPlotContext_GetChildWindowMade(ImPlotContext *self) { return self->ChildWindowMade; }
void ImPlotContext_SetStyle(ImPlotContext *ImPlotContextPtr, ImPlotStyle v) { ImPlotContextPtr->Style = v; }
ImPlotStyle ImPlotContext_GetStyle(ImPlotContext *self) { return self->Style; }
void ImPlotContext_SetColorModifiers(ImPlotContext *ImPlotContextPtr, ImVector_ImGuiColorMod v) { ImPlotContextPtr->ColorModifiers = v; }
ImVector_ImGuiColorMod ImPlotContext_GetColorModifiers(ImPlotContext *self) { return self->ColorModifiers; }
void ImPlotContext_SetStyleModifiers(ImPlotContext *ImPlotContextPtr, ImVector_ImGuiStyleMod v) { ImPlotContextPtr->StyleModifiers = v; }
ImVector_ImGuiStyleMod ImPlotContext_GetStyleModifiers(ImPlotContext *self) { return self->StyleModifiers; }
void ImPlotContext_SetColormapData(ImPlotContext *ImPlotContextPtr, ImPlotColormapData v) { ImPlotContextPtr->ColormapData = v; }
ImPlotColormapData ImPlotContext_GetColormapData(ImPlotContext *self) { return self->ColormapData; }
void ImPlotContext_SetColormapModifiers(ImPlotContext *ImPlotContextPtr, ImVector_ImPlotColormap v) { ImPlotContextPtr->ColormapModifiers = v; }
ImVector_ImPlotColormap ImPlotContext_GetColormapModifiers(ImPlotContext *self) { return self->ColormapModifiers; }
void ImPlotContext_SetTm(ImPlotContext *ImPlotContextPtr, tm v) { ImPlotContextPtr->Tm = v; }
tm ImPlotContext_GetTm(ImPlotContext *self) { return self->Tm; }
void ImPlotContext_SetTempDouble1(ImPlotContext *ImPlotContextPtr, ImVector_double v) { ImPlotContextPtr->TempDouble1 = v; }
ImVector_double ImPlotContext_GetTempDouble1(ImPlotContext *self) { return self->TempDouble1; }
void ImPlotContext_SetTempDouble2(ImPlotContext *ImPlotContextPtr, ImVector_double v) { ImPlotContextPtr->TempDouble2 = v; }
ImVector_double ImPlotContext_GetTempDouble2(ImPlotContext *self) { return self->TempDouble2; }
void ImPlotContext_SetTempInt1(ImPlotContext *ImPlotContextPtr, ImVector_int v) { ImPlotContextPtr->TempInt1 = v; }
ImVector_int ImPlotContext_GetTempInt1(ImPlotContext *self) { return self->TempInt1; }
void ImPlotContext_SetDigitalPlotItemCnt(ImPlotContext *ImPlotContextPtr, int v) { ImPlotContextPtr->DigitalPlotItemCnt = v; }
int ImPlotContext_GetDigitalPlotItemCnt(ImPlotContext *self) { return self->DigitalPlotItemCnt; }
void ImPlotContext_SetDigitalPlotOffset(ImPlotContext *ImPlotContextPtr, int v) { ImPlotContextPtr->DigitalPlotOffset = v; }
int ImPlotContext_GetDigitalPlotOffset(ImPlotContext *self) { return self->DigitalPlotOffset; }
void ImPlotContext_SetNextPlotData(ImPlotContext *ImPlotContextPtr, ImPlotNextPlotData v) { ImPlotContextPtr->NextPlotData = v; }
ImPlotNextPlotData ImPlotContext_GetNextPlotData(ImPlotContext *self) { return self->NextPlotData; }
void ImPlotContext_SetNextItemData(ImPlotContext *ImPlotContextPtr, ImPlotNextItemData v) { ImPlotContextPtr->NextItemData = v; }
ImPlotNextItemData ImPlotContext_GetNextItemData(ImPlotContext *self) { return self->NextItemData; }
void ImPlotContext_SetInputMap(ImPlotContext *ImPlotContextPtr, ImPlotInputMap v) { ImPlotContextPtr->InputMap = v; }
ImPlotInputMap ImPlotContext_GetInputMap(ImPlotContext *self) { return self->InputMap; }
void ImPlotContext_SetOpenContextThisFrame(ImPlotContext *ImPlotContextPtr, bool v) { ImPlotContextPtr->OpenContextThisFrame = v; }
bool ImPlotContext_GetOpenContextThisFrame(ImPlotContext *self) { return self->OpenContextThisFrame; }
void ImPlotContext_SetMousePosStringBuilder(ImPlotContext *ImPlotContextPtr, ImGuiTextBuffer v) { ImPlotContextPtr->MousePosStringBuilder = v; }
ImGuiTextBuffer ImPlotContext_GetMousePosStringBuilder(ImPlotContext *self) { return self->MousePosStringBuilder; }
void ImPlotContext_SetSortItems(ImPlotContext *ImPlotContextPtr, ImPlotItemGroup* v) { ImPlotContextPtr->SortItems = v; }
ImPlotItemGroup* ImPlotContext_GetSortItems(ImPlotContext *self) { return self->SortItems; }
void ImPlotContext_SetAlignmentData(ImPlotContext *ImPlotContextPtr, ImPool_ImPlotAlignmentData v) { ImPlotContextPtr->AlignmentData = v; }
ImPool_ImPlotAlignmentData ImPlotContext_GetAlignmentData(ImPlotContext *self) { return self->AlignmentData; }
void ImPlotContext_SetCurrentAlignmentH(ImPlotContext *ImPlotContextPtr, ImPlotAlignmentData* v) { ImPlotContextPtr->CurrentAlignmentH = v; }
ImPlotAlignmentData* ImPlotContext_GetCurrentAlignmentH(ImPlotContext *self) { return self->CurrentAlignmentH; }
void ImPlotContext_SetCurrentAlignmentV(ImPlotContext *ImPlotContextPtr, ImPlotAlignmentData* v) { ImPlotContextPtr->CurrentAlignmentV = v; }
ImPlotAlignmentData* ImPlotContext_GetCurrentAlignmentV(ImPlotContext *self) { return self->CurrentAlignmentV; }
void ImPlotDateTimeSpec_SetDate(ImPlotDateTimeSpec *ImPlotDateTimeSpecPtr, ImPlotDateFmt v) { ImPlotDateTimeSpecPtr->Date = v; }
ImPlotDateFmt ImPlotDateTimeSpec_GetDate(ImPlotDateTimeSpec *self) { return self->Date; }
void ImPlotDateTimeSpec_SetTime(ImPlotDateTimeSpec *ImPlotDateTimeSpecPtr, ImPlotTimeFmt v) { ImPlotDateTimeSpecPtr->Time = v; }
ImPlotTimeFmt ImPlotDateTimeSpec_GetTime(ImPlotDateTimeSpec *self) { return self->Time; }
void ImPlotDateTimeSpec_SetUseISO8601(ImPlotDateTimeSpec *ImPlotDateTimeSpecPtr, bool v) { ImPlotDateTimeSpecPtr->UseISO8601 = v; }
bool ImPlotDateTimeSpec_GetUseISO8601(ImPlotDateTimeSpec *self) { return self->UseISO8601; }
void ImPlotDateTimeSpec_SetUse24HourClock(ImPlotDateTimeSpec *ImPlotDateTimeSpecPtr, bool v) { ImPlotDateTimeSpecPtr->Use24HourClock = v; }
bool ImPlotDateTimeSpec_GetUse24HourClock(ImPlotDateTimeSpec *self) { return self->Use24HourClock; }
void ImPlotInputMap_SetPan(ImPlotInputMap *ImPlotInputMapPtr, ImGuiMouseButton v) { ImPlotInputMapPtr->Pan = v; }
ImGuiMouseButton ImPlotInputMap_GetPan(ImPlotInputMap *self) { return self->Pan; }
void ImPlotInputMap_SetPanMod(ImPlotInputMap *ImPlotInputMapPtr, int v) { ImPlotInputMapPtr->PanMod = v; }
int ImPlotInputMap_GetPanMod(ImPlotInputMap *self) { return self->PanMod; }
void ImPlotInputMap_SetFit(ImPlotInputMap *ImPlotInputMapPtr, ImGuiMouseButton v) { ImPlotInputMapPtr->Fit = v; }
ImGuiMouseButton ImPlotInputMap_GetFit(ImPlotInputMap *self) { return self->Fit; }
void ImPlotInputMap_SetSelect(ImPlotInputMap *ImPlotInputMapPtr, ImGuiMouseButton v) { ImPlotInputMapPtr->Select = v; }
ImGuiMouseButton ImPlotInputMap_GetSelect(ImPlotInputMap *self) { return self->Select; }
void ImPlotInputMap_SetSelectCancel(ImPlotInputMap *ImPlotInputMapPtr, ImGuiMouseButton v) { ImPlotInputMapPtr->SelectCancel = v; }
ImGuiMouseButton ImPlotInputMap_GetSelectCancel(ImPlotInputMap *self) { return self->SelectCancel; }
void ImPlotInputMap_SetSelectMod(ImPlotInputMap *ImPlotInputMapPtr, int v) { ImPlotInputMapPtr->SelectMod = v; }
int ImPlotInputMap_GetSelectMod(ImPlotInputMap *self) { return self->SelectMod; }
void ImPlotInputMap_SetSelectHorzMod(ImPlotInputMap *ImPlotInputMapPtr, int v) { ImPlotInputMapPtr->SelectHorzMod = v; }
int ImPlotInputMap_GetSelectHorzMod(ImPlotInputMap *self) { return self->SelectHorzMod; }
void ImPlotInputMap_SetSelectVertMod(ImPlotInputMap *ImPlotInputMapPtr, int v) { ImPlotInputMapPtr->SelectVertMod = v; }
int ImPlotInputMap_GetSelectVertMod(ImPlotInputMap *self) { return self->SelectVertMod; }
void ImPlotInputMap_SetMenu(ImPlotInputMap *ImPlotInputMapPtr, ImGuiMouseButton v) { ImPlotInputMapPtr->Menu = v; }
ImGuiMouseButton ImPlotInputMap_GetMenu(ImPlotInputMap *self) { return self->Menu; }
void ImPlotInputMap_SetOverrideMod(ImPlotInputMap *ImPlotInputMapPtr, int v) { ImPlotInputMapPtr->OverrideMod = v; }
int ImPlotInputMap_GetOverrideMod(ImPlotInputMap *self) { return self->OverrideMod; }
void ImPlotInputMap_SetZoomMod(ImPlotInputMap *ImPlotInputMapPtr, int v) { ImPlotInputMapPtr->ZoomMod = v; }
int ImPlotInputMap_GetZoomMod(ImPlotInputMap *self) { return self->ZoomMod; }
void ImPlotInputMap_SetZoomRate(ImPlotInputMap *ImPlotInputMapPtr, float v) { ImPlotInputMapPtr->ZoomRate = v; }
float ImPlotInputMap_GetZoomRate(ImPlotInputMap *self) { return self->ZoomRate; }
void ImPlotItem_SetID(ImPlotItem *ImPlotItemPtr, ImGuiID v) { ImPlotItemPtr->ID = v; }
ImGuiID ImPlotItem_GetID(ImPlotItem *self) { return self->ID; }
void ImPlotItem_SetColor(ImPlotItem *ImPlotItemPtr, ImU32 v) { ImPlotItemPtr->Color = v; }
ImU32 ImPlotItem_GetColor(ImPlotItem *self) { return self->Color; }
void ImPlotItem_SetLegendHoverRect(ImPlotItem *ImPlotItemPtr, ImRect v) { ImPlotItemPtr->LegendHoverRect = v; }
ImRect ImPlotItem_GetLegendHoverRect(ImPlotItem *self) { return self->LegendHoverRect; }
void ImPlotItem_SetNameOffset(ImPlotItem *ImPlotItemPtr, int v) { ImPlotItemPtr->NameOffset = v; }
int ImPlotItem_GetNameOffset(ImPlotItem *self) { return self->NameOffset; }
void ImPlotItem_SetShow(ImPlotItem *ImPlotItemPtr, bool v) { ImPlotItemPtr->Show = v; }
bool ImPlotItem_GetShow(ImPlotItem *self) { return self->Show; }
void ImPlotItem_SetLegendHovered(ImPlotItem *ImPlotItemPtr, bool v) { ImPlotItemPtr->LegendHovered = v; }
bool ImPlotItem_GetLegendHovered(ImPlotItem *self) { return self->LegendHovered; }
void ImPlotItem_SetSeenThisFrame(ImPlotItem *ImPlotItemPtr, bool v) { ImPlotItemPtr->SeenThisFrame = v; }
bool ImPlotItem_GetSeenThisFrame(ImPlotItem *self) { return self->SeenThisFrame; }
void ImPlotItemGroup_SetID(ImPlotItemGroup *ImPlotItemGroupPtr, ImGuiID v) { ImPlotItemGroupPtr->ID = v; }
ImGuiID ImPlotItemGroup_GetID(ImPlotItemGroup *self) { return self->ID; }
void ImPlotItemGroup_SetLegend(ImPlotItemGroup *ImPlotItemGroupPtr, ImPlotLegend v) { ImPlotItemGroupPtr->Legend = v; }
ImPlotLegend ImPlotItemGroup_GetLegend(ImPlotItemGroup *self) { return self->Legend; }
void ImPlotItemGroup_SetItemPool(ImPlotItemGroup *ImPlotItemGroupPtr, ImPool_ImPlotItem v) { ImPlotItemGroupPtr->ItemPool = v; }
ImPool_ImPlotItem ImPlotItemGroup_GetItemPool(ImPlotItemGroup *self) { return self->ItemPool; }
void ImPlotItemGroup_SetColormapIdx(ImPlotItemGroup *ImPlotItemGroupPtr, int v) { ImPlotItemGroupPtr->ColormapIdx = v; }
int ImPlotItemGroup_GetColormapIdx(ImPlotItemGroup *self) { return self->ColormapIdx; }
void ImPlotLegend_SetFlags(ImPlotLegend *ImPlotLegendPtr, ImPlotLegendFlags v) { ImPlotLegendPtr->Flags = v; }
ImPlotLegendFlags ImPlotLegend_GetFlags(ImPlotLegend *self) { return self->Flags; }
void ImPlotLegend_SetPreviousFlags(ImPlotLegend *ImPlotLegendPtr, ImPlotLegendFlags v) { ImPlotLegendPtr->PreviousFlags = v; }
ImPlotLegendFlags ImPlotLegend_GetPreviousFlags(ImPlotLegend *self) { return self->PreviousFlags; }
void ImPlotLegend_SetLocation(ImPlotLegend *ImPlotLegendPtr, ImPlotLocation v) { ImPlotLegendPtr->Location = v; }
ImPlotLocation ImPlotLegend_GetLocation(ImPlotLegend *self) { return self->Location; }
void ImPlotLegend_SetPreviousLocation(ImPlotLegend *ImPlotLegendPtr, ImPlotLocation v) { ImPlotLegendPtr->PreviousLocation = v; }
ImPlotLocation ImPlotLegend_GetPreviousLocation(ImPlotLegend *self) { return self->PreviousLocation; }
void ImPlotLegend_SetIndices(ImPlotLegend *ImPlotLegendPtr, ImVector_int v) { ImPlotLegendPtr->Indices = v; }
ImVector_int ImPlotLegend_GetIndices(ImPlotLegend *self) { return self->Indices; }
void ImPlotLegend_SetLabels(ImPlotLegend *ImPlotLegendPtr, ImGuiTextBuffer v) { ImPlotLegendPtr->Labels = v; }
ImGuiTextBuffer ImPlotLegend_GetLabels(ImPlotLegend *self) { return self->Labels; }
void ImPlotLegend_SetRect(ImPlotLegend *ImPlotLegendPtr, ImRect v) { ImPlotLegendPtr->Rect = v; }
ImRect ImPlotLegend_GetRect(ImPlotLegend *self) { return self->Rect; }
void ImPlotLegend_SetHovered(ImPlotLegend *ImPlotLegendPtr, bool v) { ImPlotLegendPtr->Hovered = v; }
bool ImPlotLegend_GetHovered(ImPlotLegend *self) { return self->Hovered; }
void ImPlotLegend_SetHeld(ImPlotLegend *ImPlotLegendPtr, bool v) { ImPlotLegendPtr->Held = v; }
bool ImPlotLegend_GetHeld(ImPlotLegend *self) { return self->Held; }
void ImPlotLegend_SetCanGoInside(ImPlotLegend *ImPlotLegendPtr, bool v) { ImPlotLegendPtr->CanGoInside = v; }
bool ImPlotLegend_GetCanGoInside(ImPlotLegend *self) { return self->CanGoInside; }
void ImPlotNextItemData_SetLineWeight(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->LineWeight = v; }
float ImPlotNextItemData_GetLineWeight(ImPlotNextItemData *self) { return self->LineWeight; }
void ImPlotNextItemData_SetMarker(ImPlotNextItemData *ImPlotNextItemDataPtr, ImPlotMarker v) { ImPlotNextItemDataPtr->Marker = v; }
ImPlotMarker ImPlotNextItemData_GetMarker(ImPlotNextItemData *self) { return self->Marker; }
void ImPlotNextItemData_SetMarkerSize(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->MarkerSize = v; }
float ImPlotNextItemData_GetMarkerSize(ImPlotNextItemData *self) { return self->MarkerSize; }
void ImPlotNextItemData_SetMarkerWeight(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->MarkerWeight = v; }
float ImPlotNextItemData_GetMarkerWeight(ImPlotNextItemData *self) { return self->MarkerWeight; }
void ImPlotNextItemData_SetFillAlpha(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->FillAlpha = v; }
float ImPlotNextItemData_GetFillAlpha(ImPlotNextItemData *self) { return self->FillAlpha; }
void ImPlotNextItemData_SetErrorBarSize(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->ErrorBarSize = v; }
float ImPlotNextItemData_GetErrorBarSize(ImPlotNextItemData *self) { return self->ErrorBarSize; }
void ImPlotNextItemData_SetErrorBarWeight(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->ErrorBarWeight = v; }
float ImPlotNextItemData_GetErrorBarWeight(ImPlotNextItemData *self) { return self->ErrorBarWeight; }
void ImPlotNextItemData_SetDigitalBitHeight(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->DigitalBitHeight = v; }
float ImPlotNextItemData_GetDigitalBitHeight(ImPlotNextItemData *self) { return self->DigitalBitHeight; }
void ImPlotNextItemData_SetDigitalBitGap(ImPlotNextItemData *ImPlotNextItemDataPtr, float v) { ImPlotNextItemDataPtr->DigitalBitGap = v; }
float ImPlotNextItemData_GetDigitalBitGap(ImPlotNextItemData *self) { return self->DigitalBitGap; }
void ImPlotNextItemData_SetRenderLine(ImPlotNextItemData *ImPlotNextItemDataPtr, bool v) { ImPlotNextItemDataPtr->RenderLine = v; }
bool ImPlotNextItemData_GetRenderLine(ImPlotNextItemData *self) { return self->RenderLine; }
void ImPlotNextItemData_SetRenderFill(ImPlotNextItemData *ImPlotNextItemDataPtr, bool v) { ImPlotNextItemDataPtr->RenderFill = v; }
bool ImPlotNextItemData_GetRenderFill(ImPlotNextItemData *self) { return self->RenderFill; }
void ImPlotNextItemData_SetRenderMarkerLine(ImPlotNextItemData *ImPlotNextItemDataPtr, bool v) { ImPlotNextItemDataPtr->RenderMarkerLine = v; }
bool ImPlotNextItemData_GetRenderMarkerLine(ImPlotNextItemData *self) { return self->RenderMarkerLine; }
void ImPlotNextItemData_SetRenderMarkerFill(ImPlotNextItemData *ImPlotNextItemDataPtr, bool v) { ImPlotNextItemDataPtr->RenderMarkerFill = v; }
bool ImPlotNextItemData_GetRenderMarkerFill(ImPlotNextItemData *self) { return self->RenderMarkerFill; }
void ImPlotNextItemData_SetHasHidden(ImPlotNextItemData *ImPlotNextItemDataPtr, bool v) { ImPlotNextItemDataPtr->HasHidden = v; }
bool ImPlotNextItemData_GetHasHidden(ImPlotNextItemData *self) { return self->HasHidden; }
void ImPlotNextItemData_SetHidden(ImPlotNextItemData *ImPlotNextItemDataPtr, bool v) { ImPlotNextItemDataPtr->Hidden = v; }
bool ImPlotNextItemData_GetHidden(ImPlotNextItemData *self) { return self->Hidden; }
void ImPlotNextItemData_SetHiddenCond(ImPlotNextItemData *ImPlotNextItemDataPtr, ImPlotCond v) { ImPlotNextItemDataPtr->HiddenCond = v; }
ImPlotCond ImPlotNextItemData_GetHiddenCond(ImPlotNextItemData *self) { return self->HiddenCond; }
void ImPlotPlot_SetID(ImPlotPlot *ImPlotPlotPtr, ImGuiID v) { ImPlotPlotPtr->ID = v; }
ImGuiID ImPlotPlot_GetID(ImPlotPlot *self) { return self->ID; }
void ImPlotPlot_SetFlags(ImPlotPlot *ImPlotPlotPtr, ImPlotFlags v) { ImPlotPlotPtr->Flags = v; }
ImPlotFlags ImPlotPlot_GetFlags(ImPlotPlot *self) { return self->Flags; }
void ImPlotPlot_SetPreviousFlags(ImPlotPlot *ImPlotPlotPtr, ImPlotFlags v) { ImPlotPlotPtr->PreviousFlags = v; }
ImPlotFlags ImPlotPlot_GetPreviousFlags(ImPlotPlot *self) { return self->PreviousFlags; }
void ImPlotPlot_SetMouseTextLocation(ImPlotPlot *ImPlotPlotPtr, ImPlotLocation v) { ImPlotPlotPtr->MouseTextLocation = v; }
ImPlotLocation ImPlotPlot_GetMouseTextLocation(ImPlotPlot *self) { return self->MouseTextLocation; }
void ImPlotPlot_SetMouseTextFlags(ImPlotPlot *ImPlotPlotPtr, ImPlotMouseTextFlags v) { ImPlotPlotPtr->MouseTextFlags = v; }
ImPlotMouseTextFlags ImPlotPlot_GetMouseTextFlags(ImPlotPlot *self) { return self->MouseTextFlags; }
void ImPlotPlot_SetTextBuffer(ImPlotPlot *ImPlotPlotPtr, ImGuiTextBuffer v) { ImPlotPlotPtr->TextBuffer = v; }
ImGuiTextBuffer ImPlotPlot_GetTextBuffer(ImPlotPlot *self) { return self->TextBuffer; }
void ImPlotPlot_SetItems(ImPlotPlot *ImPlotPlotPtr, ImPlotItemGroup v) { ImPlotPlotPtr->Items = v; }
ImPlotItemGroup ImPlotPlot_GetItems(ImPlotPlot *self) { return self->Items; }
void ImPlotPlot_SetCurrentX(ImPlotPlot *ImPlotPlotPtr, ImAxis v) { ImPlotPlotPtr->CurrentX = v; }
ImAxis ImPlotPlot_GetCurrentX(ImPlotPlot *self) { return self->CurrentX; }
void ImPlotPlot_SetCurrentY(ImPlotPlot *ImPlotPlotPtr, ImAxis v) { ImPlotPlotPtr->CurrentY = v; }
ImAxis ImPlotPlot_GetCurrentY(ImPlotPlot *self) { return self->CurrentY; }
void ImPlotPlot_SetFrameRect(ImPlotPlot *ImPlotPlotPtr, ImRect v) { ImPlotPlotPtr->FrameRect = v; }
ImRect ImPlotPlot_GetFrameRect(ImPlotPlot *self) { return self->FrameRect; }
void ImPlotPlot_SetCanvasRect(ImPlotPlot *ImPlotPlotPtr, ImRect v) { ImPlotPlotPtr->CanvasRect = v; }
ImRect ImPlotPlot_GetCanvasRect(ImPlotPlot *self) { return self->CanvasRect; }
void ImPlotPlot_SetPlotRect(ImPlotPlot *ImPlotPlotPtr, ImRect v) { ImPlotPlotPtr->PlotRect = v; }
ImRect ImPlotPlot_GetPlotRect(ImPlotPlot *self) { return self->PlotRect; }
void ImPlotPlot_SetAxesRect(ImPlotPlot *ImPlotPlotPtr, ImRect v) { ImPlotPlotPtr->AxesRect = v; }
ImRect ImPlotPlot_GetAxesRect(ImPlotPlot *self) { return self->AxesRect; }
void ImPlotPlot_SetSelectRect(ImPlotPlot *ImPlotPlotPtr, ImRect v) { ImPlotPlotPtr->SelectRect = v; }
ImRect ImPlotPlot_GetSelectRect(ImPlotPlot *self) { return self->SelectRect; }
void ImPlotPlot_SetSelectStart(ImPlotPlot *ImPlotPlotPtr, ImVec2 v) { ImPlotPlotPtr->SelectStart = v; }
ImVec2 ImPlotPlot_GetSelectStart(ImPlotPlot *self) { return self->SelectStart; }
void ImPlotPlot_SetTitleOffset(ImPlotPlot *ImPlotPlotPtr, int v) { ImPlotPlotPtr->TitleOffset = v; }
int ImPlotPlot_GetTitleOffset(ImPlotPlot *self) { return self->TitleOffset; }
void ImPlotPlot_SetJustCreated(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->JustCreated = v; }
bool ImPlotPlot_GetJustCreated(ImPlotPlot *self) { return self->JustCreated; }
void ImPlotPlot_SetInitialized(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->Initialized = v; }
bool ImPlotPlot_GetInitialized(ImPlotPlot *self) { return self->Initialized; }
void ImPlotPlot_SetSetupLocked(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->SetupLocked = v; }
bool ImPlotPlot_GetSetupLocked(ImPlotPlot *self) { return self->SetupLocked; }
void ImPlotPlot_SetFitThisFrame(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->FitThisFrame = v; }
bool ImPlotPlot_GetFitThisFrame(ImPlotPlot *self) { return self->FitThisFrame; }
void ImPlotPlot_SetHovered(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->Hovered = v; }
bool ImPlotPlot_GetHovered(ImPlotPlot *self) { return self->Hovered; }
void ImPlotPlot_SetHeld(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->Held = v; }
bool ImPlotPlot_GetHeld(ImPlotPlot *self) { return self->Held; }
void ImPlotPlot_SetSelecting(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->Selecting = v; }
bool ImPlotPlot_GetSelecting(ImPlotPlot *self) { return self->Selecting; }
void ImPlotPlot_SetSelected(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->Selected = v; }
bool ImPlotPlot_GetSelected(ImPlotPlot *self) { return self->Selected; }
void ImPlotPlot_SetContextLocked(ImPlotPlot *ImPlotPlotPtr, bool v) { ImPlotPlotPtr->ContextLocked = v; }
bool ImPlotPlot_GetContextLocked(ImPlotPlot *self) { return self->ContextLocked; }
void ImPlotPointError_SetX(ImPlotPointError *ImPlotPointErrorPtr, double v) { ImPlotPointErrorPtr->X = v; }
double ImPlotPointError_GetX(ImPlotPointError *self) { return self->X; }
void ImPlotPointError_SetY(ImPlotPointError *ImPlotPointErrorPtr, double v) { ImPlotPointErrorPtr->Y = v; }
double ImPlotPointError_GetY(ImPlotPointError *self) { return self->Y; }
void ImPlotPointError_SetNeg(ImPlotPointError *ImPlotPointErrorPtr, double v) { ImPlotPointErrorPtr->Neg = v; }
double ImPlotPointError_GetNeg(ImPlotPointError *self) { return self->Neg; }
void ImPlotPointError_SetPos(ImPlotPointError *ImPlotPointErrorPtr, double v) { ImPlotPointErrorPtr->Pos = v; }
double ImPlotPointError_GetPos(ImPlotPointError *self) { return self->Pos; }
void ImPlotRange_SetMin(ImPlotRange *ImPlotRangePtr, double v) { ImPlotRangePtr->Min = v; }
double ImPlotRange_GetMin(ImPlotRange *self) { return self->Min; }
void ImPlotRange_SetMax(ImPlotRange *ImPlotRangePtr, double v) { ImPlotRangePtr->Max = v; }
double ImPlotRange_GetMax(ImPlotRange *self) { return self->Max; }
void ImPlotRect_SetX(ImPlotRect *ImPlotRectPtr, ImPlotRange v) { ImPlotRectPtr->X = v; }
ImPlotRange ImPlotRect_GetX(ImPlotRect *self) { return self->X; }
void ImPlotRect_SetY(ImPlotRect *ImPlotRectPtr, ImPlotRange v) { ImPlotRectPtr->Y = v; }
ImPlotRange ImPlotRect_GetY(ImPlotRect *self) { return self->Y; }
void ImPlotStyle_SetLineWeight(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->LineWeight = v; }
float ImPlotStyle_GetLineWeight(ImPlotStyle *self) { return self->LineWeight; }
void ImPlotStyle_SetMarker(ImPlotStyle *ImPlotStylePtr, int v) { ImPlotStylePtr->Marker = v; }
int ImPlotStyle_GetMarker(ImPlotStyle *self) { return self->Marker; }
void ImPlotStyle_SetMarkerSize(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->MarkerSize = v; }
float ImPlotStyle_GetMarkerSize(ImPlotStyle *self) { return self->MarkerSize; }
void ImPlotStyle_SetMarkerWeight(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->MarkerWeight = v; }
float ImPlotStyle_GetMarkerWeight(ImPlotStyle *self) { return self->MarkerWeight; }
void ImPlotStyle_SetFillAlpha(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->FillAlpha = v; }
float ImPlotStyle_GetFillAlpha(ImPlotStyle *self) { return self->FillAlpha; }
void ImPlotStyle_SetErrorBarSize(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->ErrorBarSize = v; }
float ImPlotStyle_GetErrorBarSize(ImPlotStyle *self) { return self->ErrorBarSize; }
void ImPlotStyle_SetErrorBarWeight(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->ErrorBarWeight = v; }
float ImPlotStyle_GetErrorBarWeight(ImPlotStyle *self) { return self->ErrorBarWeight; }
void ImPlotStyle_SetDigitalBitHeight(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->DigitalBitHeight = v; }
float ImPlotStyle_GetDigitalBitHeight(ImPlotStyle *self) { return self->DigitalBitHeight; }
void ImPlotStyle_SetDigitalBitGap(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->DigitalBitGap = v; }
float ImPlotStyle_GetDigitalBitGap(ImPlotStyle *self) { return self->DigitalBitGap; }
void ImPlotStyle_SetPlotBorderSize(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->PlotBorderSize = v; }
float ImPlotStyle_GetPlotBorderSize(ImPlotStyle *self) { return self->PlotBorderSize; }
void ImPlotStyle_SetMinorAlpha(ImPlotStyle *ImPlotStylePtr, float v) { ImPlotStylePtr->MinorAlpha = v; }
float ImPlotStyle_GetMinorAlpha(ImPlotStyle *self) { return self->MinorAlpha; }
void ImPlotStyle_SetMajorTickLen(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MajorTickLen = v; }
ImVec2 ImPlotStyle_GetMajorTickLen(ImPlotStyle *self) { return self->MajorTickLen; }
void ImPlotStyle_SetMinorTickLen(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MinorTickLen = v; }
ImVec2 ImPlotStyle_GetMinorTickLen(ImPlotStyle *self) { return self->MinorTickLen; }
void ImPlotStyle_SetMajorTickSize(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MajorTickSize = v; }
ImVec2 ImPlotStyle_GetMajorTickSize(ImPlotStyle *self) { return self->MajorTickSize; }
void ImPlotStyle_SetMinorTickSize(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MinorTickSize = v; }
ImVec2 ImPlotStyle_GetMinorTickSize(ImPlotStyle *self) { return self->MinorTickSize; }
void ImPlotStyle_SetMajorGridSize(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MajorGridSize = v; }
ImVec2 ImPlotStyle_GetMajorGridSize(ImPlotStyle *self) { return self->MajorGridSize; }
void ImPlotStyle_SetMinorGridSize(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MinorGridSize = v; }
ImVec2 ImPlotStyle_GetMinorGridSize(ImPlotStyle *self) { return self->MinorGridSize; }
void ImPlotStyle_SetPlotPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->PlotPadding = v; }
ImVec2 ImPlotStyle_GetPlotPadding(ImPlotStyle *self) { return self->PlotPadding; }
void ImPlotStyle_SetLabelPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->LabelPadding = v; }
ImVec2 ImPlotStyle_GetLabelPadding(ImPlotStyle *self) { return self->LabelPadding; }
void ImPlotStyle_SetLegendPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->LegendPadding = v; }
ImVec2 ImPlotStyle_GetLegendPadding(ImPlotStyle *self) { return self->LegendPadding; }
void ImPlotStyle_SetLegendInnerPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->LegendInnerPadding = v; }
ImVec2 ImPlotStyle_GetLegendInnerPadding(ImPlotStyle *self) { return self->LegendInnerPadding; }
void ImPlotStyle_SetLegendSpacing(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->LegendSpacing = v; }
ImVec2 ImPlotStyle_GetLegendSpacing(ImPlotStyle *self) { return self->LegendSpacing; }
void ImPlotStyle_SetMousePosPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->MousePosPadding = v; }
ImVec2 ImPlotStyle_GetMousePosPadding(ImPlotStyle *self) { return self->MousePosPadding; }
void ImPlotStyle_SetAnnotationPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->AnnotationPadding = v; }
ImVec2 ImPlotStyle_GetAnnotationPadding(ImPlotStyle *self) { return self->AnnotationPadding; }
void ImPlotStyle_SetFitPadding(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->FitPadding = v; }
ImVec2 ImPlotStyle_GetFitPadding(ImPlotStyle *self) { return self->FitPadding; }
void ImPlotStyle_SetPlotDefaultSize(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->PlotDefaultSize = v; }
ImVec2 ImPlotStyle_GetPlotDefaultSize(ImPlotStyle *self) { return self->PlotDefaultSize; }
void ImPlotStyle_SetPlotMinSize(ImPlotStyle *ImPlotStylePtr, ImVec2 v) { ImPlotStylePtr->PlotMinSize = v; }
ImVec2 ImPlotStyle_GetPlotMinSize(ImPlotStyle *self) { return self->PlotMinSize; }
void ImPlotStyle_SetColormap(ImPlotStyle *ImPlotStylePtr, ImPlotColormap v) { ImPlotStylePtr->Colormap = v; }
ImPlotColormap ImPlotStyle_GetColormap(ImPlotStyle *self) { return self->Colormap; }
void ImPlotStyle_SetUseLocalTime(ImPlotStyle *ImPlotStylePtr, bool v) { ImPlotStylePtr->UseLocalTime = v; }
bool ImPlotStyle_GetUseLocalTime(ImPlotStyle *self) { return self->UseLocalTime; }
void ImPlotStyle_SetUseISO8601(ImPlotStyle *ImPlotStylePtr, bool v) { ImPlotStylePtr->UseISO8601 = v; }
bool ImPlotStyle_GetUseISO8601(ImPlotStyle *self) { return self->UseISO8601; }
void ImPlotStyle_SetUse24HourClock(ImPlotStyle *ImPlotStylePtr, bool v) { ImPlotStylePtr->Use24HourClock = v; }
bool ImPlotStyle_GetUse24HourClock(ImPlotStyle *self) { return self->Use24HourClock; }
void ImPlotSubplot_SetID(ImPlotSubplot *ImPlotSubplotPtr, ImGuiID v) { ImPlotSubplotPtr->ID = v; }
ImGuiID ImPlotSubplot_GetID(ImPlotSubplot *self) { return self->ID; }
void ImPlotSubplot_SetFlags(ImPlotSubplot *ImPlotSubplotPtr, ImPlotSubplotFlags v) { ImPlotSubplotPtr->Flags = v; }
ImPlotSubplotFlags ImPlotSubplot_GetFlags(ImPlotSubplot *self) { return self->Flags; }
void ImPlotSubplot_SetPreviousFlags(ImPlotSubplot *ImPlotSubplotPtr, ImPlotSubplotFlags v) { ImPlotSubplotPtr->PreviousFlags = v; }
ImPlotSubplotFlags ImPlotSubplot_GetPreviousFlags(ImPlotSubplot *self) { return self->PreviousFlags; }
void ImPlotSubplot_SetItems(ImPlotSubplot *ImPlotSubplotPtr, ImPlotItemGroup v) { ImPlotSubplotPtr->Items = v; }
ImPlotItemGroup ImPlotSubplot_GetItems(ImPlotSubplot *self) { return self->Items; }
void ImPlotSubplot_SetRows(ImPlotSubplot *ImPlotSubplotPtr, int v) { ImPlotSubplotPtr->Rows = v; }
int ImPlotSubplot_GetRows(ImPlotSubplot *self) { return self->Rows; }
void ImPlotSubplot_SetCols(ImPlotSubplot *ImPlotSubplotPtr, int v) { ImPlotSubplotPtr->Cols = v; }
int ImPlotSubplot_GetCols(ImPlotSubplot *self) { return self->Cols; }
void ImPlotSubplot_SetCurrentIdx(ImPlotSubplot *ImPlotSubplotPtr, int v) { ImPlotSubplotPtr->CurrentIdx = v; }
int ImPlotSubplot_GetCurrentIdx(ImPlotSubplot *self) { return self->CurrentIdx; }
void ImPlotSubplot_SetFrameRect(ImPlotSubplot *ImPlotSubplotPtr, ImRect v) { ImPlotSubplotPtr->FrameRect = v; }
ImRect ImPlotSubplot_GetFrameRect(ImPlotSubplot *self) { return self->FrameRect; }
void ImPlotSubplot_SetGridRect(ImPlotSubplot *ImPlotSubplotPtr, ImRect v) { ImPlotSubplotPtr->GridRect = v; }
ImRect ImPlotSubplot_GetGridRect(ImPlotSubplot *self) { return self->GridRect; }
void ImPlotSubplot_SetCellSize(ImPlotSubplot *ImPlotSubplotPtr, ImVec2 v) { ImPlotSubplotPtr->CellSize = v; }
ImVec2 ImPlotSubplot_GetCellSize(ImPlotSubplot *self) { return self->CellSize; }
void ImPlotSubplot_SetRowAlignmentData(ImPlotSubplot *ImPlotSubplotPtr, ImVector_ImPlotAlignmentData v) { ImPlotSubplotPtr->RowAlignmentData = v; }
ImVector_ImPlotAlignmentData ImPlotSubplot_GetRowAlignmentData(ImPlotSubplot *self) { return self->RowAlignmentData; }
void ImPlotSubplot_SetColAlignmentData(ImPlotSubplot *ImPlotSubplotPtr, ImVector_ImPlotAlignmentData v) { ImPlotSubplotPtr->ColAlignmentData = v; }
ImVector_ImPlotAlignmentData ImPlotSubplot_GetColAlignmentData(ImPlotSubplot *self) { return self->ColAlignmentData; }
void ImPlotSubplot_SetRowRatios(ImPlotSubplot *ImPlotSubplotPtr, ImVector_float v) { ImPlotSubplotPtr->RowRatios = v; }
ImVector_float ImPlotSubplot_GetRowRatios(ImPlotSubplot *self) { return self->RowRatios; }
void ImPlotSubplot_SetColRatios(ImPlotSubplot *ImPlotSubplotPtr, ImVector_float v) { ImPlotSubplotPtr->ColRatios = v; }
ImVector_float ImPlotSubplot_GetColRatios(ImPlotSubplot *self) { return self->ColRatios; }
void ImPlotSubplot_SetRowLinkData(ImPlotSubplot *ImPlotSubplotPtr, ImVector_ImPlotRange v) { ImPlotSubplotPtr->RowLinkData = v; }
ImVector_ImPlotRange ImPlotSubplot_GetRowLinkData(ImPlotSubplot *self) { return self->RowLinkData; }
void ImPlotSubplot_SetColLinkData(ImPlotSubplot *ImPlotSubplotPtr, ImVector_ImPlotRange v) { ImPlotSubplotPtr->ColLinkData = v; }
ImVector_ImPlotRange ImPlotSubplot_GetColLinkData(ImPlotSubplot *self) { return self->ColLinkData; }
void ImPlotSubplot_SetFrameHovered(ImPlotSubplot *ImPlotSubplotPtr, bool v) { ImPlotSubplotPtr->FrameHovered = v; }
bool ImPlotSubplot_GetFrameHovered(ImPlotSubplot *self) { return self->FrameHovered; }
void ImPlotSubplot_SetHasTitle(ImPlotSubplot *ImPlotSubplotPtr, bool v) { ImPlotSubplotPtr->HasTitle = v; }
bool ImPlotSubplot_GetHasTitle(ImPlotSubplot *self) { return self->HasTitle; }
void ImPlotTag_SetAxis(ImPlotTag *ImPlotTagPtr, ImAxis v) { ImPlotTagPtr->Axis = v; }
ImAxis ImPlotTag_GetAxis(ImPlotTag *self) { return self->Axis; }
void ImPlotTag_SetValue(ImPlotTag *ImPlotTagPtr, double v) { ImPlotTagPtr->Value = v; }
double ImPlotTag_GetValue(ImPlotTag *self) { return self->Value; }
void ImPlotTag_SetColorBg(ImPlotTag *ImPlotTagPtr, ImU32 v) { ImPlotTagPtr->ColorBg = v; }
ImU32 ImPlotTag_GetColorBg(ImPlotTag *self) { return self->ColorBg; }
void ImPlotTag_SetColorFg(ImPlotTag *ImPlotTagPtr, ImU32 v) { ImPlotTagPtr->ColorFg = v; }
ImU32 ImPlotTag_GetColorFg(ImPlotTag *self) { return self->ColorFg; }
void ImPlotTag_SetTextOffset(ImPlotTag *ImPlotTagPtr, int v) { ImPlotTagPtr->TextOffset = v; }
int ImPlotTag_GetTextOffset(ImPlotTag *self) { return self->TextOffset; }
void ImPlotTagCollection_SetTags(ImPlotTagCollection *ImPlotTagCollectionPtr, ImVector_ImPlotTag v) { ImPlotTagCollectionPtr->Tags = v; }
ImVector_ImPlotTag ImPlotTagCollection_GetTags(ImPlotTagCollection *self) { return self->Tags; }
void ImPlotTagCollection_SetTextBuffer(ImPlotTagCollection *ImPlotTagCollectionPtr, ImGuiTextBuffer v) { ImPlotTagCollectionPtr->TextBuffer = v; }
ImGuiTextBuffer ImPlotTagCollection_GetTextBuffer(ImPlotTagCollection *self) { return self->TextBuffer; }
void ImPlotTagCollection_SetSize(ImPlotTagCollection *ImPlotTagCollectionPtr, int v) { ImPlotTagCollectionPtr->Size = v; }
int ImPlotTagCollection_GetSize(ImPlotTagCollection *self) { return self->Size; }
void ImPlotTick_SetPlotPos(ImPlotTick *ImPlotTickPtr, double v) { ImPlotTickPtr->PlotPos = v; }
double ImPlotTick_GetPlotPos(ImPlotTick *self) { return self->PlotPos; }
void ImPlotTick_SetPixelPos(ImPlotTick *ImPlotTickPtr, float v) { ImPlotTickPtr->PixelPos = v; }
float ImPlotTick_GetPixelPos(ImPlotTick *self) { return self->PixelPos; }
void ImPlotTick_SetLabelSize(ImPlotTick *ImPlotTickPtr, ImVec2 v) { ImPlotTickPtr->LabelSize = v; }
ImVec2 ImPlotTick_GetLabelSize(ImPlotTick *self) { return self->LabelSize; }
void ImPlotTick_SetTextOffset(ImPlotTick *ImPlotTickPtr, int v) { ImPlotTickPtr->TextOffset = v; }
int ImPlotTick_GetTextOffset(ImPlotTick *self) { return self->TextOffset; }
void ImPlotTick_SetMajor(ImPlotTick *ImPlotTickPtr, bool v) { ImPlotTickPtr->Major = v; }
bool ImPlotTick_GetMajor(ImPlotTick *self) { return self->Major; }
void ImPlotTick_SetShowLabel(ImPlotTick *ImPlotTickPtr, bool v) { ImPlotTickPtr->ShowLabel = v; }
bool ImPlotTick_GetShowLabel(ImPlotTick *self) { return self->ShowLabel; }
void ImPlotTick_SetLevel(ImPlotTick *ImPlotTickPtr, int v) { ImPlotTickPtr->Level = v; }
int ImPlotTick_GetLevel(ImPlotTick *self) { return self->Level; }
void ImPlotTick_SetIdx(ImPlotTick *ImPlotTickPtr, int v) { ImPlotTickPtr->Idx = v; }
int ImPlotTick_GetIdx(ImPlotTick *self) { return self->Idx; }
void ImPlotTicker_SetTicks(ImPlotTicker *ImPlotTickerPtr, ImVector_ImPlotTick v) { ImPlotTickerPtr->Ticks = v; }
ImVector_ImPlotTick ImPlotTicker_GetTicks(ImPlotTicker *self) { return self->Ticks; }
void ImPlotTicker_SetTextBuffer(ImPlotTicker *ImPlotTickerPtr, ImGuiTextBuffer v) { ImPlotTickerPtr->TextBuffer = v; }
ImGuiTextBuffer ImPlotTicker_GetTextBuffer(ImPlotTicker *self) { return self->TextBuffer; }
void ImPlotTicker_SetMaxSize(ImPlotTicker *ImPlotTickerPtr, ImVec2 v) { ImPlotTickerPtr->MaxSize = v; }
ImVec2 ImPlotTicker_GetMaxSize(ImPlotTicker *self) { return self->MaxSize; }
void ImPlotTicker_SetLateSize(ImPlotTicker *ImPlotTickerPtr, ImVec2 v) { ImPlotTickerPtr->LateSize = v; }
ImVec2 ImPlotTicker_GetLateSize(ImPlotTicker *self) { return self->LateSize; }
void ImPlotTicker_SetLevels(ImPlotTicker *ImPlotTickerPtr, int v) { ImPlotTickerPtr->Levels = v; }
int ImPlotTicker_GetLevels(ImPlotTicker *self) { return self->Levels; }
void ImPlotTime_SetS(ImPlotTime *ImPlotTimePtr, time_t v) { ImPlotTimePtr->S = v; }
time_t ImPlotTime_GetS(ImPlotTime *self) { return self->S; }
void ImPlotTime_SetUs(ImPlotTime *ImPlotTimePtr, int v) { ImPlotTimePtr->Us = v; }
int ImPlotTime_GetUs(ImPlotTime *self) { return self->Us; }
