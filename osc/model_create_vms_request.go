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

// CreateVmsRequest struct for CreateVmsRequest
type CreateVmsRequest struct {
	// One or more block device mappings.
	BlockDeviceMappings []BlockDeviceMappingVmCreation `json:"BlockDeviceMappings,omitempty"`
	// By default or if true, the VM is started on creation. If false, the VM is stopped on creation.
	BootOnCreation bool `json:"BootOnCreation,omitempty"`
	// This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// A unique identifier which enables you to manage the idempotency.
	ClientToken string `json:"ClientToken,omitempty"`
	// If true, you cannot delete the VM unless you change this parameter back to false.
	DeletionProtection bool `json:"DeletionProtection,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](#readimages) method.
	ImageId string `json:"ImageId"`
	// The name of the keypair.
	KeypairName string `json:"KeypairName,omitempty"`
	// The maximum number of VMs you want to create. If all the VMs cannot be created, the largest possible number of VMs above MinVmsCount is created.
	MaxVmsCount int32 `json:"MaxVmsCount,omitempty"`
	// The minimum number of VMs you want to create. If this number of VMs cannot be created, no VMs are created.
	MinVmsCount int32 `json:"MinVmsCount,omitempty"`
	// (dedicated tenancy only) If true, nested virtualization is enabled. If false, it is disabled.
	NestedVirtualization bool `json:"NestedVirtualization,omitempty"`
	// One or more NICs. If you specify this parameter, you must not specify the `SubnetId` and `SubregionName` parameters. You also must define one NIC as the primary network interface of the VM with `0` as its device number.
	Nics []NicForVmCreation `json:"Nics,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`). By default, `high`. This parameter is ignored if you specify a performance flag directly in the `VmType` parameter.
	Performance string    `json:"Performance,omitempty"`
	Placement   Placement `json:"Placement,omitempty"`
	// One or more private IPs of the VM.
	PrivateIps []string `json:"PrivateIps,omitempty"`
	// One or more IDs of security group for the VMs.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// One or more names of security groups for the VMs.
	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	// The ID of the Subnet in which you want to create the VM. If you specify this parameter, you must not specify the `Nics` parameter.
	SubnetId string `json:"SubnetId,omitempty"`
	// Data or script used to add a specific configuration to the VM. It must be Base64-encoded and is limited to 500 kibibytes (KiB).
	UserData string `json:"UserData,omitempty"`
	// The VM behavior when you stop it. By default or if set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is terminated.
	VmInitiatedShutdownBehavior string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM. You can specify a TINA type (in the `tinavW.cXrYpZ` or `tinavW.cXrY` format), or an AWS type (for example, `t2.small`, which is the default value).<br /> If you specify an AWS type, it is converted in the background to its corresponding TINA type, but the AWS type is still returned. If the specified or converted TINA type includes a performance flag, this performance flag is applied regardless of the value you may have provided in the `Performance` parameter. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmType string `json:"VmType,omitempty"`
}
