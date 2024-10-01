# ReadUnitPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | The operation associated with the catalog entry (for example, &#x60;RunInstances-OD&#x60; or &#x60;CreateVolume&#x60;). | 
**Service** | **string** | The service associated with the catalog entry (for example, &#x60;TinaOS-FCU&#x60; or &#x60;TinaOS-OOS&#x60;). | 
**Type** | **string** | The type associated with the catalog entry (for example, &#x60;BSU:VolumeIOPS:io1&#x60; or &#x60;BoxUsage:tinav6.c6r16p3&#x60;). | 

## Methods

### NewReadUnitPriceRequest

`func NewReadUnitPriceRequest(operation string, service string, type_ string, ) *ReadUnitPriceRequest`

NewReadUnitPriceRequest instantiates a new ReadUnitPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadUnitPriceRequestWithDefaults

`func NewReadUnitPriceRequestWithDefaults() *ReadUnitPriceRequest`

NewReadUnitPriceRequestWithDefaults instantiates a new ReadUnitPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ReadUnitPriceRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ReadUnitPriceRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ReadUnitPriceRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetService

`func (o *ReadUnitPriceRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ReadUnitPriceRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ReadUnitPriceRequest) SetService(v string)`

SetService sets Service field to given value.


### GetType

`func (o *ReadUnitPriceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReadUnitPriceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReadUnitPriceRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


