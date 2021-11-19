IMAGE ?= quay.io/skhoury/kube-carbon-footprint
VERSION ?= v0.0.1
GOLANGCI_LINT_VERSION = v1.42.1
COVERPROFILE = coverage.out

ifeq (,$(shell which podman 2>/dev/null))
OCI_BIN ?= docker
else
OCI_BIN ?= podman
endif

.PHONY: prereqs
prereqs:
	@echo "### Test if prerequisites are met, and installing missing dependencies"
	test -f $(go env GOPATH)/bin/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

.PHONY: vendors
vendors:
	@echo "### Checking vendors"
	go mod tidy && go mod vendor

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint: prereqs
	@echo "### Linting code"
	golangci-lint run ./...

.PHONY: test
test:
	@echo "### Testing"
	go test ./... -coverprofile ${COVERPROFILE}

.PHONY: verify
verify: lint test

.PHONY: build
build:
	@echo "### Building"
	go build -mod vendor -o kube-carbon-footprint cmd/kube-carbon-footprint.go

.PHONY: image
image:
	@echo "### Building image with ${OCI_BIN}"
	$(OCI_BIN) build --build-arg VERSION="$(VERSION)" -t $(IMAGE):$(VERSION) .

.PHONY: push
push:
	$(OCI_BIN) push $(IMAGE):$(VERSION)
