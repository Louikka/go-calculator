package lib

import (
	"encoding/json"
	"gocalc/interpreter"
	"os"
	"path/filepath"
)

func WriteToFile(filename string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filename), 0750)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CompileAndWriteASTToFile(source string, filename string) error {
	ast, err := interpreter.CompileToAST(source)
	if err != nil {
		return err
	}

	ast_s, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		return err
	}

	err = WriteToFile(filename, ast_s)
	if err != nil {
		return err
	}

	return nil
}
