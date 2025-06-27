package main

import (
	"regexp"
	"strings"
)

type (
	// EnumIdentifier is in theory EnumName_
	EnumIdentifier string
	// CIdentifier is a string representing name of struct/func/enum in C
	CIdentifier string
	// GoIdentifier is a string representing name of struct/func/enum in Go
	// ATTENTION: GoIdentifier indicates that the name represents a valid Go name.
	// Should be created only by renameGoIdentifier
	GoIdentifier string
)

func (c CIdentifier) trimImGuiPrefix(ctx *Context) CIdentifier {
	for _, prefix := range ctx.preset.TrimPrefix {
		if HasPrefix(c, prefix) &&
			len(c) > len(prefix) && // check if removing it will not result in an empty string
			strings.ToUpper(string(c[len(prefix)])) == string(c[len(prefix)]) { // check if the method will still be exported
			c = TrimPrefix(c, prefix)
		}
	}

	return c
}

func (c CIdentifier) renameGoIdentifier(ctx *Context) GoIdentifier {
	if r, ok := ctx.preset.Replace[c]; ok {
		c = CIdentifier(r)
		return GoIdentifier(c)
	}

	c = TrimSuffix(c, "_Nil")

	c = c.trimImGuiPrefix(ctx)
	switch {
	case HasPrefix(c, "New"):
		c = "New" + c[3:].trimImGuiPrefix(ctx)
	case HasPrefix(c, "new"):
		c = "new" + c[3:].trimImGuiPrefix(ctx)
	case HasPrefix(c, "*"):
		c = "*" + c[1:].trimImGuiPrefix(ctx)
	}

	c = TrimPrefix(c, "Get")
	if c != "_" {
		c = ReplaceAll(c, "_", "")
	}

	return GoIdentifier(c)
}

func (c EnumIdentifier) renameEnum() CIdentifier {
	return CIdentifier(TrimSuffix(c, "_"))
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

func IsStructName(name CIdentifier, ctx *Context) bool {
	_, ok := ctx.typedefsNames[name]
	return ok
}

func IsTemplateTypedef(s string) bool {
	return strings.Contains(s, "<")
}
