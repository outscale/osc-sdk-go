/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// Vm Information about the VM.
type Vm struct {
	// The architecture of the VM (`i386` \\| `x86_64`).
	Architecture *string `json:"Architecture,omitempty"`
	// The block device mapping of the VM.
	BlockDeviceMappings *[]BlockDeviceMappingCreated `json:"BlockDeviceMappings,omitempty"`
	// This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.
	BsuOptimized *bool `json:"BsuOptimized,omitempty"`
	// The idempotency token provided when launching the VM.
	ClientToken *string `json:"ClientToken,omitempty"`
	// The date and time at which the VM was created.
	CreationDate *string `json:"CreationDate,omitempty"`
	// If true, you cannot delete the VM unless you change this parameter back to false.
	DeletionProtection *bool `json:"DeletionProtection,omitempty"`
	// The hypervisor type of the VMs (`ovm` \\| `xen`).
	Hypervisor *string `json:"Hypervisor,omitempty"`
	// The ID of the OMI used to create the VM.
	ImageId *string `json:"ImageId,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked *bool `json:"IsSourceDestChecked,omitempty"`
	// The name of the keypair used when launching the VM.
	KeypairName *string `json:"KeypairName,omitempty"`
	// The number for the VM when launching a group of several VMs (for example, `0`, `1`, `2`, and so on).
	LaunchNumber *int32 `json:"LaunchNumber,omitempty"`
	// If true, nested virtualization is enabled. If false, it is disabled.
	NestedVirtualization *bool `json:"NestedVirtualization,omitempty"`
	// The ID of the Net in which the VM is running.
	NetId *string `json:"NetId,omitempty"`
	// (Net only) The network interface cards (NICs) the VMs are attached to.
	Nics *[]NicLight `json:"Nics,omitempty"`
	// Indicates the operating system (OS) of the VM.
	OsFamily *string `json:"OsFamily,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`).
	Performance *string    `json:"Performance,omitempty"`
	Placement   *Placement `json:"Placement,omitempty"`
	// The name of the private DNS.
	PrivateDnsName *string `json:"PrivateDnsName,omitempty"`
	// The primary private IP of the VM.
	PrivateIp *string `json:"PrivateIp,omitempty"`
	// The product code associated with the OMI used to create the VM (`0001` Linux/Unix \\| `0002` Windows \\| `0004` Linux/Oracle \\| `0005` Windows 10).
	ProductCodes *[]string `json:"ProductCodes,omitempty"`
	// The name of the public DNS.
	PublicDnsName *string `json:"PublicDnsName,omitempty"`
	// The public IP of the VM.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The reservation ID of the VM.
	ReservationId *string `json:"ReservationId,omitempty"`
	// The name of the root device for the VM (for example, `/dev/vda1`).
	RootDeviceName *string `json:"RootDeviceName,omitempty"`
	// The type of root device used by the VM (always `bsu`).
	RootDeviceType *string `json:"RootDeviceType,omitempty"`
	// One or more security groups associated with the VM.
	SecurityGroups *[]SecurityGroupLight `json:"SecurityGroups,omitempty"`
	// The state of the VM (`pending` \\| `running` \\| `stopping` \\| `stopped` \\| `shutting-down` \\| `terminated` \\| `quarantine`).
	State *string `json:"State,omitempty"`
	// The reason explaining the current state of the VM.
	StateReason *string `json:"StateReason,omitempty"`
	// The ID of the Subnet for the VM.
	SubnetId *string `json:"SubnetId,omitempty"`
	// One or more tags associated with the VM.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The Base64-encoded MIME user data.
	UserData *string `json:"UserData,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
	// The VM behavior when you stop it. If set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is deleted.
	VmInitiatedShutdownBehavior *string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmType *string `json:"VmType,omitempty"`
}

// NewVm instantiates a new Vm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVm() *Vm {
	this := Vm{}
	return &this
}

// NewVmWithDefaults instantiates a new Vm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmWithDefaults() *Vm {
	this := Vm{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *Vm) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *Vm) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *Vm) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBlockDeviceMappings returns the BlockDeviceMappings field value if set, zero value otherwise.
func (o *Vm) GetBlockDeviceMappings() []BlockDeviceMappingCreated {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingCreated
		return ret
	}
	return *o.BlockDeviceMappings
}

// GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetBlockDeviceMappingsOk() (*[]BlockDeviceMappingCreated, bool) {
	if o == nil || o.BlockDeviceMappings == nil {
		return nil, false
	}
	return o.BlockDeviceMappings, true
}

// HasBlockDeviceMappings returns a boolean if a field has been set.
func (o *Vm) HasBlockDeviceMappings() bool {
	if o != nil && o.BlockDeviceMappings != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingCreated and assigns it to the BlockDeviceMappings field.
func (o *Vm) SetBlockDeviceMappings(v []BlockDeviceMappingCreated) {
	o.BlockDeviceMappings = &v
}

// GetBsuOptimized returns the BsuOptimized field value if set, zero value otherwise.
func (o *Vm) GetBsuOptimized() bool {
	if o == nil || o.BsuOptimized == nil {
		var ret bool
		return ret
	}
	return *o.BsuOptimized
}

// GetBsuOptimizedOk returns a tuple with the BsuOptimized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetBsuOptimizedOk() (*bool, bool) {
	if o == nil || o.BsuOptimized == nil {
		return nil, false
	}
	return o.BsuOptimized, true
}

// HasBsuOptimized returns a boolean if a field has been set.
func (o *Vm) HasBsuOptimized() bool {
	if o != nil && o.BsuOptimized != nil {
		return true
	}

	return false
}

// SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.
func (o *Vm) SetBsuOptimized(v bool) {
	o.BsuOptimized = &v
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *Vm) GetClientToken() string {
	if o == nil || o.ClientToken == nil {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetClientTokenOk() (*string, bool) {
	if o == nil || o.ClientToken == nil {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *Vm) HasClientToken() bool {
	if o != nil && o.ClientToken != nil {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *Vm) SetClientToken(v string) {
	o.ClientToken = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Vm) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Vm) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *Vm) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetDeletionProtection returns the DeletionProtection field value if set, zero value otherwise.
func (o *Vm) GetDeletionProtection() bool {
	if o == nil || o.DeletionProtection == nil {
		var ret bool
		return ret
	}
	return *o.DeletionProtection
}

// GetDeletionProtectionOk returns a tuple with the DeletionProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetDeletionProtectionOk() (*bool, bool) {
	if o == nil || o.DeletionProtection == nil {
		return nil, false
	}
	return o.DeletionProtection, true
}

// HasDeletionProtection returns a boolean if a field has been set.
func (o *Vm) HasDeletionProtection() bool {
	if o != nil && o.DeletionProtection != nil {
		return true
	}

	return false
}

// SetDeletionProtection gets a reference to the given bool and assigns it to the DeletionProtection field.
func (o *Vm) SetDeletionProtection(v bool) {
	o.DeletionProtection = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *Vm) GetHypervisor() string {
	if o == nil || o.Hypervisor == nil {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetHypervisorOk() (*string, bool) {
	if o == nil || o.Hypervisor == nil {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *Vm) HasHypervisor() bool {
	if o != nil && o.Hypervisor != nil {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *Vm) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *Vm) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *Vm) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *Vm) SetImageId(v string) {
	o.ImageId = &v
}

// GetIsSourceDestChecked returns the IsSourceDestChecked field value if set, zero value otherwise.
func (o *Vm) GetIsSourceDestChecked() bool {
	if o == nil || o.IsSourceDestChecked == nil {
		var ret bool
		return ret
	}
	return *o.IsSourceDestChecked
}

// GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetIsSourceDestCheckedOk() (*bool, bool) {
	if o == nil || o.IsSourceDestChecked == nil {
		return nil, false
	}
	return o.IsSourceDestChecked, true
}

// HasIsSourceDestChecked returns a boolean if a field has been set.
func (o *Vm) HasIsSourceDestChecked() bool {
	if o != nil && o.IsSourceDestChecked != nil {
		return true
	}

	return false
}

// SetIsSourceDestChecked gets a reference to the given bool and assigns it to the IsSourceDestChecked field.
func (o *Vm) SetIsSourceDestChecked(v bool) {
	o.IsSourceDestChecked = &v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *Vm) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetKeypairNameOk() (*string, bool) {
	if o == nil || o.KeypairName == nil {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *Vm) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *Vm) SetKeypairName(v string) {
	o.KeypairName = &v
}

// GetLaunchNumber returns the LaunchNumber field value if set, zero value otherwise.
func (o *Vm) GetLaunchNumber() int32 {
	if o == nil || o.LaunchNumber == nil {
		var ret int32
		return ret
	}
	return *o.LaunchNumber
}

// GetLaunchNumberOk returns a tuple with the LaunchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetLaunchNumberOk() (*int32, bool) {
	if o == nil || o.LaunchNumber == nil {
		return nil, false
	}
	return o.LaunchNumber, true
}

// HasLaunchNumber returns a boolean if a field has been set.
func (o *Vm) HasLaunchNumber() bool {
	if o != nil && o.LaunchNumber != nil {
		return true
	}

	return false
}

// SetLaunchNumber gets a reference to the given int32 and assigns it to the LaunchNumber field.
func (o *Vm) SetLaunchNumber(v int32) {
	o.LaunchNumber = &v
}

// GetNestedVirtualization returns the NestedVirtualization field value if set, zero value otherwise.
func (o *Vm) GetNestedVirtualization() bool {
	if o == nil || o.NestedVirtualization == nil {
		var ret bool
		return ret
	}
	return *o.NestedVirtualization
}

// GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetNestedVirtualizationOk() (*bool, bool) {
	if o == nil || o.NestedVirtualization == nil {
		return nil, false
	}
	return o.NestedVirtualization, true
}

// HasNestedVirtualization returns a boolean if a field has been set.
func (o *Vm) HasNestedVirtualization() bool {
	if o != nil && o.NestedVirtualization != nil {
		return true
	}

	return false
}

// SetNestedVirtualization gets a reference to the given bool and assigns it to the NestedVirtualization field.
func (o *Vm) SetNestedVirtualization(v bool) {
	o.NestedVirtualization = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *Vm) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *Vm) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *Vm) SetNetId(v string) {
	o.NetId = &v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *Vm) GetNics() []NicLight {
	if o == nil || o.Nics == nil {
		var ret []NicLight
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetNicsOk() (*[]NicLight, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *Vm) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []NicLight and assigns it to the Nics field.
func (o *Vm) SetNics(v []NicLight) {
	o.Nics = &v
}

// GetOsFamily returns the OsFamily field value if set, zero value otherwise.
func (o *Vm) GetOsFamily() string {
	if o == nil || o.OsFamily == nil {
		var ret string
		return ret
	}
	return *o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetOsFamilyOk() (*string, bool) {
	if o == nil || o.OsFamily == nil {
		return nil, false
	}
	return o.OsFamily, true
}

// HasOsFamily returns a boolean if a field has been set.
func (o *Vm) HasOsFamily() bool {
	if o != nil && o.OsFamily != nil {
		return true
	}

	return false
}

// SetOsFamily gets a reference to the given string and assigns it to the OsFamily field.
func (o *Vm) SetOsFamily(v string) {
	o.OsFamily = &v
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *Vm) GetPerformance() string {
	if o == nil || o.Performance == nil {
		var ret string
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetPerformanceOk() (*string, bool) {
	if o == nil || o.Performance == nil {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *Vm) HasPerformance() bool {
	if o != nil && o.Performance != nil {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given string and assigns it to the Performance field.
func (o *Vm) SetPerformance(v string) {
	o.Performance = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *Vm) GetPlacement() Placement {
	if o == nil || o.Placement == nil {
		var ret Placement
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetPlacementOk() (*Placement, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *Vm) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given Placement and assigns it to the Placement field.
func (o *Vm) SetPlacement(v Placement) {
	o.Placement = &v
}

// GetPrivateDnsName returns the PrivateDnsName field value if set, zero value otherwise.
func (o *Vm) GetPrivateDnsName() string {
	if o == nil || o.PrivateDnsName == nil {
		var ret string
		return ret
	}
	return *o.PrivateDnsName
}

// GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetPrivateDnsNameOk() (*string, bool) {
	if o == nil || o.PrivateDnsName == nil {
		return nil, false
	}
	return o.PrivateDnsName, true
}

// HasPrivateDnsName returns a boolean if a field has been set.
func (o *Vm) HasPrivateDnsName() bool {
	if o != nil && o.PrivateDnsName != nil {
		return true
	}

	return false
}

// SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.
func (o *Vm) SetPrivateDnsName(v string) {
	o.PrivateDnsName = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *Vm) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetPrivateIpOk() (*string, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *Vm) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *Vm) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetProductCodes returns the ProductCodes field value if set, zero value otherwise.
func (o *Vm) GetProductCodes() []string {
	if o == nil || o.ProductCodes == nil {
		var ret []string
		return ret
	}
	return *o.ProductCodes
}

// GetProductCodesOk returns a tuple with the ProductCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetProductCodesOk() (*[]string, bool) {
	if o == nil || o.ProductCodes == nil {
		return nil, false
	}
	return o.ProductCodes, true
}

// HasProductCodes returns a boolean if a field has been set.
func (o *Vm) HasProductCodes() bool {
	if o != nil && o.ProductCodes != nil {
		return true
	}

	return false
}

// SetProductCodes gets a reference to the given []string and assigns it to the ProductCodes field.
func (o *Vm) SetProductCodes(v []string) {
	o.ProductCodes = &v
}

// GetPublicDnsName returns the PublicDnsName field value if set, zero value otherwise.
func (o *Vm) GetPublicDnsName() string {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret
	}
	return *o.PublicDnsName
}

// GetPublicDnsNameOk returns a tuple with the PublicDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetPublicDnsNameOk() (*string, bool) {
	if o == nil || o.PublicDnsName == nil {
		return nil, false
	}
	return o.PublicDnsName, true
}

// HasPublicDnsName returns a boolean if a field has been set.
func (o *Vm) HasPublicDnsName() bool {
	if o != nil && o.PublicDnsName != nil {
		return true
	}

	return false
}

// SetPublicDnsName gets a reference to the given string and assigns it to the PublicDnsName field.
func (o *Vm) SetPublicDnsName(v string) {
	o.PublicDnsName = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *Vm) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *Vm) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *Vm) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *Vm) GetReservationId() string {
	if o == nil || o.ReservationId == nil {
		var ret string
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetReservationIdOk() (*string, bool) {
	if o == nil || o.ReservationId == nil {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *Vm) HasReservationId() bool {
	if o != nil && o.ReservationId != nil {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given string and assigns it to the ReservationId field.
func (o *Vm) SetReservationId(v string) {
	o.ReservationId = &v
}

// GetRootDeviceName returns the RootDeviceName field value if set, zero value otherwise.
func (o *Vm) GetRootDeviceName() string {
	if o == nil || o.RootDeviceName == nil {
		var ret string
		return ret
	}
	return *o.RootDeviceName
}

// GetRootDeviceNameOk returns a tuple with the RootDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetRootDeviceNameOk() (*string, bool) {
	if o == nil || o.RootDeviceName == nil {
		return nil, false
	}
	return o.RootDeviceName, true
}

// HasRootDeviceName returns a boolean if a field has been set.
func (o *Vm) HasRootDeviceName() bool {
	if o != nil && o.RootDeviceName != nil {
		return true
	}

	return false
}

// SetRootDeviceName gets a reference to the given string and assigns it to the RootDeviceName field.
func (o *Vm) SetRootDeviceName(v string) {
	o.RootDeviceName = &v
}

// GetRootDeviceType returns the RootDeviceType field value if set, zero value otherwise.
func (o *Vm) GetRootDeviceType() string {
	if o == nil || o.RootDeviceType == nil {
		var ret string
		return ret
	}
	return *o.RootDeviceType
}

// GetRootDeviceTypeOk returns a tuple with the RootDeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetRootDeviceTypeOk() (*string, bool) {
	if o == nil || o.RootDeviceType == nil {
		return nil, false
	}
	return o.RootDeviceType, true
}

// HasRootDeviceType returns a boolean if a field has been set.
func (o *Vm) HasRootDeviceType() bool {
	if o != nil && o.RootDeviceType != nil {
		return true
	}

	return false
}

// SetRootDeviceType gets a reference to the given string and assigns it to the RootDeviceType field.
func (o *Vm) SetRootDeviceType(v string) {
	o.RootDeviceType = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *Vm) GetSecurityGroups() []SecurityGroupLight {
	if o == nil || o.SecurityGroups == nil {
		var ret []SecurityGroupLight
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetSecurityGroupsOk() (*[]SecurityGroupLight, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *Vm) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []SecurityGroupLight and assigns it to the SecurityGroups field.
func (o *Vm) SetSecurityGroups(v []SecurityGroupLight) {
	o.SecurityGroups = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Vm) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Vm) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Vm) SetState(v string) {
	o.State = &v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *Vm) GetStateReason() string {
	if o == nil || o.StateReason == nil {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetStateReasonOk() (*string, bool) {
	if o == nil || o.StateReason == nil {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *Vm) HasStateReason() bool {
	if o != nil && o.StateReason != nil {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *Vm) SetStateReason(v string) {
	o.StateReason = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *Vm) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *Vm) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *Vm) SetSubnetId(v string) {
	o.SubnetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Vm) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Vm) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *Vm) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *Vm) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *Vm) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *Vm) SetUserData(v string) {
	o.UserData = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *Vm) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *Vm) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *Vm) SetVmId(v string) {
	o.VmId = &v
}

// GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field value if set, zero value otherwise.
func (o *Vm) GetVmInitiatedShutdownBehavior() string {
	if o == nil || o.VmInitiatedShutdownBehavior == nil {
		var ret string
		return ret
	}
	return *o.VmInitiatedShutdownBehavior
}

// GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetVmInitiatedShutdownBehaviorOk() (*string, bool) {
	if o == nil || o.VmInitiatedShutdownBehavior == nil {
		return nil, false
	}
	return o.VmInitiatedShutdownBehavior, true
}

// HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.
func (o *Vm) HasVmInitiatedShutdownBehavior() bool {
	if o != nil && o.VmInitiatedShutdownBehavior != nil {
		return true
	}

	return false
}

// SetVmInitiatedShutdownBehavior gets a reference to the given string and assigns it to the VmInitiatedShutdownBehavior field.
func (o *Vm) SetVmInitiatedShutdownBehavior(v string) {
	o.VmInitiatedShutdownBehavior = &v
}

// GetVmType returns the VmType field value if set, zero value otherwise.
func (o *Vm) GetVmType() string {
	if o == nil || o.VmType == nil {
		var ret string
		return ret
	}
	return *o.VmType
}

// GetVmTypeOk returns a tuple with the VmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetVmTypeOk() (*string, bool) {
	if o == nil || o.VmType == nil {
		return nil, false
	}
	return o.VmType, true
}

// HasVmType returns a boolean if a field has been set.
func (o *Vm) HasVmType() bool {
	if o != nil && o.VmType != nil {
		return true
	}

	return false
}

// SetVmType gets a reference to the given string and assigns it to the VmType field.
func (o *Vm) SetVmType(v string) {
	o.VmType = &v
}

func (o Vm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Architecture != nil {
		toSerialize["Architecture"] = o.Architecture
	}
	if o.BlockDeviceMappings != nil {
		toSerialize["BlockDeviceMappings"] = o.BlockDeviceMappings
	}
	if o.BsuOptimized != nil {
		toSerialize["BsuOptimized"] = o.BsuOptimized
	}
	if o.ClientToken != nil {
		toSerialize["ClientToken"] = o.ClientToken
	}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.DeletionProtection != nil {
		toSerialize["DeletionProtection"] = o.DeletionProtection
	}
	if o.Hypervisor != nil {
		toSerialize["Hypervisor"] = o.Hypervisor
	}
	if o.ImageId != nil {
		toSerialize["ImageId"] = o.ImageId
	}
	if o.IsSourceDestChecked != nil {
		toSerialize["IsSourceDestChecked"] = o.IsSourceDestChecked
	}
	if o.KeypairName != nil {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if o.LaunchNumber != nil {
		toSerialize["LaunchNumber"] = o.LaunchNumber
	}
	if o.NestedVirtualization != nil {
		toSerialize["NestedVirtualization"] = o.NestedVirtualization
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.Nics != nil {
		toSerialize["Nics"] = o.Nics
	}
	if o.OsFamily != nil {
		toSerialize["OsFamily"] = o.OsFamily
	}
	if o.Performance != nil {
		toSerialize["Performance"] = o.Performance
	}
	if o.Placement != nil {
		toSerialize["Placement"] = o.Placement
	}
	if o.PrivateDnsName != nil {
		toSerialize["PrivateDnsName"] = o.PrivateDnsName
	}
	if o.PrivateIp != nil {
		toSerialize["PrivateIp"] = o.PrivateIp
	}
	if o.ProductCodes != nil {
		toSerialize["ProductCodes"] = o.ProductCodes
	}
	if o.PublicDnsName != nil {
		toSerialize["PublicDnsName"] = o.PublicDnsName
	}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.ReservationId != nil {
		toSerialize["ReservationId"] = o.ReservationId
	}
	if o.RootDeviceName != nil {
		toSerialize["RootDeviceName"] = o.RootDeviceName
	}
	if o.RootDeviceType != nil {
		toSerialize["RootDeviceType"] = o.RootDeviceType
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.StateReason != nil {
		toSerialize["StateReason"] = o.StateReason
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.UserData != nil {
		toSerialize["UserData"] = o.UserData
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	if o.VmInitiatedShutdownBehavior != nil {
		toSerialize["VmInitiatedShutdownBehavior"] = o.VmInitiatedShutdownBehavior
	}
	if o.VmType != nil {
		toSerialize["VmType"] = o.VmType
	}
	return json.Marshal(toSerialize)
}

type NullableVm struct {
	value *Vm
	isSet bool
}

func (v NullableVm) Get() *Vm {
	return v.value
}

func (v *NullableVm) Set(val *Vm) {
	v.value = val
	v.isSet = true
}

func (v NullableVm) IsSet() bool {
	return v.isSet
}

func (v *NullableVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVm(val *Vm) *NullableVm {
	return &NullableVm{value: val, isSet: true}
}

func (v NullableVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
