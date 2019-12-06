# Hacking Outscale SDK

Let's talk about how to generate a new SDK version based on [OpenAPI](https://www.openapis.org/) description of [Outscale API](https://github.com/outscale/osc-api).

# Generate SDK

The generator builds last Outscale API version. To do so:

1. have some tools ready: make, git, docker
2. remove current generated sdk using `make clean`
3. launch sdk generation by running `make`
4. new sdk is now generated in `osc` folder

Under the hood, a [patched version](https://github.com/outscale-dev/openapi-generator/tree/go-awsv4-wip) of [OpenAPI-Genetator](https://github.com/OpenAPITools/openapi-generator) is used.
Patched version adds support for [AWSv4 signature](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).

# Before sending a merge request

- Examples: check that they are still running
- Licensing: run `make reuse` to verify if the project is still [REUSE](https://reuse.software/) compliant.
