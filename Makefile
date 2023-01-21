NAME=cimgui-go Code Generator

.PHONY: all

## all: generates both bindings.
all: generate

# Parameters:
# $1: prefix
# $2: include path (of header file)
# $3: definitions.json filepath
# $4: structs_and_enums.json filepath
# $5: additional agruments to codegen call (e.g. -r option)
define generate
	@echo "Generating for $(1)"
	go run github.com/AllenDang/cimgui-go/cmd/codegen -p $(1) -i $(2) -d $(3) -e $(4) $(5)
	go run mvdan.cc/gofumpt@latest -w $(1)_enums.go
	go run mvdan.cc/gofumpt@latest -w $(1)_structs.go
	go run mvdan.cc/gofumpt@latest -w $(1)_funcs.go
endef

define cimgui
	$(call generate,cimgui,cimgui/cimgui.h,cimgui/generator/output/definitions.json,cimgui/generator/output/structs_and_enums.json)
endef

## cimgui: generate cimgui binding
.PHONY: cimgui
cimgui:
	$(call cimgui)

define cimplot
	$(call generate,cimplot,cimplot/cimplot.h,cimplot/generator/output/definitions.json,cimplot/generator/output/structs_and_enums.json,-r cimgui/generator/output/structs_and_enums.json)
endef

## cimplot: generate implot binding
.PHONY: cimplot
cimplot:
	$(call cimplot)

compile_cimgui_macos:
	rm -rf ./cimgui/build
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/arm64/

## generate: generates both bindings (equal to `all`)
.PHONY: generate
generate: cimgui cimplot

# update updates sub-repos (like cimplot or cimgui)
# $1 - subrepo directory
# $2 - repository URL
# $3 - $1/<c++ repo>/
define update
	@echo "updating $1 from $2"
	rm -rf $1
	git clone --recurse-submodules $2 $1
	rm -rf $1/.git $1/$3/.git
	cd $1/generator; \
		sh generator.sh
	echo "// placeholder package used to include this code in vendor dir.\npackage doc" >> $1/doc.go
	cat $1/doc.go >> $1/$3/doc.go
endef

.PHONY: update
update:
	$(call update,cimgui,https://github.com/cimgui/cimgui,imgui)
	$(call cimgui)
	$(call update,cimplot,https://github.com/cimgui/cimplot,implot)
	$(call cimplot)
