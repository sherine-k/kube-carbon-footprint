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

- CPU usage: `/api/metrics/cpu`
- Power consumption per instance type: `/api/dataset/instancetype/{instanceType}` (e.g. `/api/dataset/instancetype/a1.medium`)

## License and credits

This software is published under the Apache v2 license (see [LICENSE file](./LICENSE)).
With the exception of the dataset, which comes from the AWS EC2 Carbon Footprint Dataset compiled by Benjamin Davy (Teads): https://docs.google.com/spreadsheets/d/1DqYgQnEDLQVQm5acMAhLgHLD8xXCG9BIrk-_Nv6jF3k and is is published under CC BY 4.0 license (see [data/LICENSE file](./data/LICENSE)).
