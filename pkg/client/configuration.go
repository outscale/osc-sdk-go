package client

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"net/http"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/outscale/osc-sdk-go/v3/internal/oks"
	"github.com/outscale/osc-sdk-go/v3/internal/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/securityprovider"
)

type OscService int

const (
	OApi OscService = iota
	LBU
	OKS
)

func NewOKSClient(opts ...OscClientOption) (*oks.Client, error) {
	if len(opts) == 0 {
		opts = append(opts, WithStandardConfiguration("", ""), WithRateLimit(), WithRetry())
	}

	oksOpts := make([]oks.ClientOption, len(opts))
	for i, o := range opts {
		oksOpts[i] = o.OKS
	}

	client, err := oks.NewClient("", oksOpts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewOapiClient(opts ...OscClientOption) (*osc.Client, error) {
	if len(opts) == 0 {
		opts = append(opts, WithStandardConfiguration("", ""), WithRateLimit(), WithRetry())
	}

	oapiOpts := make([]osc.ClientOption, len(opts))
	for i, o := range opts {
		oapiOpts[i] = o.Oapi
	}

	client, err := osc.NewClient("", oapiOpts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type OscClientOption struct {
	Oapi osc.ClientOption
	OKS  oks.ClientOption
}

func WithHTTPClient(client *http.Client) OscClientOption {
	return OscClientOption{
		Oapi: osc.WithHTTPClient(client),
		OKS:  oks.WithHTTPClient(client),
	}
}

func NewOcsClientError(e error) OscClientOption {
	return OscClientOption{
		Oapi: func(cr *osc.ClientRaw) error {
			return e
		},
		OKS: func(cr *oks.ClientRaw) error {
			return e
		},
	}
}

func WithClientCertificat(cert tls.Certificate) OscClientOption {
	transport := cleanhttp.DefaultPooledTransport()
	transport.TLSClientConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	client := &http.Client{
		Transport: transport,
	}

	return WithHTTPClient(client)
}

func WithClientCertificatFiles(certPath, keyPath string) OscClientOption {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return NewOcsClientError(err)
	}
	return WithClientCertificat(cert)
}

func WithClientCertificatBase64(certB64, keyB64 string) OscClientOption {
	certBytes, err := base64.StdEncoding.DecodeString(certB64)
	if err != nil {
		return NewOcsClientError(err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return NewOcsClientError(err)
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return NewOcsClientError(err)
	}

	return WithClientCertificat(cert)
}

func WithAkSk(accesKey, secretKey, region string) OscClientOption {
	// TODO: this is not 100% accurate, because We need to pass the correct service name for OKS
	sec, err := securityprovider.NewSecurityProviderAWSv4(accesKey, secretKey, "", "oapi", region)
	if err != nil {
		return NewOcsClientError(err)
	}

	return OscClientOption{
		Oapi: osc.WithRequestEditorFn(sec.Intercept),
		OKS:  oks.WithRequestEditorFn(sec.InterceptOks),
	}
}

func WithLoginPassword(login, password string) OscClientOption {
	sec, err := securityprovider.NewSecurityProviderLoginPassword(login, password)
	if err != nil {
		return NewOcsClientError(err)
	}

	return OscClientOption{
		Oapi: osc.WithRequestEditorFn(sec.Intercept),
		OKS:  oks.WithRequestEditorFn(sec.Intercept),
	}
}

func WithRetry(opts ...ClientWithRetryOption) OscClientOption {
	return OscClientOption{
		Oapi: func(cr *osc.ClientRaw) error {
			if cr.Client == nil {
				cr.Client = NewClientWithRetry(opts...)
				return nil
			}

			hc, ok := cr.Client.(*http.Client)
			if ok {
				opts = append(opts, WithRetryClient(hc))
				cr.Client = NewClientWithRetry(opts...)
				return nil
			}

			crl, ok := cr.Client.(ClientWithRateLimit)
			if ok {
				WithClient(NewClientWithRetry(opts...))(&crl)
				return nil
			}

			return errors.New("unsupported client type")
		},
		OKS: func(cr *oks.ClientRaw) error {
			if cr.Client == nil {
				cr.Client = NewClientWithRetry(opts...)
				return nil
			}

			hc, ok := cr.Client.(*http.Client)
			if ok {
				opts = append(opts, WithRetryClient(hc))
				cr.Client = NewClientWithRetry(opts...)
				return nil
			}

			crl, ok := cr.Client.(ClientWithRateLimit)
			if ok {
				WithClient(NewClientWithRetry(opts...))(&crl)
				return nil
			}

			return errors.New("unsupported client type")
		},
	}
}

func WithRateLimit(opts ...ClientWithRateLimitOption) OscClientOption {
	return OscClientOption{
		Oapi: func(cr *osc.ClientRaw) error {
			if cr.Client == nil {
				cr.Client = NewClientWithRateLimit(opts...)
				return nil
			}

			hc, ok := cr.Client.(*http.Client)
			if ok {
				opts = append(opts, WithClient(hc))
				cr.Client = NewClientWithRateLimit(opts...)
				return nil
			}

			rc, ok := cr.Client.(ClientWithRetry)
			if ok {
				opts = append(opts, WithClient(rc))
				cr.Client = NewClientWithRateLimit(opts...)
				return nil
			}

			return errors.New("unsupported client type")
		},
		OKS: func(cr *oks.ClientRaw) error {
			if cr.Client == nil {
				cr.Client = NewClientWithRateLimit(opts...)
				return nil
			}

			hc, ok := cr.Client.(*http.Client)
			if ok {
				opts = append(opts, WithClient(hc))
				cr.Client = NewClientWithRateLimit(opts...)
				return nil
			}

			rc, ok := cr.Client.(ClientWithRetry)
			if ok {
				opts = append(opts, WithClient(rc))
				cr.Client = NewClientWithRateLimit(opts...)
				return nil
			}

			return errors.New("unsupported client type")
		},
	}
}

func WithProfile(profile *Profile) OscClientOption {
	opts := make([]OscClientOption, 0, 2)

	// 1. Check authentication
	if profile.X509ClientCert != "" && profile.X509ClientKey != "" {
		opts = append(
			opts,
			WithClientCertificatFiles(profile.X509ClientCert, profile.X509ClientKey),
		)
	} else if profile.X509ClientCertB64 != "" && profile.X509ClientKeyB64 != "" {
		opts = append(
			opts,
			WithClientCertificatBase64(profile.X509ClientCertB64, profile.X509ClientKeyB64),
		)
	} else if profile.AccessKey != "" && profile.SecretKey != "" && profile.Region != "" {
		opts = append(
			opts,
			WithAkSk(profile.AccessKey, profile.SecretKey, profile.Region),
		)
	} else if profile.Login != "" && profile.Password != "" {
		opts = append(
			opts,
			WithLoginPassword(profile.Login, profile.Password),
		)
	} else {
		return NewOcsClientError(errors.New("no authentication provided"))
	}

	// 2. Check Endpoint
	opts = append(
		opts,
		OscClientOption{
			Oapi: func(cr *osc.ClientRaw) error {
				endpoint, err := profile.GetEndpoint(OApi)
				if err != nil {
					return err
				}
				cr.Server = endpoint
				return nil
			},
			OKS: func(cr *oks.ClientRaw) error {
				endpoint, err := profile.GetEndpoint(OKS)
				if err != nil {
					return err
				}
				cr.Server = endpoint
				return nil
			},
		},
	)

	return OscClientOption{
		Oapi: func(cr *osc.ClientRaw) error {
			for _, o := range opts {
				if err := o.Oapi(cr); err != nil {
					return err
				}
			}

			return nil
		},
		OKS: func(cr *oks.ClientRaw) error {
			for _, o := range opts {
				if err := o.OKS(cr); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func WithStandardConfiguration(profile, path string) OscClientOption {
	p, err := NewProfileFromStrandardConfuguration(profile, path)
	if err != nil {
		return NewOcsClientError(err)
	}

	return WithProfile(p)
}
