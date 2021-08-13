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

// ModelVerificationCodeResponse model verification code response
//
// swagger:model model.VerificationCodeResponse
type ModelVerificationCodeResponse struct {

	// account registration
	// Required: true
	AccountRegistration *string `json:"accountRegistration"`

	// account upgrade
	// Required: true
	AccountUpgrade *string `json:"accountUpgrade"`

	// password reset
	// Required: true
	PasswordReset *string `json:"passwordReset"`

	// update email
	// Required: true
	UpdateEmail *string `json:"updateEmail"`
}

// Validate validates this model verification code response
func (m *ModelVerificationCodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountRegistration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountUpgrade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordReset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelVerificationCodeResponse) validateAccountRegistration(formats strfmt.Registry) error {

	if err := validate.Required("accountRegistration", "body", m.AccountRegistration); err != nil {
		return err
	}

	return nil
}

func (m *ModelVerificationCodeResponse) validateAccountUpgrade(formats strfmt.Registry) error {

	if err := validate.Required("accountUpgrade", "body", m.AccountUpgrade); err != nil {
		return err
	}

	return nil
}

func (m *ModelVerificationCodeResponse) validatePasswordReset(formats strfmt.Registry) error {

	if err := validate.Required("passwordReset", "body", m.PasswordReset); err != nil {
		return err
	}

	return nil
}

func (m *ModelVerificationCodeResponse) validateUpdateEmail(formats strfmt.Registry) error {

	if err := validate.Required("updateEmail", "body", m.UpdateEmail); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelVerificationCodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelVerificationCodeResponse) UnmarshalBinary(b []byte) error {
	var res ModelVerificationCodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}