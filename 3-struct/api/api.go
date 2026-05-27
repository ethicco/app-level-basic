package api

import (
	"fmt"
	"path/filepath"

	"app-struct/bins"
	"app-struct/config"
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
	config    *config.Config
}

func NewApp(storage Storage, validator FileValidator, cfg *config.Config) *App {
	return &App{storage: storage, validator: validator, config: cfg}
}

func (a *App) CreateBin(id string, private bool, name string, path string) error {
	if !a.validator.CheckJsonExtension(path) {
		return fmt.Errorf("file must have .json extension: %s", path)
	}
	bin := bins.NewBin(id, private, name)
	return a.storage.SaveBin(filepath.Join(a.config.StoragePath, path), bin)
}

func (a *App) ListBins(path string) ([]bins.BinList, error) {
	if !a.validator.CheckJsonExtension(path) {
		return nil, fmt.Errorf("file must have .json extension: %s", path)
	}
	return a.storage.LoadBinList(filepath.Join(a.config.StoragePath, path))
}
