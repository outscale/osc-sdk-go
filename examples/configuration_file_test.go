/*
BSD 3-Clause License

Copyright (c) 2022, Outscale SAS
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

  - Redistributions of source code must retain the above copyright notice, this
    list of conditions and the following disclaimer.

  - Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

  - Neither the name of the copyright holder nor the names of its
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
package osc_test

import (
	"context"
	"fmt"
	"os"

	"github.com/outscale/osc-sdk-go/osc"
)

// This example shows how load credentials from a local configuration file.
func ExampleConfigurationFile() {
	// When running those examples, configuration file may not exist.
	// It can be manually created or generated depending of your application.
	// Here we are creating it for simplicity sake.
	configPath := "/tmp/osc-sdk-go-ExampleConfigurationFile"
	createConfigurationFile(configPath)

	// Load configuation from default location
	// You can also use osc.LoadDefaultConfigFile() (~/.osc/config.json by default on Linux or MacOS)
	configFile, err := osc.LoadConfigFile(&configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading default configuration file: %s", err.Error())
		os.Exit(1)
	}
	config, err := configFile.Configuration("default")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating configuration: %s", err.Error())
		os.Exit(1)
	}
	ctx, err := configFile.Context(context.Background(), "default")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating context: %s", err.Error())
		os.Exit(1)
	}
	client := osc.NewAPIClient(config)
	_, httpRes, err := client.SecurityGroupApi.ReadSecurityGroups(ctx, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading security groups")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status, httpRes.Body)
		}
		os.Exit(1)
	}
	fmt.Println("configuration file journey is over")
	// Output: configuration file journey is over
}

func createConfigurationFile(configPath string) {
	ak := os.Getenv("OSC_ACCESS_KEY")
	sk := os.Getenv("OSC_SECRET_KEY")
	region := os.Getenv("OSC_REGION")
	content := fmt.Sprintf(`{
	"default": {
		"access_key": "%s",
		"secret_key": "%s",
		"region": "%s"
	}}`, ak, sk, region)
	jsonFile, _ := os.Create(configPath)
	defer jsonFile.Close()
	jsonFile.WriteString(content)
}
