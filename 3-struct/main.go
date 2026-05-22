package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func NewBinList(id string, private bool, name string) BinList {
	return BinList{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func main() {}
