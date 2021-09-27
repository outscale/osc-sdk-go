/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FlexibleGpu Information about the flexible GPU (fGPU).
type FlexibleGpu struct {
	// If true, the fGPU is deleted when the VM is terminated.
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// The ID of the fGPU.
	FlexibleGpuId string `json:"FlexibleGpuId,omitempty"`
	// The compatible processor generation.
	Generation string `json:"Generation,omitempty"`
	// The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).
	ModelName string `json:"ModelName,omitempty"`
	// The state of the fGPU (`allocated` \\| `attaching` \\| `attached` \\| `detaching`).
	State string `json:"State,omitempty"`
	// The Subregion where the fGPU is located.
	SubregionName string `json:"SubregionName,omitempty"`
	// The ID of the VM the fGPU is attached to, if any.
	VmId string `json:"VmId,omitempty"`
}
