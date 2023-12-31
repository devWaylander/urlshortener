// Code generated by go-swagger; DO NOT EDIT.

package url

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/devWhisper/urlshortener/gen/models"
)

// PutLongOKCode is the HTTP code returned for type PutLongOK
const PutLongOKCode int = 200

/*
PutLongOK OK

swagger:response putLongOK
*/
type PutLongOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutLongOK creates PutLongOK with default headers values
func NewPutLongOK() *PutLongOK {

	return &PutLongOK{}
}

// WithPayload adds the payload to the put long o k response
func (o *PutLongOK) WithPayload(payload string) *PutLongOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put long o k response
func (o *PutLongOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutLongOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutLongNotFoundCode is the HTTP code returned for type PutLongNotFound
const PutLongNotFoundCode int = 404

/*
PutLongNotFound Not found.

swagger:response putLongNotFound
*/
type PutLongNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorV1 `json:"body,omitempty"`
}

// NewPutLongNotFound creates PutLongNotFound with default headers values
func NewPutLongNotFound() *PutLongNotFound {

	return &PutLongNotFound{}
}

// WithPayload adds the payload to the put long not found response
func (o *PutLongNotFound) WithPayload(payload *models.ErrorV1) *PutLongNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put long not found response
func (o *PutLongNotFound) SetPayload(payload *models.ErrorV1) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutLongNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutLongServiceUnavailableCode is the HTTP code returned for type PutLongServiceUnavailable
const PutLongServiceUnavailableCode int = 503

/*
PutLongServiceUnavailable Returned if the service is detected as unhealthy

swagger:response putLongServiceUnavailable
*/
type PutLongServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPutLongServiceUnavailable creates PutLongServiceUnavailable with default headers values
func NewPutLongServiceUnavailable() *PutLongServiceUnavailable {

	return &PutLongServiceUnavailable{}
}

// WithPayload adds the payload to the put long service unavailable response
func (o *PutLongServiceUnavailable) WithPayload(payload interface{}) *PutLongServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put long service unavailable response
func (o *PutLongServiceUnavailable) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutLongServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
