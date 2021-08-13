// Code generated by go-swagger; DO NOT EDIT.

package lobbyclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsAdminGetProfanityListFiltersV1Response models admin get profanity list filters v1 response
//
// swagger:model models.AdminGetProfanityListFiltersV1Response
type ModelsAdminGetProfanityListFiltersV1Response struct {

	// filters
	// Required: true
	Filters []*ModelsProfanityFilter `json:"filters"`
}

// Validate validates this models admin get profanity list filters v1 response
func (m *ModelsAdminGetProfanityListFiltersV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAdminGetProfanityListFiltersV1Response) validateFilters(formats strfmt.Registry) error {

	if err := validate.Required("filters", "body", m.Filters); err != nil {
		return err
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAdminGetProfanityListFiltersV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAdminGetProfanityListFiltersV1Response) UnmarshalBinary(b []byte) error {
	var res ModelsAdminGetProfanityListFiltersV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}