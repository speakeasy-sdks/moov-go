// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateRefund - Specifies a partial amount to refund. This request body is optional, an empty body will issue a refund for the full amount of the original transfer.
type CreateRefund struct {
	// Amount to refund in cents. If null, the original transfer's full amount will be refunded.
	Amount *int64 `json:"amount,omitempty"`
}

func (o *CreateRefund) GetAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.Amount
}