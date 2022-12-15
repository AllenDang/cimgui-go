NAME=cimgui-go Code Generator

.PHONY: all setup cimgui cimplot generate

## all: generates both bindings.
all: generate

## setup: pulls down dependencies.
setup:
	go get -v -d ./...

## _gencode: is an internal rule. It generates binding basing on env variables set in imgui and implot rules
_generate: setup
	echo "Generating for $(prefix)"
	cd ./cmd/codegen/build; ./codegen -p $(prefix) -i $(include_path) -d $(d_file) -e $(e_file) $(r_file)
	gofmt -w ./cmd/codegen/build/$(prefix)_enums.go
	gofmt -w ./cmd/codegen/build/$(prefix)_structs.go
	gofmt -w ./cmd/codegen/build/$(prefix)_funcs.go
	cp -f ./cmd/codegen/build/$(prefix)_wrapper.h ./
	cp -f ./cmd/codegen/build/$(prefix)_wrapper.cpp ./
	cp -f ./cmd/codegen/build/$(prefix)_structs_accessor.h ./
	cp -f ./cmd/codegen/build/$(prefix)_structs_accessor.cpp ./
	cp -f ./cmd/codegen/build/$(prefix)_enums.go ./
	cp -f ./cmd/codegen/build/$(prefix)_funcs.go ./
	cp -f ./cmd/codegen/build/$(prefix)_structs.go ./

## cimgui: generate cimgui binding
cimgui: prefix := cimgui
cimgui: include_path := cimgui/cimgui.h
cimgui: d_file :=  ../../../cimgui/generator/output/definitions.json 
cimgui: e_file := ../../../cimgui/generator/output/structs_and_enums.json
cimgui: r_file :=
cimgui: _generate _clean

## cimplot: generate implot binding
cimplot: prefix := cimplot
cimplot: include_path := cimplot/cimplot.h
cimplot: d_file := ../../../cimplot/generator/output/definitions.json
cimplot: e_file := ../../../cimplot/generator/output/structs_and_enums.json
cimplot: r_file := -r ../../../cimgui/generator/output/structs_and_enums.json
cimplot: _generate _clean

compile_cimgui_macos:
	rm -rf ./cimgui/build
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/arm64/

## generate: generates both bindings (equal to `all`)
generate: setup cimgui cimplot

_clean:
	rm ./cmd/codegen/build/$(prefix)_wrapper.h
	rm ./cmd/codegen/build/$(prefix)_wrapper.cpp
	rm ./cmd/codegen/build/$(prefix)_structs_accessor.h
	rm ./cmd/codegen/build/$(prefix)_structs_accessor.cpp
	rm ./cmd/codegen/build/$(prefix)_enums.go
	rm ./cmd/codegen/build/$(prefix)_funcs.go
	rm ./cmd/codegen/build/$(prefix)_structs.go
	rm -rf ./cmd/codegen/build
