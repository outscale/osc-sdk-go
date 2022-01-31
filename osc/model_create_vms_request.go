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

// CreateVmsRequest struct for CreateVmsRequest
type CreateVmsRequest struct {
	// One or more block device mappings.
	BlockDeviceMappings []BlockDeviceMappingVmCreation `json:"BlockDeviceMappings,omitempty"`
	// By default or if true, the VM is started on creation. If false, the VM is stopped on creation.
	BootOnCreation bool `json:"BootOnCreation,omitempty"`
	// If true, the VM is created with optimized BSU I/O.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// A unique identifier which enables you to manage the idempotency.
	ClientToken string `json:"ClientToken,omitempty"`
	// If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.
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
	// One or more NICs. If you specify this parameter, you must define one NIC as the primary network interface of the VM with `0` as its device number.
	Nics []NicForVmCreation `json:"Nics,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`).
	Performance string    `json:"Performance,omitempty"`
	Placement   Placement `json:"Placement,omitempty"`
	// One or more private IPs of the VM.
	PrivateIps []string `json:"PrivateIps,omitempty"`
	// One or more IDs of security group for the VMs.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// One or more names of security groups for the VMs.
	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	// The ID of the Subnet in which you want to create the VM.
	SubnetId string `json:"SubnetId,omitempty"`
	// Data or script used to add a specific configuration to the VM. It must be Base64-encoded and is limited to 500 kibibytes (KiB).
	UserData string `json:"UserData,omitempty"`
	// The VM behavior when you stop it. By default or if set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is terminated.
	VmInitiatedShutdownBehavior string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM (`t2.small` by default).<br /> For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmType string `json:"VmType,omitempty"`
}
