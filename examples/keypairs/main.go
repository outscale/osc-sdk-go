package main

import (
	"context"
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/outscale-dev/osc-sdk-go/osc"
)

func main() {
	client := osc.NewAPIClient(osc.NewConfiguration())
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.KeypairApi.ReadKeypairs(auth, nil)
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
	keypairName := "osc-sdk-go-example"
	creationOpts := osc.CreateKeypairOpts{
		CreateKeypairRequest: optional.NewInterface(
			osc.CreateKeypairRequest{
				KeypairName: keypairName,
				// You might want to generate your own public key.
				//PublicKey:   publicKey,
			}),
	}
	creation, httpRes, err := client.KeypairApi.CreateKeypair(auth, &creationOpts)
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
	_, httpRes, err = client.KeypairApi.DeleteKeypair(auth, &deletionOpts)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting keypair ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
}
