package armstrong

import (
	"math"
	"strconv"
)

// IsNumber determines if a number is an amstrong number or not
func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	if len(s) == 1 {
		return true
	}
	var tempSum float64
	var length = float64(len(s))
	for _, v := range s {
		tempSum += math.Pow(float64(v-48), length)
	}
	if int(tempSum) == n {
		return true
	}
	return false
}
