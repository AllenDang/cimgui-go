# codegen

Codegen is a cimgui-go-featured GO code generator.

Cod is generated basing on several JSON config files and refers to a shared library compiled from C source.

# Usage

```console
[codegen (0) ]$ go run . --help
Usage: codegen [FLAG]...

Flags:
    -lg, --log-generated         Log about functions that was generated.                                                                                                             (type: bool; default: false)
    -lng, --log-not-generated    Log about functions that was NOT generated.                                                                                                         (type: bool; default: false)
    -d, --definitions            definitions.json file path                                                                                                                          (type: string)
    -e, --enums                  structs_and_enums.json file path                                                                                                                    (type: string)
    -t, --typedefs               typedefs_dict.json file path                                                                                                                        (type: string)
    -re, --ref-enums             structs_and_enums.json file path for reference package (see README.md).                                                                             (type: string)
    -rt, --ref-typedefs          typedefs_dict.json file path for reference package (see README.md).                                                                                 (type: string)
    -p, --preset                 Preset of custom (manual) generator rules. See go doc github.com/AllenDang/cimgui-go/cmd/codegen.Preset for more details. (in form of json file)    (type: string)
    -rp, --ref-pkg               name for refTypedefs package name                                                                                                                   (type: string; default: imgui)
    -pkg, --package              name for current package                                                                                                                            (type: string)
    -i, --include                Include header file (source project's header)                                                                                                       (type: string)
    -h, --help                   show help                                                                                                                                           (type: bool)
```

# Configuration

```console
[codegen (0) ]$ go doc Preset
type Preset struct {
	// SkipFuncs functions (from definitions.json) to be skipped
	// e.g. they are temporarily hard-coded
	SkipFuncs map[CIdentifier]bool
	// SkipStructs allows to specify struct names that will be skipped.
	SkipStructs map[CIdentifier]bool
	// SkipMethods struct names from structs_and_enums.json.
	// structures that's METHODS should be skipped
	SkipMethods map[CIdentifier]bool
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
	// OriginReplace allows to force-replace function name with some other name.
	// Introduced to replace TextEditor_GetText -> TextEditor_GetText_alloc
	// but could be re-used to force use of another function than json tells us to use.
	//
	// It differs from Replace - Replace renames an identifier in general (changes its name but refers to the same function).
	// This allows to absolutely abandon the source C function and use some OTHER C function.
	OriginReplace map[CIdentifier]CIdentifier
	// DefaultArgReplace is used in C-side default args generation (gencpp.go).
	// cimgui-go uses this to change FLT_MIN with igGet_FLT_MIN()
	// NOTE: Iterated randomly!
	DefaultArgReplace map[CIdentifier]CIdentifier
	// DefaultArgArbitraryValue is similar to the above DefaultArgReplace, but
	// associates default arg name with any arbitrary value.
	// cimgui-go uses this to set text_end to 0
	DefaultArgArbitraryValue map[CIdentifier]CIdentifier
	// ExtraCGOPreamble allows to specify additional C code elements in Go files.
	// For example could be used to Include extra files.
	// For ease of use, its in form of []string. These lines will be merged and prefixed (if appliable) with '//'
	ExtraCGOPreamble []string
	// InternalFiles allows to specify files that are considered Internal.
	// If an identifier is found in such a file, it will be generated but its
	// name will be prefixed with InternalPrefix
	InternalFiles []string
	// InternalPrefix is a prefix for identifiers from InternalFiles.
	InternalPrefix string
	// PackagePath is an import path of the package.
	// This is base path. flags.PackageName will be added.
	// Example:
	//   "github.com/AllenDang/cimgui-go"
	//   If enerated with -pkg imgui, import path
	//   is supposed to be "github.com/AllenDang/cimgui-go/imgui"
	PackagePath string
	// SimpleTypes are used for simple (go-convertable) custom types (wrappers will be generated by simpleW/simpleR).
	// Example:
	//   ImS16 is defined as short in C code, so Go can easily convert it via int16()
	// Expected format is:
	//   "ImS16": ["int16", "C.ImS16", "pkgname"]
	// where:
	// - ImS16 is a C type name
	// - int16 is a Go-friendly type name
	// - C.ImS16 is a cgo compatible type name
	// - pkgname is a source package for the type (in this case int16 is a builtin so it should be empty)
	// See also: simpleW
	SimpleTypes map[CIdentifier][3]GoIdentifier
	// SimplePtrTypes are just like SimpleTypes but for pointer types.
	// Example:
	//   "ImS16": ["int16", "C.ImS16", "pkgname"]
	SimplePtrTypes map[CIdentifier][3]GoIdentifier
	// WrappableTypes are types that implement a special interface
	// github.com/AllenDang/cimgui-go/internal.WrappableType[CTYPE, SELF]
	// In short they are supposed to have 2 methods:
	// - ToC() CTYPE which returns SELF converted to CTYPE
	// - FromC(CTYPE) SELF which restores SELF from CTYPE.
	//
	// Key is a C type name, value is a list:
	// - Go name
	// - C name
	// - package where the type is defined
	// Example syntax:
	//    "ImVec2": ["Vec2", "C.ImVec2", "imgui"]
	WrappableTypes map[CIdentifier][3]GoIdentifier
}
    Preset is a set of rules iperative to XXX_templates/ json files. They are
    used to manually set some properties of the generator.
```

# How it works?

Assume you have a C project which exports its public API in a header (`.h`) file.
What this module does not cover is the first step:
1. You need JSON files describing functions in your C project. This is at the moment done by https://github.com/cimgui lua cpp2ffi.lua script which
however is highly single-purpose and you probably will not be able to use it for your project. You need to handle this yourself for now.
Let us know if you find/create an universal solution for that. For specification of the JSON files see [here](#C-Configuration)
2. Then you need a another JSON file called "preset" (see [here](#Configuration)) which will tell the generator more about your project.
e.g. its github path, ids you want to exclude (if you're need for some reason - we do for imgui), prefixes to remove, etc.
3. Execute codegen. Do `codegen -lng --preset your_preset.json -e structs_and_enums.json -t typedefs_dict.json -d definitions.json -i your_header.h -pkg your_package_name`.
This will create several files in your current directory that should contain operational code.

## Reference packages

Sometimes you may want to generate one package basing on another package wrapped by this
(I mean imgui plugin for example).
Example:
`header.h` contains symbol `A`. Then you have `header2.h` with a function using this symbol `A`.
On go sied you have `package1.A`  and you want to generate `header2.h` in `package2` and
use `package1.A` in it. This is where `--ref*` flags comes in. You specify header1's json files as reference files
and package1 as reference package. Then you generate package2 and it will use package1's symbols.

# C Configuration

:warning: Disclaimer: this is NOT our design and we are not able to modify this.
The layout of these files is enforced by https://github.com/cimgui which we're using for generation.

The minimalistic working configuration is present  in [here](./testdata) review it yourself.
