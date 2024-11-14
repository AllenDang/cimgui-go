package main

import "C"
import (
	"fmt"
	"regexp"
	"strconv"
)

type ArgumentWrapperData struct {
	// go-valid argument type (e.g. string, ImVec2, etc.)
	ArgType GoIdentifier
	// argument deffinition (e.g. arg1, arg1Fin := ...\ndefer arg1Fin())
	ArgDef string
	// This is equivalent of above (VarName will be appropiate here), but it will not define Finalizer (use this if you're not going to use Finalizer)
	ArgDefNoFin string
	// one-line go statement that should be called after calling C function
	// in order to update go value.
	// It is intended to be run in defer statement (and it will be in most cases)
	// so it should be one-line function call OR a call of custom function
	Finalizer string

	// a name of variable of wrapped C type
	VarName string

	NoFin bool

	// CType is a valid type that will have VarName.
	CType GoIdentifier
}

type argumentWrapper func(arg ArgDef) ArgumentWrapperData

func getArgWrapper(
	a *ArgDef,
	makeFirstArgReceiver, isGetter bool,
	context *Context,
) (argDeclaration string, data ArgumentWrapperData, err error) {
	argWrapperMap := map[CIdentifier]argumentWrapper{
		"char":                simpleW("rune", "C.char"),
		"char*":               constCharW,
		"const char*":         constCharW,
		"const char**":        charPtrPtrW,
		"const char* const[]": charPtrPtrW,
		"unsigned char":       simpleW("uint", "C.uchar"),
		"const unsigned char": simpleW("uint", "C.uchar"),
		"unsigned char*":      simplePtrW("uint", "C.uchar"),
		"unsigned char**":     uCharPtrW,
		"size_t":              simpleW("uint64", "C.xulong"),
		"time_t":              simpleW("uint64", "C.xulong"),
		"size_t*":             sizeTPtrW,
		"float":               simpleW("float32", "C.float"),
		"const float":         simpleW("float32", "C.float"),
		"float*":              simplePtrW("float32", "C.float"),
		"const float*":        simplePtrW("float32", "C.float"),
		"short":               simpleW("int16", "C.short"),
		"unsigned short":      simpleW("uint16", "C.ushort"),
		"unsigned short*":     simplePtrW("uint16", "C.ushort"),
		"ImU8":                simpleW("byte", "C.ImU8"),
		"ImU8*":               simplePtrW("byte", "C.ImU8"),
		"const ImU8*":         simplePtrW("byte", "C.ImU8"),
		"ImU16":               simpleW("uint16", "C.ImU16"),
		"ImU16*":              simplePtrW("uint16", "C.ImU16"),
		"const ImU16*":        simplePtrW("uint16", "C.ImU16"),
		"ImU32":               simpleW("uint32", "C.ImU32"),
		"ImU32*":              simplePtrW("uint32", "C.ImU32"),
		"const ImU32*":        simplePtrW("uint32", "C.ImU32"),
		"ImU64":               simpleW("uint64", "C.ImU64"),
		"ImU64*":              simplePtrW("uint64", "C.ImU64"),
		"const ImU64*":        uint64ArrayW,
		"ImS8":                simpleW("int", "C.ImS8"),
		"ImS8*":               simplePtrW("int8", "C.ImS8"),
		"const ImS8*":         simplePtrW("int8", "C.ImS8"),
		"ImS16":               simpleW("int16", "C.ImS16"),
		"ImS16*":              simplePtrW("int16", "C.ImS16"),
		"const ImS16*":        simplePtrW("int16", "C.ImS16"),
		"ImS32":               simpleW("int", "C.ImS32"),
		"ImS32*":              simplePtrW("int32", "C.ImS32"),
		"const ImS32*":        simplePtrW("int32", "C.ImS32"),
		"int32_t":             simpleW("int32", "C.int32_t"),
		"ImS64":               simpleW("int64", "C.ImS64"),
		"ImS64*":              simplePtrW("int64", "C.ImS64"),
		"const ImS64*":        int64ArrayW,
		"int":                 simpleW("int32", "C.int"),
		"const int":           simpleW("int32", "C.int"),
		"int*":                simplePtrW("int32", "C.int"),
		"unsigned int":        simpleW("uint32", "C.uint"),
		"unsigned int*":       simplePtrW("uint32", "C.uint"),
		"double":              simpleW("float64", "C.double"),
		"double*":             simplePtrW("float64", "C.double"),
		"const double*":       simplePtrW("float64", "C.double"),
		"bool":                simpleW("bool", "C.bool"),
		"const bool":          simpleW("bool", "C.bool"),
		"bool*":               simplePtrW("bool", "C.bool"),
		"const bool*":         simplePtrW("bool", "C.bool"),
		"ImWchar":             simpleW(prefixGoPackage("Wchar", "imgui", context), "C.ImWchar"),
		"ImWchar*":            simpleW("("+prefixGoPackage("*Wchar", "imgui", context)+")", "(*C.ImWchar)"),
		"const ImWchar*":      simpleW("("+prefixGoPackage("*Wchar", "imgui", context)+")", "(*C.ImWchar)"),
		"ImWchar16":           simpleW("uint16", "C.ImWchar16"),
		"uintptr_t":           simpleW("uintptr", "C.uintptr_t"),
		"const uintptr_t":     simpleW("uintptr", "C.uintptr_t"),
		"const ImVec2":        wrappableW(prefixGoPackage("Vec2", "imgui", context), "C.ImVec2"),
		"const ImVec2*":       wrappablePtrW(prefixGoPackage("*Vec2", "imgui", context), "C.ImVec2"),
		"ImVec2":              wrappableW(prefixGoPackage("Vec2", "imgui", context), "C.ImVec2"),
		"ImVec2*":             wrappablePtrW(prefixGoPackage("*Vec2", "imgui", context), "C.ImVec2"),
		"ImVec2[2]":           wrappablePtrArrayW(2, "C.ImVec2", prefixGoPackage("Vec2", "imgui", context)),
		"const ImVec4":        wrappableW(prefixGoPackage("Vec4", "imgui", context), "C.ImVec4"),
		"const ImVec4*":       wrappablePtrW(prefixGoPackage("*Vec4", "imgui", context), "C.ImVec4"),
		"ImVec4":              wrappableW(prefixGoPackage("Vec4", "imgui", context), "C.ImVec4"),
		"ImVec4*":             wrappablePtrW(prefixGoPackage("*Vec4", "imgui", context), "C.ImVec4"),
		"ImColor":             wrappableW(prefixGoPackage("Color", "imgui", context), "C.ImColor"),
		"ImColor*":            wrappablePtrW(prefixGoPackage("*Color", "imgui", context), "C.ImColor"),
		"ImRect":              wrappableW(prefixGoPackage("Rect", "imgui", context), "C.ImRect"),
		"const ImRect":        wrappableW(prefixGoPackage("Rect", "imgui", context), "C.ImRect"),
		"ImRect*":             wrappablePtrW(prefixGoPackage("*Rect", "imgui", context), "C.ImRect"),
		"const ImRect*":       wrappablePtrW(prefixGoPackage("*Rect", "imgui", context), "C.ImRect"),
		"ImPlotPoint":         wrappableW(prefixGoPackage("PlotPoint", "implot", context), "C.ImPlotPoint"),
		"const ImPlotPoint":   wrappableW(prefixGoPackage("PlotPoint", "implot", context), "C.ImPlotPoint"),
		"ImPlotPoint*":        wrappablePtrW(prefixGoPackage("*PlotPoint", "implot", context), "C.ImPlotPoint"),
		"ImPlotTime":          wrappableW(prefixGoPackage("PlotTime", "implot", context), "C.ImPlotTime"),
		"const ImPlotTime":    wrappableW(prefixGoPackage("PlotTime", "implot", context), "C.ImPlotTime"),
		"ImPlotTime*":         wrappablePtrW(prefixGoPackage("*PlotTime", "implot", context), "C.ImPlotTime"),
		"const ImPlotTime*":   wrappablePtrW(prefixGoPackage("*PlotTime", "implot", context), "C.ImPlotTime"),
		"tm":                  wrappableW(prefixGoPackage("Tm", "implot", context), "C.struct_tm"),
		"const tm":            wrappableW(prefixGoPackage("Tm", "implot", context), "C.struct_tm"),
		"tm*":                 wrappablePtrW(prefixGoPackage("*Tm", "imgui", context), "C.struct_tm"),
		"const tm*":           wrappablePtrW(prefixGoPackage("*Tm", "imgui", context), "C.struct_tm"),
		"void*":               simpleW("unsafe.Pointer", "unsafe.Pointer"),
	}

	if a.Name == "type" || a.Name == "range" {
		a.Name += "Arg"
	}

	if makeFirstArgReceiver {
		return "", ArgumentWrapperData{}, nil
	}

	if isGetter {
		argDeclaration = fmt.Sprintf("%s %s", a.Name, a.Type.renameGoIdentifier(context))
		data = ArgumentWrapperData{
			ArgDef:      fmt.Sprintf("%[1]sArg, %[1]sFin := %[1]s.Handle()", a.Name),
			ArgDefNoFin: fmt.Sprintf("%[1]sArg, _ := %[1]s.Handle()", a.Name),
			VarName:     fmt.Sprintf("%sArg", a.Name),
			Finalizer:   fmt.Sprintf("%sFin()", a.Name),
			NoFin:       a.RemoveFinalizer,
		}

		return
	}

	if v, ok := argWrapperMap[a.Type]; ok {
		arg := v(*a)
		data = arg
		data.NoFin = a.RemoveFinalizer

		argDeclaration = fmt.Sprintf("%s %s", a.Name, arg.ArgType)

		return argDeclaration, data, nil
	}

	pureType := TrimPrefix(a.Type, "const ")
	isPointer := false
	if HasSuffix(a.Type, "*") {
		pureType = TrimSuffix(pureType, "*")
		isPointer = true
	}

	_, isRefTypedef := context.refTypedefs[pureType]
	_, isEnum := context.enumNames[pureType]
	_, isRefEnum := context.refEnumNames[pureType]

	if isEnum || isRefEnum {
		srcPkg := context.flags.packageName
		if isRefTypedef {
			srcPkg = context.flags.refPackageName
		}

		goType := prefixGoPackage(pureType.renameGoIdentifier(context), GoIdentifier(srcPkg), context)

		if isPointer {
			argDeclaration = fmt.Sprintf("%s *%s", a.Name, goType)
			data = ArgumentWrapperData{
				ArgType: GoIdentifier(fmt.Sprintf("*%s", goType)),
				VarName: fmt.Sprintf("(*C.%s)(%s)", pureType, a.Name),
				CType:   GoIdentifier(fmt.Sprintf("*C.%s", pureType)),
			}
		} else {
			argDeclaration = fmt.Sprintf("%s %s", a.Name, goType)
			data = ArgumentWrapperData{
				ArgType: goType,
				VarName: fmt.Sprintf("C.%s(%s)", pureType, a.Name),
				CType:   GoIdentifier(fmt.Sprintf("C.%s", a.Type)),
			}
		}

		return
	}

	if HasPrefix(a.Type, "ImVector_") &&
		!(HasSuffix(a.Type, "*") || HasSuffix(a.Type, "]")) {
		pureType := TrimPrefix(a.Type, "ImVector_") + "*"
		dataName := a.Name + "Data"
		_, w, err := getArgWrapper(&ArgDef{
			Name: dataName,
			Type: pureType,
		}, false, false, context)
		if err != nil {
			return "", ArgumentWrapperData{}, fmt.Errorf("creating vector wrapper %w", err)
		}

		// NOTE Special Case
		if pureType == "char*" { // we need to handle it as *byte, not string
			charWrapper := simplePtrW("int8", "C.char")
			w = charWrapper(ArgDef{
				Name: dataName,
				Type: pureType,
			})
		}

		data = ArgumentWrapperData{
			VarName: string("*" + a.Name + "VecArg"),
			ArgType: GoIdentifier(fmt.Sprintf("vectors.Vector[%s]", Replace(w.ArgType, "*", "", 1))),
			ArgDef: fmt.Sprintf(`%[5]s := %[2]s.Data
%[1]s
%[2]sVecArg := new(C.%[3]s)
%[2]sVecArg.Size = C.int(%[2]s.Size)
%[2]sVecArg.Capacity = C.int(%[2]s.Capacity)
%[2]sVecArg.Data = %[4]s
%[2]s.Pinner().Pin(%[2]sVecArg.Data)
`, w.ArgDef, a.Name, a.Type, w.VarName, dataName),
			ArgDefNoFin: fmt.Sprintf(`%[5]s := %[2]s.Data
%[1]s
%[2]sVecArg := new(C.%[3]s)
%[2]sVecArg.Size = C.int(%[2]s.Size)
%[2]sVecArg.Capacity = C.int(%[2]s.Capacity)
%[2]sVecArg.Data = %[4]s
%[2]s.Pinner().Pin(%[2]sVecArg.Data)
`, w.ArgDefNoFin, a.Name, a.Type, w.VarName, dataName),
			Finalizer: fmt.Sprintf("%s\n%s.Pinner().Unpin()", w.Finalizer, a.Name),
			NoFin:     a.RemoveFinalizer,
		}

		argDeclaration = fmt.Sprintf("%s %s", a.Name, data.ArgType)

		return argDeclaration, data, nil
	}

	// consider array of form type[number] (check with regex)
	r, err := regexp.Compile(`\w*\[\d+\]`)
	if err != nil {
		return "", ArgumentWrapperData{}, fmt.Errorf("cannot compile regex: %w", err)
	}

	if r.MatchString(string(a.Type)) {
		// split out count and type name ("pureType")
		parts := Split(a.Type, "[")
		countStr := TrimSuffix(parts[1], "]")
		count, err := strconv.Atoi(string(countStr))
		if err != nil {
			return "", ArgumentWrapperData{}, fmt.Errorf("cannot convert count to int: %w", err)
		}
		pureType := parts[0]
		// decode pureType
		singleDefName := a.Name + "V"
		_, w, err := getArgWrapper(&ArgDef{
			Name: singleDefName,
			Type: pureType,
		}, false, false, context)
		if err != nil {
			return "", ArgumentWrapperData{}, fmt.Errorf("creating array wrapper %w", err)
		}
		// we also need to be able to convert this back
		fromC, err := getReturnWrapper(pureType, context)
		if err != nil {
			return "", ArgumentWrapperData{}, fmt.Errorf("creating array wrapper %w", err)
		}
		def := fmt.Sprintf(`
%[1]sArg := make([]%[5]s, len(%[1]s))
for i, %[3]s := range %[1]s {
%[2]s
%[1]sArg[i] = %[4]s
}`, a.Name, w.ArgDefNoFin, singleDefName, w.VarName, w.CType)
		data := ArgumentWrapperData{
			ArgType:     GoIdentifier(fmt.Sprintf("*[%d]%s", count, w.ArgType)),
			ArgDef:      def,
			ArgDefNoFin: def,
			VarName:     fmt.Sprintf("(*%s)(&%sArg[0])", w.CType, a.Name),
			Finalizer: fmt.Sprintf(`
for i, %[1]sV := range %[1]sArg {
	(*%[1]s)[i] = %[2]s
}
		
		`, a.Name, fmt.Sprintf(fromC.returnStmt, fmt.Sprintf("%sV", a.Name))),
			CType: "*" + w.CType,
		}

		argDeclaration = fmt.Sprintf("%s %s", a.Name, data.ArgType)

		return argDeclaration, data, nil
	}

	_, shouldSkipRefTypedef := context.preset.SkipTypedefs[pureType]
	if context.typedefsNames[pureType] || (isRefTypedef && !shouldSkipRefTypedef) {
		srcPkg := context.flags.packageName
		if isRefTypedef {
			srcPkg = context.flags.refPackageName
		}

		goType := prefixGoPackage(pureType.renameGoIdentifier(context), GoIdentifier(srcPkg), context)
		cType := GoIdentifier(fmt.Sprintf("C.%s", pureType))
		argType := goType
		fn := ""
		if isPointer {
			argType = "*" + argType
			cType = GoIdentifier(fmt.Sprintf("*C.%s", pureType))
			fn = "Handle"
		} else {
			fn = "C"
		}

		w := ArgumentWrapperData{
			ArgType:   argType,
			VarName:   fmt.Sprintf("internal.ReinterpretCast[%s](%sArg)", cType, a.Name),
			Finalizer: fmt.Sprintf("%sFin()", a.Name),
			NoFin:     a.RemoveFinalizer,
			CType:     cType,
		}

		w.ArgDef = fmt.Sprintf("%[1]sArg, %[1]sFin := %[1]s.%[2]s()", a.Name, fn)
		w.ArgDefNoFin = fmt.Sprintf("%[1]sArg, _ := %[1]s.%[2]s()", a.Name, fn)

		data = w

		argDeclaration = fmt.Sprintf("%s %s", a.Name, data.ArgType)

		return
	}

	return "", ArgumentWrapperData{}, fmt.Errorf("unknown argument type \"%s\"", a.Type)
}

func constCharW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType:     "string",
		VarName:     fmt.Sprintf("%sArg", arg.Name),
		ArgDef:      fmt.Sprintf("%[1]sArg, %[1]sFin := internal.WrapString[C.char](%[1]s)", arg.Name),
		ArgDefNoFin: fmt.Sprintf("%[1]sArg, _ := internal.WrapString[C.char](%[1]s)", arg.Name),
		Finalizer:   fmt.Sprintf("%sFin()", arg.Name),
		CType:       "*C.char",
	}
}

func charPtrPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType:     "[]string",
		VarName:     fmt.Sprintf("%sArg", arg.Name),
		ArgDef:      fmt.Sprintf("%[1]sArg, %[1]sFin := internal.WrapStringList[C.char](%[1]s)", arg.Name),
		ArgDefNoFin: fmt.Sprintf("%[1]sArg, _ := internal.WrapStringList[C.char](%[1]s)", arg.Name),
		Finalizer:   fmt.Sprintf("%sFin()", arg.Name),
		CType:       "**C.char",
	}
}

func uCharPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "*C.uchar",
		VarName: fmt.Sprintf("&%s", arg.Name),
		CType:   "*C.uchar",
	}
}

func sizeTPtrW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "*uint64",
		VarName: fmt.Sprintf("(*C.xulong)(%s)", arg.Name),
		CType:   "*C.xulong",
	}
}

func int64ArrayW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "[]int64",
		VarName: fmt.Sprintf("(*C.longlong)(&(%s[0]))", arg.Name),
		CType:   "*C.longlong",
	}
}

func uint64ArrayW(arg ArgDef) ArgumentWrapperData {
	return ArgumentWrapperData{
		ArgType: "[]uint64",
		VarName: fmt.Sprintf("(*C.ulonglong)(&(%s[0]))", arg.Name),
		CType:   "*C.ulonglong",
	}
}

//func voidPtrW(arg ArgDef) ArgumentWrapperData {
//	return ArgumentWrapperData{
//		ArgType:     "unsafe.Pointer",
//		ArgDef:      fmt.Sprintf("%[1]sArg, %[1]sFin := WrapVoidPtr(%[1]s)", arg.Name),
//		ArgDefNoFin: fmt.Sprintf("%[1]sArg, _ := WrapVoidPtr(%[1]s)", arg.Name),
//		VarName:     fmt.Sprintf("%sArg", arg.Name),
//		Finalizer:   fmt.Sprintf("%sFin()", arg.Name),
//		CType:       "unsafe.Pointer",
//	}
//}

// generic wrappers:

// C.int -> int32
func simpleW(goType GoIdentifier, cType GoIdentifier) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: goType,
			VarName: fmt.Sprintf("%s(%s)", cType, arg.Name),
			CType:   cType,
		}
	}
}

// C.int* -> *int32
//
//	return simplePtrW(arg.Name, "int16", "C.int")
func simplePtrW(goType GoIdentifier, cType GoIdentifier) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType:     GoIdentifier(fmt.Sprintf("*%s", goType)),
			ArgDef:      fmt.Sprintf("%[1]sArg, %[1]sFin := internal.WrapNumberPtr[%[2]s, %[3]s](%[1]s)", arg.Name, cType, goType),
			ArgDefNoFin: fmt.Sprintf("%[1]sArg, _ := internal.WrapNumberPtr[%[2]s, %[3]s](%[1]s)", arg.Name, cType, goType),
			Finalizer:   fmt.Sprintf("%[1]sFin()", arg.Name, cType, goType),
			VarName:     fmt.Sprintf("%sArg", arg.Name),
			CType:       "*" + cType,
		}
	}
}

// C.int*, C.int[] as well as C.int[2] -> [2]*int32
func simplePtrArrayW(size int, cArrayType GoIdentifier, goArrayType GoIdentifier) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		def := fmt.Sprintf(`
%[1]sArg := make([]%[2]s, len(%[1]s))
for i, %[1]sV := range %[1]s {
  %[1]sArg[i] = %[2]s(%[1]sV)
}`, arg.Name, cArrayType)
		return ArgumentWrapperData{
			ArgType:     GoIdentifier(fmt.Sprintf("*[%d]%s", size, goArrayType)),
			ArgDef:      def,
			ArgDefNoFin: def,
			VarName:     fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name),
			Finalizer: fmt.Sprintf(`
for i, %[1]sV := range %[1]sArg {
	(*%[1]s)[i] = %[3]s(%[1]sV)
}

`, arg.Name, cArrayType, goArrayType),
			CType: "*" + cArrayType,
		}
	}
}

// C.ImVec2 -> ImVec2
func wrappableW(goType, cType GoIdentifier) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType: goType,
			VarName: fmt.Sprintf("internal.ReinterpretCast[%s](%s.ToC())", cType, arg.Name),
			CType:   cType,
		}
	}
}

// C.ImVec2* -> *ImVec2
func wrappablePtrW(goType, cType GoIdentifier) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		return ArgumentWrapperData{
			ArgType:     goType,
			ArgDef:      fmt.Sprintf("%[1]sArg, %[1]sFin := internal.Wrap(%[1]s)", arg.Name, goType, cType),
			ArgDefNoFin: fmt.Sprintf("%[1]sArg, _ := internal.Wrap(%[1]s)", arg.Name, goType, cType),
			Finalizer:   fmt.Sprintf("%[1]sFin()", arg.Name, goType, cType),
			VarName:     fmt.Sprintf("internal.ReinterpretCast[*%s](%sArg)", cType, arg.Name),
			CType:       "*" + cType,
		}
	}
}

func wrappablePtrArrayW(size int, cArrayType GoIdentifier, goArrayType GoIdentifier) argumentWrapper {
	return func(arg ArgDef) ArgumentWrapperData {
		def := fmt.Sprintf(`%[1]sArg := make([]%[2]s, len(%[1]s))
%[1]sFin := make([]func(), len(%[1]s))
for i, %[1]sV := range %[1]s {
	var tmp *%[2]s
  	tmp, %[1]sFin[i] = internal.Wrap(%[1]sV)
  	%[1]sArg[i] = *tmp
}
`, arg.Name, cArrayType, goArrayType)
		return ArgumentWrapperData{
			ArgType:     GoIdentifier(fmt.Sprintf("[%d]*%s", size, goArrayType)),
			ArgDef:      def,
			ArgDefNoFin: def,
			Finalizer: fmt.Sprintf(`
  for _, %[1]sV := range %[1]sFin {
    %[1]sV()
  }
`, arg.Name, cArrayType, goArrayType),
			VarName: fmt.Sprintf("(*%s)(&%sArg[0])", cArrayType, arg.Name),
			CType:   "*" + cArrayType,
		}
	}
}
