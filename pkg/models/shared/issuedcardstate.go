// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// IssuedCardState - State of a Moov issued card
type IssuedCardState string

const (
	IssuedCardStateActive              IssuedCardState = "active"
	IssuedCardStateInactive            IssuedCardState = "inactive"
	IssuedCardStatePendingVerification IssuedCardState = "pending-verification"
	IssuedCardStateClosed              IssuedCardState = "closed"
)

func (e IssuedCardState) ToPointer() *IssuedCardState {
	return &e
}

func (e *IssuedCardState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		fallthrough
	case "pending-verification":
		fallthrough
	case "closed":
		*e = IssuedCardState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IssuedCardState: %v", v)
	}
}
