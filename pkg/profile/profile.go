// Package profile provide functions to manage standard Outscale profiles.
//
// The profile can be loaded from environment variables or from a JSON configuration file.
// The configuration file is located at ~/.osc/config.json by
package profile

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fields struct {
	AccessKey         string                `json:"access_key,omitempty" mergegroup:"aksk"`
	SecretKey         string                `json:"secret_key,omitempty" log:"sensitive" mergegroup:"aksk"`
	AccessKeyV2       string                `json:"access_key_v2,omitempty" mergegroup:"aksk_iamv2"`
	SecretKeyV2       string                `json:"secret_key_v2,omitempty" log:"sensitive" mergegroup:"aksk_iamv2"`
	IAMV2Services     []string              `json:"iam_v2_services,omitempty"`
	X509ClientCert    string                `json:"x509_client_cert,omitempty" mergegroup:"x509keypair"`
	X509ClientKey     string                `json:"x509_client_key,omitempty" log:"sensitive" mergegroup:"x509keypair"`
	X509ClientCertB64 string                `json:"x509_client_cert_b64,omitempty" mergegroup:"x509keypairb64"`
	X509ClientKeyB64  string                `json:"x509_client_key_b64,omitempty" log:"sensitive" mergegroup:"x509keypairb64"`
	TlsSkipVerify     *bool                 `json:"tls_skip_verify,omitempty"`
	Login             string                `json:"login,omitempty" mergegroup:"basicauth"`
	Password          string                `json:"password,omitempty" log:"sensitive" mergegroup:"basicauth"`
	Protocol          string                `json:"protocol,omitempty"`
	Region            string                `json:"region,omitempty"`
	Endpoints         map[OscService]string `json:"endpoints,omitzero"`
}

type FieldsGetter struct {
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

	// **HAZMAT** do not use in production unless you understand the implications.
	Default FieldsGetter
	// Nil mean os.ReadFile. for mocking
	ReadFile func(string) ([]byte, error)
}

func NewOptions() Options {
	return Options{
		Default: FieldsGetter{
			Region:        func() (string, error) { return "eu-west-2", nil },
			Protocol:      func() (string, error) { return "https", nil },
			TlsSkipVerify: func() (*bool, error) { return new(false), nil },
			SourceName:    "default",
		},
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
	Fields
	Sources     map[string]string
	ProfileName string
	File        *string
}

func New() (Profile, error) {
	return NewWith()
}

func NewWith(opts ...func(*Options) error) (Profile, error) {
	o, err := NewOptionsWith(opts...)
	if err != nil {
		return Profile{}, err
	}

	return Load(o)
}

func NewFrom(path, name string) (Profile, error) {
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

func layerFromEnv() FieldsGetter {
	var fields FieldsGetter

	genenv := func(envVar string) func() (string, error) {
		return func() (string, error) {
			return os.Getenv(envVar), nil
		}
	}

	fields.AccessKey = genenv("OSC_ACCESS_KEY")
	fields.SecretKey = genenv("OSC_SECRET_KEY")
	fields.AccessKeyV2 = genenv("OSC_ACCESS_KEY_V2")
	fields.SecretKeyV2 = genenv("OSC_SECRET_KEY_V2")
	fields.X509ClientCert = genenv("OSC_X509_CLIENT_CERT")
	fields.X509ClientCertB64 = genenv("OSC_X509_CLIENT_CERT_B64")
	fields.X509ClientKey = genenv("OSC_X509_CLIENT_KEY")
	fields.X509ClientKeyB64 = genenv("OSC_X509_CLIENT_KEY_B64")
	fields.TlsSkipVerify = func() (*bool, error) {
		v, ok := os.LookupEnv("OSC_TLS_SKIP_VERIFY")
		if ok {
			b, err := strconv.ParseBool(v)
			if err != nil {
				return nil, err
			}
			return &b, nil
		}

		return nil, nil
	}
	fields.Login = genenv("OSC_LOGIN")
	fields.Password = genenv("OSC_PASSWORD")
	fields.Protocol = genenv("OSC_PROTOCOL")
	fields.Region = genenv("OSC_REGION")
	fields.Endpoints = func() (map[OscService]string, error) {
		endpoints := make(map[OscService]string)
		for svc, envVar := range endpointEnvVars {
			v, ok := os.LookupEnv(envVar)
			if ok {
				endpoints[svc] = v
			}
		}

		return endpoints, nil
	}
	fields.IAMV2Services = func() ([]string, error) {
		if v, ok := os.LookupEnv("OSC_IAM_V2_SERVICES"); ok && v != "" {
			return strings.Split(v, ","), nil
		}

		return nil, nil
	}
	fields.SourceName = "environment"

	return fields
}

func getterFromFields(sourceName string, f Fields, err error) FieldsGetter {
	return FieldsGetter{
		AccessKey:         func() (string, error) { return f.AccessKey, err },
		SecretKey:         func() (string, error) { return f.SecretKey, err },
		AccessKeyV2:       func() (string, error) { return f.AccessKeyV2, err },
		SecretKeyV2:       func() (string, error) { return f.SecretKeyV2, err },
		IAMV2Services:     func() ([]string, error) { return f.IAMV2Services, err },
		X509ClientCert:    func() (string, error) { return f.X509ClientCert, err },
		X509ClientKey:     func() (string, error) { return f.X509ClientKey, err },
		X509ClientCertB64: func() (string, error) { return f.X509ClientCertB64, err },
		X509ClientKeyB64:  func() (string, error) { return f.X509ClientKeyB64, err },
		TlsSkipVerify:     func() (*bool, error) { return f.TlsSkipVerify, err },
		Login:             func() (string, error) { return f.Login, err },
		Password:          func() (string, error) { return f.Password, err },
		Protocol:          func() (string, error) { return f.Protocol, err },
		Region:            func() (string, error) { return f.Region, err },
		Endpoints:         func() (map[OscService]string, error) { return f.Endpoints, err },
		SourceName:        sourceName,
	}
}

func layerFromConfigFile(cf *ConfigFile) FieldsGetter {
	name, fields, err := cf.SelectedProfile()
	return getterFromFields("profile:"+name, fields, err)
}

func Load(o Options) (Profile, error) {
	if o.ReadFile == nil {
		o.ReadFile = os.ReadFile
	}

	return load(o)
}

func load(o Options) (Profile, error) {
	// 1. Load Config file and profile
	cf, err := loadConfigFile(o)
	if err != nil {
		return Profile{}, fmt.Errorf("could not load config file: %w", err)
	}

	// 4. Layer definition and ordering
	tool := getterFromFields("tool", o.Overrides, nil)
	file := layerFromConfigFile(cf)
	environment := layerFromEnv()

	layers := []FieldsGetter{tool, environment, file, o.Default}
	if cf.explicitProfile || cf.explicitFile {
		layers = []FieldsGetter{tool, file, environment, o.Default}
	}

	// 5. layer merging
	r := Profile{
		Sources:     map[string]string{},
		ProfileName: cf.SelectedProfileName,
		File:        cf.FilePath,
	}

	if err := Merge(&r.Fields, &r.Sources, layers...); err != nil {
		return Profile{}, fmt.Errorf("merging merging layers: %w", err)
	}

	// 6. Add computed default
	computedDefault := computedDefault(&r.Fields)
	if err := Merge(&r.Fields, &r.Sources, computedDefault); err != nil {
		return Profile{}, fmt.Errorf("merging layer default (computed): %w", err)
	}

	return r, nil
}

func computedDefault(fields *Fields) FieldsGetter {
	endpoints := make(map[OscService]string)
	endpoints[OscServiceApi] = fmt.Sprintf("%s://api.%s.outscale.com/api/v1", fields.Protocol, fields.Region)
	endpoints[OscServiceFCU] = fmt.Sprintf("%s://fcu.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceOOS] = fmt.Sprintf("%s://oos.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceLBU] = fmt.Sprintf("%s://lbu.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceEIM] = fmt.Sprintf("%s://eim.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceDirectLink] = fmt.Sprintf("%s://directlink.%s.outscale.com", fields.Protocol, fields.Region)
	endpoints[OscServiceOKS] = fmt.Sprintf("%s://api.%s.oks.outscale.com/api/v2", fields.Protocol, fields.Region)
	return FieldsGetter{
		Endpoints: func() (map[OscService]string, error) {
			return endpoints, nil
		},
		SourceName: "default (computed)",
	}
}
