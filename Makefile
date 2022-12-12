GO=go
export GOPRIVATE=github.com/woodwo/*
TARGET_DIR?=$(PWD)/.build
M=$(shell printf "\033[34;1m>>\033[0m")
PROTO_DIR=grpc/proto
PROTO=$(PROTO_DIR)/calculus.proto
PACKAGE=github.com/woodwo/calculus


.PHONY: build
build: 
	$(info $(M) building service...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(TARGET_DIR)/service ./cmd/*.go

.PHONY: run
run: build 
	$(TARGET_DIR)/service

.PHONY: generate
generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $(PROTO)

.PHONY: generate-mock
generate-mock: generate
	mockgen $(PACKAGE)/$(PROTO_DIR) CalculusClient >> $(PROTO_DIR)/mock/mock_calculus.go

.PHONY: evans
evans:
	evans $(PROTO) -p 8080

.PHONY: tidy
tidy:
	go mod tidy
%:
	@:

