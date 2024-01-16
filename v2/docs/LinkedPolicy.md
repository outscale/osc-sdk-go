# LinkedPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | The date and time of creation of the attached policy. | [optional] 
**LastModificationDate** | Pointer to **time.Time** | The date and time at which the attached policy was last modified. | [optional] 
**Orn** | Pointer to **string** | The Outscale Resource Name (ORN) of the policy. For more information, see [Resource Identifiers](https://docs.outscale.com/en/userguide/Resource-Identifiers.html). | [optional] 
**PolicyId** | Pointer to **string** | The ID of the policy. | [optional] 
**PolicyName** | Pointer to **string** | The name of the policy. | [optional] 

## Methods

### NewLinkedPolicy

`func NewLinkedPolicy() *LinkedPolicy`

NewLinkedPolicy instantiates a new LinkedPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedPolicyWithDefaults

`func NewLinkedPolicyWithDefaults() *LinkedPolicy`

NewLinkedPolicyWithDefaults instantiates a new LinkedPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *LinkedPolicy) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *LinkedPolicy) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *LinkedPolicy) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *LinkedPolicy) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModificationDate

`func (o *LinkedPolicy) GetLastModificationDate() time.Time`

GetLastModificationDate returns the LastModificationDate field if non-nil, zero value otherwise.

### GetLastModificationDateOk

`func (o *LinkedPolicy) GetLastModificationDateOk() (*time.Time, bool)`

GetLastModificationDateOk returns a tuple with the LastModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDate

`func (o *LinkedPolicy) SetLastModificationDate(v time.Time)`

SetLastModificationDate sets LastModificationDate field to given value.

### HasLastModificationDate

`func (o *LinkedPolicy) HasLastModificationDate() bool`

HasLastModificationDate returns a boolean if a field has been set.

### GetOrn

`func (o *LinkedPolicy) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *LinkedPolicy) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *LinkedPolicy) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *LinkedPolicy) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetPolicyId

`func (o *LinkedPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *LinkedPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *LinkedPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *LinkedPolicy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *LinkedPolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *LinkedPolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *LinkedPolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *LinkedPolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


