package bottlesong

import (
	"fmt"
	"strings"
)

var numberToWord = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	verses := make([]string, 0, takeDown*5-1)
	for i := 0; i < takeDown; i++ {
		verses = append(verses, verse(startBottles-i)...)

		if i < takeDown-1 {
			verses = append(verses, "")
		}
	}

	return verses
}

func verse(n int) []string {
	switch n {
	case 1:
		return []string{
			"One green bottle hanging on the wall,",
			"One green bottle hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be no green bottles hanging on the wall.",
		}
	case 2:
		return []string{
			"Two green bottles hanging on the wall,",
			"Two green bottles hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be one green bottle hanging on the wall.",
		}
	default:
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", numberToWord[n]),
			fmt.Sprintf("%s green bottles hanging on the wall,", numberToWord[n]),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(numberToWord[n-1])),
		}
	}
}
