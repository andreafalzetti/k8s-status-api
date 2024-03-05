package graphqlApi

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var req struct {
	Query string `json:"query"`
}

func (h handler) RunQuery(c *gin.Context) {
	h.Logger.Debug("Running graphql query")

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Execute the GraphQL query
	result := graphql.Do(graphql.Params{
		Schema:        h.gqls,
		RequestString: req.Query,
	})

	// Return the GraphQL result
	c.JSON(200, result)
}
