package main

import "C"
import (
	"fmt"
)

type ArgumentWrapperData struct {
	// go-valid argument type (e.g. string, ImVec2, etc.)
	ArgType string
	// argument deffinition (e.g. arg1, arg1Fin := ...\ndefer arg1Fin())
	ArgDef string
	// one-line go statement that should be called after calling C function
	// in order to update go value.
	// It is intended to be run in defer statement (and it will be in most cases)
	// so it should be one-line function call OR a call of custom function
	Finalizer string

	// a name of variable of wrapped C type
	VarName string
}

type argumentWrapper func(arg ArgDef) ArgumentWrapperData

func argWrapper(argType string) (wrapper argumentWrapper, err error) {
	argWrapperMap := map[string]argumentWrapper{
		"char*":                    constCharW,
		"const char*":              constCharW,
		"const char**":             charPtrPtrW,
		"const char* const[]":      charPtrPtrW,
		"unsigned char":            simpleW("uint", "C.uchar"),
		"unsigned char**":          uCharPtrW,
		"size_t":                   simpleW("uint64", "C.xlong"),
		"size_t*":                  sizeTPtrW,
		"float":                    floatW,
		"float*":                   floatPtrW,
		"const float*":             floatArrayW,
		"short":                    simpleW("int", "C.short"),
		"unsigned short":           simpleW("uint", "C.ushort"),
		"ImU8":                     simpleW("uint", "C.ImU8"),
		"const ImU8*":              simplePtrSliceW("C.ImU8", "byte"),
		"ImU16":                    simpleW("uint", "C.ImU16"),
		"const ImU16*":             simplePtrSliceW("C.ImU16", "uint16"),
		"ImU32":                    simpleW("uint32", "C.ImU32"),
		"const ImU32*":             simplePtrSliceW("C.ImU32", "uint32"),
		"ImU64":                    simpleW("uint64", "C.ImU64"),
		"const ImU64*":             uint64ArrayW,
		"ImS8":                     simpleW("int", "C.ImS8"),
		"const ImS8*":              simplePtrSliceW("C.ImS8", "int8"),
		"ImS16":                    simpleW("int", "C.ImS16"),
		"const ImS16*":             simplePtrSliceW("C.ImS16", "int"),
		"ImS32":                    simpleW("int", "C.ImS32"),
		"const ImS32*":             simplePtrSliceW("C.ImS32", "int32"),
		"const ImS64*":             int64ArrayW,
		"int":                      simpleW("int32", "C.int"),
		"int*":                     simplePtrW("int32", "C.int"),
		"unsigned int":             simpleW("uint32", "C.uint"),
		"unsigned int*":            simplePtrW("uint32", "C.uint"),
		"double":                   simpleW("float64", "C.double"),
		"double*":                  simplePtrW("float64", "C.double"),
		"bool":                     simpleW("bool", "C.bool"),
		"bool*":                    boolPtrW,
		"int[2]":                   simplePtrArrayW(2, "C.int", "int32"),
		"int[3]":                   simplePtrArrayW(3, "C.int", "int32"),
		"int[4]":                   simplePtrArrayW(4, "C.int", "int32"),
		"float[2]":                 simplePtrArrayW(2, "C.float", "float32"),
		"float[3]":                 simplePtrArrayW(3, "C.float", "float32"),
		"float[4]":                 simplePtrArrayW(4, "C.float", "float32"),
		"ImWchar":                  simpleW("Wchar", "C.ImWchar"),
		"const ImWchar*":           simpleW("*Wchar", "(*C.ImWchar)"),
		"ImGuiID":                  simpleW("ImGuiID", "C.ImGuiID"),
		"ImTextureID":              simpleW("TextureID", "C.ImTextureID"),
		"ImDrawIdx":                simpleW("DrawIdx", "C.ImDrawIdx"),
		"ImGuiTableColumnIdx":      simpleW("TableColumnIdx", "C.ImGuiTableColumnIdx"),
		"ImGuiTableDrawChannelIdx": simpleW("TableDrawChannelIdx", "C.ImGuiTableDrawChannelIdx"),
		"void*":                    simpleW("unsafe.Pointer", ""),
		"const void*":              simpleW("unsafe.Pointer", ""),
		"const ImVec2":             wrappableW("Vec2"),
		"const ImVec2*":            wrappablePtrW("*Vec2", "C.ImVec2"),
		"ImVec2":                   wrappableW("Vec2"),
		"ImVec2*":                  wrappablePtrW("*Vec2", "C.ImVec2"),
		"ImVec2[2]":                wrappablePtrArrayW(2, "C.ImVec2", "Vec2"),
		"const ImVec4":             wrappableW("Vec4"),
		"const ImVec4*":            wrappablePtrW("*Vec4", "C.ImVec4"),
		"ImVec4":                   wrappableW("Vec4"),
		"ImVec4*":                  wrappablePtrW("*Vec4", "C.ImVec4"),
		"ImColor*":                 wrappablePtrW("*Color", "C.ImColor"),
		"ImRect":                   wrappableW("Rect"),
		"ImRect*":                  wrappablePtrW("*Rect", "C.ImRect"),
		"ImPlotPoint":              wrappableW("PlotPoint"),
		"const ImPlotPoint":        wrappableW("PlotPoint"),
		"ImPlotPoint*":             wrappablePtrW("*PlotPoint", "C.ImPlotPoint"),
	}

	if wrapper, ok := argWrapperMap[argType]; ok {
		return wrapper, nil
	}

	return nil, fmt.Errorf("no wrapper for type %s", argType)
}

func constCharW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType:   "string",
		VarName:   fmt.Sprintf("%sArg", arg.Name),
		ArgDef:    fmt.Sprintf("%[1]sArg, %[1]sFin := wrapString(%[1]s)", arg.Name),
		Finalizer: fmt.Sprintf("%sFin()", arg.Name),
	}
}

func charPtrPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType:   "[]string",
		VarName:   fmt.Sprintf("%sArg", arg.Name),
		ArgDef:    fmt.Sprintf("%[1]sArg, %[1]sFin := wrapStringList(%[1]s)", arg.Name),
		Finalizer: fmt.Sprintf("%sFin()", arg.Name),
	}
}

func uCharPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "*C.uchar",
		VarName: fmt.Sprintf("&%s", arg.Name),
	}
}

func sizeTPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "*uint64",
		VarName: fmt.Sprintf("(*C.xlong)(%s)", arg.Name),
	}
}

func floatW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "float32",
		VarName: fmt.Sprintf("C.float(%s)", arg.Name),
	}
}

// leave this for now because of https://github.com/AllenDang/cimgui-go/issues/31
func floatPtrW(arg ArgDef) ArgumentWrapperData {
	return simplePtrW("float32", "C.float")(arg)
}

func floatArrayW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "[]float32",
		VarName: fmt.Sprintf("(*C.float)(&(%s[0]))", arg.Name),
	}
}

func boolPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType:   "*bool",
		ArgDef:    fmt.Sprintf("%[1]sArg, %[1]sFin := wrapBool(%[1]s)", arg.Name),
		Finalizer: fmt.Sprintf("%[1]sFin()", arg.Name),
		VarName:   fmt.Sprintf("%sArg", arg.Name),
	}
}

func int64ArrayW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "[]int64",
		VarName: fmt.Sprintf("(*C.longlong)(&(%s[0]))", arg.Name),
	}
}

func uint64ArrayW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "[]uint64",
		VarName: fmt.Sprintf("(*C.ulonglong)(&(%s[0]))", arg.Name),
	}
}

func inputeTextCallbackW(arg ArgDef) (argType string, def string, varName string) {
	argType = "ImGuiInputTextCallback"
	// TODO: implement me
	return
}

// generic wrappers:

// C.int -> int32
func simpleW(goType, cType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: goType,
			VarName: fmt.Sprintf("%s(%s)", cType, arg.Name),
		}
	}
}

// C.int* -> *int32
//
//	return simplePtrW(arg.Name, "int16", "C.int")
func simplePtrW(goType, cType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType:   fmt.Sprintf("*%s", goType),
			ArgDef:    fmt.Sprintf("%[1]sArg, %[1]sFin := wrapNumberPtr[%[2]s, %[3]s](%[1]s)", arg.Name, cType, goType),
			Finalizer: fmt.Sprintf("%[1]sFin()", arg.Name, cType, goType),
			VarName:   fmt.Sprintf("%sArg", arg.Name),
		}
	}
}

// C.int*, C.int[] as well as C.int[2] -> [2]*int32
func simplePtrArrayW(size int, cArrayType, goArrayType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: fmt.Sprintf("*[%d]%s", size, goArrayType),
			ArgDef: fmt.Sprintf(`
%[1]sArg := make([]%[2]s, len(%[1]s))
for i, %[1]sV := range %[1]s {
  %[1]sArg[i] = %[2]s(%[1]sV)
}`, arg.Name, cArrayType),
			VarName: fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name),
			Finalizer: fmt.Sprintf(`
func() {
  for i, %[1]sV := range %[1]sArg {
    (*%[1]s)[i] = %[3]s(%[1]sV)
  }
}()

`, arg.Name, cArrayType, goArrayType),
		}
	}
}

// C.int*, C.int[] -> *[]int32
func simplePtrSliceW(cArrayType, goArrayType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: fmt.Sprintf("*[]%s", goArrayType),
			ArgDef: fmt.Sprintf(`%[1]sArg := make([]%[2]s, len(*%[1]s))
for i, %[1]sV := range *%[1]s {
  %[1]sArg[i] = %[2]s(%[1]sV)
}
`, arg.Name, cArrayType, goArrayType),
			Finalizer: fmt.Sprintf(`func() {
  for i, %[1]sV := range %[1]sArg {
    (*%[1]s)[i] = %[3]s(%[1]sV)
  }
}()
`, arg.Name, cArrayType, goArrayType),
			VarName: fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name),
		}
	}
}

// C.ImVec2 -> ImVec2
func wrappableW(sType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: sType,
			VarName: fmt.Sprintf("%s.toC()", arg.Name),
		}
	}
}

// C.ImVec2* -> *ImVec2
func wrappablePtrW(goType, cType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType:   goType,
			ArgDef:    fmt.Sprintf("%[1]sArg, %[1]sFin := wrap[%[3]s, %[2]s](%[1]s)", arg.Name, goType, cType),
			Finalizer: fmt.Sprintf("%[1]sFin()", arg.Name, goType, cType),
			VarName:   fmt.Sprintf("%sArg", arg.Name),
		}
	}
}

func wrappablePtrArrayW(size int, cArrayType, goArrayType string) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: fmt.Sprintf("[%d]*%s", size, goArrayType),
			ArgDef: fmt.Sprintf(`%[1]sArg := make([]%[2]s, len(%[1]s))
%[1]sFin := make([]func(), len(%[1]s))
for i, %[1]sV := range %[1]s {
	var tmp *%[2]s
  	tmp, %[1]sFin[i] = wrap[%[2]s, *%[3]s](%[1]sV)
  	%[1]sArg[i] = *tmp
}
`, arg.Name, cArrayType, goArrayType),
			Finalizer: fmt.Sprintf(`func() {
  for _, %[1]sV := range %[1]sFin {
    %[1]sV()
  }
}()
`, arg.Name, cArrayType, goArrayType),
			VarName: fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name),
		}
	}
}
