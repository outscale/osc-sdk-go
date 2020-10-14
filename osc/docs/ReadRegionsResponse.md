# ReadRegionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**[]Region**](Region.md) | Information about one or more Regions. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadRegionsResponse

`func NewReadRegionsResponse() *ReadRegionsResponse`

NewReadRegionsResponse instantiates a new ReadRegionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRegionsResponseWithDefaults

`func NewReadRegionsResponseWithDefaults() *ReadRegionsResponse`

NewReadRegionsResponseWithDefaults instantiates a new ReadRegionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *ReadRegionsResponse) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ReadRegionsResponse) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ReadRegionsResponse) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ReadRegionsResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadRegionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadRegionsResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadRegionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadRegionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


