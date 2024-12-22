package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// FuncDef represents a definition of a function in cimgui's json format
type FuncDef struct {
	// FuncName is a cimgui name of the function
	FuncName CIdentifier `json:"ov_cimguiname"`
	// OriginalFuncName is an original Dear ImGui's name of the function
	OriginalFuncName CIdentifier `json:"original_func_name"`
	// Args represents a plain string list of function arguments
	Args string `json:"args"`
	// ArgsT represents a more accurate list of arg:type list
	ArgsT []ArgDef `json:"argsT"`
	// Defaults represents a map (arg name : value) of a default C value for
	// function's argument
	Defaults map[string]string `json:"defaults"`
	// Location represents an exact location in DearImGui's source code
	Location string `json:"location"`
	// Comment on function
	Comment string `json:"comment"`

	Constructor  bool
	Destructor   bool
	StructSetter bool
	StructGetter bool

	InvocationStmt string

	// Ret determines a type of returned value
	Ret    CIdentifier `json:"ret"`
	StName string      `json:"stname"`

	// NonUDT is != 0 whenever the function has a pointer argument
	// that could be considered as its result.
	// e.g. igCalcItemSize takes pOut ImVec2* argument, but it is used
	// to store resulting value.
	NonUDT int `json:"nonUDT"`

	// The names could collide between cimgui and the function generated
	// here, so we may need to add a prefix.
	CWrapperFuncName CIdentifier

	// AllCallArgs is used to determine if a wrapper is needed: if the cimgui
	// function takes the exact same arguments as the wrapper, the wrapper is
	// not needed.
	AllCallArgs CIdentifier
}

// ArgDef represents a deffinition of function's argument
type ArgDef struct {
	Name       CIdentifier `json:"name"`
	Type       CIdentifier `json:"type"`
	CustomType CIdentifier `json:"custom_type"`

	// this should be set to true if memory should not be freed by GO
	RemoveFinalizer bool
}

// getFunDefs takes a json file bytes as an argument and returns
// a list of cimgui functions read.
func getFunDefs(defJsonBytes []byte) ([]FuncDef, error) {
	var defJson map[string]json.RawMessage
	err := json.Unmarshal(defJsonBytes, &defJson)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %w", err)
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

	return funcs, nil
}
