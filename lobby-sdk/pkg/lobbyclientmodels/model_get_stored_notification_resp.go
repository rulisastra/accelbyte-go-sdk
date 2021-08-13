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

// ModelGetStoredNotificationResp model get stored notification resp
//
// swagger:model model.GetStoredNotificationResp
type ModelGetStoredNotificationResp struct {

	// from
	// Required: true
	From *string `json:"from"`

	// id
	// Required: true
	ID *string `json:"id"`

	// payload
	// Required: true
	Payload *string `json:"payload"`

	// sent at
	// Required: true
	SentAt *int64 `json:"sentAt"`

	// to
	// Required: true
	To *string `json:"to"`

	// topic name
	// Required: true
	TopicName *string `json:"topicName"`
}

// Validate validates this model get stored notification resp
func (m *ModelGetStoredNotificationResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSentAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopicName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelGetStoredNotificationResp) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *ModelGetStoredNotificationResp) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ModelGetStoredNotificationResp) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *ModelGetStoredNotificationResp) validateSentAt(formats strfmt.Registry) error {

	if err := validate.Required("sentAt", "body", m.SentAt); err != nil {
		return err
	}

	return nil
}

func (m *ModelGetStoredNotificationResp) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

func (m *ModelGetStoredNotificationResp) validateTopicName(formats strfmt.Registry) error {

	if err := validate.Required("topicName", "body", m.TopicName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelGetStoredNotificationResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelGetStoredNotificationResp) UnmarshalBinary(b []byte) error {
	var res ModelGetStoredNotificationResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
