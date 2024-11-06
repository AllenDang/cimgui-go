package main

// Preset is a set of rules iperative to XXX_templates/ json files.
// They are used to manually set some properties of the generator.
type Preset struct {
	// SkipFuncs functions (from definitions.json) to be skipped
	// e.g. they are temporarily hard-coded
	SkipFuncs map[CIdentifier]bool
	// SkipStructs struct names from structs_and_enums.json.
	// structures that's METHODS should be skipped
	SkipStructs map[CIdentifier]bool
	// SkipTypedefs typedefs from typedefs_dict.json to be skipped
	// e.g. for hardcoded typedefs or typedefs which are obvious (e.g. ImU16 becomes uint16 without extra type information)
	SkipTypedefs map[CIdentifier]bool
	// TypedefsPoolSize sets a default size for callbacks pool.
	// Rembmber to set this as it defaults to 0 and you'll get no callbacks!
	TypedefsPoolSize int
	// TypedefsCustomPoolSizes allows to override TypedefsPoolSize for certain types.
	TypedefsCustomPoolSizes map[CIdentifier]int
	// Replace is a map for C -> Go names conversion.
	// It allows you to force-rename anything (including functions and enums)
	Replace map[CIdentifier]GoIdentifier
	// TrimPrefix allows to remove unwanted prefixes from everything during C->Go renaming.
	// NOTE: order sensitive!
	// NOTE: Case sensitive
	TrimPrefix []string
}
