package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mpss1980/gateway/go-gateway/internal/domain"
)

// InvoiceRepository implements the domain.InvoiceRepository interface using SQL
type InvoiceRepository struct {
	db *sql.DB
}

// NewInvoiceRepository creates a new instance of InvoiceRepository
func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// Save inserts a new invoice record into the database
func (r *InvoiceRepository) Save(invoice *domain.Invoice) error {
	if _, err := r.db.Exec(`
		INSERT INTO invoices (id, account_id, status, description, payment_type, card_last_digits, amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`,
		invoice.ID,
		invoice.AccountID,
		invoice.Status,
		invoice.Description,
		invoice.PaymentType,
		invoice.CardLastDigits,
		invoice.Amount,
		invoice.CreatedAt,
		invoice.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}

// FindByID retrieves an invoice by its ID
func (r *InvoiceRepository) FindByID(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice

	err := r.db.QueryRow(`
		SELECT id, account_id, status, description, payment_type, card_last_digits, amount, created_at, updated_at
		FROM invoices
		WHERE id = $1
	`, id).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentType,
		&invoice.CardLastDigits,
		&invoice.Amount,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrInvoiceNotFound // Map to domain error
	}
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// FindByAccountID retrieves all invoices associated with a given account ID
func (r *InvoiceRepository) FindByAccountID(accountID string) ([]*domain.Invoice, error) {
	rows, err := r.db.Query(`
		SELECT id, account_id, status, description, payment_type, card_last_digits, amount, created_at, updated_at
		FROM invoices
		WHERE account_id = $1
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*domain.Invoice
	for rows.Next() {
		var invoice domain.Invoice

		err := rows.Scan(
			&invoice.ID,
			&invoice.AccountID,
			&invoice.Status,
			&invoice.Description,
			&invoice.PaymentType,
			&invoice.CardLastDigits,
			&invoice.Amount,
			&invoice.CreatedAt,
			&invoice.UpdatedAt,
		)
		if err != nil {
			return nil, err // Return on the first scan error
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

// UpdateStatus updates the status and updated_at timestamp of an invoice by ID
func (r *InvoiceRepository) UpdateStatus(id string, status domain.Status) error {
	result, err := r.db.Exec(`
		UPDATE invoices
		SET status = $1, updated_at = $2
		WHERE id = $3
	`, status, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrInvoiceNotFound // No rows updated means invoice wasn't found
	}

	return nil
}
