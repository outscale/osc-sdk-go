# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the logged call. | [optional] 
**CallDuration** | Pointer to **int32** | The duration of the logged call, in microseconds. | [optional] 
**QueryAccessKey** | Pointer to **string** | The access key used for the logged call. | [optional] 
**QueryApiName** | Pointer to **string** | The name of the API used by the logged call (always &#x60;oapi&#x60; for the OUTSCALE API). | [optional] 
**QueryApiVersion** | Pointer to **string** | The version of the API used by the logged call. | [optional] 
**QueryCallName** | Pointer to **string** | The name of the logged call. | [optional] 
**QueryDate** | Pointer to **string** | The date of the logged call, in ISO 8601 format. | [optional] 
**QueryHeaderRaw** | Pointer to **string** | The raw header of the HTTP request of the logged call. | [optional] 
**QueryHeaderSize** | Pointer to **int32** | The size of the raw header of the HTTP request of the logged call, in bytes. | [optional] 
**QueryIpAddress** | Pointer to **string** | The IP used for the logged call. | [optional] 
**QueryPayloadRaw** | Pointer to **string** | The raw payload of the HTTP request of the logged call. | [optional] 
**QueryPayloadSize** | Pointer to **int32** | The size of the raw payload of the HTTP request of the logged call, in bytes. | [optional] 
**QueryUserAgent** | Pointer to **string** | The user agent of the HTTP request of the logged call. | [optional] 
**RequestId** | Pointer to **string** | The request ID provided in the response of the logged call. | [optional] 
**ResponseSize** | Pointer to **int32** | The size of the response of the logged call, in bytes. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | The HTTP status code of the response of the logged call. | [optional] 

## Methods

### NewLog

`func NewLog() *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Log) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Log) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Log) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Log) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCallDuration

`func (o *Log) GetCallDuration() int32`

GetCallDuration returns the CallDuration field if non-nil, zero value otherwise.

### GetCallDurationOk

`func (o *Log) GetCallDurationOk() (*int32, bool)`

GetCallDurationOk returns a tuple with the CallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDuration

`func (o *Log) SetCallDuration(v int32)`

SetCallDuration sets CallDuration field to given value.

### HasCallDuration

`func (o *Log) HasCallDuration() bool`

HasCallDuration returns a boolean if a field has been set.

### GetQueryAccessKey

`func (o *Log) GetQueryAccessKey() string`

GetQueryAccessKey returns the QueryAccessKey field if non-nil, zero value otherwise.

### GetQueryAccessKeyOk

`func (o *Log) GetQueryAccessKeyOk() (*string, bool)`

GetQueryAccessKeyOk returns a tuple with the QueryAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAccessKey

`func (o *Log) SetQueryAccessKey(v string)`

SetQueryAccessKey sets QueryAccessKey field to given value.

### HasQueryAccessKey

`func (o *Log) HasQueryAccessKey() bool`

HasQueryAccessKey returns a boolean if a field has been set.

### GetQueryApiName

`func (o *Log) GetQueryApiName() string`

GetQueryApiName returns the QueryApiName field if non-nil, zero value otherwise.

### GetQueryApiNameOk

`func (o *Log) GetQueryApiNameOk() (*string, bool)`

GetQueryApiNameOk returns a tuple with the QueryApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryApiName

`func (o *Log) SetQueryApiName(v string)`

SetQueryApiName sets QueryApiName field to given value.

### HasQueryApiName

`func (o *Log) HasQueryApiName() bool`

HasQueryApiName returns a boolean if a field has been set.

### GetQueryApiVersion

`func (o *Log) GetQueryApiVersion() string`

GetQueryApiVersion returns the QueryApiVersion field if non-nil, zero value otherwise.

### GetQueryApiVersionOk

`func (o *Log) GetQueryApiVersionOk() (*string, bool)`

GetQueryApiVersionOk returns a tuple with the QueryApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryApiVersion

`func (o *Log) SetQueryApiVersion(v string)`

SetQueryApiVersion sets QueryApiVersion field to given value.

### HasQueryApiVersion

`func (o *Log) HasQueryApiVersion() bool`

HasQueryApiVersion returns a boolean if a field has been set.

### GetQueryCallName

`func (o *Log) GetQueryCallName() string`

GetQueryCallName returns the QueryCallName field if non-nil, zero value otherwise.

### GetQueryCallNameOk

`func (o *Log) GetQueryCallNameOk() (*string, bool)`

GetQueryCallNameOk returns a tuple with the QueryCallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCallName

`func (o *Log) SetQueryCallName(v string)`

SetQueryCallName sets QueryCallName field to given value.

### HasQueryCallName

`func (o *Log) HasQueryCallName() bool`

HasQueryCallName returns a boolean if a field has been set.

### GetQueryDate

`func (o *Log) GetQueryDate() string`

GetQueryDate returns the QueryDate field if non-nil, zero value otherwise.

### GetQueryDateOk

`func (o *Log) GetQueryDateOk() (*string, bool)`

GetQueryDateOk returns a tuple with the QueryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryDate

`func (o *Log) SetQueryDate(v string)`

SetQueryDate sets QueryDate field to given value.

### HasQueryDate

`func (o *Log) HasQueryDate() bool`

HasQueryDate returns a boolean if a field has been set.

### GetQueryHeaderRaw

`func (o *Log) GetQueryHeaderRaw() string`

GetQueryHeaderRaw returns the QueryHeaderRaw field if non-nil, zero value otherwise.

### GetQueryHeaderRawOk

`func (o *Log) GetQueryHeaderRawOk() (*string, bool)`

GetQueryHeaderRawOk returns a tuple with the QueryHeaderRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryHeaderRaw

`func (o *Log) SetQueryHeaderRaw(v string)`

SetQueryHeaderRaw sets QueryHeaderRaw field to given value.

### HasQueryHeaderRaw

`func (o *Log) HasQueryHeaderRaw() bool`

HasQueryHeaderRaw returns a boolean if a field has been set.

### GetQueryHeaderSize

`func (o *Log) GetQueryHeaderSize() int32`

GetQueryHeaderSize returns the QueryHeaderSize field if non-nil, zero value otherwise.

### GetQueryHeaderSizeOk

`func (o *Log) GetQueryHeaderSizeOk() (*int32, bool)`

GetQueryHeaderSizeOk returns a tuple with the QueryHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryHeaderSize

`func (o *Log) SetQueryHeaderSize(v int32)`

SetQueryHeaderSize sets QueryHeaderSize field to given value.

### HasQueryHeaderSize

`func (o *Log) HasQueryHeaderSize() bool`

HasQueryHeaderSize returns a boolean if a field has been set.

### GetQueryIpAddress

`func (o *Log) GetQueryIpAddress() string`

GetQueryIpAddress returns the QueryIpAddress field if non-nil, zero value otherwise.

### GetQueryIpAddressOk

`func (o *Log) GetQueryIpAddressOk() (*string, bool)`

GetQueryIpAddressOk returns a tuple with the QueryIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIpAddress

`func (o *Log) SetQueryIpAddress(v string)`

SetQueryIpAddress sets QueryIpAddress field to given value.

### HasQueryIpAddress

`func (o *Log) HasQueryIpAddress() bool`

HasQueryIpAddress returns a boolean if a field has been set.

### GetQueryPayloadRaw

`func (o *Log) GetQueryPayloadRaw() string`

GetQueryPayloadRaw returns the QueryPayloadRaw field if non-nil, zero value otherwise.

### GetQueryPayloadRawOk

`func (o *Log) GetQueryPayloadRawOk() (*string, bool)`

GetQueryPayloadRawOk returns a tuple with the QueryPayloadRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPayloadRaw

`func (o *Log) SetQueryPayloadRaw(v string)`

SetQueryPayloadRaw sets QueryPayloadRaw field to given value.

### HasQueryPayloadRaw

`func (o *Log) HasQueryPayloadRaw() bool`

HasQueryPayloadRaw returns a boolean if a field has been set.

### GetQueryPayloadSize

`func (o *Log) GetQueryPayloadSize() int32`

GetQueryPayloadSize returns the QueryPayloadSize field if non-nil, zero value otherwise.

### GetQueryPayloadSizeOk

`func (o *Log) GetQueryPayloadSizeOk() (*int32, bool)`

GetQueryPayloadSizeOk returns a tuple with the QueryPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPayloadSize

`func (o *Log) SetQueryPayloadSize(v int32)`

SetQueryPayloadSize sets QueryPayloadSize field to given value.

### HasQueryPayloadSize

`func (o *Log) HasQueryPayloadSize() bool`

HasQueryPayloadSize returns a boolean if a field has been set.

### GetQueryUserAgent

`func (o *Log) GetQueryUserAgent() string`

GetQueryUserAgent returns the QueryUserAgent field if non-nil, zero value otherwise.

### GetQueryUserAgentOk

`func (o *Log) GetQueryUserAgentOk() (*string, bool)`

GetQueryUserAgentOk returns a tuple with the QueryUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryUserAgent

`func (o *Log) SetQueryUserAgent(v string)`

SetQueryUserAgent sets QueryUserAgent field to given value.

### HasQueryUserAgent

`func (o *Log) HasQueryUserAgent() bool`

HasQueryUserAgent returns a boolean if a field has been set.

### GetRequestId

`func (o *Log) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Log) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Log) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Log) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResponseSize

`func (o *Log) GetResponseSize() int32`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *Log) GetResponseSizeOk() (*int32, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSize

`func (o *Log) SetResponseSize(v int32)`

SetResponseSize sets ResponseSize field to given value.

### HasResponseSize

`func (o *Log) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *Log) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *Log) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *Log) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *Log) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


