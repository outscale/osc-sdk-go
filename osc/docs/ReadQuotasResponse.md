# ReadQuotasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaTypes** | Pointer to [**[]QuotaTypes**](QuotaTypes.md) | Information about one or more quotas. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadQuotasResponse

`func NewReadQuotasResponse() *ReadQuotasResponse`

NewReadQuotasResponse instantiates a new ReadQuotasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadQuotasResponseWithDefaults

`func NewReadQuotasResponseWithDefaults() *ReadQuotasResponse`

NewReadQuotasResponseWithDefaults instantiates a new ReadQuotasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaTypes

`func (o *ReadQuotasResponse) GetQuotaTypes() []QuotaTypes`

GetQuotaTypes returns the QuotaTypes field if non-nil, zero value otherwise.

### GetQuotaTypesOk

`func (o *ReadQuotasResponse) GetQuotaTypesOk() (*[]QuotaTypes, bool)`

GetQuotaTypesOk returns a tuple with the QuotaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaTypes

`func (o *ReadQuotasResponse) SetQuotaTypes(v []QuotaTypes)`

SetQuotaTypes sets QuotaTypes field to given value.

### HasQuotaTypes

`func (o *ReadQuotasResponse) HasQuotaTypes() bool`

HasQuotaTypes returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadQuotasResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadQuotasResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadQuotasResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadQuotasResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


