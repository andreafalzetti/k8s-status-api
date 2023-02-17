DOCKER_REGISTRY ?= registry.cloud.okteto.net
DOCKER_IMAGE ?= k8s-status-api
# DOCKER_IMAGE ?= afalzetti/k8s-status-api

.PHONY: start
start:
	go run main.go

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-publish
docker-publish:
	docker tag ${DOCKER_IMAGE} ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${TAG}
	docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${TAG}

