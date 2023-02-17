package root

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func RegisterRoutes(r *gin.Engine) {
	h := &handler{}

	routes := r.Group("/")
	routes.GET("/", h.Hello)
}
