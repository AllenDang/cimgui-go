./cmd/codegen/build/codegen: clean
	go build -o ./cmd/codegen/build/codegen ./cmd/codegen

.PHONY: test_codegen
test_codegen: ./cmd/codegen/build/codegen
	cd ./cmd/codegen/build; ./codegen -d ../../../cimgui/generator/output/definitions.json -e ../../../cimgui/generator/output/structs_and_enums.json

.PHONY: clean
clean:
	rm -rf ./cmd/codegen/build/
