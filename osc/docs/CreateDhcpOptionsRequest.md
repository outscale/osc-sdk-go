# CreateDhcpOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | **string** | Specify a domain name (for example, MyCompany.com). You can specify only one domain name. You must specify at least one of the following parameters: &#x60;DomainName&#x60;, &#x60;DomainNameServers&#x60; or &#x60;NtpServers&#x60;. | [optional] 
**DomainNameServers** | **[]string** | The IPs of domain name servers. If no IPs are specified, the &#x60;OutscaleProvidedDNS&#x60; value is set by default. You must specify at least one of the following parameters: &#x60;DomainName&#x60;, &#x60;DomainNameServers&#x60; or &#x60;NtpServers&#x60;. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NtpServers** | **[]string** | The IPs of the Network Time Protocol (NTP) servers. You must specify at least one of the following parameters: &#x60;DomainName&#x60;, &#x60;DomainNameServers&#x60; or &#x60;NtpServers&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


