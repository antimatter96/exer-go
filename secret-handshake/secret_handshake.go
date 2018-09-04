package secret

var constants = map[int]string{
	1: "wink",
	2: "double blink",
	3: "close your eyes",
	4: "jump",
}

// Handshake takes a code and returns an array of strings
// Current max value of code is 31
func Handshake(code uint) []string {
	toRet := make([]string, 0)
	var power int = 1
	for code != 0 {
		lastDigit := code % 2
		if lastDigit == 1 {
			if power == 5 {
				for i, j := 0, len(toRet)-1; i < j; i, j = i+1, j-1 {
					toRet[i], toRet[j] = toRet[j], toRet[i]
				}
			} else {
				toRet = append(toRet, constants[power])
			}
		}
		code /= 2
		power++
	}

	return toRet
}
