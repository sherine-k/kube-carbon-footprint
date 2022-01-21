package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) GetInstanceTypeData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	instanceType := params["instanceType"]

	instance := h.dataset.FindInstance(instanceType)
	if instance == nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Instance type '%s' not found in dataset", instanceType))
		return
	}

	writeJSON(w, http.StatusOK, instance)
}

func (h *Handlers) GetRegionData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regionName := params["region"]

	region := h.dataset.FindRegion(regionName)
	if region == nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Region '%s' not found in dataset", regionName))
		return
	}

	writeJSON(w, http.StatusOK, region)
}
