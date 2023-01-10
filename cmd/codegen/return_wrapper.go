package main

import "fmt"

// Wrapper for return value
type returnWrapper func() (returnType string, returnStmt string)

func getReturnTypeWrapperFunc(returnType string) (returnWrapper, error) {
	returnWrapperMap := map[string]returnWrapper{
		"bool":                     boolReturnW,
		"char*":                    constCharReturnW,
		"const char*":              constCharReturnW,
		"const ImWchar*":           constWCharPtrReturnW,
		"ImWchar":                  imWcharReturnW,
		"float":                    floatReturnW,
		"double":                   doubleReturnW,
		"int":                      intReturnW,
		"unsigned int":             uintReturnW,
		"short":                    intReturnW,
		"ImS8":                     intReturnW,
		"ImS16":                    intReturnW,
		"ImS32":                    intReturnW,
		"ImU8":                     uintReturnW,
		"ImU16":                    uintReturnW,
		"ImU32":                    u32ReturnW,
		"ImU64":                    uint64ReturnW,
		"ImVec4":                   imVec4ReturnW,
		"const ImVec4*":            imVec4PtrReturnW,
		"ImGuiID":                  idReturnW,
		"ImTextureID":              textureIdReturnW,
		"ImVec2":                   imVec2ReturnW,
		"ImColor":                  imColorReturnW,
		"ImPlotPoint":              imPlotPointReturnW,
		"ImRect":                   imRectReturnW,
		"ImGuiTableColumnIdx":      imTableColumnIdxReturnW,
		"ImGuiTableDrawChannelIdx": imTableDrawChannelIdxReturnW,
		"void*":                    voidPtrReturnW,
		"size_t":                   doubleReturnW,
	}

	if v, ok := returnWrapperMap[returnType]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("return type %s not found", returnType)
}

func boolReturnW() (returnType string, returnStmt string) {
	returnType = "bool"
	returnStmt = "return %s == C.bool(true)\n"
	return
}

func constCharReturnW() (returnType string, returnStmt string) {
	returnType = "string"
	returnStmt = "return C.GoString(%s)"
	return
}

func floatReturnW() (returnType string, returnStmt string) {
	returnType = "float32"
	returnStmt = "return float32(%s)"
	return
}

func doubleReturnW() (returnType string, returnStmt string) {
	returnType = "float64"
	returnStmt = "return float64(%s)"
	return
}

func intReturnW() (returnType string, returnStmt string) {
	returnType = "int"
	returnStmt = "return int(%s)"
	return
}

func constWCharPtrReturnW() (returnType string, returnStmt string) {
	returnType = "*Wchar"
	returnStmt = "return (*Wchar)(%s)"
	return
}

func imWcharReturnW() (returnType string, returnStmt string) {
	returnType = "Wchar"
	returnStmt = "return (Wchar)(%s)"
	return
}

func imVec4PtrReturnW() (returnType string, returnStmt string) {
	// TODO: verify if it wraps correctly
	returnType = "*Vec4"
	returnStmt += `out := &Vec4{}
out.fromC(*%s)
return out
`
	return
}

func imVec4ReturnW() (returnType string, returnStmt string) {
	returnType = "Vec4"
	returnStmt = fmt.Sprintf("out := &%s{}\n", returnType)
	returnStmt += `out.fromC(%s)
return *out
`
	return
}

func imVec2ReturnW() (returnType string, returnStmt string) {
	returnType = "Vec2"
	returnStmt = fmt.Sprintf("out := &%s{}\n", returnType)
	returnStmt += `out.fromC(%s)
return *out
`
	return
}

func imColorReturnW() (returnType string, returnStmt string) {
	returnType = "Color"
	returnStmt = fmt.Sprintf("out := &%s{}\n", returnType)
	returnStmt += `out.fromC(%s)
return *out
`
	return
}

func imPlotPointReturnW() (returnType string, returnStmt string) {
	returnType = "PlotPoint"
	returnStmt = fmt.Sprintf("out := &%s{}\n", returnType)
	returnStmt += `out.fromC(%s)
return *out
`
	return
}

func imRectReturnW() (returnType string, returnStmt string) {
	returnType = "Rect"
	returnStmt = fmt.Sprintf("out := &%s{}\n", returnType)
	returnStmt += `out.fromC(%s)
return *out
`
	return
}

func imTableColumnIdxReturnW() (returnType string, returnStmt string) {
	returnType = "TableColumnIdx"
	returnStmt = "return TableColumnIdx(%s)"
	return
}

func imTableDrawChannelIdxReturnW() (returnType string, returnStmt string) {
	returnType = "TableDrawChannelIdx"
	returnStmt = "return TableDrawChannelIdx(%s)"
	return
}

func voidPtrReturnW() (returnType string, returnStmt string) {
	returnType = "unsafe.Pointer"
	returnStmt = "return unsafe.Pointer(%s)"
	return
}

func u32ReturnW() (returnType string, returnStmt string) {
	returnType = "uint32"
	returnStmt = "return uint32(%s)"
	return
}

func uintReturnW() (returnType string, returnStmt string) {
	returnType = "uint32"
	returnStmt = "return uint32(%s)"
	return
}

func uint64ReturnW() (returnType string, returnStmt string) {
	returnType = "uint64"
	returnStmt = "return uint64(%s)"
	return
}

func idReturnW() (returnType string, returnStmt string) {
	returnType = "ID"
	returnStmt = "return ID(%s)"
	return
}

func textureIdReturnW() (returnType string, returnStmt string) {
	returnType = "TextureID"
	returnStmt = "return TextureID(%s)"
	return
}
