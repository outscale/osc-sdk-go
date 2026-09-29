package profile_test

import (
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
			Env:      map[string]string{},
		})
		require.NoError(t, err)

		dst := filepath.Join(t.TempDir(), "existing.json")
		cf.FilePath = &dst
		require.NoError(t, cf.Save())

		scf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(dst),
			Env:      map[string]string{},
		})
		require.NoError(t, err)
		assert.Equal(t, "foo", scf.DefaultProfileName)

		got, err := scf.Profile("foo")
		require.NoError(t, err)
		assert.Equal(t, "ak", got.AccessKey)
		assert.Equal(t, "sk", got.SecretKey)
	})

	t.Run("A new profile file can be created in a missing directory", func(t *testing.T) {
		cf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(src),
			Env:      map[string]string{},
		})
		require.NoError(t, err)

		dst := filepath.Join(t.TempDir(), "foo", "bar", "missing.json")
		cf.FilePath = &dst
		require.NoError(t, cf.Save())

		scf, err := profile.LoadConfigFile(profile.Options{
			FilePath: new(dst),
			Env:      map[string]string{},
		})
		require.NoError(t, err)
		assert.Equal(t, "foo", scf.DefaultProfileName)

		got, err := scf.Profile("foo")
		require.NoError(t, err)
		assert.Equal(t, "ak", got.AccessKey)
		assert.Equal(t, "sk", got.SecretKey)
	})
}
