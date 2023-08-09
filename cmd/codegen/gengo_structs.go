package main

import "C"
import (
	"fmt"
	"github.com/kpango/glg"
	"os"
	"strings"
)

func generateGoStructs(prefix string, structs []StructDef, enums []EnumDef) []string {
	glg.Infof("Generating %d structs", len(structs))
	var progress int

	var sb strings.Builder

	sb.WriteString(goPackageHeader)
	sb.WriteString(fmt.Sprintf(
		`// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "%s_wrapper.h"
import "C"
import "unsafe"

`, prefix))

	// Save all struct name into a map
	var structNames []string

	for _, s := range structs {
		if shouldSkipStruct(s.Name) {
			progress++
			continue
		}

		sb.WriteString(fmt.Sprintf("%s\n", s.CommentAbove))
		if generateStruct(s, structs, enums, &sb) {
			progress++
		}

		structNames = append(structNames, s.Name)
	}

	structFile, err := os.Create(fmt.Sprintf("%s_structs.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer structFile.Close()

	_, _ = structFile.WriteString(sb.String())

	glg.Infof("%d of %d structs fully generated (%.2f %%)",
		progress,
		len(structs),
		float32(progress)/float32(len(structs))*100,
	)

	return structNames
}

func generateStruct(s StructDef, defs []StructDef, enums []EnumDef, sb *strings.Builder) (generationComplete bool) {
	generationComplete = true

	type wrapper struct {
		fromC returnWrapper
		toC   ArgumentWrapperData
	}

	var wrappers = make(map[int]wrapper)
	var isTODO bool // this will become true if some field is not implemented (e.g. union {...})
	var structBody *strings.Builder = &strings.Builder{}

	fmt.Fprintf(sb, "type %s struct {\n", renameGoIdentifier(s.Name))

	// Generate struct fields:
	for i, field := range s.Members {
		argDef := ArgDef{
			Name: renameStructField(field.Name),
			Type: field.Type,
		}

		typeName := ""

		// first of all check if the type isn't another stuct
		var isOtherStruct bool
		for _, otherS := range defs {
			if otherS.Name == field.Type && !shouldSkipStruct(field.Type) {
				isOtherStruct = true
				break
			}
		}

		// and same for enums
		var isEnum bool
		for _, enum := range enums {
			if enum.Name == field.Type && !shouldSkipStruct(field.Type) {
				isEnum = true
				break
			}
		}

		toC, toCErr := argWrapper(field.Type)
		fromC, fromCErr := getReturnTypeWrapperFunc(field.Type)
		switch {
		case toCErr == nil && fromCErr == nil:
			wrappers[i] = wrapper{
				toC:   toC(argDef),
				fromC: fromC,
			}

			typeName = wrappers[i].toC.ArgType
		case isOtherStruct:
			wrappers[i] = wrapper{
				fromC: returnWrapper{
					returnType: field.Type,
					returnStmt: fmt.Sprintf("new%sFromC(&%%s)", renameGoIdentifier(field.Type)),
				},
				toC: ArgumentWrapperData{
					ArgType:   field.Type,
					ArgDef:    fmt.Sprintf("%[1]sArg, %[1]sFin := %[2]s.c()", field.Name, renameStructField(field.Name)),
					Finalizer: fmt.Sprintf("%sFin()", field.Name),
					VarName:   fmt.Sprintf("%sArg", field.Name),
				},
			}

			typeName = renameGoIdentifier(field.Type)
		case isEnum:
			wrappers[i] = wrapper{
				fromC: returnWrapper{
					returnType: field.Type,
					returnStmt: fmt.Sprintf("%s(%%s)", renameGoIdentifier(field.Type)),
				},
				toC: ArgumentWrapperData{
					ArgType: field.Type,
					VarName: fmt.Sprintf("C.%s(%s)", field.Type, renameStructField(field.Name)),
				},
			}

			typeName = renameGoIdentifier(field.Type)
		default:
			isTODO = true
		}

		if field.Name == "" || // <- this means that type is union or something like that.
			strings.ContainsAny(field.Name, "[]") || // <- this means that it is an array; TODO
			field.Bitfield != "" ||
			isTODO {
			generationComplete = false
			isTODO = true
			/*
					This may happen due to the following things:
						- member has no name (the only case that I found is when field is "union {...}"
				        see https://github.com/cimgui/cimgui/issues/241
						- the field is an array - for some reason the [] is treated as part of name and that case
						is simply not yet implemented (TODO)
						- field is a bitfield and go does not (and is not going to) support it
						see https://github.com/golang/go/issues/59982
						- or finally (I believe the most common reason) type of this field is not implemented
						in argument_wrapper.go and return_wrapper.go. This should be present in both maps
						in order to make this generator work for that type.
			*/
			glg.Failf("Cannot generate struct \"%s\", because its member \"%s\" (bitfield %v) is of unsupported type \"%s\" ",
				s.Name, field.Name, field.Bitfield, field.Type,
			)
			// reset struct body and fill it with temporary data
			structBody = &strings.Builder{}
			fmt.Fprint(structBody, `
// TODO: contains unsupported fields
data unsafe.Pointer
`)
			break
		}

		if field.Comment.Above != "" {
			field.Comment.Above += "\n"
		}

		fmt.Fprintf(structBody, "%s%s %s %s\n",
			field.Comment.Above,
			renameStructField(field.Name), typeName,
			field.Comment.Sameline,
		)
	}

	fmt.Fprintln(sb, structBody)

	fmt.Fprint(sb, "}\n")

	// handlers:
	fmt.Fprintf(sb, `
func (data %[1]s) handle() (result *C.%[2]s, releaseFn func()) {
`, renameGoIdentifier(s.Name), s.Name)

	if isTODO {
		fmt.Fprintf(sb, `
result = (*C.%s)(data.data)
return result, func() {}
`, s.Name)
	} else {
		fmt.Fprintf(sb, "result = new(C.%s)\n", s.Name)
		for i, m := range s.Members {
			fmt.Fprintf(sb,
				"%[4]s := data.%[4]s\n%[1]s\nresult.%[2]s = %[3]s\n",
				wrappers[i].toC.ArgDef,
				m.Name, wrappers[i].toC.VarName,
				renameStructField(m.Name),
			)
		}

		// finalizer (release func)
		fmt.Fprintf(sb, "releaseFn = func() {\n")

		for i := range s.Members {
			fmt.Fprintf(sb,
				"%s\n",
				wrappers[i].toC.Finalizer)
		}

		fmt.Fprintf(sb, "}\nreturn result, releaseFn\n")
	}

	fmt.Fprintf(sb, "}\n")

	fmt.Fprintf(sb, `
	func (data %[1]s) c() (result C.%[2]s, fin func()) {
		resultPtr, finFn := data.handle()
		return *resultPtr, finFn
	}
`, renameGoIdentifier(s.Name), s.Name)

	// new X FromC
	fmt.Fprintf(sb, "func new%[1]sFromC(cvalue *C.%[2]s) %[1]s {\n", renameGoIdentifier(s.Name), s.Name)

	fmt.Fprintf(sb, "result := new(%s)\n", renameGoIdentifier(s.Name))
	if isTODO {
		fmt.Fprintf(sb, "result.data = unsafe.Pointer(cvalue)\n")
	} else {
		for i, m := range s.Members {
			w := wrappers[i].fromC
			stmt := strings.TrimPrefix(w.returnStmt, "return")
			stmt = fmt.Sprintf(stmt, fmt.Sprintf("cvalue.%s", m.Name))
			fmt.Fprintf(sb, "result.%s = %s\n", renameStructField(m.Name), stmt)
		}
	}

	fmt.Fprintf(sb, "return *result\n}\n")

	return
}

func renameStructField(original string) (result string) {
	result = strings.TrimPrefix(original, "_")
	result = "Field" + result
	return result
}
