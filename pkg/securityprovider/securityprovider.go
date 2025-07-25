package securityprovider

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	awscredentials "github.com/aws/aws-sdk-go/aws/credentials"
	awsv4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

// NewSecurityProviderLoginPassword creates a basic auth security provider
func NewSecurityProviderLoginPassword(
	login, password string,
) (*SecurityProviderLoginPassword, error) {
	return &SecurityProviderLoginPassword{login: login, password: password}, nil
}

type SecurityProviderLoginPassword struct {
	login    string
	password string
}

func (s *SecurityProviderLoginPassword) Intercept(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth(s.login, s.password)
	return nil
}

// NewSecurityProviderAWSv4 creates an AWS v4 security provider for AK/SK use
func NewSecurityProviderAWSv4(
	access_key, secret_key, session_token, service, region string,
) (*SecurityProviderAWSv4, error) {
	return &SecurityProviderAWSv4{
		access_key:    access_key,
		secret_key:    secret_key,
		session_token: session_token,
		service:       service,
		region:        region,
	}, nil
}

type SecurityProviderAWSv4 struct {
	access_key    string
	secret_key    string
	session_token string
	service       string
	region        string
}

func (s *SecurityProviderAWSv4) Intercept(ctx context.Context, req *http.Request) error {
	var fullbody []byte

	creds := awscredentials.NewStaticCredentials(s.access_key, s.secret_key, s.session_token)
	signer := awsv4.NewSigner(creds)
	timestamp := time.Now()

	if req.GetBody != nil {
		bodyreader, err := req.GetBody()
		if err != nil {
			return err
		}

		fullbody, err = io.ReadAll(bodyreader)
		if err != nil {
			return err
		}
	}

	region := s.region
	if region == "" {
		region = "eu-west-2"
	}

	service := s.service
	if service == "" {
		service = "oapi"
	}

	readseek := bytes.NewReader(fullbody)
	_, err := signer.Sign(req, readseek, service, region, timestamp)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecurityProviderAWSv4) InterceptOks(ctx context.Context, req *http.Request) error {
	req.Header.Set("AccessKey", s.access_key)
	req.Header.Set("SecretKey", s.secret_key)
	return nil
}
