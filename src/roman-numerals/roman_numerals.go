package romannumerals

import (
	"errors"
	"strings"
)

type arabicRomanPair struct {
	arabic int
	roman  string
}

var arabicRomanPairs = []arabicRomanPair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid input")
	}

	sb := strings.Builder{}
	for _, p := range arabicRomanPairs {
		for input >= p.arabic {
			sb.WriteString(p.roman)
			input -= p.arabic
		}
	}
	return sb.String(), nil
}
