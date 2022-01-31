# PrivateIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | Pointer to **bool** | If true, the IP is the primary private IP of the NIC. | [optional] 
**LinkPublicIp** | Pointer to [**LinkPublicIp**](LinkPublicIp.md) |  | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIp** | Pointer to **string** | The private IP of the NIC. | [optional] 

## Methods

### NewPrivateIp

`func NewPrivateIp() *PrivateIp`

NewPrivateIp instantiates a new PrivateIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateIpWithDefaults

`func NewPrivateIpWithDefaults() *PrivateIp`

NewPrivateIpWithDefaults instantiates a new PrivateIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrimary

`func (o *PrivateIp) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PrivateIp) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PrivateIp) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PrivateIp) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetLinkPublicIp

`func (o *PrivateIp) GetLinkPublicIp() LinkPublicIp`

GetLinkPublicIp returns the LinkPublicIp field if non-nil, zero value otherwise.

### GetLinkPublicIpOk

`func (o *PrivateIp) GetLinkPublicIpOk() (*LinkPublicIp, bool)`

GetLinkPublicIpOk returns a tuple with the LinkPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPublicIp

`func (o *PrivateIp) SetLinkPublicIp(v LinkPublicIp)`

SetLinkPublicIp sets LinkPublicIp field to given value.

### HasLinkPublicIp

`func (o *PrivateIp) HasLinkPublicIp() bool`

HasLinkPublicIp returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *PrivateIp) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *PrivateIp) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *PrivateIp) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *PrivateIp) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIp

`func (o *PrivateIp) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PrivateIp) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *PrivateIp) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *PrivateIp) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


