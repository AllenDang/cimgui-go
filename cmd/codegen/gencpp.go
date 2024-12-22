package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// Name of argument in cpp/go files.
// It is used by functions that has text and text_end arguments.
// In this case text_end is replaced by this argument (of type int)
const textLenRegisteredName = "text_len"

// Generate cpp wrapper and return valid functions
func generateCppWrapper(includePath string, funcDefs []FuncDef, ctx *Context) ([]FuncDef, error) {
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

	cppSb := &strings.Builder{}
	cppSb.WriteString(cppFileHeader)
	fmt.Fprintf(cppSb, `#include "wrapper.h"
#include "%s"

`, includePath)

	// Note for future generations: can't replace with range, because we edit funcDefs later
	for i := 0; i < len(funcDefs); i++ {
		f := funcDefs[i]

		shouldSkip := false

		// Check func names (some of them are arbitrarly skipped)
		if HasPrefix(f.FuncName, "ImSpan") ||
			HasPrefix(f.FuncName, "ImBitArray") {
			shouldSkip = true
			continue
		}

		// Process custom_type for arg
		for i, a := range f.ArgsT {
			if len(a.CustomType) > 0 {
				f.Args = ReplaceAll(f.Args, a.Type, a.CustomType)
				f.ArgsT[i].Type = a.CustomType
			}
		}

		// Check args (some arg formats are skipped)
		for _, a := range f.ArgsT {
			if Contains(a.Type, "const T*") ||
				Contains(a.Type, "const T") ||
				a.Type == "va_list" ||
				(a.Name == "self" && a.Type == "ImVector*") {
				shouldSkip = true
				break
			}
		}

		// check if func name is valid
		if len(f.FuncName) == 0 {
			shouldSkip = true
		}

		if shouldSkip {
			continue
		}

		funcName := f.FuncName

		// Check lower case member function
		funcParts := Split(funcName, "_")
		if len(funcParts) == 2 && unicode.IsLower(rune(funcParts[1][0])) {
			funcName = funcParts[0] + "_" + CIdentifier(unicode.ToUpper(rune(funcParts[1][0]))) + funcParts[1][1:]
		}

		// TODO: the majority of these wrappers, those without the V version,
		// could be skipped entirely. For those with V, the different versions
		// could be generated in Go code.
		// So, could be skip cimgui_wrapper.ccp/.h entirely?

		cWrapperFuncName := "wrap_" + f.FuncName

		// Remove all ... arg
		f.Args = strings.Replace(f.Args, ",...", "", 1)
		// Remove text_end arg
		f.Args = strings.Replace(f.Args, ",const char* text_end_", fmt.Sprintf(",const int %s", textLenRegisteredName), 1) // sometimes happens in cimmarkdown
		f.Args = strings.Replace(f.Args, ",const char* text_end", fmt.Sprintf(",const int %s", textLenRegisteredName), 1)
		ret := f.Ret
		if ret == "void*" {
			ret = "uintptr_t"
		}

		var argsT []ArgDef
		var actualCallArgs []CIdentifier

		for _, a := range f.ArgsT {
			f.AllCallArgs += a.Name + ","
			switch {
			case a.Name == "...":
				continue
			case a.Name == "text_end", a.Name == "text_end_":
				// chck if there is `text` argument
				var found bool
				for _, aa := range f.ArgsT {
					if aa.Name == "text" {
						found = true
						break
					}
				}
				if found {
					argsT = append(argsT, ArgDef{
						Name: "text_len",
						Type: "const int",
					})
					actualCallArgs = append(actualCallArgs, "(text_len > 0) ? text + text_len*sizeof(char) : 0")
				} else {
					f.Args = strings.Replace(f.Args, ",const int text_len", "", 1)
					actualCallArgs = append(actualCallArgs, "0")
					continue
				}
			// this is here, because of a BUG in GO that throws a fatal panic
			// when we attempt to print unsafe.Pointer which is valid for C but is not present
			// on the GO stack. See https://go.dev/play/p/d09eyqUlVV0
			// see:https://github.com/golang/go/issues/64467
			// case Contains(a.Type, "void*"):
			case a.Type == "void*" || a.Type == "const void*":
				actualCallArgs = append(actualCallArgs, CIdentifier(fmt.Sprintf("(%s)(uintptr_t)%s", a.Type, a.Name)))
				a.Type = "uintptr_t"
				// do some trick: we can't replace void* in a.Args immediaely by using replace, because sometimes
				// args are of form (void* user_data, (void* user_data)(int something))
				// we need to replace only "plain" arguments (not function-types)
				// we will split Args by "," and check each argument to be equal to ("void* ", a.Name)
				newArgs := TrimPrefix(f.Args, "(")
				newArgs = TrimSuffix(newArgs, ")")
				newArgsList := Split(newArgs, ",")
				for i, arg := range newArgsList {
					switch arg {
					case fmt.Sprintf("void* %s", a.Name):
						newArgsList[i] = fmt.Sprintf("uintptr_t %s", a.Name)
					case fmt.Sprintf("const void* %s", a.Name):
						newArgsList[i] = fmt.Sprintf("const uintptr_t %s", a.Name)
					}
				}

				f.Args = fmt.Sprintf("(%s)", Join(newArgsList, ","))

				argsT = append(argsT, a)
			default:
				argsT = append(argsT, a)
				actualCallArgs = append(actualCallArgs, a.Name)
			}
		}

		f.AllCallArgs = "(" + TrimSuffix(f.AllCallArgs, ",") + ")"
		f.ArgsT = argsT

		//f.Args = "("
		//for _, a := range argsT {
		//	f.Args += fmt.Sprintf("%s %s,", a.Type, a.Name)
		//}
		//f.Args = TrimSuffix(f.Args, ",")
		//f.Args += ")"

		// Generate shorter function which omits the default args
		// Skip functions as function pointer arg
		shouldSkip = false
		for _, a := range f.ArgsT {
			if Contains(a.Type, "(") {
				shouldSkip = true
			}
		}

		if len(f.Defaults) > 0 && !shouldSkip && !f.Constructor {
			var newArgsT []ArgDef
			var invocationArgs []CIdentifier

			for _, a := range f.ArgsT {
				invocationArgs = append(invocationArgs, a.Name)

				shouldIgnore := false
				for k := range f.Defaults {
					if string(a.Name) == k {
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
				if Contains(t, "[") {
					parts := Split(t, "[")
					t = parts[0]
					n = n + "[" + Join(parts[1:], "")
				}

				newArgs = append(newArgs, fmt.Sprintf("%s %s", t, n))
			}

			// Generate invocation stmt
			invocationStmt := fmt.Sprintf("(%s)", Join(invocationArgs, ","))

			for k, v := range f.Defaults {
				// match anything of form ImVec2(0, 0) or ImVec2(0.0f, 0.0f)
				re := regexp.MustCompile(`(\w*)\((.*)\)`)
				// 3, because according to FindStringSubmatch doc, it returns [ImVec2(x, y), ImVec2, "x, y"]
				// we also need to ensure aur match matches the whole string
				// we also don't want sizeof
				if m := re.FindStringSubmatch(v); len(m) == 3 && m[0] == v && m[1] != "sizeof" {
					for k, v := range ctx.preset.DefaultArgReplace {
						m[2] = ReplaceAll(m[2], k, v)
					}

					v = fmt.Sprintf("(%s){%s}", m[1], m[2])
				}

				if r, ok := ctx.preset.DefaultArgReplace[CIdentifier(v)]; ok {
					v = string(r)
				}

				if r, ok := ctx.preset.DefaultArgArbitraryValue[CIdentifier(k)]; ok {
					v = string(r)
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
				Ret:              ret,
				StName:           f.StName,
				NonUDT:           f.NonUDT,
				CWrapperFuncName: cWrapperFuncName + "V",
				AllCallArgs:      f.AllCallArgs,
			})

			// Add V as suffix to current function name
			funcName += "V"
			cWrapperFuncName += "V"
		}

		actualCallArgsStr := fmt.Sprintf("(%s)", Join(actualCallArgs, ","))
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
				Ret:              ret,
				StName:           f.StName,
				NonUDT:           f.NonUDT,
				CWrapperFuncName: cWrapperFuncName,
				AllCallArgs:      f.AllCallArgs,
				Comment:          f.Comment,
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

		// cppSb.WriteString(fmt.Sprintf("// %#v\n", f))

		// if needed, write extra stuff to cpp headers
		if string(f.AllCallArgs) == actualCallArgsStr && ret == f.Ret {
			cWrapperFuncName = f.FuncName
		} else if f.Constructor {
			headerSb.WriteString(fmt.Sprintf("extern %s* %s%s;\n", f.StName, cWrapperFuncName, f.Args))
			cppSb.WriteString(fmt.Sprintf("%s* %s%s { return %s%s; }\n", f.StName, cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
		} else if f.Destructor {
			headerSb.WriteString(fmt.Sprintf("extern void %s%s;\n", cWrapperFuncName, f.Args))
			cppSb.WriteString(fmt.Sprintf("void %s%s { %s%s; }\n", cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
		} else {
			headerSb.WriteString(fmt.Sprintf("extern %s %s%s;\n", ret, cWrapperFuncName, f.Args))

			switch f.Ret {
			case "void":
				cppSb.WriteString(fmt.Sprintf("%s %s%s { %s%s; }\n", f.Ret, cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
			case "void*":
				cppSb.WriteString(fmt.Sprintf("uintptr_t %s%s { return (uintptr_t)%s%s; }\n", cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
			default:
				cppSb.WriteString(fmt.Sprintf("%s %s%s { return %s%s; }\n", f.Ret, cWrapperFuncName, f.Args, invokeFunctionName, actualCallArgsStr))
			}
		}

		if v, ok := ctx.preset.OriginReplace[funcName]; ok {
			cWrapperFuncName = v
		}

		appendValidFunc()
	}

	headerSb.WriteString(`
#ifdef __cplusplus
}
#endif
`)

	headerPath := "wrapper.h"
	headerFile, err := os.Create(headerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create header file %v: %v", headerPath, err)
	}

	defer headerFile.Close()

	_, err = headerFile.WriteString(headerSb.String())
	if err != nil {
		return nil, fmt.Errorf("failed to write header file %v: %v", headerPath, err)
	}

	cppPath := "wrapper.cpp"
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

func generateCppStructsAccessor(validFuncs []FuncDef, structs []StructDef, context *Context) (accessors []FuncDef, err error) {
	var structAccessorFuncs []FuncDef

	// makes a setsum with context.preset.SkipFUncs
	skipFuncNames := map[CIdentifier]bool{}

	// Add all valid function's name to skipFuncNames
	for _, f := range validFuncs {
		skipFuncNames[f.FuncName] = true
	}

	sbHeader, sbCpp := &strings.Builder{}, &strings.Builder{}

	sbHeader.WriteString(cppFileHeader)
	sbCpp.WriteString(cppFileHeader)

	sbHeader.WriteString(`#pragma once

#include "wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

`)

	sbCpp.WriteString(`
#include <string.h>
#include "wrapper.h"
#include "structs_accessor.h"

`)

	for _, s := range structs {
		for _, m := range s.Members {
			if Contains(m.Type, "(") || Contains(m.Type, "union") {
				continue
			}

			setterFuncName := CIdentifier(fmt.Sprintf("%[1]s_Set%[2]s", s.Name, Capitalize(Split(m.Name, "[")[0])))
			if skipFuncNames[setterFuncName] || context.preset.SkipFuncs[setterFuncName] {
				continue
			}

			memberType := m.Type
			if memberType == "void*" {
				memberType = "uintptr_t"
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
						Name:            "v",
						Type:            memberType + CIdentifier(getSizeArg(m.Size)),
						RemoveFinalizer: true,
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

			sbHeader.WriteString(fmt.Sprintf("extern void %s(%s *%s, %s%s v);\n", setterFuncDef.CWrapperFuncName, s.Name, s.Name+"Ptr", memberType, getPtrIfSize(m.Size)))

			if m.Size > 0 {
				fmt.Fprintf(
					sbCpp,
					"void %s(%s *%s, %s%s v) { memcpy(%s->%s, v, sizeof(%[4]s)*%[8]d); }\n",
					setterFuncDef.CWrapperFuncName,
					s.Name,
					s.Name+"Ptr",
					m.Type,
					getPtrIfSize(m.Size),
					s.Name+"Ptr",
					Split(m.Name, "[")[0],
					m.Size,
				)
			} else {
				if m.Type == "void*" {
					fmt.Fprintf(
						sbCpp,
						"void %s(%s *%s, %s%s v) { %s->%s = (void*)v; }\n",
						setterFuncDef.CWrapperFuncName, s.Name, s.Name+"Ptr", memberType, getPtrIfSize(m.Size), s.Name+"Ptr", Split(m.Name, "[")[0],
					)
				} else {
					fmt.Fprintf(
						sbCpp,
						"void %s(%s *%s, %s%s v) { %s->%s = v; }\n",
						setterFuncDef.CWrapperFuncName, s.Name, s.Name+"Ptr", m.Type, getPtrIfSize(m.Size), s.Name+"Ptr", Split(m.Name, "[")[0],
					)
				}
			}

			getterFuncName := CIdentifier(fmt.Sprintf("%[1]s_Get%[2]s", s.Name, Capitalize(Split(m.Name, "[")[0])))
			if skipFuncNames[getterFuncName] || context.preset.SkipFuncs[getterFuncName] {
				continue
			}

			getterFuncDef := FuncDef{
				Args: fmt.Sprintf("(%[1]s *%[2]s)", s.Name, "self"),
				ArgsT: []ArgDef{
					{
						Name: "self",
						Type: s.Name + "*",
					},
				},
				FuncName:         getterFuncName,
				CWrapperFuncName: "wrap_" + getterFuncName,
				Location:         "",
				Constructor:      false,
				Destructor:       false,
				StructSetter:     false, StructGetter: true,
				Ret: memberType + CIdentifier(getSizeArg(m.Size)),
			}

			structAccessorFuncs = append(structAccessorFuncs, getterFuncDef)

			fmt.Fprintf(
				sbHeader,
				"extern %s%s %s(%s *self);\n",
				memberType, getPtrIfSize(m.Size), getterFuncDef.CWrapperFuncName, s.Name,
			)

			// here we change void* to uintptr_t for .go handling
			if m.Type == "void*" {
				fmt.Fprintf(sbCpp,
					"%s%s %s(%s *self) { return (uintptr_t)self->%s; }\n",
					memberType, getPtrIfSize(m.Size), getterFuncDef.CWrapperFuncName, s.Name, Split(m.Name, "[")[0],
				)
			} else {
				fmt.Fprintf(sbCpp,
					"%s%s %s(%s *self) { return self->%s; }\n",
					memberType, getPtrIfSize(m.Size), getterFuncDef.CWrapperFuncName, s.Name, Split(m.Name, "[")[0],
				)
			}

			// if is array type, need to add a special method to get certain index of array
			if _, ok := context.arrayIndexGetters[m.Type]; m.Size > 0 && !ok {
				AddArrayIndexGetter(m.Type, sbHeader, sbCpp, context)
			}
		}
	}

	sbHeader.WriteString(`
#ifdef __cplusplus
}
#endif
`)

	headerPath := "structs_accessor.h"
	headerFile, err := os.Create(headerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create cpp file %v: %v", headerPath, err)
	}

	defer headerFile.Close()

	_, err = headerFile.WriteString(sbHeader.String())
	if err != nil {
		return nil, fmt.Errorf("failed to write to file %v: %v", headerPath, err)
	}

	cppPath := "structs_accessor.cpp"
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

func getSizeArg(size int) string {
	if size > 0 {
		return fmt.Sprintf("[%d]", size)
	}

	return ""
}

func getPtrIfSize(size int) string {
	if size > 0 {
		return "*"
	}

	return ""
}

func AddArrayIndexGetter(t CIdentifier, sbHeader, sbCpp *strings.Builder, context *Context) {
	tStr := ReplaceAll(ReplaceAll(t, " ", "_"), "*", "Ptr")
	getterFuncName := CIdentifier(context.flags.PackageName) + "_" + tStr + "_GetAtIdx"
	context.arrayIndexGetters[t] = getterFuncName

	fmt.Fprintf(
		sbHeader,
		"extern %[1]s %[2]s(%[1]s *self, int index);\n",
		Split(t, "[")[0], getterFuncName,
	)

	fmt.Fprintf(sbCpp,
		"%[1]s %[2]s(%[1]s *self, int index) { return self[index]; }\n",
		t, getterFuncName,
	)
}
