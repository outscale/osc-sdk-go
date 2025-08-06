package client

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/outscale/osc-sdk-go/v3/internal/transport"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/securityprovider"
)

type OscService int

const (
	OApi OscService = iota
	LBU
	OKS
)

type oscHttpClientTransform struct {
	withRetry        bool
	withRateLimit    bool
	retryOptions     []ClientWithRetryOption
	ratelimitOptions []ClientWithRateLimitOption
}

type (
	oscTransportOption       func(*http.Transport) error
	oscHttpRequestDoerOption func(*oscHttpClientTransform) error
)

type OscClientOption struct {
	transport oscTransportOption
	doer      oscHttpRequestDoerOption
	Oapi      osc.ClientOption
	OKS       oks.ClientOption
}

func newHttpRequestDoer(opts ...OscClientOption) (HttpRequestDoer, error) {
	transport := transport.DefaultPooledTransport()
	for _, o := range opts {
		if o.transport != nil {
			err := o.transport(transport)
			if err != nil {
				return nil, err
			}
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	transform := oscHttpClientTransform{}
	for _, o := range opts {
		if o.doer != nil {
			err := o.doer(&transform)
			if err != nil {
				return nil, err
			}
		}
	}

	if transform.withRetry {
		retryOpts := append(
			[]ClientWithRetryOption{WithRetryClient(client)},
			transform.retryOptions...)
		client := NewClientWithRetry(retryOpts...)

		if transform.withRateLimit {
			ratelimitOpts := append(
				[]ClientWithRateLimitOption{WithClient(client)},
				transform.ratelimitOptions...)
			return NewClientWithRateLimit(ratelimitOpts...), nil
		} else {
			return client, nil
		}
	}

	if transform.withRateLimit {
		ratelimitOpts := append(
			[]ClientWithRateLimitOption{WithClient(client)},
			transform.ratelimitOptions...)
		return NewClientWithRateLimit(ratelimitOpts...), nil
	}

	return client, nil
}

func NewOKSClient(opts ...OscClientOption) (*oks.Client, error) {
	if len(opts) == 0 {
		opts = append(opts, WithStandardConfiguration("", ""), WithRateLimit(), WithRetry())
	}

	doer, err := newHttpRequestDoer(opts...)
	if err != nil {
		return nil, err
	}

	oksOpts := []oks.ClientOption{oks.WithHTTPClient(doer)}
	for _, o := range opts {
		if o.OKS != nil {
			oksOpts = append(oksOpts, o.OKS)
		}
	}

	client, err := oks.NewClient(oksOpts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewOapiClient(opts ...OscClientOption) (*osc.Client, error) {
	if len(opts) == 0 {
		opts = append(opts, WithStandardConfiguration("", ""), WithRateLimit(), WithRetry())
	}

	doer, err := newHttpRequestDoer(opts...)
	if err != nil {
		return nil, err
	}

	oapiOpts := []osc.ClientOption{osc.WithHTTPClient(doer)}
	for _, o := range opts {
		if o.Oapi != nil {
			oapiOpts = append(oapiOpts, o.Oapi)
		}
	}

	client, err := osc.NewClient(oapiOpts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func WithHTTPClient(client *http.Client) OscClientOption {
	return OscClientOption{
		Oapi: osc.WithHTTPClient(client),
		OKS:  oks.WithHTTPClient(client),
	}
}

func NewOscClientError(e error) OscClientOption {
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
	return OscClientOption{
		transport: func(t *http.Transport) error {
			t.TLSClientConfig.Certificates = append(t.TLSClientConfig.Certificates, cert)
			return nil
		},
	}
}

func WithClientCertificatFiles(certPath, keyPath string) OscClientOption {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return NewOscClientError(err)
	}
	return WithClientCertificat(cert)
}

func WithClientCertificatBase64(certB64, keyB64 string) OscClientOption {
	certBytes, err := base64.StdEncoding.DecodeString(certB64)
	if err != nil {
		return NewOscClientError(err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return NewOscClientError(err)
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return NewOscClientError(err)
	}

	return WithClientCertificat(cert)
}

func WithTlsSkipVerify() OscClientOption {
	return OscClientOption{
		transport: func(t *http.Transport) error {
			t.TLSClientConfig.InsecureSkipVerify = true
			return nil
		},
	}
}

func WithAkSk(accesKey, secretKey, region string) OscClientOption {
	// TODO: this is not 100% accurate, because We need to pass the correct service name for OKS
	sec, err := securityprovider.NewSecurityProviderAWSv4(accesKey, secretKey, "", "oapi", region)
	if err != nil {
		return NewOscClientError(err)
	}

	return OscClientOption{
		Oapi: osc.WithRequestEditorFn(sec.Intercept),
		OKS:  oks.WithRequestEditorFn(sec.InterceptOks),
	}
}

func WithLoginPassword(login, password string) OscClientOption {
	sec, err := securityprovider.NewSecurityProviderLoginPassword(login, password)
	if err != nil {
		return NewOscClientError(err)
	}

	return OscClientOption{
		Oapi: osc.WithRequestEditorFn(sec.Intercept),
		OKS:  oks.WithRequestEditorFn(sec.Intercept),
	}
}

func WithRetry(opts ...ClientWithRetryOption) OscClientOption {
	return OscClientOption{
		doer: func(ohct *oscHttpClientTransform) error {
			ohct.withRetry = true
			ohct.retryOptions = opts
			return nil
		},
	}
}

func WithRateLimit(opts ...ClientWithRateLimitOption) OscClientOption {
	return OscClientOption{
		doer: func(ohct *oscHttpClientTransform) error {
			ohct.withRateLimit = true
			ohct.ratelimitOptions = opts
			return nil
		},
	}
}

func WithProfile(profile *Profile) OscClientOption {
	opts := make([]OscClientOption, 0, 2)

	// 0. Skip Tls Verify.
	if profile.TlsSkipVerify {
		opts = append(opts, WithTlsSkipVerify())
	}

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
		return NewOscClientError(errors.New("no authentication provided"))
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
				println("server oks: %s", cr.Server)
				return nil
			},
		},
	)

	return OscClientOption{
		transport: func(t *http.Transport) error {
			for _, o := range opts {
				if o.transport == nil {
					continue
				}

				if err := o.transport(t); err != nil {
					return err
				}
			}

			return nil
		},
		doer: func(ohct *oscHttpClientTransform) error {
			for _, o := range opts {
				if o.doer == nil {
					continue
				}

				if err := o.doer(ohct); err != nil {
					return err
				}
			}

			return nil
		},
		Oapi: func(cr *osc.ClientRaw) error {
			for _, o := range opts {
				if o.Oapi == nil {
					continue
				}

				if err := o.Oapi(cr); err != nil {
					return err
				}
			}

			return nil
		},
		OKS: func(cr *oks.ClientRaw) error {
			for _, o := range opts {
				if o.OKS == nil {
					continue
				}

				if err := o.OKS(cr); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func WithStandardConfiguration(profile, path string) OscClientOption {
	p, err := NewProfileFromStrandardConfiguration(profile, path)
	if err != nil {
		return NewOscClientError(err)
	}

	return WithProfile(p)
}
