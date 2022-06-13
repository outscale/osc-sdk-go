# ApiAccessRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccessRuleId** | Pointer to **string** |  The ID of the API access rule. | [optional] 
**CaIds** | Pointer to **[]string** | One or more IDs of Client Certificate Authorities (CAs) used for the API access rule. | [optional] 
**Cns** | Pointer to **[]string** | One or more Client Certificate Common Names (CNs). | [optional] 
**Description** | Pointer to **string** | The description of the API access rule. | [optional] 
**IpRanges** | Pointer to **[]string** | One or more IP ranges used for the API access rule, in CIDR notation (for example, &#x60;192.0.2.0/16&#x60;). | [optional] 

## Methods

### NewApiAccessRule

`func NewApiAccessRule() *ApiAccessRule`

NewApiAccessRule instantiates a new ApiAccessRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccessRuleWithDefaults

`func NewApiAccessRuleWithDefaults() *ApiAccessRule`

NewApiAccessRuleWithDefaults instantiates a new ApiAccessRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccessRuleId

`func (o *ApiAccessRule) GetApiAccessRuleId() string`

GetApiAccessRuleId returns the ApiAccessRuleId field if non-nil, zero value otherwise.

### GetApiAccessRuleIdOk

`func (o *ApiAccessRule) GetApiAccessRuleIdOk() (*string, bool)`

GetApiAccessRuleIdOk returns a tuple with the ApiAccessRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessRuleId

`func (o *ApiAccessRule) SetApiAccessRuleId(v string)`

SetApiAccessRuleId sets ApiAccessRuleId field to given value.

### HasApiAccessRuleId

`func (o *ApiAccessRule) HasApiAccessRuleId() bool`

HasApiAccessRuleId returns a boolean if a field has been set.

### GetCaIds

`func (o *ApiAccessRule) GetCaIds() []string`

GetCaIds returns the CaIds field if non-nil, zero value otherwise.

### GetCaIdsOk

`func (o *ApiAccessRule) GetCaIdsOk() (*[]string, bool)`

GetCaIdsOk returns a tuple with the CaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaIds

`func (o *ApiAccessRule) SetCaIds(v []string)`

SetCaIds sets CaIds field to given value.

### HasCaIds

`func (o *ApiAccessRule) HasCaIds() bool`

HasCaIds returns a boolean if a field has been set.

### GetCns

`func (o *ApiAccessRule) GetCns() []string`

GetCns returns the Cns field if non-nil, zero value otherwise.

### GetCnsOk

`func (o *ApiAccessRule) GetCnsOk() (*[]string, bool)`

GetCnsOk returns a tuple with the Cns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCns

`func (o *ApiAccessRule) SetCns(v []string)`

SetCns sets Cns field to given value.

### HasCns

`func (o *ApiAccessRule) HasCns() bool`

HasCns returns a boolean if a field has been set.

### GetDescription

`func (o *ApiAccessRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAccessRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAccessRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAccessRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpRanges

`func (o *ApiAccessRule) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *ApiAccessRule) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *ApiAccessRule) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *ApiAccessRule) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


