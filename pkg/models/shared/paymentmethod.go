// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

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

type PaymentMethod struct {
	PaymentMethodWallet      *PaymentMethodWallet
	PaymentMethodBankAccount *PaymentMethodBankAccount
	PaymentMethodCard        *PaymentMethodCard
	PaymentMethodApplePay    *PaymentMethodApplePay

	Type PaymentMethodType
}

func CreatePaymentMethodAchCreditSameDay(achCreditSameDay PaymentMethodBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchCreditSameDay

	typStr := PaymentMethodsType(typ)
	achCreditSameDay.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodBankAccount: &achCreditSameDay,
		Type:                     typ,
	}
}

func CreatePaymentMethodAchCreditStandard(achCreditStandard PaymentMethodBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchCreditStandard

	typStr := PaymentMethodsType(typ)
	achCreditStandard.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodBankAccount: &achCreditStandard,
		Type:                     typ,
	}
}

func CreatePaymentMethodAchDebitCollect(achDebitCollect PaymentMethodBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchDebitCollect

	typStr := PaymentMethodsType(typ)
	achDebitCollect.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodBankAccount: &achDebitCollect,
		Type:                     typ,
	}
}

func CreatePaymentMethodAchDebitFund(achDebitFund PaymentMethodBankAccount) PaymentMethod {
	typ := PaymentMethodTypeAchDebitFund

	typStr := PaymentMethodsType(typ)
	achDebitFund.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodBankAccount: &achDebitFund,
		Type:                     typ,
	}
}

func CreatePaymentMethodApplePay(applePay PaymentMethodApplePay) PaymentMethod {
	typ := PaymentMethodTypeApplePay

	typStr := PaymentMethodsType(typ)
	applePay.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodApplePay: &applePay,
		Type:                  typ,
	}
}

func CreatePaymentMethodCardPayment(cardPayment PaymentMethodCard) PaymentMethod {
	typ := PaymentMethodTypeCardPayment

	typStr := PaymentMethodsType(typ)
	cardPayment.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodCard: &cardPayment,
		Type:              typ,
	}
}

func CreatePaymentMethodMoovWallet(moovWallet PaymentMethodWallet) PaymentMethod {
	typ := PaymentMethodTypeMoovWallet

	typStr := PaymentMethodsType(typ)
	moovWallet.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodWallet: &moovWallet,
		Type:                typ,
	}
}

func CreatePaymentMethodRtpCredit(rtpCredit PaymentMethodBankAccount) PaymentMethod {
	typ := PaymentMethodTypeRtpCredit

	typStr := PaymentMethodsType(typ)
	rtpCredit.PaymentMethodType = &typStr

	return PaymentMethod{
		PaymentMethodBankAccount: &rtpCredit,
		Type:                     typ,
	}
}

func (u *PaymentMethod) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		PaymentMethodType string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.PaymentMethodType {
	case "ach-credit-same-day":
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchDebitFund
		return nil
	case "apple-pay":
		paymentMethodApplePay := new(PaymentMethodApplePay)
		if err := utils.UnmarshalJSON(data, &paymentMethodApplePay, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodApplePay = paymentMethodApplePay
		u.Type = PaymentMethodTypeApplePay
		return nil
	case "card-payment":
		paymentMethodCard := new(PaymentMethodCard)
		if err := utils.UnmarshalJSON(data, &paymentMethodCard, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCard = paymentMethodCard
		u.Type = PaymentMethodTypeCardPayment
		return nil
	case "moov-wallet":
		paymentMethodWallet := new(PaymentMethodWallet)
		if err := utils.UnmarshalJSON(data, &paymentMethodWallet, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodWallet = paymentMethodWallet
		u.Type = PaymentMethodTypeMoovWallet
		return nil
	case "rtp-credit":
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodBankAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeRtpCredit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentMethod) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodWallet != nil {
		return utils.MarshalJSON(u.PaymentMethodWallet, "", true)
	}

	if u.PaymentMethodBankAccount != nil {
		return utils.MarshalJSON(u.PaymentMethodBankAccount, "", true)
	}

	if u.PaymentMethodCard != nil {
		return utils.MarshalJSON(u.PaymentMethodCard, "", true)
	}

	if u.PaymentMethodApplePay != nil {
		return utils.MarshalJSON(u.PaymentMethodApplePay, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
