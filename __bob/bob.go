// Package bob implements Bob, a teengager with limited responses
package bob

import (
	"strings"
	"unicode"
)

// Hey returns Bob's response to the provided string
func Hey(remark string) string {
	trimmedString := strings.TrimFunc(remark, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	length := len(trimmedString)

	if length == 0 {
		return "Fine. Be that way!"
	}

	allUpper := true

	hasNumbers := false
	hasLetters := false
	for index := 0; (!hasLetters || !hasNumbers || allUpper) && index < length; index++ {

	}

	for _, v := range trimmedString {
		if unicode.IsNumber(v) {
			hasNumbers = true
		} else if unicode.IsLetter(v) {
			hasLetters = true
			if unicode.IsLower(v) {
				allUpper = false
			}
		}
		if hasLetters && hasNumbers && !allUpper {
			break
		}
	}

	if trimmedString[length-1] == '?' {
		if !allUpper || (allUpper && !hasLetters) {
			return "Sure."
		} else {
			return "Calm down, I know what I'm doing!"
		}
	}

	if !allUpper {
		return "Whatever."
	}

	if allUpper && hasLetters {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
