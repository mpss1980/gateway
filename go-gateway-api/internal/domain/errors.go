package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrDuplicatedAPIKey   = errors.New("API key already exists")
	ErrInvoiceNotFound    = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidStatus      = errors.New("invalid status")
	ErrAmountTooLarge     = errors.New("amount exceeds processing limit")
)
