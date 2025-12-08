IFACEMAKER_VERSION?=v1.3.0

all: help

.PHONY: help
help:
	@echo "help:"
	@echo "- make gen  		: regenerate SDK"
	@echo "- make test		: run all tests"
	@echo "- make ifacemaker: install ifacemaker"

go.sum: go.mod
	go mod tidy

tmpl_srcs:=$(wildcard templates/*.tmpl)
pkg/oks/client.gen.go: $(tmpl_srcs) pkg/oks/api.yaml pkg/oks/cfg.yaml pkg/oks/patch.yaml pkg/oks/generate.go go.sum
	go generate ./pkg/oks/

pkg/oks/interface.go: pkg/oks/client.gen.go
	go generate ./pkg/oks/

pkg/osc/client.gen.go: $(tmpl_srcs) pkg/osc/api.yaml pkg/osc/cfg.yaml pkg/osc/patch.yaml pkg/osc/generate.go go.sum
	go generate ./pkg/osc/

pkg/osc/interface.go: pkg/osc/client.gen.go
	go generate ./pkg/osc/

.PHONY: gen
gen: pkg/oks/client.gen.go pkg/oks/interface.go pkg/osc/client.gen.go pkg/osc/interface.go
	go mod tidy

.PHONY: test
test: gen
	go test -v ./examples

.PHONY: ifacemaker
ifacemaker:
	go install github.com/vburenin/ifacemaker@$(IFACEMAKER_VERSION)
