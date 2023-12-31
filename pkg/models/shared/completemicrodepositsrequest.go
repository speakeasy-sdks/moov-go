// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CompleteMicroDepositsRequest - Request to complete the micro-deposit verification workflow.
type CompleteMicroDepositsRequest struct {
	// Two positive integers, in cents, equal to the values of the micro-deposits sent to the bank account.
	Amounts []int64 `json:"amounts,omitempty"`
}

func (o *CompleteMicroDepositsRequest) GetAmounts() []int64 {
	if o == nil {
		return nil
	}
	return o.Amounts
}
