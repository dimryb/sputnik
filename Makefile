#Makefile
NAME := "sputnik"

GIT_HASH := $(shell git log --format="%h" -n 1)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	BIN := ./bin/$(NAME)
    DATE_CMD = date -u +'%Y-%m-%dT%H:%M:%S'
    GO_PATH := $(shell go env GOPATH)
else #windows
	BIN := ./bin/$(NAME).exe
    DATE_CMD = powershell.exe -Command "Get-Date -Format 'yyyy-MM-ddTHH:mm:ss'"
    GO_PATH := $(shell go env GOPATH | tr '\\' '/')
endif

LDFLAGS := -X main.release="develop" \
    -X main.buildDate=$(shell $(DATE_CMD)) \
    -X main.gitHash=$(GIT_HASH)

.PHONY: build
build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/app

.PHONY: run
run: build
	$(BIN) -config ./configs/config.yaml

.PHONY: version
version: build
	$(BIN) version

.PHONY: test
test:
	go test -race ./internal/...

.PHONY: install-lint-deps
install-lint-deps:
	(which golangci-lint > /dev/null) || \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	sh -s -- -b $(GO_PATH)/bin v1.64.8

.PHONY: lint
lint: install-lint-deps
	golangci-lint run --config golangci.yml ./...

.PHONY: build-service
build-service:
	docker build \
		--build-arg LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f $(DOCKERFILE_PATH) .

.PHONY: build-sputnik-img
build-sputnik-img:
	$(MAKE) build-service \
		DOCKER_IMG=sputnik:develop \
		DOCKERFILE_PATH=build/Dockerfile


.PHONY: run-docker
run-docker:
	docker run --rm \
		-v $(shell pwd)/configs:/etc/sputnik \
		sputnik:develop /opt/sputnik/sputnik-app -config /etc/sputnik/config.yaml

.PHONY: build-ghcr
build-sputnik:
	$(MAKE) build-service \
		DOCKER_IMG=ghcr.io/dimryb/sputnik:latest \
		DOCKERFILE_PATH=build/Dockerfile

.PHONY: push-ghcr
push-sputnik: build-ghcr
	docker push ghcr.io/dimryb/sputnik:latest

.PHONY: run-ghcr
run-sputnik:
	docker run --rm \
		-v $(shell pwd)/configs:/etc/sputnik \
		ghcr.io/dimryb/sputnik:latest /opt/sputnik/sputnik-app -config /etc/sputnik/config.yaml

.PHONY: update-sputnik
update-sputnik:
	docker pull ghcr.io/dimryb/sputnik:latest
