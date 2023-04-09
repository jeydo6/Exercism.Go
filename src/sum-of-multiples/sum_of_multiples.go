package summultiples

func SumMultiples(limit int, divisors ...int) int {
	result := 0
	for number := 1; number < limit; number++ {
		for _, divisor := range divisors {
			if divisor != 0 && number%divisor == 0 {
				result += number
				break
			}
		}
	}
	return result
}
