# DeleteDirectLinkInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkInterfaceId** | **string** | The ID of the DirectLink interface you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewDeleteDirectLinkInterfaceRequest

`func NewDeleteDirectLinkInterfaceRequest(directLinkInterfaceId string, ) *DeleteDirectLinkInterfaceRequest`

NewDeleteDirectLinkInterfaceRequest instantiates a new DeleteDirectLinkInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDirectLinkInterfaceRequestWithDefaults

`func NewDeleteDirectLinkInterfaceRequestWithDefaults() *DeleteDirectLinkInterfaceRequest`

NewDeleteDirectLinkInterfaceRequestWithDefaults instantiates a new DeleteDirectLinkInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLinkInterfaceId

`func (o *DeleteDirectLinkInterfaceRequest) GetDirectLinkInterfaceId() string`

GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceIdOk

`func (o *DeleteDirectLinkInterfaceRequest) GetDirectLinkInterfaceIdOk() (*string, bool)`

GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterfaceId

`func (o *DeleteDirectLinkInterfaceRequest) SetDirectLinkInterfaceId(v string)`

SetDirectLinkInterfaceId sets DirectLinkInterfaceId field to given value.


### GetDryRun

`func (o *DeleteDirectLinkInterfaceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteDirectLinkInterfaceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteDirectLinkInterfaceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteDirectLinkInterfaceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


