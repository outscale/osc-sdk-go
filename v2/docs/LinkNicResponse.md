# LinkNicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkNicId** | Pointer to **string** | The ID of the NIC attachment. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewLinkNicResponse

`func NewLinkNicResponse() *LinkNicResponse`

NewLinkNicResponse instantiates a new LinkNicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkNicResponseWithDefaults

`func NewLinkNicResponseWithDefaults() *LinkNicResponse`

NewLinkNicResponseWithDefaults instantiates a new LinkNicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkNicId

`func (o *LinkNicResponse) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *LinkNicResponse) GetLinkNicIdOk() (*string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicId

`func (o *LinkNicResponse) SetLinkNicId(v string)`

SetLinkNicId sets LinkNicId field to given value.

### HasLinkNicId

`func (o *LinkNicResponse) HasLinkNicId() bool`

HasLinkNicId returns a boolean if a field has been set.

### GetResponseContext

`func (o *LinkNicResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *LinkNicResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *LinkNicResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *LinkNicResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


