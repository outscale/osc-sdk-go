VERSION=$(shell cat version)
all: osc

osc: osc-api/outscale.yaml
	rm -rf .sdk || true
	mkdir .sdk
	docker run -v $(PWD)/.sdk:/local -v $(PWD)/osc-api:/api --rm outscale/openapi-generator:osc-v1 generate -i /api/outscale.yaml -g go --package-name osc --git-user-id outscale-dev --git-repo-id osc-sdk-go/osc -o /local
	mv .sdk osc

osc-api/outscale.yaml:
	git clone https://github.com/outscale/osc-api.git && cd osc-api && git checkout -b $(VERSION) $(VERSION)

.PHONY: clean
clean:
	rm -rf .sdk osc-api osc || true

.PHONY: reuse
reuse:
	docker run -v $(PWD):/code --rm fsfe/reuse /bin/sh -c "cd /code && reuse lint"
