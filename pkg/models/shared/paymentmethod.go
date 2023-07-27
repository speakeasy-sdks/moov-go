// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	var d *json.Decoder

	type discriminator struct {
		PaymentMethodType string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.PaymentMethodType {
	case "ach-credit-same-day":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = PaymentMethodTypeAchDebitFund
		return nil
	case "apple-pay":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodApplePay := new(PaymentMethodApplePay)
		if err := d.Decode(&paymentMethodApplePay); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodApplePay = paymentMethodApplePay
		u.Type = PaymentMethodTypeApplePay
		return nil
	case "card-payment":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodCard := new(PaymentMethodCard)
		if err := d.Decode(&paymentMethodCard); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCard = paymentMethodCard
		u.Type = PaymentMethodTypeCardPayment
		return nil
	case "moov-wallet":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodWallet := new(PaymentMethodWallet)
		if err := d.Decode(&paymentMethodWallet); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodWallet = paymentMethodWallet
		u.Type = PaymentMethodTypeMoovWallet
		return nil
	case "rtp-credit":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
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
		return json.Marshal(u.PaymentMethodWallet)
	}

	if u.PaymentMethodBankAccount != nil {
		return json.Marshal(u.PaymentMethodBankAccount)
	}

	if u.PaymentMethodCard != nil {
		return json.Marshal(u.PaymentMethodCard)
	}

	if u.PaymentMethodApplePay != nil {
		return json.Marshal(u.PaymentMethodApplePay)
	}

	return nil, nil
}