package anagram

import (
	"reflect"
	"unicode"
)

// Detect is used to determine which words from a given list are anagarams to a given word
func Detect(subject string, candidates []string) []string {
	var valid []string

	subjectMap := make(map[rune]int)
	var subjectHasLower bool
	var subjectHasUpper bool
	for _, v := range subject {
		if unicode.IsUpper(v) {
			v = unicode.ToLower(v)
			subjectHasUpper = true
		} else {
			subjectHasLower = true
		}
		subjectMap[v]++
	}

	for _, candidate := range candidates {
		candidateMap := make(map[rune]int)
		var candidateHasLower bool
		var candidateHasUpper bool
		for _, v := range candidate {
			if unicode.IsUpper(v) {
				v = unicode.ToLower(v)
				candidateHasUpper = true
			} else {
				candidateHasLower = true
			}
			candidateMap[v]++
		}
		if reflect.DeepEqual(subjectMap, candidateMap) {
			if (subjectHasUpper == candidateHasUpper) && (subjectHasLower == candidateHasLower) {
				valid = append(valid, candidate)
			} else if subjectHasLower {
				valid = append(valid, candidate)
			}
		}
	}

	return valid
}
