package pods

import (
	"context"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func createPodMock(name string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:              name,
			Namespace:         "default",
			CreationTimestamp: metav1.Time{Time: time.Now()},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "container-",
					Image: "some-random-image",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					RestartCount: 0,
				},
			},
		},
	}
}

func TestGetSortedPods(t *testing.T) {
	ctx := context.Background()

	// Creating a mock client
	clientset := fake.NewSimpleClientset()

	// Creating some mock data
	pod1 := createPodMock("pod-1")
	pod2 := createPodMock("pod-2")

	// Injecting the data
	clientset.CoreV1().Pods("default").Create(ctx, pod1, metav1.CreateOptions{})
	clientset.CoreV1().Pods("default").Create(ctx, pod2, metav1.CreateOptions{})

	// Injecting the client into the handler
	h := handler{
		SharedHandlerContext: &util.SharedHandlerContext{
			Kubernetes_client: clientset,
			Default_Namespace: "default",
		},
	}

	// Check that it returns the right amount of pods
	pods, err := GetSortedPods(ctx, h, "")
	assert.Nil(t, err)
	assert.Len(t, pods, 2)
}
