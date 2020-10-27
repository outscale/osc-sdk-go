# CreatePublicIpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIp** | Pointer to [**PublicIp**](PublicIp.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreatePublicIpResponse

`func NewCreatePublicIpResponse() *CreatePublicIpResponse`

NewCreatePublicIpResponse instantiates a new CreatePublicIpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePublicIpResponseWithDefaults

`func NewCreatePublicIpResponseWithDefaults() *CreatePublicIpResponse`

NewCreatePublicIpResponseWithDefaults instantiates a new CreatePublicIpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIp

`func (o *CreatePublicIpResponse) GetPublicIp() PublicIp`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CreatePublicIpResponse) GetPublicIpOk() (*PublicIp, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CreatePublicIpResponse) SetPublicIp(v PublicIp)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CreatePublicIpResponse) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreatePublicIpResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreatePublicIpResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreatePublicIpResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreatePublicIpResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


