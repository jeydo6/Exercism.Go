package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var sb = strings.Builder{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			sb.WriteRune('z' + 'a' - unicode.ToLower(r))
		} else if unicode.IsDigit(r) {
			sb.WriteRune(r)
		}

		if sb.Len()%6 == 5 {
			sb.WriteRune(' ')
		}
	}

	return strings.TrimSpace(sb.String())
}
