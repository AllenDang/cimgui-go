package cimgui

// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

type ImFontAtlasCustomRect uintptr

func (data ImFontAtlasCustomRect) handle() *C.ImFontAtlasCustomRect {
	return (*C.ImFontAtlasCustomRect)(unsafe.Pointer(data))
}

func (data ImFontAtlasCustomRect) c() C.ImFontAtlasCustomRect {
	return *(data.handle())
}

func newImFontAtlasCustomRectFromC(cvalue C.ImFontAtlasCustomRect) ImFontAtlasCustomRect {
	return ImFontAtlasCustomRect(unsafe.Pointer(&cvalue))
}

type ImGuiComboPreviewData uintptr

func (data ImGuiComboPreviewData) handle() *C.ImGuiComboPreviewData {
	return (*C.ImGuiComboPreviewData)(unsafe.Pointer(data))
}

func (data ImGuiComboPreviewData) c() C.ImGuiComboPreviewData {
	return *(data.handle())
}

func newImGuiComboPreviewDataFromC(cvalue C.ImGuiComboPreviewData) ImGuiComboPreviewData {
	return ImGuiComboPreviewData(unsafe.Pointer(&cvalue))
}

type ImGuiTableTempData uintptr

func (data ImGuiTableTempData) handle() *C.ImGuiTableTempData {
	return (*C.ImGuiTableTempData)(unsafe.Pointer(data))
}

func (data ImGuiTableTempData) c() C.ImGuiTableTempData {
	return *(data.handle())
}

func newImGuiTableTempDataFromC(cvalue C.ImGuiTableTempData) ImGuiTableTempData {
	return ImGuiTableTempData(unsafe.Pointer(&cvalue))
}

type ImGuiViewportP uintptr

func (data ImGuiViewportP) handle() *C.ImGuiViewportP {
	return (*C.ImGuiViewportP)(unsafe.Pointer(data))
}

func (data ImGuiViewportP) c() C.ImGuiViewportP {
	return *(data.handle())
}

func newImGuiViewportPFromC(cvalue C.ImGuiViewportP) ImGuiViewportP {
	return ImGuiViewportP(unsafe.Pointer(&cvalue))
}

type ImFontAtlas uintptr

func (data ImFontAtlas) handle() *C.ImFontAtlas {
	return (*C.ImFontAtlas)(unsafe.Pointer(data))
}

func (data ImFontAtlas) c() C.ImFontAtlas {
	return *(data.handle())
}

func newImFontAtlasFromC(cvalue C.ImFontAtlas) ImFontAtlas {
	return ImFontAtlas(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventMouseWheel uintptr

func (data ImGuiInputEventMouseWheel) handle() *C.ImGuiInputEventMouseWheel {
	return (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMouseWheel) c() C.ImGuiInputEventMouseWheel {
	return *(data.handle())
}

func newImGuiInputEventMouseWheelFromC(cvalue C.ImGuiInputEventMouseWheel) ImGuiInputEventMouseWheel {
	return ImGuiInputEventMouseWheel(unsafe.Pointer(&cvalue))
}

type ImGuiMetricsConfig uintptr

func (data ImGuiMetricsConfig) handle() *C.ImGuiMetricsConfig {
	return (*C.ImGuiMetricsConfig)(unsafe.Pointer(data))
}

func (data ImGuiMetricsConfig) c() C.ImGuiMetricsConfig {
	return *(data.handle())
}

func newImGuiMetricsConfigFromC(cvalue C.ImGuiMetricsConfig) ImGuiMetricsConfig {
	return ImGuiMetricsConfig(unsafe.Pointer(&cvalue))
}

type ImGuiNavItemData uintptr

func (data ImGuiNavItemData) handle() *C.ImGuiNavItemData {
	return (*C.ImGuiNavItemData)(unsafe.Pointer(data))
}

func (data ImGuiNavItemData) c() C.ImGuiNavItemData {
	return *(data.handle())
}

func newImGuiNavItemDataFromC(cvalue C.ImGuiNavItemData) ImGuiNavItemData {
	return ImGuiNavItemData(unsafe.Pointer(&cvalue))
}

type ImGuiOldColumnData uintptr

func (data ImGuiOldColumnData) handle() *C.ImGuiOldColumnData {
	return (*C.ImGuiOldColumnData)(unsafe.Pointer(data))
}

func (data ImGuiOldColumnData) c() C.ImGuiOldColumnData {
	return *(data.handle())
}

func newImGuiOldColumnDataFromC(cvalue C.ImGuiOldColumnData) ImGuiOldColumnData {
	return ImGuiOldColumnData(unsafe.Pointer(&cvalue))
}

type ImGuiShrinkWidthItem uintptr

func (data ImGuiShrinkWidthItem) handle() *C.ImGuiShrinkWidthItem {
	return (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(data))
}

func (data ImGuiShrinkWidthItem) c() C.ImGuiShrinkWidthItem {
	return *(data.handle())
}

func newImGuiShrinkWidthItemFromC(cvalue C.ImGuiShrinkWidthItem) ImGuiShrinkWidthItem {
	return ImGuiShrinkWidthItem(unsafe.Pointer(&cvalue))
}

type ImGuiTable uintptr

func (data ImGuiTable) handle() *C.ImGuiTable {
	return (*C.ImGuiTable)(unsafe.Pointer(data))
}

func (data ImGuiTable) c() C.ImGuiTable {
	return *(data.handle())
}

func newImGuiTableFromC(cvalue C.ImGuiTable) ImGuiTable {
	return ImGuiTable(unsafe.Pointer(&cvalue))
}

type ImGuiColorMod uintptr

func (data ImGuiColorMod) handle() *C.ImGuiColorMod {
	return (*C.ImGuiColorMod)(unsafe.Pointer(data))
}

func (data ImGuiColorMod) c() C.ImGuiColorMod {
	return *(data.handle())
}

func newImGuiColorModFromC(cvalue C.ImGuiColorMod) ImGuiColorMod {
	return ImGuiColorMod(unsafe.Pointer(&cvalue))
}

type ImGuiDockNode uintptr

func (data ImGuiDockNode) handle() *C.ImGuiDockNode {
	return (*C.ImGuiDockNode)(unsafe.Pointer(data))
}

func (data ImGuiDockNode) c() C.ImGuiDockNode {
	return *(data.handle())
}

func newImGuiDockNodeFromC(cvalue C.ImGuiDockNode) ImGuiDockNode {
	return ImGuiDockNode(unsafe.Pointer(&cvalue))
}

type ImGuiKeyData uintptr

func (data ImGuiKeyData) handle() *C.ImGuiKeyData {
	return (*C.ImGuiKeyData)(unsafe.Pointer(data))
}

func (data ImGuiKeyData) c() C.ImGuiKeyData {
	return *(data.handle())
}

func newImGuiKeyDataFromC(cvalue C.ImGuiKeyData) ImGuiKeyData {
	return ImGuiKeyData(unsafe.Pointer(&cvalue))
}

type ImGuiListClipperRange uintptr

func (data ImGuiListClipperRange) handle() *C.ImGuiListClipperRange {
	return (*C.ImGuiListClipperRange)(unsafe.Pointer(data))
}

func (data ImGuiListClipperRange) c() C.ImGuiListClipperRange {
	return *(data.handle())
}

func newImGuiListClipperRangeFromC(cvalue C.ImGuiListClipperRange) ImGuiListClipperRange {
	return ImGuiListClipperRange(unsafe.Pointer(&cvalue))
}

type ImGuiPtrOrIndex uintptr

func (data ImGuiPtrOrIndex) handle() *C.ImGuiPtrOrIndex {
	return (*C.ImGuiPtrOrIndex)(unsafe.Pointer(data))
}

func (data ImGuiPtrOrIndex) c() C.ImGuiPtrOrIndex {
	return *(data.handle())
}

func newImGuiPtrOrIndexFromC(cvalue C.ImGuiPtrOrIndex) ImGuiPtrOrIndex {
	return ImGuiPtrOrIndex(unsafe.Pointer(&cvalue))
}

type ImGuiWindowTempData uintptr

func (data ImGuiWindowTempData) handle() *C.ImGuiWindowTempData {
	return (*C.ImGuiWindowTempData)(unsafe.Pointer(data))
}

func (data ImGuiWindowTempData) c() C.ImGuiWindowTempData {
	return *(data.handle())
}

func newImGuiWindowTempDataFromC(cvalue C.ImGuiWindowTempData) ImGuiWindowTempData {
	return ImGuiWindowTempData(unsafe.Pointer(&cvalue))
}

type ImFont uintptr

func (data ImFont) handle() *C.ImFont {
	return (*C.ImFont)(unsafe.Pointer(data))
}

func (data ImFont) c() C.ImFont {
	return *(data.handle())
}

func newImFontFromC(cvalue C.ImFont) ImFont {
	return ImFont(unsafe.Pointer(&cvalue))
}

type ImDrawData uintptr

func (data ImDrawData) handle() *C.ImDrawData {
	return (*C.ImDrawData)(unsafe.Pointer(data))
}

func (data ImDrawData) c() C.ImDrawData {
	return *(data.handle())
}

func newImDrawDataFromC(cvalue C.ImDrawData) ImDrawData {
	return ImDrawData(unsafe.Pointer(&cvalue))
}

type ImGuiDataTypeInfo uintptr

func (data ImGuiDataTypeInfo) handle() *C.ImGuiDataTypeInfo {
	return (*C.ImGuiDataTypeInfo)(unsafe.Pointer(data))
}

func (data ImGuiDataTypeInfo) c() C.ImGuiDataTypeInfo {
	return *(data.handle())
}

func newImGuiDataTypeInfoFromC(cvalue C.ImGuiDataTypeInfo) ImGuiDataTypeInfo {
	return ImGuiDataTypeInfo(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventMouseButton uintptr

func (data ImGuiInputEventMouseButton) handle() *C.ImGuiInputEventMouseButton {
	return (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMouseButton) c() C.ImGuiInputEventMouseButton {
	return *(data.handle())
}

func newImGuiInputEventMouseButtonFromC(cvalue C.ImGuiInputEventMouseButton) ImGuiInputEventMouseButton {
	return ImGuiInputEventMouseButton(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventMousePos uintptr

func (data ImGuiInputEventMousePos) handle() *C.ImGuiInputEventMousePos {
	return (*C.ImGuiInputEventMousePos)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMousePos) c() C.ImGuiInputEventMousePos {
	return *(data.handle())
}

func newImGuiInputEventMousePosFromC(cvalue C.ImGuiInputEventMousePos) ImGuiInputEventMousePos {
	return ImGuiInputEventMousePos(unsafe.Pointer(&cvalue))
}

type ImGuiOldColumns uintptr

func (data ImGuiOldColumns) handle() *C.ImGuiOldColumns {
	return (*C.ImGuiOldColumns)(unsafe.Pointer(data))
}

func (data ImGuiOldColumns) c() C.ImGuiOldColumns {
	return *(data.handle())
}

func newImGuiOldColumnsFromC(cvalue C.ImGuiOldColumns) ImGuiOldColumns {
	return ImGuiOldColumns(unsafe.Pointer(&cvalue))
}

type ImGuiTableColumnSortSpecs uintptr

func (data ImGuiTableColumnSortSpecs) handle() *C.ImGuiTableColumnSortSpecs {
	return (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(data))
}

func (data ImGuiTableColumnSortSpecs) c() C.ImGuiTableColumnSortSpecs {
	return *(data.handle())
}

func newImGuiTableColumnSortSpecsFromC(cvalue C.ImGuiTableColumnSortSpecs) ImGuiTableColumnSortSpecs {
	return ImGuiTableColumnSortSpecs(unsafe.Pointer(&cvalue))
}

type ImGuiTableSettings uintptr

func (data ImGuiTableSettings) handle() *C.ImGuiTableSettings {
	return (*C.ImGuiTableSettings)(unsafe.Pointer(data))
}

func (data ImGuiTableSettings) c() C.ImGuiTableSettings {
	return *(data.handle())
}

func newImGuiTableSettingsFromC(cvalue C.ImGuiTableSettings) ImGuiTableSettings {
	return ImGuiTableSettings(unsafe.Pointer(&cvalue))
}

type ImDrawChannel uintptr

func (data ImDrawChannel) handle() *C.ImDrawChannel {
	return (*C.ImDrawChannel)(unsafe.Pointer(data))
}

func (data ImDrawChannel) c() C.ImDrawChannel {
	return *(data.handle())
}

func newImDrawChannelFromC(cvalue C.ImDrawChannel) ImDrawChannel {
	return ImDrawChannel(unsafe.Pointer(&cvalue))
}

type ImGuiTextFilter uintptr

func (data ImGuiTextFilter) handle() *C.ImGuiTextFilter {
	return (*C.ImGuiTextFilter)(unsafe.Pointer(data))
}

func (data ImGuiTextFilter) c() C.ImGuiTextFilter {
	return *(data.handle())
}

func newImGuiTextFilterFromC(cvalue C.ImGuiTextFilter) ImGuiTextFilter {
	return ImGuiTextFilter(unsafe.Pointer(&cvalue))
}

type ImGuiTableSortSpecs uintptr

func (data ImGuiTableSortSpecs) handle() *C.ImGuiTableSortSpecs {
	return (*C.ImGuiTableSortSpecs)(unsafe.Pointer(data))
}

func (data ImGuiTableSortSpecs) c() C.ImGuiTableSortSpecs {
	return *(data.handle())
}

func newImGuiTableSortSpecsFromC(cvalue C.ImGuiTableSortSpecs) ImGuiTableSortSpecs {
	return ImGuiTableSortSpecs(unsafe.Pointer(&cvalue))
}

type ImGuiTabBar uintptr

func (data ImGuiTabBar) handle() *C.ImGuiTabBar {
	return (*C.ImGuiTabBar)(unsafe.Pointer(data))
}

func (data ImGuiTabBar) c() C.ImGuiTabBar {
	return *(data.handle())
}

func newImGuiTabBarFromC(cvalue C.ImGuiTabBar) ImGuiTabBar {
	return ImGuiTabBar(unsafe.Pointer(&cvalue))
}

type ImGuiViewport uintptr

func (data ImGuiViewport) handle() *C.ImGuiViewport {
	return (*C.ImGuiViewport)(unsafe.Pointer(data))
}

func (data ImGuiViewport) c() C.ImGuiViewport {
	return *(data.handle())
}

func newImGuiViewportFromC(cvalue C.ImGuiViewport) ImGuiViewport {
	return ImGuiViewport(unsafe.Pointer(&cvalue))
}

type ImGuiStackTool uintptr

func (data ImGuiStackTool) handle() *C.ImGuiStackTool {
	return (*C.ImGuiStackTool)(unsafe.Pointer(data))
}

func (data ImGuiStackTool) c() C.ImGuiStackTool {
	return *(data.handle())
}

func newImGuiStackToolFromC(cvalue C.ImGuiStackTool) ImGuiStackTool {
	return ImGuiStackTool(unsafe.Pointer(&cvalue))
}

type ImGuiStyle uintptr

func (data ImGuiStyle) handle() *C.ImGuiStyle {
	return (*C.ImGuiStyle)(unsafe.Pointer(data))
}

func (data ImGuiStyle) c() C.ImGuiStyle {
	return *(data.handle())
}

func newImGuiStyleFromC(cvalue C.ImGuiStyle) ImGuiStyle {
	return ImGuiStyle(unsafe.Pointer(&cvalue))
}

type ImGuiWindowStackData uintptr

func (data ImGuiWindowStackData) handle() *C.ImGuiWindowStackData {
	return (*C.ImGuiWindowStackData)(unsafe.Pointer(data))
}

func (data ImGuiWindowStackData) c() C.ImGuiWindowStackData {
	return *(data.handle())
}

func newImGuiWindowStackDataFromC(cvalue C.ImGuiWindowStackData) ImGuiWindowStackData {
	return ImGuiWindowStackData(unsafe.Pointer(&cvalue))
}

type ImGuiStoragePair uintptr

func (data ImGuiStoragePair) handle() *C.ImGuiStoragePair {
	return (*C.ImGuiStoragePair)(unsafe.Pointer(data))
}

func (data ImGuiStoragePair) c() C.ImGuiStoragePair {
	return *(data.handle())
}

func newImGuiStoragePairFromC(cvalue C.ImGuiStoragePair) ImGuiStoragePair {
	return ImGuiStoragePair(unsafe.Pointer(&cvalue))
}

type ImGuiIO uintptr

func (data ImGuiIO) handle() *C.ImGuiIO {
	return (*C.ImGuiIO)(unsafe.Pointer(data))
}

func (data ImGuiIO) c() C.ImGuiIO {
	return *(data.handle())
}

func newImGuiIOFromC(cvalue C.ImGuiIO) ImGuiIO {
	return ImGuiIO(unsafe.Pointer(&cvalue))
}

type ImGuiTableCellData uintptr

func (data ImGuiTableCellData) handle() *C.ImGuiTableCellData {
	return (*C.ImGuiTableCellData)(unsafe.Pointer(data))
}

func (data ImGuiTableCellData) c() C.ImGuiTableCellData {
	return *(data.handle())
}

func newImGuiTableCellDataFromC(cvalue C.ImGuiTableCellData) ImGuiTableCellData {
	return ImGuiTableCellData(unsafe.Pointer(&cvalue))
}

type ImFontBuilderIO uintptr

func (data ImFontBuilderIO) handle() *C.ImFontBuilderIO {
	return (*C.ImFontBuilderIO)(unsafe.Pointer(data))
}

func (data ImFontBuilderIO) c() C.ImFontBuilderIO {
	return *(data.handle())
}

func newImFontBuilderIOFromC(cvalue C.ImFontBuilderIO) ImFontBuilderIO {
	return ImFontBuilderIO(unsafe.Pointer(&cvalue))
}

type ImFontGlyph uintptr

func (data ImFontGlyph) handle() *C.ImFontGlyph {
	return (*C.ImFontGlyph)(unsafe.Pointer(data))
}

func (data ImFontGlyph) c() C.ImFontGlyph {
	return *(data.handle())
}

func newImFontGlyphFromC(cvalue C.ImFontGlyph) ImFontGlyph {
	return ImFontGlyph(unsafe.Pointer(&cvalue))
}

type ImGuiOnceUponAFrame uintptr

func (data ImGuiOnceUponAFrame) handle() *C.ImGuiOnceUponAFrame {
	return (*C.ImGuiOnceUponAFrame)(unsafe.Pointer(data))
}

func (data ImGuiOnceUponAFrame) c() C.ImGuiOnceUponAFrame {
	return *(data.handle())
}

func newImGuiOnceUponAFrameFromC(cvalue C.ImGuiOnceUponAFrame) ImGuiOnceUponAFrame {
	return ImGuiOnceUponAFrame(unsafe.Pointer(&cvalue))
}

type ImGuiPlatformIO uintptr

func (data ImGuiPlatformIO) handle() *C.ImGuiPlatformIO {
	return (*C.ImGuiPlatformIO)(unsafe.Pointer(data))
}

func (data ImGuiPlatformIO) c() C.ImGuiPlatformIO {
	return *(data.handle())
}

func newImGuiPlatformIOFromC(cvalue C.ImGuiPlatformIO) ImGuiPlatformIO {
	return ImGuiPlatformIO(unsafe.Pointer(&cvalue))
}

type ImGuiPopupData uintptr

func (data ImGuiPopupData) handle() *C.ImGuiPopupData {
	return (*C.ImGuiPopupData)(unsafe.Pointer(data))
}

func (data ImGuiPopupData) c() C.ImGuiPopupData {
	return *(data.handle())
}

func newImGuiPopupDataFromC(cvalue C.ImGuiPopupData) ImGuiPopupData {
	return ImGuiPopupData(unsafe.Pointer(&cvalue))
}

type ImDrawListSharedData uintptr

func (data ImDrawListSharedData) handle() *C.ImDrawListSharedData {
	return (*C.ImDrawListSharedData)(unsafe.Pointer(data))
}

func (data ImDrawListSharedData) c() C.ImDrawListSharedData {
	return *(data.handle())
}

func newImDrawListSharedDataFromC(cvalue C.ImDrawListSharedData) ImDrawListSharedData {
	return ImDrawListSharedData(unsafe.Pointer(&cvalue))
}

type ImDrawCmdHeader uintptr

func (data ImDrawCmdHeader) handle() *C.ImDrawCmdHeader {
	return (*C.ImDrawCmdHeader)(unsafe.Pointer(data))
}

func (data ImDrawCmdHeader) c() C.ImDrawCmdHeader {
	return *(data.handle())
}

func newImDrawCmdHeaderFromC(cvalue C.ImDrawCmdHeader) ImDrawCmdHeader {
	return ImDrawCmdHeader(unsafe.Pointer(&cvalue))
}

type ImDrawListSplitter uintptr

func (data ImDrawListSplitter) handle() *C.ImDrawListSplitter {
	return (*C.ImDrawListSplitter)(unsafe.Pointer(data))
}

func (data ImDrawListSplitter) c() C.ImDrawListSplitter {
	return *(data.handle())
}

func newImDrawListSplitterFromC(cvalue C.ImDrawListSplitter) ImDrawListSplitter {
	return ImDrawListSplitter(unsafe.Pointer(&cvalue))
}

type ImGuiPlatformImeData uintptr

func (data ImGuiPlatformImeData) handle() *C.ImGuiPlatformImeData {
	return (*C.ImGuiPlatformImeData)(unsafe.Pointer(data))
}

func (data ImGuiPlatformImeData) c() C.ImGuiPlatformImeData {
	return *(data.handle())
}

func newImGuiPlatformImeDataFromC(cvalue C.ImGuiPlatformImeData) ImGuiPlatformImeData {
	return ImGuiPlatformImeData(unsafe.Pointer(&cvalue))
}

type ImGuiSizeCallbackData uintptr

func (data ImGuiSizeCallbackData) handle() *C.ImGuiSizeCallbackData {
	return (*C.ImGuiSizeCallbackData)(unsafe.Pointer(data))
}

func (data ImGuiSizeCallbackData) c() C.ImGuiSizeCallbackData {
	return *(data.handle())
}

func newImGuiSizeCallbackDataFromC(cvalue C.ImGuiSizeCallbackData) ImGuiSizeCallbackData {
	return ImGuiSizeCallbackData(unsafe.Pointer(&cvalue))
}

type ImGuiTableColumn uintptr

func (data ImGuiTableColumn) handle() *C.ImGuiTableColumn {
	return (*C.ImGuiTableColumn)(unsafe.Pointer(data))
}

func (data ImGuiTableColumn) c() C.ImGuiTableColumn {
	return *(data.handle())
}

func newImGuiTableColumnFromC(cvalue C.ImGuiTableColumn) ImGuiTableColumn {
	return ImGuiTableColumn(unsafe.Pointer(&cvalue))
}

type ImGuiNextWindowData uintptr

func (data ImGuiNextWindowData) handle() *C.ImGuiNextWindowData {
	return (*C.ImGuiNextWindowData)(unsafe.Pointer(data))
}

func (data ImGuiNextWindowData) c() C.ImGuiNextWindowData {
	return *(data.handle())
}

func newImGuiNextWindowDataFromC(cvalue C.ImGuiNextWindowData) ImGuiNextWindowData {
	return ImGuiNextWindowData(unsafe.Pointer(&cvalue))
}

type ImGuiStyleMod uintptr

func (data ImGuiStyleMod) handle() *C.ImGuiStyleMod {
	return (*C.ImGuiStyleMod)(unsafe.Pointer(data))
}

func (data ImGuiStyleMod) c() C.ImGuiStyleMod {
	return *(data.handle())
}

func newImGuiStyleModFromC(cvalue C.ImGuiStyleMod) ImGuiStyleMod {
	return ImGuiStyleMod(unsafe.Pointer(&cvalue))
}

type ImGuiTextRange uintptr

func (data ImGuiTextRange) handle() *C.ImGuiTextRange {
	return (*C.ImGuiTextRange)(unsafe.Pointer(data))
}

func (data ImGuiTextRange) c() C.ImGuiTextRange {
	return *(data.handle())
}

func newImGuiTextRangeFromC(cvalue C.ImGuiTextRange) ImGuiTextRange {
	return ImGuiTextRange(unsafe.Pointer(&cvalue))
}

type ImGuiWindow uintptr

func (data ImGuiWindow) handle() *C.ImGuiWindow {
	return (*C.ImGuiWindow)(unsafe.Pointer(data))
}

func (data ImGuiWindow) c() C.ImGuiWindow {
	return *(data.handle())
}

func newImGuiWindowFromC(cvalue C.ImGuiWindow) ImGuiWindow {
	return ImGuiWindow(unsafe.Pointer(&cvalue))
}

type ImGuiWindowDockStyle uintptr

func (data ImGuiWindowDockStyle) handle() *C.ImGuiWindowDockStyle {
	return (*C.ImGuiWindowDockStyle)(unsafe.Pointer(data))
}

func (data ImGuiWindowDockStyle) c() C.ImGuiWindowDockStyle {
	return *(data.handle())
}

func newImGuiWindowDockStyleFromC(cvalue C.ImGuiWindowDockStyle) ImGuiWindowDockStyle {
	return ImGuiWindowDockStyle(unsafe.Pointer(&cvalue))
}

type ImDrawDataBuilder uintptr

func (data ImDrawDataBuilder) handle() *C.ImDrawDataBuilder {
	return (*C.ImDrawDataBuilder)(unsafe.Pointer(data))
}

func (data ImDrawDataBuilder) c() C.ImDrawDataBuilder {
	return *(data.handle())
}

func newImDrawDataBuilderFromC(cvalue C.ImDrawDataBuilder) ImDrawDataBuilder {
	return ImDrawDataBuilder(unsafe.Pointer(&cvalue))
}

type ImGuiLastItemData uintptr

func (data ImGuiLastItemData) handle() *C.ImGuiLastItemData {
	return (*C.ImGuiLastItemData)(unsafe.Pointer(data))
}

func (data ImGuiLastItemData) c() C.ImGuiLastItemData {
	return *(data.handle())
}

func newImGuiLastItemDataFromC(cvalue C.ImGuiLastItemData) ImGuiLastItemData {
	return ImGuiLastItemData(unsafe.Pointer(&cvalue))
}

type ImGuiTextBuffer uintptr

func (data ImGuiTextBuffer) handle() *C.ImGuiTextBuffer {
	return (*C.ImGuiTextBuffer)(unsafe.Pointer(data))
}

func (data ImGuiTextBuffer) c() C.ImGuiTextBuffer {
	return *(data.handle())
}

func newImGuiTextBufferFromC(cvalue C.ImGuiTextBuffer) ImGuiTextBuffer {
	return ImGuiTextBuffer(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventText uintptr

func (data ImGuiInputEventText) handle() *C.ImGuiInputEventText {
	return (*C.ImGuiInputEventText)(unsafe.Pointer(data))
}

func (data ImGuiInputEventText) c() C.ImGuiInputEventText {
	return *(data.handle())
}

func newImGuiInputEventTextFromC(cvalue C.ImGuiInputEventText) ImGuiInputEventText {
	return ImGuiInputEventText(unsafe.Pointer(&cvalue))
}

type ImDrawList uintptr

func (data ImDrawList) handle() *C.ImDrawList {
	return (*C.ImDrawList)(unsafe.Pointer(data))
}

func (data ImDrawList) c() C.ImDrawList {
	return *(data.handle())
}

func newImDrawListFromC(cvalue C.ImDrawList) ImDrawList {
	return ImDrawList(unsafe.Pointer(&cvalue))
}

type ImGuiNextItemData uintptr

func (data ImGuiNextItemData) handle() *C.ImGuiNextItemData {
	return (*C.ImGuiNextItemData)(unsafe.Pointer(data))
}

func (data ImGuiNextItemData) c() C.ImGuiNextItemData {
	return *(data.handle())
}

func newImGuiNextItemDataFromC(cvalue C.ImGuiNextItemData) ImGuiNextItemData {
	return ImGuiNextItemData(unsafe.Pointer(&cvalue))
}

type ImGuiWindowSettings uintptr

func (data ImGuiWindowSettings) handle() *C.ImGuiWindowSettings {
	return (*C.ImGuiWindowSettings)(unsafe.Pointer(data))
}

func (data ImGuiWindowSettings) c() C.ImGuiWindowSettings {
	return *(data.handle())
}

func newImGuiWindowSettingsFromC(cvalue C.ImGuiWindowSettings) ImGuiWindowSettings {
	return ImGuiWindowSettings(unsafe.Pointer(&cvalue))
}

type ImDrawCmd uintptr

func (data ImDrawCmd) handle() *C.ImDrawCmd {
	return (*C.ImDrawCmd)(unsafe.Pointer(data))
}

func (data ImDrawCmd) c() C.ImDrawCmd {
	return *(data.handle())
}

func newImDrawCmdFromC(cvalue C.ImDrawCmd) ImDrawCmd {
	return ImDrawCmd(unsafe.Pointer(&cvalue))
}

type ImGuiContextHook uintptr

func (data ImGuiContextHook) handle() *C.ImGuiContextHook {
	return (*C.ImGuiContextHook)(unsafe.Pointer(data))
}

func (data ImGuiContextHook) c() C.ImGuiContextHook {
	return *(data.handle())
}

func newImGuiContextHookFromC(cvalue C.ImGuiContextHook) ImGuiContextHook {
	return ImGuiContextHook(unsafe.Pointer(&cvalue))
}

type ImGuiDataTypeTempStorage uintptr

func (data ImGuiDataTypeTempStorage) handle() *C.ImGuiDataTypeTempStorage {
	return (*C.ImGuiDataTypeTempStorage)(unsafe.Pointer(data))
}

func (data ImGuiDataTypeTempStorage) c() C.ImGuiDataTypeTempStorage {
	return *(data.handle())
}

func newImGuiDataTypeTempStorageFromC(cvalue C.ImGuiDataTypeTempStorage) ImGuiDataTypeTempStorage {
	return ImGuiDataTypeTempStorage(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventAppFocused uintptr

func (data ImGuiInputEventAppFocused) handle() *C.ImGuiInputEventAppFocused {
	return (*C.ImGuiInputEventAppFocused)(unsafe.Pointer(data))
}

func (data ImGuiInputEventAppFocused) c() C.ImGuiInputEventAppFocused {
	return *(data.handle())
}

func newImGuiInputEventAppFocusedFromC(cvalue C.ImGuiInputEventAppFocused) ImGuiInputEventAppFocused {
	return ImGuiInputEventAppFocused(unsafe.Pointer(&cvalue))
}

type ImGuiListClipper uintptr

func (data ImGuiListClipper) handle() *C.ImGuiListClipper {
	return (*C.ImGuiListClipper)(unsafe.Pointer(data))
}

func (data ImGuiListClipper) c() C.ImGuiListClipper {
	return *(data.handle())
}

func newImGuiListClipperFromC(cvalue C.ImGuiListClipper) ImGuiListClipper {
	return ImGuiListClipper(unsafe.Pointer(&cvalue))
}

type ImGuiPlatformMonitor uintptr

func (data ImGuiPlatformMonitor) handle() *C.ImGuiPlatformMonitor {
	return (*C.ImGuiPlatformMonitor)(unsafe.Pointer(data))
}

func (data ImGuiPlatformMonitor) c() C.ImGuiPlatformMonitor {
	return *(data.handle())
}

func newImGuiPlatformMonitorFromC(cvalue C.ImGuiPlatformMonitor) ImGuiPlatformMonitor {
	return ImGuiPlatformMonitor(unsafe.Pointer(&cvalue))
}

type ImGuiSettingsHandler uintptr

func (data ImGuiSettingsHandler) handle() *C.ImGuiSettingsHandler {
	return (*C.ImGuiSettingsHandler)(unsafe.Pointer(data))
}

func (data ImGuiSettingsHandler) c() C.ImGuiSettingsHandler {
	return *(data.handle())
}

func newImGuiSettingsHandlerFromC(cvalue C.ImGuiSettingsHandler) ImGuiSettingsHandler {
	return ImGuiSettingsHandler(unsafe.Pointer(&cvalue))
}

type ImGuiStackLevelInfo uintptr

func (data ImGuiStackLevelInfo) handle() *C.ImGuiStackLevelInfo {
	return (*C.ImGuiStackLevelInfo)(unsafe.Pointer(data))
}

func (data ImGuiStackLevelInfo) c() C.ImGuiStackLevelInfo {
	return *(data.handle())
}

func newImGuiStackLevelInfoFromC(cvalue C.ImGuiStackLevelInfo) ImGuiStackLevelInfo {
	return ImGuiStackLevelInfo(unsafe.Pointer(&cvalue))
}

type ImBitVector uintptr

func (data ImBitVector) handle() *C.ImBitVector {
	return (*C.ImBitVector)(unsafe.Pointer(data))
}

func (data ImBitVector) c() C.ImBitVector {
	return *(data.handle())
}

func newImBitVectorFromC(cvalue C.ImBitVector) ImBitVector {
	return ImBitVector(unsafe.Pointer(&cvalue))
}

type ImGuiStorage uintptr

func (data ImGuiStorage) handle() *C.ImGuiStorage {
	return (*C.ImGuiStorage)(unsafe.Pointer(data))
}

func (data ImGuiStorage) c() C.ImGuiStorage {
	return *(data.handle())
}

func newImGuiStorageFromC(cvalue C.ImGuiStorage) ImGuiStorage {
	return ImGuiStorage(unsafe.Pointer(&cvalue))
}

type ImGuiInputEvent uintptr

func (data ImGuiInputEvent) handle() *C.ImGuiInputEvent {
	return (*C.ImGuiInputEvent)(unsafe.Pointer(data))
}

func (data ImGuiInputEvent) c() C.ImGuiInputEvent {
	return *(data.handle())
}

func newImGuiInputEventFromC(cvalue C.ImGuiInputEvent) ImGuiInputEvent {
	return ImGuiInputEvent(unsafe.Pointer(&cvalue))
}

type ImGuiMenuColumns uintptr

func (data ImGuiMenuColumns) handle() *C.ImGuiMenuColumns {
	return (*C.ImGuiMenuColumns)(unsafe.Pointer(data))
}

func (data ImGuiMenuColumns) c() C.ImGuiMenuColumns {
	return *(data.handle())
}

func newImGuiMenuColumnsFromC(cvalue C.ImGuiMenuColumns) ImGuiMenuColumns {
	return ImGuiMenuColumns(unsafe.Pointer(&cvalue))
}

type ImGuiPayload uintptr

func (data ImGuiPayload) handle() *C.ImGuiPayload {
	return (*C.ImGuiPayload)(unsafe.Pointer(data))
}

func (data ImGuiPayload) c() C.ImGuiPayload {
	return *(data.handle())
}

func newImGuiPayloadFromC(cvalue C.ImGuiPayload) ImGuiPayload {
	return ImGuiPayload(unsafe.Pointer(&cvalue))
}

type ImGuiTableColumnSettings uintptr

func (data ImGuiTableColumnSettings) handle() *C.ImGuiTableColumnSettings {
	return (*C.ImGuiTableColumnSettings)(unsafe.Pointer(data))
}

func (data ImGuiTableColumnSettings) c() C.ImGuiTableColumnSettings {
	return *(data.handle())
}

func newImGuiTableColumnSettingsFromC(cvalue C.ImGuiTableColumnSettings) ImGuiTableColumnSettings {
	return ImGuiTableColumnSettings(unsafe.Pointer(&cvalue))
}

type ImDrawVert uintptr

func (data ImDrawVert) handle() *C.ImDrawVert {
	return (*C.ImDrawVert)(unsafe.Pointer(data))
}

func (data ImDrawVert) c() C.ImDrawVert {
	return *(data.handle())
}

func newImDrawVertFromC(cvalue C.ImDrawVert) ImDrawVert {
	return ImDrawVert(unsafe.Pointer(&cvalue))
}

type ImGuiContext uintptr

func (data ImGuiContext) handle() *C.ImGuiContext {
	return (*C.ImGuiContext)(unsafe.Pointer(data))
}

func (data ImGuiContext) c() C.ImGuiContext {
	return *(data.handle())
}

func newImGuiContextFromC(cvalue C.ImGuiContext) ImGuiContext {
	return ImGuiContext(unsafe.Pointer(&cvalue))
}

type ImGuiDockContext uintptr

func (data ImGuiDockContext) handle() *C.ImGuiDockContext {
	return (*C.ImGuiDockContext)(unsafe.Pointer(data))
}

func (data ImGuiDockContext) c() C.ImGuiDockContext {
	return *(data.handle())
}

func newImGuiDockContextFromC(cvalue C.ImGuiDockContext) ImGuiDockContext {
	return ImGuiDockContext(unsafe.Pointer(&cvalue))
}

type ImGuiGroupData uintptr

func (data ImGuiGroupData) handle() *C.ImGuiGroupData {
	return (*C.ImGuiGroupData)(unsafe.Pointer(data))
}

func (data ImGuiGroupData) c() C.ImGuiGroupData {
	return *(data.handle())
}

func newImGuiGroupDataFromC(cvalue C.ImGuiGroupData) ImGuiGroupData {
	return ImGuiGroupData(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventKey uintptr

func (data ImGuiInputEventKey) handle() *C.ImGuiInputEventKey {
	return (*C.ImGuiInputEventKey)(unsafe.Pointer(data))
}

func (data ImGuiInputEventKey) c() C.ImGuiInputEventKey {
	return *(data.handle())
}

func newImGuiInputEventKeyFromC(cvalue C.ImGuiInputEventKey) ImGuiInputEventKey {
	return ImGuiInputEventKey(unsafe.Pointer(&cvalue))
}

type ImGuiListClipperData uintptr

func (data ImGuiListClipperData) handle() *C.ImGuiListClipperData {
	return (*C.ImGuiListClipperData)(unsafe.Pointer(data))
}

func (data ImGuiListClipperData) c() C.ImGuiListClipperData {
	return *(data.handle())
}

func newImGuiListClipperDataFromC(cvalue C.ImGuiListClipperData) ImGuiListClipperData {
	return ImGuiListClipperData(unsafe.Pointer(&cvalue))
}

type ImGuiTableInstanceData uintptr

func (data ImGuiTableInstanceData) handle() *C.ImGuiTableInstanceData {
	return (*C.ImGuiTableInstanceData)(unsafe.Pointer(data))
}

func (data ImGuiTableInstanceData) c() C.ImGuiTableInstanceData {
	return *(data.handle())
}

func newImGuiTableInstanceDataFromC(cvalue C.ImGuiTableInstanceData) ImGuiTableInstanceData {
	return ImGuiTableInstanceData(unsafe.Pointer(&cvalue))
}

type ImFontGlyphRangesBuilder uintptr

func (data ImFontGlyphRangesBuilder) handle() *C.ImFontGlyphRangesBuilder {
	return (*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(data))
}

func (data ImFontGlyphRangesBuilder) c() C.ImFontGlyphRangesBuilder {
	return *(data.handle())
}

func newImFontGlyphRangesBuilderFromC(cvalue C.ImFontGlyphRangesBuilder) ImFontGlyphRangesBuilder {
	return ImFontGlyphRangesBuilder(unsafe.Pointer(&cvalue))
}

type ImGuiInputEventMouseViewport uintptr

func (data ImGuiInputEventMouseViewport) handle() *C.ImGuiInputEventMouseViewport {
	return (*C.ImGuiInputEventMouseViewport)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMouseViewport) c() C.ImGuiInputEventMouseViewport {
	return *(data.handle())
}

func newImGuiInputEventMouseViewportFromC(cvalue C.ImGuiInputEventMouseViewport) ImGuiInputEventMouseViewport {
	return ImGuiInputEventMouseViewport(unsafe.Pointer(&cvalue))
}

type ImGuiInputTextCallbackData uintptr

func (data ImGuiInputTextCallbackData) handle() *C.ImGuiInputTextCallbackData {
	return (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(data))
}

func (data ImGuiInputTextCallbackData) c() C.ImGuiInputTextCallbackData {
	return *(data.handle())
}

func newImGuiInputTextCallbackDataFromC(cvalue C.ImGuiInputTextCallbackData) ImGuiInputTextCallbackData {
	return ImGuiInputTextCallbackData(unsafe.Pointer(&cvalue))
}

type ImGuiInputTextState uintptr

func (data ImGuiInputTextState) handle() *C.ImGuiInputTextState {
	return (*C.ImGuiInputTextState)(unsafe.Pointer(data))
}

func (data ImGuiInputTextState) c() C.ImGuiInputTextState {
	return *(data.handle())
}

func newImGuiInputTextStateFromC(cvalue C.ImGuiInputTextState) ImGuiInputTextState {
	return ImGuiInputTextState(unsafe.Pointer(&cvalue))
}

type ImGuiStackSizes uintptr

func (data ImGuiStackSizes) handle() *C.ImGuiStackSizes {
	return (*C.ImGuiStackSizes)(unsafe.Pointer(data))
}

func (data ImGuiStackSizes) c() C.ImGuiStackSizes {
	return *(data.handle())
}

func newImGuiStackSizesFromC(cvalue C.ImGuiStackSizes) ImGuiStackSizes {
	return ImGuiStackSizes(unsafe.Pointer(&cvalue))
}

type ImGuiTabItem uintptr

func (data ImGuiTabItem) handle() *C.ImGuiTabItem {
	return (*C.ImGuiTabItem)(unsafe.Pointer(data))
}

func (data ImGuiTabItem) c() C.ImGuiTabItem {
	return *(data.handle())
}

func newImGuiTabItemFromC(cvalue C.ImGuiTabItem) ImGuiTabItem {
	return ImGuiTabItem(unsafe.Pointer(&cvalue))
}

type ImGuiWindowClass uintptr

func (data ImGuiWindowClass) handle() *C.ImGuiWindowClass {
	return (*C.ImGuiWindowClass)(unsafe.Pointer(data))
}

func (data ImGuiWindowClass) c() C.ImGuiWindowClass {
	return *(data.handle())
}

func newImGuiWindowClassFromC(cvalue C.ImGuiWindowClass) ImGuiWindowClass {
	return ImGuiWindowClass(unsafe.Pointer(&cvalue))
}

type ImFontConfig uintptr

func (data ImFontConfig) handle() *C.ImFontConfig {
	return (*C.ImFontConfig)(unsafe.Pointer(data))
}

func (data ImFontConfig) c() C.ImFontConfig {
	return *(data.handle())
}

func newImFontConfigFromC(cvalue C.ImFontConfig) ImFontConfig {
	return ImFontConfig(unsafe.Pointer(&cvalue))
}
