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
		"bool":            {"bool", "%s == C.bool(true)", "C.bool"},
		"bool*":           simplePtrR("bool", "C.bool"),
		"char":            simpleR("rune", "C.char"),
		"char*":           {"string", "C.GoString(%s)", "*C.char"},
		"unsigned char":   simpleR("uint", "C.char"),
		"unsigned char*":  {"*uint", "(*uint)(unsafe.Pointer(%s))", "C.uchar"}, // NOTE: This should work but I'm not 100% sure
		"float":           simpleR("float32", "C.float"),
		"float*":          simplePtrR("float32", "C.float"),
		"double":          simpleR("float64", "C.double"),
		"double*":         simplePtrR("float64", "C.double"),
		"int":             simpleR("int32", "C.int"),
		"int32_t":         simpleR("int32", "C.int"),
		"int*":            simplePtrR("int32", "C.int"),
		"unsigned int":    simpleR("uint32", "C.uint"),
		"unsigned int*":   simplePtrR("uint32", "C.uint"),
		"short":           simpleR("int16", "C.short"),
		"unsigned short":  simpleR("uint16", "C.ushort"),
		"unsigned short*": simplePtrR("uint16", "C.ushort"),
		"uintptr_t":       simpleR("uintptr", "C.uintptr_t"),
		"size_t":          simpleR("uint64", "C.size_t"),
		"time_t":          simpleR("uint64", "C.time_t"),
		"void*":           simpleR("unsafe.Pointer", "unsafe.Pointer"),
		"tm":              wrappableR(prefixGoPackage("Tm", "implot", context), "C.tm"),

		// TODO: this should be generalized and loaded from preset
		"ImVec4_c*": imVec4PtrReturnW(context),
	}

	// import preset
	for k, v := range context.preset.SimpleTypes {
		returnWrapperMap[k] = simpleR("("+prefixGoPackage(v[0], v[2], context)+")", v[1])
	}

	for k, v := range context.preset.SimplePtrTypes {
		returnWrapperMap[k+"*"] = simplePtrR(prefixGoPackage(v[0], v[2], context), v[1])
	}

	for k, v := range context.preset.WrappableTypes {
		returnWrapperMap[k] = wrappableR(prefixGoPackage(v[0], v[2], context), v[1])
		// TODO: now wrappablePtrR yet - implement.
		// returnWrapperMap[k+"*"] = wrappablePtrR(prefixGoPackage("*"+v[0], v[2], context), v[1])
	}

	isPointer := HasSuffix(t, "*")
	pureType := TrimPrefix(TrimSuffix(t, "*"), "const ")
	// isNonPODUsed := HasSuffix(pureType, context.preset.NonPODUsedSuffix) // FIXME
	pureType = TrimSuffix(pureType, context.preset.NonPODUsedSuffix)
	// check if pureType is a declared type (struct or something else from typedefs)
	_, isRefStruct := context.refStructNames[pureType]
	_, isRefTypedef := context.refTypedefs[pureType]
	_, isEnum := context.enumNames[pureType]
	_, isRefEnum := context.refEnumNames[pureType]
	shouldSkipRefTypedef := context.ShouldSkipTypedef(pureType)
	_, isStruct := context.typedefsNames[pureType]
	isStruct = isStruct || ((isRefStruct || (isRefTypedef && !isRefEnum)) && !shouldSkipRefTypedef)
	w, known := returnWrapperMap[TrimPrefix(TrimSuffix(t, context.preset.NonPODUsedSuffix), "const ")]
	// check if is array (match regex)
	isArray, err := regexp.Match(".*\\[\\d+\\]", []byte(t))
	if err != nil {
		glg.Fatalf("Error in regex: %s", err)
	}

	/*
		if isNonPODUsed {
			fmt.Println(returnWrapperMap[pureType])
			fmt.Println(pureType, known)
			panic("")
		}
	*/

	srcPackage := GoIdentifier(context.flags.PackageName)
	if isRefTypedef {
		srcPackage = GoIdentifier(context.flags.RefPackageName)
	}

	shouldSkipStruct := context.ShouldSkipStruct(pureType)

	switch {
	case known:
		return w, nil
		// case (context.typedefsNames[t] || context.refStructNames[t]) && !shouldSkipStruct(t):
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
		!(HasSuffix(t, "]")):
		pureType := CIdentifier(TrimSuffix(TrimPrefix(t, "ImVector_"), "*"))
		if HasSuffix(pureType, "Ptr") {
			pureType = TrimSuffix(pureType, "Ptr")
		}

		for _, st := range context.structs {
			if st.Name == pureType {
				pureType = st.podName(context)
				break
			}
		}

		pureType += "*"

		rw, err := getReturnWrapper(pureType, context)
		if err != nil {
			return returnWrapper{}, fmt.Errorf("creating vector wrapper %w", err)
		}

		// NOTE: Special Case
		if pureType == "char*" {
			rw = simplePtrR("int8", "C.char")
		}

		if isPointer {
			return returnWrapper{
				returnType: GoIdentifier(fmt.Sprintf("vectors.Vector[%s]", Replace(rw.returnType, "*", "", 1))),
				returnStmt: fmt.Sprintf("func() vectors.Vector[%s] {result := %%[1]s; return vectors.NewVectorFromC(result.Size, result.Capacity, %s)}()", Replace(rw.returnType, "*", "", 1), fmt.Sprintf(rw.returnStmt, "*result.Data")),
				CType:      GoIdentifier(fmt.Sprintf("*C.%s", pureType)),
			}, nil
		} else {
			return returnWrapper{
				returnType: GoIdentifier(fmt.Sprintf("vectors.Vector[%s]", Replace(rw.returnType, "*", "", 1))),
				returnStmt: fmt.Sprintf("func() vectors.Vector[%s] {result := %%[1]s; return vectors.NewVectorFromC(result.Size, result.Capacity, %s)}()", Replace(rw.returnType, "*", "", 1), fmt.Sprintf(rw.returnStmt, "result.Data")),
				CType:      GoIdentifier(fmt.Sprintf("*C.%s", pureType)),
			}, nil
		}
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
