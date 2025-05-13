package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type ClientService interface {
	FindById(int)(domain.Client, error)
}

type ClientHandler struct {
	Service ClientService
}

func NewClientHandler(s ClientService) *ClientHandler {
	return &ClientHandler{Service: s}
}

func (h *ClientHandler) GetClientById(id int) (domain.Client, error){
	client, err := h.Service.FindById(id)
	return client, err
}
