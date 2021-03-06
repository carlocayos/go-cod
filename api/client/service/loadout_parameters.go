// Code generated by go-swagger; DO NOT EDIT.

package service

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
)

// NewLoadoutParams creates a new LoadoutParams object
// with the default values initialized.
func NewLoadoutParams() *LoadoutParams {
	var ()
	return &LoadoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadoutParamsWithTimeout creates a new LoadoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadoutParamsWithTimeout(timeout time.Duration) *LoadoutParams {
	var ()
	return &LoadoutParams{

		timeout: timeout,
	}
}

// NewLoadoutParamsWithContext creates a new LoadoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadoutParamsWithContext(ctx context.Context) *LoadoutParams {
	var ()
	return &LoadoutParams{

		Context: ctx,
	}
}

// NewLoadoutParamsWithHTTPClient creates a new LoadoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadoutParamsWithHTTPClient(client *http.Client) *LoadoutParams {
	var ()
	return &LoadoutParams{
		HTTPClient: client,
	}
}

/*LoadoutParams contains all the parameters to send to the API endpoint
for the loadout operation typically these are written to a http.Request
*/
type LoadoutParams struct {

	/*GameType
	  Game Type

	*/
	GameType string
	/*Title
	  Game title

	*/
	Title string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the loadout params
func (o *LoadoutParams) WithTimeout(timeout time.Duration) *LoadoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the loadout params
func (o *LoadoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the loadout params
func (o *LoadoutParams) WithContext(ctx context.Context) *LoadoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the loadout params
func (o *LoadoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the loadout params
func (o *LoadoutParams) WithHTTPClient(client *http.Client) *LoadoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the loadout params
func (o *LoadoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGameType adds the gameType to the loadout params
func (o *LoadoutParams) WithGameType(gameType string) *LoadoutParams {
	o.SetGameType(gameType)
	return o
}

// SetGameType adds the gameType to the loadout params
func (o *LoadoutParams) SetGameType(gameType string) {
	o.GameType = gameType
}

// WithTitle adds the title to the loadout params
func (o *LoadoutParams) WithTitle(title string) *LoadoutParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the loadout params
func (o *LoadoutParams) SetTitle(title string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *LoadoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gameType
	if err := r.SetPathParam("gameType", o.GameType); err != nil {
		return err
	}

	// path param title
	if err := r.SetPathParam("title", o.Title); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
