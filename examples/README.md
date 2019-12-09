# Run Examples

In order to run examples, you will need to get few modules:
```shell
go get github.com/antihax/optional
go get github.com/outscale-dev/osc-sdk-go/osc
```

You can then setup your credentials with environement variables:
```bash
export OSC_ACCESS_KEY=<ACCESS_KEY>
export OSC_SECRET_KEY=<SECRET_KEY>
```

Then you can `cd` in each example and hit `go run .`