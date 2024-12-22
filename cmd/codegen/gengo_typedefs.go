package main

import "C"
import (
	"fmt"
	"os"
	"strings"

	"github.com/kpango/glg"
)

type typedefsGenerator struct {
	GoSb  *strings.Builder
	HSb   *strings.Builder
	CppSb *strings.Builder
	ctx   *Context
}

// GenerateTypedefs will proceed all typedefs from typedefs_dict.json
func GenerateTypedefs(
	typedefs *Typedefs,
	structs []StructDef,
	ctx *Context,
) (validTypeNames, callbacks []CIdentifier, err error) {
	// quick counter for coverage control
	generatedTypedefs := 0
	maxTypedefs := len(typedefs.data)

	generator := &typedefsGenerator{
		// we need FILES
		GoSb:  &strings.Builder{},
		HSb:   &strings.Builder{},
		CppSb: &strings.Builder{},
		ctx:   ctx,
	}

	generator.writeHeaders()

	// because go ranges through maps as if it was drunken, we need to sort keys.
	keys := make([]CIdentifier, 0, len(typedefs.data))
	for k := range typedefs.data {
		keys = append(keys, k)
	}

	// sort keys
	SortStrings(keys)

	for _, k := range keys {
		typedef := typedefs.data[k]
		if shouldSkip, ok := ctx.preset.SkipTypedefs[k]; ok && shouldSkip {
			if ctx.flags.showNotGenerated {
				glg.Infof("Arbitrarly skipping typedef %s", k)
			}

			maxTypedefs--
			continue
		}

		if _, exists := ctx.refTypedefs[k]; exists {
			if ctx.flags.showNotGenerated {
				glg.Infof("Duplicate of %s in reference typedefs. Skipping.", k)
			}

			maxTypedefs--
			continue
		}

		if shouldSkipStruct := ctx.preset.SkipStructs[k]; shouldSkipStruct {
			if ctx.flags.showNotGenerated {
				glg.Infof("Arbitrarly skipping struct %s", k)
			}

			maxTypedefs--
			continue
		}

		_, isEnum := ctx.enumNames[k]
		_, isRefEnum := ctx.refEnumNames[k]
		if isEnum || isRefEnum {
			if ctx.flags.showGenerated {
				glg.Infof("typedef %s has extended deffinition in structs_and_enums.json. Will generate later", k)
			}

			maxTypedefs--
			continue
		}

		if IsTemplateTypedef(typedef) {
			if ctx.flags.showNotGenerated {
				glg.Warnf("typedef %s is a template. not implemented yet", k)
			}

			continue
		}

		isPtr := HasSuffix(typedef, "*")

		known := generator.parseArgDef(CIdentifier(typedef), ctx)

		switch {
		case typedefs.data[k] == "void*":
			if ctx.flags.showGenerated {
				glg.Successf("typedef %s is an alias to void*.", k)
			}

			generator.writeVoidPtrTypedef(k, known)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case !isPtr &&
			known.ptrReturnTypeErr == nil &&
			known.argTypeErr == nil &&
			known.ptrArgTypeErr == nil:
			if ctx.flags.showGenerated {
				glg.Successf("typedef %s is an alias typedef.", k)
			}

			generator.writeAliasTypedef(k, known)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case isPtr &&
			known.returnTypeErr == nil &&
			known.argTypeErr == nil:

			generator.writePtrTypedef(k, known)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case IsCallbackTypedef(typedefs.data[k]):
			maxTypedefs--
			callbacks = append(callbacks, k)
		case HasPrefix(typedefs.data[k], "struct"):
			isOpaque := !IsStructName(k, ctx)
			if ctx.flags.showGenerated {
				glg.Successf("typedef %s is a struct (is opaque? %v).", k, isOpaque)
			}

			generator.writeStruct(k, isOpaque)
			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		default:
			if ctx.flags.showNotGenerated {
				glg.Errorf("unknown situation happened for type %s; not implemented. Probably unknown Arg (err: %v), Ret (err; %v) PtrArg (err: %v) or PtrRet (err: %v) type wrappers for isPointer: %v for %s. Check out source code for more details",
					k, known.argTypeErr, known.returnTypeErr, known.ptrArgTypeErr, known.ptrReturnTypeErr, isPtr, typedefs.data[k])
			}
		}
	}

	if err := generator.saveToDisk(); err != nil {
		return nil, nil, fmt.Errorf("cannot save typedefs to disk: %w", err)
	}

	glg.Infof("Typedefs generation complete. Generated %d/%d (%.2f%%) typedefs.", generatedTypedefs, maxTypedefs, float32(generatedTypedefs*100)/float32(maxTypedefs))
	return validTypeNames, callbacks, nil
}

// Let's say our pureType is of form "short"
// the following code needs to handle two things:
// - int16 -> short (to know go type AND know how to proceed in C() func)
// - *int16 -> short* (for Handle())
// - short* -> *int16 (for NewXXXFromC)
type typedefTypeContext struct {
	argType,
	ptrArgType ArgumentWrapperData
	returnType,
	ptrReturnType returnWrapper

	argTypeErr, ptrArgTypeErr,
	returnTypeErr, ptrReturnTypeErr error
}

func (g *typedefsGenerator) parseArgDef(typedef CIdentifier, ctx *Context) *typedefTypeContext {
	result := &typedefTypeContext{}
	result.returnType, result.returnTypeErr = getReturnWrapper(
		CIdentifier(typedef),
		ctx,
	)

	result.ptrReturnType, result.ptrReturnTypeErr = getReturnWrapper(
		CIdentifier(typedef)+"*",
		ctx,
	)

	_, result.argType, result.argTypeErr = getArgWrapper(
		&ArgDef{
			Name: "self",
			Type: CIdentifier(typedef),
		},
		false, false,
		ctx,
	)

	_, result.ptrArgType, result.ptrArgTypeErr = getArgWrapper(
		&ArgDef{
			Name: "self",
			Type: CIdentifier(typedef) + "*",
		},
		false, false,
		ctx,
	)

	return result
}

func (g *typedefsGenerator) writeHeaders() {
	g.HSb.WriteString(cppFileHeader)
	fmt.Fprintf(g.HSb,
		`
#pragma once

#include "%s"

#ifdef __cplusplus
extern "C" {
#endif
`, g.ctx.flags.include)
	g.CppSb.WriteString(cppFileHeader)
	fmt.Fprintf(g.CppSb,
		`
#include "typedefs.h"
#include "%[1]s"
`, g.ctx.flags.include)

	g.GoSb.WriteString(getGoPackageHeader(g.ctx))
	fmt.Fprintf(g.GoSb,
		`// #include <stdlib.h>
// #include <memory.h>
// #include "wrapper.h"
// #include "typedefs.h"
%s
import "C"
import "unsafe"
`, g.ctx.preset.MergeCGoPreamble())
}

// k is plain C name of the typedef (key in typedefs_dict.json)
// known is parsed value of k
func (g *typedefsGenerator) writeVoidPtrTypedef(k CIdentifier, known *typedefTypeContext) {
	fmt.Fprintf(g.CppSb,
		`
uintptr_t %[1]s_toUintptr(%[1]s ptr) {
	return (uintptr_t)ptr;
}

%[1]s %[1]s_fromUintptr(uintptr_t ptr) {
	return (%[1]s)ptr;
}
`, k)
	fmt.Fprintf(g.HSb, `extern uintptr_t %[1]s_toUintptr(%[1]s ptr);
extern %[1]s %[1]s_fromUintptr(uintptr_t ptr);`, k)

	// NOTE: in case of problems e.g. with Textures, here might be potential issue:
	// Handle() is incomplete - it doesn't have right finalizer (for now I think this will not affect code)
	fmt.Fprintf(g.GoSb, `
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
	return &%[1]s{Data: (uintptr)(C.%[6]s_toUintptr(*internal.ReinterpretCast[*C.%[6]s](cvalue)))}
}
`,
		k.renameGoIdentifier(g.ctx),
		known.argType.ArgType,

		known.ptrArgType.ArgDef,
		known.ptrArgType.VarName,
		known.ptrArgType.Finalizer,

		k,

		known.argType.ArgDef,
		known.argType.VarName,
		known.argType.Finalizer,

		fmt.Sprintf(known.ptrReturnType.returnStmt, "cvalue"),
	)
}

// k is plain C name of the typedef (key in typedefs_dict.json)
// known is parsed value of k
func (g *typedefsGenerator) writeAliasTypedef(k CIdentifier, known *typedefTypeContext) {
	fmt.Fprintf(g.GoSb, `
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
		k.renameGoIdentifier(g.ctx),
		known.argType.ArgType,

		known.ptrArgType.ArgDef,
		known.ptrArgType.VarName,
		known.ptrArgType.Finalizer,

		k,

		known.argType.ArgDef,
		known.argType.VarName,
		known.argType.Finalizer,

		fmt.Sprintf(known.ptrReturnType.returnStmt, fmt.Sprintf("internal.ReinterpretCast[*C.%s](cvalue)", k)),
	)
}

func (g *typedefsGenerator) writePtrTypedef(k CIdentifier, known *typedefTypeContext) {
	fmt.Fprintf(g.GoSb, `
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
	v := (%[8]s)(*internal.ReinterpretCast[*C.%[6]s](cvalue))
	return &%[1]s{Data: %[7]s}
}
`,
		k.renameGoIdentifier(g.ctx),
		known.argType.ArgType,

		known.argType.ArgDef,
		known.argType.VarName,
		known.argType.Finalizer,

		k,

		fmt.Sprintf(known.returnType.returnStmt, "v"),
		known.argType.CType,
	)
}

func (g *typedefsGenerator) writeStruct(
	name CIdentifier,
	isOpaque bool, // if this is true, we assume that we cannot use the exact value of a struct (only pointer)
) {
	// this will be put only for structs that are NOT opaque (w can know the exact definition)
	var toPlainValue string
	if !isOpaque {
		toPlainValue = fmt.Sprintf(`
// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[2]s, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmpty%[1]s creates %[1]s with its 0 value.
func NewEmpty%[1]s() *%[1]s {
	return &%[1]s{CData: new(C.%[2]s)}
}
`, name.renameGoIdentifier(g.ctx), name)
	}

	// we need to make it a struct, because we need to hide C type (otherwise it will duplicate methods)
	fmt.Fprintf(g.GoSb, `
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
	return &%[1]s{CData: internal.ReinterpretCast[*C.%[2]s](cvalue)}
}
`, name.renameGoIdentifier(g.ctx), name, toPlainValue)
}

func (g *typedefsGenerator) saveToDisk() error {
	fmt.Fprint(g.HSb,
		`
#ifdef __cplusplus
}
#endif`)

	if err := os.WriteFile("typedefs.go", []byte(FormatGo(g.GoSb.String(), g.ctx)), 0644); err != nil {
		return fmt.Errorf("cannot write typedefs.go: %w", err)
	}

	if err := os.WriteFile("typedefs.cpp", []byte(g.CppSb.String()), 0644); err != nil {
		return fmt.Errorf("cannot write typedefs.cpp: %w", err)
	}

	if err := os.WriteFile("typedefs.h", []byte(g.HSb.String()), 0644); err != nil {
		return fmt.Errorf("cannot write typedefs.h: %w", err)
	}

	return nil
}
