package examples_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Create an EIM user.
// 2. Read the EIM user by ID.
// 3. Validate the returned user data.
// 4. Delete the EIM user.
func TestEIMUser(t *testing.T) {
	ctx := t.Context()

	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	userName := "osc-sdk-go-test-" + RandomString(10)
	userPath := "/"
	userEmail := userName + "@example.com"

	createResp, err := client.CreateUser(ctx, osc.CreateUserRequest{
		Path:      &userPath,
		UserEmail: &userEmail,
		UserName:  userName,
	})
	require.NoError(t, err)
	require.NotNil(t, createResp.User)
	require.NotNil(t, createResp.User.UserId)
	require.NotNil(t, createResp.User.UserName)

	userID := *createResp.User.UserId
	require.NotEmpty(t, userID)

	t.Logf("Created EIM user: %s", userID)

	defer func() {
		_, _ = client.DeleteUser(ctx, osc.DeleteUserRequest{
			UserName: userName,
		})
	}()

	readResp, err := client.ReadUsers(ctx, osc.ReadUsersRequest{
		Filters: &osc.FiltersUsers{
			UserIds: &[]string{userID},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, readResp.Users)
	require.Len(t, *readResp.Users, 1)

	user := (*readResp.Users)[0]
	require.NotNil(t, user.UserId)
	require.Equal(t, userID, *user.UserId)
	require.NotNil(t, user.UserName)
	require.Equal(t, userName, *user.UserName)
	require.NotNil(t, user.UserEmail)
	require.Equal(t, userEmail, *user.UserEmail)

	t.Logf("Successfully read EIM user: %s", userID)

	deleteResp, err := client.DeleteUser(ctx, osc.DeleteUserRequest{
		UserName: userName,
	})
	require.NoError(t, err)
	require.NotNil(t, deleteResp.ResponseContext)
	require.NotNil(t, deleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted EIM user: %s", userID)
}
