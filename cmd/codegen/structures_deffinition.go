package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// StructSection appears in json file on top of structs definition section.
type StructSection struct {
	StructComments json.RawMessage `json:"struct_comments"`
	Structs        json.RawMessage `json:"structs"`
}

// StructDef represents a definition of an ImGui struct.
type StructDef struct {
	Name         CIdentifier `json:"name"`
	CommentAbove string
	Members      []StructMemberDef `json:"members"`
}

// StructMemberDef represents a definition of an ImGui struct member.
type StructMemberDef struct {
	Name         CIdentifier `json:"name"`
	TemplateType CIdentifier `json:"template_type"`
	Type         CIdentifier `json:"type"`
	Size         int         `json:"size"`
	Comment      CommentDef
	CommentData  json.RawMessage `json:"comment"`
	Bitfield     string          `json:"bitfield"`
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
		structs           []StructDef
		structJson        map[string]json.RawMessage
		structCommentJson map[string]json.RawMessage = make(map[string]json.RawMessage)
	)

	err = json.Unmarshal(structSectionJson.Structs, &structJson)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal structs section: %w", err)
	}

	if structSectionJson.StructComments != nil && string(structSectionJson.StructComments) != "[]" {
		err = json.Unmarshal(structSectionJson.StructComments, &structCommentJson)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal struct's comments section: %w", err)
		}
	}

	for k, v := range structJson {
		var memberDefs []StructMemberDef

		err := json.Unmarshal(v, &memberDefs)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal struct %s: %w", k, err)
		}

		for i, member := range memberDefs {
			if member.CommentData != nil {
				var comment CommentDef
				err = json.Unmarshal(member.CommentData, &comment)
				if err != nil {
					return nil, fmt.Errorf("cannot unmarshal struct comment %s: %w", member.CommentData, err)
				}

				memberDefs[i].Comment = comment
			}
		}

		str := StructDef{
			Name:    CIdentifier(k),
			Members: memberDefs,
		}

		if commentData, ok := structCommentJson[k]; ok {
			var structCommentValue CommentDef
			err := json.Unmarshal(commentData, &structCommentValue)
			if err != nil {
				return nil, fmt.Errorf("cannot unmarshal enum comment data %s: %w", k, err)
			}

			str.CommentAbove = structCommentValue.Above
		}

		structs = append(structs, str)
	}

	// sort lexicographically for determenistic generation
	sort.Slice(structs, func(i, j int) bool {
		return structs[i].Name < structs[j].Name
	})

	return structs, nil
}

func shouldSkipStruct(name CIdentifier) bool {
	valueTypeStructs := map[CIdentifier]bool{
		"ImVec2ih":    true,
		"ImVec2":      true,
		"ImVec4":      true,
		"ImRect":      true,
		"ImColor":     true,
		"ImPlotPoint": true,
		"ImPlotTime":  true,
	}

	// Skip all value type struct
	if valueTypeStructs[name] {
		return true
	}

	return false
}
