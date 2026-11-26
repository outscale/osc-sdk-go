package profile_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigFile_Save(t *testing.T) {
	src := writeConfig(t, map[string]any{
		"foo": map[string]any{"default": true, "access_key": "ak", "secret_key": "sk"},
	})

	t.Run("A new profile file can be created in an existing directory", func(t *testing.T) {
		cf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(src),
		})
		require.NoError(t, err)

		dst := filepath.Join(t.TempDir(), "existing.json")
		cf.FilePath = &dst
		require.NoError(t, cf.Save())

		scf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(dst),
		})
		require.NoError(t, err)

		got, err := scf.DefaultProfile()
		require.NoError(t, err)
		assert.Equal(t, "foo", got.Name)
		assert.Equal(t, "ak", got.AccessKey)
		assert.Equal(t, "sk", got.SecretKey)
	})

	t.Run("A new profile file can be created in a missing directory", func(t *testing.T) {
		cf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(src),
		})
		require.NoError(t, err)

		dst := filepath.Join(t.TempDir(), "foo", "bar", "missing.json")
		cf.FilePath = &dst
		require.NoError(t, cf.Save())

		scf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(dst),
		})
		require.NoError(t, err)

		got, err := scf.DefaultProfile()
		require.NoError(t, err)
		assert.Equal(t, "foo", got.Name)
		assert.Equal(t, "ak", got.AccessKey)
		assert.Equal(t, "sk", got.SecretKey)
	})
}

func TestConfigFile_ProfileManagement(t *testing.T) {
	path := writeConfig(t, map[string]any{
		"z":       map[string]any{"default": true, "region": "z-region"},
		"default": map[string]any{"region": "fallback-region"},
	})
	t.Setenv(profile.ProfileNameEnvName, "z")
	o := profile.Options{FilePath: &path}
	cf, err := profile.LoadConfigFile(o)
	require.NoError(t, err)

	require.ErrorIs(t, cf.SetDefault("missing"), profile.ErrProfileNotFound)
	require.NoError(t, cf.Save())
	cf, err = profile.LoadConfigFile(o)
	require.NoError(t, err)
	profiles, err := cf.ListProfiles()
	require.NoError(t, err)
	assert.Equal(t, []profile.Profile{
		{Name: "default", Fields: profile.Fields{Region: "fallback-region"}},
		{Name: "z", Default: true, Fields: profile.Fields{Region: "z-region"}},
	}, profiles)

	require.NoError(t, cf.AddProfile(profile.Profile{
		Name: "a", Default: true, Fields: profile.Fields{Region: "a-region"},
	}))
	require.NoError(t, cf.Save())
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.JSONEq(t, `{"a":{"default":true,"region":"a-region"},"default":{"region":"fallback-region"},"z":{"region":"z-region"}}`, string(data))

	cf, err = profile.LoadConfigFile(o)
	require.NoError(t, err)
	require.NoError(t, cf.SetDefault("z"))
	p, err := cf.DefaultProfile()
	require.NoError(t, err)
	assert.Equal(t, "z", p.Name)
	require.NoError(t, cf.RemoveProfile("z")) // OSC_PROFILE does not constrain the registry.
	p, err = cf.DefaultProfile()
	require.NoError(t, err)
	assert.Equal(t, "default", p.Name)
	assert.True(t, p.Default)
}

func TestConfigFile_ListProfilesDecodeError(t *testing.T) {
	path := writeConfig(t, map[string]any{
		"a": map[string]any{"region": "valid"},
		"z": map[string]any{"region": 123},
	})
	cf, err := profile.LoadConfigFile(profile.Options{FilePath: &path})
	require.NoError(t, err)
	profiles, err := cf.ListProfiles()
	require.ErrorContains(t, err, `decode profile "z"`)
	assert.Nil(t, profiles)
}
