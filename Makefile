# Parameters to compile and run application
GOOS?=linux
GOARCH?=amd64

PROJECT?=github.com/ekrukov/nsk2
BUILD_PATH?=cmd/nsk2
APP?=nsk2

PORT?=8080

# Current version and commit
RELEASE?=0.0.4
#COMMIT?=$(shell git rev-parse --short HEAD)
#BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Parameters to push images and release app to Kubernetes or try it with Docker
REGISTRY?=docker.io/webdeva
NAMESPACE?=ekrukov
CONTAINER_NAME?=${NAMESPACE}-${APP}
CONTAINER_IMAGE?=${REGISTRY}/${CONTAINER_NAME}
VALUES?=values-stable

build:
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: build
	docker push $(CONTAINER_IMAGE):$(RELEASE)

