package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/thoas/go-funk"
)

// StructSection appears in json file on top of structs definition section.
type StructSection struct {
	Structs json.RawMessage `json:"structs"`
}

// StructDef represents a definition of an ImGui struct.
type StructDef struct {
	Name    string            `json:"name"`
	Members []StructMemberDef `json:"members"`
}

// StructMemberDef represents a definition of an ImGui struct member.
type StructMemberDef struct {
	Name         string `json:"name"`
	TemplateType string `json:"template_type"`
	Type         string `json:"type"`
	Size         int    `json:"size"`
}

// getStructDefs takes a json file bytes (struct_and_enums.json) and returns a slice of StructDef.
func getStructDefs(enumJsonBytes []byte) []StructDef {
	// extract only structs section from json
	var structSectionJson StructSection
	err := json.Unmarshal(enumJsonBytes, &structSectionJson)
	if err != nil {
		panic(err.Error())
	}

	var (
		structs    []StructDef
		structJson map[string]json.RawMessage
	)

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

	// sort lexicographically for determenistic generation
	sort.Slice(structs, func(i, j int) bool {
		return structs[i].Name < structs[j].Name
	})

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
