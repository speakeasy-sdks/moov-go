// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CardUpdateRequest - Describes properties of a card to update
type CardUpdateRequest struct {
	// Provide address fields as necessary to patch the currently saved address.
	// Omit any fields that should not be changed.
	//
	BillingAddress *UpdateAddress `json:"billingAddress,omitempty"`
	// Provide a CVV to trigger a re-verification of this card.
	// Omit CVV to update LinkedCard fields without re-verification.
	//
	CardCvv *string `json:"cardCvv,omitempty"`
	// Indicates cardholder has authorized card to be stored for future payments
	CardOnFile *bool `json:"cardOnFile,omitempty"`
	// Provide expiration date fields as necessary to patch the currently saved date.
	// Omit any fields that should not be changed.
	//
	Expiration *UpdateCardExpiration `json:"expiration,omitempty"`
}

func (o *CardUpdateRequest) GetBillingAddress() *UpdateAddress {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *CardUpdateRequest) GetCardCvv() *string {
	if o == nil {
		return nil
	}
	return o.CardCvv
}

func (o *CardUpdateRequest) GetCardOnFile() *bool {
	if o == nil {
		return nil
	}
	return o.CardOnFile
}

func (o *CardUpdateRequest) GetExpiration() *UpdateCardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}