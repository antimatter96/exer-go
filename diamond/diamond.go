package diamond

import (
	"fmt"
	"strings"
)

// Gen creates the diamond for a particular character
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("Out of range Error")
	}

	n := int(char - 'A')
	if n == 0 {
		return "A", nil
	}

	maxL := (2 * n) + 1

	common := strings.Repeat(" ", maxL)
	srcc := make([][]byte, maxL)
	for index := 0; index < maxL; index++ {
		srcc[index] = []byte(common)
	}

	for index := 0; index <= maxL/2; index++ {
		srcc[index][(maxL/2)-index] = byte(int('A') + index)
		srcc[index][(maxL/2)+index] = byte(int('A') + index)
	}
	for index := maxL / 2; index < maxL; index++ {
		srcc[index] = srcc[maxL-1-index]
	}

	var strBuil strings.Builder
	for index := 0; index < maxL; index++ {
		strBuil.WriteString(string(srcc[index]))
		if index < maxL-1 {
			strBuil.WriteByte('\n')
		}
	}

	return strBuil.String(), nil
}
