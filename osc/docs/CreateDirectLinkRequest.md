# CreateDirectLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | **string** | The bandwidth of the DirectLink (&#x60;1Gbps&#x60; \\| &#x60;10Gbps&#x60;). | 
**DirectLinkName** | **string** | The name of the DirectLink. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Location** | **string** | The code of the requested location for the DirectLink, returned by the [ReadLocations](#readlocations) method. | 

## Methods

### NewCreateDirectLinkRequest

`func NewCreateDirectLinkRequest(bandwidth string, directLinkName string, location string, ) *CreateDirectLinkRequest`

NewCreateDirectLinkRequest instantiates a new CreateDirectLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDirectLinkRequestWithDefaults

`func NewCreateDirectLinkRequestWithDefaults() *CreateDirectLinkRequest`

NewCreateDirectLinkRequestWithDefaults instantiates a new CreateDirectLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *CreateDirectLinkRequest) GetBandwidth() string`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CreateDirectLinkRequest) GetBandwidthOk() (*string, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CreateDirectLinkRequest) SetBandwidth(v string)`

SetBandwidth sets Bandwidth field to given value.


### GetDirectLinkName

`func (o *CreateDirectLinkRequest) GetDirectLinkName() string`

GetDirectLinkName returns the DirectLinkName field if non-nil, zero value otherwise.

### GetDirectLinkNameOk

`func (o *CreateDirectLinkRequest) GetDirectLinkNameOk() (*string, bool)`

GetDirectLinkNameOk returns a tuple with the DirectLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkName

`func (o *CreateDirectLinkRequest) SetDirectLinkName(v string)`

SetDirectLinkName sets DirectLinkName field to given value.


### GetDryRun

`func (o *CreateDirectLinkRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDirectLinkRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateDirectLinkRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateDirectLinkRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLocation

`func (o *CreateDirectLinkRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateDirectLinkRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateDirectLinkRequest) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


