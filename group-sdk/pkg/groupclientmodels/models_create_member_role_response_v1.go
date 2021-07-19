// Code generated by go-swagger; DO NOT EDIT.

package groupclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsCreateMemberRoleResponseV1 models create member role response v1
//
// swagger:model models.CreateMemberRoleResponseV1
type ModelsCreateMemberRoleResponseV1 struct {

	// member role Id
	// Required: true
	MemberRoleID *string `json:"memberRoleId"`

	// member role name
	// Required: true
	MemberRoleName *string `json:"memberRoleName"`

	// member role permissions
	// Required: true
	MemberRolePermissions []*ModelsRolePermission `json:"memberRolePermissions"`
}

// Validate validates this models create member role response v1
func (m *ModelsCreateMemberRoleResponseV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemberRoleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberRoleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberRolePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCreateMemberRoleResponseV1) validateMemberRoleID(formats strfmt.Registry) error {

	if err := validate.Required("memberRoleId", "body", m.MemberRoleID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateMemberRoleResponseV1) validateMemberRoleName(formats strfmt.Registry) error {

	if err := validate.Required("memberRoleName", "body", m.MemberRoleName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateMemberRoleResponseV1) validateMemberRolePermissions(formats strfmt.Registry) error {

	if err := validate.Required("memberRolePermissions", "body", m.MemberRolePermissions); err != nil {
		return err
	}

	for i := 0; i < len(m.MemberRolePermissions); i++ {
		if swag.IsZero(m.MemberRolePermissions[i]) { // not required
			continue
		}

		if m.MemberRolePermissions[i] != nil {
			if err := m.MemberRolePermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberRolePermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCreateMemberRoleResponseV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCreateMemberRoleResponseV1) UnmarshalBinary(b []byte) error {
	var res ModelsCreateMemberRoleResponseV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}