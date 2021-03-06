// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataOutput Data Output
//
// The properties that are included when fetching a list of Data.
// swagger:model data-output
type DataOutput struct {
	DataOutputAllOf0

	// job id
	// Required: true
	JobID *string `json:"job_id"`

	// script id
	// Required: true
	ScriptID *string `json:"script_id"`

	// script rev
	ScriptRev string `json:"script_rev,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DataOutput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DataOutputAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DataOutputAllOf0 = aO0

	// AO1
	var dataAO1 struct {
		JobID *string `json:"job_id"`

		ScriptID *string `json:"script_id"`

		ScriptRev string `json:"script_rev,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.JobID = dataAO1.JobID

	m.ScriptID = dataAO1.ScriptID

	m.ScriptRev = dataAO1.ScriptRev

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DataOutput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.DataOutputAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		JobID *string `json:"job_id"`

		ScriptID *string `json:"script_id"`

		ScriptRev string `json:"script_rev,omitempty"`
	}

	dataAO1.JobID = m.JobID

	dataAO1.ScriptID = m.ScriptID

	dataAO1.ScriptRev = m.ScriptRev

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data output
func (m *DataOutput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DataOutputAllOf0
	if err := m.DataOutputAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataOutput) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", m.JobID); err != nil {
		return err
	}

	return nil
}

func (m *DataOutput) validateScriptID(formats strfmt.Registry) error {

	if err := validate.Required("script_id", "body", m.ScriptID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataOutput) UnmarshalBinary(b []byte) error {
	var res DataOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataOutputAllOf0 Entity
//
// Represents a database entity
// swagger:model DataOutputAllOf0
type DataOutputAllOf0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// rev
	// Required: true
	Rev *string `json:"rev"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DataOutputAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ID = dataAO0.ID

	m.Rev = dataAO0.Rev

	// AO1
	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CreatedAt = dataAO1.CreatedAt

	m.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DataOutputAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}

	dataAO0.ID = m.ID

	dataAO0.Rev = m.Rev

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}

	dataAO1.CreatedAt = m.CreatedAt

	dataAO1.UpdatedAt = m.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data output all of0
func (m *DataOutputAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataOutputAllOf0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DataOutputAllOf0) validateRev(formats strfmt.Registry) error {

	if err := validate.Required("rev", "body", m.Rev); err != nil {
		return err
	}

	return nil
}

func (m *DataOutputAllOf0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataOutputAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataOutputAllOf0) UnmarshalBinary(b []byte) error {
	var res DataOutputAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
