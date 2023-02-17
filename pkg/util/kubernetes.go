package util

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// TODO: extend this function to support creating the client outside the cluster
func GetKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
