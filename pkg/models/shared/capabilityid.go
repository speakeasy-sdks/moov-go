// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CapabilityID - ID of Capability
type CapabilityID string

const (
	CapabilityIDTransfers    CapabilityID = "transfers"
	CapabilityIDSendFunds    CapabilityID = "send-funds"
	CapabilityIDCollectFunds CapabilityID = "collect-funds"
	CapabilityIDWallet       CapabilityID = "wallet"
	CapabilityIDCardIssuing  CapabilityID = "card-issuing"
)

func (e CapabilityID) ToPointer() *CapabilityID {
	return &e
}

func (e *CapabilityID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "transfers":
		fallthrough
	case "send-funds":
		fallthrough
	case "collect-funds":
		fallthrough
	case "wallet":
		fallthrough
	case "card-issuing":
		*e = CapabilityID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CapabilityID: %v", v)
	}
}
