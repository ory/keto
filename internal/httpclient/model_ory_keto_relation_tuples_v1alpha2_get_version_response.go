/*
 * Ory Keto API
 *
 * Documentation for all of Ory Keto's REST APIs. gRPC is documented separately.
 *
 * API version:
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OryKetoRelationTuplesV1alpha2GetVersionResponse Response of the VersionService.GetVersion RPC.
type OryKetoRelationTuplesV1alpha2GetVersionResponse struct {
	// The version string of the Ory Keto instance.
	Version *string `json:"version,omitempty"`
}

// NewOryKetoRelationTuplesV1alpha2GetVersionResponse instantiates a new OryKetoRelationTuplesV1alpha2GetVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOryKetoRelationTuplesV1alpha2GetVersionResponse() *OryKetoRelationTuplesV1alpha2GetVersionResponse {
	this := OryKetoRelationTuplesV1alpha2GetVersionResponse{}
	return &this
}

// NewOryKetoRelationTuplesV1alpha2GetVersionResponseWithDefaults instantiates a new OryKetoRelationTuplesV1alpha2GetVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOryKetoRelationTuplesV1alpha2GetVersionResponseWithDefaults() *OryKetoRelationTuplesV1alpha2GetVersionResponse {
	this := OryKetoRelationTuplesV1alpha2GetVersionResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OryKetoRelationTuplesV1alpha2GetVersionResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OryKetoRelationTuplesV1alpha2GetVersionResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OryKetoRelationTuplesV1alpha2GetVersionResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OryKetoRelationTuplesV1alpha2GetVersionResponse) SetVersion(v string) {
	o.Version = &v
}

func (o OryKetoRelationTuplesV1alpha2GetVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableOryKetoRelationTuplesV1alpha2GetVersionResponse struct {
	value *OryKetoRelationTuplesV1alpha2GetVersionResponse
	isSet bool
}

func (v NullableOryKetoRelationTuplesV1alpha2GetVersionResponse) Get() *OryKetoRelationTuplesV1alpha2GetVersionResponse {
	return v.value
}

func (v *NullableOryKetoRelationTuplesV1alpha2GetVersionResponse) Set(val *OryKetoRelationTuplesV1alpha2GetVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOryKetoRelationTuplesV1alpha2GetVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOryKetoRelationTuplesV1alpha2GetVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOryKetoRelationTuplesV1alpha2GetVersionResponse(val *OryKetoRelationTuplesV1alpha2GetVersionResponse) *NullableOryKetoRelationTuplesV1alpha2GetVersionResponse {
	return &NullableOryKetoRelationTuplesV1alpha2GetVersionResponse{value: val, isSet: true}
}

func (v NullableOryKetoRelationTuplesV1alpha2GetVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOryKetoRelationTuplesV1alpha2GetVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}