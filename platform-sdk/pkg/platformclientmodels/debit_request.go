// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DebitRequest A DTO for wallet's debit
//
// swagger:model DebitRequest
type DebitRequest struct {

	// amount
	// Required: true
	// Minimum: 1
	Amount *int64 `json:"amount"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this debit request
func (m *DebitRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebitRequest) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.MinimumInt("amount", "body", int64(*m.Amount), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebitRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebitRequest) UnmarshalBinary(b []byte) error {
	var res DebitRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}