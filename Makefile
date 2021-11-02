GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

.PHONY: init
# init env
init:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/apps/protoc-gen-openapi@latest
	go install github.com/google/gnostic@latest
	go install entgo.io/ent/cmd/ent@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
		   --openapi_out=paths=source_relative:. \
	       $(API_PROTO_FILES)
	sed -i -e 's/title: EquipmentService/title: Inventory/g' openapi.yaml

.PHONY: build
# build binary
build:
	mkdir -p bin/
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/inventory
	go build -o ./bin/ ./cmd/inventoryctl

.PHONY: migrate
migrate:
	go run cmd/inventoryctl/main.go migrate

.PHONY: generate
# run go generators
generate:
	go generate ./...

.PHONY: proto
proto:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           $(INTERNAL_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;
	make proto;
	make generate;

.PHONY: test
# run test suite
test:
	go test -v ./... -cover

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
