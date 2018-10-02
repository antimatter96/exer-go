package romannumerals

func ToRomanNumeral(n int) (string, bool) {
	if n < 1 || n > 3000 {
		return "", true
	}
	var str string

	for ; n > 999; n -= 1000 {
		str += "M"
	}

	if n > 900 {
		str += "CM"
		n -= 900
	}

	if n > 500 {

	}

	if n > 0 {
		remaining, _ := ToRomanNumeral(n)
		str += remaining
	}

	return str, false
}
