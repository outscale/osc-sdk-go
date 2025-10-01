# VolumeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to [**VolumeUpdateParameters**](VolumeUpdateParameters.md) |  | [optional] 
**Target** | Pointer to [**VolumeUpdateParameters**](VolumeUpdateParameters.md) |  | [optional] 

## Methods

### NewVolumeUpdate

`func NewVolumeUpdate() *VolumeUpdate`

NewVolumeUpdate instantiates a new VolumeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateWithDefaults

`func NewVolumeUpdateWithDefaults() *VolumeUpdate`

NewVolumeUpdateWithDefaults instantiates a new VolumeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *VolumeUpdate) GetOrigin() VolumeUpdateParameters`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *VolumeUpdate) GetOriginOk() (*VolumeUpdateParameters, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *VolumeUpdate) SetOrigin(v VolumeUpdateParameters)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *VolumeUpdate) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTarget

`func (o *VolumeUpdate) GetTarget() VolumeUpdateParameters`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VolumeUpdate) GetTargetOk() (*VolumeUpdateParameters, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VolumeUpdate) SetTarget(v VolumeUpdateParameters)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VolumeUpdate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


