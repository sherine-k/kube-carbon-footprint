package server

import (
	"github.com/gorilla/mux"

	"github.com/sherine-k/kube-carbon-footprint/pkg/handlers"
	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

func setupRoutes(cfg prometheus.Config) *mux.Router {
	h := handlers.NewHandlers(cfg)
	r := mux.NewRouter()
	r.HandleFunc("/api/status", handlers.Status)
	r.HandleFunc("/api/metrics/cpu", h.GetCPUMetrics)
	r.HandleFunc("/api/dataset/instancetype/{instanceType}", h.GetInstanceTypeData)
	return r
}
