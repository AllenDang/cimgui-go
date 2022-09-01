package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/thoas/go-funk"
)

// Generate cpp wrapper and return valid functions
func generateCppWrapper(prefix, includePath string, funcDefs []FuncDef) []FuncDef {
	var validFuncs []FuncDef

	// Generate header
	var headerSb strings.Builder
	headerSb.WriteString(fmt.Sprintf(`#pragma once

#include "%s"

#ifdef __cplusplus
extern "C" {
#endif

`, includePath))

	var cppSb strings.Builder
	cppSb.WriteString(fmt.Sprintf(`#include "%s_wrapper.h"
#include "%s"

`, prefix, includePath))

	for i := 0; i < len(funcDefs); i++ {
		f := funcDefs[i]
		shouldSkip := false

		// Check func names
		if strings.HasPrefix(f.FuncName, "ImSpan") ||
			strings.HasPrefix(f.FuncName, "ImBitArray") ||
			strings.Contains(f.FuncName, "Storage") ||
			strings.Contains(f.FuncName, "TextRange") ||
			strings.Contains(f.FuncName, "ImVector") ||
			strings.Contains(f.FuncName, "Allocator") ||
			strings.Contains(f.FuncName, "__") {
			shouldSkip = true
			continue
		}

		// Process custom_type for arg
		for i, a := range f.ArgsT {
			if len(a.CustomType) > 0 {
				f.Args = strings.Replace(f.Args, a.Type, a.CustomType, -1)
				f.ArgsT[i].Type = a.CustomType
			}
		}

		// Check args
		for _, a := range f.ArgsT {
			if strings.Contains(a.Type, "const T*") ||
				strings.Contains(a.Type, "const T") ||
				a.Type == "va_list" ||
				(a.Name == "self" && a.Type == "ImVector*") {
				shouldSkip = true
				break
			}
		}

		if len(f.FuncName) == 0 || strings.Contains(f.Location, "internal") {
			shouldSkip = true
		}

		funcName := strings.TrimPrefix(f.FuncName, "ImGui")
		funcName = strings.TrimPrefix(funcName, "Im")
		funcName = strings.TrimPrefix(funcName, "ig")

		// Check lower case for function
		if unicode.IsLower(rune(funcName[0])) {
			shouldSkip = true
		}

		if shouldSkip {
			continue
		}

		// Check lower case member function
		funcParts := strings.Split(funcName, "_")
		if len(funcParts) == 2 && unicode.IsLower(rune(funcParts[1][0])) {
			funcName = funcParts[0] + "_" + string(unicode.ToUpper(rune(funcParts[1][0]))) + funcParts[1][1:]
		}

		// Transform some function names
		funcName = strings.Replace(funcName, "GetCursor", "GetDrawCursor", 1)
		funcName = strings.Replace(funcName, "SetCursor", "SetDrawCursor", 1)

		// Remove all ... arg
		f.Args = strings.Replace(f.Args, ",...", "", 1)
		// Remoe text_end arg
		f.Args = strings.Replace(f.Args, ",const char* text_end", "", 1)

		var argsT []ArgDef
		var actualCallArgs []string

		for _, a := range f.ArgsT {
			switch {
			case a.Name == "...":
				continue
			case a.Name == "text_end":
				actualCallArgs = append(actualCallArgs, "0")
				continue
			default:
				argsT = append(argsT, a)
				actualCallArgs = append(actualCallArgs, a.Name)
			}
		}

		f.ArgsT = argsT

		actualCallArgsStr := fmt.Sprintf("(%s)", strings.Join(actualCallArgs, ","))

		appendValidFunc := func() {
			validFuncs = append(validFuncs, FuncDef{
				Args:        f.Args,
				ArgsT:       f.ArgsT,
				FuncName:    funcName,
				Defaults:    f.Defaults,
				Location:    f.Location,
				Constructor: f.Constructor,
				Destructor:  f.Destructor,
				Ret:         f.Ret,
			})
		}

		if !f.Constructor && !f.Destructor {
			headerSb.WriteString(fmt.Sprintf("extern %s %s%s;\n", f.Ret, funcName, f.Args))

			if f.Ret == "void" {
				cppSb.WriteString(fmt.Sprintf("%s %s%s { %s%s; }\n", f.Ret, funcName, f.Args, f.FuncName, actualCallArgsStr))
			} else {
				cppSb.WriteString(fmt.Sprintf("%s %s%s { return %s%s; }\n", f.Ret, funcName, f.Args, f.FuncName, actualCallArgsStr))
			}

			appendValidFunc()
		}

		if f.Constructor {
			ret := strings.Split(f.FuncName, "_")[0]
			headerSb.WriteString(fmt.Sprintf("extern %s* %s%s;\n", ret, funcName, f.Args))
			cppSb.WriteString(fmt.Sprintf("%s* %s%s { return %s%s; }\n", ret, funcName, f.Args, f.FuncName, actualCallArgsStr))

			appendValidFunc()
		}

		if f.Destructor {
			headerSb.WriteString(fmt.Sprintf("extern void %s%s;\n", funcName, f.Args))
			cppSb.WriteString(fmt.Sprintf("void %s%s { %s%s; }\n", funcName, f.Args, f.FuncName, actualCallArgsStr))

			appendValidFunc()
		}
	}

	headerSb.WriteString(`
#ifdef __cplusplus
}
#endif
`)

	headerFile, err := os.Create(fmt.Sprintf("%s_wrapper.h", prefix))
	if err != nil {
		panic(err.Error())
	}
	defer headerFile.Close()

	_, err = headerFile.WriteString(headerSb.String())
	if err != nil {
		panic(err.Error())
	}

	cppFile, err := os.Create(fmt.Sprintf("%s_wrapper.cpp", prefix))
	if err != nil {
		panic(err.Error())
	}
	defer cppFile.Close()

	_, err = cppFile.WriteString(cppSb.String())
	if err != nil {
		panic(err.Error())
	}

	return validFuncs
}

func generateCppStructsAccessor(prefix string, validFuncs []FuncDef, structs []StructDef) []FuncDef {
	var structAccessorFuncs []FuncDef

	skipFuncNames := []string{
		"ImGuiIO_SetAppAcceptingEvents",
		"ImGuiDockNode_SetLocalFlags",
		"ImFontAtlas_SetTexID",
		"ImVec2_Getx",
		"ImVec2_Gety",
		"ImVec4_Getx",
		"ImVec4_Gety",
		"ImVec4_Getw",
		"ImVec4_Getz",
		"ImRect_GetMin",
		"ImRect_GetMax",
		"ImPlotColormapData_GetKeys",
	}

	// Add all valid function's name to skipFuncNames
	for _, f := range validFuncs {
		skipFuncNames = append(skipFuncNames, f.FuncName)
	}

	var sbHeader strings.Builder
	var sbCpp strings.Builder

	sbHeader.WriteString(fmt.Sprintf(`#pragma once

#include "%s_wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

`, prefix))

	sbCpp.WriteString(fmt.Sprintf(`
#include "%[1]s_wrapper.h"
#include "%[1]s_structs_accessor.h"

`, prefix))

	for _, s := range structs {
		for _, m := range s.Members {
			if strings.Contains(m.Name, "[") || strings.Contains(m.Type, "(") || strings.Contains(m.Type, "union") {
				continue
			}

			setterFuncName := fmt.Sprintf("%[1]s_Set%[2]s", s.Name, m.Name)
			if funk.ContainsString(skipFuncNames, setterFuncName) {
				continue
			}

			// Generate setter function
			structAccessorFuncs = append(structAccessorFuncs, FuncDef{
				Args: fmt.Sprintf("(%[1]s *%[2]s, %[3]s v)", s.Name, s.Name+"Ptr", m.Type),
				ArgsT: []ArgDef{
					{
						Name: s.Name + "Ptr",
						Type: s.Name,
					},
					{
						Name: "v",
						Type: m.Type,
					},
				},
				FuncName:     setterFuncName,
				Location:     "",
				Constructor:  false,
				Destructor:   false,
				StructSetter: true,
				Ret:          "void",
			})

			getterFuncName := fmt.Sprintf("%[1]s_Get%[2]s", s.Name, m.Name)
			if funk.ContainsString(skipFuncNames, getterFuncName) {
				continue
			}

			structAccessorFuncs = append(structAccessorFuncs, FuncDef{
				Args: fmt.Sprintf("(%[1]s *%[2]s)", s.Name, "self"),
				ArgsT: []ArgDef{
					{
						Name: "self",
						Type: s.Name,
					},
				},
				FuncName:     getterFuncName,
				Location:     "",
				Constructor:  false,
				Destructor:   false,
				StructSetter: false,
				StructGetter: true,
				Ret:          m.Type,
			})

			sbHeader.WriteString(fmt.Sprintf("extern void %[1]s_Set%[2]s(%[1]s *%[3]s, %[4]s v);\n", s.Name, m.Name, s.Name+"Ptr", m.Type))
			sbHeader.WriteString(fmt.Sprintf("extern %[4]s %[1]s_Get%[2]s(%[1]s *%[3]s);\n", s.Name, m.Name, "self", m.Type))

			sbCpp.WriteString(fmt.Sprintf("void %[1]s_Set%[2]s(%[1]s *%[3]s, %[4]s v) { %[3]s->%[2]s = v; }\n", s.Name, m.Name, s.Name+"Ptr", m.Type))
			sbCpp.WriteString(fmt.Sprintf("%[4]s %[1]s_Get%[2]s(%[1]s *%[3]s) { return %[3]s->%[2]s; }\n", s.Name, m.Name, "self", m.Type))
		}
	}

	sbHeader.WriteString(`
#ifdef __cplusplus
}
#endif
`)

	cppFile, err := os.Create(fmt.Sprintf("%s_structs_accessor.h", prefix))
	if err != nil {
		panic(err.Error())
	}
	defer cppFile.Close()

	_, err = cppFile.WriteString(sbHeader.String())
	if err != nil {
		panic(err.Error())
	}

	cppFile, err = os.Create(fmt.Sprintf("%s_structs_accessor.cpp", prefix))
	if err != nil {
		panic(err)
	}
	defer cppFile.Close()

	_, err = cppFile.WriteString(sbCpp.String())
	if err != nil {
		panic(err)
	}

	return structAccessorFuncs
}
