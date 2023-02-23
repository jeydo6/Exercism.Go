package pangram

func IsPangram(input string) bool {
	var letters [26]int
	for i := 0; i < len(input); i++ {
		letterIndex := getLetterIndex(input[i])
		if letterIndex > -1 {
			letters[letterIndex]++
		}
	}

	for i := 0; i < len(letters); i++ {
		if letters[i] == 0 {
			return false
		}
	}

	return true
}

func getLetterIndex(ch uint8) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch - 'a')
	} else if ch >= 'A' && ch <= 'Z' {
		return int(ch - 'A')
	}
	return -1
}
