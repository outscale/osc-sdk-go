package examples_test

import (
	"context"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
)

func TestKeypair(t *testing.T) {
	userProfile, err := profile.NewProfileFromStrandardConfiguration("", "")
	if err != nil {
		panic(err)
	}

	client, err := osc.NewClient(userProfile)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	keypairName := "osc-sdk-go-example"
	faux := false
	keypair, err := client.CreateKeypair(ctx, osc.CreateKeypairRequest{
		DryRun:      &faux,
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
