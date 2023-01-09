package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Visitor struct {
	fset *token.FileSet
	file *ast.File
}

var replace = map[string]string{
	"GetDrawData":           "CurrentDrawData",
	"GetDrawListSharedData": "CurrentDrawListSharedData",
	"GetFont":               "CurrentFont",
	"GetIO":                 "CurrentIO",
	"GetPlatformIO":         "CurrentPlatformIO",
	"GetStyle":              "CurrentStyle",
	"GetMouseCursor":        "CurrentMouseCursor",
	"ImAxis":                "AxisName",
	"GetDrawCursor":         "Cursor",
	"SetDrawCursor":         "SetCursor",
}

func removeImGuiIm(n string) string {
	old := n
	if strings.HasPrefix(n, "ImGui") {
		n = strings.TrimPrefix(n, "ImGui")
	} else if strings.HasPrefix(n, "Im") {
		n = strings.TrimPrefix(n, "Im")
	}
	c, _ := utf8.DecodeRuneInString(n)
	if !unicode.IsUpper(c) {
		n = old
	}
	return n
}

func renameIdent(n string) string {
	if r, ok := replace[n]; ok {
		n = r
	}
	n = removeImGuiIm(n)
	if strings.HasPrefix(n, "New") {
		n = "New" + removeImGuiIm(n[3:])
	} else if strings.HasPrefix(n, "new") {
		n = "new" + removeImGuiIm(n[3:])
	}
	n = strings.TrimPrefix(n, "Get")
	if n != "_" {
		n = strings.ReplaceAll(n, "_", "")
	}
	return n
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.SelectorExpr:
		x, ok := n.X.(*ast.Ident)
		if ok && x.Name == "C" {
			return nil
		}
	case *ast.Ident:
		if !n.IsExported() {
			break
		}
		n.Name = renameIdent(n.Name)
	}
	return v
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: rename files")
		return
	}

	for _, fname := range os.Args[1:] {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			fmt.Println(err)
			return
		}

		//ast.Print(fset, file)

		v := &Visitor{fset: fset, file: file}
		for _, decl := range file.Decls {
			ast.Walk(v, decl)
		}

		out, err := os.Create(fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		printer.Fprint(out, fset, file)
		out.Close()
	}
}
