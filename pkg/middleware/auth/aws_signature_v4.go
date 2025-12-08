package auth

import (
	"crypto/sha256"
	"io"
	"net/http"

	"github.com/aws/smithy-go/aws-http-auth/credentials"
	"github.com/aws/smithy-go/aws-http-auth/sigv4"
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
		signer:                sigv4.New(),
		SecurityProviderAWSv4: s,
	}
}

type awsSignatureV4RoundTripper struct {
	inner  http.RoundTripper
	signer *sigv4.Signer
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
	sigReq := sigv4.SignRequestInput{
		Request: req,
		Credentials: credentials.Credentials{
			AccessKeyID:     a.accessKey,
			SecretAccessKey: a.secretKey,
			SessionToken:    a.sessionToken,
		},
		Service: a.service,
		Region:  a.region,
	}

	if req.GetBody != nil {
		h := sha256.New()

		bodyreader, err := req.GetBody()
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(h, bodyreader); err != nil {
			return nil, err
		}

		sigReq.PayloadHash = h.Sum(nil)
	}

	err := a.signer.SignRequest(&sigReq)
	if err != nil {
		return nil, err
	}

	return a.inner.RoundTrip(req)
}
