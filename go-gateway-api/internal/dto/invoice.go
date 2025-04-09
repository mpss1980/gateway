package dto

import (
	"time"

	"github.com/mpss1980/gateway/go-gateway/internal/domain"
)

// Status constants
const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

// CreateInvoiceInput represents the request body for creating a new invoice
type CreateInvoiceInput struct {
	APIKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardholderName string  `json:"cardholder_name"`
}

// InvoiceOutput represents the response body for invoice operations
type InvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	Amount         float64   `json:"amount"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ToInvoice converts CreateInvoiceInput to a domain.Invoice
func ToInvoice(accountID string, input CreateInvoiceInput) (*domain.Invoice, error) {
	creditCard := domain.CreditCard{
		Number:          input.CardNumber,
		CVV:             input.CVV,
		ExpirationMonth: input.ExpiryMonth,
		ExpirationYear:  input.ExpiryYear,
		HolderName:      input.CardholderName,
	}

	invoice, err := domain.NewInvoice(
		accountID,
		input.Description,
		input.PaymentType,
		input.Amount,
		creditCard,
	)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// FromInvoice converts a domain.Invoice to InvoiceOutput
func FromInvoice(invoice *domain.Invoice) InvoiceOutput {
	return InvoiceOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		Amount:         invoice.Amount,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
