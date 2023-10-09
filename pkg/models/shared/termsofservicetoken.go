// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

// TermsOfServiceToken - An encrypted value used to record acceptance of Moov's Terms of Service
type TermsOfServiceToken struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Token                *string                `json:"token,omitempty"`
}

func (t TermsOfServiceToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TermsOfServiceToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TermsOfServiceToken) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *TermsOfServiceToken) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
