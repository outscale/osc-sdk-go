/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.26
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersVmType One or more filters.
type FiltersVmType struct {
	// This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// The amounts of memory, in gibibytes (GiB).
	MemorySizes []float32 `json:"MemorySizes,omitempty"`
	// The numbers of vCores.
	VcoreCounts []int32 `json:"VcoreCounts,omitempty"`
	// The names of the VM types. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmTypeNames []string `json:"VmTypeNames,omitempty"`
	// The maximum number of ephemeral storage disks.
	VolumeCounts []int32 `json:"VolumeCounts,omitempty"`
	// The size of one ephemeral storage disk, in gibibytes (GiB).
	VolumeSizes []int32 `json:"VolumeSizes,omitempty"`
}
