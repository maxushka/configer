package configer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Save[T any](s T, path string) error {
	var fileContent []byte
	var err error

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		fileContent, err = yaml.Marshal(s)
	case ".json":
		fileContent, err = json.Marshal(s)
	}

	if err != nil {
		return err
	}
	err = os.WriteFile(path, fileContent, 0755)

	return err
}

func Load[T any](path string) (ret T, err error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return
	}

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(fileContent, &ret)
	case ".json":
		err = json.Unmarshal(fileContent, &ret)
	}

	return
}
