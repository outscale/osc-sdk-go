package profile

import (
	"slices"
)

func (p *Profile) GetAccessKeys(svc OscService) (string, string) {
	if slices.Contains(p.IAMV2Services, svc) {
		return p.AccessKeyV2, p.SecretKeyV2
	} else {
		return p.AccessKey, p.SecretKey
	}
}
