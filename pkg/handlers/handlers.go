package handlers

import (
	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

type Handlers struct {
	promConfig prometheus.Config
}

func NewHandlers(promConfig prometheus.Config) *Handlers {
	return &Handlers{
		promConfig: promConfig,
	}
}
