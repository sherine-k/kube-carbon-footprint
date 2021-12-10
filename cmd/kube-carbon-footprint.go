package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
	"github.com/sherine-k/kube-carbon-footprint/pkg/server"
)

var (
	version      = "unknown"
	app          = "kube-carbon-footprint"
	port         = flag.Int("port", 9000, "server port to listen on (default: 9000)")
	cert         = flag.String("cert", "", "cert file path to enable TLS (disabled by default)")
	key          = flag.String("key", "", "private key file path to enable TLS (disabled by default)")
	promURL      = flag.String("prom", "http://prometheus:9090", "Prometheus URL")
	promToken    = flag.String("prom-token", "", "Bearer token for Prometheus")
	promInsecure = flag.Bool("prom-insecure", false, "TLS skip verify")
	promTimeout  = flag.Duration("prom-timeout", 10*time.Second, "Timeout for Prometheus client calls")
	logLevel     = flag.String("loglevel", "info", "log level")
	versionFlag  = flag.Bool("v", false, "print version")
	appVersion   = fmt.Sprintf("%s %s", app, version)
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		log.Errorf("Log level %s not recognized, using info", *logLevel)
		*logLevel = "info"
		lvl = log.InfoLevel
	}
	log.SetLevel(lvl)
	log.Infof("Starting %s at log level %s", appVersion, *logLevel)

	server.Start(server.Config{
		Port:           *port,
		CertFile:       *cert,
		PrivateKeyFile: *key,
	}, prometheus.Config{
		URL:                *promURL,
		Timeout:            *promTimeout,
		Token:              *promToken,
		InsecureSkipVerify: *promInsecure,
	})
}
