# Examples

In order to run examples, you will need to setup your credentials (eu-west-2 region) with environement variables and run `go test`:

```bash
export OSC_ACCESS_KEY=<ACCESS_KEY>
export OSC_SECRET_KEY=<SECRET_KEY>
cd v2/
go test -run Volume
go test -run Keypair
...
```
