/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadApiLogsResponse struct for ReadApiLogsResponse
type ReadApiLogsResponse struct {
	// Information displayed in one or more API logs.
	Logs *[]Log `json:"Logs,omitempty"`
	// The token to request the next page of results.
	NextPageToken *string `json:"NextPageToken,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadApiLogsResponse instantiates a new ReadApiLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApiLogsResponse() *ReadApiLogsResponse {
	this := ReadApiLogsResponse{}
	return &this
}

// NewReadApiLogsResponseWithDefaults instantiates a new ReadApiLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApiLogsResponseWithDefaults() *ReadApiLogsResponse {
	this := ReadApiLogsResponse{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ReadApiLogsResponse) GetLogs() []Log {
	if o == nil || o.Logs == nil {
		var ret []Log
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsResponse) GetLogsOk() (*[]Log, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ReadApiLogsResponse) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []Log and assigns it to the Logs field.
func (o *ReadApiLogsResponse) SetLogs(v []Log) {
	o.Logs = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ReadApiLogsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ReadApiLogsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ReadApiLogsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadApiLogsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadApiLogsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadApiLogsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadApiLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Logs != nil {
		toSerialize["Logs"] = o.Logs
	}
	if o.NextPageToken != nil {
		toSerialize["NextPageToken"] = o.NextPageToken
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadApiLogsResponse struct {
	value *ReadApiLogsResponse
	isSet bool
}

func (v NullableReadApiLogsResponse) Get() *ReadApiLogsResponse {
	return v.value
}

func (v *NullableReadApiLogsResponse) Set(val *ReadApiLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApiLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApiLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApiLogsResponse(val *ReadApiLogsResponse) *NullableReadApiLogsResponse {
	return &NullableReadApiLogsResponse{value: val, isSet: true}
}

func (v NullableReadApiLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApiLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


