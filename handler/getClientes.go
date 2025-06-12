package handler

import (
	"fmt"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClientes(c *gin.Context) {
	nameFilter := c.Query("name")
	emailFilter := c.Query("email")

	query := "SELECT id, name, email FROM clientes"
	args := []interface{}{}
	// variavel de controle p where ou and
	keyword := " WHERE"

	if nameFilter != "" {
		query += fmt.Sprintf("%s name ILIKE $%d", keyword, len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", nameFilter))
		keyword = " AND"
	}

	if emailFilter != "" {
		query += fmt.Sprintf("%s email ILIKE $%d", keyword, len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", emailFilter))
		keyword = " AND"
	}

	rows, err := h.DB.QueryContext(c.Request.Context(), query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes: " + err.Error()})
		return
	}
	defer rows.Close()

	var clientes []models.Cliente
	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Name, &cliente.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao escanear cliente"})
			return
		}
		clientes = append(clientes, cliente)
	}

	if len(clientes) == 0 {
		c.JSON(http.StatusOK, []models.Cliente{})
		return
	}

	c.JSON(http.StatusOK, clientes)
}
