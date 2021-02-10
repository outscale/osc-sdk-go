API_VERSION=$(shell cat api_version)
SDK_VERSION=$(shell cat sdk_version)
USER_ID=$(shell id -u)
GROUP_ID=$(shell id -g)
OPENAPI_GEN_VERSION=:v5.0.1

all: help

.PHONY: help
help:
	@echo "help:"
	@echo "- make gen  : regenerate SDK"
	@echo "- make test : run all tests"

.PHONY: gen
gen: clean v2

v2: osc-generate update-examples go-modules-update gofmt

.PHONY: osc-generate
osc-generate: osc-api/outscale.yaml
	rm -rf .sdk || true
	mkdir .sdk
	docker run -v $(PWD):/sdk --rm openapitools/openapi-generator-cli$(OPENAPI_GEN_VERSION) generate -i /sdk/osc-api/outscale.yaml -g go -c /sdk/gen.yml -o /sdk/.sdk --additional-properties=packageVersion=$(SDK_VERSION)
	docker run -v $(PWD):/sdk --rm openapitools/openapi-generator-cli$(OPENAPI_GEN_VERSION) chown -R $(USER_ID).$(GROUP_ID) /sdk/.sdk
	mv .sdk v2

osc-api/outscale.yaml:
	git clone https://github.com/outscale/osc-api.git && cd osc-api && git checkout -b $(API_VERSION) $(API_VERSION)

.PHONY: go-modules-update
go.sum-udpate:
	go get -u ./...
	go mod tidy
	cd v2 && go get -u ./... && go mod tidy

.PHONY: clean
clean:
	rm -rf .sdk osc-api v2 || true

.PHONY: update-examples
update-examples:
	@find $(PWD)/examples/ -type f -name "*.go" -exec ln -sr {} v2/ \;

.PHONY: test
test: reuse-test go-test
	@echo all tests OK

.PHONY: reuse-test
reuse-test:
	docker run --volume $(PWD):/data fsfe/reuse:0.11.1 lint

.PHONY: go-test
go-test:
	cd v2 && go test .

# try to regen, should not have any difference
.PHONY: regen-test
regen-test: gen
	git add v2
	git add examples
	git diff --cached -s --exit-code
	git diff -s --exit-code

.PHONY: gofmt
gofmt:
	@find $(PWD)/examples/ -type f -name "*.go" -exec gofmt -l {} \;
	@find $(PWD)/examples/ -type f -name "*.go" -exec gofmt -w {} \;
	@find $(PWD)/v2/ -type f -name "*.go" -exec gofmt -l {} \;
	@find $(PWD)/v2/ -type f -name "*.go" -exec gofmt -w {} \;

# Used by bot to auto-release
# GH_TOKEN and SSH_PRIVATE_KEY are needed
.PHONY: auto-release
auto-release: auto-release-cleanup osc-api-check release-check-duplicate release-build release-push release-pr
	@echo OK

.PHONY: auto-release-cleanup
auto-release-cleanup:
	rm -rf .auto-release-abort || true

.PHONY: osc-api-check
osc-api-check:
	bash .github/scripts/osc-api-check.sh

.PHONY: release-check-duplicate
release-check-duplicate:
	bash .github/scripts/release-check-duplicate.sh

.PHONY: release-build
release-build:
	bash .github/scripts/release-build.sh

.PHONY: release-push
release-push:
	bash .github/scripts/release-push.sh

.PHONY: release-pr
release-pr:
	bash .github/scripts/release-pr.sh
