package middleware

import (
	"net/http"

	"github.com/mpss1980/gateway/go-gateway/internal/domain"
	"github.com/mpss1980/gateway/go-gateway/internal/service"
)

// AuthMiddleware is a middleware that checks for a valid APIKey in the request header
type AuthMiddleware struct {
	accountService *service.AccountService
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

// Authenticate returns the HTTP handler that performs the authentication check
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API_KEY")
		if apiKey == "" {
			http.Error(w, "X-API_KEY is required", http.StatusUnauthorized)
			return
		}

		// Validate the API key against the account service
		_, err := m.accountService.FindByAPIKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
