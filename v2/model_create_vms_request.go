/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateVmsRequest struct for CreateVmsRequest
type CreateVmsRequest struct {
	// One or more block device mappings.
	BlockDeviceMappings *[]BlockDeviceMappingVmCreation `json:"BlockDeviceMappings,omitempty"`
	// By default or if true, the VM is started on creation. If false, the VM is stopped on creation.
	BootOnCreation *bool `json:"BootOnCreation,omitempty"`
	// If true, the VM is created with optimized BSU I/O.
	BsuOptimized *bool `json:"BsuOptimized,omitempty"`
	// A unique identifier which enables you to manage the idempotency.
	ClientToken *string `json:"ClientToken,omitempty"`
	// If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.
	DeletionProtection *bool `json:"DeletionProtection,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](#readimages) method.
	ImageId string `json:"ImageId"`
	// The name of the keypair.
	KeypairName *string `json:"KeypairName,omitempty"`
	// The maximum number of VMs you want to create. If all the VMs cannot be created, the largest possible number of VMs above MinVmsCount is created.
	MaxVmsCount *int32 `json:"MaxVmsCount,omitempty"`
	// The minimum number of VMs you want to create. If this number of VMs cannot be created, no VMs are created.
	MinVmsCount *int32 `json:"MinVmsCount,omitempty"`
	// One or more NICs. If you specify this parameter, you must define one NIC as the primary network interface of the VM with `0` as its device number.
	Nics *[]NicForVmCreation `json:"Nics,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`).
	Performance *string    `json:"Performance,omitempty"`
	Placement   *Placement `json:"Placement,omitempty"`
	// One or more private IPs of the VM.
	PrivateIps *[]string `json:"PrivateIps,omitempty"`
	// One or more IDs of security group for the VMs.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// One or more names of security groups for the VMs.
	SecurityGroups *[]string `json:"SecurityGroups,omitempty"`
	// The ID of the Subnet in which you want to create the VM.
	SubnetId *string `json:"SubnetId,omitempty"`
	// Data or script used to add a specific configuration to the VM. It must be Base64-encoded and is limited to 500 kibibytes (KiB).
	UserData *string `json:"UserData,omitempty"`
	// The VM behavior when you stop it. By default or if set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is terminated.
	VmInitiatedShutdownBehavior *string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM (`t2.small` by default).<br /> For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).
	VmType *string `json:"VmType,omitempty"`
}

// NewCreateVmsRequest instantiates a new CreateVmsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVmsRequest(imageId string) *CreateVmsRequest {
	this := CreateVmsRequest{}
	this.ImageId = imageId
	var performance string = "high"
	this.Performance = &performance
	return &this
}

// NewCreateVmsRequestWithDefaults instantiates a new CreateVmsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVmsRequestWithDefaults() *CreateVmsRequest {
	this := CreateVmsRequest{}
	var performance string = "high"
	this.Performance = &performance
	return &this
}

// GetBlockDeviceMappings returns the BlockDeviceMappings field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetBlockDeviceMappings() []BlockDeviceMappingVmCreation {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingVmCreation
		return ret
	}
	return *o.BlockDeviceMappings
}

// GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetBlockDeviceMappingsOk() (*[]BlockDeviceMappingVmCreation, bool) {
	if o == nil || o.BlockDeviceMappings == nil {
		return nil, false
	}
	return o.BlockDeviceMappings, true
}

// HasBlockDeviceMappings returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasBlockDeviceMappings() bool {
	if o != nil && o.BlockDeviceMappings != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingVmCreation and assigns it to the BlockDeviceMappings field.
func (o *CreateVmsRequest) SetBlockDeviceMappings(v []BlockDeviceMappingVmCreation) {
	o.BlockDeviceMappings = &v
}

// GetBootOnCreation returns the BootOnCreation field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetBootOnCreation() bool {
	if o == nil || o.BootOnCreation == nil {
		var ret bool
		return ret
	}
	return *o.BootOnCreation
}

// GetBootOnCreationOk returns a tuple with the BootOnCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetBootOnCreationOk() (*bool, bool) {
	if o == nil || o.BootOnCreation == nil {
		return nil, false
	}
	return o.BootOnCreation, true
}

// HasBootOnCreation returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasBootOnCreation() bool {
	if o != nil && o.BootOnCreation != nil {
		return true
	}

	return false
}

// SetBootOnCreation gets a reference to the given bool and assigns it to the BootOnCreation field.
func (o *CreateVmsRequest) SetBootOnCreation(v bool) {
	o.BootOnCreation = &v
}

// GetBsuOptimized returns the BsuOptimized field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetBsuOptimized() bool {
	if o == nil || o.BsuOptimized == nil {
		var ret bool
		return ret
	}
	return *o.BsuOptimized
}

// GetBsuOptimizedOk returns a tuple with the BsuOptimized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetBsuOptimizedOk() (*bool, bool) {
	if o == nil || o.BsuOptimized == nil {
		return nil, false
	}
	return o.BsuOptimized, true
}

// HasBsuOptimized returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasBsuOptimized() bool {
	if o != nil && o.BsuOptimized != nil {
		return true
	}

	return false
}

// SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.
func (o *CreateVmsRequest) SetBsuOptimized(v bool) {
	o.BsuOptimized = &v
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetClientToken() string {
	if o == nil || o.ClientToken == nil {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetClientTokenOk() (*string, bool) {
	if o == nil || o.ClientToken == nil {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasClientToken() bool {
	if o != nil && o.ClientToken != nil {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *CreateVmsRequest) SetClientToken(v string) {
	o.ClientToken = &v
}

// GetDeletionProtection returns the DeletionProtection field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetDeletionProtection() bool {
	if o == nil || o.DeletionProtection == nil {
		var ret bool
		return ret
	}
	return *o.DeletionProtection
}

// GetDeletionProtectionOk returns a tuple with the DeletionProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetDeletionProtectionOk() (*bool, bool) {
	if o == nil || o.DeletionProtection == nil {
		return nil, false
	}
	return o.DeletionProtection, true
}

// HasDeletionProtection returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasDeletionProtection() bool {
	if o != nil && o.DeletionProtection != nil {
		return true
	}

	return false
}

// SetDeletionProtection gets a reference to the given bool and assigns it to the DeletionProtection field.
func (o *CreateVmsRequest) SetDeletionProtection(v bool) {
	o.DeletionProtection = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateVmsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetImageId returns the ImageId field value
func (o *CreateVmsRequest) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *CreateVmsRequest) SetImageId(v string) {
	o.ImageId = v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetKeypairNameOk() (*string, bool) {
	if o == nil || o.KeypairName == nil {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *CreateVmsRequest) SetKeypairName(v string) {
	o.KeypairName = &v
}

// GetMaxVmsCount returns the MaxVmsCount field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetMaxVmsCount() int32 {
	if o == nil || o.MaxVmsCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxVmsCount
}

// GetMaxVmsCountOk returns a tuple with the MaxVmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetMaxVmsCountOk() (*int32, bool) {
	if o == nil || o.MaxVmsCount == nil {
		return nil, false
	}
	return o.MaxVmsCount, true
}

// HasMaxVmsCount returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasMaxVmsCount() bool {
	if o != nil && o.MaxVmsCount != nil {
		return true
	}

	return false
}

// SetMaxVmsCount gets a reference to the given int32 and assigns it to the MaxVmsCount field.
func (o *CreateVmsRequest) SetMaxVmsCount(v int32) {
	o.MaxVmsCount = &v
}

// GetMinVmsCount returns the MinVmsCount field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetMinVmsCount() int32 {
	if o == nil || o.MinVmsCount == nil {
		var ret int32
		return ret
	}
	return *o.MinVmsCount
}

// GetMinVmsCountOk returns a tuple with the MinVmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetMinVmsCountOk() (*int32, bool) {
	if o == nil || o.MinVmsCount == nil {
		return nil, false
	}
	return o.MinVmsCount, true
}

// HasMinVmsCount returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasMinVmsCount() bool {
	if o != nil && o.MinVmsCount != nil {
		return true
	}

	return false
}

// SetMinVmsCount gets a reference to the given int32 and assigns it to the MinVmsCount field.
func (o *CreateVmsRequest) SetMinVmsCount(v int32) {
	o.MinVmsCount = &v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetNics() []NicForVmCreation {
	if o == nil || o.Nics == nil {
		var ret []NicForVmCreation
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetNicsOk() (*[]NicForVmCreation, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []NicForVmCreation and assigns it to the Nics field.
func (o *CreateVmsRequest) SetNics(v []NicForVmCreation) {
	o.Nics = &v
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetPerformance() string {
	if o == nil || o.Performance == nil {
		var ret string
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetPerformanceOk() (*string, bool) {
	if o == nil || o.Performance == nil {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasPerformance() bool {
	if o != nil && o.Performance != nil {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given string and assigns it to the Performance field.
func (o *CreateVmsRequest) SetPerformance(v string) {
	o.Performance = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetPlacement() Placement {
	if o == nil || o.Placement == nil {
		var ret Placement
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetPlacementOk() (*Placement, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given Placement and assigns it to the Placement field.
func (o *CreateVmsRequest) SetPlacement(v Placement) {
	o.Placement = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetPrivateIps() []string {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetPrivateIpsOk() (*[]string, bool) {
	if o == nil || o.PrivateIps == nil {
		return nil, false
	}
	return o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.
func (o *CreateVmsRequest) SetPrivateIps(v []string) {
	o.PrivateIps = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetSecurityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *CreateVmsRequest) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetSecurityGroups() []string {
	if o == nil || o.SecurityGroups == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *CreateVmsRequest) SetSecurityGroups(v []string) {
	o.SecurityGroups = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *CreateVmsRequest) SetSubnetId(v string) {
	o.SubnetId = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *CreateVmsRequest) SetUserData(v string) {
	o.UserData = &v
}

// GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetVmInitiatedShutdownBehavior() string {
	if o == nil || o.VmInitiatedShutdownBehavior == nil {
		var ret string
		return ret
	}
	return *o.VmInitiatedShutdownBehavior
}

// GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetVmInitiatedShutdownBehaviorOk() (*string, bool) {
	if o == nil || o.VmInitiatedShutdownBehavior == nil {
		return nil, false
	}
	return o.VmInitiatedShutdownBehavior, true
}

// HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasVmInitiatedShutdownBehavior() bool {
	if o != nil && o.VmInitiatedShutdownBehavior != nil {
		return true
	}

	return false
}

// SetVmInitiatedShutdownBehavior gets a reference to the given string and assigns it to the VmInitiatedShutdownBehavior field.
func (o *CreateVmsRequest) SetVmInitiatedShutdownBehavior(v string) {
	o.VmInitiatedShutdownBehavior = &v
}

// GetVmType returns the VmType field value if set, zero value otherwise.
func (o *CreateVmsRequest) GetVmType() string {
	if o == nil || o.VmType == nil {
		var ret string
		return ret
	}
	return *o.VmType
}

// GetVmTypeOk returns a tuple with the VmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmsRequest) GetVmTypeOk() (*string, bool) {
	if o == nil || o.VmType == nil {
		return nil, false
	}
	return o.VmType, true
}

// HasVmType returns a boolean if a field has been set.
func (o *CreateVmsRequest) HasVmType() bool {
	if o != nil && o.VmType != nil {
		return true
	}

	return false
}

// SetVmType gets a reference to the given string and assigns it to the VmType field.
func (o *CreateVmsRequest) SetVmType(v string) {
	o.VmType = &v
}

func (o CreateVmsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockDeviceMappings != nil {
		toSerialize["BlockDeviceMappings"] = o.BlockDeviceMappings
	}
	if o.BootOnCreation != nil {
		toSerialize["BootOnCreation"] = o.BootOnCreation
	}
	if o.BsuOptimized != nil {
		toSerialize["BsuOptimized"] = o.BsuOptimized
	}
	if o.ClientToken != nil {
		toSerialize["ClientToken"] = o.ClientToken
	}
	if o.DeletionProtection != nil {
		toSerialize["DeletionProtection"] = o.DeletionProtection
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ImageId"] = o.ImageId
	}
	if o.KeypairName != nil {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if o.MaxVmsCount != nil {
		toSerialize["MaxVmsCount"] = o.MaxVmsCount
	}
	if o.MinVmsCount != nil {
		toSerialize["MinVmsCount"] = o.MinVmsCount
	}
	if o.Nics != nil {
		toSerialize["Nics"] = o.Nics
	}
	if o.Performance != nil {
		toSerialize["Performance"] = o.Performance
	}
	if o.Placement != nil {
		toSerialize["Placement"] = o.Placement
	}
	if o.PrivateIps != nil {
		toSerialize["PrivateIps"] = o.PrivateIps
	}
	if o.SecurityGroupIds != nil {
		toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	if o.UserData != nil {
		toSerialize["UserData"] = o.UserData
	}
	if o.VmInitiatedShutdownBehavior != nil {
		toSerialize["VmInitiatedShutdownBehavior"] = o.VmInitiatedShutdownBehavior
	}
	if o.VmType != nil {
		toSerialize["VmType"] = o.VmType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVmsRequest struct {
	value *CreateVmsRequest
	isSet bool
}

func (v NullableCreateVmsRequest) Get() *CreateVmsRequest {
	return v.value
}

func (v *NullableCreateVmsRequest) Set(val *CreateVmsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVmsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVmsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVmsRequest(val *CreateVmsRequest) *NullableCreateVmsRequest {
	return &NullableCreateVmsRequest{value: val, isSet: true}
}

func (v NullableCreateVmsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVmsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
