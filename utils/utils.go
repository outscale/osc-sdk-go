package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// DebugRequest ...
func DebugRequest(req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string("####################"))
	fmt.Println(string("###### REQUEST #######"))
	fmt.Println(string(requestDump))
}

// DebugResponse ...
func DebugResponse(req *http.Response) {
	requestDump, err := httputil.DumpResponse(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string("####################"))
	fmt.Println(string("###### RESPONSE #######"))
	fmt.Println(string(requestDump))
}
