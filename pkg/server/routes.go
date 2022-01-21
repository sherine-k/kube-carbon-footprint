package server

import (
	"github.com/gorilla/mux"

	"github.com/sherine-k/kube-carbon-footprint/pkg/handlers"
	"github.com/sherine-k/kube-carbon-footprint/pkg/kube"
	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

func setupRoutes(promCfg prometheus.Config, kubeClient *kube.Client) *mux.Router {
	h := handlers.NewHandlers(promCfg, kubeClient)
	r := mux.NewRouter()
	r.HandleFunc("/api/status", handlers.Status)
	r.HandleFunc("/api/metrics/cpu", h.GetCPUMetrics)
	r.HandleFunc("/api/metrics/carbonfootprint", h.GetCarbonFootprint)
	r.HandleFunc("/api/dataset/instancetype/{instanceType}", h.GetInstanceTypeData)
	r.HandleFunc("/api/dataset/region/{region}", h.GetRegionData)
	r.HandleFunc("/api/datacenter/namespace/{namespace}/pod/{pod}", h.GetDatacenterInfo)
	return r
}
