package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteClientes(c *gin.Context) {
	id := c.Param("id")

	err := h.ClienteRepo.Delete(c.Request.Context(), id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente"})
		return
	}

	c.Status(http.StatusNoContent)
}
