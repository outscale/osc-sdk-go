/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.21
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// VmType Information about the VM type.
type VmType struct {
	// This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// The maximum number of private IPs per network interface card (NIC).
	MaxPrivateIps int32 `json:"MaxPrivateIps,omitempty"`
	// The amount of memory, in gibibytes.
	MemorySize float32 `json:"MemorySize,omitempty"`
	// The number of vCores.
	VcoreCount int32 `json:"VcoreCount,omitempty"`
	// The name of the VM type.
	VmTypeName string `json:"VmTypeName,omitempty"`
	// The maximum number of ephemeral storage disks.
	VolumeCount int32 `json:"VolumeCount,omitempty"`
	// The size of one ephemeral storage disk, in gibibytes (GiB).
	VolumeSize int32 `json:"VolumeSize,omitempty"`
}
