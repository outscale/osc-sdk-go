# MinimalPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the entity. | [optional] 
**Name** | Pointer to **string** | The name of the entity. | [optional] 
**Orn** | Pointer to **string** | The OUTSCALE Resource Name (ORN) of the entity. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | [optional] 

## Methods

### NewMinimalPolicy

`func NewMinimalPolicy() *MinimalPolicy`

NewMinimalPolicy instantiates a new MinimalPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalPolicyWithDefaults

`func NewMinimalPolicyWithDefaults() *MinimalPolicy`

NewMinimalPolicyWithDefaults instantiates a new MinimalPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MinimalPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MinimalPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MinimalPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrn

`func (o *MinimalPolicy) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *MinimalPolicy) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *MinimalPolicy) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *MinimalPolicy) HasOrn() bool`

HasOrn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


