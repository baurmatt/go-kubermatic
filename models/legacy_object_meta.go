// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LegacyObjectMeta LegacyObjectMeta is an object storing common metadata for persistable objects.
//
// Deprecated: LegacyObjectMeta is deprecated use ObjectMeta instead.
//
// swagger:model LegacyObjectMeta
type LegacyObjectMeta struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource version
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID
	UID string `json:"uid,omitempty"`
}

// Validate validates this legacy object meta
func (m *LegacyObjectMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LegacyObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyObjectMeta) UnmarshalBinary(b []byte) error {
	var res LegacyObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
