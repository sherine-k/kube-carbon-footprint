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

func (h *Handlers) GetCarbonFootprint(w http.ResponseWriter, r *http.Request) {

	// collect request parameters
	parameters := r.URL.Query()
	regionName := parameters.Get("region")
	if regionName == "" {
		writeError(w, http.StatusBadRequest, "missing paramter: region")
	}
	// instanceType := parameters.Get("instancetype")
	// if instanceType == "" {
	// 	writeError(w, http.StatusBadRequest, "missing paramter: instancetype")
	// }
	pod := parameters.Get("pod")
	namespace := parameters.Get("namespace")

	//get PUE and CO2e for region
	region := h.dataset.FindRegion(regionName)
	fmt.Printf("%s", region.Name)
	//get load
	if pod == "" && namespace == "" {
		//TODO mais pas du tout
		h.GetCPUMetrics(w, r)
		return
	} else if pod != "" && namespace != "" {
		client, err := prometheus.NewClient(h.promConfig)
		if err != nil {
			message := fmt.Sprintf("Error when creating the prometheus client: %v", err)
			hlog.Error(message)
			writeError(w, http.StatusServiceUnavailable, message)
			return
		}
		matrix, err := client.GetCPUAvgByPod(pod, namespace)
		writeJSON(w, http.StatusOK, matrix)
		if err != nil {
			message := fmt.Sprintf("Error when getting CPU metrics for pod %s: %v", pod, err)
			hlog.Error(message)
			writeError(w, http.StatusServiceUnavailable, message)
			return
		}
	}

	//compute C footprint

	//gCO₂eq = PUE * Power * ZoneCO2e / 1000

}
