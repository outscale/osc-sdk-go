all: help

.PHONY: help
help:
	@echo "help:"
	@echo "- make gen  : regenerate SDK"
	@echo "- make test : run all tests"

go.sum: go.mod
	go mod tidy

tmpl_srcs:=$(wildcard templates/*.tmpl)
pkg/oks/client.gen.go: $(tmpl_srcs) pkg/oks/api.yaml pkg/oks/cfg.yaml pkg/oks/patch.yaml pkg/oks/generate.go go.sum
	go generate ./pkg/oks/

pkg/osc/client.gen.go: $(tmpl_srcs) pkg/osc/api.yaml pkg/osc/cfg.yaml pkg/osc/patch.yaml pkg/osc/generate.go go.sum
	go generate ./pkg/osc/

.PHONY: gen
gen: pkg/oks/client.gen.go pkg/osc/client.gen.go
	go mod tidy
ifdef SDK_VERSION
	sed -i 's/Version = "[^"]\+"/Version = "$(SDK_VERSION)"/' pkg/version/version.go
endif
ifdef OAPI_VERSION
	sed -i 's/Version = "[^"]\+"/Version = "$(OAPI_VERSION)"/' pkg/osc/version.go
endif
ifdef OKS_VERSION
	sed -i 's/Version = "[^"]\+"/Version = "$(OKS_VERSION)"/' pkg/oks/version.go
endif

.PHONY: test
test: gen
	go test -v ./...
