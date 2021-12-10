# kube-carbon-footprint

WORK IN PROGRESS
(early stage)

## Running

Using Prometheus in OpenShift, pass the Prometheus address and admin token as program arguments:

E.g.:

```bash
./kube-carbon-footprint -prom=https://prometheus.mycluster.openshift.com -prom-insecure=true -prom-token="sha256~XXXXXXXX"
```

Endpoints:

- CPU usage: http://localhost:9000/api/metrics/cpu
