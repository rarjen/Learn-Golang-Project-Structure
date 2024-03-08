DOCKER-VERSION := 1.0.0
PROJECT-NAME := template-ulamm-backend-go

.PHONY:	docker-build
docker-build:
	docker build -t ${PROJECT-NAME}:${DOCKER-VERSION} .