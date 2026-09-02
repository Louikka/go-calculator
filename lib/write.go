package lib

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Pretty-formats JSON or JSON-like value and writes it to the file.
func WriteJSONToFile(v any, file string) error {
	v_asBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(file), 0750)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, v_asBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
