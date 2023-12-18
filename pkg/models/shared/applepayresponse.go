// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ApplePayResponse - Describes an Apple Pay token on a Moov account.
type ApplePayResponse struct {
	// The card brand
	Brand *CardBrand `json:"brand,omitempty"`
	// User-friendly name of the tokenized card returned by Apple.
	// It usually contains the brand and the last four digits of the underlying card for example, "Visa 1256".
	// There is no standard format.
	//
	CardDisplayName *string `json:"cardDisplayName,omitempty"`
	// The type of the card
	CardType *CardType `json:"cardType,omitempty"`
	// The last four digits of the Apple Pay token, which may differ from the tokenized card's last four digits
	DynamicLastFour *string `json:"dynamicLastFour,omitempty"`
	// The expiration date of the linked card or token
	Expiration *CardExpiration `json:"expiration,omitempty"`
	// Uniquely identifies a linked payment card or token.
	// For Apple Pay, the fingerprint is based on the tokenized card number and may vary based on the user's device.
	// This field can be used to identify specific payment methods across multiple accounts on your platform.
	//
	Fingerprint *string `json:"fingerprint,omitempty"`
}

func (o *ApplePayResponse) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *ApplePayResponse) GetCardDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.CardDisplayName
}

func (o *ApplePayResponse) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *ApplePayResponse) GetDynamicLastFour() *string {
	if o == nil {
		return nil
	}
	return o.DynamicLastFour
}

func (o *ApplePayResponse) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *ApplePayResponse) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}
