# CreateNetAccessPointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAccessPoint** | Pointer to [**NetAccessPoint**](NetAccessPoint.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateNetAccessPointResponse

`func NewCreateNetAccessPointResponse() *CreateNetAccessPointResponse`

NewCreateNetAccessPointResponse instantiates a new CreateNetAccessPointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetAccessPointResponseWithDefaults

`func NewCreateNetAccessPointResponseWithDefaults() *CreateNetAccessPointResponse`

NewCreateNetAccessPointResponseWithDefaults instantiates a new CreateNetAccessPointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAccessPoint

`func (o *CreateNetAccessPointResponse) GetNetAccessPoint() NetAccessPoint`

GetNetAccessPoint returns the NetAccessPoint field if non-nil, zero value otherwise.

### GetNetAccessPointOk

`func (o *CreateNetAccessPointResponse) GetNetAccessPointOk() (*NetAccessPoint, bool)`

GetNetAccessPointOk returns a tuple with the NetAccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPoint

`func (o *CreateNetAccessPointResponse) SetNetAccessPoint(v NetAccessPoint)`

SetNetAccessPoint sets NetAccessPoint field to given value.

### HasNetAccessPoint

`func (o *CreateNetAccessPointResponse) HasNetAccessPoint() bool`

HasNetAccessPoint returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateNetAccessPointResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateNetAccessPointResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateNetAccessPointResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateNetAccessPointResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


