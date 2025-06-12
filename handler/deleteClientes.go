package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteClientes(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM clientes WHERE id=$1"
	result, err := h.DB.ExecContext(c.Request.Context(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.Status(http.StatusNoContent)
}
