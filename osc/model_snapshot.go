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

// Snapshot Information about the snapshot.
type Snapshot struct {
	// The account alias of the owner of the snapshot.
	AccountAlias string `json:"AccountAlias,omitempty"`
	// The account ID of the owner of the snapshot.
	AccountId string `json:"AccountId,omitempty"`
	// The date and time of creation of the snapshot.
	CreationDate string `json:"CreationDate,omitempty"`
	// The description of the snapshot.
	Description               string                `json:"Description,omitempty"`
	PermissionsToCreateVolume PermissionsOnResource `json:"PermissionsToCreateVolume,omitempty"`
	// The progress of the snapshot, as a percentage.
	Progress int32 `json:"Progress,omitempty"`
	// The ID of the snapshot.
	SnapshotId string `json:"SnapshotId,omitempty"`
	// The state of the snapshot (`in-queue` \\| `completed` \\| `error`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the snapshot.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The ID of the volume used to create the snapshot.
	VolumeId string `json:"VolumeId,omitempty"`
	// The size of the volume used to create the snapshot, in gibibytes (GiB).
	VolumeSize int32 `json:"VolumeSize,omitempty"`
}
