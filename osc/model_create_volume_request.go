/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CreateVolumeRequest struct for CreateVolumeRequest
type CreateVolumeRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The number of I/O operations per second (IOPS). This parameter must be specified only if you create an `io1` volume. The maximum number of IOPS allowed for `io1` volumes is `13000`.
	Iops int32 `json:"Iops,omitempty"`
	// The size of the volume, in gibibytes (GiB). The maximum allowed size for a volume is 14901 GiB. This parameter is required if the volume is not created from a snapshot (`SnapshotId` unspecified).
	Size int32 `json:"Size,omitempty"`
	// The ID of the snapshot from which you want to create the volume.
	SnapshotId string `json:"SnapshotId,omitempty"`
	// The Subregion in which you want to create the volume.
	SubregionName string `json:"SubregionName"`
	// The type of volume you want to create (`io1` \\| `gp2` \\| `standard`). If not specified, a `standard` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).
	VolumeType string `json:"VolumeType,omitempty"`
}
