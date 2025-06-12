package routers

import (
	"gin/handler"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func initializeRouters(r *gin.Engine, h *handler.Handler) {
	myGroup := r.Group("/myapi")
	{
		myGroup.POST("/login", h.PostLogin)
	}

	protected := myGroup.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		myGroup.GET("/clientes", h.GetClientes)
		myGroup.GET("/clientes/:id", h.GetClientesById)
		myGroup.POST("/clientes", h.PostClientes)
		myGroup.PUT("/clientes/:id", h.PutClientes)
		myGroup.DELETE("/clientes/:id", h.DeleteClientes)
	}

}
