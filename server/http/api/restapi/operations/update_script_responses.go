// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateScriptOKCode is the HTTP code returned for type UpdateScriptOK
const UpdateScriptOKCode int = 200

/*UpdateScriptOK update script o k

swagger:response updateScriptOK
*/
type UpdateScriptOK struct {

	/*
	  In: Body
	*/
	Payload *UpdateScriptOKBody `json:"body,omitempty"`
}

// NewUpdateScriptOK creates UpdateScriptOK with default headers values
func NewUpdateScriptOK() *UpdateScriptOK {

	return &UpdateScriptOK{}
}

// WithPayload adds the payload to the update script o k response
func (o *UpdateScriptOK) WithPayload(payload *UpdateScriptOKBody) *UpdateScriptOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update script o k response
func (o *UpdateScriptOK) SetPayload(payload *UpdateScriptOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateScriptOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}