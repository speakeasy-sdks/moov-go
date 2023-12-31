// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

// ACHDetailsSource - ACH specific details about the transaction
type ACHDetailsSource struct {
	// An optional override of the default NACHA Company Entry Description for a transfer.
	CompanyEntryDescription *string       `json:"companyEntryDescription,omitempty"`
	Correction              *ACHException `json:"correction,omitempty"`
	// An optional override of your default ACH hold period in banking days. The hold period must be longer than or equal to your default setting.
	DebitHoldPeriod *DebitHoldPeriod `json:"debitHoldPeriod,omitempty"`
	// An optional override of the default NACHA Company Name for a transfer.
	OriginatingCompanyName *string       `json:"originatingCompanyName,omitempty"`
	Return                 *ACHException `json:"return,omitempty"`
	// ACH status
	Status      ACHStatus `json:"status"`
	TraceNumber string    `json:"traceNumber"`
}

func (o *ACHDetailsSource) GetCompanyEntryDescription() *string {
	if o == nil {
		return nil
	}
	return o.CompanyEntryDescription
}

func (o *ACHDetailsSource) GetCorrection() *ACHException {
	if o == nil {
		return nil
	}
	return o.Correction
}

func (o *ACHDetailsSource) GetDebitHoldPeriod() *DebitHoldPeriod {
	if o == nil {
		return nil
	}
	return o.DebitHoldPeriod
}

func (o *ACHDetailsSource) GetOriginatingCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.OriginatingCompanyName
}

func (o *ACHDetailsSource) GetReturn() *ACHException {
	if o == nil {
		return nil
	}
	return o.Return
}

func (o *ACHDetailsSource) GetStatus() ACHStatus {
	if o == nil {
		return ACHStatus("")
	}
	return o.Status
}

func (o *ACHDetailsSource) GetTraceNumber() string {
	if o == nil {
		return ""
	}
	return o.TraceNumber
}
