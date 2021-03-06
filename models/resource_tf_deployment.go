// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	hal "github.com/getlunaform/lunaform/models/hal"
)

// ResourceTfDeployment resource tf deployment
// swagger:model resource-tf-deployment
type ResourceTfDeployment struct {

	// links
	Links *hal.HalRscLinks `json:"_links,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// status
	// Enum: [DEPLOYING SUCCESS FAIL]
	Status *string `json:"status,omitempty"`

	// workspace
	Workspace *string `json:"workspace,omitempty"`
}

// Validate validates this resource tf deployment
func (m *ResourceTfDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceTfDeployment) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var resourceTfDeploymentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEPLOYING","SUCCESS","FAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTfDeploymentTypeStatusPropEnum = append(resourceTfDeploymentTypeStatusPropEnum, v)
	}
}

const (

	// ResourceTfDeploymentStatusDEPLOYING captures enum value "DEPLOYING"
	ResourceTfDeploymentStatusDEPLOYING string = "DEPLOYING"

	// ResourceTfDeploymentStatusSUCCESS captures enum value "SUCCESS"
	ResourceTfDeploymentStatusSUCCESS string = "SUCCESS"

	// ResourceTfDeploymentStatusFAIL captures enum value "FAIL"
	ResourceTfDeploymentStatusFAIL string = "FAIL"
)

// prop value enum
func (m *ResourceTfDeployment) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceTfDeploymentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceTfDeployment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceTfDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceTfDeployment) UnmarshalBinary(b []byte) error {
	var res ResourceTfDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
