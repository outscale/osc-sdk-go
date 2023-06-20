# VgwTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedRouteCount** | Pointer to **int32** | The number of routes accepted through BGP (Border Gateway Protocol) route exchanges. | [optional] 
**LastStateChangeDate** | Pointer to **time.Time** | The date and time (UTC) of the latest state update. | [optional] 
**OutsideIpAddress** | Pointer to **string** | The IP on the OUTSCALE side of the tunnel. | [optional] 
**State** | Pointer to **string** | The state of the IPSEC tunnel (&#x60;UP&#x60; \\| &#x60;DOWN&#x60;). | [optional] 
**StateDescription** | Pointer to **string** | A description of the current state of the tunnel. | [optional] 

## Methods

### NewVgwTelemetry

`func NewVgwTelemetry() *VgwTelemetry`

NewVgwTelemetry instantiates a new VgwTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVgwTelemetryWithDefaults

`func NewVgwTelemetryWithDefaults() *VgwTelemetry`

NewVgwTelemetryWithDefaults instantiates a new VgwTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedRouteCount

`func (o *VgwTelemetry) GetAcceptedRouteCount() int32`

GetAcceptedRouteCount returns the AcceptedRouteCount field if non-nil, zero value otherwise.

### GetAcceptedRouteCountOk

`func (o *VgwTelemetry) GetAcceptedRouteCountOk() (*int32, bool)`

GetAcceptedRouteCountOk returns a tuple with the AcceptedRouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedRouteCount

`func (o *VgwTelemetry) SetAcceptedRouteCount(v int32)`

SetAcceptedRouteCount sets AcceptedRouteCount field to given value.

### HasAcceptedRouteCount

`func (o *VgwTelemetry) HasAcceptedRouteCount() bool`

HasAcceptedRouteCount returns a boolean if a field has been set.

### GetLastStateChangeDate

`func (o *VgwTelemetry) GetLastStateChangeDate() time.Time`

GetLastStateChangeDate returns the LastStateChangeDate field if non-nil, zero value otherwise.

### GetLastStateChangeDateOk

`func (o *VgwTelemetry) GetLastStateChangeDateOk() (*time.Time, bool)`

GetLastStateChangeDateOk returns a tuple with the LastStateChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeDate

`func (o *VgwTelemetry) SetLastStateChangeDate(v time.Time)`

SetLastStateChangeDate sets LastStateChangeDate field to given value.

### HasLastStateChangeDate

`func (o *VgwTelemetry) HasLastStateChangeDate() bool`

HasLastStateChangeDate returns a boolean if a field has been set.

### GetOutsideIpAddress

`func (o *VgwTelemetry) GetOutsideIpAddress() string`

GetOutsideIpAddress returns the OutsideIpAddress field if non-nil, zero value otherwise.

### GetOutsideIpAddressOk

`func (o *VgwTelemetry) GetOutsideIpAddressOk() (*string, bool)`

GetOutsideIpAddressOk returns a tuple with the OutsideIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIpAddress

`func (o *VgwTelemetry) SetOutsideIpAddress(v string)`

SetOutsideIpAddress sets OutsideIpAddress field to given value.

### HasOutsideIpAddress

`func (o *VgwTelemetry) HasOutsideIpAddress() bool`

HasOutsideIpAddress returns a boolean if a field has been set.

### GetState

`func (o *VgwTelemetry) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VgwTelemetry) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VgwTelemetry) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VgwTelemetry) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDescription

`func (o *VgwTelemetry) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *VgwTelemetry) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *VgwTelemetry) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *VgwTelemetry) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


