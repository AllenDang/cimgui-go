package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/kpango/glg"
)

var callbackNotGeneratedError = errors.New("callback was not generated")

// This is an underlying function for gengo_typedefs.go for now.
// May be reworked to be a separated "genertor" in the future.
// - includes logging
// - returns CallbackNotGeneratedError
func (g *typedefsGenerator) writeCallback(typedefName CIdentifier, def string) error {
	// see https://github.com/AllenDang/cimgui-go/issues/224#issuecomment-2452156237
	// 1: preprocessing - parse typedefs.data[k] to get return type and arguments

	// now let me use a bit of regex.
	// We have 2 possibilities:
	// - returnType(*<funcName>)(args1 arg1Name, args2 arg2Name, args3 arg3Name);
	// - returnType <funcName>(args1 arg1Name, args2 arg2Name, args3 arg3Name);
	// NOTE: the second is uesed mainly in immarkdown
	// NOTE: in the 1st, spaces does not matter so we'll trim them
	expr1, err := regexp.Compile("([a-zA-Z0-9_]+\\*?)\\(\\*.*\\)\\((.*)\\);")
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

	if ok := expr1.MatchString(def); ok {
		glg.Debugf("callback typedef \"%s\" is in form 1", typedefName)
		// now split by "("
		// it should be something like this:
		// ["returnType", "*<optional func name)", "args1 arg1Name, args2 arg2Name, args3 arg3Name);"]
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
		glg.Warnf("Callback option 2 for %s not implemented yet", typedefName)
		return callbackNotGeneratedError
	} else {
		if g.ctx.flags.showNotGenerated {
			return fmt.Errorf("cannot parse callback typedef: %w", callbackNotGeneratedError)
		}
	}

	// 2: Find wrappers:
	// We need to figure out how to wrap returnType and args.
	// In fact, we need to swap meaning of them, because we want to convert C argument type to Go argument type
	// so we are supposed to use returnWrapper for that.
	glg.Debugf("From %s got \"%s\" and \"%v\"", typedefName, returnTypeC, argsC)

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
			return fmt.Errorf("unable to get arg wrapper for \"%s\" - \"%s\" %w: %w", typedefName, arg.Type, callbackNotGeneratedError, err)
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
	fmt.Fprintf(g.GoSb, `
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
		typedefName.renameGoIdentifier(),
		returnType.ArgType,
		returnType.CType,
		goCallStmt,
		cCallStmt,
		typedefName,
	)

	cCallStmt2 := cCallStmt
	if cCallStmt2 != "" {
		cCallStmt2 = ", " + cCallStmt2
	}

	if returnType.ArgType == "" {
		fmt.Fprintf(g.GoSb, `
func wrap%[1]s(cb %[1]s %[3]s) %[2]s {
	cb(%[4]s)
}
`,
			typedefName.renameGoIdentifier(),
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
		fmt.Fprintf(g.GoSb, `
func wrap%[1]s(cb %[1]s %[5]s) %[2]s {
	result := cb(%[6]s)
	%[3]s
	return %[4]s
}
`,
			typedefName.renameGoIdentifier(),
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
		)
	}

	size := TypedefsPoolSize
	if h, ok := customPoolSize[typedefName]; ok {
		size = h
	}

	poolNames := make([]string, size)
	// now write N functions
	for i := 0; i < size; i++ {
		fmt.Fprintf(g.GoSb,
			`//export callback%[1]s%[2]d
func callback%[1]s%[2]d(%[5]s) %[3]s { %[4]s wrap%[1]s(pool%[1]s.Get(%[2]d), %[6]s) }
					`,
			typedefName.renameGoIdentifier(),
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

		fmt.Fprintf(g.CGoHeaderSb,
			`// extern %[1]s callback%[2]s%[3]d(%[4]s);
`,
			returnTypeC,
			typedefName.renameGoIdentifier(),
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

		poolNames[i] = fmt.Sprintf("C.%[3]s(C.callback%[1]s%[2]d)", typedefName.renameGoIdentifier(), i, typedefName)
	}

	fmt.Fprintf(g.GoSb, `
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
		typedefName.renameGoIdentifier(),
		strings.Join(poolNames, ",\n")+",\n",
		typedefName,
	)

	return nil
}
