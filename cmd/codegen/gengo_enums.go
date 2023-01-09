package main

import (
	"fmt"
	"os"
	"strings"
)

// Generate enums and return enum type names
func generateGoEnums(prefix string, enums []EnumDef) []string {
	var sb strings.Builder

	sb.WriteString(goPackageHeader)

	var enumNames []string
	for _, e := range enums {
		originalName := e.Name
		eName := strings.TrimSuffix(e.Name, "_")
		eName = trimImGuiPrefix(eName)

		enumNames = append(enumNames, eName)

		sb.WriteString(fmt.Sprintf("// original name: %s\n", originalName))
		sb.WriteString(fmt.Sprintf("type %s int\n", eName))
		sb.WriteString("const (\n")

		for _, v := range e.Values {
			vName := trimImGuiPrefix(v.Name)
			vName = strings.TrimSuffix(vName, "_")
			sb.WriteString(fmt.Sprintf("\t%s = %d\n", vName, v.Value))
		}

		sb.WriteString(")\n\n")
	}

	enumFile, err := os.Create(fmt.Sprintf("%s_enums.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer enumFile.Close()

	_, _ = enumFile.WriteString(sb.String())

	return enumNames
}
