package main

import "C"
import "fmt"

type typeWrapper func(arg ArgDef) (argType string, def string, varName string)

func constCharW(arg ArgDef) (argType string, def string, varName string) {
	argType = "string"
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrapString(%[1]s)
defer %[1]sFin()`, arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func charPtrPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]string"
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrapStringList(%[1]s)
defer %[1]sFin()`, arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func ucharW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint", "C.uchar")
}

func uCharPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*C.uchar"
	varName = fmt.Sprintf("&%s", arg.Name)
	return
}

func sizeTW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint64", "C.xlong")
}

func sizeTPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*uint64"
	varName = fmt.Sprintf("(*C.xlong)(%s)", arg.Name)
	return
}

func floatW(arg ArgDef) (argType string, def string, varName string) {
	argType = "float32"
	varName = fmt.Sprintf("C.float(%s)", arg.Name)
	return
}

func floatPtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "float32", "C.float")
}

func floatArrayW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]float32"
	varName = fmt.Sprintf("(*C.float)(&(%s[0]))", arg.Name)
	return
}

func boolW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "bool", "C.bool")
}

func boolPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*bool"
	def = fmt.Sprintf("%[1]sArg, %[1]sFin := wrapBool(%[1]s)\ndefer %[1]sFin()", arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func shortW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "int", "C.short")
}

func ushortW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint", "C.ushort")
}

func uintW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint32", "C.uint")
}

func uintPtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "uint32", "C.uint")
}

func u8W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint", "C.ImU8")
}

func u8PtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "byte", "C.ImU8")
}

func u8SliceW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrSliceW("C.ImU8", "byte", arg)
}

func u16W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint", "C.ImU16")
}

func u16PtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "uint16", "C.uint")
}

func u16SliceW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrSliceW("C.ImU16", "uint16", arg)
}

func u32W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint32", "C.ImU32")
}

func u32PtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "uint16", "C.ImU32")
}

func u32SliceW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrSliceW("C.ImU32", "uint32", arg)
}

func u64W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "uint64", "C.ImU64")
}

func int2W(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrArrayW(2, "C.int", "int32", arg)
}

func int3W(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrArrayW(3, "C.int", "int32", arg)
}

func int4W(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrArrayW(4, "C.int", "int32", arg)
}

func s8W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "int", "C.ImS8")
}

func s8SliceW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrSliceW("C.ImS8", "int8", arg)
}

func s16W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "int", "C.ImS16")
}

func s16SliceW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrSliceW("C.ImS16", "int", arg)
}

func s32W(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "int", "C.ImS32")
}

func s32SliceW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrSliceW("C.ImS32", "int32", arg)
}

func float2W(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrArrayW(2, "C.float", "float32", arg)
}

func float3W(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrArrayW(3, "C.float", "float32", arg)
}

func float4W(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrArrayW(4, "C.float", "float32", arg)
}

func imWcharW(arg ArgDef) (argType string, def string, varName string) {
	argType = "Wchar"
	varName = fmt.Sprintf("C.ImWchar(%s)", arg.Name)
	return
}

func imWcharPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*Wchar"
	varName = fmt.Sprintf("(*C.ImWchar)(%s)", arg.Name)
	return
}

func intW(arg ArgDef) (argType string, def string, varName string) {
	argType = "int32"
	varName = fmt.Sprintf("C.int(%s)", arg.Name)
	return
}

func int8PtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "int8", "C.int")
}

func int16PtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "int16", "C.int")
}

func intPtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "int32", "C.int")
}

func int64ArrayW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]int64"
	varName = fmt.Sprintf("(*C.longlong)(&(%s[0]))", arg.Name)
	return
}

func uint64ArrayW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]uint64"
	varName = fmt.Sprintf("(*C.ulonglong)(&(%s[0]))", arg.Name)
	return
}

func doubleW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "float64", "C.double")
}

func doublePtrW(arg ArgDef) (argType string, def string, varName string) {
	return simplePtrW(arg.Name, "float64", "C.double")
}

func imGuiIDW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "ImGuiID", "C.ImGuiID")
}

func imTextureIDW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "TextureID", "C.ImTextureID")
}

func imDrawIdxW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "DrawIdx", "C.ImDrawIdx")
}

func imTableColumnIdxW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "TableColumnIdx", "C.ImGuiTableColumnIdx")
}

func imTableDrawChannelIdxW(arg ArgDef) (argType string, def string, varName string) {
	return simpleW(arg.Name, "TableDrawChannelIdx", "C.ImGuiTableDrawChannelIdx")
}

func voidPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "unsafe.Pointer"
	varName = arg.Name
	return
}

func imPlotPointW(arg ArgDef) (argType string, def string, varName string) {
	return wrappableW(arg.Name, "PlotPoint")
}

func imPlotPointPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrappablePtrW(arg.Name, "*PlotPoint", "C.ImPlotPoint")
}

func imVec2W(arg ArgDef) (argType string, def string, varName string) {
	return wrappableW(arg.Name, "Vec2")
}

func imVec2PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrappablePtrW(arg.Name, "*Vec2", "C.ImVec2")
}

// ImVec2[2] -> [2]ImVec2
func imVec22W(arg ArgDef) (argType string, def string, varName string) {
	return wrappablePtrArrayW(2, "C.ImVec2", "Vec2", arg)
}

func imVec4W(arg ArgDef) (argType string, def string, varName string) {
	return wrappableW(arg.Name, "Vec4")
}

func imVec4PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrappablePtrW(arg.Name, "*Vec4", "C.ImVec4")
}

func imRectW(arg ArgDef) (argType string, def string, varName string) {
	return wrappableW(arg.Name, "Rect")
}

func imRectPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrappablePtrW(arg.Name, "*Rect", "C.ImRect")
}

func imColorPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrappablePtrW(arg.Name, "*Color", "C.ImColor")
}

func inputeTextCallbackW(arg ArgDef) (argType string, def string, varName string) {
	argType = "ImGuiInputTextCallback"
	// TODO: implement me
	return
}

// generic wrappers:

// C.int -> int32
func simpleW(argName, goType, cType string) (argType string, def string, varName string) {
	argType = goType
	varName = fmt.Sprintf("%s(%s)", cType, argName)
	return
}

// C.int* -> *int32
func simplePtrW(argName, goType, cType string) (argType, def, varName string) {
	argType = fmt.Sprintf("*%s", goType)
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrapNumberPtr[%[2]s, %[3]s](%[1]s)
defer %[1]sFin()`, argName, cType, goType)
	varName = fmt.Sprintf("%sArg", argName)
	return
}

// C.int*, C.int[] as well as C.int[2] -> [2]*int32
func simplePtrArrayW(size int, cArrayType, goArrayType string, arg ArgDef) (argType string, def string, varName string) {
	argType = fmt.Sprintf("[%d]*%s", size, goArrayType)
	def = fmt.Sprintf(`%[1]sArg := make([]%[2]s, len(%[1]s))
for i, %[1]sV := range %[1]s {
  %[1]sArg[i] = %[2]s(*%[1]sV)
}
defer func() {
  for i, %[1]sV := range %[1]sArg {
    *%[1]s[i] = %[3]s(%[1]sV)
  }
}()

`, arg.Name, cArrayType, goArrayType)
	varName = fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name)
	return
}

// C.int*, C.int[] -> []*int32
func simplePtrSliceW(cArrayType, goArrayType string, arg ArgDef) (argType string, def string, varName string) {
	argType = fmt.Sprintf("[]*%s", goArrayType)
	def = fmt.Sprintf(`%[1]sArg := make([]%[2]s, len(%[1]s))
for i, %[1]sV := range %[1]s {
  %[1]sArg[i] = %[2]s(*%[1]sV)
}
defer func() {
  for i, %[1]sV := range %[1]sArg {
    *%[1]s[i] = %[3]s(%[1]sV)
  }
}()

`, arg.Name, cArrayType, goArrayType)
	varName = fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name)
	return
}

// C.ImVec2 -> ImVec2
func wrappableW(sName, sType string) (argType string, def string, varName string) {
	argType = sType
	varName = fmt.Sprintf("%s.toC()", sName)
	return
}

// C.ImVec2* -> *ImVec2
func wrappablePtrW(argName, goType, cType string) (argType, def, varName string) {
	argType = goType
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrap[%[3]s, %[2]s](%[1]s)
defer %[1]sFin()`, argName, goType, cType)
	varName = fmt.Sprintf("%sArg", argName)
	return
}

func wrappablePtrArrayW(size int, cArrayType, goArrayType string, arg ArgDef) (argType string, def string, varName string) {
	argType = fmt.Sprintf("[%d]*%s", size, goArrayType)
	def = fmt.Sprintf(`%[1]sArg := make([]%[2]s, len(%[1]s))
%[1]sFin := make([]func(), len(%[1]s))
for i, %[1]sV := range %[1]s {
	var tmp *%[2]s
  	tmp, %[1]sFin[i] = wrap[%[2]s, *%[3]s](%[1]sV)
  	%[1]sArg[i] = *tmp
}
defer func() {
  for _, %[1]sV := range %[1]sFin {
    %[1]sV()
  }
}()

`, arg.Name, cArrayType, goArrayType)
	varName = fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name)
	return
}
