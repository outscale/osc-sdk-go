# AccepterNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the accepter Net. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the accepter Net, in CIDR notation (for example, &#x60;10.0.0.0/16&#x60;). | [optional] 
**NetId** | Pointer to **string** | The ID of the accepter Net. | [optional] 

## Methods

### NewAccepterNet

`func NewAccepterNet() *AccepterNet`

NewAccepterNet instantiates a new AccepterNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccepterNetWithDefaults

`func NewAccepterNetWithDefaults() *AccepterNet`

NewAccepterNetWithDefaults instantiates a new AccepterNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccepterNet) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccepterNet) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccepterNet) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccepterNet) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIpRange

`func (o *AccepterNet) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *AccepterNet) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *AccepterNet) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *AccepterNet) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetNetId

`func (o *AccepterNet) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *AccepterNet) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *AccepterNet) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *AccepterNet) HasNetId() bool`

HasNetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


