package j_provisioner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJProvisionerUpdateParams creates a new JProvisionerUpdateParams object
// with the default values initialized.
func NewJProvisionerUpdateParams() *JProvisionerUpdateParams {
	var ()
	return &JProvisionerUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJProvisionerUpdateParamsWithTimeout creates a new JProvisionerUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJProvisionerUpdateParamsWithTimeout(timeout time.Duration) *JProvisionerUpdateParams {
	var ()
	return &JProvisionerUpdateParams{

		timeout: timeout,
	}
}

// NewJProvisionerUpdateParamsWithContext creates a new JProvisionerUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJProvisionerUpdateParamsWithContext(ctx context.Context) *JProvisionerUpdateParams {
	var ()
	return &JProvisionerUpdateParams{

		Context: ctx,
	}
}

/*JProvisionerUpdateParams contains all the parameters to send to the API endpoint
for the j provisioner update operation typically these are written to a http.Request
*/
type JProvisionerUpdateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j provisioner update params
func (o *JProvisionerUpdateParams) WithTimeout(timeout time.Duration) *JProvisionerUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j provisioner update params
func (o *JProvisionerUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j provisioner update params
func (o *JProvisionerUpdateParams) WithContext(ctx context.Context) *JProvisionerUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j provisioner update params
func (o *JProvisionerUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j provisioner update params
func (o *JProvisionerUpdateParams) WithBody(body models.DefaultSelector) *JProvisionerUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j provisioner update params
func (o *JProvisionerUpdateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j provisioner update params
func (o *JProvisionerUpdateParams) WithID(id string) *JProvisionerUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j provisioner update params
func (o *JProvisionerUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JProvisionerUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
