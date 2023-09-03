package main

import (
	"strings"
)

type (
	// CIdentifier is a string representing name of struct/func/enum in C
	CIdentifier string
	// GoIdentifier is a string representing name of struct/func/enum in Go
	// ATTENTION: GoIdentifier indicates that the name represents a valid Go name.
	// Should be created only by renameGoIdentifier
	GoIdentifier string
)

// Skip functions
// e.g. they are temporarily hard-coded
var skippedFuncs = map[CIdentifier]bool{
	"igInputText":                     true,
	"igInputTextWithHint":             true,
	"igInputTextMultiline":            true,
	"ImFontAtlas_GetTexDataAsAlpha8":  true,
	"ImFontAtlas_GetTexDataAsAlpha8V": true,
	"ImFontAtlas_GetTexDataAsRGBA32":  true,
	"ImFontAtlas_GetTexDataAsRGBA32V": true,
}

// structures that's methods should be skipped
var skippedStructs = map[CIdentifier]bool{
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

var replace = map[CIdentifier]GoIdentifier{
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

func (n CIdentifier) trimImGuiPrefix() CIdentifier {
	if HasPrefix(n, "ImGui") {
		n = TrimPrefix(n, "ImGui")
	} else if HasPrefix(n, "Im") && len(n) > 2 && strings.ToUpper(string(n[2])) == string(n[2]) {
		n = TrimPrefix(n, "Im")
	} else if HasPrefix(n, "ig") && len(n) > 2 && strings.ToUpper(string(n[2])) == string(n[2]) {
		n = TrimPrefix(n, "ig")
	}

	return n
}

func (n CIdentifier) renameGoIdentifier() GoIdentifier {
	if r, ok := replace[n]; ok {
		n = CIdentifier(r)
	}

	n = n.trimImGuiPrefix()
	switch {
	case HasPrefix(n, "New"):
		n = "New" + n[3:].trimImGuiPrefix()
	case HasPrefix(n, "new"):
		n = "new" + n[3:].trimImGuiPrefix()
	case HasPrefix(n, "*"):
		n = "*" + n[1:].trimImGuiPrefix()
	}

	n = TrimPrefix(n, "Get")
	if n != "_" {
		n = ReplaceAll(n, "_", "")
	}
	return GoIdentifier(n)
}

func (e CIdentifier) renameEnum() GoIdentifier {
	return TrimSuffix(e, "_").renameGoIdentifier()
}
