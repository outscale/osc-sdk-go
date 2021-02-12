# CreateNatServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIpId** | **string** | The allocation ID of the EIP to associate with the NAT service.&lt;br /&gt; If the EIP is already associated with another resource, you must first disassociate it. | 
**SubnetId** | **string** | The ID of the Subnet in which you want to create the NAT service. | 

## Methods

### NewCreateNatServiceRequest

`func NewCreateNatServiceRequest(publicIpId string, subnetId string, ) *CreateNatServiceRequest`

NewCreateNatServiceRequest instantiates a new CreateNatServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNatServiceRequestWithDefaults

`func NewCreateNatServiceRequestWithDefaults() *CreateNatServiceRequest`

NewCreateNatServiceRequestWithDefaults instantiates a new CreateNatServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateNatServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNatServiceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateNatServiceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateNatServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPublicIpId

`func (o *CreateNatServiceRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *CreateNatServiceRequest) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *CreateNatServiceRequest) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.


### GetSubnetId

`func (o *CreateNatServiceRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateNatServiceRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateNatServiceRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


