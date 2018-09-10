package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is used represent the frequency of words in sentences
type Frequency map[string]int

// WordCount returns the frequency of words in a string
func WordCount(phrase string) Frequency {

	phrase = strings.ToLower(phrase)
	reg1 := regexp.MustCompile("'([a-z0-9]+)'")
	phrase = reg1.ReplaceAllString(phrase, "$1")

	reg := regexp.MustCompile("[^a-z0-9']+")
	words := reg.Split(phrase, -1)

	freq := make(Frequency)
	for _, word := range words {
		if len(word) > 0 {
			freq[word]++
		}
	}
	return freq
}
