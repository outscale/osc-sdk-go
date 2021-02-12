# DeleteNetAccessPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetAccessPointId** | **string** | The ID of the Net access point. | 

## Methods

### NewDeleteNetAccessPointRequest

`func NewDeleteNetAccessPointRequest(netAccessPointId string, ) *DeleteNetAccessPointRequest`

NewDeleteNetAccessPointRequest instantiates a new DeleteNetAccessPointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteNetAccessPointRequestWithDefaults

`func NewDeleteNetAccessPointRequestWithDefaults() *DeleteNetAccessPointRequest`

NewDeleteNetAccessPointRequestWithDefaults instantiates a new DeleteNetAccessPointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DeleteNetAccessPointRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteNetAccessPointRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteNetAccessPointRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteNetAccessPointRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetAccessPointId

`func (o *DeleteNetAccessPointRequest) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *DeleteNetAccessPointRequest) GetNetAccessPointIdOk() (*string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAccessPointId

`func (o *DeleteNetAccessPointRequest) SetNetAccessPointId(v string)`

SetNetAccessPointId sets NetAccessPointId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


