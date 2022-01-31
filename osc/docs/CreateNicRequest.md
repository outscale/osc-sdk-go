# CreateNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description for the NIC. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PrivateIps** | [**[]PrivateIpLight**](PrivateIpLight.md) | The primary private IP for the NIC.&lt;br /&gt; This IP must be within the IP range of the Subnet that you specify with the &#x60;SubnetId&#x60; attribute.&lt;br /&gt; If you do not specify this attribute, a random private IP is selected within the IP range of the Subnet. | [optional] 
**SecurityGroupIds** | **[]string** | One or more IDs of security groups for the NIC. | [optional] 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the NIC. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


