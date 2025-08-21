package profile

import (
	"errors"
	"fmt"
)

type Endpoint struct {
	API string `json:"api"`
	OKS string `json:"oks"`
}

func getDefaultEndpointTemplate(service OscService) (string, error) {
	switch service {
	case OscServiceApi:
		return "%s://api.%s.outscale.com/api/v1", nil
	case OscServiceOKS:
		return "%s://api.%s.oks.outscale.com/api/v2", nil
	default:
		return "", errors.New("unsupported service")
	}
}

func (p *Profile) getDefaultEndpoint(service OscService) (string, error) {
	temp, err := getDefaultEndpointTemplate(service)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(temp, p.Protocol, p.Region), nil
}

func (p *Profile) GetEndpoint(service OscService) (string, error) {
	var endpoint string

	switch service {
	case OscServiceApi:
		endpoint = p.Endpoints.API
	case OscServiceOKS:
		endpoint = p.Endpoints.OKS
	}

	if endpoint == "" {
		return p.getDefaultEndpoint(service)
	}

	return endpoint, nil
}
