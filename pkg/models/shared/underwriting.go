// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Underwriting - Describes underwriting values (in USD) used for card payment acceptance
type Underwriting struct {
	AverageMonthlyTransactionVolume *int64 `json:"averageMonthlyTransactionVolume,omitempty"`
	AverageTransactionSize          *int64 `json:"averageTransactionSize,omitempty"`
	MaxTransactionSize              *int64 `json:"maxTransactionSize,omitempty"`
	// The status of underwriting for an account
	Status *UnderwritingStatus `json:"status,omitempty"`
}

func (o *Underwriting) GetAverageMonthlyTransactionVolume() *int64 {
	if o == nil {
		return nil
	}
	return o.AverageMonthlyTransactionVolume
}

func (o *Underwriting) GetAverageTransactionSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AverageTransactionSize
}

func (o *Underwriting) GetMaxTransactionSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTransactionSize
}

func (o *Underwriting) GetStatus() *UnderwritingStatus {
	if o == nil {
		return nil
	}
	return o.Status
}