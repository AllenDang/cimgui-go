package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/thoas/go-funk"
)

// Generate cpp wrapper and return valid functions
func generateCppWrapper(funcDefs []FuncDef) []FuncDef {
	var validFuncs []FuncDef

	// Generate header
	var headerSb strings.Builder
	headerSb.WriteString(`#pragma once

#include "cimgui/cimgui.h"

#ifdef __cplusplus
extern "C" {
#endif

`)

	var cppSb strings.Builder
	cppSb.WriteString(`#include "cimgui_wrapper.h"
#include "cimgui/cimgui.h"

`)

	for _, f := range funcDefs {
		shouldSkip := false

		// Check func names
		if f.FuncName == "ImSpan_destroy" ||
			f.FuncName == "ImBitArray_destroy" ||
			f.FuncName == "ImSpanAllocator_destroy" ||
			strings.Contains(f.FuncName, "Allocator") ||
			strings.Contains(f.FuncName, "__") {
			shouldSkip = true
			continue
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

		if shouldSkip {
			continue
		}

		if len(f.FuncName) == 0 || strings.Contains(f.Location, "internal") {
			continue
		}

		funcName := strings.Replace(f.FuncName, "ig", "", 1)
		funcName = strings.Replace(funcName, "ImGui", "", -1)
		funcName = strings.Replace(funcName, "Im", "", -1)

		// Check lower case for function
		if unicode.IsLower(rune(funcName[0])) {
			continue
		}

		// Check lower case member function
		funcParts := strings.Split(funcName, "_")
		if len(funcParts) == 2 && unicode.IsLower(rune(funcParts[1][0])) {
			continue
		}

		// Remove all ... arg
		f.Args = strings.Replace(f.Args, ",...", "", 1)
		var argsT []ArgDef
		for _, a := range f.ArgsT {
			if a.Name != "..." {
				argsT = append(argsT, a)
			}
		}

		f.ArgsT = argsT

		actualCallArgs := []string{}

		for _, a := range f.ArgsT {
			actualCallArgs = append(actualCallArgs, a.Name)
		}

		actualCallArgsStr := fmt.Sprintf("(%s)", strings.Join(actualCallArgs, ","))

		appendValidFunc := func() {
			validFuncs = append(validFuncs, FuncDef{
				Args:        f.Args,
				ArgsT:       f.ArgsT,
				FuncName:    funcName,
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

	headerFile, err := os.Create("cimgui_wrapper.h")
	if err != nil {
		panic(err.Error())
	}
	defer headerFile.Close()

	_, err = headerFile.WriteString(headerSb.String())
	if err != nil {
		panic(err.Error())
	}

	cppFile, err := os.Create("cimgui_wrapper.cpp")
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

func generateCppStructsAccessor(structs []StructDef) []FuncDef {
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
	}
	var sbHeader strings.Builder
	var sbCpp strings.Builder

	sbHeader.WriteString(`#pragma once

#include "cimgui_wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

`)

	sbCpp.WriteString(`
#include "cimgui_wrapper.h"
#include "cimgui_structs_accessor.h"

`)

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

	cppFile, err := os.Create("cimgui_structs_accessor.h")
	if err != nil {
		panic(err.Error())
	}
	defer cppFile.Close()

	_, err = cppFile.WriteString(sbHeader.String())
	if err != nil {
		panic(err.Error())
	}

	cppFile, err = os.Create("cimgui_structs_accessor.cpp")
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
