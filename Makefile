OUT_DIR = out/

help:
	@echo "=====[ Go Cast ]====="
	@echo "test ...... Run tests and show the code coverage."
	@echo "lint ...... Run linter to check gofmt, govet, etc."
	@echo "ci ........ Run common flow: vendor, test, lint, build."
.PHONY: help

# Run tests and show code coverage statistics.
test:
	mkdir -p $(OUT_DIR)
	go test -coverprofile $(OUT_DIR)coverage.out
	go tool cover -func $(OUT_DIR)coverage.out
.PHONY: test

# Check that all is ok.
# see: .golangci.yaml
# (!!!) INSTALL golangci to complete checking:
# go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint:
	golangci-lint run
.PHONY: lint

# CI automation to check local deps, tests/lint and build.
ci: test lint
.PHONY: ci