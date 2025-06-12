package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) PostLogin(c *gin.Context) {
	var body map[string]string
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido"})
		return
	}

	if body["username"] != "admin" || body["password"] != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais invalidas"})
		return
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": body["username"],
		"exp": time.Now().Add(time.Hour).Unix(),
	}).SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
