# ReadFlexibleGpuCatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlexibleGpuCatalog** | Pointer to [**[]FlexibleGpuCatalog**](FlexibleGpuCatalog.md) | Information about one or more fGPUs available in the public catalog. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadFlexibleGpuCatalogResponse

`func NewReadFlexibleGpuCatalogResponse() *ReadFlexibleGpuCatalogResponse`

NewReadFlexibleGpuCatalogResponse instantiates a new ReadFlexibleGpuCatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadFlexibleGpuCatalogResponseWithDefaults

`func NewReadFlexibleGpuCatalogResponseWithDefaults() *ReadFlexibleGpuCatalogResponse`

NewReadFlexibleGpuCatalogResponseWithDefaults instantiates a new ReadFlexibleGpuCatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlexibleGpuCatalog

`func (o *ReadFlexibleGpuCatalogResponse) GetFlexibleGpuCatalog() []FlexibleGpuCatalog`

GetFlexibleGpuCatalog returns the FlexibleGpuCatalog field if non-nil, zero value otherwise.

### GetFlexibleGpuCatalogOk

`func (o *ReadFlexibleGpuCatalogResponse) GetFlexibleGpuCatalogOk() (*[]FlexibleGpuCatalog, bool)`

GetFlexibleGpuCatalogOk returns a tuple with the FlexibleGpuCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleGpuCatalog

`func (o *ReadFlexibleGpuCatalogResponse) SetFlexibleGpuCatalog(v []FlexibleGpuCatalog)`

SetFlexibleGpuCatalog sets FlexibleGpuCatalog field to given value.

### HasFlexibleGpuCatalog

`func (o *ReadFlexibleGpuCatalogResponse) HasFlexibleGpuCatalog() bool`

HasFlexibleGpuCatalog returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadFlexibleGpuCatalogResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadFlexibleGpuCatalogResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadFlexibleGpuCatalogResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadFlexibleGpuCatalogResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


