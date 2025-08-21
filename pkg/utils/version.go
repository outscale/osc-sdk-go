package utils

import (
	"runtime/debug"
)

const (
	defaultVersion = "v3.0.0+dev"
	path           = "github.com/outscale/osc-sdk-go"
)

var cachedVersion *string

func SdkVersion() string {
	if cachedVersion == nil {
		versionFromDebugInfo := ""
		b, ok := debug.ReadBuildInfo()
		if ok {
			for _, dep := range b.Deps {
				if dep.Path == path {
					versionFromDebugInfo = dep.Version
				}
			}
		}

		cachedVersion = &versionFromDebugInfo
	}

	if *cachedVersion != "" {
		return *cachedVersion
	}

	return defaultVersion
}
