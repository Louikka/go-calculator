OUTPUT_DIR := bin



test :
	go run .

build :
	go build -o $(OUTPUT_DIR)/



.PHONY : test build
