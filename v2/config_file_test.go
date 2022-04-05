package osc

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestBasicConfigFileWithValidEndpoint1(t *testing.T) {
	ak := os.Getenv("OSC_ACCESS_KEY")
	sk := os.Getenv("OSC_SECRET_KEY")
	region := os.Getenv("OSC_REGION")
	content := fmt.Sprintf(`{
	"SomeProfile": {
		"access_key": "%s",
		"secret_key": "%s",
		"protocol": "https",
		"method": "post",
		"region": "bad-region-that-should-not-impact",
		"endpoints": {
			"api": "api.%s.outscale.com/api/v1"
		}
	}}`, ak, sk, region)
	configPath := "/tmp/osc-sdk-go-TestBasicConfigFileWithValidEndpoint1"
	if err := testDumpToFile(configPath, content); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	defer os.Remove(configPath)
	testApiOk(t, configPath)
}

func TestBasicConfigFileWithValidEndpoint2(t *testing.T) {
	ak := os.Getenv("OSC_ACCESS_KEY")
	sk := os.Getenv("OSC_SECRET_KEY")
	region := os.Getenv("OSC_REGION")
	content := fmt.Sprintf(`{
		"SomeProfile": {
			"access_key": "%s",
			"secret_key": "%s",
			"endpoints": {
				"api": "api.%s.outscale.com/api/v1"
			}
		}}`, ak, sk, region)
	configPath := "/tmp/osc-sdk-go-TestBasicConfigFileWithValidEndpoint2"
	if err := testDumpToFile(configPath, content); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	defer os.Remove(configPath)
	testApiOk(t, configPath)
}

func TestBasicConfigFileWithValidRegion2(t *testing.T) {
	ak := os.Getenv("OSC_ACCESS_KEY")
	sk := os.Getenv("OSC_SECRET_KEY")
	region := os.Getenv("OSC_REGION")
	content := fmt.Sprintf(`{
		"SomeProfile": {
			"access_key": "%s",
			"secret_key": "%s",
			"region": "%s"
		}}`, ak, sk, region)
	configPath := "/tmp/osc-sdk-go-TestBasicConfigFileWithValidRegion2"
	if err := testDumpToFile(configPath, content); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	defer os.Remove(configPath)
	testApiOk(t, configPath)
}

func testDumpToFile(path string, content string) error {
	jsonFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	if _, err := jsonFile.WriteString(content); err != nil {
		return err
	}
	return nil
}

func testApiOk(t *testing.T, configPath string) {
	configFile, err := LoadConfigFile(&configPath)
	if err != nil {
		t.Fatalf("Cannot load configuration file: %s", err.Error())
	}
	_, err = configFile.Configuration("non-existant")
	if err == nil {
		t.Fatal("profile 'non-existant' should not exist")
	}
	config, err := configFile.Configuration("SomeProfile")
	if err != nil {
		t.Errorf("profile 'SomeProfile' should exist (ConfigFromProfileEndpoint()): %s", err.Error())
	}
	_, err = configFile.Context(context.Background(), "non-existant")
	if err == nil {
		t.Fatal("profile 'non-existant' should not exist (Context())")
	}
	ctx, err := configFile.Context(context.Background(), "SomeProfile")
	if err != nil {
		t.Fatalf("profile 'SomeProfile' should exist (Context()): %s", err.Error())
	}
	client := NewAPIClient(config)
	_, _, err = client.SubregionApi.ReadSubregions(ctx).ReadSubregionsRequest(ReadSubregionsRequest{}).Execute()
	if err != nil {
		t.Fatalf("Cannot call ReadSubregions: %s", err.Error())
	}
}

func TestBadConfigFile1(t *testing.T) {
	content := `{
		"SomeProfile": {
			Garbage
		}}`
	configPath := "/tmp/osc-sdk-go-TestBadConfigFile2"
	if err := testDumpToFile(configPath, content); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	defer os.Remove(configPath)
	testBadConfigFile(t, configPath)
}

func testBadConfigFile(t *testing.T, configPath string) {
	_, err := LoadConfigFile(&configPath)
	if err == nil {
		t.Fatal("Config file should not have loaded")
	}
}
