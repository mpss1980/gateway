package handlers

import "github.com/mpss1980/gateway/go-gateway/internal/service"

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}
