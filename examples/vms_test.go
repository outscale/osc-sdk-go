package examples_test

import (
	"context"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/internal/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/client"
)

func TestReadVms(t *testing.T) {
	builder, err := client.Builder("", "")
	if err != nil {
		panic(err)
	}

	client, err := builder.OApi()
	if err != nil {
		panic(err)
	}

	read, err := client.ReadVms(context.TODO(), osc.ReadVmsRequest{Filters: nil})
	if err != nil {
		panic(err)
	}

	println(*((*read.Vms)[0].VmId))
}
