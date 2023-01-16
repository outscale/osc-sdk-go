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

// LinkNicToUpdate Information about the NIC attachment. If you are modifying the `DeleteOnVmDeletion` attribute, you must specify the ID of the NIC attachment.
type LinkNicToUpdate struct {
	// If true, the NIC is deleted when the VM is terminated. If false, the NIC is detached from the VM.
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// The ID of the NIC attachment.
	LinkNicId string `json:"LinkNicId,omitempty"`
}
