include .env
export

DOCKER-VERSION := 1.0.0
PROJECT-NAME := template-ulamm-backend-go
GITLAB-PNM-TOKEN := ${GITLAB_TOKEN}

.PHONY:	swag
swag:
	swag init -d . -g server.go -o ./docs

# Go Mock
mock: mock.usecase mock.repository

mock.usecase:
	mockgen -source=./pkg/usecase/common.go -destination=./mock/usecase/common.go

mock.repository:
	mockgen -source=./pkg/repository/common.go -destination=./mock/repository/common.go

mock.test:
	go test -coverprofile ./tmp/cover.out ./...

.PHONY:	docker-build
docker-build:
	docker build -t ${PROJECT-NAME}:${DOCKER-VERSION} .

# Proto
# Download .proto from mt-be-grpc repository, manually download one by one, should find the alternative to download all files
.PHONY: download-proto
download-proto:
	curl -o pkg/proto/common.proto --header "Private-Token: ${GITLAB-PNM-TOKEN}" "http://10.61.4.35/api/v4/projects/465/repository/files/proto%2Fcommon.proto/raw?ref=main"
	curl -o pkg/proto/master_service.proto --header "Private-Token: ${GITLAB-PNM-TOKEN}" "http://10.61.4.35/api/v4/projects/465/repository/files/proto%2Fmaster_service.proto/raw?ref=main"

.PHONY: inject-tag
inject-tag:
	protoc-go-inject-tag -input="pkg/grpc/**/*.pb.go"

.PHONY: generate-proto
generate-proto:
	protoc \
		--proto_path=pkg \
		--go_out=pkg \
		--go_grpc_out=pkg \
		pkg/proto/*.proto

.PHONY: proto
proto: download-proto generate-proto inject-tag
