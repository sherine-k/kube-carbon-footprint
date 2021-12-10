package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

var slog = logrus.WithField("module", "server")

type Config struct {
	Port           int
	CertFile       string
	PrivateKeyFile string
}

func Start(cfg Config, promCfg prometheus.Config) {
	router := setupRoutes(promCfg)

	// Clients must use TLS 1.2 or higher
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	httpServer := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		TLSConfig:    tlsConfig,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if cfg.CertFile != "" && cfg.PrivateKeyFile != "" {
		slog.Infof("listening on https://:%d", cfg.Port)
		panic(httpServer.ListenAndServeTLS(cfg.CertFile, cfg.PrivateKeyFile))
	} else {
		slog.Infof("listening on http://:%d", cfg.Port)
		panic(httpServer.ListenAndServe())
	}
}
