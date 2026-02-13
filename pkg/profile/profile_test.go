package profile_test

import (
	"path/filepath"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromEnv(t *testing.T) {
	t.Run("profile is loaded from env", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "foo")
		t.Setenv("OSC_SECRET_KEY", "bar")
		p := profile.Profile{}
		err := profile.FromEnv()(&p)
		require.NoError(t, err)
		assert.True(t, p.IsSet())
		assert.Equal(t, "foo", p.AccessKey)
		assert.Equal(t, "bar", p.SecretKey)
	})
	t.Run("env is ignored if already loaded", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "foo")
		t.Setenv("OSC_SECRET_KEY", "bar")
		p := profile.Profile{
			AccessKey: "bar",
			SecretKey: "baz",
		}
		err := profile.FromEnv()(&p)
		require.NoError(t, err)
		assert.True(t, p.IsSet())
		assert.Equal(t, "bar", p.AccessKey)
		assert.Equal(t, "baz", p.SecretKey)
	})
}

func TestFromFile(t *testing.T) {
	t.Run("the default profile is loaded", func(t *testing.T) {
		cf := newConfigFile(t)
		p := profile.Profile{}
		err := profile.FromFile("", cf.Path)(&p)
		require.NoError(t, err)
		assert.Equal(t, "defaultak", p.AccessKey)
		assert.Equal(t, "defaultsk", p.SecretKey)
	})
	t.Run("a profile is loaded", func(t *testing.T) {
		cf := newConfigFile(t)
		p := profile.Profile{}
		err := profile.FromFile("foo", cf.Path)(&p)
		require.NoError(t, err)
		assert.Equal(t, "fooak", p.AccessKey)
		assert.Equal(t, "foosk", p.SecretKey)
	})
	t.Run("a custom default is loaded", func(t *testing.T) {
		cf := newConfigFile(t)
		err := cf.SetDefault("foo")
		require.NoError(t, err)
		err = cf.Save()
		require.NoError(t, err)
		p := profile.Profile{}
		err = profile.FromFile("", cf.Path)(&p)
		require.NoError(t, err)
		assert.Equal(t, "fooak", p.AccessKey)
		assert.Equal(t, "foosk", p.SecretKey)
	})
}

func TestNewFrom(t *testing.T) {
	cf := newConfigFile(t)
	t.Run("env is loaded first", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "foo")
		t.Setenv("OSC_SECRET_KEY", "bar")
		p, err := profile.NewFrom("", cf.Path)
		require.NoError(t, err)
		assert.Equal(t, "foo", p.AccessKey)
		assert.Equal(t, "bar", p.SecretKey)
	})
	t.Run("env is loaded first, even if profile is set", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "foo")
		t.Setenv("OSC_SECRET_KEY", "bar")
		p, err := profile.NewFrom("foo", cf.Path)
		require.NoError(t, err)
		assert.Equal(t, "foo", p.AccessKey)
		assert.Equal(t, "bar", p.SecretKey)
	})
	t.Run("file is used if env is empty", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "")
		t.Setenv("OSC_SECRET_KEY", "")
		p, err := profile.NewFrom("", cf.Path)
		require.NoError(t, err)
		assert.Equal(t, "defaultak", p.AccessKey)
		assert.Equal(t, "defaultsk", p.SecretKey)
	})
	t.Run("OSC_CONFIG_FILE can define the config file path", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "")
		t.Setenv("OSC_SECRET_KEY", "")
		t.Setenv("OSC_CONFIG_FILE", cf.Path)
		p, err := profile.NewFrom("", "")
		require.NoError(t, err)
		assert.Equal(t, "defaultak", p.AccessKey)
		assert.Equal(t, "defaultsk", p.SecretKey)
	})
	t.Run("OSC_PROFILE can define the profile path", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "")
		t.Setenv("OSC_SECRET_KEY", "")
		t.Setenv("OSC_PROFILE", "foo")
		p, err := profile.NewFrom("", cf.Path)
		require.NoError(t, err)
		assert.Equal(t, "fooak", p.AccessKey)
		assert.Equal(t, "foosk", p.SecretKey)
	})
	t.Run("file is merged with env", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "foo")
		t.Setenv("OSC_SECRET_KEY", "bar")
		p, err := profile.NewFrom("", cf.Path)
		require.NoError(t, err)
		assert.Equal(t, "foo", p.AccessKey)
		assert.Equal(t, "bar", p.SecretKey)
		assert.Equal(t, "defaultakv2", p.AccessKeyV2)
		assert.Equal(t, "defaultskv2", p.SecretKeyV2)
	})
	t.Run("NewFrom does not fail if no credentials are found", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "")
		t.Setenv("OSC_SECRET_KEY", "")
		p, err := profile.NewFrom("", "")
		require.NoError(t, err)
		assert.Equal(t, "", p.AccessKey)
		assert.Equal(t, "", p.SecretKey)
	})
	t.Run("NewFrom fails if a config file was requested and not found", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "")
		t.Setenv("OSC_SECRET_KEY", "")
		_, err := profile.NewFrom("", "/does/not/exist")
		require.Error(t, err)
	})
	t.Run("NewFrom fails if a profile was requested and not found", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "")
		t.Setenv("OSC_SECRET_KEY", "")
		_, err := profile.NewFrom("does not exist", cf.Path)
		require.Error(t, err)
	})
}

func newConfigFile(t *testing.T) *profile.ConfigFile {
	path := filepath.Join(t.TempDir(), "config.json")

	cf := profile.ConfigFile{
		Path: path,
		Profiles: map[string]profile.Profile{
			profile.DefaultProfile: profile.Profile{
				AccessKey:   "defaultak",
				SecretKey:   "defaultsk",
				AccessKeyV2: "defaultakv2",
				SecretKeyV2: "defaultskv2",
			},
			"foo": profile.Profile{
				AccessKey: "fooak",
				SecretKey: "foosk",
			},
		},
	}
	err := cf.Save()
	require.NoError(t, err)
	return &cf
}
