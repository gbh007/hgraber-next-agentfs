.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen@v1.2.1 --target open_api/agentAPI -package agentAPI --clean open_api/agent.yaml
	go run github.com/ogen-go/ogen/cmd/ogen@v1.2.1 --target open_api/serverAPI -package serverAPI --clean open_api/server.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: build-arm64
build-arm64: create_build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./_build/hgraber-agentfs-arm64 ./cmd/server

.PHONY: run-example
run-example: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	./_build/server --config config-example.yaml

.PHONY: run
run: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	./_build/server