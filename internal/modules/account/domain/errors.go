package domain

import "errors"

var (
	ErrInvalidAmount      = errors.New("amount must be greater than zero")
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrAccountNotFound    = errors.New("account not found")
	ErrNameInvalid        = errors.New("The name must contain more than 3 and up to 120 characters")
	ErrEmailInvalid       = errors.New("Email format not accepted")
	ErrInvalidCPF         = errors.New("This CPF is invalid.")
	ErrNumberPhoneInvalid = errors.New("The phone number must have 10 or 11 digits, including the area code (DDD).")
)
