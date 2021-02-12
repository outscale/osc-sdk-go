# DeleteSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Flow** | **string** | The direction of the flow: &#x60;Inbound&#x60; or &#x60;Outbound&#x60;. You can specify &#x60;Outbound&#x60; for Nets only. | 
**FromPortRange** | Pointer to **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | Pointer to **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;) or protocol number. By default, &#x60;-1&#x60;, which means all protocols. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | One or more rules you want to delete from the security group. | [optional] 
**SecurityGroupAccountIdToUnlink** | Pointer to **string** | The account ID of the owner of the security group you want to delete a rule from. | [optional] 
**SecurityGroupId** | **string** | The ID of the security group you want to delete a rule from. | 
**SecurityGroupNameToUnlink** | Pointer to **string** | The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group. | [optional] 
**ToPortRange** | Pointer to **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 

## Methods

### NewDeleteSecurityGroupRuleRequest

`func NewDeleteSecurityGroupRuleRequest(flow string, securityGroupId string, ) *DeleteSecurityGroupRuleRequest`

NewDeleteSecurityGroupRuleRequest instantiates a new DeleteSecurityGroupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSecurityGroupRuleRequestWithDefaults

`func NewDeleteSecurityGroupRuleRequestWithDefaults() *DeleteSecurityGroupRuleRequest`

NewDeleteSecurityGroupRuleRequestWithDefaults instantiates a new DeleteSecurityGroupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteSecurityGroupRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteSecurityGroupRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteSecurityGroupRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteSecurityGroupRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFlow

`func (o *DeleteSecurityGroupRuleRequest) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DeleteSecurityGroupRuleRequest) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DeleteSecurityGroupRuleRequest) SetFlow(v string)`

SetFlow sets Flow field to given value.


### GetFromPortRange

`func (o *DeleteSecurityGroupRuleRequest) GetFromPortRange() int32`

GetFromPortRange returns the FromPortRange field if non-nil, zero value otherwise.

### GetFromPortRangeOk

`func (o *DeleteSecurityGroupRuleRequest) GetFromPortRangeOk() (*int32, bool)`

GetFromPortRangeOk returns a tuple with the FromPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPortRange

`func (o *DeleteSecurityGroupRuleRequest) SetFromPortRange(v int32)`

SetFromPortRange sets FromPortRange field to given value.

### HasFromPortRange

`func (o *DeleteSecurityGroupRuleRequest) HasFromPortRange() bool`

HasFromPortRange returns a boolean if a field has been set.

### GetIpProtocol

`func (o *DeleteSecurityGroupRuleRequest) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *DeleteSecurityGroupRuleRequest) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *DeleteSecurityGroupRuleRequest) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.

### HasIpProtocol

`func (o *DeleteSecurityGroupRuleRequest) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### GetIpRange

`func (o *DeleteSecurityGroupRuleRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *DeleteSecurityGroupRuleRequest) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *DeleteSecurityGroupRuleRequest) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *DeleteSecurityGroupRuleRequest) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetRules

`func (o *DeleteSecurityGroupRuleRequest) GetRules() []SecurityGroupRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DeleteSecurityGroupRuleRequest) GetRulesOk() (*[]SecurityGroupRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DeleteSecurityGroupRuleRequest) SetRules(v []SecurityGroupRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DeleteSecurityGroupRuleRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSecurityGroupAccountIdToUnlink

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupAccountIdToUnlink() string`

GetSecurityGroupAccountIdToUnlink returns the SecurityGroupAccountIdToUnlink field if non-nil, zero value otherwise.

### GetSecurityGroupAccountIdToUnlinkOk

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupAccountIdToUnlinkOk() (*string, bool)`

GetSecurityGroupAccountIdToUnlinkOk returns a tuple with the SecurityGroupAccountIdToUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupAccountIdToUnlink

`func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupAccountIdToUnlink(v string)`

SetSecurityGroupAccountIdToUnlink sets SecurityGroupAccountIdToUnlink field to given value.

### HasSecurityGroupAccountIdToUnlink

`func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupAccountIdToUnlink() bool`

HasSecurityGroupAccountIdToUnlink returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### GetSecurityGroupNameToUnlink

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupNameToUnlink() string`

GetSecurityGroupNameToUnlink returns the SecurityGroupNameToUnlink field if non-nil, zero value otherwise.

### GetSecurityGroupNameToUnlinkOk

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupNameToUnlinkOk() (*string, bool)`

GetSecurityGroupNameToUnlinkOk returns a tuple with the SecurityGroupNameToUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupNameToUnlink

`func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupNameToUnlink(v string)`

SetSecurityGroupNameToUnlink sets SecurityGroupNameToUnlink field to given value.

### HasSecurityGroupNameToUnlink

`func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupNameToUnlink() bool`

HasSecurityGroupNameToUnlink returns a boolean if a field has been set.

### GetToPortRange

`func (o *DeleteSecurityGroupRuleRequest) GetToPortRange() int32`

GetToPortRange returns the ToPortRange field if non-nil, zero value otherwise.

### GetToPortRangeOk

`func (o *DeleteSecurityGroupRuleRequest) GetToPortRangeOk() (*int32, bool)`

GetToPortRangeOk returns a tuple with the ToPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPortRange

`func (o *DeleteSecurityGroupRuleRequest) SetToPortRange(v int32)`

SetToPortRange sets ToPortRange field to given value.

### HasToPortRange

`func (o *DeleteSecurityGroupRuleRequest) HasToPortRange() bool`

HasToPortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


