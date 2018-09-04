package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram tells wheteher a string is an isogram or not
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	present := make(map[rune]bool)

	for _, v := range s {
		if unicode.IsLetter(v) {
			if present[v] {
				return false
			}
			present[v] = true
		}
	}
	return true
}
