// Package profile provide functions to manage standard Outscale profiles.
//
// The profile can be loaded from environment variables or from a JSON configuration file.
// The configuration file is located at ~/.osc/config.json by
package profile

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"

	"dario.cat/mergo"
)

const (
	defaultProfile = "default"
)

type configFile struct {
	profiles map[string]Profile
}

type Profile struct {
	AccessKey         string   `json:"access_key"`
	SecretKey         string   `json:"secret_key"`
	AccessKeyV2       string   `json:"access_key_v2"`
	SecretKeyV2       string   `json:"secret_key_v2"`
	IAMV2Services     []string `json:"iam_v2_services"`
	X509ClientCert    string   `json:"x509_client_cert"`
	X509ClientCertB64 string   `json:"x509_client_cert_b64"`
	X509ClientKey     string   `json:"x509_client_key"`
	X509ClientKeyB64  string   `json:"x509_client_key_b64"`
	TlsSkipVerify     bool     `json:"tls_skip_verify"`
	Login             string
	Password          string
	Protocol          string   `json:"protocol"`
	Region            string   `json:"region"`
	Endpoints         Endpoint `json:"endpoints"`
}

type Option func(*Profile) error

func newConfigFile() *configFile {
	return &configFile{
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

func getenvBool(key string) bool {
	env := os.Getenv(key)
	return env == "yes" || env == "Yes" || env == "true" || env == "True"
}

func LoadProfileFromEnv() Profile {
	var profile Profile

	profile.AccessKey = os.Getenv("OSC_ACCESS_KEY")
	profile.SecretKey = os.Getenv("OSC_SECRET_KEY")
	profile.AccessKeyV2 = os.Getenv("OSC_ACCESS_KEY_V2")
	profile.SecretKeyV2 = os.Getenv("OSC_SECRET_KEY_V2")
	profile.X509ClientCert = os.Getenv("OSC_X509_CLIENT_CERT")
	profile.X509ClientCertB64 = os.Getenv("OSC_X509_CLIENT_CERT_B64")
	profile.X509ClientKey = os.Getenv("OSC_X509_CLIENT_KEY")
	profile.X509ClientKeyB64 = os.Getenv("OSC_X509_CLIENT_KEY_B64")
	profile.TlsSkipVerify = getenvBool("OSC_TLS_SKIP_VERIFY")
	profile.Login = os.Getenv("OSC_LOGIN")
	profile.Password = os.Getenv("OSC_PASSWORD")
	profile.Protocol = os.Getenv("OSC_PROTOCOL")
	profile.Region = os.Getenv("OSC_REGION")
	profile.Endpoints.API = os.Getenv("OSC_ENDPOINT_API")
	profile.Endpoints.OKS = os.Getenv("OSC_ENDPOINT_OKS")
	profile.Endpoints.LBU = os.Getenv("OSC_ENDPOINT_LBU")
	profile.Endpoints.OOS = os.Getenv("OSC_ENDPOINT_OOS")
	profile.Endpoints.FCU = os.Getenv("OSC_ENDPOINT_FCU")
	profile.Endpoints.EIM = os.Getenv("OSC_ENDPOINT_EIM")
	profile.Endpoints.DirectLink = os.Getenv("OSC_ENDPOINT_DIRECT_LINK")

	if iamV2Services, ok := os.LookupEnv("OSC_IAM_V2_SERVICES"); ok {
		profile.IAMV2Services = strings.Split(iamV2Services, ",")
	}

	return profile
}

// New loads the profile from the environment or a profile file.
// The profile file name is fetched from the OSC_CONFIG_FILE env var, or `~/.osc/config.json` if not set.
// The profile name is fetched from OSC_PROFILE, or `default` if not set.
func New(opts ...Option) (*Profile, error) {
	var profile string
	if value, present := os.LookupEnv("OSC_PROFILE"); present {
		profile = value
	} else {
		profile = defaultProfile
	}

	var path string
	if value, present := os.LookupEnv("OSC_CONFIG_FILE"); present {
		path = value
	} else {
		path, _ = defaultConfigPath()
	}
	return NewFrom(profile, path, opts...)
}

// NewFrom loads the profile from the environment or a profile file.
func NewFrom(profile, path string, opts ...Option) (*Profile, error) {
	// 1. Load profile from environment
	mergedProfile := LoadProfileFromEnv()

	// 2. Load profile for config file and merge
	configFile, err := loadConfigFile(path)
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

	// 3. Apply options
	for _, opt := range opts {
		err := opt(&mergedProfile)
		if err != nil {
			return nil, err
		}
	}

	// 4. Set defaults
	if mergedProfile.Protocol == "" {
		mergedProfile.Protocol = "https"
	}

	if mergedProfile.Region == "" {
		mergedProfile.Region = "eu-west-2"
	}

	return &mergedProfile, nil
}
