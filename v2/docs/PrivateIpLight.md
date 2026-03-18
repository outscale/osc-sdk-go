# PrivateIpLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | Pointer to **bool** | If true, the IP is the primary private IP of the NIC. | [optional] 
**PrivateIp** | Pointer to **string** | A private IP for the NIC. This IP must be within the IP range of the Subnet that you specify with the &#x60;SubnetId&#x60; parameter. However, it cannot be one of the first four IPs (ending in &#x60;.0&#x60;, &#x60;.1&#x60;, &#x60;.2&#x60;, &#x60;.3&#x60;) or the last IP (ending in &#x60;.255&#x60;) of the Subnet, as these are reserved by 3DS OUTSCALE. For more information, see [About Nets](https://docs.outscale.com/en/userguide/About-Nets.html). | [optional] 

## Methods

### NewPrivateIpLight

`func NewPrivateIpLight() *PrivateIpLight`

NewPrivateIpLight instantiates a new PrivateIpLight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateIpLightWithDefaults

`func NewPrivateIpLightWithDefaults() *PrivateIpLight`

NewPrivateIpLightWithDefaults instantiates a new PrivateIpLight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrimary

`func (o *PrivateIpLight) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PrivateIpLight) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PrivateIpLight) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PrivateIpLight) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetPrivateIp

`func (o *PrivateIpLight) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PrivateIpLight) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *PrivateIpLight) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *PrivateIpLight) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


