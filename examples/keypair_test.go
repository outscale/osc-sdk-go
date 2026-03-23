package examples_test

import (
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Create a keypair.
// 2. Validate the keypair data returned by the API.
// 3. Delete the keypair.
func TestKeypair(t *testing.T) {
	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	ctx := t.Context()

	keypairName := "osc-sdk-go-test-" + RandomString(10)
	faux := false
	resp, err := client.CreateKeypair(ctx, osc.CreateKeypairRequest{
		DryRun:      &faux,
		KeypairName: keypairName,
	}, options.WithRetryTimeout(time.Minute*10))
	require.NoError(t, err)
	require.NotNil(t, resp.Keypair)
	require.NotNil(t, resp.Keypair.KeypairId)

	t.Logf("Keypair created: %s", *resp.Keypair.KeypairId)

	_, err = client.DeleteKeypair(ctx, osc.DeleteKeypairRequest{
		KeypairId: resp.Keypair.KeypairId,
	})
	require.NoError(t, err)
}
