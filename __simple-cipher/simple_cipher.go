package cipher

import (
	"strings"
	"unicode"
)

func getStringToEncode(in string) string {
	var out strings.Builder
	for _, r := range in {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			out.WriteRune(r)
		}
	}
	return out.String()
}

// type ceaserCipher struct{}

// func (c *ceaserCipher) Encode(plainText string) string {
// 	toEncode := getStringToEncode(plainText)
// 	var strB strings.Builder
// 	for _, r := range toEncode {
// 		if r+3 > 'z' {
// 			strB.WriteRune(r + 3 - 26)
// 		} else {
// 			strB.WriteRune(r + 3)
// 		}
// 	}
// 	return strB.String()
// }

// func (c *ceaserCipher) Decode(cipherText string) string {
// 	var strB strings.Builder
// 	for _, r := range cipherText {
// 		if r-3 < 'a' {
// 			strB.WriteRune(r - 3 + 26)
// 		} else {
// 			strB.WriteRune(r - 3)
// 		}
// 	}
// 	return strB.String()
// }

func NewCaesar() Cipher {
	return &shiftCipher{3}
}

type shiftCipher struct {
	shiftDistance int
}

func (c *shiftCipher) Encode(plainText string) string {
	toEncode := getStringToEncode(plainText)
	var strB strings.Builder
	var temp rune
	for _, r := range toEncode {
		temp = r + rune(c.shiftDistance)
		if temp > 'z' {
			strB.WriteRune(temp - 26)
		} else if temp < 'a' {
			strB.WriteRune(temp + 26)
		} else {
			strB.WriteRune(temp)
		}
	}
	return strB.String()
}

func (c *shiftCipher) Decode(cipherText string) string {
	var strB strings.Builder
	var temp rune
	for _, r := range cipherText {
		temp = r - rune(c.shiftDistance)
		if temp < 'a' {
			strB.WriteRune(temp + 26)
		} else if temp > 'z' {
			strB.WriteRune(temp - 26)
		} else {
			strB.WriteRune(temp)
		}
	}
	return strB.String()
}

// NewShift returns a shift cipher with shift distance allowed between [1,25] and [-25,-1]
func NewShift(distance int) Cipher {
	if !((distance > 0 && distance < 26) || (distance < 0 && distance > -26)) {
		return nil
	}
	return &shiftCipher{distance}
}

type vigenereCipher struct {
	key string
}

func (c *vigenereCipher) Encode(plainText string) string {
	var keyMod int = len(c.key)
	var keyIndex int
	toEncode := getStringToEncode(plainText)
	var strB strings.Builder
	var temp rune
	for _, r := range toEncode {
		shift := c.key[keyIndex] - 'a'
		temp = r + rune(shift)
		if temp > 'z' {
			strB.WriteRune(temp - 26)
		} else {
			strB.WriteRune(temp)
		}
		keyIndex++
		keyIndex %= keyMod
	}
	return strB.String()
}

func (c *vigenereCipher) Decode(cipherText string) string {
	var keyMod int = len(c.key)
	var keyIndex int
	var strB strings.Builder
	var temp rune
	for _, r := range cipherText {
		shift := c.key[keyIndex] - 'a'
		temp = r - rune(shift)
		if temp < 'a' {
			strB.WriteRune(temp + 26)
		} else {
			strB.WriteRune(temp)
		}
		keyIndex++
		keyIndex %= keyMod
	}
	return strB.String()
}

// NewVigenere returns a vigerene cipher. The key is lower letters, can't all be 'a's
func NewVigenere(key string) Cipher {
	var notA bool
	for _, r := range key {
		if !unicode.IsLetter(r) || !unicode.IsLower(r) {
			return nil
		}
		if r != 'a' {
			notA = true
		}
	}
	if !notA {
		return nil
	}
	return &vigenereCipher{key}
}
