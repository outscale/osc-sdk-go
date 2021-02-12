# DhcpOptionsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | If true, the DHCP options set is a default one. If false, it is not. | [optional] 
**DhcpOptionsSetId** | Pointer to **string** | The ID of the DHCP options set. | [optional] 
**DomainName** | Pointer to **string** | The domain name. | [optional] 
**DomainNameServers** | Pointer to **[]string** | One or more IP addresses for the domain name servers. | [optional] 
**NtpServers** | Pointer to **[]string** | One or more IP addresses for the NTP servers. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the DHCP options set. | [optional] 

## Methods

### NewDhcpOptionsSet

`func NewDhcpOptionsSet() *DhcpOptionsSet`

NewDhcpOptionsSet instantiates a new DhcpOptionsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpOptionsSetWithDefaults

`func NewDhcpOptionsSetWithDefaults() *DhcpOptionsSet`

NewDhcpOptionsSetWithDefaults instantiates a new DhcpOptionsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *DhcpOptionsSet) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DhcpOptionsSet) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DhcpOptionsSet) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DhcpOptionsSet) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDhcpOptionsSetId

`func (o *DhcpOptionsSet) GetDhcpOptionsSetId() string`

GetDhcpOptionsSetId returns the DhcpOptionsSetId field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdOk

`func (o *DhcpOptionsSet) GetDhcpOptionsSetIdOk() (*string, bool)`

GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsSetId

`func (o *DhcpOptionsSet) SetDhcpOptionsSetId(v string)`

SetDhcpOptionsSetId sets DhcpOptionsSetId field to given value.

### HasDhcpOptionsSetId

`func (o *DhcpOptionsSet) HasDhcpOptionsSetId() bool`

HasDhcpOptionsSetId returns a boolean if a field has been set.

### GetDomainName

`func (o *DhcpOptionsSet) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DhcpOptionsSet) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DhcpOptionsSet) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DhcpOptionsSet) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *DhcpOptionsSet) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *DhcpOptionsSet) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *DhcpOptionsSet) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *DhcpOptionsSet) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetNtpServers

`func (o *DhcpOptionsSet) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *DhcpOptionsSet) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *DhcpOptionsSet) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *DhcpOptionsSet) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetTags

`func (o *DhcpOptionsSet) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DhcpOptionsSet) GetTagsOk() (*[]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DhcpOptionsSet) SetTags(v []ResourceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DhcpOptionsSet) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


