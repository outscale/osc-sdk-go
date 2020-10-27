# AccessLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If &#x60;true&#x60;, access logs are enabled for your load balancer. If &#x60;false&#x60;, they are not. If you set this to &#x60;true&#x60; in your request, the &#x60;OsuBucketName&#x60; parameter is required. | [optional] 
**OsuBucketName** | Pointer to **string** | The name of the Object Storage Unit (OSU) bucket for the access logs. | [optional] 
**OsuBucketPrefix** | Pointer to **string** | The path to the folder of the access logs in your Object Storage Unit (OSU) bucket (by default, the &#x60;root&#x60; level of your bucket). | [optional] 
**PublicationInterval** | Pointer to **int32** | The time interval for the publication of access logs in the Object Storage Unit (OSU) bucket, in minutes. This value can be either 5 or 60 (by default, 60). | [optional] 

## Methods

### NewAccessLog

`func NewAccessLog() *AccessLog`

NewAccessLog instantiates a new AccessLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessLogWithDefaults

`func NewAccessLogWithDefaults() *AccessLog`

NewAccessLogWithDefaults instantiates a new AccessLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AccessLog) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AccessLog) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AccessLog) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AccessLog) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetOsuBucketName

`func (o *AccessLog) GetOsuBucketName() string`

GetOsuBucketName returns the OsuBucketName field if non-nil, zero value otherwise.

### GetOsuBucketNameOk

`func (o *AccessLog) GetOsuBucketNameOk() (*string, bool)`

GetOsuBucketNameOk returns a tuple with the OsuBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuBucketName

`func (o *AccessLog) SetOsuBucketName(v string)`

SetOsuBucketName sets OsuBucketName field to given value.

### HasOsuBucketName

`func (o *AccessLog) HasOsuBucketName() bool`

HasOsuBucketName returns a boolean if a field has been set.

### GetOsuBucketPrefix

`func (o *AccessLog) GetOsuBucketPrefix() string`

GetOsuBucketPrefix returns the OsuBucketPrefix field if non-nil, zero value otherwise.

### GetOsuBucketPrefixOk

`func (o *AccessLog) GetOsuBucketPrefixOk() (*string, bool)`

GetOsuBucketPrefixOk returns a tuple with the OsuBucketPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuBucketPrefix

`func (o *AccessLog) SetOsuBucketPrefix(v string)`

SetOsuBucketPrefix sets OsuBucketPrefix field to given value.

### HasOsuBucketPrefix

`func (o *AccessLog) HasOsuBucketPrefix() bool`

HasOsuBucketPrefix returns a boolean if a field has been set.

### GetPublicationInterval

`func (o *AccessLog) GetPublicationInterval() int32`

GetPublicationInterval returns the PublicationInterval field if non-nil, zero value otherwise.

### GetPublicationIntervalOk

`func (o *AccessLog) GetPublicationIntervalOk() (*int32, bool)`

GetPublicationIntervalOk returns a tuple with the PublicationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationInterval

`func (o *AccessLog) SetPublicationInterval(v int32)`

SetPublicationInterval sets PublicationInterval field to given value.

### HasPublicationInterval

`func (o *AccessLog) HasPublicationInterval() bool`

HasPublicationInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


