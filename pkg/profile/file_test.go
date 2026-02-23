package profile_test

import (
	"path/filepath"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigFile_Save(t *testing.T) {
	t.Run("A new profile file can be created in an existing directory", func(t *testing.T) {
		cf := &profile.ConfigFile{
			Path: filepath.Join(t.TempDir(), "existing.json"),
			Profiles: map[string]profile.Profile{
				"foo": {Default: true},
			},
		}
		err := cf.Save()
		require.NoError(t, err)

		scf, err := profile.LoadConfigFile(cf.Path)
		require.NoError(t, err)
		assert.Equal(t, cf.Profiles, scf.Profiles)
	})
	t.Run("A new profile file can be created in a missing directory", func(t *testing.T) {
		cf := &profile.ConfigFile{
			Path: filepath.Join(t.TempDir(), "foo", "bar", "missing.json"),
			Profiles: map[string]profile.Profile{
				"foo": {Default: true},
			},
		}
		err := cf.Save()
		require.NoError(t, err)

		scf, err := profile.LoadConfigFile(cf.Path)
		require.NoError(t, err)
		assert.Equal(t, cf.Profiles, scf.Profiles)
	})
}
