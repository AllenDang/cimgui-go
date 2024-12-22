package main

import (
	"github.com/kpango/glg"
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
		glg.Fatalf("Unable to format go code: %v", err)
	}

	return string(sBytes)
}
