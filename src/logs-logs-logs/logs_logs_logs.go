package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var runes = map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, ch := range log {
		if _, exists := runes[ch]; exists {
			return runes[ch]
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""

	for _, ch := range log {
		if ch == oldRune {
			ch = newRune
		}

		result += string(ch)
	}

	return result
}

// WithinLimit determines whether the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
