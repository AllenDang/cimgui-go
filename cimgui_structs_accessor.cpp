
#include "cimgui_wrapper.h"
#include "cimgui_structs_accessor.h"

void ImFontAtlasCustomRect_SetWidth(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, unsigned short v) { ImFontAtlasCustomRectPtr->Width = v; }
unsigned short ImFontAtlasCustomRect_GetWidth(ImFontAtlasCustomRect *self) { return self->Width; }
void ImFontAtlasCustomRect_SetHeight(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, unsigned short v) { ImFontAtlasCustomRectPtr->Height = v; }
unsigned short ImFontAtlasCustomRect_GetHeight(ImFontAtlasCustomRect *self) { return self->Height; }
void ImFontAtlasCustomRect_SetX(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, unsigned short v) { ImFontAtlasCustomRectPtr->X = v; }
unsigned short ImFontAtlasCustomRect_GetX(ImFontAtlasCustomRect *self) { return self->X; }
void ImFontAtlasCustomRect_SetY(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, unsigned short v) { ImFontAtlasCustomRectPtr->Y = v; }
unsigned short ImFontAtlasCustomRect_GetY(ImFontAtlasCustomRect *self) { return self->Y; }
void ImFontAtlasCustomRect_SetGlyphID(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, unsigned int v) { ImFontAtlasCustomRectPtr->GlyphID = v; }
unsigned int ImFontAtlasCustomRect_GetGlyphID(ImFontAtlasCustomRect *self) { return self->GlyphID; }
void ImFontAtlasCustomRect_SetGlyphAdvanceX(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, float v) { ImFontAtlasCustomRectPtr->GlyphAdvanceX = v; }
float ImFontAtlasCustomRect_GetGlyphAdvanceX(ImFontAtlasCustomRect *self) { return self->GlyphAdvanceX; }
void ImFontAtlasCustomRect_SetGlyphOffset(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, ImVec2 v) { ImFontAtlasCustomRectPtr->GlyphOffset = v; }
ImVec2 ImFontAtlasCustomRect_GetGlyphOffset(ImFontAtlasCustomRect *self) { return self->GlyphOffset; }
void ImFontAtlasCustomRect_SetFont(ImFontAtlasCustomRect *ImFontAtlasCustomRectPtr, ImFont* v) { ImFontAtlasCustomRectPtr->Font = v; }
ImFont* ImFontAtlasCustomRect_GetFont(ImFontAtlasCustomRect *self) { return self->Font; }
void ImGuiComboPreviewData_SetPreviewRect(ImGuiComboPreviewData *ImGuiComboPreviewDataPtr, ImRect v) { ImGuiComboPreviewDataPtr->PreviewRect = v; }
ImRect ImGuiComboPreviewData_GetPreviewRect(ImGuiComboPreviewData *self) { return self->PreviewRect; }
void ImGuiComboPreviewData_SetBackupCursorPos(ImGuiComboPreviewData *ImGuiComboPreviewDataPtr, ImVec2 v) { ImGuiComboPreviewDataPtr->BackupCursorPos = v; }
ImVec2 ImGuiComboPreviewData_GetBackupCursorPos(ImGuiComboPreviewData *self) { return self->BackupCursorPos; }
void ImGuiComboPreviewData_SetBackupCursorMaxPos(ImGuiComboPreviewData *ImGuiComboPreviewDataPtr, ImVec2 v) { ImGuiComboPreviewDataPtr->BackupCursorMaxPos = v; }
ImVec2 ImGuiComboPreviewData_GetBackupCursorMaxPos(ImGuiComboPreviewData *self) { return self->BackupCursorMaxPos; }
void ImGuiComboPreviewData_SetBackupCursorPosPrevLine(ImGuiComboPreviewData *ImGuiComboPreviewDataPtr, ImVec2 v) { ImGuiComboPreviewDataPtr->BackupCursorPosPrevLine = v; }
ImVec2 ImGuiComboPreviewData_GetBackupCursorPosPrevLine(ImGuiComboPreviewData *self) { return self->BackupCursorPosPrevLine; }
void ImGuiComboPreviewData_SetBackupPrevLineTextBaseOffset(ImGuiComboPreviewData *ImGuiComboPreviewDataPtr, float v) { ImGuiComboPreviewDataPtr->BackupPrevLineTextBaseOffset = v; }
float ImGuiComboPreviewData_GetBackupPrevLineTextBaseOffset(ImGuiComboPreviewData *self) { return self->BackupPrevLineTextBaseOffset; }
void ImGuiComboPreviewData_SetBackupLayout(ImGuiComboPreviewData *ImGuiComboPreviewDataPtr, ImGuiLayoutType v) { ImGuiComboPreviewDataPtr->BackupLayout = v; }
ImGuiLayoutType ImGuiComboPreviewData_GetBackupLayout(ImGuiComboPreviewData *self) { return self->BackupLayout; }
void ImGuiTableTempData_SetTableIndex(ImGuiTableTempData *ImGuiTableTempDataPtr, int v) { ImGuiTableTempDataPtr->TableIndex = v; }
int ImGuiTableTempData_GetTableIndex(ImGuiTableTempData *self) { return self->TableIndex; }
void ImGuiTableTempData_SetLastTimeActive(ImGuiTableTempData *ImGuiTableTempDataPtr, float v) { ImGuiTableTempDataPtr->LastTimeActive = v; }
float ImGuiTableTempData_GetLastTimeActive(ImGuiTableTempData *self) { return self->LastTimeActive; }
void ImGuiTableTempData_SetUserOuterSize(ImGuiTableTempData *ImGuiTableTempDataPtr, ImVec2 v) { ImGuiTableTempDataPtr->UserOuterSize = v; }
ImVec2 ImGuiTableTempData_GetUserOuterSize(ImGuiTableTempData *self) { return self->UserOuterSize; }
void ImGuiTableTempData_SetDrawSplitter(ImGuiTableTempData *ImGuiTableTempDataPtr, ImDrawListSplitter v) { ImGuiTableTempDataPtr->DrawSplitter = v; }
ImDrawListSplitter ImGuiTableTempData_GetDrawSplitter(ImGuiTableTempData *self) { return self->DrawSplitter; }
void ImGuiTableTempData_SetHostBackupWorkRect(ImGuiTableTempData *ImGuiTableTempDataPtr, ImRect v) { ImGuiTableTempDataPtr->HostBackupWorkRect = v; }
ImRect ImGuiTableTempData_GetHostBackupWorkRect(ImGuiTableTempData *self) { return self->HostBackupWorkRect; }
void ImGuiTableTempData_SetHostBackupParentWorkRect(ImGuiTableTempData *ImGuiTableTempDataPtr, ImRect v) { ImGuiTableTempDataPtr->HostBackupParentWorkRect = v; }
ImRect ImGuiTableTempData_GetHostBackupParentWorkRect(ImGuiTableTempData *self) { return self->HostBackupParentWorkRect; }
void ImGuiTableTempData_SetHostBackupPrevLineSize(ImGuiTableTempData *ImGuiTableTempDataPtr, ImVec2 v) { ImGuiTableTempDataPtr->HostBackupPrevLineSize = v; }
ImVec2 ImGuiTableTempData_GetHostBackupPrevLineSize(ImGuiTableTempData *self) { return self->HostBackupPrevLineSize; }
void ImGuiTableTempData_SetHostBackupCurrLineSize(ImGuiTableTempData *ImGuiTableTempDataPtr, ImVec2 v) { ImGuiTableTempDataPtr->HostBackupCurrLineSize = v; }
ImVec2 ImGuiTableTempData_GetHostBackupCurrLineSize(ImGuiTableTempData *self) { return self->HostBackupCurrLineSize; }
void ImGuiTableTempData_SetHostBackupCursorMaxPos(ImGuiTableTempData *ImGuiTableTempDataPtr, ImVec2 v) { ImGuiTableTempDataPtr->HostBackupCursorMaxPos = v; }
ImVec2 ImGuiTableTempData_GetHostBackupCursorMaxPos(ImGuiTableTempData *self) { return self->HostBackupCursorMaxPos; }
void ImGuiTableTempData_SetHostBackupColumnsOffset(ImGuiTableTempData *ImGuiTableTempDataPtr, ImVec1 v) { ImGuiTableTempDataPtr->HostBackupColumnsOffset = v; }
ImVec1 ImGuiTableTempData_GetHostBackupColumnsOffset(ImGuiTableTempData *self) { return self->HostBackupColumnsOffset; }
void ImGuiTableTempData_SetHostBackupItemWidth(ImGuiTableTempData *ImGuiTableTempDataPtr, float v) { ImGuiTableTempDataPtr->HostBackupItemWidth = v; }
float ImGuiTableTempData_GetHostBackupItemWidth(ImGuiTableTempData *self) { return self->HostBackupItemWidth; }
void ImGuiTableTempData_SetHostBackupItemWidthStackSize(ImGuiTableTempData *ImGuiTableTempDataPtr, int v) { ImGuiTableTempDataPtr->HostBackupItemWidthStackSize = v; }
int ImGuiTableTempData_GetHostBackupItemWidthStackSize(ImGuiTableTempData *self) { return self->HostBackupItemWidthStackSize; }
void ImGuiViewportP_Set_ImGuiViewport(ImGuiViewportP *ImGuiViewportPPtr, ImGuiViewport v) { ImGuiViewportPPtr->_ImGuiViewport = v; }
ImGuiViewport ImGuiViewportP_Get_ImGuiViewport(ImGuiViewportP *self) { return self->_ImGuiViewport; }
void ImGuiViewportP_SetIdx(ImGuiViewportP *ImGuiViewportPPtr, int v) { ImGuiViewportPPtr->Idx = v; }
int ImGuiViewportP_GetIdx(ImGuiViewportP *self) { return self->Idx; }
void ImGuiViewportP_SetLastFrameActive(ImGuiViewportP *ImGuiViewportPPtr, int v) { ImGuiViewportPPtr->LastFrameActive = v; }
int ImGuiViewportP_GetLastFrameActive(ImGuiViewportP *self) { return self->LastFrameActive; }
void ImGuiViewportP_SetLastFrontMostStampCount(ImGuiViewportP *ImGuiViewportPPtr, int v) { ImGuiViewportPPtr->LastFrontMostStampCount = v; }
int ImGuiViewportP_GetLastFrontMostStampCount(ImGuiViewportP *self) { return self->LastFrontMostStampCount; }
void ImGuiViewportP_SetLastNameHash(ImGuiViewportP *ImGuiViewportPPtr, ImGuiID v) { ImGuiViewportPPtr->LastNameHash = v; }
ImGuiID ImGuiViewportP_GetLastNameHash(ImGuiViewportP *self) { return self->LastNameHash; }
void ImGuiViewportP_SetLastPos(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->LastPos = v; }
ImVec2 ImGuiViewportP_GetLastPos(ImGuiViewportP *self) { return self->LastPos; }
void ImGuiViewportP_SetAlpha(ImGuiViewportP *ImGuiViewportPPtr, float v) { ImGuiViewportPPtr->Alpha = v; }
float ImGuiViewportP_GetAlpha(ImGuiViewportP *self) { return self->Alpha; }
void ImGuiViewportP_SetLastAlpha(ImGuiViewportP *ImGuiViewportPPtr, float v) { ImGuiViewportPPtr->LastAlpha = v; }
float ImGuiViewportP_GetLastAlpha(ImGuiViewportP *self) { return self->LastAlpha; }
void ImGuiViewportP_SetPlatformMonitor(ImGuiViewportP *ImGuiViewportPPtr, short v) { ImGuiViewportPPtr->PlatformMonitor = v; }
short ImGuiViewportP_GetPlatformMonitor(ImGuiViewportP *self) { return self->PlatformMonitor; }
void ImGuiViewportP_SetPlatformWindowCreated(ImGuiViewportP *ImGuiViewportPPtr, bool v) { ImGuiViewportPPtr->PlatformWindowCreated = v; }
bool ImGuiViewportP_GetPlatformWindowCreated(ImGuiViewportP *self) { return self->PlatformWindowCreated; }
void ImGuiViewportP_SetWindow(ImGuiViewportP *ImGuiViewportPPtr, ImGuiWindow* v) { ImGuiViewportPPtr->Window = v; }
ImGuiWindow* ImGuiViewportP_GetWindow(ImGuiViewportP *self) { return self->Window; }
void ImGuiViewportP_SetDrawDataP(ImGuiViewportP *ImGuiViewportPPtr, ImDrawData v) { ImGuiViewportPPtr->DrawDataP = v; }
ImDrawData ImGuiViewportP_GetDrawDataP(ImGuiViewportP *self) { return self->DrawDataP; }
void ImGuiViewportP_SetDrawDataBuilder(ImGuiViewportP *ImGuiViewportPPtr, ImDrawDataBuilder v) { ImGuiViewportPPtr->DrawDataBuilder = v; }
ImDrawDataBuilder ImGuiViewportP_GetDrawDataBuilder(ImGuiViewportP *self) { return self->DrawDataBuilder; }
void ImGuiViewportP_SetLastPlatformPos(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->LastPlatformPos = v; }
ImVec2 ImGuiViewportP_GetLastPlatformPos(ImGuiViewportP *self) { return self->LastPlatformPos; }
void ImGuiViewportP_SetLastPlatformSize(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->LastPlatformSize = v; }
ImVec2 ImGuiViewportP_GetLastPlatformSize(ImGuiViewportP *self) { return self->LastPlatformSize; }
void ImGuiViewportP_SetLastRendererSize(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->LastRendererSize = v; }
ImVec2 ImGuiViewportP_GetLastRendererSize(ImGuiViewportP *self) { return self->LastRendererSize; }
void ImGuiViewportP_SetWorkOffsetMin(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->WorkOffsetMin = v; }
ImVec2 ImGuiViewportP_GetWorkOffsetMin(ImGuiViewportP *self) { return self->WorkOffsetMin; }
void ImGuiViewportP_SetWorkOffsetMax(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->WorkOffsetMax = v; }
ImVec2 ImGuiViewportP_GetWorkOffsetMax(ImGuiViewportP *self) { return self->WorkOffsetMax; }
void ImGuiViewportP_SetBuildWorkOffsetMin(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->BuildWorkOffsetMin = v; }
ImVec2 ImGuiViewportP_GetBuildWorkOffsetMin(ImGuiViewportP *self) { return self->BuildWorkOffsetMin; }
void ImGuiViewportP_SetBuildWorkOffsetMax(ImGuiViewportP *ImGuiViewportPPtr, ImVec2 v) { ImGuiViewportPPtr->BuildWorkOffsetMax = v; }
ImVec2 ImGuiViewportP_GetBuildWorkOffsetMax(ImGuiViewportP *self) { return self->BuildWorkOffsetMax; }
void StbTexteditRow_Setx0(StbTexteditRow *StbTexteditRowPtr, float v) { StbTexteditRowPtr->x0 = v; }
float StbTexteditRow_Getx0(StbTexteditRow *self) { return self->x0; }
void StbTexteditRow_Setx1(StbTexteditRow *StbTexteditRowPtr, float v) { StbTexteditRowPtr->x1 = v; }
float StbTexteditRow_Getx1(StbTexteditRow *self) { return self->x1; }
void StbTexteditRow_Setbaseline_y_delta(StbTexteditRow *StbTexteditRowPtr, float v) { StbTexteditRowPtr->baseline_y_delta = v; }
float StbTexteditRow_Getbaseline_y_delta(StbTexteditRow *self) { return self->baseline_y_delta; }
void StbTexteditRow_Setymin(StbTexteditRow *StbTexteditRowPtr, float v) { StbTexteditRowPtr->ymin = v; }
float StbTexteditRow_Getymin(StbTexteditRow *self) { return self->ymin; }
void StbTexteditRow_Setymax(StbTexteditRow *StbTexteditRowPtr, float v) { StbTexteditRowPtr->ymax = v; }
float StbTexteditRow_Getymax(StbTexteditRow *self) { return self->ymax; }
void StbTexteditRow_Setnum_chars(StbTexteditRow *StbTexteditRowPtr, int v) { StbTexteditRowPtr->num_chars = v; }
int StbTexteditRow_Getnum_chars(StbTexteditRow *self) { return self->num_chars; }
void ImFontAtlas_SetFlags(ImFontAtlas *ImFontAtlasPtr, ImFontAtlasFlags v) { ImFontAtlasPtr->Flags = v; }
ImFontAtlasFlags ImFontAtlas_GetFlags(ImFontAtlas *self) { return self->Flags; }
void ImFontAtlas_SetTexDesiredWidth(ImFontAtlas *ImFontAtlasPtr, int v) { ImFontAtlasPtr->TexDesiredWidth = v; }
int ImFontAtlas_GetTexDesiredWidth(ImFontAtlas *self) { return self->TexDesiredWidth; }
void ImFontAtlas_SetTexGlyphPadding(ImFontAtlas *ImFontAtlasPtr, int v) { ImFontAtlasPtr->TexGlyphPadding = v; }
int ImFontAtlas_GetTexGlyphPadding(ImFontAtlas *self) { return self->TexGlyphPadding; }
void ImFontAtlas_SetLocked(ImFontAtlas *ImFontAtlasPtr, bool v) { ImFontAtlasPtr->Locked = v; }
bool ImFontAtlas_GetLocked(ImFontAtlas *self) { return self->Locked; }
void ImFontAtlas_SetTexReady(ImFontAtlas *ImFontAtlasPtr, bool v) { ImFontAtlasPtr->TexReady = v; }
bool ImFontAtlas_GetTexReady(ImFontAtlas *self) { return self->TexReady; }
void ImFontAtlas_SetTexPixelsUseColors(ImFontAtlas *ImFontAtlasPtr, bool v) { ImFontAtlasPtr->TexPixelsUseColors = v; }
bool ImFontAtlas_GetTexPixelsUseColors(ImFontAtlas *self) { return self->TexPixelsUseColors; }
void ImFontAtlas_SetTexPixelsAlpha8(ImFontAtlas *ImFontAtlasPtr, unsigned char* v) { ImFontAtlasPtr->TexPixelsAlpha8 = v; }
unsigned char* ImFontAtlas_GetTexPixelsAlpha8(ImFontAtlas *self) { return self->TexPixelsAlpha8; }
void ImFontAtlas_SetTexPixelsRGBA32(ImFontAtlas *ImFontAtlasPtr, unsigned int* v) { ImFontAtlasPtr->TexPixelsRGBA32 = v; }
unsigned int* ImFontAtlas_GetTexPixelsRGBA32(ImFontAtlas *self) { return self->TexPixelsRGBA32; }
void ImFontAtlas_SetTexWidth(ImFontAtlas *ImFontAtlasPtr, int v) { ImFontAtlasPtr->TexWidth = v; }
int ImFontAtlas_GetTexWidth(ImFontAtlas *self) { return self->TexWidth; }
void ImFontAtlas_SetTexHeight(ImFontAtlas *ImFontAtlasPtr, int v) { ImFontAtlasPtr->TexHeight = v; }
int ImFontAtlas_GetTexHeight(ImFontAtlas *self) { return self->TexHeight; }
void ImFontAtlas_SetTexUvScale(ImFontAtlas *ImFontAtlasPtr, ImVec2 v) { ImFontAtlasPtr->TexUvScale = v; }
ImVec2 ImFontAtlas_GetTexUvScale(ImFontAtlas *self) { return self->TexUvScale; }
void ImFontAtlas_SetTexUvWhitePixel(ImFontAtlas *ImFontAtlasPtr, ImVec2 v) { ImFontAtlasPtr->TexUvWhitePixel = v; }
ImVec2 ImFontAtlas_GetTexUvWhitePixel(ImFontAtlas *self) { return self->TexUvWhitePixel; }
void ImFontAtlas_SetFonts(ImFontAtlas *ImFontAtlasPtr, ImVector_ImFontPtr v) { ImFontAtlasPtr->Fonts = v; }
ImVector_ImFontPtr ImFontAtlas_GetFonts(ImFontAtlas *self) { return self->Fonts; }
void ImFontAtlas_SetCustomRects(ImFontAtlas *ImFontAtlasPtr, ImVector_ImFontAtlasCustomRect v) { ImFontAtlasPtr->CustomRects = v; }
ImVector_ImFontAtlasCustomRect ImFontAtlas_GetCustomRects(ImFontAtlas *self) { return self->CustomRects; }
void ImFontAtlas_SetConfigData(ImFontAtlas *ImFontAtlasPtr, ImVector_ImFontConfig v) { ImFontAtlasPtr->ConfigData = v; }
ImVector_ImFontConfig ImFontAtlas_GetConfigData(ImFontAtlas *self) { return self->ConfigData; }
void ImFontAtlas_SetFontBuilderIO(ImFontAtlas *ImFontAtlasPtr, const ImFontBuilderIO* v) { ImFontAtlasPtr->FontBuilderIO = v; }
const ImFontBuilderIO* ImFontAtlas_GetFontBuilderIO(ImFontAtlas *self) { return self->FontBuilderIO; }
void ImFontAtlas_SetFontBuilderFlags(ImFontAtlas *ImFontAtlasPtr, unsigned int v) { ImFontAtlasPtr->FontBuilderFlags = v; }
unsigned int ImFontAtlas_GetFontBuilderFlags(ImFontAtlas *self) { return self->FontBuilderFlags; }
void ImFontAtlas_SetPackIdMouseCursors(ImFontAtlas *ImFontAtlasPtr, int v) { ImFontAtlasPtr->PackIdMouseCursors = v; }
int ImFontAtlas_GetPackIdMouseCursors(ImFontAtlas *self) { return self->PackIdMouseCursors; }
void ImFontAtlas_SetPackIdLines(ImFontAtlas *ImFontAtlasPtr, int v) { ImFontAtlasPtr->PackIdLines = v; }
int ImFontAtlas_GetPackIdLines(ImFontAtlas *self) { return self->PackIdLines; }
void ImGuiInputEventMouseWheel_SetWheelX(ImGuiInputEventMouseWheel *ImGuiInputEventMouseWheelPtr, float v) { ImGuiInputEventMouseWheelPtr->WheelX = v; }
float ImGuiInputEventMouseWheel_GetWheelX(ImGuiInputEventMouseWheel *self) { return self->WheelX; }
void ImGuiInputEventMouseWheel_SetWheelY(ImGuiInputEventMouseWheel *ImGuiInputEventMouseWheelPtr, float v) { ImGuiInputEventMouseWheelPtr->WheelY = v; }
float ImGuiInputEventMouseWheel_GetWheelY(ImGuiInputEventMouseWheel *self) { return self->WheelY; }
void ImGuiMetricsConfig_SetShowDebugLog(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowDebugLog = v; }
bool ImGuiMetricsConfig_GetShowDebugLog(ImGuiMetricsConfig *self) { return self->ShowDebugLog; }
void ImGuiMetricsConfig_SetShowStackTool(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowStackTool = v; }
bool ImGuiMetricsConfig_GetShowStackTool(ImGuiMetricsConfig *self) { return self->ShowStackTool; }
void ImGuiMetricsConfig_SetShowWindowsRects(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowWindowsRects = v; }
bool ImGuiMetricsConfig_GetShowWindowsRects(ImGuiMetricsConfig *self) { return self->ShowWindowsRects; }
void ImGuiMetricsConfig_SetShowWindowsBeginOrder(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowWindowsBeginOrder = v; }
bool ImGuiMetricsConfig_GetShowWindowsBeginOrder(ImGuiMetricsConfig *self) { return self->ShowWindowsBeginOrder; }
void ImGuiMetricsConfig_SetShowTablesRects(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowTablesRects = v; }
bool ImGuiMetricsConfig_GetShowTablesRects(ImGuiMetricsConfig *self) { return self->ShowTablesRects; }
void ImGuiMetricsConfig_SetShowDrawCmdMesh(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowDrawCmdMesh = v; }
bool ImGuiMetricsConfig_GetShowDrawCmdMesh(ImGuiMetricsConfig *self) { return self->ShowDrawCmdMesh; }
void ImGuiMetricsConfig_SetShowDrawCmdBoundingBoxes(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowDrawCmdBoundingBoxes = v; }
bool ImGuiMetricsConfig_GetShowDrawCmdBoundingBoxes(ImGuiMetricsConfig *self) { return self->ShowDrawCmdBoundingBoxes; }
void ImGuiMetricsConfig_SetShowDockingNodes(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, bool v) { ImGuiMetricsConfigPtr->ShowDockingNodes = v; }
bool ImGuiMetricsConfig_GetShowDockingNodes(ImGuiMetricsConfig *self) { return self->ShowDockingNodes; }
void ImGuiMetricsConfig_SetShowWindowsRectsType(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, int v) { ImGuiMetricsConfigPtr->ShowWindowsRectsType = v; }
int ImGuiMetricsConfig_GetShowWindowsRectsType(ImGuiMetricsConfig *self) { return self->ShowWindowsRectsType; }
void ImGuiMetricsConfig_SetShowTablesRectsType(ImGuiMetricsConfig *ImGuiMetricsConfigPtr, int v) { ImGuiMetricsConfigPtr->ShowTablesRectsType = v; }
int ImGuiMetricsConfig_GetShowTablesRectsType(ImGuiMetricsConfig *self) { return self->ShowTablesRectsType; }
void ImGuiNavItemData_SetWindow(ImGuiNavItemData *ImGuiNavItemDataPtr, ImGuiWindow* v) { ImGuiNavItemDataPtr->Window = v; }
ImGuiWindow* ImGuiNavItemData_GetWindow(ImGuiNavItemData *self) { return self->Window; }
void ImGuiNavItemData_SetID(ImGuiNavItemData *ImGuiNavItemDataPtr, ImGuiID v) { ImGuiNavItemDataPtr->ID = v; }
ImGuiID ImGuiNavItemData_GetID(ImGuiNavItemData *self) { return self->ID; }
void ImGuiNavItemData_SetFocusScopeId(ImGuiNavItemData *ImGuiNavItemDataPtr, ImGuiID v) { ImGuiNavItemDataPtr->FocusScopeId = v; }
ImGuiID ImGuiNavItemData_GetFocusScopeId(ImGuiNavItemData *self) { return self->FocusScopeId; }
void ImGuiNavItemData_SetRectRel(ImGuiNavItemData *ImGuiNavItemDataPtr, ImRect v) { ImGuiNavItemDataPtr->RectRel = v; }
ImRect ImGuiNavItemData_GetRectRel(ImGuiNavItemData *self) { return self->RectRel; }
void ImGuiNavItemData_SetInFlags(ImGuiNavItemData *ImGuiNavItemDataPtr, ImGuiItemFlags v) { ImGuiNavItemDataPtr->InFlags = v; }
ImGuiItemFlags ImGuiNavItemData_GetInFlags(ImGuiNavItemData *self) { return self->InFlags; }
void ImGuiNavItemData_SetDistBox(ImGuiNavItemData *ImGuiNavItemDataPtr, float v) { ImGuiNavItemDataPtr->DistBox = v; }
float ImGuiNavItemData_GetDistBox(ImGuiNavItemData *self) { return self->DistBox; }
void ImGuiNavItemData_SetDistCenter(ImGuiNavItemData *ImGuiNavItemDataPtr, float v) { ImGuiNavItemDataPtr->DistCenter = v; }
float ImGuiNavItemData_GetDistCenter(ImGuiNavItemData *self) { return self->DistCenter; }
void ImGuiNavItemData_SetDistAxial(ImGuiNavItemData *ImGuiNavItemDataPtr, float v) { ImGuiNavItemDataPtr->DistAxial = v; }
float ImGuiNavItemData_GetDistAxial(ImGuiNavItemData *self) { return self->DistAxial; }
void ImGuiOldColumnData_SetOffsetNorm(ImGuiOldColumnData *ImGuiOldColumnDataPtr, float v) { ImGuiOldColumnDataPtr->OffsetNorm = v; }
float ImGuiOldColumnData_GetOffsetNorm(ImGuiOldColumnData *self) { return self->OffsetNorm; }
void ImGuiOldColumnData_SetOffsetNormBeforeResize(ImGuiOldColumnData *ImGuiOldColumnDataPtr, float v) { ImGuiOldColumnDataPtr->OffsetNormBeforeResize = v; }
float ImGuiOldColumnData_GetOffsetNormBeforeResize(ImGuiOldColumnData *self) { return self->OffsetNormBeforeResize; }
void ImGuiOldColumnData_SetFlags(ImGuiOldColumnData *ImGuiOldColumnDataPtr, ImGuiOldColumnFlags v) { ImGuiOldColumnDataPtr->Flags = v; }
ImGuiOldColumnFlags ImGuiOldColumnData_GetFlags(ImGuiOldColumnData *self) { return self->Flags; }
void ImGuiOldColumnData_SetClipRect(ImGuiOldColumnData *ImGuiOldColumnDataPtr, ImRect v) { ImGuiOldColumnDataPtr->ClipRect = v; }
ImRect ImGuiOldColumnData_GetClipRect(ImGuiOldColumnData *self) { return self->ClipRect; }
void ImGuiShrinkWidthItem_SetIndex(ImGuiShrinkWidthItem *ImGuiShrinkWidthItemPtr, int v) { ImGuiShrinkWidthItemPtr->Index = v; }
int ImGuiShrinkWidthItem_GetIndex(ImGuiShrinkWidthItem *self) { return self->Index; }
void ImGuiShrinkWidthItem_SetWidth(ImGuiShrinkWidthItem *ImGuiShrinkWidthItemPtr, float v) { ImGuiShrinkWidthItemPtr->Width = v; }
float ImGuiShrinkWidthItem_GetWidth(ImGuiShrinkWidthItem *self) { return self->Width; }
void ImGuiShrinkWidthItem_SetInitialWidth(ImGuiShrinkWidthItem *ImGuiShrinkWidthItemPtr, float v) { ImGuiShrinkWidthItemPtr->InitialWidth = v; }
float ImGuiShrinkWidthItem_GetInitialWidth(ImGuiShrinkWidthItem *self) { return self->InitialWidth; }
void ImGuiTable_SetID(ImGuiTable *ImGuiTablePtr, ImGuiID v) { ImGuiTablePtr->ID = v; }
ImGuiID ImGuiTable_GetID(ImGuiTable *self) { return self->ID; }
void ImGuiTable_SetFlags(ImGuiTable *ImGuiTablePtr, ImGuiTableFlags v) { ImGuiTablePtr->Flags = v; }
ImGuiTableFlags ImGuiTable_GetFlags(ImGuiTable *self) { return self->Flags; }
void ImGuiTable_SetRawData(ImGuiTable *ImGuiTablePtr, void* v) { ImGuiTablePtr->RawData = v; }
void* ImGuiTable_GetRawData(ImGuiTable *self) { return self->RawData; }
void ImGuiTable_SetTempData(ImGuiTable *ImGuiTablePtr, ImGuiTableTempData* v) { ImGuiTablePtr->TempData = v; }
ImGuiTableTempData* ImGuiTable_GetTempData(ImGuiTable *self) { return self->TempData; }
void ImGuiTable_SetColumns(ImGuiTable *ImGuiTablePtr, ImSpan_ImGuiTableColumn v) { ImGuiTablePtr->Columns = v; }
ImSpan_ImGuiTableColumn ImGuiTable_GetColumns(ImGuiTable *self) { return self->Columns; }
void ImGuiTable_SetDisplayOrderToIndex(ImGuiTable *ImGuiTablePtr, ImSpan_ImGuiTableColumnIdx v) { ImGuiTablePtr->DisplayOrderToIndex = v; }
ImSpan_ImGuiTableColumnIdx ImGuiTable_GetDisplayOrderToIndex(ImGuiTable *self) { return self->DisplayOrderToIndex; }
void ImGuiTable_SetRowCellData(ImGuiTable *ImGuiTablePtr, ImSpan_ImGuiTableCellData v) { ImGuiTablePtr->RowCellData = v; }
ImSpan_ImGuiTableCellData ImGuiTable_GetRowCellData(ImGuiTable *self) { return self->RowCellData; }
void ImGuiTable_SetEnabledMaskByDisplayOrder(ImGuiTable *ImGuiTablePtr, ImU64 v) { ImGuiTablePtr->EnabledMaskByDisplayOrder = v; }
ImU64 ImGuiTable_GetEnabledMaskByDisplayOrder(ImGuiTable *self) { return self->EnabledMaskByDisplayOrder; }
void ImGuiTable_SetEnabledMaskByIndex(ImGuiTable *ImGuiTablePtr, ImU64 v) { ImGuiTablePtr->EnabledMaskByIndex = v; }
ImU64 ImGuiTable_GetEnabledMaskByIndex(ImGuiTable *self) { return self->EnabledMaskByIndex; }
void ImGuiTable_SetVisibleMaskByIndex(ImGuiTable *ImGuiTablePtr, ImU64 v) { ImGuiTablePtr->VisibleMaskByIndex = v; }
ImU64 ImGuiTable_GetVisibleMaskByIndex(ImGuiTable *self) { return self->VisibleMaskByIndex; }
void ImGuiTable_SetRequestOutputMaskByIndex(ImGuiTable *ImGuiTablePtr, ImU64 v) { ImGuiTablePtr->RequestOutputMaskByIndex = v; }
ImU64 ImGuiTable_GetRequestOutputMaskByIndex(ImGuiTable *self) { return self->RequestOutputMaskByIndex; }
void ImGuiTable_SetSettingsLoadedFlags(ImGuiTable *ImGuiTablePtr, ImGuiTableFlags v) { ImGuiTablePtr->SettingsLoadedFlags = v; }
ImGuiTableFlags ImGuiTable_GetSettingsLoadedFlags(ImGuiTable *self) { return self->SettingsLoadedFlags; }
void ImGuiTable_SetSettingsOffset(ImGuiTable *ImGuiTablePtr, int v) { ImGuiTablePtr->SettingsOffset = v; }
int ImGuiTable_GetSettingsOffset(ImGuiTable *self) { return self->SettingsOffset; }
void ImGuiTable_SetLastFrameActive(ImGuiTable *ImGuiTablePtr, int v) { ImGuiTablePtr->LastFrameActive = v; }
int ImGuiTable_GetLastFrameActive(ImGuiTable *self) { return self->LastFrameActive; }
void ImGuiTable_SetColumnsCount(ImGuiTable *ImGuiTablePtr, int v) { ImGuiTablePtr->ColumnsCount = v; }
int ImGuiTable_GetColumnsCount(ImGuiTable *self) { return self->ColumnsCount; }
void ImGuiTable_SetCurrentRow(ImGuiTable *ImGuiTablePtr, int v) { ImGuiTablePtr->CurrentRow = v; }
int ImGuiTable_GetCurrentRow(ImGuiTable *self) { return self->CurrentRow; }
void ImGuiTable_SetCurrentColumn(ImGuiTable *ImGuiTablePtr, int v) { ImGuiTablePtr->CurrentColumn = v; }
int ImGuiTable_GetCurrentColumn(ImGuiTable *self) { return self->CurrentColumn; }
void ImGuiTable_SetInstanceCurrent(ImGuiTable *ImGuiTablePtr, ImS16 v) { ImGuiTablePtr->InstanceCurrent = v; }
ImS16 ImGuiTable_GetInstanceCurrent(ImGuiTable *self) { return self->InstanceCurrent; }
void ImGuiTable_SetInstanceInteracted(ImGuiTable *ImGuiTablePtr, ImS16 v) { ImGuiTablePtr->InstanceInteracted = v; }
ImS16 ImGuiTable_GetInstanceInteracted(ImGuiTable *self) { return self->InstanceInteracted; }
void ImGuiTable_SetRowPosY1(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->RowPosY1 = v; }
float ImGuiTable_GetRowPosY1(ImGuiTable *self) { return self->RowPosY1; }
void ImGuiTable_SetRowPosY2(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->RowPosY2 = v; }
float ImGuiTable_GetRowPosY2(ImGuiTable *self) { return self->RowPosY2; }
void ImGuiTable_SetRowMinHeight(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->RowMinHeight = v; }
float ImGuiTable_GetRowMinHeight(ImGuiTable *self) { return self->RowMinHeight; }
void ImGuiTable_SetRowTextBaseline(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->RowTextBaseline = v; }
float ImGuiTable_GetRowTextBaseline(ImGuiTable *self) { return self->RowTextBaseline; }
void ImGuiTable_SetRowIndentOffsetX(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->RowIndentOffsetX = v; }
float ImGuiTable_GetRowIndentOffsetX(ImGuiTable *self) { return self->RowIndentOffsetX; }
void ImGuiTable_SetRowFlags(ImGuiTable *ImGuiTablePtr, ImGuiTableRowFlags v) { ImGuiTablePtr->RowFlags = v; }
ImGuiTableRowFlags ImGuiTable_GetRowFlags(ImGuiTable *self) { return self->RowFlags; }
void ImGuiTable_SetLastRowFlags(ImGuiTable *ImGuiTablePtr, ImGuiTableRowFlags v) { ImGuiTablePtr->LastRowFlags = v; }
ImGuiTableRowFlags ImGuiTable_GetLastRowFlags(ImGuiTable *self) { return self->LastRowFlags; }
void ImGuiTable_SetRowBgColorCounter(ImGuiTable *ImGuiTablePtr, int v) { ImGuiTablePtr->RowBgColorCounter = v; }
int ImGuiTable_GetRowBgColorCounter(ImGuiTable *self) { return self->RowBgColorCounter; }
void ImGuiTable_SetBorderColorStrong(ImGuiTable *ImGuiTablePtr, ImU32 v) { ImGuiTablePtr->BorderColorStrong = v; }
ImU32 ImGuiTable_GetBorderColorStrong(ImGuiTable *self) { return self->BorderColorStrong; }
void ImGuiTable_SetBorderColorLight(ImGuiTable *ImGuiTablePtr, ImU32 v) { ImGuiTablePtr->BorderColorLight = v; }
ImU32 ImGuiTable_GetBorderColorLight(ImGuiTable *self) { return self->BorderColorLight; }
void ImGuiTable_SetBorderX1(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->BorderX1 = v; }
float ImGuiTable_GetBorderX1(ImGuiTable *self) { return self->BorderX1; }
void ImGuiTable_SetBorderX2(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->BorderX2 = v; }
float ImGuiTable_GetBorderX2(ImGuiTable *self) { return self->BorderX2; }
void ImGuiTable_SetHostIndentX(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->HostIndentX = v; }
float ImGuiTable_GetHostIndentX(ImGuiTable *self) { return self->HostIndentX; }
void ImGuiTable_SetMinColumnWidth(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->MinColumnWidth = v; }
float ImGuiTable_GetMinColumnWidth(ImGuiTable *self) { return self->MinColumnWidth; }
void ImGuiTable_SetOuterPaddingX(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->OuterPaddingX = v; }
float ImGuiTable_GetOuterPaddingX(ImGuiTable *self) { return self->OuterPaddingX; }
void ImGuiTable_SetCellPaddingX(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->CellPaddingX = v; }
float ImGuiTable_GetCellPaddingX(ImGuiTable *self) { return self->CellPaddingX; }
void ImGuiTable_SetCellPaddingY(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->CellPaddingY = v; }
float ImGuiTable_GetCellPaddingY(ImGuiTable *self) { return self->CellPaddingY; }
void ImGuiTable_SetCellSpacingX1(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->CellSpacingX1 = v; }
float ImGuiTable_GetCellSpacingX1(ImGuiTable *self) { return self->CellSpacingX1; }
void ImGuiTable_SetCellSpacingX2(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->CellSpacingX2 = v; }
float ImGuiTable_GetCellSpacingX2(ImGuiTable *self) { return self->CellSpacingX2; }
void ImGuiTable_SetInnerWidth(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->InnerWidth = v; }
float ImGuiTable_GetInnerWidth(ImGuiTable *self) { return self->InnerWidth; }
void ImGuiTable_SetColumnsGivenWidth(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->ColumnsGivenWidth = v; }
float ImGuiTable_GetColumnsGivenWidth(ImGuiTable *self) { return self->ColumnsGivenWidth; }
void ImGuiTable_SetColumnsAutoFitWidth(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->ColumnsAutoFitWidth = v; }
float ImGuiTable_GetColumnsAutoFitWidth(ImGuiTable *self) { return self->ColumnsAutoFitWidth; }
void ImGuiTable_SetColumnsStretchSumWeights(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->ColumnsStretchSumWeights = v; }
float ImGuiTable_GetColumnsStretchSumWeights(ImGuiTable *self) { return self->ColumnsStretchSumWeights; }
void ImGuiTable_SetResizedColumnNextWidth(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->ResizedColumnNextWidth = v; }
float ImGuiTable_GetResizedColumnNextWidth(ImGuiTable *self) { return self->ResizedColumnNextWidth; }
void ImGuiTable_SetResizeLockMinContentsX2(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->ResizeLockMinContentsX2 = v; }
float ImGuiTable_GetResizeLockMinContentsX2(ImGuiTable *self) { return self->ResizeLockMinContentsX2; }
void ImGuiTable_SetRefScale(ImGuiTable *ImGuiTablePtr, float v) { ImGuiTablePtr->RefScale = v; }
float ImGuiTable_GetRefScale(ImGuiTable *self) { return self->RefScale; }
void ImGuiTable_SetOuterRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->OuterRect = v; }
ImRect ImGuiTable_GetOuterRect(ImGuiTable *self) { return self->OuterRect; }
void ImGuiTable_SetInnerRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->InnerRect = v; }
ImRect ImGuiTable_GetInnerRect(ImGuiTable *self) { return self->InnerRect; }
void ImGuiTable_SetWorkRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->WorkRect = v; }
ImRect ImGuiTable_GetWorkRect(ImGuiTable *self) { return self->WorkRect; }
void ImGuiTable_SetInnerClipRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->InnerClipRect = v; }
ImRect ImGuiTable_GetInnerClipRect(ImGuiTable *self) { return self->InnerClipRect; }
void ImGuiTable_SetBgClipRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->BgClipRect = v; }
ImRect ImGuiTable_GetBgClipRect(ImGuiTable *self) { return self->BgClipRect; }
void ImGuiTable_SetBg0ClipRectForDrawCmd(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->Bg0ClipRectForDrawCmd = v; }
ImRect ImGuiTable_GetBg0ClipRectForDrawCmd(ImGuiTable *self) { return self->Bg0ClipRectForDrawCmd; }
void ImGuiTable_SetBg2ClipRectForDrawCmd(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->Bg2ClipRectForDrawCmd = v; }
ImRect ImGuiTable_GetBg2ClipRectForDrawCmd(ImGuiTable *self) { return self->Bg2ClipRectForDrawCmd; }
void ImGuiTable_SetHostClipRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->HostClipRect = v; }
ImRect ImGuiTable_GetHostClipRect(ImGuiTable *self) { return self->HostClipRect; }
void ImGuiTable_SetHostBackupInnerClipRect(ImGuiTable *ImGuiTablePtr, ImRect v) { ImGuiTablePtr->HostBackupInnerClipRect = v; }
ImRect ImGuiTable_GetHostBackupInnerClipRect(ImGuiTable *self) { return self->HostBackupInnerClipRect; }
void ImGuiTable_SetOuterWindow(ImGuiTable *ImGuiTablePtr, ImGuiWindow* v) { ImGuiTablePtr->OuterWindow = v; }
ImGuiWindow* ImGuiTable_GetOuterWindow(ImGuiTable *self) { return self->OuterWindow; }
void ImGuiTable_SetInnerWindow(ImGuiTable *ImGuiTablePtr, ImGuiWindow* v) { ImGuiTablePtr->InnerWindow = v; }
ImGuiWindow* ImGuiTable_GetInnerWindow(ImGuiTable *self) { return self->InnerWindow; }
void ImGuiTable_SetColumnsNames(ImGuiTable *ImGuiTablePtr, ImGuiTextBuffer v) { ImGuiTablePtr->ColumnsNames = v; }
ImGuiTextBuffer ImGuiTable_GetColumnsNames(ImGuiTable *self) { return self->ColumnsNames; }
void ImGuiTable_SetDrawSplitter(ImGuiTable *ImGuiTablePtr, ImDrawListSplitter* v) { ImGuiTablePtr->DrawSplitter = v; }
ImDrawListSplitter* ImGuiTable_GetDrawSplitter(ImGuiTable *self) { return self->DrawSplitter; }
void ImGuiTable_SetInstanceDataFirst(ImGuiTable *ImGuiTablePtr, ImGuiTableInstanceData v) { ImGuiTablePtr->InstanceDataFirst = v; }
ImGuiTableInstanceData ImGuiTable_GetInstanceDataFirst(ImGuiTable *self) { return self->InstanceDataFirst; }
void ImGuiTable_SetInstanceDataExtra(ImGuiTable *ImGuiTablePtr, ImVector_ImGuiTableInstanceData v) { ImGuiTablePtr->InstanceDataExtra = v; }
ImVector_ImGuiTableInstanceData ImGuiTable_GetInstanceDataExtra(ImGuiTable *self) { return self->InstanceDataExtra; }
void ImGuiTable_SetSortSpecsSingle(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnSortSpecs v) { ImGuiTablePtr->SortSpecsSingle = v; }
ImGuiTableColumnSortSpecs ImGuiTable_GetSortSpecsSingle(ImGuiTable *self) { return self->SortSpecsSingle; }
void ImGuiTable_SetSortSpecsMulti(ImGuiTable *ImGuiTablePtr, ImVector_ImGuiTableColumnSortSpecs v) { ImGuiTablePtr->SortSpecsMulti = v; }
ImVector_ImGuiTableColumnSortSpecs ImGuiTable_GetSortSpecsMulti(ImGuiTable *self) { return self->SortSpecsMulti; }
void ImGuiTable_SetSortSpecs(ImGuiTable *ImGuiTablePtr, ImGuiTableSortSpecs v) { ImGuiTablePtr->SortSpecs = v; }
ImGuiTableSortSpecs ImGuiTable_GetSortSpecs(ImGuiTable *self) { return self->SortSpecs; }
void ImGuiTable_SetSortSpecsCount(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->SortSpecsCount = v; }
ImGuiTableColumnIdx ImGuiTable_GetSortSpecsCount(ImGuiTable *self) { return self->SortSpecsCount; }
void ImGuiTable_SetColumnsEnabledCount(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->ColumnsEnabledCount = v; }
ImGuiTableColumnIdx ImGuiTable_GetColumnsEnabledCount(ImGuiTable *self) { return self->ColumnsEnabledCount; }
void ImGuiTable_SetColumnsEnabledFixedCount(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->ColumnsEnabledFixedCount = v; }
ImGuiTableColumnIdx ImGuiTable_GetColumnsEnabledFixedCount(ImGuiTable *self) { return self->ColumnsEnabledFixedCount; }
void ImGuiTable_SetDeclColumnsCount(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->DeclColumnsCount = v; }
ImGuiTableColumnIdx ImGuiTable_GetDeclColumnsCount(ImGuiTable *self) { return self->DeclColumnsCount; }
void ImGuiTable_SetHoveredColumnBody(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->HoveredColumnBody = v; }
ImGuiTableColumnIdx ImGuiTable_GetHoveredColumnBody(ImGuiTable *self) { return self->HoveredColumnBody; }
void ImGuiTable_SetHoveredColumnBorder(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->HoveredColumnBorder = v; }
ImGuiTableColumnIdx ImGuiTable_GetHoveredColumnBorder(ImGuiTable *self) { return self->HoveredColumnBorder; }
void ImGuiTable_SetAutoFitSingleColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->AutoFitSingleColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetAutoFitSingleColumn(ImGuiTable *self) { return self->AutoFitSingleColumn; }
void ImGuiTable_SetResizedColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->ResizedColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetResizedColumn(ImGuiTable *self) { return self->ResizedColumn; }
void ImGuiTable_SetLastResizedColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->LastResizedColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetLastResizedColumn(ImGuiTable *self) { return self->LastResizedColumn; }
void ImGuiTable_SetHeldHeaderColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->HeldHeaderColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetHeldHeaderColumn(ImGuiTable *self) { return self->HeldHeaderColumn; }
void ImGuiTable_SetReorderColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->ReorderColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetReorderColumn(ImGuiTable *self) { return self->ReorderColumn; }
void ImGuiTable_SetReorderColumnDir(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->ReorderColumnDir = v; }
ImGuiTableColumnIdx ImGuiTable_GetReorderColumnDir(ImGuiTable *self) { return self->ReorderColumnDir; }
void ImGuiTable_SetLeftMostEnabledColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->LeftMostEnabledColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetLeftMostEnabledColumn(ImGuiTable *self) { return self->LeftMostEnabledColumn; }
void ImGuiTable_SetRightMostEnabledColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->RightMostEnabledColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetRightMostEnabledColumn(ImGuiTable *self) { return self->RightMostEnabledColumn; }
void ImGuiTable_SetLeftMostStretchedColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->LeftMostStretchedColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetLeftMostStretchedColumn(ImGuiTable *self) { return self->LeftMostStretchedColumn; }
void ImGuiTable_SetRightMostStretchedColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->RightMostStretchedColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetRightMostStretchedColumn(ImGuiTable *self) { return self->RightMostStretchedColumn; }
void ImGuiTable_SetContextPopupColumn(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->ContextPopupColumn = v; }
ImGuiTableColumnIdx ImGuiTable_GetContextPopupColumn(ImGuiTable *self) { return self->ContextPopupColumn; }
void ImGuiTable_SetFreezeRowsRequest(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->FreezeRowsRequest = v; }
ImGuiTableColumnIdx ImGuiTable_GetFreezeRowsRequest(ImGuiTable *self) { return self->FreezeRowsRequest; }
void ImGuiTable_SetFreezeRowsCount(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->FreezeRowsCount = v; }
ImGuiTableColumnIdx ImGuiTable_GetFreezeRowsCount(ImGuiTable *self) { return self->FreezeRowsCount; }
void ImGuiTable_SetFreezeColumnsRequest(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->FreezeColumnsRequest = v; }
ImGuiTableColumnIdx ImGuiTable_GetFreezeColumnsRequest(ImGuiTable *self) { return self->FreezeColumnsRequest; }
void ImGuiTable_SetFreezeColumnsCount(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->FreezeColumnsCount = v; }
ImGuiTableColumnIdx ImGuiTable_GetFreezeColumnsCount(ImGuiTable *self) { return self->FreezeColumnsCount; }
void ImGuiTable_SetRowCellDataCurrent(ImGuiTable *ImGuiTablePtr, ImGuiTableColumnIdx v) { ImGuiTablePtr->RowCellDataCurrent = v; }
ImGuiTableColumnIdx ImGuiTable_GetRowCellDataCurrent(ImGuiTable *self) { return self->RowCellDataCurrent; }
void ImGuiTable_SetDummyDrawChannel(ImGuiTable *ImGuiTablePtr, ImGuiTableDrawChannelIdx v) { ImGuiTablePtr->DummyDrawChannel = v; }
ImGuiTableDrawChannelIdx ImGuiTable_GetDummyDrawChannel(ImGuiTable *self) { return self->DummyDrawChannel; }
void ImGuiTable_SetBg2DrawChannelCurrent(ImGuiTable *ImGuiTablePtr, ImGuiTableDrawChannelIdx v) { ImGuiTablePtr->Bg2DrawChannelCurrent = v; }
ImGuiTableDrawChannelIdx ImGuiTable_GetBg2DrawChannelCurrent(ImGuiTable *self) { return self->Bg2DrawChannelCurrent; }
void ImGuiTable_SetBg2DrawChannelUnfrozen(ImGuiTable *ImGuiTablePtr, ImGuiTableDrawChannelIdx v) { ImGuiTablePtr->Bg2DrawChannelUnfrozen = v; }
ImGuiTableDrawChannelIdx ImGuiTable_GetBg2DrawChannelUnfrozen(ImGuiTable *self) { return self->Bg2DrawChannelUnfrozen; }
void ImGuiTable_SetIsLayoutLocked(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsLayoutLocked = v; }
bool ImGuiTable_GetIsLayoutLocked(ImGuiTable *self) { return self->IsLayoutLocked; }
void ImGuiTable_SetIsInsideRow(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsInsideRow = v; }
bool ImGuiTable_GetIsInsideRow(ImGuiTable *self) { return self->IsInsideRow; }
void ImGuiTable_SetIsInitializing(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsInitializing = v; }
bool ImGuiTable_GetIsInitializing(ImGuiTable *self) { return self->IsInitializing; }
void ImGuiTable_SetIsSortSpecsDirty(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsSortSpecsDirty = v; }
bool ImGuiTable_GetIsSortSpecsDirty(ImGuiTable *self) { return self->IsSortSpecsDirty; }
void ImGuiTable_SetIsUsingHeaders(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsUsingHeaders = v; }
bool ImGuiTable_GetIsUsingHeaders(ImGuiTable *self) { return self->IsUsingHeaders; }
void ImGuiTable_SetIsContextPopupOpen(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsContextPopupOpen = v; }
bool ImGuiTable_GetIsContextPopupOpen(ImGuiTable *self) { return self->IsContextPopupOpen; }
void ImGuiTable_SetIsSettingsRequestLoad(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsSettingsRequestLoad = v; }
bool ImGuiTable_GetIsSettingsRequestLoad(ImGuiTable *self) { return self->IsSettingsRequestLoad; }
void ImGuiTable_SetIsSettingsDirty(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsSettingsDirty = v; }
bool ImGuiTable_GetIsSettingsDirty(ImGuiTable *self) { return self->IsSettingsDirty; }
void ImGuiTable_SetIsDefaultDisplayOrder(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsDefaultDisplayOrder = v; }
bool ImGuiTable_GetIsDefaultDisplayOrder(ImGuiTable *self) { return self->IsDefaultDisplayOrder; }
void ImGuiTable_SetIsResetAllRequest(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsResetAllRequest = v; }
bool ImGuiTable_GetIsResetAllRequest(ImGuiTable *self) { return self->IsResetAllRequest; }
void ImGuiTable_SetIsResetDisplayOrderRequest(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsResetDisplayOrderRequest = v; }
bool ImGuiTable_GetIsResetDisplayOrderRequest(ImGuiTable *self) { return self->IsResetDisplayOrderRequest; }
void ImGuiTable_SetIsUnfrozenRows(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsUnfrozenRows = v; }
bool ImGuiTable_GetIsUnfrozenRows(ImGuiTable *self) { return self->IsUnfrozenRows; }
void ImGuiTable_SetIsDefaultSizingPolicy(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->IsDefaultSizingPolicy = v; }
bool ImGuiTable_GetIsDefaultSizingPolicy(ImGuiTable *self) { return self->IsDefaultSizingPolicy; }
void ImGuiTable_SetMemoryCompacted(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->MemoryCompacted = v; }
bool ImGuiTable_GetMemoryCompacted(ImGuiTable *self) { return self->MemoryCompacted; }
void ImGuiTable_SetHostSkipItems(ImGuiTable *ImGuiTablePtr, bool v) { ImGuiTablePtr->HostSkipItems = v; }
bool ImGuiTable_GetHostSkipItems(ImGuiTable *self) { return self->HostSkipItems; }
void ImGuiColorMod_SetCol(ImGuiColorMod *ImGuiColorModPtr, ImGuiCol v) { ImGuiColorModPtr->Col = v; }
ImGuiCol ImGuiColorMod_GetCol(ImGuiColorMod *self) { return self->Col; }
void ImGuiColorMod_SetBackupValue(ImGuiColorMod *ImGuiColorModPtr, ImVec4 v) { ImGuiColorModPtr->BackupValue = v; }
ImVec4 ImGuiColorMod_GetBackupValue(ImGuiColorMod *self) { return self->BackupValue; }
void ImGuiDockNode_SetID(ImGuiDockNode *ImGuiDockNodePtr, ImGuiID v) { ImGuiDockNodePtr->ID = v; }
ImGuiID ImGuiDockNode_GetID(ImGuiDockNode *self) { return self->ID; }
void ImGuiDockNode_SetSharedFlags(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNodeFlags v) { ImGuiDockNodePtr->SharedFlags = v; }
ImGuiDockNodeFlags ImGuiDockNode_GetSharedFlags(ImGuiDockNode *self) { return self->SharedFlags; }
void ImGuiDockNode_SetLocalFlagsInWindows(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNodeFlags v) { ImGuiDockNodePtr->LocalFlagsInWindows = v; }
ImGuiDockNodeFlags ImGuiDockNode_GetLocalFlagsInWindows(ImGuiDockNode *self) { return self->LocalFlagsInWindows; }
void ImGuiDockNode_SetMergedFlags(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNodeFlags v) { ImGuiDockNodePtr->MergedFlags = v; }
ImGuiDockNodeFlags ImGuiDockNode_GetMergedFlags(ImGuiDockNode *self) { return self->MergedFlags; }
void ImGuiDockNode_SetState(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNodeState v) { ImGuiDockNodePtr->State = v; }
ImGuiDockNodeState ImGuiDockNode_GetState(ImGuiDockNode *self) { return self->State; }
void ImGuiDockNode_SetParentNode(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNode* v) { ImGuiDockNodePtr->ParentNode = v; }
ImGuiDockNode* ImGuiDockNode_GetParentNode(ImGuiDockNode *self) { return self->ParentNode; }
void ImGuiDockNode_SetWindows(ImGuiDockNode *ImGuiDockNodePtr, ImVector_ImGuiWindowPtr v) { ImGuiDockNodePtr->Windows = v; }
ImVector_ImGuiWindowPtr ImGuiDockNode_GetWindows(ImGuiDockNode *self) { return self->Windows; }
void ImGuiDockNode_SetTabBar(ImGuiDockNode *ImGuiDockNodePtr, ImGuiTabBar* v) { ImGuiDockNodePtr->TabBar = v; }
ImGuiTabBar* ImGuiDockNode_GetTabBar(ImGuiDockNode *self) { return self->TabBar; }
void ImGuiDockNode_SetPos(ImGuiDockNode *ImGuiDockNodePtr, ImVec2 v) { ImGuiDockNodePtr->Pos = v; }
ImVec2 ImGuiDockNode_GetPos(ImGuiDockNode *self) { return self->Pos; }
void ImGuiDockNode_SetSize(ImGuiDockNode *ImGuiDockNodePtr, ImVec2 v) { ImGuiDockNodePtr->Size = v; }
ImVec2 ImGuiDockNode_GetSize(ImGuiDockNode *self) { return self->Size; }
void ImGuiDockNode_SetSizeRef(ImGuiDockNode *ImGuiDockNodePtr, ImVec2 v) { ImGuiDockNodePtr->SizeRef = v; }
ImVec2 ImGuiDockNode_GetSizeRef(ImGuiDockNode *self) { return self->SizeRef; }
void ImGuiDockNode_SetSplitAxis(ImGuiDockNode *ImGuiDockNodePtr, ImGuiAxis v) { ImGuiDockNodePtr->SplitAxis = v; }
ImGuiAxis ImGuiDockNode_GetSplitAxis(ImGuiDockNode *self) { return self->SplitAxis; }
void ImGuiDockNode_SetWindowClass(ImGuiDockNode *ImGuiDockNodePtr, ImGuiWindowClass v) { ImGuiDockNodePtr->WindowClass = v; }
ImGuiWindowClass ImGuiDockNode_GetWindowClass(ImGuiDockNode *self) { return self->WindowClass; }
void ImGuiDockNode_SetLastBgColor(ImGuiDockNode *ImGuiDockNodePtr, ImU32 v) { ImGuiDockNodePtr->LastBgColor = v; }
ImU32 ImGuiDockNode_GetLastBgColor(ImGuiDockNode *self) { return self->LastBgColor; }
void ImGuiDockNode_SetHostWindow(ImGuiDockNode *ImGuiDockNodePtr, ImGuiWindow* v) { ImGuiDockNodePtr->HostWindow = v; }
ImGuiWindow* ImGuiDockNode_GetHostWindow(ImGuiDockNode *self) { return self->HostWindow; }
void ImGuiDockNode_SetVisibleWindow(ImGuiDockNode *ImGuiDockNodePtr, ImGuiWindow* v) { ImGuiDockNodePtr->VisibleWindow = v; }
ImGuiWindow* ImGuiDockNode_GetVisibleWindow(ImGuiDockNode *self) { return self->VisibleWindow; }
void ImGuiDockNode_SetCentralNode(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNode* v) { ImGuiDockNodePtr->CentralNode = v; }
ImGuiDockNode* ImGuiDockNode_GetCentralNode(ImGuiDockNode *self) { return self->CentralNode; }
void ImGuiDockNode_SetOnlyNodeWithWindows(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDockNode* v) { ImGuiDockNodePtr->OnlyNodeWithWindows = v; }
ImGuiDockNode* ImGuiDockNode_GetOnlyNodeWithWindows(ImGuiDockNode *self) { return self->OnlyNodeWithWindows; }
void ImGuiDockNode_SetCountNodeWithWindows(ImGuiDockNode *ImGuiDockNodePtr, int v) { ImGuiDockNodePtr->CountNodeWithWindows = v; }
int ImGuiDockNode_GetCountNodeWithWindows(ImGuiDockNode *self) { return self->CountNodeWithWindows; }
void ImGuiDockNode_SetLastFrameAlive(ImGuiDockNode *ImGuiDockNodePtr, int v) { ImGuiDockNodePtr->LastFrameAlive = v; }
int ImGuiDockNode_GetLastFrameAlive(ImGuiDockNode *self) { return self->LastFrameAlive; }
void ImGuiDockNode_SetLastFrameActive(ImGuiDockNode *ImGuiDockNodePtr, int v) { ImGuiDockNodePtr->LastFrameActive = v; }
int ImGuiDockNode_GetLastFrameActive(ImGuiDockNode *self) { return self->LastFrameActive; }
void ImGuiDockNode_SetLastFrameFocused(ImGuiDockNode *ImGuiDockNodePtr, int v) { ImGuiDockNodePtr->LastFrameFocused = v; }
int ImGuiDockNode_GetLastFrameFocused(ImGuiDockNode *self) { return self->LastFrameFocused; }
void ImGuiDockNode_SetLastFocusedNodeId(ImGuiDockNode *ImGuiDockNodePtr, ImGuiID v) { ImGuiDockNodePtr->LastFocusedNodeId = v; }
ImGuiID ImGuiDockNode_GetLastFocusedNodeId(ImGuiDockNode *self) { return self->LastFocusedNodeId; }
void ImGuiDockNode_SetSelectedTabId(ImGuiDockNode *ImGuiDockNodePtr, ImGuiID v) { ImGuiDockNodePtr->SelectedTabId = v; }
ImGuiID ImGuiDockNode_GetSelectedTabId(ImGuiDockNode *self) { return self->SelectedTabId; }
void ImGuiDockNode_SetWantCloseTabId(ImGuiDockNode *ImGuiDockNodePtr, ImGuiID v) { ImGuiDockNodePtr->WantCloseTabId = v; }
ImGuiID ImGuiDockNode_GetWantCloseTabId(ImGuiDockNode *self) { return self->WantCloseTabId; }
void ImGuiDockNode_SetAuthorityForPos(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDataAuthority v) { ImGuiDockNodePtr->AuthorityForPos = v; }
ImGuiDataAuthority ImGuiDockNode_GetAuthorityForPos(ImGuiDockNode *self) { return self->AuthorityForPos; }
void ImGuiDockNode_SetAuthorityForSize(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDataAuthority v) { ImGuiDockNodePtr->AuthorityForSize = v; }
ImGuiDataAuthority ImGuiDockNode_GetAuthorityForSize(ImGuiDockNode *self) { return self->AuthorityForSize; }
void ImGuiDockNode_SetAuthorityForViewport(ImGuiDockNode *ImGuiDockNodePtr, ImGuiDataAuthority v) { ImGuiDockNodePtr->AuthorityForViewport = v; }
ImGuiDataAuthority ImGuiDockNode_GetAuthorityForViewport(ImGuiDockNode *self) { return self->AuthorityForViewport; }
void ImGuiDockNode_SetIsVisible(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->IsVisible = v; }
bool ImGuiDockNode_GetIsVisible(ImGuiDockNode *self) { return self->IsVisible; }
void ImGuiDockNode_SetIsFocused(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->IsFocused = v; }
bool ImGuiDockNode_GetIsFocused(ImGuiDockNode *self) { return self->IsFocused; }
void ImGuiDockNode_SetIsBgDrawnThisFrame(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->IsBgDrawnThisFrame = v; }
bool ImGuiDockNode_GetIsBgDrawnThisFrame(ImGuiDockNode *self) { return self->IsBgDrawnThisFrame; }
void ImGuiDockNode_SetHasCloseButton(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->HasCloseButton = v; }
bool ImGuiDockNode_GetHasCloseButton(ImGuiDockNode *self) { return self->HasCloseButton; }
void ImGuiDockNode_SetHasWindowMenuButton(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->HasWindowMenuButton = v; }
bool ImGuiDockNode_GetHasWindowMenuButton(ImGuiDockNode *self) { return self->HasWindowMenuButton; }
void ImGuiDockNode_SetHasCentralNodeChild(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->HasCentralNodeChild = v; }
bool ImGuiDockNode_GetHasCentralNodeChild(ImGuiDockNode *self) { return self->HasCentralNodeChild; }
void ImGuiDockNode_SetWantCloseAll(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->WantCloseAll = v; }
bool ImGuiDockNode_GetWantCloseAll(ImGuiDockNode *self) { return self->WantCloseAll; }
void ImGuiDockNode_SetWantLockSizeOnce(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->WantLockSizeOnce = v; }
bool ImGuiDockNode_GetWantLockSizeOnce(ImGuiDockNode *self) { return self->WantLockSizeOnce; }
void ImGuiDockNode_SetWantMouseMove(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->WantMouseMove = v; }
bool ImGuiDockNode_GetWantMouseMove(ImGuiDockNode *self) { return self->WantMouseMove; }
void ImGuiDockNode_SetWantHiddenTabBarUpdate(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->WantHiddenTabBarUpdate = v; }
bool ImGuiDockNode_GetWantHiddenTabBarUpdate(ImGuiDockNode *self) { return self->WantHiddenTabBarUpdate; }
void ImGuiDockNode_SetWantHiddenTabBarToggle(ImGuiDockNode *ImGuiDockNodePtr, bool v) { ImGuiDockNodePtr->WantHiddenTabBarToggle = v; }
bool ImGuiDockNode_GetWantHiddenTabBarToggle(ImGuiDockNode *self) { return self->WantHiddenTabBarToggle; }
void ImGuiKeyData_SetDown(ImGuiKeyData *ImGuiKeyDataPtr, bool v) { ImGuiKeyDataPtr->Down = v; }
bool ImGuiKeyData_GetDown(ImGuiKeyData *self) { return self->Down; }
void ImGuiKeyData_SetDownDuration(ImGuiKeyData *ImGuiKeyDataPtr, float v) { ImGuiKeyDataPtr->DownDuration = v; }
float ImGuiKeyData_GetDownDuration(ImGuiKeyData *self) { return self->DownDuration; }
void ImGuiKeyData_SetDownDurationPrev(ImGuiKeyData *ImGuiKeyDataPtr, float v) { ImGuiKeyDataPtr->DownDurationPrev = v; }
float ImGuiKeyData_GetDownDurationPrev(ImGuiKeyData *self) { return self->DownDurationPrev; }
void ImGuiKeyData_SetAnalogValue(ImGuiKeyData *ImGuiKeyDataPtr, float v) { ImGuiKeyDataPtr->AnalogValue = v; }
float ImGuiKeyData_GetAnalogValue(ImGuiKeyData *self) { return self->AnalogValue; }
void ImGuiListClipperRange_SetMin(ImGuiListClipperRange *ImGuiListClipperRangePtr, int v) { ImGuiListClipperRangePtr->Min = v; }
int ImGuiListClipperRange_GetMin(ImGuiListClipperRange *self) { return self->Min; }
void ImGuiListClipperRange_SetMax(ImGuiListClipperRange *ImGuiListClipperRangePtr, int v) { ImGuiListClipperRangePtr->Max = v; }
int ImGuiListClipperRange_GetMax(ImGuiListClipperRange *self) { return self->Max; }
void ImGuiListClipperRange_SetPosToIndexConvert(ImGuiListClipperRange *ImGuiListClipperRangePtr, bool v) { ImGuiListClipperRangePtr->PosToIndexConvert = v; }
bool ImGuiListClipperRange_GetPosToIndexConvert(ImGuiListClipperRange *self) { return self->PosToIndexConvert; }
void ImGuiListClipperRange_SetPosToIndexOffsetMin(ImGuiListClipperRange *ImGuiListClipperRangePtr, ImS8 v) { ImGuiListClipperRangePtr->PosToIndexOffsetMin = v; }
ImS8 ImGuiListClipperRange_GetPosToIndexOffsetMin(ImGuiListClipperRange *self) { return self->PosToIndexOffsetMin; }
void ImGuiListClipperRange_SetPosToIndexOffsetMax(ImGuiListClipperRange *ImGuiListClipperRangePtr, ImS8 v) { ImGuiListClipperRangePtr->PosToIndexOffsetMax = v; }
ImS8 ImGuiListClipperRange_GetPosToIndexOffsetMax(ImGuiListClipperRange *self) { return self->PosToIndexOffsetMax; }
void ImGuiPtrOrIndex_SetPtr(ImGuiPtrOrIndex *ImGuiPtrOrIndexPtr, void* v) { ImGuiPtrOrIndexPtr->Ptr = v; }
void* ImGuiPtrOrIndex_GetPtr(ImGuiPtrOrIndex *self) { return self->Ptr; }
void ImGuiPtrOrIndex_SetIndex(ImGuiPtrOrIndex *ImGuiPtrOrIndexPtr, int v) { ImGuiPtrOrIndexPtr->Index = v; }
int ImGuiPtrOrIndex_GetIndex(ImGuiPtrOrIndex *self) { return self->Index; }
void ImGuiWindowTempData_SetCursorPos(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->CursorPos = v; }
ImVec2 ImGuiWindowTempData_GetCursorPos(ImGuiWindowTempData *self) { return self->CursorPos; }
void ImGuiWindowTempData_SetCursorPosPrevLine(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->CursorPosPrevLine = v; }
ImVec2 ImGuiWindowTempData_GetCursorPosPrevLine(ImGuiWindowTempData *self) { return self->CursorPosPrevLine; }
void ImGuiWindowTempData_SetCursorStartPos(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->CursorStartPos = v; }
ImVec2 ImGuiWindowTempData_GetCursorStartPos(ImGuiWindowTempData *self) { return self->CursorStartPos; }
void ImGuiWindowTempData_SetCursorMaxPos(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->CursorMaxPos = v; }
ImVec2 ImGuiWindowTempData_GetCursorMaxPos(ImGuiWindowTempData *self) { return self->CursorMaxPos; }
void ImGuiWindowTempData_SetIdealMaxPos(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->IdealMaxPos = v; }
ImVec2 ImGuiWindowTempData_GetIdealMaxPos(ImGuiWindowTempData *self) { return self->IdealMaxPos; }
void ImGuiWindowTempData_SetCurrLineSize(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->CurrLineSize = v; }
ImVec2 ImGuiWindowTempData_GetCurrLineSize(ImGuiWindowTempData *self) { return self->CurrLineSize; }
void ImGuiWindowTempData_SetPrevLineSize(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->PrevLineSize = v; }
ImVec2 ImGuiWindowTempData_GetPrevLineSize(ImGuiWindowTempData *self) { return self->PrevLineSize; }
void ImGuiWindowTempData_SetCurrLineTextBaseOffset(ImGuiWindowTempData *ImGuiWindowTempDataPtr, float v) { ImGuiWindowTempDataPtr->CurrLineTextBaseOffset = v; }
float ImGuiWindowTempData_GetCurrLineTextBaseOffset(ImGuiWindowTempData *self) { return self->CurrLineTextBaseOffset; }
void ImGuiWindowTempData_SetPrevLineTextBaseOffset(ImGuiWindowTempData *ImGuiWindowTempDataPtr, float v) { ImGuiWindowTempDataPtr->PrevLineTextBaseOffset = v; }
float ImGuiWindowTempData_GetPrevLineTextBaseOffset(ImGuiWindowTempData *self) { return self->PrevLineTextBaseOffset; }
void ImGuiWindowTempData_SetIsSameLine(ImGuiWindowTempData *ImGuiWindowTempDataPtr, bool v) { ImGuiWindowTempDataPtr->IsSameLine = v; }
bool ImGuiWindowTempData_GetIsSameLine(ImGuiWindowTempData *self) { return self->IsSameLine; }
void ImGuiWindowTempData_SetIndent(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec1 v) { ImGuiWindowTempDataPtr->Indent = v; }
ImVec1 ImGuiWindowTempData_GetIndent(ImGuiWindowTempData *self) { return self->Indent; }
void ImGuiWindowTempData_SetColumnsOffset(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec1 v) { ImGuiWindowTempDataPtr->ColumnsOffset = v; }
ImVec1 ImGuiWindowTempData_GetColumnsOffset(ImGuiWindowTempData *self) { return self->ColumnsOffset; }
void ImGuiWindowTempData_SetGroupOffset(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec1 v) { ImGuiWindowTempDataPtr->GroupOffset = v; }
ImVec1 ImGuiWindowTempData_GetGroupOffset(ImGuiWindowTempData *self) { return self->GroupOffset; }
void ImGuiWindowTempData_SetCursorStartPosLossyness(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->CursorStartPosLossyness = v; }
ImVec2 ImGuiWindowTempData_GetCursorStartPosLossyness(ImGuiWindowTempData *self) { return self->CursorStartPosLossyness; }
void ImGuiWindowTempData_SetNavLayerCurrent(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiNavLayer v) { ImGuiWindowTempDataPtr->NavLayerCurrent = v; }
ImGuiNavLayer ImGuiWindowTempData_GetNavLayerCurrent(ImGuiWindowTempData *self) { return self->NavLayerCurrent; }
void ImGuiWindowTempData_SetNavLayersActiveMask(ImGuiWindowTempData *ImGuiWindowTempDataPtr, short v) { ImGuiWindowTempDataPtr->NavLayersActiveMask = v; }
short ImGuiWindowTempData_GetNavLayersActiveMask(ImGuiWindowTempData *self) { return self->NavLayersActiveMask; }
void ImGuiWindowTempData_SetNavLayersActiveMaskNext(ImGuiWindowTempData *ImGuiWindowTempDataPtr, short v) { ImGuiWindowTempDataPtr->NavLayersActiveMaskNext = v; }
short ImGuiWindowTempData_GetNavLayersActiveMaskNext(ImGuiWindowTempData *self) { return self->NavLayersActiveMaskNext; }
void ImGuiWindowTempData_SetNavFocusScopeIdCurrent(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiID v) { ImGuiWindowTempDataPtr->NavFocusScopeIdCurrent = v; }
ImGuiID ImGuiWindowTempData_GetNavFocusScopeIdCurrent(ImGuiWindowTempData *self) { return self->NavFocusScopeIdCurrent; }
void ImGuiWindowTempData_SetNavHideHighlightOneFrame(ImGuiWindowTempData *ImGuiWindowTempDataPtr, bool v) { ImGuiWindowTempDataPtr->NavHideHighlightOneFrame = v; }
bool ImGuiWindowTempData_GetNavHideHighlightOneFrame(ImGuiWindowTempData *self) { return self->NavHideHighlightOneFrame; }
void ImGuiWindowTempData_SetNavHasScroll(ImGuiWindowTempData *ImGuiWindowTempDataPtr, bool v) { ImGuiWindowTempDataPtr->NavHasScroll = v; }
bool ImGuiWindowTempData_GetNavHasScroll(ImGuiWindowTempData *self) { return self->NavHasScroll; }
void ImGuiWindowTempData_SetMenuBarAppending(ImGuiWindowTempData *ImGuiWindowTempDataPtr, bool v) { ImGuiWindowTempDataPtr->MenuBarAppending = v; }
bool ImGuiWindowTempData_GetMenuBarAppending(ImGuiWindowTempData *self) { return self->MenuBarAppending; }
void ImGuiWindowTempData_SetMenuBarOffset(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVec2 v) { ImGuiWindowTempDataPtr->MenuBarOffset = v; }
ImVec2 ImGuiWindowTempData_GetMenuBarOffset(ImGuiWindowTempData *self) { return self->MenuBarOffset; }
void ImGuiWindowTempData_SetMenuColumns(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiMenuColumns v) { ImGuiWindowTempDataPtr->MenuColumns = v; }
ImGuiMenuColumns ImGuiWindowTempData_GetMenuColumns(ImGuiWindowTempData *self) { return self->MenuColumns; }
void ImGuiWindowTempData_SetTreeDepth(ImGuiWindowTempData *ImGuiWindowTempDataPtr, int v) { ImGuiWindowTempDataPtr->TreeDepth = v; }
int ImGuiWindowTempData_GetTreeDepth(ImGuiWindowTempData *self) { return self->TreeDepth; }
void ImGuiWindowTempData_SetTreeJumpToParentOnPopMask(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImU32 v) { ImGuiWindowTempDataPtr->TreeJumpToParentOnPopMask = v; }
ImU32 ImGuiWindowTempData_GetTreeJumpToParentOnPopMask(ImGuiWindowTempData *self) { return self->TreeJumpToParentOnPopMask; }
void ImGuiWindowTempData_SetChildWindows(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVector_ImGuiWindowPtr v) { ImGuiWindowTempDataPtr->ChildWindows = v; }
ImVector_ImGuiWindowPtr ImGuiWindowTempData_GetChildWindows(ImGuiWindowTempData *self) { return self->ChildWindows; }
void ImGuiWindowTempData_SetStateStorage(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiStorage* v) { ImGuiWindowTempDataPtr->StateStorage = v; }
ImGuiStorage* ImGuiWindowTempData_GetStateStorage(ImGuiWindowTempData *self) { return self->StateStorage; }
void ImGuiWindowTempData_SetCurrentColumns(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiOldColumns* v) { ImGuiWindowTempDataPtr->CurrentColumns = v; }
ImGuiOldColumns* ImGuiWindowTempData_GetCurrentColumns(ImGuiWindowTempData *self) { return self->CurrentColumns; }
void ImGuiWindowTempData_SetCurrentTableIdx(ImGuiWindowTempData *ImGuiWindowTempDataPtr, int v) { ImGuiWindowTempDataPtr->CurrentTableIdx = v; }
int ImGuiWindowTempData_GetCurrentTableIdx(ImGuiWindowTempData *self) { return self->CurrentTableIdx; }
void ImGuiWindowTempData_SetLayoutType(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiLayoutType v) { ImGuiWindowTempDataPtr->LayoutType = v; }
ImGuiLayoutType ImGuiWindowTempData_GetLayoutType(ImGuiWindowTempData *self) { return self->LayoutType; }
void ImGuiWindowTempData_SetParentLayoutType(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImGuiLayoutType v) { ImGuiWindowTempDataPtr->ParentLayoutType = v; }
ImGuiLayoutType ImGuiWindowTempData_GetParentLayoutType(ImGuiWindowTempData *self) { return self->ParentLayoutType; }
void ImGuiWindowTempData_SetItemWidth(ImGuiWindowTempData *ImGuiWindowTempDataPtr, float v) { ImGuiWindowTempDataPtr->ItemWidth = v; }
float ImGuiWindowTempData_GetItemWidth(ImGuiWindowTempData *self) { return self->ItemWidth; }
void ImGuiWindowTempData_SetTextWrapPos(ImGuiWindowTempData *ImGuiWindowTempDataPtr, float v) { ImGuiWindowTempDataPtr->TextWrapPos = v; }
float ImGuiWindowTempData_GetTextWrapPos(ImGuiWindowTempData *self) { return self->TextWrapPos; }
void ImGuiWindowTempData_SetItemWidthStack(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVector_float v) { ImGuiWindowTempDataPtr->ItemWidthStack = v; }
ImVector_float ImGuiWindowTempData_GetItemWidthStack(ImGuiWindowTempData *self) { return self->ItemWidthStack; }
void ImGuiWindowTempData_SetTextWrapPosStack(ImGuiWindowTempData *ImGuiWindowTempDataPtr, ImVector_float v) { ImGuiWindowTempDataPtr->TextWrapPosStack = v; }
ImVector_float ImGuiWindowTempData_GetTextWrapPosStack(ImGuiWindowTempData *self) { return self->TextWrapPosStack; }
void ImVec1_Setx(ImVec1 *ImVec1Ptr, float v) { ImVec1Ptr->x = v; }
float ImVec1_Getx(ImVec1 *self) { return self->x; }
void ImFont_SetIndexAdvanceX(ImFont *ImFontPtr, ImVector_float v) { ImFontPtr->IndexAdvanceX = v; }
ImVector_float ImFont_GetIndexAdvanceX(ImFont *self) { return self->IndexAdvanceX; }
void ImFont_SetFallbackAdvanceX(ImFont *ImFontPtr, float v) { ImFontPtr->FallbackAdvanceX = v; }
float ImFont_GetFallbackAdvanceX(ImFont *self) { return self->FallbackAdvanceX; }
void ImFont_SetFontSize(ImFont *ImFontPtr, float v) { ImFontPtr->FontSize = v; }
float ImFont_GetFontSize(ImFont *self) { return self->FontSize; }
void ImFont_SetIndexLookup(ImFont *ImFontPtr, ImVector_ImWchar v) { ImFontPtr->IndexLookup = v; }
ImVector_ImWchar ImFont_GetIndexLookup(ImFont *self) { return self->IndexLookup; }
void ImFont_SetGlyphs(ImFont *ImFontPtr, ImVector_ImFontGlyph v) { ImFontPtr->Glyphs = v; }
ImVector_ImFontGlyph ImFont_GetGlyphs(ImFont *self) { return self->Glyphs; }
void ImFont_SetFallbackGlyph(ImFont *ImFontPtr, const ImFontGlyph* v) { ImFontPtr->FallbackGlyph = v; }
const ImFontGlyph* ImFont_GetFallbackGlyph(ImFont *self) { return self->FallbackGlyph; }
void ImFont_SetContainerAtlas(ImFont *ImFontPtr, ImFontAtlas* v) { ImFontPtr->ContainerAtlas = v; }
ImFontAtlas* ImFont_GetContainerAtlas(ImFont *self) { return self->ContainerAtlas; }
void ImFont_SetConfigData(ImFont *ImFontPtr, const ImFontConfig* v) { ImFontPtr->ConfigData = v; }
const ImFontConfig* ImFont_GetConfigData(ImFont *self) { return self->ConfigData; }
void ImFont_SetConfigDataCount(ImFont *ImFontPtr, short v) { ImFontPtr->ConfigDataCount = v; }
short ImFont_GetConfigDataCount(ImFont *self) { return self->ConfigDataCount; }
void ImFont_SetFallbackChar(ImFont *ImFontPtr, ImWchar v) { ImFontPtr->FallbackChar = v; }
ImWchar ImFont_GetFallbackChar(ImFont *self) { return self->FallbackChar; }
void ImFont_SetEllipsisChar(ImFont *ImFontPtr, ImWchar v) { ImFontPtr->EllipsisChar = v; }
ImWchar ImFont_GetEllipsisChar(ImFont *self) { return self->EllipsisChar; }
void ImFont_SetDotChar(ImFont *ImFontPtr, ImWchar v) { ImFontPtr->DotChar = v; }
ImWchar ImFont_GetDotChar(ImFont *self) { return self->DotChar; }
void ImFont_SetDirtyLookupTables(ImFont *ImFontPtr, bool v) { ImFontPtr->DirtyLookupTables = v; }
bool ImFont_GetDirtyLookupTables(ImFont *self) { return self->DirtyLookupTables; }
void ImFont_SetScale(ImFont *ImFontPtr, float v) { ImFontPtr->Scale = v; }
float ImFont_GetScale(ImFont *self) { return self->Scale; }
void ImFont_SetAscent(ImFont *ImFontPtr, float v) { ImFontPtr->Ascent = v; }
float ImFont_GetAscent(ImFont *self) { return self->Ascent; }
void ImFont_SetDescent(ImFont *ImFontPtr, float v) { ImFontPtr->Descent = v; }
float ImFont_GetDescent(ImFont *self) { return self->Descent; }
void ImFont_SetMetricsTotalSurface(ImFont *ImFontPtr, int v) { ImFontPtr->MetricsTotalSurface = v; }
int ImFont_GetMetricsTotalSurface(ImFont *self) { return self->MetricsTotalSurface; }
void ImDrawData_SetValid(ImDrawData *ImDrawDataPtr, bool v) { ImDrawDataPtr->Valid = v; }
bool ImDrawData_GetValid(ImDrawData *self) { return self->Valid; }
void ImDrawData_SetCmdListsCount(ImDrawData *ImDrawDataPtr, int v) { ImDrawDataPtr->CmdListsCount = v; }
int ImDrawData_GetCmdListsCount(ImDrawData *self) { return self->CmdListsCount; }
void ImDrawData_SetTotalIdxCount(ImDrawData *ImDrawDataPtr, int v) { ImDrawDataPtr->TotalIdxCount = v; }
int ImDrawData_GetTotalIdxCount(ImDrawData *self) { return self->TotalIdxCount; }
void ImDrawData_SetTotalVtxCount(ImDrawData *ImDrawDataPtr, int v) { ImDrawDataPtr->TotalVtxCount = v; }
int ImDrawData_GetTotalVtxCount(ImDrawData *self) { return self->TotalVtxCount; }
void ImDrawData_SetCmdLists(ImDrawData *ImDrawDataPtr, ImDrawList** v) { ImDrawDataPtr->CmdLists = v; }
ImDrawList** ImDrawData_GetCmdLists(ImDrawData *self) { return self->CmdLists; }
void ImDrawData_SetDisplayPos(ImDrawData *ImDrawDataPtr, ImVec2 v) { ImDrawDataPtr->DisplayPos = v; }
ImVec2 ImDrawData_GetDisplayPos(ImDrawData *self) { return self->DisplayPos; }
void ImDrawData_SetDisplaySize(ImDrawData *ImDrawDataPtr, ImVec2 v) { ImDrawDataPtr->DisplaySize = v; }
ImVec2 ImDrawData_GetDisplaySize(ImDrawData *self) { return self->DisplaySize; }
void ImDrawData_SetFramebufferScale(ImDrawData *ImDrawDataPtr, ImVec2 v) { ImDrawDataPtr->FramebufferScale = v; }
ImVec2 ImDrawData_GetFramebufferScale(ImDrawData *self) { return self->FramebufferScale; }
void ImDrawData_SetOwnerViewport(ImDrawData *ImDrawDataPtr, ImGuiViewport* v) { ImDrawDataPtr->OwnerViewport = v; }
ImGuiViewport* ImDrawData_GetOwnerViewport(ImDrawData *self) { return self->OwnerViewport; }
void ImGuiDataTypeInfo_SetSize(ImGuiDataTypeInfo *ImGuiDataTypeInfoPtr, size_t v) { ImGuiDataTypeInfoPtr->Size = v; }
size_t ImGuiDataTypeInfo_GetSize(ImGuiDataTypeInfo *self) { return self->Size; }
void ImGuiDataTypeInfo_SetName(ImGuiDataTypeInfo *ImGuiDataTypeInfoPtr, const char* v) { ImGuiDataTypeInfoPtr->Name = v; }
const char* ImGuiDataTypeInfo_GetName(ImGuiDataTypeInfo *self) { return self->Name; }
void ImGuiDataTypeInfo_SetPrintFmt(ImGuiDataTypeInfo *ImGuiDataTypeInfoPtr, const char* v) { ImGuiDataTypeInfoPtr->PrintFmt = v; }
const char* ImGuiDataTypeInfo_GetPrintFmt(ImGuiDataTypeInfo *self) { return self->PrintFmt; }
void ImGuiDataTypeInfo_SetScanFmt(ImGuiDataTypeInfo *ImGuiDataTypeInfoPtr, const char* v) { ImGuiDataTypeInfoPtr->ScanFmt = v; }
const char* ImGuiDataTypeInfo_GetScanFmt(ImGuiDataTypeInfo *self) { return self->ScanFmt; }
void ImGuiInputEventMouseButton_SetButton(ImGuiInputEventMouseButton *ImGuiInputEventMouseButtonPtr, int v) { ImGuiInputEventMouseButtonPtr->Button = v; }
int ImGuiInputEventMouseButton_GetButton(ImGuiInputEventMouseButton *self) { return self->Button; }
void ImGuiInputEventMouseButton_SetDown(ImGuiInputEventMouseButton *ImGuiInputEventMouseButtonPtr, bool v) { ImGuiInputEventMouseButtonPtr->Down = v; }
bool ImGuiInputEventMouseButton_GetDown(ImGuiInputEventMouseButton *self) { return self->Down; }
void ImGuiInputEventMousePos_SetPosX(ImGuiInputEventMousePos *ImGuiInputEventMousePosPtr, float v) { ImGuiInputEventMousePosPtr->PosX = v; }
float ImGuiInputEventMousePos_GetPosX(ImGuiInputEventMousePos *self) { return self->PosX; }
void ImGuiInputEventMousePos_SetPosY(ImGuiInputEventMousePos *ImGuiInputEventMousePosPtr, float v) { ImGuiInputEventMousePosPtr->PosY = v; }
float ImGuiInputEventMousePos_GetPosY(ImGuiInputEventMousePos *self) { return self->PosY; }
void ImGuiOldColumns_SetID(ImGuiOldColumns *ImGuiOldColumnsPtr, ImGuiID v) { ImGuiOldColumnsPtr->ID = v; }
ImGuiID ImGuiOldColumns_GetID(ImGuiOldColumns *self) { return self->ID; }
void ImGuiOldColumns_SetFlags(ImGuiOldColumns *ImGuiOldColumnsPtr, ImGuiOldColumnFlags v) { ImGuiOldColumnsPtr->Flags = v; }
ImGuiOldColumnFlags ImGuiOldColumns_GetFlags(ImGuiOldColumns *self) { return self->Flags; }
void ImGuiOldColumns_SetIsFirstFrame(ImGuiOldColumns *ImGuiOldColumnsPtr, bool v) { ImGuiOldColumnsPtr->IsFirstFrame = v; }
bool ImGuiOldColumns_GetIsFirstFrame(ImGuiOldColumns *self) { return self->IsFirstFrame; }
void ImGuiOldColumns_SetIsBeingResized(ImGuiOldColumns *ImGuiOldColumnsPtr, bool v) { ImGuiOldColumnsPtr->IsBeingResized = v; }
bool ImGuiOldColumns_GetIsBeingResized(ImGuiOldColumns *self) { return self->IsBeingResized; }
void ImGuiOldColumns_SetCurrent(ImGuiOldColumns *ImGuiOldColumnsPtr, int v) { ImGuiOldColumnsPtr->Current = v; }
int ImGuiOldColumns_GetCurrent(ImGuiOldColumns *self) { return self->Current; }
void ImGuiOldColumns_SetCount(ImGuiOldColumns *ImGuiOldColumnsPtr, int v) { ImGuiOldColumnsPtr->Count = v; }
int ImGuiOldColumns_GetCount(ImGuiOldColumns *self) { return self->Count; }
void ImGuiOldColumns_SetOffMinX(ImGuiOldColumns *ImGuiOldColumnsPtr, float v) { ImGuiOldColumnsPtr->OffMinX = v; }
float ImGuiOldColumns_GetOffMinX(ImGuiOldColumns *self) { return self->OffMinX; }
void ImGuiOldColumns_SetOffMaxX(ImGuiOldColumns *ImGuiOldColumnsPtr, float v) { ImGuiOldColumnsPtr->OffMaxX = v; }
float ImGuiOldColumns_GetOffMaxX(ImGuiOldColumns *self) { return self->OffMaxX; }
void ImGuiOldColumns_SetLineMinY(ImGuiOldColumns *ImGuiOldColumnsPtr, float v) { ImGuiOldColumnsPtr->LineMinY = v; }
float ImGuiOldColumns_GetLineMinY(ImGuiOldColumns *self) { return self->LineMinY; }
void ImGuiOldColumns_SetLineMaxY(ImGuiOldColumns *ImGuiOldColumnsPtr, float v) { ImGuiOldColumnsPtr->LineMaxY = v; }
float ImGuiOldColumns_GetLineMaxY(ImGuiOldColumns *self) { return self->LineMaxY; }
void ImGuiOldColumns_SetHostCursorPosY(ImGuiOldColumns *ImGuiOldColumnsPtr, float v) { ImGuiOldColumnsPtr->HostCursorPosY = v; }
float ImGuiOldColumns_GetHostCursorPosY(ImGuiOldColumns *self) { return self->HostCursorPosY; }
void ImGuiOldColumns_SetHostCursorMaxPosX(ImGuiOldColumns *ImGuiOldColumnsPtr, float v) { ImGuiOldColumnsPtr->HostCursorMaxPosX = v; }
float ImGuiOldColumns_GetHostCursorMaxPosX(ImGuiOldColumns *self) { return self->HostCursorMaxPosX; }
void ImGuiOldColumns_SetHostInitialClipRect(ImGuiOldColumns *ImGuiOldColumnsPtr, ImRect v) { ImGuiOldColumnsPtr->HostInitialClipRect = v; }
ImRect ImGuiOldColumns_GetHostInitialClipRect(ImGuiOldColumns *self) { return self->HostInitialClipRect; }
void ImGuiOldColumns_SetHostBackupClipRect(ImGuiOldColumns *ImGuiOldColumnsPtr, ImRect v) { ImGuiOldColumnsPtr->HostBackupClipRect = v; }
ImRect ImGuiOldColumns_GetHostBackupClipRect(ImGuiOldColumns *self) { return self->HostBackupClipRect; }
void ImGuiOldColumns_SetHostBackupParentWorkRect(ImGuiOldColumns *ImGuiOldColumnsPtr, ImRect v) { ImGuiOldColumnsPtr->HostBackupParentWorkRect = v; }
ImRect ImGuiOldColumns_GetHostBackupParentWorkRect(ImGuiOldColumns *self) { return self->HostBackupParentWorkRect; }
void ImGuiOldColumns_SetColumns(ImGuiOldColumns *ImGuiOldColumnsPtr, ImVector_ImGuiOldColumnData v) { ImGuiOldColumnsPtr->Columns = v; }
ImVector_ImGuiOldColumnData ImGuiOldColumns_GetColumns(ImGuiOldColumns *self) { return self->Columns; }
void ImGuiOldColumns_SetSplitter(ImGuiOldColumns *ImGuiOldColumnsPtr, ImDrawListSplitter v) { ImGuiOldColumnsPtr->Splitter = v; }
ImDrawListSplitter ImGuiOldColumns_GetSplitter(ImGuiOldColumns *self) { return self->Splitter; }
void ImGuiTableColumnSortSpecs_SetColumnUserID(ImGuiTableColumnSortSpecs *ImGuiTableColumnSortSpecsPtr, ImGuiID v) { ImGuiTableColumnSortSpecsPtr->ColumnUserID = v; }
ImGuiID ImGuiTableColumnSortSpecs_GetColumnUserID(ImGuiTableColumnSortSpecs *self) { return self->ColumnUserID; }
void ImGuiTableColumnSortSpecs_SetColumnIndex(ImGuiTableColumnSortSpecs *ImGuiTableColumnSortSpecsPtr, ImS16 v) { ImGuiTableColumnSortSpecsPtr->ColumnIndex = v; }
ImS16 ImGuiTableColumnSortSpecs_GetColumnIndex(ImGuiTableColumnSortSpecs *self) { return self->ColumnIndex; }
void ImGuiTableColumnSortSpecs_SetSortOrder(ImGuiTableColumnSortSpecs *ImGuiTableColumnSortSpecsPtr, ImS16 v) { ImGuiTableColumnSortSpecsPtr->SortOrder = v; }
ImS16 ImGuiTableColumnSortSpecs_GetSortOrder(ImGuiTableColumnSortSpecs *self) { return self->SortOrder; }
void ImGuiTableColumnSortSpecs_SetSortDirection(ImGuiTableColumnSortSpecs *ImGuiTableColumnSortSpecsPtr, ImGuiSortDirection v) { ImGuiTableColumnSortSpecsPtr->SortDirection = v; }
ImGuiSortDirection ImGuiTableColumnSortSpecs_GetSortDirection(ImGuiTableColumnSortSpecs *self) { return self->SortDirection; }
void ImGuiTableSettings_SetID(ImGuiTableSettings *ImGuiTableSettingsPtr, ImGuiID v) { ImGuiTableSettingsPtr->ID = v; }
ImGuiID ImGuiTableSettings_GetID(ImGuiTableSettings *self) { return self->ID; }
void ImGuiTableSettings_SetSaveFlags(ImGuiTableSettings *ImGuiTableSettingsPtr, ImGuiTableFlags v) { ImGuiTableSettingsPtr->SaveFlags = v; }
ImGuiTableFlags ImGuiTableSettings_GetSaveFlags(ImGuiTableSettings *self) { return self->SaveFlags; }
void ImGuiTableSettings_SetRefScale(ImGuiTableSettings *ImGuiTableSettingsPtr, float v) { ImGuiTableSettingsPtr->RefScale = v; }
float ImGuiTableSettings_GetRefScale(ImGuiTableSettings *self) { return self->RefScale; }
void ImGuiTableSettings_SetColumnsCount(ImGuiTableSettings *ImGuiTableSettingsPtr, ImGuiTableColumnIdx v) { ImGuiTableSettingsPtr->ColumnsCount = v; }
ImGuiTableColumnIdx ImGuiTableSettings_GetColumnsCount(ImGuiTableSettings *self) { return self->ColumnsCount; }
void ImGuiTableSettings_SetColumnsCountMax(ImGuiTableSettings *ImGuiTableSettingsPtr, ImGuiTableColumnIdx v) { ImGuiTableSettingsPtr->ColumnsCountMax = v; }
ImGuiTableColumnIdx ImGuiTableSettings_GetColumnsCountMax(ImGuiTableSettings *self) { return self->ColumnsCountMax; }
void ImGuiTableSettings_SetWantApply(ImGuiTableSettings *ImGuiTableSettingsPtr, bool v) { ImGuiTableSettingsPtr->WantApply = v; }
bool ImGuiTableSettings_GetWantApply(ImGuiTableSettings *self) { return self->WantApply; }
void ImDrawChannel_Set_CmdBuffer(ImDrawChannel *ImDrawChannelPtr, ImVector_ImDrawCmd v) { ImDrawChannelPtr->_CmdBuffer = v; }
ImVector_ImDrawCmd ImDrawChannel_Get_CmdBuffer(ImDrawChannel *self) { return self->_CmdBuffer; }
void ImDrawChannel_Set_IdxBuffer(ImDrawChannel *ImDrawChannelPtr, ImVector_ImDrawIdx v) { ImDrawChannelPtr->_IdxBuffer = v; }
ImVector_ImDrawIdx ImDrawChannel_Get_IdxBuffer(ImDrawChannel *self) { return self->_IdxBuffer; }
void ImGuiTextFilter_SetFilters(ImGuiTextFilter *ImGuiTextFilterPtr, ImVector_ImGuiTextRange v) { ImGuiTextFilterPtr->Filters = v; }
ImVector_ImGuiTextRange ImGuiTextFilter_GetFilters(ImGuiTextFilter *self) { return self->Filters; }
void ImGuiTextFilter_SetCountGrep(ImGuiTextFilter *ImGuiTextFilterPtr, int v) { ImGuiTextFilterPtr->CountGrep = v; }
int ImGuiTextFilter_GetCountGrep(ImGuiTextFilter *self) { return self->CountGrep; }
void ImGuiTableSortSpecs_SetSpecs(ImGuiTableSortSpecs *ImGuiTableSortSpecsPtr, const ImGuiTableColumnSortSpecs* v) { ImGuiTableSortSpecsPtr->Specs = v; }
const ImGuiTableColumnSortSpecs* ImGuiTableSortSpecs_GetSpecs(ImGuiTableSortSpecs *self) { return self->Specs; }
void ImGuiTableSortSpecs_SetSpecsCount(ImGuiTableSortSpecs *ImGuiTableSortSpecsPtr, int v) { ImGuiTableSortSpecsPtr->SpecsCount = v; }
int ImGuiTableSortSpecs_GetSpecsCount(ImGuiTableSortSpecs *self) { return self->SpecsCount; }
void ImGuiTableSortSpecs_SetSpecsDirty(ImGuiTableSortSpecs *ImGuiTableSortSpecsPtr, bool v) { ImGuiTableSortSpecsPtr->SpecsDirty = v; }
bool ImGuiTableSortSpecs_GetSpecsDirty(ImGuiTableSortSpecs *self) { return self->SpecsDirty; }
void ImGuiTabBar_SetTabs(ImGuiTabBar *ImGuiTabBarPtr, ImVector_ImGuiTabItem v) { ImGuiTabBarPtr->Tabs = v; }
ImVector_ImGuiTabItem ImGuiTabBar_GetTabs(ImGuiTabBar *self) { return self->Tabs; }
void ImGuiTabBar_SetFlags(ImGuiTabBar *ImGuiTabBarPtr, ImGuiTabBarFlags v) { ImGuiTabBarPtr->Flags = v; }
ImGuiTabBarFlags ImGuiTabBar_GetFlags(ImGuiTabBar *self) { return self->Flags; }
void ImGuiTabBar_SetID(ImGuiTabBar *ImGuiTabBarPtr, ImGuiID v) { ImGuiTabBarPtr->ID = v; }
ImGuiID ImGuiTabBar_GetID(ImGuiTabBar *self) { return self->ID; }
void ImGuiTabBar_SetSelectedTabId(ImGuiTabBar *ImGuiTabBarPtr, ImGuiID v) { ImGuiTabBarPtr->SelectedTabId = v; }
ImGuiID ImGuiTabBar_GetSelectedTabId(ImGuiTabBar *self) { return self->SelectedTabId; }
void ImGuiTabBar_SetNextSelectedTabId(ImGuiTabBar *ImGuiTabBarPtr, ImGuiID v) { ImGuiTabBarPtr->NextSelectedTabId = v; }
ImGuiID ImGuiTabBar_GetNextSelectedTabId(ImGuiTabBar *self) { return self->NextSelectedTabId; }
void ImGuiTabBar_SetVisibleTabId(ImGuiTabBar *ImGuiTabBarPtr, ImGuiID v) { ImGuiTabBarPtr->VisibleTabId = v; }
ImGuiID ImGuiTabBar_GetVisibleTabId(ImGuiTabBar *self) { return self->VisibleTabId; }
void ImGuiTabBar_SetCurrFrameVisible(ImGuiTabBar *ImGuiTabBarPtr, int v) { ImGuiTabBarPtr->CurrFrameVisible = v; }
int ImGuiTabBar_GetCurrFrameVisible(ImGuiTabBar *self) { return self->CurrFrameVisible; }
void ImGuiTabBar_SetPrevFrameVisible(ImGuiTabBar *ImGuiTabBarPtr, int v) { ImGuiTabBarPtr->PrevFrameVisible = v; }
int ImGuiTabBar_GetPrevFrameVisible(ImGuiTabBar *self) { return self->PrevFrameVisible; }
void ImGuiTabBar_SetBarRect(ImGuiTabBar *ImGuiTabBarPtr, ImRect v) { ImGuiTabBarPtr->BarRect = v; }
ImRect ImGuiTabBar_GetBarRect(ImGuiTabBar *self) { return self->BarRect; }
void ImGuiTabBar_SetCurrTabsContentsHeight(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->CurrTabsContentsHeight = v; }
float ImGuiTabBar_GetCurrTabsContentsHeight(ImGuiTabBar *self) { return self->CurrTabsContentsHeight; }
void ImGuiTabBar_SetPrevTabsContentsHeight(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->PrevTabsContentsHeight = v; }
float ImGuiTabBar_GetPrevTabsContentsHeight(ImGuiTabBar *self) { return self->PrevTabsContentsHeight; }
void ImGuiTabBar_SetWidthAllTabs(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->WidthAllTabs = v; }
float ImGuiTabBar_GetWidthAllTabs(ImGuiTabBar *self) { return self->WidthAllTabs; }
void ImGuiTabBar_SetWidthAllTabsIdeal(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->WidthAllTabsIdeal = v; }
float ImGuiTabBar_GetWidthAllTabsIdeal(ImGuiTabBar *self) { return self->WidthAllTabsIdeal; }
void ImGuiTabBar_SetScrollingAnim(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ScrollingAnim = v; }
float ImGuiTabBar_GetScrollingAnim(ImGuiTabBar *self) { return self->ScrollingAnim; }
void ImGuiTabBar_SetScrollingTarget(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ScrollingTarget = v; }
float ImGuiTabBar_GetScrollingTarget(ImGuiTabBar *self) { return self->ScrollingTarget; }
void ImGuiTabBar_SetScrollingTargetDistToVisibility(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ScrollingTargetDistToVisibility = v; }
float ImGuiTabBar_GetScrollingTargetDistToVisibility(ImGuiTabBar *self) { return self->ScrollingTargetDistToVisibility; }
void ImGuiTabBar_SetScrollingSpeed(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ScrollingSpeed = v; }
float ImGuiTabBar_GetScrollingSpeed(ImGuiTabBar *self) { return self->ScrollingSpeed; }
void ImGuiTabBar_SetScrollingRectMinX(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ScrollingRectMinX = v; }
float ImGuiTabBar_GetScrollingRectMinX(ImGuiTabBar *self) { return self->ScrollingRectMinX; }
void ImGuiTabBar_SetScrollingRectMaxX(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ScrollingRectMaxX = v; }
float ImGuiTabBar_GetScrollingRectMaxX(ImGuiTabBar *self) { return self->ScrollingRectMaxX; }
void ImGuiTabBar_SetReorderRequestTabId(ImGuiTabBar *ImGuiTabBarPtr, ImGuiID v) { ImGuiTabBarPtr->ReorderRequestTabId = v; }
ImGuiID ImGuiTabBar_GetReorderRequestTabId(ImGuiTabBar *self) { return self->ReorderRequestTabId; }
void ImGuiTabBar_SetReorderRequestOffset(ImGuiTabBar *ImGuiTabBarPtr, ImS16 v) { ImGuiTabBarPtr->ReorderRequestOffset = v; }
ImS16 ImGuiTabBar_GetReorderRequestOffset(ImGuiTabBar *self) { return self->ReorderRequestOffset; }
void ImGuiTabBar_SetBeginCount(ImGuiTabBar *ImGuiTabBarPtr, ImS8 v) { ImGuiTabBarPtr->BeginCount = v; }
ImS8 ImGuiTabBar_GetBeginCount(ImGuiTabBar *self) { return self->BeginCount; }
void ImGuiTabBar_SetWantLayout(ImGuiTabBar *ImGuiTabBarPtr, bool v) { ImGuiTabBarPtr->WantLayout = v; }
bool ImGuiTabBar_GetWantLayout(ImGuiTabBar *self) { return self->WantLayout; }
void ImGuiTabBar_SetVisibleTabWasSubmitted(ImGuiTabBar *ImGuiTabBarPtr, bool v) { ImGuiTabBarPtr->VisibleTabWasSubmitted = v; }
bool ImGuiTabBar_GetVisibleTabWasSubmitted(ImGuiTabBar *self) { return self->VisibleTabWasSubmitted; }
void ImGuiTabBar_SetTabsAddedNew(ImGuiTabBar *ImGuiTabBarPtr, bool v) { ImGuiTabBarPtr->TabsAddedNew = v; }
bool ImGuiTabBar_GetTabsAddedNew(ImGuiTabBar *self) { return self->TabsAddedNew; }
void ImGuiTabBar_SetTabsActiveCount(ImGuiTabBar *ImGuiTabBarPtr, ImS16 v) { ImGuiTabBarPtr->TabsActiveCount = v; }
ImS16 ImGuiTabBar_GetTabsActiveCount(ImGuiTabBar *self) { return self->TabsActiveCount; }
void ImGuiTabBar_SetLastTabItemIdx(ImGuiTabBar *ImGuiTabBarPtr, ImS16 v) { ImGuiTabBarPtr->LastTabItemIdx = v; }
ImS16 ImGuiTabBar_GetLastTabItemIdx(ImGuiTabBar *self) { return self->LastTabItemIdx; }
void ImGuiTabBar_SetItemSpacingY(ImGuiTabBar *ImGuiTabBarPtr, float v) { ImGuiTabBarPtr->ItemSpacingY = v; }
float ImGuiTabBar_GetItemSpacingY(ImGuiTabBar *self) { return self->ItemSpacingY; }
void ImGuiTabBar_SetFramePadding(ImGuiTabBar *ImGuiTabBarPtr, ImVec2 v) { ImGuiTabBarPtr->FramePadding = v; }
ImVec2 ImGuiTabBar_GetFramePadding(ImGuiTabBar *self) { return self->FramePadding; }
void ImGuiTabBar_SetBackupCursorPos(ImGuiTabBar *ImGuiTabBarPtr, ImVec2 v) { ImGuiTabBarPtr->BackupCursorPos = v; }
ImVec2 ImGuiTabBar_GetBackupCursorPos(ImGuiTabBar *self) { return self->BackupCursorPos; }
void ImGuiTabBar_SetTabsNames(ImGuiTabBar *ImGuiTabBarPtr, ImGuiTextBuffer v) { ImGuiTabBarPtr->TabsNames = v; }
ImGuiTextBuffer ImGuiTabBar_GetTabsNames(ImGuiTabBar *self) { return self->TabsNames; }
void ImGuiViewport_SetID(ImGuiViewport *ImGuiViewportPtr, ImGuiID v) { ImGuiViewportPtr->ID = v; }
ImGuiID ImGuiViewport_GetID(ImGuiViewport *self) { return self->ID; }
void ImGuiViewport_SetFlags(ImGuiViewport *ImGuiViewportPtr, ImGuiViewportFlags v) { ImGuiViewportPtr->Flags = v; }
ImGuiViewportFlags ImGuiViewport_GetFlags(ImGuiViewport *self) { return self->Flags; }
void ImGuiViewport_SetPos(ImGuiViewport *ImGuiViewportPtr, ImVec2 v) { ImGuiViewportPtr->Pos = v; }
ImVec2 ImGuiViewport_GetPos(ImGuiViewport *self) { return self->Pos; }
void ImGuiViewport_SetSize(ImGuiViewport *ImGuiViewportPtr, ImVec2 v) { ImGuiViewportPtr->Size = v; }
ImVec2 ImGuiViewport_GetSize(ImGuiViewport *self) { return self->Size; }
void ImGuiViewport_SetWorkPos(ImGuiViewport *ImGuiViewportPtr, ImVec2 v) { ImGuiViewportPtr->WorkPos = v; }
ImVec2 ImGuiViewport_GetWorkPos(ImGuiViewport *self) { return self->WorkPos; }
void ImGuiViewport_SetWorkSize(ImGuiViewport *ImGuiViewportPtr, ImVec2 v) { ImGuiViewportPtr->WorkSize = v; }
ImVec2 ImGuiViewport_GetWorkSize(ImGuiViewport *self) { return self->WorkSize; }
void ImGuiViewport_SetDpiScale(ImGuiViewport *ImGuiViewportPtr, float v) { ImGuiViewportPtr->DpiScale = v; }
float ImGuiViewport_GetDpiScale(ImGuiViewport *self) { return self->DpiScale; }
void ImGuiViewport_SetParentViewportId(ImGuiViewport *ImGuiViewportPtr, ImGuiID v) { ImGuiViewportPtr->ParentViewportId = v; }
ImGuiID ImGuiViewport_GetParentViewportId(ImGuiViewport *self) { return self->ParentViewportId; }
void ImGuiViewport_SetDrawData(ImGuiViewport *ImGuiViewportPtr, ImDrawData* v) { ImGuiViewportPtr->DrawData = v; }
ImDrawData* ImGuiViewport_GetDrawData(ImGuiViewport *self) { return self->DrawData; }
void ImGuiViewport_SetRendererUserData(ImGuiViewport *ImGuiViewportPtr, void* v) { ImGuiViewportPtr->RendererUserData = v; }
void* ImGuiViewport_GetRendererUserData(ImGuiViewport *self) { return self->RendererUserData; }
void ImGuiViewport_SetPlatformUserData(ImGuiViewport *ImGuiViewportPtr, void* v) { ImGuiViewportPtr->PlatformUserData = v; }
void* ImGuiViewport_GetPlatformUserData(ImGuiViewport *self) { return self->PlatformUserData; }
void ImGuiViewport_SetPlatformHandle(ImGuiViewport *ImGuiViewportPtr, void* v) { ImGuiViewportPtr->PlatformHandle = v; }
void* ImGuiViewport_GetPlatformHandle(ImGuiViewport *self) { return self->PlatformHandle; }
void ImGuiViewport_SetPlatformHandleRaw(ImGuiViewport *ImGuiViewportPtr, void* v) { ImGuiViewportPtr->PlatformHandleRaw = v; }
void* ImGuiViewport_GetPlatformHandleRaw(ImGuiViewport *self) { return self->PlatformHandleRaw; }
void ImGuiViewport_SetPlatformRequestMove(ImGuiViewport *ImGuiViewportPtr, bool v) { ImGuiViewportPtr->PlatformRequestMove = v; }
bool ImGuiViewport_GetPlatformRequestMove(ImGuiViewport *self) { return self->PlatformRequestMove; }
void ImGuiViewport_SetPlatformRequestResize(ImGuiViewport *ImGuiViewportPtr, bool v) { ImGuiViewportPtr->PlatformRequestResize = v; }
bool ImGuiViewport_GetPlatformRequestResize(ImGuiViewport *self) { return self->PlatformRequestResize; }
void ImGuiViewport_SetPlatformRequestClose(ImGuiViewport *ImGuiViewportPtr, bool v) { ImGuiViewportPtr->PlatformRequestClose = v; }
bool ImGuiViewport_GetPlatformRequestClose(ImGuiViewport *self) { return self->PlatformRequestClose; }
void ImGuiStackTool_SetLastActiveFrame(ImGuiStackTool *ImGuiStackToolPtr, int v) { ImGuiStackToolPtr->LastActiveFrame = v; }
int ImGuiStackTool_GetLastActiveFrame(ImGuiStackTool *self) { return self->LastActiveFrame; }
void ImGuiStackTool_SetStackLevel(ImGuiStackTool *ImGuiStackToolPtr, int v) { ImGuiStackToolPtr->StackLevel = v; }
int ImGuiStackTool_GetStackLevel(ImGuiStackTool *self) { return self->StackLevel; }
void ImGuiStackTool_SetQueryId(ImGuiStackTool *ImGuiStackToolPtr, ImGuiID v) { ImGuiStackToolPtr->QueryId = v; }
ImGuiID ImGuiStackTool_GetQueryId(ImGuiStackTool *self) { return self->QueryId; }
void ImGuiStackTool_SetResults(ImGuiStackTool *ImGuiStackToolPtr, ImVector_ImGuiStackLevelInfo v) { ImGuiStackToolPtr->Results = v; }
ImVector_ImGuiStackLevelInfo ImGuiStackTool_GetResults(ImGuiStackTool *self) { return self->Results; }
void ImGuiStackTool_SetCopyToClipboardOnCtrlC(ImGuiStackTool *ImGuiStackToolPtr, bool v) { ImGuiStackToolPtr->CopyToClipboardOnCtrlC = v; }
bool ImGuiStackTool_GetCopyToClipboardOnCtrlC(ImGuiStackTool *self) { return self->CopyToClipboardOnCtrlC; }
void ImGuiStackTool_SetCopyToClipboardLastTime(ImGuiStackTool *ImGuiStackToolPtr, float v) { ImGuiStackToolPtr->CopyToClipboardLastTime = v; }
float ImGuiStackTool_GetCopyToClipboardLastTime(ImGuiStackTool *self) { return self->CopyToClipboardLastTime; }
void ImGuiStyle_SetAlpha(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->Alpha = v; }
float ImGuiStyle_GetAlpha(ImGuiStyle *self) { return self->Alpha; }
void ImGuiStyle_SetDisabledAlpha(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->DisabledAlpha = v; }
float ImGuiStyle_GetDisabledAlpha(ImGuiStyle *self) { return self->DisabledAlpha; }
void ImGuiStyle_SetWindowPadding(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->WindowPadding = v; }
ImVec2 ImGuiStyle_GetWindowPadding(ImGuiStyle *self) { return self->WindowPadding; }
void ImGuiStyle_SetWindowRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->WindowRounding = v; }
float ImGuiStyle_GetWindowRounding(ImGuiStyle *self) { return self->WindowRounding; }
void ImGuiStyle_SetWindowBorderSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->WindowBorderSize = v; }
float ImGuiStyle_GetWindowBorderSize(ImGuiStyle *self) { return self->WindowBorderSize; }
void ImGuiStyle_SetWindowMinSize(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->WindowMinSize = v; }
ImVec2 ImGuiStyle_GetWindowMinSize(ImGuiStyle *self) { return self->WindowMinSize; }
void ImGuiStyle_SetWindowTitleAlign(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->WindowTitleAlign = v; }
ImVec2 ImGuiStyle_GetWindowTitleAlign(ImGuiStyle *self) { return self->WindowTitleAlign; }
void ImGuiStyle_SetWindowMenuButtonPosition(ImGuiStyle *ImGuiStylePtr, ImGuiDir v) { ImGuiStylePtr->WindowMenuButtonPosition = v; }
ImGuiDir ImGuiStyle_GetWindowMenuButtonPosition(ImGuiStyle *self) { return self->WindowMenuButtonPosition; }
void ImGuiStyle_SetChildRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->ChildRounding = v; }
float ImGuiStyle_GetChildRounding(ImGuiStyle *self) { return self->ChildRounding; }
void ImGuiStyle_SetChildBorderSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->ChildBorderSize = v; }
float ImGuiStyle_GetChildBorderSize(ImGuiStyle *self) { return self->ChildBorderSize; }
void ImGuiStyle_SetPopupRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->PopupRounding = v; }
float ImGuiStyle_GetPopupRounding(ImGuiStyle *self) { return self->PopupRounding; }
void ImGuiStyle_SetPopupBorderSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->PopupBorderSize = v; }
float ImGuiStyle_GetPopupBorderSize(ImGuiStyle *self) { return self->PopupBorderSize; }
void ImGuiStyle_SetFramePadding(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->FramePadding = v; }
ImVec2 ImGuiStyle_GetFramePadding(ImGuiStyle *self) { return self->FramePadding; }
void ImGuiStyle_SetFrameRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->FrameRounding = v; }
float ImGuiStyle_GetFrameRounding(ImGuiStyle *self) { return self->FrameRounding; }
void ImGuiStyle_SetFrameBorderSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->FrameBorderSize = v; }
float ImGuiStyle_GetFrameBorderSize(ImGuiStyle *self) { return self->FrameBorderSize; }
void ImGuiStyle_SetItemSpacing(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->ItemSpacing = v; }
ImVec2 ImGuiStyle_GetItemSpacing(ImGuiStyle *self) { return self->ItemSpacing; }
void ImGuiStyle_SetItemInnerSpacing(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->ItemInnerSpacing = v; }
ImVec2 ImGuiStyle_GetItemInnerSpacing(ImGuiStyle *self) { return self->ItemInnerSpacing; }
void ImGuiStyle_SetCellPadding(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->CellPadding = v; }
ImVec2 ImGuiStyle_GetCellPadding(ImGuiStyle *self) { return self->CellPadding; }
void ImGuiStyle_SetTouchExtraPadding(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->TouchExtraPadding = v; }
ImVec2 ImGuiStyle_GetTouchExtraPadding(ImGuiStyle *self) { return self->TouchExtraPadding; }
void ImGuiStyle_SetIndentSpacing(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->IndentSpacing = v; }
float ImGuiStyle_GetIndentSpacing(ImGuiStyle *self) { return self->IndentSpacing; }
void ImGuiStyle_SetColumnsMinSpacing(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->ColumnsMinSpacing = v; }
float ImGuiStyle_GetColumnsMinSpacing(ImGuiStyle *self) { return self->ColumnsMinSpacing; }
void ImGuiStyle_SetScrollbarSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->ScrollbarSize = v; }
float ImGuiStyle_GetScrollbarSize(ImGuiStyle *self) { return self->ScrollbarSize; }
void ImGuiStyle_SetScrollbarRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->ScrollbarRounding = v; }
float ImGuiStyle_GetScrollbarRounding(ImGuiStyle *self) { return self->ScrollbarRounding; }
void ImGuiStyle_SetGrabMinSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->GrabMinSize = v; }
float ImGuiStyle_GetGrabMinSize(ImGuiStyle *self) { return self->GrabMinSize; }
void ImGuiStyle_SetGrabRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->GrabRounding = v; }
float ImGuiStyle_GetGrabRounding(ImGuiStyle *self) { return self->GrabRounding; }
void ImGuiStyle_SetLogSliderDeadzone(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->LogSliderDeadzone = v; }
float ImGuiStyle_GetLogSliderDeadzone(ImGuiStyle *self) { return self->LogSliderDeadzone; }
void ImGuiStyle_SetTabRounding(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->TabRounding = v; }
float ImGuiStyle_GetTabRounding(ImGuiStyle *self) { return self->TabRounding; }
void ImGuiStyle_SetTabBorderSize(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->TabBorderSize = v; }
float ImGuiStyle_GetTabBorderSize(ImGuiStyle *self) { return self->TabBorderSize; }
void ImGuiStyle_SetTabMinWidthForCloseButton(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->TabMinWidthForCloseButton = v; }
float ImGuiStyle_GetTabMinWidthForCloseButton(ImGuiStyle *self) { return self->TabMinWidthForCloseButton; }
void ImGuiStyle_SetColorButtonPosition(ImGuiStyle *ImGuiStylePtr, ImGuiDir v) { ImGuiStylePtr->ColorButtonPosition = v; }
ImGuiDir ImGuiStyle_GetColorButtonPosition(ImGuiStyle *self) { return self->ColorButtonPosition; }
void ImGuiStyle_SetButtonTextAlign(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->ButtonTextAlign = v; }
ImVec2 ImGuiStyle_GetButtonTextAlign(ImGuiStyle *self) { return self->ButtonTextAlign; }
void ImGuiStyle_SetSelectableTextAlign(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->SelectableTextAlign = v; }
ImVec2 ImGuiStyle_GetSelectableTextAlign(ImGuiStyle *self) { return self->SelectableTextAlign; }
void ImGuiStyle_SetDisplayWindowPadding(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->DisplayWindowPadding = v; }
ImVec2 ImGuiStyle_GetDisplayWindowPadding(ImGuiStyle *self) { return self->DisplayWindowPadding; }
void ImGuiStyle_SetDisplaySafeAreaPadding(ImGuiStyle *ImGuiStylePtr, ImVec2 v) { ImGuiStylePtr->DisplaySafeAreaPadding = v; }
ImVec2 ImGuiStyle_GetDisplaySafeAreaPadding(ImGuiStyle *self) { return self->DisplaySafeAreaPadding; }
void ImGuiStyle_SetMouseCursorScale(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->MouseCursorScale = v; }
float ImGuiStyle_GetMouseCursorScale(ImGuiStyle *self) { return self->MouseCursorScale; }
void ImGuiStyle_SetAntiAliasedLines(ImGuiStyle *ImGuiStylePtr, bool v) { ImGuiStylePtr->AntiAliasedLines = v; }
bool ImGuiStyle_GetAntiAliasedLines(ImGuiStyle *self) { return self->AntiAliasedLines; }
void ImGuiStyle_SetAntiAliasedLinesUseTex(ImGuiStyle *ImGuiStylePtr, bool v) { ImGuiStylePtr->AntiAliasedLinesUseTex = v; }
bool ImGuiStyle_GetAntiAliasedLinesUseTex(ImGuiStyle *self) { return self->AntiAliasedLinesUseTex; }
void ImGuiStyle_SetAntiAliasedFill(ImGuiStyle *ImGuiStylePtr, bool v) { ImGuiStylePtr->AntiAliasedFill = v; }
bool ImGuiStyle_GetAntiAliasedFill(ImGuiStyle *self) { return self->AntiAliasedFill; }
void ImGuiStyle_SetCurveTessellationTol(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->CurveTessellationTol = v; }
float ImGuiStyle_GetCurveTessellationTol(ImGuiStyle *self) { return self->CurveTessellationTol; }
void ImGuiStyle_SetCircleTessellationMaxError(ImGuiStyle *ImGuiStylePtr, float v) { ImGuiStylePtr->CircleTessellationMaxError = v; }
float ImGuiStyle_GetCircleTessellationMaxError(ImGuiStyle *self) { return self->CircleTessellationMaxError; }
void ImGuiWindowStackData_SetWindow(ImGuiWindowStackData *ImGuiWindowStackDataPtr, ImGuiWindow* v) { ImGuiWindowStackDataPtr->Window = v; }
ImGuiWindow* ImGuiWindowStackData_GetWindow(ImGuiWindowStackData *self) { return self->Window; }
void ImGuiWindowStackData_SetParentLastItemDataBackup(ImGuiWindowStackData *ImGuiWindowStackDataPtr, ImGuiLastItemData v) { ImGuiWindowStackDataPtr->ParentLastItemDataBackup = v; }
ImGuiLastItemData ImGuiWindowStackData_GetParentLastItemDataBackup(ImGuiWindowStackData *self) { return self->ParentLastItemDataBackup; }
void ImGuiWindowStackData_SetStackSizesOnBegin(ImGuiWindowStackData *ImGuiWindowStackDataPtr, ImGuiStackSizes v) { ImGuiWindowStackDataPtr->StackSizesOnBegin = v; }
ImGuiStackSizes ImGuiWindowStackData_GetStackSizesOnBegin(ImGuiWindowStackData *self) { return self->StackSizesOnBegin; }
void StbUndoRecord_Setwhere(StbUndoRecord *StbUndoRecordPtr, int v) { StbUndoRecordPtr->where = v; }
int StbUndoRecord_Getwhere(StbUndoRecord *self) { return self->where; }
void StbUndoRecord_Setinsert_length(StbUndoRecord *StbUndoRecordPtr, int v) { StbUndoRecordPtr->insert_length = v; }
int StbUndoRecord_Getinsert_length(StbUndoRecord *self) { return self->insert_length; }
void StbUndoRecord_Setdelete_length(StbUndoRecord *StbUndoRecordPtr, int v) { StbUndoRecordPtr->delete_length = v; }
int StbUndoRecord_Getdelete_length(StbUndoRecord *self) { return self->delete_length; }
void StbUndoRecord_Setchar_storage(StbUndoRecord *StbUndoRecordPtr, int v) { StbUndoRecordPtr->char_storage = v; }
int StbUndoRecord_Getchar_storage(StbUndoRecord *self) { return self->char_storage; }
void ImGuiStoragePair_Setkey(ImGuiStoragePair *ImGuiStoragePairPtr, ImGuiID v) { ImGuiStoragePairPtr->key = v; }
ImGuiID ImGuiStoragePair_Getkey(ImGuiStoragePair *self) { return self->key; }
void ImGuiIO_SetConfigFlags(ImGuiIO *ImGuiIOPtr, ImGuiConfigFlags v) { ImGuiIOPtr->ConfigFlags = v; }
ImGuiConfigFlags ImGuiIO_GetConfigFlags(ImGuiIO *self) { return self->ConfigFlags; }
void ImGuiIO_SetBackendFlags(ImGuiIO *ImGuiIOPtr, ImGuiBackendFlags v) { ImGuiIOPtr->BackendFlags = v; }
ImGuiBackendFlags ImGuiIO_GetBackendFlags(ImGuiIO *self) { return self->BackendFlags; }
void ImGuiIO_SetDisplaySize(ImGuiIO *ImGuiIOPtr, ImVec2 v) { ImGuiIOPtr->DisplaySize = v; }
ImVec2 ImGuiIO_GetDisplaySize(ImGuiIO *self) { return self->DisplaySize; }
void ImGuiIO_SetDeltaTime(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->DeltaTime = v; }
float ImGuiIO_GetDeltaTime(ImGuiIO *self) { return self->DeltaTime; }
void ImGuiIO_SetIniSavingRate(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->IniSavingRate = v; }
float ImGuiIO_GetIniSavingRate(ImGuiIO *self) { return self->IniSavingRate; }
void ImGuiIO_SetIniFilename(ImGuiIO *ImGuiIOPtr, const char* v) { ImGuiIOPtr->IniFilename = v; }
const char* ImGuiIO_GetIniFilename(ImGuiIO *self) { return self->IniFilename; }
void ImGuiIO_SetLogFilename(ImGuiIO *ImGuiIOPtr, const char* v) { ImGuiIOPtr->LogFilename = v; }
const char* ImGuiIO_GetLogFilename(ImGuiIO *self) { return self->LogFilename; }
void ImGuiIO_SetMouseDoubleClickTime(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->MouseDoubleClickTime = v; }
float ImGuiIO_GetMouseDoubleClickTime(ImGuiIO *self) { return self->MouseDoubleClickTime; }
void ImGuiIO_SetMouseDoubleClickMaxDist(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->MouseDoubleClickMaxDist = v; }
float ImGuiIO_GetMouseDoubleClickMaxDist(ImGuiIO *self) { return self->MouseDoubleClickMaxDist; }
void ImGuiIO_SetMouseDragThreshold(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->MouseDragThreshold = v; }
float ImGuiIO_GetMouseDragThreshold(ImGuiIO *self) { return self->MouseDragThreshold; }
void ImGuiIO_SetKeyRepeatDelay(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->KeyRepeatDelay = v; }
float ImGuiIO_GetKeyRepeatDelay(ImGuiIO *self) { return self->KeyRepeatDelay; }
void ImGuiIO_SetKeyRepeatRate(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->KeyRepeatRate = v; }
float ImGuiIO_GetKeyRepeatRate(ImGuiIO *self) { return self->KeyRepeatRate; }
void ImGuiIO_SetUserData(ImGuiIO *ImGuiIOPtr, void* v) { ImGuiIOPtr->UserData = v; }
void* ImGuiIO_GetUserData(ImGuiIO *self) { return self->UserData; }
void ImGuiIO_SetFonts(ImGuiIO *ImGuiIOPtr, ImFontAtlas* v) { ImGuiIOPtr->Fonts = v; }
ImFontAtlas* ImGuiIO_GetFonts(ImGuiIO *self) { return self->Fonts; }
void ImGuiIO_SetFontGlobalScale(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->FontGlobalScale = v; }
float ImGuiIO_GetFontGlobalScale(ImGuiIO *self) { return self->FontGlobalScale; }
void ImGuiIO_SetFontAllowUserScaling(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->FontAllowUserScaling = v; }
bool ImGuiIO_GetFontAllowUserScaling(ImGuiIO *self) { return self->FontAllowUserScaling; }
void ImGuiIO_SetFontDefault(ImGuiIO *ImGuiIOPtr, ImFont* v) { ImGuiIOPtr->FontDefault = v; }
ImFont* ImGuiIO_GetFontDefault(ImGuiIO *self) { return self->FontDefault; }
void ImGuiIO_SetDisplayFramebufferScale(ImGuiIO *ImGuiIOPtr, ImVec2 v) { ImGuiIOPtr->DisplayFramebufferScale = v; }
ImVec2 ImGuiIO_GetDisplayFramebufferScale(ImGuiIO *self) { return self->DisplayFramebufferScale; }
void ImGuiIO_SetConfigDockingNoSplit(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigDockingNoSplit = v; }
bool ImGuiIO_GetConfigDockingNoSplit(ImGuiIO *self) { return self->ConfigDockingNoSplit; }
void ImGuiIO_SetConfigDockingWithShift(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigDockingWithShift = v; }
bool ImGuiIO_GetConfigDockingWithShift(ImGuiIO *self) { return self->ConfigDockingWithShift; }
void ImGuiIO_SetConfigDockingAlwaysTabBar(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigDockingAlwaysTabBar = v; }
bool ImGuiIO_GetConfigDockingAlwaysTabBar(ImGuiIO *self) { return self->ConfigDockingAlwaysTabBar; }
void ImGuiIO_SetConfigDockingTransparentPayload(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigDockingTransparentPayload = v; }
bool ImGuiIO_GetConfigDockingTransparentPayload(ImGuiIO *self) { return self->ConfigDockingTransparentPayload; }
void ImGuiIO_SetConfigViewportsNoAutoMerge(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigViewportsNoAutoMerge = v; }
bool ImGuiIO_GetConfigViewportsNoAutoMerge(ImGuiIO *self) { return self->ConfigViewportsNoAutoMerge; }
void ImGuiIO_SetConfigViewportsNoTaskBarIcon(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigViewportsNoTaskBarIcon = v; }
bool ImGuiIO_GetConfigViewportsNoTaskBarIcon(ImGuiIO *self) { return self->ConfigViewportsNoTaskBarIcon; }
void ImGuiIO_SetConfigViewportsNoDecoration(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigViewportsNoDecoration = v; }
bool ImGuiIO_GetConfigViewportsNoDecoration(ImGuiIO *self) { return self->ConfigViewportsNoDecoration; }
void ImGuiIO_SetConfigViewportsNoDefaultParent(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigViewportsNoDefaultParent = v; }
bool ImGuiIO_GetConfigViewportsNoDefaultParent(ImGuiIO *self) { return self->ConfigViewportsNoDefaultParent; }
void ImGuiIO_SetMouseDrawCursor(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->MouseDrawCursor = v; }
bool ImGuiIO_GetMouseDrawCursor(ImGuiIO *self) { return self->MouseDrawCursor; }
void ImGuiIO_SetConfigMacOSXBehaviors(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigMacOSXBehaviors = v; }
bool ImGuiIO_GetConfigMacOSXBehaviors(ImGuiIO *self) { return self->ConfigMacOSXBehaviors; }
void ImGuiIO_SetConfigInputTrickleEventQueue(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigInputTrickleEventQueue = v; }
bool ImGuiIO_GetConfigInputTrickleEventQueue(ImGuiIO *self) { return self->ConfigInputTrickleEventQueue; }
void ImGuiIO_SetConfigInputTextCursorBlink(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigInputTextCursorBlink = v; }
bool ImGuiIO_GetConfigInputTextCursorBlink(ImGuiIO *self) { return self->ConfigInputTextCursorBlink; }
void ImGuiIO_SetConfigInputTextEnterKeepActive(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigInputTextEnterKeepActive = v; }
bool ImGuiIO_GetConfigInputTextEnterKeepActive(ImGuiIO *self) { return self->ConfigInputTextEnterKeepActive; }
void ImGuiIO_SetConfigDragClickToInputText(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigDragClickToInputText = v; }
bool ImGuiIO_GetConfigDragClickToInputText(ImGuiIO *self) { return self->ConfigDragClickToInputText; }
void ImGuiIO_SetConfigWindowsResizeFromEdges(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigWindowsResizeFromEdges = v; }
bool ImGuiIO_GetConfigWindowsResizeFromEdges(ImGuiIO *self) { return self->ConfigWindowsResizeFromEdges; }
void ImGuiIO_SetConfigWindowsMoveFromTitleBarOnly(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->ConfigWindowsMoveFromTitleBarOnly = v; }
bool ImGuiIO_GetConfigWindowsMoveFromTitleBarOnly(ImGuiIO *self) { return self->ConfigWindowsMoveFromTitleBarOnly; }
void ImGuiIO_SetConfigMemoryCompactTimer(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->ConfigMemoryCompactTimer = v; }
float ImGuiIO_GetConfigMemoryCompactTimer(ImGuiIO *self) { return self->ConfigMemoryCompactTimer; }
void ImGuiIO_SetBackendPlatformName(ImGuiIO *ImGuiIOPtr, const char* v) { ImGuiIOPtr->BackendPlatformName = v; }
const char* ImGuiIO_GetBackendPlatformName(ImGuiIO *self) { return self->BackendPlatformName; }
void ImGuiIO_SetBackendRendererName(ImGuiIO *ImGuiIOPtr, const char* v) { ImGuiIOPtr->BackendRendererName = v; }
const char* ImGuiIO_GetBackendRendererName(ImGuiIO *self) { return self->BackendRendererName; }
void ImGuiIO_SetBackendPlatformUserData(ImGuiIO *ImGuiIOPtr, void* v) { ImGuiIOPtr->BackendPlatformUserData = v; }
void* ImGuiIO_GetBackendPlatformUserData(ImGuiIO *self) { return self->BackendPlatformUserData; }
void ImGuiIO_SetBackendRendererUserData(ImGuiIO *ImGuiIOPtr, void* v) { ImGuiIOPtr->BackendRendererUserData = v; }
void* ImGuiIO_GetBackendRendererUserData(ImGuiIO *self) { return self->BackendRendererUserData; }
void ImGuiIO_SetBackendLanguageUserData(ImGuiIO *ImGuiIOPtr, void* v) { ImGuiIOPtr->BackendLanguageUserData = v; }
void* ImGuiIO_GetBackendLanguageUserData(ImGuiIO *self) { return self->BackendLanguageUserData; }
void ImGuiIO_SetClipboardUserData(ImGuiIO *ImGuiIOPtr, void* v) { ImGuiIOPtr->ClipboardUserData = v; }
void* ImGuiIO_GetClipboardUserData(ImGuiIO *self) { return self->ClipboardUserData; }
void ImGuiIO_Set_UnusedPadding(ImGuiIO *ImGuiIOPtr, void* v) { ImGuiIOPtr->_UnusedPadding = v; }
void* ImGuiIO_Get_UnusedPadding(ImGuiIO *self) { return self->_UnusedPadding; }
void ImGuiIO_SetWantCaptureMouse(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->WantCaptureMouse = v; }
bool ImGuiIO_GetWantCaptureMouse(ImGuiIO *self) { return self->WantCaptureMouse; }
void ImGuiIO_SetWantCaptureKeyboard(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->WantCaptureKeyboard = v; }
bool ImGuiIO_GetWantCaptureKeyboard(ImGuiIO *self) { return self->WantCaptureKeyboard; }
void ImGuiIO_SetWantTextInput(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->WantTextInput = v; }
bool ImGuiIO_GetWantTextInput(ImGuiIO *self) { return self->WantTextInput; }
void ImGuiIO_SetWantSetMousePos(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->WantSetMousePos = v; }
bool ImGuiIO_GetWantSetMousePos(ImGuiIO *self) { return self->WantSetMousePos; }
void ImGuiIO_SetWantSaveIniSettings(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->WantSaveIniSettings = v; }
bool ImGuiIO_GetWantSaveIniSettings(ImGuiIO *self) { return self->WantSaveIniSettings; }
void ImGuiIO_SetNavActive(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->NavActive = v; }
bool ImGuiIO_GetNavActive(ImGuiIO *self) { return self->NavActive; }
void ImGuiIO_SetNavVisible(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->NavVisible = v; }
bool ImGuiIO_GetNavVisible(ImGuiIO *self) { return self->NavVisible; }
void ImGuiIO_SetFramerate(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->Framerate = v; }
float ImGuiIO_GetFramerate(ImGuiIO *self) { return self->Framerate; }
void ImGuiIO_SetMetricsRenderVertices(ImGuiIO *ImGuiIOPtr, int v) { ImGuiIOPtr->MetricsRenderVertices = v; }
int ImGuiIO_GetMetricsRenderVertices(ImGuiIO *self) { return self->MetricsRenderVertices; }
void ImGuiIO_SetMetricsRenderIndices(ImGuiIO *ImGuiIOPtr, int v) { ImGuiIOPtr->MetricsRenderIndices = v; }
int ImGuiIO_GetMetricsRenderIndices(ImGuiIO *self) { return self->MetricsRenderIndices; }
void ImGuiIO_SetMetricsRenderWindows(ImGuiIO *ImGuiIOPtr, int v) { ImGuiIOPtr->MetricsRenderWindows = v; }
int ImGuiIO_GetMetricsRenderWindows(ImGuiIO *self) { return self->MetricsRenderWindows; }
void ImGuiIO_SetMetricsActiveWindows(ImGuiIO *ImGuiIOPtr, int v) { ImGuiIOPtr->MetricsActiveWindows = v; }
int ImGuiIO_GetMetricsActiveWindows(ImGuiIO *self) { return self->MetricsActiveWindows; }
void ImGuiIO_SetMetricsActiveAllocations(ImGuiIO *ImGuiIOPtr, int v) { ImGuiIOPtr->MetricsActiveAllocations = v; }
int ImGuiIO_GetMetricsActiveAllocations(ImGuiIO *self) { return self->MetricsActiveAllocations; }
void ImGuiIO_SetMouseDelta(ImGuiIO *ImGuiIOPtr, ImVec2 v) { ImGuiIOPtr->MouseDelta = v; }
ImVec2 ImGuiIO_GetMouseDelta(ImGuiIO *self) { return self->MouseDelta; }
void ImGuiIO_SetMousePos(ImGuiIO *ImGuiIOPtr, ImVec2 v) { ImGuiIOPtr->MousePos = v; }
ImVec2 ImGuiIO_GetMousePos(ImGuiIO *self) { return self->MousePos; }
void ImGuiIO_SetMouseWheel(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->MouseWheel = v; }
float ImGuiIO_GetMouseWheel(ImGuiIO *self) { return self->MouseWheel; }
void ImGuiIO_SetMouseWheelH(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->MouseWheelH = v; }
float ImGuiIO_GetMouseWheelH(ImGuiIO *self) { return self->MouseWheelH; }
void ImGuiIO_SetMouseHoveredViewport(ImGuiIO *ImGuiIOPtr, ImGuiID v) { ImGuiIOPtr->MouseHoveredViewport = v; }
ImGuiID ImGuiIO_GetMouseHoveredViewport(ImGuiIO *self) { return self->MouseHoveredViewport; }
void ImGuiIO_SetKeyCtrl(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->KeyCtrl = v; }
bool ImGuiIO_GetKeyCtrl(ImGuiIO *self) { return self->KeyCtrl; }
void ImGuiIO_SetKeyShift(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->KeyShift = v; }
bool ImGuiIO_GetKeyShift(ImGuiIO *self) { return self->KeyShift; }
void ImGuiIO_SetKeyAlt(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->KeyAlt = v; }
bool ImGuiIO_GetKeyAlt(ImGuiIO *self) { return self->KeyAlt; }
void ImGuiIO_SetKeySuper(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->KeySuper = v; }
bool ImGuiIO_GetKeySuper(ImGuiIO *self) { return self->KeySuper; }
void ImGuiIO_SetKeyMods(ImGuiIO *ImGuiIOPtr, ImGuiModFlags v) { ImGuiIOPtr->KeyMods = v; }
ImGuiModFlags ImGuiIO_GetKeyMods(ImGuiIO *self) { return self->KeyMods; }
void ImGuiIO_SetWantCaptureMouseUnlessPopupClose(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->WantCaptureMouseUnlessPopupClose = v; }
bool ImGuiIO_GetWantCaptureMouseUnlessPopupClose(ImGuiIO *self) { return self->WantCaptureMouseUnlessPopupClose; }
void ImGuiIO_SetMousePosPrev(ImGuiIO *ImGuiIOPtr, ImVec2 v) { ImGuiIOPtr->MousePosPrev = v; }
ImVec2 ImGuiIO_GetMousePosPrev(ImGuiIO *self) { return self->MousePosPrev; }
void ImGuiIO_SetPenPressure(ImGuiIO *ImGuiIOPtr, float v) { ImGuiIOPtr->PenPressure = v; }
float ImGuiIO_GetPenPressure(ImGuiIO *self) { return self->PenPressure; }
void ImGuiIO_SetAppFocusLost(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->AppFocusLost = v; }
bool ImGuiIO_GetAppFocusLost(ImGuiIO *self) { return self->AppFocusLost; }
void ImGuiIO_SetBackendUsingLegacyKeyArrays(ImGuiIO *ImGuiIOPtr, ImS8 v) { ImGuiIOPtr->BackendUsingLegacyKeyArrays = v; }
ImS8 ImGuiIO_GetBackendUsingLegacyKeyArrays(ImGuiIO *self) { return self->BackendUsingLegacyKeyArrays; }
void ImGuiIO_SetBackendUsingLegacyNavInputArray(ImGuiIO *ImGuiIOPtr, bool v) { ImGuiIOPtr->BackendUsingLegacyNavInputArray = v; }
bool ImGuiIO_GetBackendUsingLegacyNavInputArray(ImGuiIO *self) { return self->BackendUsingLegacyNavInputArray; }
void ImGuiIO_SetInputQueueSurrogate(ImGuiIO *ImGuiIOPtr, ImWchar16 v) { ImGuiIOPtr->InputQueueSurrogate = v; }
ImWchar16 ImGuiIO_GetInputQueueSurrogate(ImGuiIO *self) { return self->InputQueueSurrogate; }
void ImGuiIO_SetInputQueueCharacters(ImGuiIO *ImGuiIOPtr, ImVector_ImWchar v) { ImGuiIOPtr->InputQueueCharacters = v; }
ImVector_ImWchar ImGuiIO_GetInputQueueCharacters(ImGuiIO *self) { return self->InputQueueCharacters; }
void ImGuiTableCellData_SetBgColor(ImGuiTableCellData *ImGuiTableCellDataPtr, ImU32 v) { ImGuiTableCellDataPtr->BgColor = v; }
ImU32 ImGuiTableCellData_GetBgColor(ImGuiTableCellData *self) { return self->BgColor; }
void ImGuiTableCellData_SetColumn(ImGuiTableCellData *ImGuiTableCellDataPtr, ImGuiTableColumnIdx v) { ImGuiTableCellDataPtr->Column = v; }
ImGuiTableColumnIdx ImGuiTableCellData_GetColumn(ImGuiTableCellData *self) { return self->Column; }
void ImFontGlyph_SetColored(ImFontGlyph *ImFontGlyphPtr, unsigned int v) { ImFontGlyphPtr->Colored = v; }
unsigned int ImFontGlyph_GetColored(ImFontGlyph *self) { return self->Colored; }
void ImFontGlyph_SetVisible(ImFontGlyph *ImFontGlyphPtr, unsigned int v) { ImFontGlyphPtr->Visible = v; }
unsigned int ImFontGlyph_GetVisible(ImFontGlyph *self) { return self->Visible; }
void ImFontGlyph_SetCodepoint(ImFontGlyph *ImFontGlyphPtr, unsigned int v) { ImFontGlyphPtr->Codepoint = v; }
unsigned int ImFontGlyph_GetCodepoint(ImFontGlyph *self) { return self->Codepoint; }
void ImFontGlyph_SetAdvanceX(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->AdvanceX = v; }
float ImFontGlyph_GetAdvanceX(ImFontGlyph *self) { return self->AdvanceX; }
void ImFontGlyph_SetX0(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->X0 = v; }
float ImFontGlyph_GetX0(ImFontGlyph *self) { return self->X0; }
void ImFontGlyph_SetY0(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->Y0 = v; }
float ImFontGlyph_GetY0(ImFontGlyph *self) { return self->Y0; }
void ImFontGlyph_SetX1(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->X1 = v; }
float ImFontGlyph_GetX1(ImFontGlyph *self) { return self->X1; }
void ImFontGlyph_SetY1(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->Y1 = v; }
float ImFontGlyph_GetY1(ImFontGlyph *self) { return self->Y1; }
void ImFontGlyph_SetU0(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->U0 = v; }
float ImFontGlyph_GetU0(ImFontGlyph *self) { return self->U0; }
void ImFontGlyph_SetV0(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->V0 = v; }
float ImFontGlyph_GetV0(ImFontGlyph *self) { return self->V0; }
void ImFontGlyph_SetU1(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->U1 = v; }
float ImFontGlyph_GetU1(ImFontGlyph *self) { return self->U1; }
void ImFontGlyph_SetV1(ImFontGlyph *ImFontGlyphPtr, float v) { ImFontGlyphPtr->V1 = v; }
float ImFontGlyph_GetV1(ImFontGlyph *self) { return self->V1; }
void ImGuiOnceUponAFrame_SetRefFrame(ImGuiOnceUponAFrame *ImGuiOnceUponAFramePtr, int v) { ImGuiOnceUponAFramePtr->RefFrame = v; }
int ImGuiOnceUponAFrame_GetRefFrame(ImGuiOnceUponAFrame *self) { return self->RefFrame; }
void ImGuiPlatformIO_SetMonitors(ImGuiPlatformIO *ImGuiPlatformIOPtr, ImVector_ImGuiPlatformMonitor v) { ImGuiPlatformIOPtr->Monitors = v; }
ImVector_ImGuiPlatformMonitor ImGuiPlatformIO_GetMonitors(ImGuiPlatformIO *self) { return self->Monitors; }
void ImGuiPlatformIO_SetViewports(ImGuiPlatformIO *ImGuiPlatformIOPtr, ImVector_ImGuiViewportPtr v) { ImGuiPlatformIOPtr->Viewports = v; }
ImVector_ImGuiViewportPtr ImGuiPlatformIO_GetViewports(ImGuiPlatformIO *self) { return self->Viewports; }
void ImGuiPopupData_SetPopupId(ImGuiPopupData *ImGuiPopupDataPtr, ImGuiID v) { ImGuiPopupDataPtr->PopupId = v; }
ImGuiID ImGuiPopupData_GetPopupId(ImGuiPopupData *self) { return self->PopupId; }
void ImGuiPopupData_SetWindow(ImGuiPopupData *ImGuiPopupDataPtr, ImGuiWindow* v) { ImGuiPopupDataPtr->Window = v; }
ImGuiWindow* ImGuiPopupData_GetWindow(ImGuiPopupData *self) { return self->Window; }
void ImGuiPopupData_SetSourceWindow(ImGuiPopupData *ImGuiPopupDataPtr, ImGuiWindow* v) { ImGuiPopupDataPtr->SourceWindow = v; }
ImGuiWindow* ImGuiPopupData_GetSourceWindow(ImGuiPopupData *self) { return self->SourceWindow; }
void ImGuiPopupData_SetParentNavLayer(ImGuiPopupData *ImGuiPopupDataPtr, int v) { ImGuiPopupDataPtr->ParentNavLayer = v; }
int ImGuiPopupData_GetParentNavLayer(ImGuiPopupData *self) { return self->ParentNavLayer; }
void ImGuiPopupData_SetOpenFrameCount(ImGuiPopupData *ImGuiPopupDataPtr, int v) { ImGuiPopupDataPtr->OpenFrameCount = v; }
int ImGuiPopupData_GetOpenFrameCount(ImGuiPopupData *self) { return self->OpenFrameCount; }
void ImGuiPopupData_SetOpenParentId(ImGuiPopupData *ImGuiPopupDataPtr, ImGuiID v) { ImGuiPopupDataPtr->OpenParentId = v; }
ImGuiID ImGuiPopupData_GetOpenParentId(ImGuiPopupData *self) { return self->OpenParentId; }
void ImGuiPopupData_SetOpenPopupPos(ImGuiPopupData *ImGuiPopupDataPtr, ImVec2 v) { ImGuiPopupDataPtr->OpenPopupPos = v; }
ImVec2 ImGuiPopupData_GetOpenPopupPos(ImGuiPopupData *self) { return self->OpenPopupPos; }
void ImGuiPopupData_SetOpenMousePos(ImGuiPopupData *ImGuiPopupDataPtr, ImVec2 v) { ImGuiPopupDataPtr->OpenMousePos = v; }
ImVec2 ImGuiPopupData_GetOpenMousePos(ImGuiPopupData *self) { return self->OpenMousePos; }
void ImDrawListSharedData_SetTexUvWhitePixel(ImDrawListSharedData *ImDrawListSharedDataPtr, ImVec2 v) { ImDrawListSharedDataPtr->TexUvWhitePixel = v; }
ImVec2 ImDrawListSharedData_GetTexUvWhitePixel(ImDrawListSharedData *self) { return self->TexUvWhitePixel; }
void ImDrawListSharedData_SetFont(ImDrawListSharedData *ImDrawListSharedDataPtr, ImFont* v) { ImDrawListSharedDataPtr->Font = v; }
ImFont* ImDrawListSharedData_GetFont(ImDrawListSharedData *self) { return self->Font; }
void ImDrawListSharedData_SetFontSize(ImDrawListSharedData *ImDrawListSharedDataPtr, float v) { ImDrawListSharedDataPtr->FontSize = v; }
float ImDrawListSharedData_GetFontSize(ImDrawListSharedData *self) { return self->FontSize; }
void ImDrawListSharedData_SetCurveTessellationTol(ImDrawListSharedData *ImDrawListSharedDataPtr, float v) { ImDrawListSharedDataPtr->CurveTessellationTol = v; }
float ImDrawListSharedData_GetCurveTessellationTol(ImDrawListSharedData *self) { return self->CurveTessellationTol; }
void ImDrawListSharedData_SetCircleSegmentMaxError(ImDrawListSharedData *ImDrawListSharedDataPtr, float v) { ImDrawListSharedDataPtr->CircleSegmentMaxError = v; }
float ImDrawListSharedData_GetCircleSegmentMaxError(ImDrawListSharedData *self) { return self->CircleSegmentMaxError; }
void ImDrawListSharedData_SetClipRectFullscreen(ImDrawListSharedData *ImDrawListSharedDataPtr, ImVec4 v) { ImDrawListSharedDataPtr->ClipRectFullscreen = v; }
ImVec4 ImDrawListSharedData_GetClipRectFullscreen(ImDrawListSharedData *self) { return self->ClipRectFullscreen; }
void ImDrawListSharedData_SetInitialFlags(ImDrawListSharedData *ImDrawListSharedDataPtr, ImDrawListFlags v) { ImDrawListSharedDataPtr->InitialFlags = v; }
ImDrawListFlags ImDrawListSharedData_GetInitialFlags(ImDrawListSharedData *self) { return self->InitialFlags; }
void ImDrawListSharedData_SetArcFastRadiusCutoff(ImDrawListSharedData *ImDrawListSharedDataPtr, float v) { ImDrawListSharedDataPtr->ArcFastRadiusCutoff = v; }
float ImDrawListSharedData_GetArcFastRadiusCutoff(ImDrawListSharedData *self) { return self->ArcFastRadiusCutoff; }
void ImDrawListSharedData_SetTexUvLines(ImDrawListSharedData *ImDrawListSharedDataPtr, const ImVec4* v) { ImDrawListSharedDataPtr->TexUvLines = v; }
const ImVec4* ImDrawListSharedData_GetTexUvLines(ImDrawListSharedData *self) { return self->TexUvLines; }
void ImDrawCmdHeader_SetClipRect(ImDrawCmdHeader *ImDrawCmdHeaderPtr, ImVec4 v) { ImDrawCmdHeaderPtr->ClipRect = v; }
ImVec4 ImDrawCmdHeader_GetClipRect(ImDrawCmdHeader *self) { return self->ClipRect; }
void ImDrawCmdHeader_SetTextureId(ImDrawCmdHeader *ImDrawCmdHeaderPtr, ImTextureID v) { ImDrawCmdHeaderPtr->TextureId = v; }
ImTextureID ImDrawCmdHeader_GetTextureId(ImDrawCmdHeader *self) { return self->TextureId; }
void ImDrawCmdHeader_SetVtxOffset(ImDrawCmdHeader *ImDrawCmdHeaderPtr, unsigned int v) { ImDrawCmdHeaderPtr->VtxOffset = v; }
unsigned int ImDrawCmdHeader_GetVtxOffset(ImDrawCmdHeader *self) { return self->VtxOffset; }
void ImDrawListSplitter_Set_Current(ImDrawListSplitter *ImDrawListSplitterPtr, int v) { ImDrawListSplitterPtr->_Current = v; }
int ImDrawListSplitter_Get_Current(ImDrawListSplitter *self) { return self->_Current; }
void ImDrawListSplitter_Set_Count(ImDrawListSplitter *ImDrawListSplitterPtr, int v) { ImDrawListSplitterPtr->_Count = v; }
int ImDrawListSplitter_Get_Count(ImDrawListSplitter *self) { return self->_Count; }
void ImDrawListSplitter_Set_Channels(ImDrawListSplitter *ImDrawListSplitterPtr, ImVector_ImDrawChannel v) { ImDrawListSplitterPtr->_Channels = v; }
ImVector_ImDrawChannel ImDrawListSplitter_Get_Channels(ImDrawListSplitter *self) { return self->_Channels; }
void ImGuiPlatformImeData_SetWantVisible(ImGuiPlatformImeData *ImGuiPlatformImeDataPtr, bool v) { ImGuiPlatformImeDataPtr->WantVisible = v; }
bool ImGuiPlatformImeData_GetWantVisible(ImGuiPlatformImeData *self) { return self->WantVisible; }
void ImGuiPlatformImeData_SetInputPos(ImGuiPlatformImeData *ImGuiPlatformImeDataPtr, ImVec2 v) { ImGuiPlatformImeDataPtr->InputPos = v; }
ImVec2 ImGuiPlatformImeData_GetInputPos(ImGuiPlatformImeData *self) { return self->InputPos; }
void ImGuiPlatformImeData_SetInputLineHeight(ImGuiPlatformImeData *ImGuiPlatformImeDataPtr, float v) { ImGuiPlatformImeDataPtr->InputLineHeight = v; }
float ImGuiPlatformImeData_GetInputLineHeight(ImGuiPlatformImeData *self) { return self->InputLineHeight; }
void ImGuiSizeCallbackData_SetUserData(ImGuiSizeCallbackData *ImGuiSizeCallbackDataPtr, void* v) { ImGuiSizeCallbackDataPtr->UserData = v; }
void* ImGuiSizeCallbackData_GetUserData(ImGuiSizeCallbackData *self) { return self->UserData; }
void ImGuiSizeCallbackData_SetPos(ImGuiSizeCallbackData *ImGuiSizeCallbackDataPtr, ImVec2 v) { ImGuiSizeCallbackDataPtr->Pos = v; }
ImVec2 ImGuiSizeCallbackData_GetPos(ImGuiSizeCallbackData *self) { return self->Pos; }
void ImGuiSizeCallbackData_SetCurrentSize(ImGuiSizeCallbackData *ImGuiSizeCallbackDataPtr, ImVec2 v) { ImGuiSizeCallbackDataPtr->CurrentSize = v; }
ImVec2 ImGuiSizeCallbackData_GetCurrentSize(ImGuiSizeCallbackData *self) { return self->CurrentSize; }
void ImGuiSizeCallbackData_SetDesiredSize(ImGuiSizeCallbackData *ImGuiSizeCallbackDataPtr, ImVec2 v) { ImGuiSizeCallbackDataPtr->DesiredSize = v; }
ImVec2 ImGuiSizeCallbackData_GetDesiredSize(ImGuiSizeCallbackData *self) { return self->DesiredSize; }
void ImGuiTableColumn_SetFlags(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableColumnFlags v) { ImGuiTableColumnPtr->Flags = v; }
ImGuiTableColumnFlags ImGuiTableColumn_GetFlags(ImGuiTableColumn *self) { return self->Flags; }
void ImGuiTableColumn_SetWidthGiven(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->WidthGiven = v; }
float ImGuiTableColumn_GetWidthGiven(ImGuiTableColumn *self) { return self->WidthGiven; }
void ImGuiTableColumn_SetMinX(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->MinX = v; }
float ImGuiTableColumn_GetMinX(ImGuiTableColumn *self) { return self->MinX; }
void ImGuiTableColumn_SetMaxX(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->MaxX = v; }
float ImGuiTableColumn_GetMaxX(ImGuiTableColumn *self) { return self->MaxX; }
void ImGuiTableColumn_SetWidthRequest(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->WidthRequest = v; }
float ImGuiTableColumn_GetWidthRequest(ImGuiTableColumn *self) { return self->WidthRequest; }
void ImGuiTableColumn_SetWidthAuto(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->WidthAuto = v; }
float ImGuiTableColumn_GetWidthAuto(ImGuiTableColumn *self) { return self->WidthAuto; }
void ImGuiTableColumn_SetStretchWeight(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->StretchWeight = v; }
float ImGuiTableColumn_GetStretchWeight(ImGuiTableColumn *self) { return self->StretchWeight; }
void ImGuiTableColumn_SetInitStretchWeightOrWidth(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->InitStretchWeightOrWidth = v; }
float ImGuiTableColumn_GetInitStretchWeightOrWidth(ImGuiTableColumn *self) { return self->InitStretchWeightOrWidth; }
void ImGuiTableColumn_SetClipRect(ImGuiTableColumn *ImGuiTableColumnPtr, ImRect v) { ImGuiTableColumnPtr->ClipRect = v; }
ImRect ImGuiTableColumn_GetClipRect(ImGuiTableColumn *self) { return self->ClipRect; }
void ImGuiTableColumn_SetUserID(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiID v) { ImGuiTableColumnPtr->UserID = v; }
ImGuiID ImGuiTableColumn_GetUserID(ImGuiTableColumn *self) { return self->UserID; }
void ImGuiTableColumn_SetWorkMinX(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->WorkMinX = v; }
float ImGuiTableColumn_GetWorkMinX(ImGuiTableColumn *self) { return self->WorkMinX; }
void ImGuiTableColumn_SetWorkMaxX(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->WorkMaxX = v; }
float ImGuiTableColumn_GetWorkMaxX(ImGuiTableColumn *self) { return self->WorkMaxX; }
void ImGuiTableColumn_SetItemWidth(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->ItemWidth = v; }
float ImGuiTableColumn_GetItemWidth(ImGuiTableColumn *self) { return self->ItemWidth; }
void ImGuiTableColumn_SetContentMaxXFrozen(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->ContentMaxXFrozen = v; }
float ImGuiTableColumn_GetContentMaxXFrozen(ImGuiTableColumn *self) { return self->ContentMaxXFrozen; }
void ImGuiTableColumn_SetContentMaxXUnfrozen(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->ContentMaxXUnfrozen = v; }
float ImGuiTableColumn_GetContentMaxXUnfrozen(ImGuiTableColumn *self) { return self->ContentMaxXUnfrozen; }
void ImGuiTableColumn_SetContentMaxXHeadersUsed(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->ContentMaxXHeadersUsed = v; }
float ImGuiTableColumn_GetContentMaxXHeadersUsed(ImGuiTableColumn *self) { return self->ContentMaxXHeadersUsed; }
void ImGuiTableColumn_SetContentMaxXHeadersIdeal(ImGuiTableColumn *ImGuiTableColumnPtr, float v) { ImGuiTableColumnPtr->ContentMaxXHeadersIdeal = v; }
float ImGuiTableColumn_GetContentMaxXHeadersIdeal(ImGuiTableColumn *self) { return self->ContentMaxXHeadersIdeal; }
void ImGuiTableColumn_SetNameOffset(ImGuiTableColumn *ImGuiTableColumnPtr, ImS16 v) { ImGuiTableColumnPtr->NameOffset = v; }
ImS16 ImGuiTableColumn_GetNameOffset(ImGuiTableColumn *self) { return self->NameOffset; }
void ImGuiTableColumn_SetDisplayOrder(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnPtr->DisplayOrder = v; }
ImGuiTableColumnIdx ImGuiTableColumn_GetDisplayOrder(ImGuiTableColumn *self) { return self->DisplayOrder; }
void ImGuiTableColumn_SetIndexWithinEnabledSet(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnPtr->IndexWithinEnabledSet = v; }
ImGuiTableColumnIdx ImGuiTableColumn_GetIndexWithinEnabledSet(ImGuiTableColumn *self) { return self->IndexWithinEnabledSet; }
void ImGuiTableColumn_SetPrevEnabledColumn(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnPtr->PrevEnabledColumn = v; }
ImGuiTableColumnIdx ImGuiTableColumn_GetPrevEnabledColumn(ImGuiTableColumn *self) { return self->PrevEnabledColumn; }
void ImGuiTableColumn_SetNextEnabledColumn(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnPtr->NextEnabledColumn = v; }
ImGuiTableColumnIdx ImGuiTableColumn_GetNextEnabledColumn(ImGuiTableColumn *self) { return self->NextEnabledColumn; }
void ImGuiTableColumn_SetSortOrder(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnPtr->SortOrder = v; }
ImGuiTableColumnIdx ImGuiTableColumn_GetSortOrder(ImGuiTableColumn *self) { return self->SortOrder; }
void ImGuiTableColumn_SetDrawChannelCurrent(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableDrawChannelIdx v) { ImGuiTableColumnPtr->DrawChannelCurrent = v; }
ImGuiTableDrawChannelIdx ImGuiTableColumn_GetDrawChannelCurrent(ImGuiTableColumn *self) { return self->DrawChannelCurrent; }
void ImGuiTableColumn_SetDrawChannelFrozen(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableDrawChannelIdx v) { ImGuiTableColumnPtr->DrawChannelFrozen = v; }
ImGuiTableDrawChannelIdx ImGuiTableColumn_GetDrawChannelFrozen(ImGuiTableColumn *self) { return self->DrawChannelFrozen; }
void ImGuiTableColumn_SetDrawChannelUnfrozen(ImGuiTableColumn *ImGuiTableColumnPtr, ImGuiTableDrawChannelIdx v) { ImGuiTableColumnPtr->DrawChannelUnfrozen = v; }
ImGuiTableDrawChannelIdx ImGuiTableColumn_GetDrawChannelUnfrozen(ImGuiTableColumn *self) { return self->DrawChannelUnfrozen; }
void ImGuiTableColumn_SetIsEnabled(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsEnabled = v; }
bool ImGuiTableColumn_GetIsEnabled(ImGuiTableColumn *self) { return self->IsEnabled; }
void ImGuiTableColumn_SetIsUserEnabled(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsUserEnabled = v; }
bool ImGuiTableColumn_GetIsUserEnabled(ImGuiTableColumn *self) { return self->IsUserEnabled; }
void ImGuiTableColumn_SetIsUserEnabledNextFrame(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsUserEnabledNextFrame = v; }
bool ImGuiTableColumn_GetIsUserEnabledNextFrame(ImGuiTableColumn *self) { return self->IsUserEnabledNextFrame; }
void ImGuiTableColumn_SetIsVisibleX(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsVisibleX = v; }
bool ImGuiTableColumn_GetIsVisibleX(ImGuiTableColumn *self) { return self->IsVisibleX; }
void ImGuiTableColumn_SetIsVisibleY(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsVisibleY = v; }
bool ImGuiTableColumn_GetIsVisibleY(ImGuiTableColumn *self) { return self->IsVisibleY; }
void ImGuiTableColumn_SetIsRequestOutput(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsRequestOutput = v; }
bool ImGuiTableColumn_GetIsRequestOutput(ImGuiTableColumn *self) { return self->IsRequestOutput; }
void ImGuiTableColumn_SetIsSkipItems(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsSkipItems = v; }
bool ImGuiTableColumn_GetIsSkipItems(ImGuiTableColumn *self) { return self->IsSkipItems; }
void ImGuiTableColumn_SetIsPreserveWidthAuto(ImGuiTableColumn *ImGuiTableColumnPtr, bool v) { ImGuiTableColumnPtr->IsPreserveWidthAuto = v; }
bool ImGuiTableColumn_GetIsPreserveWidthAuto(ImGuiTableColumn *self) { return self->IsPreserveWidthAuto; }
void ImGuiTableColumn_SetNavLayerCurrent(ImGuiTableColumn *ImGuiTableColumnPtr, ImS8 v) { ImGuiTableColumnPtr->NavLayerCurrent = v; }
ImS8 ImGuiTableColumn_GetNavLayerCurrent(ImGuiTableColumn *self) { return self->NavLayerCurrent; }
void ImGuiTableColumn_SetAutoFitQueue(ImGuiTableColumn *ImGuiTableColumnPtr, ImU8 v) { ImGuiTableColumnPtr->AutoFitQueue = v; }
ImU8 ImGuiTableColumn_GetAutoFitQueue(ImGuiTableColumn *self) { return self->AutoFitQueue; }
void ImGuiTableColumn_SetCannotSkipItemsQueue(ImGuiTableColumn *ImGuiTableColumnPtr, ImU8 v) { ImGuiTableColumnPtr->CannotSkipItemsQueue = v; }
ImU8 ImGuiTableColumn_GetCannotSkipItemsQueue(ImGuiTableColumn *self) { return self->CannotSkipItemsQueue; }
void ImGuiTableColumn_SetSortDirection(ImGuiTableColumn *ImGuiTableColumnPtr, ImU8 v) { ImGuiTableColumnPtr->SortDirection = v; }
ImU8 ImGuiTableColumn_GetSortDirection(ImGuiTableColumn *self) { return self->SortDirection; }
void ImGuiTableColumn_SetSortDirectionsAvailCount(ImGuiTableColumn *ImGuiTableColumnPtr, ImU8 v) { ImGuiTableColumnPtr->SortDirectionsAvailCount = v; }
ImU8 ImGuiTableColumn_GetSortDirectionsAvailCount(ImGuiTableColumn *self) { return self->SortDirectionsAvailCount; }
void ImGuiTableColumn_SetSortDirectionsAvailMask(ImGuiTableColumn *ImGuiTableColumnPtr, ImU8 v) { ImGuiTableColumnPtr->SortDirectionsAvailMask = v; }
ImU8 ImGuiTableColumn_GetSortDirectionsAvailMask(ImGuiTableColumn *self) { return self->SortDirectionsAvailMask; }
void ImGuiTableColumn_SetSortDirectionsAvailList(ImGuiTableColumn *ImGuiTableColumnPtr, ImU8 v) { ImGuiTableColumnPtr->SortDirectionsAvailList = v; }
ImU8 ImGuiTableColumn_GetSortDirectionsAvailList(ImGuiTableColumn *self) { return self->SortDirectionsAvailList; }
void ImColor_SetValue(ImColor *ImColorPtr, ImVec4 v) { ImColorPtr->Value = v; }
ImVec4 ImColor_GetValue(ImColor *self) { return self->Value; }
void ImGuiNextWindowData_SetFlags(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiNextWindowDataFlags v) { ImGuiNextWindowDataPtr->Flags = v; }
ImGuiNextWindowDataFlags ImGuiNextWindowData_GetFlags(ImGuiNextWindowData *self) { return self->Flags; }
void ImGuiNextWindowData_SetPosCond(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiCond v) { ImGuiNextWindowDataPtr->PosCond = v; }
ImGuiCond ImGuiNextWindowData_GetPosCond(ImGuiNextWindowData *self) { return self->PosCond; }
void ImGuiNextWindowData_SetSizeCond(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiCond v) { ImGuiNextWindowDataPtr->SizeCond = v; }
ImGuiCond ImGuiNextWindowData_GetSizeCond(ImGuiNextWindowData *self) { return self->SizeCond; }
void ImGuiNextWindowData_SetCollapsedCond(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiCond v) { ImGuiNextWindowDataPtr->CollapsedCond = v; }
ImGuiCond ImGuiNextWindowData_GetCollapsedCond(ImGuiNextWindowData *self) { return self->CollapsedCond; }
void ImGuiNextWindowData_SetDockCond(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiCond v) { ImGuiNextWindowDataPtr->DockCond = v; }
ImGuiCond ImGuiNextWindowData_GetDockCond(ImGuiNextWindowData *self) { return self->DockCond; }
void ImGuiNextWindowData_SetPosVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImVec2 v) { ImGuiNextWindowDataPtr->PosVal = v; }
ImVec2 ImGuiNextWindowData_GetPosVal(ImGuiNextWindowData *self) { return self->PosVal; }
void ImGuiNextWindowData_SetPosPivotVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImVec2 v) { ImGuiNextWindowDataPtr->PosPivotVal = v; }
ImVec2 ImGuiNextWindowData_GetPosPivotVal(ImGuiNextWindowData *self) { return self->PosPivotVal; }
void ImGuiNextWindowData_SetSizeVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImVec2 v) { ImGuiNextWindowDataPtr->SizeVal = v; }
ImVec2 ImGuiNextWindowData_GetSizeVal(ImGuiNextWindowData *self) { return self->SizeVal; }
void ImGuiNextWindowData_SetContentSizeVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImVec2 v) { ImGuiNextWindowDataPtr->ContentSizeVal = v; }
ImVec2 ImGuiNextWindowData_GetContentSizeVal(ImGuiNextWindowData *self) { return self->ContentSizeVal; }
void ImGuiNextWindowData_SetScrollVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImVec2 v) { ImGuiNextWindowDataPtr->ScrollVal = v; }
ImVec2 ImGuiNextWindowData_GetScrollVal(ImGuiNextWindowData *self) { return self->ScrollVal; }
void ImGuiNextWindowData_SetPosUndock(ImGuiNextWindowData *ImGuiNextWindowDataPtr, bool v) { ImGuiNextWindowDataPtr->PosUndock = v; }
bool ImGuiNextWindowData_GetPosUndock(ImGuiNextWindowData *self) { return self->PosUndock; }
void ImGuiNextWindowData_SetCollapsedVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, bool v) { ImGuiNextWindowDataPtr->CollapsedVal = v; }
bool ImGuiNextWindowData_GetCollapsedVal(ImGuiNextWindowData *self) { return self->CollapsedVal; }
void ImGuiNextWindowData_SetSizeConstraintRect(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImRect v) { ImGuiNextWindowDataPtr->SizeConstraintRect = v; }
ImRect ImGuiNextWindowData_GetSizeConstraintRect(ImGuiNextWindowData *self) { return self->SizeConstraintRect; }
void ImGuiNextWindowData_SetSizeCallback(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiSizeCallback v) { ImGuiNextWindowDataPtr->SizeCallback = v; }
ImGuiSizeCallback ImGuiNextWindowData_GetSizeCallback(ImGuiNextWindowData *self) { return self->SizeCallback; }
void ImGuiNextWindowData_SetSizeCallbackUserData(ImGuiNextWindowData *ImGuiNextWindowDataPtr, void* v) { ImGuiNextWindowDataPtr->SizeCallbackUserData = v; }
void* ImGuiNextWindowData_GetSizeCallbackUserData(ImGuiNextWindowData *self) { return self->SizeCallbackUserData; }
void ImGuiNextWindowData_SetBgAlphaVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, float v) { ImGuiNextWindowDataPtr->BgAlphaVal = v; }
float ImGuiNextWindowData_GetBgAlphaVal(ImGuiNextWindowData *self) { return self->BgAlphaVal; }
void ImGuiNextWindowData_SetViewportId(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiID v) { ImGuiNextWindowDataPtr->ViewportId = v; }
ImGuiID ImGuiNextWindowData_GetViewportId(ImGuiNextWindowData *self) { return self->ViewportId; }
void ImGuiNextWindowData_SetDockId(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiID v) { ImGuiNextWindowDataPtr->DockId = v; }
ImGuiID ImGuiNextWindowData_GetDockId(ImGuiNextWindowData *self) { return self->DockId; }
void ImGuiNextWindowData_SetWindowClass(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImGuiWindowClass v) { ImGuiNextWindowDataPtr->WindowClass = v; }
ImGuiWindowClass ImGuiNextWindowData_GetWindowClass(ImGuiNextWindowData *self) { return self->WindowClass; }
void ImGuiNextWindowData_SetMenuBarOffsetMinVal(ImGuiNextWindowData *ImGuiNextWindowDataPtr, ImVec2 v) { ImGuiNextWindowDataPtr->MenuBarOffsetMinVal = v; }
ImVec2 ImGuiNextWindowData_GetMenuBarOffsetMinVal(ImGuiNextWindowData *self) { return self->MenuBarOffsetMinVal; }
void ImGuiStyleMod_SetVarIdx(ImGuiStyleMod *ImGuiStyleModPtr, ImGuiStyleVar v) { ImGuiStyleModPtr->VarIdx = v; }
ImGuiStyleVar ImGuiStyleMod_GetVarIdx(ImGuiStyleMod *self) { return self->VarIdx; }
void ImGuiTextRange_Setb(ImGuiTextRange *ImGuiTextRangePtr, const char* v) { ImGuiTextRangePtr->b = v; }
const char* ImGuiTextRange_Getb(ImGuiTextRange *self) { return self->b; }
void ImGuiTextRange_Sete(ImGuiTextRange *ImGuiTextRangePtr, const char* v) { ImGuiTextRangePtr->e = v; }
const char* ImGuiTextRange_Gete(ImGuiTextRange *self) { return self->e; }
void ImGuiWindow_SetName(ImGuiWindow *ImGuiWindowPtr, char* v) { ImGuiWindowPtr->Name = v; }
char* ImGuiWindow_GetName(ImGuiWindow *self) { return self->Name; }
void ImGuiWindow_SetID(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->ID = v; }
ImGuiID ImGuiWindow_GetID(ImGuiWindow *self) { return self->ID; }
void ImGuiWindow_SetFlags(ImGuiWindow *ImGuiWindowPtr, ImGuiWindowFlags v) { ImGuiWindowPtr->Flags = v; }
ImGuiWindowFlags ImGuiWindow_GetFlags(ImGuiWindow *self) { return self->Flags; }
void ImGuiWindow_SetFlagsPreviousFrame(ImGuiWindow *ImGuiWindowPtr, ImGuiWindowFlags v) { ImGuiWindowPtr->FlagsPreviousFrame = v; }
ImGuiWindowFlags ImGuiWindow_GetFlagsPreviousFrame(ImGuiWindow *self) { return self->FlagsPreviousFrame; }
void ImGuiWindow_SetWindowClass(ImGuiWindow *ImGuiWindowPtr, ImGuiWindowClass v) { ImGuiWindowPtr->WindowClass = v; }
ImGuiWindowClass ImGuiWindow_GetWindowClass(ImGuiWindow *self) { return self->WindowClass; }
void ImGuiWindow_SetViewport(ImGuiWindow *ImGuiWindowPtr, ImGuiViewportP* v) { ImGuiWindowPtr->Viewport = v; }
ImGuiViewportP* ImGuiWindow_GetViewport(ImGuiWindow *self) { return self->Viewport; }
void ImGuiWindow_SetViewportId(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->ViewportId = v; }
ImGuiID ImGuiWindow_GetViewportId(ImGuiWindow *self) { return self->ViewportId; }
void ImGuiWindow_SetViewportPos(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ViewportPos = v; }
ImVec2 ImGuiWindow_GetViewportPos(ImGuiWindow *self) { return self->ViewportPos; }
void ImGuiWindow_SetViewportAllowPlatformMonitorExtend(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->ViewportAllowPlatformMonitorExtend = v; }
int ImGuiWindow_GetViewportAllowPlatformMonitorExtend(ImGuiWindow *self) { return self->ViewportAllowPlatformMonitorExtend; }
void ImGuiWindow_SetPos(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->Pos = v; }
ImVec2 ImGuiWindow_GetPos(ImGuiWindow *self) { return self->Pos; }
void ImGuiWindow_SetSize(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->Size = v; }
ImVec2 ImGuiWindow_GetSize(ImGuiWindow *self) { return self->Size; }
void ImGuiWindow_SetSizeFull(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->SizeFull = v; }
ImVec2 ImGuiWindow_GetSizeFull(ImGuiWindow *self) { return self->SizeFull; }
void ImGuiWindow_SetContentSize(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ContentSize = v; }
ImVec2 ImGuiWindow_GetContentSize(ImGuiWindow *self) { return self->ContentSize; }
void ImGuiWindow_SetContentSizeIdeal(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ContentSizeIdeal = v; }
ImVec2 ImGuiWindow_GetContentSizeIdeal(ImGuiWindow *self) { return self->ContentSizeIdeal; }
void ImGuiWindow_SetContentSizeExplicit(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ContentSizeExplicit = v; }
ImVec2 ImGuiWindow_GetContentSizeExplicit(ImGuiWindow *self) { return self->ContentSizeExplicit; }
void ImGuiWindow_SetWindowPadding(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->WindowPadding = v; }
ImVec2 ImGuiWindow_GetWindowPadding(ImGuiWindow *self) { return self->WindowPadding; }
void ImGuiWindow_SetWindowRounding(ImGuiWindow *ImGuiWindowPtr, float v) { ImGuiWindowPtr->WindowRounding = v; }
float ImGuiWindow_GetWindowRounding(ImGuiWindow *self) { return self->WindowRounding; }
void ImGuiWindow_SetWindowBorderSize(ImGuiWindow *ImGuiWindowPtr, float v) { ImGuiWindowPtr->WindowBorderSize = v; }
float ImGuiWindow_GetWindowBorderSize(ImGuiWindow *self) { return self->WindowBorderSize; }
void ImGuiWindow_SetNameBufLen(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->NameBufLen = v; }
int ImGuiWindow_GetNameBufLen(ImGuiWindow *self) { return self->NameBufLen; }
void ImGuiWindow_SetMoveId(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->MoveId = v; }
ImGuiID ImGuiWindow_GetMoveId(ImGuiWindow *self) { return self->MoveId; }
void ImGuiWindow_SetTabId(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->TabId = v; }
ImGuiID ImGuiWindow_GetTabId(ImGuiWindow *self) { return self->TabId; }
void ImGuiWindow_SetChildId(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->ChildId = v; }
ImGuiID ImGuiWindow_GetChildId(ImGuiWindow *self) { return self->ChildId; }
void ImGuiWindow_SetScroll(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->Scroll = v; }
ImVec2 ImGuiWindow_GetScroll(ImGuiWindow *self) { return self->Scroll; }
void ImGuiWindow_SetScrollMax(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ScrollMax = v; }
ImVec2 ImGuiWindow_GetScrollMax(ImGuiWindow *self) { return self->ScrollMax; }
void ImGuiWindow_SetScrollTarget(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ScrollTarget = v; }
ImVec2 ImGuiWindow_GetScrollTarget(ImGuiWindow *self) { return self->ScrollTarget; }
void ImGuiWindow_SetScrollTargetCenterRatio(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ScrollTargetCenterRatio = v; }
ImVec2 ImGuiWindow_GetScrollTargetCenterRatio(ImGuiWindow *self) { return self->ScrollTargetCenterRatio; }
void ImGuiWindow_SetScrollTargetEdgeSnapDist(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ScrollTargetEdgeSnapDist = v; }
ImVec2 ImGuiWindow_GetScrollTargetEdgeSnapDist(ImGuiWindow *self) { return self->ScrollTargetEdgeSnapDist; }
void ImGuiWindow_SetScrollbarSizes(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->ScrollbarSizes = v; }
ImVec2 ImGuiWindow_GetScrollbarSizes(ImGuiWindow *self) { return self->ScrollbarSizes; }
void ImGuiWindow_SetScrollbarX(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->ScrollbarX = v; }
bool ImGuiWindow_GetScrollbarX(ImGuiWindow *self) { return self->ScrollbarX; }
void ImGuiWindow_SetScrollbarY(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->ScrollbarY = v; }
bool ImGuiWindow_GetScrollbarY(ImGuiWindow *self) { return self->ScrollbarY; }
void ImGuiWindow_SetViewportOwned(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->ViewportOwned = v; }
bool ImGuiWindow_GetViewportOwned(ImGuiWindow *self) { return self->ViewportOwned; }
void ImGuiWindow_SetActive(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->Active = v; }
bool ImGuiWindow_GetActive(ImGuiWindow *self) { return self->Active; }
void ImGuiWindow_SetWasActive(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->WasActive = v; }
bool ImGuiWindow_GetWasActive(ImGuiWindow *self) { return self->WasActive; }
void ImGuiWindow_SetWriteAccessed(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->WriteAccessed = v; }
bool ImGuiWindow_GetWriteAccessed(ImGuiWindow *self) { return self->WriteAccessed; }
void ImGuiWindow_SetCollapsed(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->Collapsed = v; }
bool ImGuiWindow_GetCollapsed(ImGuiWindow *self) { return self->Collapsed; }
void ImGuiWindow_SetWantCollapseToggle(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->WantCollapseToggle = v; }
bool ImGuiWindow_GetWantCollapseToggle(ImGuiWindow *self) { return self->WantCollapseToggle; }
void ImGuiWindow_SetSkipItems(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->SkipItems = v; }
bool ImGuiWindow_GetSkipItems(ImGuiWindow *self) { return self->SkipItems; }
void ImGuiWindow_SetAppearing(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->Appearing = v; }
bool ImGuiWindow_GetAppearing(ImGuiWindow *self) { return self->Appearing; }
void ImGuiWindow_SetHidden(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->Hidden = v; }
bool ImGuiWindow_GetHidden(ImGuiWindow *self) { return self->Hidden; }
void ImGuiWindow_SetIsFallbackWindow(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->IsFallbackWindow = v; }
bool ImGuiWindow_GetIsFallbackWindow(ImGuiWindow *self) { return self->IsFallbackWindow; }
void ImGuiWindow_SetIsExplicitChild(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->IsExplicitChild = v; }
bool ImGuiWindow_GetIsExplicitChild(ImGuiWindow *self) { return self->IsExplicitChild; }
void ImGuiWindow_SetHasCloseButton(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->HasCloseButton = v; }
bool ImGuiWindow_GetHasCloseButton(ImGuiWindow *self) { return self->HasCloseButton; }
void ImGuiWindow_SetResizeBorderHeld(ImGuiWindow *ImGuiWindowPtr, signed char v) { ImGuiWindowPtr->ResizeBorderHeld = v; }
signed char ImGuiWindow_GetResizeBorderHeld(ImGuiWindow *self) { return self->ResizeBorderHeld; }
void ImGuiWindow_SetBeginCount(ImGuiWindow *ImGuiWindowPtr, short v) { ImGuiWindowPtr->BeginCount = v; }
short ImGuiWindow_GetBeginCount(ImGuiWindow *self) { return self->BeginCount; }
void ImGuiWindow_SetBeginOrderWithinParent(ImGuiWindow *ImGuiWindowPtr, short v) { ImGuiWindowPtr->BeginOrderWithinParent = v; }
short ImGuiWindow_GetBeginOrderWithinParent(ImGuiWindow *self) { return self->BeginOrderWithinParent; }
void ImGuiWindow_SetBeginOrderWithinContext(ImGuiWindow *ImGuiWindowPtr, short v) { ImGuiWindowPtr->BeginOrderWithinContext = v; }
short ImGuiWindow_GetBeginOrderWithinContext(ImGuiWindow *self) { return self->BeginOrderWithinContext; }
void ImGuiWindow_SetFocusOrder(ImGuiWindow *ImGuiWindowPtr, short v) { ImGuiWindowPtr->FocusOrder = v; }
short ImGuiWindow_GetFocusOrder(ImGuiWindow *self) { return self->FocusOrder; }
void ImGuiWindow_SetPopupId(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->PopupId = v; }
ImGuiID ImGuiWindow_GetPopupId(ImGuiWindow *self) { return self->PopupId; }
void ImGuiWindow_SetAutoFitFramesX(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->AutoFitFramesX = v; }
ImS8 ImGuiWindow_GetAutoFitFramesX(ImGuiWindow *self) { return self->AutoFitFramesX; }
void ImGuiWindow_SetAutoFitFramesY(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->AutoFitFramesY = v; }
ImS8 ImGuiWindow_GetAutoFitFramesY(ImGuiWindow *self) { return self->AutoFitFramesY; }
void ImGuiWindow_SetAutoFitChildAxises(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->AutoFitChildAxises = v; }
ImS8 ImGuiWindow_GetAutoFitChildAxises(ImGuiWindow *self) { return self->AutoFitChildAxises; }
void ImGuiWindow_SetAutoFitOnlyGrows(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->AutoFitOnlyGrows = v; }
bool ImGuiWindow_GetAutoFitOnlyGrows(ImGuiWindow *self) { return self->AutoFitOnlyGrows; }
void ImGuiWindow_SetAutoPosLastDirection(ImGuiWindow *ImGuiWindowPtr, ImGuiDir v) { ImGuiWindowPtr->AutoPosLastDirection = v; }
ImGuiDir ImGuiWindow_GetAutoPosLastDirection(ImGuiWindow *self) { return self->AutoPosLastDirection; }
void ImGuiWindow_SetHiddenFramesCanSkipItems(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->HiddenFramesCanSkipItems = v; }
ImS8 ImGuiWindow_GetHiddenFramesCanSkipItems(ImGuiWindow *self) { return self->HiddenFramesCanSkipItems; }
void ImGuiWindow_SetHiddenFramesCannotSkipItems(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->HiddenFramesCannotSkipItems = v; }
ImS8 ImGuiWindow_GetHiddenFramesCannotSkipItems(ImGuiWindow *self) { return self->HiddenFramesCannotSkipItems; }
void ImGuiWindow_SetHiddenFramesForRenderOnly(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->HiddenFramesForRenderOnly = v; }
ImS8 ImGuiWindow_GetHiddenFramesForRenderOnly(ImGuiWindow *self) { return self->HiddenFramesForRenderOnly; }
void ImGuiWindow_SetDisableInputsFrames(ImGuiWindow *ImGuiWindowPtr, ImS8 v) { ImGuiWindowPtr->DisableInputsFrames = v; }
ImS8 ImGuiWindow_GetDisableInputsFrames(ImGuiWindow *self) { return self->DisableInputsFrames; }
void ImGuiWindow_SetSetWindowPosAllowFlags(ImGuiWindow *ImGuiWindowPtr, ImGuiCond v) { ImGuiWindowPtr->SetWindowPosAllowFlags = v; }
ImGuiCond ImGuiWindow_GetSetWindowPosAllowFlags(ImGuiWindow *self) { return self->SetWindowPosAllowFlags; }
void ImGuiWindow_SetSetWindowSizeAllowFlags(ImGuiWindow *ImGuiWindowPtr, ImGuiCond v) { ImGuiWindowPtr->SetWindowSizeAllowFlags = v; }
ImGuiCond ImGuiWindow_GetSetWindowSizeAllowFlags(ImGuiWindow *self) { return self->SetWindowSizeAllowFlags; }
void ImGuiWindow_SetSetWindowCollapsedAllowFlags(ImGuiWindow *ImGuiWindowPtr, ImGuiCond v) { ImGuiWindowPtr->SetWindowCollapsedAllowFlags = v; }
ImGuiCond ImGuiWindow_GetSetWindowCollapsedAllowFlags(ImGuiWindow *self) { return self->SetWindowCollapsedAllowFlags; }
void ImGuiWindow_SetSetWindowDockAllowFlags(ImGuiWindow *ImGuiWindowPtr, ImGuiCond v) { ImGuiWindowPtr->SetWindowDockAllowFlags = v; }
ImGuiCond ImGuiWindow_GetSetWindowDockAllowFlags(ImGuiWindow *self) { return self->SetWindowDockAllowFlags; }
void ImGuiWindow_SetSetWindowPosVal(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->SetWindowPosVal = v; }
ImVec2 ImGuiWindow_GetSetWindowPosVal(ImGuiWindow *self) { return self->SetWindowPosVal; }
void ImGuiWindow_SetSetWindowPosPivot(ImGuiWindow *ImGuiWindowPtr, ImVec2 v) { ImGuiWindowPtr->SetWindowPosPivot = v; }
ImVec2 ImGuiWindow_GetSetWindowPosPivot(ImGuiWindow *self) { return self->SetWindowPosPivot; }
void ImGuiWindow_SetIDStack(ImGuiWindow *ImGuiWindowPtr, ImVector_ImGuiID v) { ImGuiWindowPtr->IDStack = v; }
ImVector_ImGuiID ImGuiWindow_GetIDStack(ImGuiWindow *self) { return self->IDStack; }
void ImGuiWindow_SetDC(ImGuiWindow *ImGuiWindowPtr, ImGuiWindowTempData v) { ImGuiWindowPtr->DC = v; }
ImGuiWindowTempData ImGuiWindow_GetDC(ImGuiWindow *self) { return self->DC; }
void ImGuiWindow_SetOuterRectClipped(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->OuterRectClipped = v; }
ImRect ImGuiWindow_GetOuterRectClipped(ImGuiWindow *self) { return self->OuterRectClipped; }
void ImGuiWindow_SetInnerRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->InnerRect = v; }
ImRect ImGuiWindow_GetInnerRect(ImGuiWindow *self) { return self->InnerRect; }
void ImGuiWindow_SetInnerClipRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->InnerClipRect = v; }
ImRect ImGuiWindow_GetInnerClipRect(ImGuiWindow *self) { return self->InnerClipRect; }
void ImGuiWindow_SetWorkRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->WorkRect = v; }
ImRect ImGuiWindow_GetWorkRect(ImGuiWindow *self) { return self->WorkRect; }
void ImGuiWindow_SetParentWorkRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->ParentWorkRect = v; }
ImRect ImGuiWindow_GetParentWorkRect(ImGuiWindow *self) { return self->ParentWorkRect; }
void ImGuiWindow_SetClipRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->ClipRect = v; }
ImRect ImGuiWindow_GetClipRect(ImGuiWindow *self) { return self->ClipRect; }
void ImGuiWindow_SetContentRegionRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->ContentRegionRect = v; }
ImRect ImGuiWindow_GetContentRegionRect(ImGuiWindow *self) { return self->ContentRegionRect; }
void ImGuiWindow_SetHitTestHoleSize(ImGuiWindow *ImGuiWindowPtr, ImVec2ih v) { ImGuiWindowPtr->HitTestHoleSize = v; }
ImVec2ih ImGuiWindow_GetHitTestHoleSize(ImGuiWindow *self) { return self->HitTestHoleSize; }
void ImGuiWindow_SetHitTestHoleOffset(ImGuiWindow *ImGuiWindowPtr, ImVec2ih v) { ImGuiWindowPtr->HitTestHoleOffset = v; }
ImVec2ih ImGuiWindow_GetHitTestHoleOffset(ImGuiWindow *self) { return self->HitTestHoleOffset; }
void ImGuiWindow_SetLastFrameActive(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->LastFrameActive = v; }
int ImGuiWindow_GetLastFrameActive(ImGuiWindow *self) { return self->LastFrameActive; }
void ImGuiWindow_SetLastFrameJustFocused(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->LastFrameJustFocused = v; }
int ImGuiWindow_GetLastFrameJustFocused(ImGuiWindow *self) { return self->LastFrameJustFocused; }
void ImGuiWindow_SetLastTimeActive(ImGuiWindow *ImGuiWindowPtr, float v) { ImGuiWindowPtr->LastTimeActive = v; }
float ImGuiWindow_GetLastTimeActive(ImGuiWindow *self) { return self->LastTimeActive; }
void ImGuiWindow_SetItemWidthDefault(ImGuiWindow *ImGuiWindowPtr, float v) { ImGuiWindowPtr->ItemWidthDefault = v; }
float ImGuiWindow_GetItemWidthDefault(ImGuiWindow *self) { return self->ItemWidthDefault; }
void ImGuiWindow_SetStateStorage(ImGuiWindow *ImGuiWindowPtr, ImGuiStorage v) { ImGuiWindowPtr->StateStorage = v; }
ImGuiStorage ImGuiWindow_GetStateStorage(ImGuiWindow *self) { return self->StateStorage; }
void ImGuiWindow_SetColumnsStorage(ImGuiWindow *ImGuiWindowPtr, ImVector_ImGuiOldColumns v) { ImGuiWindowPtr->ColumnsStorage = v; }
ImVector_ImGuiOldColumns ImGuiWindow_GetColumnsStorage(ImGuiWindow *self) { return self->ColumnsStorage; }
void ImGuiWindow_SetFontWindowScale(ImGuiWindow *ImGuiWindowPtr, float v) { ImGuiWindowPtr->FontWindowScale = v; }
float ImGuiWindow_GetFontWindowScale(ImGuiWindow *self) { return self->FontWindowScale; }
void ImGuiWindow_SetFontDpiScale(ImGuiWindow *ImGuiWindowPtr, float v) { ImGuiWindowPtr->FontDpiScale = v; }
float ImGuiWindow_GetFontDpiScale(ImGuiWindow *self) { return self->FontDpiScale; }
void ImGuiWindow_SetSettingsOffset(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->SettingsOffset = v; }
int ImGuiWindow_GetSettingsOffset(ImGuiWindow *self) { return self->SettingsOffset; }
void ImGuiWindow_SetDrawList(ImGuiWindow *ImGuiWindowPtr, ImDrawList* v) { ImGuiWindowPtr->DrawList = v; }
ImDrawList* ImGuiWindow_GetDrawList(ImGuiWindow *self) { return self->DrawList; }
void ImGuiWindow_SetDrawListInst(ImGuiWindow *ImGuiWindowPtr, ImDrawList v) { ImGuiWindowPtr->DrawListInst = v; }
ImDrawList ImGuiWindow_GetDrawListInst(ImGuiWindow *self) { return self->DrawListInst; }
void ImGuiWindow_SetParentWindow(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->ParentWindow = v; }
ImGuiWindow* ImGuiWindow_GetParentWindow(ImGuiWindow *self) { return self->ParentWindow; }
void ImGuiWindow_SetParentWindowInBeginStack(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->ParentWindowInBeginStack = v; }
ImGuiWindow* ImGuiWindow_GetParentWindowInBeginStack(ImGuiWindow *self) { return self->ParentWindowInBeginStack; }
void ImGuiWindow_SetRootWindow(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->RootWindow = v; }
ImGuiWindow* ImGuiWindow_GetRootWindow(ImGuiWindow *self) { return self->RootWindow; }
void ImGuiWindow_SetRootWindowPopupTree(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->RootWindowPopupTree = v; }
ImGuiWindow* ImGuiWindow_GetRootWindowPopupTree(ImGuiWindow *self) { return self->RootWindowPopupTree; }
void ImGuiWindow_SetRootWindowDockTree(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->RootWindowDockTree = v; }
ImGuiWindow* ImGuiWindow_GetRootWindowDockTree(ImGuiWindow *self) { return self->RootWindowDockTree; }
void ImGuiWindow_SetRootWindowForTitleBarHighlight(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->RootWindowForTitleBarHighlight = v; }
ImGuiWindow* ImGuiWindow_GetRootWindowForTitleBarHighlight(ImGuiWindow *self) { return self->RootWindowForTitleBarHighlight; }
void ImGuiWindow_SetRootWindowForNav(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->RootWindowForNav = v; }
ImGuiWindow* ImGuiWindow_GetRootWindowForNav(ImGuiWindow *self) { return self->RootWindowForNav; }
void ImGuiWindow_SetNavLastChildNavWindow(ImGuiWindow *ImGuiWindowPtr, ImGuiWindow* v) { ImGuiWindowPtr->NavLastChildNavWindow = v; }
ImGuiWindow* ImGuiWindow_GetNavLastChildNavWindow(ImGuiWindow *self) { return self->NavLastChildNavWindow; }
void ImGuiWindow_SetMemoryDrawListIdxCapacity(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->MemoryDrawListIdxCapacity = v; }
int ImGuiWindow_GetMemoryDrawListIdxCapacity(ImGuiWindow *self) { return self->MemoryDrawListIdxCapacity; }
void ImGuiWindow_SetMemoryDrawListVtxCapacity(ImGuiWindow *ImGuiWindowPtr, int v) { ImGuiWindowPtr->MemoryDrawListVtxCapacity = v; }
int ImGuiWindow_GetMemoryDrawListVtxCapacity(ImGuiWindow *self) { return self->MemoryDrawListVtxCapacity; }
void ImGuiWindow_SetMemoryCompacted(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->MemoryCompacted = v; }
bool ImGuiWindow_GetMemoryCompacted(ImGuiWindow *self) { return self->MemoryCompacted; }
void ImGuiWindow_SetDockIsActive(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->DockIsActive = v; }
bool ImGuiWindow_GetDockIsActive(ImGuiWindow *self) { return self->DockIsActive; }
void ImGuiWindow_SetDockNodeIsVisible(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->DockNodeIsVisible = v; }
bool ImGuiWindow_GetDockNodeIsVisible(ImGuiWindow *self) { return self->DockNodeIsVisible; }
void ImGuiWindow_SetDockTabIsVisible(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->DockTabIsVisible = v; }
bool ImGuiWindow_GetDockTabIsVisible(ImGuiWindow *self) { return self->DockTabIsVisible; }
void ImGuiWindow_SetDockTabWantClose(ImGuiWindow *ImGuiWindowPtr, bool v) { ImGuiWindowPtr->DockTabWantClose = v; }
bool ImGuiWindow_GetDockTabWantClose(ImGuiWindow *self) { return self->DockTabWantClose; }
void ImGuiWindow_SetDockOrder(ImGuiWindow *ImGuiWindowPtr, short v) { ImGuiWindowPtr->DockOrder = v; }
short ImGuiWindow_GetDockOrder(ImGuiWindow *self) { return self->DockOrder; }
void ImGuiWindow_SetDockStyle(ImGuiWindow *ImGuiWindowPtr, ImGuiWindowDockStyle v) { ImGuiWindowPtr->DockStyle = v; }
ImGuiWindowDockStyle ImGuiWindow_GetDockStyle(ImGuiWindow *self) { return self->DockStyle; }
void ImGuiWindow_SetDockNode(ImGuiWindow *ImGuiWindowPtr, ImGuiDockNode* v) { ImGuiWindowPtr->DockNode = v; }
ImGuiDockNode* ImGuiWindow_GetDockNode(ImGuiWindow *self) { return self->DockNode; }
void ImGuiWindow_SetDockNodeAsHost(ImGuiWindow *ImGuiWindowPtr, ImGuiDockNode* v) { ImGuiWindowPtr->DockNodeAsHost = v; }
ImGuiDockNode* ImGuiWindow_GetDockNodeAsHost(ImGuiWindow *self) { return self->DockNodeAsHost; }
void ImGuiWindow_SetDockId(ImGuiWindow *ImGuiWindowPtr, ImGuiID v) { ImGuiWindowPtr->DockId = v; }
ImGuiID ImGuiWindow_GetDockId(ImGuiWindow *self) { return self->DockId; }
void ImGuiWindow_SetDockTabItemStatusFlags(ImGuiWindow *ImGuiWindowPtr, ImGuiItemStatusFlags v) { ImGuiWindowPtr->DockTabItemStatusFlags = v; }
ImGuiItemStatusFlags ImGuiWindow_GetDockTabItemStatusFlags(ImGuiWindow *self) { return self->DockTabItemStatusFlags; }
void ImGuiWindow_SetDockTabItemRect(ImGuiWindow *ImGuiWindowPtr, ImRect v) { ImGuiWindowPtr->DockTabItemRect = v; }
ImRect ImGuiWindow_GetDockTabItemRect(ImGuiWindow *self) { return self->DockTabItemRect; }
void StbUndoState_Setundo_point(StbUndoState *StbUndoStatePtr, short v) { StbUndoStatePtr->undo_point = v; }
short StbUndoState_Getundo_point(StbUndoState *self) { return self->undo_point; }
void StbUndoState_Setredo_point(StbUndoState *StbUndoStatePtr, short v) { StbUndoStatePtr->redo_point = v; }
short StbUndoState_Getredo_point(StbUndoState *self) { return self->redo_point; }
void StbUndoState_Setundo_char_point(StbUndoState *StbUndoStatePtr, int v) { StbUndoStatePtr->undo_char_point = v; }
int StbUndoState_Getundo_char_point(StbUndoState *self) { return self->undo_char_point; }
void StbUndoState_Setredo_char_point(StbUndoState *StbUndoStatePtr, int v) { StbUndoStatePtr->redo_char_point = v; }
int StbUndoState_Getredo_char_point(StbUndoState *self) { return self->redo_char_point; }
void ImGuiLastItemData_SetID(ImGuiLastItemData *ImGuiLastItemDataPtr, ImGuiID v) { ImGuiLastItemDataPtr->ID = v; }
ImGuiID ImGuiLastItemData_GetID(ImGuiLastItemData *self) { return self->ID; }
void ImGuiLastItemData_SetInFlags(ImGuiLastItemData *ImGuiLastItemDataPtr, ImGuiItemFlags v) { ImGuiLastItemDataPtr->InFlags = v; }
ImGuiItemFlags ImGuiLastItemData_GetInFlags(ImGuiLastItemData *self) { return self->InFlags; }
void ImGuiLastItemData_SetStatusFlags(ImGuiLastItemData *ImGuiLastItemDataPtr, ImGuiItemStatusFlags v) { ImGuiLastItemDataPtr->StatusFlags = v; }
ImGuiItemStatusFlags ImGuiLastItemData_GetStatusFlags(ImGuiLastItemData *self) { return self->StatusFlags; }
void ImGuiLastItemData_SetRect(ImGuiLastItemData *ImGuiLastItemDataPtr, ImRect v) { ImGuiLastItemDataPtr->Rect = v; }
ImRect ImGuiLastItemData_GetRect(ImGuiLastItemData *self) { return self->Rect; }
void ImGuiLastItemData_SetNavRect(ImGuiLastItemData *ImGuiLastItemDataPtr, ImRect v) { ImGuiLastItemDataPtr->NavRect = v; }
ImRect ImGuiLastItemData_GetNavRect(ImGuiLastItemData *self) { return self->NavRect; }
void ImGuiLastItemData_SetDisplayRect(ImGuiLastItemData *ImGuiLastItemDataPtr, ImRect v) { ImGuiLastItemDataPtr->DisplayRect = v; }
ImRect ImGuiLastItemData_GetDisplayRect(ImGuiLastItemData *self) { return self->DisplayRect; }
void ImGuiTextBuffer_SetBuf(ImGuiTextBuffer *ImGuiTextBufferPtr, ImVector_char v) { ImGuiTextBufferPtr->Buf = v; }
ImVector_char ImGuiTextBuffer_GetBuf(ImGuiTextBuffer *self) { return self->Buf; }
void ImGuiInputEventText_SetChar(ImGuiInputEventText *ImGuiInputEventTextPtr, unsigned int v) { ImGuiInputEventTextPtr->Char = v; }
unsigned int ImGuiInputEventText_GetChar(ImGuiInputEventText *self) { return self->Char; }
void ImDrawList_SetCmdBuffer(ImDrawList *ImDrawListPtr, ImVector_ImDrawCmd v) { ImDrawListPtr->CmdBuffer = v; }
ImVector_ImDrawCmd ImDrawList_GetCmdBuffer(ImDrawList *self) { return self->CmdBuffer; }
void ImDrawList_SetIdxBuffer(ImDrawList *ImDrawListPtr, ImVector_ImDrawIdx v) { ImDrawListPtr->IdxBuffer = v; }
ImVector_ImDrawIdx ImDrawList_GetIdxBuffer(ImDrawList *self) { return self->IdxBuffer; }
void ImDrawList_SetVtxBuffer(ImDrawList *ImDrawListPtr, ImVector_ImDrawVert v) { ImDrawListPtr->VtxBuffer = v; }
ImVector_ImDrawVert ImDrawList_GetVtxBuffer(ImDrawList *self) { return self->VtxBuffer; }
void ImDrawList_SetFlags(ImDrawList *ImDrawListPtr, ImDrawListFlags v) { ImDrawListPtr->Flags = v; }
ImDrawListFlags ImDrawList_GetFlags(ImDrawList *self) { return self->Flags; }
void ImDrawList_Set_VtxCurrentIdx(ImDrawList *ImDrawListPtr, unsigned int v) { ImDrawListPtr->_VtxCurrentIdx = v; }
unsigned int ImDrawList_Get_VtxCurrentIdx(ImDrawList *self) { return self->_VtxCurrentIdx; }
void ImDrawList_Set_Data(ImDrawList *ImDrawListPtr, const ImDrawListSharedData* v) { ImDrawListPtr->_Data = v; }
const ImDrawListSharedData* ImDrawList_Get_Data(ImDrawList *self) { return self->_Data; }
void ImDrawList_Set_OwnerName(ImDrawList *ImDrawListPtr, const char* v) { ImDrawListPtr->_OwnerName = v; }
const char* ImDrawList_Get_OwnerName(ImDrawList *self) { return self->_OwnerName; }
void ImDrawList_Set_VtxWritePtr(ImDrawList *ImDrawListPtr, ImDrawVert* v) { ImDrawListPtr->_VtxWritePtr = v; }
ImDrawVert* ImDrawList_Get_VtxWritePtr(ImDrawList *self) { return self->_VtxWritePtr; }
void ImDrawList_Set_IdxWritePtr(ImDrawList *ImDrawListPtr, ImDrawIdx* v) { ImDrawListPtr->_IdxWritePtr = v; }
ImDrawIdx* ImDrawList_Get_IdxWritePtr(ImDrawList *self) { return self->_IdxWritePtr; }
void ImDrawList_Set_ClipRectStack(ImDrawList *ImDrawListPtr, ImVector_ImVec4 v) { ImDrawListPtr->_ClipRectStack = v; }
ImVector_ImVec4 ImDrawList_Get_ClipRectStack(ImDrawList *self) { return self->_ClipRectStack; }
void ImDrawList_Set_TextureIdStack(ImDrawList *ImDrawListPtr, ImVector_ImTextureID v) { ImDrawListPtr->_TextureIdStack = v; }
ImVector_ImTextureID ImDrawList_Get_TextureIdStack(ImDrawList *self) { return self->_TextureIdStack; }
void ImDrawList_Set_Path(ImDrawList *ImDrawListPtr, ImVector_ImVec2 v) { ImDrawListPtr->_Path = v; }
ImVector_ImVec2 ImDrawList_Get_Path(ImDrawList *self) { return self->_Path; }
void ImDrawList_Set_CmdHeader(ImDrawList *ImDrawListPtr, ImDrawCmdHeader v) { ImDrawListPtr->_CmdHeader = v; }
ImDrawCmdHeader ImDrawList_Get_CmdHeader(ImDrawList *self) { return self->_CmdHeader; }
void ImDrawList_Set_Splitter(ImDrawList *ImDrawListPtr, ImDrawListSplitter v) { ImDrawListPtr->_Splitter = v; }
ImDrawListSplitter ImDrawList_Get_Splitter(ImDrawList *self) { return self->_Splitter; }
void ImDrawList_Set_FringeScale(ImDrawList *ImDrawListPtr, float v) { ImDrawListPtr->_FringeScale = v; }
float ImDrawList_Get_FringeScale(ImDrawList *self) { return self->_FringeScale; }
void ImGuiNextItemData_SetFlags(ImGuiNextItemData *ImGuiNextItemDataPtr, ImGuiNextItemDataFlags v) { ImGuiNextItemDataPtr->Flags = v; }
ImGuiNextItemDataFlags ImGuiNextItemData_GetFlags(ImGuiNextItemData *self) { return self->Flags; }
void ImGuiNextItemData_SetWidth(ImGuiNextItemData *ImGuiNextItemDataPtr, float v) { ImGuiNextItemDataPtr->Width = v; }
float ImGuiNextItemData_GetWidth(ImGuiNextItemData *self) { return self->Width; }
void ImGuiNextItemData_SetFocusScopeId(ImGuiNextItemData *ImGuiNextItemDataPtr, ImGuiID v) { ImGuiNextItemDataPtr->FocusScopeId = v; }
ImGuiID ImGuiNextItemData_GetFocusScopeId(ImGuiNextItemData *self) { return self->FocusScopeId; }
void ImGuiNextItemData_SetOpenCond(ImGuiNextItemData *ImGuiNextItemDataPtr, ImGuiCond v) { ImGuiNextItemDataPtr->OpenCond = v; }
ImGuiCond ImGuiNextItemData_GetOpenCond(ImGuiNextItemData *self) { return self->OpenCond; }
void ImGuiNextItemData_SetOpenVal(ImGuiNextItemData *ImGuiNextItemDataPtr, bool v) { ImGuiNextItemDataPtr->OpenVal = v; }
bool ImGuiNextItemData_GetOpenVal(ImGuiNextItemData *self) { return self->OpenVal; }
void ImGuiWindowSettings_SetID(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImGuiID v) { ImGuiWindowSettingsPtr->ID = v; }
ImGuiID ImGuiWindowSettings_GetID(ImGuiWindowSettings *self) { return self->ID; }
void ImGuiWindowSettings_SetPos(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImVec2ih v) { ImGuiWindowSettingsPtr->Pos = v; }
ImVec2ih ImGuiWindowSettings_GetPos(ImGuiWindowSettings *self) { return self->Pos; }
void ImGuiWindowSettings_SetSize(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImVec2ih v) { ImGuiWindowSettingsPtr->Size = v; }
ImVec2ih ImGuiWindowSettings_GetSize(ImGuiWindowSettings *self) { return self->Size; }
void ImGuiWindowSettings_SetViewportPos(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImVec2ih v) { ImGuiWindowSettingsPtr->ViewportPos = v; }
ImVec2ih ImGuiWindowSettings_GetViewportPos(ImGuiWindowSettings *self) { return self->ViewportPos; }
void ImGuiWindowSettings_SetViewportId(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImGuiID v) { ImGuiWindowSettingsPtr->ViewportId = v; }
ImGuiID ImGuiWindowSettings_GetViewportId(ImGuiWindowSettings *self) { return self->ViewportId; }
void ImGuiWindowSettings_SetDockId(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImGuiID v) { ImGuiWindowSettingsPtr->DockId = v; }
ImGuiID ImGuiWindowSettings_GetDockId(ImGuiWindowSettings *self) { return self->DockId; }
void ImGuiWindowSettings_SetClassId(ImGuiWindowSettings *ImGuiWindowSettingsPtr, ImGuiID v) { ImGuiWindowSettingsPtr->ClassId = v; }
ImGuiID ImGuiWindowSettings_GetClassId(ImGuiWindowSettings *self) { return self->ClassId; }
void ImGuiWindowSettings_SetDockOrder(ImGuiWindowSettings *ImGuiWindowSettingsPtr, short v) { ImGuiWindowSettingsPtr->DockOrder = v; }
short ImGuiWindowSettings_GetDockOrder(ImGuiWindowSettings *self) { return self->DockOrder; }
void ImGuiWindowSettings_SetCollapsed(ImGuiWindowSettings *ImGuiWindowSettingsPtr, bool v) { ImGuiWindowSettingsPtr->Collapsed = v; }
bool ImGuiWindowSettings_GetCollapsed(ImGuiWindowSettings *self) { return self->Collapsed; }
void ImGuiWindowSettings_SetWantApply(ImGuiWindowSettings *ImGuiWindowSettingsPtr, bool v) { ImGuiWindowSettingsPtr->WantApply = v; }
bool ImGuiWindowSettings_GetWantApply(ImGuiWindowSettings *self) { return self->WantApply; }
void ImDrawCmd_SetClipRect(ImDrawCmd *ImDrawCmdPtr, ImVec4 v) { ImDrawCmdPtr->ClipRect = v; }
ImVec4 ImDrawCmd_GetClipRect(ImDrawCmd *self) { return self->ClipRect; }
void ImDrawCmd_SetTextureId(ImDrawCmd *ImDrawCmdPtr, ImTextureID v) { ImDrawCmdPtr->TextureId = v; }
ImTextureID ImDrawCmd_GetTextureId(ImDrawCmd *self) { return self->TextureId; }
void ImDrawCmd_SetVtxOffset(ImDrawCmd *ImDrawCmdPtr, unsigned int v) { ImDrawCmdPtr->VtxOffset = v; }
unsigned int ImDrawCmd_GetVtxOffset(ImDrawCmd *self) { return self->VtxOffset; }
void ImDrawCmd_SetIdxOffset(ImDrawCmd *ImDrawCmdPtr, unsigned int v) { ImDrawCmdPtr->IdxOffset = v; }
unsigned int ImDrawCmd_GetIdxOffset(ImDrawCmd *self) { return self->IdxOffset; }
void ImDrawCmd_SetElemCount(ImDrawCmd *ImDrawCmdPtr, unsigned int v) { ImDrawCmdPtr->ElemCount = v; }
unsigned int ImDrawCmd_GetElemCount(ImDrawCmd *self) { return self->ElemCount; }
void ImDrawCmd_SetUserCallback(ImDrawCmd *ImDrawCmdPtr, ImDrawCallback v) { ImDrawCmdPtr->UserCallback = v; }
ImDrawCallback ImDrawCmd_GetUserCallback(ImDrawCmd *self) { return self->UserCallback; }
void ImDrawCmd_SetUserCallbackData(ImDrawCmd *ImDrawCmdPtr, void* v) { ImDrawCmdPtr->UserCallbackData = v; }
void* ImDrawCmd_GetUserCallbackData(ImDrawCmd *self) { return self->UserCallbackData; }
void ImGuiContextHook_SetHookId(ImGuiContextHook *ImGuiContextHookPtr, ImGuiID v) { ImGuiContextHookPtr->HookId = v; }
ImGuiID ImGuiContextHook_GetHookId(ImGuiContextHook *self) { return self->HookId; }
void ImGuiContextHook_SetType(ImGuiContextHook *ImGuiContextHookPtr, ImGuiContextHookType v) { ImGuiContextHookPtr->Type = v; }
ImGuiContextHookType ImGuiContextHook_GetType(ImGuiContextHook *self) { return self->Type; }
void ImGuiContextHook_SetOwner(ImGuiContextHook *ImGuiContextHookPtr, ImGuiID v) { ImGuiContextHookPtr->Owner = v; }
ImGuiID ImGuiContextHook_GetOwner(ImGuiContextHook *self) { return self->Owner; }
void ImGuiContextHook_SetCallback(ImGuiContextHook *ImGuiContextHookPtr, ImGuiContextHookCallback v) { ImGuiContextHookPtr->Callback = v; }
ImGuiContextHookCallback ImGuiContextHook_GetCallback(ImGuiContextHook *self) { return self->Callback; }
void ImGuiContextHook_SetUserData(ImGuiContextHook *ImGuiContextHookPtr, void* v) { ImGuiContextHookPtr->UserData = v; }
void* ImGuiContextHook_GetUserData(ImGuiContextHook *self) { return self->UserData; }
void ImGuiInputEventAppFocused_SetFocused(ImGuiInputEventAppFocused *ImGuiInputEventAppFocusedPtr, bool v) { ImGuiInputEventAppFocusedPtr->Focused = v; }
bool ImGuiInputEventAppFocused_GetFocused(ImGuiInputEventAppFocused *self) { return self->Focused; }
void ImGuiListClipper_SetDisplayStart(ImGuiListClipper *ImGuiListClipperPtr, int v) { ImGuiListClipperPtr->DisplayStart = v; }
int ImGuiListClipper_GetDisplayStart(ImGuiListClipper *self) { return self->DisplayStart; }
void ImGuiListClipper_SetDisplayEnd(ImGuiListClipper *ImGuiListClipperPtr, int v) { ImGuiListClipperPtr->DisplayEnd = v; }
int ImGuiListClipper_GetDisplayEnd(ImGuiListClipper *self) { return self->DisplayEnd; }
void ImGuiListClipper_SetItemsCount(ImGuiListClipper *ImGuiListClipperPtr, int v) { ImGuiListClipperPtr->ItemsCount = v; }
int ImGuiListClipper_GetItemsCount(ImGuiListClipper *self) { return self->ItemsCount; }
void ImGuiListClipper_SetItemsHeight(ImGuiListClipper *ImGuiListClipperPtr, float v) { ImGuiListClipperPtr->ItemsHeight = v; }
float ImGuiListClipper_GetItemsHeight(ImGuiListClipper *self) { return self->ItemsHeight; }
void ImGuiListClipper_SetStartPosY(ImGuiListClipper *ImGuiListClipperPtr, float v) { ImGuiListClipperPtr->StartPosY = v; }
float ImGuiListClipper_GetStartPosY(ImGuiListClipper *self) { return self->StartPosY; }
void ImGuiListClipper_SetTempData(ImGuiListClipper *ImGuiListClipperPtr, void* v) { ImGuiListClipperPtr->TempData = v; }
void* ImGuiListClipper_GetTempData(ImGuiListClipper *self) { return self->TempData; }
void ImGuiPlatformMonitor_SetMainPos(ImGuiPlatformMonitor *ImGuiPlatformMonitorPtr, ImVec2 v) { ImGuiPlatformMonitorPtr->MainPos = v; }
ImVec2 ImGuiPlatformMonitor_GetMainPos(ImGuiPlatformMonitor *self) { return self->MainPos; }
void ImGuiPlatformMonitor_SetMainSize(ImGuiPlatformMonitor *ImGuiPlatformMonitorPtr, ImVec2 v) { ImGuiPlatformMonitorPtr->MainSize = v; }
ImVec2 ImGuiPlatformMonitor_GetMainSize(ImGuiPlatformMonitor *self) { return self->MainSize; }
void ImGuiPlatformMonitor_SetWorkPos(ImGuiPlatformMonitor *ImGuiPlatformMonitorPtr, ImVec2 v) { ImGuiPlatformMonitorPtr->WorkPos = v; }
ImVec2 ImGuiPlatformMonitor_GetWorkPos(ImGuiPlatformMonitor *self) { return self->WorkPos; }
void ImGuiPlatformMonitor_SetWorkSize(ImGuiPlatformMonitor *ImGuiPlatformMonitorPtr, ImVec2 v) { ImGuiPlatformMonitorPtr->WorkSize = v; }
ImVec2 ImGuiPlatformMonitor_GetWorkSize(ImGuiPlatformMonitor *self) { return self->WorkSize; }
void ImGuiPlatformMonitor_SetDpiScale(ImGuiPlatformMonitor *ImGuiPlatformMonitorPtr, float v) { ImGuiPlatformMonitorPtr->DpiScale = v; }
float ImGuiPlatformMonitor_GetDpiScale(ImGuiPlatformMonitor *self) { return self->DpiScale; }
void ImGuiSettingsHandler_SetTypeName(ImGuiSettingsHandler *ImGuiSettingsHandlerPtr, const char* v) { ImGuiSettingsHandlerPtr->TypeName = v; }
const char* ImGuiSettingsHandler_GetTypeName(ImGuiSettingsHandler *self) { return self->TypeName; }
void ImGuiSettingsHandler_SetTypeHash(ImGuiSettingsHandler *ImGuiSettingsHandlerPtr, ImGuiID v) { ImGuiSettingsHandlerPtr->TypeHash = v; }
ImGuiID ImGuiSettingsHandler_GetTypeHash(ImGuiSettingsHandler *self) { return self->TypeHash; }
void ImGuiSettingsHandler_SetUserData(ImGuiSettingsHandler *ImGuiSettingsHandlerPtr, void* v) { ImGuiSettingsHandlerPtr->UserData = v; }
void* ImGuiSettingsHandler_GetUserData(ImGuiSettingsHandler *self) { return self->UserData; }
void ImGuiStackLevelInfo_SetID(ImGuiStackLevelInfo *ImGuiStackLevelInfoPtr, ImGuiID v) { ImGuiStackLevelInfoPtr->ID = v; }
ImGuiID ImGuiStackLevelInfo_GetID(ImGuiStackLevelInfo *self) { return self->ID; }
void ImGuiStackLevelInfo_SetQueryFrameCount(ImGuiStackLevelInfo *ImGuiStackLevelInfoPtr, ImS8 v) { ImGuiStackLevelInfoPtr->QueryFrameCount = v; }
ImS8 ImGuiStackLevelInfo_GetQueryFrameCount(ImGuiStackLevelInfo *self) { return self->QueryFrameCount; }
void ImGuiStackLevelInfo_SetQuerySuccess(ImGuiStackLevelInfo *ImGuiStackLevelInfoPtr, bool v) { ImGuiStackLevelInfoPtr->QuerySuccess = v; }
bool ImGuiStackLevelInfo_GetQuerySuccess(ImGuiStackLevelInfo *self) { return self->QuerySuccess; }
void ImGuiStackLevelInfo_SetDataType(ImGuiStackLevelInfo *ImGuiStackLevelInfoPtr, ImGuiDataType v) { ImGuiStackLevelInfoPtr->DataType = v; }
ImGuiDataType ImGuiStackLevelInfo_GetDataType(ImGuiStackLevelInfo *self) { return self->DataType; }
void ImBitVector_SetStorage(ImBitVector *ImBitVectorPtr, ImVector_ImU32 v) { ImBitVectorPtr->Storage = v; }
ImVector_ImU32 ImBitVector_GetStorage(ImBitVector *self) { return self->Storage; }
void STB_TexteditState_Setcursor(STB_TexteditState *STB_TexteditStatePtr, int v) { STB_TexteditStatePtr->cursor = v; }
int STB_TexteditState_Getcursor(STB_TexteditState *self) { return self->cursor; }
void STB_TexteditState_Setselect_start(STB_TexteditState *STB_TexteditStatePtr, int v) { STB_TexteditStatePtr->select_start = v; }
int STB_TexteditState_Getselect_start(STB_TexteditState *self) { return self->select_start; }
void STB_TexteditState_Setselect_end(STB_TexteditState *STB_TexteditStatePtr, int v) { STB_TexteditStatePtr->select_end = v; }
int STB_TexteditState_Getselect_end(STB_TexteditState *self) { return self->select_end; }
void STB_TexteditState_Setinsert_mode(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->insert_mode = v; }
unsigned char STB_TexteditState_Getinsert_mode(STB_TexteditState *self) { return self->insert_mode; }
void STB_TexteditState_Setrow_count_per_page(STB_TexteditState *STB_TexteditStatePtr, int v) { STB_TexteditStatePtr->row_count_per_page = v; }
int STB_TexteditState_Getrow_count_per_page(STB_TexteditState *self) { return self->row_count_per_page; }
void STB_TexteditState_Setcursor_at_end_of_line(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->cursor_at_end_of_line = v; }
unsigned char STB_TexteditState_Getcursor_at_end_of_line(STB_TexteditState *self) { return self->cursor_at_end_of_line; }
void STB_TexteditState_Setinitialized(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->initialized = v; }
unsigned char STB_TexteditState_Getinitialized(STB_TexteditState *self) { return self->initialized; }
void STB_TexteditState_Sethas_preferred_x(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->has_preferred_x = v; }
unsigned char STB_TexteditState_Gethas_preferred_x(STB_TexteditState *self) { return self->has_preferred_x; }
void STB_TexteditState_Setsingle_line(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->single_line = v; }
unsigned char STB_TexteditState_Getsingle_line(STB_TexteditState *self) { return self->single_line; }
void STB_TexteditState_Setpadding1(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->padding1 = v; }
unsigned char STB_TexteditState_Getpadding1(STB_TexteditState *self) { return self->padding1; }
void STB_TexteditState_Setpadding2(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->padding2 = v; }
unsigned char STB_TexteditState_Getpadding2(STB_TexteditState *self) { return self->padding2; }
void STB_TexteditState_Setpadding3(STB_TexteditState *STB_TexteditStatePtr, unsigned char v) { STB_TexteditStatePtr->padding3 = v; }
unsigned char STB_TexteditState_Getpadding3(STB_TexteditState *self) { return self->padding3; }
void STB_TexteditState_Setpreferred_x(STB_TexteditState *STB_TexteditStatePtr, float v) { STB_TexteditStatePtr->preferred_x = v; }
float STB_TexteditState_Getpreferred_x(STB_TexteditState *self) { return self->preferred_x; }
void STB_TexteditState_Setundostate(STB_TexteditState *STB_TexteditStatePtr, StbUndoState v) { STB_TexteditStatePtr->undostate = v; }
StbUndoState STB_TexteditState_Getundostate(STB_TexteditState *self) { return self->undostate; }
void ImGuiStorage_SetData(ImGuiStorage *ImGuiStoragePtr, ImVector_ImGuiStoragePair v) { ImGuiStoragePtr->Data = v; }
ImVector_ImGuiStoragePair ImGuiStorage_GetData(ImGuiStorage *self) { return self->Data; }
void ImGuiInputEvent_SetType(ImGuiInputEvent *ImGuiInputEventPtr, ImGuiInputEventType v) { ImGuiInputEventPtr->Type = v; }
ImGuiInputEventType ImGuiInputEvent_GetType(ImGuiInputEvent *self) { return self->Type; }
void ImGuiInputEvent_SetSource(ImGuiInputEvent *ImGuiInputEventPtr, ImGuiInputSource v) { ImGuiInputEventPtr->Source = v; }
ImGuiInputSource ImGuiInputEvent_GetSource(ImGuiInputEvent *self) { return self->Source; }
void ImGuiInputEvent_SetAddedByTestEngine(ImGuiInputEvent *ImGuiInputEventPtr, bool v) { ImGuiInputEventPtr->AddedByTestEngine = v; }
bool ImGuiInputEvent_GetAddedByTestEngine(ImGuiInputEvent *self) { return self->AddedByTestEngine; }
void ImGuiMenuColumns_SetTotalWidth(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU32 v) { ImGuiMenuColumnsPtr->TotalWidth = v; }
ImU32 ImGuiMenuColumns_GetTotalWidth(ImGuiMenuColumns *self) { return self->TotalWidth; }
void ImGuiMenuColumns_SetNextTotalWidth(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU32 v) { ImGuiMenuColumnsPtr->NextTotalWidth = v; }
ImU32 ImGuiMenuColumns_GetNextTotalWidth(ImGuiMenuColumns *self) { return self->NextTotalWidth; }
void ImGuiMenuColumns_SetSpacing(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU16 v) { ImGuiMenuColumnsPtr->Spacing = v; }
ImU16 ImGuiMenuColumns_GetSpacing(ImGuiMenuColumns *self) { return self->Spacing; }
void ImGuiMenuColumns_SetOffsetIcon(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU16 v) { ImGuiMenuColumnsPtr->OffsetIcon = v; }
ImU16 ImGuiMenuColumns_GetOffsetIcon(ImGuiMenuColumns *self) { return self->OffsetIcon; }
void ImGuiMenuColumns_SetOffsetLabel(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU16 v) { ImGuiMenuColumnsPtr->OffsetLabel = v; }
ImU16 ImGuiMenuColumns_GetOffsetLabel(ImGuiMenuColumns *self) { return self->OffsetLabel; }
void ImGuiMenuColumns_SetOffsetShortcut(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU16 v) { ImGuiMenuColumnsPtr->OffsetShortcut = v; }
ImU16 ImGuiMenuColumns_GetOffsetShortcut(ImGuiMenuColumns *self) { return self->OffsetShortcut; }
void ImGuiMenuColumns_SetOffsetMark(ImGuiMenuColumns *ImGuiMenuColumnsPtr, ImU16 v) { ImGuiMenuColumnsPtr->OffsetMark = v; }
ImU16 ImGuiMenuColumns_GetOffsetMark(ImGuiMenuColumns *self) { return self->OffsetMark; }
void ImGuiPayload_SetData(ImGuiPayload *ImGuiPayloadPtr, void* v) { ImGuiPayloadPtr->Data = v; }
void* ImGuiPayload_GetData(ImGuiPayload *self) { return self->Data; }
void ImGuiPayload_SetDataSize(ImGuiPayload *ImGuiPayloadPtr, int v) { ImGuiPayloadPtr->DataSize = v; }
int ImGuiPayload_GetDataSize(ImGuiPayload *self) { return self->DataSize; }
void ImGuiPayload_SetSourceId(ImGuiPayload *ImGuiPayloadPtr, ImGuiID v) { ImGuiPayloadPtr->SourceId = v; }
ImGuiID ImGuiPayload_GetSourceId(ImGuiPayload *self) { return self->SourceId; }
void ImGuiPayload_SetSourceParentId(ImGuiPayload *ImGuiPayloadPtr, ImGuiID v) { ImGuiPayloadPtr->SourceParentId = v; }
ImGuiID ImGuiPayload_GetSourceParentId(ImGuiPayload *self) { return self->SourceParentId; }
void ImGuiPayload_SetDataFrameCount(ImGuiPayload *ImGuiPayloadPtr, int v) { ImGuiPayloadPtr->DataFrameCount = v; }
int ImGuiPayload_GetDataFrameCount(ImGuiPayload *self) { return self->DataFrameCount; }
void ImGuiPayload_SetPreview(ImGuiPayload *ImGuiPayloadPtr, bool v) { ImGuiPayloadPtr->Preview = v; }
bool ImGuiPayload_GetPreview(ImGuiPayload *self) { return self->Preview; }
void ImGuiPayload_SetDelivery(ImGuiPayload *ImGuiPayloadPtr, bool v) { ImGuiPayloadPtr->Delivery = v; }
bool ImGuiPayload_GetDelivery(ImGuiPayload *self) { return self->Delivery; }
void ImGuiTableColumnSettings_SetWidthOrWeight(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, float v) { ImGuiTableColumnSettingsPtr->WidthOrWeight = v; }
float ImGuiTableColumnSettings_GetWidthOrWeight(ImGuiTableColumnSettings *self) { return self->WidthOrWeight; }
void ImGuiTableColumnSettings_SetUserID(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImGuiID v) { ImGuiTableColumnSettingsPtr->UserID = v; }
ImGuiID ImGuiTableColumnSettings_GetUserID(ImGuiTableColumnSettings *self) { return self->UserID; }
void ImGuiTableColumnSettings_SetIndex(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnSettingsPtr->Index = v; }
ImGuiTableColumnIdx ImGuiTableColumnSettings_GetIndex(ImGuiTableColumnSettings *self) { return self->Index; }
void ImGuiTableColumnSettings_SetDisplayOrder(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnSettingsPtr->DisplayOrder = v; }
ImGuiTableColumnIdx ImGuiTableColumnSettings_GetDisplayOrder(ImGuiTableColumnSettings *self) { return self->DisplayOrder; }
void ImGuiTableColumnSettings_SetSortOrder(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImGuiTableColumnIdx v) { ImGuiTableColumnSettingsPtr->SortOrder = v; }
ImGuiTableColumnIdx ImGuiTableColumnSettings_GetSortOrder(ImGuiTableColumnSettings *self) { return self->SortOrder; }
void ImGuiTableColumnSettings_SetSortDirection(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImU8 v) { ImGuiTableColumnSettingsPtr->SortDirection = v; }
ImU8 ImGuiTableColumnSettings_GetSortDirection(ImGuiTableColumnSettings *self) { return self->SortDirection; }
void ImGuiTableColumnSettings_SetIsEnabled(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImU8 v) { ImGuiTableColumnSettingsPtr->IsEnabled = v; }
ImU8 ImGuiTableColumnSettings_GetIsEnabled(ImGuiTableColumnSettings *self) { return self->IsEnabled; }
void ImGuiTableColumnSettings_SetIsStretch(ImGuiTableColumnSettings *ImGuiTableColumnSettingsPtr, ImU8 v) { ImGuiTableColumnSettingsPtr->IsStretch = v; }
ImU8 ImGuiTableColumnSettings_GetIsStretch(ImGuiTableColumnSettings *self) { return self->IsStretch; }
void ImDrawVert_Setpos(ImDrawVert *ImDrawVertPtr, ImVec2 v) { ImDrawVertPtr->pos = v; }
ImVec2 ImDrawVert_Getpos(ImDrawVert *self) { return self->pos; }
void ImDrawVert_Setuv(ImDrawVert *ImDrawVertPtr, ImVec2 v) { ImDrawVertPtr->uv = v; }
ImVec2 ImDrawVert_Getuv(ImDrawVert *self) { return self->uv; }
void ImDrawVert_Setcol(ImDrawVert *ImDrawVertPtr, ImU32 v) { ImDrawVertPtr->col = v; }
ImU32 ImDrawVert_Getcol(ImDrawVert *self) { return self->col; }
void ImGuiContext_SetInitialized(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->Initialized = v; }
bool ImGuiContext_GetInitialized(ImGuiContext *self) { return self->Initialized; }
void ImGuiContext_SetFontAtlasOwnedByContext(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->FontAtlasOwnedByContext = v; }
bool ImGuiContext_GetFontAtlasOwnedByContext(ImGuiContext *self) { return self->FontAtlasOwnedByContext; }
void ImGuiContext_SetIO(ImGuiContext *ImGuiContextPtr, ImGuiIO v) { ImGuiContextPtr->IO = v; }
ImGuiIO ImGuiContext_GetIO(ImGuiContext *self) { return self->IO; }
void ImGuiContext_SetPlatformIO(ImGuiContext *ImGuiContextPtr, ImGuiPlatformIO v) { ImGuiContextPtr->PlatformIO = v; }
ImGuiPlatformIO ImGuiContext_GetPlatformIO(ImGuiContext *self) { return self->PlatformIO; }
void ImGuiContext_SetInputEventsQueue(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiInputEvent v) { ImGuiContextPtr->InputEventsQueue = v; }
ImVector_ImGuiInputEvent ImGuiContext_GetInputEventsQueue(ImGuiContext *self) { return self->InputEventsQueue; }
void ImGuiContext_SetInputEventsTrail(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiInputEvent v) { ImGuiContextPtr->InputEventsTrail = v; }
ImVector_ImGuiInputEvent ImGuiContext_GetInputEventsTrail(ImGuiContext *self) { return self->InputEventsTrail; }
void ImGuiContext_SetStyle(ImGuiContext *ImGuiContextPtr, ImGuiStyle v) { ImGuiContextPtr->Style = v; }
ImGuiStyle ImGuiContext_GetStyle(ImGuiContext *self) { return self->Style; }
void ImGuiContext_SetConfigFlagsCurrFrame(ImGuiContext *ImGuiContextPtr, ImGuiConfigFlags v) { ImGuiContextPtr->ConfigFlagsCurrFrame = v; }
ImGuiConfigFlags ImGuiContext_GetConfigFlagsCurrFrame(ImGuiContext *self) { return self->ConfigFlagsCurrFrame; }
void ImGuiContext_SetConfigFlagsLastFrame(ImGuiContext *ImGuiContextPtr, ImGuiConfigFlags v) { ImGuiContextPtr->ConfigFlagsLastFrame = v; }
ImGuiConfigFlags ImGuiContext_GetConfigFlagsLastFrame(ImGuiContext *self) { return self->ConfigFlagsLastFrame; }
void ImGuiContext_SetFont(ImGuiContext *ImGuiContextPtr, ImFont* v) { ImGuiContextPtr->Font = v; }
ImFont* ImGuiContext_GetFont(ImGuiContext *self) { return self->Font; }
void ImGuiContext_SetFontSize(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->FontSize = v; }
float ImGuiContext_GetFontSize(ImGuiContext *self) { return self->FontSize; }
void ImGuiContext_SetFontBaseSize(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->FontBaseSize = v; }
float ImGuiContext_GetFontBaseSize(ImGuiContext *self) { return self->FontBaseSize; }
void ImGuiContext_SetDrawListSharedData(ImGuiContext *ImGuiContextPtr, ImDrawListSharedData v) { ImGuiContextPtr->DrawListSharedData = v; }
ImDrawListSharedData ImGuiContext_GetDrawListSharedData(ImGuiContext *self) { return self->DrawListSharedData; }
void ImGuiContext_SetTime(ImGuiContext *ImGuiContextPtr, double v) { ImGuiContextPtr->Time = v; }
double ImGuiContext_GetTime(ImGuiContext *self) { return self->Time; }
void ImGuiContext_SetFrameCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->FrameCount = v; }
int ImGuiContext_GetFrameCount(ImGuiContext *self) { return self->FrameCount; }
void ImGuiContext_SetFrameCountEnded(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->FrameCountEnded = v; }
int ImGuiContext_GetFrameCountEnded(ImGuiContext *self) { return self->FrameCountEnded; }
void ImGuiContext_SetFrameCountPlatformEnded(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->FrameCountPlatformEnded = v; }
int ImGuiContext_GetFrameCountPlatformEnded(ImGuiContext *self) { return self->FrameCountPlatformEnded; }
void ImGuiContext_SetFrameCountRendered(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->FrameCountRendered = v; }
int ImGuiContext_GetFrameCountRendered(ImGuiContext *self) { return self->FrameCountRendered; }
void ImGuiContext_SetWithinFrameScope(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->WithinFrameScope = v; }
bool ImGuiContext_GetWithinFrameScope(ImGuiContext *self) { return self->WithinFrameScope; }
void ImGuiContext_SetWithinFrameScopeWithImplicitWindow(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->WithinFrameScopeWithImplicitWindow = v; }
bool ImGuiContext_GetWithinFrameScopeWithImplicitWindow(ImGuiContext *self) { return self->WithinFrameScopeWithImplicitWindow; }
void ImGuiContext_SetWithinEndChild(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->WithinEndChild = v; }
bool ImGuiContext_GetWithinEndChild(ImGuiContext *self) { return self->WithinEndChild; }
void ImGuiContext_SetGcCompactAll(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->GcCompactAll = v; }
bool ImGuiContext_GetGcCompactAll(ImGuiContext *self) { return self->GcCompactAll; }
void ImGuiContext_SetTestEngineHookItems(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->TestEngineHookItems = v; }
bool ImGuiContext_GetTestEngineHookItems(ImGuiContext *self) { return self->TestEngineHookItems; }
void ImGuiContext_SetTestEngine(ImGuiContext *ImGuiContextPtr, void* v) { ImGuiContextPtr->TestEngine = v; }
void* ImGuiContext_GetTestEngine(ImGuiContext *self) { return self->TestEngine; }
void ImGuiContext_SetWindows(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiWindowPtr v) { ImGuiContextPtr->Windows = v; }
ImVector_ImGuiWindowPtr ImGuiContext_GetWindows(ImGuiContext *self) { return self->Windows; }
void ImGuiContext_SetWindowsFocusOrder(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiWindowPtr v) { ImGuiContextPtr->WindowsFocusOrder = v; }
ImVector_ImGuiWindowPtr ImGuiContext_GetWindowsFocusOrder(ImGuiContext *self) { return self->WindowsFocusOrder; }
void ImGuiContext_SetWindowsTempSortBuffer(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiWindowPtr v) { ImGuiContextPtr->WindowsTempSortBuffer = v; }
ImVector_ImGuiWindowPtr ImGuiContext_GetWindowsTempSortBuffer(ImGuiContext *self) { return self->WindowsTempSortBuffer; }
void ImGuiContext_SetCurrentWindowStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiWindowStackData v) { ImGuiContextPtr->CurrentWindowStack = v; }
ImVector_ImGuiWindowStackData ImGuiContext_GetCurrentWindowStack(ImGuiContext *self) { return self->CurrentWindowStack; }
void ImGuiContext_SetWindowsById(ImGuiContext *ImGuiContextPtr, ImGuiStorage v) { ImGuiContextPtr->WindowsById = v; }
ImGuiStorage ImGuiContext_GetWindowsById(ImGuiContext *self) { return self->WindowsById; }
void ImGuiContext_SetWindowsActiveCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->WindowsActiveCount = v; }
int ImGuiContext_GetWindowsActiveCount(ImGuiContext *self) { return self->WindowsActiveCount; }
void ImGuiContext_SetWindowsHoverPadding(ImGuiContext *ImGuiContextPtr, ImVec2 v) { ImGuiContextPtr->WindowsHoverPadding = v; }
ImVec2 ImGuiContext_GetWindowsHoverPadding(ImGuiContext *self) { return self->WindowsHoverPadding; }
void ImGuiContext_SetCurrentWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->CurrentWindow = v; }
ImGuiWindow* ImGuiContext_GetCurrentWindow(ImGuiContext *self) { return self->CurrentWindow; }
void ImGuiContext_SetHoveredWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->HoveredWindow = v; }
ImGuiWindow* ImGuiContext_GetHoveredWindow(ImGuiContext *self) { return self->HoveredWindow; }
void ImGuiContext_SetHoveredWindowUnderMovingWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->HoveredWindowUnderMovingWindow = v; }
ImGuiWindow* ImGuiContext_GetHoveredWindowUnderMovingWindow(ImGuiContext *self) { return self->HoveredWindowUnderMovingWindow; }
void ImGuiContext_SetHoveredDockNode(ImGuiContext *ImGuiContextPtr, ImGuiDockNode* v) { ImGuiContextPtr->HoveredDockNode = v; }
ImGuiDockNode* ImGuiContext_GetHoveredDockNode(ImGuiContext *self) { return self->HoveredDockNode; }
void ImGuiContext_SetMovingWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->MovingWindow = v; }
ImGuiWindow* ImGuiContext_GetMovingWindow(ImGuiContext *self) { return self->MovingWindow; }
void ImGuiContext_SetWheelingWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->WheelingWindow = v; }
ImGuiWindow* ImGuiContext_GetWheelingWindow(ImGuiContext *self) { return self->WheelingWindow; }
void ImGuiContext_SetWheelingWindowRefMousePos(ImGuiContext *ImGuiContextPtr, ImVec2 v) { ImGuiContextPtr->WheelingWindowRefMousePos = v; }
ImVec2 ImGuiContext_GetWheelingWindowRefMousePos(ImGuiContext *self) { return self->WheelingWindowRefMousePos; }
void ImGuiContext_SetWheelingWindowTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->WheelingWindowTimer = v; }
float ImGuiContext_GetWheelingWindowTimer(ImGuiContext *self) { return self->WheelingWindowTimer; }
void ImGuiContext_SetDebugHookIdInfo(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->DebugHookIdInfo = v; }
ImGuiID ImGuiContext_GetDebugHookIdInfo(ImGuiContext *self) { return self->DebugHookIdInfo; }
void ImGuiContext_SetHoveredId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->HoveredId = v; }
ImGuiID ImGuiContext_GetHoveredId(ImGuiContext *self) { return self->HoveredId; }
void ImGuiContext_SetHoveredIdPreviousFrame(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->HoveredIdPreviousFrame = v; }
ImGuiID ImGuiContext_GetHoveredIdPreviousFrame(ImGuiContext *self) { return self->HoveredIdPreviousFrame; }
void ImGuiContext_SetHoveredIdAllowOverlap(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->HoveredIdAllowOverlap = v; }
bool ImGuiContext_GetHoveredIdAllowOverlap(ImGuiContext *self) { return self->HoveredIdAllowOverlap; }
void ImGuiContext_SetHoveredIdUsingMouseWheel(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->HoveredIdUsingMouseWheel = v; }
bool ImGuiContext_GetHoveredIdUsingMouseWheel(ImGuiContext *self) { return self->HoveredIdUsingMouseWheel; }
void ImGuiContext_SetHoveredIdPreviousFrameUsingMouseWheel(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->HoveredIdPreviousFrameUsingMouseWheel = v; }
bool ImGuiContext_GetHoveredIdPreviousFrameUsingMouseWheel(ImGuiContext *self) { return self->HoveredIdPreviousFrameUsingMouseWheel; }
void ImGuiContext_SetHoveredIdDisabled(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->HoveredIdDisabled = v; }
bool ImGuiContext_GetHoveredIdDisabled(ImGuiContext *self) { return self->HoveredIdDisabled; }
void ImGuiContext_SetHoveredIdTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->HoveredIdTimer = v; }
float ImGuiContext_GetHoveredIdTimer(ImGuiContext *self) { return self->HoveredIdTimer; }
void ImGuiContext_SetHoveredIdNotActiveTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->HoveredIdNotActiveTimer = v; }
float ImGuiContext_GetHoveredIdNotActiveTimer(ImGuiContext *self) { return self->HoveredIdNotActiveTimer; }
void ImGuiContext_SetActiveId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->ActiveId = v; }
ImGuiID ImGuiContext_GetActiveId(ImGuiContext *self) { return self->ActiveId; }
void ImGuiContext_SetActiveIdIsAlive(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->ActiveIdIsAlive = v; }
ImGuiID ImGuiContext_GetActiveIdIsAlive(ImGuiContext *self) { return self->ActiveIdIsAlive; }
void ImGuiContext_SetActiveIdTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->ActiveIdTimer = v; }
float ImGuiContext_GetActiveIdTimer(ImGuiContext *self) { return self->ActiveIdTimer; }
void ImGuiContext_SetActiveIdIsJustActivated(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdIsJustActivated = v; }
bool ImGuiContext_GetActiveIdIsJustActivated(ImGuiContext *self) { return self->ActiveIdIsJustActivated; }
void ImGuiContext_SetActiveIdAllowOverlap(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdAllowOverlap = v; }
bool ImGuiContext_GetActiveIdAllowOverlap(ImGuiContext *self) { return self->ActiveIdAllowOverlap; }
void ImGuiContext_SetActiveIdNoClearOnFocusLoss(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdNoClearOnFocusLoss = v; }
bool ImGuiContext_GetActiveIdNoClearOnFocusLoss(ImGuiContext *self) { return self->ActiveIdNoClearOnFocusLoss; }
void ImGuiContext_SetActiveIdHasBeenPressedBefore(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdHasBeenPressedBefore = v; }
bool ImGuiContext_GetActiveIdHasBeenPressedBefore(ImGuiContext *self) { return self->ActiveIdHasBeenPressedBefore; }
void ImGuiContext_SetActiveIdHasBeenEditedBefore(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdHasBeenEditedBefore = v; }
bool ImGuiContext_GetActiveIdHasBeenEditedBefore(ImGuiContext *self) { return self->ActiveIdHasBeenEditedBefore; }
void ImGuiContext_SetActiveIdHasBeenEditedThisFrame(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdHasBeenEditedThisFrame = v; }
bool ImGuiContext_GetActiveIdHasBeenEditedThisFrame(ImGuiContext *self) { return self->ActiveIdHasBeenEditedThisFrame; }
void ImGuiContext_SetActiveIdClickOffset(ImGuiContext *ImGuiContextPtr, ImVec2 v) { ImGuiContextPtr->ActiveIdClickOffset = v; }
ImVec2 ImGuiContext_GetActiveIdClickOffset(ImGuiContext *self) { return self->ActiveIdClickOffset; }
void ImGuiContext_SetActiveIdWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->ActiveIdWindow = v; }
ImGuiWindow* ImGuiContext_GetActiveIdWindow(ImGuiContext *self) { return self->ActiveIdWindow; }
void ImGuiContext_SetActiveIdSource(ImGuiContext *ImGuiContextPtr, ImGuiInputSource v) { ImGuiContextPtr->ActiveIdSource = v; }
ImGuiInputSource ImGuiContext_GetActiveIdSource(ImGuiContext *self) { return self->ActiveIdSource; }
void ImGuiContext_SetActiveIdMouseButton(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->ActiveIdMouseButton = v; }
int ImGuiContext_GetActiveIdMouseButton(ImGuiContext *self) { return self->ActiveIdMouseButton; }
void ImGuiContext_SetActiveIdPreviousFrame(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->ActiveIdPreviousFrame = v; }
ImGuiID ImGuiContext_GetActiveIdPreviousFrame(ImGuiContext *self) { return self->ActiveIdPreviousFrame; }
void ImGuiContext_SetActiveIdPreviousFrameIsAlive(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdPreviousFrameIsAlive = v; }
bool ImGuiContext_GetActiveIdPreviousFrameIsAlive(ImGuiContext *self) { return self->ActiveIdPreviousFrameIsAlive; }
void ImGuiContext_SetActiveIdPreviousFrameHasBeenEditedBefore(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->ActiveIdPreviousFrameHasBeenEditedBefore = v; }
bool ImGuiContext_GetActiveIdPreviousFrameHasBeenEditedBefore(ImGuiContext *self) { return self->ActiveIdPreviousFrameHasBeenEditedBefore; }
void ImGuiContext_SetActiveIdPreviousFrameWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->ActiveIdPreviousFrameWindow = v; }
ImGuiWindow* ImGuiContext_GetActiveIdPreviousFrameWindow(ImGuiContext *self) { return self->ActiveIdPreviousFrameWindow; }
void ImGuiContext_SetLastActiveId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->LastActiveId = v; }
ImGuiID ImGuiContext_GetLastActiveId(ImGuiContext *self) { return self->LastActiveId; }
void ImGuiContext_SetLastActiveIdTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->LastActiveIdTimer = v; }
float ImGuiContext_GetLastActiveIdTimer(ImGuiContext *self) { return self->LastActiveIdTimer; }
void ImGuiContext_SetActiveIdUsingNavDirMask(ImGuiContext *ImGuiContextPtr, ImU32 v) { ImGuiContextPtr->ActiveIdUsingNavDirMask = v; }
ImU32 ImGuiContext_GetActiveIdUsingNavDirMask(ImGuiContext *self) { return self->ActiveIdUsingNavDirMask; }
void ImGuiContext_SetActiveIdUsingKeyInputMask(ImGuiContext *ImGuiContextPtr, ImBitArrayForNamedKeys v) { ImGuiContextPtr->ActiveIdUsingKeyInputMask = v; }
ImBitArrayForNamedKeys ImGuiContext_GetActiveIdUsingKeyInputMask(ImGuiContext *self) { return self->ActiveIdUsingKeyInputMask; }
void ImGuiContext_SetActiveIdUsingNavInputMask(ImGuiContext *ImGuiContextPtr, ImU32 v) { ImGuiContextPtr->ActiveIdUsingNavInputMask = v; }
ImU32 ImGuiContext_GetActiveIdUsingNavInputMask(ImGuiContext *self) { return self->ActiveIdUsingNavInputMask; }
void ImGuiContext_SetCurrentItemFlags(ImGuiContext *ImGuiContextPtr, ImGuiItemFlags v) { ImGuiContextPtr->CurrentItemFlags = v; }
ImGuiItemFlags ImGuiContext_GetCurrentItemFlags(ImGuiContext *self) { return self->CurrentItemFlags; }
void ImGuiContext_SetNextItemData(ImGuiContext *ImGuiContextPtr, ImGuiNextItemData v) { ImGuiContextPtr->NextItemData = v; }
ImGuiNextItemData ImGuiContext_GetNextItemData(ImGuiContext *self) { return self->NextItemData; }
void ImGuiContext_SetLastItemData(ImGuiContext *ImGuiContextPtr, ImGuiLastItemData v) { ImGuiContextPtr->LastItemData = v; }
ImGuiLastItemData ImGuiContext_GetLastItemData(ImGuiContext *self) { return self->LastItemData; }
void ImGuiContext_SetNextWindowData(ImGuiContext *ImGuiContextPtr, ImGuiNextWindowData v) { ImGuiContextPtr->NextWindowData = v; }
ImGuiNextWindowData ImGuiContext_GetNextWindowData(ImGuiContext *self) { return self->NextWindowData; }
void ImGuiContext_SetColorStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiColorMod v) { ImGuiContextPtr->ColorStack = v; }
ImVector_ImGuiColorMod ImGuiContext_GetColorStack(ImGuiContext *self) { return self->ColorStack; }
void ImGuiContext_SetStyleVarStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiStyleMod v) { ImGuiContextPtr->StyleVarStack = v; }
ImVector_ImGuiStyleMod ImGuiContext_GetStyleVarStack(ImGuiContext *self) { return self->StyleVarStack; }
void ImGuiContext_SetFontStack(ImGuiContext *ImGuiContextPtr, ImVector_ImFontPtr v) { ImGuiContextPtr->FontStack = v; }
ImVector_ImFontPtr ImGuiContext_GetFontStack(ImGuiContext *self) { return self->FontStack; }
void ImGuiContext_SetFocusScopeStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiID v) { ImGuiContextPtr->FocusScopeStack = v; }
ImVector_ImGuiID ImGuiContext_GetFocusScopeStack(ImGuiContext *self) { return self->FocusScopeStack; }
void ImGuiContext_SetItemFlagsStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiItemFlags v) { ImGuiContextPtr->ItemFlagsStack = v; }
ImVector_ImGuiItemFlags ImGuiContext_GetItemFlagsStack(ImGuiContext *self) { return self->ItemFlagsStack; }
void ImGuiContext_SetGroupStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiGroupData v) { ImGuiContextPtr->GroupStack = v; }
ImVector_ImGuiGroupData ImGuiContext_GetGroupStack(ImGuiContext *self) { return self->GroupStack; }
void ImGuiContext_SetOpenPopupStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiPopupData v) { ImGuiContextPtr->OpenPopupStack = v; }
ImVector_ImGuiPopupData ImGuiContext_GetOpenPopupStack(ImGuiContext *self) { return self->OpenPopupStack; }
void ImGuiContext_SetBeginPopupStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiPopupData v) { ImGuiContextPtr->BeginPopupStack = v; }
ImVector_ImGuiPopupData ImGuiContext_GetBeginPopupStack(ImGuiContext *self) { return self->BeginPopupStack; }
void ImGuiContext_SetBeginMenuCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->BeginMenuCount = v; }
int ImGuiContext_GetBeginMenuCount(ImGuiContext *self) { return self->BeginMenuCount; }
void ImGuiContext_SetViewports(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiViewportPPtr v) { ImGuiContextPtr->Viewports = v; }
ImVector_ImGuiViewportPPtr ImGuiContext_GetViewports(ImGuiContext *self) { return self->Viewports; }
void ImGuiContext_SetCurrentDpiScale(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->CurrentDpiScale = v; }
float ImGuiContext_GetCurrentDpiScale(ImGuiContext *self) { return self->CurrentDpiScale; }
void ImGuiContext_SetCurrentViewport(ImGuiContext *ImGuiContextPtr, ImGuiViewportP* v) { ImGuiContextPtr->CurrentViewport = v; }
ImGuiViewportP* ImGuiContext_GetCurrentViewport(ImGuiContext *self) { return self->CurrentViewport; }
void ImGuiContext_SetMouseViewport(ImGuiContext *ImGuiContextPtr, ImGuiViewportP* v) { ImGuiContextPtr->MouseViewport = v; }
ImGuiViewportP* ImGuiContext_GetMouseViewport(ImGuiContext *self) { return self->MouseViewport; }
void ImGuiContext_SetMouseLastHoveredViewport(ImGuiContext *ImGuiContextPtr, ImGuiViewportP* v) { ImGuiContextPtr->MouseLastHoveredViewport = v; }
ImGuiViewportP* ImGuiContext_GetMouseLastHoveredViewport(ImGuiContext *self) { return self->MouseLastHoveredViewport; }
void ImGuiContext_SetPlatformLastFocusedViewportId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->PlatformLastFocusedViewportId = v; }
ImGuiID ImGuiContext_GetPlatformLastFocusedViewportId(ImGuiContext *self) { return self->PlatformLastFocusedViewportId; }
void ImGuiContext_SetFallbackMonitor(ImGuiContext *ImGuiContextPtr, ImGuiPlatformMonitor v) { ImGuiContextPtr->FallbackMonitor = v; }
ImGuiPlatformMonitor ImGuiContext_GetFallbackMonitor(ImGuiContext *self) { return self->FallbackMonitor; }
void ImGuiContext_SetViewportFrontMostStampCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->ViewportFrontMostStampCount = v; }
int ImGuiContext_GetViewportFrontMostStampCount(ImGuiContext *self) { return self->ViewportFrontMostStampCount; }
void ImGuiContext_SetNavWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->NavWindow = v; }
ImGuiWindow* ImGuiContext_GetNavWindow(ImGuiContext *self) { return self->NavWindow; }
void ImGuiContext_SetNavId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavId = v; }
ImGuiID ImGuiContext_GetNavId(ImGuiContext *self) { return self->NavId; }
void ImGuiContext_SetNavFocusScopeId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavFocusScopeId = v; }
ImGuiID ImGuiContext_GetNavFocusScopeId(ImGuiContext *self) { return self->NavFocusScopeId; }
void ImGuiContext_SetNavActivateId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavActivateId = v; }
ImGuiID ImGuiContext_GetNavActivateId(ImGuiContext *self) { return self->NavActivateId; }
void ImGuiContext_SetNavActivateDownId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavActivateDownId = v; }
ImGuiID ImGuiContext_GetNavActivateDownId(ImGuiContext *self) { return self->NavActivateDownId; }
void ImGuiContext_SetNavActivatePressedId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavActivatePressedId = v; }
ImGuiID ImGuiContext_GetNavActivatePressedId(ImGuiContext *self) { return self->NavActivatePressedId; }
void ImGuiContext_SetNavActivateInputId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavActivateInputId = v; }
ImGuiID ImGuiContext_GetNavActivateInputId(ImGuiContext *self) { return self->NavActivateInputId; }
void ImGuiContext_SetNavActivateFlags(ImGuiContext *ImGuiContextPtr, ImGuiActivateFlags v) { ImGuiContextPtr->NavActivateFlags = v; }
ImGuiActivateFlags ImGuiContext_GetNavActivateFlags(ImGuiContext *self) { return self->NavActivateFlags; }
void ImGuiContext_SetNavJustMovedToId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavJustMovedToId = v; }
ImGuiID ImGuiContext_GetNavJustMovedToId(ImGuiContext *self) { return self->NavJustMovedToId; }
void ImGuiContext_SetNavJustMovedToFocusScopeId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavJustMovedToFocusScopeId = v; }
ImGuiID ImGuiContext_GetNavJustMovedToFocusScopeId(ImGuiContext *self) { return self->NavJustMovedToFocusScopeId; }
void ImGuiContext_SetNavJustMovedToKeyMods(ImGuiContext *ImGuiContextPtr, ImGuiModFlags v) { ImGuiContextPtr->NavJustMovedToKeyMods = v; }
ImGuiModFlags ImGuiContext_GetNavJustMovedToKeyMods(ImGuiContext *self) { return self->NavJustMovedToKeyMods; }
void ImGuiContext_SetNavNextActivateId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavNextActivateId = v; }
ImGuiID ImGuiContext_GetNavNextActivateId(ImGuiContext *self) { return self->NavNextActivateId; }
void ImGuiContext_SetNavNextActivateFlags(ImGuiContext *ImGuiContextPtr, ImGuiActivateFlags v) { ImGuiContextPtr->NavNextActivateFlags = v; }
ImGuiActivateFlags ImGuiContext_GetNavNextActivateFlags(ImGuiContext *self) { return self->NavNextActivateFlags; }
void ImGuiContext_SetNavInputSource(ImGuiContext *ImGuiContextPtr, ImGuiInputSource v) { ImGuiContextPtr->NavInputSource = v; }
ImGuiInputSource ImGuiContext_GetNavInputSource(ImGuiContext *self) { return self->NavInputSource; }
void ImGuiContext_SetNavLayer(ImGuiContext *ImGuiContextPtr, ImGuiNavLayer v) { ImGuiContextPtr->NavLayer = v; }
ImGuiNavLayer ImGuiContext_GetNavLayer(ImGuiContext *self) { return self->NavLayer; }
void ImGuiContext_SetNavIdIsAlive(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavIdIsAlive = v; }
bool ImGuiContext_GetNavIdIsAlive(ImGuiContext *self) { return self->NavIdIsAlive; }
void ImGuiContext_SetNavMousePosDirty(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavMousePosDirty = v; }
bool ImGuiContext_GetNavMousePosDirty(ImGuiContext *self) { return self->NavMousePosDirty; }
void ImGuiContext_SetNavDisableHighlight(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavDisableHighlight = v; }
bool ImGuiContext_GetNavDisableHighlight(ImGuiContext *self) { return self->NavDisableHighlight; }
void ImGuiContext_SetNavDisableMouseHover(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavDisableMouseHover = v; }
bool ImGuiContext_GetNavDisableMouseHover(ImGuiContext *self) { return self->NavDisableMouseHover; }
void ImGuiContext_SetNavAnyRequest(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavAnyRequest = v; }
bool ImGuiContext_GetNavAnyRequest(ImGuiContext *self) { return self->NavAnyRequest; }
void ImGuiContext_SetNavInitRequest(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavInitRequest = v; }
bool ImGuiContext_GetNavInitRequest(ImGuiContext *self) { return self->NavInitRequest; }
void ImGuiContext_SetNavInitRequestFromMove(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavInitRequestFromMove = v; }
bool ImGuiContext_GetNavInitRequestFromMove(ImGuiContext *self) { return self->NavInitRequestFromMove; }
void ImGuiContext_SetNavInitResultId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->NavInitResultId = v; }
ImGuiID ImGuiContext_GetNavInitResultId(ImGuiContext *self) { return self->NavInitResultId; }
void ImGuiContext_SetNavInitResultRectRel(ImGuiContext *ImGuiContextPtr, ImRect v) { ImGuiContextPtr->NavInitResultRectRel = v; }
ImRect ImGuiContext_GetNavInitResultRectRel(ImGuiContext *self) { return self->NavInitResultRectRel; }
void ImGuiContext_SetNavMoveSubmitted(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavMoveSubmitted = v; }
bool ImGuiContext_GetNavMoveSubmitted(ImGuiContext *self) { return self->NavMoveSubmitted; }
void ImGuiContext_SetNavMoveScoringItems(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavMoveScoringItems = v; }
bool ImGuiContext_GetNavMoveScoringItems(ImGuiContext *self) { return self->NavMoveScoringItems; }
void ImGuiContext_SetNavMoveForwardToNextFrame(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavMoveForwardToNextFrame = v; }
bool ImGuiContext_GetNavMoveForwardToNextFrame(ImGuiContext *self) { return self->NavMoveForwardToNextFrame; }
void ImGuiContext_SetNavMoveFlags(ImGuiContext *ImGuiContextPtr, ImGuiNavMoveFlags v) { ImGuiContextPtr->NavMoveFlags = v; }
ImGuiNavMoveFlags ImGuiContext_GetNavMoveFlags(ImGuiContext *self) { return self->NavMoveFlags; }
void ImGuiContext_SetNavMoveScrollFlags(ImGuiContext *ImGuiContextPtr, ImGuiScrollFlags v) { ImGuiContextPtr->NavMoveScrollFlags = v; }
ImGuiScrollFlags ImGuiContext_GetNavMoveScrollFlags(ImGuiContext *self) { return self->NavMoveScrollFlags; }
void ImGuiContext_SetNavMoveKeyMods(ImGuiContext *ImGuiContextPtr, ImGuiModFlags v) { ImGuiContextPtr->NavMoveKeyMods = v; }
ImGuiModFlags ImGuiContext_GetNavMoveKeyMods(ImGuiContext *self) { return self->NavMoveKeyMods; }
void ImGuiContext_SetNavMoveDir(ImGuiContext *ImGuiContextPtr, ImGuiDir v) { ImGuiContextPtr->NavMoveDir = v; }
ImGuiDir ImGuiContext_GetNavMoveDir(ImGuiContext *self) { return self->NavMoveDir; }
void ImGuiContext_SetNavMoveDirForDebug(ImGuiContext *ImGuiContextPtr, ImGuiDir v) { ImGuiContextPtr->NavMoveDirForDebug = v; }
ImGuiDir ImGuiContext_GetNavMoveDirForDebug(ImGuiContext *self) { return self->NavMoveDirForDebug; }
void ImGuiContext_SetNavMoveClipDir(ImGuiContext *ImGuiContextPtr, ImGuiDir v) { ImGuiContextPtr->NavMoveClipDir = v; }
ImGuiDir ImGuiContext_GetNavMoveClipDir(ImGuiContext *self) { return self->NavMoveClipDir; }
void ImGuiContext_SetNavScoringRect(ImGuiContext *ImGuiContextPtr, ImRect v) { ImGuiContextPtr->NavScoringRect = v; }
ImRect ImGuiContext_GetNavScoringRect(ImGuiContext *self) { return self->NavScoringRect; }
void ImGuiContext_SetNavScoringNoClipRect(ImGuiContext *ImGuiContextPtr, ImRect v) { ImGuiContextPtr->NavScoringNoClipRect = v; }
ImRect ImGuiContext_GetNavScoringNoClipRect(ImGuiContext *self) { return self->NavScoringNoClipRect; }
void ImGuiContext_SetNavScoringDebugCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->NavScoringDebugCount = v; }
int ImGuiContext_GetNavScoringDebugCount(ImGuiContext *self) { return self->NavScoringDebugCount; }
void ImGuiContext_SetNavTabbingDir(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->NavTabbingDir = v; }
int ImGuiContext_GetNavTabbingDir(ImGuiContext *self) { return self->NavTabbingDir; }
void ImGuiContext_SetNavTabbingCounter(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->NavTabbingCounter = v; }
int ImGuiContext_GetNavTabbingCounter(ImGuiContext *self) { return self->NavTabbingCounter; }
void ImGuiContext_SetNavMoveResultLocal(ImGuiContext *ImGuiContextPtr, ImGuiNavItemData v) { ImGuiContextPtr->NavMoveResultLocal = v; }
ImGuiNavItemData ImGuiContext_GetNavMoveResultLocal(ImGuiContext *self) { return self->NavMoveResultLocal; }
void ImGuiContext_SetNavMoveResultLocalVisible(ImGuiContext *ImGuiContextPtr, ImGuiNavItemData v) { ImGuiContextPtr->NavMoveResultLocalVisible = v; }
ImGuiNavItemData ImGuiContext_GetNavMoveResultLocalVisible(ImGuiContext *self) { return self->NavMoveResultLocalVisible; }
void ImGuiContext_SetNavMoveResultOther(ImGuiContext *ImGuiContextPtr, ImGuiNavItemData v) { ImGuiContextPtr->NavMoveResultOther = v; }
ImGuiNavItemData ImGuiContext_GetNavMoveResultOther(ImGuiContext *self) { return self->NavMoveResultOther; }
void ImGuiContext_SetNavTabbingResultFirst(ImGuiContext *ImGuiContextPtr, ImGuiNavItemData v) { ImGuiContextPtr->NavTabbingResultFirst = v; }
ImGuiNavItemData ImGuiContext_GetNavTabbingResultFirst(ImGuiContext *self) { return self->NavTabbingResultFirst; }
void ImGuiContext_SetNavWindowingTarget(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->NavWindowingTarget = v; }
ImGuiWindow* ImGuiContext_GetNavWindowingTarget(ImGuiContext *self) { return self->NavWindowingTarget; }
void ImGuiContext_SetNavWindowingTargetAnim(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->NavWindowingTargetAnim = v; }
ImGuiWindow* ImGuiContext_GetNavWindowingTargetAnim(ImGuiContext *self) { return self->NavWindowingTargetAnim; }
void ImGuiContext_SetNavWindowingListWindow(ImGuiContext *ImGuiContextPtr, ImGuiWindow* v) { ImGuiContextPtr->NavWindowingListWindow = v; }
ImGuiWindow* ImGuiContext_GetNavWindowingListWindow(ImGuiContext *self) { return self->NavWindowingListWindow; }
void ImGuiContext_SetNavWindowingTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->NavWindowingTimer = v; }
float ImGuiContext_GetNavWindowingTimer(ImGuiContext *self) { return self->NavWindowingTimer; }
void ImGuiContext_SetNavWindowingHighlightAlpha(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->NavWindowingHighlightAlpha = v; }
float ImGuiContext_GetNavWindowingHighlightAlpha(ImGuiContext *self) { return self->NavWindowingHighlightAlpha; }
void ImGuiContext_SetNavWindowingToggleLayer(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->NavWindowingToggleLayer = v; }
bool ImGuiContext_GetNavWindowingToggleLayer(ImGuiContext *self) { return self->NavWindowingToggleLayer; }
void ImGuiContext_SetNavWindowingAccumDeltaPos(ImGuiContext *ImGuiContextPtr, ImVec2 v) { ImGuiContextPtr->NavWindowingAccumDeltaPos = v; }
ImVec2 ImGuiContext_GetNavWindowingAccumDeltaPos(ImGuiContext *self) { return self->NavWindowingAccumDeltaPos; }
void ImGuiContext_SetNavWindowingAccumDeltaSize(ImGuiContext *ImGuiContextPtr, ImVec2 v) { ImGuiContextPtr->NavWindowingAccumDeltaSize = v; }
ImVec2 ImGuiContext_GetNavWindowingAccumDeltaSize(ImGuiContext *self) { return self->NavWindowingAccumDeltaSize; }
void ImGuiContext_SetDimBgRatio(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->DimBgRatio = v; }
float ImGuiContext_GetDimBgRatio(ImGuiContext *self) { return self->DimBgRatio; }
void ImGuiContext_SetMouseCursor(ImGuiContext *ImGuiContextPtr, ImGuiMouseCursor v) { ImGuiContextPtr->MouseCursor = v; }
ImGuiMouseCursor ImGuiContext_GetMouseCursor(ImGuiContext *self) { return self->MouseCursor; }
void ImGuiContext_SetDragDropActive(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->DragDropActive = v; }
bool ImGuiContext_GetDragDropActive(ImGuiContext *self) { return self->DragDropActive; }
void ImGuiContext_SetDragDropWithinSource(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->DragDropWithinSource = v; }
bool ImGuiContext_GetDragDropWithinSource(ImGuiContext *self) { return self->DragDropWithinSource; }
void ImGuiContext_SetDragDropWithinTarget(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->DragDropWithinTarget = v; }
bool ImGuiContext_GetDragDropWithinTarget(ImGuiContext *self) { return self->DragDropWithinTarget; }
void ImGuiContext_SetDragDropSourceFlags(ImGuiContext *ImGuiContextPtr, ImGuiDragDropFlags v) { ImGuiContextPtr->DragDropSourceFlags = v; }
ImGuiDragDropFlags ImGuiContext_GetDragDropSourceFlags(ImGuiContext *self) { return self->DragDropSourceFlags; }
void ImGuiContext_SetDragDropSourceFrameCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->DragDropSourceFrameCount = v; }
int ImGuiContext_GetDragDropSourceFrameCount(ImGuiContext *self) { return self->DragDropSourceFrameCount; }
void ImGuiContext_SetDragDropMouseButton(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->DragDropMouseButton = v; }
int ImGuiContext_GetDragDropMouseButton(ImGuiContext *self) { return self->DragDropMouseButton; }
void ImGuiContext_SetDragDropPayload(ImGuiContext *ImGuiContextPtr, ImGuiPayload v) { ImGuiContextPtr->DragDropPayload = v; }
ImGuiPayload ImGuiContext_GetDragDropPayload(ImGuiContext *self) { return self->DragDropPayload; }
void ImGuiContext_SetDragDropTargetRect(ImGuiContext *ImGuiContextPtr, ImRect v) { ImGuiContextPtr->DragDropTargetRect = v; }
ImRect ImGuiContext_GetDragDropTargetRect(ImGuiContext *self) { return self->DragDropTargetRect; }
void ImGuiContext_SetDragDropTargetId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->DragDropTargetId = v; }
ImGuiID ImGuiContext_GetDragDropTargetId(ImGuiContext *self) { return self->DragDropTargetId; }
void ImGuiContext_SetDragDropAcceptFlags(ImGuiContext *ImGuiContextPtr, ImGuiDragDropFlags v) { ImGuiContextPtr->DragDropAcceptFlags = v; }
ImGuiDragDropFlags ImGuiContext_GetDragDropAcceptFlags(ImGuiContext *self) { return self->DragDropAcceptFlags; }
void ImGuiContext_SetDragDropAcceptIdCurrRectSurface(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->DragDropAcceptIdCurrRectSurface = v; }
float ImGuiContext_GetDragDropAcceptIdCurrRectSurface(ImGuiContext *self) { return self->DragDropAcceptIdCurrRectSurface; }
void ImGuiContext_SetDragDropAcceptIdCurr(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->DragDropAcceptIdCurr = v; }
ImGuiID ImGuiContext_GetDragDropAcceptIdCurr(ImGuiContext *self) { return self->DragDropAcceptIdCurr; }
void ImGuiContext_SetDragDropAcceptIdPrev(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->DragDropAcceptIdPrev = v; }
ImGuiID ImGuiContext_GetDragDropAcceptIdPrev(ImGuiContext *self) { return self->DragDropAcceptIdPrev; }
void ImGuiContext_SetDragDropAcceptFrameCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->DragDropAcceptFrameCount = v; }
int ImGuiContext_GetDragDropAcceptFrameCount(ImGuiContext *self) { return self->DragDropAcceptFrameCount; }
void ImGuiContext_SetDragDropHoldJustPressedId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->DragDropHoldJustPressedId = v; }
ImGuiID ImGuiContext_GetDragDropHoldJustPressedId(ImGuiContext *self) { return self->DragDropHoldJustPressedId; }
void ImGuiContext_SetDragDropPayloadBufHeap(ImGuiContext *ImGuiContextPtr, ImVector_unsigned_char v) { ImGuiContextPtr->DragDropPayloadBufHeap = v; }
ImVector_unsigned_char ImGuiContext_GetDragDropPayloadBufHeap(ImGuiContext *self) { return self->DragDropPayloadBufHeap; }
void ImGuiContext_SetClipperTempDataStacked(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->ClipperTempDataStacked = v; }
int ImGuiContext_GetClipperTempDataStacked(ImGuiContext *self) { return self->ClipperTempDataStacked; }
void ImGuiContext_SetClipperTempData(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiListClipperData v) { ImGuiContextPtr->ClipperTempData = v; }
ImVector_ImGuiListClipperData ImGuiContext_GetClipperTempData(ImGuiContext *self) { return self->ClipperTempData; }
void ImGuiContext_SetCurrentTable(ImGuiContext *ImGuiContextPtr, ImGuiTable* v) { ImGuiContextPtr->CurrentTable = v; }
ImGuiTable* ImGuiContext_GetCurrentTable(ImGuiContext *self) { return self->CurrentTable; }
void ImGuiContext_SetTablesTempDataStacked(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->TablesTempDataStacked = v; }
int ImGuiContext_GetTablesTempDataStacked(ImGuiContext *self) { return self->TablesTempDataStacked; }
void ImGuiContext_SetTablesTempData(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiTableTempData v) { ImGuiContextPtr->TablesTempData = v; }
ImVector_ImGuiTableTempData ImGuiContext_GetTablesTempData(ImGuiContext *self) { return self->TablesTempData; }
void ImGuiContext_SetTables(ImGuiContext *ImGuiContextPtr, ImPool_ImGuiTable v) { ImGuiContextPtr->Tables = v; }
ImPool_ImGuiTable ImGuiContext_GetTables(ImGuiContext *self) { return self->Tables; }
void ImGuiContext_SetTablesLastTimeActive(ImGuiContext *ImGuiContextPtr, ImVector_float v) { ImGuiContextPtr->TablesLastTimeActive = v; }
ImVector_float ImGuiContext_GetTablesLastTimeActive(ImGuiContext *self) { return self->TablesLastTimeActive; }
void ImGuiContext_SetDrawChannelsTempMergeBuffer(ImGuiContext *ImGuiContextPtr, ImVector_ImDrawChannel v) { ImGuiContextPtr->DrawChannelsTempMergeBuffer = v; }
ImVector_ImDrawChannel ImGuiContext_GetDrawChannelsTempMergeBuffer(ImGuiContext *self) { return self->DrawChannelsTempMergeBuffer; }
void ImGuiContext_SetCurrentTabBar(ImGuiContext *ImGuiContextPtr, ImGuiTabBar* v) { ImGuiContextPtr->CurrentTabBar = v; }
ImGuiTabBar* ImGuiContext_GetCurrentTabBar(ImGuiContext *self) { return self->CurrentTabBar; }
void ImGuiContext_SetTabBars(ImGuiContext *ImGuiContextPtr, ImPool_ImGuiTabBar v) { ImGuiContextPtr->TabBars = v; }
ImPool_ImGuiTabBar ImGuiContext_GetTabBars(ImGuiContext *self) { return self->TabBars; }
void ImGuiContext_SetCurrentTabBarStack(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiPtrOrIndex v) { ImGuiContextPtr->CurrentTabBarStack = v; }
ImVector_ImGuiPtrOrIndex ImGuiContext_GetCurrentTabBarStack(ImGuiContext *self) { return self->CurrentTabBarStack; }
void ImGuiContext_SetShrinkWidthBuffer(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiShrinkWidthItem v) { ImGuiContextPtr->ShrinkWidthBuffer = v; }
ImVector_ImGuiShrinkWidthItem ImGuiContext_GetShrinkWidthBuffer(ImGuiContext *self) { return self->ShrinkWidthBuffer; }
void ImGuiContext_SetMouseLastValidPos(ImGuiContext *ImGuiContextPtr, ImVec2 v) { ImGuiContextPtr->MouseLastValidPos = v; }
ImVec2 ImGuiContext_GetMouseLastValidPos(ImGuiContext *self) { return self->MouseLastValidPos; }
void ImGuiContext_SetInputTextState(ImGuiContext *ImGuiContextPtr, ImGuiInputTextState v) { ImGuiContextPtr->InputTextState = v; }
ImGuiInputTextState ImGuiContext_GetInputTextState(ImGuiContext *self) { return self->InputTextState; }
void ImGuiContext_SetInputTextPasswordFont(ImGuiContext *ImGuiContextPtr, ImFont v) { ImGuiContextPtr->InputTextPasswordFont = v; }
ImFont ImGuiContext_GetInputTextPasswordFont(ImGuiContext *self) { return self->InputTextPasswordFont; }
void ImGuiContext_SetTempInputId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->TempInputId = v; }
ImGuiID ImGuiContext_GetTempInputId(ImGuiContext *self) { return self->TempInputId; }
void ImGuiContext_SetColorEditOptions(ImGuiContext *ImGuiContextPtr, ImGuiColorEditFlags v) { ImGuiContextPtr->ColorEditOptions = v; }
ImGuiColorEditFlags ImGuiContext_GetColorEditOptions(ImGuiContext *self) { return self->ColorEditOptions; }
void ImGuiContext_SetColorEditLastHue(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->ColorEditLastHue = v; }
float ImGuiContext_GetColorEditLastHue(ImGuiContext *self) { return self->ColorEditLastHue; }
void ImGuiContext_SetColorEditLastSat(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->ColorEditLastSat = v; }
float ImGuiContext_GetColorEditLastSat(ImGuiContext *self) { return self->ColorEditLastSat; }
void ImGuiContext_SetColorEditLastColor(ImGuiContext *ImGuiContextPtr, ImU32 v) { ImGuiContextPtr->ColorEditLastColor = v; }
ImU32 ImGuiContext_GetColorEditLastColor(ImGuiContext *self) { return self->ColorEditLastColor; }
void ImGuiContext_SetColorPickerRef(ImGuiContext *ImGuiContextPtr, ImVec4 v) { ImGuiContextPtr->ColorPickerRef = v; }
ImVec4 ImGuiContext_GetColorPickerRef(ImGuiContext *self) { return self->ColorPickerRef; }
void ImGuiContext_SetComboPreviewData(ImGuiContext *ImGuiContextPtr, ImGuiComboPreviewData v) { ImGuiContextPtr->ComboPreviewData = v; }
ImGuiComboPreviewData ImGuiContext_GetComboPreviewData(ImGuiContext *self) { return self->ComboPreviewData; }
void ImGuiContext_SetSliderGrabClickOffset(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->SliderGrabClickOffset = v; }
float ImGuiContext_GetSliderGrabClickOffset(ImGuiContext *self) { return self->SliderGrabClickOffset; }
void ImGuiContext_SetSliderCurrentAccum(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->SliderCurrentAccum = v; }
float ImGuiContext_GetSliderCurrentAccum(ImGuiContext *self) { return self->SliderCurrentAccum; }
void ImGuiContext_SetSliderCurrentAccumDirty(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->SliderCurrentAccumDirty = v; }
bool ImGuiContext_GetSliderCurrentAccumDirty(ImGuiContext *self) { return self->SliderCurrentAccumDirty; }
void ImGuiContext_SetDragCurrentAccumDirty(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->DragCurrentAccumDirty = v; }
bool ImGuiContext_GetDragCurrentAccumDirty(ImGuiContext *self) { return self->DragCurrentAccumDirty; }
void ImGuiContext_SetDragCurrentAccum(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->DragCurrentAccum = v; }
float ImGuiContext_GetDragCurrentAccum(ImGuiContext *self) { return self->DragCurrentAccum; }
void ImGuiContext_SetDragSpeedDefaultRatio(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->DragSpeedDefaultRatio = v; }
float ImGuiContext_GetDragSpeedDefaultRatio(ImGuiContext *self) { return self->DragSpeedDefaultRatio; }
void ImGuiContext_SetScrollbarClickDeltaToGrabCenter(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->ScrollbarClickDeltaToGrabCenter = v; }
float ImGuiContext_GetScrollbarClickDeltaToGrabCenter(ImGuiContext *self) { return self->ScrollbarClickDeltaToGrabCenter; }
void ImGuiContext_SetDisabledAlphaBackup(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->DisabledAlphaBackup = v; }
float ImGuiContext_GetDisabledAlphaBackup(ImGuiContext *self) { return self->DisabledAlphaBackup; }
void ImGuiContext_SetDisabledStackSize(ImGuiContext *ImGuiContextPtr, short v) { ImGuiContextPtr->DisabledStackSize = v; }
short ImGuiContext_GetDisabledStackSize(ImGuiContext *self) { return self->DisabledStackSize; }
void ImGuiContext_SetTooltipOverrideCount(ImGuiContext *ImGuiContextPtr, short v) { ImGuiContextPtr->TooltipOverrideCount = v; }
short ImGuiContext_GetTooltipOverrideCount(ImGuiContext *self) { return self->TooltipOverrideCount; }
void ImGuiContext_SetTooltipSlowDelay(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->TooltipSlowDelay = v; }
float ImGuiContext_GetTooltipSlowDelay(ImGuiContext *self) { return self->TooltipSlowDelay; }
void ImGuiContext_SetClipboardHandlerData(ImGuiContext *ImGuiContextPtr, ImVector_char v) { ImGuiContextPtr->ClipboardHandlerData = v; }
ImVector_char ImGuiContext_GetClipboardHandlerData(ImGuiContext *self) { return self->ClipboardHandlerData; }
void ImGuiContext_SetMenusIdSubmittedThisFrame(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiID v) { ImGuiContextPtr->MenusIdSubmittedThisFrame = v; }
ImVector_ImGuiID ImGuiContext_GetMenusIdSubmittedThisFrame(ImGuiContext *self) { return self->MenusIdSubmittedThisFrame; }
void ImGuiContext_SetPlatformImeData(ImGuiContext *ImGuiContextPtr, ImGuiPlatformImeData v) { ImGuiContextPtr->PlatformImeData = v; }
ImGuiPlatformImeData ImGuiContext_GetPlatformImeData(ImGuiContext *self) { return self->PlatformImeData; }
void ImGuiContext_SetPlatformImeDataPrev(ImGuiContext *ImGuiContextPtr, ImGuiPlatformImeData v) { ImGuiContextPtr->PlatformImeDataPrev = v; }
ImGuiPlatformImeData ImGuiContext_GetPlatformImeDataPrev(ImGuiContext *self) { return self->PlatformImeDataPrev; }
void ImGuiContext_SetPlatformImeViewport(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->PlatformImeViewport = v; }
ImGuiID ImGuiContext_GetPlatformImeViewport(ImGuiContext *self) { return self->PlatformImeViewport; }
void ImGuiContext_SetPlatformLocaleDecimalPoint(ImGuiContext *ImGuiContextPtr, char v) { ImGuiContextPtr->PlatformLocaleDecimalPoint = v; }
char ImGuiContext_GetPlatformLocaleDecimalPoint(ImGuiContext *self) { return self->PlatformLocaleDecimalPoint; }
void ImGuiContext_SetDockContext(ImGuiContext *ImGuiContextPtr, ImGuiDockContext v) { ImGuiContextPtr->DockContext = v; }
ImGuiDockContext ImGuiContext_GetDockContext(ImGuiContext *self) { return self->DockContext; }
void ImGuiContext_SetSettingsLoaded(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->SettingsLoaded = v; }
bool ImGuiContext_GetSettingsLoaded(ImGuiContext *self) { return self->SettingsLoaded; }
void ImGuiContext_SetSettingsDirtyTimer(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->SettingsDirtyTimer = v; }
float ImGuiContext_GetSettingsDirtyTimer(ImGuiContext *self) { return self->SettingsDirtyTimer; }
void ImGuiContext_SetSettingsIniData(ImGuiContext *ImGuiContextPtr, ImGuiTextBuffer v) { ImGuiContextPtr->SettingsIniData = v; }
ImGuiTextBuffer ImGuiContext_GetSettingsIniData(ImGuiContext *self) { return self->SettingsIniData; }
void ImGuiContext_SetSettingsHandlers(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiSettingsHandler v) { ImGuiContextPtr->SettingsHandlers = v; }
ImVector_ImGuiSettingsHandler ImGuiContext_GetSettingsHandlers(ImGuiContext *self) { return self->SettingsHandlers; }
void ImGuiContext_SetSettingsWindows(ImGuiContext *ImGuiContextPtr, ImChunkStream_ImGuiWindowSettings v) { ImGuiContextPtr->SettingsWindows = v; }
ImChunkStream_ImGuiWindowSettings ImGuiContext_GetSettingsWindows(ImGuiContext *self) { return self->SettingsWindows; }
void ImGuiContext_SetSettingsTables(ImGuiContext *ImGuiContextPtr, ImChunkStream_ImGuiTableSettings v) { ImGuiContextPtr->SettingsTables = v; }
ImChunkStream_ImGuiTableSettings ImGuiContext_GetSettingsTables(ImGuiContext *self) { return self->SettingsTables; }
void ImGuiContext_SetHooks(ImGuiContext *ImGuiContextPtr, ImVector_ImGuiContextHook v) { ImGuiContextPtr->Hooks = v; }
ImVector_ImGuiContextHook ImGuiContext_GetHooks(ImGuiContext *self) { return self->Hooks; }
void ImGuiContext_SetHookIdNext(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->HookIdNext = v; }
ImGuiID ImGuiContext_GetHookIdNext(ImGuiContext *self) { return self->HookIdNext; }
void ImGuiContext_SetLogEnabled(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->LogEnabled = v; }
bool ImGuiContext_GetLogEnabled(ImGuiContext *self) { return self->LogEnabled; }
void ImGuiContext_SetLogType(ImGuiContext *ImGuiContextPtr, ImGuiLogType v) { ImGuiContextPtr->LogType = v; }
ImGuiLogType ImGuiContext_GetLogType(ImGuiContext *self) { return self->LogType; }
void ImGuiContext_SetLogFile(ImGuiContext *ImGuiContextPtr, ImFileHandle v) { ImGuiContextPtr->LogFile = v; }
ImFileHandle ImGuiContext_GetLogFile(ImGuiContext *self) { return self->LogFile; }
void ImGuiContext_SetLogBuffer(ImGuiContext *ImGuiContextPtr, ImGuiTextBuffer v) { ImGuiContextPtr->LogBuffer = v; }
ImGuiTextBuffer ImGuiContext_GetLogBuffer(ImGuiContext *self) { return self->LogBuffer; }
void ImGuiContext_SetLogNextPrefix(ImGuiContext *ImGuiContextPtr, const char* v) { ImGuiContextPtr->LogNextPrefix = v; }
const char* ImGuiContext_GetLogNextPrefix(ImGuiContext *self) { return self->LogNextPrefix; }
void ImGuiContext_SetLogNextSuffix(ImGuiContext *ImGuiContextPtr, const char* v) { ImGuiContextPtr->LogNextSuffix = v; }
const char* ImGuiContext_GetLogNextSuffix(ImGuiContext *self) { return self->LogNextSuffix; }
void ImGuiContext_SetLogLinePosY(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->LogLinePosY = v; }
float ImGuiContext_GetLogLinePosY(ImGuiContext *self) { return self->LogLinePosY; }
void ImGuiContext_SetLogLineFirstItem(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->LogLineFirstItem = v; }
bool ImGuiContext_GetLogLineFirstItem(ImGuiContext *self) { return self->LogLineFirstItem; }
void ImGuiContext_SetLogDepthRef(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->LogDepthRef = v; }
int ImGuiContext_GetLogDepthRef(ImGuiContext *self) { return self->LogDepthRef; }
void ImGuiContext_SetLogDepthToExpand(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->LogDepthToExpand = v; }
int ImGuiContext_GetLogDepthToExpand(ImGuiContext *self) { return self->LogDepthToExpand; }
void ImGuiContext_SetLogDepthToExpandDefault(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->LogDepthToExpandDefault = v; }
int ImGuiContext_GetLogDepthToExpandDefault(ImGuiContext *self) { return self->LogDepthToExpandDefault; }
void ImGuiContext_SetDebugLogFlags(ImGuiContext *ImGuiContextPtr, ImGuiDebugLogFlags v) { ImGuiContextPtr->DebugLogFlags = v; }
ImGuiDebugLogFlags ImGuiContext_GetDebugLogFlags(ImGuiContext *self) { return self->DebugLogFlags; }
void ImGuiContext_SetDebugLogBuf(ImGuiContext *ImGuiContextPtr, ImGuiTextBuffer v) { ImGuiContextPtr->DebugLogBuf = v; }
ImGuiTextBuffer ImGuiContext_GetDebugLogBuf(ImGuiContext *self) { return self->DebugLogBuf; }
void ImGuiContext_SetDebugItemPickerActive(ImGuiContext *ImGuiContextPtr, bool v) { ImGuiContextPtr->DebugItemPickerActive = v; }
bool ImGuiContext_GetDebugItemPickerActive(ImGuiContext *self) { return self->DebugItemPickerActive; }
void ImGuiContext_SetDebugItemPickerMouseButton(ImGuiContext *ImGuiContextPtr, ImU8 v) { ImGuiContextPtr->DebugItemPickerMouseButton = v; }
ImU8 ImGuiContext_GetDebugItemPickerMouseButton(ImGuiContext *self) { return self->DebugItemPickerMouseButton; }
void ImGuiContext_SetDebugItemPickerBreakId(ImGuiContext *ImGuiContextPtr, ImGuiID v) { ImGuiContextPtr->DebugItemPickerBreakId = v; }
ImGuiID ImGuiContext_GetDebugItemPickerBreakId(ImGuiContext *self) { return self->DebugItemPickerBreakId; }
void ImGuiContext_SetDebugMetricsConfig(ImGuiContext *ImGuiContextPtr, ImGuiMetricsConfig v) { ImGuiContextPtr->DebugMetricsConfig = v; }
ImGuiMetricsConfig ImGuiContext_GetDebugMetricsConfig(ImGuiContext *self) { return self->DebugMetricsConfig; }
void ImGuiContext_SetDebugStackTool(ImGuiContext *ImGuiContextPtr, ImGuiStackTool v) { ImGuiContextPtr->DebugStackTool = v; }
ImGuiStackTool ImGuiContext_GetDebugStackTool(ImGuiContext *self) { return self->DebugStackTool; }
void ImGuiContext_SetFramerateSecPerFrameIdx(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->FramerateSecPerFrameIdx = v; }
int ImGuiContext_GetFramerateSecPerFrameIdx(ImGuiContext *self) { return self->FramerateSecPerFrameIdx; }
void ImGuiContext_SetFramerateSecPerFrameCount(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->FramerateSecPerFrameCount = v; }
int ImGuiContext_GetFramerateSecPerFrameCount(ImGuiContext *self) { return self->FramerateSecPerFrameCount; }
void ImGuiContext_SetFramerateSecPerFrameAccum(ImGuiContext *ImGuiContextPtr, float v) { ImGuiContextPtr->FramerateSecPerFrameAccum = v; }
float ImGuiContext_GetFramerateSecPerFrameAccum(ImGuiContext *self) { return self->FramerateSecPerFrameAccum; }
void ImGuiContext_SetWantCaptureMouseNextFrame(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->WantCaptureMouseNextFrame = v; }
int ImGuiContext_GetWantCaptureMouseNextFrame(ImGuiContext *self) { return self->WantCaptureMouseNextFrame; }
void ImGuiContext_SetWantCaptureKeyboardNextFrame(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->WantCaptureKeyboardNextFrame = v; }
int ImGuiContext_GetWantCaptureKeyboardNextFrame(ImGuiContext *self) { return self->WantCaptureKeyboardNextFrame; }
void ImGuiContext_SetWantTextInputNextFrame(ImGuiContext *ImGuiContextPtr, int v) { ImGuiContextPtr->WantTextInputNextFrame = v; }
int ImGuiContext_GetWantTextInputNextFrame(ImGuiContext *self) { return self->WantTextInputNextFrame; }
void ImGuiContext_SetTempBuffer(ImGuiContext *ImGuiContextPtr, ImVector_char v) { ImGuiContextPtr->TempBuffer = v; }
ImVector_char ImGuiContext_GetTempBuffer(ImGuiContext *self) { return self->TempBuffer; }
void ImGuiDockContext_SetNodes(ImGuiDockContext *ImGuiDockContextPtr, ImGuiStorage v) { ImGuiDockContextPtr->Nodes = v; }
ImGuiStorage ImGuiDockContext_GetNodes(ImGuiDockContext *self) { return self->Nodes; }
void ImGuiDockContext_SetRequests(ImGuiDockContext *ImGuiDockContextPtr, ImVector_ImGuiDockRequest v) { ImGuiDockContextPtr->Requests = v; }
ImVector_ImGuiDockRequest ImGuiDockContext_GetRequests(ImGuiDockContext *self) { return self->Requests; }
void ImGuiDockContext_SetNodesSettings(ImGuiDockContext *ImGuiDockContextPtr, ImVector_ImGuiDockNodeSettings v) { ImGuiDockContextPtr->NodesSettings = v; }
ImVector_ImGuiDockNodeSettings ImGuiDockContext_GetNodesSettings(ImGuiDockContext *self) { return self->NodesSettings; }
void ImGuiDockContext_SetWantFullRebuild(ImGuiDockContext *ImGuiDockContextPtr, bool v) { ImGuiDockContextPtr->WantFullRebuild = v; }
bool ImGuiDockContext_GetWantFullRebuild(ImGuiDockContext *self) { return self->WantFullRebuild; }
void ImGuiGroupData_SetWindowID(ImGuiGroupData *ImGuiGroupDataPtr, ImGuiID v) { ImGuiGroupDataPtr->WindowID = v; }
ImGuiID ImGuiGroupData_GetWindowID(ImGuiGroupData *self) { return self->WindowID; }
void ImGuiGroupData_SetBackupCursorPos(ImGuiGroupData *ImGuiGroupDataPtr, ImVec2 v) { ImGuiGroupDataPtr->BackupCursorPos = v; }
ImVec2 ImGuiGroupData_GetBackupCursorPos(ImGuiGroupData *self) { return self->BackupCursorPos; }
void ImGuiGroupData_SetBackupCursorMaxPos(ImGuiGroupData *ImGuiGroupDataPtr, ImVec2 v) { ImGuiGroupDataPtr->BackupCursorMaxPos = v; }
ImVec2 ImGuiGroupData_GetBackupCursorMaxPos(ImGuiGroupData *self) { return self->BackupCursorMaxPos; }
void ImGuiGroupData_SetBackupIndent(ImGuiGroupData *ImGuiGroupDataPtr, ImVec1 v) { ImGuiGroupDataPtr->BackupIndent = v; }
ImVec1 ImGuiGroupData_GetBackupIndent(ImGuiGroupData *self) { return self->BackupIndent; }
void ImGuiGroupData_SetBackupGroupOffset(ImGuiGroupData *ImGuiGroupDataPtr, ImVec1 v) { ImGuiGroupDataPtr->BackupGroupOffset = v; }
ImVec1 ImGuiGroupData_GetBackupGroupOffset(ImGuiGroupData *self) { return self->BackupGroupOffset; }
void ImGuiGroupData_SetBackupCurrLineSize(ImGuiGroupData *ImGuiGroupDataPtr, ImVec2 v) { ImGuiGroupDataPtr->BackupCurrLineSize = v; }
ImVec2 ImGuiGroupData_GetBackupCurrLineSize(ImGuiGroupData *self) { return self->BackupCurrLineSize; }
void ImGuiGroupData_SetBackupCurrLineTextBaseOffset(ImGuiGroupData *ImGuiGroupDataPtr, float v) { ImGuiGroupDataPtr->BackupCurrLineTextBaseOffset = v; }
float ImGuiGroupData_GetBackupCurrLineTextBaseOffset(ImGuiGroupData *self) { return self->BackupCurrLineTextBaseOffset; }
void ImGuiGroupData_SetBackupActiveIdIsAlive(ImGuiGroupData *ImGuiGroupDataPtr, ImGuiID v) { ImGuiGroupDataPtr->BackupActiveIdIsAlive = v; }
ImGuiID ImGuiGroupData_GetBackupActiveIdIsAlive(ImGuiGroupData *self) { return self->BackupActiveIdIsAlive; }
void ImGuiGroupData_SetBackupActiveIdPreviousFrameIsAlive(ImGuiGroupData *ImGuiGroupDataPtr, bool v) { ImGuiGroupDataPtr->BackupActiveIdPreviousFrameIsAlive = v; }
bool ImGuiGroupData_GetBackupActiveIdPreviousFrameIsAlive(ImGuiGroupData *self) { return self->BackupActiveIdPreviousFrameIsAlive; }
void ImGuiGroupData_SetBackupHoveredIdIsAlive(ImGuiGroupData *ImGuiGroupDataPtr, bool v) { ImGuiGroupDataPtr->BackupHoveredIdIsAlive = v; }
bool ImGuiGroupData_GetBackupHoveredIdIsAlive(ImGuiGroupData *self) { return self->BackupHoveredIdIsAlive; }
void ImGuiGroupData_SetEmitItem(ImGuiGroupData *ImGuiGroupDataPtr, bool v) { ImGuiGroupDataPtr->EmitItem = v; }
bool ImGuiGroupData_GetEmitItem(ImGuiGroupData *self) { return self->EmitItem; }
void ImGuiInputEventKey_SetKey(ImGuiInputEventKey *ImGuiInputEventKeyPtr, ImGuiKey v) { ImGuiInputEventKeyPtr->Key = v; }
ImGuiKey ImGuiInputEventKey_GetKey(ImGuiInputEventKey *self) { return self->Key; }
void ImGuiInputEventKey_SetDown(ImGuiInputEventKey *ImGuiInputEventKeyPtr, bool v) { ImGuiInputEventKeyPtr->Down = v; }
bool ImGuiInputEventKey_GetDown(ImGuiInputEventKey *self) { return self->Down; }
void ImGuiInputEventKey_SetAnalogValue(ImGuiInputEventKey *ImGuiInputEventKeyPtr, float v) { ImGuiInputEventKeyPtr->AnalogValue = v; }
float ImGuiInputEventKey_GetAnalogValue(ImGuiInputEventKey *self) { return self->AnalogValue; }
void ImGuiListClipperData_SetListClipper(ImGuiListClipperData *ImGuiListClipperDataPtr, ImGuiListClipper* v) { ImGuiListClipperDataPtr->ListClipper = v; }
ImGuiListClipper* ImGuiListClipperData_GetListClipper(ImGuiListClipperData *self) { return self->ListClipper; }
void ImGuiListClipperData_SetLossynessOffset(ImGuiListClipperData *ImGuiListClipperDataPtr, float v) { ImGuiListClipperDataPtr->LossynessOffset = v; }
float ImGuiListClipperData_GetLossynessOffset(ImGuiListClipperData *self) { return self->LossynessOffset; }
void ImGuiListClipperData_SetStepNo(ImGuiListClipperData *ImGuiListClipperDataPtr, int v) { ImGuiListClipperDataPtr->StepNo = v; }
int ImGuiListClipperData_GetStepNo(ImGuiListClipperData *self) { return self->StepNo; }
void ImGuiListClipperData_SetItemsFrozen(ImGuiListClipperData *ImGuiListClipperDataPtr, int v) { ImGuiListClipperDataPtr->ItemsFrozen = v; }
int ImGuiListClipperData_GetItemsFrozen(ImGuiListClipperData *self) { return self->ItemsFrozen; }
void ImGuiListClipperData_SetRanges(ImGuiListClipperData *ImGuiListClipperDataPtr, ImVector_ImGuiListClipperRange v) { ImGuiListClipperDataPtr->Ranges = v; }
ImVector_ImGuiListClipperRange ImGuiListClipperData_GetRanges(ImGuiListClipperData *self) { return self->Ranges; }
void ImGuiTableInstanceData_SetLastOuterHeight(ImGuiTableInstanceData *ImGuiTableInstanceDataPtr, float v) { ImGuiTableInstanceDataPtr->LastOuterHeight = v; }
float ImGuiTableInstanceData_GetLastOuterHeight(ImGuiTableInstanceData *self) { return self->LastOuterHeight; }
void ImGuiTableInstanceData_SetLastFirstRowHeight(ImGuiTableInstanceData *ImGuiTableInstanceDataPtr, float v) { ImGuiTableInstanceDataPtr->LastFirstRowHeight = v; }
float ImGuiTableInstanceData_GetLastFirstRowHeight(ImGuiTableInstanceData *self) { return self->LastFirstRowHeight; }
void ImFontGlyphRangesBuilder_SetUsedChars(ImFontGlyphRangesBuilder *ImFontGlyphRangesBuilderPtr, ImVector_ImU32 v) { ImFontGlyphRangesBuilderPtr->UsedChars = v; }
ImVector_ImU32 ImFontGlyphRangesBuilder_GetUsedChars(ImFontGlyphRangesBuilder *self) { return self->UsedChars; }
void ImGuiInputEventMouseViewport_SetHoveredViewportID(ImGuiInputEventMouseViewport *ImGuiInputEventMouseViewportPtr, ImGuiID v) { ImGuiInputEventMouseViewportPtr->HoveredViewportID = v; }
ImGuiID ImGuiInputEventMouseViewport_GetHoveredViewportID(ImGuiInputEventMouseViewport *self) { return self->HoveredViewportID; }
void ImGuiInputTextCallbackData_SetEventFlag(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, ImGuiInputTextFlags v) { ImGuiInputTextCallbackDataPtr->EventFlag = v; }
ImGuiInputTextFlags ImGuiInputTextCallbackData_GetEventFlag(ImGuiInputTextCallbackData *self) { return self->EventFlag; }
void ImGuiInputTextCallbackData_SetFlags(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, ImGuiInputTextFlags v) { ImGuiInputTextCallbackDataPtr->Flags = v; }
ImGuiInputTextFlags ImGuiInputTextCallbackData_GetFlags(ImGuiInputTextCallbackData *self) { return self->Flags; }
void ImGuiInputTextCallbackData_SetUserData(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, void* v) { ImGuiInputTextCallbackDataPtr->UserData = v; }
void* ImGuiInputTextCallbackData_GetUserData(ImGuiInputTextCallbackData *self) { return self->UserData; }
void ImGuiInputTextCallbackData_SetEventChar(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, ImWchar v) { ImGuiInputTextCallbackDataPtr->EventChar = v; }
ImWchar ImGuiInputTextCallbackData_GetEventChar(ImGuiInputTextCallbackData *self) { return self->EventChar; }
void ImGuiInputTextCallbackData_SetEventKey(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, ImGuiKey v) { ImGuiInputTextCallbackDataPtr->EventKey = v; }
ImGuiKey ImGuiInputTextCallbackData_GetEventKey(ImGuiInputTextCallbackData *self) { return self->EventKey; }
void ImGuiInputTextCallbackData_SetBuf(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, char* v) { ImGuiInputTextCallbackDataPtr->Buf = v; }
char* ImGuiInputTextCallbackData_GetBuf(ImGuiInputTextCallbackData *self) { return self->Buf; }
void ImGuiInputTextCallbackData_SetBufTextLen(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, int v) { ImGuiInputTextCallbackDataPtr->BufTextLen = v; }
int ImGuiInputTextCallbackData_GetBufTextLen(ImGuiInputTextCallbackData *self) { return self->BufTextLen; }
void ImGuiInputTextCallbackData_SetBufSize(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, int v) { ImGuiInputTextCallbackDataPtr->BufSize = v; }
int ImGuiInputTextCallbackData_GetBufSize(ImGuiInputTextCallbackData *self) { return self->BufSize; }
void ImGuiInputTextCallbackData_SetBufDirty(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, bool v) { ImGuiInputTextCallbackDataPtr->BufDirty = v; }
bool ImGuiInputTextCallbackData_GetBufDirty(ImGuiInputTextCallbackData *self) { return self->BufDirty; }
void ImGuiInputTextCallbackData_SetCursorPos(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, int v) { ImGuiInputTextCallbackDataPtr->CursorPos = v; }
int ImGuiInputTextCallbackData_GetCursorPos(ImGuiInputTextCallbackData *self) { return self->CursorPos; }
void ImGuiInputTextCallbackData_SetSelectionStart(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, int v) { ImGuiInputTextCallbackDataPtr->SelectionStart = v; }
int ImGuiInputTextCallbackData_GetSelectionStart(ImGuiInputTextCallbackData *self) { return self->SelectionStart; }
void ImGuiInputTextCallbackData_SetSelectionEnd(ImGuiInputTextCallbackData *ImGuiInputTextCallbackDataPtr, int v) { ImGuiInputTextCallbackDataPtr->SelectionEnd = v; }
int ImGuiInputTextCallbackData_GetSelectionEnd(ImGuiInputTextCallbackData *self) { return self->SelectionEnd; }
void ImGuiInputTextState_SetID(ImGuiInputTextState *ImGuiInputTextStatePtr, ImGuiID v) { ImGuiInputTextStatePtr->ID = v; }
ImGuiID ImGuiInputTextState_GetID(ImGuiInputTextState *self) { return self->ID; }
void ImGuiInputTextState_SetCurLenW(ImGuiInputTextState *ImGuiInputTextStatePtr, int v) { ImGuiInputTextStatePtr->CurLenW = v; }
int ImGuiInputTextState_GetCurLenW(ImGuiInputTextState *self) { return self->CurLenW; }
void ImGuiInputTextState_SetCurLenA(ImGuiInputTextState *ImGuiInputTextStatePtr, int v) { ImGuiInputTextStatePtr->CurLenA = v; }
int ImGuiInputTextState_GetCurLenA(ImGuiInputTextState *self) { return self->CurLenA; }
void ImGuiInputTextState_SetTextW(ImGuiInputTextState *ImGuiInputTextStatePtr, ImVector_ImWchar v) { ImGuiInputTextStatePtr->TextW = v; }
ImVector_ImWchar ImGuiInputTextState_GetTextW(ImGuiInputTextState *self) { return self->TextW; }
void ImGuiInputTextState_SetTextA(ImGuiInputTextState *ImGuiInputTextStatePtr, ImVector_char v) { ImGuiInputTextStatePtr->TextA = v; }
ImVector_char ImGuiInputTextState_GetTextA(ImGuiInputTextState *self) { return self->TextA; }
void ImGuiInputTextState_SetInitialTextA(ImGuiInputTextState *ImGuiInputTextStatePtr, ImVector_char v) { ImGuiInputTextStatePtr->InitialTextA = v; }
ImVector_char ImGuiInputTextState_GetInitialTextA(ImGuiInputTextState *self) { return self->InitialTextA; }
void ImGuiInputTextState_SetTextAIsValid(ImGuiInputTextState *ImGuiInputTextStatePtr, bool v) { ImGuiInputTextStatePtr->TextAIsValid = v; }
bool ImGuiInputTextState_GetTextAIsValid(ImGuiInputTextState *self) { return self->TextAIsValid; }
void ImGuiInputTextState_SetBufCapacityA(ImGuiInputTextState *ImGuiInputTextStatePtr, int v) { ImGuiInputTextStatePtr->BufCapacityA = v; }
int ImGuiInputTextState_GetBufCapacityA(ImGuiInputTextState *self) { return self->BufCapacityA; }
void ImGuiInputTextState_SetScrollX(ImGuiInputTextState *ImGuiInputTextStatePtr, float v) { ImGuiInputTextStatePtr->ScrollX = v; }
float ImGuiInputTextState_GetScrollX(ImGuiInputTextState *self) { return self->ScrollX; }
void ImGuiInputTextState_SetStb(ImGuiInputTextState *ImGuiInputTextStatePtr, STB_TexteditState v) { ImGuiInputTextStatePtr->Stb = v; }
STB_TexteditState ImGuiInputTextState_GetStb(ImGuiInputTextState *self) { return self->Stb; }
void ImGuiInputTextState_SetCursorAnim(ImGuiInputTextState *ImGuiInputTextStatePtr, float v) { ImGuiInputTextStatePtr->CursorAnim = v; }
float ImGuiInputTextState_GetCursorAnim(ImGuiInputTextState *self) { return self->CursorAnim; }
void ImGuiInputTextState_SetCursorFollow(ImGuiInputTextState *ImGuiInputTextStatePtr, bool v) { ImGuiInputTextStatePtr->CursorFollow = v; }
bool ImGuiInputTextState_GetCursorFollow(ImGuiInputTextState *self) { return self->CursorFollow; }
void ImGuiInputTextState_SetSelectedAllMouseLock(ImGuiInputTextState *ImGuiInputTextStatePtr, bool v) { ImGuiInputTextStatePtr->SelectedAllMouseLock = v; }
bool ImGuiInputTextState_GetSelectedAllMouseLock(ImGuiInputTextState *self) { return self->SelectedAllMouseLock; }
void ImGuiInputTextState_SetEdited(ImGuiInputTextState *ImGuiInputTextStatePtr, bool v) { ImGuiInputTextStatePtr->Edited = v; }
bool ImGuiInputTextState_GetEdited(ImGuiInputTextState *self) { return self->Edited; }
void ImGuiInputTextState_SetFlags(ImGuiInputTextState *ImGuiInputTextStatePtr, ImGuiInputTextFlags v) { ImGuiInputTextStatePtr->Flags = v; }
ImGuiInputTextFlags ImGuiInputTextState_GetFlags(ImGuiInputTextState *self) { return self->Flags; }
void ImGuiStackSizes_SetSizeOfIDStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfIDStack = v; }
short ImGuiStackSizes_GetSizeOfIDStack(ImGuiStackSizes *self) { return self->SizeOfIDStack; }
void ImGuiStackSizes_SetSizeOfColorStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfColorStack = v; }
short ImGuiStackSizes_GetSizeOfColorStack(ImGuiStackSizes *self) { return self->SizeOfColorStack; }
void ImGuiStackSizes_SetSizeOfStyleVarStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfStyleVarStack = v; }
short ImGuiStackSizes_GetSizeOfStyleVarStack(ImGuiStackSizes *self) { return self->SizeOfStyleVarStack; }
void ImGuiStackSizes_SetSizeOfFontStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfFontStack = v; }
short ImGuiStackSizes_GetSizeOfFontStack(ImGuiStackSizes *self) { return self->SizeOfFontStack; }
void ImGuiStackSizes_SetSizeOfFocusScopeStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfFocusScopeStack = v; }
short ImGuiStackSizes_GetSizeOfFocusScopeStack(ImGuiStackSizes *self) { return self->SizeOfFocusScopeStack; }
void ImGuiStackSizes_SetSizeOfGroupStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfGroupStack = v; }
short ImGuiStackSizes_GetSizeOfGroupStack(ImGuiStackSizes *self) { return self->SizeOfGroupStack; }
void ImGuiStackSizes_SetSizeOfItemFlagsStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfItemFlagsStack = v; }
short ImGuiStackSizes_GetSizeOfItemFlagsStack(ImGuiStackSizes *self) { return self->SizeOfItemFlagsStack; }
void ImGuiStackSizes_SetSizeOfBeginPopupStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfBeginPopupStack = v; }
short ImGuiStackSizes_GetSizeOfBeginPopupStack(ImGuiStackSizes *self) { return self->SizeOfBeginPopupStack; }
void ImGuiStackSizes_SetSizeOfDisabledStack(ImGuiStackSizes *ImGuiStackSizesPtr, short v) { ImGuiStackSizesPtr->SizeOfDisabledStack = v; }
short ImGuiStackSizes_GetSizeOfDisabledStack(ImGuiStackSizes *self) { return self->SizeOfDisabledStack; }
void ImGuiTabItem_SetID(ImGuiTabItem *ImGuiTabItemPtr, ImGuiID v) { ImGuiTabItemPtr->ID = v; }
ImGuiID ImGuiTabItem_GetID(ImGuiTabItem *self) { return self->ID; }
void ImGuiTabItem_SetFlags(ImGuiTabItem *ImGuiTabItemPtr, ImGuiTabItemFlags v) { ImGuiTabItemPtr->Flags = v; }
ImGuiTabItemFlags ImGuiTabItem_GetFlags(ImGuiTabItem *self) { return self->Flags; }
void ImGuiTabItem_SetWindow(ImGuiTabItem *ImGuiTabItemPtr, ImGuiWindow* v) { ImGuiTabItemPtr->Window = v; }
ImGuiWindow* ImGuiTabItem_GetWindow(ImGuiTabItem *self) { return self->Window; }
void ImGuiTabItem_SetLastFrameVisible(ImGuiTabItem *ImGuiTabItemPtr, int v) { ImGuiTabItemPtr->LastFrameVisible = v; }
int ImGuiTabItem_GetLastFrameVisible(ImGuiTabItem *self) { return self->LastFrameVisible; }
void ImGuiTabItem_SetLastFrameSelected(ImGuiTabItem *ImGuiTabItemPtr, int v) { ImGuiTabItemPtr->LastFrameSelected = v; }
int ImGuiTabItem_GetLastFrameSelected(ImGuiTabItem *self) { return self->LastFrameSelected; }
void ImGuiTabItem_SetOffset(ImGuiTabItem *ImGuiTabItemPtr, float v) { ImGuiTabItemPtr->Offset = v; }
float ImGuiTabItem_GetOffset(ImGuiTabItem *self) { return self->Offset; }
void ImGuiTabItem_SetWidth(ImGuiTabItem *ImGuiTabItemPtr, float v) { ImGuiTabItemPtr->Width = v; }
float ImGuiTabItem_GetWidth(ImGuiTabItem *self) { return self->Width; }
void ImGuiTabItem_SetContentWidth(ImGuiTabItem *ImGuiTabItemPtr, float v) { ImGuiTabItemPtr->ContentWidth = v; }
float ImGuiTabItem_GetContentWidth(ImGuiTabItem *self) { return self->ContentWidth; }
void ImGuiTabItem_SetRequestedWidth(ImGuiTabItem *ImGuiTabItemPtr, float v) { ImGuiTabItemPtr->RequestedWidth = v; }
float ImGuiTabItem_GetRequestedWidth(ImGuiTabItem *self) { return self->RequestedWidth; }
void ImGuiTabItem_SetNameOffset(ImGuiTabItem *ImGuiTabItemPtr, ImS32 v) { ImGuiTabItemPtr->NameOffset = v; }
ImS32 ImGuiTabItem_GetNameOffset(ImGuiTabItem *self) { return self->NameOffset; }
void ImGuiTabItem_SetBeginOrder(ImGuiTabItem *ImGuiTabItemPtr, ImS16 v) { ImGuiTabItemPtr->BeginOrder = v; }
ImS16 ImGuiTabItem_GetBeginOrder(ImGuiTabItem *self) { return self->BeginOrder; }
void ImGuiTabItem_SetIndexDuringLayout(ImGuiTabItem *ImGuiTabItemPtr, ImS16 v) { ImGuiTabItemPtr->IndexDuringLayout = v; }
ImS16 ImGuiTabItem_GetIndexDuringLayout(ImGuiTabItem *self) { return self->IndexDuringLayout; }
void ImGuiTabItem_SetWantClose(ImGuiTabItem *ImGuiTabItemPtr, bool v) { ImGuiTabItemPtr->WantClose = v; }
bool ImGuiTabItem_GetWantClose(ImGuiTabItem *self) { return self->WantClose; }
void ImGuiWindowClass_SetClassId(ImGuiWindowClass *ImGuiWindowClassPtr, ImGuiID v) { ImGuiWindowClassPtr->ClassId = v; }
ImGuiID ImGuiWindowClass_GetClassId(ImGuiWindowClass *self) { return self->ClassId; }
void ImGuiWindowClass_SetParentViewportId(ImGuiWindowClass *ImGuiWindowClassPtr, ImGuiID v) { ImGuiWindowClassPtr->ParentViewportId = v; }
ImGuiID ImGuiWindowClass_GetParentViewportId(ImGuiWindowClass *self) { return self->ParentViewportId; }
void ImGuiWindowClass_SetViewportFlagsOverrideSet(ImGuiWindowClass *ImGuiWindowClassPtr, ImGuiViewportFlags v) { ImGuiWindowClassPtr->ViewportFlagsOverrideSet = v; }
ImGuiViewportFlags ImGuiWindowClass_GetViewportFlagsOverrideSet(ImGuiWindowClass *self) { return self->ViewportFlagsOverrideSet; }
void ImGuiWindowClass_SetViewportFlagsOverrideClear(ImGuiWindowClass *ImGuiWindowClassPtr, ImGuiViewportFlags v) { ImGuiWindowClassPtr->ViewportFlagsOverrideClear = v; }
ImGuiViewportFlags ImGuiWindowClass_GetViewportFlagsOverrideClear(ImGuiWindowClass *self) { return self->ViewportFlagsOverrideClear; }
void ImGuiWindowClass_SetTabItemFlagsOverrideSet(ImGuiWindowClass *ImGuiWindowClassPtr, ImGuiTabItemFlags v) { ImGuiWindowClassPtr->TabItemFlagsOverrideSet = v; }
ImGuiTabItemFlags ImGuiWindowClass_GetTabItemFlagsOverrideSet(ImGuiWindowClass *self) { return self->TabItemFlagsOverrideSet; }
void ImGuiWindowClass_SetDockNodeFlagsOverrideSet(ImGuiWindowClass *ImGuiWindowClassPtr, ImGuiDockNodeFlags v) { ImGuiWindowClassPtr->DockNodeFlagsOverrideSet = v; }
ImGuiDockNodeFlags ImGuiWindowClass_GetDockNodeFlagsOverrideSet(ImGuiWindowClass *self) { return self->DockNodeFlagsOverrideSet; }
void ImGuiWindowClass_SetDockingAlwaysTabBar(ImGuiWindowClass *ImGuiWindowClassPtr, bool v) { ImGuiWindowClassPtr->DockingAlwaysTabBar = v; }
bool ImGuiWindowClass_GetDockingAlwaysTabBar(ImGuiWindowClass *self) { return self->DockingAlwaysTabBar; }
void ImGuiWindowClass_SetDockingAllowUnclassed(ImGuiWindowClass *ImGuiWindowClassPtr, bool v) { ImGuiWindowClassPtr->DockingAllowUnclassed = v; }
bool ImGuiWindowClass_GetDockingAllowUnclassed(ImGuiWindowClass *self) { return self->DockingAllowUnclassed; }
void ImVec2ih_Setx(ImVec2ih *ImVec2ihPtr, short v) { ImVec2ihPtr->x = v; }
short ImVec2ih_Getx(ImVec2ih *self) { return self->x; }
void ImVec2ih_Sety(ImVec2ih *ImVec2ihPtr, short v) { ImVec2ihPtr->y = v; }
short ImVec2ih_Gety(ImVec2ih *self) { return self->y; }
void ImFontConfig_SetFontData(ImFontConfig *ImFontConfigPtr, void* v) { ImFontConfigPtr->FontData = v; }
void* ImFontConfig_GetFontData(ImFontConfig *self) { return self->FontData; }
void ImFontConfig_SetFontDataSize(ImFontConfig *ImFontConfigPtr, int v) { ImFontConfigPtr->FontDataSize = v; }
int ImFontConfig_GetFontDataSize(ImFontConfig *self) { return self->FontDataSize; }
void ImFontConfig_SetFontDataOwnedByAtlas(ImFontConfig *ImFontConfigPtr, bool v) { ImFontConfigPtr->FontDataOwnedByAtlas = v; }
bool ImFontConfig_GetFontDataOwnedByAtlas(ImFontConfig *self) { return self->FontDataOwnedByAtlas; }
void ImFontConfig_SetFontNo(ImFontConfig *ImFontConfigPtr, int v) { ImFontConfigPtr->FontNo = v; }
int ImFontConfig_GetFontNo(ImFontConfig *self) { return self->FontNo; }
void ImFontConfig_SetSizePixels(ImFontConfig *ImFontConfigPtr, float v) { ImFontConfigPtr->SizePixels = v; }
float ImFontConfig_GetSizePixels(ImFontConfig *self) { return self->SizePixels; }
void ImFontConfig_SetOversampleH(ImFontConfig *ImFontConfigPtr, int v) { ImFontConfigPtr->OversampleH = v; }
int ImFontConfig_GetOversampleH(ImFontConfig *self) { return self->OversampleH; }
void ImFontConfig_SetOversampleV(ImFontConfig *ImFontConfigPtr, int v) { ImFontConfigPtr->OversampleV = v; }
int ImFontConfig_GetOversampleV(ImFontConfig *self) { return self->OversampleV; }
void ImFontConfig_SetPixelSnapH(ImFontConfig *ImFontConfigPtr, bool v) { ImFontConfigPtr->PixelSnapH = v; }
bool ImFontConfig_GetPixelSnapH(ImFontConfig *self) { return self->PixelSnapH; }
void ImFontConfig_SetGlyphExtraSpacing(ImFontConfig *ImFontConfigPtr, ImVec2 v) { ImFontConfigPtr->GlyphExtraSpacing = v; }
ImVec2 ImFontConfig_GetGlyphExtraSpacing(ImFontConfig *self) { return self->GlyphExtraSpacing; }
void ImFontConfig_SetGlyphOffset(ImFontConfig *ImFontConfigPtr, ImVec2 v) { ImFontConfigPtr->GlyphOffset = v; }
ImVec2 ImFontConfig_GetGlyphOffset(ImFontConfig *self) { return self->GlyphOffset; }
void ImFontConfig_SetGlyphRanges(ImFontConfig *ImFontConfigPtr, const ImWchar* v) { ImFontConfigPtr->GlyphRanges = v; }
const ImWchar* ImFontConfig_GetGlyphRanges(ImFontConfig *self) { return self->GlyphRanges; }
void ImFontConfig_SetGlyphMinAdvanceX(ImFontConfig *ImFontConfigPtr, float v) { ImFontConfigPtr->GlyphMinAdvanceX = v; }
float ImFontConfig_GetGlyphMinAdvanceX(ImFontConfig *self) { return self->GlyphMinAdvanceX; }
void ImFontConfig_SetGlyphMaxAdvanceX(ImFontConfig *ImFontConfigPtr, float v) { ImFontConfigPtr->GlyphMaxAdvanceX = v; }
float ImFontConfig_GetGlyphMaxAdvanceX(ImFontConfig *self) { return self->GlyphMaxAdvanceX; }
void ImFontConfig_SetMergeMode(ImFontConfig *ImFontConfigPtr, bool v) { ImFontConfigPtr->MergeMode = v; }
bool ImFontConfig_GetMergeMode(ImFontConfig *self) { return self->MergeMode; }
void ImFontConfig_SetFontBuilderFlags(ImFontConfig *ImFontConfigPtr, unsigned int v) { ImFontConfigPtr->FontBuilderFlags = v; }
unsigned int ImFontConfig_GetFontBuilderFlags(ImFontConfig *self) { return self->FontBuilderFlags; }
void ImFontConfig_SetRasterizerMultiply(ImFontConfig *ImFontConfigPtr, float v) { ImFontConfigPtr->RasterizerMultiply = v; }
float ImFontConfig_GetRasterizerMultiply(ImFontConfig *self) { return self->RasterizerMultiply; }
void ImFontConfig_SetEllipsisChar(ImFontConfig *ImFontConfigPtr, ImWchar v) { ImFontConfigPtr->EllipsisChar = v; }
ImWchar ImFontConfig_GetEllipsisChar(ImFontConfig *self) { return self->EllipsisChar; }
void ImFontConfig_SetDstFont(ImFontConfig *ImFontConfigPtr, ImFont* v) { ImFontConfigPtr->DstFont = v; }
ImFont* ImFontConfig_GetDstFont(ImFontConfig *self) { return self->DstFont; }
