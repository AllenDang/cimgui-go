package main

// Preset is a set of rules iperative to XXX_templates/ json files.
// They are used to manually set some properties of the generator.
type Preset struct {
	// SkipFuncs functions (from definitions.json) to be skipped
	// e.g. they are temporarily hard-coded
	SkipFuncs map[CIdentifier]bool
	// SkipStructs structs from structs_and_enums.json to be skipped
	SkipStructs map[CIdentifier]bool
	// SkipTypedefs typedefs from typedefs_dict.json to be skipped
	SkipTypedefs map[CIdentifier]bool
}
