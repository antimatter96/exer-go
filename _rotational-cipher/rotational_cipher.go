package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher is used to implement a rot cipher with given key
func RotationalCipher(plainText string, key int) string {
	var strB strings.Builder
	var temp rune
	for _, r := range plainText {
		if !unicode.IsLetter(r) {
			strB.WriteRune(r)
			continue
		}
		temp = r + rune(key)
		if unicode.IsLower(r) && temp > 'z' {
			strB.WriteRune(temp - 26)
		} else if unicode.IsUpper(r) && temp > 'Z' {
			strB.WriteRune(temp - 26)
		} else {
			strB.WriteRune(temp)
		}
	}
	return strB.String()
}
