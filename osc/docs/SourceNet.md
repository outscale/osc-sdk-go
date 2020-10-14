# SourceNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the source Net. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**NetId** | Pointer to **string** | The ID of the source Net. | [optional] 

## Methods

### NewSourceNet

`func NewSourceNet() *SourceNet`

NewSourceNet instantiates a new SourceNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceNetWithDefaults

`func NewSourceNetWithDefaults() *SourceNet`

NewSourceNetWithDefaults instantiates a new SourceNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SourceNet) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SourceNet) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SourceNet) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SourceNet) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIpRange

`func (o *SourceNet) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SourceNet) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SourceNet) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *SourceNet) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetNetId

`func (o *SourceNet) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *SourceNet) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *SourceNet) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *SourceNet) HasNetId() bool`

HasNetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


