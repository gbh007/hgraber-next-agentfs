.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/controller/api/internal/server --package server --clean internal/controller/api/internal/agent.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: build-arm64
build-arm64: create_build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./_build/hgraber-agentfs-arm64 ./cmd/server

.PHONY: run
run: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	./_build/server --debug --addr 127.0.0.1:8081 --token agent-token --export-path ./.hidden/export --data-path ./.hidden/filedata --trace-endpoint http://localhost:4318/v1/traces