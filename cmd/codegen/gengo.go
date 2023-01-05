package main

import (
	"fmt"
	"strings"
)

// Skip functions
// e.g. they are temporarily hard-coded
func skippedFuncs() []string {
	return []string{
		"InputText",
		"InputTextWithHint",
		"InputTextMultiline",
		"FontAtlas_GetTexDataAsAlpha8",
		"FontAtlas_GetTexDataAsAlpha8V",
		"FontAtlas_GetTexDataAsRGBA32",
		"FontAtlas_GetTexDataAsRGBA32V",
	}
}

func trimImGuiPrefix(id string) string {
	// don't trim prefixes for implot's ImAxis - it conflicts with ImGuIAxis (from imgui_internal.h)
	if strings.HasPrefix(id, "ImAxis") {
		return id
	}

	id = strings.TrimPrefix(id, "ImGui")
	id = strings.TrimPrefix(id, "Im")
	id = strings.TrimPrefix(id, "ig")
	return id
}

type argOutput struct {
	ArgType string
	ArgDef  string
	VarName string
}

// Generate function args
func argStmtFunc(argWrappers []argOutput, sb *strings.Builder) string {
	var invokeStmt []string
	for _, aw := range argWrappers {
		invokeStmt = append(invokeStmt, aw.VarName)
		if len(aw.ArgDef) > 0 {
			sb.WriteString(fmt.Sprintf("%s\n\n", aw.ArgDef))
		}
	}

	return strings.Join(invokeStmt, ",")
}
