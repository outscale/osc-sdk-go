package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"dario.cat/mergo"
)

const (
	defaultProfile = "default"
)

type ConfigFile struct {
	profiles map[string]Profile
}

type Profile struct {
	AccessKey         string `json:"access_key"`
	SecretKey         string `json:"secret_key"`
	X509ClientCert    string `json:"x509_client_cert"`
	X509ClientCertB64 string `json:"x509_client_cert_b64"`
	X509ClientKey     string `json:"x509_client_key"`
	X509ClientKeyB64  string `json:"x509_client_key_b64"`
	Login             string
	Password          string
	Protocol          string   `json:"protocol"`
	Region            string   `json:"region"`
	Endpoints         endpoint `json:"endpoints"`
}

type endpoint struct {
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

func (p *Profile) getDefaultEndpoint(service OscService) (string, error) {
	temp, err := getDefaultEndpointTemplate(service)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(temp, p.Protocol, p.Region), nil
}

func (p *Profile) GetEndpoint(service OscService) (string, error) {
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

func NewConfigFile() *ConfigFile {
	return &ConfigFile{
		profiles: make(map[string]Profile),
	}
}

func defaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".osc", "config.json"), nil
}

func LoadConfigFile(path string) (*ConfigFile, error) {
	if path == "" {
		return nil, errors.New("no path provided")
	}

	configJSON, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configFile := NewConfigFile()
	if err := json.Unmarshal(configJSON, &configFile.profiles); err != nil {
		return nil, err
	}

	return configFile, nil
}

func LoadProfileFromEnv() Profile {
	var profile Profile

	profile.AccessKey = os.Getenv("OSC_ACCESS_KEY")
	profile.SecretKey = os.Getenv("OSC_SECRET_KEY")
	profile.X509ClientCert = os.Getenv("OSC_X509_CLIENT_CERT")
	profile.X509ClientCertB64 = os.Getenv("OSC_X509_CLIENT_CERT_B64")
	profile.X509ClientKey = os.Getenv("OSC_X509_CLIENT_KEY")
	profile.X509ClientKeyB64 = os.Getenv("OSC_X509_CLIENT_KEY_B64")
	profile.Login = os.Getenv("OSC_LOGIN")
	profile.Password = os.Getenv("OSC_PASSWORD")
	profile.Protocol = os.Getenv("OSC_PROTOCOL")
	profile.Region = os.Getenv("OSC_REGION")
	profile.Endpoints.API = os.Getenv("OSC_ENDPOINT_API")
	profile.Endpoints.LBU = os.Getenv("OSC_ENDPOINT_LBU")
	profile.Endpoints.OKS = os.Getenv("OSC_ENDPOINT_OKS")

	return profile
}

func NewProfileFromStrandardConfuguration(profile, path string) (*Profile, error) {
	// 1. Load profile from environment
	mergedProfile := LoadProfileFromEnv()

	// 2. Load additional config from environment
	if profile == "" {
		if value, present := os.LookupEnv("OSC_PROFILE"); present {
			profile = value
		} else {
			profile = defaultProfile
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
	configFile, err := LoadConfigFile(path)
	if err == nil {
		if fileprofile, ok := configFile.profiles[profile]; ok {
			err := mergo.Merge(&mergedProfile, fileprofile)
			if err != nil {
				return nil, err
			}
		} else if profile != defaultProfile {
			return nil, errors.New("specified profile not found")
		}
	}

	// 4. Load default
	if mergedProfile.Protocol == "" {
		mergedProfile.Protocol = "https"
	}

	if mergedProfile.Region == "" {
		mergedProfile.Region = "eu-west-2"
	}

	return &mergedProfile, nil
}
