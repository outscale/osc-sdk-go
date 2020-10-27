# UnlinkInternetServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**InternetServiceId** | **string** | The ID of the Internet service you want to detach. | 
**NetId** | **string** | The ID of the Net from which you want to detach the Internet service. | 

## Methods

### NewUnlinkInternetServiceRequest

`func NewUnlinkInternetServiceRequest(internetServiceId string, netId string, ) *UnlinkInternetServiceRequest`

NewUnlinkInternetServiceRequest instantiates a new UnlinkInternetServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkInternetServiceRequestWithDefaults

`func NewUnlinkInternetServiceRequestWithDefaults() *UnlinkInternetServiceRequest`

NewUnlinkInternetServiceRequestWithDefaults instantiates a new UnlinkInternetServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkInternetServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkInternetServiceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkInternetServiceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkInternetServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetInternetServiceId

`func (o *UnlinkInternetServiceRequest) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *UnlinkInternetServiceRequest) GetInternetServiceIdOk() (*string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetServiceId

`func (o *UnlinkInternetServiceRequest) SetInternetServiceId(v string)`

SetInternetServiceId sets InternetServiceId field to given value.


### GetNetId

`func (o *UnlinkInternetServiceRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UnlinkInternetServiceRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *UnlinkInternetServiceRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


