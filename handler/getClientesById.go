package handler

import (
	"database/sql"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClientesById(c *gin.Context) {
	id := c.Param("id")
	var cliente models.Cliente

	query := "SELECT id, name, email FROM clientes WHERE id=$1"
	err := h.DB.QueryRowContext(c.Request.Context(), query, id).Scan(&cliente.ID, &cliente.Name, &cliente.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro no servidor"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}
