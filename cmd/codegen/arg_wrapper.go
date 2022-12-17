package main

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
	return simpleValueW(arg.Name, "uint", "uchar")
}

func uCharPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*C.uchar"
	varName = fmt.Sprintf("&%s", arg.Name)
	return
}

func sizeTW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint64", "xlong")
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
	return wrapNumberPtr(arg.Name, "float32", "C.float")
}

func floatArrayW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]float32"
	varName = fmt.Sprintf("(*C.float)(&(%s[0]))", arg.Name)
	return
}

func boolW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "bool", "bool")
}

func boolPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*bool"
	def = fmt.Sprintf("%[1]sArg, %[1]sFin := wrapBool(%[1]s)\ndefer %[1]sFin()", arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func shortW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "int", "short")
}

func ushortW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint", "ushort")
}

func u8W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint", "ImU8")
}

func u8PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "byte", "C.uint")
}

func u16PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "uint16", "C.uint")
}

func u16W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint", "ImU16")
}

func int2W(arg ArgDef) (argType string, def string, varName string) {
	return arrayW(2, "int", "int32", arg)
}

func int3W(arg ArgDef) (argType string, def string, varName string) {
	return arrayW(3, "int", "int32", arg)
}

func int4W(arg ArgDef) (argType string, def string, varName string) {
	return arrayW(4, "int", "int32", arg)
}

func u32W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint32", "ImU32")
}

func u64W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint64", "ImU64")
}

func s8W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "int", "ImS8")
}

func s16W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "int", "ImS16")
}

func s32W(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "int", "ImS32")
}

func float2W(arg ArgDef) (argType string, def string, varName string) {
	return arrayW(2, "float", "float32", arg)
}

func float3W(arg ArgDef) (argType string, def string, varName string) {
	return arrayW(3, "float", "float32", arg)
}

func float4W(arg ArgDef) (argType string, def string, varName string) {
	return arrayW(4, "float", "float32", arg)
}

func imWcharW(arg ArgDef) (argType string, def string, varName string) {
	argType = "ImWchar"
	varName = fmt.Sprintf("C.ImWchar(%s)", arg.Name)
	return
}

func imWcharPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*ImWchar"
	varName = fmt.Sprintf("(*C.ImWchar)(%s)", arg.Name)
	return
}

func intW(arg ArgDef) (argType string, def string, varName string) {
	argType = "int32"
	varName = fmt.Sprintf("C.int(%s)", arg.Name)
	return
}

func int8PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "int8", "C.int")
}

func int16PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "int16", "C.int")
}

func intPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "int32", "C.int")
}

func int64ArrayW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]int64"
	varName = fmt.Sprintf("(*C.longlong)(&(%s[0]))", arg.Name)
	return
}

func uintW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "uint32", "uint")
}

func uintPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "uint32", "C.uint")
}

func uint64ArrayW(arg ArgDef) (argType string, def string, varName string) {
	argType = "[]uint64"
	varName = fmt.Sprintf("(*C.ulonglong)(&(%s[0]))", arg.Name)
	return
}

func doubleW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "float64", "double")
}

func doublePtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapNumberPtr(arg.Name, "float64", "C.double")
}

func imGuiIDW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "ImGuiID", "ImGuiID")
}

func imTextureIDW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "ImTextureID", "ImTextureID")
}

func imDrawIdxW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "ImDrawIdx", "ImDrawIdx")
}

func imTableColumnIdxW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "ImGuiTableColumnIdx", "ImGuiTableColumnIdx")
}

func imTableDrawChannelIdxW(arg ArgDef) (argType string, def string, varName string) {
	return simpleValueW(arg.Name, "ImGuiTableDrawChannelIdx", "ImGuiTableDrawChannelIdx")
}

func voidPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "unsafe.Pointer"
	varName = arg.Name
	return
}

func imPlotPointW(arg ArgDef) (argType string, def string, varName string) {
	return valueStructW(arg.Name, "ImPlotPoint")
}

func imPlotPointPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapImGuiTypePtr(arg.Name, "*ImPlotPoint")
}

func imVec2W(arg ArgDef) (argType string, def string, varName string) {
	return valueStructW(arg.Name, "ImVec2")
}

func imVec2PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapImGuiTypePtr(arg.Name, "*ImVec2")
}

// ImVec2[2] -> [2]ImVec2
func imVec22W(arg ArgDef) (argType string, def string, varName string) {
	argType = "[2]*ImVec2"
	varName = fmt.Sprintf("(*C.ImVec2)(unsafe.Pointer(&%s[0]))", arg.Name)
	return
}

func imVec4W(arg ArgDef) (argType string, def string, varName string) {
	return valueStructW(arg.Name, "ImVec4")
}

func imVec4PtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapImGuiTypePtr(arg.Name, "*ImVec4")
}

func imRectW(arg ArgDef) (argType string, def string, varName string) {
	return valueStructW(arg.Name, "ImRect")
}

func imRectPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapImGuiTypePtr(arg.Name, "ImRect")
}

func imColorPtrW(arg ArgDef) (argType string, def string, varName string) {
	return wrapImGuiTypePtr(arg.Name, "*ImColor")
}

func inputeTextCallbackW(arg ArgDef) (argType string, def string, varName string) {
	argType = "ImGuiInputTextCallback"
	// TODO: implement me
	return
}

// generic wrappers:

func simpleValueW(argName, goType, cType string) (argType string, def string, varName string) {
	argType = goType
	varName = fmt.Sprintf("C.%s(%s)", cType, argName)
	return
}

func wrapImGuiTypePtr(argName, goType string) (argType, def, varName string) {
	argType = goType
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := %[1]s.wrap()
defer %[1]sFin()`, argName)
	varName = fmt.Sprintf("%sArg", argName)
	return
}

func wrapNumberPtr(argName, goType, cType string) (argType, def, varName string) {
	argType = fmt.Sprintf("*%s", goType)
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrapNumberPtr[%[2]s, %[3]s](%[1]s)
defer %[1]sFin()`, argName, cType, goType)
	varName = fmt.Sprintf("%sArg", argName)
	return
}

func arrayW(size int, cArrayType, goArrayType string, arg ArgDef) (argType string, def string, varName string) {
	argType = fmt.Sprintf("[%d]*%s", size, goArrayType)
	def = fmt.Sprintf(`%[1]sArg := make([]C.%[2]s, len(%[1]s))
for i, %[1]sV := range %[1]s {
  %[1]sArg[i] = C.%[2]s(*%[1]sV)
}
defer func() {
  for i, %[1]sV := range %[1]sArg {
    *%[1]s[i] = %[3]s(%[1]sV)
  }
}()

`, arg.Name, cArrayType, goArrayType)
	varName = fmt.Sprintf("(*C.%s)(&%sArg[0])", cArrayType, arg.Name)
	return
}

func valueStructW(sName, sType string) (argType string, def string, varName string) {
	argType = sType
	varName = fmt.Sprintf("%s.toC()", sName)
	return
}
