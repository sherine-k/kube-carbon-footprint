FROM registry.access.redhat.com/ubi8/go-toolset:1.15.14-14 as builder
ARG VERSION=""

WORKDIR /opt/app-root
COPY . .

RUN go build -ldflags "-X main.version=${VERSION}" -mod vendor -o kube-carbon-footprint cmd/kube-carbon-footprint.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4-210

COPY --from=builder /opt/app-root/kube-carbon-footprint ./

ENTRYPOINT ["./kube-carbon-footprint"]
