package useragent

import "net/http"

const userAgentHeader = "User-Agent"

type UseragentMiddleware struct {
	Useragent string
}

func (r *UseragentMiddleware) Decorate(next http.RoundTripper) http.RoundTripper {
	return &innerUseragent{
		inner:               next,
		UseragentMiddleware: r,
	}
}

type innerUseragent struct {
	inner http.RoundTripper
	*UseragentMiddleware
}

func (u *innerUseragent) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get(userAgentHeader) == "" {
		req.Header.Add(userAgentHeader, u.Useragent)
	} else {
		req.Header.Set(userAgentHeader, u.Useragent)
	}
	return u.inner.RoundTrip(req)
}
