package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Status(w http.ResponseWriter, r *http.Request) {
	// TODO: check prom connectivity...
	_, err := w.Write([]byte("OK"))
	if err != nil {
		logrus.Errorf("could not write response: %v", err)
	}
}
