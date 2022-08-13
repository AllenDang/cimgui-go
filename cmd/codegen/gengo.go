package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thoas/go-funk"
)

type TypeMap struct {
	GoType     string
	CgoWrapper string
}

func tm(goType string, cgoWrapper string) *TypeMap {
	return &TypeMap{
		GoType:     goType,
		CgoWrapper: cgoWrapper,
	}
}

var (
	structMemberTypeMap = map[string]*TypeMap{
		"unsigned int": tm("uint32", "C.uint(%s)"),
		"float":        tm("float32", "C.float(%s)"),
		"int":          tm("int32", "C.int(%s)"),
	}
)

// Generate enums and return enum type names
func generateGoEnums(enums []EnumDef) []string {
	var sb strings.Builder

	sb.WriteString("package cimgui\n\n")

	var enumNames []string
	for _, e := range enums {
		eName := strings.TrimRight(e.Name, "_")

		enumNames = append(enumNames, eName)

		sb.WriteString(fmt.Sprintf("type %s int\n", eName))
		sb.WriteString("const (\n")

		for _, v := range e.Values {
			sb.WriteString(fmt.Sprintf("\t%s = %d\n", v.Name, v.Value))
		}

		sb.WriteString(")\n\n")
	}

	enumFile, err := os.Create("enums.go")
	if err != nil {
		panic(err.Error())
	}
	defer enumFile.Close()

	_, _ = enumFile.WriteString(sb.String())

	return enumNames
}

func generateGoStructs(structs []StructDef) []string {
	var sb strings.Builder

	sb.WriteString(`package cimgui

 // #include "cimgui_wrapper.h"
 import "C"

`)

	// Save all struct name into a map
	var structNames []string

	for _, s := range structs {
		if strings.HasPrefix(s.Name, "Im") {
			sb.WriteString(fmt.Sprintf("type %[1]s C.%[1]s\n", s.Name))
			structNames = append(structNames, s.Name)
		}
	}

	structFile, err := os.Create("structs.go")
	if err != nil {
		panic(err.Error())
	}
	defer structFile.Close()

	_, _ = structFile.WriteString(sb.String())

	return structNames
}

type typeWrapper func(arg ArgDef) (argType string, def string, varName string)

func constCharW(arg ArgDef) (argType string, def string, varName string) {
	argType = "string"
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrapString(%[1]s)
defer %[1]sFin()`, arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func sizeTW(arg ArgDef) (argType string, def string, varName string) {
	argType = "uint64"
	varName = fmt.Sprintf("C.ulong(%s)", arg.Name)
	return
}

func floatW(arg ArgDef) (argType string, def string, varName string) {
	argType = "float32"
	varName = fmt.Sprintf("C.float(%s)", arg.Name)
	return
}

func floatPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*float32"
	def = fmt.Sprintf(`%[1]sArg, %[1]sFin := wrapFloat(%[1]s)
defer %[1]sFin()`, arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func boolW(arg ArgDef) (argType string, def string, varName string) {
	argType = "bool"
	def = fmt.Sprintf("%[1]sArg := C.bool(%[1]s)", arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func boolPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*bool"
	def = fmt.Sprintf("%[1]sArg, %[1]sFin := wrapBool(%[1]s)\ndefer %[1]sFin()", arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func arrayW(size int, arrayType, goArrayType string, arg ArgDef) (argType string, def string, varName string) {
	argType = fmt.Sprintf("*[%d]%s", size, goArrayType)
	def = fmt.Sprintf("%[1]sArg := (*C.%[2]s)(&%[1]s[0])", arg.Name, arrayType)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
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
	argType = "uint32"
	varName = fmt.Sprintf("C.uint(%s)", arg.Name)
	return
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

func intW(arg ArgDef) (argType string, def string, varName string) {
	argType = "int32"
	varName = fmt.Sprintf("C.int(%s)", arg.Name)
	return
}

func intPtrW(arg ArgDef) (argType string, def string, varName string) {
	argType = "*int32"
	def = fmt.Sprintf("%[1]sArg, %[1]sFin := wrapInt32(%[1]s)\ndefer %[1]sFin()", arg.Name)
	varName = fmt.Sprintf("%sArg", arg.Name)
	return
}

func imGuiIDW(arg ArgDef) (argType string, def string, varName string) {
	argType = "ImGuiID"
	varName = fmt.Sprintf("C.ImGuiID(%s)", arg.Name)
	return
}

func imTextureIDW(arg ArgDef) (argType string, def string, varName string) {
	argType = "ImTextureID"
	varName = fmt.Sprintf("C.ImTextureID(%s)", arg.Name)
	return
}

type returnWrapper func(f FuncDef) (returnType string, returnStmt string)

func boolReturnW(f FuncDef) (returnType string, returnStmt string) {
	returnType = "bool"
	returnStmt = "return %s != C.bool(true)\n"
	return
}

func constCharReturnW(f FuncDef) (returnType string, returnStmt string) {
	returnType = "string"
	returnStmt = "return C.GoString(%s)"
	return
}

func floatReturnW(f FuncDef) (returnType string, returnStmt string) {
	returnType = "float32"
	returnStmt = "return float32(%s)"
	return
}

func intReturnW(f FuncDef) (returnType string, returnStmt string) {
	returnType = "int"
	returnStmt = "return int(%s)"
	return
}

func constWCharPtrReturnW(f FuncDef) (returnType string, returnStmt string) {
	returnType = "*ImWchar"
	returnStmt = "return (*ImWchar)(%s)"
	return
}

func generateGoWrapper(validFuncs []FuncDef, enumNames []string, structNames []string) {
	var sb strings.Builder
	convertedFuncCount := 0

	sb.WriteString(`package cimgui

// #include "cimgui_wrapper.h"
import "C"

`)

	argWrapperMap := map[string]typeWrapper{
		"const char*": constCharW,
		"size_t":      sizeTW,
		"float":       floatW,
		"float*":      floatPtrW,
		"int":         intW,
		"int*":        intPtrW,
		"bool":        boolW,
		"bool*":       boolPtrW,
		"int[2]":      int2W,
		"int[3]":      int3W,
		"int[4]":      int4W,
		"ImU32":       u32W,
		"float[2]":    float2W,
		"float[3]":    float3W,
		"float[4]":    float4W,
		"ImWchar":     imWcharW,
		"ImGuiID":     imGuiIDW,
		"ImTextureID": imTextureIDW,
	}

	returnWrapperMap := map[string]returnWrapper{
		"bool":           boolReturnW,
		"const char*":    constCharReturnW,
		"const ImWchar*": constWCharPtrReturnW,
		"float":          floatReturnW,
		"int":            intReturnW,
	}

	type argOutput struct {
		ArgType string
		ArgDef  string
		VarName string
	}

	isEnum := func(argType string) bool {
		for _, en := range enumNames {
			if argType == en {
				return true
			}
		}

		return false
	}

	for _, f := range validFuncs {
		var args []string
		var argWrappers []argOutput

		shouldGenerate := false

		for _, a := range f.ArgsT {
			shouldGenerate = false

			if a.Name == "type" {
				a.Name = "typeArg"
			}

			if v, ok := argWrapperMap[a.Type]; ok {
				argType, argDef, varName := v(a)
				argWrappers = append(argWrappers, argOutput{
					ArgType: argType,
					ArgDef:  argDef,
					VarName: varName,
				})

				args = append(args, fmt.Sprintf("%s %s", a.Name, argType))

				shouldGenerate = true
			}

			if isEnum(a.Type) {
				args = append(args, fmt.Sprintf("%s %s", a.Name, a.Type))
				argWrappers = append(argWrappers, argOutput{
					VarName: fmt.Sprintf("C.%s(%s)", a.Type, a.Name),
				})

				shouldGenerate = true
			}

			if funk.ContainsString(structNames, a.Type) || funk.Contains(structNames, strings.TrimLeft(a.Type, "const ")) {
				aType := strings.TrimLeft(a.Type, "const ")
				args = append(args, fmt.Sprintf("%s %s", a.Name, aType))
				argWrappers = append(argWrappers, argOutput{
					VarName: fmt.Sprintf("C.%s(%s)", aType, a.Name),
				})

				shouldGenerate = true
			}

			if strings.HasSuffix(a.Type, "*") {
				pureType := strings.TrimLeft(a.Type, "const ")
				pureType = strings.TrimRight(pureType, "*")

				if funk.ContainsString(structNames, pureType) {
					args = append(args, fmt.Sprintf("%s *%s", a.Name, pureType))
					argWrappers = append(argWrappers, argOutput{
						VarName: fmt.Sprintf("(*C.%s)(%s)", pureType, a.Name),
					})

					shouldGenerate = true
				}
			}

			if !shouldGenerate {
				break
			}
		}

		if len(f.ArgsT) == 0 {
			shouldGenerate = true
		}

		if !shouldGenerate {
			fmt.Printf("%s%s\n", f.FuncName, f.Args)
			continue
		}

		// Generate function args
		argStmtFunc := func() string {
			var invokeStmt []string
			for _, aw := range argWrappers {
				invokeStmt = append(invokeStmt, aw.VarName)
				if len(aw.ArgDef) > 0 {
					sb.WriteString(fmt.Sprintf("%s\n\n", aw.ArgDef))
				}
			}

			return strings.Join(invokeStmt, ",")
		}

		if f.Ret == "void" {
			sb.WriteString(fmt.Sprintf("func %s(%s) {\n", f.FuncName, strings.Join(args, ",")))

			argInvokeStmt := argStmtFunc()

			sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.FuncName, argInvokeStmt))
			sb.WriteString("}\n\n")

			convertedFuncCount += 1
		} else {
			if rf, ok := returnWrapperMap[f.Ret]; ok {
				returnType, returnStmt := rf(f)

				sb.WriteString(fmt.Sprintf("func %s(%s) %s {\n", f.FuncName, strings.Join(args, ","), returnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf(returnStmt, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else if strings.HasSuffix(f.Ret, "*") && (funk.Contains(structNames, strings.TrimRight(f.Ret, "*")) || funk.Contains(structNames, strings.TrimRight(strings.TrimLeft(f.Ret, "const "), "*"))) {
				// return Im struct ptr
				pureReturnType := strings.TrimLeft(f.Ret, "const ")
				pureReturnType = strings.TrimRight(pureReturnType, "*")

				sb.WriteString(fmt.Sprintf("func %s(%s) %s {\n", f.FuncName, strings.Join(args, ","), "*"+pureReturnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("return (*%s)(%s)", pureReturnType, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else if f.Constructor {
				returnType := strings.Split(f.FuncName, "_")[0]

				if funk.ContainsString(structNames, "Im"+returnType) {
					returnType = "Im" + returnType
				} else if funk.ContainsString(structNames, "ImGui"+returnType) {
					returnType = "ImGui" + returnType
				} else {
					continue
				}

				sb.WriteString(fmt.Sprintf("func %s(%s) %s {\n", f.FuncName, strings.Join(args, ","), "*"+returnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("return (*%s)(C.%s(%s))", returnType, f.FuncName, argInvokeStmt))

				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else {
				fmt.Printf("%s%s -> %s\n", f.FuncName, f.Args, f.Ret)
			}
		}
	}

	fmt.Printf("Convert progress: %d/%d\n", convertedFuncCount, len(validFuncs))

	goFile, err := os.Create("funcs.go")
	if err != nil {
		panic(err.Error())
	}
	defer goFile.Close()

	_, _ = goFile.WriteString(sb.String())
}
