package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 {
		return 1, nil
	}

	if span < 0 || span > len(digits) {
		return 0, fmt.Errorf("argument exception: %d", span)
	}

	if len(digits) == 0 {
		return 0, fmt.Errorf("argument exception: %s", digits)
	}

	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return 0, fmt.Errorf("argument exception: %v", r)
		}
	}

	result := int64(0)
	for i := 0; i < len(digits)-span+1; i++ {
		product := int64(1)
		for j := 0; j < span; j++ {
			product *= int64(digits[i+j] - '0')
		}

		if product > result {
			result = product
		}
	}
	return result, nil
}
