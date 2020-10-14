# DeleteClientGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayId** | **string** | The ID of the client gateway you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewDeleteClientGatewayRequest

`func NewDeleteClientGatewayRequest(clientGatewayId string, ) *DeleteClientGatewayRequest`

NewDeleteClientGatewayRequest instantiates a new DeleteClientGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteClientGatewayRequestWithDefaults

`func NewDeleteClientGatewayRequestWithDefaults() *DeleteClientGatewayRequest`

NewDeleteClientGatewayRequestWithDefaults instantiates a new DeleteClientGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGatewayId

`func (o *DeleteClientGatewayRequest) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *DeleteClientGatewayRequest) GetClientGatewayIdOk() (*string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGatewayId

`func (o *DeleteClientGatewayRequest) SetClientGatewayId(v string)`

SetClientGatewayId sets ClientGatewayId field to given value.


### GetDryRun

`func (o *DeleteClientGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteClientGatewayRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteClientGatewayRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteClientGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


