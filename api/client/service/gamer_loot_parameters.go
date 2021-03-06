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

// NewGamerLootParams creates a new GamerLootParams object
// with the default values initialized.
func NewGamerLootParams() *GamerLootParams {
	var ()
	return &GamerLootParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGamerLootParamsWithTimeout creates a new GamerLootParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGamerLootParamsWithTimeout(timeout time.Duration) *GamerLootParams {
	var ()
	return &GamerLootParams{

		timeout: timeout,
	}
}

// NewGamerLootParamsWithContext creates a new GamerLootParams object
// with the default values initialized, and the ability to set a context for a request
func NewGamerLootParamsWithContext(ctx context.Context) *GamerLootParams {
	var ()
	return &GamerLootParams{

		Context: ctx,
	}
}

// NewGamerLootParamsWithHTTPClient creates a new GamerLootParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGamerLootParamsWithHTTPClient(client *http.Client) *GamerLootParams {
	var ()
	return &GamerLootParams{
		HTTPClient: client,
	}
}

/*GamerLootParams contains all the parameters to send to the API endpoint
for the gamer loot operation typically these are written to a http.Request
*/
type GamerLootParams struct {

	/*Cookie
	  ACT_SSO_COOKIE required for authenticated requests. Set with the value returned in login

	*/
	Cookie string
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

// WithTimeout adds the timeout to the gamer loot params
func (o *GamerLootParams) WithTimeout(timeout time.Duration) *GamerLootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gamer loot params
func (o *GamerLootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gamer loot params
func (o *GamerLootParams) WithContext(ctx context.Context) *GamerLootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gamer loot params
func (o *GamerLootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gamer loot params
func (o *GamerLootParams) WithHTTPClient(client *http.Client) *GamerLootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gamer loot params
func (o *GamerLootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the gamer loot params
func (o *GamerLootParams) WithCookie(cookie string) *GamerLootParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the gamer loot params
func (o *GamerLootParams) SetCookie(cookie string) {
	o.Cookie = cookie
}

// WithGamertag adds the gamertag to the gamer loot params
func (o *GamerLootParams) WithGamertag(gamertag string) *GamerLootParams {
	o.SetGamertag(gamertag)
	return o
}

// SetGamertag adds the gamertag to the gamer loot params
func (o *GamerLootParams) SetGamertag(gamertag string) {
	o.Gamertag = gamertag
}

// WithLookupType adds the lookupType to the gamer loot params
func (o *GamerLootParams) WithLookupType(lookupType string) *GamerLootParams {
	o.SetLookupType(lookupType)
	return o
}

// SetLookupType adds the lookupType to the gamer loot params
func (o *GamerLootParams) SetLookupType(lookupType string) {
	o.LookupType = lookupType
}

// WithPlatform adds the platform to the gamer loot params
func (o *GamerLootParams) WithPlatform(platform string) *GamerLootParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the gamer loot params
func (o *GamerLootParams) SetPlatform(platform string) {
	o.Platform = platform
}

// WithTitle adds the title to the gamer loot params
func (o *GamerLootParams) WithTitle(title string) *GamerLootParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the gamer loot params
func (o *GamerLootParams) SetTitle(title string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *GamerLootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Cookie
	if err := r.SetHeaderParam("Cookie", o.Cookie); err != nil {
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
