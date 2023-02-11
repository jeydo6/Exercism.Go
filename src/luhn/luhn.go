package luhn

func Valid(id string) bool {

	digits := make([]int, 0, len(id))
	for _, ch := range id {
		if ch >= '0' && ch <= '9' {
			digits = append(digits, int(ch-'0'))
		} else if ch != ' ' {
			return false
		}
	}

	if len(digits) < 2 {
		return false
	}

	result := 0
	for i := 0; i < len(digits); i++ {
		digit := digits[len(digits)-1-i]
		if i%2 == 1 {
			if digit*2 > 9 {
				result += digit*2 - 9
			} else {
				result += digit * 2
			}
		} else {
			result += digit
		}
	}

	return result%10 == 0
}
