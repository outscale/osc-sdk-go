# DeleteVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VolumeId** | **string** | The ID of the volume you want to delete. | 

## Methods

### NewDeleteVolumeRequest

`func NewDeleteVolumeRequest(volumeId string, ) *DeleteVolumeRequest`

NewDeleteVolumeRequest instantiates a new DeleteVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVolumeRequestWithDefaults

`func NewDeleteVolumeRequestWithDefaults() *DeleteVolumeRequest`

NewDeleteVolumeRequestWithDefaults instantiates a new DeleteVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVolumeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteVolumeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetVolumeId

`func (o *DeleteVolumeRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *DeleteVolumeRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *DeleteVolumeRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


