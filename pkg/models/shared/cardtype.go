// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardType - The type of the card
type CardType string

const (
	CardTypeDebit   CardType = "debit"
	CardTypeCredit  CardType = "credit"
	CardTypePrepaid CardType = "prepaid"
	CardTypeUnknown CardType = "unknown"
)

func (e CardType) ToPointer() *CardType {
	return &e
}

func (e *CardType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debit":
		fallthrough
	case "credit":
		fallthrough
	case "prepaid":
		fallthrough
	case "unknown":
		*e = CardType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardType: %v", v)
	}
}