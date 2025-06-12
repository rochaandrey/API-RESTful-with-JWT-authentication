package handler

import (
	"gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostClientes(c *gin.Context) {
	var newCliente models.Cliente

	if err := c.ShouldBindJSON(&newCliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	query := "INSERT INTO clientes (name, email) VALUES ($1, $2) RETURNING id"
	err := h.DB.QueryRowContext(c.Request.Context(), query, newCliente.Name, newCliente.Email).Scan(&newCliente.ID)

	if err != nil {
		log.Printf("Erro ao criar cliente: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, newCliente)
}
