# CreateDirectLinkInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkId** | **string** | The ID of the existing DirectLink for which you want to create the DirectLink interface. | 
**DirectLinkInterface** | [**DirectLinkInterface**](DirectLinkInterface.md) |  | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewCreateDirectLinkInterfaceRequest

`func NewCreateDirectLinkInterfaceRequest(directLinkId string, directLinkInterface DirectLinkInterface, ) *CreateDirectLinkInterfaceRequest`

NewCreateDirectLinkInterfaceRequest instantiates a new CreateDirectLinkInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDirectLinkInterfaceRequestWithDefaults

`func NewCreateDirectLinkInterfaceRequestWithDefaults() *CreateDirectLinkInterfaceRequest`

NewCreateDirectLinkInterfaceRequestWithDefaults instantiates a new CreateDirectLinkInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLinkId

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkIdOk() (*string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkId

`func (o *CreateDirectLinkInterfaceRequest) SetDirectLinkId(v string)`

SetDirectLinkId sets DirectLinkId field to given value.


### GetDirectLinkInterface

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkInterface() DirectLinkInterface`

GetDirectLinkInterface returns the DirectLinkInterface field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceOk

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkInterfaceOk() (*DirectLinkInterface, bool)`

GetDirectLinkInterfaceOk returns a tuple with the DirectLinkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLinkInterface

`func (o *CreateDirectLinkInterfaceRequest) SetDirectLinkInterface(v DirectLinkInterface)`

SetDirectLinkInterface sets DirectLinkInterface field to given value.


### GetDryRun

`func (o *CreateDirectLinkInterfaceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDirectLinkInterfaceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateDirectLinkInterfaceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *CreateDirectLinkInterfaceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


