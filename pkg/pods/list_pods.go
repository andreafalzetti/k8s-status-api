package pods

import (
	"context"
	"net/http"
	"time"

	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	"github.com/gin-gonic/gin"
)

func (h handler) ListPods(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.Logger.Debug("Fetching pods from Kubernetes")
	pods, err := util.GetPods(ctx, h.Kubernetes_client, h.Default_Namespace)
	if err != nil {
		h.Logger.WithError(err).Error("Error getting pods")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var items []Pod
	for _, pod := range pods.Items {
		age := time.Since(pod.CreationTimestamp.Time).Truncate(time.Second)
		restarts := 0
		for _, containerStatus := range pod.Status.ContainerStatuses {
			restarts += int(containerStatus.RestartCount)
		}
		items = append(items, Pod{
			Name:     pod.Name,
			Age:      age.String(),
			Restarts: restarts,
		})
	}

	// fmt.Println("Pods: ", items)
	// c.JSON(http.StatusOK, gin.H{
	// 	"items": pods,
	// })

	c.JSON(http.StatusOK, ListPodsResponse{
		Meta: Meta{
			Count: len(items),
		},
		Data: items,
	})
}
