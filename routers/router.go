package routers

import (
	"gin/handler"

	"github.com/gin-gonic/gin"
)

func Initialize(h *handler.Handler) {
	r := gin.Default()
	initializeRouters(r, h)
	r.Run("8080")
}
