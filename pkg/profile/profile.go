// Package profile provide functions to manage standard Outscale profiles.
//
// The profile can be loaded from environment variables or from a JSON configuration file.
// The configuration file is located at ~/.osc/config.json by
package profile

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Fields struct {
	AccessKey         string                `json:"access_key,omitempty" mergegroup:"aksk"`
	SecretKey         string                `json:"secret_key,omitempty" log:"sensitive" mergegroup:"aksk"`
	AccessKeyV2       string                `json:"access_key_v2,omitempty" mergegroup:"aksk_iamv2"`
	SecretKeyV2       string                `json:"secret_key_v2,omitempty" log:"sensitive" mergegroup:"aksk_iamv2"`
	IAMV2Services     []string              `json:"iam_v2_services,omitempty"`
	X509ClientCert    string                `json:"x509_client_cert,omitempty" mergegroup:"x509keypair"`                       // Deprecated
	X509ClientKey     string                `json:"x509_client_key,omitempty" log:"sensitive" mergegroup:"x509keypair"`        // Deprecated
	X509ClientCertB64 string                `json:"x509_client_cert_b64,omitempty" mergegroup:"x509keypairb64"`                // Deprecated
	X509ClientKeyB64  string                `json:"x509_client_key_b64,omitempty" log:"sensitive" mergegroup:"x509keypairb64"` // Deprecated
	TlsSkipVerify     *bool                 `json:"tls_skip_verify,omitempty"`
	Login             string                `json:"login,omitempty" mergegroup:"basicauth"`
	Password          string                `json:"password,omitempty" log:"sensitive" mergegroup:"basicauth"`
	Protocol          string                `json:"protocol,omitempty"`
	Region            string                `json:"region,omitempty"`
	Endpoints         map[OscService]string `json:"endpoints,omitzero"`
}

type fieldsGetter struct {
	AccessKey         func() (string, error)
	SecretKey         func() (string, error)
	AccessKeyV2       func() (string, error)
	SecretKeyV2       func() (string, error)
	IAMV2Services     func() ([]string, error)
	X509ClientCert    func() (string, error)
	X509ClientKey     func() (string, error)
	X509ClientCertB64 func() (string, error)
	X509ClientKeyB64  func() (string, error)
	TlsSkipVerify     func() (*bool, error)
	Login             func() (string, error)
	Password          func() (string, error)
	Protocol          func() (string, error)
	Region            func() (string, error)
	Endpoints         func() (map[OscService]string, error)
	SourceName        string
}

type Options struct {
	FilePath    *string
	ProfileName *string
	Overrides   Fields

	// ReadFile defaults to os.ReadFile.
	ReadFile func(string) ([]byte, error)

	defaultRegion func() (string, error)
}

func NewOptions() Options {
	return Options{}
}

// WithDefaultRegion resolves the region only when no configured or static default
// region is set. An empty result uses the SDK default; an error aborts loading.
// A nil resolver disables the option.
func WithDefaultRegion(resolve func() (string, error)) func(*Options) error {
	return func(o *Options) error {
		o.defaultRegion = resolve
		return nil
	}
}

func NewOptionsWith(opts ...func(*Options) error) (Options, error) {
	o := NewOptions()

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return o, err
		}
	}
	return o, nil
}

type Profile struct {
	Name    string `json:"-"`
	Default bool   `json:"default,omitempty"`
	Fields
}

// ProfileExplanation records the sources of resolved fields and the config file read.
type ProfileExplanation struct {
	Sources map[string]string
	File    *string
}

func New() (*Profile, error) {
	return NewWith()
}

func NewWith(opts ...func(*Options) error) (*Profile, error) {
	o, err := NewOptionsWith(opts...)
	if err != nil {
		return nil, err
	}

	p, err := load(o, nil)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// NewFrom loads the named profile from path. Empty arguments use the usual
// environment and default selection rules.
func NewFrom(name, path string) (*Profile, error) {
	return NewWith(func(o *Options) error {
		if path != "" {
			o.FilePath = &path
		}

		if name != "" {
			o.ProfileName = &name
		}

		return nil
	})
}

var endpointEnvVars = map[OscService]string{
	OscServiceApi:        "OSC_ENDPOINT_API",
	OscServiceOKS:        "OSC_ENDPOINT_OKS",
	OscServiceFCU:        "OSC_ENDPOINT_FCU",
	OscServiceLBU:        "OSC_ENDPOINT_LBU",
	OscServiceOOS:        "OSC_ENDPOINT_OOS",
	OscServiceEIM:        "OSC_ENDPOINT_EIM",
	OscServiceDirectLink: "OSC_ENDPOINT_DIRECTLINK",
}

func layerFromEnv() fieldsGetter {
	var fields fieldsGetter

	getenv := func(envVar string) func() (string, error) {
		return func() (string, error) {
			return os.Getenv(envVar), nil
		}
	}

	fields.AccessKey = getenv("OSC_ACCESS_KEY")
	fields.SecretKey = getenv("OSC_SECRET_KEY")
	fields.AccessKeyV2 = getenv("OSC_ACCESS_KEY_V2")
	fields.SecretKeyV2 = getenv("OSC_SECRET_KEY_V2")
	fields.X509ClientCert = getenv("OSC_X509_CLIENT_CERT")
	fields.X509ClientCertB64 = getenv("OSC_X509_CLIENT_CERT_B64")
	fields.X509ClientKey = getenv("OSC_X509_CLIENT_KEY")
	fields.X509ClientKeyB64 = getenv("OSC_X509_CLIENT_KEY_B64")
	fields.TlsSkipVerify = func() (*bool, error) {
		v, ok := os.LookupEnv("OSC_TLS_SKIP_VERIFY")
		if !ok {
			return nil, nil
		}

		// EqualFold alone also matches the long-s lookalike in "yeſ".
		return new(strings.EqualFold(v, "true") || (len(v) == len("yes") && strings.EqualFold(v, "yes"))), nil
	}
	fields.Login = getenv("OSC_LOGIN")
	fields.Password = getenv("OSC_PASSWORD")
	fields.Protocol = getenv("OSC_PROTOCOL")
	fields.Region = getenv("OSC_REGION")
	fields.Endpoints = func() (map[OscService]string, error) {
		endpoints := make(map[OscService]string)
		for svc, envVar := range endpointEnvVars {
			if v := os.Getenv(envVar); v != "" {
				endpoints[svc] = v
			}
		}

		return endpoints, nil
	}
	fields.IAMV2Services = func() ([]string, error) {
		if v := os.Getenv("OSC_IAM_V2_SERVICES"); v != "" {
			return strings.Split(v, ","), nil
		}

		return nil, nil
	}
	fields.SourceName = "environment"

	return fields
}

func layerFromFields(sourceName string, f *Fields) fieldsGetter {
	return fieldsGetter{
		AccessKey:         func() (string, error) { return f.AccessKey, nil },
		SecretKey:         func() (string, error) { return f.SecretKey, nil },
		AccessKeyV2:       func() (string, error) { return f.AccessKeyV2, nil },
		SecretKeyV2:       func() (string, error) { return f.SecretKeyV2, nil },
		IAMV2Services:     func() ([]string, error) { return f.IAMV2Services, nil },
		X509ClientCert:    func() (string, error) { return f.X509ClientCert, nil },
		X509ClientKey:     func() (string, error) { return f.X509ClientKey, nil },
		X509ClientCertB64: func() (string, error) { return f.X509ClientCertB64, nil },
		X509ClientKeyB64:  func() (string, error) { return f.X509ClientKeyB64, nil },
		TlsSkipVerify:     func() (*bool, error) { return f.TlsSkipVerify, nil },
		Login:             func() (string, error) { return f.Login, nil },
		Password:          func() (string, error) { return f.Password, nil },
		Protocol:          func() (string, error) { return f.Protocol, nil },
		Region:            func() (string, error) { return f.Region, nil },
		Endpoints:         func() (map[OscService]string, error) { return f.Endpoints, nil },
		SourceName:        sourceName,
	}
}

// Load resolves a profile and reports the origin of its fields.
func Load(o Options) (Profile, ProfileExplanation, error) {
	var explanation ProfileExplanation
	p, err := load(o, &explanation)
	if err != nil {
		return Profile{}, ProfileExplanation{}, err
	}
	return p, explanation, nil
}

func load(o Options, explanation *ProfileExplanation) (Profile, error) {
	cf, explicitFile, err := loadConfigFile(o)
	if err != nil {
		return Profile{}, fmt.Errorf("could not load config file: %w", err)
	}

	name, explicitProfile := choose(o.ProfileName, ProfileNameEnvName)
	if !explicitProfile {
		name = cf.defaultName()
	}
	saved, err := cf.Profile(name)
	var file fieldsGetter
	switch {
	case err == nil:
		file = layerFromFields("profile:"+name, &saved.Fields)
	case explicitFile || explicitProfile || !errors.Is(err, ErrProfileNotFound):
		return Profile{}, fmt.Errorf("loading profile %q: %w", name, err)
	}

	var sources map[string]string
	if explanation != nil {
		sources = make(map[string]string)
	}
	r := Profile{Name: name, Default: saved.Default}
	layers := [...]fieldsGetter{
		layerFromFields("tool", &o.Overrides),
		layerFromEnv(),
		file,
		{
			Protocol:      func() (string, error) { return "https", nil },
			TlsSkipVerify: func() (*bool, error) { return new(false), nil },
			Region: func() (string, error) {
				if o.defaultRegion != nil {
					region, err := o.defaultRegion()
					if err != nil {
						return "", fmt.Errorf("resolving default region: %w", err)
					}
					return region, nil
				}
				return "eu-west-2", nil
			},
			SourceName: "default",
		},
	}
	if explicitFile || explicitProfile {
		layers[1], layers[2] = layers[2], layers[1]
	}
	if err := merge(&r.Fields, sources, layers[:]...); err != nil {
		return Profile{}, fmt.Errorf("merging layers: %w", err)
	}

	if err := merge(&r.Fields, sources, layerFromComputedDefaults(&r.Fields)); err != nil {
		return Profile{}, fmt.Errorf("merging layer default (computed): %w", err)
	}
	if explanation != nil {
		*explanation = ProfileExplanation{Sources: sources, File: cf.FilePath}
	}

	return r, nil
}

func layerFromComputedDefaults(fields *Fields) fieldsGetter {
	endpoints := make(map[OscService]string)
	endpoints[OscServiceApi] = fmt.Sprintf("%s://api.%s.outscale.com/api/v1", fields.Protocol, fields.Region)
	endpoints[OscServiceFCU] = fmt.Sprintf("%s://fcu.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceOOS] = fmt.Sprintf("%s://oos.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceLBU] = fmt.Sprintf("%s://lbu.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceEIM] = fmt.Sprintf("%s://eim.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceDirectLink] = fmt.Sprintf("%s://directlink.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceOKS] = fmt.Sprintf("%s://api.%s.oks.outscale.com/api/v2", fields.Protocol, fields.Region)
	return fieldsGetter{
		Endpoints: func() (map[OscService]string, error) {
			return endpoints, nil
		},
		SourceName: "default (computed)",
	}
}
