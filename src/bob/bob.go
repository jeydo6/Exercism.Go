package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isEmpty(remark) {
		return "Fine. Be that way!"
	}

	var isQuestionValue = isQuestion(remark)
	if isUpper(remark) {
		if isQuestionValue {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	}
	if isQuestionValue {
		return "Sure."
	} else {
		return "Whatever."
	}
}

func isUpper(remark string) bool {
	hasUpper := false
	for _, ch := range remark {
		if !hasUpper && unicode.IsUpper(ch) {
			hasUpper = true
		}

		if unicode.IsLower(ch) {
			return false
		}
	}
	return hasUpper
}

func isQuestion(remark string) bool {
	return remark[len(remark)-1] == '?'
}

func isEmpty(remark string) bool {
	return remark == ""
}
