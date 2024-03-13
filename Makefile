DOCKER-VERSION := 1.0.0
PROJECT-NAME := template-ulamm-backend-go
PROJECT-LOCATION := /d/PNM/projects/go-project/${PROJECT-NAME}
PROJECT-LOCATION-WINDOWS := D:/PNM/projects/go-project/template-ulamm-backend-go

.PHONY:	docker-build
docker-build:
	docker build -t ${PROJECT-NAME}:${DOCKER-VERSION} .

# run this command to update the swagger
swag:
	cd ~/go/bin; swag init -d ${PROJECT-LOCATION-WINDOWS} -o ${PROJECT-LOCATION-WINDOWS}/docs	