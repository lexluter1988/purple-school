package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func NewBin(id string, private bool, name string) *Bin {
	createdAt := time.Now()
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

func NewBinList() *BinList {
	return &BinList{}
}

func AddBin(binlist *BinList, bin *Bin) {
	binlist.bins = append(binlist.bins, *bin)
}

func main() {
	fmt.Println("Hello Purple School")
	bins := NewBinList()
	bin1 := NewBin("1", false, "one")
	bin2 := NewBin("2", false, "two")
	bin3 := NewBin("3", false, "three")
	fmt.Println(bins)
	AddBin(bins, bin1)
	fmt.Println(bins)
	AddBin(bins, bin2)
	fmt.Println(bins)
	AddBin(bins, bin3)
	fmt.Println(bins)

}
