package main

import (
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

// structures that's methods should be skipped
func skippedStructs() []string {
	return []string{
		"ImVec1",
		"ImVec2",
		"ImVec2ih",
		"ImVec4",
		"ImColor",
		"ImRect",
		"StbUndoRecord",
		"StbUndoState",
		"StbTexteditRow",
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
	// go-valid argument type (e.g. string, ImVec2, etc.)
	ArgType string
	// argument deffinition (e.g. arg1, arg1Fin := ...\ndefer arg1Fin())
	ArgDef string
	// name of argument (e.g. arg1)
	VarName string
}
