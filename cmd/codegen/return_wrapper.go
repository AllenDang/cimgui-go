package main

// Wrapper for return value
type returnWrapper func() (returnType string, returnStmt string)

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
	returnType = "*ImWchar"
	returnStmt = "return (*ImWchar)(%s)"
	return
}

func imVec4PtrReturnW() (returnType string, returnStmt string) {
	returnType = "ImVec4"
	returnStmt = "return newImVec4FromCPtr(%s)"
	return
}

func imVec4ReturnW() (returnType string, returnStmt string) {
	returnType = "ImVec4"
	returnStmt = "return newImVec4FromC(%s)"
	return
}

func imVec2ReturnW() (returnType string, returnStmt string) {
	returnType = "ImVec2"
	returnStmt = "return newImVec2FromC(%s)"
	return
}

func imColorReturnW() (returnType string, returnStmt string) {
	returnType = "ImColor"
	returnStmt = "return newImColorFromC(%s)"
	return
}

func imPlotPointReturnW() (returnType string, returnStmt string) {
	returnType = "ImPlotPoint"
	returnStmt = "return newImPlotPointFromC(%s)"
	return
}

func imRectReturnW() (returnType string, returnStmt string) {
	returnType = "ImRect"
	returnStmt = "return newImRectFromC(%s)"
	return
}

func imTableColumnIdxReturnW() (returnType string, returnStmt string) {
	returnType = "ImGuiTableColumnIdx"
	returnStmt = "return ImGuiTableColumnIdx(%s)"
	return
}

func imTableDrawChannelIdxReturnW() (returnType string, returnStmt string) {
	returnType = "ImGuiTableDrawChannelIdx"
	returnStmt = "return ImGuiTableDrawChannelIdx(%s)"
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
	returnType = "ImGuiID"
	returnStmt = "return ImGuiID(%s)"
	return
}

func textureIdReturnW() (returnType string, returnStmt string) {
	returnType = "ImTextureID"
	returnStmt = "return ImTextureID(%s)"
	return
}
