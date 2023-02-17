package util

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type SharedHandlerContext struct {
	Logger            *logrus.Logger
	Kubernetes_client *kubernetes.Clientset
}
