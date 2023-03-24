package cipher

type shift int
type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	d := shift(distance)

	switch {
	case d > 0 && d < 26:
		return d
	case d > -26 && d < 0:
		return d + 26
	}
	return nil
}

func (c shift) Encode(input string) string {
	var result []rune
	for _, r := range input {
		if r = encode(r, int(c)); r >= 0 {
			result = append(result, r)
		}
	}
	return string(result)
}

func (c shift) Decode(input string) string {
	var result []rune
	for _, r := range input {
		if r = decode(r, int(c)); r >= 0 {
			result = append(result, r)
		}
	}
	return string(result)
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}

	ok := false
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		} else if r > 'a' {
			ok = true
		}
	}

	if !ok {
		return nil
	}

	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	var result []rune
	p := 0
	for _, r := range input {
		if r = encode(r, int(v[p]-'a')); r >= 0 {
			result = append(result, r)
			p = (p + 1) % len(v)
		}
	}
	return string(result)
}

func (v vigenere) Decode(input string) string {
	var result []rune
	p := 0
	for _, r := range input {
		if r = decode(r, int(v[p]-'a')); r >= 0 {
			result = append(result, r)
			p = (p + 1) % len(v)
		}
	}
	return string(result)
}

func encode(r rune, s int) rune {
	if r >= 'a' && r <= 'z' {
		return (r-'a'+rune(s))%26 + 'a'
	} else if r >= 'A' && r <= 'Z' {
		return (r-'A'+rune(s))%26 + 'a'
	}
	return -1
}

func decode(r rune, s int) rune {
	if r >= 'a' && r <= 'z' {
		return (r-'a'+rune(26-s))%26 + 'a'
	}
	return -1
}
