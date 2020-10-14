# ReadRouteTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**RouteTables** | Pointer to [**[]RouteTable**](RouteTable.md) | Information about one or more route tables. | [optional] 

## Methods

### NewReadRouteTablesResponse

`func NewReadRouteTablesResponse() *ReadRouteTablesResponse`

NewReadRouteTablesResponse instantiates a new ReadRouteTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRouteTablesResponseWithDefaults

`func NewReadRouteTablesResponseWithDefaults() *ReadRouteTablesResponse`

NewReadRouteTablesResponseWithDefaults instantiates a new ReadRouteTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadRouteTablesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadRouteTablesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadRouteTablesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadRouteTablesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetRouteTables

`func (o *ReadRouteTablesResponse) GetRouteTables() []RouteTable`

GetRouteTables returns the RouteTables field if non-nil, zero value otherwise.

### GetRouteTablesOk

`func (o *ReadRouteTablesResponse) GetRouteTablesOk() (*[]RouteTable, bool)`

GetRouteTablesOk returns a tuple with the RouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTables

`func (o *ReadRouteTablesResponse) SetRouteTables(v []RouteTable)`

SetRouteTables sets RouteTables field to given value.

### HasRouteTables

`func (o *ReadRouteTablesResponse) HasRouteTables() bool`

HasRouteTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


