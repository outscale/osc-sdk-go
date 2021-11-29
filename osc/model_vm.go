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

// Vm Information about the VM.
type Vm struct {
	// The architecture of the VM (`i386` \\| `x86_64`).
	Architecture string `json:"Architecture,omitempty"`
	// The block device mapping of the VM.
	BlockDeviceMappings []BlockDeviceMappingCreated `json:"BlockDeviceMappings,omitempty"`
	// If true, the VM is optimized for BSU I/O.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// The idempotency token provided when launching the VM.
	ClientToken string `json:"ClientToken,omitempty"`
	// The date and time at which the VM was created.
	CreationDate string `json:"CreationDate,omitempty"`
	// If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.
	DeletionProtection bool `json:"DeletionProtection,omitempty"`
	// The hypervisor type of the VMs (`ovm` \\| `xen`).
	Hypervisor string `json:"Hypervisor,omitempty"`
	// The ID of the OMI used to create the VM.
	ImageId string `json:"ImageId,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked bool `json:"IsSourceDestChecked,omitempty"`
	// The name of the keypair used when launching the VM.
	KeypairName string `json:"KeypairName,omitempty"`
	// The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).
	LaunchNumber int32 `json:"LaunchNumber,omitempty"`
	// The ID of the Net in which the VM is running.
	NetId string `json:"NetId,omitempty"`
	// (Net only) The network interface cards (NICs) the VMs are attached to.
	Nics []NicLight `json:"Nics,omitempty"`
	// Indicates the operating system (OS) of the VM.
	OsFamily string `json:"OsFamily,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`).
	Performance string    `json:"Performance,omitempty"`
	Placement   Placement `json:"Placement,omitempty"`
	// The name of the private DNS.
	PrivateDnsName string `json:"PrivateDnsName,omitempty"`
	// The primary private IP address of the VM.
	PrivateIp string `json:"PrivateIp,omitempty"`
	// The product code associated with the OMI used to create the VM (`0001` Linux/Unix \\| `0002` Windows \\| `0004` Linux/Oracle \\| `0005` Windows 10).
	ProductCodes []string `json:"ProductCodes,omitempty"`
	// The name of the public DNS.
	PublicDnsName string `json:"PublicDnsName,omitempty"`
	// The public IP of the VM.
	PublicIp string `json:"PublicIp,omitempty"`
	// The reservation ID of the VM.
	ReservationId string `json:"ReservationId,omitempty"`
	// The name of the root device for the VM (for example, /dev/vda1).
	RootDeviceName string `json:"RootDeviceName,omitempty"`
	// The type of root device used by the VM (always `bsu`).
	RootDeviceType string `json:"RootDeviceType,omitempty"`
	// One or more security groups associated with the VM.
	SecurityGroups []SecurityGroupLight `json:"SecurityGroups,omitempty"`
	// The state of the VM (`pending` \\| `running` \\| `stopping` \\| `stopped` \\| `shutting-down` \\| `terminated` \\| `quarantine`).
	State string `json:"State,omitempty"`
	// The reason explaining the current state of the VM.
	StateReason string `json:"StateReason,omitempty"`
	// The ID of the Subnet for the VM.
	SubnetId string `json:"SubnetId,omitempty"`
	// One or more tags associated with the VM.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The Base64-encoded MIME user data.
	UserData string `json:"UserData,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId,omitempty"`
	// The VM behavior when you stop it. By default or if set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is deleted.
	VmInitiatedShutdownBehavior string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).
	VmType string `json:"VmType,omitempty"`
}
