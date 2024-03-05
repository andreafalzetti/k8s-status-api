package graphqlApi

import "github.com/graphql-go/graphql"

func GenerateSchema(pods *graphql.Field, health *graphql.Field) (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"pods":   pods,
			"health": health,
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	return schema, err
}

func podsSchema(resolver graphql.FieldResolveFn) *graphql.Field {
	fields := graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.String,
		},
		"restarts": &graphql.Field{
			Type: graphql.Int,
		},
	}
	podType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Pod",
		Fields: fields,
	})

	return &graphql.Field{
		Type:    graphql.NewList(podType),
		Resolve: resolver,
	}
}

func healthSchema(resolver graphql.FieldResolveFn) *graphql.Field {
	fields := graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
	}

	healthType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Health",
		Fields: fields,
	})

	return &graphql.Field{
		Type:    healthType,
		Resolve: resolver,
	}
}
