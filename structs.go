package cimgui

// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

type ImGuiInputEventMouseViewport uintptr

func (data ImGuiInputEventMouseViewport) handle() *C.ImGuiInputEventMouseViewport {
  return (*C.ImGuiInputEventMouseViewport)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMouseViewport) C() C.ImGuiInputEventMouseViewport {
  return *(data.handle())
}

type ImGuiInputTextCallbackData uintptr

func (data ImGuiInputTextCallbackData) handle() *C.ImGuiInputTextCallbackData {
  return (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(data))
}

func (data ImGuiInputTextCallbackData) C() C.ImGuiInputTextCallbackData {
  return *(data.handle())
}

type ImGuiPlatformMonitor uintptr

func (data ImGuiPlatformMonitor) handle() *C.ImGuiPlatformMonitor {
  return (*C.ImGuiPlatformMonitor)(unsafe.Pointer(data))
}

func (data ImGuiPlatformMonitor) C() C.ImGuiPlatformMonitor {
  return *(data.handle())
}

type ImGuiShrinkWidthItem uintptr

func (data ImGuiShrinkWidthItem) handle() *C.ImGuiShrinkWidthItem {
  return (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(data))
}

func (data ImGuiShrinkWidthItem) C() C.ImGuiShrinkWidthItem {
  return *(data.handle())
}

type ImGuiViewportP uintptr

func (data ImGuiViewportP) handle() *C.ImGuiViewportP {
  return (*C.ImGuiViewportP)(unsafe.Pointer(data))
}

func (data ImGuiViewportP) C() C.ImGuiViewportP {
  return *(data.handle())
}

type ImGuiWindowSettings uintptr

func (data ImGuiWindowSettings) handle() *C.ImGuiWindowSettings {
  return (*C.ImGuiWindowSettings)(unsafe.Pointer(data))
}

func (data ImGuiWindowSettings) C() C.ImGuiWindowSettings {
  return *(data.handle())
}

type ImGuiContextHook uintptr

func (data ImGuiContextHook) handle() *C.ImGuiContextHook {
  return (*C.ImGuiContextHook)(unsafe.Pointer(data))
}

func (data ImGuiContextHook) C() C.ImGuiContextHook {
  return *(data.handle())
}

type ImGuiMetricsConfig uintptr

func (data ImGuiMetricsConfig) handle() *C.ImGuiMetricsConfig {
  return (*C.ImGuiMetricsConfig)(unsafe.Pointer(data))
}

func (data ImGuiMetricsConfig) C() C.ImGuiMetricsConfig {
  return *(data.handle())
}

type ImGuiIO uintptr

func (data ImGuiIO) handle() *C.ImGuiIO {
  return (*C.ImGuiIO)(unsafe.Pointer(data))
}

func (data ImGuiIO) C() C.ImGuiIO {
  return *(data.handle())
}

type ImGuiPayload uintptr

func (data ImGuiPayload) handle() *C.ImGuiPayload {
  return (*C.ImGuiPayload)(unsafe.Pointer(data))
}

func (data ImGuiPayload) C() C.ImGuiPayload {
  return *(data.handle())
}

type ImGuiSizeCallbackData uintptr

func (data ImGuiSizeCallbackData) handle() *C.ImGuiSizeCallbackData {
  return (*C.ImGuiSizeCallbackData)(unsafe.Pointer(data))
}

func (data ImGuiSizeCallbackData) C() C.ImGuiSizeCallbackData {
  return *(data.handle())
}

type ImGuiWindowClass uintptr

func (data ImGuiWindowClass) handle() *C.ImGuiWindowClass {
  return (*C.ImGuiWindowClass)(unsafe.Pointer(data))
}

func (data ImGuiWindowClass) C() C.ImGuiWindowClass {
  return *(data.handle())
}

type ImGuiWindowDockStyle uintptr

func (data ImGuiWindowDockStyle) handle() *C.ImGuiWindowDockStyle {
  return (*C.ImGuiWindowDockStyle)(unsafe.Pointer(data))
}

func (data ImGuiWindowDockStyle) C() C.ImGuiWindowDockStyle {
  return *(data.handle())
}

type ImDrawVert uintptr

func (data ImDrawVert) handle() *C.ImDrawVert {
  return (*C.ImDrawVert)(unsafe.Pointer(data))
}

func (data ImDrawVert) C() C.ImDrawVert {
  return *(data.handle())
}

type ImGuiTableInstanceData uintptr

func (data ImGuiTableInstanceData) handle() *C.ImGuiTableInstanceData {
  return (*C.ImGuiTableInstanceData)(unsafe.Pointer(data))
}

func (data ImGuiTableInstanceData) C() C.ImGuiTableInstanceData {
  return *(data.handle())
}

type ImGuiWindowStackData uintptr

func (data ImGuiWindowStackData) handle() *C.ImGuiWindowStackData {
  return (*C.ImGuiWindowStackData)(unsafe.Pointer(data))
}

func (data ImGuiWindowStackData) C() C.ImGuiWindowStackData {
  return *(data.handle())
}

type ImFont uintptr

func (data ImFont) handle() *C.ImFont {
  return (*C.ImFont)(unsafe.Pointer(data))
}

func (data ImFont) C() C.ImFont {
  return *(data.handle())
}

type ImGuiColorMod uintptr

func (data ImGuiColorMod) handle() *C.ImGuiColorMod {
  return (*C.ImGuiColorMod)(unsafe.Pointer(data))
}

func (data ImGuiColorMod) C() C.ImGuiColorMod {
  return *(data.handle())
}

type ImGuiLastItemData uintptr

func (data ImGuiLastItemData) handle() *C.ImGuiLastItemData {
  return (*C.ImGuiLastItemData)(unsafe.Pointer(data))
}

func (data ImGuiLastItemData) C() C.ImGuiLastItemData {
  return *(data.handle())
}

type ImGuiTableCellData uintptr

func (data ImGuiTableCellData) handle() *C.ImGuiTableCellData {
  return (*C.ImGuiTableCellData)(unsafe.Pointer(data))
}

func (data ImGuiTableCellData) C() C.ImGuiTableCellData {
  return *(data.handle())
}

type ImGuiTableColumn uintptr

func (data ImGuiTableColumn) handle() *C.ImGuiTableColumn {
  return (*C.ImGuiTableColumn)(unsafe.Pointer(data))
}

func (data ImGuiTableColumn) C() C.ImGuiTableColumn {
  return *(data.handle())
}

type ImGuiTextBuffer uintptr

func (data ImGuiTextBuffer) handle() *C.ImGuiTextBuffer {
  return (*C.ImGuiTextBuffer)(unsafe.Pointer(data))
}

func (data ImGuiTextBuffer) C() C.ImGuiTextBuffer {
  return *(data.handle())
}

type ImDrawDataBuilder uintptr

func (data ImDrawDataBuilder) handle() *C.ImDrawDataBuilder {
  return (*C.ImDrawDataBuilder)(unsafe.Pointer(data))
}

func (data ImDrawDataBuilder) C() C.ImDrawDataBuilder {
  return *(data.handle())
}

type ImFontConfig uintptr

func (data ImFontConfig) handle() *C.ImFontConfig {
  return (*C.ImFontConfig)(unsafe.Pointer(data))
}

func (data ImFontConfig) C() C.ImFontConfig {
  return *(data.handle())
}

type ImGuiDockContext uintptr

func (data ImGuiDockContext) handle() *C.ImGuiDockContext {
  return (*C.ImGuiDockContext)(unsafe.Pointer(data))
}

func (data ImGuiDockContext) C() C.ImGuiDockContext {
  return *(data.handle())
}

type ImGuiListClipperRange uintptr

func (data ImGuiListClipperRange) handle() *C.ImGuiListClipperRange {
  return (*C.ImGuiListClipperRange)(unsafe.Pointer(data))
}

func (data ImGuiListClipperRange) C() C.ImGuiListClipperRange {
  return *(data.handle())
}

type ImGuiOnceUponAFrame uintptr

func (data ImGuiOnceUponAFrame) handle() *C.ImGuiOnceUponAFrame {
  return (*C.ImGuiOnceUponAFrame)(unsafe.Pointer(data))
}

func (data ImGuiOnceUponAFrame) C() C.ImGuiOnceUponAFrame {
  return *(data.handle())
}

type ImGuiTabBar uintptr

func (data ImGuiTabBar) handle() *C.ImGuiTabBar {
  return (*C.ImGuiTabBar)(unsafe.Pointer(data))
}

func (data ImGuiTabBar) C() C.ImGuiTabBar {
  return *(data.handle())
}

type ImGuiTextRange uintptr

func (data ImGuiTextRange) handle() *C.ImGuiTextRange {
  return (*C.ImGuiTextRange)(unsafe.Pointer(data))
}

func (data ImGuiTextRange) C() C.ImGuiTextRange {
  return *(data.handle())
}

type ImFontBuilderIO uintptr

func (data ImFontBuilderIO) handle() *C.ImFontBuilderIO {
  return (*C.ImFontBuilderIO)(unsafe.Pointer(data))
}

func (data ImFontBuilderIO) C() C.ImFontBuilderIO {
  return *(data.handle())
}

type ImGuiTableColumnSettings uintptr

func (data ImGuiTableColumnSettings) handle() *C.ImGuiTableColumnSettings {
  return (*C.ImGuiTableColumnSettings)(unsafe.Pointer(data))
}

func (data ImGuiTableColumnSettings) C() C.ImGuiTableColumnSettings {
  return *(data.handle())
}

type ImGuiTableColumnSortSpecs uintptr

func (data ImGuiTableColumnSortSpecs) handle() *C.ImGuiTableColumnSortSpecs {
  return (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(data))
}

func (data ImGuiTableColumnSortSpecs) C() C.ImGuiTableColumnSortSpecs {
  return *(data.handle())
}

type ImGuiTextFilter uintptr

func (data ImGuiTextFilter) handle() *C.ImGuiTextFilter {
  return (*C.ImGuiTextFilter)(unsafe.Pointer(data))
}

func (data ImGuiTextFilter) C() C.ImGuiTextFilter {
  return *(data.handle())
}

type ImDrawListSharedData uintptr

func (data ImDrawListSharedData) handle() *C.ImDrawListSharedData {
  return (*C.ImDrawListSharedData)(unsafe.Pointer(data))
}

func (data ImDrawListSharedData) C() C.ImDrawListSharedData {
  return *(data.handle())
}

type ImGuiDataTypeTempStorage uintptr

func (data ImGuiDataTypeTempStorage) handle() *C.ImGuiDataTypeTempStorage {
  return (*C.ImGuiDataTypeTempStorage)(unsafe.Pointer(data))
}

func (data ImGuiDataTypeTempStorage) C() C.ImGuiDataTypeTempStorage {
  return *(data.handle())
}

type ImGuiInputTextState uintptr

func (data ImGuiInputTextState) handle() *C.ImGuiInputTextState {
  return (*C.ImGuiInputTextState)(unsafe.Pointer(data))
}

func (data ImGuiInputTextState) C() C.ImGuiInputTextState {
  return *(data.handle())
}

type ImGuiKeyData uintptr

func (data ImGuiKeyData) handle() *C.ImGuiKeyData {
  return (*C.ImGuiKeyData)(unsafe.Pointer(data))
}

func (data ImGuiKeyData) C() C.ImGuiKeyData {
  return *(data.handle())
}

type ImGuiNextItemData uintptr

func (data ImGuiNextItemData) handle() *C.ImGuiNextItemData {
  return (*C.ImGuiNextItemData)(unsafe.Pointer(data))
}

func (data ImGuiNextItemData) C() C.ImGuiNextItemData {
  return *(data.handle())
}

type ImGuiNextWindowData uintptr

func (data ImGuiNextWindowData) handle() *C.ImGuiNextWindowData {
  return (*C.ImGuiNextWindowData)(unsafe.Pointer(data))
}

func (data ImGuiNextWindowData) C() C.ImGuiNextWindowData {
  return *(data.handle())
}

type ImGuiWindow uintptr

func (data ImGuiWindow) handle() *C.ImGuiWindow {
  return (*C.ImGuiWindow)(unsafe.Pointer(data))
}

func (data ImGuiWindow) C() C.ImGuiWindow {
  return *(data.handle())
}

type ImGuiDataTypeInfo uintptr

func (data ImGuiDataTypeInfo) handle() *C.ImGuiDataTypeInfo {
  return (*C.ImGuiDataTypeInfo)(unsafe.Pointer(data))
}

func (data ImGuiDataTypeInfo) C() C.ImGuiDataTypeInfo {
  return *(data.handle())
}

type ImFontAtlasCustomRect uintptr

func (data ImFontAtlasCustomRect) handle() *C.ImFontAtlasCustomRect {
  return (*C.ImFontAtlasCustomRect)(unsafe.Pointer(data))
}

func (data ImFontAtlasCustomRect) C() C.ImFontAtlasCustomRect {
  return *(data.handle())
}

type ImFontGlyph uintptr

func (data ImFontGlyph) handle() *C.ImFontGlyph {
  return (*C.ImFontGlyph)(unsafe.Pointer(data))
}

func (data ImFontGlyph) C() C.ImFontGlyph {
  return *(data.handle())
}

type ImGuiOldColumns uintptr

func (data ImGuiOldColumns) handle() *C.ImGuiOldColumns {
  return (*C.ImGuiOldColumns)(unsafe.Pointer(data))
}

func (data ImGuiOldColumns) C() C.ImGuiOldColumns {
  return *(data.handle())
}

type ImGuiTableTempData uintptr

func (data ImGuiTableTempData) handle() *C.ImGuiTableTempData {
  return (*C.ImGuiTableTempData)(unsafe.Pointer(data))
}

func (data ImGuiTableTempData) C() C.ImGuiTableTempData {
  return *(data.handle())
}

type ImDrawChannel uintptr

func (data ImDrawChannel) handle() *C.ImDrawChannel {
  return (*C.ImDrawChannel)(unsafe.Pointer(data))
}

func (data ImDrawChannel) C() C.ImDrawChannel {
  return *(data.handle())
}

type ImDrawCmd uintptr

func (data ImDrawCmd) handle() *C.ImDrawCmd {
  return (*C.ImDrawCmd)(unsafe.Pointer(data))
}

func (data ImDrawCmd) C() C.ImDrawCmd {
  return *(data.handle())
}

type ImDrawCmdHeader uintptr

func (data ImDrawCmdHeader) handle() *C.ImDrawCmdHeader {
  return (*C.ImDrawCmdHeader)(unsafe.Pointer(data))
}

func (data ImDrawCmdHeader) C() C.ImDrawCmdHeader {
  return *(data.handle())
}

type ImFontAtlas uintptr

func (data ImFontAtlas) handle() *C.ImFontAtlas {
  return (*C.ImFontAtlas)(unsafe.Pointer(data))
}

func (data ImFontAtlas) C() C.ImFontAtlas {
  return *(data.handle())
}

type ImGuiInputEventMouseWheel uintptr

func (data ImGuiInputEventMouseWheel) handle() *C.ImGuiInputEventMouseWheel {
  return (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMouseWheel) C() C.ImGuiInputEventMouseWheel {
  return *(data.handle())
}

type ImGuiPtrOrIndex uintptr

func (data ImGuiPtrOrIndex) handle() *C.ImGuiPtrOrIndex {
  return (*C.ImGuiPtrOrIndex)(unsafe.Pointer(data))
}

func (data ImGuiPtrOrIndex) C() C.ImGuiPtrOrIndex {
  return *(data.handle())
}

type ImGuiStackSizes uintptr

func (data ImGuiStackSizes) handle() *C.ImGuiStackSizes {
  return (*C.ImGuiStackSizes)(unsafe.Pointer(data))
}

func (data ImGuiStackSizes) C() C.ImGuiStackSizes {
  return *(data.handle())
}

type ImGuiStackTool uintptr

func (data ImGuiStackTool) handle() *C.ImGuiStackTool {
  return (*C.ImGuiStackTool)(unsafe.Pointer(data))
}

func (data ImGuiStackTool) C() C.ImGuiStackTool {
  return *(data.handle())
}

type ImBitVector uintptr

func (data ImBitVector) handle() *C.ImBitVector {
  return (*C.ImBitVector)(unsafe.Pointer(data))
}

func (data ImBitVector) C() C.ImBitVector {
  return *(data.handle())
}

type ImGuiTabItem uintptr

func (data ImGuiTabItem) handle() *C.ImGuiTabItem {
  return (*C.ImGuiTabItem)(unsafe.Pointer(data))
}

func (data ImGuiTabItem) C() C.ImGuiTabItem {
  return *(data.handle())
}

type ImGuiStyleMod uintptr

func (data ImGuiStyleMod) handle() *C.ImGuiStyleMod {
  return (*C.ImGuiStyleMod)(unsafe.Pointer(data))
}

func (data ImGuiStyleMod) C() C.ImGuiStyleMod {
  return *(data.handle())
}

type ImDrawListSplitter uintptr

func (data ImDrawListSplitter) handle() *C.ImDrawListSplitter {
  return (*C.ImDrawListSplitter)(unsafe.Pointer(data))
}

func (data ImDrawListSplitter) C() C.ImDrawListSplitter {
  return *(data.handle())
}

type ImGuiListClipper uintptr

func (data ImGuiListClipper) handle() *C.ImGuiListClipper {
  return (*C.ImGuiListClipper)(unsafe.Pointer(data))
}

func (data ImGuiListClipper) C() C.ImGuiListClipper {
  return *(data.handle())
}

type ImGuiStyle uintptr

func (data ImGuiStyle) handle() *C.ImGuiStyle {
  return (*C.ImGuiStyle)(unsafe.Pointer(data))
}

func (data ImGuiStyle) C() C.ImGuiStyle {
  return *(data.handle())
}

type ImGuiViewport uintptr

func (data ImGuiViewport) handle() *C.ImGuiViewport {
  return (*C.ImGuiViewport)(unsafe.Pointer(data))
}

func (data ImGuiViewport) C() C.ImGuiViewport {
  return *(data.handle())
}

type ImDrawList uintptr

func (data ImDrawList) handle() *C.ImDrawList {
  return (*C.ImDrawList)(unsafe.Pointer(data))
}

func (data ImDrawList) C() C.ImDrawList {
  return *(data.handle())
}

type ImGuiDockNode uintptr

func (data ImGuiDockNode) handle() *C.ImGuiDockNode {
  return (*C.ImGuiDockNode)(unsafe.Pointer(data))
}

func (data ImGuiDockNode) C() C.ImGuiDockNode {
  return *(data.handle())
}

type ImGuiInputEvent uintptr

func (data ImGuiInputEvent) handle() *C.ImGuiInputEvent {
  return (*C.ImGuiInputEvent)(unsafe.Pointer(data))
}

func (data ImGuiInputEvent) C() C.ImGuiInputEvent {
  return *(data.handle())
}

type ImGuiInputEventMouseButton uintptr

func (data ImGuiInputEventMouseButton) handle() *C.ImGuiInputEventMouseButton {
  return (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMouseButton) C() C.ImGuiInputEventMouseButton {
  return *(data.handle())
}

type ImGuiInputEventText uintptr

func (data ImGuiInputEventText) handle() *C.ImGuiInputEventText {
  return (*C.ImGuiInputEventText)(unsafe.Pointer(data))
}

func (data ImGuiInputEventText) C() C.ImGuiInputEventText {
  return *(data.handle())
}

type ImGuiOldColumnData uintptr

func (data ImGuiOldColumnData) handle() *C.ImGuiOldColumnData {
  return (*C.ImGuiOldColumnData)(unsafe.Pointer(data))
}

func (data ImGuiOldColumnData) C() C.ImGuiOldColumnData {
  return *(data.handle())
}

type ImGuiSettingsHandler uintptr

func (data ImGuiSettingsHandler) handle() *C.ImGuiSettingsHandler {
  return (*C.ImGuiSettingsHandler)(unsafe.Pointer(data))
}

func (data ImGuiSettingsHandler) C() C.ImGuiSettingsHandler {
  return *(data.handle())
}

type ImGuiTableSettings uintptr

func (data ImGuiTableSettings) handle() *C.ImGuiTableSettings {
  return (*C.ImGuiTableSettings)(unsafe.Pointer(data))
}

func (data ImGuiTableSettings) C() C.ImGuiTableSettings {
  return *(data.handle())
}

type ImFontGlyphRangesBuilder uintptr

func (data ImFontGlyphRangesBuilder) handle() *C.ImFontGlyphRangesBuilder {
  return (*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(data))
}

func (data ImFontGlyphRangesBuilder) C() C.ImFontGlyphRangesBuilder {
  return *(data.handle())
}

type ImGuiTableSortSpecs uintptr

func (data ImGuiTableSortSpecs) handle() *C.ImGuiTableSortSpecs {
  return (*C.ImGuiTableSortSpecs)(unsafe.Pointer(data))
}

func (data ImGuiTableSortSpecs) C() C.ImGuiTableSortSpecs {
  return *(data.handle())
}

type ImGuiInputEventMousePos uintptr

func (data ImGuiInputEventMousePos) handle() *C.ImGuiInputEventMousePos {
  return (*C.ImGuiInputEventMousePos)(unsafe.Pointer(data))
}

func (data ImGuiInputEventMousePos) C() C.ImGuiInputEventMousePos {
  return *(data.handle())
}

type ImGuiNavItemData uintptr

func (data ImGuiNavItemData) handle() *C.ImGuiNavItemData {
  return (*C.ImGuiNavItemData)(unsafe.Pointer(data))
}

func (data ImGuiNavItemData) C() C.ImGuiNavItemData {
  return *(data.handle())
}

type ImGuiStorage uintptr

func (data ImGuiStorage) handle() *C.ImGuiStorage {
  return (*C.ImGuiStorage)(unsafe.Pointer(data))
}

func (data ImGuiStorage) C() C.ImGuiStorage {
  return *(data.handle())
}

type ImDrawData uintptr

func (data ImDrawData) handle() *C.ImDrawData {
  return (*C.ImDrawData)(unsafe.Pointer(data))
}

func (data ImDrawData) C() C.ImDrawData {
  return *(data.handle())
}

type ImGuiInputEventAppFocused uintptr

func (data ImGuiInputEventAppFocused) handle() *C.ImGuiInputEventAppFocused {
  return (*C.ImGuiInputEventAppFocused)(unsafe.Pointer(data))
}

func (data ImGuiInputEventAppFocused) C() C.ImGuiInputEventAppFocused {
  return *(data.handle())
}

type ImGuiListClipperData uintptr

func (data ImGuiListClipperData) handle() *C.ImGuiListClipperData {
  return (*C.ImGuiListClipperData)(unsafe.Pointer(data))
}

func (data ImGuiListClipperData) C() C.ImGuiListClipperData {
  return *(data.handle())
}

type ImGuiMenuColumns uintptr

func (data ImGuiMenuColumns) handle() *C.ImGuiMenuColumns {
  return (*C.ImGuiMenuColumns)(unsafe.Pointer(data))
}

func (data ImGuiMenuColumns) C() C.ImGuiMenuColumns {
  return *(data.handle())
}

type ImGuiPlatformIO uintptr

func (data ImGuiPlatformIO) handle() *C.ImGuiPlatformIO {
  return (*C.ImGuiPlatformIO)(unsafe.Pointer(data))
}

func (data ImGuiPlatformIO) C() C.ImGuiPlatformIO {
  return *(data.handle())
}

type ImGuiPopupData uintptr

func (data ImGuiPopupData) handle() *C.ImGuiPopupData {
  return (*C.ImGuiPopupData)(unsafe.Pointer(data))
}

func (data ImGuiPopupData) C() C.ImGuiPopupData {
  return *(data.handle())
}

type ImGuiStoragePair uintptr

func (data ImGuiStoragePair) handle() *C.ImGuiStoragePair {
  return (*C.ImGuiStoragePair)(unsafe.Pointer(data))
}

func (data ImGuiStoragePair) C() C.ImGuiStoragePair {
  return *(data.handle())
}

type ImGuiWindowTempData uintptr

func (data ImGuiWindowTempData) handle() *C.ImGuiWindowTempData {
  return (*C.ImGuiWindowTempData)(unsafe.Pointer(data))
}

func (data ImGuiWindowTempData) C() C.ImGuiWindowTempData {
  return *(data.handle())
}

type ImGuiComboPreviewData uintptr

func (data ImGuiComboPreviewData) handle() *C.ImGuiComboPreviewData {
  return (*C.ImGuiComboPreviewData)(unsafe.Pointer(data))
}

func (data ImGuiComboPreviewData) C() C.ImGuiComboPreviewData {
  return *(data.handle())
}

type ImGuiPlatformImeData uintptr

func (data ImGuiPlatformImeData) handle() *C.ImGuiPlatformImeData {
  return (*C.ImGuiPlatformImeData)(unsafe.Pointer(data))
}

func (data ImGuiPlatformImeData) C() C.ImGuiPlatformImeData {
  return *(data.handle())
}

type ImGuiStackLevelInfo uintptr

func (data ImGuiStackLevelInfo) handle() *C.ImGuiStackLevelInfo {
  return (*C.ImGuiStackLevelInfo)(unsafe.Pointer(data))
}

func (data ImGuiStackLevelInfo) C() C.ImGuiStackLevelInfo {
  return *(data.handle())
}

type ImGuiGroupData uintptr

func (data ImGuiGroupData) handle() *C.ImGuiGroupData {
  return (*C.ImGuiGroupData)(unsafe.Pointer(data))
}

func (data ImGuiGroupData) C() C.ImGuiGroupData {
  return *(data.handle())
}

type ImGuiContext uintptr

func (data ImGuiContext) handle() *C.ImGuiContext {
  return (*C.ImGuiContext)(unsafe.Pointer(data))
}

func (data ImGuiContext) C() C.ImGuiContext {
  return *(data.handle())
}

type ImGuiInputEventKey uintptr

func (data ImGuiInputEventKey) handle() *C.ImGuiInputEventKey {
  return (*C.ImGuiInputEventKey)(unsafe.Pointer(data))
}

func (data ImGuiInputEventKey) C() C.ImGuiInputEventKey {
  return *(data.handle())
}

type ImGuiTable uintptr

func (data ImGuiTable) handle() *C.ImGuiTable {
  return (*C.ImGuiTable)(unsafe.Pointer(data))
}

func (data ImGuiTable) C() C.ImGuiTable {
  return *(data.handle())
}

