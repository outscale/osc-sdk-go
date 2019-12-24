# Hacking Outscale SDK

SDK itself is generated from Outscale's [OpenAPI description](https://github.com/outscale/osc-api) in [osc](osc/) folder.

Other contributions like examples and tests are welcome!

# Generate SDK

1. have some tools ready: GNU make, git, docker
2. edit `version` file and to the latest version
3. clean old sdk using `make clean`
4. launch sdk generation by running `make`
5. new sdk is now generated in `osc` folder

Under the hood, a [patched version](https://github.com/outscale-dev/openapi-generator/tree/osc-v2) of [OpenAPI-Genetator](https://github.com/OpenAPITools/openapi-generator) is used.
Patched version adds support for [AWSv4 signature](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).

# Sending a Merge Request

- your merge request must be rebased on `next-release` branch
- be sure that tests still pass by running `make test`
