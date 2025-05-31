package validation

import (
	"errors"
	"regexp"
	"unicode"
)

func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("e-mail inválido")
	}
	return nil
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasDigit, hasSpecial bool

	for _, c := range password {
		switch {

		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasDigit = true
		}
	}

	specialCharRegex := regexp.MustCompile(`[!@#$%^&*()\-_=+\[\]{}|;:'",.<>?/\\` + "`~]")

	hasSpecial = specialCharRegex.MatchString(password)

	if !hasUpper || !hasDigit || !hasSpecial {
		return false
	}

	return true
}

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^(?:\+55\s?)?(?:\(?\d{2}\)?\s?)?(?:9\d{4}|\d{4})-?\d{4}$`)
	return re.MatchString(phone)
}
