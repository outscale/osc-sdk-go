/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersVmsState One or more filters.
type FiltersVmsState struct {
	// The code for the scheduled event (`system-reboot` \\| `system-maintenance`).
	MaintenanceEventCodes []string `json:"MaintenanceEventCodes,omitempty"`
	// The description of the scheduled event.
	MaintenanceEventDescriptions []string `json:"MaintenanceEventDescriptions,omitempty"`
	// The latest time the event can end.
	MaintenanceEventsNotAfter []string `json:"MaintenanceEventsNotAfter,omitempty"`
	// The earliest time the event can start.
	MaintenanceEventsNotBefore []string `json:"MaintenanceEventsNotBefore,omitempty"`
	// The names of the Subregions of the VMs.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// One or more IDs of VMs.
	VmIds []string `json:"VmIds,omitempty"`
	// The states of the VMs (`pending` \\| `running` \\| `stopping` \\| `stopped` \\| `shutting-down` \\| `terminated` \\| `quarantine`).
	VmStates []string `json:"VmStates,omitempty"`
}
