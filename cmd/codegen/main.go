package main

import (
	"encoding/json"
	"flag"
	"os"
	"strings"

	"github.com/thoas/go-funk"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type ArgDef struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	CustomType string `json:"custom_type"`
}

type FuncDef struct {
	FuncName         string            `json:"ov_cimguiname"`
	OriginalFuncName string            `json:"original_func_name"`
	Args             string            `json:"args"`
	ArgsT            []ArgDef          `json:"argsT"`
	Defaults         map[string]string `json:"defaults"`
	Location         string            `json:"location"`
	Constructor      bool              `json:"constructor"`
	Destructor       bool              `json:"destructor"`
	StructSetter     bool              `json:"struct_setter"`
	StructGetter     bool              `json:"struct_getter"`
	InvocationStmt   string            `json:"invocation_stmt"`
	Ret              string            `json:"ret"`
	StName           string            `json:"stname"`
}

type EnumValueDef struct {
	Name  string `json:"name"`
	Value int    `json:"calc_value"`
}

type EnumDef struct {
	Name   string
	Values []EnumValueDef
}

type EnumsSection struct {
	Enums json.RawMessage `json:"enums"`
}

type StructMemberDef struct {
	Name         string `json:"name"`
	TemplateType string `json:"template_type"`
	Type         string `json:"type"`
	Size         int    `json:"size"`
}

type StructDef struct {
	Name    string            `json:"name"`
	Members []StructMemberDef `json:"members"`
}

type StructSection struct {
	Structs json.RawMessage `json:"structs"`
}

func getFunDefs(defJsonBytes []byte) []FuncDef {
	var defJson map[string]json.RawMessage
	err := json.Unmarshal(defJsonBytes, &defJson)
	if err != nil {
		panic(err.Error())
	}

	var funcs []FuncDef

	keys := maps.Keys(defJson)
	slices.Sort(keys)
	for _, k := range keys {
		v := defJson[k]
		var funcDefs []FuncDef
		_ = json.Unmarshal(v, &funcDefs)

		funcs = append(funcs, funcDefs...)
	}

	return funcs
}

func getEnumDefs(enumJsonBytes []byte) []EnumDef {
	var enumSectionJson EnumsSection
	err := json.Unmarshal(enumJsonBytes, &enumSectionJson)
	if err != nil {
		panic(err.Error())
	}

	var enums []EnumDef

	var enumJson map[string]json.RawMessage
	err = json.Unmarshal(enumSectionJson.Enums, &enumJson)
	if err != nil {
		panic(err.Error())
	}

	keys := maps.Keys(enumJson)
	slices.Sort(keys)
	for _, k := range keys {
		v := enumJson[k]
		var enumValues []EnumValueDef
		err := json.Unmarshal(v, &enumValues)
		if err != nil {
			panic(err.Error())
		}

		enums = append(enums, EnumDef{
			Name:   k,
			Values: enumValues,
		})
	}

	return enums
}

func getStructDefs(enumJsonBytes []byte) []StructDef {
	var structSectionJson StructSection
	err := json.Unmarshal(enumJsonBytes, &structSectionJson)
	if err != nil {
		panic(err.Error())
	}

	var structs []StructDef

	var structJson map[string]json.RawMessage
	err = json.Unmarshal(structSectionJson.Structs, &structJson)
	if err != nil {
		panic(err.Error())
	}
	keys := maps.Keys(structJson)
	slices.Sort(keys)
	for _, k := range keys {
		v := structJson[k]
		var memberDefs []StructMemberDef
		err := json.Unmarshal(v, &memberDefs)
		if err != nil {
			panic(err.Error())
		}

		structs = append(structs, StructDef{
			Name:    k,
			Members: memberDefs,
		})
	}

	return structs
}

func shouldSkipStruct(name string) bool {
	valueTypeStructs := []string{
		"ImVec1",
		"ImVec2ih",
		"ImVec2",
		"ImVec4",
		"ImRect",
		"ImColor",
		"ImPlotPoint",
	}

	if !strings.HasPrefix(name, "Im") {
		return true
	}

	// Skip all value type struct
	if funk.ContainsString(valueTypeStructs, name) {
		return true
	}

	return false
}

func getEnumAndStructNames(enumJsonBytes []byte) (enumNames []string, structNames []string) {
	enums := getEnumDefs(enumJsonBytes)
	structs := getStructDefs(enumJsonBytes)

	for _, e := range enums {
		enumNames = append(enumNames, strings.TrimSuffix(e.Name, "_"))
	}

	for _, s := range structs {
		if !shouldSkipStruct(s.Name) {
			structNames = append(structNames, s.Name)
		}
	}

	return
}

func main() {
	defJsonPath := flag.String("d", "", "definitions json file path")
	enumsJsonpath := flag.String("e", "", "structs and enums json file path")
	refEnumsJsonPath := flag.String("r", "", "reference structs and enums json file path")
	prefix := flag.String("p", "", "prefix for the generated file")
	include := flag.String("i", "", "include header file")

	flag.Parse()

	stat, err := os.Stat(*defJsonPath)
	if err != nil || stat.IsDir() {
		panic("Invalid definitions json file path")
	}

	stat, err = os.Stat(*enumsJsonpath)
	if err != nil || stat.IsDir() {
		panic("Invalid enum json file path")
	}

	defJsonBytes, err := os.ReadFile(*defJsonPath)
	if err != nil {
		panic(err.Error())
	}

	enumJsonBytes, err := os.ReadFile(*enumsJsonpath)
	if err != nil {
		panic(err.Error())
	}

	var refEnumJsonBytes []byte
	if len(*refEnumsJsonPath) > 0 {
		refEnumJsonBytes, err = os.ReadFile(*refEnumsJsonPath)
		if err != nil {
			panic(err.Error())
		}
	}

	// get definitions from json file
	funcs := getFunDefs(defJsonBytes)

	enums := getEnumDefs(enumJsonBytes)

	structs := getStructDefs(enumJsonBytes)

	validFuncs := generateCppWrapper(*prefix, *include, funcs)

	// generate code
	enumNames := generateGoEnums(*prefix, enums)
	structNames := generateGoStructs(*prefix, structs)

	structAccessorFuncs := generateCppStructsAccessor(*prefix, validFuncs, structs)
	validFuncs = append(validFuncs, structAccessorFuncs...)

	// generate reference only enum and struct names
	if len(refEnumJsonBytes) > 0 {
		es, ss := getEnumAndStructNames(refEnumJsonBytes)
		enumNames = append(enumNames, es...)
		structNames = append(structNames, ss...)
	}

	generateGoFuncs(*prefix, validFuncs, enumNames, structNames)
}
