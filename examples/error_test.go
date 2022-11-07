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
	"net/http"
	"os"

	osc "github.com/outscale/osc-sdk-go/v2"
)

/*
	Error example

A quick example which show how to retrieve api error in the SDK in case you have to see what's happening in requests.
This examples just delete unexisting volumes and shows API errors.
*/
func ExampleError() {
	// few things which might be useful for debugging
	config := osc.NewConfiguration()
	config.UserAgent = "osc-sdk-go-example/0.0.0"
	// If you need to setup a specific HTTPClient for SSL client, check https://gist.github.com/michaljemala/d6f4e01c4834bf47a9c4
	config.HTTPClient = new(http.Client)

	client := osc.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
		AccessKey: os.Getenv("OSC_ACCESS_KEY"),
		SecretKey: os.Getenv("OSC_SECRET_KEY"),
	})

	_, httpRes, err := client.VolumeApi.DeleteVolume(ctx).DeleteVolumeRequest(osc.DeleteVolumeRequest{VolumeId: "volume-toto"}).Execute()
	if err != nil {
		if ok, apiErr := extractError(err); ok {
			fmt.Fprintf(os.Stderr, "API Error : Type %v, Message '%v', Code %v\n", apiErr.GetErrors()[0].GetType(), apiErr.GetErrors()[0].GetDetails(), apiErr.GetErrors()[0].GetCode())
		}
		fmt.Fprintf(os.Stderr, "Error while deleting volumes. Status Code: '%v', Status Text: '%v'\n", httpRes.StatusCode, httpRes.Status)
	} else {
		os.Exit(1)
	}

	fmt.Println("Debug journey is over")
	// Output: Debug journey is over
}

func extractError(err error) (bool, *osc.ErrorResponse) {
	genericError, ok := err.(osc.GenericOpenAPIError)
	if ok {
		errorsResponse, ok := genericError.Model().(osc.ErrorResponse)
		if ok {
			return true, &errorsResponse
		}
		return false, nil
	}
	return false, nil
}
