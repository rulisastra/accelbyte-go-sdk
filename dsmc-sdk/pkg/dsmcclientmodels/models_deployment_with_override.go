// Code generated by go-swagger; DO NOT EDIT.

package dsmcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsDeploymentWithOverride models deployment with override
//
// swagger:model models.DeploymentWithOverride
type ModelsDeploymentWithOverride struct {

	// allow version override
	// Required: true
	AllowVersionOverride *bool `json:"allow_version_override"`

	// buffer count
	// Required: true
	BufferCount *int32 `json:"buffer_count"`

	// configuration
	// Required: true
	Configuration *string `json:"configuration"`

	// game version
	// Required: true
	GameVersion *string `json:"game_version"`

	// max count
	// Required: true
	MaxCount *int32 `json:"max_count"`

	// min count
	// Required: true
	MinCount *int32 `json:"min_count"`

	// overrides
	// Required: true
	Overrides map[string]ModelsDeploymentConfig `json:"overrides"`

	// regions
	// Required: true
	Regions []string `json:"regions"`
}

// Validate validates this models deployment with override
func (m *ModelsDeploymentWithOverride) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowVersionOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBufferCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsDeploymentWithOverride) validateAllowVersionOverride(formats strfmt.Registry) error {

	if err := validate.Required("allow_version_override", "body", m.AllowVersionOverride); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateBufferCount(formats strfmt.Registry) error {

	if err := validate.Required("buffer_count", "body", m.BufferCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateGameVersion(formats strfmt.Registry) error {

	if err := validate.Required("game_version", "body", m.GameVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateMaxCount(formats strfmt.Registry) error {

	if err := validate.Required("max_count", "body", m.MaxCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateMinCount(formats strfmt.Registry) error {

	if err := validate.Required("min_count", "body", m.MinCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateOverrides(formats strfmt.Registry) error {

	for k := range m.Overrides {

		if err := validate.Required("overrides"+"."+k, "body", m.Overrides[k]); err != nil {
			return err
		}
		if val, ok := m.Overrides[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ModelsDeploymentWithOverride) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsDeploymentWithOverride) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsDeploymentWithOverride) UnmarshalBinary(b []byte) error {
	var res ModelsDeploymentWithOverride
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
