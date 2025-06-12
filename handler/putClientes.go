package handler

import (
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

	query := "UPDATE clientes SET name=$1, email=$2 WHERE id=$3"
	result, err := h.DB.ExecContext(c.Request.Context(), query, cliente.Name, cliente.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	cliente.ID = id
	c.JSON(http.StatusOK, cliente)
}
