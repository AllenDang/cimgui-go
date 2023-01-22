package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
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
func getStructDefs(enumJsonBytes []byte) ([]StructDef, error) {
	// extract only structs section from json
	var structSectionJson StructSection
	err := json.Unmarshal(enumJsonBytes, &structSectionJson)
	if err != nil {
		return nil, fmt.Errorf("cannot extract structs section from json: %w", err)
	}

	var (
		structs    []StructDef
		structJson map[string]json.RawMessage
	)

	err = json.Unmarshal(structSectionJson.Structs, &structJson)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal structs section: %w", err)
	}

	for k, v := range structJson {
		var memberDefs []StructMemberDef

		err := json.Unmarshal(v, &memberDefs)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal struct %s: %w", k, err)
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

	return structs, nil
}

func shouldSkipStruct(name string) bool {
	valueTypeStructs := map[string]bool{
		"ImVec1":      true,
		"ImVec2ih":    true,
		"ImVec2":      true,
		"ImVec4":      true,
		"ImRect":      true,
		"ImColor":     true,
		"ImPlotPoint": true,
		"ImPlotTime":  true,
	}

	if !strings.HasPrefix(name, "Im") {
		return true
	}

	// Skip all value type struct
	if valueTypeStructs[name] {
		return true
	}

	return false
}
