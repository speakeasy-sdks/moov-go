// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/types"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

type AuthorizationSpendLimitControl struct {
	// Maximum value in cents allowed per duration
	Amount *int64 `json:"amount,omitempty"`
	// Unit of authorization limit control
	duration *string `const:"transaction" json:"duration,omitempty"`
}

func (a AuthorizationSpendLimitControl) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthorizationSpendLimitControl) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthorizationSpendLimitControl) GetAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AuthorizationSpendLimitControl) GetDuration() *string {
	return types.String("transaction")
}
