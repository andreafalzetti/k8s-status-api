package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Hello(c *gin.Context) {
	response := map[string]string{"message": "Hello! Just trying out Okteto Preview Environments!"}
	c.JSON(http.StatusOK, response)
}
