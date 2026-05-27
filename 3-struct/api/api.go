package api

import (
	"fmt"

	"app-struct/bins"
)

type Storage interface {
	SaveBin(path string, bin bins.Bin) error
	LoadBinList(path string) ([]bins.BinList, error)
}

type FileValidator interface {
	CheckJsonExtension(name string) bool
}

type App struct {
	storage   Storage
	validator FileValidator
}

func NewApp(storage Storage, validator FileValidator) *App {
	return &App{storage: storage, validator: validator}
}

func (a *App) CreateBin(id string, private bool, name string, path string) error {
	if !a.validator.CheckJsonExtension(path) {
		return fmt.Errorf("file must have .json extension: %s", path)
	}
	bin := bins.NewBin(id, private, name)
	return a.storage.SaveBin(path, bin)
}

func (a *App) ListBins(path string) ([]bins.BinList, error) {
	if !a.validator.CheckJsonExtension(path) {
		return nil, fmt.Errorf("file must have .json extension: %s", path)
	}
	return a.storage.LoadBinList(path)
}
