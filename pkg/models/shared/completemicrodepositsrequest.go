// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

// CompleteMicroDepositsRequest - Request to complete the micro-deposit verification workflow.
type CompleteMicroDepositsRequest struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Two positive integers, in cents, equal to the values of the micro-deposits sent to the bank account.
	Amounts []int64 `json:"amounts,omitempty"`
}

func (c CompleteMicroDepositsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CompleteMicroDepositsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CompleteMicroDepositsRequest) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *CompleteMicroDepositsRequest) GetAmounts() []int64 {
	if o == nil {
		return nil
	}
	return o.Amounts
}
