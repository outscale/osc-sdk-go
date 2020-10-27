# Hacking Outscale SDK

SDK itself is generated from Outscale's [OpenAPI description](https://github.com/outscale/osc-api) using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator/).

Other contributions like examples and tests are welcome!

# Versioning

This SDK follows [semantic versioning](https://semver.org/) from the SDK perspective (not API).
Some events may trigger a major (breaking) version of the SDK:
1. OpenAPI generator introduce a new major version
2. Outscale introduce a new major version of its API

When OpenAPI generator introduce a breaking change, SDK can be generated in several versions (see corresponding branches)

# Generate SDK

1. have some tools ready: GNU make, git, docker
2. edit `api_version` file and to the latest Outscale API version
3. edit `sdk_version` file and change it according to [semantic versioning](https://semver.org/)
4. launch sdk generation by running `make gen`
5. new sdk is now generated in `osc` folder

Under the hood:
- we get official Outscale yaml
- apply some patch if needed (.patch folder)
- run openapi-generator through docker to build the SDK

# Sending a Merge Request

As SDK is generated at each release. If you plan to make some change inside, consider making a pull request in [openapi-generator project](https://github.com/OpenAPITools/openapi-generator/).

Otherwise:
- your merge request must be rebased on the corresponding major version branch (v1, v2, ...)
- be sure that tests still pass by running `make test` and `make run-examples`

# How to release

1. rebase on corresponding major version branch (v1, v2, ...)
2. edit `sdk_version` and `api_version` file
3. `make gen` to update the sdk
4. `make test` and `make run-examples` and fix any issue
5. update `changelog.md` file
6. commit changes
7. tag version
8. push to corresponding branch
