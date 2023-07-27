// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FilePurpose - The file purpose
type FilePurpose string

const (
	FilePurposeIdentityVerification FilePurpose = "identity_verification"
	FilePurposeBusinessVerification FilePurpose = "business_verification"
	FilePurposeMerchantUnderwriting FilePurpose = "merchant_underwriting"
	FilePurposeAccountRequirement   FilePurpose = "account_requirement"
)

func (e FilePurpose) ToPointer() *FilePurpose {
	return &e
}

func (e *FilePurpose) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "identity_verification":
		fallthrough
	case "business_verification":
		fallthrough
	case "merchant_underwriting":
		fallthrough
	case "account_requirement":
		*e = FilePurpose(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FilePurpose: %v", v)
	}
}