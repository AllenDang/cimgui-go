_build_cmd_codegen:
	rm -rf ./cmd/codegen/build
	go build -o ./cmd/codegen/build/codegen ./cmd/codegen

_build_cmd_rename:
	rm -rf ./cmd/codegen/rename
	go build -o ./cmd/codegen/build/rename ./cmd/rename

_handwritten: 
	cp -f ./handwritten/*.cpp .
	cp -f ./handwritten/*.h .
	cp -f ./handwritten/*.go .

_gencode prefix include_path d_file e_file *r_file: _build_cmd_codegen
	cd ./cmd/codegen/build; ./codegen -p {{prefix}} -i {{include_path}} -d {{d_file}} -e {{e_file}} {{r_file}}
	gofmt -w ./cmd/codegen/build/{{prefix}}_enums.go
	gofmt -w ./cmd/codegen/build/{{prefix}}_structs.go
	gofmt -w ./cmd/codegen/build/{{prefix}}_funcs.go

	cp -f ./cmd/codegen/build/{{prefix}}_wrapper.h ./
	cp -f ./cmd/codegen/build/{{prefix}}_wrapper.cpp ./
	cp -f ./cmd/codegen/build/{{prefix}}_structs_accessor.h ./
	cp -f ./cmd/codegen/build/{{prefix}}_structs_accessor.cpp ./
	cp -f ./cmd/codegen/build/{{prefix}}_enums.go ./
	cp -f ./cmd/codegen/build/{{prefix}}_funcs.go ./
	cp -f ./cmd/codegen/build/{{prefix}}_structs.go ./

	rm ./cmd/codegen/build/{{prefix}}_wrapper.h
	rm ./cmd/codegen/build/{{prefix}}_wrapper.cpp
	rm ./cmd/codegen/build/{{prefix}}_structs_accessor.h
	rm ./cmd/codegen/build/{{prefix}}_structs_accessor.cpp
	rm ./cmd/codegen/build/{{prefix}}_enums.go
	rm ./cmd/codegen/build/{{prefix}}_funcs.go
	rm ./cmd/codegen/build/{{prefix}}_structs.go

_idiomatic: _build_cmd_rename
	rm -rf ./idiomatic
	mkdir ./idiomatic
	mkdir ./idiomatic/cimgui
	mkdir ./idiomatic/cimplot
	mkdir ./idiomatic/thirdparty
	mkdir ./idiomatic/thirdparty/glfw
	cp ./*.go ./idiomatic
	cp ./*.cpp ./idiomatic
	cp ./*.h ./idiomatic
	cp ./cimgui/*.cpp ./idiomatic/cimgui
	cp ./cimgui/*.h ./idiomatic/cimgui
	cp ./cimplot/*.cpp ./idiomatic/cimplot
	cp ./cimplot/*.h ./idiomatic/cimplot
	cp -r ./thirdparty/glfw/* ./idiomatic/thirdparty/glfw
	./cmd/codegen/build/rename ./idiomatic/*.go

gencode_cimgui args="":
	just _gencode cimgui cimgui/cimgui.h ../../../cimgui/generator/output/definitions.json ../../../cimgui/generator/output/structs_and_enums.json
	just _handwritten
	just _idiomatic

gencode_cimplot args="":
	just _gencode cimplot cimplot/cimplot.h ../../../cimplot/generator/output/definitions.json ../../../cimplot/generator/output/structs_and_enums.json -r ../../../cimgui/generator/output/structs_and_enums.json
	just _handwritten

compile_cimgui_macos:
	rm -rf ./cimgui/build
	cd ./cimgui; cmake -Bbuild -DCMAKE_BUILD_TYPE=Release -DIMGUI_STATIC=On -DCMAKE_OSX_ARCHITECTURES=arm64
	cd ./cimgui/build; make
	cp -f ./cimgui/build/cimgui.a ./lib/macos/arm64/
