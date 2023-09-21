// Code generated by go-swagger; DO NOT EDIT.

package url

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetLongParams creates a new GetLongParams object
//
// There are no default values defined in the spec.
func NewGetLongParams() GetLongParams {

	return GetLongParams{}
}

// GetLongParams contains all the bound params for the get long operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLong
type GetLongParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Token string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLongParams() beforehand.
func (o *GetLongParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rToken, rhkToken, _ := route.Params.GetOK("token")
	if err := o.bindToken(rToken, rhkToken, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindToken binds and validates parameter Token from path.
func (o *GetLongParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Token = raw

	return nil
}
