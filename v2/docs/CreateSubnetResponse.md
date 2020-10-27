# CreateSubnetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Subnet** | Pointer to [**Subnet**](Subnet.md) |  | [optional] 

## Methods

### NewCreateSubnetResponse

`func NewCreateSubnetResponse() *CreateSubnetResponse`

NewCreateSubnetResponse instantiates a new CreateSubnetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetResponseWithDefaults

`func NewCreateSubnetResponseWithDefaults() *CreateSubnetResponse`

NewCreateSubnetResponseWithDefaults instantiates a new CreateSubnetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContext

`func (o *CreateSubnetResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *CreateSubnetResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *CreateSubnetResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *CreateSubnetResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetSubnet

`func (o *CreateSubnetResponse) GetSubnet() Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateSubnetResponse) GetSubnetOk() (*Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateSubnetResponse) SetSubnet(v Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CreateSubnetResponse) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


