package main

import (
	"fmt"
	"strings"
)

// Wrapper for return value
type returnWrapper struct {
	returnType string
	returnStmt string
}

func getReturnWrapper(
	t string,
	structNames,
	enumNames map[string]bool,
) (returnWrapper, error) {
	returnWrapperMap := map[string]returnWrapper{
		"bool":                     {"bool", "%s == C.bool(true)"},
		"char":                     simpleR("rune"),
		"unsigned char":            simpleR("uint"),
		"unsigned char*":           {"(*uint)", "(*uint)(unsafe.Pointer(%s))"}, // NOTE: This should work but I'm not 100% sure
		"char*":                    {"string", "C.GoString(%s)"},
		"const char*":              {"string", "C.GoString(%s)"},
		"const ImWchar*":           simpleR("(*Wchar)"),
		"ImWchar":                  simpleR("Wchar"),
		"ImWchar16":                simpleR("uint16"),
		"float":                    simpleR("float32"),
		"double":                   simpleR("float64"),
		"int":                      simpleR("int32"),
		"int*":                     simplePtrR("*int32"),
		"unsigned int":             simpleR("uint32"),
		"unsigned int*":            simplePtrR("*uint32"),
		"short":                    simpleR("int16"),
		"unsigned short":           simpleR("uint16"),
		"ImS8":                     simpleR("int"),
		"ImS16":                    simpleR("int"),
		"ImS32":                    simpleR("int"),
		"ImS64":                    simpleR("int64"),
		"ImU8":                     simpleR("byte"),
		"ImU16":                    simpleR("uint16"),
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
		"size_t":                   simpleR("uint64"),
	}

	w, known := returnWrapperMap[t]
	switch {
	case known:
		return w, nil
	case structNames[t] && !shouldSkipStruct(t):
		return returnWrapper{
			returnType: renameGoIdentifier(t),
			returnStmt: fmt.Sprintf(`*new%sFromC(func() *C.%s {result := %%s; return &result}())
`, renameGoIdentifier(t), t),
		}, nil
	case isEnum(t, enumNames):
		return returnWrapper{
			returnType: renameEnum(t),
			returnStmt: fmt.Sprintf("%s(%%s)", renameEnum(t)),
		}, nil
	case strings.HasPrefix(t, "ImVector_") &&
		!(strings.HasSuffix(t, "*") || strings.HasSuffix(t, "]")):
		pureType := strings.TrimPrefix(t, "ImVector_") + "*"
		rw, err := getReturnWrapper(pureType, structNames, enumNames)
		if err != nil {
			return returnWrapper{}, fmt.Errorf("creating vector wrapper %w", err)
		}
		return returnWrapper{
			returnType: fmt.Sprintf("Vector[%s]", rw.returnType),
			returnStmt: fmt.Sprintf("newVectorFromC(%%[1]s.Size, %%[1]s.Capacity, %s)", fmt.Sprintf(rw.returnStmt, "%[1]s.Data")),
		}, nil
	case strings.HasSuffix(t, "*") && structNames[strings.TrimPrefix(strings.TrimSuffix(t, "*"), "const ")] && !shouldSkipStruct(strings.TrimPrefix(strings.TrimSuffix(t, "*"), "const ")):
		return returnWrapper{
			returnType: "*" + renameGoIdentifier(strings.TrimPrefix(strings.TrimSuffix(t, "*"), "const ")),
			returnStmt: fmt.Sprintf("new%sFromC(%%s)", renameGoIdentifier(strings.TrimPrefix(strings.TrimSuffix(t, "*"), "const "))),
		}, nil
	case strings.HasSuffix(t, "*") && isEnum(strings.TrimSuffix(t, "*"), enumNames):
		return returnWrapper{
			returnType: "*" + renameEnum(strings.TrimSuffix(t, "*")),
			returnStmt: fmt.Sprintf("(*%s)(%%s)", renameEnum(strings.TrimSuffix(t, "*"))),
		}, nil
	default:
		return returnWrapper{}, fmt.Errorf("unknown return type %s", t)
	}
}

func imVec4PtrReturnW() returnWrapper {
	// TODO: verify if it wraps correctly
	return returnWrapper{
		returnType: "*Vec4",
		returnStmt: "(&Vec4{}).fromC(*%s)",
	}
}

func simpleR(goType string) returnWrapper {
	return returnWrapper{goType, fmt.Sprintf("%s(%s)", goType, "%s")}
}

func simplePtrR(goType string) returnWrapper {
	return returnWrapper{goType, fmt.Sprintf("(%s)(%s)", goType, "%s")}
}

func wrappableR(goType string) returnWrapper {
	return returnWrapper{
		returnType: goType,
		returnStmt: fmt.Sprintf("*(&%s{}).fromC(%s)", goType, "%s"),
	}
}
