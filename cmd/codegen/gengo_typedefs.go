package main

import "C"
import (
	"fmt"
	"github.com/kpango/glg"
	"os"
	"strings"
)

// this function will proceed the following typedefs:
// - all structs thatare not present in struct_and_enums.json (they are supposed to be epaque)
// - everything that satisfies IsCallbackTypedef
func proceedTypedefs(prefix string, typedefs *Typedefs, structs []StructDef, enums []EnumDef) (validTypeNames []CIdentifier, err error) {
	// we need FILES
	callbacksGoSb := &strings.Builder{}
	callbacksGoSb.WriteString(goPackageHeader)
	fmt.Fprintf(callbacksGoSb,
		`// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "%s_wrapper.h"
import "C"
import "unsafe"

`, prefix)
	//callbacksCSb := &strings.Builder{}

	// because go ranges through maps as if it was drunken, we need to sort keys.
	keys := make([]CIdentifier, 0, len(typedefs.data))
	for k := range typedefs.data {
		keys = append(keys, k)
	}

	// sort keys
	SortStrings(keys)

	for _, k := range keys {
		if IsEnumName(k, enums) || IsStructName(k, structs) {
			glg.Infof("typedef %s has extended deffinition in structs_and_enums.json. Will generate later", k)
			continue
		}

		if IsTemplateTypedef(typedefs.data[k]) {
			glg.Infof("typedef %s is a template. not implemented yet", k)
			continue
		}

		knownReturnType, returnTypeErr := getReturnWrapper(
			CIdentifier(typedefs.data[k]),
			map[CIdentifier]bool{},
			map[GoIdentifier]bool{},
		)

		_, knownArgType, argTypeErr := getArgWrapper(
			&ArgDef{
				Name: "self",
				Type: CIdentifier(typedefs.data[k]),
			},
			false, false,
			map[CIdentifier]bool{},
			map[GoIdentifier]bool{},
		)

		// check if k is a name of struct from structDefs
		switch {
		case returnTypeErr == nil && argTypeErr == nil:
			glg.Infof("typedef %s is an alias typedef.", k)
			fmt.Fprintf(callbacksGoSb, `
type %[1]s %[2]s

func (self %[1]s) handle() (result *%[2]s, fin func()) {
    %[3]s
    return %[4]s, func() { %[5]s }
}

func (self %[1]s) c() (%[2]s, func()) {
	result, fin := self.handle()
	return *result, fin
}

func new%[1]sFromC(cvalue %[6]s) {
	return %[7]s
}
`, k.renameGoIdentifier(), knownArgType.ArgType, knownArgType.ArgDef, knownArgType.VarName, knownArgType.Finalizer, knownArgType.CType, fmt.Sprintf(knownReturnType.returnStmt, "cvalue"))

			validTypeNames = append(validTypeNames, k)
		case IsCallbackTypedef(typedefs.data[k]):
			glg.Infof("typedef %s is a callback. Not implemented yet", k)
		case HasPrefix(typedefs.data[k], "struct"):
			glg.Infof("typedef %s is an opaque struct.", k)
			writeOpaqueStruct(k, callbacksGoSb)
			validTypeNames = append(validTypeNames, k)
		}
	}

	fmt.Println(callbacksGoSb.String())

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.go", prefix), []byte(callbacksGoSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.go: %w", prefix, err)
	}

	return validTypeNames, nil
}

func writeOpaqueStruct(name CIdentifier, sb *strings.Builder) {
	fmt.Fprintf(sb, `
type %[1]s C.%[2]s

func (self %[1]s) handle() (result *C.%[2]s, fin func()) {
	result = (*C.%[2]s)(unsafe.Pointer(&self))
	return result, func() {}
}

func (self %[1]s) c() (C.%[2]s, func()) {
	result, fin := self.handle()
	return *result, fin
}

func new%[2]sFromC(cvalue *C.%[2]s) *%[1]s {
	return (*%[1]s)(cvalue)
}
`, name.renameGoIdentifier(), name)
}

func IsStructName(name CIdentifier, structs []StructDef) bool {
	for _, s := range structs {
		if s.Name == name {
			return true
		}
	}

	return false
}

func IsEnumName(name CIdentifier, enums []EnumDef) bool {
	for _, e := range enums {
		if e.Name.renameEnum() == name.renameEnum() { // compare GO equivalents because C names has _ at their end
			return true
		}
	}

	return false
}

func IsTemplateTypedef(s string) bool {
	return strings.Contains(s, "<")
}
