package file

import (
	"os"
	"strings"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (f *File) ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *File) CheckJsonExtension(name string) bool {
	return strings.HasSuffix(name, ".json")
}
