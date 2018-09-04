package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(s string) string {
	if len(s) == 0 {
		return s
	}
	var toRet strings.Builder
	var count int = 1
	var lastWas byte = s[0]
	for i := 1; i < len(s); i++ {
		if lastWas == s[i] {
			count++
		} else {
			if count != 1 {
				toRet.WriteString(strconv.Itoa(count))
			}
			toRet.WriteByte(lastWas)
			lastWas = s[i]
			count = 1
		}
	}
	if count != 1 {
		toRet.WriteString(strconv.Itoa(count))
	}
	toRet.WriteByte(lastWas)
	return toRet.String()
}

func RunLengthDecode(s string) string {
	if len(s) == 0 {
		return s
	}
	var toRet strings.Builder
	var count = 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count = (count * 10) + (int(r) - 48)
			continue
		} else {
			if count < 1 {
				count = 1
			}
			for i := 0; i < count; i++ {
				toRet.WriteRune(r)
			}
			count = 0
		}
	}
	return toRet.String()
}
