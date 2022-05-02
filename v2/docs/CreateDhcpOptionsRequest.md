# CreateDhcpOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | Specify a domain name (for example, MyCompany.com). You can specify only one domain name. You must specify at least one of the following parameters: &#x60;DomainName&#x60;, &#x60;DomainNameServers&#x60; or &#x60;NtpServers&#x60;. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The IPs of domain name servers. If no IPs are specified, the &#x60;OutscaleProvidedDNS&#x60; value is set by default. You must specify at least one of the following parameters: &#x60;DomainName&#x60;, &#x60;DomainNameServers&#x60; or &#x60;NtpServers&#x60;. | [optional] 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NtpServers** | Pointer to **[]string** | The IPs of the Network Time Protocol (NTP) servers. You must specify at least one of the following parameters: &#x60;DomainName&#x60;, &#x60;DomainNameServers&#x60; or &#x60;NtpServers&#x60;. | [optional] 

## Methods

### NewCreateDhcpOptionsRequest

`func NewCreateDhcpOptionsRequest() *CreateDhcpOptionsRequest`

NewCreateDhcpOptionsRequest instantiates a new CreateDhcpOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDhcpOptionsRequestWithDefaults

`func NewCreateDhcpOptionsRequestWithDefaults() *CreateDhcpOptionsRequest`

NewCreateDhcpOptionsRequestWithDefaults instantiates a new CreateDhcpOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *CreateDhcpOptionsRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateDhcpOptionsRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateDhcpOptionsRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateDhcpOptionsRequest) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *CreateDhcpOptionsRequest) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *CreateDhcpOptionsRequest) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *CreateDhcpOptionsRequest) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *CreateDhcpOptionsRequest) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetDryRun

`func (o *CreateDhcpOptionsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDhcpOptionsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateDhcpOptionsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateDhcpOptionsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNtpServers

`func (o *CreateDhcpOptionsRequest) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *CreateDhcpOptionsRequest) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *CreateDhcpOptionsRequest) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *CreateDhcpOptionsRequest) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


