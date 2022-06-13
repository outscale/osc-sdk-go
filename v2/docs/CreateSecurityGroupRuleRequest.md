# CreateSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Flow** | **string** | The direction of the flow: &#x60;Inbound&#x60; or &#x60;Outbound&#x60;. You can specify &#x60;Outbound&#x60; for Nets only. | 
**FromPortRange** | Pointer to **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 
**IpProtocol** | Pointer to **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;, or &#x60;-1&#x60; for all protocols). By default, &#x60;-1&#x60;. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | Information about the security group rule to create. If you specify this parent parameter and its subparameters, you cannot specify the following parent parameters: &#x60;FromPortRange&#x60;, &#x60;IpProtocol&#x60;, &#x60;IpRange&#x60;, and &#x60;ToPortRange&#x60;. | [optional] 
**SecurityGroupAccountIdToLink** | Pointer to **string** | The account ID of the owner of the security group for which you want to create a rule. | [optional] 
**SecurityGroupId** | **string** | The ID of the security group for which you want to create a rule. | 
**SecurityGroupNameToLink** | Pointer to **string** | The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group. | [optional] 
**ToPortRange** | Pointer to **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP code number. If you specify this parameter, you cannot specify the &#x60;Rules&#x60; parameter and its subparameters. | [optional] 

## Methods

### NewCreateSecurityGroupRuleRequest

`func NewCreateSecurityGroupRuleRequest(flow string, securityGroupId string, ) *CreateSecurityGroupRuleRequest`

NewCreateSecurityGroupRuleRequest instantiates a new CreateSecurityGroupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleRequestWithDefaults

`func NewCreateSecurityGroupRuleRequestWithDefaults() *CreateSecurityGroupRuleRequest`

NewCreateSecurityGroupRuleRequestWithDefaults instantiates a new CreateSecurityGroupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateSecurityGroupRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSecurityGroupRuleRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateSecurityGroupRuleRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateSecurityGroupRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFlow

`func (o *CreateSecurityGroupRuleRequest) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *CreateSecurityGroupRuleRequest) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *CreateSecurityGroupRuleRequest) SetFlow(v string)`

SetFlow sets Flow field to given value.


### GetFromPortRange

`func (o *CreateSecurityGroupRuleRequest) GetFromPortRange() int32`

GetFromPortRange returns the FromPortRange field if non-nil, zero value otherwise.

### GetFromPortRangeOk

`func (o *CreateSecurityGroupRuleRequest) GetFromPortRangeOk() (*int32, bool)`

GetFromPortRangeOk returns a tuple with the FromPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPortRange

`func (o *CreateSecurityGroupRuleRequest) SetFromPortRange(v int32)`

SetFromPortRange sets FromPortRange field to given value.

### HasFromPortRange

`func (o *CreateSecurityGroupRuleRequest) HasFromPortRange() bool`

HasFromPortRange returns a boolean if a field has been set.

### GetIpProtocol

`func (o *CreateSecurityGroupRuleRequest) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *CreateSecurityGroupRuleRequest) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *CreateSecurityGroupRuleRequest) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.

### HasIpProtocol

`func (o *CreateSecurityGroupRuleRequest) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### GetIpRange

`func (o *CreateSecurityGroupRuleRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *CreateSecurityGroupRuleRequest) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *CreateSecurityGroupRuleRequest) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *CreateSecurityGroupRuleRequest) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetRules

`func (o *CreateSecurityGroupRuleRequest) GetRules() []SecurityGroupRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateSecurityGroupRuleRequest) GetRulesOk() (*[]SecurityGroupRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateSecurityGroupRuleRequest) SetRules(v []SecurityGroupRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateSecurityGroupRuleRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSecurityGroupAccountIdToLink

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupAccountIdToLink() string`

GetSecurityGroupAccountIdToLink returns the SecurityGroupAccountIdToLink field if non-nil, zero value otherwise.

### GetSecurityGroupAccountIdToLinkOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupAccountIdToLinkOk() (*string, bool)`

GetSecurityGroupAccountIdToLinkOk returns a tuple with the SecurityGroupAccountIdToLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupAccountIdToLink

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupAccountIdToLink(v string)`

SetSecurityGroupAccountIdToLink sets SecurityGroupAccountIdToLink field to given value.

### HasSecurityGroupAccountIdToLink

`func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupAccountIdToLink() bool`

HasSecurityGroupAccountIdToLink returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### GetSecurityGroupNameToLink

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupNameToLink() string`

GetSecurityGroupNameToLink returns the SecurityGroupNameToLink field if non-nil, zero value otherwise.

### GetSecurityGroupNameToLinkOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupNameToLinkOk() (*string, bool)`

GetSecurityGroupNameToLinkOk returns a tuple with the SecurityGroupNameToLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupNameToLink

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupNameToLink(v string)`

SetSecurityGroupNameToLink sets SecurityGroupNameToLink field to given value.

### HasSecurityGroupNameToLink

`func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupNameToLink() bool`

HasSecurityGroupNameToLink returns a boolean if a field has been set.

### GetToPortRange

`func (o *CreateSecurityGroupRuleRequest) GetToPortRange() int32`

GetToPortRange returns the ToPortRange field if non-nil, zero value otherwise.

### GetToPortRangeOk

`func (o *CreateSecurityGroupRuleRequest) GetToPortRangeOk() (*int32, bool)`

GetToPortRangeOk returns a tuple with the ToPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPortRange

`func (o *CreateSecurityGroupRuleRequest) SetToPortRange(v int32)`

SetToPortRange sets ToPortRange field to given value.

### HasToPortRange

`func (o *CreateSecurityGroupRuleRequest) HasToPortRange() bool`

HasToPortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


