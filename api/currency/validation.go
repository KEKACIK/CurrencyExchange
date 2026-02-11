package currency

import (
	"CurrencyExchange/config"
	"errors"
	"unicode/utf8"
)

var (
	ErrValidationNameEmpty = errors.New("Bad request: name empty")
)

func NameValidation(name string) error {
	if name == "" {
		return ErrValidationNameEmpty
	}

	return nil
}

var (
	ErrValidationCodeEmpty    = errors.New("Bad request: code empty")
	ErrValidationCodeTooSmall = errors.New("Bad request: code too small")
	ErrValidationCodeTooBig   = errors.New("Bad request: code too big")
)

func CodeValidation(code string) error {
	if code == "" {
		return ErrValidationNameEmpty
	}
	if utf8.RuneCountInString(code) < config.CurrencyCodeMinLen {
		return ErrValidationCodeTooSmall
	}
	if utf8.RuneCountInString(code) > config.CurrencyCodeMaxLen {
		return ErrValidationCodeTooBig
	}

	return nil
}

var (
	ErrValidationSignEmpty = errors.New("Bad request: sign empty")
)

func SignValidation(sign string) error {
	if sign == "" {
		return ErrValidationSignEmpty
	}

	return nil
}
