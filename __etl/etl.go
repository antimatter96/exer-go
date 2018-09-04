package etl

import (
	"strings"
)

func Transform(current map[int][]string) map[string]int {
	mp := make(map[string]int)
	for k, v := range current {
		for _, digit := range v {
			mp[strings.ToLower(digit)] = k
		}
	}
	return mp
}
