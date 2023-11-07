// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
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

// SynchronousTransferResponse - Transfer details
type SynchronousTransferResponse struct {
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

func (s SynchronousTransferResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SynchronousTransferResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SynchronousTransferResponse) GetAmount() *Amount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *SynchronousTransferResponse) GetCreatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *SynchronousTransferResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SynchronousTransferResponse) GetDestination() *GetTransferFullSourceDestination {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *SynchronousTransferResponse) GetDisputedAmount() *DisputedAmount {
	if o == nil {
		return nil
	}
	return o.DisputedAmount
}

func (o *SynchronousTransferResponse) GetDisputes() []GetDispute {
	if o == nil {
		return nil
	}
	return o.Disputes
}

func (o *SynchronousTransferResponse) GetFacilitatorFee() *GetFacilitatorFee {
	if o == nil {
		return nil
	}
	return o.FacilitatorFee
}

func (o *SynchronousTransferResponse) GetFailureReason() *FailureReason {
	if o == nil {
		return nil
	}
	return o.FailureReason
}

func (o *SynchronousTransferResponse) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}

func (o *SynchronousTransferResponse) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *SynchronousTransferResponse) GetMoovFee() *int64 {
	if o == nil {
		return nil
	}
	return o.MoovFee
}

func (o *SynchronousTransferResponse) GetRefundedAmount() *RefundedAmount {
	if o == nil {
		return nil
	}
	return o.RefundedAmount
}

func (o *SynchronousTransferResponse) GetRefunds() []GetRefund {
	if o == nil {
		return nil
	}
	return o.Refunds
}

func (o *SynchronousTransferResponse) GetSource() *GetTransferFullSource {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *SynchronousTransferResponse) GetStatus() *TransferStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SynchronousTransferResponse) GetTransferID() *string {
	if o == nil {
		return nil
	}
	return o.TransferID
}

type TransferPostResponseType string

const (
	TransferPostResponseTypeCreatedTransfer             TransferPostResponseType = "CreatedTransfer"
	TransferPostResponseTypeSynchronousTransferResponse TransferPostResponseType = "Synchronous transfer response"
)

type TransferPostResponse struct {
	CreatedTransfer             *CreatedTransfer
	SynchronousTransferResponse *SynchronousTransferResponse

	Type TransferPostResponseType
}

func CreateTransferPostResponseCreatedTransfer(createdTransfer CreatedTransfer) TransferPostResponse {
	typ := TransferPostResponseTypeCreatedTransfer

	return TransferPostResponse{
		CreatedTransfer: &createdTransfer,
		Type:            typ,
	}
}

func CreateTransferPostResponseSynchronousTransferResponse(synchronousTransferResponse SynchronousTransferResponse) TransferPostResponse {
	typ := TransferPostResponseTypeSynchronousTransferResponse

	return TransferPostResponse{
		SynchronousTransferResponse: &synchronousTransferResponse,
		Type:                        typ,
	}
}

func (u *TransferPostResponse) UnmarshalJSON(data []byte) error {

	createdTransfer := CreatedTransfer{}
	if err := utils.UnmarshalJSON(data, &createdTransfer, "", true, true); err == nil {
		u.CreatedTransfer = &createdTransfer
		u.Type = TransferPostResponseTypeCreatedTransfer
		return nil
	}

	synchronousTransferResponse := SynchronousTransferResponse{}
	if err := utils.UnmarshalJSON(data, &synchronousTransferResponse, "", true, true); err == nil {
		u.SynchronousTransferResponse = &synchronousTransferResponse
		u.Type = TransferPostResponseTypeSynchronousTransferResponse
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TransferPostResponse) MarshalJSON() ([]byte, error) {
	if u.CreatedTransfer != nil {
		return utils.MarshalJSON(u.CreatedTransfer, "", true)
	}

	if u.SynchronousTransferResponse != nil {
		return utils.MarshalJSON(u.SynchronousTransferResponse, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
