package osc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type ConfigFile struct {
	profiles map[string]Profile
}

type Profile struct {
	AccessKey       string   `json:"access_key"`
	SecretKey       string   `json:"secret_key"`
	X509ClientCert  string   `json:"x509_client_cert"`
	X509_client_key string   `json:"x509_client_key"`
	Protocol        string   `json:"protocol"`
	Method          string   `json:"method"`
	Region          string   `json:"region"`
	Endpoints       Endpoint `json:"endpoints"`
}

type Endpoint struct {
	API string `json:"api"`
	FCU string `json:"fcu"`
	LBU string `json:"lbu"`
	EIM string `json:"eim"`
	ICU string `json:"icu"`
	OOS string `json:"oos"`
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

func LoadDefaultConfigFile() (*ConfigFile, error) {
	configPath, err := defaultConfigPath()
	if err != nil {
		return nil, err
	}
	return LoadConfigFile(&configPath)
}

func LoadConfigFile(path *string) (*ConfigFile, error) {
	if path == nil || len(*path) == 0 {
		return nil, errors.New("no path provided")
	}
	config_json, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}
	configFile := NewConfigFile()
	if err := json.Unmarshal(config_json, &configFile.profiles); err != nil {
		return nil, err
	}
	return configFile, nil
}

func profile(profileName string) string {
	if profileName == "" {
		profileName = os.Getenv("OSC_PROFILE")
	}
	if profileName == "" {
		profileName = "default"
	}
	return profileName
}

func (configFile *ConfigFile) Configuration(profileName string) (*Configuration, error) {
	profile, ok := configFile.profiles[profile(profileName)]
	if !ok {
		return nil, errors.New("profile not found for creating Context, did you load config file?")
	}
	var url string
	if len(profile.Endpoints.API) > 0 {
		if len(profile.Protocol) > 0 {
			url = fmt.Sprintf("%s://%s", profile.Protocol, profile.Endpoints.API)
		} else {
			url = fmt.Sprintf("https://%s", profile.Endpoints.API)
		}
	} else {
		url = "https://api.{region}.outscale.com/api/v1"
	}

	var region string
	if len(profile.Region) > 0 {
		region = profile.Region
	} else {
		region = "eu-west-2"
	}

	config := NewConfiguration()
	config.Servers = ServerConfigurations{
		{
			URL:         url,
			Description: "Loaded from profile",
			Variables: map[string]ServerVariable{
				"region": ServerVariable{
					Description:  "Loaded from profile",
					DefaultValue: region,
					EnumValues:   []string{region},
				},
			},
		},
	}
	return config, nil
}

func (configFile *ConfigFile) Context(ctx context.Context, profileName string) (context.Context, error) {
	profile, ok := configFile.profiles[profileName]
	if !ok {
		return nil, errors.New("profile not found for creating Context, did you load config file?")
	}
	if len(profile.AccessKey) > 0 && len(profile.SecretKey) > 0 {
		ctx = context.WithValue(ctx, ContextAWSv4, AWSv4{
			AccessKey: profile.AccessKey,
			SecretKey: profile.SecretKey,
		})
	}
	return ctx, nil
}
