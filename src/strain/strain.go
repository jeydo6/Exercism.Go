package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (result Ints) {
	for _, item := range i {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

func (i Ints) Discard(filter func(int) bool) (result Ints) {
	for _, item := range i {
		if !filter(item) {
			result = append(result, item)
		}
	}

	return result
}

func (l Lists) Keep(filter func([]int) bool) (result Lists) {
	for _, item := range l {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

func (s Strings) Keep(filter func(string) bool) (result Strings) {
	for _, item := range s {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}
