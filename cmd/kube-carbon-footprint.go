package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/sherine-k/kube-carbon-footprint/pkg/prometheus"
)

var (
	version     = "unknown"
	app         = "kube-carbon-footprint"
	logLevel    = flag.String("loglevel", "info", "log level")
	versionFlag = flag.Bool("v", false, "print version")
	appVersion  = fmt.Sprintf("%s %s", app, version)
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

	prometheus.Query()
}
