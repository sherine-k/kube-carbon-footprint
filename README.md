# kube-carbon-footprint

WORK IN PROGRESS
(early stage)

## Running

Using Prometheus in OpenShift, pass the Prometheus address and admin token as program arguments:

E.g.:

```bash
./kube-carbon-footprint -prom=https://prometheus.mycluster.openshift.com -prom-insecure=true -prom-token="sha256~XXXXXXXX" -kube=/home/$USER/.kube/config
```

Endpoints:

- CPU usage: `/api/metrics/cpu`
- Power consumption per instance type: `/api/dataset/instancetype/{instanceType}` (e.g. `/api/dataset/instancetype/a1.medium`)
- Stats per region: `/api/dataset/region/{region}` (e.g. `/api/dataset/region/us-east-1`)
- Datacenter info for pod: `api/datacenter/namespace/{namespace}/pod/{pod}` (e.g. `api/datacenter/namespace/default/pod/my-pod`)

## Running in cluster
```bash
kubectl apply -f test_resources/kcf-role.yaml
kubectl apply -f test_resources/kcf-deployment.yaml
```
PS: what we're missing is a serviceaccount that has permissions to prometheus and the correct prometheus URL

## License and credits

This software is published under the Apache v2 license (see [LICENSE file](./LICENSE)).
With the exception of the dataset, which comes from the AWS EC2 Carbon Footprint Dataset compiled by Benjamin Davy (Teads): https://docs.google.com/spreadsheets/d/1DqYgQnEDLQVQm5acMAhLgHLD8xXCG9BIrk-_Nv6jF3k and is is published under CC BY 4.0 license (see [data/LICENSE file](./data/LICENSE)).
