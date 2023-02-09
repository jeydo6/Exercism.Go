package hamming

import "errors"

var errLengthShouldBeEqual = errors.New("the length should be equal")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errLengthShouldBeEqual
	}

	result := 0
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			result++
		}
	}

	return result, nil
}
