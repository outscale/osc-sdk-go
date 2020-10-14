# AcceptNetPeeringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetPeeringId** | **string** | The ID of the Net peering connection you want to accept. | 

## Methods

### NewAcceptNetPeeringRequest

`func NewAcceptNetPeeringRequest(netPeeringId string, ) *AcceptNetPeeringRequest`

NewAcceptNetPeeringRequest instantiates a new AcceptNetPeeringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptNetPeeringRequestWithDefaults

`func NewAcceptNetPeeringRequestWithDefaults() *AcceptNetPeeringRequest`

NewAcceptNetPeeringRequestWithDefaults instantiates a new AcceptNetPeeringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *AcceptNetPeeringRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AcceptNetPeeringRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AcceptNetPeeringRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AcceptNetPeeringRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetNetPeeringId

`func (o *AcceptNetPeeringRequest) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *AcceptNetPeeringRequest) GetNetPeeringIdOk() (*string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPeeringId

`func (o *AcceptNetPeeringRequest) SetNetPeeringId(v string)`

SetNetPeeringId sets NetPeeringId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


