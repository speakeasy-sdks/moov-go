// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

// GetTransferFullSourceDestinationApplePay - Describes an Apple Pay token on a Moov account.
type GetTransferFullSourceDestinationApplePay struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
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

func (g GetTransferFullSourceDestinationApplePay) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransferFullSourceDestinationApplePay) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTransferFullSourceDestinationApplePay) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetTransferFullSourceDestinationApplePay) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *GetTransferFullSourceDestinationApplePay) GetCardDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.CardDisplayName
}

func (o *GetTransferFullSourceDestinationApplePay) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *GetTransferFullSourceDestinationApplePay) GetDynamicLastFour() *string {
	if o == nil {
		return nil
	}
	return o.DynamicLastFour
}

func (o *GetTransferFullSourceDestinationApplePay) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *GetTransferFullSourceDestinationApplePay) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

// GetTransferFullSourceDestinationBankAccount - Describes a bank account on a Moov account.
type GetTransferFullSourceDestinationBankAccount struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
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

func (g GetTransferFullSourceDestinationBankAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransferFullSourceDestinationBankAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTransferFullSourceDestinationBankAccount) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetTransferFullSourceDestinationBankAccount) GetBankAccountID() *string {
	if o == nil {
		return nil
	}
	return o.BankAccountID
}

func (o *GetTransferFullSourceDestinationBankAccount) GetBankAccountType() *BankAccountType {
	if o == nil {
		return nil
	}
	return o.BankAccountType
}

func (o *GetTransferFullSourceDestinationBankAccount) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

func (o *GetTransferFullSourceDestinationBankAccount) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *GetTransferFullSourceDestinationBankAccount) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *GetTransferFullSourceDestinationBankAccount) GetHolderType() *HolderType {
	if o == nil {
		return nil
	}
	return o.HolderType
}

func (o *GetTransferFullSourceDestinationBankAccount) GetLastFourAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourAccountNumber
}

func (o *GetTransferFullSourceDestinationBankAccount) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *GetTransferFullSourceDestinationBankAccount) GetStatus() *BankAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTransferFullSourceDestinationCard - Describes a card on a Moov account
type GetTransferFullSourceDestinationCard struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	BillingAddress       *Address               `json:"billingAddress,omitempty"`
	Bin                  *string                `json:"bin,omitempty"`
	// The card brand
	Brand *CardBrand `json:"brand,omitempty"`
	// The results of the most recent card update request
	CardAccountUpdater *CardAccountUpdater `json:"cardAccountUpdater,omitempty"`
	// UUID v4
	CardID *string `json:"cardID,omitempty"`
	// Indicates cardholder has authorized card to be stored for future payments
	CardOnFile *bool `default:"false" json:"cardOnFile"`
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

func (g GetTransferFullSourceDestinationCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransferFullSourceDestinationCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTransferFullSourceDestinationCard) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetTransferFullSourceDestinationCard) GetBillingAddress() *Address {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *GetTransferFullSourceDestinationCard) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *GetTransferFullSourceDestinationCard) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *GetTransferFullSourceDestinationCard) GetCardAccountUpdater() *CardAccountUpdater {
	if o == nil {
		return nil
	}
	return o.CardAccountUpdater
}

func (o *GetTransferFullSourceDestinationCard) GetCardID() *string {
	if o == nil {
		return nil
	}
	return o.CardID
}

func (o *GetTransferFullSourceDestinationCard) GetCardOnFile() *bool {
	if o == nil {
		return nil
	}
	return o.CardOnFile
}

func (o *GetTransferFullSourceDestinationCard) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *GetTransferFullSourceDestinationCard) GetCardVerification() *CardVerifications {
	if o == nil {
		return nil
	}
	return o.CardVerification
}

func (o *GetTransferFullSourceDestinationCard) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *GetTransferFullSourceDestinationCard) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *GetTransferFullSourceDestinationCard) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *GetTransferFullSourceDestinationCard) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *GetTransferFullSourceDestinationCard) GetIssuerCountry() *string {
	if o == nil {
		return nil
	}
	return o.IssuerCountry
}

func (o *GetTransferFullSourceDestinationCard) GetLastFourCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourCardNumber
}

func (o *GetTransferFullSourceDestinationCard) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

// GetTransferFullSourceDestinationWallet - A Moov wallet to store funds for transfers.
type GetTransferFullSourceDestinationWallet struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// UUID v4
	WalletID *string `json:"walletID,omitempty"`
}

func (g GetTransferFullSourceDestinationWallet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransferFullSourceDestinationWallet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTransferFullSourceDestinationWallet) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetTransferFullSourceDestinationWallet) GetWalletID() *string {
	if o == nil {
		return nil
	}
	return o.WalletID
}

type GetTransferFullSourceDestination struct {
	Account *GetTransferPartialAccount `json:"account,omitempty"`
	// ACH specific details about the transaction
	AchDetails  *ACHDetailsBase                              `json:"achDetails,omitempty"`
	ApplePay    *GetTransferFullSourceDestinationApplePay    `json:"applePay,omitempty"`
	BankAccount *GetTransferFullSourceDestinationBankAccount `json:"bankAccount,omitempty"`
	Card        *GetTransferFullSourceDestinationCard        `json:"card,omitempty"`
	CardDetails *CardDetails                                 `json:"cardDetails,omitempty"`
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType                     `json:"paymentMethodType,omitempty"`
	Wallet            *GetTransferFullSourceDestinationWallet `json:"wallet,omitempty"`
}

func (o *GetTransferFullSourceDestination) GetAccount() *GetTransferPartialAccount {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *GetTransferFullSourceDestination) GetAchDetails() *ACHDetailsBase {
	if o == nil {
		return nil
	}
	return o.AchDetails
}

func (o *GetTransferFullSourceDestination) GetApplePay() *GetTransferFullSourceDestinationApplePay {
	if o == nil {
		return nil
	}
	return o.ApplePay
}

func (o *GetTransferFullSourceDestination) GetBankAccount() *GetTransferFullSourceDestinationBankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *GetTransferFullSourceDestination) GetCard() *GetTransferFullSourceDestinationCard {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *GetTransferFullSourceDestination) GetCardDetails() *CardDetails {
	if o == nil {
		return nil
	}
	return o.CardDetails
}

func (o *GetTransferFullSourceDestination) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *GetTransferFullSourceDestination) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}

func (o *GetTransferFullSourceDestination) GetWallet() *GetTransferFullSourceDestinationWallet {
	if o == nil {
		return nil
	}
	return o.Wallet
}
