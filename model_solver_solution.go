/*
api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package UpStudyGoSdk

import (
	"encoding/json"
)

// checks if the SolverSolution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolverSolution{}

// SolverSolution struct for SolverSolution
type SolverSolution struct {
	// Block name: e.g., Function/Solve the equation
	BlockName *SolverDescription `json:"block_name,omitempty"`
	// Results, possibly multiple, e.g., {\\frac{1}{4},0.25}
	Results []SolverStep `json:"results,omitempty"`
	// Specific solution name under the block classification: e.g., Block: Function, Solution: Find the slope
	SolutionName *SolverDescription `json:"solution_name,omitempty"`
}

// NewSolverSolution instantiates a new SolverSolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolverSolution() *SolverSolution {
	this := SolverSolution{}
	return &this
}

// NewSolverSolutionWithDefaults instantiates a new SolverSolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolverSolutionWithDefaults() *SolverSolution {
	this := SolverSolution{}
	return &this
}

// GetBlockName returns the BlockName field value if set, zero value otherwise.
func (o *SolverSolution) GetBlockName() SolverDescription {
	if o == nil || IsNil(o.BlockName) {
		var ret SolverDescription
		return ret
	}
	return *o.BlockName
}

// GetBlockNameOk returns a tuple with the BlockName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverSolution) GetBlockNameOk() (*SolverDescription, bool) {
	if o == nil || IsNil(o.BlockName) {
		return nil, false
	}
	return o.BlockName, true
}

// HasBlockName returns a boolean if a field has been set.
func (o *SolverSolution) HasBlockName() bool {
	if o != nil && !IsNil(o.BlockName) {
		return true
	}

	return false
}

// SetBlockName gets a reference to the given SolverDescription and assigns it to the BlockName field.
func (o *SolverSolution) SetBlockName(v SolverDescription) {
	o.BlockName = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SolverSolution) GetResults() []SolverStep {
	if o == nil || IsNil(o.Results) {
		var ret []SolverStep
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverSolution) GetResultsOk() ([]SolverStep, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SolverSolution) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SolverStep and assigns it to the Results field.
func (o *SolverSolution) SetResults(v []SolverStep) {
	o.Results = v
}

// GetSolutionName returns the SolutionName field value if set, zero value otherwise.
func (o *SolverSolution) GetSolutionName() SolverDescription {
	if o == nil || IsNil(o.SolutionName) {
		var ret SolverDescription
		return ret
	}
	return *o.SolutionName
}

// GetSolutionNameOk returns a tuple with the SolutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverSolution) GetSolutionNameOk() (*SolverDescription, bool) {
	if o == nil || IsNil(o.SolutionName) {
		return nil, false
	}
	return o.SolutionName, true
}

// HasSolutionName returns a boolean if a field has been set.
func (o *SolverSolution) HasSolutionName() bool {
	if o != nil && !IsNil(o.SolutionName) {
		return true
	}

	return false
}

// SetSolutionName gets a reference to the given SolverDescription and assigns it to the SolutionName field.
func (o *SolverSolution) SetSolutionName(v SolverDescription) {
	o.SolutionName = &v
}

func (o SolverSolution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolverSolution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockName) {
		toSerialize["block_name"] = o.BlockName
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.SolutionName) {
		toSerialize["solution_name"] = o.SolutionName
	}
	return toSerialize, nil
}

type NullableSolverSolution struct {
	value *SolverSolution
	isSet bool
}

func (v NullableSolverSolution) Get() *SolverSolution {
	return v.value
}

func (v *NullableSolverSolution) Set(val *SolverSolution) {
	v.value = val
	v.isSet = true
}

func (v NullableSolverSolution) IsSet() bool {
	return v.isSet
}

func (v *NullableSolverSolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolverSolution(val *SolverSolution) *NullableSolverSolution {
	return &NullableSolverSolution{value: val, isSet: true}
}

func (v NullableSolverSolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolverSolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


