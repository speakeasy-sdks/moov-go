// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BankAccountIntegration - Describes the account to link to the Moov account
type BankAccountIntegration struct {
	AccountNumber string `json:"accountNumber"`
	// The bank account type
	BankAccountType BankAccountType `json:"bankAccountType"`
	HolderName      string          `json:"holderName"`
	// The type of holder on a funding source
	HolderType    HolderType `json:"holderType"`
	RoutingNumber string     `json:"routingNumber"`
}

func (o *BankAccountIntegration) GetAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.AccountNumber
}

func (o *BankAccountIntegration) GetBankAccountType() BankAccountType {
	if o == nil {
		return BankAccountType("")
	}
	return o.BankAccountType
}

func (o *BankAccountIntegration) GetHolderName() string {
	if o == nil {
		return ""
	}
	return o.HolderName
}

func (o *BankAccountIntegration) GetHolderType() HolderType {
	if o == nil {
		return HolderType("")
	}
	return o.HolderType
}

func (o *BankAccountIntegration) GetRoutingNumber() string {
	if o == nil {
		return ""
	}
	return o.RoutingNumber
}
