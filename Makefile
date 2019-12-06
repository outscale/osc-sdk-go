all: osc

osc: osc-api/outscale.yaml
	rm -rf .sdk || true
	mkdir .sdk
	docker run -v ${PWD}/.sdk:/local -v ${PWD}/osc-api:/api --rm outscale/openapi-generator:osc-v0 generate -i /api/outscale.yaml -g go --package-name osc --git-user-id outscale-dev --git-repo-id osc-sdk-go/osc -o /local
	mv .sdk osc

osc-api/outscale.yaml:
	git clone https://github.com/outscale/osc-api.git

.PHONY: clean
clean:
	rm -rf .sdk osc-api osc || true

.PHONY: reuse
reuse:
	docker run -v ${PWD}:/code --rm fsfe/reuse /bin/sh -c "cd /code && reuse lint"
