// Code generated by go-swagger; DO NOT EDIT.

package socialclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BulkStatItemUpdate bulk stat item update
//
// swagger:model BulkStatItemUpdate
type BulkStatItemUpdate struct {

	// Additional data to be published in event payload
	AdditionalData map[string]interface{} `json:"additionalData,omitempty"`

	// stat code
	// Required: true
	StatCode *string `json:"statCode"`

	// update strategy
	// Required: true
	// Enum: [OVERRIDE INCREMENT MAX MIN]
	UpdateStrategy *string `json:"updateStrategy"`

	// value
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this bulk stat item update
func (m *BulkStatItemUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkStatItemUpdate) validateStatCode(formats strfmt.Registry) error {

	if err := validate.Required("statCode", "body", m.StatCode); err != nil {
		return err
	}

	return nil
}

var bulkStatItemUpdateTypeUpdateStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OVERRIDE","INCREMENT","MAX","MIN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bulkStatItemUpdateTypeUpdateStrategyPropEnum = append(bulkStatItemUpdateTypeUpdateStrategyPropEnum, v)
	}
}

const (

	// BulkStatItemUpdateUpdateStrategyOVERRIDE captures enum value "OVERRIDE"
	BulkStatItemUpdateUpdateStrategyOVERRIDE string = "OVERRIDE"

	// BulkStatItemUpdateUpdateStrategyINCREMENT captures enum value "INCREMENT"
	BulkStatItemUpdateUpdateStrategyINCREMENT string = "INCREMENT"

	// BulkStatItemUpdateUpdateStrategyMAX captures enum value "MAX"
	BulkStatItemUpdateUpdateStrategyMAX string = "MAX"

	// BulkStatItemUpdateUpdateStrategyMIN captures enum value "MIN"
	BulkStatItemUpdateUpdateStrategyMIN string = "MIN"
)

// prop value enum
func (m *BulkStatItemUpdate) validateUpdateStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bulkStatItemUpdateTypeUpdateStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BulkStatItemUpdate) validateUpdateStrategy(formats strfmt.Registry) error {

	if err := validate.Required("updateStrategy", "body", m.UpdateStrategy); err != nil {
		return err
	}

	// value enum
	if err := m.validateUpdateStrategyEnum("updateStrategy", "body", *m.UpdateStrategy); err != nil {
		return err
	}

	return nil
}

func (m *BulkStatItemUpdate) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkStatItemUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkStatItemUpdate) UnmarshalBinary(b []byte) error {
	var res BulkStatItemUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}