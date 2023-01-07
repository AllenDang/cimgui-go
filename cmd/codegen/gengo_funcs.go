package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"sort"
	"strings"
)

func generateGoFuncs(prefix string, validFuncs []FuncDef, enumNames []string, structNames []string) error {
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

		args, argWrappers := generator.generateFuncArgs(f)

		if len(f.ArgsT) == 0 {
			generator.shouldGenerate = true
		}

		// stop, when the function should not be generated
		if !generator.shouldGenerate {
			fmt.Printf("not generated: %s%s\n", f.FuncName, f.Args)
			continue
		} else {
			fmt.Printf("generated: %s%s\n", f.FuncName, f.Args)
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

			generator.sb.WriteString(generator.generateFuncDeclarationStmt(f.FuncName, args[1:], returnType, f))

			// temporary out arg definition
			generator.sb.WriteString(fmt.Sprintf("%s := &%s{}\n", outArg.Name, returnType))

			argInvokeStmt := generator.generateFuncBody(argWrappers)

			// C function call
			generator.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.FuncName, argInvokeStmt))

			// return statement
			generator.sb.WriteString(fmt.Sprintf("return *%s", outArg.Name))

			generator.sb.WriteString("}\n\n")

			generator.convertedFuncCount += 1
		case f.Ret == "void":
			if shouldContinue := generator.generateVoidFuncBody(f, args, argWrappers); !shouldContinue {
				continue
			}
		default:
			if rf, err := getReturnTypeWrapperFunc(f.Ret); err == nil {
				returnType, returnStmt := rf()

				generator.sb.WriteString(generator.generateFuncDeclarationStmt(f.FuncName, args, returnType, f))

				argInvokeStmt := generator.generateFuncBody(argWrappers)

				generator.sb.WriteString(fmt.Sprintf(returnStmt, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if goEnumName := trimImGuiPrefix(f.Ret); funk.ContainsString(enumNames, goEnumName) {
				returnType := goEnumName

				generator.sb.WriteString(generator.generateFuncDeclarationStmt(f.FuncName, args, returnType, f))

				argInvokeStmt := generator.generateFuncBody(argWrappers)

				generator.sb.WriteString(fmt.Sprintf("return %s(%s)", returnType, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if strings.HasSuffix(f.Ret, "*") && (funk.Contains(structNames, strings.TrimSuffix(f.Ret, "*")) || funk.Contains(structNames, strings.TrimSuffix(strings.TrimPrefix(f.Ret, "const "), "*"))) {
				// return Im struct ptr
				pureReturnType := strings.TrimPrefix(f.Ret, "const ")
				pureReturnType = strings.TrimSuffix(pureReturnType, "*")

				generator.sb.WriteString(generator.generateFuncDeclarationStmt(f.FuncName, args, pureReturnType, f))

				argInvokeStmt := generator.generateFuncBody(argWrappers)

				generator.sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(%s))", pureReturnType, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				generator.sb.WriteString("}\n\n")

				generator.convertedFuncCount += 1
			} else if f.StructGetter && funk.ContainsString(structNames, f.Ret) {
				generator.sb.WriteString(generator.generateFuncDeclarationStmt(f.FuncName, args, f.Ret, f))

				argInvokeStmt := generator.generateFuncBody(argWrappers)

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

				argInvokeStmt := generator.generateFuncBody(argWrappers)

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

	_, err = goFile.WriteString(generator.sb.String())
	if err != nil {
		return fmt.Errorf("failed to write content of GO file: %w", err)
	}

	return nil
}

// goFuncsGenerator is an internal state of GO funcs' generator
type goFuncsGenerator struct {
	prefix                 string
	structNames, enumNames []string

	sb                 strings.Builder
	convertedFuncCount int

	shouldGenerate bool
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

func (g *goFuncsGenerator) generateFuncArgs(f FuncDef) (args []string, argWrappers []argOutput) {
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

// this method is responsible for createing a function declaration statement.
// it takes function name, list of arguments and return type and returns go statement.
// e.g.: func (self *ImGuiType) FuncName(arg1 type1, arg2 type2) returnType {
func (g *goFuncsGenerator) generateFuncDeclarationStmt(funcName string, args []string, returnType string, f FuncDef) (functionDeclaration string) {
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
		!funk.ContainsString(skippedStructs(), typeName) {
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

func (g *goFuncsGenerator) generateVoidFuncBody(f FuncDef, args []string, argWrappers []argOutput) (shouldContinue bool) {
	if f.StructSetter {
		funcParts := strings.Split(f.FuncName, "_")
		funcName := strings.TrimPrefix(f.FuncName, funcParts[0]+"_")
		if len(funcName) == 0 || !strings.HasPrefix(funcName, "Set") || funk.ContainsString(skippedStructs(), funcParts[0]) {
			return false
		}

		g.sb.WriteString(fmt.Sprintf("func (self %[1]s) %[2]s(%[3]s) {\n", funcParts[0], funcName, strings.Join(args, ",")))

		argInvokeStmt := g.generateFuncBody(argWrappers)

		g.sb.WriteString(fmt.Sprintf("C.%s(self.handle(), %s)\n", f.FuncName, argInvokeStmt))
		g.sb.WriteString("}\n\n")
	} else {
		g.sb.WriteString(g.generateFuncDeclarationStmt(f.FuncName, args, "", f))

		argInvokeStmt := g.generateFuncBody(argWrappers)

		g.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.FuncName, argInvokeStmt))
		g.sb.WriteString("}\n\n")
	}

	g.convertedFuncCount += 1

	return true
}

// Generate function body
// and returns function call arguments
// e.g.:
// it will write the following into the buffer:
func (g *goFuncsGenerator) generateFuncBody(argWrappers []argOutput) string {
	var invokeStmt []string
	for _, aw := range argWrappers {
		invokeStmt = append(invokeStmt, aw.VarName)
		if len(aw.ArgDef) > 0 {
			g.sb.WriteString(fmt.Sprintf("%s\n\n", aw.ArgDef))
		}
	}

	return strings.Join(invokeStmt, ",")
}
