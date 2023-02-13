package robotname

import (
	"errors"
	"math/rand"
)

type Robot struct {
	name string
}

const capacity = 676 * 1000

var usedNames = map[string]bool{"": true}

func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}

	nameCandidate := ""
	for exists := true; exists; _, exists = usedNames[nameCandidate] {
		if len(usedNames) == capacity+1 {
			return "", errors.New("namespace is exhausted")
		}
		nameCandidate = getRandomName()
	}

	r.name = nameCandidate
	usedNames[nameCandidate] = true
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func getRandomName() string {
	number := 100 + rand.Intn(capacity+100)

	result := make([]rune, 5)

	for i := len(result) - 1; i >= 2; i-- {
		result[i] = rune('0' + number%10)
		number /= 10
	}

	for i := 1; i >= 0; i-- {
		result[i] = rune('A' + number%26)
		number /= 26
	}

	return string(result)
}
