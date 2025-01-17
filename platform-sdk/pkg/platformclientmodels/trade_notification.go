// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TradeNotification trade notification
//
// swagger:model TradeNotification
type TradeNotification struct {

	// additional data
	AdditionalData *AdditionalData `json:"additionalData,omitempty"`

	// The time of the order authorised
	// Format: date-time
	AuthorisedTime strfmt.DateTime `json:"authorisedTime,omitempty"`

	// The time of the order chargeback reversed
	// Format: date-time
	ChargebackReversedTime strfmt.DateTime `json:"chargebackReversedTime,omitempty"`

	// The time of the order chargeback
	// Format: date-time
	ChargebackTime strfmt.DateTime `json:"chargebackTime,omitempty"`

	// The time of the order charged
	// Format: date-time
	ChargedTime strfmt.DateTime `json:"chargedTime,omitempty"`

	// The time of the order created
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// Payment order currency info
	// Required: true
	Currency *CurrencySummary `json:"currency"`

	// User custom parameters
	CustomParameters map[string]interface{} `json:"customParameters,omitempty"`

	// Order number
	// Required: true
	ExtOrderNo *string `json:"extOrderNo"`

	// External transaction id
	ExtTxID string `json:"extTxId,omitempty"`

	// optional, external user id, can be the character id
	ExtUserID string `json:"extUserId,omitempty"`

	// event issued at
	// Required: true
	// Format: date-time
	IssuedAt *strfmt.DateTime `json:"issuedAt"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// namespace which event issues from
	// Required: true
	Namespace *string `json:"namespace"`

	// Random string
	// Required: true
	NonceStr *string `json:"nonceStr"`

	// Payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// Payment method fee
	PaymentMethodFee int32 `json:"paymentMethodFee,omitempty"`

	// Payment order number
	// Required: true
	PaymentOrderNo *string `json:"paymentOrderNo"`

	// Payment provider
	// Required: true
	// Enum: [WALLET XSOLLA ADYEN STRIPE CHECKOUT ALIPAY WXPAY PAYPAL]
	PaymentProvider *string `json:"paymentProvider"`

	// Payment provider fee
	PaymentProviderFee int32 `json:"paymentProviderFee,omitempty"`

	// Payment station url
	PaymentStationURL string `json:"paymentStationUrl,omitempty"`

	// Payment order price
	// Required: true
	Price *int32 `json:"price"`

	// The time of the order refunded
	// Format: date-time
	RefundedTime strfmt.DateTime `json:"refundedTime,omitempty"`

	// Payment order sales tax
	SalesTax int32 `json:"salesTax,omitempty"`

	// isSandbox, indicate if order is sandbox
	// Required: true
	Sandbox *bool `json:"sandbox"`

	// optional, unique identity for the item
	Sku string `json:"sku,omitempty"`

	// Payment order status
	// Required: true
	// Enum: [INIT AUTHORISED AUTHORISE_FAILED CHARGED CHARGE_FAILED NOTIFICATION_OF_CHARGEBACK REQUEST_FOR_INFORMATION CHARGEBACK CHARGEBACK_REVERSED REFUNDING REFUNDED REFUND_FAILED DELETED]
	Status *string `json:"status"`

	// Payment order status reason
	StatusReason string `json:"statusReason,omitempty"`

	// Subscription id if exist
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// subtotal price
	SubtotalPrice int32 `json:"subtotalPrice,omitempty"`

	// target namespace, usually it's the game namespace
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// target user id, usually it's the user id in game namespace
	TargetUserID string `json:"targetUserId,omitempty"`

	// Payment total tax
	Tax int32 `json:"tax,omitempty"`

	// total price
	TotalPrice int32 `json:"totalPrice,omitempty"`

	// total tax
	TotalTax int32 `json:"totalTax,omitempty"`

	// Transaction end date time
	// Format: date-time
	TxEndTime strfmt.DateTime `json:"txEndTime,omitempty"`

	// notification type: distribution, payment
	// Required: true
	Type *string `json:"type"`

	// user id in namespace, will be null if there's targetUserId
	UserID string `json:"userId,omitempty"`

	// Payment order VAT
	Vat int32 `json:"vat,omitempty"`
}

// Validate validates this trade notification
func (m *TradeNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorisedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargebackReversedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargebackTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtOrderNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonceStr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentOrderNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefundedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSandbox(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TradeNotification) validateAdditionalData(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalData) { // not required
		return nil
	}

	if m.AdditionalData != nil {
		if err := m.AdditionalData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additionalData")
			}
			return err
		}
	}

	return nil
}

func (m *TradeNotification) validateAuthorisedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorisedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("authorisedTime", "body", "date-time", m.AuthorisedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateChargebackReversedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargebackReversedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("chargebackReversedTime", "body", "date-time", m.ChargebackReversedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateChargebackTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargebackTime) { // not required
		return nil
	}

	if err := validate.FormatOf("chargebackTime", "body", "date-time", m.ChargebackTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateChargedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("chargedTime", "body", "date-time", m.ChargedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *TradeNotification) validateExtOrderNo(formats strfmt.Registry) error {

	if err := validate.Required("extOrderNo", "body", m.ExtOrderNo); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateIssuedAt(formats strfmt.Registry) error {

	if err := validate.Required("issuedAt", "body", m.IssuedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("issuedAt", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateNonceStr(formats strfmt.Registry) error {

	if err := validate.Required("nonceStr", "body", m.NonceStr); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validatePaymentOrderNo(formats strfmt.Registry) error {

	if err := validate.Required("paymentOrderNo", "body", m.PaymentOrderNo); err != nil {
		return err
	}

	return nil
}

var tradeNotificationTypePaymentProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WALLET","XSOLLA","ADYEN","STRIPE","CHECKOUT","ALIPAY","WXPAY","PAYPAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tradeNotificationTypePaymentProviderPropEnum = append(tradeNotificationTypePaymentProviderPropEnum, v)
	}
}

const (

	// TradeNotificationPaymentProviderWALLET captures enum value "WALLET"
	TradeNotificationPaymentProviderWALLET string = "WALLET"

	// TradeNotificationPaymentProviderXSOLLA captures enum value "XSOLLA"
	TradeNotificationPaymentProviderXSOLLA string = "XSOLLA"

	// TradeNotificationPaymentProviderADYEN captures enum value "ADYEN"
	TradeNotificationPaymentProviderADYEN string = "ADYEN"

	// TradeNotificationPaymentProviderSTRIPE captures enum value "STRIPE"
	TradeNotificationPaymentProviderSTRIPE string = "STRIPE"

	// TradeNotificationPaymentProviderCHECKOUT captures enum value "CHECKOUT"
	TradeNotificationPaymentProviderCHECKOUT string = "CHECKOUT"

	// TradeNotificationPaymentProviderALIPAY captures enum value "ALIPAY"
	TradeNotificationPaymentProviderALIPAY string = "ALIPAY"

	// TradeNotificationPaymentProviderWXPAY captures enum value "WXPAY"
	TradeNotificationPaymentProviderWXPAY string = "WXPAY"

	// TradeNotificationPaymentProviderPAYPAL captures enum value "PAYPAL"
	TradeNotificationPaymentProviderPAYPAL string = "PAYPAL"
)

// prop value enum
func (m *TradeNotification) validatePaymentProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tradeNotificationTypePaymentProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TradeNotification) validatePaymentProvider(formats strfmt.Registry) error {

	if err := validate.Required("paymentProvider", "body", m.PaymentProvider); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentProviderEnum("paymentProvider", "body", *m.PaymentProvider); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateRefundedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RefundedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("refundedTime", "body", "date-time", m.RefundedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateSandbox(formats strfmt.Registry) error {

	if err := validate.Required("sandbox", "body", m.Sandbox); err != nil {
		return err
	}

	return nil
}

var tradeNotificationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INIT","AUTHORISED","AUTHORISE_FAILED","CHARGED","CHARGE_FAILED","NOTIFICATION_OF_CHARGEBACK","REQUEST_FOR_INFORMATION","CHARGEBACK","CHARGEBACK_REVERSED","REFUNDING","REFUNDED","REFUND_FAILED","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tradeNotificationTypeStatusPropEnum = append(tradeNotificationTypeStatusPropEnum, v)
	}
}

const (

	// TradeNotificationStatusINIT captures enum value "INIT"
	TradeNotificationStatusINIT string = "INIT"

	// TradeNotificationStatusAUTHORISED captures enum value "AUTHORISED"
	TradeNotificationStatusAUTHORISED string = "AUTHORISED"

	// TradeNotificationStatusAUTHORISEFAILED captures enum value "AUTHORISE_FAILED"
	TradeNotificationStatusAUTHORISEFAILED string = "AUTHORISE_FAILED"

	// TradeNotificationStatusCHARGED captures enum value "CHARGED"
	TradeNotificationStatusCHARGED string = "CHARGED"

	// TradeNotificationStatusCHARGEFAILED captures enum value "CHARGE_FAILED"
	TradeNotificationStatusCHARGEFAILED string = "CHARGE_FAILED"

	// TradeNotificationStatusNOTIFICATIONOFCHARGEBACK captures enum value "NOTIFICATION_OF_CHARGEBACK"
	TradeNotificationStatusNOTIFICATIONOFCHARGEBACK string = "NOTIFICATION_OF_CHARGEBACK"

	// TradeNotificationStatusREQUESTFORINFORMATION captures enum value "REQUEST_FOR_INFORMATION"
	TradeNotificationStatusREQUESTFORINFORMATION string = "REQUEST_FOR_INFORMATION"

	// TradeNotificationStatusCHARGEBACK captures enum value "CHARGEBACK"
	TradeNotificationStatusCHARGEBACK string = "CHARGEBACK"

	// TradeNotificationStatusCHARGEBACKREVERSED captures enum value "CHARGEBACK_REVERSED"
	TradeNotificationStatusCHARGEBACKREVERSED string = "CHARGEBACK_REVERSED"

	// TradeNotificationStatusREFUNDING captures enum value "REFUNDING"
	TradeNotificationStatusREFUNDING string = "REFUNDING"

	// TradeNotificationStatusREFUNDED captures enum value "REFUNDED"
	TradeNotificationStatusREFUNDED string = "REFUNDED"

	// TradeNotificationStatusREFUNDFAILED captures enum value "REFUND_FAILED"
	TradeNotificationStatusREFUNDFAILED string = "REFUND_FAILED"

	// TradeNotificationStatusDELETED captures enum value "DELETED"
	TradeNotificationStatusDELETED string = "DELETED"
)

// prop value enum
func (m *TradeNotification) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tradeNotificationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TradeNotification) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateTxEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.TxEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("txEndTime", "body", "date-time", m.TxEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TradeNotification) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TradeNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TradeNotification) UnmarshalBinary(b []byte) error {
	var res TradeNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
