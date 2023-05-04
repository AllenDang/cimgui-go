package main

import "fmt"

// Wrapper for return value
type returnWrapper struct {
	returnType string
	returnStmt string
}

func getReturnTypeWrapperFunc(returnType string) (returnWrapper, error) {
	returnWrapperMap := map[string]returnWrapper{
		"bool":                     {"bool", "return %s == C.bool(true)"},
		"char*":                    {"string", "return C.GoString(%s)"},
		"const char*":              {"string", "return C.GoString(%s)"},
		"const ImWchar*":           simpleR("(*Wchar)"),
		"ImWchar":                  simpleR("Wchar"),
		"ImWchar16":                simpleR("uint16"),
		"float":                    simpleR("float32"),
		"double":                   simpleR("float64"),
		"int":                      simpleR("int"),
		"unsigned int":             simpleR("uint32"),
		"short":                    simpleR("int"),
		"ImS8":                     simpleR("int"),
		"ImS16":                    simpleR("int"),
		"ImS32":                    simpleR("int"),
		"ImU8":                     simpleR("uint32"),
		"ImU16":                    simpleR("uint32"),
		"ImU32":                    simpleR("uint32"),
		"ImU64":                    simpleR("uint64"),
		"ImVec4":                   wrappableR("Vec4"),
		"const ImVec4*":            imVec4PtrReturnW(),
		"ImGuiID":                  simpleR("ID"),
		"ImTextureID":              simpleR("TextureID"),
		"ImVec2":                   wrappableR("Vec2"),
		"ImColor":                  wrappableR("Color"),
		"ImPlotPoint":              wrappableR("PlotPoint"),
		"ImRect":                   wrappableR("Rect"),
		"ImPlotTime":               wrappableR("PlotTime"),
		"ImGuiTableColumnIdx":      simpleR("TableColumnIdx"),
		"ImGuiTableDrawChannelIdx": simpleR("TableDrawChannelIdx"),
		"void*":                    simpleR("unsafe.Pointer"),
		"size_t":                   simpleR("float64"),
	}

	if v, ok := returnWrapperMap[returnType]; ok {
		return v, nil
	}

	return returnWrapper{}, fmt.Errorf("return type %s not found", returnType)
}

func imVec4PtrReturnW() returnWrapper {
	// TODO: verify if it wraps correctly
	return returnWrapper{
		returnType: "*Vec4",
		returnStmt: `out := &Vec4{}
out.fromC(*%s)
return out
`,
	}
}

func simpleR(goType string) returnWrapper {
	return returnWrapper{goType, fmt.Sprintf("return %s(%s)", goType, "%s")}
}

func wrappableR(goType string) returnWrapper {
	return returnWrapper{
		returnType: goType,
		returnStmt: fmt.Sprintf("return *(&%s{}).fromC(%s)", goType, "%s"),
	}
}
