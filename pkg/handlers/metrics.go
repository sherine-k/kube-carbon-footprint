package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

var hlog = logrus.WithField("module", "handler")

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

func writeJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		hlog.Errorf("Marshalling error while responding JSON: %v", err)
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		hlog.Errorf("Error while responding JSON: %v", err)
	}
}

type errorResponse struct{ Message string }

func writeError(w http.ResponseWriter, code int, message string) {
	response, err := json.Marshal(errorResponse{Message: message})
	if err != nil {
		hlog.Errorf("Marshalling error while responding an error: %v (message was: %s)", err, message)
		code = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		hlog.Errorf("Error while responding an error: %v (message was: %s)", err, message)
	}
}
