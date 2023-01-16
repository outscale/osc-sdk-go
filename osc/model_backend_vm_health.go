/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.24
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// BackendVmHealth Information about the health of a back-end VM.
type BackendVmHealth struct {
	// The description of the state of the back-end VM.
	Description string `json:"Description,omitempty"`
	// The state of the back-end VM (`InService` \\| `OutOfService` \\| `Unknown`).
	State string `json:"State,omitempty"`
	// Information about the cause of `OutOfService` VMs.<br /> Specifically, whether the cause is Elastic Load Balancing or the VM (`ELB` \\| `Instance` \\| `N/A`).
	StateReason string `json:"StateReason,omitempty"`
	// The ID of the back-end VM.
	VmId string `json:"VmId,omitempty"`
}
