package file

import (
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func CheckJsonExtension(name string) bool {
	return strings.HasSuffix(name, ".json")
}
