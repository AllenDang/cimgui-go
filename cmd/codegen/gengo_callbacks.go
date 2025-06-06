package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/kpango/glg"
)

type callbackType int

const (
	callbackType1 callbackType = iota
	callbackType2
)

var callbackNotGeneratedError = errors.New("callback was not generated")

type callbacksGenerator struct {
	goSb, cgoSb *strings.Builder
	ctx         *Context
}

func GenerateCallbacks(callbacks []CIdentifier, context *Context) (validTypes []CIdentifier, err error) {
	result := &callbacksGenerator{
		goSb:  &strings.Builder{},
		cgoSb: &strings.Builder{},
		ctx:   context,
	}

	generatedCallbacks := 0
	callbacksToGenerate := len(callbacks)

	validTypes = make([]CIdentifier, 0)

	// sort just in case
	sort.Slice(callbacks, func(i, j int) bool {
		return callbacks[i] < callbacks[j]
	})

	result.writeHeaer()

	for _, typedefName := range callbacks {
		if err := result.writeCallback(typedefName, result.ctx.typedefs.data[typedefName]); err != nil {
			if errors.Is(err, callbackNotGeneratedError) {
				if context.flags.ShowNotGenerated {
					glg.Warnf("Callback \"%s\" was not generated: %v", typedefName, err)
				}

				continue
			}

			return nil, err
		}

		if context.flags.ShowGenerated {
			glg.Successf("Callback \"%s\" was generated", typedefName)
		}

		validTypes = append(validTypes, typedefName)
		generatedCallbacks++
	}

	if err := result.saveToDisk(); err != nil {
		return nil, fmt.Errorf("cannot save to disk: %w", err)
	}

	glg.Infof("Callbacks generation complete. Generated %d/%d callbacks (%.2f%%)", generatedCallbacks, callbacksToGenerate, float32(generatedCallbacks)/float32(callbacksToGenerate)*100)
	return validTypes, nil
}

// This is an underlying function for gengo_typedefs.go for now.
// May be reworked to be a separated "genertor" in the future.
// - includes logging
// - returns CallbackNotGeneratedError
func (g *callbacksGenerator) writeCallback(typedefName CIdentifier, def string) error {
	// see https://github.com/AllenDang/cimgui-go/issues/224#issuecomment-2452156237
	// 1: preprocessing - parse typedefs.data[k] to get return type and arguments

	// now let me use a bit of regex.
	// We have 2 possibilities:
	// - returnType(*<funcName>)(args1 arg1Name, args2 arg2Name, args3 arg3Name);
	// - returnType <funcName>(args1 arg1Name, args2 arg2Name, args3 arg3Name);
	// NOTE: the second is uesed mainly in immarkdown
	// NOTE: in the 1st, spaces does not matter so we'll trim them
	expr1, err := regexp.Compile("(typedef )?([a-zA-Z0-9_]+\\*?) *\\(\\*.*\\)\\((.*)\\);")
	if err != nil {
		panic(fmt.Sprintf("Cannot compile regex expr1!: %v", err))
	}

	expr2, err := regexp.Compile("([a-zA-Z0-9_]+)\\s+([a-zA-Z0-9_]+)\\((.*)\\);")
	if err != nil {
		panic(fmt.Sprintf("Cannot compile regex expr2!: %v", err))
	}

	// we need the following from them:
	var returnTypeC CIdentifier
	argsC := make([]ArgDef, 0)

	var cbType callbackType

	if ok := expr1.MatchString(def); ok {
		cbType = callbackType1
		if g.ctx.flags.Verbose {
			glg.Debugf("callback typedef \"%s\" is in form 1", typedefName)
		}

		// now split by "("
		// it should be something like this:
		// ["returnType", "*<optional func name)", "args1 arg1Name, args2 arg2Name, args3 arg3Name);"]
		def = TrimPrefix(def, "typedef ")
		parts := Split(def, "(")
		if len(parts) != 3 {
			panic("Cannot split by (, check implementation in cmd/codegen!")
		}

		returnTypeC = TrimSuffix(CIdentifier(parts[0]), " ")
		argsStr := parts[2]
		argsStr = TrimSuffix(argsStr, ");")
		argsStr = ReplaceAll(argsStr, ", ", ",")
		argsStr = ReplaceAll(argsStr, "&", "")
		argsListStr := Split(argsStr, ",")
		for a, argStr := range argsListStr {
			// get name
			argParts := Split(argStr, " ")
			var name, typeName string
			switch len(argParts) {
			case 1:
				name = fmt.Sprintf("arg%d", a)
				typeName = strings.Join(argParts, " ")
			case 2: // like "int arg1" or "const int"
				if argParts[0] == "const" {
					name = fmt.Sprintf("arg%d", a)
					typeName = strings.Join(argParts, " ")
					break
				}

				fallthrough
			case 3: // something like "int" or "const int arg1"
				name = argParts[len(argParts)-1]
				typeName = strings.Join(argParts[:len(argParts)-1], " ")
			}

			argsC = append(argsC, ArgDef{
				Name: CIdentifier(name),
				Type: CIdentifier(typeName),
			})
		}
	} else if ok := expr2.MatchString(def); ok {
		cbType = callbackType2
		parts := Split(def, "(")
		if len(parts) != 2 {
			panic("Cannot split by (, check implementation in cmd/codegen!")
		}

		returnTypeC = TrimSuffix(CIdentifier(parts[0]), " ")
		returnTypeC = TrimSuffix(CIdentifier(parts[0]), fmt.Sprintf(" %s", typedefName))
		argsStr := parts[1]
		argsStr = TrimSuffix(argsStr, ");")
		argsStr = ReplaceAll(argsStr, ", ", ",")
		argsStr = ReplaceAll(argsStr, "&", "*")
		argsListStr := Split(argsStr, ",")
		for a, argStr := range argsListStr {
			// get name
			argParts := Split(argStr, " ")
			var name, typeName string
			switch len(argParts) {
			case 1:
				name = fmt.Sprintf("arg%d", a)
				typeName = strings.Join(argParts, " ")
			case 2: // like "int arg1" or "const int"
				if argParts[0] == "const" {
					name = fmt.Sprintf("arg%d", a)
					typeName = strings.Join(argParts, " ")
					break
				}

				fallthrough
			case 3: // something like "int" or "const int arg1"
				name = argParts[len(argParts)-1]
				typeName = strings.Join(argParts[:len(argParts)-1], " ")
			}

			argsC = append(argsC, ArgDef{
				Name: CIdentifier(name),
				Type: CIdentifier(typeName),
			})
		}
	} else {
		if g.ctx.flags.ShowNotGenerated {
			return fmt.Errorf("cannot parse callback typedef: %w", callbackNotGeneratedError)
		}
	}

	// 2: Find wrappers:
	// We need to figure out how to wrap returnType and args.
	// In fact, we need to swap meaning of them, because we want to convert C argument type to Go argument type
	// so we are supposed to use returnWrapper for that.
	if g.ctx.flags.Verbose {
		glg.Debugf("From %s got \"%s\" and \"%v\"", typedefName, returnTypeC, argsC)
	}

	// 2.1: get return wrapper
	var returnType ArgumentWrapperData
	if returnTypeC == "void" {
		returnType = ArgumentWrapperData{}
	} else {
		_, returnType, err = getArgWrapper(
			&ArgDef{
				Name: "result",
				Type: returnTypeC,
			},
			false, false,
			g.ctx,
		)
		if err != nil {
			return fmt.Errorf("unable to get return wrapper %w: %w", callbackNotGeneratedError, err)
		}
	}

	// 2.1: get arg wrappers
	args := make([]returnWrapper, len(argsC))
	for i, arg := range argsC {
		rw, err := getReturnWrapper(arg.Type, g.ctx)
		if err != nil {
			return fmt.Errorf("unable to get arg wrapper for \"%s\" (%w): %w", arg.Type, callbackNotGeneratedError, err)
		}

		// fill rw return stmt
		rw.returnStmt = fmt.Sprintf(rw.returnStmt, arg.Name)

		args[i] = rw
	}

	// 3: Prepare call statement
	cCallStmt := ""
	goCallStmt := ""
	valuePassStmt := ""
	for i, arg := range args {
		cCallStmt += fmt.Sprintf("%s %s, ", argsC[i].Name, arg.CType)
		goCallStmt += fmt.Sprintf("%s %s, ", argsC[i].Name, arg.returnType)
		valuePassStmt += fmt.Sprintf("%s, ", argsC[i].Name)
	}

	cCallStmt = TrimSuffix(cCallStmt, ", ")
	goCallStmt = TrimSuffix(goCallStmt, ", ")
	valuePassStmt = TrimSuffix(valuePassStmt, ", ")

	// 4: Write code
	switch cbType {
	case callbackType1:
		fmt.Fprintf(g.goSb, `
type %[1]s func(%[4]s) %[2]s
type c%[1]s func(%[5]s) %[3]s

func New%[1]sFromC(cvalue *C.%[6]s) *%[1]s {
	result := pool%[1]s.Find(*cvalue)
	return &result
}

func (c %[1]s) C() (C.%[6]s, func()) {
	return pool%[1]s.Allocate(c), func() { }
}
`,
			typedefName.renameGoIdentifier(g.ctx),
			returnType.ArgType,
			returnType.CType,
			goCallStmt,
			cCallStmt,
			typedefName,
		)
	case callbackType2:
		fmt.Fprintf(g.goSb, `
type %[1]s func(%[4]s) %[2]s
type c%[1]s func(%[5]s) %[3]s

func New%[1]sFromC(cvalue *C.%[6]s) *%[1]s {
	result := pool%[1]s.Find(cvalue)
	return &result
}

func (c %[1]s) Handle() (*C.%[6]s, func()) {
	result := pool%[1]s.Allocate(c)
	return result, func() {}
}
`,
			typedefName.renameGoIdentifier(g.ctx),
			returnType.ArgType,
			returnType.CType,
			goCallStmt,
			cCallStmt,
			typedefName,
		)
	}

	cCallStmt2 := cCallStmt
	if cCallStmt2 != "" {
		cCallStmt2 = ", " + cCallStmt2
	}

	if returnType.ArgType == "" {
		fmt.Fprintf(g.goSb, `
func wrap%[1]s(cb %[1]s %[3]s) %[2]s {
	cb(%[4]s)
}
`,
			typedefName.renameGoIdentifier(g.ctx),
			returnType.CType,
			cCallStmt2,
			func() string {
				result := ""
				for _, a := range args {
					result += a.returnStmt + ", "
				}
				result = TrimSuffix(result, ", ")
				return result
			}(),
		)
	} else {
		fmt.Fprintf(g.goSb, `
func wrap%[1]s(cb %[1]s %[5]s) %[2]s {
	result := cb(%[6]s)
	%[3]s
	defer func() {
		%[7]s
	}()
	return %[4]s
}
`,
			typedefName.renameGoIdentifier(g.ctx),
			returnType.CType,
			returnType.ArgDef,
			returnType.VarName,
			cCallStmt2,
			func() string {
				result := ""
				for _, a := range args {
					result += a.returnStmt + ", "
				}
				result = TrimSuffix(result, ", ")
				return result
			}(),
			returnType.Finalizer,
		)
	}

	size := g.ctx.preset.TypedefsPoolSize
	if h, ok := g.ctx.preset.TypedefsCustomPoolSizes[typedefName]; ok {
		size = h
	}

	poolNames := make([]string, size)
	// now write N functions
	for i := 0; i < size; i++ {
		switch cbType {
		case callbackType1:
			fmt.Fprintf(g.goSb,
				`//export callback%[1]s%[2]d
func callback%[1]s%[2]d(%[5]s) %[3]s { %[4]s wrap%[1]s(pool%[1]s.Get(%[2]d), %[6]s) }
					`,
				typedefName.renameGoIdentifier(g.ctx),
				i,
				returnType.CType,
				func() string {
					if returnType.ArgType != "" {
						return "return"
					}

					return ""
				}(),
				cCallStmt,
				valuePassStmt,
			)
		case callbackType2:

			fmt.Fprintf(g.goSb,
				`//export callback%[1]s%[2]d
func callback%[1]s%[2]d(%[5]s) %[3]s {%[4]s wrap%[1]s(pool%[1]s.Get(%[2]d), %[6]s) }
					`,
				typedefName.renameGoIdentifier(g.ctx),
				i,
				returnType.CType,
				func() string {
					if returnType.ArgType != "" {
						return "return"
					}

					return ""
				}(),
				cCallStmt,
				valuePassStmt,
			)
		}

		fmt.Fprintf(g.cgoSb,
			`// extern %[1]s callback%[2]s%[3]d(%[4]s);
`,
			returnTypeC,
			typedefName.renameGoIdentifier(g.ctx),
			i,
			func() string {
				result := ""
				for _, a := range argsC {
					result += TrimPrefix(string(a.Type), "const ") + ", "
				}

				return TrimSuffix(result, ", ")
				return result
			}(),
		)

		switch cbType {
		case callbackType1:
			poolNames[i] = fmt.Sprintf("C.%[3]s(C.callback%[1]s%[2]d)", typedefName.renameGoIdentifier(g.ctx), i, typedefName)
		case callbackType2:
			poolNames[i] = fmt.Sprintf("(*C.%[3]s)(C.callback%[1]s%[2]d)", typedefName.renameGoIdentifier(g.ctx), i, typedefName)
		}
	}

	switch cbType {
	case callbackType1:
		fmt.Fprintf(g.goSb, `
var pool%[1]s *internal.Pool[%[1]s, C.%[3]s]
func init() {
	pool%[1]s = internal.NewPool[%[1]s, C.%[3]s](
%[2]s
)
}

func Clear%[1]sPool() {
	pool%[1]s.Clear()
}
`,
			typedefName.renameGoIdentifier(g.ctx),
			strings.Join(poolNames, ",\n")+",\n",
			typedefName,
		)
	case callbackType2:
		fmt.Fprintf(g.goSb, `
var pool%[1]s *internal.Pool[%[1]s, *C.%[3]s]
func init() {
	pool%[1]s = internal.NewPool[%[1]s, *C.%[3]s](
%[2]s
)
}

func Clear%[1]sPool() {
	pool%[1]s.Clear()
}
`,
			typedefName.renameGoIdentifier(g.ctx),
			strings.Join(poolNames, ",\n")+",\n",
			typedefName,
		)
	}

	return nil
}

func (g *callbacksGenerator) writeHeaer() {
}

func (g *callbacksGenerator) saveToDisk() error {
	result := &strings.Builder{}
	result.WriteString(getGoPackageHeader(g.ctx))
	fmt.Fprintf(result,
		`// #include <stdlib.h>
// #include <memory.h>
// #include "wrapper.h"
// #include "typedefs.h"
%s
`, g.ctx.preset.MergeCGoPreamble())

	fmt.Fprintf(result, g.cgoSb.String())

	fmt.Fprintf(result, `import "C"
import "unsafe"
`)

	fmt.Fprintf(result, g.goSb.String())

	if err := os.WriteFile("callbacks.go", []byte(FormatGo(result.String(), g.ctx)), 0644); err != nil {
		return fmt.Errorf("cannot write typedefs.go: %w", err)
	}

	return nil
}
