all: help

.PHONY: help
help:
	@echo "help:"
	@echo "- make gen  : regenerate SDK"
	@echo "- make test : run all tests"

go.sum: go.mod
	go mod tidy

tmpl_srcs:=$(wildcard templates/*.tmpl)
internal/oks/client.gen.go: $(tmpl_srcs) pkg/oks/api.yaml pkg/oks/cfg.yaml pkg/oks/generate.go go.sum
	go generate ./pkg/oks/

internal/osc/client.gen.go: $(tmpl_srcs) pkg/osc/api.yaml pkg/osc/cfg.yaml pkgs/osc/generate.go go.sum
	go generate ./pkg/osc/

.PHONY: gen
gen: pkg/oks/client.gen.go pkg/osc/client.gen.go

.PHONY: test
test: gen
	go test -v ./examples
