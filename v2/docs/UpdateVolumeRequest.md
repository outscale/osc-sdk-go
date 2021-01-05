# UpdateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Size** | Pointer to **int32** | The new size of the volume, in gibibytes (GiB). This value must be equal to or greater than the current size of the volume. | [optional] 
**VolumeId** | **string** | The ID of the volume you want to update. | 

## Methods

### NewUpdateVolumeRequest

`func NewUpdateVolumeRequest(volumeId string, ) *UpdateVolumeRequest`

NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeRequestWithDefaults

`func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest`

NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UpdateVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateVolumeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateVolumeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSize

`func (o *UpdateVolumeRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateVolumeRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateVolumeRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateVolumeRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeId

`func (o *UpdateVolumeRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *UpdateVolumeRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *UpdateVolumeRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


