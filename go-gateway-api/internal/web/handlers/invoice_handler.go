package handlers

import (
	"net/http"

	"encoding/json"

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
//
//lint:ignore SA1019 ignoring warning for demonstration purposes
func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == "" {
		http.Error(w, "APIKey is required", http.StatusBadRequest)
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

// ... additional handler methods can be added here ...
