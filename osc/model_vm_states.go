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

// VmStates Information about the states of the VMs.
type VmStates struct {
	// One or more scheduled events associated with the VM.
	MaintenanceEvents []MaintenanceEvent `json:"MaintenanceEvents,omitempty"`
	// The name of the Subregion of the VM.
	SubregionName string `json:"SubregionName,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId,omitempty"`
	// The state of the VM (`pending` \\| `running` \\| `stopping` \\| `stopped` \\| `shutting-down` \\| `terminated` \\| `quarantine`).
	VmState string `json:"VmState,omitempty"`
}
