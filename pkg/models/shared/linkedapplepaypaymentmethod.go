// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LinkedApplePayPaymentMethod struct {
	// Describes an Apple Pay token on a Moov account.
	ApplePay ApplePayResponse `json:"applePay"`
	// ID of the payment method
	PaymentMethodID string `json:"paymentMethodID"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType PaymentMethodsType `json:"paymentMethodType"`
}

func (o *LinkedApplePayPaymentMethod) GetApplePay() ApplePayResponse {
	if o == nil {
		return ApplePayResponse{}
	}
	return o.ApplePay
}

func (o *LinkedApplePayPaymentMethod) GetPaymentMethodID() string {
	if o == nil {
		return ""
	}
	return o.PaymentMethodID
}

func (o *LinkedApplePayPaymentMethod) GetPaymentMethodType() PaymentMethodsType {
	if o == nil {
		return PaymentMethodsType("")
	}
	return o.PaymentMethodType
}
