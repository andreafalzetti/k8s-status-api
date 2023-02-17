package pods

import (
	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	"github.com/gin-gonic/gin"
)

type handler struct {
	*util.SharedHandlerContext
}

func RegisterRoutes(r *gin.Engine, shc *util.SharedHandlerContext) {
	h := &handler{shc}
	routes := r.Group("/pods")
	routes.GET("/", h.ListPods)
}
