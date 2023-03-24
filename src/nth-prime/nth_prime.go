package prime

import "errors"

var primes = []int{2}

// Nth returns the nth prime number.
// An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input must be greater than 1")
	}
	for i := 3; len(primes) < n; i++ {
		isPrime(i)
	}
	return primes[n-1], nil
}

func isPrime(n int) {
	for _, prime := range primes {
		if n%prime == 0 {
			return
		}
	}
	primes = append(primes, n)
}
