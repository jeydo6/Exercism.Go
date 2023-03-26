package phonenumber

import (
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	number, err := clean(phoneNumber)
	if err != nil {
		return "", err
	}
	numberLen := len(number)
	switch {
	case numberLen < 10:
		return "", fmt.Errorf("phone number %q contains too few digits", phoneNumber)
	case numberLen == 10 && number[0] >= '2' && number[3] >= '2':
		return number, nil
	case numberLen == 11 && number[0] == '1' && number[1] >= '2' && number[4] >= '2':
		return number[1:], nil
	}
	return "", fmt.Errorf("phone number %q is invalid", phoneNumber)
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}

func clean(phoneNumber string) (string, error) {
	number := make([]rune, 0, len(phoneNumber))
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			number = append(number, r)
		} else if r == ' ' ||
			r == '+' ||
			r == '(' ||
			r == ')' ||
			r == '-' ||
			r == '.' {
			//
		} else {
			return "", fmt.Errorf("phone number %q contains invalid characters", phoneNumber)
		}
	}

	return string(number), nil
}
