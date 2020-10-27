# LinkRouteTableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkRouteTableId** | Pointer to **string** | The ID of the association between the route table and the Subnet. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewLinkRouteTableResponse

`func NewLinkRouteTableResponse() *LinkRouteTableResponse`

NewLinkRouteTableResponse instantiates a new LinkRouteTableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRouteTableResponseWithDefaults

`func NewLinkRouteTableResponseWithDefaults() *LinkRouteTableResponse`

NewLinkRouteTableResponseWithDefaults instantiates a new LinkRouteTableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkRouteTableId

`func (o *LinkRouteTableResponse) GetLinkRouteTableId() string`

GetLinkRouteTableId returns the LinkRouteTableId field if non-nil, zero value otherwise.

### GetLinkRouteTableIdOk

`func (o *LinkRouteTableResponse) GetLinkRouteTableIdOk() (*string, bool)`

GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRouteTableId

`func (o *LinkRouteTableResponse) SetLinkRouteTableId(v string)`

SetLinkRouteTableId sets LinkRouteTableId field to given value.

### HasLinkRouteTableId

`func (o *LinkRouteTableResponse) HasLinkRouteTableId() bool`

HasLinkRouteTableId returns a boolean if a field has been set.

### GetResponseContext

`func (o *LinkRouteTableResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *LinkRouteTableResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *LinkRouteTableResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *LinkRouteTableResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


