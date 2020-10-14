# CreateNicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nic** | Pointer to [**Nic**](Nic.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateNicResponse

`func NewCreateNicResponse() *CreateNicResponse`

NewCreateNicResponse instantiates a new CreateNicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNicResponseWithDefaults

`func NewCreateNicResponseWithDefaults() *CreateNicResponse`

NewCreateNicResponseWithDefaults instantiates a new CreateNicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNic

`func (o *CreateNicResponse) GetNic() Nic`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *CreateNicResponse) GetNicOk() (*Nic, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *CreateNicResponse) SetNic(v Nic)`

SetNic sets Nic field to given value.

### HasNic

`func (o *CreateNicResponse) HasNic() bool`

HasNic returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateNicResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateNicResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateNicResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateNicResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


