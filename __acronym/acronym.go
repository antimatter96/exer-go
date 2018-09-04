// Package acronym is used to get acronyms
package acronym

import "strings"

// Abbreviate returns an acronym of the given string.
func Abbreviate(s string) string {
	if len(s) == 0 {
		return ""
	}
	var acronym string
	acronym += string(s[0])
	for index := 1; index < len(s)-1; index++ {
		if s[index] == ' ' || s[index] == '-' {
			acronym += string(s[index+1])
		}
	}
	return strings.ToUpper(acronym)
}
