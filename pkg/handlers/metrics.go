package handlers

import (
	"fmt"
	"net/http"

	"github.com/sherine-k/kube-carbon-footprint/pkg/compute"
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

	pod := parameters.Get("pod")
	namespace := parameters.Get("namespace")

	if pod == "" || namespace == "" {
		message := fmt.Sprintf("Missing parameters: pod=%s, namespace=%s", pod, namespace)
		hlog.Error(message)
		writeError(w, http.StatusBadRequest, message)
		return
	}

	dcInfo, err := h.kubeClient.GetPodDatacenterInfo(pod, namespace)
	if err != nil {
		message := fmt.Sprintf("Error getting datacenter info: %v", err)
		hlog.Error(message)
		writeError(w, http.StatusInternalServerError, message)
		return
	}
	//get region data
	region := h.dataset.FindRegion(dcInfo.Region)
	if region == nil {
		message := fmt.Sprintf("Data for region %s not found", dcInfo.Region)
		hlog.Error(message)
		writeError(w, http.StatusInternalServerError, message)
		return
	}

	//get instancetype data of the underlying node
	instanceType := h.dataset.FindInstance(dcInfo.InstanceType)
	if instanceType == nil {
		message := fmt.Sprintf("Data for instance type  %s not found", dcInfo.InstanceType)
		hlog.Error(message)
		writeError(w, http.StatusInternalServerError, message)
		return
	}

	//get load
	client, err := prometheus.NewClient(h.promConfig)
	if err != nil {
		message := fmt.Sprintf("Error when creating the prometheus client: %v", err)
		hlog.Error(message)
		writeError(w, http.StatusServiceUnavailable, message)
		return
	}
	matrix, err := client.GetCPUByPod(pod, namespace)
	if err != nil {
		message := fmt.Sprintf("Error when getting CPU metrics for pod %s: %v", pod, err)
		hlog.Error(message)
		writeError(w, http.StatusServiceUnavailable, message)
		return
	}

	//compute C footprint
	cfpMatrix := compute.ComputeCarbonFootprint(matrix, instanceType, region)

	writeJSON(w, http.StatusOK, cfpMatrix)

}
