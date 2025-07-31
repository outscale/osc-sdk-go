all: help

.PHONY: help
help:
	@echo "help:"
	@echo "- make gen  : regenerate SDK"
	@echo "- make test : run all tests"

go.sum: go.mod
	go mod tidy

tmpl_srcs:=$(wildcard templates/*.tmpl)
internal/oks/client.gen.go: $(tmpl_srcs) internal/oks/api.yaml internal/oks/cfg.yaml internal/oks/generate.go go.sum
	go generate ./internal/oks/

internal/osc/client.gen.go: $(tmpl_srcs) internal/osc/api.yaml internal/osc/cfg.yaml internal/osc/generate.go go.sum
	go generate ./internal/osc/

.PHONY: gen
gen: internal/oks/client.gen.go internal/osc/client.gen.go

.PHONY: test
test: gen
	go test -v ./examples
