# DeleteInternetServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**InternetServiceId** | **string** | The ID of the Internet service you want to delete. | 

## Methods

### NewDeleteInternetServiceRequest

`func NewDeleteInternetServiceRequest(internetServiceId string, ) *DeleteInternetServiceRequest`

NewDeleteInternetServiceRequest instantiates a new DeleteInternetServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteInternetServiceRequestWithDefaults

`func NewDeleteInternetServiceRequestWithDefaults() *DeleteInternetServiceRequest`

NewDeleteInternetServiceRequestWithDefaults instantiates a new DeleteInternetServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteInternetServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteInternetServiceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteInternetServiceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteInternetServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetInternetServiceId

`func (o *DeleteInternetServiceRequest) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *DeleteInternetServiceRequest) GetInternetServiceIdOk() (*string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetServiceId

`func (o *DeleteInternetServiceRequest) SetInternetServiceId(v string)`

SetInternetServiceId sets InternetServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


