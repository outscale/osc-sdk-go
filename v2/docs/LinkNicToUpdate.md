# LinkNicToUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If true, the NIC is deleted when the VM is terminated. | [optional] 
**LinkNicId** | Pointer to **string** | The ID of the NIC attachment. | [optional] 

## Methods

### NewLinkNicToUpdate

`func NewLinkNicToUpdate() *LinkNicToUpdate`

NewLinkNicToUpdate instantiates a new LinkNicToUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkNicToUpdateWithDefaults

`func NewLinkNicToUpdateWithDefaults() *LinkNicToUpdate`

NewLinkNicToUpdateWithDefaults instantiates a new LinkNicToUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnVmDeletion

`func (o *LinkNicToUpdate) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *LinkNicToUpdate) GetDeleteOnVmDeletionOk() (*bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnVmDeletion

`func (o *LinkNicToUpdate) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion sets DeleteOnVmDeletion field to given value.

### HasDeleteOnVmDeletion

`func (o *LinkNicToUpdate) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### GetLinkNicId

`func (o *LinkNicToUpdate) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *LinkNicToUpdate) GetLinkNicIdOk() (*string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNicId

`func (o *LinkNicToUpdate) SetLinkNicId(v string)`

SetLinkNicId sets LinkNicId field to given value.

### HasLinkNicId

`func (o *LinkNicToUpdate) HasLinkNicId() bool`

HasLinkNicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


