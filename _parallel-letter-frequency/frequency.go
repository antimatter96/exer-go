package letter

type FreqMap map[rune]int

// Frequency is the already implemented function
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func tempFunc(st string, ch chan FreqMap) {
	ch <- Frequency(st)
}

// ConcurrentFrequency calls Frequency concurrently on an array of strings
// and combines their results
func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}

	var n = len(ss)
	ch := make(chan FreqMap)

	for i := 0; i < len(ss); i++ {
		go tempFunc(ss[i], ch)
	}

	for tempMap := range ch {
		n--
		for k, v := range tempMap {
			m[k] += v
		}
		if n == 0 {
			close(ch)
		}
	}

	return m
}
