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
	"errors"
	"os"
)

type ConfigEnv struct {
	AccessKey           *string
	SecretKey           *string
	OutscaleApiEndpoint *string
	ProfileName         *string
	Region              *string
}

func NewConfigEnv() *ConfigEnv {
	var configEnv ConfigEnv
	if value, present := os.LookupEnv("OSC_ACCESS_KEY"); present {
		configEnv.AccessKey = &value
	}
	if value, present := os.LookupEnv("OSC_SECRET_KEY"); present {
		configEnv.SecretKey = &value
	}
	if value, present := os.LookupEnv("OSC_ENDPOINT_API"); present {
		configEnv.OutscaleApiEndpoint = &value
	}
	if value, present := os.LookupEnv("OSC_PROFILE"); present {
		configEnv.ProfileName = &value
	}
	if value, present := os.LookupEnv("OSC_REGION"); present {
		configEnv.Region = &value
	}
	return &configEnv
}

func (configEnv *ConfigEnv) Configuration() (*Configuration, error) {
	var config *Configuration

	if configEnv.ProfileName != nil {
		configFile, err := LoadDefaultConfigFile()
		if err != nil {
			return nil, err
		}
		config, err = configFile.Configuration(*configEnv.ProfileName)
		if err != nil {
			return nil, err
		}
	} else {
		config = NewConfiguration()
		default_region := "eu-west-2"
		config.Servers = []ServerConfiguration{
			{
				Url:         "https://api.{region}.outscale.com/api/v1",
				Description: "Loaded from profile",
				Variables: map[string]ServerVariable{
					"region": ServerVariable{
						Description:  "Loaded from env variables",
						DefaultValue: default_region,
						EnumValues:   []string{default_region},
					},
				},
			},
		}
	}

	// Overload with remaining environement variable values
	if configEnv.OutscaleApiEndpoint != nil {
		config.Servers[0].Url = *configEnv.OutscaleApiEndpoint
	}

	if configEnv.Region != nil && len(config.Servers) > 0 {
		config.Servers[0].Variables["region"] = ServerVariable{
			Description:  "Loaded from env variables",
			DefaultValue: *configEnv.Region,
			EnumValues:   []string{*configEnv.Region},
		}
	}
	return config, nil
}

func (configEnv *ConfigEnv) Context(ctx context.Context) (context.Context, error) {
	var accessKey *string
	var secretKey *string

	if configEnv.ProfileName != nil {
		configFile, err := LoadDefaultConfigFile()
		if err != nil {
			return nil, err
		}
		profile, ok := configFile.profiles[*configEnv.ProfileName]
		if !ok {
			return nil, errors.New("profile not found for creating Context, did you load config file?")
		}
		accessKey = &profile.AccessKey
		secretKey = &profile.SecretKey
	}
	// Overload with environement variable values
	if configEnv.AccessKey != nil {
		accessKey = configEnv.AccessKey
	}
	if configEnv.SecretKey != nil {
		secretKey = configEnv.SecretKey
	}

	if accessKey == nil || len(*accessKey) == 0 || secretKey == nil || len(*secretKey) == 0 {
		return ctx, nil
	}
	ctx = context.WithValue(ctx, ContextAWSv4, AWSv4{
		AccessKey: *accessKey,
		SecretKey: *secretKey,
	})
	return ctx, nil
}
