package bins

import "time"

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
