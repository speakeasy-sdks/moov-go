// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

type BankAccountPayloadType string

const (
	BankAccountPayloadTypeBankAccount BankAccountPayloadType = "BankAccount"
	BankAccountPayloadTypePlaid       BankAccountPayloadType = "Plaid"
	BankAccountPayloadTypePlaidLink   BankAccountPayloadType = "PlaidLink"
	BankAccountPayloadTypeMx          BankAccountPayloadType = "MX"
)

type BankAccountPayload struct {
	BankAccount *BankAccount
	Plaid       *Plaid
	PlaidLink   *PlaidLink
	Mx          *Mx

	Type BankAccountPayloadType
}

func CreateBankAccountPayloadBankAccount(bankAccount BankAccount) BankAccountPayload {
	typ := BankAccountPayloadTypeBankAccount

	return BankAccountPayload{
		BankAccount: &bankAccount,
		Type:        typ,
	}
}

func CreateBankAccountPayloadPlaid(plaid Plaid) BankAccountPayload {
	typ := BankAccountPayloadTypePlaid

	return BankAccountPayload{
		Plaid: &plaid,
		Type:  typ,
	}
}

func CreateBankAccountPayloadPlaidLink(plaidLink PlaidLink) BankAccountPayload {
	typ := BankAccountPayloadTypePlaidLink

	return BankAccountPayload{
		PlaidLink: &plaidLink,
		Type:      typ,
	}
}

func CreateBankAccountPayloadMx(mx Mx) BankAccountPayload {
	typ := BankAccountPayloadTypeMx

	return BankAccountPayload{
		Mx:   &mx,
		Type: typ,
	}
}

func (u *BankAccountPayload) UnmarshalJSON(data []byte) error {

	bankAccount := BankAccount{}
	if err := utils.UnmarshalJSON(data, &bankAccount, "", true, true); err == nil {
		u.BankAccount = &bankAccount
		u.Type = BankAccountPayloadTypeBankAccount
		return nil
	}

	plaid := Plaid{}
	if err := utils.UnmarshalJSON(data, &plaid, "", true, true); err == nil {
		u.Plaid = &plaid
		u.Type = BankAccountPayloadTypePlaid
		return nil
	}

	plaidLink := PlaidLink{}
	if err := utils.UnmarshalJSON(data, &plaidLink, "", true, true); err == nil {
		u.PlaidLink = &plaidLink
		u.Type = BankAccountPayloadTypePlaidLink
		return nil
	}

	mx := Mx{}
	if err := utils.UnmarshalJSON(data, &mx, "", true, true); err == nil {
		u.Mx = &mx
		u.Type = BankAccountPayloadTypeMx
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u BankAccountPayload) MarshalJSON() ([]byte, error) {
	if u.BankAccount != nil {
		return utils.MarshalJSON(u.BankAccount, "", true)
	}

	if u.Plaid != nil {
		return utils.MarshalJSON(u.Plaid, "", true)
	}

	if u.PlaidLink != nil {
		return utils.MarshalJSON(u.PlaidLink, "", true)
	}

	if u.Mx != nil {
		return utils.MarshalJSON(u.Mx, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
