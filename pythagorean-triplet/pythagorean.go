package main

import (
	"fmt"
)

var tt []int = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 53, 87}

type Triplet [3]int

func Range(min, max int) []Triplet {
	var t []Triplet
	return t
}
func Sum(p int) []Triplet {
	var t []Triplet
	return t
}

func main() {
	for i := 0; i < len(tt)-3; i++ {
		hi := tt[i]
		hi1 := tt[i+1]
		hi2 := tt[i+2]
		hi3 := tt[i+3]
		x := 2 * hi1 * hi2
		y := hi * hi3
		z := x + (hi * hi)
		fmt.Println(x, y, z)
	}
}
