package storage

import (
	"encoding/json"
	"os"

	"app-struct/bins"
)

func SaveBin(path string, bin bins.Bin) error {
	data, err := json.Marshal(bin)

	if err != nil {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	_, err = file.Write(data)
	file.Close()

	return err
}

func LoadBinList(path string) ([]bins.BinList, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var list []bins.BinList

	err = json.Unmarshal(data, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}
