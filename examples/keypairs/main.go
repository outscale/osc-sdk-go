package main

import (
	"context"
	"fmt"
	"github.com/outscale/osc-sdk-go/v2"
	"os"
)

func main() {
	client := osc.NewAPIClient(osc.NewConfiguration())
	ctx := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.KeypairApi.ReadKeypairs(ctx).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading keypairs")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}

	println("We currently have", len(*read.Keypairs), "keypairs:")
	for _, keypair := range *read.Keypairs {
		println("-", keypair.KeypairName)
	}

	println("Creating a new keypair")
	keypairName := "osc-sdk-go-example"
	createKeypair := osc.NewCreateKeypairRequest(keypairName)
	creation_req := client.KeypairApi.CreateKeypair(ctx)
	creation_req.CreateKeypairRequest(*createKeypair) // anoying
	creation, httpRes, err := creation_req.Execute()

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

	deleteKeypair := osc.NewDeleteKeypairRequest(keypairName)
	deletion_req := client.KeypairApi.DeleteKeypair(ctx)
	deletion_req.DeleteKeypairRequest(*deleteKeypair) // anoying
	_, httpRes, err = deletion_req.Execute()

	if err != nil {
		fmt.Fprint(os.Stderr, "Error while deleting keypair ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
}
