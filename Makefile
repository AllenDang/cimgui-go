NAME=cimgui-go Code Generator

.PHONY: all

## all: generates both bindings.
all: generate

.PHONY: setup
## setup: pulls down dependencies.
setup:
	go get -v -d ./...

# Parameters:
# $1: prefix
# $2: include path (of header file)
# $3: definitions.json filepath
# $4: structs_and_enums.json filepath
# $5: additional agruments to codegen call (e.g. -r option)
define generate
	@echo "Generating for $(1)"
	go run github.com/AllenDang/cimgui-go/cmd/codegen -p $(1) -i $(2) -d $(3) -e $(4) $(5)
	gofmt -w $(1)_enums.go
	gofmt -w $(1)_structs.go
	gofmt -w $(1)_funcs.go
endef

## cimgui: generate cimgui binding
.PHONY: cimgui
cimgui: setup
	$(call generate,cimgui,cimgui/cimgui.h,cimgui/generator/output/definitions.json,cimgui/generator/output/structs_and_enums.json)

## cimplot: generate implot binding
.PHONY: cimplot
cimplot: setup
	$(call generate,cimplot,cimplot/cimplot.h,cimplot/generator/output/definitions.json,cimplot/generator/output/structs_and_enums.json,-r cimgui/generator/output/structs_and_enums.json)

compile_cimgui_macos:
	rm -rf ./cimgui/build
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/arm64/

## generate: generates both bindings (equal to `all`)
.PHONY: generate
generate: setup cimgui cimplot
