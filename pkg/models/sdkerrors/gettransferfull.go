// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"time"
)

// DisputedAmount - A representation of money containing an integer value and it's currency.
type DisputedAmount struct {
	// A 3-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Quantity in the smallest unit of the specified currency. In USD this is cents, so $12.04 is 1204 and $0.99 would be 99.
	Value int64 `json:"value"`
}

func (o *DisputedAmount) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *DisputedAmount) GetValue() int64 {
	if o == nil {
		return 0
	}
	return o.Value
}

// RefundedAmount - A representation of money containing an integer value and it's currency.
type RefundedAmount struct {
	// A 3-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Quantity in the smallest unit of the specified currency. In USD this is cents, so $12.04 is 1204 and $0.99 would be 99.
	Value int64 `json:"value"`
}

func (o *RefundedAmount) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *RefundedAmount) GetValue() int64 {
	if o == nil {
		return 0
	}
	return o.Value
}

// GetTransferFull - Transfer details
type GetTransferFull struct {
	// A representation of money containing an integer value and it's currency.
	Amount    *Amount    `json:"amount,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// A description of the transfer
	Description *string                           `json:"description,omitempty"`
	Destination *GetTransferFullSourceDestination `json:"destination,omitempty"`
	// The total disputed amount for a card transfer
	DisputedAmount *DisputedAmount `json:"disputedAmount,omitempty"`
	// A list of disputes for a card transfer
	Disputes []GetDispute `json:"disputes,omitempty"`
	// Fee you charged your customer for the transfer
	FacilitatorFee *GetFacilitatorFee `json:"facilitatorFee,omitempty"`
	// Transfer failure reason
	FailureReason *FailureReason `json:"failureReason,omitempty"`
	GroupID       *string        `json:"groupID,omitempty"`
	// Free-form key-value pair list. Useful for storing information that is not captured elsewhere.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Fee charged to your platform account for card transfers
	MoovFee *int64 `json:"moovFee,omitempty"`
	// The total refunded amount for a card transfer
	RefundedAmount *RefundedAmount `json:"refundedAmount,omitempty"`
	// A list of refunds for a card transfer
	Refunds []GetRefund            `json:"refunds,omitempty"`
	Source  *GetTransferFullSource `json:"source,omitempty"`
	// Current status of a transfer
	Status *TransferStatus `json:"status,omitempty"`
	// UUID v4
	TransferID *string `json:"transferID,omitempty"`
}

var _ error = &GetTransferFull{}

func (e *GetTransferFull) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}