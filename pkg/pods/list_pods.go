package pods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) ListPods(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
