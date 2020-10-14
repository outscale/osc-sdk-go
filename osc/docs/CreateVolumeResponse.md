# CreateVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Volume** | Pointer to [**Volume**](Volume.md) |  | [optional] 

## Methods

### NewCreateVolumeResponse

`func NewCreateVolumeResponse() *CreateVolumeResponse`

NewCreateVolumeResponse instantiates a new CreateVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeResponseWithDefaults

`func NewCreateVolumeResponseWithDefaults() *CreateVolumeResponse`

NewCreateVolumeResponseWithDefaults instantiates a new CreateVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateVolumeResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateVolumeResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateVolumeResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateVolumeResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVolume

`func (o *CreateVolumeResponse) GetVolume() Volume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateVolumeResponse) GetVolumeOk() (*Volume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateVolumeResponse) SetVolume(v Volume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CreateVolumeResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


