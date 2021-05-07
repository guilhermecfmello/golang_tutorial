package main

import "fmt"

// Both are boolean		Returns 3 boolean values
func purchases(work1, work2 bool) (bool, bool, bool) {
	buyTv50 := work1 && work2
	buyTv32 := work1 != work2 // XOR
	buyIcecream := work1 || work2

	return buyTv50, buyTv32, buyIcecream
}

func main() {
	tv50, tv32, icecream := purchases(true, false)

	fmt.Printf("TV50: %s, TV32: %s, Icecream: %s, Healthy: %s", tv50, tv32, icecream, !icecream)
}
