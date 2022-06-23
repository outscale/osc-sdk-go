/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.20
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CatalogEntry Information about the catalog entry.
type CatalogEntry struct {
	// The category of the catalog entry (for example, `network`).
	Category string `json:"Category,omitempty"`
	// When returned and equal to `PER_MONTH`, the price of the catalog entry is calculated on a monthly basis.
	Flags string `json:"Flags,omitempty"`
	// The API call associated with the catalog entry (for example, `CreateVms` or `RunInstances`).
	Operation string `json:"Operation,omitempty"`
	// The service associated with the catalog entry (`TinaOS-FCU`, `TinaOS-LBU`, `TinaOS-DirectLink`, or `TinaOS-OOS`).
	Service string `json:"Service,omitempty"`
	// The Subregion associated with the catalog entry.
	SubregionName string `json:"SubregionName,omitempty"`
	// The description of the catalog entry.
	Title string `json:"Title,omitempty"`
	// The type of resource associated with the catalog entry.
	Type string `json:"Type,omitempty"`
	// The unit price of the catalog entry, in the currency of the catalog of the Region where the API method was used.
	UnitPrice float32 `json:"UnitPrice,omitempty"`
}