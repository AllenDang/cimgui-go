package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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
			strings.Contains(f.FuncName, "ImPlotTime") ||
			strings.Contains(f.FuncName, "MakeTime") ||
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

		if len(f.FuncName) == 0 || strings.Contains(f.Location, "imgui_internal") {
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

		// TODO: the majority of these wrappers, those without the V version,
		// could be skipped entirely. For those with V, the different versions
		// could be generated in Go code.
		// So, could be skip cimgui_wrapper.ccp/.h entirely?

		cWrapperFuncName := "wrap_" + funcName

		// Remove all ... arg
		f.Args = strings.Replace(f.Args, ",...", "", 1)
		// Remove text_end arg
		f.Args = strings.Replace(f.Args, ",const char* text_end", "", 1)

		var argsT []ArgDef
		var actualCallArgs []string

		for _, a := range f.ArgsT {
			f.AllCallArgs += a.Name + ","
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
		f.AllCallArgs = "(" + strings.TrimSuffix(f.AllCallArgs, ",") + ")"

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
				CWrapperFuncName: cWrapperFuncName + "V",
				AllCallArgs:      f.AllCallArgs,
			})

			// Add V as suffix to current function name
			funcName += "V"
			cWrapperFuncName += "V"
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
				CWrapperFuncName: cWrapperFuncName,
				AllCallArgs:      f.AllCallArgs,
			})
		}

		invokeFunctionName := f.FuncName
		if f.OriginalFuncName != "" {
			for _, ff := range validFuncs {
				if ff.FuncName == f.OriginalFuncName {
					invokeFunctionName = ff.CWrapperFuncName
					break
				}
			}
		}

		//cppSb.WriteString(fmt.Sprintf("// %#v\n", f))

		if f.AllCallArgs == actualCallArgsStr {
			cWrapperFuncName = f.FuncName
		} else if f.Constructor {
			headerSb.WriteString(fmt.Sprintf("extern %s* %s%s;\n", f.StName, cWrapperFuncName, f.Args))
			cppSb.WriteString(fmt.Sprintf("%s* %s%s { return %s%s; }\n", f.StName, cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
		} else if f.Destructor {
			headerSb.WriteString(fmt.Sprintf("extern void %s%s;\n", cWrapperFuncName, f.Args))
			cppSb.WriteString(fmt.Sprintf("void %s%s { %s%s; }\n", cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
		} else {
			headerSb.WriteString(fmt.Sprintf("extern %s %s%s;\n", f.Ret, cWrapperFuncName, f.Args))

			if f.Ret == "void" {
				cppSb.WriteString(fmt.Sprintf("%s %s%s { %s%s; }\n", f.Ret, cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
			} else {
				cppSb.WriteString(fmt.Sprintf("%s %s%s { return %s%s; }\n", f.Ret, cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
			}
		}
		appendValidFunc()
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

	skipFuncNames := map[string]bool{
		"ImGuiIO_SetAppAcceptingEvents": true,
		"ImGuiDockNode_SetLocalFlags":   true,
		"ImFontAtlas_SetTexID":          true,
		"ImVec2_Getx":                   true,
		"ImVec2_Gety":                   true,
		"ImVec4_Getx":                   true,
		"ImVec4_Gety":                   true,
		"ImVec4_Getw":                   true,
		"ImVec4_Getz":                   true,
		"ImRect_GetMin":                 true,
		"ImRect_GetMax":                 true,
		"ImPlotPoint_Setx":              true,
		"ImPlotPoint_Sety":              true,
		"ImPlotColormapData_GetKeys":    true,
	}

	// Add all valid function's name to skipFuncNames
	for _, f := range validFuncs {
		skipFuncNames[f.FuncName] = true
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
			if skipFuncNames[setterFuncName] {
				continue
			}

			// Generate setter function
			setterFuncDef := FuncDef{
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
				FuncName:         setterFuncName,
				CWrapperFuncName: "wrap_" + setterFuncName,
				Location:         "",
				Constructor:      false,
				Destructor:       false,
				StructSetter:     true,
				Ret:              "void",
			}
			structAccessorFuncs = append(structAccessorFuncs, setterFuncDef)

			getterFuncName := fmt.Sprintf("%[1]s_Get%[2]s", s.Name, m.Name)
			if skipFuncNames[getterFuncName] {
				continue
			}

			getterFuncDef := FuncDef{
				Args: fmt.Sprintf("(%[1]s *%[2]s)", s.Name, "self"),
				ArgsT: []ArgDef{
					{
						Name: "self",
						Type: s.Name,
					},
				},
				FuncName:         getterFuncName,
				CWrapperFuncName: "wrap_" + getterFuncName,
				Location:         "",
				Constructor:      false,
				Destructor:       false,
				StructSetter:     false,
				StructGetter:     true,
				Ret:              m.Type,
			}
			structAccessorFuncs = append(structAccessorFuncs, getterFuncDef)

			sbHeader.WriteString(fmt.Sprintf("extern void %s(%s *%s, %s v);\n", setterFuncDef.CWrapperFuncName, s.Name, s.Name+"Ptr", m.Type))
			sbHeader.WriteString(fmt.Sprintf("extern %s %s(%s *self);\n", m.Type, getterFuncDef.CWrapperFuncName, s.Name))

			sbCpp.WriteString(fmt.Sprintf("void %s(%s *%s, %s v) { %s->%s = v; }\n",
				setterFuncDef.CWrapperFuncName, s.Name, s.Name+"Ptr", m.Type, s.Name+"Ptr", m.Name))
			sbCpp.WriteString(fmt.Sprintf("%s %s(%s *self) { return self->%s; }\n",
				m.Type, getterFuncDef.CWrapperFuncName, s.Name, m.Name))
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
