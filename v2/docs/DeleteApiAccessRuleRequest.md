# DeleteApiAccessRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccessRuleId** | **string** | The ID of the API access rule you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewDeleteApiAccessRuleRequest

`func NewDeleteApiAccessRuleRequest(apiAccessRuleId string, ) *DeleteApiAccessRuleRequest`

NewDeleteApiAccessRuleRequest instantiates a new DeleteApiAccessRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteApiAccessRuleRequestWithDefaults

`func NewDeleteApiAccessRuleRequestWithDefaults() *DeleteApiAccessRuleRequest`

NewDeleteApiAccessRuleRequestWithDefaults instantiates a new DeleteApiAccessRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccessRuleId

`func (o *DeleteApiAccessRuleRequest) GetApiAccessRuleId() string`

GetApiAccessRuleId returns the ApiAccessRuleId field if non-nil, zero value otherwise.

### GetApiAccessRuleIdOk

`func (o *DeleteApiAccessRuleRequest) GetApiAccessRuleIdOk() (*string, bool)`

GetApiAccessRuleIdOk returns a tuple with the ApiAccessRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessRuleId

`func (o *DeleteApiAccessRuleRequest) SetApiAccessRuleId(v string)`

SetApiAccessRuleId sets ApiAccessRuleId field to given value.


### GetDryRun

`func (o *DeleteApiAccessRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteApiAccessRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteApiAccessRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteApiAccessRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


