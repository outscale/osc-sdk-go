package main

import (
	"context"
	"fmt"
	"os"

	"github.com/outscale-dev/osc-sdk-go/osc"
)

func main() {
	config := osc.NewConfiguration()
	config.BasePath, _ = config.ServerUrl(0, map[string]string{"region": "eu-west-2"})
	client := osc.NewAPIClient(config)
	auth := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	read, httpRes, err := client.SubregionApi.ReadSubregions(auth, nil)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error while reading subregions ")
		if httpRes != nil {
			fmt.Fprintln(os.Stderr, httpRes.Status)
		}
		os.Exit(1)
	}
	println("We have ", len(read.Subregions), "subregions (Availability Zones):")
	for _, sub := range read.Subregions {
		println("- (", sub.RegionName, ")", sub.SubregionName, ":", sub.State)
	}
}
