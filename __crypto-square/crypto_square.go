package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var regex *regexp.Regexp = regexp.MustCompile(`[^a-z0-9]`)

func getRC(length int) (int, int) {
	sqrt := int(math.Sqrt(float64(length)))
	if sqrt*sqrt == length {
		return sqrt, sqrt
	}
	if (sqrt * (sqrt + 1)) >= length {
		return sqrt, sqrt + 1
	}
	return sqrt + 1, sqrt + 1
}

func Encode(pt string) string {

	pt = strings.ToLower(pt)
	pt = regex.ReplaceAllString(pt, "")

	r, c := getRC(len(pt))
	if r*c > len(pt) {
		pt += strings.Repeat(" ", (r*c)-len(pt))
	}

	var newLen int = len(pt)

	var toRet strings.Builder
	for i := 0; i < c; i++ {
		var written int = 0
		for j := i; j < newLen && written < r; j, written = j+c, written+1 {
			toRet.WriteByte(pt[j])
		}
		if i < c-1 {
			toRet.WriteByte(' ')
		}
	}

	return toRet.String()
}
