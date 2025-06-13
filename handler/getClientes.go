package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClientes(c *gin.Context) {
	nameFilter := c.Query("name")
	emailFilter := c.Query("email")

	clientes, err := h.ClienteRepo.GetAll(c.Request.Context(), nameFilter, emailFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
		return
	}

	c.JSON(http.StatusOK, clientes)
}
