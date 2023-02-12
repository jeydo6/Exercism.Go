package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(l))

	ch := make(chan FreqMap)
	for _, s := range l {
		go func(s string) {
			ch <- Frequency(s)
			wg.Done()
		}(s)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := FreqMap{}
	for item := range ch {
		for letter, frequency := range item {
			result[letter] += frequency
		}
	}

	return result
}
