package armstrong

import "strconv"

func IsNumber(n int) bool {
	str := strconv.Itoa(n)

	result := 0
	for i := 0; i < len(str); i++ {
		result += pow(int(str[i]-'0'), len(str))
	}

	return result == n
}

func pow(n int, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= n
	}
	return result
}
