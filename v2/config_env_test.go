package osc

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"
)

func TestEnvVariablesAkSk(t *testing.T) {
	configEnv := NewConfigEnv()
	config, err := configEnv.Configuration()
	if err != nil {
		t.Fatalf("Cannot create configutation: %s", err.Error())
	}
	ctx, err := configEnv.Context(context.Background())
	if err != nil {
		t.Fatalf("Cannot create context for making a query: %s", err.Error())
	}
	testConfAndContextOk(t, config, &ctx)
}

func TestEnvVariablesWithProfile(t *testing.T) {
	if err := os.Setenv("OSC_PROFILE", "SomeProfile"); err != nil {
		t.Fatalf("Cannot set OSC_PROFILE: %s", err.Error())
	}
	ak := os.Getenv("OSC_ACCESS_KEY")
	sk := os.Getenv("OSC_SECRET_KEY")
	region := os.Getenv("OSC_REGION")
	os.Unsetenv("OSC_ACCESS_KEY")
	os.Unsetenv("OSC_SECRET_KEY")
	os.Unsetenv("OSC_REGION")
	defer os.Setenv("OSC_ACCESS_KEY", ak)
	defer os.Setenv("OSC_SECRET_KEY", sk)
	defer os.Setenv("OSC_REGION", region)
	endpoint, endpoint_found := os.LookupEnv("OSC_ENDPOINT_API")
	if !endpoint_found {
		endpoint = fmt.Sprintf("api.%s.outscale.com/api/v1", region)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Cannot get user home dir: %s", err.Error())
	}
	configFolderPath := path.Join(home, ".osc")
	configPath := path.Join(configFolderPath, "config.json")

	os.RemoveAll(configFolderPath)
	if err := os.Mkdir(configFolderPath, os.ModePerm); err != nil {
		t.Fatalf("Cannot create .osc folder: %s", err.Error())
	}
	defer os.RemoveAll(configFolderPath)

	content := fmt.Sprintf(`{
		"SomeProfile": {
			"access_key": "%s",
			"secret_key": "%s",
			"endpoints": {
				"api": "%s"
			}
		}}`, ak, sk, endpoint)
	if err := testDumpToFile(configPath, content); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	configEnv := NewConfigEnv()
	config, err := configEnv.Configuration()
	if err != nil {
		t.Fatalf("Cannot create configutation: %s", err.Error())
	}
	ctx, err := configEnv.Context(context.Background())
	if err != nil {
		t.Fatalf("Cannot create context for making a query: %s", err.Error())
	}
	testConfAndContextOk(t, config, &ctx)
}

func testConfAndContextOk(t *testing.T, config *Configuration, ctx *context.Context) {
	client := NewAPIClient(config)
	_, _, err := client.SubregionApi.ReadSubregions(*ctx).ReadSubregionsRequest(ReadSubregionsRequest{}).Execute()
	if err != nil {
		t.Fatalf("Cannot call ReadSubregions: %s", err.Error())
	}
}
