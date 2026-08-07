package main

import (
	"fmt"
	"purple/dz-3/bins"
)

func main() {
	binslist := bins.NewBinList()
	bin1 := bins.NewBin("1", false, "one")
	bin2 := bins.NewBin("2", false, "two")
	bin3 := bins.NewBin("3", false, "three")
	fmt.Println(binslist)
	bins.AddBin(binslist, bin1)
	fmt.Println(binslist)
	bins.AddBin(binslist, bin2)
	fmt.Println(binslist)
	bins.AddBin(binslist, bin3)
	fmt.Println(binslist)

}
