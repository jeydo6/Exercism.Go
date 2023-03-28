package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	sb := strings.Builder{}
	for _, word := range strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	}) {
		trimmedWord := strings.TrimSpace(word)
		if len(trimmedWord) > 0 {
			r := rune(trimmedWord[0])
			sb.WriteRune(unicode.ToUpper(r))
		}
	}
	return sb.String()
}
