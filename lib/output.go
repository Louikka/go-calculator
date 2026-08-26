package lib

import (
	"encoding/json"
	"gocalc/interpreter"
	"os"
	"path/filepath"
)

func CompileAndWriteASTToFile(source string, file string) error {
	ast, err := interpreter.CompileToAST(source)
	if err != nil {
		return err
	}

	ast_s, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(file), 0750)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, ast_s, 0644)
	if err != nil {
		return err
	}

	return nil
}
