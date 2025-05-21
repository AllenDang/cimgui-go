package main

import "github.com/cosiner/flag"

type flags struct {
	ShowGenerated    bool `names:"-lg, --log-generated" usage:"Log about functions that was generated." default:"false"`
	ShowNotGenerated bool `names:"-lng, --log-not-generated" usage:"Log about functions that was NOT generated." default:"false"`
	Verbose          bool `names:"--verbose" usage:"Verbose output (dump literally everything it could dump)." default:"false`

	DefJsonPath      string `names:"-d, --definitions" usage:"definitions.json file path"`
	EnumsJsonpath    string `names:"-e, --enums" usage:"structs_and_enums.json file path"`
	TypedefsJsonpath string `names:"-t, --typedefs" usage:"typedefs_dict.json file path"`

	RefEnumsJsonPath    string `names:"-re, --ref-enums" usage:"structs_and_enums.json file path for reference package (see README.md)."`
	RefTypedefsJsonPath string `names:"-rt, --ref-typedefs" usage:"typedefs_dict.json file path for reference package (see README.md)."`

	PresetJsonPath string `names:"-p, --preset" usage:"Preset of custom (manual) generator rules. See go doc github.com/AllenDang/cimgui-go/cmd/codegen.Preset for more details. (in form of json file)"`

	// name for refTypedefs (default: imgui)
	RefPackageName string `names:"-rp, --ref-pkg" usage:"name for refTypedefs package name" default:"imgui"`
	// name for current package (e.g. imgui, implot)
	PackageName string `names:"-pkg, --package" usage:"name for current package"`
	Include     string `names:"-i, --include" usage:"Include header file (source project's header)"`
	// RefInclude is a reference include path (e.g. for cimgui.h if the extension doesn't include it by default.
	RefInclude string `names:"-ri, --ref-include" usage:"Include reference header file (reference project's header - e.g. cimgui.h)"`
}

func parse() *flags {
	flags := &flags{}
	flag.ParseStruct(flags)

	return flags
}
