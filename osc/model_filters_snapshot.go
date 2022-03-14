/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// FiltersSnapshot One or more filters.
type FiltersSnapshot struct {
	// The account aliases of the owners of the snapshots.
	AccountAliases []string `json:"AccountAliases,omitempty"`
	// The account IDs of the owners of the snapshots.
	AccountIds []string `json:"AccountIds,omitempty"`
	// The descriptions of the snapshots.
	Descriptions []string `json:"Descriptions,omitempty"`
	// The account IDs of one or more users who have permissions to create volumes.
	PermissionsToCreateVolumeAccountIds []string `json:"PermissionsToCreateVolumeAccountIds,omitempty"`
	// If true, lists all public volumes. If false, lists all private volumes.
	PermissionsToCreateVolumeGlobalPermission bool `json:"PermissionsToCreateVolumeGlobalPermission,omitempty"`
	// The progresses of the snapshots, as a percentage.
	Progresses []int32 `json:"Progresses,omitempty"`
	// The IDs of the snapshots.
	SnapshotIds []string `json:"SnapshotIds,omitempty"`
	// The states of the snapshots (`in-queue` \\| `completed` \\| `error`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the snapshots.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the snapshots.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the snapshots, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
	// The IDs of the volumes used to create the snapshots.
	VolumeIds []string `json:"VolumeIds,omitempty"`
	// The sizes of the volumes used to create the snapshots, in gibibytes (GiB).
	VolumeSizes []int32 `json:"VolumeSizes,omitempty"`
}
