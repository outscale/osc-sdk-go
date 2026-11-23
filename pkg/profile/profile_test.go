package profile_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// clearProfileEnv removes inherited profile settings for this test and its subtests.
func clearProfileEnv(t *testing.T) {
	t.Helper()
	for _, e := range os.Environ() {
		if k, _, ok := strings.Cut(e, "="); ok && len(k) >= len("OSC_") && strings.EqualFold(k[:len("OSC_")], "OSC_") {
			// Register restoration with t.Setenv before unsetting. Some settings distinguish empty from unset.
			t.Setenv(k, "")
			require.NoError(t, os.Unsetenv(k))
		}
	}
}

func writeConfig(t *testing.T, profiles map[string]any) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	data, err := json.Marshal(profiles)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(path, data, 0o600))
	return path
}

func load(t *testing.T, o profile.Options) (profile.Profile, profile.ProfileExplanation) {
	t.Helper()
	r, explanation, err := profile.Load(o)
	require.NoError(t, err)
	return r, explanation
}

func TestLoad_FileSelection(t *testing.T) {
	clearProfileEnv(t)
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"access_key": "file-ak", "secret_key": "file-sk"},
	})

	t.Run("tool path wins over env and default", func(t *testing.T) {
		other := writeConfig(t, map[string]any{
			"default": map[string]any{"access_key": "other-ak", "secret_key": "other-sk"},
		})
		t.Setenv(profile.ConfigFileEnvName, other)
		r, explanation := load(t, profile.Options{
			FilePath: new(cfg),
		})
		assert.Equal(t, "file-ak", r.AccessKey)
		assert.Equal(t, &cfg, explanation.File)
	})

	t.Run("OSC_CONFIG_FILE selects path", func(t *testing.T) {
		t.Setenv("OSC_CONFIG_FILE", cfg)
		r, explanation := load(t, profile.Options{})
		assert.Equal(t, "file-ak", r.AccessKey)
		assert.Equal(t, &cfg, explanation.File)
	})

	t.Run("missing default path is ok", func(t *testing.T) {
		home := t.TempDir()
		t.Setenv("HOME", home)
		t.Setenv("USERPROFILE", home)
		o, _ := profile.NewOptionsWith(func(o *profile.Options) error {
			o.ReadFile = func(path string) ([]byte, error) {
				assert.Equal(t, filepath.Join(home, ".osc", "config.json"), path)
				return nil, os.ErrNotExist
			}
			return nil
		})
		r, explanation := load(t, o)

		assert.Empty(t, r.AccessKey)
		assert.Empty(t, explanation.File)
		assert.Equal(t, "https", r.Protocol)
		assert.Equal(t, "eu-west-2", r.Region)
	})

	t.Run("explicit missing path is error", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{
			FilePath: new("/does/not/exist.json"),
			ReadFile: func(s string) ([]byte, error) {
				return nil, os.ErrNotExist
			},
		})
		require.Error(t, err)
	})

	t.Run("OSC_CONFIG_FILE missing path is error", func(t *testing.T) {
		t.Setenv("OSC_CONFIG_FILE", filepath.Join(t.TempDir(), "missing.json"))
		_, _, err := profile.Load(profile.Options{})
		require.ErrorIs(t, err, os.ErrNotExist)
	})

	t.Run("explicit empty path is error", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{FilePath: new("")})
		require.Error(t, err)
	})

	t.Run("invalid config content is error", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{
			FilePath: new(cfg),
			ReadFile: func(string) ([]byte, error) {
				return []byte(`{"default": {"nope": true}}`), nil
			},
		})
		require.Error(t, err)
	})
}

func TestLoad_ProfileSelection(t *testing.T) {
	clearProfileEnv(t)
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"access_key": "def-ak", "secret_key": "def-sk"},
		"foo":     map[string]any{"access_key": "foo-ak", "secret_key": "foo-sk"},
		"bar":     map[string]any{"access_key": "bar-ak", "secret_key": "bar-sk", "default": true},
	})

	t.Run("tool profile name wins", func(t *testing.T) {
		t.Setenv(profile.ProfileNameEnvName, "bar")
		r, _ := load(t, profile.Options{
			FilePath:    new(cfg),
			ProfileName: new("foo"),
		})
		assert.Equal(t, "foo", r.Name)
		assert.False(t, r.Default)
		assert.Equal(t, "foo-ak", r.AccessKey)
	})

	t.Run("OSC_PROFILE selects name", func(t *testing.T) {
		t.Setenv(profile.ProfileNameEnvName, "foo")
		r, _ := load(t, profile.Options{FilePath: new(cfg)})
		assert.Equal(t, "foo", r.Name)
		assert.Equal(t, "foo-ak", r.AccessKey)
	})

	t.Run("marked default is selected", func(t *testing.T) {
		r, _ := load(t, profile.Options{FilePath: new(cfg)})
		assert.Equal(t, "bar", r.Name)
		assert.True(t, r.Default)
		assert.Equal(t, "bar-ak", r.AccessKey)
	})

	t.Run("named default is fallback", func(t *testing.T) {
		only := writeConfig(t, map[string]any{
			"default": map[string]any{"access_key": "def-ak", "secret_key": "def-sk"},
			"foo":     map[string]any{"access_key": "foo-ak", "secret_key": "foo-sk"},
		})
		r, _ := load(t, profile.Options{FilePath: new(only)})
		assert.Equal(t, "default", r.Name)
		assert.True(t, r.Default)
		assert.Equal(t, "def-ak", r.AccessKey)
	})

	t.Run("multiple marked defaults is error", func(t *testing.T) {
		bad := writeConfig(t, map[string]any{
			"a": map[string]any{"default": true},
			"b": map[string]any{"default": true},
		})
		_, _, err := profile.Load(profile.Options{FilePath: new(bad)})
		require.Error(t, err)
	})

	t.Run("explicit missing profile is error", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{
			FilePath:    new(cfg),
			ProfileName: new("missing"),
		})
		require.ErrorIs(t, err, profile.ErrProfileNotFound)
	})

	t.Run("explicit profile missing with no file is error", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{
			ProfileName: new("foo"),
			ReadFile: func(string) ([]byte, error) {
				return nil, os.ErrNotExist
			},
		})
		require.Error(t, err)
	})

	t.Run("empty profile name is invalid", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{
			FilePath:    new(cfg),
			ProfileName: new(""),
		})
		require.Error(t, err)
	})

	t.Run("empty OSC_PROFILE is invalid", func(t *testing.T) {
		t.Setenv(profile.ProfileNameEnvName, "")
		_, _, err := profile.Load(profile.Options{
			FilePath: new(cfg),
		})
		require.Error(t, err)
	})
}

func TestLoad_Precedence(t *testing.T) {
	clearProfileEnv(t)
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{
			"access_key": "file-ak",
			"secret_key": "file-sk",
			"region":     "file-region",
			"protocol":   "http",
		},
	})

	t.Run("without explicit file/profile: tool > env > profile > default", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "env-ak")
		t.Setenv("OSC_SECRET_KEY", "env-sk")
		t.Setenv("OSC_PROTOCOL", "https")
		r, explanation := load(t, profile.Options{
			Overrides: profile.Fields{Region: "tool-region"},
			ReadFile: func(string) ([]byte, error) {
				return os.ReadFile(cfg)
			},
		})
		assert.Equal(t, "tool-region", r.Region)
		assert.Equal(t, "tool", explanation.Sources["Region"])
		assert.Equal(t, "env-ak", r.AccessKey)
		assert.Equal(t, "environment", explanation.Sources["AccessKey"])
		assert.Equal(t, "https", r.Protocol)
		assert.Equal(t, "environment", explanation.Sources["Protocol"])
	})

	t.Run("without explicit: env beats profile", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "env-ak")
		t.Setenv("OSC_SECRET_KEY", "env-sk")
		r, explanation := load(t, profile.Options{
			ReadFile: func(string) ([]byte, error) { return os.ReadFile(cfg) },
		})
		assert.Equal(t, "env-ak", r.AccessKey)
		assert.Equal(t, "file-region", r.Region)
		assert.Equal(t, "profile:default", explanation.Sources["Region"])
	})

	t.Run("with explicit profile: profile beats env", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "env-ak")
		t.Setenv("OSC_SECRET_KEY", "env-sk")
		t.Setenv("OSC_REGION", "env-region")
		r, explanation := load(t, profile.Options{
			FilePath:    new(cfg),
			ProfileName: new("default"),
		})
		assert.Equal(t, "file-ak", r.AccessKey)
		assert.Equal(t, "profile:default", explanation.Sources["AccessKey"])
		assert.Equal(t, "file-region", r.Region)
	})

	t.Run("with explicit file path: profile beats env", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "env-ak")
		t.Setenv("OSC_SECRET_KEY", "env-sk")
		r, explanation := load(t, profile.Options{FilePath: new(cfg)})
		assert.Equal(t, "file-ak", r.AccessKey)
		assert.Equal(t, "profile:default", explanation.Sources["AccessKey"])
	})

	t.Run("with OSC_CONFIG_FILE: profile beats env", func(t *testing.T) {
		t.Setenv(profile.ConfigFileEnvName, cfg)
		t.Setenv("OSC_ACCESS_KEY", "env-ak")
		t.Setenv("OSC_SECRET_KEY", "env-sk")
		r, _ := load(t, profile.Options{})
		assert.Equal(t, "file-ak", r.AccessKey)
	})

	t.Run("tool overrides always win", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "env-ak")
		t.Setenv("OSC_SECRET_KEY", "env-sk")
		r, explanation := load(t, profile.Options{
			FilePath:  new(cfg),
			Overrides: profile.Fields{AccessKey: "tool-ak", SecretKey: "tool-sk"},
		})
		assert.Equal(t, "tool-ak", r.AccessKey)
		assert.Equal(t, "tool", explanation.Sources["AccessKey"])
	})

	t.Run("built-in defaults apply last", func(t *testing.T) {
		o, _ := profile.NewOptionsWith(func(o *profile.Options) error {
			o.ReadFile = func(string) ([]byte, error) {
				return nil, os.ErrNotExist
			}
			return nil
		})
		r, explanation := load(t, o)
		assert.Equal(t, "eu-west-2", r.Region)
		assert.Equal(t, "https", r.Protocol)
		assert.Equal(t, "default", explanation.Sources["Region"])
		assert.Equal(t, "default", explanation.Sources["Protocol"])
	})
}

func TestLoad_AtomicPairs(t *testing.T) {
	clearProfileEnv(t)
	t.Run("incomplete ak/sk pair from env is error", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "only-ak")
		_, _, err := profile.Load(profile.Options{
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "aksk")
	})

	t.Run("incomplete ak/sk pair from tool is error", func(t *testing.T) {
		_, _, err := profile.Load(profile.Options{
			Overrides: profile.Fields{AccessKey: "only-ak"},
			ReadFile:  func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
	})

	t.Run("incomplete pair from profile is error", func(t *testing.T) {
		cfg := writeConfig(t, map[string]any{
			"default": map[string]any{"access_key": "only-ak"},
		})
		_, _, err := profile.Load(profile.Options{FilePath: new(cfg)})
		require.Error(t, err)
	})

	t.Run("complete pair from one source is ok", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "ak")
		t.Setenv("OSC_SECRET_KEY", "sk")
		r, _ := load(t, profile.Options{
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		assert.Equal(t, "ak", r.AccessKey)
		assert.Equal(t, "sk", r.SecretKey)
	})

	t.Run("login/password pair is atomic", func(t *testing.T) {
		t.Setenv("OSC_LOGIN", "user")
		_, _, err := profile.Load(profile.Options{
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "basicauth")
	})

	t.Run("x509 path pair is atomic", func(t *testing.T) {
		t.Setenv("OSC_X509_CLIENT_CERT", "/tmp/cert")
		_, _, err := profile.Load(profile.Options{
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "x509keypair")
	})
}

func TestLoad_EnvAndEndpoints(t *testing.T) {
	clearProfileEnv(t)
	t.Run("env loads standard fields", func(t *testing.T) {
		t.Setenv("OSC_ACCESS_KEY", "ak")
		t.Setenv("OSC_SECRET_KEY", "sk")
		t.Setenv("OSC_REGION", "us-east-2")
		t.Setenv("OSC_PROTOCOL", "http")
		t.Setenv("OSC_ENDPOINT_API", "https://api.example")
		t.Setenv("OSC_ENDPOINT_FCU", "https://fcu.example")
		t.Setenv("OSC_TLS_SKIP_VERIFY", "true")
		t.Setenv("OSC_IAM_V2_SERVICES", "oos,oks")
		t.Setenv("OSC_ENDPOINT_DIRECTLINK", "https://dl.example")
		r, _ := load(t, profile.Options{
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		assert.Equal(t, "ak", r.AccessKey)
		assert.Equal(t, "us-east-2", r.Region)
		assert.Equal(t, "http", r.Protocol)
		assert.Equal(t, "https://api.example", r.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "https://fcu.example", r.Endpoints[profile.OscServiceFCU])
		assert.Equal(t, "https://dl.example", r.Endpoints[profile.OscServiceDirectLink])
		require.NotNil(t, r.TlsSkipVerify)
		assert.True(t, *r.TlsSkipVerify)
		assert.Equal(t, []string{"oos", "oks"}, r.IAMV2Services)
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
		t.Setenv(profile.ConfigFileEnvName, cfg)
		t.Setenv("OSC_ENDPOINT_API", "https://env-api")
		r, _ := load(t, profile.Options{})
		assert.Equal(t, "https://file-api", r.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "https://file-fcu", r.Endpoints[profile.OscServiceFCU])
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
		t.Setenv("OSC_ENDPOINT_API", "https://env-api")
		t.Setenv("OSC_ENDPOINT_FCU", "")
		r, explanation := load(t, profile.Options{
			ReadFile: func(string) ([]byte, error) {
				return os.ReadFile(cfg)
			},
		})
		assert.Equal(t, "https://env-api", r.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "https://file-fcu", r.Endpoints[profile.OscServiceFCU])
		assert.Equal(t, "environment", explanation.Sources["Endpoints.api"])
		assert.Equal(t, "profile:default", explanation.Sources["Endpoints.fcu"])
	})

	t.Run("empty OSC_IAM_V2_SERVICES is unset", func(t *testing.T) {
		t.Setenv("OSC_IAM_V2_SERVICES", "")
		r, _ := load(t, profile.Options{
			ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
		})
		assert.Empty(t, r.IAMV2Services)
	})
}

func TestLoad_TLSSkipVerify(t *testing.T) {
	clearProfileEnv(t)
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"tls_skip_verify": true},
	})
	for _, tt := range []struct {
		name  string
		value *string
		want  bool
	}{
		{name: "unset uses profile", want: true},
		{name: "yes enables case insensitively", value: new("YeS"), want: true},
		{name: "empty disables", value: new("")},
		{name: "false disables", value: new("false")},
		{name: "unknown disables", value: new("maybe")},
		{name: "numeric value disables", value: new("1")},
		{name: "non-ASCII lookalike disables", value: new("yeſ")},
	} {
		t.Run(tt.name, func(t *testing.T) {
			wantSource := "profile:default"
			if tt.value != nil {
				t.Setenv("OSC_TLS_SKIP_VERIFY", *tt.value)
				wantSource = "environment"
			}
			r, explanation := load(t, profile.Options{
				ReadFile: func(string) ([]byte, error) { return os.ReadFile(cfg) },
			})
			require.NotNil(t, r.TlsSkipVerify)
			assert.Equal(t, tt.want, *r.TlsSkipVerify)
			assert.Equal(t, wantSource, explanation.Sources["TlsSkipVerify"])
		})
	}
}

func TestLoad_NoCredentialsOk(t *testing.T) {
	clearProfileEnv(t)
	r, _ := load(t, profile.Options{
		ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist },
	})
	assert.Empty(t, r.AccessKey)
	assert.Empty(t, r.SecretKey)
}

func TestWithDefaultRegion(t *testing.T) {
	clearProfileEnv(t)
	cfg := writeConfig(t, map[string]any{"default": map[string]any{"region": "file-region"}})
	for _, tt := range []struct {
		name string
		env  string
		opts profile.Options
		want string
	}{
		{name: "override", opts: profile.Options{Overrides: profile.Fields{Region: "tool-region"}}, want: "tool-region"},
		{name: "environment", env: "env-region", want: "env-region"},
		{name: "file", opts: profile.Options{FilePath: &cfg}, want: "file-region"},
	} {
		t.Run(tt.name+" skips resolver", func(t *testing.T) {
			t.Setenv("OSC_REGION", tt.env)
			o := tt.opts
			if o.FilePath == nil {
				o.ReadFile = func(string) ([]byte, error) { return nil, os.ErrNotExist }
			}
			require.NoError(t, profile.WithDefaultRegion(func() (string, error) {
				return "", errors.New("should not query metadata")
			})(&o))
			p, _ := load(t, o)
			assert.Equal(t, tt.want, p.Region)
		})
	}

	t.Run("resolver supplies endpoint region", func(t *testing.T) {
		o := profile.Options{ReadFile: func(string) ([]byte, error) { return nil, os.ErrNotExist }}
		calls := 0
		require.NoError(t, profile.WithDefaultRegion(func() (string, error) {
			calls++
			return "vm-region", nil
		})(&o))
		p, explanation := load(t, o)
		assert.Equal(t, 1, calls)
		assert.Equal(t, "vm-region", p.Region)
		assert.Equal(t, "https://api.vm-region.outscale.com/api/v1", p.Endpoints[profile.OscServiceApi])
		assert.Equal(t, "default", explanation.Sources["Region"])
	})

	t.Run("resolver error reaches constructor caller", func(t *testing.T) {
		sentinel := errors.New("metadata unavailable")
		p, err := profile.NewWith(func(o *profile.Options) error {
			o.ReadFile = func(string) ([]byte, error) { return nil, os.ErrNotExist }
			return nil
		}, profile.WithDefaultRegion(func() (string, error) { return "", sentinel }))
		require.ErrorIs(t, err, sentinel)
		assert.Nil(t, p)
	})
}

func TestNewFrom_Compat(t *testing.T) {
	clearProfileEnv(t)
	cfg := writeConfig(t, map[string]any{
		"default": map[string]any{"access_key": "def-ak", "secret_key": "def-sk"},
		"foo":     map[string]any{"access_key": "foo-ak", "secret_key": "foo-sk"},
	})

	t.Run("loads named profile", func(t *testing.T) {
		p, err := profile.NewFrom("foo", cfg)
		require.NoError(t, err)
		assert.Equal(t, "foo", p.Name)
		assert.Equal(t, "foo-ak", p.AccessKey)
	})

	t.Run("empty name uses default selection", func(t *testing.T) {
		p, err := profile.NewFrom("", cfg)
		require.NoError(t, err)
		assert.Equal(t, "def-ak", p.AccessKey)
	})
}

func TestNewWith_Error(t *testing.T) {
	clearProfileEnv(t)
	t.Run("option fails", func(t *testing.T) {
		sentinel := errors.New("option failed")
		p, err := profile.NewWith(func(*profile.Options) error { return sentinel })
		require.ErrorIs(t, err, sentinel)
		assert.Nil(t, p)
	})

	t.Run("explicit config file is missing", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.json")
		p, err := profile.NewWith(func(o *profile.Options) error {
			o.FilePath = &path
			return nil
		})
		require.ErrorIs(t, err, os.ErrNotExist)
		assert.Nil(t, p)
	})
}
