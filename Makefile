NAME=cimgui-go Code Generator

.PHONY: all

## all: generates both bindings.
all: generate

.PHONY: setup
## setup: pulls down dependencies.
setup:
	go get -v -d ./...

## _gencode: is an internal rule. It generates binding basing on env variables set in imgui and implot rules
_generate: setup
	@echo "Generating for $(prefix)"
	go run github.com/AllenDang/cimgui-go/cmd/codegen -p $(prefix) -i $(include_path) -d $(d_file) -e $(e_file) $(r_file)
	gofmt -w $(prefix)_enums.go
	gofmt -w $(prefix)_structs.go
	gofmt -w $(prefix)_funcs.go

## cimgui: generate cimgui binding
.PHONY: cimgui
cimgui: prefix := cimgui
cimgui: include_path := cimgui/cimgui.h
cimgui: d_file :=  cimgui/generator/output/definitions.json 
cimgui: e_file := cimgui/generator/output/structs_and_enums.json
cimgui: r_file :=
cimgui: _generate

## cimplot: generate implot binding
.PHONY: cimplot
cimplot: prefix := cimplot
cimplot: include_path := cimplot/cimplot.h
cimplot: d_file := cimplot/generator/output/definitions.json
cimplot: e_file := cimplot/generator/output/structs_and_enums.json
cimplot: r_file := -r cimgui/generator/output/structs_and_enums.json
cimplot: _generate

compile_cimgui_macos:
	rm -rf ./cimgui/build
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/arm64/

## generate: generates both bindings (equal to `all`)
.PHONY: generate
generate: setup cimgui cimplot
