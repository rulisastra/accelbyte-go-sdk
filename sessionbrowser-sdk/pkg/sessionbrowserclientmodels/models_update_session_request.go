// Code generated by go-swagger; DO NOT EDIT.

package sessionbrowserclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsUpdateSessionRequest models update session request
//
// swagger:model models.UpdateSessionRequest
type ModelsUpdateSessionRequest struct {

	// game current player
	// Required: true
	GameCurrentPlayer *int32 `json:"game_current_player"`

	// game max player
	// Required: true
	GameMaxPlayer *int32 `json:"game_max_player"`
}

// Validate validates this models update session request
func (m *ModelsUpdateSessionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameCurrentPlayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameMaxPlayer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsUpdateSessionRequest) validateGameCurrentPlayer(formats strfmt.Registry) error {

	if err := validate.Required("game_current_player", "body", m.GameCurrentPlayer); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdateSessionRequest) validateGameMaxPlayer(formats strfmt.Registry) error {

	if err := validate.Required("game_max_player", "body", m.GameMaxPlayer); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsUpdateSessionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsUpdateSessionRequest) UnmarshalBinary(b []byte) error {
	var res ModelsUpdateSessionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}