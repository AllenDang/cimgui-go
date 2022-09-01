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
func generateGoEnums(prefix string, enums []EnumDef) []string {
	var sb strings.Builder

	sb.WriteString("package cimgui\n\n")

	var enumNames []string
	for _, e := range enums {
		eName := strings.TrimSuffix(e.Name, "_")

		enumNames = append(enumNames, eName)

		sb.WriteString(fmt.Sprintf("type %s int\n", eName))
		sb.WriteString("const (\n")

		for _, v := range e.Values {
			sb.WriteString(fmt.Sprintf("\t%s = %d\n", v.Name, v.Value))
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

	sb.WriteString(fmt.Sprintf(`package cimgui

// #include "%s_wrapper.h"
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

func generateGoFuncs(prefix string, validFuncs []FuncDef, enumNames []string, structNames []string) {
	var sb strings.Builder
	convertedFuncCount := 0

	sb.WriteString(fmt.Sprintf(`package cimgui

// #include "extra_types.h"
// #include "%[1]s_structs_accessor.h"
// #include "%[1]s_wrapper.h"
import "C"
import "unsafe"

`, prefix))

	argWrapperMap := map[string]typeWrapper{
		"char*":                    constCharW,
		"const char*":              constCharW,
		"unsigned char":            ucharW,
		"unsigned char**":          uCharPtrW,
		"size_t":                   sizeTW,
		"size_t*":                  sizeTPtrW,
		"float":                    floatW,
		"float*":                   floatPtrW,
		"const float*":             floatArrayW,
		"short":                    shortW,
		"unsigned short":           ushortW,
		"ImU8":                     u8W,
		"ImU16":                    u16W,
		"ImU64":                    u64W,
		"const ImU64*":             uint64ArrayW,
		"ImS8":                     s8W,
		"ImS16":                    s16W,
		"ImS32":                    s32W,
		"const ImS64*":             int64ArrayW,
		"int":                      intW,
		"int*":                     intPtrW,
		"unsigned int":             uintW,
		"double":                   doubleW,
		"double*":                  doublePtrW,
		"bool":                     boolW,
		"bool*":                    boolPtrW,
		"int[2]":                   int2W,
		"int[3]":                   int3W,
		"int[4]":                   int4W,
		"ImU32":                    u32W,
		"float[2]":                 float2W,
		"float[3]":                 float3W,
		"float[4]":                 float4W,
		"ImWchar":                  imWcharW,
		"const ImWchar*":           imWcharPtrW,
		"ImGuiID":                  imGuiIDW,
		"ImTextureID":              imTextureIDW,
		"ImDrawIdx":                imDrawIdxW,
		"ImGuiTableColumnIdx":      imTableColumnIdxW,
		"ImGuiTableDrawChannelIdx": imTableDrawChannelIdxW,
		"void*":                    voidPtrW,
		"const void*":              voidPtrW,
		"const ImVec2":             imVec2W,
		"const ImVec2*":            imVec2PtrW,
		"ImVec2":                   imVec2W,
		"ImVec2*":                  imVec2PtrW,
		"const ImVec4":             imVec4W,
		"const ImVec4*":            imVec4PtrW,
		"ImVec4":                   imVec4W,
		"ImVec4*":                  imVec4PtrW,
		"ImColor*":                 imColorPtrW,
		"ImRect":                   imRectW,
	}

	returnWrapperMap := map[string]returnWrapper{
		"bool":                     boolReturnW,
		"const char*":              constCharReturnW,
		"const ImWchar*":           constWCharPtrReturnW,
		"float":                    floatReturnW,
		"double":                   doubleReturnW,
		"int":                      intReturnW,
		"unsigned int":             uintReturnW,
		"short":                    intReturnW,
		"ImS8":                     intReturnW,
		"ImS16":                    intReturnW,
		"ImS32":                    intReturnW,
		"ImU8":                     uintReturnW,
		"ImU16":                    uintReturnW,
		"ImU32":                    u32ReturnW,
		"ImU64":                    uint64ReturnW,
		"ImVec4":                   imVec4ReturnW,
		"const ImVec4*":            imVec4PtrReturnW,
		"ImGuiID":                  idReturnW,
		"ImTextureID":              textureIdReturnW,
		"ImVec2":                   imVec2ReturnW,
		"ImRect":                   imRectReturnW,
		"ImGuiTableColumnIdx":      imTableColumnIdxReturnW,
		"ImGuiTableDrawChannelIdx": imTableDrawChannelIdxReturnW,
		"void*":                    voidPtrReturnW,
		"size_t":                   doubleReturnW,
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

		for i, a := range f.ArgsT {
			shouldGenerate = false

			if a.Name == "type" {
				a.Name = "typeArg"
			}

			if i == 0 && f.StructSetter {
				shouldGenerate = true
			}

			if f.StructGetter && funk.ContainsString(structNames, a.Type) {
				args = append(args, fmt.Sprintf("%s %s", a.Name, a.Type))
				argWrappers = append(argWrappers, argOutput{
					VarName: fmt.Sprintf("%s.handle()", a.Name),
				})
				shouldGenerate = true
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

			if strings.HasSuffix(a.Type, "*") {
				pureType := strings.TrimPrefix(a.Type, "const ")
				pureType = strings.TrimSuffix(pureType, "*")

				if funk.ContainsString(structNames, pureType) {
					args = append(args, fmt.Sprintf("%s %s", a.Name, pureType))
					argWrappers = append(argWrappers, argOutput{
						VarName: fmt.Sprintf("%s.handle()", a.Name),
					})

					shouldGenerate = true
				}
			}

			if !shouldGenerate {
				fmt.Println("Unknown arg: ", a.Type)
				break
			}
		}

		if len(f.ArgsT) == 0 {
			shouldGenerate = true
		}

		if !shouldGenerate {
			// fmt.Printf("%s%s\n", f.FuncName, f.Args)
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

				for n, v := range f.Defaults {
					commentSb.WriteString(fmt.Sprintf("// %s: %s\n", n, v))
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

		if f.Ret == "void" {
			if f.StructSetter {
				funcParts := strings.Split(f.FuncName, "_")
				funcName := strings.TrimPrefix(f.FuncName, funcParts[0]+"_")
				if len(funcName) == 0 || !strings.HasPrefix(funcName, "Set") || funk.ContainsString(skipStructs, funcParts[0]) {
					continue
				}

				sb.WriteString(fmt.Sprintf("func (self %[1]s) %[2]s(%[3]s) {\n", funcParts[0], funcName, strings.Join(args, ",")))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("C.%s(self.handle(), %s)\n", f.FuncName, argInvokeStmt))
				sb.WriteString("}\n\n")
			} else {
				sb.WriteString(funcSignatureFunc(f.FuncName, args, ""))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.FuncName, argInvokeStmt))
				sb.WriteString("}\n\n")
			}

			convertedFuncCount += 1
		} else {
			if rf, ok := returnWrapperMap[f.Ret]; ok {
				returnType, returnStmt := rf(f)

				sb.WriteString(funcSignatureFunc(f.FuncName, args, returnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf(returnStmt, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else if funk.ContainsString(enumNames, f.Ret) {
				returnType := f.Ret

				sb.WriteString(funcSignatureFunc(f.FuncName, args, returnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("return %s(%s)", f.Ret, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else if strings.HasSuffix(f.Ret, "*") && (funk.Contains(structNames, strings.TrimSuffix(f.Ret, "*")) || funk.Contains(structNames, strings.TrimSuffix(strings.TrimPrefix(f.Ret, "const "), "*"))) {
				// return Im struct ptr
				pureReturnType := strings.TrimPrefix(f.Ret, "const ")
				pureReturnType = strings.TrimSuffix(pureReturnType, "*")

				sb.WriteString(funcSignatureFunc(f.FuncName, args, pureReturnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(%s))", pureReturnType, fmt.Sprintf("C.%s(%s)", f.FuncName, argInvokeStmt)))
				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else if f.StructGetter && funk.ContainsString(structNames, f.Ret) {
				sb.WriteString(funcSignatureFunc(f.FuncName, args, f.Ret))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("return new%sFromC(C.%s(%s))", f.Ret, f.FuncName, argInvokeStmt))
				sb.WriteString("}\n\n")

				convertedFuncCount += 1
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

				sb.WriteString(fmt.Sprintf("func %s(%s) %s {\n", newFuncName, strings.Join(args, ","), returnType))

				argInvokeStmt := argStmtFunc()

				sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(C.%s(%s)))", returnType, f.FuncName, argInvokeStmt))

				sb.WriteString("}\n\n")

				convertedFuncCount += 1
			} else {
				fmt.Println("Unknown ret: ", f.Ret)
			}
		}
	}

	fmt.Printf("Convert progress: %d/%d\n", convertedFuncCount, len(validFuncs))

	goFile, err := os.Create(fmt.Sprintf("%s_funcs.go", prefix))
	if err != nil {
		panic(err.Error())
	}
	defer goFile.Close()

	_, _ = goFile.WriteString(sb.String())
}
