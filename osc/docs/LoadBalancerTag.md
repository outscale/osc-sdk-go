# LoadBalancerTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the tag. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer. | [optional] 
**Value** | Pointer to **string** | The value of the tag. | [optional] 

## Methods

### NewLoadBalancerTag

`func NewLoadBalancerTag() *LoadBalancerTag`

NewLoadBalancerTag instantiates a new LoadBalancerTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerTagWithDefaults

`func NewLoadBalancerTagWithDefaults() *LoadBalancerTag`

NewLoadBalancerTagWithDefaults instantiates a new LoadBalancerTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LoadBalancerTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LoadBalancerTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LoadBalancerTag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LoadBalancerTag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLoadBalancerName

`func (o *LoadBalancerTag) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *LoadBalancerTag) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *LoadBalancerTag) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.

### HasLoadBalancerName

`func (o *LoadBalancerTag) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### GetValue

`func (o *LoadBalancerTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoadBalancerTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoadBalancerTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LoadBalancerTag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


