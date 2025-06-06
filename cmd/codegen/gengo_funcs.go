package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/kpango/glg"
)

// returnTypeType represents an arbitrary type of return value of the function.
// for example Known reffers to returnTypeWrappersMap (see below)
type returnTypeType uint16

const (
	// default value - will cause the function to be skipped and an error will be printed to stdout
	returnTypeUnknown returnTypeType = 1 << iota
	// return type is void (in go - the function returns nothing)
	returnTypeVoid
	// METHOD returns nothing, but it has receiver (called self)
	returnTypeStructSetter
	// Known - reffers to getReturnTypeWrapperFunc
	returnTypeKnown
	// return type is a pointer to ImGui struct
	returnTypeStructPtr
	// returns ImGui struct
	returnTypeStruct
	// the method is a constructor
	returnTypeConstructor
	// function with first arugment as pointer of return value
	returnTypeNonUDT
	// will be treated as struct fieldd getter
	// TODO: This is convirmed to work only with returnTypeKnown.
	returnTypeCustomFin
)

// Is returns true if r contains t.
func (r *returnTypeType) Is(t returnTypeType) bool {
	return *r&t != 0
}

// GenerateGoFuncs generates given list of functions and writes them to file
func GenerateGoFuncs(
	validFuncs []FuncDef,
	context *Context,
) error {
	funcsToConvert := len(validFuncs)

	generator := &goFuncsGenerator{
		typedefsNames: context.typedefsNames,
		enumNames:     context.enumNames,
		refTypedefs:   context.refTypedefs,
		context:       context,
		sb:            &strings.Builder{},
	}

	generator.writeFuncsFileHeader()

	for _, f := range validFuncs {
		// check whether the function shouldn't be skipped
		if context.ShouldSkipFunc(f.FuncName) {
			funcsToConvert--
			if context.flags.ShowNotGenerated {
				glg.Debugf("Func skipped: %s%s", f.FuncName, f.Args)
			}

			continue
		}

		if context.ShouldSkipFile(f.Location) {
			if context.flags.ShowNotGenerated {
				glg.Warnf("File %s skipped: %s%s", f.Location, f.FuncName, f.Args)
			}

			funcsToConvert--

			continue
		}

		args, argWrappers := generator.generateFuncArgs(f)

		if len(f.ArgsT) == 0 {
			generator.shouldGenerate = true
		}

		// stop, when the function should not be generated
		if !generator.shouldGenerate {
			if context.flags.ShowNotGenerated {
				glg.Warnf("not generated: %s%s", f.FuncName, f.Args)
			}

			continue
		} else {
			if context.flags.ShowGenerated {
				glg.Successf("generated: %s%s", f.FuncName, f.Args)
			}
		}

		if noErrors := generator.GenerateFunction(f, args, argWrappers); !noErrors {
			continue
		}
	}

	glg.Infof("GO Functions generation complete. Generated %d/%d (%.2f%%)",
		generator.convertedFuncCount,
		funcsToConvert,
		100*float32(generator.convertedFuncCount)/float32(funcsToConvert),
	)

	filename := "funcs.go"
	glg.Infof("Running Formatting on %s", filename)
	filecontent := FormatGo(generator.sb.String(), context)

	if err := os.WriteFile(filename, []byte(filecontent), 0o666); err != nil {
		return fmt.Errorf("failed to write content of GO file %s: %w", filename, err)
	}

	return nil
}

// goFuncsGenerator is an internal state of GO funcs' generator
type goFuncsGenerator struct {
	typedefsNames map[CIdentifier]bool
	enumNames     map[CIdentifier]bool
	refTypedefs   map[CIdentifier]bool

	sb                 *strings.Builder
	convertedFuncCount int
	shouldGenerate     bool

	context *Context
}

// writeFuncsFileHeader writes a header of the generated file
func (g *goFuncsGenerator) writeFuncsFileHeader() {
	g.sb.WriteString(getGoPackageHeader(g.context))

	fmt.Fprintf(g.sb,
		`// #include "structs_accessor.h"
// #include "wrapper.h"
%s
import "C"
import "unsafe"

`, g.context.preset.MergeCGoPreamble())
}

func (g *goFuncsGenerator) GenerateFunction(f FuncDef, args []GoIdentifier, argWrappers []ArgumentWrapperData) (noErrors bool) {
	var cReturnType CIdentifier
	var returnType, receiver GoIdentifier
	var cfuncCall string
	funcName := f.FuncName
	shouldDefer := false

	// determine kind of function:
	returnTypeType := returnTypeUnknown
	_, err := getReturnWrapper(f.Ret, g.context) // TODO: we call this twice now
	if err == nil {
		returnTypeType = returnTypeKnown
	}

	// attention! order is _probably_ important here so consider that
	// before changing anything here
	if f.NonUDT == 1 {
		returnTypeType = returnTypeNonUDT
	} else if f.Ret == "void" {
		if f.StructSetter {
			returnTypeType = returnTypeStructSetter
		} else {
			returnTypeType = returnTypeVoid
		}
	} else if HasSuffix(f.Ret, "*") && (g.typedefsNames[TrimSuffix(f.Ret, "*")] || g.typedefsNames[TrimSuffix(TrimPrefix(f.Ret, "const "), "*")]) {
		returnTypeType = returnTypeStructPtr
	} else if f.StructGetter && g.typedefsNames[f.Ret] {
		returnTypeType = returnTypeStruct
	} else if f.Constructor {
		returnTypeType = returnTypeConstructor
	}

	if _, ok := g.context.preset.CustomFInalizer[f.FuncName]; ok {
		returnTypeType |= returnTypeCustomFin
	}

	// determine function name, return type (and return statement)
	switch {
	case returnTypeType.Is(returnTypeVoid):
		// noop
	case returnTypeType.Is(returnTypeNonUDT):
		outArg := argWrappers[0]
		returnType = TrimPrefix(outArg.ArgType, "*")

		cfuncCall = fmt.Sprintf("*%s", f.ArgsT[0].Name)

		argWrappers[0].ArgDef = fmt.Sprintf(`%s := new(%s)
%s
		`, f.ArgsT[0].Name, returnType, outArg.ArgDef)
		args = args[1:]
	case returnTypeType.Is(returnTypeStructSetter):
		funcParts := Split(f.FuncName, "_")
		funcName = TrimPrefix(f.FuncName, string(funcParts[0]+"_"))
		if len(funcName) == 0 || !HasPrefix(funcName, "Set") || g.context.ShouldSkipMethod(funcParts[0]) {
			return false
		}

		receiver = funcParts[0].renameGoIdentifier(g.context)
	case returnTypeType.Is(returnTypeKnown):
		shouldDefer = true
		cReturnType = f.Ret
		returnType = f.Ret.renameGoIdentifier(g.context)
	case returnTypeType.Is(returnTypeStructPtr):
		// return Im struct ptr
		shouldDefer = true
		cReturnType = TrimPrefix(f.Ret, "const ")
		returnType = cReturnType.renameGoIdentifier(g.context)
	case returnTypeType.Is(returnTypeStruct):
		shouldDefer = true
		cReturnType = f.Ret
		returnType = cReturnType.renameGoIdentifier(g.context)
	case returnTypeType.Is(returnTypeConstructor):
		shouldDefer = true
		parts := Split(f.FuncName, "_")
		cReturnType = parts[0] + "*"
		returnType = cReturnType.renameGoIdentifier(g.context)

		suffix := ""
		if len(parts) > 2 {
			suffix = string(Join(parts[2:], ""))
		}

		funcName = CIdentifier("New" + string(parts[0]) + suffix)
	default:
		if g.context.flags.Verbose {
			glg.Debugf("Unknown return type \"%s\" in function %s", f.Ret, f.FuncName)
		}

		return false
	}

	rw, err := getReturnWrapper(cReturnType, g.context)
	if err != nil {
		if returnTypeType.Is(returnTypeKnown | returnTypeStructPtr | returnTypeConstructor | returnTypeStruct) {
			return false
		}
	}

	if rw.returnType != "" {
		returnType = rw.returnType
	}

	g.sb.WriteString(g.generateFuncDeclarationStmt(receiver, funcName, args, returnType, f))
	argInvokeStmt, declarations, finishers := g.generateFuncBody(argWrappers)
	g.sb.WriteString(strings.Join(declarations, "\n"))
	if len(declarations) > 0 {
		g.sb.WriteString("\n")
	}

	// assume there is a variable called "result" which should be freed by calling its CustomFinalizer
	if returnTypeType.Is(returnTypeCustomFin) {
		finishers = append(finishers, fmt.Sprintf("C.%s(result)", g.context.preset.CustomFInalizer[f.FuncName]))
	}

	// return Var switch
	switch {
	case returnTypeType.Is(returnTypeStruct | returnTypeCustomFin):
		g.sb.WriteString(fmt.Sprintf(`
result := C.%s(%s)
`,
			f.CWrapperFuncName,
			argInvokeStmt,
		))
		cfuncCall = "result"
	case returnTypeType.Is(returnTypeKnown | returnTypeConstructor | returnTypeStructPtr):
		cfuncCall = fmt.Sprintf("C.%s(%s)", f.CWrapperFuncName, argInvokeStmt)
	}

	if shouldDefer {
		g.writeFinishers(shouldDefer, finishers)
	}

	// write non-return function calls (finalizers called normally)
	switch {
	case returnTypeType.Is(returnTypeVoid | returnTypeNonUDT):
		g.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.CWrapperFuncName, argInvokeStmt))
	case returnTypeType.Is(returnTypeStructSetter):
		g.sb.WriteString(fmt.Sprintf(`
selfArg, selfFin := self.Handle()
defer selfFin()
C.%s(selfArg, %s)
`, f.CWrapperFuncName, argInvokeStmt))
	}

	if !shouldDefer {
		g.writeFinishers(shouldDefer, finishers)
	}

	switch {
	case returnTypeType.Is(returnTypeNonUDT):
		g.sb.WriteString(fmt.Sprintf("return %s", cfuncCall))
	case returnTypeType.Is(returnTypeKnown | returnTypeStructPtr | returnTypeConstructor | returnTypeStruct):
		g.sb.WriteString("return " + fmt.Sprintf(rw.returnStmt, cfuncCall))
	}

	g.sb.WriteString("}\n\n")
	g.convertedFuncCount += 1

	return true
}

// this method is responsible for createing a function declaration statement.
// it takes function name, list of arguments and return type and returns go statement.
// e.g.: func (self *ImGuiType) FuncName(arg1 type1, arg2 type2) returnType {
func (g *goFuncsGenerator) generateFuncDeclarationStmt(receiver GoIdentifier, funcName CIdentifier, args []GoIdentifier, returnType GoIdentifier, f FuncDef) (functionDeclaration string) {
	funcParts := Split(funcName, "_")
	typeName := funcParts[0]

	// convert func(self *receiverType) into a method
	if len(funcParts) > 1 &&
		len(args) > 0 &&
		Contains(args[0], "self ") {

		funcName = TrimPrefix(funcName, string(typeName+"_"))
		receiver = TrimPrefix(args[0], "self ")
		args = args[1:]
	}

	if len(receiver) > 0 {
		receiver = GoIdentifier(fmt.Sprintf("(self %s)", receiver))
	}

	goFuncName := funcName.renameGoIdentifier(g.context)
	// make sure goFuncName is exported
	goFuncName = GoIdentifier(unicode.ToUpper(rune(goFuncName[0]))) + goFuncName[1:]

	// if file comes from imgui_internal.h,prefix Internal is added.
	// ref: https://github.com/AllenDang/cimgui-go/pull/118
	for _, internalFile := range g.context.preset.InternalFiles {
		if strings.HasPrefix(f.Location, internalFile) {
			goFuncName = GoIdentifier(g.context.preset.InternalPrefix) + goFuncName
			break
		}
	}

	// Generate default param value hint
	commentSb := &strings.Builder{}
	comments := strings.Split(f.Comment, "\n")
	for i, comment := range comments {
		if !strings.HasPrefix(comment, "//") {
			comments[i] = "// " + comments[i]
		}
	}

	commentSb.WriteString(fmt.Sprintf("%s\n", strings.Join(comments, "\n")))
	if len(f.Defaults) > 0 {
		fmt.Fprintf(commentSb, "// %s parameter default value hint:\n", goFuncName)

		type defaultParam struct {
			name  GoIdentifier
			value string
		}
		defaults := make([]defaultParam, 0, len(f.Defaults))
		// sort according to the order of the arguments
		for _, arg := range args {
			if idx := Index(arg, " "); idx != -1 {
				arg = arg[:idx]
			}
			d, ok := f.Defaults[string(arg)]
			if !ok {
				continue
			}

			defaults = append(defaults, defaultParam{name: arg, value: d})
		}

		for _, p := range defaults {
			commentSb.WriteString(fmt.Sprintf("// %s: %s\n", p.name, p.value))
		}
	}

	return fmt.Sprintf("%sfunc %s %s(%s) %s {\n",
		commentSb.String(),
		receiver,
		goFuncName,
		Join(args, ","),
		returnType)
}

func (g *goFuncsGenerator) generateFuncArgs(f FuncDef) (args []GoIdentifier, argWrappers []ArgumentWrapperData) {
	for i, a := range f.ArgsT {
		g.shouldGenerate = false

		if a.Name == textLenRegisteredName {
			g.shouldGenerate = true
			argWrappers = append(argWrappers, ArgumentWrapperData{
				VarName: "C.int(len(text))",
			})

			continue
		}

		decl, wrapper, err := getArgWrapper(
			&a,
			i == 0 && f.StructSetter,
			f.StructGetter && g.typedefsNames[a.Type],
			g.context,
		)
		if err != nil {
			if g.context.flags.Verbose {
				glg.Debugf("Unknown argument type \"%s\" in function %s", a.Type, f.FuncName)
			}

			break
		}

		g.shouldGenerate = true
		if len(decl) > 0 {
			args = append(args, GoIdentifier(decl))
			argWrappers = append(argWrappers, wrapper)
		}
	}

	return args, argWrappers
}

// Generate function body
// and returns function call arguments
// e.g.:
// it will write the following into the buffer:
func (g *goFuncsGenerator) generateFuncBody(argWrappers []ArgumentWrapperData) (invokeStatement string, declarations, finishers []string) {
	var invokeStmt []string
	declarations, finishers = make([]string, 0, len(argWrappers)), make([]string, 0, len(argWrappers))

	for _, aw := range argWrappers {
		invokeStmt = append(invokeStmt, aw.VarName)
		var def string
		if aw.NoFin {
			def = aw.ArgDefNoFin
		} else {
			def = aw.ArgDef
		}

		if len(def) > 0 {
			declarations = append(declarations, def)
			if aw.Finalizer != "" && !aw.NoFin {
				finishers = append(finishers, aw.Finalizer)
			}
		}
	}

	return strings.Join(invokeStmt, ","), declarations, finishers
}

func (g *goFuncsGenerator) writeFinishers(shouldDefer bool, finishers []string) {
	if len(finishers) == 0 {
		return
	}

	g.sb.WriteString("\n")

	if shouldDefer {
		g.sb.WriteString("defer func() {\n")
		defer g.sb.WriteString("\n}()\n")
	}

	g.sb.WriteString(strings.Join(finishers, "\n"))
	g.sb.WriteString("\n\n")
}
