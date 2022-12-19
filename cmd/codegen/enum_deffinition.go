package main

import (
	"encoding/json"
	"sort"
)

// EnumsSection appears in json file on top of enums deffinition section.
type EnumsSection struct {
	Enums json.RawMessage `json:"enums"`
}

// EnumDef represents a definition of an ImGui enum (like flags).
type EnumDef struct {
	Name   string
	Values []EnumValueDef
}

// EnumValueDef represents a definition of an ImGui enum value.
type EnumValueDef struct {
	Name  string `json:"name"`
	Value int    `json:"calc_value"`
}

// getEnumDefs takes a json file bytes (struct_and_enums.json) and returns a slice of EnumDef.
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

	// sort lexicographically for determenistic generation
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	return enums
}
