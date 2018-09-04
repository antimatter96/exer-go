package pangram

import (
	"unicode"
)

// IsPangram checks if the string is a pangram ie contains all letters
func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}

	var count int = 0
	present := make(map[rune]bool)

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if _, ok := present[r]; ok {
			continue
		} else {
			present[r] = true
			count++
		}
		if count == 26 {
			return true
		}
	}

	return false
}
