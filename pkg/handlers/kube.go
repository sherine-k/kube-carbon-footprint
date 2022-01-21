package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (h *Handlers) GetDatacenterInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	namespace := params["namespace"]
	pod := params["pod"]

	info, err := h.kubeClient.GetPodDatacenterInfo(pod, namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, info)
}
