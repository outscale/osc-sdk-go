# Hacking Outscale SDK

SDK itself is generated from Outscale's [OpenAPI description](https://github.com/outscale/osc-api) in [v2](v2/) folder using OpenAPI Genetator.

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
5. new sdk is now generated in `v2` folder

Under the hood:
- we get official Outscale yaml
- run openapi-generator through docker to build osc folder

# Sending a Merge Request

Content in `v2` folder is generated at each release.
If you plan to make some change here, consider making a pull request in [openapi-generator project](https://github.com/OpenAPITools/openapi-generator/).

Otherwise:
- your merge request must be rebased on the corresponding major version branch (v1, v2, ...)
- be sure that tests still pass by running `make test`

# How to release

For each major version (v1, v2, ...):
1. rebase on corresponding major version branch
2. update `api_version` to the last Outscale API version
3. update `sdk_version` following [semantic versioning](https://semver.org/) logic.
4. eventually update go.mod file
5. `make gen` to re-build the sdk
6. `make test` and fix any issue
7. update `changelog.md` file
8. commit changes
9. tag version
10. push to corresponding branch

Note that CI should automatically detect new release on osc-api, update the SDK and push a new version.
