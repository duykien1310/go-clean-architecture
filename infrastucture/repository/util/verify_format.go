package util

import (
	"regexp"
)

func CheckMailFormat(email string) (bool, error) {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false, err
	}
	return match, nil
}

func CheckNameFormat(name string, maxLength int) (bool, error) {
	nameRegex := `^[ぁ-んァ-ン一-龯]+$`

	// Check length
	if len(name) > maxLength {
		return false, nil
	}

	match, err := regexp.MatchString(nameRegex, name)
	if err != nil {
		return false, err
	}
	return match, nil
}

func CheckCodeFormat(code string, maxLength int) (bool, error) {
	codeRegex := `^[A-Za-z0-9]+$`

	// Check length
	if len(code) > maxLength {
		return false, nil
	}

	// Kiểm tra ký tự đặc biệt, dấu cách, và kí tự không phải là chữ cái và số bằng regex
	match, err := regexp.MatchString(codeRegex, code)
	if err != nil {
		return false, err
	}
	return match, nil
}
