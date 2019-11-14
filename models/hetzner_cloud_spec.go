// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HetznerCloudSpec HetznerCloudSpec specifies access data to hetzner cloud.
// swagger:model HetznerCloudSpec
type HetznerCloudSpec struct {

	// token
	Token string `json:"token,omitempty"`

	// credentials reference
	CredentialsReference GlobalSecretKeySelector `json:"credentialsReference,omitempty"`
}

// Validate validates this hetzner cloud spec
func (m *HetznerCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HetznerCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsReference) { // not required
		return nil
	}

	if err := m.CredentialsReference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentialsReference")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HetznerCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HetznerCloudSpec) UnmarshalBinary(b []byte) error {
	var res HetznerCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
