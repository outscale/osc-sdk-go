# PolicyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The policy document as a json string. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time of creation of the version. | [optional] 
**DefaultVersion** | Pointer to **bool** | If true, the version is the default one. | [optional] 
**VersionId** | Pointer to **string** | The ID of the version. | [optional] 

## Methods

### NewPolicyVersion

`func NewPolicyVersion() *PolicyVersion`

NewPolicyVersion instantiates a new PolicyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyVersionWithDefaults

`func NewPolicyVersionWithDefaults() *PolicyVersion`

NewPolicyVersionWithDefaults instantiates a new PolicyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PolicyVersion) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PolicyVersion) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PolicyVersion) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PolicyVersion) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreationDate

`func (o *PolicyVersion) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PolicyVersion) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PolicyVersion) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *PolicyVersion) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDefaultVersion

`func (o *PolicyVersion) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *PolicyVersion) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *PolicyVersion) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *PolicyVersion) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetVersionId

`func (o *PolicyVersion) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *PolicyVersion) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *PolicyVersion) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *PolicyVersion) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


