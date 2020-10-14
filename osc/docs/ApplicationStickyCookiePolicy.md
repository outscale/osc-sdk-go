# ApplicationStickyCookiePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CookieName** | Pointer to **string** | The name of the application cookie used for stickiness. | [optional] 
**PolicyName** | Pointer to **string** | The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer. | [optional] 

## Methods

### NewApplicationStickyCookiePolicy

`func NewApplicationStickyCookiePolicy() *ApplicationStickyCookiePolicy`

NewApplicationStickyCookiePolicy instantiates a new ApplicationStickyCookiePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStickyCookiePolicyWithDefaults

`func NewApplicationStickyCookiePolicyWithDefaults() *ApplicationStickyCookiePolicy`

NewApplicationStickyCookiePolicyWithDefaults instantiates a new ApplicationStickyCookiePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookieName

`func (o *ApplicationStickyCookiePolicy) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *ApplicationStickyCookiePolicy) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *ApplicationStickyCookiePolicy) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *ApplicationStickyCookiePolicy) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApplicationStickyCookiePolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApplicationStickyCookiePolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApplicationStickyCookiePolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApplicationStickyCookiePolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


