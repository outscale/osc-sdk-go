# UpdateNicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nic** | Pointer to [**Nic**](Nic.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewUpdateNicResponse

`func NewUpdateNicResponse() *UpdateNicResponse`

NewUpdateNicResponse instantiates a new UpdateNicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNicResponseWithDefaults

`func NewUpdateNicResponseWithDefaults() *UpdateNicResponse`

NewUpdateNicResponseWithDefaults instantiates a new UpdateNicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNic

`func (o *UpdateNicResponse) GetNic() Nic`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *UpdateNicResponse) GetNicOk() (*Nic, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *UpdateNicResponse) SetNic(v Nic)`

SetNic sets Nic field to given value.

### HasNic

`func (o *UpdateNicResponse) HasNic() bool`

HasNic returns a boolean if a field has been set.

### GetResponseContext

`func (o *UpdateNicResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UpdateNicResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UpdateNicResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UpdateNicResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


