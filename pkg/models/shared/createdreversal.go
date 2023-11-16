// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

// CreatedReversalFailureCode - This field is deprecated and will be removed in December 2023.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type CreatedReversalFailureCode string

const (
	CreatedReversalFailureCodeCallIssuer              CreatedReversalFailureCode = "call-issuer"
	CreatedReversalFailureCodeDoNotHonor              CreatedReversalFailureCode = "do-not-honor"
	CreatedReversalFailureCodeProcessingError         CreatedReversalFailureCode = "processing-error"
	CreatedReversalFailureCodeInvalidTransaction      CreatedReversalFailureCode = "invalid-transaction"
	CreatedReversalFailureCodeInvalidAmount           CreatedReversalFailureCode = "invalid-amount"
	CreatedReversalFailureCodeNoSuchIssuer            CreatedReversalFailureCode = "no-such-issuer"
	CreatedReversalFailureCodeReenterTransaction      CreatedReversalFailureCode = "reenter-transaction"
	CreatedReversalFailureCodeCvvMismatch             CreatedReversalFailureCode = "cvv-mismatch"
	CreatedReversalFailureCodeLostOrStolen            CreatedReversalFailureCode = "lost-or-stolen"
	CreatedReversalFailureCodeInsufficientFunds       CreatedReversalFailureCode = "insufficient-funds"
	CreatedReversalFailureCodeInvalidCardNumber       CreatedReversalFailureCode = "invalid-card-number"
	CreatedReversalFailureCodeExpiredCard             CreatedReversalFailureCode = "expired-card"
	CreatedReversalFailureCodeIncorrectPin            CreatedReversalFailureCode = "incorrect-pin"
	CreatedReversalFailureCodeTransactionNotAllowed   CreatedReversalFailureCode = "transaction-not-allowed"
	CreatedReversalFailureCodeSuspectedFraud          CreatedReversalFailureCode = "suspected-fraud"
	CreatedReversalFailureCodeAmountLimitExceeded     CreatedReversalFailureCode = "amount-limit-exceeded"
	CreatedReversalFailureCodeVelocityLimitExceeded   CreatedReversalFailureCode = "velocity-limit-exceeded"
	CreatedReversalFailureCodeCardNotActivated        CreatedReversalFailureCode = "card-not-activated"
	CreatedReversalFailureCodeIssuerNotAvailable      CreatedReversalFailureCode = "issuer-not-available"
	CreatedReversalFailureCodeCouldNotRoute           CreatedReversalFailureCode = "could-not-route"
	CreatedReversalFailureCodeCardholderAccountClosed CreatedReversalFailureCode = "cardholder-account-closed"
	CreatedReversalFailureCodeUnknownIssue            CreatedReversalFailureCode = "unknown-issue"
	CreatedReversalFailureCodeDuplicateTransaction    CreatedReversalFailureCode = "duplicate-transaction"
)

func (e CreatedReversalFailureCode) ToPointer() *CreatedReversalFailureCode {
	return &e
}

func (e *CreatedReversalFailureCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "call-issuer":
		fallthrough
	case "do-not-honor":
		fallthrough
	case "processing-error":
		fallthrough
	case "invalid-transaction":
		fallthrough
	case "invalid-amount":
		fallthrough
	case "no-such-issuer":
		fallthrough
	case "reenter-transaction":
		fallthrough
	case "cvv-mismatch":
		fallthrough
	case "lost-or-stolen":
		fallthrough
	case "insufficient-funds":
		fallthrough
	case "invalid-card-number":
		fallthrough
	case "expired-card":
		fallthrough
	case "incorrect-pin":
		fallthrough
	case "transaction-not-allowed":
		fallthrough
	case "suspected-fraud":
		fallthrough
	case "amount-limit-exceeded":
		fallthrough
	case "velocity-limit-exceeded":
		fallthrough
	case "card-not-activated":
		fallthrough
	case "issuer-not-available":
		fallthrough
	case "could-not-route":
		fallthrough
	case "cardholder-account-closed":
		fallthrough
	case "unknown-issue":
		fallthrough
	case "duplicate-transaction":
		*e = CreatedReversalFailureCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatedReversalFailureCode: %v", v)
	}
}

// SynchronousRefundResponse - Details of a card refund
type SynchronousRefundResponse struct {
	// A representation of money containing an integer value and it's currency.
	Amount      *Amount            `json:"amount,omitempty"`
	CardDetails *RefundCardDetails `json:"cardDetails,omitempty"`
	CreatedOn   *time.Time         `json:"createdOn,omitempty"`
	// This field is deprecated and will be removed in December 2023.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	FailureCode *CreatedReversalFailureCode `json:"failureCode,omitempty"`
	// UUID v4
	RefundID  *string       `json:"refundID,omitempty"`
	Status    *RefundStatus `json:"status,omitempty"`
	UpdatedOn *time.Time    `json:"updatedOn,omitempty"`
}

func (s SynchronousRefundResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SynchronousRefundResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SynchronousRefundResponse) GetAmount() *Amount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *SynchronousRefundResponse) GetCardDetails() *RefundCardDetails {
	if o == nil {
		return nil
	}
	return o.CardDetails
}

func (o *SynchronousRefundResponse) GetCreatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *SynchronousRefundResponse) GetFailureCode() *CreatedReversalFailureCode {
	if o == nil {
		return nil
	}
	return o.FailureCode
}

func (o *SynchronousRefundResponse) GetRefundID() *string {
	if o == nil {
		return nil
	}
	return o.RefundID
}

func (o *SynchronousRefundResponse) GetStatus() *RefundStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SynchronousRefundResponse) GetUpdatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedOn
}

type CreatedReversal struct {
	Cancellation *CreatedCancellation       `json:"cancellation,omitempty"`
	Refund       *SynchronousRefundResponse `json:"refund,omitempty"`
}

func (o *CreatedReversal) GetCancellation() *CreatedCancellation {
	if o == nil {
		return nil
	}
	return o.Cancellation
}

func (o *CreatedReversal) GetRefund() *SynchronousRefundResponse {
	if o == nil {
		return nil
	}
	return o.Refund
}
