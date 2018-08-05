// Code generated by go-swagger; DO NOT EDIT.

package providers

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

	models "github.com/getlunaform/lunaform/models"
)

// NewGetProviderConfigurationParams creates a new GetProviderConfigurationParams object
// with the default values initialized.
func NewGetProviderConfigurationParams() *GetProviderConfigurationParams {
	var ()
	return &GetProviderConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProviderConfigurationParamsWithTimeout creates a new GetProviderConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProviderConfigurationParamsWithTimeout(timeout time.Duration) *GetProviderConfigurationParams {
	var ()
	return &GetProviderConfigurationParams{

		timeout: timeout,
	}
}

// NewGetProviderConfigurationParamsWithContext creates a new GetProviderConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProviderConfigurationParamsWithContext(ctx context.Context) *GetProviderConfigurationParams {
	var ()
	return &GetProviderConfigurationParams{

		Context: ctx,
	}
}

// NewGetProviderConfigurationParamsWithHTTPClient creates a new GetProviderConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProviderConfigurationParamsWithHTTPClient(client *http.Client) *GetProviderConfigurationParams {
	var ()
	return &GetProviderConfigurationParams{
		HTTPClient: client,
	}
}

/*GetProviderConfigurationParams contains all the parameters to send to the API endpoint
for the get provider configuration operation typically these are written to a http.Request
*/
type GetProviderConfigurationParams struct {

	/*ID
	  Configuration for a Terraform Provider

	*/
	ID string
	/*ProviderName
	  Terraform Provider ID

	*/
	ProviderName string
	/*TerraformProvider
	  A terraform module

	*/
	TerraformProvider *models.ResourceTfProviderConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get provider configuration params
func (o *GetProviderConfigurationParams) WithTimeout(timeout time.Duration) *GetProviderConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get provider configuration params
func (o *GetProviderConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get provider configuration params
func (o *GetProviderConfigurationParams) WithContext(ctx context.Context) *GetProviderConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get provider configuration params
func (o *GetProviderConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get provider configuration params
func (o *GetProviderConfigurationParams) WithHTTPClient(client *http.Client) *GetProviderConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get provider configuration params
func (o *GetProviderConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get provider configuration params
func (o *GetProviderConfigurationParams) WithID(id string) *GetProviderConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get provider configuration params
func (o *GetProviderConfigurationParams) SetID(id string) {
	o.ID = id
}

// WithProviderName adds the providerName to the get provider configuration params
func (o *GetProviderConfigurationParams) WithProviderName(providerName string) *GetProviderConfigurationParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the get provider configuration params
func (o *GetProviderConfigurationParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WithTerraformProvider adds the terraformProvider to the get provider configuration params
func (o *GetProviderConfigurationParams) WithTerraformProvider(terraformProvider *models.ResourceTfProviderConfiguration) *GetProviderConfigurationParams {
	o.SetTerraformProvider(terraformProvider)
	return o
}

// SetTerraformProvider adds the terraformProvider to the get provider configuration params
func (o *GetProviderConfigurationParams) SetTerraformProvider(terraformProvider *models.ResourceTfProviderConfiguration) {
	o.TerraformProvider = terraformProvider
}

// WriteToRequest writes these params to a swagger request
func (o *GetProviderConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param provider-name
	if err := r.SetPathParam("provider-name", o.ProviderName); err != nil {
		return err
	}

	if o.TerraformProvider != nil {
		if err := r.SetBodyParam(o.TerraformProvider); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
