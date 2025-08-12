package examples_test

import (
	"context"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/client"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
)

func TestKeypair(t *testing.T) {
	client, err := client.NewOapiClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	keypairName := "osc-sdk-go-example"
	keypair, err := client.CreateKeypair(ctx, osc.CreateKeypairRequest{
		KeypairName: keypairName,
	})
	if err != nil {
		panic(err)
	}

	t.Logf("Keypair created: %v", keypair)

	_, err = client.DeleteKeypair(ctx, osc.DeleteKeypairRequest{
		KeypairId: keypair.Keypair.KeypairId,
	})
	if err != nil {
		panic(err)
	}
}
