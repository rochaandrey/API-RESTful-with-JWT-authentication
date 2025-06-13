package handler

import (
	"gin/repository"
)

type Handler struct {
	ClienteRepo repository.ClienteRepository
}

func NewHandler(repo repository.ClienteRepository) *Handler {
	return &Handler{ClienteRepo: repo}
}
