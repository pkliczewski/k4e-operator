// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/jakub-dzon/k4e-operator/models"
)

// NewPostControlMessageForDeviceParams creates a new PostControlMessageForDeviceParams object
// with the default values initialized.
func NewPostControlMessageForDeviceParams() *PostControlMessageForDeviceParams {
	var ()
	return &PostControlMessageForDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostControlMessageForDeviceParamsWithTimeout creates a new PostControlMessageForDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostControlMessageForDeviceParamsWithTimeout(timeout time.Duration) *PostControlMessageForDeviceParams {
	var ()
	return &PostControlMessageForDeviceParams{

		timeout: timeout,
	}
}

// NewPostControlMessageForDeviceParamsWithContext creates a new PostControlMessageForDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostControlMessageForDeviceParamsWithContext(ctx context.Context) *PostControlMessageForDeviceParams {
	var ()
	return &PostControlMessageForDeviceParams{

		Context: ctx,
	}
}

// NewPostControlMessageForDeviceParamsWithHTTPClient creates a new PostControlMessageForDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostControlMessageForDeviceParamsWithHTTPClient(client *http.Client) *PostControlMessageForDeviceParams {
	var ()
	return &PostControlMessageForDeviceParams{
		HTTPClient: client,
	}
}

/*PostControlMessageForDeviceParams contains all the parameters to send to the API endpoint
for the post control message for device operation typically these are written to a http.Request
*/
type PostControlMessageForDeviceParams struct {

	/*DeviceID
	  Device ID

	*/
	DeviceID string
	/*Message*/
	Message *models.Message

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post control message for device params
func (o *PostControlMessageForDeviceParams) WithTimeout(timeout time.Duration) *PostControlMessageForDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post control message for device params
func (o *PostControlMessageForDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post control message for device params
func (o *PostControlMessageForDeviceParams) WithContext(ctx context.Context) *PostControlMessageForDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post control message for device params
func (o *PostControlMessageForDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post control message for device params
func (o *PostControlMessageForDeviceParams) WithHTTPClient(client *http.Client) *PostControlMessageForDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post control message for device params
func (o *PostControlMessageForDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the post control message for device params
func (o *PostControlMessageForDeviceParams) WithDeviceID(deviceID string) *PostControlMessageForDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post control message for device params
func (o *PostControlMessageForDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithMessage adds the message to the post control message for device params
func (o *PostControlMessageForDeviceParams) WithMessage(message *models.Message) *PostControlMessageForDeviceParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the post control message for device params
func (o *PostControlMessageForDeviceParams) SetMessage(message *models.Message) {
	o.Message = message
}

// WriteToRequest writes these params to a swagger request
func (o *PostControlMessageForDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	if o.Message != nil {
		if err := r.SetBodyParam(o.Message); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
