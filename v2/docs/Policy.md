# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | The date and time of creation of the policy. | [optional] 
**Description** | Pointer to **string** | A friendly name for the policy (between 0 and 1000 characters). | [optional] 
**IsLinkable** | Pointer to **bool** | Indicates whether the policy can be linked to a group or an EIM user. | [optional] 
**LastModificationDate** | Pointer to **time.Time** | The date and time at which the policy was last modified. | [optional] 
**Orn** | Pointer to **string** | The OUTSCALE Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | [optional] 
**Path** | Pointer to **string** | The path to the policy. | [optional] 
**PolicyDefaultVersionId** | Pointer to **string** | The ID of the policy default version. | [optional] 
**PolicyId** | Pointer to **string** | The ID of the policy. | [optional] 
**PolicyName** | Pointer to **string** | The name of the policy. | [optional] 
**ResourcesCount** | Pointer to **int32** | The number of resources attached to the policy. | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *Policy) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Policy) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Policy) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Policy) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsLinkable

`func (o *Policy) GetIsLinkable() bool`

GetIsLinkable returns the IsLinkable field if non-nil, zero value otherwise.

### GetIsLinkableOk

`func (o *Policy) GetIsLinkableOk() (*bool, bool)`

GetIsLinkableOk returns a tuple with the IsLinkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkable

`func (o *Policy) SetIsLinkable(v bool)`

SetIsLinkable sets IsLinkable field to given value.

### HasIsLinkable

`func (o *Policy) HasIsLinkable() bool`

HasIsLinkable returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *Policy) GetLastModificationDate() time.Time`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *Policy) GetLastModificationDateOk() (*time.Time, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *Policy) SetLastModificationDate(v time.Time)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *Policy) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetOrn

`func (o *Policy) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *Policy) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *Policy) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *Policy) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetPath

`func (o *Policy) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Policy) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Policy) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Policy) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPolicyDefaultVersionId

`func (o *Policy) GetPolicyDefaultVersionId() string`

GetPolicyDefaultVersionId returns the PolicyDefaultVersionId field if non-nil, zero value otherwise.

### GetPolicyDefaultVersionIdOk

`func (o *Policy) GetPolicyDefaultVersionIdOk() (*string, bool)`

GetPolicyDefaultVersionIdOk returns a tuple with the PolicyDefaultVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDefaultVersionId

`func (o *Policy) SetPolicyDefaultVersionId(v string)`

SetPolicyDefaultVersionId sets PolicyDefaultVersionId field to given value.

### HasPolicyDefaultVersionId

`func (o *Policy) HasPolicyDefaultVersionId() bool`

HasPolicyDefaultVersionId returns a boolean if a field has been set.

### GetPolicyId

`func (o *Policy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Policy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Policy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Policy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *Policy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *Policy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *Policy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *Policy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetResourcesCount

`func (o *Policy) GetResourcesCount() int32`

GetResourcesCount returns the ResourcesCount field if non-nil, zero value otherwise.

### GetResourcesCountOk

`func (o *Policy) GetResourcesCountOk() (*int32, bool)`

GetResourcesCountOk returns a tuple with the ResourcesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesCount

`func (o *Policy) SetResourcesCount(v int32)`

SetResourcesCount sets ResourcesCount field to given value.

### HasResourcesCount

`func (o *Policy) HasResourcesCount() bool`

HasResourcesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


