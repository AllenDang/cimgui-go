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
	CType      GoIdentifier
}

func getReturnWrapper(
	t CIdentifier,
	context *Context,
) (returnWrapper, error) {
	returnWrapperMap := map[CIdentifier]returnWrapper{
		"bool":                {"bool", "%s == C.bool(true)", "C.bool"},
		"bool*":               simplePtrR("bool", "C.bool"),
		"const bool*":         simplePtrR("bool", "c.bool"),
		"char":                simpleR("rune", "C.char"),
		"unsigned char":       simpleR("uint", "C.char"),
		"const unsigned char": simpleR("uint", "C.char"),
		"unsigned char*":      {"*uint", "(*uint)(unsafe.Pointer(%s))", "C.uchar"}, // NOTE: This should work but I'm not 100% sure
		"char*":               {"string", "C.GoString(%s)", "*C.char"},
		"const char*":         {"string", "C.GoString(%s)", "*C.char"},
		"const ImWchar*":      simpleR("(*Wchar)", "*C.ImWchar"),
		"ImWchar*":            simpleR("(*Wchar)", "*C.ImWchar"),
		"ImWchar":             simpleR("Wchar", "C.ImWchar"),
		"ImWchar16":           simpleR("uint16", "C.ImWchar16"),
		"float":               simpleR("float32", "C.float"),
		"float*":              simplePtrR("float32", "C.float"),
		"double":              simpleR("float64", "C.double"),
		"double*":             simplePtrR("float64", "C.double"),
		"int":                 simpleR("int32", "C.int"),
		"int32_t":             simpleR("int32", "C.int"),
		"int*":                simplePtrR("int32", "C.int"),
		"unsigned int":        simpleR("uint32", "C.uint"),
		"unsigned int*":       simplePtrR("uint32", "C.uint"),
		"short":               simpleR("int16", "C.short"),
		"unsigned short":      simpleR("uint16", "C.ushort"),
		"unsigned short*":     simplePtrR("uint16", "C.ushort"),
		"ImS8":                simpleR("int", "C.ImS8"),
		"ImS16":               simpleR("int16", "C.ImS16"),
		"ImS16*":              simplePtrR("int16", "C.ImS16"),
		"ImS32":               simpleR("int", "C.ImS32"),
		"ImS64":               simpleR("int64", "C.ImS64"),
		"ImS64*":              simplePtrR("int64", "C.ImS64"),
		"ImU8":                simpleR("byte", "C.ImU8"),
		"ImU8*":               simplePtrR("byte", "C.ImU8"),
		"ImU16":               simpleR("uint16", "C.ImU16"),
		"ImU16*":              simplePtrR("uint16", "C.ImU16"),
		"ImU32":               simpleR("uint32", "C.ImU32"),
		"ImU32*":              simplePtrR("uint32", "C.ImU32"),
		"const ImU32*":        simplePtrR("uint32", "C.ImU32"),
		"ImU64":               simpleR("uint64", "C.ImU64"),
		"ImU64*":              simplePtrR("uint64", "C.ImU64"),
		"ImVec4":              wrappableR(prefixGoPackage("Vec4", "imgui", context), "C.ImVec4"),
		"const ImVec4*":       imVec4PtrReturnW(context),
		"ImVec2":              wrappableR(prefixGoPackage("Vec2", "imgui", context), "C.ImVec2"),
		"ImColor":             wrappableR(prefixGoPackage("Color", "imgui", context), "C.ImColor"),
		"ImPlotPoint":         wrappableR(prefixGoPackage("PlotPoint", "implot", context), "C.ImPlotPoint"),
		"ImRect":              wrappableR(prefixGoPackage("Rect", "imgui", context), "C.ImRect"),
		"ImPlotTime":          wrappableR(prefixGoPackage("PlotTime", "implot", context), "C.ImPlotTime"),
		"tm":                  wrappableR(prefixGoPackage("Tm", "implot", context), "C.tm"),
		"const tm":            wrappableR(prefixGoPackage("Tm", "implot", context), "C.tm"),
		"uintptr_t":           simpleR("uintptr", "C.uintptr_t"),
		"size_t":              simpleR("uint64", "C.size_t"),
		"time_t":              simpleR("uint64", "C.time_t"),
		"void*":               simpleR("unsafe.Pointer", "unsafe.Pointer"),
	}

	isPointer := HasSuffix(t, "*")
	pureType := TrimPrefix(TrimSuffix(t, "*"), "const ")
	// check if pureType is a declared type (struct or something else from typedefs)
	_, isRefStruct := context.refStructNames[pureType]
	_, isRefTypedef := context.refTypedefs[pureType]
	_, isEnum := context.enumNames[pureType]
	_, isRefEnum := context.refEnumNames[pureType]
	_, shouldSkipRefTypedef := context.preset.SkipTypedefs[pureType]
	_, isStruct := context.structNames[pureType]
	isStruct = isStruct || ((isRefStruct || (isRefTypedef && !IsEnum(pureType, context.refEnumNames))) && !shouldSkipRefTypedef)
	w, known := returnWrapperMap[t]
	// check if is array (match regex)
	isArray, err := regexp.Match(".*\\[\\d+\\]", []byte(t))
	if err != nil {
		glg.Fatalf("Error in regex: %s", err)
	}

	srcPackage := GoIdentifier(context.flags.packageName)
	if isRefTypedef {
		srcPackage = GoIdentifier(context.flags.refPackageName)
	}

	_, shouldSkipStruct := context.preset.SkipStructs[pureType]

	switch {
	case known:
		return w, nil
		// case (context.structNames[t] || context.refStructNames[t]) && !shouldSkipStruct(t):
	case !isPointer && isStruct && !shouldSkipStruct:
		return returnWrapper{
			returnType: prefixGoPackage(pureType.renameGoIdentifier(context), srcPackage, context),
			// this is a small trick as using prefixGoPackage isn't in fact intended to be used in such a way, but it should treat the whole string as a "type" and prefix it correctly
			returnStmt: string(prefixGoPackage(GoIdentifier(fmt.Sprintf("*New%sFromC(func() *C.%s {result := %%s; return &result}())", pureType.renameGoIdentifier(context), pureType)), srcPackage, context)),
			CType:      GoIdentifier(fmt.Sprintf("C.%s", pureType)),
		}, nil
	case isEnum || isRefEnum:
		goType := prefixGoPackage(pureType.renameGoIdentifier(context), srcPackage, context)
		if isPointer {
			return returnWrapper{
				returnType: "*" + goType,
				returnStmt: fmt.Sprintf("(*%s)(%%s)", goType),
				CType:      GoIdentifier(fmt.Sprintf("*C.%s", TrimSuffix(pureType, "*"))),
			}, nil
		} else {
			return returnWrapper{
				returnType: goType,
				returnStmt: fmt.Sprintf("%s(%%s)", goType),
				CType:      GoIdentifier(fmt.Sprintf("C.%s", pureType)),
			}, nil
		}
	case HasPrefix(t, "ImVector_") &&
		!(HasSuffix(t, "*") || HasSuffix(t, "]")):
		pureType := CIdentifier(TrimPrefix(t, "ImVector_") + "*")
		rw, err := getReturnWrapper(pureType, context)
		if err != nil {
			return returnWrapper{}, fmt.Errorf("creating vector wrapper %w", err)
		}

		// NOTE: Special Case
		if pureType == "char*" {
			rw = simplePtrR("int8", "C.char")
		}

		return returnWrapper{
			returnType: GoIdentifier(fmt.Sprintf("vectors.Vector[%s]", Replace(rw.returnType, "*", "", 1))),
			returnStmt: fmt.Sprintf("vectors.NewVectorFromC(%%[1]s.Size, %%[1]s.Capacity, %s)", fmt.Sprintf(rw.returnStmt, "%[1]s.Data")),
			CType:      GoIdentifier(fmt.Sprintf("*C.%s", pureType)),
		}, nil
	case HasSuffix(t, "*") && isStruct && !shouldSkipStruct:
		goType := prefixGoPackage("*"+TrimSuffix(pureType, "*").renameGoIdentifier(context), srcPackage, context)
		return returnWrapper{
			returnType: goType,
			returnStmt: string(prefixGoPackage(GoIdentifier(fmt.Sprintf("New%sFromC(%%s)", TrimSuffix(pureType, "*").renameGoIdentifier(context))), srcPackage, context)),
			CType:      GoIdentifier(fmt.Sprintf("*C.%s", TrimSuffix(pureType, "*"))),
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
			CType: GoIdentifier(fmt.Sprintf("*C.%s", typeName)),
		}

		return result, nil
	default:
		return returnWrapper{}, fmt.Errorf("unknown return type %s", pureType)
	}
}

func imVec4PtrReturnW(ctx *Context) returnWrapper {
	// TODO: verify if it wraps correctly
	goType := prefixGoPackage("Vec4", "imgui", ctx)
	return returnWrapper{
		returnType: "*" + goType,
		returnStmt: "(&" + string(goType) + "{}).FromC(unsafe.Pointer(%s))",
		CType:      "*C.ImVec4",
	}
}

func simpleR(goType GoIdentifier, cType GoIdentifier) returnWrapper {
	return returnWrapper{
		returnType: goType,
		returnStmt: fmt.Sprintf("%s(%s)", goType, "%s"),
		CType:      cType,
	}
}

func simplePtrR(goType, cType GoIdentifier) returnWrapper {
	return returnWrapper{
		returnType: GoIdentifier(fmt.Sprintf("*%s", goType)),
		returnStmt: fmt.Sprintf("(*%s)(%s)", goType, "%s"),
		CType:      GoIdentifier(fmt.Sprintf("*%s", cType)),
	}
}

func wrappableR(goType, cType GoIdentifier) returnWrapper {
	return returnWrapper{
		returnType: goType,
		returnStmt: fmt.Sprintf("func() %[1]s {out := %%s ; return *(&%[1]s{}).FromC(unsafe.Pointer(&out))}()", goType),
		CType:      cType,
	}
}
