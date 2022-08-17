package main

import (
	"encoding/json"
	"flag"
	"os"
)

type ArgDef struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type FuncDef struct {
	Args         string   `json:"args"`
	ArgsT        []ArgDef `json:"argsT"`
	FuncName     string   `json:"cimguiname"`
	Location     string   `json:"location"`
	Constructor  bool     `json:"constructor"`
	Destructor   bool     `json:"destructor"`
	StructSetter bool     `json:"struct_setter"`
	Ret          string   `json:"ret"`
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

func main() {
	defJsonPath := flag.String("d", "", "definitions json file path")
	enumsJsonpath := flag.String("e", "", "enums json file path")

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

	var defJson map[string]json.RawMessage
	err = json.Unmarshal(defJsonBytes, &defJson)
	if err != nil {
		panic(err.Error())
	}

	var funcs []FuncDef

	for _, v := range defJson {
		var funcDefs []FuncDef
		_ = json.Unmarshal(v, &funcDefs)

		if len(funcDefs) == 1 {
			funcs = append(funcs, funcDefs[0])
		}
	}

	var enumSectionJson EnumsSection
	err = json.Unmarshal(enumJsonBytes, &enumSectionJson)
	if err != nil {
		panic(err.Error())
	}

	var enums []EnumDef

	var enumJson map[string]json.RawMessage
	err = json.Unmarshal(enumSectionJson.Enums, &enumJson)
	if err != nil {
		panic(err.Error())
	}

	for k, v := range enumJson {
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

	var structSectionJson StructSection
	err = json.Unmarshal(enumJsonBytes, &structSectionJson)
	if err != nil {
		panic(err.Error())
	}

	var structs []StructDef

	var structJson map[string]json.RawMessage
	err = json.Unmarshal(structSectionJson.Structs, &structJson)
	if err != nil {
		panic(err.Error())
	}
	for k, v := range structJson {
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

	validFuncs := generateCppWrapper(funcs)

	enumNames := generateGoEnums(enums)
	structNames := generateGoStructs(structs)

	structAccessorFuncs := generateCppStructsAccessor(structs)
	validFuncs = append(validFuncs, structAccessorFuncs...)

	generateGoFuncs(validFuncs, enumNames, structNames)
}
