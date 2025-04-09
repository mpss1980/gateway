package handlers

import (
	"net/http"

	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/mpss1980/gateway/go-gateway/internal/domain"
	"github.com/mpss1980/gateway/go-gateway/internal/dto"
	"github.com/mpss1980/gateway/go-gateway/internal/service"
)

// InvoiceHandler handles HTTP requests for invoices
type InvoiceHandler struct {
	service *service.InvoiceService
}

// NewInvoiceHandler creates a new InvoiceHandler
func NewInvoiceHandler(service *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

// Create handles the creation of a new invoice
func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API_KEY")
	if apiKey == "" {
		http.Error(w, "X-API_KEY is required", http.StatusBadRequest)
		return
	}

	var input dto.CreateInvoiceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = apiKey

	output, err := h.service.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// GetByID handles retrieving an invoice by ID
func (h *InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	apiKey := r.Header.Get("X-API_KEY")
	if apiKey == "" {
		http.Error(w, "X-API_KEY is required", http.StatusBadRequest)
		return
	}

	output, err := h.service.GetById(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// ListByAccount handles listing invoices by account ID
func (h *InvoiceHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API_KEY")
	if apiKey == "" {
		http.Error(w, "X-API_KEY is required", http.StatusBadRequest)
		return
	}

	outputs, err := h.service.ListByAccountAPIKey(apiKey)
	if err != nil {
		switch err {
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputs)
}

// ... additional handler methods can be added here ...
