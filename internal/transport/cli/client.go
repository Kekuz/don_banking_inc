package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type ClientService interface {
	FindById(int) domain.Client
}

type ClientHandler struct {
	Service ClientService
}

func NewClientHandler(s ClientService) *ClientHandler {
	return &ClientHandler{Service: s}
}

func (h *ClientHandler) GetClientById(id int) domain.Client{
	client := h.Service.FindById(id)
	return client
}
