# Examples

The files in this folder are integration-style SDK example tests for `osc-sdk-go/v3`.

## Prerequisites

Before running the examples, make sure your OUTSCALE credentials are configured in a way that `profile.New()` can load.

You can export credentials as environment variables:

```bash
export OSC_ACCESS_KEY=<ACCESS_KEY>
export OSC_SECRET_KEY=<SECRET_KEY>
```

Or configure them in `~/.osc/config.json` and optionally select a profile with `OSC_PROFILE`.

## Run All Examples

From the repository root:

```bash
go test ./examples
```

To skip OKS example tests:

```bash
SKIP_OKS_TESTS=true go test ./examples
```

## Run One Example

```bash
go test ./examples -run TestSecurityGroup
go test ./examples -run TestKeypair
go test ./examples -run TestVm
go test ./examples -run TestLoadBalancerBackend
go test ./examples -run TestProject
```

## Notes

- These tests call real APIs.
- Some tests create and delete cloud resources.
- Use valid credentials and a suitable test environment before running them.
