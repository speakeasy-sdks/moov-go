// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// GetDispute - Details of a card dispute
type GetDispute struct {
	// A representation of money containing an integer value and it's currency.
	Amount    *Amount    `json:"amount,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// UUID v4
	DisputeID *string `json:"disputeID,omitempty"`
}

func (o *GetDispute) GetAmount() *Amount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *GetDispute) GetCreatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *GetDispute) GetDisputeID() *string {
	if o == nil {
		return nil
	}
	return o.DisputeID
}
