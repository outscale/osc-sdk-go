package useragent

import "net/http"

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
	req.Header.Add("User-Agent", u.Useragent)
	return u.inner.RoundTrip(req)
}

