# FiltersApiAccessRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccessRuleIds** | Pointer to **[]string** | One or more IDs of API access rules. | [optional] 
**CaIds** | Pointer to **[]string** | One or more IDs of Client Certificate Authorities (CAs). | [optional] 
**Cns** | Pointer to **[]string** | One or more Client Certificate Common Names (CNs). | [optional] 
**Descriptions** | Pointer to **[]string** | One or more descriptions of API access rules. | [optional] 
**IpRanges** | Pointer to **[]string** | One or more IP ranges, in CIDR notation (for example, 192.0.2.0/16). | [optional] 

## Methods

### NewFiltersApiAccessRule

`func NewFiltersApiAccessRule() *FiltersApiAccessRule`

NewFiltersApiAccessRule instantiates a new FiltersApiAccessRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersApiAccessRuleWithDefaults

`func NewFiltersApiAccessRuleWithDefaults() *FiltersApiAccessRule`

NewFiltersApiAccessRuleWithDefaults instantiates a new FiltersApiAccessRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccessRuleIds

`func (o *FiltersApiAccessRule) GetApiAccessRuleIds() []string`

GetApiAccessRuleIds returns the ApiAccessRuleIds field if non-nil, zero value otherwise.

### GetApiAccessRuleIdsOk

`func (o *FiltersApiAccessRule) GetApiAccessRuleIdsOk() (*[]string, bool)`

GetApiAccessRuleIdsOk returns a tuple with the ApiAccessRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessRuleIds

`func (o *FiltersApiAccessRule) SetApiAccessRuleIds(v []string)`

SetApiAccessRuleIds sets ApiAccessRuleIds field to given value.

### HasApiAccessRuleIds

`func (o *FiltersApiAccessRule) HasApiAccessRuleIds() bool`

HasApiAccessRuleIds returns a boolean if a field has been set.

### GetCaIds

`func (o *FiltersApiAccessRule) GetCaIds() []string`

GetCaIds returns the CaIds field if non-nil, zero value otherwise.

### GetCaIdsOk

`func (o *FiltersApiAccessRule) GetCaIdsOk() (*[]string, bool)`

GetCaIdsOk returns a tuple with the CaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaIds

`func (o *FiltersApiAccessRule) SetCaIds(v []string)`

SetCaIds sets CaIds field to given value.

### HasCaIds

`func (o *FiltersApiAccessRule) HasCaIds() bool`

HasCaIds returns a boolean if a field has been set.

### GetCns

`func (o *FiltersApiAccessRule) GetCns() []string`

GetCns returns the Cns field if non-nil, zero value otherwise.

### GetCnsOk

`func (o *FiltersApiAccessRule) GetCnsOk() (*[]string, bool)`

GetCnsOk returns a tuple with the Cns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCns

`func (o *FiltersApiAccessRule) SetCns(v []string)`

SetCns sets Cns field to given value.

### HasCns

`func (o *FiltersApiAccessRule) HasCns() bool`

HasCns returns a boolean if a field has been set.

### GetDescriptions

`func (o *FiltersApiAccessRule) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersApiAccessRule) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *FiltersApiAccessRule) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *FiltersApiAccessRule) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetIpRanges

`func (o *FiltersApiAccessRule) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FiltersApiAccessRule) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *FiltersApiAccessRule) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *FiltersApiAccessRule) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


