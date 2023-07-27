// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BusinessType - The type of entity represented by this Business
type BusinessType string

const (
	BusinessTypeSoleProprietorship        BusinessType = "soleProprietorship"
	BusinessTypeUnincorporatedAssociation BusinessType = "unincorporatedAssociation"
	BusinessTypeTrust                     BusinessType = "trust"
	BusinessTypePublicCorporation         BusinessType = "publicCorporation"
	BusinessTypePrivateCorporation        BusinessType = "privateCorporation"
	BusinessTypeLlc                       BusinessType = "llc"
	BusinessTypePartnership               BusinessType = "partnership"
	BusinessTypeUnincorporatedNonProfit   BusinessType = "unincorporatedNonProfit"
	BusinessTypeIncorporatedNonProfit     BusinessType = "incorporatedNonProfit"
)

func (e BusinessType) ToPointer() *BusinessType {
	return &e
}

func (e *BusinessType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "soleProprietorship":
		fallthrough
	case "unincorporatedAssociation":
		fallthrough
	case "trust":
		fallthrough
	case "publicCorporation":
		fallthrough
	case "privateCorporation":
		fallthrough
	case "llc":
		fallthrough
	case "partnership":
		fallthrough
	case "unincorporatedNonProfit":
		fallthrough
	case "incorporatedNonProfit":
		*e = BusinessType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BusinessType: %v", v)
	}
}