SRC_DIR_RELATIVE=./api
DST_DIR_RELATIVE=./pkg

TASK_PROTO=./api/task/task.proto
STATUS_PROTO=./api/status/status.proto
SPRINT_PROTO=./api/sprint/sprint.proto
PROTO_FILES=$(STATUS_PROTO) $(SPRINT_PROTO) $(TASK_PROTO)

LOCAL_BIN=$(CURDIR)/bin

.PHONY: .generate
.generate: .gen-go-pb .gen-go-grpc .gen-gateway .gen-openapi

.PHONY: .gen-go-pb
.gen-go-pb:
	protoc -I $(SRC_DIR_RELATIVE) \
      --go-grpc_out $(DST_DIR_RELATIVE) --go-grpc_opt paths=source_relative --go-grpc_opt require_unimplemented_servers=false\
      $(PROTO_FILES)

.PHONY: .gen-go-grpc
.gen-go-grpc:
	protoc -I $(SRC_DIR_RELATIVE) \
      --go_out $(DST_DIR_RELATIVE) --go_opt paths=source_relative \
      $(PROTO_FILES)

.PHONY: .gen-gateway
.gen-gateway:
	protoc -I $(SRC_DIR_RELATIVE) --grpc-gateway_out $(DST_DIR_RELATIVE) \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        $(PROTO_FILES)

.PHONY: .gen-openapi
.gen-openapi:
	protoc -I $(SRC_DIR_RELATIVE) --openapiv2_out $(DST_DIR_RELATIVE) \
        --openapiv2_opt logtostderr=true \
         $(PROTO_FILES)

.PHONY: .gen-vtproto
.gen-vtproto:
	protoc -I $(SRC_DIR_RELATIVE) --go-vtproto_out $(DST_DIR_RELATIVE) \
        --go-vtproto_opt features=marshal+unmarshal+size+pool --go-vtproto_opt paths=source_relative\
         $(PROTO_FILES)

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies)
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.2
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.3.0
	go install github.com/mitchellh/gox@v1.0.1
	go install golang.org/x/tools/cmd/goimports@v0.1.9
	go install github.com/bufbuild/buf/cmd/buf

.PHONY: .install-lint
.install-lint:
	$(info Installing linter)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: generate
generate: .bin-deps .generate

.PHONY: fast-generate
fast-generate: .generate

.PHONY: run
run:
	go run ./cmd/crnt-data-manager/main.go --local-config-enabled --bind-localhost