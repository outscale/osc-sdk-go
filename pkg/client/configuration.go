package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"dario.cat/mergo"
	oks "github.com/outscale/osc-sdk-go/v3/internal/oks"
	osc "github.com/outscale/osc-sdk-go/v3/internal/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/securityprovider"
)

type OscService int

const (
	OApi OscService = iota
	LBU
	OKS
)

type configFile struct {
	profiles map[string]ClientBuilder
}

type ClientBuilder struct {
	AccessKey         string   `json:"access_key"`
	SecretKey         string   `json:"secret_key"`
	X509ClientCert    string   `json:"x509_client_cert"`
	X509ClientCertB64 string   `json:"x509_client_cert_b64"`
	X509ClientKey     string   `json:"x509_client_key"`
	X509ClientKeyB64  string   `json:"x509_client_key_b64"`
	Protocol          string   `json:"protocol"`
	Region            string   `json:"region"`
	Endpoints         Endpoint `json:"endpoints"`
}

type Endpoint struct {
	API string `json:"api"`
	LBU string `json:"lbu"`
	OKS string `json:"oks"`
}

func getDefaultEndpointTemplate(service OscService) (string, error) {
	switch service {
	case OApi:
		return "%s://api.%s.outscale.com/api/v1", nil
	case LBU:
		return "%s://lbu.%s.outscale.com", nil
	case OKS:
		return "%s://api.%s.oks.outscale.com/api/v2", nil
	default:
		return "", errors.New("unsupported service")
	}
}

func (p *ClientBuilder) getDefaultEndpoint(service OscService) (string, error) {
	temp, err := getDefaultEndpointTemplate(service)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(temp, p.Protocol, p.Region), nil
}

func (p *ClientBuilder) GetEndpoint(service OscService) (string, error) {
	var endpoint string

	switch service {
	case OApi:
		endpoint = p.Endpoints.API
	case LBU:
		endpoint = p.Endpoints.LBU
	case OKS:
		endpoint = p.Endpoints.OKS
	}

	if endpoint == "" {
		return p.getDefaultEndpoint(service)
	}

	return endpoint, nil
}

func (p *ClientBuilder) OKS() (*oks.Client, error) {
	aksk, err := securityprovider.NewSecurityProviderAWSv4(
		p.AccessKey,
		p.SecretKey,
		"",
		"oks",
		p.Region,
	)
	if err != nil {
		return nil, err
	}

	server, err := p.GetEndpoint(OKS)
	if err != nil {
		return nil, err
	}

	client, err := oks.NewClient(
		server,
		oks.WithRequestEditorFn(aksk.InterceptOks),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (p *ClientBuilder) OApi() (*osc.Client, error) {
	aksk, err := securityprovider.NewSecurityProviderAWSv4(
		p.AccessKey,
		p.SecretKey,
		"",
		"oapi",
		p.Region,
	)
	if err != nil {
		return nil, err
	}

	server, err := p.GetEndpoint(OApi)
	if err != nil {
		return nil, err
	}

	c := NewClientWithRateLimit(WithClient(NewClientWithRetry()))

	client, err := osc.NewClient(
		server,
		osc.WithRequestEditorFn(aksk.Intercept),
		osc.WithHTTPClient(c),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type CopyOption func(p *ClientBuilder)

func (p *ClientBuilder) Copy(opts ...CopyOption) *ClientBuilder {
	c := ClientBuilder{
		AccessKey:         p.AccessKey,
		SecretKey:         p.SecretKey,
		X509ClientCertB64: p.X509ClientCertB64,
		X509ClientKeyB64:  p.X509ClientKeyB64,
		X509ClientCert:    p.X509ClientCert,
		X509ClientKey:     p.X509ClientKey,
		Protocol:          p.Protocol,
		Region:            p.Region,
		Endpoints:         p.Endpoints,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

func newConfigFile() *configFile {
	return &configFile{
		profiles: make(map[string]ClientBuilder),
	}
}

func defaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".osc", "config.json"), nil
}

func loadConfigFile(path string) (*configFile, error) {
	if path == "" {
		return nil, errors.New("no path provided")
	}

	configJSON, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configFile := newConfigFile()
	if err := json.Unmarshal(configJSON, &configFile.profiles); err != nil {
		return nil, err
	}

	return configFile, nil
}

func LoadProfileFromEnv() ClientBuilder {
	var profile ClientBuilder

	profile.AccessKey = os.Getenv("OSC_ACCESS_KEY")
	profile.SecretKey = os.Getenv("OSC_SECRET_KEY")
	profile.X509ClientCert = os.Getenv("OSC_X509_CLIENT_CERT")
	profile.X509ClientCertB64 = os.Getenv("OSC_X509_CLIENT_CERT_B64")
	profile.X509ClientKey = os.Getenv("OSC_X509_CLIENT_KEY")
	profile.X509ClientKeyB64 = os.Getenv("OSC_X509_CLIENT_KEY_B64")
	profile.Protocol = os.Getenv("OSC_PROTOCOL")
	profile.Region = os.Getenv("OSC_REGION")
	profile.Endpoints.API = os.Getenv("OSC_ENDPOINT_API")
	profile.Endpoints.LBU = os.Getenv("OSC_ENDPOINT_LBU")
	profile.Endpoints.OKS = os.Getenv("OSC_ENDPOINT_OKS")

	return profile
}

func Builder(profile, path string) (*ClientBuilder, error) {
	// 1. Load profile from environment
	azerty := LoadProfileFromEnv()

	// 2. Load additional config from environment
	if profile == "" {
		if value, present := os.LookupEnv("OSC_PROFILE"); present {
			profile = value
		} else {
			profile = "default"
		}
	}

	if path == "" {
		if value, present := os.LookupEnv("OSC_CONFIG_FILE"); present {
			path = value
		} else {
			path, _ = defaultConfigPath()
		}
	}

	// 3. Load profile for config file
	configFile, err := loadConfigFile(path)
	if err != nil {
		if fileprofile, ok := configFile.profiles[profile]; ok {
			err := mergo.Merge(&azerty, fileprofile)
			if err != nil {
				return nil, err
			}
		}
	}

	// 4. Load default
	if azerty.Protocol == "" {
		azerty.Protocol = "https"
	}

	if azerty.Region == "" {
		azerty.Region = "eu-west-2"
	}

	return &azerty, nil
}
