package utils

import (
	"jatin/pkg/errors"
	"unicode"
)

func IsPasswordStrong(p string) error {
	errs := []error{
		errors.ErrPasswordIsShort,
		errors.ErrPasswordNotHaveUpperCase,
		errors.ErrPasswordNotHaveLowerCase,
		errors.ErrPasswordNotHaveNumber,
		errors.ErrPasswordNotHaveSpecialChar,
	}
	validation := []bool{
		false,
		false,
		false,
		false,
		false,
	}
	s := []rune(p)
	if len(s) >= 7 {
		validation[0] = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			validation[1] = true
		case unicode.IsLower(char):
			validation[2] = true
		case unicode.IsNumber(char):
			validation[3] = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			validation[4] = true
		}
	}
	for i, vdt := range validation {
		if !vdt {
			return errs[i]
		}
	}
	return nil
}
