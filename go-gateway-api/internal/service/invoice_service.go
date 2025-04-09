package service

import (
	"github.com/mpss1980/gateway/go-gateway/internal/domain"
	"github.com/mpss1980/gateway/go-gateway/internal/dto"
)

// InvoiceService provides operations on invoices
type InvoiceService struct {
	repo           domain.InvoiceRepository
	accountService AccountService
}

// NewInvoiceService creates a new InvoiceService
func NewInvoiceService(repo domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{repo: repo, accountService: accountService}
}

// Create creates a new invoice
func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(accountOutput.ID, input)
	if err != nil {
		return nil, err
	}

	// Process the invoice before saving
	err = invoice.Process()
	if err != nil {
		return nil, err
	}

	// Update balance if status is approved
	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(accountOutput.ID, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	// Save the invoice
	if err := s.repo.Save(invoice); err != nil {
		return nil, err
	}

	output := dto.FromInvoice(invoice)
	return &output, nil
}

// GetById retrieves an invoice by ID and API key
func (s *InvoiceService) GetById(id, apiKey string) (*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	output := dto.FromInvoice(invoice)
	return &output, nil
}

// ListByAccount retrieves all invoices for a given account ID
func (s *InvoiceService) ListByAccount(accountId string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.repo.FindByAccountID(accountId)
	if err != nil {
		return nil, err
	}

	outputs := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output := dto.FromInvoice(invoice)
		outputs[i] = &output
	}

	return outputs, nil
}

// ListByAccountAPIKey retrieves all invoices for a given API key
func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(accountOutput.ID)
}
