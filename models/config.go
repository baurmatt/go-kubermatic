// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Config Config holds the information needed to build connect to remote kubernetes clusters as a given user
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
// swagger:model Config
type Config struct {

	// Legacy field from pkg/api/types.go TypeMeta.
	// TODO(jlowdermilk): remove this after eliminating downstream dependencies.
	// +k8s:conversion-gen=false
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// AuthInfos is a map of referencable names to user configs
	AuthInfos []*NamedAuthInfo `json:"users"`

	// Clusters is a map of referencable names to cluster configs
	Clusters []*NamedCluster `json:"clusters"`

	// Contexts is a map of referencable names to context configs
	Contexts []*NamedContext `json:"contexts"`

	// CurrentContext is the name of the context that you would like to use by default
	CurrentContext string `json:"current-context,omitempty"`

	// Extensions holds additional information. This is useful for extenders so that reads and writes don't clobber unknown fields
	// +optional
	Extensions []*NamedExtension `json:"extensions"`

	// Legacy field from pkg/api/types.go TypeMeta.
	// TODO(jlowdermilk): remove this after eliminating downstream dependencies.
	// +k8s:conversion-gen=false
	// +optional
	Kind string `json:"kind,omitempty"`

	// preferences
	Preferences *Preferences `json:"preferences,omitempty"`
}

// Validate validates this config
func (m *Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContexts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Config) validateAuthInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthInfos); i++ {
		if swag.IsZero(m.AuthInfos[i]) { // not required
			continue
		}

		if m.AuthInfos[i] != nil {
			if err := m.AuthInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Config) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Config) validateContexts(formats strfmt.Registry) error {

	if swag.IsZero(m.Contexts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contexts); i++ {
		if swag.IsZero(m.Contexts[i]) { // not required
			continue
		}

		if m.Contexts[i] != nil {
			if err := m.Contexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Config) validateExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Config) validatePreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	if m.Preferences != nil {
		if err := m.Preferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Config) UnmarshalBinary(b []byte) error {
	var res Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
