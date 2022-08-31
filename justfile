_build_cmd_codegen:
	rm -rf ./cmd/codegen/build
	go build -o ./cmd/codegen/build/codegen ./cmd/codegen

_gencode prefix include_path d_file e_file: _build_cmd_codegen
	cd ./cmd/codegen/build; ./codegen -p {{prefix}} -i {{include_path}} -d {{d_file}} -e {{e_file}}
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

gencode_cimgui:
	just _gencode cimgui cimgui/cimgui.h ../../../cimgui/generator/output/definitions.json ../../../cimgui/generator/output/structs_and_enums.json

gencode_cimplot:
	just _gencode cimplot cimplot/cimplot.h ../../../cimplot/generator/output/definitions.json ../../../cimplot/generator/output/structs_and_enums.json
