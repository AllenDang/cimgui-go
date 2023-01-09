package main

import (
	"fmt"
	"os"
	"strings"
)

func generateGoStructs(prefix string, structs []StructDef) []string {
	var sb strings.Builder

	sb.WriteString(goPackageHeader)
	sb.WriteString(fmt.Sprintf(
		`// #include "%s_wrapper.h"
import "C"
import "unsafe"

`, prefix))

	// Save all struct name into a map
	var structNames []string

	for _, s := range structs {
		if shouldSkipStruct(s.Name) {
			continue
		}

		goName := trimImGuiPrefix(s.Name)

		sb.WriteString(fmt.Sprintf(`type %[2]s uintptr

func (data %[2]s) handle() *C.%[1]s {
  return (*C.%[1]s)(unsafe.Pointer(data))
}

func (data %[2]s) c() C.%[1]s {
  return *(data.handle())
}

func new%[2]sFromC(cvalue C.%[1]s) %[2]s {
  return %[2]s(unsafe.Pointer(&cvalue))
}

`, s.Name, goName))

		structNames = append(structNames, goName)
	}

	structFile, err := os.Create(fmt.Sprintf("%s_structs.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer structFile.Close()

	_, _ = structFile.WriteString(sb.String())

	return structNames
}
