package osc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMutateBody(t *testing.T) {
	data := CreateVmsRequest{
		ImageId: "123",
	}

	require.NoError(t, data.SetClientToken())
	require.NotNil(t, data.ClientToken, "ClientToken should be set by mutateBody")
	require.Equal(t, "123", data.ImageId, "ImageId should not be modified")
}

func TestMutateBodyClientTokenAlreadySet(t *testing.T) {
	token := "azerty"

	data := CreateVmsRequest{
		ClientToken: &token,
		ImageId:     "123",
	}

	require.NoError(t, data.SetClientToken())
	require.Equal(t, &token, data.ClientToken, "ClientToken should not be modified by mutateBody if already set")
	require.Equal(t, "123", data.ImageId, "ImageId should not be modified")
}
