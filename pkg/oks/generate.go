package oks

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml api.yaml
//go:generate ifacemaker -f client.gen.go -s Client -i Interface -p oks -o interface.go

var _ Interface = (*Client)(nil)
