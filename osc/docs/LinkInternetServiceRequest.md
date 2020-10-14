# LinkInternetServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**InternetServiceId** | **string** | The ID of the Internet service you want to attach. | 
**NetId** | **string** | The ID of the Net to which you want to attach the Internet service. | 

## Methods

### NewLinkInternetServiceRequest

`func NewLinkInternetServiceRequest(internetServiceId string, netId string, ) *LinkInternetServiceRequest`

NewLinkInternetServiceRequest instantiates a new LinkInternetServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkInternetServiceRequestWithDefaults

`func NewLinkInternetServiceRequestWithDefaults() *LinkInternetServiceRequest`

NewLinkInternetServiceRequestWithDefaults instantiates a new LinkInternetServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *LinkInternetServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkInternetServiceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LinkInternetServiceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LinkInternetServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetInternetServiceId

`func (o *LinkInternetServiceRequest) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *LinkInternetServiceRequest) GetInternetServiceIdOk() (*string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetServiceId

`func (o *LinkInternetServiceRequest) SetInternetServiceId(v string)`

SetInternetServiceId sets InternetServiceId field to given value.


### GetNetId

`func (o *LinkInternetServiceRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *LinkInternetServiceRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *LinkInternetServiceRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


