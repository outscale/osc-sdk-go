# DeleteNatServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NatServiceId** | **string** | The ID of the NAT service you want to delete. | 

## Methods

### NewDeleteNatServiceRequest

`func NewDeleteNatServiceRequest(natServiceId string, ) *DeleteNatServiceRequest`

NewDeleteNatServiceRequest instantiates a new DeleteNatServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteNatServiceRequestWithDefaults

`func NewDeleteNatServiceRequestWithDefaults() *DeleteNatServiceRequest`

NewDeleteNatServiceRequestWithDefaults instantiates a new DeleteNatServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteNatServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteNatServiceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteNatServiceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteNatServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNatServiceId

`func (o *DeleteNatServiceRequest) GetNatServiceId() string`

GetNatServiceId returns the NatServiceId field if non-nil, zero value otherwise.

### GetNatServiceIdOk

`func (o *DeleteNatServiceRequest) GetNatServiceIdOk() (*string, bool)`

GetNatServiceIdOk returns a tuple with the NatServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatServiceId

`func (o *DeleteNatServiceRequest) SetNatServiceId(v string)`

SetNatServiceId sets NatServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


