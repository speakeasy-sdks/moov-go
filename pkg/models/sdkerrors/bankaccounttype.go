// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// BankAccountType - The bank account type
type BankAccountType string

const (
	BankAccountTypeChecking BankAccountType = "checking"
	BankAccountTypeSavings  BankAccountType = "savings"
	BankAccountTypeUnknown  BankAccountType = "unknown"
)

func (e BankAccountType) ToPointer() *BankAccountType {
	return &e
}

func (e *BankAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "checking":
		fallthrough
	case "savings":
		fallthrough
	case "unknown":
		*e = BankAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BankAccountType: %v", v)
	}
}