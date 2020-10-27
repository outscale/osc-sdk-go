# ReadNetAccessPointServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Services** | Pointer to [**[]Service**](Service.md) | The names of the services you can use for Net access points. | [optional] 

## Methods

### NewReadNetAccessPointServicesResponse

`func NewReadNetAccessPointServicesResponse() *ReadNetAccessPointServicesResponse`

NewReadNetAccessPointServicesResponse instantiates a new ReadNetAccessPointServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadNetAccessPointServicesResponseWithDefaults

`func NewReadNetAccessPointServicesResponseWithDefaults() *ReadNetAccessPointServicesResponse`

NewReadNetAccessPointServicesResponseWithDefaults instantiates a new ReadNetAccessPointServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *ReadNetAccessPointServicesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetAccessPointServicesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadNetAccessPointServicesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadNetAccessPointServicesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetServices

`func (o *ReadNetAccessPointServicesResponse) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReadNetAccessPointServicesResponse) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ReadNetAccessPointServicesResponse) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *ReadNetAccessPointServicesResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


