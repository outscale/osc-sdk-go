# RejectNetPeeringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**NetPeeringId** | **string** | The ID of the Net peering connection you want to reject. | 

## Methods

### NewRejectNetPeeringRequest

`func NewRejectNetPeeringRequest(netPeeringId string, ) *RejectNetPeeringRequest`

NewRejectNetPeeringRequest instantiates a new RejectNetPeeringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectNetPeeringRequestWithDefaults

`func NewRejectNetPeeringRequestWithDefaults() *RejectNetPeeringRequest`

NewRejectNetPeeringRequestWithDefaults instantiates a new RejectNetPeeringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *RejectNetPeeringRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RejectNetPeeringRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RejectNetPeeringRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RejectNetPeeringRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetPeeringId

`func (o *RejectNetPeeringRequest) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *RejectNetPeeringRequest) GetNetPeeringIdOk() (*string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeeringId

`func (o *RejectNetPeeringRequest) SetNetPeeringId(v string)`

SetNetPeeringId sets NetPeeringId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


