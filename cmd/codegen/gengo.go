package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/thoas/go-funk"
)

// Skip functions
// e.g. they are temporarily hard-coded
func skippedFuncs() []string {
	return []string{
		"InputText",
		"InputTextWithHint",
		"InputTextMultiline",
		"FontAtlas_GetTexDataAsAlpha8",
		"FontAtlas_GetTexDataAsAlpha8V",
		"FontAtlas_GetTexDataAsRGBA32",
		"FontAtlas_GetTexDataAsRGBA32V",
	}
}

func trimImGuiPrefix(id string) string {
	// don't trim prefixes for implot's ImAxis - it conflicts with ImGuIAxis (from imgui_internal.h)
	if strings.HasPrefix(id, "ImAxis") {
		return id
	}

	id = strings.TrimPrefix(id, "ImGui")
	id = strings.TrimPrefix(id, "Im")
	id = strings.TrimPrefix(id, "ig")
	return id
}

// Generate enums and return enum type names
func generateGoEnums(prefix string, enums []EnumDef) []string {
	var sb strings.Builder

	sb.WriteString(goPackageHeader)

	var enumNames []string
	for _, e := range enums {
		originalName := e.Name
		eName := strings.TrimSuffix(e.Name, "_")
		eName = trimImGuiPrefix(eName)

		enumNames = append(enumNames, eName)

		sb.WriteString(fmt.Sprintf("// original name: %s\n", originalName))
		sb.WriteString(fmt.Sprintf("type %s int\n", eName))
		sb.WriteString("const (\n")

		for _, v := range e.Values {
			vName := trimImGuiPrefix(v.Name)
			vName = strings.TrimSuffix(vName, "_")
			sb.WriteString(fmt.Sprintf("\t%s = %d\n", vName, v.Value))
		}

		sb.WriteString(")\n\n")
	}

	enumFile, err := os.Create(fmt.Sprintf("%s_enums.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer enumFile.Close()

	_, _ = enumFile.WriteString(sb.String())

	return enumNames
}

func generateGoStructs(prefix string, structs []StructDef) []string {
	var sb strings.Builder

	sb.WriteString(goPackageHeader)
	sb.WriteString(fmt.Sprintf(
		`// #include "%s_wrapper.h"
import "C"
import "unsafe"

`, prefix))

	// Save all struct name into a map
	var structNames []string

	for _, s := range structs {
		if shouldSkipStruct(s.Name) {
			continue
		}

		sb.WriteString(fmt.Sprintf(`type %[1]s uintptr

func (data %[1]s) handle() *C.%[1]s {
  return (*C.%[1]s)(unsafe.Pointer(data))
}

func (data %[1]s) c() C.%[1]s {
  return *(data.handle())
}

func new%[1]sFromC(cvalue C.%[1]s) %[1]s {
  return %[1]s(unsafe.Pointer(&cvalue))
}

`, s.Name))

		structNames = append(structNames, s.Name)
	}

	structFile, err := os.Create(fmt.Sprintf("%s_structs.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer structFile.Close()

	_, _ = structFile.WriteString(sb.String())

	return structNames
}

type goFuncsGenerator struct {
	prefix                 string
	structNames, enumNames []string

	sb                 strings.Builder
	convertedFuncCount int

	shouldGenerate bool
}

func generateGoFuncs(prefix string, validFuncs []FuncDef, enumNames []string, structNames []string) {
	generator := &goFuncsGenerator{
		prefix:      prefix,
		structNames: structNames,
		enumNames:   enumNames,
	}

	generator.writeFuncsFileHeader()

	for _, f := range validFuncs {
		// check whether the function shouldn't be skipped
		if funk.ContainsString(skippedFuncs(), f.FuncName) {
			continue
		}

		args, argWrappers := generator.generateFunccArgs(f)

		if len(f.ArgsT) == 0 {
			generator.shouldGenerate = true
		}

		if !generator.shouldGenerate {
			fmt.Printf("not generated: %s%s\n", f.FuncName, f.Args)
			continue
		} else {
			fmt.Printf("generated: %s%s\n", f.FuncName, f.Args)
		}

		skipStructs := []string{
			"ImVec1",
			"ImVec2",
			"ImVec2ih",
			"ImVec4",
			"ImColor",
			"ImRect",
			"StbUndoRecord",
			"StbUndoState",
			"StbTexteditRow",
		}

		funcSignatureFunc := func(funcName string, args []string, returnType string) string {
			funcParts := strings.Split(funcName, "_")
			typeName := funcParts[0]

			// Generate default param value hint
			var commentSb strings.Builder
			if len(f.Defaults) > 0 {
				commentSb.WriteString(fmt.Sprintf("// %s parameter default value hint:\n", funcName))

				// sort lexicographically for determenistic generation
				type defaultParam struct {
					name  string
					value string
				}
				defaults := make([]defaultParam, 0, len(f.Defaults))
				for n, v := range f.Defaults {
					defaults = append(defaults, defaultParam{name: n, value: v})
				}
				sort.Slice(defaults, func(i, j int) bool {
					return defaults[i].name < defaults[j].name
				})

				for _, p := range defaults {
					commentSb.WriteString(fmt.Sprintf("// %s: %s\n", p.name, p.value))
				}
			}

			if strings.Contains(funcName, "_") &&
				len(funcParts) > 1 &&
				len(args) > 0 && strings.Contains(args[0], "self ") &&
				!funk.ContainsString(skipStructs, typeName) {
				newFuncName := strings.TrimPrefix(funcName, typeName+"_")
				newArgs := args
				if len(newArgs) > 0 {
					newArgs = args[1:]
				}

				typeName = strings.TrimPrefix(args[0], "self ")
				return fmt.Sprintf("%sfunc (self %s) %s(%s) %s {\n", commentSb.String(), typeName, newFuncName, strings.Join(newArgs, ","), returnType)
			}

			return fmt.Sprintf("%sfunc %s(%s) %s {\n", commentSb.String(), funcName, strings.Join(args, ","), returnType)
		}

		switch {
		case f.NonUDT == 1:
			/*
				template:
				func FuncName(arg2 type2) typeOfArg1 {
					pOut := &typeOfArg1{}
					pOutArg, pOutFin := pOut.wrapped()
					defer pOutFin()
					C.FuncName(pOutArg, arg2)
					return *pOut
				}
			*/

			// find out the return type
			outArg := f.ArgsT[0]
			outArgT := strings.TrimSuffix(outArg.Type, "*")
			returnWrapper, err := getReturnTypeWrapperFunc(outArgT)
			if err != nil {
				fmt.Printf("Unknown return type \"%s\" in function %s\n", f.Ret, f.FuncName)
				continue
			}

			returnType, _ := returnWrapper()

			generator.sb.WriteString(funcSignatureFunc(f.FuncName, args[1:], returnType))

			// temporary out arg definition
			generator.sb.WriteString(fmt.Sprintf("%s := &%s{}\n", outArg.Name, returnType))

			argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

			// C function call
			generator.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.FuncName, argInvokeStmt))

			// return statement
			generator.sb.WriteString(fmt.Sprintf("return *%s", outArg.Name))

			generator.sb.WriteString("}\n\n")

			generator.convertedFuncCount += 1
		case f.Ret == "void":
			if f.StructSetter {
				funcParts := strings.Split(f.FuncName, "_")
				funcName := strings.TrimPrefix(f.FuncName, funcParts[0]+"_")
				if len(funcName) == 0 || !strings.HasPrefix(funcName, "Set") || funk.ContainsString(skipStructs, funcParts[0]) {
					continue
				}

				generator.sb.WriteString(fmt.Sprintf("func (self %[1]s) %[2]s(%[3]s) {\n", funcParts[0], funcName, strings.Join(args, ",")))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf("C.%s(self.handle(), %s)\n", f.FuncName, argInvokeStmt))
				generator.sb.WriteString("}\n\n")
			} else {
				generator.sb.WriteString(funcSignatureFunc(f.FuncName, args, ""))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.FuncName, argInvokeStmt))
				generator.sb.WriteString("}\n\n")
			}

			generator.convertedFuncCount += 1
		default:
			if rf, err := getReturnTypeWrapperFunc(f.Ret); err == nil {
				returnType, returnStmt := rf()

				generator.sb.WriteString(funcSignatureFunc(f.FuncName, args, returnType))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf(returnStmt, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if goEnumName := trimImGuiPrefix(f.Ret); funk.ContainsString(enumNames, goEnumName) {
				returnType := goEnumName

				generator.sb.WriteString(funcSignatureFunc(f.FuncName, args, returnType))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf("return %s(%s)", returnType, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if strings.HasSuffix(f.Ret, "*") && (funk.Contains(structNames, strings.TrimSuffix(f.Ret, "*")) || funk.Contains(structNames, strings.TrimSuffix(strings.TrimPrefix(f.Ret, "const "), "*"))) {
				// return Im struct ptr
				pureReturnType := strings.TrimPrefix(f.Ret, "const ")
				pureReturnType = strings.TrimSuffix(pureReturnType, "*")

				generator.sb.WriteString(funcSignatureFunc(f.FuncName, args, pureReturnType))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(%s))", pureReturnType, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if f.StructGetter && funk.ContainsString(structNames, f.Ret) {
				generator.sb.WriteString(funcSignatureFunc(f.FuncName, args, f.Ret))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf("return new%sFromC(C.%s(%s))", f.Ret, f.FuncName, argInvokeStmt))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if f.Constructor {
				parts := strings.Split(f.FuncName, "_")

				returnType := parts[0]

				if funk.ContainsString(structNames, "Im"+returnType) {
					returnType = "Im" + returnType
				} else if funk.ContainsString(structNames, "ImGui"+returnType) {
					returnType = "ImGui" + returnType
				} else {
					continue
				}

				suffix := ""
				if len(parts) > 2 {
					suffix = strings.Join(parts[2:], "")
				}

				newFuncName := "New" + returnType + suffix

				generator.sb.WriteString(fmt.Sprintf("func %s(%s) %s {\n", newFuncName, strings.Join(args, ","), returnType))

				argInvokeStmt := argStmtFunc(argWrappers, &generator.sb)

				generator.sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(C.%s(%s)))", returnType, f.FuncName, argInvokeStmt))

				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else {
				fmt.Printf("Unknown return type \"%s\" in function %s\n", f.Ret, f.FuncName)
			}
		}
	}

	fmt.Printf("Convert progress: %d/%d\n", generator.convertedFuncCount, len(validFuncs))

	goFile, err := os.Create(fmt.Sprintf("%s_funcs.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer goFile.Close()

	_, _ = goFile.WriteString(generator.sb.String())
}

func (g *goFuncsGenerator) writeFuncsFileHeader() {
	g.sb.WriteString(goPackageHeader)

	g.sb.WriteString(fmt.Sprintf(
		`// #include "extra_types.h"
// #include "%[1]s_structs_accessor.h"
// #include "%[1]s_wrapper.h"
import "C"
import "unsafe"

`, g.prefix))
}

func (g *goFuncsGenerator) isEnum(argType string) bool {
	for _, en := range g.enumNames {
		if argType == en {
			return true
		}
	}

	return false
}

func (g *goFuncsGenerator) generateFunccArgs(f FuncDef) (args []string, argWrappers []argOutput) {
	for i, a := range f.ArgsT {
		g.shouldGenerate = false

		if a.Name == "type" {
			a.Name = "typeArg"
		}

		if i == 0 && f.StructSetter {
			g.shouldGenerate = true
		}

		if f.StructGetter && funk.ContainsString(g.structNames, a.Type) {
			args = append(args, fmt.Sprintf("%s %s", a.Name, a.Type))
			argWrappers = append(argWrappers, argOutput{
				VarName: fmt.Sprintf("%s.handle()", a.Name),
			})

			g.shouldGenerate = true

			continue
		}

		if v, err := argWrapper(a.Type); err == nil {
			argType, argDef, varName := v(a)
			if goEnumName := trimImGuiPrefix(argType); g.isEnum(goEnumName) {
				argType = goEnumName
			}

			argWrappers = append(argWrappers, argOutput{
				ArgType: argType,
				ArgDef:  argDef,
				VarName: varName,
			})

			args = append(args, fmt.Sprintf("%s %s", a.Name, argType))

			g.shouldGenerate = true
			continue
		}

		if goEnumName := trimImGuiPrefix(a.Type); g.isEnum(goEnumName) {
			args = append(args, fmt.Sprintf("%s %s", a.Name, goEnumName))
			argWrappers = append(argWrappers, argOutput{
				VarName: fmt.Sprintf("C.%s(%s)", a.Type, a.Name),
			})

			g.shouldGenerate = true
			continue
		}

		if strings.HasSuffix(a.Type, "*") {
			pureType := strings.TrimPrefix(a.Type, "const ")
			pureType = strings.TrimSuffix(pureType, "*")

			if funk.ContainsString(g.structNames, pureType) {
				args = append(args, fmt.Sprintf("%s %s", a.Name, pureType))
				argWrappers = append(argWrappers, argOutput{
					VarName: fmt.Sprintf("%s.handle()", a.Name),
				})

				g.shouldGenerate = true
				continue
			}
		}

		if !g.shouldGenerate {
			fmt.Printf("Unknown argument type \"%s\" in function %s\n", a.Type, f.FuncName)
			break
		}
	}

	return args, argWrappers
}

type argOutput struct {
	ArgType string
	ArgDef  string
	VarName string
}

// Generate function args
func argStmtFunc(argWrappers []argOutput, sb *strings.Builder) string {
	var invokeStmt []string
	for _, aw := range argWrappers {
		invokeStmt = append(invokeStmt, aw.VarName)
		if len(aw.ArgDef) > 0 {
			sb.WriteString(fmt.Sprintf("%s\n\n", aw.ArgDef))
		}
	}

	return strings.Join(invokeStmt, ",")
}
