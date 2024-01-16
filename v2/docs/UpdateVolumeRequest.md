# UpdateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Iops** | Pointer to **int32** | **Cold volume**: the new number of I/O operations per second (IOPS). This parameter can be specified only if you update an &#x60;io1&#x60; volume or if you change the type of the volume for an &#x60;io1&#x60;. This modification is instantaneous. &lt;br /&gt; **Hot volume**: the new number of I/O operations per second (IOPS). This parameter can be specified only if you update an &#x60;io1&#x60; volume. This modification is not instantaneous. &lt;br /&gt;&lt;br /&gt; The maximum number of IOPS allowed for &#x60;io1&#x60; volumes is &#x60;13000&#x60; with a maximum performance ratio of 300 IOPS per gibibyte. | [optional] 
**Size** | Pointer to **int32** | **Cold volume**: the new size of the volume, in gibibytes (GiB). This value must be equal to or greater than the current size of the volume. This modification is not instantaneous. &lt;br /&gt; **Hot volume**: you cannot change the size of a hot volume. | [optional] 
**VolumeId** | **string** | The ID of the volume you want to update. | 
**VolumeType** | Pointer to **string** | **Cold volume**: the new type of the volume (&#x60;standard&#x60; \\| &#x60;io1&#x60; \\| &#x60;gp2&#x60;). This modification is instantaneous. If you update to an &#x60;io1&#x60; volume, you must also specify the &#x60;Iops&#x60; parameter.&lt;br /&gt; **Hot volume**: you cannot change the type of a hot volume. | [optional] 

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

### GetIops

`func (o *UpdateVolumeRequest) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *UpdateVolumeRequest) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *UpdateVolumeRequest) SetIops(v int32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *UpdateVolumeRequest) HasIops() bool`

HasIops returns a boolean if a field has been set.

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


### GetVolumeType

`func (o *UpdateVolumeRequest) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *UpdateVolumeRequest) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *UpdateVolumeRequest) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *UpdateVolumeRequest) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


