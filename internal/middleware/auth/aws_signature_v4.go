package auth

import (
	"bytes"
	"io"
	"net/http"
	"time"

	awscredentials "github.com/aws/aws-sdk-go/aws/credentials"
	awsv4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

// NewSecurityProviderAWSv4 creates an AWS v4 security provider for AK/SK use
func NewSecurityProviderAWSv4(
	accessKey, secretKey, sessionToken, service, region string,
) (*SecurityProviderAWSv4, error) {
	return &SecurityProviderAWSv4{
		accessKey:    accessKey,
		secretKey:    secretKey,
		sessionToken: sessionToken,
		service:      service,
		region:       region,
	}, nil
}

type SecurityProviderAWSv4 struct {
	accessKey    string
	secretKey    string
	sessionToken string
	service      string
	region       string
}

func (s *SecurityProviderAWSv4) Decorate(next http.RoundTripper) http.RoundTripper {
	if s.region == "" {
		s.region = "eu-west-2"
	}

	if s.service == "" {
		s.service = "oapi"
	}

	return &awsSignatureV4RoundTripper{
		inner:                 next,
		SecurityProviderAWSv4: s,
	}
}

type awsSignatureV4RoundTripper struct {
	inner http.RoundTripper
	*SecurityProviderAWSv4
}

func (a *awsSignatureV4RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	switch a.service {
	case "oks":
		return a.roundTripOks(req)
	default:
		return a.roundTrip(req)
	}
}

func (a *awsSignatureV4RoundTripper) roundTripOks(req *http.Request) (*http.Response, error) {
	req.Header.Set("AccessKey", a.accessKey)
	req.Header.Set("SecretKey", a.secretKey)
	return a.inner.RoundTrip(req)
}

func (a *awsSignatureV4RoundTripper) roundTrip(req *http.Request) (*http.Response, error) {
	var fullbody []byte

	creds := awscredentials.NewStaticCredentials(a.accessKey, a.secretKey, a.sessionToken)
	signer := awsv4.NewSigner(creds)
	timestamp := time.Now()

	if req.GetBody != nil {
		bodyreader, err := req.GetBody()
		if err != nil {
			return nil, err
		}

		fullbody, err = io.ReadAll(bodyreader)
		if err != nil {
			return nil, err
		}
	}

	readseek := bytes.NewReader(fullbody)
	_, err := signer.Sign(req, readseek, a.service, a.region, timestamp)
	if err != nil {
		return nil, err
	}

	return a.inner.RoundTrip(req)
}
