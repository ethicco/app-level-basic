package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"app-struct/api"
	"app-struct/bins"
	"app-struct/config"
	"app-struct/file"
	"app-struct/storage"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		panic("Не найден .env файл")
	}

	fileService := file.NewFile()
	store := storage.NewStorage(fileService)
	app := api.NewApp(store, fileService, cfg)

	if err := app.CreateBin("bin-1", false, "My First Bin", "bin-1.json"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("bin saved: bin-1.json")

	list := []bins.BinList{
		bins.NewBinList("bin-1", false, "My First Bin"),
		bins.NewBinList("bin-2", true, "Private Bin"),
	}
	data, err := json.Marshal(list)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile("bins.json", data, 0644); err != nil {
		log.Fatal(err)
	}

	loaded, err := app.ListBins("bins.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("loaded bins:")
	for _, b := range loaded {
		fmt.Printf("  id=%s name=%s private=%v\n", b.ID, b.Name, b.Private)
	}
}
