package profile_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeConfig(t *testing.T, profiles map[string]any) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	data, err := json.Marshal(profiles)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(path, data, 0o600))
	return path
}

func load(t *testing.T, o profile.Options) profile.Result {
	t.Helper()
	r, err := profile.Load(o)
	require.NoError(t, err)
	return r
}

func TestLoad_FileSelection(t *testing.T) {
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"access_key": "file-ak", "secret_key": "file-sk"},
	})

	t.Run("tool path wins over env and default", func(t *testing.T) {
		other := writeConfig(t, map[string]any{
			"default": map[string]any{"access_key": "other-ak", "secret_key": "other-sk"},
		})
		r := load(t, profile.Options{
			FilePath: new(cfg),
			Env:      map[string]string{profile.ConfigFileEnvName: other},
		})
		assert.Equal(t, "file-ak", r.Values.AccessKey)
		assert.Equal(t, &cfg, r.File)
	})

	t.Run("OSC_CONFIG_FILE selects path", func(t *testing.T) {
		r := load(t, profile.Options{
			Env: map[string]string{profile.ConfigFileEnvName: cfg},
		})
		assert.Equal(t, "file-ak", r.Values.AccessKey)
		assert.Equal(t, &cfg, r.File)
	})

	t.Run("missing default path is ok", func(t *testing.T) {
		r := load(t, profile.Options{
			Env: map[string]string{},
			ReadFile: func(string) ([]byte, error) {
				return nil, os.ErrNotExist
			},
		})
		assert.Empty(t, r.Values.AccessKey)
		assert.Empty(t, r.File)
		assert.Equal(t, "https", r.Values.Protocol)
		assert.Equal(t, "eu-west-2", r.Values.Region)
	})

	t.Run("explicit missing path is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			FilePath: new("/does/not/exist.json"),
			Env:      map[string]string{},
			ReadFile: func(s string) ([]byte, error) {
				return nil, os.ErrNotExist
			},
		})
		require.Error(t, err)
	})

	t.Run("OSC_CONFIG_FILE missing path is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			Env: map[string]string{profile.ConfigFileEnvName: "/does/not/exist.json"},
			ReadFile: func(s string) ([]byte, error) {
				return nil, os.ErrNotExist
			},
		})
		require.Error(t, err)
	})

	t.Run("explicit empty path is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{FilePath: new(""), Env: map[string]string{}})
		require.Error(t, err)
	})

	t.Run("invalid config content is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			FilePath: new(cfg),
			Env:      map[string]string{},
			ReadFile: func(string) ([]byte, error) {
				return []byte(`{"default": {"nope": true}}`), nil
			},
		})
		require.Error(t, err)
	})
}

func TestLoad_ProfileSelection(t *testing.T) {
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"access_key": "def-ak", "secret_key": "def-sk"},
		"foo":     map[string]any{"access_key": "foo-ak", "secret_key": "foo-sk"},
		"bar":     map[string]any{"access_key": "bar-ak", "secret_key": "bar-sk", "default": true},
	})

	t.Run("tool profile name wins", func(t *testing.T) {
		r := load(t, profile.Options{
			FilePath:    new(cfg),
			ProfileName: new("foo"),
			Env:         map[string]string{profile.ProfileNameEnvName: "bar"},
		})
		assert.Equal(t, "foo", r.ProfileName)
		assert.Equal(t, "foo-ak", r.Values.AccessKey)
	})

	t.Run("OSC_PROFILE selects name", func(t *testing.T) {
		r := load(t, profile.Options{
			FilePath: new(cfg),
			Env:      map[string]string{profile.ProfileNameEnvName: "foo"},
		})
		assert.Equal(t, "foo", r.ProfileName)
		assert.Equal(t, "foo-ak", r.Values.AccessKey)
	})

	t.Run("marked default is selected", func(t *testing.T) {
		r := load(t, profile.Options{FilePath: new(cfg), Env: map[string]string{}})
		assert.Equal(t, "bar", r.ProfileName)
		assert.Equal(t, "bar-ak", r.Values.AccessKey)
	})

	t.Run("named default is fallback", func(t *testing.T) {
		only := writeConfig(t, map[string]any{
			"default": map[string]any{"access_key": "def-ak", "secret_key": "def-sk"},
			"foo":     map[string]any{"access_key": "foo-ak", "secret_key": "foo-sk"},
		})
		r := load(t, profile.Options{FilePath: new(only), Env: map[string]string{}})
		assert.Equal(t, "default", r.ProfileName)
		assert.Equal(t, "def-ak", r.Values.AccessKey)
	})

	t.Run("multiple marked defaults is error", func(t *testing.T) {
		bad := writeConfig(t, map[string]any{
			"a": map[string]any{"default": true},
			"b": map[string]any{"default": true},
		})
		_, err := profile.Load(profile.Options{FilePath: new(bad), Env: map[string]string{}})
		require.Error(t, err)
	})

	t.Run("explicit missing profile is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			FilePath:    new(cfg),
			ProfileName: new("missing"),
			Env:         map[string]string{},
		})
		require.Error(t, err)
	})

	t.Run("explicit profile missing with no file is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			ProfileName: new("foo"),
			Env:         map[string]string{},
			ReadFile: func(string) ([]byte, error) {
				return nil, os.ErrNotExist
			},
		})
		require.Error(t, err)
	})

	t.Run("empty profile name is invalid", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			FilePath:    new(cfg),
			ProfileName: new(""),
			Env:         map[string]string{},
		})
		require.Error(t, err)
	})

	t.Run("empty OSC_PROFILE is invalid", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			FilePath: new(cfg),
			Env:      map[string]string{profile.ProfileNameEnvName: ""},
		})
		require.Error(t, err)
	})
}

func TestLoad_Precedence(t *testing.T) {
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{
			"access_key": "file-ak",
			"secret_key": "file-sk",
			"region":     "file-region",
			"protocol":   "http",
		},
	})

	t.Run("without explicit file/profile: tool > env > profile > default", func(t *testing.T) {
		r := load(t, profile.Options{
			Overrides: profile.Fields{Region: "tool-region"},
			Env: map[string]string{
				"OSC_ACCESS_KEY": "env-ak",
				"OSC_SECRET_KEY": "env-sk",
				"OSC_PROTOCOL":   "https",
			},
			ReadFile: func(string) ([]byte, error) {
				return os.ReadFile(cfg)
			},
		})
		assert.Equal(t, "tool-region", r.Values.Region)
		assert.Equal(t, "tool", r.Sources["Region"])
		assert.Equal(t, "env-ak", r.Values.AccessKey)
		assert.Equal(t, "environment", r.Sources["AccessKey"])
		assert.Equal(t, "https", r.Values.Protocol)
		assert.Equal(t, "environment", r.Sources["Protocol"])
	})

	t.Run("without explicit: env beats profile", func(t *testing.T) {
		r := load(t, profile.Options{
			Env: map[string]string{
				"OSC_ACCESS_KEY": "env-ak",
				"OSC_SECRET_KEY": "env-sk",
			},
			ReadFile: func(string) ([]byte, error) { return os.ReadFile(cfg) },
		})
		assert.Equal(t, "env-ak", r.Values.AccessKey)
		assert.Equal(t, "file-region", r.Values.Region)
		assert.Equal(t, "profile:default", r.Sources["Region"])
	})

	t.Run("with explicit profile: profile beats env", func(t *testing.T) {
		r := load(t, profile.Options{
			FilePath:    new(cfg),
			ProfileName: new("default"),
			Env: map[string]string{
				"OSC_ACCESS_KEY": "env-ak",
				"OSC_SECRET_KEY": "env-sk",
				"OSC_REGION":     "env-region",
			},
		})
		assert.Equal(t, "file-ak", r.Values.AccessKey)
		assert.Equal(t, "profile:default", r.Sources["AccessKey"])
		assert.Equal(t, "file-region", r.Values.Region)
	})

	t.Run("with explicit file path: profile beats env", func(t *testing.T) {
		r := load(t, profile.Options{
			FilePath: new(cfg),
			Env: map[string]string{
				"OSC_ACCESS_KEY": "env-ak",
				"OSC_SECRET_KEY": "env-sk",
			},
		})
		assert.Equal(t, "file-ak", r.Values.AccessKey)
		assert.Equal(t, "profile:default", r.Sources["AccessKey"])
	})

	t.Run("with OSC_CONFIG_FILE: profile beats env", func(t *testing.T) {
		r := load(t, profile.Options{
			Env: map[string]string{
				profile.ConfigFileEnvName: cfg,
				"OSC_ACCESS_KEY":          "env-ak",
				"OSC_SECRET_KEY":          "env-sk",
			},
		})
		assert.Equal(t, "file-ak", r.Values.AccessKey)
	})

	t.Run("tool overrides always win", func(t *testing.T) {
		r := load(t, profile.Options{
			FilePath:  new(cfg),
			Overrides: profile.Fields{AccessKey: "tool-ak", SecretKey: "tool-sk"},
			Env: map[string]string{
				"OSC_ACCESS_KEY": "env-ak",
				"OSC_SECRET_KEY": "env-sk",
			},
		})
		assert.Equal(t, "tool-ak", r.Values.AccessKey)
		assert.Equal(t, "tool", r.Sources["AccessKey"])
	})

	t.Run("built-in defaults apply last", func(t *testing.T) {
		r := load(t, profile.Options{
			Env:      map[string]string{},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		assert.Equal(t, "eu-west-2", r.Values.Region)
		assert.Equal(t, "https", r.Values.Protocol)
		assert.Equal(t, "default", r.Sources["Region"])
		assert.Equal(t, "default", r.Sources["Protocol"])
	})
}

func TestLoad_AtomicPairs(t *testing.T) {
	t.Run("incomplete ak/sk pair from env is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			Env:      map[string]string{"OSC_ACCESS_KEY": "only-ak"},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "aksk")
	})

	t.Run("incomplete ak/sk pair from tool is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			Overrides: profile.Fields{AccessKey: "only-ak"},
			Env:       map[string]string{},
			ReadFile:  func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
	})

	t.Run("incomplete pair from profile is error", func(t *testing.T) {
		cfg := writeConfig(t, map[string]any{
			"default": map[string]any{"access_key": "only-ak"},
		})
		_, err := profile.Load(profile.Options{FilePath: new(cfg), Env: map[string]string{}})
		require.Error(t, err)
	})

	t.Run("complete pair from one source is ok", func(t *testing.T) {
		r := load(t, profile.Options{
			Env: map[string]string{
				"OSC_ACCESS_KEY": "ak",
				"OSC_SECRET_KEY": "sk",
			},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		assert.Equal(t, "ak", r.Values.AccessKey)
		assert.Equal(t, "sk", r.Values.SecretKey)
	})

	t.Run("login/password pair is atomic", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			Env:      map[string]string{"OSC_LOGIN": "user"},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "basicauth")
	})

	t.Run("x509 path pair is atomic", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			Env:      map[string]string{"OSC_X509_CLIENT_CERT": "/tmp/cert"},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "x509keypair")
	})
}

func TestLoad_EnvAndEndpoints(t *testing.T) {
	t.Run("env loads standard fields", func(t *testing.T) {
		r := load(t, profile.Options{
			Env: map[string]string{
				"OSC_ACCESS_KEY":          "ak",
				"OSC_SECRET_KEY":          "sk",
				"OSC_REGION":              "us-east-2",
				"OSC_PROTOCOL":            "http",
				"OSC_ENDPOINT_API":        "https://api.example",
				"OSC_ENDPOINT_FCU":        "https://fcu.example",
				"OSC_TLS_SKIP_VERIFY":     "true",
				"OSC_IAM_V2_SERVICES":     "oos,oks",
				"OSC_ENDPOINT_DIRECTLINK": "https://dl.example",
			},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		assert.Equal(t, "ak", r.Values.AccessKey)
		assert.Equal(t, "us-east-2", r.Values.Region)
		assert.Equal(t, "http", r.Values.Protocol)
		assert.Equal(t, "https://api.example", r.Values.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "https://fcu.example", r.Values.Endpoints[profile.OscServiceFCU])
		assert.Equal(t, "https://dl.example", r.Values.Endpoints[profile.OscServiceDirectLink])
		require.NotNil(t, r.Values.TlsSkipVerify)
		assert.True(t, *r.Values.TlsSkipVerify)
		assert.Equal(t, []string{"oos", "oks"}, r.Values.IAMV2Services)
	})

	t.Run("endpoints merge field by field", func(t *testing.T) {
		cfg := writeConfig(t, map[string]any{
			"default": map[string]any{
				"endpoints": map[string]any{
					"api": "https://file-api",
					"fcu": "https://file-fcu",
				},
			},
		})
		r := load(t, profile.Options{
			Env: map[string]string{
				profile.ConfigFileEnvName: cfg,
				"OSC_ENDPOINT_API":        "https://env-api",
			},
		})
		assert.Equal(t, "https://file-api", r.Values.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "https://file-fcu", r.Values.Endpoints[profile.OscServiceFCU])
	})

	t.Run("endpoints merge field by field without explicit: env beats profile", func(t *testing.T) {
		cfg := writeConfig(t, map[string]any{
			"default": map[string]any{
				"endpoints": map[string]any{
					"api": "https://file-api",
					"fcu": "https://file-fcu",
				},
			},
		})
		r := load(t, profile.Options{
			Env: map[string]string{"OSC_ENDPOINT_API": "https://env-api"},
			ReadFile: func(string) ([]byte, error) {
				return os.ReadFile(cfg)
			},
		})
		assert.Equal(t, "https://env-api", r.Values.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "https://file-fcu", r.Values.Endpoints[profile.OscServiceFCU])
		assert.Equal(t, "environment", r.Sources["Endpoints.api"])
		assert.Equal(t, "profile:default", r.Sources["Endpoints.fcu"])
	})

	t.Run("invalid TLS_SKIP_VERIFY is error", func(t *testing.T) {
		_, err := profile.Load(profile.Options{
			Env:      map[string]string{"OSC_TLS_SKIP_VERIFY": "maybe"},
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
	})

	t.Run("tls false from env overrides profile true when env has precedence", func(t *testing.T) {
		cfg := writeConfig(t, map[string]any{
			"default": map[string]any{"tls_skip_verify": true},
		})
		r := load(t, profile.Options{
			Env: map[string]string{"OSC_TLS_SKIP_VERIFY": "false"},
			ReadFile: func(string) ([]byte, error) {
				return os.ReadFile(cfg)
			},
		})
		require.NotNil(t, r.Values.TlsSkipVerify)
		assert.False(t, *r.Values.TlsSkipVerify)
		assert.Equal(t, "environment", r.Sources["TlsSkipVerify"])
	})
}

func TestLoad_NoCredentialsOk(t *testing.T) {
	r := load(t, profile.Options{
		Env:      map[string]string{},
		ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
	})
	assert.Empty(t, r.Values.AccessKey)
	assert.Empty(t, r.Values.SecretKey)
}

func TestNewFrom_Compat(t *testing.T) {
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"access_key": "def-ak", "secret_key": "def-sk"},
		"foo":     map[string]any{"access_key": "foo-ak", "secret_key": "foo-sk"},
	})

	t.Run("loads named profile", func(t *testing.T) {
		p, err := profile.NewFrom(cfg, "foo")
		require.NoError(t, err)
		assert.Equal(t, "foo-ak", p.Values.AccessKey)
	})

	t.Run("empty name uses default selection", func(t *testing.T) {
		p, err := profile.NewFrom(cfg, "")
		require.NoError(t, err)
		assert.Equal(t, "def-ak", p.Values.AccessKey)
	})
}
