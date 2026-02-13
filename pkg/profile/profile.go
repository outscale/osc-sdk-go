// Package profile provide functions to manage standard Outscale profiles.
//
// The profile can be loaded from environment variables or from a JSON configuration file.
// The configuration file is located at ~/.osc/config.json by
package profile

import (
	"os"
	"path"
	"strings"

	"dario.cat/mergo"
)

const (
	DefaultProfile = "default"
)

type Profile struct {
	Default           bool     `json:"default,omitempty"`
	AccessKey         string   `json:"access_key,omitempty"`
	SecretKey         string   `json:"secret_key,omitempty"`
	AccessKeyV2       string   `json:"access_key_v2,omitempty"`
	SecretKeyV2       string   `json:"secret_key_v2,omitempty"`
	IAMV2Services     []string `json:"iam_v2_services,omitempty"`
	X509ClientCert    string   `json:"x509_client_cert,omitempty"`
	X509ClientCertB64 string   `json:"x509_client_cert_b64,omitempty"`
	X509ClientKey     string   `json:"x509_client_key,omitempty"`
	X509ClientKeyB64  string   `json:"x509_client_key_b64,omitempty"`
	TlsSkipVerify     bool     `json:"tls_skip_verify,omitempty"`
	Login             string   `json:"login,omitempty"`
	Password          string   `json:"password,omitempty"`
	Protocol          string   `json:"protocol,omitempty"`
	Region            string   `json:"region,omitempty"`
	Endpoints         Endpoint `json:"endpoints,omitempty"`
}

func (p *Profile) IsSet() bool {
	return (p.AccessKey != "" && p.SecretKey != "") || (p.AccessKeyV2 != "" && p.SecretKeyV2 != "")
}

type Option func(*Profile) error

func DefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".osc", "config.json"), nil
}

// FromFile is an option that loads env from a file.
func FromFile(profile, path string) Option {
	return func(p *Profile) error {
		if p.IsSet() {
			return nil
		}

		if profile == "" {
			profile = os.Getenv("OSC_PROFILE")
		}
		if path == "" {
			path = os.Getenv("OSC_CONFIG_FILE")
		}
		configFile, err := LoadConfigFile(path)
		switch {
		// no path/profile given, do not fail if default file is missing
		case path == "" && profile == "" && err != nil:
			return nil
		case err != nil:
			return err
		}
		cfp, err := configFile.Profile(profile)
		switch {
		// no path/profile given, do not fail if default profile is missing
		case path == "" && profile == "" && err != nil:
			return nil
		case err != nil:
			return err
		}
		*p = cfp
		return nil
	}
}

func getenvBool(key string) bool {
	env := strings.ToLower(os.Getenv(key))
	return env == "yes" || env == "true"
}

// FromEnv is an option that loads a profile from env vars.
func FromEnv() Option {
	return func(p *Profile) error {
		if p.IsSet() {
			return nil
		}

		p.AccessKey = os.Getenv("OSC_ACCESS_KEY")
		p.SecretKey = os.Getenv("OSC_SECRET_KEY")
		p.AccessKeyV2 = os.Getenv("OSC_ACCESS_KEY_V2")
		p.SecretKeyV2 = os.Getenv("OSC_SECRET_KEY_V2")
		p.X509ClientCert = os.Getenv("OSC_X509_CLIENT_CERT")
		p.X509ClientCertB64 = os.Getenv("OSC_X509_CLIENT_CERT_B64")
		p.X509ClientKey = os.Getenv("OSC_X509_CLIENT_KEY")
		p.X509ClientKeyB64 = os.Getenv("OSC_X509_CLIENT_KEY_B64")
		p.TlsSkipVerify = getenvBool("OSC_TLS_SKIP_VERIFY")
		p.Login = os.Getenv("OSC_LOGIN")
		p.Password = os.Getenv("OSC_PASSWORD")
		p.Protocol = os.Getenv("OSC_PROTOCOL")
		p.Region = os.Getenv("OSC_REGION")
		p.Endpoints.API = os.Getenv("OSC_ENDPOINT_API")
		p.Endpoints.OKS = os.Getenv("OSC_ENDPOINT_OKS")
		p.Endpoints.LBU = os.Getenv("OSC_ENDPOINT_LBU")
		p.Endpoints.OOS = os.Getenv("OSC_ENDPOINT_OOS")
		p.Endpoints.FCU = os.Getenv("OSC_ENDPOINT_FCU")
		p.Endpoints.EIM = os.Getenv("OSC_ENDPOINT_EIM")
		p.Endpoints.DirectLink = os.Getenv("OSC_ENDPOINT_DIRECT_LINK")

		if iamV2Services, ok := os.LookupEnv("OSC_IAM_V2_SERVICES"); ok {
			p.IAMV2Services = strings.Split(iamV2Services, ",")
		}

		return nil
	}
}

// MergeWith is a profile option that merges any previously loaded profile with
// another.
func MergeWith(opt Option) Option {
	return func(p *Profile) error {
		var other Profile
		err := opt(&other)
		if err != nil {
			return err
		}
		return mergo.Merge(p, other)
	}
}

// New loads the profile from the environment or a profile file.
// The profile file name is fetched from the OSC_CONFIG_FILE env var, or `~/.osc/config.json` if not set.
// The profile name is fetched from OSC_PROFILE, or `default` if not set.
func New(opts ...Option) (*Profile, error) {
	if len(opts) == 0 {
		opts = []Option{FromEnv(), MergeWith(FromFile("", ""))}
	}
	var p Profile
	for _, opt := range opts {
		err := opt(&p)
		if err != nil {
			return nil, err
		}
	}
	if p.Protocol == "" {
		p.Protocol = "https"
	}

	if p.Region == "" {
		p.Region = "eu-west-2"
	}
	return &p, nil
}

// NewFrom loads the profile from the environment or a profile file.
// If empty, the profile file name is fetched from the OSC_CONFIG_FILE env var, or `~/.osc/config.json` if not set.
// If empty, the profile name is fetched from OSC_PROFILE, or `default` if not set.
func NewFrom(profile, path string) (*Profile, error) {
	return New(FromEnv(), MergeWith(FromFile(profile, path)))
}
