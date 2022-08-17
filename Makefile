.PHONY: clean_gencode
clean_gencode:
	rm -rf ./cmd/codegen/build/

./cmd/codegen/build/codegen: clean_gencode
	go build -o ./cmd/codegen/build/codegen ./cmd/codegen

.PHONY: gencode
gencode: ./cmd/codegen/build/codegen
	cd ./cmd/codegen/build; ./codegen -d ../../../cimgui/generator/output/definitions.json -e ../../../cimgui/generator/output/structs_and_enums.json
	cp -f ./cmd/codegen/build/cimgui_wrapper.cpp ./
	cp -f ./cmd/codegen/build/cimgui_wrapper.h ./
	cp -f ./cmd/codegen/build/cimgui_structs_accessor.h ./
	cp -f ./cmd/codegen/build/enums.go ./
	cp -f ./cmd/codegen/build/funcs.go ./
	cp -f ./cmd/codegen/build/structs.go ./


.PHONY: gen_cimgui
gen_cimgui:
	cd ./cimgui/generator; ./generator.sh

.PHONY: clean_cimgui
clean_cimgui:
	rm -rf ./cimgui/build

.PHONY: compile_cimgui_macos_x86_64
compile_cimgui_macos_x86_64: clean_cimgui
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=x86_64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/x64/

.PHONY: compile_cimgui_macos_arm64
compile_cimgui_macos_arm64: clean_cimgui
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/arm64/

.PHONY: compile_cimgui_macos
compile_cimgui_macos: compile_cimgui_macos_x86_64 compile_cimgui_macos_arm64
	echo "Compile macos x86_64 and arm64"

.PHONY: clean_glfw
clean_glfw:
	rm -rf ./thirdparty/glfw/build

.PHONY: compile_glfw_macos_x86_64
compile_glfw_macos_x86_64: clean_glfw
	cd ./thirdparty/glfw; cmake -S . -Bbuild -DCMAKE_OSX_ARCHITECTURES=x86_64
	cd ./thirdparty/glfw/build; make
	cp -f ./thirdparty/glfw/build/src/libglfw3.a ./lib/macos/x64/

.PHONY: compile_glfw_macos_arm64
compile_glfw_macos_arm64: clean_glfw
	cd ./thirdparty/glfw; cmake -S . -Bbuild -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./thirdparty/glfw/build; make
	cp -f ./thirdparty/glfw/build/src/libglfw3.a ./lib/macos/arm64/
