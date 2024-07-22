.PHONY: generate
generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/controller/api/internal/server --package server --clean internal/controller/api/internal/agent.yaml

create_build_dir:
	mkdir -p ./_build

.PHONY: run
run: create_build_dir
	go build -trimpath -o ./_build/server  ./cmd/server

	./_build/server --debug