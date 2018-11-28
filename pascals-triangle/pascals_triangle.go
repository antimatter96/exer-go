package pascal

import "math/big"

func row(n int) []int {
	var x []int
	if n < 0 {
		return x
	}

	if n == 0 {
		return []int{1}
	}

	x = make([]int, n+1)
	x[0] = 1
	if n > -1 {
		x[n] = 1
	}
	for i := 1; i <= n/2+1; i++ {
		x[i] = compute(n, i)
	}
	for i := 1; i < n/2+1; i++ {
		x[n-i] = x[i]
	}
	return x
}

// Triangle creates a pascals triangle upto the nth level
func Triangle(n int) [][]int {
	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = row(i)
	}
	return x
}

var computeCache = map[int]map[int]int{1: map[int]int{1: 1}}

func compute(m, n int) int {
	if m == n {
		return 1
	}
	if n == 1 {
		return m
	}

	if _, presentMap := computeCache[m]; !presentMap {
		computeCache[m] = make(map[int]int)
	}

	value, present := computeCache[m][n]
	if present {
		return value
	}

	x := new(big.Int).Binomial(int64(m), int64(n))
	toRet := int(x.Int64())
	computeCache[m][n] = toRet
	return toRet
}
