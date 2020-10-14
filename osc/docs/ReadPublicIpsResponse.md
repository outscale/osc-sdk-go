# ReadPublicIpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIps** | Pointer to [**[]PublicIp**](PublicIp.md) | Information about one or more EIPs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadPublicIpsResponse

`func NewReadPublicIpsResponse() *ReadPublicIpsResponse`

NewReadPublicIpsResponse instantiates a new ReadPublicIpsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPublicIpsResponseWithDefaults

`func NewReadPublicIpsResponseWithDefaults() *ReadPublicIpsResponse`

NewReadPublicIpsResponseWithDefaults instantiates a new ReadPublicIpsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIps

`func (o *ReadPublicIpsResponse) GetPublicIps() []PublicIp`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *ReadPublicIpsResponse) GetPublicIpsOk() (*[]PublicIp, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *ReadPublicIpsResponse) SetPublicIps(v []PublicIp)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *ReadPublicIpsResponse) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadPublicIpsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPublicIpsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadPublicIpsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadPublicIpsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


