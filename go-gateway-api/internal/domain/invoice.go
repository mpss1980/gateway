package domain

import (
	"time"

	"math/rand"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	Amount         float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewInvoice creates a new Invoice instance, validating amount
func NewInvoice(accountID, description, paymentType string, amount float64, creditCard CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := ""
	if len(creditCard.Number) >= 4 {
		lastDigits = creditCard.Number[len(creditCard.Number)-4:]
	}

	// Return struct literal directly
	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		Amount:         amount,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// CreditCard represents credit card details (potentially sensitive)
type CreditCard struct {
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	CVV             string
	HolderName      string
}

// Process handles the processing logic for an invoice
func (i *Invoice) Process() error {
	if i.Amount > 1000 {
		return nil
	}

	// Combine source creation, rand instance, and Float64 call
	randomSourceValue := rand.New(rand.NewSource(time.Now().Unix())).Float64()

	if randomSourceValue > 0.7 {
		i.Status = StatusApproved
	} else {
		i.Status = StatusRejected
	}

	i.UpdatedAt = time.Now()

	return nil
}

// UpdateStatus updates the status only if the current status is not pending.
func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status == StatusPending {
		return ErrInvalidStatus // Cannot update status if it's still pending
	}

	// Optional: Add validation for newStatus if needed (e.g., check if it's a valid known status)

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
