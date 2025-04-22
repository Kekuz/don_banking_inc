package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type ClientService interface {
	FindById(string) domain.Client
}

type ClientHandler struct {
	service ClientService
}

func NewClientHandler(s ClientService) *ClientHandler {
	return &ClientHandler{service: s}
}

func (h *ClientHandler) GetClientById(id string) domain.Client{
	client := h.service.FindById(id)
	return client
}
