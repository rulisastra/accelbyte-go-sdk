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

// ModelUserBanResponseV3 model user ban response v3
//
// swagger:model model.UserBanResponseV3
type ModelUserBanResponseV3 struct {

	// ban
	// Required: true
	Ban *string `json:"ban"`

	// ban Id
	// Required: true
	BanID *string `json:"banId"`

	// banned by
	// Required: true
	BannedBy *AccountcommonBannedByV3 `json:"bannedBy"`

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// disabled date
	// Required: true
	// Format: date-time
	DisabledDate *strfmt.DateTime `json:"disabledDate"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// end date
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// reason
	// Required: true
	Reason *string `json:"reason"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this model user ban response v3
func (m *ModelUserBanResponseV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBannedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabledDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelUserBanResponseV3) validateBan(formats strfmt.Registry) error {

	if err := validate.Required("ban", "body", m.Ban); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateBanID(formats strfmt.Registry) error {

	if err := validate.Required("banId", "body", m.BanID); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateBannedBy(formats strfmt.Registry) error {

	if err := validate.Required("bannedBy", "body", m.BannedBy); err != nil {
		return err
	}

	if m.BannedBy != nil {
		if err := m.BannedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bannedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateDisabledDate(formats strfmt.Registry) error {

	if err := validate.Required("disabledDate", "body", m.DisabledDate); err != nil {
		return err
	}

	if err := validate.FormatOf("disabledDate", "body", "date-time", m.DisabledDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserBanResponseV3) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelUserBanResponseV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelUserBanResponseV3) UnmarshalBinary(b []byte) error {
	var res ModelUserBanResponseV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
