package main

import (
	"fmt"

	"github.com/andrealfalzetti/k8s-status-api/pkg/pods"
	"github.com/andrealfalzetti/k8s-status-api/pkg/root"
	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Debug("Starting k8s-status-api")

	kubernetes_client, err := util.GetKubernetesClient()
	if err != nil {
		logger.Fatalf("Failed to create Kubernetes client %d", err)
	}

	fmt.Println(kubernetes_client)

	r := gin.Default()

	h := &util.SharedHandlerContext{
		Logger:            logger,
		Kubernetes_client: kubernetes_client,
	}

	root.RegisterRoutes(r, h)
	pods.RegisterRoutes(r, h)

	r.Run()
}
