// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RefundCardStatus - Card status for refunds
type RefundCardStatus string

const (
	RefundCardStatusInitiated RefundCardStatus = "initiated"
	RefundCardStatusConfirmed RefundCardStatus = "confirmed"
	RefundCardStatusSettled   RefundCardStatus = "settled"
	RefundCardStatusFailed    RefundCardStatus = "failed"
	RefundCardStatusCompleted RefundCardStatus = "completed"
)

func (e RefundCardStatus) ToPointer() *RefundCardStatus {
	return &e
}

func (e *RefundCardStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "initiated":
		fallthrough
	case "confirmed":
		fallthrough
	case "settled":
		fallthrough
	case "failed":
		fallthrough
	case "completed":
		*e = RefundCardStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefundCardStatus: %v", v)
	}
}
