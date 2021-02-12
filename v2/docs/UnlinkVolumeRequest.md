# UnlinkVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**ForceUnlink** | Pointer to **bool** | Forces the detachment of the volume in case of previous failure. Important: This action may damage your data or file systems. | [optional] 
**VolumeId** | **string** | The ID of the volume you want to detach. | 

## Methods

### NewUnlinkVolumeRequest

`func NewUnlinkVolumeRequest(volumeId string, ) *UnlinkVolumeRequest`

NewUnlinkVolumeRequest instantiates a new UnlinkVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkVolumeRequestWithDefaults

`func NewUnlinkVolumeRequestWithDefaults() *UnlinkVolumeRequest`

NewUnlinkVolumeRequestWithDefaults instantiates a new UnlinkVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkVolumeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkVolumeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForceUnlink

`func (o *UnlinkVolumeRequest) GetForceUnlink() bool`

GetForceUnlink returns the ForceUnlink field if non-nil, zero value otherwise.

### GetForceUnlinkOk

`func (o *UnlinkVolumeRequest) GetForceUnlinkOk() (*bool, bool)`

GetForceUnlinkOk returns a tuple with the ForceUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnlink

`func (o *UnlinkVolumeRequest) SetForceUnlink(v bool)`

SetForceUnlink sets ForceUnlink field to given value.

### HasForceUnlink

`func (o *UnlinkVolumeRequest) HasForceUnlink() bool`

HasForceUnlink returns a boolean if a field has been set.

### GetVolumeId

`func (o *UnlinkVolumeRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *UnlinkVolumeRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *UnlinkVolumeRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


