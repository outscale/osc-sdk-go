# UpdateDirectLinkInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkInterfaceId** | **string** | The ID of the DirectLink interface you want to update. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Mtu** | **int32** | The maximum transmission unit (MTU) of the DirectLink interface, in bytes (between &#x60;1500&#x60; and &#x60;9001&#x60;, both included). | [default to 1500]

## Methods

### NewUpdateDirectLinkInterfaceRequest

`func NewUpdateDirectLinkInterfaceRequest(directLinkInterfaceId string, mtu int32, ) *UpdateDirectLinkInterfaceRequest`

NewUpdateDirectLinkInterfaceRequest instantiates a new UpdateDirectLinkInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDirectLinkInterfaceRequestWithDefaults

`func NewUpdateDirectLinkInterfaceRequestWithDefaults() *UpdateDirectLinkInterfaceRequest`

NewUpdateDirectLinkInterfaceRequestWithDefaults instantiates a new UpdateDirectLinkInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLinkInterfaceId

`func (o *UpdateDirectLinkInterfaceRequest) GetDirectLinkInterfaceId() string`

GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceIdOk

`func (o *UpdateDirectLinkInterfaceRequest) GetDirectLinkInterfaceIdOk() (*string, bool)`

GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterfaceId

`func (o *UpdateDirectLinkInterfaceRequest) SetDirectLinkInterfaceId(v string)`

SetDirectLinkInterfaceId sets DirectLinkInterfaceId field to given value.


### GetDryRun

`func (o *UpdateDirectLinkInterfaceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateDirectLinkInterfaceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateDirectLinkInterfaceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateDirectLinkInterfaceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMtu

`func (o *UpdateDirectLinkInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *UpdateDirectLinkInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *UpdateDirectLinkInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


