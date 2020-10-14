# PublicIpLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIp** | Pointer to **string** | The External IP address (EIP) associated with the NAT service. | [optional] 
**PublicIpId** | Pointer to **string** | The allocation ID of the EIP associated with the NAT service. | [optional] 

## Methods

### NewPublicIpLight

`func NewPublicIpLight() *PublicIpLight`

NewPublicIpLight instantiates a new PublicIpLight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIpLightWithDefaults

`func NewPublicIpLightWithDefaults() *PublicIpLight`

NewPublicIpLightWithDefaults instantiates a new PublicIpLight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIp

`func (o *PublicIpLight) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *PublicIpLight) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *PublicIpLight) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *PublicIpLight) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicIpId

`func (o *PublicIpLight) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *PublicIpLight) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *PublicIpLight) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *PublicIpLight) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


