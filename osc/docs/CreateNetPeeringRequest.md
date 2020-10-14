# CreateNetPeeringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterNetId** | **string** | The ID of the Net you want to connect with. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**SourceNetId** | **string** | The ID of the Net you send the peering request from. | 

## Methods

### NewCreateNetPeeringRequest

`func NewCreateNetPeeringRequest(accepterNetId string, sourceNetId string, ) *CreateNetPeeringRequest`

NewCreateNetPeeringRequest instantiates a new CreateNetPeeringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetPeeringRequestWithDefaults

`func NewCreateNetPeeringRequestWithDefaults() *CreateNetPeeringRequest`

NewCreateNetPeeringRequestWithDefaults instantiates a new CreateNetPeeringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepterNetId

`func (o *CreateNetPeeringRequest) GetAccepterNetId() string`

GetAccepterNetId returns the AccepterNetId field if non-nil, zero value otherwise.

### GetAccepterNetIdOk

`func (o *CreateNetPeeringRequest) GetAccepterNetIdOk() (*string, bool)`

GetAccepterNetIdOk returns a tuple with the AccepterNetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterNetId

`func (o *CreateNetPeeringRequest) SetAccepterNetId(v string)`

SetAccepterNetId sets AccepterNetId field to given value.


### GetDryRun

`func (o *CreateNetPeeringRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNetPeeringRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateNetPeeringRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateNetPeeringRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSourceNetId

`func (o *CreateNetPeeringRequest) GetSourceNetId() string`

GetSourceNetId returns the SourceNetId field if non-nil, zero value otherwise.

### GetSourceNetIdOk

`func (o *CreateNetPeeringRequest) GetSourceNetIdOk() (*string, bool)`

GetSourceNetIdOk returns a tuple with the SourceNetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetId

`func (o *CreateNetPeeringRequest) SetSourceNetId(v string)`

SetSourceNetId sets SourceNetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


