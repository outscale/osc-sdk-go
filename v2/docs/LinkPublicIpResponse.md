# LinkPublicIpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkPublicIpId** | Pointer to **string** | (Net only) The ID representing the association of the EIP with the VM or the NIC. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewLinkPublicIpResponse

`func NewLinkPublicIpResponse() *LinkPublicIpResponse`

NewLinkPublicIpResponse instantiates a new LinkPublicIpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPublicIpResponseWithDefaults

`func NewLinkPublicIpResponseWithDefaults() *LinkPublicIpResponse`

NewLinkPublicIpResponseWithDefaults instantiates a new LinkPublicIpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkPublicIpId

`func (o *LinkPublicIpResponse) GetLinkPublicIpId() string`

GetLinkPublicIpId returns the LinkPublicIpId field if non-nil, zero value otherwise.

### GetLinkPublicIpIdOk

`func (o *LinkPublicIpResponse) GetLinkPublicIpIdOk() (*string, bool)`

GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIpId

`func (o *LinkPublicIpResponse) SetLinkPublicIpId(v string)`

SetLinkPublicIpId sets LinkPublicIpId field to given value.

### HasLinkPublicIpId

`func (o *LinkPublicIpResponse) HasLinkPublicIpId() bool`

HasLinkPublicIpId returns a boolean if a field has been set.

### GetResponseContext

`func (o *LinkPublicIpResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *LinkPublicIpResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *LinkPublicIpResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *LinkPublicIpResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


