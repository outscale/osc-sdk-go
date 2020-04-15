package main

import (
	"context"
	"fmt"
	"github.com/outscale-dev/osc-sdk-go/osc"
	"os"
)

func main() {
     	config := osc.NewConfiguration()
	config.Debug = true
	client := osc.NewAPIClient(config)
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.VolumeApi.ReadVolumes(auth, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading volumes")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	for _, volume := range read.Volumes {
		println("-", volume.VolumeId)
	}
}
