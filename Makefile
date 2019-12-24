VERSION=$(shell cat version)
all: osc

osc: osc-api/outscale.yaml
	rm -rf .sdk || true
	mkdir .sdk
	docker run -v $(PWD):/sdk --rm outscale/openapi-generator:osc-v2 generate -i /sdk/osc-api/outscale.yaml -g go -c /sdk/gen.yml -o /sdk/.sdk
	mv .sdk osc

osc-api/outscale.yaml:
	git clone https://github.com/outscale/osc-api.git && cd osc-api && git checkout -b $(VERSION) $(VERSION)

.PHONY: clean
clean:
	rm -rf .sdk osc-api osc || true


.PHONY: test
test: build-examples reuse
	@echo all tests OK

.PHONY: reuse
reuse:
	docker run -v $(PWD):/code --rm fsfe/reuse /bin/sh -c "cd /code && reuse lint"

.PHONY: build-examples
build-examples: examples
	find ./examples -type d -depth 1 -exec go build -o /dev/null {} \;

.PHONY: gofmt
gofmt:
	@find $(PWD)/examples/ -type f -name *.go -exec gofmt -l {} \;
	@find $(PWD)/examples/ -type f -name *.go -exec gofmt -w {} \;
