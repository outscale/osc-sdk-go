# UnlinkRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LinkRouteTableId** | **string** | The ID of the association between the route table and the Subnet. | 

## Methods

### NewUnlinkRouteTableRequest

`func NewUnlinkRouteTableRequest(linkRouteTableId string, ) *UnlinkRouteTableRequest`

NewUnlinkRouteTableRequest instantiates a new UnlinkRouteTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkRouteTableRequestWithDefaults

`func NewUnlinkRouteTableRequestWithDefaults() *UnlinkRouteTableRequest`

NewUnlinkRouteTableRequestWithDefaults instantiates a new UnlinkRouteTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *UnlinkRouteTableRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkRouteTableRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *UnlinkRouteTableRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *UnlinkRouteTableRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetLinkRouteTableId

`func (o *UnlinkRouteTableRequest) GetLinkRouteTableId() string`

GetLinkRouteTableId returns the LinkRouteTableId field if non-nil, zero value otherwise.

### GetLinkRouteTableIdOk

`func (o *UnlinkRouteTableRequest) GetLinkRouteTableIdOk() (*string, bool)`

GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRouteTableId

`func (o *UnlinkRouteTableRequest) SetLinkRouteTableId(v string)`

SetLinkRouteTableId sets LinkRouteTableId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


