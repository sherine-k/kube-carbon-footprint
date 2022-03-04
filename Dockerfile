FROM golang:1.17 as builder
ARG VERSION="1.17"

WORKDIR /opt/app-root
COPY . .

RUN go build -ldflags "-X main.version=${VERSION}" -mod vendor -o kube-carbon-footprint cmd/kube-carbon-footprint.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.5

COPY --from=builder /opt/app-root/kube-carbon-footprint ./

ENTRYPOINT ["sh", "-c", "./kube-carbon-footprint --prom-insecure=true --prom-token-file=/var/run/secrets/kubernetes.io/serviceaccount/token --prom=https://prometheus-k8s.openshift-monitoring.apps-crc.testing:9091/" ]
