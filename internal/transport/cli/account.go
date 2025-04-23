package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type AccountService interface {
	FindById(id int) []domain.Account
	WriteAccount(client domain.Client, currency domain.Currency)
	DeleteAccount(id int, currency domain.Currency)
	UpdateAccountBalance(id int, currency domain.Currency)
}

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(s AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) GetAccountsById(id int) []domain.Account {
	accounts := h.service.FindById(id)
	return accounts
}

func (h *AccountHandler) CreateNewAccount(client domain.Client, currency domain.Currency) {
	h.service.WriteAccount(client, currency)
}
