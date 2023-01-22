package main

import (
	"strings"
)

// Skip functions
// e.g. they are temporarily hard-coded
var skippedFuncs = map[string]bool{
	"igInputText":                     true,
	"igInputTextWithHint":             true,
	"igInputTextMultiline":            true,
	"ImFontAtlas_GetTexDataAsAlpha8":  true,
	"ImFontAtlas_GetTexDataAsAlpha8V": true,
	"ImFontAtlas_GetTexDataAsRGBA32":  true,
	"ImFontAtlas_GetTexDataAsRGBA32V": true,
}

// structures that's methods should be skipped
var skippedStructs = map[string]bool{
	"ImVec1":         true,
	"ImVec2":         true,
	"ImVec2ih":       true,
	"ImVec4":         true,
	"ImColor":        true,
	"ImRect":         true,
	"ImPlotTime":     true,
	"StbUndoRecord":  true,
	"StbUndoState":   true,
	"StbTexteditRow": true,
}

var replace = map[string]string{
	"igGetDrawData":           "CurrentDrawData",
	"igGetDrawListSharedData": "CurrentDrawListSharedData",
	"igGetFont":               "CurrentFont",
	"igGetIO":                 "CurrentIO",
	"igGetPlatformIO":         "CurrentPlatformIO",
	"igGetStyle":              "CurrentStyle",
	"igGetMouseCursor":        "CurrentMouseCursor",
	"ImAxis":                  "PlotAxisEnum",
	"GetItem_ID":              "ItemByID",
	//"ImGetDrawCursor":         "Cursor",
	//"ImSetDrawCursor":         "SetCursor",
}

func trimImGuiPrefix(n string) string {
	if strings.HasPrefix(n, "ImGui") {
		n = strings.TrimPrefix(n, "ImGui")
	} else if strings.HasPrefix(n, "Im") && len(n) > 2 && strings.ToUpper(string(n[2])) == string(n[2]) {
		n = strings.TrimPrefix(n, "Im")
	} else if strings.HasPrefix(n, "ig") && len(n) > 2 && strings.ToUpper(string(n[2])) == string(n[2]) {
		n = strings.TrimPrefix(n, "ig")
	}
	return n
}

func renameGoIdentifier(n string) string {
	if r, ok := replace[n]; ok {
		n = r
	}

	n = trimImGuiPrefix(n)
	switch {
	case strings.HasPrefix(n, "New"):
		n = "New" + trimImGuiPrefix(n[3:])
	case strings.HasPrefix(n, "new"):
		n = "new" + trimImGuiPrefix(n[3:])
	case strings.HasPrefix(n, "*"):
		n = "*" + trimImGuiPrefix(n[1:])
	}

	n = strings.TrimPrefix(n, "Get")
	if n != "_" {
		n = strings.ReplaceAll(n, "_", "")
	}
	return n
}
