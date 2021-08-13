// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelUserActiveBanResponse model user active ban response
//
// swagger:model model.UserActiveBanResponse
type ModelUserActiveBanResponse struct {

	// ban
	// Required: true
	Ban *string `json:"Ban"`

	// ban Id
	// Required: true
	BanID *string `json:"BanId"`

	// end date
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"EndDate"`
}

// Validate validates this model user active ban response
func (m *ModelUserActiveBanResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelUserActiveBanResponse) validateBan(formats strfmt.Registry) error {

	if err := validate.Required("Ban", "body", m.Ban); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserActiveBanResponse) validateBanID(formats strfmt.Registry) error {

	if err := validate.Required("BanId", "body", m.BanID); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserActiveBanResponse) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("EndDate", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelUserActiveBanResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelUserActiveBanResponse) UnmarshalBinary(b []byte) error {
	var res ModelUserActiveBanResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}