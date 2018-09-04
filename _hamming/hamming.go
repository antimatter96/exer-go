package hamming

import (
	"errors"
)

// Distance returns hamming distance between two strands of DNA
// Returns error if strands are not equal in length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Lengths not equal")
	}
	count := 0
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			count++
		}
	}
	return count, nil
}
