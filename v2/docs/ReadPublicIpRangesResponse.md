# ReadPublicIpRangesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIps** | Pointer to **[]string** | The list of public IPv4 addresses used in the Region, in CIDR notation. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadPublicIpRangesResponse

`func NewReadPublicIpRangesResponse() *ReadPublicIpRangesResponse`

NewReadPublicIpRangesResponse instantiates a new ReadPublicIpRangesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadPublicIpRangesResponseWithDefaults

`func NewReadPublicIpRangesResponseWithDefaults() *ReadPublicIpRangesResponse`

NewReadPublicIpRangesResponseWithDefaults instantiates a new ReadPublicIpRangesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIps

`func (o *ReadPublicIpRangesResponse) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *ReadPublicIpRangesResponse) GetPublicIpsOk() (*[]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *ReadPublicIpRangesResponse) SetPublicIps(v []string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *ReadPublicIpRangesResponse) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadPublicIpRangesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPublicIpRangesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadPublicIpRangesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadPublicIpRangesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


