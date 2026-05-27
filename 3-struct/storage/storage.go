package storage

import (
	"encoding/json"
	"os"

	"app-struct/bins"
)

type FileReader interface {
	ReadFile(name string) ([]byte, error)
}

type Storage struct {
	reader FileReader
}

func NewStorage(reader FileReader) *Storage {
	return &Storage{reader: reader}
}

func (s *Storage) SaveBin(path string, key string, bin bins.Bin) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	f.Close()

	return err
}

func (s *Storage) LoadBinList(path string, key string) ([]bins.BinList, error) {
	data, err := s.reader.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var list []bins.BinList
	if err = json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	return list, nil
}
