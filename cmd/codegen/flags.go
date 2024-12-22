package main

import "flag"

type flags struct {
	showGenerated    bool
	showNotGenerated bool

	defJsonPath,
	enumsJsonpath,
	typedefsJsonpath,
	refEnumsJsonPath,
	refTypedefsJsonPath,
	presetJsonPath,
	refPackageName, // name for refTypedefs (default: imgui)
	packageName, // name for current package (e.g. imgui, implot)
	include string
}

func parse() *flags {
	flags := &flags{}
	flag.BoolVar(&flags.showGenerated, "generated", false, "Log about functions that was generated.")
	flag.BoolVar(&flags.showNotGenerated, "not-generated", true, "Log about functions that was NOT generated.")

	flag.StringVar(&flags.defJsonPath, "d", "", "definitions json file path")
	flag.StringVar(&flags.enumsJsonpath, "e", "", "structs and enums json file path")
	flag.StringVar(&flags.typedefsJsonpath, "t", "", "typedefs dict json file path")
	flag.StringVar(&flags.refEnumsJsonPath, "r", "", "reference structs and enums json file path")
	flag.StringVar(&flags.refTypedefsJsonPath, "rt", "", "reference typedefs_dict.json file path")
	flag.StringVar(&flags.presetJsonPath, "preset", "", "Preset of custom (manual) generator rules. See go doc github.com/AllenDang/cimgui-go/cmd/codegen.Preset for more details.")
	flag.StringVar(&flags.refPackageName, "refPkg", "imgui", "name for refTypedefs package name")
	flag.StringVar(&flags.packageName, "pkg", "", "name for current package")
	flag.StringVar(&flags.include, "i", "", "include header file (source project's header)")
	flag.Parse()

	return flags
}
