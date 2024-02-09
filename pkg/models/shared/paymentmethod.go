// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

// SchemasApplePay - Describes an Apple Pay token on a Moov account.
type SchemasApplePay struct {
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

func (o *SchemasApplePay) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *SchemasApplePay) GetCardDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.CardDisplayName
}

func (o *SchemasApplePay) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *SchemasApplePay) GetDynamicLastFour() *string {
	if o == nil {
		return nil
	}
	return o.DynamicLastFour
}

func (o *SchemasApplePay) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *SchemasApplePay) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

// ApplePay - A method of moving money using an Apple Pay token.
type ApplePay struct {
	ApplePay *SchemasApplePay `json:"applePay,omitempty"`
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType `json:"paymentMethodType,omitempty"`
}

func (o *ApplePay) GetApplePay() *SchemasApplePay {
	if o == nil {
		return nil
	}
	return o.ApplePay
}

func (o *ApplePay) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *ApplePay) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}

// SchemasCard - Describes a card on a Moov account
type SchemasCard struct {
	BillingAddress *Address `json:"billingAddress,omitempty"`
	Bin            *string  `json:"bin,omitempty"`
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

func (s SchemasCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemasCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemasCard) GetBillingAddress() *Address {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *SchemasCard) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *SchemasCard) GetBrand() *CardBrand {
	if o == nil {
		return nil
	}
	return o.Brand
}

func (o *SchemasCard) GetCardAccountUpdater() *CardAccountUpdater {
	if o == nil {
		return nil
	}
	return o.CardAccountUpdater
}

func (o *SchemasCard) GetCardID() *string {
	if o == nil {
		return nil
	}
	return o.CardID
}

func (o *SchemasCard) GetCardOnFile() *bool {
	if o == nil {
		return nil
	}
	return o.CardOnFile
}

func (o *SchemasCard) GetCardType() *CardType {
	if o == nil {
		return nil
	}
	return o.CardType
}

func (o *SchemasCard) GetCardVerification() *CardVerifications {
	if o == nil {
		return nil
	}
	return o.CardVerification
}

func (o *SchemasCard) GetExpiration() *CardExpiration {
	if o == nil {
		return nil
	}
	return o.Expiration
}

func (o *SchemasCard) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *SchemasCard) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *SchemasCard) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *SchemasCard) GetIssuerCountry() *string {
	if o == nil {
		return nil
	}
	return o.IssuerCountry
}

func (o *SchemasCard) GetLastFourCardNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourCardNumber
}

func (o *SchemasCard) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

// SchemasPaymentMethodCardCard - A method of moving money that is a credit or debit card
type SchemasPaymentMethodCardCard struct {
	Card *SchemasCard `json:"card,omitempty"`
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType `json:"paymentMethodType,omitempty"`
}

func (o *SchemasPaymentMethodCardCard) GetCard() *SchemasCard {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *SchemasPaymentMethodCardCard) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *SchemasPaymentMethodCardCard) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}

// SchemasBankAccount - Describes a bank account on a Moov account.
type SchemasBankAccount struct {
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

func (o *SchemasBankAccount) GetBankAccountID() *string {
	if o == nil {
		return nil
	}
	return o.BankAccountID
}

func (o *SchemasBankAccount) GetBankAccountType() *BankAccountType {
	if o == nil {
		return nil
	}
	return o.BankAccountType
}

func (o *SchemasBankAccount) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

func (o *SchemasBankAccount) GetFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.Fingerprint
}

func (o *SchemasBankAccount) GetHolderName() *string {
	if o == nil {
		return nil
	}
	return o.HolderName
}

func (o *SchemasBankAccount) GetHolderType() *HolderType {
	if o == nil {
		return nil
	}
	return o.HolderType
}

func (o *SchemasBankAccount) GetLastFourAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.LastFourAccountNumber
}

func (o *SchemasBankAccount) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *SchemasBankAccount) GetStatus() *BankAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// SchemasPaymentMethodBankAccountBankAccount - A method of moving money that is a bank account
type SchemasPaymentMethodBankAccountBankAccount struct {
	BankAccount *SchemasBankAccount `json:"bankAccount,omitempty"`
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType `json:"paymentMethodType,omitempty"`
}

func (o *SchemasPaymentMethodBankAccountBankAccount) GetBankAccount() *SchemasBankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *SchemasPaymentMethodBankAccountBankAccount) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *SchemasPaymentMethodBankAccountBankAccount) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}

// SchemasWallet - A Moov wallet to store funds for transfers.
type SchemasWallet struct {
	// UUID v4
	WalletID *string `json:"walletID,omitempty"`
}

func (o *SchemasWallet) GetWalletID() *string {
	if o == nil {
		return nil
	}
	return o.WalletID
}

// SchemasPaymentMethodWalletWallet - A method of moving money that is a Moov wallet
type SchemasPaymentMethodWalletWallet struct {
	// UUID v4
	PaymentMethodID *string `json:"paymentMethodID,omitempty"`
	// The payment method type that represents a payment rail and directionality
	PaymentMethodType *PaymentMethodsType `json:"paymentMethodType,omitempty"`
	Wallet            *SchemasWallet      `json:"wallet,omitempty"`
}

func (o *SchemasPaymentMethodWalletWallet) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

func (o *SchemasPaymentMethodWalletWallet) GetPaymentMethodType() *PaymentMethodsType {
	if o == nil {
		return nil
	}
	return o.PaymentMethodType
}

func (o *SchemasPaymentMethodWalletWallet) GetWallet() *SchemasWallet {
	if o == nil {
		return nil
	}
	return o.Wallet
}

type PaymentMethodType string

const (
	PaymentMethodTypeAchCreditSameDay  PaymentMethodType = "ach-credit-same-day"
	PaymentMethodTypeAchCreditStandard PaymentMethodType = "ach-credit-standard"
	PaymentMethodTypeAchDebitCollect   PaymentMethodType = "ach-debit-collect"
	PaymentMethodTypeAchDebitFund      PaymentMethodType = "ach-debit-fund"
	PaymentMethodTypeApplePay          PaymentMethodType = "apple-pay"
	PaymentMethodTypeCardPayment       PaymentMethodType = "card-payment"
	PaymentMethodTypeMoovWallet        PaymentMethodType = "moov-wallet"
	PaymentMethodTypeRtpCredit         PaymentMethodType = "rtp-credit"
)

// PaymentMethod - A method of moving money
type PaymentMethod struct {
	SchemasPaymentMethodWalletWallet           *SchemasPaymentMethodWalletWallet
	SchemasPaymentMethodBankAccountBankAccount *SchemasPaymentMethodBankAccountBankAccount
	SchemasPaymentMethodCardCard               *SchemasPaymentMethodCardCard
	ApplePay                                   *ApplePay

	Type PaymentMethodType
}

func CreatePaymentMethodAchCreditSameDay(achCreditSameDay SchemasPaymentMethodBankAccountBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchCreditSameDay

	typStr := PaymentMethodsType(typ)
	achCreditSameDay.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodBankAccountBankAccount: &achCreditSameDay,
		Type: typ,
	}
}

func CreatePaymentMethodAchCreditStandard(achCreditStandard SchemasPaymentMethodBankAccountBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchCreditStandard

	typStr := PaymentMethodsType(typ)
	achCreditStandard.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodBankAccountBankAccount: &achCreditStandard,
		Type: typ,
	}
}

func CreatePaymentMethodAchDebitCollect(achDebitCollect SchemasPaymentMethodBankAccountBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchDebitCollect

	typStr := PaymentMethodsType(typ)
	achDebitCollect.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodBankAccountBankAccount: &achDebitCollect,
		Type: typ,
	}
}

func CreatePaymentMethodAchDebitFund(achDebitFund SchemasPaymentMethodBankAccountBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchDebitFund

	typStr := PaymentMethodsType(typ)
	achDebitFund.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodBankAccountBankAccount: &achDebitFund,
		Type: typ,
	}
}

func CreatePaymentMethodApplePay(applePay ApplePay) PaymentMethod {
	typ := PaymentMethodTypeApplePay

	typStr := PaymentMethodsType(typ)
	applePay.PaymentMethodType = &typStr

	return PaymentMethod{
		ApplePay: &applePay,
		Type:     typ,
	}
}

func CreatePaymentMethodCardPayment(cardPayment SchemasPaymentMethodCardCard) PaymentMethod {
	typ := PaymentMethodTypeCardPayment

	typStr := PaymentMethodsType(typ)
	cardPayment.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodCardCard: &cardPayment,
		Type:                         typ,
	}
}

func CreatePaymentMethodMoovWallet(moovWallet SchemasPaymentMethodWalletWallet) PaymentMethod {
	typ := PaymentMethodTypeMoovWallet

	typStr := PaymentMethodsType(typ)
	moovWallet.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodWalletWallet: &moovWallet,
		Type:                             typ,
	}
}

func CreatePaymentMethodRtpCredit(rtpCredit SchemasPaymentMethodBankAccountBankAccount) PaymentMethod {
	typ := PaymentMethodTypeRtpCredit

	typStr := PaymentMethodsType(typ)
	rtpCredit.PaymentMethodType = &typStr

	return PaymentMethod{
		SchemasPaymentMethodBankAccountBankAccount: &rtpCredit,
		Type: typ,
	}
}

func (u *PaymentMethod) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		PaymentMethodType string `json:"paymentMethodType"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.PaymentMethodType {
	case "ach-credit-same-day":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = PaymentMethodTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = PaymentMethodTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = PaymentMethodTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = PaymentMethodTypeAchDebitFund
		return nil
	case "apple-pay":
		applePay := new(ApplePay)
		if err := utils.UnmarshalJSON(data, &applePay, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ApplePay = applePay
		u.Type = PaymentMethodTypeApplePay
		return nil
	case "card-payment":
		schemasPaymentMethodCardCard := new(SchemasPaymentMethodCardCard)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodCardCard, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodCardCard = schemasPaymentMethodCardCard
		u.Type = PaymentMethodTypeCardPayment
		return nil
	case "moov-wallet":
		schemasPaymentMethodWalletWallet := new(SchemasPaymentMethodWalletWallet)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodWalletWallet, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodWalletWallet = schemasPaymentMethodWalletWallet
		u.Type = PaymentMethodTypeMoovWallet
		return nil
	case "rtp-credit":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = PaymentMethodTypeRtpCredit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentMethod) MarshalJSON() ([]byte, error) {
	if u.SchemasPaymentMethodWalletWallet != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodWalletWallet, "", true)
	}

	if u.SchemasPaymentMethodBankAccountBankAccount != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodBankAccountBankAccount, "", true)
	}

	if u.SchemasPaymentMethodCardCard != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodCardCard, "", true)
	}

	if u.ApplePay != nil {
		return utils.MarshalJSON(u.ApplePay, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationOptionsType string

const (
	DestinationOptionsTypeAchCreditSameDay  DestinationOptionsType = "ach-credit-same-day"
	DestinationOptionsTypeAchCreditStandard DestinationOptionsType = "ach-credit-standard"
	DestinationOptionsTypeAchDebitCollect   DestinationOptionsType = "ach-debit-collect"
	DestinationOptionsTypeAchDebitFund      DestinationOptionsType = "ach-debit-fund"
	DestinationOptionsTypeApplePay          DestinationOptionsType = "apple-pay"
	DestinationOptionsTypeCardPayment       DestinationOptionsType = "card-payment"
	DestinationOptionsTypeMoovWallet        DestinationOptionsType = "moov-wallet"
	DestinationOptionsTypeRtpCredit         DestinationOptionsType = "rtp-credit"
)

// DestinationOptions - A method of moving money
type DestinationOptions struct {
	SchemasPaymentMethodWalletWallet           *SchemasPaymentMethodWalletWallet
	SchemasPaymentMethodBankAccountBankAccount *SchemasPaymentMethodBankAccountBankAccount
	SchemasPaymentMethodCardCard               *SchemasPaymentMethodCardCard
	ApplePay                                   *ApplePay

	Type DestinationOptionsType
}

func CreateDestinationOptionsAchCreditSameDay(achCreditSameDay SchemasPaymentMethodBankAccountBankAccount) DestinationOptions {
	typ := DestinationOptionsTypeAchCreditSameDay

	typStr := PaymentMethodsType(typ)
	achCreditSameDay.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achCreditSameDay,
		Type: typ,
	}
}

func CreateDestinationOptionsAchCreditStandard(achCreditStandard SchemasPaymentMethodBankAccountBankAccount) DestinationOptions {
	typ := DestinationOptionsTypeAchCreditStandard

	typStr := PaymentMethodsType(typ)
	achCreditStandard.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achCreditStandard,
		Type: typ,
	}
}

func CreateDestinationOptionsAchDebitCollect(achDebitCollect SchemasPaymentMethodBankAccountBankAccount) DestinationOptions {
	typ := DestinationOptionsTypeAchDebitCollect

	typStr := PaymentMethodsType(typ)
	achDebitCollect.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achDebitCollect,
		Type: typ,
	}
}

func CreateDestinationOptionsAchDebitFund(achDebitFund SchemasPaymentMethodBankAccountBankAccount) DestinationOptions {
	typ := DestinationOptionsTypeAchDebitFund

	typStr := PaymentMethodsType(typ)
	achDebitFund.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achDebitFund,
		Type: typ,
	}
}

func CreateDestinationOptionsApplePay(applePay ApplePay) DestinationOptions {
	typ := DestinationOptionsTypeApplePay

	typStr := PaymentMethodsType(typ)
	applePay.PaymentMethodType = &typStr

	return DestinationOptions{
		ApplePay: &applePay,
		Type:     typ,
	}
}

func CreateDestinationOptionsCardPayment(cardPayment SchemasPaymentMethodCardCard) DestinationOptions {
	typ := DestinationOptionsTypeCardPayment

	typStr := PaymentMethodsType(typ)
	cardPayment.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodCardCard: &cardPayment,
		Type:                         typ,
	}
}

func CreateDestinationOptionsMoovWallet(moovWallet SchemasPaymentMethodWalletWallet) DestinationOptions {
	typ := DestinationOptionsTypeMoovWallet

	typStr := PaymentMethodsType(typ)
	moovWallet.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodWalletWallet: &moovWallet,
		Type:                             typ,
	}
}

func CreateDestinationOptionsRtpCredit(rtpCredit SchemasPaymentMethodBankAccountBankAccount) DestinationOptions {
	typ := DestinationOptionsTypeRtpCredit

	typStr := PaymentMethodsType(typ)
	rtpCredit.PaymentMethodType = &typStr

	return DestinationOptions{
		SchemasPaymentMethodBankAccountBankAccount: &rtpCredit,
		Type: typ,
	}
}

func (u *DestinationOptions) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		PaymentMethodType string `json:"paymentMethodType"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.PaymentMethodType {
	case "ach-credit-same-day":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = DestinationOptionsTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = DestinationOptionsTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = DestinationOptionsTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = DestinationOptionsTypeAchDebitFund
		return nil
	case "apple-pay":
		applePay := new(ApplePay)
		if err := utils.UnmarshalJSON(data, &applePay, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ApplePay = applePay
		u.Type = DestinationOptionsTypeApplePay
		return nil
	case "card-payment":
		schemasPaymentMethodCardCard := new(SchemasPaymentMethodCardCard)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodCardCard, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodCardCard = schemasPaymentMethodCardCard
		u.Type = DestinationOptionsTypeCardPayment
		return nil
	case "moov-wallet":
		schemasPaymentMethodWalletWallet := new(SchemasPaymentMethodWalletWallet)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodWalletWallet, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodWalletWallet = schemasPaymentMethodWalletWallet
		u.Type = DestinationOptionsTypeMoovWallet
		return nil
	case "rtp-credit":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = DestinationOptionsTypeRtpCredit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationOptions) MarshalJSON() ([]byte, error) {
	if u.SchemasPaymentMethodWalletWallet != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodWalletWallet, "", true)
	}

	if u.SchemasPaymentMethodBankAccountBankAccount != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodBankAccountBankAccount, "", true)
	}

	if u.SchemasPaymentMethodCardCard != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodCardCard, "", true)
	}

	if u.ApplePay != nil {
		return utils.MarshalJSON(u.ApplePay, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceOptionsType string

const (
	SourceOptionsTypeAchCreditSameDay  SourceOptionsType = "ach-credit-same-day"
	SourceOptionsTypeAchCreditStandard SourceOptionsType = "ach-credit-standard"
	SourceOptionsTypeAchDebitCollect   SourceOptionsType = "ach-debit-collect"
	SourceOptionsTypeAchDebitFund      SourceOptionsType = "ach-debit-fund"
	SourceOptionsTypeApplePay          SourceOptionsType = "apple-pay"
	SourceOptionsTypeCardPayment       SourceOptionsType = "card-payment"
	SourceOptionsTypeMoovWallet        SourceOptionsType = "moov-wallet"
	SourceOptionsTypeRtpCredit         SourceOptionsType = "rtp-credit"
)

// SourceOptions - A method of moving money
type SourceOptions struct {
	SchemasPaymentMethodWalletWallet           *SchemasPaymentMethodWalletWallet
	SchemasPaymentMethodBankAccountBankAccount *SchemasPaymentMethodBankAccountBankAccount
	SchemasPaymentMethodCardCard               *SchemasPaymentMethodCardCard
	ApplePay                                   *ApplePay

	Type SourceOptionsType
}

func CreateSourceOptionsAchCreditSameDay(achCreditSameDay SchemasPaymentMethodBankAccountBankAccount) SourceOptions {
	typ := SourceOptionsTypeAchCreditSameDay

	typStr := PaymentMethodsType(typ)
	achCreditSameDay.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achCreditSameDay,
		Type: typ,
	}
}

func CreateSourceOptionsAchCreditStandard(achCreditStandard SchemasPaymentMethodBankAccountBankAccount) SourceOptions {
	typ := SourceOptionsTypeAchCreditStandard

	typStr := PaymentMethodsType(typ)
	achCreditStandard.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achCreditStandard,
		Type: typ,
	}
}

func CreateSourceOptionsAchDebitCollect(achDebitCollect SchemasPaymentMethodBankAccountBankAccount) SourceOptions {
	typ := SourceOptionsTypeAchDebitCollect

	typStr := PaymentMethodsType(typ)
	achDebitCollect.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achDebitCollect,
		Type: typ,
	}
}

func CreateSourceOptionsAchDebitFund(achDebitFund SchemasPaymentMethodBankAccountBankAccount) SourceOptions {
	typ := SourceOptionsTypeAchDebitFund

	typStr := PaymentMethodsType(typ)
	achDebitFund.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodBankAccountBankAccount: &achDebitFund,
		Type: typ,
	}
}

func CreateSourceOptionsApplePay(applePay ApplePay) SourceOptions {
	typ := SourceOptionsTypeApplePay

	typStr := PaymentMethodsType(typ)
	applePay.PaymentMethodType = &typStr

	return SourceOptions{
		ApplePay: &applePay,
		Type:     typ,
	}
}

func CreateSourceOptionsCardPayment(cardPayment SchemasPaymentMethodCardCard) SourceOptions {
	typ := SourceOptionsTypeCardPayment

	typStr := PaymentMethodsType(typ)
	cardPayment.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodCardCard: &cardPayment,
		Type:                         typ,
	}
}

func CreateSourceOptionsMoovWallet(moovWallet SchemasPaymentMethodWalletWallet) SourceOptions {
	typ := SourceOptionsTypeMoovWallet

	typStr := PaymentMethodsType(typ)
	moovWallet.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodWalletWallet: &moovWallet,
		Type:                             typ,
	}
}

func CreateSourceOptionsRtpCredit(rtpCredit SchemasPaymentMethodBankAccountBankAccount) SourceOptions {
	typ := SourceOptionsTypeRtpCredit

	typStr := PaymentMethodsType(typ)
	rtpCredit.PaymentMethodType = &typStr

	return SourceOptions{
		SchemasPaymentMethodBankAccountBankAccount: &rtpCredit,
		Type: typ,
	}
}

func (u *SourceOptions) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		PaymentMethodType string `json:"paymentMethodType"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.PaymentMethodType {
	case "ach-credit-same-day":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = SourceOptionsTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = SourceOptionsTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = SourceOptionsTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = SourceOptionsTypeAchDebitFund
		return nil
	case "apple-pay":
		applePay := new(ApplePay)
		if err := utils.UnmarshalJSON(data, &applePay, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ApplePay = applePay
		u.Type = SourceOptionsTypeApplePay
		return nil
	case "card-payment":
		schemasPaymentMethodCardCard := new(SchemasPaymentMethodCardCard)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodCardCard, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodCardCard = schemasPaymentMethodCardCard
		u.Type = SourceOptionsTypeCardPayment
		return nil
	case "moov-wallet":
		schemasPaymentMethodWalletWallet := new(SchemasPaymentMethodWalletWallet)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodWalletWallet, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodWalletWallet = schemasPaymentMethodWalletWallet
		u.Type = SourceOptionsTypeMoovWallet
		return nil
	case "rtp-credit":
		schemasPaymentMethodBankAccountBankAccount := new(SchemasPaymentMethodBankAccountBankAccount)
		if err := utils.UnmarshalJSON(data, &schemasPaymentMethodBankAccountBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.SchemasPaymentMethodBankAccountBankAccount = schemasPaymentMethodBankAccountBankAccount
		u.Type = SourceOptionsTypeRtpCredit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOptions) MarshalJSON() ([]byte, error) {
	if u.SchemasPaymentMethodWalletWallet != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodWalletWallet, "", true)
	}

	if u.SchemasPaymentMethodBankAccountBankAccount != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodBankAccountBankAccount, "", true)
	}

	if u.SchemasPaymentMethodCardCard != nil {
		return utils.MarshalJSON(u.SchemasPaymentMethodCardCard, "", true)
	}

	if u.ApplePay != nil {
		return utils.MarshalJSON(u.ApplePay, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreatedTransferOptions struct {
	DestinationOptions []DestinationOptions `json:"destinationOptions,omitempty"`
	SourceOptions      []SourceOptions      `json:"sourceOptions,omitempty"`
}

func (o *CreatedTransferOptions) GetDestinationOptions() []DestinationOptions {
	if o == nil {
		return nil
	}
	return o.DestinationOptions
}

func (o *CreatedTransferOptions) GetSourceOptions() []SourceOptions {
	if o == nil {
		return nil
	}
	return o.SourceOptions
}
