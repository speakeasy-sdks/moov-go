// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

type Type string

const (
	TypeDriversLicense Type = "DriversLicense"
	TypePassport       Type = "Passport"
	TypeUtilityBill    Type = "UtilityBill"
	TypeBankStatement  Type = "BankStatement"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DriversLicense":
		fallthrough
	case "Passport":
		fallthrough
	case "UtilityBill":
		fallthrough
	case "BankStatement":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// Document - Describes an uploaded file
type Document struct {
	ContentType string `json:"contentType"`
	// A unique identifier for this document
	DocumentID string `json:"documentID"`
	// Optional array of errors encountered dring automated parsing.
	ParseErrors []string  `json:"parseErrors,omitempty"`
	Type        Type      `json:"type"`
	UploadedAt  time.Time `json:"uploadedAt"`
}

func (d Document) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Document) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Document) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Document) GetDocumentID() string {
	if o == nil {
		return ""
	}
	return o.DocumentID
}

func (o *Document) GetParseErrors() []string {
	if o == nil {
		return nil
	}
	return o.ParseErrors
}

func (o *Document) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *Document) GetUploadedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UploadedAt
}
