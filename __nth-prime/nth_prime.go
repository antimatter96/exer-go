package prime

import (
	"math"
	"math/big"
)

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}
	if n == 2 {
		return 3, true
	}
	var maxTill int64 = int64(math.MaxInt32)
	var i int64 = 5
	var tillNow int = 2
	var k *big.Int
	for ; i < maxTill; i += 2 {
		if i%3 == 0 {
			continue
		}
		k = big.NewInt(i)
		if k.ProbablyPrime(1) {
			tillNow++
			if tillNow == n {
				return int(i), true
			}
		}
	}
	return 0, true
}
