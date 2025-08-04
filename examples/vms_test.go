package examples_test

import (
	"context"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/client"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
)

func TestReadVms(t *testing.T) {
	client, err := client.NewOapiClient()
	if err != nil {
		panic(err)
	}

	read, err := client.ReadVms(context.TODO(), osc.ReadVmsRequest{Filters: nil})
	if err != nil {
		panic(err)
	}

	println(*((*read.Vms)[0].VmId))
}
