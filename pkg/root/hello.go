package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Hello(c *gin.Context) {
	response := map[string]string{"message": "Hello! Check out the /pods endpoint for more interesting stuff :)"}
	c.JSON(http.StatusOK, response)
}
