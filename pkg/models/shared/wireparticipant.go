// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// WireParticipant - Financial institution information regarding a wire participant
type WireParticipant struct {
	BookEntrySecuritiesTransferStatus *string       `json:"bookEntrySecuritiesTransferStatus,omitempty"`
	CustomerName                      *string       `json:"customerName,omitempty"`
	Date                              *string       `json:"date,omitempty"`
	FundsSettlementOnlyStatus         *string       `json:"fundsSettlementOnlyStatus,omitempty"`
	FundsTransferStatus               *string       `json:"fundsTransferStatus,omitempty"`
	Location                          *WireLocation `json:"location,omitempty"`
	RoutingNumber                     *string       `json:"routingNumber,omitempty"`
	TelegraphicName                   *string       `json:"telegraphicName,omitempty"`
}

func (o *WireParticipant) GetBookEntrySecuritiesTransferStatus() *string {
	if o == nil {
		return nil
	}
	return o.BookEntrySecuritiesTransferStatus
}

func (o *WireParticipant) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *WireParticipant) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *WireParticipant) GetFundsSettlementOnlyStatus() *string {
	if o == nil {
		return nil
	}
	return o.FundsSettlementOnlyStatus
}

func (o *WireParticipant) GetFundsTransferStatus() *string {
	if o == nil {
		return nil
	}
	return o.FundsTransferStatus
}

func (o *WireParticipant) GetLocation() *WireLocation {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *WireParticipant) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *WireParticipant) GetTelegraphicName() *string {
	if o == nil {
		return nil
	}
	return o.TelegraphicName
}
