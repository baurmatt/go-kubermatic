// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AWSVpcCidrBlockAssociation AWSVpcCidrBlockAssociation describes an IPv4 CIDR block associated with a VPC.
// swagger:model AWSVpcCidrBlockAssociation
type AWSVpcCidrBlockAssociation struct {

	// The association ID for the IPv4 CIDR block.
	AssociationID string `json:"associationId,omitempty"`

	// The IPv4 CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// The state of the CIDR block.
	State string `json:"state,omitempty"`

	// A message about the status of the CIDR block, if applicable.
	StatusMessage string `json:"statusMessage,omitempty"`
}

// Validate validates this a w s vpc cidr block association
func (m *AWSVpcCidrBlockAssociation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AWSVpcCidrBlockAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSVpcCidrBlockAssociation) UnmarshalBinary(b []byte) error {
	var res AWSVpcCidrBlockAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
