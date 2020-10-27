# ReadDirectLinkInterfacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersDirectLinkInterface**](FiltersDirectLinkInterface.md) |  | [optional] 

## Methods

### NewReadDirectLinkInterfacesRequest

`func NewReadDirectLinkInterfacesRequest() *ReadDirectLinkInterfacesRequest`

NewReadDirectLinkInterfacesRequest instantiates a new ReadDirectLinkInterfacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDirectLinkInterfacesRequestWithDefaults

`func NewReadDirectLinkInterfacesRequestWithDefaults() *ReadDirectLinkInterfacesRequest`

NewReadDirectLinkInterfacesRequestWithDefaults instantiates a new ReadDirectLinkInterfacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadDirectLinkInterfacesRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadDirectLinkInterfacesRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadDirectLinkInterfacesRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadDirectLinkInterfacesRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFilters

`func (o *ReadDirectLinkInterfacesRequest) GetFilters() FiltersDirectLinkInterface`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadDirectLinkInterfacesRequest) GetFiltersOk() (*FiltersDirectLinkInterface, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReadDirectLinkInterfacesRequest) SetFilters(v FiltersDirectLinkInterface)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReadDirectLinkInterfacesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


