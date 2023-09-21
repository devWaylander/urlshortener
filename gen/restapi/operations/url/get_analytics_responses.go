// Code generated by go-swagger; DO NOT EDIT.

package url

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/devWhisper/urlshortener/gen/models"
)

// GetAnalyticsOKCode is the HTTP code returned for type GetAnalyticsOK
const GetAnalyticsOKCode int = 200

/*
GetAnalyticsOK OK

swagger:response getAnalyticsOK
*/
type GetAnalyticsOK struct {

	/*
	  In: Body
	*/
	Payload *GetAnalyticsOKBody `json:"body,omitempty"`
}

// NewGetAnalyticsOK creates GetAnalyticsOK with default headers values
func NewGetAnalyticsOK() *GetAnalyticsOK {

	return &GetAnalyticsOK{}
}

// WithPayload adds the payload to the get analytics o k response
func (o *GetAnalyticsOK) WithPayload(payload *GetAnalyticsOKBody) *GetAnalyticsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics o k response
func (o *GetAnalyticsOK) SetPayload(payload *GetAnalyticsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnalyticsNotFoundCode is the HTTP code returned for type GetAnalyticsNotFound
const GetAnalyticsNotFoundCode int = 404

/*
GetAnalyticsNotFound Not found.

swagger:response getAnalyticsNotFound
*/
type GetAnalyticsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorV1 `json:"body,omitempty"`
}

// NewGetAnalyticsNotFound creates GetAnalyticsNotFound with default headers values
func NewGetAnalyticsNotFound() *GetAnalyticsNotFound {

	return &GetAnalyticsNotFound{}
}

// WithPayload adds the payload to the get analytics not found response
func (o *GetAnalyticsNotFound) WithPayload(payload *models.ErrorV1) *GetAnalyticsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics not found response
func (o *GetAnalyticsNotFound) SetPayload(payload *models.ErrorV1) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnalyticsServiceUnavailableCode is the HTTP code returned for type GetAnalyticsServiceUnavailable
const GetAnalyticsServiceUnavailableCode int = 503

/*
GetAnalyticsServiceUnavailable Returned if the service is detected as unhealthy

swagger:response getAnalyticsServiceUnavailable
*/
type GetAnalyticsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAnalyticsServiceUnavailable creates GetAnalyticsServiceUnavailable with default headers values
func NewGetAnalyticsServiceUnavailable() *GetAnalyticsServiceUnavailable {

	return &GetAnalyticsServiceUnavailable{}
}

// WithPayload adds the payload to the get analytics service unavailable response
func (o *GetAnalyticsServiceUnavailable) WithPayload(payload interface{}) *GetAnalyticsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics service unavailable response
func (o *GetAnalyticsServiceUnavailable) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}