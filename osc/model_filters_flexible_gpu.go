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

// FiltersFlexibleGpu One or more filters.
type FiltersFlexibleGpu struct {
	// Indicates whether the fGPU is deleted when terminating the VM.
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// One or more IDs of fGPUs.
	FlexibleGpuIds []string `json:"FlexibleGpuIds,omitempty"`
	// The processor generations that the fGPUs are compatible with.
	Generations []string `json:"Generations,omitempty"`
	// One or more models of fGPUs. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).
	ModelNames []string `json:"ModelNames,omitempty"`
	// The states of the fGPUs (`allocated` \\| `attaching` \\| `attached` \\| `detaching`).
	States []string `json:"States,omitempty"`
	// The Subregions where the fGPUs are located.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// One or more IDs of VMs.
	VmIds []string `json:"VmIds,omitempty"`
}
