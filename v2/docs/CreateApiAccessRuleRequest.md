# CreateApiAccessRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaIds** | Pointer to **[]string** |  One or more IDs of Client Certificate Authorities (CAs). | [optional] 
**Cns** | Pointer to **[]string** | One or more Client Certificate Common Names (CNs). If this parameter is specified, you must also specify the &#x60;CaIds&#x60; parameter. | [optional] 
**Description** | Pointer to **string** | A description for the API access rule. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**IpRanges** | Pointer to **[]string** | One or more IP ranges, in CIDR notation (for example, 192.0.2.0/16). | [optional] 

## Methods

### NewCreateApiAccessRuleRequest

`func NewCreateApiAccessRuleRequest() *CreateApiAccessRuleRequest`

NewCreateApiAccessRuleRequest instantiates a new CreateApiAccessRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiAccessRuleRequestWithDefaults

`func NewCreateApiAccessRuleRequestWithDefaults() *CreateApiAccessRuleRequest`

NewCreateApiAccessRuleRequestWithDefaults instantiates a new CreateApiAccessRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaIds

`func (o *CreateApiAccessRuleRequest) GetCaIds() []string`

GetCaIds returns the CaIds field if non-nil, zero value otherwise.

### GetCaIdsOk

`func (o *CreateApiAccessRuleRequest) GetCaIdsOk() (*[]string, bool)`

GetCaIdsOk returns a tuple with the CaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaIds

`func (o *CreateApiAccessRuleRequest) SetCaIds(v []string)`

SetCaIds sets CaIds field to given value.

### HasCaIds

`func (o *CreateApiAccessRuleRequest) HasCaIds() bool`

HasCaIds returns a boolean if a field has been set.

### GetCns

`func (o *CreateApiAccessRuleRequest) GetCns() []string`

GetCns returns the Cns field if non-nil, zero value otherwise.

### GetCnsOk

`func (o *CreateApiAccessRuleRequest) GetCnsOk() (*[]string, bool)`

GetCnsOk returns a tuple with the Cns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCns

`func (o *CreateApiAccessRuleRequest) SetCns(v []string)`

SetCns sets Cns field to given value.

### HasCns

`func (o *CreateApiAccessRuleRequest) HasCns() bool`

HasCns returns a boolean if a field has been set.

### GetDescription

`func (o *CreateApiAccessRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApiAccessRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApiAccessRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApiAccessRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateApiAccessRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateApiAccessRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateApiAccessRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateApiAccessRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetIpRanges

`func (o *CreateApiAccessRuleRequest) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *CreateApiAccessRuleRequest) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *CreateApiAccessRuleRequest) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *CreateApiAccessRuleRequest) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


