package main

import (
	"encoding/json"
	"sort"
)

// FuncDef represents a definition of a function in cimgui's json format
type FuncDef struct {
	// FuncName is a cimgui name of the function
	FuncName string `json:"ov_cimguiname"`
	// FuncName is an original Dear ImGui's name of the function
	OriginalFuncName string `json:"original_func_name"`
	// Args represents a plain string list of function arguments
	Args string `json:"args"`
	// ArgsT represents a more accurate list of arg:type list
	ArgsT []ArgDef `json:"argsT"`
	// Defaults represents a map (arg name : value) of a default C value for
	// function's argument
	Defaults map[string]string `json:"defaults"`
	// Location represents an exact location in DearImGui's source code
	Location string `json:"location"`

	Constructor  bool `json:"constructor"`
	Destructor   bool `json:"destructor"`
	StructSetter bool `json:"struct_setter"`
	StructGetter bool `json:"struct_getter"`

	InvocationStmt string `json:"invocation_stmt"`

	// Ret determines a type of returned value
	Ret    string `json:"ret"`
	StName string `json:"stname"`

	// NonUDT is != 0 whenever the function has a pointer argument
	// that could be considered as its result.
	// e.g. igCalcItemSize takes pOut ImVec2* argument, but it is used
	// to store resulting value.
	NonUDT int `json:"nonUDT"`
}

// ArgDef represents a deffinition of function's argument
type ArgDef struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	CustomType string `json:"custom_type"`
}

// getFunDefs takes a json file bytes as an argument and returns
// a list of cimgui functions read.
func getFunDefs(defJsonBytes []byte) []FuncDef {
	var defJson map[string]json.RawMessage
	err := json.Unmarshal(defJsonBytes, &defJson)
	if err != nil {
		panic(err.Error())
	}

	var funcs []FuncDef

	for _, v := range defJson {
		var funcDefs []FuncDef
		_ = json.Unmarshal(v, &funcDefs)

		funcs = append(funcs, funcDefs...)
	}

	// sort lexicographically for determenistic generation
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].FuncName < funcs[j].FuncName
	})

	return funcs
}
