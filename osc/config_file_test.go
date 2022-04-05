/*
  BSD 3-Clause License

  Copyright (c) 2022, Outscale SAS
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:

  * Redistributions of source code must retain the above copyright notice, this
    list of conditions and the following disclaimer.

  * Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

  * Neither the name of the copyright holder nor the names of its
    contributors may be used to endorse or promote products derived from
    this software without specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
  DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
  OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

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
	_, _, err = client.SubregionApi.ReadSubregions(ctx, nil)
	if err != nil {
		t.Fatalf("Error while reading subregions: %s", err.Error())
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
