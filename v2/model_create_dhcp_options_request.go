/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.<br /><br /> An OpenAPI description of the OUTSCALE API is also available in this [GitHub repository](https://github.com/outscale/osc-api).
 *
 * API version: 1.23
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateDhcpOptionsRequest struct for CreateDhcpOptionsRequest
type CreateDhcpOptionsRequest struct {
	// Specify a domain name (for example, `MyCompany.com`). You can specify only one domain name. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers`, `LogServers`, or `NtpServers`.
	DomainName *string `json:"DomainName,omitempty"`
	// The IPs of domain name servers. If no IPs are specified, the `OutscaleProvidedDNS` value is set by default. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers`, `LogServers`, or `NtpServers`.
	DomainNameServers *[]string `json:"DomainNameServers,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The IPs of the log servers. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers`, `LogServers`, or `NtpServers`.
	LogServers *[]string `json:"LogServers,omitempty"`
	// The IPs of the Network Time Protocol (NTP) servers. You must specify at least one of the following parameters: `DomainName`, `DomainNameServers`, `LogServers`, or `NtpServers`.
	NtpServers *[]string `json:"NtpServers,omitempty"`
}

// NewCreateDhcpOptionsRequest instantiates a new CreateDhcpOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDhcpOptionsRequest() *CreateDhcpOptionsRequest {
	this := CreateDhcpOptionsRequest{}
	return &this
}

// NewCreateDhcpOptionsRequestWithDefaults instantiates a new CreateDhcpOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDhcpOptionsRequestWithDefaults() *CreateDhcpOptionsRequest {
	this := CreateDhcpOptionsRequest{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *CreateDhcpOptionsRequest) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsRequest) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *CreateDhcpOptionsRequest) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *CreateDhcpOptionsRequest) SetDomainName(v string) {
	o.DomainName = &v
}

// GetDomainNameServers returns the DomainNameServers field value if set, zero value otherwise.
func (o *CreateDhcpOptionsRequest) GetDomainNameServers() []string {
	if o == nil || o.DomainNameServers == nil {
		var ret []string
		return ret
	}
	return *o.DomainNameServers
}

// GetDomainNameServersOk returns a tuple with the DomainNameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsRequest) GetDomainNameServersOk() (*[]string, bool) {
	if o == nil || o.DomainNameServers == nil {
		return nil, false
	}
	return o.DomainNameServers, true
}

// HasDomainNameServers returns a boolean if a field has been set.
func (o *CreateDhcpOptionsRequest) HasDomainNameServers() bool {
	if o != nil && o.DomainNameServers != nil {
		return true
	}

	return false
}

// SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.
func (o *CreateDhcpOptionsRequest) SetDomainNameServers(v []string) {
	o.DomainNameServers = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateDhcpOptionsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateDhcpOptionsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateDhcpOptionsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLogServers returns the LogServers field value if set, zero value otherwise.
func (o *CreateDhcpOptionsRequest) GetLogServers() []string {
	if o == nil || o.LogServers == nil {
		var ret []string
		return ret
	}
	return *o.LogServers
}

// GetLogServersOk returns a tuple with the LogServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsRequest) GetLogServersOk() (*[]string, bool) {
	if o == nil || o.LogServers == nil {
		return nil, false
	}
	return o.LogServers, true
}

// HasLogServers returns a boolean if a field has been set.
func (o *CreateDhcpOptionsRequest) HasLogServers() bool {
	if o != nil && o.LogServers != nil {
		return true
	}

	return false
}

// SetLogServers gets a reference to the given []string and assigns it to the LogServers field.
func (o *CreateDhcpOptionsRequest) SetLogServers(v []string) {
	o.LogServers = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *CreateDhcpOptionsRequest) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsRequest) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *CreateDhcpOptionsRequest) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *CreateDhcpOptionsRequest) SetNtpServers(v []string) {
	o.NtpServers = &v
}

func (o CreateDhcpOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.DomainNameServers != nil {
		toSerialize["DomainNameServers"] = o.DomainNameServers
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.LogServers != nil {
		toSerialize["LogServers"] = o.LogServers
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDhcpOptionsRequest struct {
	value *CreateDhcpOptionsRequest
	isSet bool
}

func (v NullableCreateDhcpOptionsRequest) Get() *CreateDhcpOptionsRequest {
	return v.value
}

func (v *NullableCreateDhcpOptionsRequest) Set(val *CreateDhcpOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDhcpOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDhcpOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDhcpOptionsRequest(val *CreateDhcpOptionsRequest) *NullableCreateDhcpOptionsRequest {
	return &NullableCreateDhcpOptionsRequest{value: val, isSet: true}
}

func (v NullableCreateDhcpOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDhcpOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
