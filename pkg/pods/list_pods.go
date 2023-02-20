package pods

import (
	"context"
	"net/http"
	"time"

	"github.com/andrealfalzetti/k8s-status-api/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetSortedPods(ctx context.Context, h handler, sortParam string) ([]Pod, error) {
	pods, err := util.GetPods(ctx, h.Kubernetes_client, h.Default_Namespace)
	if err != nil {
		return nil, err
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

	if sortParam == "" {
		return items, nil
	}

	sortedItems, err := SortPodsByParam(items, sortParam)
	if err != nil {
		return nil, err
	}

	return sortedItems, nil
}

func (h handler) ListPods(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h.Logger.Debug("Fetching pods from Kubernetes")

	sortParam := c.Query("sort")

	// TODO: handle both 4xx and 5xx errors
	items, err := GetSortedPods(ctx, h, sortParam)
	if err != nil {
		h.Logger.WithError(err).Error("Error retrieving pods")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, ListPodsResponse{
		Meta: Meta{
			Count: len(items),
		},
		Data: items,
	})
}
