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

// UpdateVmRequest struct for UpdateVmRequest
type UpdateVmRequest struct {
	// One or more block device mappings of the VM.
	BlockDeviceMappings []BlockDeviceMappingVmUpdate `json:"BlockDeviceMappings,omitempty"`
	// This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// If true, you cannot delete the VM unless you change this parameter back to false.
	DeletionProtection bool `json:"DeletionProtection,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked bool `json:"IsSourceDestChecked,omitempty"`
	// The name of a keypair you want to associate with the VM.<br /> When you replace the keypair of a VM with another one, the metadata of the VM is modified to reflect the new public key, but the replacement is still not effective in the operating system of the VM. To complete the replacement and effectively apply the new keypair, you need to perform other actions inside the VM. For more information, see [Modifying the Keypair of an Instance](https://docs.outscale.com/en/userguide/Modifying-the-Keypair-of-an-Instance.html).
	KeypairName string `json:"KeypairName,omitempty"`
	// (dedicated tenancy only) If true, nested virtualization is enabled. If false, it is disabled.
	NestedVirtualization bool `json:"NestedVirtualization,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`).
	Performance string `json:"Performance,omitempty"`
	// One or more IDs of security groups for the VM.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// The Base64-encoded MIME user data, limited to 500 kibibytes (KiB).
	UserData string `json:"UserData,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId"`
	// The VM behavior when you stop it. If set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is terminated.
	VmInitiatedShutdownBehavior string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmType string `json:"VmType,omitempty"`
}
