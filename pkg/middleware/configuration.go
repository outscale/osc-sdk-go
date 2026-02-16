package middleware

import (
	"crypto/tls"
	"encoding/base64"

	"github.com/outscale/osc-sdk-go/v3/pkg/middleware/auth"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
)

func newOscClientError(err error) MiddlewareChainOption {
	return func(mc *MiddlewareChain) error {
		return err
	}
}

func withClientCertificat(cert tls.Certificate) MiddlewareChainOption {
	return func(mc *MiddlewareChain) error {
		mc.Base.TLSClientConfig.Certificates = append(mc.Base.TLSClientConfig.Certificates, cert)
		return nil
	}
}

func withClientCertificatFiles(certPath, keyPath string) MiddlewareChainOption {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return newOscClientError(err)
	}
	return withClientCertificat(cert)
}

func withClientCertificatBase64(certB64, keyB64 string) MiddlewareChainOption {
	certBytes, err := base64.StdEncoding.DecodeString(certB64)
	if err != nil {
		return newOscClientError(err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return newOscClientError(err)
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return newOscClientError(err)
	}

	return withClientCertificat(cert)
}

func withTlsSkipVerify() MiddlewareChainOption {
	return func(mc *MiddlewareChain) error {
		mc.Base.TLSClientConfig.InsecureSkipVerify = true
		return nil
	}
}

func withAkSk(accesKey, secretKey, region, service string) MiddlewareChainOption {
	// TODO: this is not 100% accurate, because We need to pass the correct service name for OKS
	sec, err := auth.NewSecurityProviderAWSv4(accesKey, secretKey, "", service, region)
	if err != nil {
		return newOscClientError(err)
	}

	return WithMiddleware(MiddlewareSlotAuth, sec)
}

func withLoginPassword(login, password string) MiddlewareChainOption {
	sec, err := auth.NewSecurityProviderLoginPassword(login, password)
	if err != nil {
		return newOscClientError(err)
	}

	return WithMiddleware(MiddlewareSlotAuth, sec)
}

func FromProfile(userProfile *profile.Profile, service profile.OscService) MiddlewareChainOption {
	opts := make([]MiddlewareChainOption, 0, 2)

	// 0. Skip Tls Verify.
	if userProfile.TlsSkipVerify {
		opts = append(opts, withTlsSkipVerify())
	}

	// 1. Check authentication
	ak, sk := userProfile.GetAccessKeys(service)

	if userProfile.X509ClientCert != "" && userProfile.X509ClientKey != "" {
		opts = append(
			opts,
			withClientCertificatFiles(userProfile.X509ClientCert, userProfile.X509ClientKey),
		)
	} else if userProfile.X509ClientCertB64 != "" && userProfile.X509ClientKeyB64 != "" {
		opts = append(
			opts,
			withClientCertificatBase64(userProfile.X509ClientCertB64, userProfile.X509ClientKeyB64),
		)
	}

	if ak != "" && sk != "" && userProfile.Region != "" {
		opts = append(
			opts,
			withAkSk(ak, sk, userProfile.Region, service),
		)
	} else if userProfile.Login != "" && userProfile.Password != "" {
		opts = append(
			opts,
			withLoginPassword(userProfile.Login, userProfile.Password),
		)
	}

	return func(mc *MiddlewareChain) error {
		return mc.applyOptions(opts...)
	}
}
