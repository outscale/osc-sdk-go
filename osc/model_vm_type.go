/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// VmType Information about the VM type.
type VmType struct {
	// Indicates whether the VM is optimized for BSU I/O.
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
