// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

type CardDetails struct {
	// An optional override of the default card statement descriptor for a transfer.
	DynamicDescriptor *string          `json:"dynamicDescriptor,omitempty"`
	FailureCode       *CardFailureCode `json:"failureCode,omitempty"`
	// Card status
	Status *CardStatus `json:"status,omitempty"`
	// Describes how the card transaction was initiated
	TransactionSource *TransactionSource `json:"transactionSource,omitempty"`
}

func (o *CardDetails) GetDynamicDescriptor() *string {
	if o == nil {
		return nil
	}
	return o.DynamicDescriptor
}

func (o *CardDetails) GetFailureCode() *CardFailureCode {
	if o == nil {
		return nil
	}
	return o.FailureCode
}

func (o *CardDetails) GetStatus() *CardStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CardDetails) GetTransactionSource() *TransactionSource {
	if o == nil {
		return nil
	}
	return o.TransactionSource
}
