package routers

import (
	"gin/handler"

	"github.com/gin-gonic/gin"
)

func initializeRouters(r *gin.Engine, h *handler.Handler) {
	myGroup := r.Group("/myapi")
	{
		myGroup.GET("/clientes", h.GetClientes)
		myGroup.GET("/clientes/:id", h.GetClientesById)
		myGroup.POST("/login", h.PostClientesLogin)
		myGroup.POST("/clientes", h.PostClientesLogin)
		myGroup.PUT("/clientes/:id", h.PutClientes)
		myGroup.DELETE("/clientes/:id", h.DeleteClientes)
	}
}
