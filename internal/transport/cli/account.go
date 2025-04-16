package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type AccountService interface {
	FindById(string) []domain.Account
}

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(s AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) GetAccountsById(id string) []domain.Account{
	accounts := h.service.FindById(id)
	return accounts
}
