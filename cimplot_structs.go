package cimgui

// #include "cimplot_wrapper.h"
import "C"
import "unsafe"

type ImPlotContext uintptr

func (data ImPlotContext) handle() *C.ImPlotContext {
	return (*C.ImPlotContext)(unsafe.Pointer(data))
}

func (data ImPlotContext) c() C.ImPlotContext {
	return *(data.handle())
}

func newImPlotContextFromC(cvalue C.ImPlotContext) ImPlotContext {
	return ImPlotContext(unsafe.Pointer(&cvalue))
}

type ImPlotInputMap uintptr

func (data ImPlotInputMap) handle() *C.ImPlotInputMap {
	return (*C.ImPlotInputMap)(unsafe.Pointer(data))
}

func (data ImPlotInputMap) c() C.ImPlotInputMap {
	return *(data.handle())
}

func newImPlotInputMapFromC(cvalue C.ImPlotInputMap) ImPlotInputMap {
	return ImPlotInputMap(unsafe.Pointer(&cvalue))
}

type ImPlotNextItemData uintptr

func (data ImPlotNextItemData) handle() *C.ImPlotNextItemData {
	return (*C.ImPlotNextItemData)(unsafe.Pointer(data))
}

func (data ImPlotNextItemData) c() C.ImPlotNextItemData {
	return *(data.handle())
}

func newImPlotNextItemDataFromC(cvalue C.ImPlotNextItemData) ImPlotNextItemData {
	return ImPlotNextItemData(unsafe.Pointer(&cvalue))
}

type ImPlotRect uintptr

func (data ImPlotRect) handle() *C.ImPlotRect {
	return (*C.ImPlotRect)(unsafe.Pointer(data))
}

func (data ImPlotRect) c() C.ImPlotRect {
	return *(data.handle())
}

func newImPlotRectFromC(cvalue C.ImPlotRect) ImPlotRect {
	return ImPlotRect(unsafe.Pointer(&cvalue))
}

type ImPlotSubplot uintptr

func (data ImPlotSubplot) handle() *C.ImPlotSubplot {
	return (*C.ImPlotSubplot)(unsafe.Pointer(data))
}

func (data ImPlotSubplot) c() C.ImPlotSubplot {
	return *(data.handle())
}

func newImPlotSubplotFromC(cvalue C.ImPlotSubplot) ImPlotSubplot {
	return ImPlotSubplot(unsafe.Pointer(&cvalue))
}

type ImPlotAnnotation uintptr

func (data ImPlotAnnotation) handle() *C.ImPlotAnnotation {
	return (*C.ImPlotAnnotation)(unsafe.Pointer(data))
}

func (data ImPlotAnnotation) c() C.ImPlotAnnotation {
	return *(data.handle())
}

func newImPlotAnnotationFromC(cvalue C.ImPlotAnnotation) ImPlotAnnotation {
	return ImPlotAnnotation(unsafe.Pointer(&cvalue))
}

type ImPlotAxis uintptr

func (data ImPlotAxis) handle() *C.ImPlotAxis {
	return (*C.ImPlotAxis)(unsafe.Pointer(data))
}

func (data ImPlotAxis) c() C.ImPlotAxis {
	return *(data.handle())
}

func newImPlotAxisFromC(cvalue C.ImPlotAxis) ImPlotAxis {
	return ImPlotAxis(unsafe.Pointer(&cvalue))
}

type ImPlotTicker uintptr

func (data ImPlotTicker) handle() *C.ImPlotTicker {
	return (*C.ImPlotTicker)(unsafe.Pointer(data))
}

func (data ImPlotTicker) c() C.ImPlotTicker {
	return *(data.handle())
}

func newImPlotTickerFromC(cvalue C.ImPlotTicker) ImPlotTicker {
	return ImPlotTicker(unsafe.Pointer(&cvalue))
}

type ImPlotNextPlotData uintptr

func (data ImPlotNextPlotData) handle() *C.ImPlotNextPlotData {
	return (*C.ImPlotNextPlotData)(unsafe.Pointer(data))
}

func (data ImPlotNextPlotData) c() C.ImPlotNextPlotData {
	return *(data.handle())
}

func newImPlotNextPlotDataFromC(cvalue C.ImPlotNextPlotData) ImPlotNextPlotData {
	return ImPlotNextPlotData(unsafe.Pointer(&cvalue))
}

type ImPlotStyle uintptr

func (data ImPlotStyle) handle() *C.ImPlotStyle {
	return (*C.ImPlotStyle)(unsafe.Pointer(data))
}

func (data ImPlotStyle) c() C.ImPlotStyle {
	return *(data.handle())
}

func newImPlotStyleFromC(cvalue C.ImPlotStyle) ImPlotStyle {
	return ImPlotStyle(unsafe.Pointer(&cvalue))
}

type ImPlotTag uintptr

func (data ImPlotTag) handle() *C.ImPlotTag {
	return (*C.ImPlotTag)(unsafe.Pointer(data))
}

func (data ImPlotTag) c() C.ImPlotTag {
	return *(data.handle())
}

func newImPlotTagFromC(cvalue C.ImPlotTag) ImPlotTag {
	return ImPlotTag(unsafe.Pointer(&cvalue))
}

type ImPlotColormapData uintptr

func (data ImPlotColormapData) handle() *C.ImPlotColormapData {
	return (*C.ImPlotColormapData)(unsafe.Pointer(data))
}

func (data ImPlotColormapData) c() C.ImPlotColormapData {
	return *(data.handle())
}

func newImPlotColormapDataFromC(cvalue C.ImPlotColormapData) ImPlotColormapData {
	return ImPlotColormapData(unsafe.Pointer(&cvalue))
}

type ImPlotDateTimeSpec uintptr

func (data ImPlotDateTimeSpec) handle() *C.ImPlotDateTimeSpec {
	return (*C.ImPlotDateTimeSpec)(unsafe.Pointer(data))
}

func (data ImPlotDateTimeSpec) c() C.ImPlotDateTimeSpec {
	return *(data.handle())
}

func newImPlotDateTimeSpecFromC(cvalue C.ImPlotDateTimeSpec) ImPlotDateTimeSpec {
	return ImPlotDateTimeSpec(unsafe.Pointer(&cvalue))
}

type ImPlotItemGroup uintptr

func (data ImPlotItemGroup) handle() *C.ImPlotItemGroup {
	return (*C.ImPlotItemGroup)(unsafe.Pointer(data))
}

func (data ImPlotItemGroup) c() C.ImPlotItemGroup {
	return *(data.handle())
}

func newImPlotItemGroupFromC(cvalue C.ImPlotItemGroup) ImPlotItemGroup {
	return ImPlotItemGroup(unsafe.Pointer(&cvalue))
}

type ImPlotTime uintptr

func (data ImPlotTime) handle() *C.ImPlotTime {
	return (*C.ImPlotTime)(unsafe.Pointer(data))
}

func (data ImPlotTime) c() C.ImPlotTime {
	return *(data.handle())
}

func newImPlotTimeFromC(cvalue C.ImPlotTime) ImPlotTime {
	return ImPlotTime(unsafe.Pointer(&cvalue))
}

type ImPlotLegend uintptr

func (data ImPlotLegend) handle() *C.ImPlotLegend {
	return (*C.ImPlotLegend)(unsafe.Pointer(data))
}

func (data ImPlotLegend) c() C.ImPlotLegend {
	return *(data.handle())
}

func newImPlotLegendFromC(cvalue C.ImPlotLegend) ImPlotLegend {
	return ImPlotLegend(unsafe.Pointer(&cvalue))
}

type ImPlotTagCollection uintptr

func (data ImPlotTagCollection) handle() *C.ImPlotTagCollection {
	return (*C.ImPlotTagCollection)(unsafe.Pointer(data))
}

func (data ImPlotTagCollection) c() C.ImPlotTagCollection {
	return *(data.handle())
}

func newImPlotTagCollectionFromC(cvalue C.ImPlotTagCollection) ImPlotTagCollection {
	return ImPlotTagCollection(unsafe.Pointer(&cvalue))
}

type ImPlotTick uintptr

func (data ImPlotTick) handle() *C.ImPlotTick {
	return (*C.ImPlotTick)(unsafe.Pointer(data))
}

func (data ImPlotTick) c() C.ImPlotTick {
	return *(data.handle())
}

func newImPlotTickFromC(cvalue C.ImPlotTick) ImPlotTick {
	return ImPlotTick(unsafe.Pointer(&cvalue))
}

type ImPlotPlot uintptr

func (data ImPlotPlot) handle() *C.ImPlotPlot {
	return (*C.ImPlotPlot)(unsafe.Pointer(data))
}

func (data ImPlotPlot) c() C.ImPlotPlot {
	return *(data.handle())
}

func newImPlotPlotFromC(cvalue C.ImPlotPlot) ImPlotPlot {
	return ImPlotPlot(unsafe.Pointer(&cvalue))
}

type ImPlotPoint uintptr

func (data ImPlotPoint) handle() *C.ImPlotPoint {
	return (*C.ImPlotPoint)(unsafe.Pointer(data))
}

func (data ImPlotPoint) c() C.ImPlotPoint {
	return *(data.handle())
}

func newImPlotPointFromC(cvalue C.ImPlotPoint) ImPlotPoint {
	return ImPlotPoint(unsafe.Pointer(&cvalue))
}

type ImPlotPointError uintptr

func (data ImPlotPointError) handle() *C.ImPlotPointError {
	return (*C.ImPlotPointError)(unsafe.Pointer(data))
}

func (data ImPlotPointError) c() C.ImPlotPointError {
	return *(data.handle())
}

func newImPlotPointErrorFromC(cvalue C.ImPlotPointError) ImPlotPointError {
	return ImPlotPointError(unsafe.Pointer(&cvalue))
}

type ImPlotRange uintptr

func (data ImPlotRange) handle() *C.ImPlotRange {
	return (*C.ImPlotRange)(unsafe.Pointer(data))
}

func (data ImPlotRange) c() C.ImPlotRange {
	return *(data.handle())
}

func newImPlotRangeFromC(cvalue C.ImPlotRange) ImPlotRange {
	return ImPlotRange(unsafe.Pointer(&cvalue))
}

type ImPlotAlignmentData uintptr

func (data ImPlotAlignmentData) handle() *C.ImPlotAlignmentData {
	return (*C.ImPlotAlignmentData)(unsafe.Pointer(data))
}

func (data ImPlotAlignmentData) c() C.ImPlotAlignmentData {
	return *(data.handle())
}

func newImPlotAlignmentDataFromC(cvalue C.ImPlotAlignmentData) ImPlotAlignmentData {
	return ImPlotAlignmentData(unsafe.Pointer(&cvalue))
}

type ImPlotAnnotationCollection uintptr

func (data ImPlotAnnotationCollection) handle() *C.ImPlotAnnotationCollection {
	return (*C.ImPlotAnnotationCollection)(unsafe.Pointer(data))
}

func (data ImPlotAnnotationCollection) c() C.ImPlotAnnotationCollection {
	return *(data.handle())
}

func newImPlotAnnotationCollectionFromC(cvalue C.ImPlotAnnotationCollection) ImPlotAnnotationCollection {
	return ImPlotAnnotationCollection(unsafe.Pointer(&cvalue))
}

type ImPlotItem uintptr

func (data ImPlotItem) handle() *C.ImPlotItem {
	return (*C.ImPlotItem)(unsafe.Pointer(data))
}

func (data ImPlotItem) c() C.ImPlotItem {
	return *(data.handle())
}

func newImPlotItemFromC(cvalue C.ImPlotItem) ImPlotItem {
	return ImPlotItem(unsafe.Pointer(&cvalue))
}
