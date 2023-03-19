package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	result := make([]rune, len(plain))
	shiftKeyInt32 := int32(shiftKey)
	for i, r := range plain {
		var shift rune
		if unicode.IsLetter(r) {
			current := unicode.ToLower(r) - 'a'
			shift = (current+shiftKeyInt32)%26 - current
		}
		result[i] = r + shift
	}
	return string(result)
}
