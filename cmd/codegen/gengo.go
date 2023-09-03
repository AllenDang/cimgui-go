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

func trimImGuiPrefix(n CIdentifier) CIdentifier {
	tmpName := string(n)
	if strings.HasPrefix(tmpName, "ImGui") {
		tmpName = strings.TrimPrefix(tmpName, "ImGui")
	} else if strings.HasPrefix(tmpName, "Im") && len(n) > 2 && strings.ToUpper(string(n[2])) == string(n[2]) {
		tmpName = strings.TrimPrefix(tmpName, "Im")
	} else if strings.HasPrefix(tmpName, "ig") && len(n) > 2 && strings.ToUpper(string(n[2])) == string(n[2]) {
		tmpName = strings.TrimPrefix(tmpName, "ig")
	}

	return CIdentifier(tmpName)
}

func (n CIdentifier) renameGoIdentifier() GoIdentifier {
	if r, ok := replace[n]; ok {
		n = CIdentifier(r)
	}

	n = trimImGuiPrefix(n)
	switch {
	case strings.HasPrefix(string(n), "New"):
		n = "New" + trimImGuiPrefix(n[3:])
	case strings.HasPrefix(string(n), "new"):
		n = "new" + trimImGuiPrefix(n[3:])
	case strings.HasPrefix(string(n), "*"):
		n = "*" + trimImGuiPrefix(n[1:])
	}

	n = CIdentifier(strings.TrimPrefix(string(n), "Get"))
	if n != "_" {
		n = CIdentifier(strings.ReplaceAll(string(n), "_", ""))
	}
	return GoIdentifier(n)
}

func (e CIdentifier) renameEnum() GoIdentifier {
	e = CIdentifier(strings.TrimSuffix(string(e), "_"))
	return e.renameGoIdentifier()
}
