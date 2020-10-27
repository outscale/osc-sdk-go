# ReadDirectLinksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinks** | Pointer to [**[]DirectLink**](DirectLink.md) | Information about one or more DirectLinks. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadDirectLinksResponse

`func NewReadDirectLinksResponse() *ReadDirectLinksResponse`

NewReadDirectLinksResponse instantiates a new ReadDirectLinksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDirectLinksResponseWithDefaults

`func NewReadDirectLinksResponseWithDefaults() *ReadDirectLinksResponse`

NewReadDirectLinksResponseWithDefaults instantiates a new ReadDirectLinksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLinks

`func (o *ReadDirectLinksResponse) GetDirectLinks() []DirectLink`

GetDirectLinks returns the DirectLinks field if non-nil, zero value otherwise.

### GetDirectLinksOk

`func (o *ReadDirectLinksResponse) GetDirectLinksOk() (*[]DirectLink, bool)`

GetDirectLinksOk returns a tuple with the DirectLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinks

`func (o *ReadDirectLinksResponse) SetDirectLinks(v []DirectLink)`

SetDirectLinks sets DirectLinks field to given value.

### HasDirectLinks

`func (o *ReadDirectLinksResponse) HasDirectLinks() bool`

HasDirectLinks returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadDirectLinksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadDirectLinksResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadDirectLinksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadDirectLinksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


