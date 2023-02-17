package root

import (
	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	"github.com/gin-gonic/gin"
)

type handler struct {
	*util.SharedHandlerContext
}

func RegisterRoutes(r *gin.Engine, shc *util.SharedHandlerContext) {
	h := &handler{shc}
	routes := r.Group("/")
	routes.GET("/", h.Hello)
}
