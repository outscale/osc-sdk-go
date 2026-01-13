package examples_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/outscale/osc-sdk-go/v3/pkg/utils"
)

func TestKeypair(t *testing.T) {
	userProfile, err := profile.New()
	if err != nil {
		panic(err)
	}

	client, err := osc.NewClient(userProfile, utils.WithLogging(&testingLogger{t}))
	if err != nil {
		panic(err)
	}

	ctx := t.Context()

	keypairName := "osc-sdk-go-example"
	faux := false
	resp, err := client.CreateKeypair(ctx, osc.CreateKeypairRequest{
		DryRun:      &faux,
		KeypairName: keypairName,
	})
	if err != nil {
		panic(err)
	}

	t.Logf("Keypair created: %s", *resp.Keypair.KeypairId)

	_, err = client.DeleteKeypair(ctx, osc.DeleteKeypairRequest{
		KeypairId: resp.Keypair.KeypairId,
	})
	if err != nil {
		panic(err)
	}
}
