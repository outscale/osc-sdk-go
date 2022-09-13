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
	"fmt"
)

type Credentials struct {
	ConfigEnv
}

func (c Credentials) retrieveCreds() (*Configuration, error) {
	configEnv := NewConfigEnv()
	if configEnv.AccessKey != nil && configEnv.SecretKey != nil {
		config, err := configEnv.Configuration(); err!= nil {
			return nil, errors.New("error in loading env variables")
		}
		return config, nil
	} else {
		profileName := *configEnv.ProfileName
		if profileName == "" {
			profileName = "Default"
		}
		fmt.Println("profileName is ", profileName)
		configFile, err := LoadDefaultConfigFile() ; err != nil {
			return nil, err
		}
		config, err := configFile.Configuration(profileName) ; err != nil {
			return nil, errors.New("error in loading variables fromm config file")
		}
		return config, nil
	}
}

func (c *Credentials) Context(ctx context.Context) (context.Context, error) {
	configEnv := *NewConfigEnv()
	if configEnv.AccessKey != nil && configEnv.SecretKey != nil {
		ctx, err := configEnv.Context(ctx) ; err != nil {
			return nil, errors.New("cannot create context from env var")
		}
		return ctx, nil
	} else {
		configFile, err := LoadDefaultConfigFile() ; err != nil {
			return nil, err
		}
		profileName := configEnv.ProfileName
		if len(*profileName) == 0 {
			*profileName = "Default"
		}
		ctx, err := configFile.Context(ctx, *profileName) ; err != nil {
			return nil, errors.New("cannot create context from file, did you load config file")
		}
		return ctx, nil
	}
}
