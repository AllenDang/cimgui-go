package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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
				a.Name == "..." ||
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
