package main

import (
	"regexp"
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
	"ImVec1":     true,
	"ImVec2":     true,
	"ImVec2ih":   true,
	"ImVec4":     true,
	"ImColor":    true,
	"ImRect":     true,
	"ImPlotTime": true,
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
}

func (c CIdentifier) trimImGuiPrefix() CIdentifier {
	if HasPrefix(c, "ImGui") {
		c = TrimPrefix(c, "ImGui")
	} else if HasPrefix(c, "Im") && len(c) > 2 && strings.ToUpper(string(c[2])) == string(c[2]) {
		c = TrimPrefix(c, "Im")
	} else if HasPrefix(c, "ig") && len(c) > 2 && strings.ToUpper(string(c[2])) == string(c[2]) {
		c = TrimPrefix(c, "ig")
	}

	return c
}

func (c CIdentifier) renameGoIdentifier() GoIdentifier {
	if r, ok := replace[c]; ok {
		c = CIdentifier(r)
	}

	c = c.trimImGuiPrefix()
	switch {
	case HasPrefix(c, "New"):
		c = "New" + c[3:].trimImGuiPrefix()
	case HasPrefix(c, "new"):
		c = "new" + c[3:].trimImGuiPrefix()
	case HasPrefix(c, "*"):
		c = "*" + c[1:].trimImGuiPrefix()
	}

	c = TrimPrefix(c, "Get")
	if c != "_" {
		c = ReplaceAll(c, "_", "")
	}
	return GoIdentifier(c)
}

func (c CIdentifier) renameEnum() GoIdentifier {
	return TrimSuffix(c, "_").renameGoIdentifier()
}

// returns true if s is of form TypeName<*> <(>Name<*><)>(args)
// (fragments in <> are optional)
func IsCallbackTypedef(s string) bool {
	pattern := `\w*\**\(*\w*\**\)*\(.*\);`
	b, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}

	return b
}
