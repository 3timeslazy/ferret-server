// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExecutionCommon Execution Common
//
// The properties that are shared amongst all versions of the Execution model.
// swagger:model execution-common
type ExecutionCommon struct {

	// Execution Cause
	//
	// Execution cause
	// Required: true
	// Enum: [unknown manual schedule hook]
	Cause *string `json:"cause"`

	// job id
	// Required: true
	// Pattern: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	JobID *string `json:"job_id"`

	// script id
	// Required: true
	// Pattern: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	ScriptID *string `json:"script_id"`

	// script rev
	// Required: true
	ScriptRev *string `json:"script_rev"`

	// Execution Status
	//
	// Execution stats
	// Required: true
	// Enum: [unknown queued running completed cancelled errored]
	Status *string `json:"status"`
}

// Validate validates this execution common
func (m *ExecutionCommon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptRev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var executionCommonTypeCausePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","manual","schedule","hook"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		executionCommonTypeCausePropEnum = append(executionCommonTypeCausePropEnum, v)
	}
}

const (

	// ExecutionCommonCauseUnknown captures enum value "unknown"
	ExecutionCommonCauseUnknown string = "unknown"

	// ExecutionCommonCauseManual captures enum value "manual"
	ExecutionCommonCauseManual string = "manual"

	// ExecutionCommonCauseSchedule captures enum value "schedule"
	ExecutionCommonCauseSchedule string = "schedule"

	// ExecutionCommonCauseHook captures enum value "hook"
	ExecutionCommonCauseHook string = "hook"
)

// prop value enum
func (m *ExecutionCommon) validateCauseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, executionCommonTypeCausePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExecutionCommon) validateCause(formats strfmt.Registry) error {

	if err := validate.Required("cause", "body", m.Cause); err != nil {
		return err
	}

	// value enum
	if err := m.validateCauseEnum("cause", "body", *m.Cause); err != nil {
		return err
	}

	return nil
}

func (m *ExecutionCommon) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", m.JobID); err != nil {
		return err
	}

	if err := validate.Pattern("job_id", "body", string(*m.JobID), `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *ExecutionCommon) validateScriptID(formats strfmt.Registry) error {

	if err := validate.Required("script_id", "body", m.ScriptID); err != nil {
		return err
	}

	if err := validate.Pattern("script_id", "body", string(*m.ScriptID), `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *ExecutionCommon) validateScriptRev(formats strfmt.Registry) error {

	if err := validate.Required("script_rev", "body", m.ScriptRev); err != nil {
		return err
	}

	return nil
}

var executionCommonTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","queued","running","completed","cancelled","errored"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		executionCommonTypeStatusPropEnum = append(executionCommonTypeStatusPropEnum, v)
	}
}

const (

	// ExecutionCommonStatusUnknown captures enum value "unknown"
	ExecutionCommonStatusUnknown string = "unknown"

	// ExecutionCommonStatusQueued captures enum value "queued"
	ExecutionCommonStatusQueued string = "queued"

	// ExecutionCommonStatusRunning captures enum value "running"
	ExecutionCommonStatusRunning string = "running"

	// ExecutionCommonStatusCompleted captures enum value "completed"
	ExecutionCommonStatusCompleted string = "completed"

	// ExecutionCommonStatusCancelled captures enum value "cancelled"
	ExecutionCommonStatusCancelled string = "cancelled"

	// ExecutionCommonStatusErrored captures enum value "errored"
	ExecutionCommonStatusErrored string = "errored"
)

// prop value enum
func (m *ExecutionCommon) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, executionCommonTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExecutionCommon) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecutionCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecutionCommon) UnmarshalBinary(b []byte) error {
	var res ExecutionCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
