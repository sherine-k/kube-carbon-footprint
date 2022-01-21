package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/sherine-k/kube-carbon-footprint/pkg/dataset"
	"github.com/sherine-k/kube-carbon-footprint/pkg/kube"
	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

var hlog = logrus.WithField("module", "handler")

type Handlers struct {
	promConfig prometheus.Config
	kubeClient *kube.Client
	dataset    *dataset.Dataset
}

func NewHandlers(promConfig prometheus.Config, kubeClient *kube.Client) *Handlers {
	ds, err := dataset.Load()
	if err != nil {
		hlog.Errorf("Cannot load dataset: %v", err)
		ds = &dataset.Dataset{}
	}
	return &Handlers{
		promConfig: promConfig,
		kubeClient: kubeClient,
		dataset:    ds,
	}
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
