package profile

import (
	"errors"
	"fmt"
)

type Endpoint struct {
	API        string `json:"api"`
	OKS        string `json:"oks"`
	LBU        string `json:"lbu"`
	OOS        string `json:"oos"`
	FCU        string `json:"fcu"`
	EIM        string `json:"eim"`
	DirectLink string `json:"direct_link"`
}

func getDefaultEndpointTemplate(service OscService) (string, error) {
	switch service {
	case OscServiceApi:
		return "%s://api.%s.outscale.com/api/v1", nil
	case OscServiceOKS:
		return "%s://api.%s.oks.outscale.com/api/v2", nil
	case OscServiceLBU:
		return "%s://lbu.%s.outscale.com", nil
	case OscServiceOOS:
		return "%s://oos.%s.outscale.com", nil
	case OscServiceFCU:
		return "%s://fcu.%s.outscale.com", nil
	case OscServiceEIM:
		return "%s://eim.%s.outscale.com", nil
	case OscServiceDirectLink:
		return "%s://directlink.%s.outscale.com", nil
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
	case OscServiceLBU:
		endpoint = p.Endpoints.LBU
	case OscServiceOOS:
		endpoint = p.Endpoints.OOS
	case OscServiceFCU:
		endpoint = p.Endpoints.FCU
	case OscServiceEIM:
		endpoint = p.Endpoints.EIM
	case OscServiceDirectLink:
		endpoint = p.Endpoints.DirectLink
	}

	if endpoint == "" {
		return p.getDefaultEndpoint(service)
	}

	return endpoint, nil
}
