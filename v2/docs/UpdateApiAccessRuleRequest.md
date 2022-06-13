# UpdateApiAccessRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccessRuleId** | **string** | The ID of the API access rule you want to update. | 
**CaIds** | Pointer to **[]string** | One or more IDs of Client Certificate Authorities (CAs). | [optional] 
**Cns** | Pointer to **[]string** | One or more Client Certificate Common Names (CNs). | [optional] 
**Description** | Pointer to **string** | A new description for the API access rule. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IpRanges** | Pointer to **[]string** | One or more IP ranges, in CIDR notation (for example, &#x60;192.0.2.0/16&#x60;). | [optional] 

## Methods

### NewUpdateApiAccessRuleRequest

`func NewUpdateApiAccessRuleRequest(apiAccessRuleId string, ) *UpdateApiAccessRuleRequest`

NewUpdateApiAccessRuleRequest instantiates a new UpdateApiAccessRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiAccessRuleRequestWithDefaults

`func NewUpdateApiAccessRuleRequestWithDefaults() *UpdateApiAccessRuleRequest`

NewUpdateApiAccessRuleRequestWithDefaults instantiates a new UpdateApiAccessRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccessRuleId

`func (o *UpdateApiAccessRuleRequest) GetApiAccessRuleId() string`

GetApiAccessRuleId returns the ApiAccessRuleId field if non-nil, zero value otherwise.

### GetApiAccessRuleIdOk

`func (o *UpdateApiAccessRuleRequest) GetApiAccessRuleIdOk() (*string, bool)`

GetApiAccessRuleIdOk returns a tuple with the ApiAccessRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessRuleId

`func (o *UpdateApiAccessRuleRequest) SetApiAccessRuleId(v string)`

SetApiAccessRuleId sets ApiAccessRuleId field to given value.


### GetCaIds

`func (o *UpdateApiAccessRuleRequest) GetCaIds() []string`

GetCaIds returns the CaIds field if non-nil, zero value otherwise.

### GetCaIdsOk

`func (o *UpdateApiAccessRuleRequest) GetCaIdsOk() (*[]string, bool)`

GetCaIdsOk returns a tuple with the CaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaIds

`func (o *UpdateApiAccessRuleRequest) SetCaIds(v []string)`

SetCaIds sets CaIds field to given value.

### HasCaIds

`func (o *UpdateApiAccessRuleRequest) HasCaIds() bool`

HasCaIds returns a boolean if a field has been set.

### GetCns

`func (o *UpdateApiAccessRuleRequest) GetCns() []string`

GetCns returns the Cns field if non-nil, zero value otherwise.

### GetCnsOk

`func (o *UpdateApiAccessRuleRequest) GetCnsOk() (*[]string, bool)`

GetCnsOk returns a tuple with the Cns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCns

`func (o *UpdateApiAccessRuleRequest) SetCns(v []string)`

SetCns sets Cns field to given value.

### HasCns

`func (o *UpdateApiAccessRuleRequest) HasCns() bool`

HasCns returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateApiAccessRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApiAccessRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApiAccessRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApiAccessRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *UpdateApiAccessRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateApiAccessRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UpdateApiAccessRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UpdateApiAccessRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetIpRanges

`func (o *UpdateApiAccessRuleRequest) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *UpdateApiAccessRuleRequest) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *UpdateApiAccessRuleRequest) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *UpdateApiAccessRuleRequest) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


