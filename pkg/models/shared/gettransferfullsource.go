// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetTransferFullSourceApplePay - Describes an Apple Pay token on a Moov account.
type GetTransferFullSourceApplePay struct {
	// The card brand
	Brand *CardBrand `json:"brand,omitempty"`
	// User-friendly name of the tokenized card returned by Apple.
	// It usually contains the brand and the last four digits of the underlying card for example, "Visa 1256".
	// There is no standard format.
	//
	CardDisplayName *string `json:"cardDisplayName,omitempty"`
	// The type of the card
	CardType *CardType `json:"cardType,omitempty"`
	// The last four digits of the Apple Pay token, which may differ from the tokenized card's last four digits
	DynamicLastFour *string `json:"dynamicLastFour,omitempty"`
	// The expiration date of the linked card or token
	Expiration *CardExpiration `json:"expiration,omitempty"`
	// Uniquely identifies a linked payment card or token.
	// For Apple Pay, the fingerprint is based on the tokenized card number and may vary based on the user's device.
	// This field can be used to identify specific payment methods across multiple accounts on your platform.
	//
	Fingerprint *string `json:"fingerprint,omitempty"`
}

func (o *GetTransferFullSourceApplePay) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *GetTransferFullSourceApplePay) GetCardDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.CardDisplayName
}

func (o *GetTransferFullSourceApplePay) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *GetTransferFullSourceApplePay) GetDynamicLastFour() *string {
	if o == nil {
		return nil
	}
	return o.DynamicLastFour
}

func (o *GetTransferFullSourceApplePay) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *GetTransferFullSourceApplePay) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

// GetTransferFullSourceBankAccount - Describes a bank account on a Moov account.
type GetTransferFullSourceBankAccount struct {
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

func (o *GetTransferFullSourceBankAccount) GetBankAccountID() *string {
	if o == nil {
		return nil
	}
	return o.BankAccountID
}

func (o *GetTransferFullSourceBankAccount) GetBankAccountType() *BankAccountType {
	if o == nil {
		return nil
	}
	return o.BankAccountType
}

func (o *GetTransferFullSourceBankAccount) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

func (o *GetTransferFullSourceBankAccount) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *GetTransferFullSourceBankAccount) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *GetTransferFullSourceBankAccount) GetHolderType() *HolderType {
	if o == nil {
		return nil
	}
	return o.HolderType
}

func (o *GetTransferFullSourceBankAccount) GetLastFourAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourAccountNumber
}

func (o *GetTransferFullSourceBankAccount) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *GetTransferFullSourceBankAccount) GetStatus() *BankAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTransferFullSourceCard - Describes a card on a Moov account
type GetTransferFullSourceCard struct {
	BillingAddress *Address `json:"billingAddress,omitempty"`
	Bin            *string  `json:"bin,omitempty"`
	// The card brand
	Brand *CardBrand `json:"brand,omitempty"`
	// The results of the most recent card update request
	CardAccountUpdater *CardAccountUpdater `json:"cardAccountUpdater,omitempty"`
	// UUID v4
	CardID *string `json:"cardID,omitempty"`
	// Indicates cardholder has authorized card to be stored for future payments
	CardOnFile *bool `json:"cardOnFile,omitempty"`
	// The type of the card
	CardType *CardType `json:"cardType,omitempty"`
	// The results of submitting cardholder data to a card network for verification
	CardVerification *CardVerifications `json:"cardVerification,omitempty"`
	// The expiration date of the linked card or token
	Expiration *CardExpiration `json:"expiration,omitempty"`
	// Uniquely identifies a linked payment card or token.
	// For Apple Pay, the fingerprint is based on the tokenized card number and may vary based on the user's device.
	// This field can be used to identify specific payment methods across multiple accounts on your platform.
	//
	Fingerprint        *string `json:"fingerprint,omitempty"`
	HolderName         *string `json:"holderName,omitempty"`
	Issuer             *string `json:"issuer,omitempty"`
	IssuerCountry      *string `json:"issuerCountry,omitempty"`
	LastFourCardNumber *string `json:"lastFourCardNumber,omitempty"`
	// Moov account ID of the merchant or entity authorized to store the card. Defaults to your platform account ID if cardOnFile is set to true and no other account is provided
	MerchantAccountID *string `json:"merchantAccountID,omitempty"`
}

func (o *GetTransferFullSourceCard) GetBillingAddress() *Address {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *GetTransferFullSourceCard) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *GetTransferFullSourceCard) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *GetTransferFullSourceCard) GetCardAccountUpdater() *CardAccountUpdater {
	if o == nil {
		return nil
	}
	return o.CardAccountUpdater
}

func (o *GetTransferFullSourceCard) GetCardID() *string {
	if o == nil {
		return nil
	}
	return o.CardID
}

func (o *GetTransferFullSourceCard) GetCardOnFile() *bool {
	if o == nil {
		return nil
	}
	return o.CardOnFile
}

func (o *GetTransferFullSourceCard) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *GetTransferFullSourceCard) GetCardVerification() *CardVerifications {
	if o == nil {
		return nil
	}
	return o.CardVerification
}

func (o *GetTransferFullSourceCard) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *GetTransferFullSourceCard) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *GetTransferFullSourceCard) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *GetTransferFullSourceCard) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *GetTransferFullSourceCard) GetIssuerCountry() *string {
	if o == nil {
		return nil
	}
	return o.IssuerCountry
}

func (o *GetTransferFullSourceCard) GetLastFourCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourCardNumber
}

func (o *GetTransferFullSourceCard) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

// GetTransferFullSourceWallet - A Moov wallet to store funds for transfers.
type GetTransferFullSourceWallet struct {
	// UUID v4
	WalletID *string `json:"walletID,omitempty"`
}

func (o *GetTransferFullSourceWallet) GetWalletID() *string {
	if o == nil {
		return nil
	}
	return o.WalletID
}

type GetTransferFullSource struct {
	Account     *GetTransferPartialAccount        `json:"account,omitempty"`
	AchDetails  *ACHDetailsSource                 `json:"achDetails,omitempty"`
	ApplePay    *GetTransferFullSourceApplePay    `json:"applePay,omitempty"`
	BankAccount *GetTransferFullSourceBankAccount `json:"bankAccount,omitempty"`
	Card        *GetTransferFullSourceCard        `json:"card,omitempty"`
	CardDetails *CardDetails                      `json:"cardDetails,omitempty"`
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType `json:"paymentMethodType,omitempty"`
	// UUID v4
	TransferID *string                      `json:"transferID,omitempty"`
	Wallet     *GetTransferFullSourceWallet `json:"wallet,omitempty"`
}

func (o *GetTransferFullSource) GetAccount() *GetTransferPartialAccount {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *GetTransferFullSource) GetAchDetails() *ACHDetailsSource {
	if o == nil {
		return nil
	}
	return o.AchDetails
}

func (o *GetTransferFullSource) GetApplePay() *GetTransferFullSourceApplePay {
	if o == nil {
		return nil
	}
	return o.ApplePay
}

func (o *GetTransferFullSource) GetBankAccount() *GetTransferFullSourceBankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *GetTransferFullSource) GetCard() *GetTransferFullSourceCard {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *GetTransferFullSource) GetCardDetails() *CardDetails {
	if o == nil {
		return nil
	}
	return o.CardDetails
}

func (o *GetTransferFullSource) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *GetTransferFullSource) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}

func (o *GetTransferFullSource) GetTransferID() *string {
	if o == nil {
		return nil
	}
	return o.TransferID
}

func (o *GetTransferFullSource) GetWallet() *GetTransferFullSourceWallet {
	if o == nil {
		return nil
	}
	return o.Wallet
}