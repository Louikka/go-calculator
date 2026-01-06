OUTPUT_DIR := bin



run :
	go run .

build :
	go build -o $(OUTPUT_DIR)/



.PHONY : run build
