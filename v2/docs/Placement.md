# Placement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubregionName** | Pointer to **string** | The name of the Subregion. If you specify this parameter, you must not specify the &#x60;Nics&#x60; parameter. | [optional] 
**Tenancy** | Pointer to **string** | The tenancy of the VM (&#x60;default&#x60; \\| &#x60;dedicated&#x60;). | [optional] 

## Methods

### NewPlacement

`func NewPlacement() *Placement`

NewPlacement instantiates a new Placement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementWithDefaults

`func NewPlacementWithDefaults() *Placement`

NewPlacementWithDefaults instantiates a new Placement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubregionName

`func (o *Placement) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *Placement) GetSubregionNameOk() (*string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregionName

`func (o *Placement) SetSubregionName(v string)`

SetSubregionName sets SubregionName field to given value.

### HasSubregionName

`func (o *Placement) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### GetTenancy

`func (o *Placement) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *Placement) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *Placement) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *Placement) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


