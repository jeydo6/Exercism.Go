package isogram

import "strings"

func IsIsogram(word string) bool {
	chars := make([]int, 26)
	for _, ch := range strings.ToLower(word) {
		if ch < 'a' || ch > 'z' {
			continue
		}

		charIndex := ch - 'a'
		if chars[charIndex] > 0 {
			return false
		}

		chars[charIndex]++
	}

	return true
}
