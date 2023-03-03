/*
BSD 3-Clause License

Copyright (c) 2021, Outscale SAS
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

	"github.com/antihax/optional"
	"github.com/outscale/osc-sdk-go/osc"
)

/*
	Keypairs Example

- List all keypairs
- Create new keypair
- Show newly created keypair details
- Delete keypair
*/
func ExampleKeypair() {
	configEnv := osc.NewConfigEnv()
	config, err := configEnv.Configuration()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot load configuration from environement variables")
		os.Exit(1)
	}
	ctx, err := configEnv.Context(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot set context from environement variables")
		os.Exit(1)
	}
	client := osc.NewAPIClient(config)

	read, httpRes, err := client.KeypairApi.ReadKeypairs(ctx, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading keypairs")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}

	println("We currently have", len(read.Keypairs), "keypairs:")
	for _, keypair := range read.Keypairs {
		println("-", keypair.KeypairName)
	}

	println("Creating a new keypair")
	keypairName := "osc-sdk-go-example-" + RandomString(7)
	creationOpts := osc.CreateKeypairOpts{
		CreateKeypairRequest: optional.NewInterface(
			osc.CreateKeypairRequest{
				KeypairName: keypairName,
				// You might want to generate your own public key.
				//PublicKey:   publicKey,
			}),
	}
	creation, httpRes, err := client.KeypairApi.CreateKeypair(ctx, &creationOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while creating keypair ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("Created keypair with fingerprint", creation.Keypair.KeypairFingerprint)
	println("Private key content:")
	println(creation.Keypair.PrivateKey)

	println("Deleting", keypairName, "keypair")
	deletionOpts := osc.DeleteKeypairOpts{
		DeleteKeypairRequest: optional.NewInterface(
			osc.DeleteKeypairRequest{
				KeypairName: keypairName,
			}),
	}
	_, httpRes, err = client.KeypairApi.DeleteKeypair(ctx, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting keypair ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	fmt.Println("Keypair journey is over")
	// Output: Keypair journey is over
}
