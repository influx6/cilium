// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFqdnCacheParams creates a new DeleteFqdnCacheParams object
// with the default values initialized.
func NewDeleteFqdnCacheParams() DeleteFqdnCacheParams {
	var ()
	return DeleteFqdnCacheParams{}
}

// DeleteFqdnCacheParams contains all the bound params for the delete fqdn cache operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteFqdnCache
type DeleteFqdnCacheParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*A toFQDNs compatible matchPattern expression
	  In: query
	*/
	Matchpattern *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteFqdnCacheParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMatchpattern, qhkMatchpattern, _ := qs.GetOK("matchpattern")
	if err := o.bindMatchpattern(qMatchpattern, qhkMatchpattern, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteFqdnCacheParams) bindMatchpattern(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Matchpattern = &raw

	return nil
}