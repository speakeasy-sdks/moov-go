// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type CreatedTransferOptionsDestinationOptionsType string

const (
	CreatedTransferOptionsDestinationOptionsTypeAchCreditSameDay  CreatedTransferOptionsDestinationOptionsType = "ach-credit-same-day"
	CreatedTransferOptionsDestinationOptionsTypeAchCreditStandard CreatedTransferOptionsDestinationOptionsType = "ach-credit-standard"
	CreatedTransferOptionsDestinationOptionsTypeAchDebitCollect   CreatedTransferOptionsDestinationOptionsType = "ach-debit-collect"
	CreatedTransferOptionsDestinationOptionsTypeAchDebitFund      CreatedTransferOptionsDestinationOptionsType = "ach-debit-fund"
	CreatedTransferOptionsDestinationOptionsTypeApplePay          CreatedTransferOptionsDestinationOptionsType = "apple-pay"
	CreatedTransferOptionsDestinationOptionsTypeCardPayment       CreatedTransferOptionsDestinationOptionsType = "card-payment"
	CreatedTransferOptionsDestinationOptionsTypeMoovWallet        CreatedTransferOptionsDestinationOptionsType = "moov-wallet"
	CreatedTransferOptionsDestinationOptionsTypeRtpCredit         CreatedTransferOptionsDestinationOptionsType = "rtp-credit"
)

type CreatedTransferOptionsDestinationOptions struct {
	PaymentMethodWallet      *PaymentMethodWallet
	PaymentMethodBankAccount *PaymentMethodBankAccount
	PaymentMethodCard        *PaymentMethodCard
	PaymentMethodApplePay    *PaymentMethodApplePay

	Type CreatedTransferOptionsDestinationOptionsType
}

func CreateCreatedTransferOptionsDestinationOptionsAchCreditSameDay(achCreditSameDay PaymentMethodBankAccount) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeAchCreditSameDay
	typStr := PaymentMethodsType(typ)
	achCreditSameDay.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodBankAccount: &achCreditSameDay,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsAchCreditStandard(achCreditStandard PaymentMethodBankAccount) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeAchCreditStandard
	typStr := PaymentMethodsType(typ)
	achCreditStandard.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodBankAccount: &achCreditStandard,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsAchDebitCollect(achDebitCollect PaymentMethodBankAccount) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeAchDebitCollect
	typStr := PaymentMethodsType(typ)
	achDebitCollect.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodBankAccount: &achDebitCollect,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsAchDebitFund(achDebitFund PaymentMethodBankAccount) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeAchDebitFund
	typStr := PaymentMethodsType(typ)
	achDebitFund.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodBankAccount: &achDebitFund,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsApplePay(applePay PaymentMethodApplePay) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeApplePay
	typStr := PaymentMethodsType(typ)
	applePay.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodApplePay: &applePay,
		Type:                  typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsCardPayment(cardPayment PaymentMethodCard) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeCardPayment
	typStr := PaymentMethodsType(typ)
	cardPayment.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodCard: &cardPayment,
		Type:              typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsMoovWallet(moovWallet PaymentMethodWallet) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeMoovWallet
	typStr := PaymentMethodsType(typ)
	moovWallet.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodWallet: &moovWallet,
		Type:                typ,
	}
}

func CreateCreatedTransferOptionsDestinationOptionsRtpCredit(rtpCredit PaymentMethodBankAccount) CreatedTransferOptionsDestinationOptions {
	typ := CreatedTransferOptionsDestinationOptionsTypeRtpCredit
	typStr := PaymentMethodsType(typ)
	rtpCredit.PaymentMethodType = &typStr

	return CreatedTransferOptionsDestinationOptions{
		PaymentMethodBankAccount: &rtpCredit,
		Type:                     typ,
	}
}

func (u *CreatedTransferOptionsDestinationOptions) UnmarshalJSON(data []byte) error {
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
		u.Type = CreatedTransferOptionsDestinationOptionsTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsDestinationOptionsTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsDestinationOptionsTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsDestinationOptionsTypeAchDebitFund
		return nil
	case "apple-pay":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodApplePay := new(PaymentMethodApplePay)
		if err := d.Decode(&paymentMethodApplePay); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodApplePay = paymentMethodApplePay
		u.Type = CreatedTransferOptionsDestinationOptionsTypeApplePay
		return nil
	case "card-payment":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodCard := new(PaymentMethodCard)
		if err := d.Decode(&paymentMethodCard); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCard = paymentMethodCard
		u.Type = CreatedTransferOptionsDestinationOptionsTypeCardPayment
		return nil
	case "moov-wallet":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodWallet := new(PaymentMethodWallet)
		if err := d.Decode(&paymentMethodWallet); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodWallet = paymentMethodWallet
		u.Type = CreatedTransferOptionsDestinationOptionsTypeMoovWallet
		return nil
	case "rtp-credit":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsDestinationOptionsTypeRtpCredit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreatedTransferOptionsDestinationOptions) MarshalJSON() ([]byte, error) {
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

type CreatedTransferOptionsSourceOptionsType string

const (
	CreatedTransferOptionsSourceOptionsTypeAchCreditSameDay  CreatedTransferOptionsSourceOptionsType = "ach-credit-same-day"
	CreatedTransferOptionsSourceOptionsTypeAchCreditStandard CreatedTransferOptionsSourceOptionsType = "ach-credit-standard"
	CreatedTransferOptionsSourceOptionsTypeAchDebitCollect   CreatedTransferOptionsSourceOptionsType = "ach-debit-collect"
	CreatedTransferOptionsSourceOptionsTypeAchDebitFund      CreatedTransferOptionsSourceOptionsType = "ach-debit-fund"
	CreatedTransferOptionsSourceOptionsTypeApplePay          CreatedTransferOptionsSourceOptionsType = "apple-pay"
	CreatedTransferOptionsSourceOptionsTypeCardPayment       CreatedTransferOptionsSourceOptionsType = "card-payment"
	CreatedTransferOptionsSourceOptionsTypeMoovWallet        CreatedTransferOptionsSourceOptionsType = "moov-wallet"
	CreatedTransferOptionsSourceOptionsTypeRtpCredit         CreatedTransferOptionsSourceOptionsType = "rtp-credit"
)

type CreatedTransferOptionsSourceOptions struct {
	PaymentMethodWallet      *PaymentMethodWallet
	PaymentMethodBankAccount *PaymentMethodBankAccount
	PaymentMethodCard        *PaymentMethodCard
	PaymentMethodApplePay    *PaymentMethodApplePay

	Type CreatedTransferOptionsSourceOptionsType
}

func CreateCreatedTransferOptionsSourceOptionsAchCreditSameDay(achCreditSameDay PaymentMethodBankAccount) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeAchCreditSameDay
	typStr := PaymentMethodsType(typ)
	achCreditSameDay.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodBankAccount: &achCreditSameDay,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsAchCreditStandard(achCreditStandard PaymentMethodBankAccount) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeAchCreditStandard
	typStr := PaymentMethodsType(typ)
	achCreditStandard.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodBankAccount: &achCreditStandard,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsAchDebitCollect(achDebitCollect PaymentMethodBankAccount) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeAchDebitCollect
	typStr := PaymentMethodsType(typ)
	achDebitCollect.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodBankAccount: &achDebitCollect,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsAchDebitFund(achDebitFund PaymentMethodBankAccount) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeAchDebitFund
	typStr := PaymentMethodsType(typ)
	achDebitFund.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodBankAccount: &achDebitFund,
		Type:                     typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsApplePay(applePay PaymentMethodApplePay) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeApplePay
	typStr := PaymentMethodsType(typ)
	applePay.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodApplePay: &applePay,
		Type:                  typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsCardPayment(cardPayment PaymentMethodCard) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeCardPayment
	typStr := PaymentMethodsType(typ)
	cardPayment.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodCard: &cardPayment,
		Type:              typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsMoovWallet(moovWallet PaymentMethodWallet) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeMoovWallet
	typStr := PaymentMethodsType(typ)
	moovWallet.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodWallet: &moovWallet,
		Type:                typ,
	}
}

func CreateCreatedTransferOptionsSourceOptionsRtpCredit(rtpCredit PaymentMethodBankAccount) CreatedTransferOptionsSourceOptions {
	typ := CreatedTransferOptionsSourceOptionsTypeRtpCredit
	typStr := PaymentMethodsType(typ)
	rtpCredit.PaymentMethodType = &typStr

	return CreatedTransferOptionsSourceOptions{
		PaymentMethodBankAccount: &rtpCredit,
		Type:                     typ,
	}
}

func (u *CreatedTransferOptionsSourceOptions) UnmarshalJSON(data []byte) error {
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
		u.Type = CreatedTransferOptionsSourceOptionsTypeAchCreditSameDay
		return nil
	case "ach-credit-standard":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsSourceOptionsTypeAchCreditStandard
		return nil
	case "ach-debit-collect":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsSourceOptionsTypeAchDebitCollect
		return nil
	case "ach-debit-fund":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsSourceOptionsTypeAchDebitFund
		return nil
	case "apple-pay":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodApplePay := new(PaymentMethodApplePay)
		if err := d.Decode(&paymentMethodApplePay); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodApplePay = paymentMethodApplePay
		u.Type = CreatedTransferOptionsSourceOptionsTypeApplePay
		return nil
	case "card-payment":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodCard := new(PaymentMethodCard)
		if err := d.Decode(&paymentMethodCard); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCard = paymentMethodCard
		u.Type = CreatedTransferOptionsSourceOptionsTypeCardPayment
		return nil
	case "moov-wallet":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodWallet := new(PaymentMethodWallet)
		if err := d.Decode(&paymentMethodWallet); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodWallet = paymentMethodWallet
		u.Type = CreatedTransferOptionsSourceOptionsTypeMoovWallet
		return nil
	case "rtp-credit":
		d = json.NewDecoder(bytes.NewReader(data))
		paymentMethodBankAccount := new(PaymentMethodBankAccount)
		if err := d.Decode(&paymentMethodBankAccount); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodBankAccount = paymentMethodBankAccount
		u.Type = CreatedTransferOptionsSourceOptionsTypeRtpCredit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreatedTransferOptionsSourceOptions) MarshalJSON() ([]byte, error) {
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

type CreatedTransferOptions struct {
	DestinationOptions []CreatedTransferOptionsDestinationOptions `json:"destinationOptions,omitempty"`
	SourceOptions      []CreatedTransferOptionsSourceOptions      `json:"sourceOptions,omitempty"`
}

func (o *CreatedTransferOptions) GetDestinationOptions() []CreatedTransferOptionsDestinationOptions {
	if o == nil {
		return nil
	}
	return o.DestinationOptions
}

func (o *CreatedTransferOptions) GetSourceOptions() []CreatedTransferOptionsSourceOptions {
	if o == nil {
		return nil
	}
	return o.SourceOptions
}
