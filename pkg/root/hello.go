package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Hello(c *gin.Context) {
	response := map[string]string{"message": "Hola Hola! Check out the /pods endpoint for more interesting stuff :)"}
	c.JSON(http.StatusOK, response)
}
