# FiltersKeypair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeypairFingerprints** | Pointer to **[]string** | The fingerprints of the keypairs. | [optional] 
**KeypairIds** | Pointer to **[]string** | The IDs of the keypairs. | [optional] 
**KeypairNames** | Pointer to **[]string** | The names of the keypairs. | [optional] 
**KeypairTypes** | Pointer to **[]string** | The types of the keypairs (&#x60;ssh-rsa&#x60;, &#x60;ssh-ed25519&#x60;, &#x60;ecdsa-sha2-nistp256&#x60;, &#x60;ecdsa-sha2-nistp384&#x60;, or &#x60;ecdsa-sha2-nistp521&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the keypairs. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the keypairs. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the keypairs, in the following format: &amp;quot;Filters&amp;quot;:{&amp;quot;Tags&amp;quot;:[&amp;quot;TAGKEY&#x3D;TAGVALUE&amp;quot;]}. | [optional] 

## Methods

### NewFiltersKeypair

`func NewFiltersKeypair() *FiltersKeypair`

NewFiltersKeypair instantiates a new FiltersKeypair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersKeypairWithDefaults

`func NewFiltersKeypairWithDefaults() *FiltersKeypair`

NewFiltersKeypairWithDefaults instantiates a new FiltersKeypair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypairFingerprints

`func (o *FiltersKeypair) GetKeypairFingerprints() []string`

GetKeypairFingerprints returns the KeypairFingerprints field if non-nil, zero value otherwise.

### GetKeypairFingerprintsOk

`func (o *FiltersKeypair) GetKeypairFingerprintsOk() (*[]string, bool)`

GetKeypairFingerprintsOk returns a tuple with the KeypairFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairFingerprints

`func (o *FiltersKeypair) SetKeypairFingerprints(v []string)`

SetKeypairFingerprints sets KeypairFingerprints field to given value.

### HasKeypairFingerprints

`func (o *FiltersKeypair) HasKeypairFingerprints() bool`

HasKeypairFingerprints returns a boolean if a field has been set.

### GetKeypairIds

`func (o *FiltersKeypair) GetKeypairIds() []string`

GetKeypairIds returns the KeypairIds field if non-nil, zero value otherwise.

### GetKeypairIdsOk

`func (o *FiltersKeypair) GetKeypairIdsOk() (*[]string, bool)`

GetKeypairIdsOk returns a tuple with the KeypairIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairIds

`func (o *FiltersKeypair) SetKeypairIds(v []string)`

SetKeypairIds sets KeypairIds field to given value.

### HasKeypairIds

`func (o *FiltersKeypair) HasKeypairIds() bool`

HasKeypairIds returns a boolean if a field has been set.

### GetKeypairNames

`func (o *FiltersKeypair) GetKeypairNames() []string`

GetKeypairNames returns the KeypairNames field if non-nil, zero value otherwise.

### GetKeypairNamesOk

`func (o *FiltersKeypair) GetKeypairNamesOk() (*[]string, bool)`

GetKeypairNamesOk returns a tuple with the KeypairNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairNames

`func (o *FiltersKeypair) SetKeypairNames(v []string)`

SetKeypairNames sets KeypairNames field to given value.

### HasKeypairNames

`func (o *FiltersKeypair) HasKeypairNames() bool`

HasKeypairNames returns a boolean if a field has been set.

### GetKeypairTypes

`func (o *FiltersKeypair) GetKeypairTypes() []string`

GetKeypairTypes returns the KeypairTypes field if non-nil, zero value otherwise.

### GetKeypairTypesOk

`func (o *FiltersKeypair) GetKeypairTypesOk() (*[]string, bool)`

GetKeypairTypesOk returns a tuple with the KeypairTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairTypes

`func (o *FiltersKeypair) SetKeypairTypes(v []string)`

SetKeypairTypes sets KeypairTypes field to given value.

### HasKeypairTypes

`func (o *FiltersKeypair) HasKeypairTypes() bool`

HasKeypairTypes returns a boolean if a field has been set.

### GetTagKeys

`func (o *FiltersKeypair) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersKeypair) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *FiltersKeypair) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *FiltersKeypair) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagValues

`func (o *FiltersKeypair) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersKeypair) GetTagValuesOk() (*[]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValues

`func (o *FiltersKeypair) SetTagValues(v []string)`

SetTagValues sets TagValues field to given value.

### HasTagValues

`func (o *FiltersKeypair) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### GetTags

`func (o *FiltersKeypair) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersKeypair) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FiltersKeypair) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FiltersKeypair) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


