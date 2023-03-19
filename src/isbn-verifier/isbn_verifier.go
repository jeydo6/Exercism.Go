package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	digits := make([]int, 0, len(isbn))
	for i, ch := range isbn {
		if unicode.IsDigit(ch) {
			digits = append(digits, int(ch-'0'))
		} else if i == len(isbn)-1 && ch == 'X' {
			digits = append(digits, 10)
		} else if ch != '-' {
			return false
		}
	}

	digitsLen := len(digits)
	if digitsLen != 10 {
		return false
	}

	result := 0
	for i := 0; i < digitsLen; i++ {
		result += digits[digitsLen-1-i] * (i + 1)
	}

	return result%(digitsLen+1) == 0
}
