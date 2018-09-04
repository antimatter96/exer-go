// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverbs := make([]string, len(rhyme))
	for index := 1; index < len(rhyme); index++ {
		proverbs[index-1] = fmt.Sprintf("For want of a %v the %v was lost.", rhyme[index-1], rhyme[index])
	}
	if len(rhyme) > 0 {
		proverbs[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])
	}
	return proverbs
}
