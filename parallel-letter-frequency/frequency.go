package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}
	n := len(ss)
	ch := make(chan FreqMap)
	for _, s := range ss {
		go func(st string) {
			ch <- Frequency(st)
		}(s)
	}

	for true {
		switch expression {
		case condition:

		}
	}
	return m
}
