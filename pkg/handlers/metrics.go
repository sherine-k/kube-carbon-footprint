package handlers

import (
	"fmt"
	"net/http"

	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

func (h *Handlers) GetCPUMetrics(w http.ResponseWriter, r *http.Request) {
	client, err := prometheus.NewClient(h.promConfig)
	if err != nil {
		message := fmt.Sprintf("Error when creating the prometheus client: %v", err)
		hlog.Error(message)
		writeError(w, http.StatusServiceUnavailable, message)
		return
	}
	matrix, err := client.GetCPUMetrics()
	if err != nil {
		message := fmt.Sprintf("Error when getting CPU metrics: %v", err)
		hlog.Error(message)
		writeError(w, http.StatusServiceUnavailable, message)
		return
	}
	writeJSON(w, http.StatusOK, matrix)
}
