package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/thoas/go-funk"
)

// Returns if should export func
func shouldExportFunc(funcName string) bool {
	if unicode.IsUpper(rune(funcName[0])) {
		return true
	}
	if strings.HasPrefix(funcName, "ig") {
		return len(funcName) > 2 && unicode.IsUpper(rune(funcName[2]))
	}
	return false
}

// Generate cpp wrapper and return valid functions
func generateCppWrapper(prefix, includePath string, funcDefs []FuncDef) ([]FuncDef, error) {
	var validFuncs []FuncDef

	// Generate header
	var headerSb strings.Builder
	headerSb.WriteString(cppFileHeader)
	headerSb.WriteString(fmt.Sprintf(`#pragma once

#include "%s"

#ifdef __cplusplus
extern "C" {
#endif

`, includePath))

	var cppSb strings.Builder
	cppSb.WriteString(cppFileHeader)
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

		funcName := f.FuncName

		// Check lower case for function
		if !shouldExportFunc(funcName) {
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

		// Generate shotter function which omits the default args
		// Skip functions as function pointer arg
		shouldSkip = false
		for _, a := range f.ArgsT {
			if strings.Contains(a.Type, "(") {
				shouldSkip = true
			}
		}

		if len(f.Defaults) > 0 && !shouldSkip && !f.Constructor {
			var newArgsT []ArgDef
			var invocationArgs []string

			for _, a := range f.ArgsT {
				invocationArgs = append(invocationArgs, a.Name)

				shouldIgnore := false
				for k := range f.Defaults {
					if a.Name == k {
						shouldIgnore = true
						break
					}
				}

				if !shouldIgnore {
					newArgsT = append(newArgsT, a)
				}
			}

			var newArgs []string
			for _, a := range newArgsT {
				t := a.Type
				n := a.Name

				// Process array type
				if strings.Contains(t, "[") {
					parts := strings.Split(t, "[")
					t = parts[0]
					n = n + "[" + strings.Join(parts[1:], "")
				}

				newArgs = append(newArgs, fmt.Sprintf("%s %s", t, n))
			}

			// Generate invocation stmt
			invocationStmt := fmt.Sprintf("(%s)", strings.Join(invocationArgs, ","))

			for k, v := range f.Defaults {
				if v == "ImVec2(0,0)" || v == "ImVec2(0.0f,0.0f)" {
					v = "(ImVec2){.x=0, .y=0}"
				}

				if v == "ImVec2(1,1)" {
					v = "(ImVec2){.x=1, .y=1}"
				}

				if v == "ImVec2(1,0)" {
					v = "(ImVec2){.x=1, .y=0}"
				}

				if v == "ImVec2(0,1)" {
					v = "(ImVec2){.x=0, .y=1}"
				}

				if v == "ImVec2(-1,0)" {
					v = "(ImVec2){.x=-1, .y=0}"
				}

				if v == "ImVec2(-FLT_MIN,0)" {
					v = "(ImVec2){.x=-1*igGET_FLT_MIN(), .y=0}"
				}

				if v == "ImVec4(0,0,0,0)" {
					v = "(ImVec4){.x=0, .y=0, .z=0, .w=0}"
				}

				if v == "ImVec4(1,1,1,1)" {
					v = "(ImVec4){.x=1, .y=1, .z=1, .w=1}"
				}

				if v == "ImVec4(0,0,0,-1)" {
					v = "(ImVec4){.x=0, .y=0, .z=0, .w=-1}"
				}

				if v == "ImPlotPoint(0,0)" {
					v = "(ImPlotPoint){.x=0, .y=0}"
				}

				if v == "ImPlotPoint(1,1)" {
					v = "(ImPlotPoint){.x=1, .y=1}"
				}

				if v == "FLT_MAX" {
					v = "igGET_FLT_MAX()"
				}

				if v == "((void*)0)" {
					v = "NULL"
				}

				if strings.Contains(invocationStmt, ","+k) {
					invocationStmt = strings.Replace(invocationStmt, ","+k, ","+v, 1)
				} else {
					invocationStmt = strings.Replace(invocationStmt, k, v, 1)
				}
			}

			// Generate new function
			funcDefs = append(funcDefs, FuncDef{
				FuncName:         funcName,
				OriginalFuncName: funcName + "V",
				Args:             fmt.Sprintf("(%s)", strings.Join(newArgs, ",")),
				ArgsT:            newArgsT,
				InvocationStmt:   invocationStmt,
				Defaults:         nil,
				Location:         f.Location,
				Constructor:      f.Constructor,
				Destructor:       f.Destructor,
				StructSetter:     false,
				StructGetter:     false,
				Ret:              f.Ret,
				StName:           f.StName,
				NonUDT:           f.NonUDT,
			})

			// Add V as suffix to current function name
			funcName += "V"
		}

		actualCallArgsStr := fmt.Sprintf("(%s)", strings.Join(actualCallArgs, ","))

		if len(f.InvocationStmt) > 0 {
			actualCallArgsStr = f.InvocationStmt
		}

		appendValidFunc := func() {
			validFuncs = append(validFuncs, FuncDef{
				Args:             f.Args,
				ArgsT:            f.ArgsT,
				FuncName:         funcName,
				OriginalFuncName: f.OriginalFuncName,
				Defaults:         f.Defaults,
				Location:         f.Location,
				Constructor:      f.Constructor,
				Destructor:       f.Destructor,
				InvocationStmt:   f.InvocationStmt,
				Ret:              f.Ret,
				StName:           f.StName,
				NonUDT:           f.NonUDT,
			})
		}

		invokeFunctionName := f.FuncName
		if len(f.OriginalFuncName) > 0 {
			invokeFunctionName = "wrap_" + f.OriginalFuncName
		}

		if !f.Constructor && !f.Destructor {
			headerSb.WriteString(fmt.Sprintf("extern %s %s%s;\n", f.Ret, "wrap_"+funcName, f.Args))

			if f.Ret == "void" {
				cppSb.WriteString(fmt.Sprintf("%s %s%s { %s%s; }\n", f.Ret, "wrap_"+funcName, f.Args, invokeFunctionName, actualCallArgsStr))
			} else {
				cppSb.WriteString(fmt.Sprintf("%s %s%s { return %s%s; }\n", f.Ret, "wrap_"+funcName, f.Args, invokeFunctionName, actualCallArgsStr))
			}

			appendValidFunc()
		}

		if f.Constructor {
			ret := f.StName
			headerSb.WriteString(fmt.Sprintf("extern %s* %s%s;\n", ret, "wrap_"+funcName, f.Args))
			cppSb.WriteString(fmt.Sprintf("%s* %s%s { return %s%s; }\n", ret, "wrap_"+funcName, f.Args, invokeFunctionName, actualCallArgsStr))

			appendValidFunc()
		}

		if f.Destructor {
			headerSb.WriteString(fmt.Sprintf("extern void %s%s;\n", "wrap_"+funcName, f.Args))
			cppSb.WriteString(fmt.Sprintf("void %s%s { %s%s; }\n", "wrap_"+funcName, f.Args, invokeFunctionName, actualCallArgsStr))

			appendValidFunc()
		}
	}

	headerSb.WriteString(`
#ifdef __cplusplus
}
#endif
`)

	headerPath := fmt.Sprintf("%s_wrapper.h", prefix)
	headerFile, err := os.Create(headerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create header file %v: %v", headerPath, err)
	}

	defer headerFile.Close()

	_, err = headerFile.WriteString(headerSb.String())
	if err != nil {
		return nil, fmt.Errorf("failed to write header file %v: %v", headerPath, err)
	}

	cppPath := fmt.Sprintf("%s_wrapper.cpp", prefix)
	cppFile, err := os.Create(cppPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create cpp file %v: %v", cppPath, err)
	}

	defer cppFile.Close()

	_, err = cppFile.WriteString(cppSb.String())
	if err != nil {
		return nil, fmt.Errorf("failed to write cpp file %v: %v", cppPath, err)
	}

	return validFuncs, nil
}

func generateCppStructsAccessor(prefix string, validFuncs []FuncDef, structs []StructDef) ([]FuncDef, error) {
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
		"ImPlotPoint_Setx",
		"ImPlotPoint_Sety",
		"ImPlotColormapData_GetKeys",
	}

	// Add all valid function's name to skipFuncNames
	for _, f := range validFuncs {
		skipFuncNames = append(skipFuncNames, f.FuncName)
	}

	var sbHeader, sbCpp strings.Builder

	sbHeader.WriteString(cppFileHeader)
	sbCpp.WriteString(cppFileHeader)

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

			sbHeader.WriteString(fmt.Sprintf("extern void wrap_%[1]s_Set%[2]s(%[1]s *%[3]s, %[4]s v);\n", s.Name, m.Name, s.Name+"Ptr", m.Type))
			sbHeader.WriteString(fmt.Sprintf("extern %[4]s wrap_%[1]s_Get%[2]s(%[1]s *%[3]s);\n", s.Name, m.Name, "self", m.Type))

			sbCpp.WriteString(fmt.Sprintf("void wrap_%[1]s_Set%[2]s(%[1]s *%[3]s, %[4]s v) { %[3]s->%[2]s = v; }\n", s.Name, m.Name, s.Name+"Ptr", m.Type))
			sbCpp.WriteString(fmt.Sprintf("%[4]s wrap_%[1]s_Get%[2]s(%[1]s *%[3]s) { return %[3]s->%[2]s; }\n", s.Name, m.Name, "self", m.Type))
		}
	}

	sbHeader.WriteString(`
#ifdef __cplusplus
}
#endif
`)

	headerPath := fmt.Sprintf("%s_structs_accessor.h", prefix)
	headerFile, err := os.Create(headerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create cpp file %v: %v", headerPath, err)
	}

	defer headerFile.Close()

	_, err = headerFile.WriteString(sbHeader.String())
	if err != nil {
		return nil, fmt.Errorf("failed to write to file %v: %v", headerPath, err)
	}

	cppPath := fmt.Sprintf("%s_structs_accessor.cpp", prefix)
	cppFile, err := os.Create(cppPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create cpp file %v: %v", cppPath, err)
	}

	defer cppFile.Close()

	_, err = cppFile.WriteString(sbCpp.String())
	if err != nil {
		return nil, fmt.Errorf("failed to write to file %v: %v", cppPath, err)
	}

	return structAccessorFuncs, nil
}
