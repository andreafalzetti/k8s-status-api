package main

import (
	"github.com/andrealfalzetti/k8s-status-api/pkg/pods"
	"github.com/andrealfalzetti/k8s-status-api/pkg/root"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	root.RegisterRoutes(r)
	pods.RegisterRoutes(r)

	r.Run()
}
