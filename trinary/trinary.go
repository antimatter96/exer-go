package trinary

import (
	"fmt"
)

func ParseTrinary(s string) (int64, error) {
	var pow int64 = 1
	var n int64 = 0
	for i := len(s) - 1; i > -1; i-- {
		n += int64(s[i]-48) * pow
		pow *= 3
	}
	if n < 0 {
		return 0, fmt.Errorf("Overflow")
	}
	return n, nil
}
