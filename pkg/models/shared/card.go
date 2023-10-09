// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

// Card - Describes a card on a Moov account
type Card struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	BillingAddress       *Address               `json:"billingAddress,omitempty"`
	Bin                  *string                `json:"bin,omitempty"`
	// The card brand
	Brand *CardBrand `json:"brand,omitempty"`
	// The results of the most recent card update request
	CardAccountUpdater *CardAccountUpdater `json:"cardAccountUpdater,omitempty"`
	// UUID v4
	CardID *string `json:"cardID,omitempty"`
	// Indicates cardholder has authorized card to be stored for future payments
	CardOnFile *bool `default:"false" json:"cardOnFile"`
	// The type of the card
	CardType *CardType `json:"cardType,omitempty"`
	// The results of submitting cardholder data to a card network for verification
	CardVerification *CardVerifications `json:"cardVerification,omitempty"`
	// The expiration date of the linked card or token
	Expiration *CardExpiration `json:"expiration,omitempty"`
	// Uniquely identifies a linked payment card or token.
	// For Apple Pay, the fingerprint is based on the tokenized card number and may vary based on the user's device.
	// This field can be used to identify specific payment methods across multiple accounts on your platform.
	//
	Fingerprint        *string `json:"fingerprint,omitempty"`
	HolderName         *string `json:"holderName,omitempty"`
	Issuer             *string `json:"issuer,omitempty"`
	IssuerCountry      *string `json:"issuerCountry,omitempty"`
	LastFourCardNumber *string `json:"lastFourCardNumber,omitempty"`
	// Moov account ID of the merchant or entity authorized to store the card. Defaults to your platform account ID if cardOnFile is set to true and no other account is provided
	MerchantAccountID *string `json:"merchantAccountID,omitempty"`
}

func (c Card) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Card) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Card) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Card) GetBillingAddress() *Address {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *Card) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *Card) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *Card) GetCardAccountUpdater() *CardAccountUpdater {
	if o == nil {
		return nil
	}
	return o.CardAccountUpdater
}

func (o *Card) GetCardID() *string {
	if o == nil {
		return nil
	}
	return o.CardID
}

func (o *Card) GetCardOnFile() *bool {
	if o == nil {
		return nil
	}
	return o.CardOnFile
}

func (o *Card) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *Card) GetCardVerification() *CardVerifications {
	if o == nil {
		return nil
	}
	return o.CardVerification
}

func (o *Card) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *Card) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *Card) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *Card) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *Card) GetIssuerCountry() *string {
	if o == nil {
		return nil
	}
	return o.IssuerCountry
}

func (o *Card) GetLastFourCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourCardNumber
}

func (o *Card) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}
