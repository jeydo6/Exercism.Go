package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	normalizedSubject := normalize(subject)
	var matches []string
	for _, candidate := range candidates {
		if strings.EqualFold(candidate, subject) {
			continue
		}

		normalizedCandidate := normalize(candidate)
		if normalizedCandidate == normalizedSubject {
			matches = append(matches, candidate)
		}
	}

	return matches
}

func normalize(s string) string {
	chars := strings.Split(strings.ToLower(s), "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
