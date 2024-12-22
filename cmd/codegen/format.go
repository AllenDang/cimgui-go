package main

import (
	"github.com/kpango/glg"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

// FormatGo takes an output GO string and formats it using gofumpt and goimports returning the formatted string.
func FormatGo(s string, ctx *Context) string {
	var err error
	sBytes := []byte(s)
	sBytes, err = format.Source(sBytes, format.Options{
		ModulePath: ctx.preset.PackagePath,
		ExtraRules: true,
	})
	if err != nil {
		if ctx.flags.Verbose {
			glg.Debugf(string(s))
		}

		glg.Fatalf("Unable to format go code: %v", err)
	}

	sBytes, err = imports.Process("", sBytes, &imports.Options{
		Fragment:   false,
		AllErrors:  true,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: false,
	})
	if err != nil {
		if ctx.flags.Verbose {
			glg.Debugf(string(s))
		}

		glg.Fatalf("Unable to goimports code: %v", err)
	}

	return string(sBytes)
}
