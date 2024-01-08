// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

type ApplePayMerchantDomains struct {
	// ID of Account
	AccountID *string    `json:"accountID,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// A string of 64 or fewer UTF-8 characters. Displayed in the Buy button.
	//
	DisplayName *string `json:"displayName,omitempty"`
	// A list of fully qualified top-level or sub-domain names where you will accept Apple Pay.
	//
	Domains   []string   `json:"domains,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

func (a ApplePayMerchantDomains) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplePayMerchantDomains) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplePayMerchantDomains) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *ApplePayMerchantDomains) GetCreatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *ApplePayMerchantDomains) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ApplePayMerchantDomains) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *ApplePayMerchantDomains) GetUpdatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedOn
}
