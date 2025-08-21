package auth

import (
	"net/http"
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

func (s *SecurityProviderLoginPassword) Decorate(next http.RoundTripper) http.RoundTripper {
	return &loginPasswordRoundTripper{
		inner: next,
		creds: s,
	}
}

type loginPasswordRoundTripper struct {
	inner http.RoundTripper
	creds *SecurityProviderLoginPassword
}

func (l *loginPasswordRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(l.creds.login, l.creds.password)
	return l.inner.RoundTrip(req)
}
