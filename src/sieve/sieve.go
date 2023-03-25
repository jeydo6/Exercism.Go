package sieve

import "math"

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	numbers := make([]int, limit-1)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 2
	}

	for step := 2; step <= int(math.Sqrt(float64(limit))); step++ {
		for i := 2*step - 2; i < len(numbers); i += step {
			if numbers[i] > 0 {
				numbers[i] = 0
			}
		}
	}

	var result []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			result = append(result, numbers[i])
		}
	}

	return result
}
