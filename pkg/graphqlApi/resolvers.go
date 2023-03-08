package graphqlApi

import (
	"context"
	"time"

	"github.com/andrealfalzetti/k8s-status-api/pkg/pods"
	"github.com/graphql-go/graphql"
)

func (h handler) ListPodsResolver(p graphql.ResolveParams) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sortParam := "name"
	pods, err := pods.GetSortedPods(ctx, h.Kubernetes_client, h.Default_Namespace, sortParam)
	return pods, err
}

type Response struct {
	Status string
}

func (h handler) HealthResolver(p graphql.ResolveParams) (interface{}, error) {
	res := Response{Status: "ok"}
	return res, nil
}
