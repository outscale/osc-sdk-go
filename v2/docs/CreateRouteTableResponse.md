# CreateRouteTableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**RouteTable** | Pointer to [**RouteTable**](RouteTable.md) |  | [optional] 

## Methods

### NewCreateRouteTableResponse

`func NewCreateRouteTableResponse() *CreateRouteTableResponse`

NewCreateRouteTableResponse instantiates a new CreateRouteTableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteTableResponseWithDefaults

`func NewCreateRouteTableResponseWithDefaults() *CreateRouteTableResponse`

NewCreateRouteTableResponseWithDefaults instantiates a new CreateRouteTableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateRouteTableResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateRouteTableResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateRouteTableResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateRouteTableResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetRouteTable

`func (o *CreateRouteTableResponse) GetRouteTable() RouteTable`

GetRouteTable returns the RouteTable field if non-nil, zero value otherwise.

### GetRouteTableOk

`func (o *CreateRouteTableResponse) GetRouteTableOk() (*RouteTable, bool)`

GetRouteTableOk returns a tuple with the RouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTable

`func (o *CreateRouteTableResponse) SetRouteTable(v RouteTable)`

SetRouteTable sets RouteTable field to given value.

### HasRouteTable

`func (o *CreateRouteTableResponse) HasRouteTable() bool`

HasRouteTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


