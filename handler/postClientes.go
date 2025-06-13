package handler

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostClientes(c *gin.Context) {
	var newCliente models.Cliente
	if err := c.ShouldBindJSON(&newCliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	createdCliente, err := h.ClienteRepo.Create(c.Request.Context(), newCliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, createdCliente)
}
