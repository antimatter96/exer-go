package atbash

import (
	"strings"
	"unicode"
)

// Atbash applies the atbash cipher to the given plaintext
func Atbash(plaintext string) string {
	var cipher strings.Builder

	var count int

	for _, r := range plaintext {
		if unicode.IsNumber(r) {
			count++
			cipher.WriteRune(r)
			continue
		}
		if !unicode.IsLetter(r) {
			continue
		}
		if count > 4 && count%5 == 0 {
			cipher.WriteRune(rune(' '))
		}
		count++
		r = unicode.ToLower(r)
		cipher.WriteRune(rune('a') + rune('z') - r)
	}

	return cipher.String()
}
