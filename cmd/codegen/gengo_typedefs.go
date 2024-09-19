package main

import "C"
import (
	"fmt"
	"os"
	"strings"

	"github.com/kpango/glg"
)

// this function will proceed the following typedefs:
// - all structs thatare not present in struct_and_enums.json (they are supposed to be epaque)
// - everything that satisfies IsCallbackTypedef
func proceedTypedefs(
	typedefs *Typedefs,
	structs []StructDef,
	data *Context,
) (validTypeNames []CIdentifier, err error) {
	// quick counter for coverage control
	generatedTypedefs := 0
	maxTypedefs := len(typedefs.data)

	structsMap := make(map[CIdentifier]StructDef)
	for _, s := range structs {
		structsMap[s.Name] = s
	}

	// we need FILES
	callbacksGoSb := &strings.Builder{}
	callbacksGoSb.WriteString(goPackageHeader)
	fmt.Fprintf(callbacksGoSb,
		`// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "%[1]s_wrapper.h"
// #include "%[1]s_typedefs.h"
import "C"
import "unsafe"

`, data.prefix)

	typedefsHeaderSb := &strings.Builder{}
	typedefsHeaderSb.WriteString(cppFileHeader)
	fmt.Fprintf(typedefsHeaderSb,
		`
#pragma once

#include "cwrappers/%s.h"

#ifdef __cplusplus
extern "C" {
#endif
`, data.prefix)
	typedefsCppSb := &strings.Builder{}
	typedefsCppSb.WriteString(cppFileHeader)
	fmt.Fprintf(typedefsCppSb,
		`
#include "%[1]s_typedefs.h"
#include "cwrappers/%[1]s.h"
`, data.prefix)

	// because go ranges through maps as if it was drunken, we need to sort keys.
	keys := make([]CIdentifier, 0, len(typedefs.data))
	for k := range typedefs.data {
		keys = append(keys, k)
	}

	// sort keys
	SortStrings(keys)

	for _, k := range keys {
		typedef := typedefs.data[k]
		if shouldSkip, ok := skippedTypedefs[k]; ok && shouldSkip {
			if data.flags.showNotGenerated {
				glg.Infof("Arbitrarly skipping typedef %s", k)
			}

			maxTypedefs--
			continue
		}

		if _, exists := data.refTypedefs[k]; exists {
			if data.flags.showNotGenerated {
				glg.Infof("Duplicate of %s in reference typedefs. Skipping.", k)
			}

			maxTypedefs--
			continue
		}

		if shouldSkipStruct(k) {
			if data.flags.showNotGenerated {
				glg.Infof("Arbitrarly skipping struct %s", k)
			}

			maxTypedefs--
			continue
		}

		if IsEnumName(k, data.enumNames) /*|| IsStructName(k, structs)*/ {
			if data.flags.showGenerated {
				glg.Infof("typedef %s has extended deffinition in structs_and_enums.json. Will generate later", k)
			}

			maxTypedefs--
			continue
		}

		if IsTemplateTypedef(typedef) {
			if data.flags.showNotGenerated {
				glg.Failf("typedef %s is a template. not implemented yet", k)
			}

			continue
		}

		isPtr := HasSuffix(typedef, "*")

		var knownReturnType, knownPtrReturnType returnWrapper
		var knownArgType, knownPtrArgType ArgumentWrapperData
		var argTypeErr, ptrArgTypeErr, returnTypeErr, ptrReturnTypeErr error

		if typedef == "void*" {
			typedef = "uintptr_t"
		}

		// Let's say our pureType is of form short
		// the following code needs to handle two things:
		// - int16 -> short (to know go type AND know how to proceed in c() func)
		// - *int16 -> short* (for Handle())
		// - short* -> *int16 (for NewXXXFromC)
		knownReturnType, returnTypeErr = getReturnWrapper(
			CIdentifier(typedef),
			data, // TODO: this might be empty
		)

		knownPtrReturnType, ptrReturnTypeErr = getReturnWrapper(
			CIdentifier(typedef)+"*",
			data, // TODO: this might be empty
		)

		_, knownArgType, argTypeErr = getArgWrapper(
			&ArgDef{
				Name: "self",
				Type: CIdentifier(typedef),
			},
			false, false,
			data, // TODO: this might be empty
		)

		_, knownPtrArgType, ptrArgTypeErr = getArgWrapper(
			&ArgDef{
				Name: "self",
				Type: CIdentifier(typedef) + "*",
			},
			false, false,
			data, // TODO: this might be empty
		)

		// check if k is a name of struct from structDefs
		switch {
		case typedefs.data[k] == "void*":
			if data.flags.showGenerated {
				glg.Successf("typedef %s is an alias to void*.", k)
			}

			fmt.Fprintf(typedefsCppSb,
				`
uintptr_t %[1]s_toUintptr(%[1]s ptr) {
	return (uintptr_t)ptr;
}

%[1]s %[1]s_fromUintptr(uintptr_t ptr) {
	return (%[1]s)ptr;
}
`, k)
			fmt.Fprintf(typedefsHeaderSb, `extern uintptr_t %[1]s_toUintptr(%[1]s ptr);
extern %[1]s %[1]s_fromUintptr(uintptr_t ptr);`, k)

			// NOTE: in case of problems e.g. with Textures, here might be potential issue:
			// Handle() is incomplete - it doesn't have right finalizer (for now I think this will not affect code)
			fmt.Fprintf(callbacksGoSb, `
type %[1]s struct {
	Data uintptr
}

// Handle returns C version of %[1]s and its finalizer func.
func (self *%[1]s) Handle() (result *C.%[6]s, fin func()) {
	r, f := self.C()
    return &r, f
}

// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[6]s, func()) {
    return (C.%[6]s)(C.%[6]s_fromUintptr(C.uintptr_t(self.Data))), func() { }
}

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[6]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	return &%[1]s{Data: (uintptr)(C.%[6]s_toUintptr(*datautils.ConvertCTypes[*C.%[6]s](cvalue)))}
}
`,
				k.renameGoIdentifier(),
				knownArgType.ArgType,

				knownPtrArgType.ArgDef,
				knownPtrArgType.VarName,
				knownPtrArgType.Finalizer,

				k,

				knownArgType.ArgDef,
				knownArgType.VarName,
				knownArgType.Finalizer,

				fmt.Sprintf(knownPtrReturnType.returnStmt, "cvalue"),
			)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case ptrReturnTypeErr == nil && argTypeErr == nil && ptrArgTypeErr == nil && !isPtr:
			if data.flags.showGenerated {
				glg.Successf("typedef %s is an alias typedef.", k)
			}

			fmt.Fprintf(callbacksGoSb, `
type %[1]s %[2]s

// Handle returns C version of %[1]s and its finalizer func.
func (selfSrc *%[1]s) Handle() (result *C.%[6]s, fin func()) {
	self := (*%[2]s)(selfSrc)
    %[3]s
    return (*C.%[6]s)(%[4]s), func() { %[5]s }
}

// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[6]s, func()) {
    %[7]s
    return (C.%[6]s)(%[8]s), func() { %[9]s }
}

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[6]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	return (*%[1]s)(%[10]s)
}
`,
				k.renameGoIdentifier(),
				knownArgType.ArgType,

				knownPtrArgType.ArgDef,
				knownPtrArgType.VarName,
				knownPtrArgType.Finalizer,

				k,

				knownArgType.ArgDef,
				knownArgType.VarName,
				knownArgType.Finalizer,

				fmt.Sprintf(knownPtrReturnType.returnStmt, fmt.Sprintf("datautils.ConvertCTypes[*C.%s](cvalue)", k)),
			)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case returnTypeErr == nil && argTypeErr == nil && isPtr:
			// if it's a pointer type, I think we can proceed as above, but without Handle() method...
			// (handle proceeds pointer values and we don't want double pointers, really)
			fmt.Fprintf(callbacksGoSb, `
type %[1]s  struct {
	Data %[2]s
}

// Handle returns C version of %[1]s and its finalizer func.
func (self *%[1]s) Handle() (*C.%[6]s, func()) {
	result, fn := self.C()
	return &result, fn
}

// C is like Handle but returns plain type instead of pointer.
func (selfStruct *%[1]s) C() (result C.%[6]s, fin func()) {
	self := selfStruct.Data
    %[3]s
    return (C.%[6]s)(%[4]s), func() { %[5]s }
}

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[6]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	v := (%[8]s)(*datautils.ConvertCTypes[*C.%[6]s](cvalue))
	return &%[1]s{Data: %[7]s}
}
`,
				k.renameGoIdentifier(),
				knownArgType.ArgType,

				knownArgType.ArgDef,
				knownArgType.VarName,
				knownArgType.Finalizer,

				k,

				fmt.Sprintf(knownReturnType.returnStmt, "v"),
				knownArgType.CType,
			)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case IsCallbackTypedef(typedefs.data[k]):
			if data.flags.showNotGenerated {
				glg.Failf("typedef %s is a callback. Not implemented yet", k)
			}
		case HasPrefix(typedefs.data[k], "struct"):
			isOpaque := !IsStructName(k, structsMap)
			if data.flags.showGenerated {
				glg.Successf("typedef %s is a struct (is opaque? %v).", k, isOpaque)
			}

			writeOpaqueStruct(k, isOpaque, callbacksGoSb)
			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		default:
			if data.flags.showNotGenerated {
				glg.Failf("unknown situation happened for type %s; not implemented. Probably unknown Arg (err: %v), Ret (err; %v) PtrArg (err: %v) or PtrRet (err: %v) type wrappers for isPointer: %v for %s. Check out source code for more details", k, argTypeErr, returnTypeErr, ptrArgTypeErr, ptrReturnTypeErr, isPtr, typedefs.data[k])
			}
		}
	}

	fmt.Fprint(typedefsHeaderSb,
		`
#ifdef __cplusplus
}
#endif`)

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.go", data.prefix), []byte(callbacksGoSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.go: %w", data.prefix, err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.cpp", data.prefix), []byte(typedefsCppSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.cpp: %w", data.prefix, err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.h", data.prefix), []byte(typedefsHeaderSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.h: %w", data.prefix, err)
	}

	glg.Infof("Typedefs generation complete. Generated %d/%d (%.2f%%) typedefs.", generatedTypedefs, maxTypedefs, float32(generatedTypedefs*100)/float32(maxTypedefs))
	return validTypeNames, nil
}

func writeOpaqueStruct(name CIdentifier, isOpaque bool, sb *strings.Builder) {
	// this will be put only for structs that are NOT opaque (w can know the exact definition)
	var toPlainValue string
	if !isOpaque {
		toPlainValue = fmt.Sprintf(`
// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[2]s, func()) {
	result, fn := self.Handle()
	return *result, fn
}
`, name.renameGoIdentifier(), name)
	}

	// we need to make it a struct, because we need to hide C type (otherwise it will duplicate methods)
	fmt.Fprintf(sb, `
type %[1]s struct {
	CData *C.%[2]s
}

// Handle returns C version of %[1]s and its finalizer func.
func (self *%[1]s) Handle() (result *C.%[2]s, fin func()) {
	return self.CData, func() {}
}

%[3]s

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[2]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	return &%[1]s{CData: datautils.ConvertCTypes[*C.%[2]s](cvalue)}
}
`, name.renameGoIdentifier(), name, toPlainValue)
}

func IsStructName[T any](name CIdentifier, structs map[CIdentifier]T) bool {
	_, ok := structs[name]
	return ok
}

func IsEnumName(name CIdentifier, enums map[GoIdentifier]bool) bool {
	_, ok := enums[name.renameEnum()]
	return ok
}

func IsTemplateTypedef(s string) bool {
	return strings.Contains(s, "<")
}
