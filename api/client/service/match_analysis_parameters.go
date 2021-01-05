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
	"github.com/go-openapi/swag"
)

// NewMatchAnalysisParams creates a new MatchAnalysisParams object
// with the default values initialized.
func NewMatchAnalysisParams() *MatchAnalysisParams {
	var ()
	return &MatchAnalysisParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMatchAnalysisParamsWithTimeout creates a new MatchAnalysisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMatchAnalysisParamsWithTimeout(timeout time.Duration) *MatchAnalysisParams {
	var ()
	return &MatchAnalysisParams{

		timeout: timeout,
	}
}

// NewMatchAnalysisParamsWithContext creates a new MatchAnalysisParams object
// with the default values initialized, and the ability to set a context for a request
func NewMatchAnalysisParamsWithContext(ctx context.Context) *MatchAnalysisParams {
	var ()
	return &MatchAnalysisParams{

		Context: ctx,
	}
}

// NewMatchAnalysisParamsWithHTTPClient creates a new MatchAnalysisParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMatchAnalysisParamsWithHTTPClient(client *http.Client) *MatchAnalysisParams {
	var ()
	return &MatchAnalysisParams{
		HTTPClient: client,
	}
}

/*MatchAnalysisParams contains all the parameters to send to the API endpoint
for the match analysis operation typically these are written to a http.Request
*/
type MatchAnalysisParams struct {

	/*Cookie
	  ACT_SSO_COOKIE required for authenticated requests. Set with the value returned in login

	*/
	Cookie string
	/*End
	  Match end unix timestamp in milliseconds ex. 1608407321000

	*/
	End int64
	/*Gamertag
	  Gamer tag

	*/
	Gamertag string
	/*LookupType
	  Lookup Type

	*/
	LookupType string
	/*Platform
	  Game platform

	*/
	Platform string
	/*Title
	  Game title

	*/
	Title string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the match analysis params
func (o *MatchAnalysisParams) WithTimeout(timeout time.Duration) *MatchAnalysisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the match analysis params
func (o *MatchAnalysisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the match analysis params
func (o *MatchAnalysisParams) WithContext(ctx context.Context) *MatchAnalysisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the match analysis params
func (o *MatchAnalysisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the match analysis params
func (o *MatchAnalysisParams) WithHTTPClient(client *http.Client) *MatchAnalysisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the match analysis params
func (o *MatchAnalysisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the match analysis params
func (o *MatchAnalysisParams) WithCookie(cookie string) *MatchAnalysisParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the match analysis params
func (o *MatchAnalysisParams) SetCookie(cookie string) {
	o.Cookie = cookie
}

// WithEnd adds the end to the match analysis params
func (o *MatchAnalysisParams) WithEnd(end int64) *MatchAnalysisParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the match analysis params
func (o *MatchAnalysisParams) SetEnd(end int64) {
	o.End = end
}

// WithGamertag adds the gamertag to the match analysis params
func (o *MatchAnalysisParams) WithGamertag(gamertag string) *MatchAnalysisParams {
	o.SetGamertag(gamertag)
	return o
}

// SetGamertag adds the gamertag to the match analysis params
func (o *MatchAnalysisParams) SetGamertag(gamertag string) {
	o.Gamertag = gamertag
}

// WithLookupType adds the lookupType to the match analysis params
func (o *MatchAnalysisParams) WithLookupType(lookupType string) *MatchAnalysisParams {
	o.SetLookupType(lookupType)
	return o
}

// SetLookupType adds the lookupType to the match analysis params
func (o *MatchAnalysisParams) SetLookupType(lookupType string) {
	o.LookupType = lookupType
}

// WithPlatform adds the platform to the match analysis params
func (o *MatchAnalysisParams) WithPlatform(platform string) *MatchAnalysisParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the match analysis params
func (o *MatchAnalysisParams) SetPlatform(platform string) {
	o.Platform = platform
}

// WithTitle adds the title to the match analysis params
func (o *MatchAnalysisParams) WithTitle(title string) *MatchAnalysisParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the match analysis params
func (o *MatchAnalysisParams) SetTitle(title string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *MatchAnalysisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Cookie
	if err := r.SetHeaderParam("Cookie", o.Cookie); err != nil {
		return err
	}

	// path param end
	if err := r.SetPathParam("end", swag.FormatInt64(o.End)); err != nil {
		return err
	}

	// path param gamertag
	if err := r.SetPathParam("gamertag", o.Gamertag); err != nil {
		return err
	}

	// path param lookupType
	if err := r.SetPathParam("lookupType", o.LookupType); err != nil {
		return err
	}

	// path param platform
	if err := r.SetPathParam("platform", o.Platform); err != nil {
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
