# Hacking Outscale SDK

SDK itself is generated from Outscale's [OpenAPI description](https://github.com/outscale/osc-api) in [osc](osc/) folder.

Other contributions like examples and tests are welcome!

# Generate SDK

1. have some tools ready: GNU make, git, docker
2. edit `api_version` file and to the latest Outscale API version
3. edit `sdk_version` file and change it according to [semantic versioning](https://semver.org/) from the SDK perspective (not API)
4. launch sdk generation by running `make gen`
5. new sdk is now generated in `osc` folder
6. make sure that examples can still run using `make run-examples`

Under the hood:
- we get official Outscale yaml
- apply some patch if needed (.patch folder)
- run openapi-generator through docker to build osc folder

# Sending a Merge Request

Content in `osc` folder is generated at each release.
If you plan to make some change here, consider making a pull request in [openapi-generator project](https://github.com/OpenAPITools/openapi-generator/).

Otherwise:
- your merge request must be rebased on `next-release` branch
- be sure that tests still pass by running `make test`

# How to release

1. rebase on next-release branch
2. edit `api_version` file
3. `make gen` to update the sdk
4. `make test` and `make run-examples` and fix any issue
5. update `changelog.md` file
6. commit changes
7. tag version
8. push to next-release and master
