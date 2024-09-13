package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/kpango/glg"
)

// Wrapper for return value
type returnWrapper struct {
	returnType GoIdentifier
	returnStmt string
}

func getReturnWrapper(
	t CIdentifier,
	context *Context,
) (returnWrapper, error) {
	returnWrapperMap := map[CIdentifier]returnWrapper{
		"bool":           {"bool", "%s == C.bool(true)"},
		"char":           simpleR("rune"),
		"unsigned char":  simpleR("uint"),
		"unsigned char*": {"*uint", "(*uint)(unsafe.Pointer(%s))"}, // NOTE: This should work but I'm not 100% sure
		"char*":          {"string", "C.GoString(%s)"},
		"const char*":    {"string", "C.GoString(%s)"},
		"const ImWchar*": simpleR("(*Wchar)"),
		"ImWchar*":       simpleR("(*Wchar)"),
		"ImWchar":        simpleR("Wchar"),
		"ImWchar16":      simpleR("uint16"),
		"float":          simpleR("float32"),
		"float*":         simplePtrR("float32"),
		"double":         simpleR("float64"),
		"double*":        simplePtrR("float64"),
		"int":            simpleR("int32"),
		"int*":           simplePtrR("int32"),
		"unsigned int":   simpleR("uint32"),
		"unsigned int*":  simplePtrR("uint32"),
		"short":          simpleR("int16"),
		"unsigned short": simpleR("uint16"),
		"ImS8":           simpleR("int"),
		"ImS16":          simpleR("int"),
		"ImS32":          simpleR("int"),
		"ImS64":          simpleR("int64"),
		"ImU8":           simpleR("byte"),
		"ImU8*":          simplePtrR("byte"),
		"ImU16":          simpleR("uint16"),
		"ImU16*":         simplePtrR("uint16"),
		"ImU32":          simpleR("uint32"),
		"ImU32*":         simplePtrR("uint32"),
		"ImU64":          simpleR("uint64"),
		"ImU64*":         simplePtrR("uint64"),
		"ImVec4":         wrappableR("Vec4"),
		"const ImVec4*":  imVec4PtrReturnW(),
		"ImVec2":         wrappableR("Vec2"),
		"ImColor":        wrappableR("Color"),
		"ImPlotPoint":    wrappableR("PlotPoint"),
		"ImRect":         wrappableR("Rect"),
		"ImPlotTime":     wrappableR("PlotTime"),
		"uintptr_t":      simpleR("uintptr"),
		"size_t":         simpleR("uint64"),
	}

	pureType := TrimPrefix(TrimSuffix(t, "*"), "const ")
	// check if pureType is a declared type (struct or something else from typedefs)
	_, isRefStruct := context.refStructNames[pureType]
	_, shouldSkipRefTypedef := skippedTypedefs[pureType]
	_, isStruct := context.structNames[pureType]
	isStruct = isStruct || (isRefStruct && !shouldSkipRefTypedef)
	w, known := returnWrapperMap[t]
	// check if is array (match regex)
	isArray, err := regexp.Match(".*\\[\\d+\\]", []byte(t))
	if err != nil {
		glg.Fatalf("Error in regex: %s", err)
	}

	switch {
	case known:
		return w, nil
	case (context.structNames[t] || context.refStructNames[t]) && !shouldSkipStruct(t):
		return returnWrapper{
			returnType: t.renameGoIdentifier(context),
			returnStmt: fmt.Sprintf(`*new%sFromC(func() *C.%s {result := %%s; return &result}())`, t.renameGoIdentifier(context), t),
		}, nil
	case isEnum(t, context.enumNames, context):
		return returnWrapper{
			returnType: t.renameEnum(context),
			returnStmt: fmt.Sprintf("%s(%%s)", t.renameEnum(context)),
		}, nil
	case HasPrefix(t, "ImVector_") &&
		!(HasSuffix(t, "*") || HasSuffix(t, "]")):
		pureType := CIdentifier(TrimPrefix(t, "ImVector_") + "*")
		rw, err := getReturnWrapper(pureType, context)
		if err != nil {
			return returnWrapper{}, fmt.Errorf("creating vector wrapper %w", err)
		}
		return returnWrapper{
			returnType: GoIdentifier(fmt.Sprintf("Vector[%s]", rw.returnType)),
			returnStmt: fmt.Sprintf("newVectorFromC(%%[1]s.Size, %%[1]s.Capacity, %s)", fmt.Sprintf(rw.returnStmt, "%[1]s.Data")),
		}, nil
	case HasSuffix(t, "*") && isEnum(TrimSuffix(t, "*"), context.enumNames, context):
		return returnWrapper{
			returnType: "*" + TrimSuffix(t, "*").renameEnum(context),
			returnStmt: fmt.Sprintf("(*%s)(%%s)", TrimSuffix(t, "*").renameEnum(context)),
		}, nil
	case HasSuffix(t, "*") && isStruct && !shouldSkipStruct(pureType):
		return returnWrapper{
			returnType: "*" + TrimPrefix(TrimSuffix(t, "*"), "const ").renameGoIdentifier(context),
			returnStmt: fmt.Sprintf("new%sFromC(%%s)", TrimPrefix(TrimSuffix(t, "*"), "const ").renameGoIdentifier(context)),
		}, nil
	case isArray:
		typeCount := Split(t, "[")
		count, err := strconv.Atoi(string(TrimSuffix(typeCount[1], "]")))
		if err != nil {
			return returnWrapper{}, fmt.Errorf("parsing array len: %w", err)
		}

		typeName := typeCount[0]
		// check if array index getter exists
		getterName, ok := context.arrayIndexGetters[typeName]
		if !ok {
			return returnWrapper{}, fmt.Errorf("no array index getter for %s", typeName)
		}

		rw, err := getReturnWrapper(typeName, context)
		if err != nil {
			return returnWrapper{}, fmt.Errorf("creating array wrapper %w", err)
		}

		result := returnWrapper{
			returnType: GoIdentifier(fmt.Sprintf("[%d]%s", count, rw.returnType)),
			returnStmt: fmt.Sprintf(`func() [%[1]d]%[2]s {
result := [%[1]d]%[2]s{}
	resultMirr := %%s
	for i := range result {
		result[i] = %[3]s
	}

	return result
}()`, count, rw.returnType, fmt.Sprintf(rw.returnStmt, fmt.Sprintf("C.%s(resultMirr, C.int(i))", getterName))),
		}

		return result, nil
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

func simpleR(goType GoIdentifier) returnWrapper {
	return returnWrapper{goType, fmt.Sprintf("%s(%s)", goType, "%s")}
}

func simplePtrR(goType GoIdentifier) returnWrapper {
	return returnWrapper{
		returnType: GoIdentifier(fmt.Sprintf("*%s", goType)),
		returnStmt: fmt.Sprintf("(*%s)(%s)", goType, "%s"),
	}
}

func wrappableR(goType GoIdentifier) returnWrapper {
	return returnWrapper{
		returnType: goType,
		returnStmt: fmt.Sprintf("*(&%s{}).fromC(%s)", goType, "%s"),
	}
}
