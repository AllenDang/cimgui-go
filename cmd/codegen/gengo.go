package main

import (
	"fmt"
	"os"
	"strings"
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

type returnWrapper func(f FuncDef) (returnType string, returnStmt string)

func boolReturnW(f FuncDef) (returnType string, returnStmt string) {
	returnType = "bool"
	returnStmt = "return %s != C.bool(true)\n"
	return
}

func generateGoWrapper(validFuncs []FuncDef, enumNames []string) {
	var sb strings.Builder

	sb.WriteString(`package cimgui

// #include "cimgui_wrapper.h"
import "C"

`)

	argWrapperMap := map[string]typeWrapper{
		"const char*": constCharW,
		"size_t":      sizeTW,
		"float":       floatW,
		"float*":      floatPtrW,
		"bool":        boolW,
		"bool*":       boolPtrW,
	}

	returnWrapperMap := map[string]returnWrapper{
		"bool": boolReturnW,
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

			if !shouldGenerate {
				break
			}
		}

		if !shouldGenerate {
			fmt.Println("Err unknown arg type: ", f.FuncName)
			continue
		}

		// Generate function return void
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
		} else {
			if rf, ok := returnWrapperMap[f.Ret]; ok {
				returnType, returnStmt := rf(f)

				sb.WriteString(fmt.Sprintf("func %s(%s) %s {\n", f.FuncName, strings.Join(args, ","), returnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf(returnStmt, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				sb.WriteString("}\n\n")
			} else {
				fmt.Println("Err unkown return type: ", f.FuncName)
			}
		}
	}

	goFile, err := os.Create("funcs.go")
	if err != nil {
		panic(err.Error())
	}
	defer goFile.Close()

	_, _ = goFile.WriteString(sb.String())
}
