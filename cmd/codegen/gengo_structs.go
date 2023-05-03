package main

import "C"
import (
	"fmt"
	"os"
	"strings"
)

func generateGoStructs(prefix string, structs []StructDef) []string {
	var sb strings.Builder

	sb.WriteString(goPackageHeader)
	sb.WriteString(fmt.Sprintf(
		`// #include <stdlib.h>
// #include <memory.h>
// #include "%s_wrapper.h"
import "C"
import "unsafe"

`, prefix))

	// Save all struct name into a map
	var structNames []string

	for _, s := range structs {
		if shouldSkipStruct(s.Name) {
			continue
		}

		sb.WriteString(fmt.Sprintf("%s\n", s.CommentAbove))
		generateStruct(s, &sb)

		structNames = append(structNames, s.Name)
	}

	structFile, err := os.Create(fmt.Sprintf("%s_structs.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer structFile.Close()

	_, _ = structFile.WriteString(sb.String())

	return structNames
}

func generateStruct(s StructDef, sb *strings.Builder) {
	type wrapper struct {
		fromC returnWrapper
		toC   argumentWrapper
	}

	var wrappers = make(map[int]wrapper)
	var isTODO bool // this will become true if some field is not implemented (e.g. union {...})
	var structBody *strings.Builder = &strings.Builder{}

	fmt.Fprintf(sb, "type %s struct {\n", renameGoIdentifier(s.Name))
	// Generate struct fields:
	for i, field := range s.Members {
		if field.Name == "" || // <- this means that type is union or something like that.
			strings.ContainsAny(field.Name, "[]") { // <- this means that it is an array; TODO
			isTODO = true
			// reset struct body and fill it with temporary data
			structBody = &strings.Builder{}
			fmt.Fprint(structBody, `
// TODO: contains unsupported fields
data uintptr
`)
			continue
		}

		typeName := ""
		switch field.Type {
		default:
			wrappers[i] = wrapper{
				fromC: simpleR("uintptr"),
				toC:   simpleW("uintptr", field.Name),
			}

			typeName = "uintptr"
		}

		fmt.Fprintf(structBody, "%s %s\n", renameStructField(field.Name), typeName)
	}

	fmt.Fprintln(sb, structBody)

	fmt.Fprint(sb, "}\n")

	// handlers:
	fmt.Fprintf(sb, `
func (data %[1]s) handle() *C.%[2]s {
`, renameGoIdentifier(s.Name), s.Name)

	if isTODO {
		fmt.Fprintf(sb, "return (*C.%s)(unsafe.Pointer(data.data))", s.Name)
	} else {
		fmt.Fprintf(sb, "result := new(C.%s)\n", s.Name)
		for i, m := range s.Members {
			wr := wrappers[i].toC(ArgDef{
				Name: fmt.Sprintf("data.%s", m.Name),
				Type: "uintptr", // TODO
			})

			fmt.Fprintf(sb, "%s\nresult.%s = %s\n", wr.ArgDef, renameStructField(m.Name), wr.VarName)
		}

		fmt.Fprintf(sb, "return result\n")
	}

	fmt.Fprintf(sb, "}\n")

	fmt.Fprintf(sb, `
	func (data %[1]s) c() C.%[2]s {
		return *(data.handle())
	}
`, renameGoIdentifier(s.Name), s.Name)

	// release func:
	fmt.Fprintf(sb, `
func (data %[1]s) release(result *C.%[2]s) {
`, renameGoIdentifier(s.Name), s.Name)

	if isTODO {
		fmt.Fprintf(sb, "C.free(unsafe.Pointer(result.data))")
	} else {
		for i, m := range s.Members {
			wr := wrappers[i].toC(ArgDef{
				Name: fmt.Sprintf("data.%s", m.Name),
				Type: "uintptr", // TODO
			})

			fmt.Fprintf(sb, "%s\n", wr.Finalizer)
		}
	}

	fmt.Fprintf(sb, "}\n")

	// new X FromC
	fmt.Fprintf(sb, "func new%[1]sFromC(cvalue C.%[2]s) %[1]s {\n", renameGoIdentifier(s.Name), s.Name)

	if isTODO {
		fmt.Fprintf(sb, "return %[1]s(unsafe.Pointer(&cvalue.data))", s.Name)
	} else {
		fmt.Fprintf(sb, "result := new(%s)\n", renameGoIdentifier(s.Name))
		for i, m := range s.Members {
			w := wrappers[i].fromC
			stmt := strings.TrimPrefix(w.returnStmt, "return")
			stmt = fmt.Sprintf(stmt, fmt.Sprintf("cvalue.%s", s.Name))
			fmt.Fprintf(sb, "result.%s = %s\n", renameStructField(m.Name), stmt)
		}

		fmt.Fprintf(sb, "return result\n")
	}

	//sb.WriteString(fmt.Sprintf(`
	//return %[1]s(unsafe.Pointer(&cvalue))

	//`, renameGoIdentifier(s.Name), s.Name))

	fmt.Fprintf(sb, "}\n")
}

func renameStructField(original string) (result string) {
	result = strings.TrimPrefix(original, "_")
	result = "Field" + result
	return result
}
