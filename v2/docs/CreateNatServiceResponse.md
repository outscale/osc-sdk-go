# CreateNatServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatService** | Pointer to [**NatService**](NatService.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewCreateNatServiceResponse

`func NewCreateNatServiceResponse() *CreateNatServiceResponse`

NewCreateNatServiceResponse instantiates a new CreateNatServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNatServiceResponseWithDefaults

`func NewCreateNatServiceResponseWithDefaults() *CreateNatServiceResponse`

NewCreateNatServiceResponseWithDefaults instantiates a new CreateNatServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatService

`func (o *CreateNatServiceResponse) GetNatService() NatService`

GetNatService returns the NatService field if non-nil, zero value otherwise.

### GetNatServiceOk

`func (o *CreateNatServiceResponse) GetNatServiceOk() (*NatService, bool)`

GetNatServiceOk returns a tuple with the NatService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatService

`func (o *CreateNatServiceResponse) SetNatService(v NatService)`

SetNatService sets NatService field to given value.

### HasNatService

`func (o *CreateNatServiceResponse) HasNatService() bool`

HasNatService returns a boolean if a field has been set.

### GetResponseContext

`func (o *CreateNatServiceResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateNatServiceResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateNatServiceResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateNatServiceResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


