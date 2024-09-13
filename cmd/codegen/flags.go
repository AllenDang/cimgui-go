package main

import "flag"

type flags struct {
	// TODO
	showGenerated bool
	// TODO
	showNotGenerated bool

	defJsonPath,
	enumsJsonpath,
	typedefsJsonpath,
	refEnumsJsonPath,
	refTypedefsJsonPath,
	prefix,
	codePrefix,
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
	flag.StringVar(&flags.prefix, "p", "", "prefix for the generated file")
	flag.StringVar(&flags.include, "i", "", "include header file")
	flag.StringVar(&flags.codePrefix, "cp", "", "Add this to evereyhing what is generated")
	flag.Parse()

	return flags
}
