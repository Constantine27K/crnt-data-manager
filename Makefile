SRC_DIR_RELATIVE=./api
DST_DIR_RELATIVE=./pkg

PROTO_FILES=$(shell find $(SRC_DIR_RELATIVE) \
	-not \( -path ./api/google -prune \) \
	-not \( -path ./api/protoc* -prune \) \
	-name '*.proto')

SWAGGER_FILES=$(shell find $(DST_DIR_RELATIVE) \
                     	-name '*.swagger.json*')

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

.PHONY: .copy-swagger
.copy-swagger:
	cp $(SWAGGER_FILES) ./swagger

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

.PHONY: buf-install
buf-install:
	brew install bufbuild/buf/buf

.PHONY: .install-lint
.install-lint:
	$(info Installing linter)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: generate-protoc
generate-protoc: .bin-deps .generate .copy-swagger

.PHONY: generate
generate:
	buf mod update
	buf generate
	make .copy-swagger
	go mod tidy

.PHONY: build
build:
	go build -o ./bin/crnt-data-manager ./cmd/crnt-data-manager/main.go

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: .run
.run:
	go run ./cmd/crnt-data-manager/main.go

.PHONY: .run-env
.run-env:
	cd ./scripts/dev && make up-local-db

.PHONY: run
run: build .run

.PHONY: run-all
run-all: .run-env run

PHONY: test-unit
test-unit:
	go test ./internal/pkg/...

PHONY: test-integration
test-integration:
	$(info Running inegration tests...)
	go test -v -count=1 -tags integration ./integration_tests/suites/...

PHONY: test-all
test-all: test-unit test-integration
