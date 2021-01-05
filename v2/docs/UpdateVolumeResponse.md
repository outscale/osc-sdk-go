# UpdateVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Volume** | Pointer to [**Volume**](Volume.md) |  | [optional] 

## Methods

### NewUpdateVolumeResponse

`func NewUpdateVolumeResponse() *UpdateVolumeResponse`

NewUpdateVolumeResponse instantiates a new UpdateVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeResponseWithDefaults

`func NewUpdateVolumeResponseWithDefaults() *UpdateVolumeResponse`

NewUpdateVolumeResponseWithDefaults instantiates a new UpdateVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *UpdateVolumeResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateVolumeResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateVolumeResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateVolumeResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVolume

`func (o *UpdateVolumeResponse) GetVolume() Volume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *UpdateVolumeResponse) GetVolumeOk() (*Volume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *UpdateVolumeResponse) SetVolume(v Volume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *UpdateVolumeResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


