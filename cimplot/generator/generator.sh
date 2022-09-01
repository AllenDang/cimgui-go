#!/bin/sh

# this script must be executed in this directory
# all the output goes to generator/output folder
# .cpp and .h files:
# cimplot.h and cimplot.cpp with gcc preprocess
# lua and json files:
# definitions.lua for function definitions
# structs_and_enums.lua with struct and enum information-definitions
# impl_definitions.lua for implementation function definitions

#process  files
# arg[1] compiler name gcc, clang, or cl
# arg[2] options as words in one string: internal for imgui_internal generation
# examples: "" "internal"
luajit ./generator.lua gcc "internal"
