package isbn

import (
	"strings"
)

// IsValidISBN is used to check if given string is a valid ISBN number
func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) < 10 {
		return false
	}
	var sum int
	for i := 0; i < len(isbn)-1; i++ {
		sum += (10 - i) * (int(isbn[i]) - 48)
	}
	if isbn[9] == 'X' {
		sum += 10
	} else {
		sum += (int(isbn[9]) - 48)
	}
	return sum%11 == 0
}
