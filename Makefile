.PHONY: .install-lint
.install-lint:
	$(info Installing linter)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: generate
generate: .bin-deps fast-generate

.PHONY: fast-generate
fast-generate: .generate .copy-swagger

.PHONY: run
run:
	go run ./cmd/crnt-data-manager/main.go