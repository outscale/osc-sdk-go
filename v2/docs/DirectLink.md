# DirectLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the DirectLink. | [optional] 
**Bandwidth** | Pointer to **string** | The physical link bandwidth (either 1 Gbps or 10 Gbps). | [optional] 
**DirectLinkId** | Pointer to **string** | The ID of the DirectLink (for example, dxcon-xxxxxxxx). | [optional] 
**DirectLinkName** | Pointer to **string** | The name of the DirectLink. | [optional] 
**Location** | Pointer to **string** | The datacenter where the DirectLink is located. | [optional] 
**RegionName** | Pointer to **string** | The Region in which the DirectLink has been created. | [optional] 
**State** | Pointer to **string** | The state of the DirectLink.&lt;br /&gt; * &#x60;requested&#x60;: The DirectLink is requested but the request has not been validated yet.&lt;br /&gt; * &#x60;pending&#x60;: The DirectLink request has been validated. It remains in the &#x60;pending&#x60; state until you establish the physical link.&lt;br /&gt; * &#x60;available&#x60;: The physical link is established and the connection is ready to use.&lt;br /&gt; * &#x60;deleting&#x60;: The deletion process is in progress.&lt;br /&gt; * &#x60;deleted&#x60;: The DirectLink is deleted. | [optional] 

## Methods

### NewDirectLink

`func NewDirectLink() *DirectLink`

NewDirectLink instantiates a new DirectLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectLinkWithDefaults

`func NewDirectLinkWithDefaults() *DirectLink`

NewDirectLinkWithDefaults instantiates a new DirectLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DirectLink) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DirectLink) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DirectLink) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DirectLink) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBandwidth

`func (o *DirectLink) GetBandwidth() string`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *DirectLink) GetBandwidthOk() (*string, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *DirectLink) SetBandwidth(v string)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *DirectLink) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetDirectLinkId

`func (o *DirectLink) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *DirectLink) GetDirectLinkIdOk() (*string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkId

`func (o *DirectLink) SetDirectLinkId(v string)`

SetDirectLinkId sets DirectLinkId field to given value.

### HasDirectLinkId

`func (o *DirectLink) HasDirectLinkId() bool`

HasDirectLinkId returns a boolean if a field has been set.

### GetDirectLinkName

`func (o *DirectLink) GetDirectLinkName() string`

GetDirectLinkName returns the DirectLinkName field if non-nil, zero value otherwise.

### GetDirectLinkNameOk

`func (o *DirectLink) GetDirectLinkNameOk() (*string, bool)`

GetDirectLinkNameOk returns a tuple with the DirectLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkName

`func (o *DirectLink) SetDirectLinkName(v string)`

SetDirectLinkName sets DirectLinkName field to given value.

### HasDirectLinkName

`func (o *DirectLink) HasDirectLinkName() bool`

HasDirectLinkName returns a boolean if a field has been set.

### GetLocation

`func (o *DirectLink) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DirectLink) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DirectLink) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DirectLink) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRegionName

`func (o *DirectLink) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DirectLink) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DirectLink) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *DirectLink) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetState

`func (o *DirectLink) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectLink) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectLink) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectLink) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


