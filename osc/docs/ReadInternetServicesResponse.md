# ReadInternetServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetServices** | Pointer to [**[]InternetService**](InternetService.md) | Information about one or more Internet services. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewReadInternetServicesResponse

`func NewReadInternetServicesResponse() *ReadInternetServicesResponse`

NewReadInternetServicesResponse instantiates a new ReadInternetServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadInternetServicesResponseWithDefaults

`func NewReadInternetServicesResponseWithDefaults() *ReadInternetServicesResponse`

NewReadInternetServicesResponseWithDefaults instantiates a new ReadInternetServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternetServices

`func (o *ReadInternetServicesResponse) GetInternetServices() []InternetService`

GetInternetServices returns the InternetServices field if non-nil, zero value otherwise.

### GetInternetServicesOk

`func (o *ReadInternetServicesResponse) GetInternetServicesOk() (*[]InternetService, bool)`

GetInternetServicesOk returns a tuple with the InternetServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetServices

`func (o *ReadInternetServicesResponse) SetInternetServices(v []InternetService)`

SetInternetServices sets InternetServices field to given value.

### HasInternetServices

`func (o *ReadInternetServicesResponse) HasInternetServices() bool`

HasInternetServices returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadInternetServicesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadInternetServicesResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadInternetServicesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadInternetServicesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


