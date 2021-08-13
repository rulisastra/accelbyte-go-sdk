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

// AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4 account upgrade headless account with verification code request v4
//
// swagger:model account.upgradeHeadlessAccountWithVerificationCodeRequestV4
type AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4 struct {

	// code
	// Required: true
	Code *string `json:"code"`

	// country
	Country string `json:"country,omitempty"`

	// date of birth
	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// email address
	// Required: true
	EmailAddress *string `json:"emailAddress"`

	// password
	// Required: true
	Password *string `json:"password"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this account upgrade headless account with verification code request v4
func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("emailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4) UnmarshalBinary(b []byte) error {
	var res AccountUpgradeHeadlessAccountWithVerificationCodeRequestV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}