# ReadConsoleOutputResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsoleOutput** | Pointer to **string** | The Base64-encoded output of the console. If a command line tool is used, the output is decoded by the tool. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 

## Methods

### NewReadConsoleOutputResponse

`func NewReadConsoleOutputResponse() *ReadConsoleOutputResponse`

NewReadConsoleOutputResponse instantiates a new ReadConsoleOutputResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadConsoleOutputResponseWithDefaults

`func NewReadConsoleOutputResponseWithDefaults() *ReadConsoleOutputResponse`

NewReadConsoleOutputResponseWithDefaults instantiates a new ReadConsoleOutputResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsoleOutput

`func (o *ReadConsoleOutputResponse) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *ReadConsoleOutputResponse) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *ReadConsoleOutputResponse) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *ReadConsoleOutputResponse) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetResponseContext

`func (o *ReadConsoleOutputResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadConsoleOutputResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *ReadConsoleOutputResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *ReadConsoleOutputResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### GetVmId

`func (o *ReadConsoleOutputResponse) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ReadConsoleOutputResponse) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ReadConsoleOutputResponse) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ReadConsoleOutputResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


