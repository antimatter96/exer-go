package letter

import (
	"sync"
)

type FreqMap map[rune]int

// Frequency is the already implemented function
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func tempFunc(st string, mutex *sync.Mutex, wg *sync.WaitGroup, m *FreqMap) {
	defer wg.Done()
	tempMap := Frequency(st)
	mutex.Lock()
	for k, v := range tempMap {
		(*m)[k] += v
	}
	mutex.Unlock()
}

// ConcurrentFrequency calls Frequency concurrently on an array of strings
// and combines their results
func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}

	var mutex = &sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(len(ss))

	for i := 0; i < len(ss); i++ {
		go tempFunc(ss[i], mutex, &wg, &m)
	}

	wg.Wait()

	return m
}
