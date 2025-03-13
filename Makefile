NAME=cimgui-go Code Generator

.PHONY: all

## all: generates both bindings.
all: generate

## setup: prepare some dependencies
.PHONY: setup
setup:
	cd cmd/codegen; \
		go build -o ../../codegen .

# Parameters:
# $1: go package name
# $2: include path (of header file)
# $3: definitions.json filepath
# $4: structs_and_enums.json filepath
# $5: typedefs_dict.json filepath
# $6: additional agruments to codegen call (e.g. -r option)
define generate
	@echo "Generating for $(1)"
	mkdir -p $(1)
	cat templates/cflags.go |sed -e "s/^package.*/package $(1)/g" > $(1)/cflags.go
	cd $(1); \
		../codegen --preset ../cmd/codegen/cimgui-go-preset.json --package $(1) --include ../$(2) --definitions ../$(3) --enums ../$(4) --typedefs ../$(5) $(6)
endef

define imgui
	$(call generate,imgui,cwrappers/cimgui.h,cwrappers/cimgui_templates/definitions.json,cwrappers/cimgui_templates/structs_and_enums.json,cwrappers/cimgui_templates/typedefs_dict.json)
endef

## cimgui: generate imgui binding
.PHONY: imgui
imgui: setup
	$(call imgui)

define implot
	$(call generate,implot,cwrappers/cimplot.h,cwrappers/cimplot_templates/definitions.json,cwrappers/cimplot_templates/structs_and_enums.json,cwrappers/cimplot_templates/typedefs_dict.json,--ref-enums ../cwrappers/cimgui_templates/structs_and_enums.json --ref-typedefs ../cwrappers/cimgui_templates/typedefs_dict.json)
endef

## implot: generate implot binding
.PHONY: implot
implot: setup
	$(call implot)

define imnodes
	$(call generate,imnodes,cwrappers/cimnodes.h,cwrappers/cimnodes_templates/definitions.json,cwrappers/cimnodes_templates/structs_and_enums.json,cwrappers/cimnodes_templates/typedefs_dict.json,--ref-enums ../cwrappers/cimgui_templates/structs_and_enums.json --ref-typedefs ../cwrappers/cimgui_templates/typedefs_dict.json)
endef

## imnodes: generate imnodes binding
.PHONY: imnodes
imnodes: setup
	$(call imnodes)

define immarkdown
	$(call generate,immarkdown,cwrappers/cimmarkdown.h,cwrappers/cimmarkdown_templates/definitions.json,cwrappers/cimmarkdown_templates/structs_and_enums.json,cwrappers/cimmarkdown_templates/typedefs_dict.json,--ref-enums ../cwrappers/cimgui_templates/structs_and_enums.json -rt ../cwrappers/cimgui_templates/typedefs_dict.json)
endef

## immarkdown: generate immarkdown binding
.PHONY: immarkdown
immarkdown: setup
	$(call immarkdown)

define imguizmo
	$(call generate,imguizmo,cwrappers/cimguizmo.h,cwrappers/cimguizmo_templates/definitions.json,cwrappers/cimguizmo_templates/structs_and_enums.json,cwrappers/cimguizmo_templates/typedefs_dict.json,-re ../cwrappers/cimgui_templates/structs_and_enums.json -rt ../cwrappers/cimgui_templates/typedefs_dict.json)
endef

## imguizmo: generate imguizmo binding
.PHONY: imguizmo
imguizmo: setup
	$(call imguizmo)

define imcte
	$(call generate,ImGuiColorTextEdit,cwrappers/cimCTE.h,cwrappers/cimCTE_templates/definitions.json,cwrappers/cimCTE_templates/structs_and_enums.json,cwrappers/cimCTE_templates/typedefs_dict.json,-re ../cwrappers/cimgui_templates/structs_and_enums.json -rt ../cwrappers/cimgui_templates/typedefs_dict.json)
endef

## imcte: generate ImGuiColorTextEdit binding
.PHONY: imcte
imcte: setup
	$(call imcte)

## generate: generates both bindings (equal to `all`)
.PHONY: generate
generate: imgui implot imnodes immarkdown imguizmo imcte

# update updates sub-repos (like cimplot or cimgui)
# $1 - subrepo directory
# $2 - repository URL
# $3 - $1/<c++ repo>/
# $4 - branch in $3 (cd tmp/$1/$3 && git checkout $4)
# $5 - additional generator options
define update
	@echo "updating $1 from $2"
	mkdir -p tmp/
	if test -e tmp/$1; then \
		rm -rf tmp/*; \
	fi
	git clone --recurse-submodules $2 tmp/$1
	cd tmp/$1/$3; \
		git checkout $4
	cd tmp/$1/generator; \
		bash generator.sh --target "internal noimstrv comments" $5
	cp -f tmp/$1/$1* cwrappers/
	if test -e tmp/$1/generator/output/$1*; then \
		cp -f tmp/$1/generator/output/$1* cwrappers/; \
	fi
	mkdir cwrappers/$1_templates
	cp -f tmp/$1/generator/output/*json cwrappers/$1_templates
	mkdir -p cwrappers/$3
	cp -rf tmp/$1/$3/* cwrappers/$3
	cd tmp/$1; \
		echo "$1 ($2) HEAD is on: `git rev-parse HEAD`" >> ../../cwrappers/VERSION.txt
	cd tmp/$1/$3; \
		echo "$1/$3 HEAD is on: `git rev-parse HEAD`" >> ../../../cwrappers/VERSION.txt
endef

.PHONY: update
update: setup
	rm -rf cwrappers/*
	$(call update,cimgui,https://github.com/cimgui/cimgui,imgui,docking, --cflags "glfw opengl3 opengl2 sdl2 -DIMGUI_USE_WCHAR32")
	cat templates/assert.h >> cwrappers/imgui/imconfig.h
	$(call imgui)
	$(call update,cimplot,https://github.com/cimgui/cimplot,implot,master)
	$(call implot)
	$(call update,cimnodes,https://github.com/cimgui/cimnodes,imnodes,master)
	$(call imnodes)
	$(call update,cimmarkdown,https://github.com/gucio321/cimmarkdown,imgui_markdown,main)
	$(call immarkdown)
	$(call update,cimguizmo,https://github.com/cimgui/cimguizmo,ImGuizmo,master)
	$(call imguizmo)
	$(call update,cimCTE,https://github.com/cimgui/cimcte,ImGuiColorTextEdit,master)
	$(call imcte)
	$(call vendor-eliminate)
	$(call dummy)

# dummy creates dummy.go files to baypass GO vendor policy that excludes everything that has no .go files (including our C source).
define dummy
	echo -e "//go:build required\n// +build required\n\npackage imgui\n\nimport (\n" > dummy.go
	for i in `find cwrappers -type f \( -name "*.h" -o -name "*.cpp" \) -exec dirname {} \; | sort -u`; do \
		cp templates/dummy.go.template $$i/dummy.go; \
		echo -e "\t_ \"github.com/AllenDang/cimgui-go/$$i\"" >> dummy.go; \
		done
	echo ")" >> dummy.go
endef

define vendor-eliminate
	for i in `find . -name vendor -type d`; do \
		mv $$i $${i}1; \
		done
endef

.PHONY: dummy
dummy:
	$(call dummy)
