// Code generated by go-swagger; DO NOT EDIT.

package lobbyclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelNotificationWithTemplateRequest model notification with template request
//
// swagger:model model.NotificationWithTemplateRequest
type ModelNotificationWithTemplateRequest struct {

	// template context
	// Required: true
	TemplateContext map[string]string `json:"templateContext"`

	// template language
	// Required: true
	TemplateLanguage *string `json:"templateLanguage"`

	// template slug
	// Required: true
	TemplateSlug *string `json:"templateSlug"`

	// topic
	// Required: true
	Topic *string `json:"topic"`
}

// Validate validates this model notification with template request
func (m *ModelNotificationWithTemplateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelNotificationWithTemplateRequest) validateTemplateContext(formats strfmt.Registry) error {

	return nil
}

func (m *ModelNotificationWithTemplateRequest) validateTemplateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("templateLanguage", "body", m.TemplateLanguage); err != nil {
		return err
	}

	return nil
}

func (m *ModelNotificationWithTemplateRequest) validateTemplateSlug(formats strfmt.Registry) error {

	if err := validate.Required("templateSlug", "body", m.TemplateSlug); err != nil {
		return err
	}

	return nil
}

func (m *ModelNotificationWithTemplateRequest) validateTopic(formats strfmt.Registry) error {

	if err := validate.Required("topic", "body", m.Topic); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelNotificationWithTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelNotificationWithTemplateRequest) UnmarshalBinary(b []byte) error {
	var res ModelNotificationWithTemplateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}