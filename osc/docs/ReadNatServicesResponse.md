# ReadNatServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatServices** | Pointer to [**[]NatService**](NatService.md) | Information about one or more NAT services. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadNatServicesResponse

`func NewReadNatServicesResponse() *ReadNatServicesResponse`

NewReadNatServicesResponse instantiates a new ReadNatServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadNatServicesResponseWithDefaults

`func NewReadNatServicesResponseWithDefaults() *ReadNatServicesResponse`

NewReadNatServicesResponseWithDefaults instantiates a new ReadNatServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatServices

`func (o *ReadNatServicesResponse) GetNatServices() []NatService`

GetNatServices returns the NatServices field if non-nil, zero value otherwise.

### GetNatServicesOk

`func (o *ReadNatServicesResponse) GetNatServicesOk() (*[]NatService, bool)`

GetNatServicesOk returns a tuple with the NatServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatServices

`func (o *ReadNatServicesResponse) SetNatServices(v []NatService)`

SetNatServices sets NatServices field to given value.

### HasNatServices

`func (o *ReadNatServicesResponse) HasNatServices() bool`

HasNatServices returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadNatServicesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNatServicesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadNatServicesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadNatServicesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


