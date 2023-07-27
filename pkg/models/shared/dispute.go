// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Dispute - Details about card dispute
type Dispute struct {
	// A representation of money containing an integer value and it's currency.
	Amount    *Amount    `json:"amount,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// UUID v4
	DisputeID                *string    `json:"disputeID,omitempty"`
	NetworkReasonCode        *string    `json:"networkReasonCode,omitempty"`
	NetworkReasonDescription *string    `json:"networkReasonDescription,omitempty"`
	RespondBy                *time.Time `json:"respondBy,omitempty"`
	// Dispute status
	Status *DisputeStatus `json:"status,omitempty"`
	// Details about transfer
	Transfer *Transfer `json:"transfer,omitempty"`
}

func (o *Dispute) GetAmount() *Amount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Dispute) GetCreatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *Dispute) GetDisputeID() *string {
	if o == nil {
		return nil
	}
	return o.DisputeID
}

func (o *Dispute) GetNetworkReasonCode() *string {
	if o == nil {
		return nil
	}
	return o.NetworkReasonCode
}

func (o *Dispute) GetNetworkReasonDescription() *string {
	if o == nil {
		return nil
	}
	return o.NetworkReasonDescription
}

func (o *Dispute) GetRespondBy() *time.Time {
	if o == nil {
		return nil
	}
	return o.RespondBy
}

func (o *Dispute) GetStatus() *DisputeStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Dispute) GetTransfer() *Transfer {
	if o == nil {
		return nil
	}
	return o.Transfer
}