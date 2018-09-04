package luhn

import (
	"strings"
	"unicode"
)

// Valid returns whether or not given number is valid per the Luhn formula.
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) < 2 {
		return false
	}

	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}

	var sum int
	var temp int

	target := len(s) % 2
	for i := len(s) - 1; i > -1; i-- {
		temp = int(s[i]) - 48
		if i%2 == target {
			temp *= 2
			if temp > 9 {
				temp -= 9
			}
		}
		sum += temp
	}

	return (sum % 10) == 0
}
