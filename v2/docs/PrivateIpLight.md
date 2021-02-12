# PrivateIpLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | Pointer to **bool** | If true, the IP address is the primary private IP address of the NIC. | [optional] 
**PrivateIp** | Pointer to **string** | The private IP address of the NIC. | [optional] 

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


