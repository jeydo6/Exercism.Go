package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var result = Frequency{}
	for _, word := range getWords(phrase) {
		if _, exists := result[word]; !exists {
			result[word] = 0
		}
		result[word]++
	}

	return result
}

func getWords(phrase string) []string {
	var result []string
	sb := strings.Builder{}

	for _, r := range phrase {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\\' || r == '\'' {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			if sb.Len() > 0 {
				word := strings.Trim(sb.String(), "'")
				if len(word) > 0 {
					result = append(result, word)
				}

				sb.Reset()
			}
		}
	}

	if sb.Len() > 0 {
		word := strings.Trim(sb.String(), "'")
		if len(word) > 0 {
			result = append(result, word)
		}

		sb.Reset()
	}

	return result
}
