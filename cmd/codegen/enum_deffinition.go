package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// EnumsSection appears in json file on top of enums deffinition section.
type EnumsSection struct {
	EnumComments json.RawMessage `json:"enum_comments"`
	Enums        json.RawMessage `json:"enums"`
}

// EnumDef represents a definition of an ImGui enum (like flags).
type EnumDef struct {
	Name         EnumIdentifier
	CommentAbove string
	Values       []EnumValueDef
}

// EnumValueDef represents a definition of an ImGui enum value.
type EnumValueDef struct {
	Name    CIdentifier `json:"name"`
	Value   int         `json:"calc_value"`
	Comment string      `json:"comment"`
}

type CommentDef struct {
	Above    string `json:"above"`
	Sameline string `json:"sameline"`
}

// getEnumDefs takes a json file bytes (struct_and_enums.json) and returns a slice of EnumDef.
func getEnumDefs(enumJsonBytes []byte) ([]EnumDef, error) {
	var enumSectionJson EnumsSection
	err := json.Unmarshal(enumJsonBytes, &enumSectionJson)
	if err != nil {
		return nil, fmt.Errorf("cannot extract enums section from json: %w", err)
	}

	var enums []EnumDef

	var enumJson map[string]json.RawMessage
	err = json.Unmarshal(enumSectionJson.Enums, &enumJson)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal enums section: %w", err)
	}

	enumCommentsJson := make(map[string]json.RawMessage)
	if enumSectionJson.EnumComments != nil {
		err = json.Unmarshal(enumSectionJson.EnumComments, &enumCommentsJson)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal enum comments section: %w", err)
		}
	}

	for k, v := range enumJson {
		var enumValues []EnumValueDef
		err := json.Unmarshal(v, &enumValues)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal enum %s: %w", k, err)
		}

		enum := EnumDef{
			Name:   EnumIdentifier(k),
			Values: enumValues,
		}

		if commentData, ok := enumCommentsJson[k]; ok {
			var enumCommentValue CommentDef
			err := json.Unmarshal(commentData, &enumCommentValue)
			if err != nil {
				return nil, fmt.Errorf("cannot unmarshal enum comment data %s: %w", k, err)
			}

			enum.CommentAbove = enumCommentValue.Above
		}

		enums = append(enums, enum)
	}

	// sort lexicographically for determenistic generation
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	return enums, nil
}
