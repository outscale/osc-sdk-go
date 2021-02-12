# UpdateNetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSetId** | **string** | The ID of the DHCP options set (or &#x60;default&#x60; if you want to associate the default one). | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | **string** | The ID of the Net. | 

## Methods

### NewUpdateNetRequest

`func NewUpdateNetRequest(dhcpOptionsSetId string, netId string, ) *UpdateNetRequest`

NewUpdateNetRequest instantiates a new UpdateNetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetRequestWithDefaults

`func NewUpdateNetRequestWithDefaults() *UpdateNetRequest`

NewUpdateNetRequestWithDefaults instantiates a new UpdateNetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptionsSetId

`func (o *UpdateNetRequest) GetDhcpOptionsSetId() string`

GetDhcpOptionsSetId returns the DhcpOptionsSetId field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdOk

`func (o *UpdateNetRequest) GetDhcpOptionsSetIdOk() (*string, bool)`

GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsSetId

`func (o *UpdateNetRequest) SetDhcpOptionsSetId(v string)`

SetDhcpOptionsSetId sets DhcpOptionsSetId field to given value.


### GetDryRun

`func (o *UpdateNetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateNetRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateNetRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateNetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetId

`func (o *UpdateNetRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UpdateNetRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *UpdateNetRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


