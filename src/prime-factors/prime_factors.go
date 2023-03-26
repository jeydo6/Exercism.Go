package prime

import "math"

func Factors(n int64) []int64 {
	var result []int64
	for factor := int64(2); factor <= n; {
		if n%factor == 0 && isPrime(factor) {
			result = append(result, factor)
			n /= factor
		} else {
			factor++
		}
	}
	return result
}

func isPrime(factor int64) bool {
	if factor < 2 {
		return false
	}

	for i := int64(2); i <= int64(math.Sqrt(float64(factor))); i++ {
		if factor%i == 0 {
			return false
		}
	}

	return true
}
