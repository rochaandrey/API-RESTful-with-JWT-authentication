package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClientesById(c *gin.Context) {
	id := c.Param("id")

	cliente, err := h.ClienteRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar cliente"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}
