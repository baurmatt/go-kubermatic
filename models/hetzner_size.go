// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HetznerSize HetznerSize is the object representing Hetzner sizes.
//
// swagger:model HetznerSize
type HetznerSize struct {

	// cores
	Cores int64 `json:"cores,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// disk
	Disk int64 `json:"disk,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// memory
	Memory float32 `json:"memory,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hetzner size
func (m *HetznerSize) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HetznerSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HetznerSize) UnmarshalBinary(b []byte) error {
	var res HetznerSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
