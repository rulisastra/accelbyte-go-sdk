// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppLocalization app localization
//
// swagger:model AppLocalization
type AppLocalization struct {

	// announcement
	Announcement string `json:"announcement,omitempty"`

	// slogan
	Slogan string `json:"slogan,omitempty"`
}

// Validate validates this app localization
func (m *AppLocalization) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppLocalization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppLocalization) UnmarshalBinary(b []byte) error {
	var res AppLocalization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}