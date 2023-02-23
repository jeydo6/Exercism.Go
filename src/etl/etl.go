package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var out = map[string]int{}
	for key, values := range in {
		for _, value := range values {
			out[strings.ToLower(value)] = key
		}
	}
	return out
}
