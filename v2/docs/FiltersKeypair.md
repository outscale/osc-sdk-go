# FiltersKeypair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeypairFingerprints** | Pointer to **[]string** | The fingerprints of the keypairs. | [optional] 
**KeypairNames** | Pointer to **[]string** | The names of the keypairs. | [optional] 
**KeypairTypes** | Pointer to **[]string** | The types of the keypairs (&#x60;ssh-rsa&#x60;, &#x60;ssh-ed25519&#x60;, &#x60;ecdsa-sha2-nistp256&#x60;, &#x60;ecdsa-sha2-nistp384&#x60;, or &#x60;ecdsa-sha2-nistp521&#x60;). | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


