// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PaymentMethodBankAccountBankAccount - Describes a bank account on a Moov account.
type PaymentMethodBankAccountBankAccount struct {
	// UUID v4
	BankAccountID *string `json:"bankAccountID,omitempty"`
	// The bank account type
	BankAccountType *BankAccountType `json:"bankAccountType,omitempty"`
	BankName        *string          `json:"bankName,omitempty"`
	// Once the bank account is linked, we don't reveal the full bank account number. The fingerprint acts as a way to identify whether two linked bank accounts are the same.
	Fingerprint *string `json:"fingerprint,omitempty"`
	HolderName  *string `json:"holderName,omitempty"`
	// The type of holder on a funding source
	HolderType            *HolderType `json:"holderType,omitempty"`
	LastFourAccountNumber *string     `json:"lastFourAccountNumber,omitempty"`
	RoutingNumber         *string     `json:"routingNumber,omitempty"`
	// The bank account status
	Status *BankAccountStatus `json:"status,omitempty"`
}

func (o *PaymentMethodBankAccountBankAccount) GetBankAccountID() *string {
	if o == nil {
		return nil
	}
	return o.BankAccountID
}

func (o *PaymentMethodBankAccountBankAccount) GetBankAccountType() *BankAccountType {
	if o == nil {
		return nil
	}
	return o.BankAccountType
}

func (o *PaymentMethodBankAccountBankAccount) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

func (o *PaymentMethodBankAccountBankAccount) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *PaymentMethodBankAccountBankAccount) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *PaymentMethodBankAccountBankAccount) GetHolderType() *HolderType {
	if o == nil {
		return nil
	}
	return o.HolderType
}

func (o *PaymentMethodBankAccountBankAccount) GetLastFourAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourAccountNumber
}

func (o *PaymentMethodBankAccountBankAccount) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *PaymentMethodBankAccountBankAccount) GetStatus() *BankAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// PaymentMethodBankAccount - A method of moving money that is a bank account
type PaymentMethodBankAccount struct {
	BankAccount *PaymentMethodBankAccountBankAccount `json:"bankAccount,omitempty"`
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType `json:"paymentMethodType,omitempty"`
}

func (o *PaymentMethodBankAccount) GetBankAccount() *PaymentMethodBankAccountBankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *PaymentMethodBankAccount) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *PaymentMethodBankAccount) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}