// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RoleUser RoleUser defines associated user with role
// swagger:model RoleUser
type RoleUser struct {

	// group
	Group string `json:"group,omitempty"`

	// user email
	UserEmail string `json:"userEmail,omitempty"`
}

// Validate validates this role user
func (m *RoleUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleUser) UnmarshalBinary(b []byte) error {
	var res RoleUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
