# CreateNetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Net** | Pointer to [**Net**](Net.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateNetResponse

`func NewCreateNetResponse() *CreateNetResponse`

NewCreateNetResponse instantiates a new CreateNetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetResponseWithDefaults

`func NewCreateNetResponseWithDefaults() *CreateNetResponse`

NewCreateNetResponseWithDefaults instantiates a new CreateNetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNet

`func (o *CreateNetResponse) GetNet() Net`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *CreateNetResponse) GetNetOk() (*Net, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *CreateNetResponse) SetNet(v Net)`

SetNet sets Net field to given value.

### HasNet

`func (o *CreateNetResponse) HasNet() bool`

HasNet returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateNetResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateNetResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateNetResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateNetResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


