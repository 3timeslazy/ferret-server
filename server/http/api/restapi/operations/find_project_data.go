// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// FindProjectDataHandlerFunc turns a function with the right signature into a find project data handler
type FindProjectDataHandlerFunc func(FindProjectDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindProjectDataHandlerFunc) Handle(params FindProjectDataParams) middleware.Responder {
	return fn(params)
}

// FindProjectDataHandler interface for that can handle valid find project data params
type FindProjectDataHandler interface {
	Handle(FindProjectDataParams) middleware.Responder
}

// NewFindProjectData creates a new http.Handler for the find project data operation
func NewFindProjectData(ctx *middleware.Context, handler FindProjectDataHandler) *FindProjectData {
	return &FindProjectData{Context: ctx, Handler: handler}
}

/*FindProjectData swagger:route GET /projects/{projectID}/data findProjectData

List Data

*/
type FindProjectData struct {
	Context *middleware.Context
	Handler FindProjectDataHandler
}

func (o *FindProjectData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindProjectDataParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FindProjectDataOKBodyItems0 Data Output
//
// The properties that are included when fetching a list of Data.
// swagger:model FindProjectDataOKBodyItems0
type FindProjectDataOKBodyItems0 struct {
	FindProjectDataOKBodyItems0AllOf0

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
func (o *FindProjectDataOKBodyItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FindProjectDataOKBodyItems0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.FindProjectDataOKBodyItems0AllOf0 = aO0

	// AO1
	var dataAO1 struct {
		JobID *string `json:"job_id"`

		ScriptID *string `json:"script_id"`

		ScriptRev string `json:"script_rev,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.JobID = dataAO1.JobID

	o.ScriptID = dataAO1.ScriptID

	o.ScriptRev = dataAO1.ScriptRev

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o FindProjectDataOKBodyItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(o.FindProjectDataOKBodyItems0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		JobID *string `json:"job_id"`

		ScriptID *string `json:"script_id"`

		ScriptRev string `json:"script_rev,omitempty"`
	}

	dataAO1.JobID = o.JobID

	dataAO1.ScriptID = o.ScriptID

	dataAO1.ScriptRev = o.ScriptRev

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this find project data o k body items0
func (o *FindProjectDataOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FindProjectDataOKBodyItems0AllOf0
	if err := o.FindProjectDataOKBodyItems0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScriptID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindProjectDataOKBodyItems0) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", o.JobID); err != nil {
		return err
	}

	return nil
}

func (o *FindProjectDataOKBodyItems0) validateScriptID(formats strfmt.Registry) error {

	if err := validate.Required("script_id", "body", o.ScriptID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FindProjectDataOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindProjectDataOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res FindProjectDataOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// FindProjectDataOKBodyItems0AllOf0 Entity
//
// Represents a database entity
// swagger:model FindProjectDataOKBodyItems0AllOf0
type FindProjectDataOKBodyItems0AllOf0 struct {

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
func (o *FindProjectDataOKBodyItems0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	o.Rev = dataAO0.Rev

	// AO1
	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.CreatedAt = dataAO1.CreatedAt

	o.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o FindProjectDataOKBodyItems0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}

	dataAO0.ID = o.ID

	dataAO0.Rev = o.Rev

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}

	dataAO1.CreatedAt = o.CreatedAt

	dataAO1.UpdatedAt = o.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this find project data o k body items0 all of0
func (o *FindProjectDataOKBodyItems0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRev(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindProjectDataOKBodyItems0AllOf0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *FindProjectDataOKBodyItems0AllOf0) validateRev(formats strfmt.Registry) error {

	if err := validate.Required("rev", "body", o.Rev); err != nil {
		return err
	}

	return nil
}

func (o *FindProjectDataOKBodyItems0AllOf0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FindProjectDataOKBodyItems0AllOf0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindProjectDataOKBodyItems0AllOf0) UnmarshalBinary(b []byte) error {
	var res FindProjectDataOKBodyItems0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}