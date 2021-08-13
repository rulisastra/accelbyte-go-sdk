// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountcommonRole accountcommon role
//
// swagger:model accountcommon.Role
type AccountcommonRole struct {

	// admin role
	// Required: true
	AdminRole *bool `json:"AdminRole"`

	// is wildcard
	// Required: true
	IsWildcard *bool `json:"IsWildcard"`

	// managers
	// Required: true
	Managers []*AccountcommonRoleManager `json:"Managers"`

	// members
	// Required: true
	Members []*AccountcommonRoleMember `json:"Members"`

	// permissions
	// Required: true
	Permissions []*AccountcommonPermission `json:"Permissions"`

	// role Id
	// Required: true
	RoleID *string `json:"RoleId"`

	// role name
	// Required: true
	RoleName *string `json:"RoleName"`
}

// Validate validates this accountcommon role
func (m *AccountcommonRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsWildcard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountcommonRole) validateAdminRole(formats strfmt.Registry) error {

	if err := validate.Required("AdminRole", "body", m.AdminRole); err != nil {
		return err
	}

	return nil
}

func (m *AccountcommonRole) validateIsWildcard(formats strfmt.Registry) error {

	if err := validate.Required("IsWildcard", "body", m.IsWildcard); err != nil {
		return err
	}

	return nil
}

func (m *AccountcommonRole) validateManagers(formats strfmt.Registry) error {

	if err := validate.Required("Managers", "body", m.Managers); err != nil {
		return err
	}

	for i := 0; i < len(m.Managers); i++ {
		if swag.IsZero(m.Managers[i]) { // not required
			continue
		}

		if m.Managers[i] != nil {
			if err := m.Managers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Managers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountcommonRole) validateMembers(formats strfmt.Registry) error {

	if err := validate.Required("Members", "body", m.Members); err != nil {
		return err
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountcommonRole) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("Permissions", "body", m.Permissions); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountcommonRole) validateRoleID(formats strfmt.Registry) error {

	if err := validate.Required("RoleId", "body", m.RoleID); err != nil {
		return err
	}

	return nil
}

func (m *AccountcommonRole) validateRoleName(formats strfmt.Registry) error {

	if err := validate.Required("RoleName", "body", m.RoleName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountcommonRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountcommonRole) UnmarshalBinary(b []byte) error {
	var res AccountcommonRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}