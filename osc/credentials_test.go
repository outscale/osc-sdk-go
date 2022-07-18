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
	"path"
	"testing"
)

func TestWithProfileName(t *testing.T) {
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
				"api": "api.%s.outscale.com/api/v1"
			}
		}}`, ak, sk, region)
	if err := testDumpToFile(configPath, content); err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	var creds Credentials
	config, err := creds.retrieveCreds()
	if err != nil {
		t.Fatalf("Cannot create configutation: %s", err.Error())
	}
	ctx, err := creds.Context(context.Background())
	if err != nil {
		t.Fatalf("Cannot create context for making a query: %s", err.Error())
	}
	client := NewAPIClient(config)
	_, _, err = client.SubregionApi.ReadSubregions(ctx, nil)
	if err != nil {
		t.Fatalf("Cannot call ReadSubregions: %s", err.Error())
	}
}
