# CreateRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**RouteTable** | Pointer to [**RouteTable**](RouteTable.md) |  | [optional] 

## Methods

### NewCreateRouteResponse

`func NewCreateRouteResponse() *CreateRouteResponse`

NewCreateRouteResponse instantiates a new CreateRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteResponseWithDefaults

`func NewCreateRouteResponseWithDefaults() *CreateRouteResponse`

NewCreateRouteResponseWithDefaults instantiates a new CreateRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateRouteResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateRouteResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateRouteResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateRouteResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetRouteTable

`func (o *CreateRouteResponse) GetRouteTable() RouteTable`

GetRouteTable returns the RouteTable field if non-nil, zero value otherwise.

### GetRouteTableOk

`func (o *CreateRouteResponse) GetRouteTableOk() (*RouteTable, bool)`

GetRouteTableOk returns a tuple with the RouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTable

`func (o *CreateRouteResponse) SetRouteTable(v RouteTable)`

SetRouteTable sets RouteTable field to given value.

### HasRouteTable

`func (o *CreateRouteResponse) HasRouteTable() bool`

HasRouteTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


