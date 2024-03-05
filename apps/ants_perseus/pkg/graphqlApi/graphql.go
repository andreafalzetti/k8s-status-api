package graphqlApi

import (
	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type handler struct {
	*util.SharedHandlerContext
	gqls graphql.Schema
}

func RegisterRoutes(r *gin.Engine, shc *util.SharedHandlerContext) {
	h := &handler{
		SharedHandlerContext: shc,
		gqls:                 graphql.Schema{},
	}

	// init resolvers
	pods := podsSchema(h.ListPodsResolver)
	health := healthSchema(h.HealthResolver)

	// create schema
	schema, err := GenerateSchema(pods, health)
	if err != nil {
		panic(err)
	}
	h.gqls = schema

	// declare route
	routes := r.Group("/graphql")
	routes.POST("/", h.RunQuery)
}
