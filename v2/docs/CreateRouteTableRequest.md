# CreateRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | **string** | The ID of the Net for which you want to create a route table. | 

## Methods

### NewCreateRouteTableRequest

`func NewCreateRouteTableRequest(netId string, ) *CreateRouteTableRequest`

NewCreateRouteTableRequest instantiates a new CreateRouteTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteTableRequestWithDefaults

`func NewCreateRouteTableRequestWithDefaults() *CreateRouteTableRequest`

NewCreateRouteTableRequestWithDefaults instantiates a new CreateRouteTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *CreateRouteTableRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateRouteTableRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateRouteTableRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateRouteTableRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetId

`func (o *CreateRouteTableRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateRouteTableRequest) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *CreateRouteTableRequest) SetNetId(v string)`

SetNetId sets NetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


