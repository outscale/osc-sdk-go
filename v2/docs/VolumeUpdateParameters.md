# VolumeUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iops** | **NullableInt32** | The new number of I/O operations per second (IOPS):&lt;br /&gt; - For &#x60;io1&#x60; volumes, the number of provisioned IOPS.&lt;br /&gt; - For &#x60;gp2&#x60; volumes, the baseline performance of the volume. | 
**Size** | **int32** | The new size of the volume, in gibibytes (GiB). | 
**VolumeType** | **string** | The type of the volume (&#x60;standard&#x60; \\| &#x60;io1&#x60; \\| &#x60;gp2&#x60;). | 

## Methods

### NewVolumeUpdateParameters

`func NewVolumeUpdateParameters(iops NullableInt32, size int32, volumeType string, ) *VolumeUpdateParameters`

NewVolumeUpdateParameters instantiates a new VolumeUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateParametersWithDefaults

`func NewVolumeUpdateParametersWithDefaults() *VolumeUpdateParameters`

NewVolumeUpdateParametersWithDefaults instantiates a new VolumeUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIops

`func (o *VolumeUpdateParameters) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *VolumeUpdateParameters) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *VolumeUpdateParameters) SetIops(v int32)`

SetIops sets Iops field to given value.


### SetIopsNil

`func (o *VolumeUpdateParameters) SetIopsNil(b bool)`

 SetIopsNil sets the value for Iops to be an explicit nil

### UnsetIops
`func (o *VolumeUpdateParameters) UnsetIops()`

UnsetIops ensures that no value is present for Iops, not even an explicit nil
### GetSize

`func (o *VolumeUpdateParameters) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeUpdateParameters) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeUpdateParameters) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVolumeType

`func (o *VolumeUpdateParameters) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeUpdateParameters) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeUpdateParameters) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


