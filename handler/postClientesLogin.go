package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostClientesLogin(c *gin.Context) {
	// retorna token jwt
	c.JSON(http.StatusOK, gin.H{
		"message ": "post ok",
	})
}
