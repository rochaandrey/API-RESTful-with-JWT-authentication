package handler

import (
	"database/sql"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PutClientes(c *gin.Context) {
	id := c.Param("id")
	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	err := h.ClienteRepo.Update(c.Request.Context(), id, cliente)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	cliente.ID = id
	c.JSON(http.StatusOK, cliente)
}
